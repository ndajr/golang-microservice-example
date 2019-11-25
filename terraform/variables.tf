variable "profile" {
  type        = string
  description = "AWS profile in use (e.g. 374047294805)"
}

variable "region" {
  type        = string
  description = "AWS region such as (eu-central-1, us-west-1)"
}

variable "env" {
  type        = string
  description = "The environment name (dev, test, stage, prod)"
}

variable "docker_image_tag" {
  type = string
}
