# Custom domain name in API Gateway
resource "aws_api_gateway_domain_name" "project_api_domain_name" {
  certificate_arn = data.aws_ssm_parameter.certificate_arn_regional.value
  domain_name     = var.api_domain_name
}

# Entry in Route 53
resource "aws_route53_record" "api_domain_routing" {
  name    = aws_api_gateway_domain_name.project_api_domain_name.domain_name
  type    = "A"
  zone_id = data.aws_route53_zone.jponc_dev.zone_id

  alias {
    evaluate_target_health = false
    name                   = aws_api_gateway_domain_name.project_api_domain_name.cloudfront_domain_name
    zone_id                = aws_api_gateway_domain_name.project_api_domain_name.cloudfront_zone_id
  }
}

# Connecting the custom domain name to api gateway via custom domain name mapping
resource "aws_api_gateway_base_path_mapping" "project_api_domain_name_mapping" {
  api_id      = data.aws_api_gateway_rest_api.project_api.id
  stage_name  = var.environment
  domain_name = aws_api_gateway_domain_name.project_api_domain_name.domain_name
}
