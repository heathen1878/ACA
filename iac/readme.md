# IAC

[Home](../README.md)

## Prerequsities

### Github

[GitHub Cli](https://cli.github.com/)

### Github Variables

Create github variables to set the Terraform backend

```shell
# Example syntax
gh variable set NAME -b "secret value" -r "repo name"

gh variable set TERRAFORM_VERSION -b "core executable version e.g. 1.9.3" -r account_name/repo_name
gh variable set STATE_STORAGE_ACCOUNT -b storage_account_name -r "account_name/repo_name"
gh variable set STATE_STORAGE_CONTAINER -b state_storage_container -r "account_name/repo_name"
gh variable set STATE_FILE_NAME -b state_file_name -r "account_name/repo_name"
gh variable set ARM_SUBSCRIPTION_ID -b sub_guid -r "account_name/repo_name"
gh variable set APP_ID -b token_app_id -r "account_name/repo_name"
```

### Github Secrets

Create github secrets to authenticate with Azure

```shell
gh secret set ARM_CLIENT_ID -b "<Client ID>" -r "account_name/repo_name"
gh secret set ARM_CLIENT_SECRET -b "<Client Secret>" -r "account_name/repo_name"
gh secret set ARM_TENANT_ID -b "<Tenant ID>" -r "account_name/repo_name"
gh secret set CREATE_TOKEN_PRIVATE_KEY -b CREATE TOKEN APP PRIVATE KEY -r "account_name/repo_name"
```

#### Infracost

See here for instructions but in summary you need an API key secret and a infracost workflow template.

```shell
gh secret set INFRACOST_API_KEY -b "Infracost API key" -r "account_name/repo_name"
```

Workflow template [here](../.github/workflows/)

#### Snyk

Snyk is a vulnerability scanner for IaC amongst other things.

```shell
gh secret set SNYK_TOKEN -b "Snyk API key" -r "account_name/repo_name"
```

## Validation

The workflow [validate infra](../.github/workflows/validate_fibonacci_v1_infra.yaml) does the following:

- Authenticates the pipeline
- Gets the runners IP address
- Set an ACL on the state storage account
- Check Terraform formatting
- Validates the Terraform syntax
- Plans against any `.tf` within the repo
- Displays a cost for the proposed Terraform changes - See Infracost
- Deploys the infrastructure and associated Docker Images
- Tests the Web App responds via the Ingress Controller
- Destroys the Infrastructure
