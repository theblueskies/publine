data "archive_file" "runner_archive" {
  source_file = "../lambdas/runner/main"
  output_path = "../lambdas/runner/main.zip"
  type        = "zip"
}

resource "aws_lambda_function" "runner" {
  function_name = "ddb_events_consumer"
  role          = aws_iam_role.lambda_consumer_role.arn
  handler       = "main"
  memory_size   = "128"
  timeout       = 10
  runtime       = "go1.x"
  architectures = ["x86_64"]

  filename         = data.archive_file.runner_archive.output_path
  source_code_hash = data.archive_file.runner_archive.output_path

  depends_on = [ aws_iam_role.lambda_consumer_role ]
}

resource "aws_lambda_event_source_mapping" "runner_event_source" {
  event_source_arn  = aws_dynamodb_table.pubcoredb.stream_arn
  function_name     = aws_lambda_function.runner.arn
  starting_position = "LATEST"
}