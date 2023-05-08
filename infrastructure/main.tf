terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
    }
  }

  required_version = ">= 1.2.0"

  backend "s3" {
    bucket = "publinestate" #Note: This bucket needs to preexist
    key    = "terraform.tfstate"
    region = "us-east-1"
  }
}

provider "aws" {
  region      = "us-east-1"
  max_retries = 3
}
