resource "aws_s3_bucket" "backend-s3" {
  bucket = "golang-microservice-example-bucket"
  tags   = local.tags
}
