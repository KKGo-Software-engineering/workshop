terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "4.29.0"
    }
  }
}

variable "group_name" {}

provider "aws" {
  region = "ap-southeast-1"
}