resource "aws_ssm_parameter" "dynamodb_table" {
  name  = "/${var.project_name}/${var.environment}/DYNAMODB_TABLE_NAME"
  type  = "String"
  value = aws_dynamodb_table.project_table.id
}
