# These are the different configurable variables we currently support
variable "environment" {
  type        = string
  description = "Environment"
}

variable "project_name" {
  type        = string
  description = "This is the name of the project"
}

variable "aws_profile" {
  type        = string
  description = "AWS Profile"
}

variable "aws_region" {
  type        = string
  description = "AWS Region"
}

variable "api_domain_name" {
  type        = string
  description = "API Domain name"
}

variable "frontend_domain_name" {
  type        = string
  description = "Frontend Domain name"
}
