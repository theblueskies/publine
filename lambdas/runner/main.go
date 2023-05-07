package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, event events.DynamoDBEvent) (string, error) {
	return fmt.Sprintf("Hello world!"), nil
}

func main() {
	lambda.Start(HandleRequest)
}
