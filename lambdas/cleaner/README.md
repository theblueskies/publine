# cleaner

cleaner is a lambda which is triggered by Cloudwatch Alarm once every minute. Dynamo streams receive items on it only after they have been deleted from the table. This lambda scans DynamoDB for any items that should have been deleted on expiry and deletes them in batches. It ensures that any items that should have been deleted ARE actually deleted.

Items from DynamoDB are placed on the DynamoDB stream only after they are deleted. The DynamoDB TTL expiry is done by a background task that may not guarantee deletions at the exact time. This lambda is a backup/insurance against any delays.

## Build

1. `make build`  - This will get a new executable called `main` in the current folder

More info on the DynamoDB event here -

- [AWS SDK Golang](https://docs.aws.amazon.com/sdk-for-go/api/)
- [Golang DynamoDB struct](https://github.com/aws/aws-lambda-go/blob/main/events/README_DynamoDB.md)

Note: This lambda uses v1 of the AWS SDK for Go. An implementation with AWS SDK Go v2 can be made by implementing the DBClean interface.