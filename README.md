# publine

## Build lambda (runner)

1. `make runner`

## Deploy infrastructure  

Prerequisite: Setup your AWS access key and secret in `~/.aws/credentials`

1. export AWS_PROFILE=name_of_your_profile_in_aws_credentials
2. `terraform init`  
3. `terraform apply`  

Note: A deploy of a lambda requires both the major steps listed above - building the lambda and then redoing `terraform apply`
