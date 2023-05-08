resource "aws_cloudwatch_event_rule" "every_one_minute" {
  name                = "every-one-minute"
  description         = "Fires every one minute"
  schedule_expression = "rate(1 minute)" //Plural: 5 minutes
  // Rate expressions: https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html#RateExpressions
}

resource "aws_cloudwatch_event_target" "check_every_minute" {
  arn       = aws_lambda_function.cleaner.arn
  rule      = aws_cloudwatch_event_rule.every_one_minute.name
  target_id = "clean_ddb"
}

resource "aws_lambda_permission" "allow_cloudwatch_to_call_check_foo" {
  statement_id  = "AllowExecutionFromCloudWatch"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.cleaner.function_name
  principal     = "events.amazonaws.com"
  source_arn    = aws_cloudwatch_event_rule.every_one_minute.arn
}
