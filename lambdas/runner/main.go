package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.uber.org/zap"
)

type MyEvent struct {
	Name string `json:"name"`
}

// HandleRequest implements the core behaviors of the lambda
func HandleRequest(ctx context.Context, ddbEvents events.DynamoDBEvent) (string, error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	for _, e := range ddbEvents.Records {
		sugar.Info("Got DDB record", e)
		// The current filter pattern being applied only sends events to Lambda when eventName = REMOVE
		// Filter pattern is defined in infrastucture/lambdas.tf
		// e.Change.NewImage is nil because values get erased when item is deleted from DynamoDB
		userId := e.Change.OldImage["userId"].String()
		todoTitle := e.Change.OldImage["todoTitle"].String()

		sugar.Info("userID: ", userId)
		sugar.Info("todoTitle: ", todoTitle)

		// INSERT YOUR BUSINESS LOGIC HERE
	}

	return fmt.Sprintf("run complete!"), nil
}

func main() {
	lambda.Start(HandleRequest)
}
