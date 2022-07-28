data "aws_ssm_parameter" "certificate_arn_regional" {
  name = "/jponc.dev/CERTIFICATE_ARN"
}

data "aws_api_gateway_rest_api" "project_api" {
  name = "${var.project_name}-${var.environment}"
}

data "aws_route53_zone" "jponc_dev" {
  name = "jponc.dev"
}
