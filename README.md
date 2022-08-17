# AWS Credentials Cloner

This small tool was written to copy credentials from a `credential_process` in `~/.aws/config` to `~/.aws/credentials`. 
This makes dealing with tools that only read `~/.aws/credentials` (such as `serverless`) simpler.

# Installation

## Makefile
```shell
make install
```
## Manually
run these commands inside the project root
```shell
go build .
```
```shell
sudo cp aws-credentials-cloner /usr/local/bin/.
```

# Usage
`aws-credentials-cloner` will take the currently selected profile from the `AWS_PROFILE` env var and run 
the credential process configured for that profile in `~/.aws/config`. 

It is written specifically for credential 
processes that output as json and do so without the user supplying input (e.g. 
`aws-adfs login --authfile=<some_creds_file> --adfs-host <some_host> --role-arn <some_role> --region <some_region> --stdout`).
Credential processes that require user input will not work.