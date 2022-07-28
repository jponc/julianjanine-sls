// This is the general terraform configuration including AWS profile used, dynamodb table, and AWS region
terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "4.17.1"
    }
  }

  backend "s3" {
    bucket         = "julianjanine-terraform-state-ap-southeast-1"
    key            = "global/s3/terraform.tfstate"
    region         = "ap-southeast-1"
    profile        = "jponc"
    dynamodb_table = "julianjanine-terraform-locks-ap-southeast-1"
    encrypt        = true
  }
}

provider "aws" {
  profile = var.aws_profile
  region  = var.aws_region
}
