# az-devops-assigner
## Description
The tool assigns user/group privileges on all existing ( or, at least, accessible by a passed PAT ) Variable Groups matched by a pattern in Azure DevOps Organization.
## Requirements
### Environment Variables
The following environment variables must be set:
* AZ_DEVOPS_PAT: Azure DevOps Personal Access Token
* AZ_DEVOPS_ORG: Azure DevOps Organization name
### CLI Arguments
The following CLI arguments must be passed:
  * --name: User email or a group name you want give permissions to (test@example.com, TestGroup, etc.)
  * --pattern string: Case-insensitive part of a variable group name to filter by (VarGroup, var, vaRG matches 'var' pattern))
  * --role: Role you want to assign to the user or group (Administrator, Reader, etc.)
## Usage
```shell
USAGE:
  az-devops-assigner [flags]

FLAGS:
  -name string (REQUIRED)       User email or a group name you want give permissions to (test@example.com, TestGroup, etc.)
  -pattern string (REQUIRED)    Case-insensitive part of a variable group name to filter by (VarGroup, var, vaRG matches 'var' pattern))
  -role string (REQUIRED)       Role you want to assign to the user or group (Administrator, Reader, etc.)
```
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
