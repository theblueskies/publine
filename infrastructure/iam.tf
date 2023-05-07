resource "aws_iam_role" "lambda_consumer_role" {
  name = "lambda_consumer_role"

  # Terraform's "jsonencode" function converts a
  # Terraform expression result to valid JSON syntax.
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = ""
        Principal = {
          Service = "ec2.amazonaws.com"
        }
      },
    ]
  })

  tags = {
    Name        = "lambda_consumer_role"
    Description = "lambda role to consume data from dynamodb streams"
  }
}

resource "aws_iam_role_policy_attachment" "lambda_execution_policy" {
    role = aws_iam_role.lambda_consumer_role.name
    policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole" 
}

resource "aws_iam_role_policy_attachment" "ddbstream_execution_policy" {
    role = aws_iam_role.lambda_consumer_role.name
    policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaDynamoDBExecutionRole" 
}
