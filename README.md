# az-devops-assigner
## Description
The tool assigns user/group privileges on all existing ( or, at least, accessible by a passed PAT ) Variable Groups matched by a pattern in Azure DevOps Organization.
## Requirements
The following environment variables must be set:
* AZ_DEVOPS_PAT: Azure DevOps Personal Access Token
* AZ_DEVOPS_ORG: Azure DevOps Organization name

## Usage
### Developing locally
```shell
❯ go run ./cmd/azassigner  --name test@example.com --pattern var --role Reader
2023/07/20 15:04:40 Running with username: test@example.com, desired role: Reader, variable group pattern: var
....
```
### Building and running as a Docker container
> :warning: AZ_DEVOPS_PAT, AZ_DEVOPS_ORG variables must be set.
```shell
❯ docker build . -t azassigner
❯ docker run --rm -it \
-e AZ_DEVOPS_PAT \
-e AZ_DEVOPS_ORG \
azassigner \
--name test@example.com --pattern var --role Reader
```
### Building and running as a Docker container with Docker Compose
> :warning: AZ_DEVOPS_PAT, AZ_DEVOPS_ORG variables must be set.
```shell
❯ docker compose up
```
