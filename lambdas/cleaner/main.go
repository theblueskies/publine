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

	// INSERT YOUR BUSINESS LOGIC HERE

	sugar.Info("run complete")
	return fmt.Sprintf("run complete!"), nil
}

func main() {
	lambda.Start(HandleRequest)
}
