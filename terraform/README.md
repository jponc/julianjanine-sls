# Infrastructure

- This uses Terraform with `hashicorp/aws` version `4.17.1`

# Starting

Install dependencies

```
brew tap hashicorp/tap
brew install hashicorp/tap/terraform
```

Initialise remote

```
cd ap-southeast-1/remote
```

Initialise infrastructure

```
cd ap-southeast-1/infrastructure

# initialise terraform
terraform init

# initialise workspaces
make init-workspaces

# init all workspaces
make init-dev
make init-prod
```
