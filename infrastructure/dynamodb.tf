resource "aws_dynamodb_table" "pubcoredb" {
  name             = "pubcore_db"
  billing_mode     = "PAY_PER_REQUEST"
  stream_enabled   = true
  stream_view_type = "NEW_AND_OLD_IMAGES"
  hash_key         = "userId"
  range_key        = "todoTitle"

  attribute {
    name = "userId"
    type = "S"
  }

  attribute {
    name = "todoTitle"
    type = "S"
  }

  ttl {
    attribute_name = "expiryTime"
    enabled        = true
  }

  tags = {
    Name        = "pubcore_db"
    Environment = "production"
  }
}