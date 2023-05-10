# runner

This lambda executes any side effects resulting from the TTL expiry event of an item coming in through the DynamoDB Stream.

This is a skeleton and only logs that it received event(s). Any business logic required is left for the user to make.

## Build

1. `make build`  - This will get a new executable called `main` in the current folder

More info on the DynamoDB event here -

- [DynamoDB JSON Payload](https://github.com/aws/aws-lambda-go/blob/main/events/testdata/dynamodb-event.json)
- [Golang DynamoDB struct](https://github.com/aws/aws-lambda-go/blob/main/events/README_DynamoDB.md)
