# publine

Publine is a proof-of-concept for -

1. DynamoDB events getting sent over the DynamoDB stream on TTL expiry
2. an AWS Lambda ingesting events from a DynamoDB stream

There are two major steps to deploy - building the lambda and deploying the lambda alongwith DynamoDB and related infrastructure

## Build lambda (runner)

This is the lambda (named runner) that ingests events from DynamoDB streams. There is a Makefile in the folder of runner that has the instruction to build the lambda

1. `make runner`

## Deploy infrastructure  

Prerequisite: Setup your AWS access key and secret in `~/.aws/credentials`

1. `export AWS_PROFILE=name_of_your_profile_in_aws_credentials`
2. `terraform init`  
3. `terraform apply`  

Note: A deploy of a lambda requires both the major steps listed above - building the lambda and then redoing `terraform apply`

### Branch descriptions

These branches are in increasing order of development. After the first branch, the others are pretty incremental in change

- 1.0: Basic infrastructure setup and lambda deploy
- 1.1: Update lambda function structure
- 1.2: Attribute mapping from DynamoDB streamed event into Lambda
- 1.3: Filter criteria on events before they are sent to the Lambda
- 2.0: Lambda triggered as a cronjob by Cloudwatch Alarm
