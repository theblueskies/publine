package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/theblueskies/publine/lambdas/cleaner/api"

	"go.uber.org/zap"
)

// HandleRequest is the entrypoint to the lambda
func HandleRequest(ctx context.Context, cwEvent events.CloudWatchEvent) (string, error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	// INSERT YOUR BUSINESS LOGIC HERE
	apix, err := api.NewUserRepository("pubcore_db")
	if err != nil {
		sugar.Error("error creating cleaner service", err)
	}

	err = apix.BatchDeleteExpiredItems(ctx, cwEvent)
	if err != nil {
		sugar.Error("error deleting expired items", err)
	}

	sugar.Info("run complete")
	return fmt.Sprintf("run complete!"), nil
}

func main() {
	lambda.Start(HandleRequest)
}
