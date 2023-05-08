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

func HandleRequest(ctx context.Context, ddbEvents events.DynamoDBEvent) (string, error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	for _, e := range ddbEvents.Records {
		sugar.Info("Got DDB record", e)
	}

	return fmt.Sprintf("run complete!"), nil
}

func main() {
	lambda.Start(HandleRequest)
}
