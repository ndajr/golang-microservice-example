terraform {
  required_version = "~> 0.12.0"

  backend "s3" {}
}

provider "aws" {
  region  = "eu-central-1"
  profile = var.profile
}
