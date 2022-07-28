# Create workspaces per env
init-workspaces:
	terraform workspace new prod

# Initialise per env

init-prod:
	terraform workspace select prod
	terraform init

# Plan per env
plan-prod:
	terraform workspace select prod
	terraform plan -var-file="prod.tfvars"

# Apply per env
apply-prod:
	terraform workspace select prod
	terraform apply -var-file="prod.tfvars"
