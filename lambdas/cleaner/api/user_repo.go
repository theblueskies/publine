package api

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"go.uber.org/zap"
)

// This is a max upper bound limit that can be sent to DynamoDB
// https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#DynamoDB.BatchWriteItem
const DYNAMO_MAX_BATCH_CALL_LIMIT = 25

type UserEntry struct {
	UserId    string `json:"userId"`
	TodoTitle string `json:"todoTitle"`
}

// UserRepository implements the DBClean interface
type UserRepository struct {
	svc       *dynamodb.DynamoDB
	tableName string
	logger    *zap.SugaredLogger
	callLimit int
}

func NewUserRepository(tableName string) (*UserRepository, error) {
	svc := dynamodb.New(session.New())
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	return &UserRepository{svc: svc, tableName: tableName, logger: sugar}, nil
}

// BatchDeleteExpiredItems takes all items that should have beel deleted and enforces that delete
func (u *UserRepository) BatchDeleteExpiredItems(ctx context.Context, cwEvent events.CloudWatchEvent) error {
	expiryTime := cwEvent.Time.Unix()
	filter := expression.Name("expiryTime").LessThan(expression.Value(expiryTime))
	projection := expression.NamesList(expression.Name("userId"), expression.Name("todoTitle"), expression.Name("expiryTime"))

	expr, err := expression.NewBuilder().WithFilter(filter).WithProjection(projection).Build()
	if err != nil {
		u.logger.Error("error building expression: %s", err)
	}

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(u.tableName),
	}

	result, err := u.svc.Scan(params)
	if err != nil {
		u.logger.Error("error scanning table", err)
		return err
	}

	var wr []*dynamodb.WriteRequest
	batchCount := 1
	writtenCount := 0
	for index, v := range result.Items {
		userEntry := UserEntry{}
		err := dynamodbattribute.UnmarshalMap(v, &userEntry)
		if err != nil {
			u.logger.Error("failed to unmarshal Dynamodb Scan Items", err)
		}

		wr = append(wr, &dynamodb.WriteRequest{
			DeleteRequest: &dynamodb.DeleteRequest{
				Key: map[string]*dynamodb.AttributeValue{
					"userId": {
						S: aws.String(userEntry.UserId),
					},
					"todoTitle": {
						S: aws.String(userEntry.TodoTitle),
					},
				},
			}})

		// Write batch request if it hits max batch limit or is the last element in the scan
		if writtenCount == u.callLimit || index == int(*result.Count)-1 {
			u.logger.Info("deleting batch ", map[string]int{"batch": batchCount, "batchSize": writtenCount})

			input := &dynamodb.BatchWriteItemInput{
				RequestItems: map[string][]*dynamodb.WriteRequest{
					u.tableName: wr,
				},
			}

			// potential for offloading this call to a goroutine
			err = u.performDelete(input)
			if err != nil {
				return err
			}

			// reset the counts for the next batch
			wr = nil
			writtenCount = 0
		}
	}
	return nil
}

func (u *UserRepository) performDelete(input *dynamodb.BatchWriteItemInput) error {
	_, err := u.svc.BatchWriteItem(input)
	if err != nil {
		u.logger.Error("error batch writing the delete operation", err)
		return err
	}
	return nil
}
