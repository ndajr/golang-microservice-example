locals {
  service_name = "golang-microservice-example"

  tags = {
    Name        = "${var.profile}-${local.service_name}"
    Environment = var.env
    Application = local.service_name
    Service     = local.service_name
    GitRepo     = "github.com/neemiasjnr/${local.service_name}"
    Owner       = "neemias.junior@gmail.com"
    ManagedBy   = "terraform"
  }
}
