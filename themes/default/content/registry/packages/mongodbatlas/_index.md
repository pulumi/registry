---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-mongodbatlas/v3.37.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi/pulumi-mongodbatlas/blob/v3.37.0/docs/_index.md
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Mongodbatlas Provider
meta_desc: Provides an overview on how to configure the Pulumi Mongodbatlas provider.
layout: package
---

## Installation

The Mongodbatlas provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/mongodbatlas`](https://www.npmjs.com/package/@pulumi/mongodbatlas)
* Python: [`pulumi-mongodbatlas`](https://pypi.org/project/pulumi-mongodbatlas/)
* Go: [`github.com/pulumi/pulumi-mongodbatlas/sdk/v3/go/mongodbatlas`](https://github.com/pulumi/pulumi-mongodbatlas)
* .NET: [`Pulumi.Mongodbatlas`](https://www.nuget.org/packages/Pulumi.Mongodbatlas)
* Java: [`com.pulumi/mongodbatlas`](https://central.sonatype.com/artifact/com.pulumi/mongodbatlas)

## Overview

You can use the MongoDB Atlas provider to interact with the resources supported by [MongoDB Atlas](https://www.mongodb.com/cloud/atlas).
The provider needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available provider resources and functions.

You may want to consider pinning the provider version to ensure you have a chance to review and prepare for changes.
## Example Usage

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    mongodbatlas:privateKey:
        value: 'TODO: var.mongodbatlas_private_key'
    mongodbatlas:publicKey:
        value: 'TODO: var.mongodbatlas_public_key'

```
## Configure Atlas Programmatic Access

In order to set up authentication with the MongoDB Atlas provider, you must generate a programmatic API key for MongoDB Atlas with the appropriate [role](https://docs.atlas.mongodb.com/reference/user-roles/).
The [MongoDB Atlas documentation](https://docs.atlas.mongodb.com/tutorial/manage-programmatic-access/index.html) contains the most up-to-date instructions for creating and managing your key(s), setting the appropriate role, and optionally configuring IP access.

**Role**: If unsure of which role level to grant your key, we suggest creating an organization API Key with an Organization Owner role. This ensures that you have sufficient access for all actions.
## Configure MongoDB Atlas for Government

In order to enable the Pulumi MongoDB Atlas Provider for use with MongoDB Atlas for Government add isMongodbgovCloud = true to your provider configuration:
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    mongodbatlas:isMongodbgovCloud:
        value: true
    mongodbatlas:privateKey:
        value: 'TODO: var.mongodbatlas_private_key'
    mongodbatlas:publicKey:
        value: 'TODO: var.mongodbatlas_public_key'

```

Also see [`Atlas for Government Considerations`](https://www.mongodb.com/docs/atlas/government/api/#atlas-for-government-considerations).
## Authenticate the Provider

The MongoDB Atlas provider offers a flexible means of providing credentials for authentication.
You can use any the following methods:
### Environment Variables

You can also provide your credentials via the environment variables,
`MONGODB_ATLAS_PUBLIC_API_KEY` and `MONGODB_ATLAS_PRIVATE_API_KEY`,
for your public and private MongoDB Atlas programmatic API key pair respectively:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:

```

Usage (prefix the export commands with a space to avoid the keys being recorded in OS history):

```shell
$  export MONGODB_ATLAS_PUBLIC_API_KEY="<ATLAS_PUBLIC_KEY>"
$  export MONGODB_ATLAS_PRIVATE_API_KEY="<ATLAS_PRIVATE_KEY>"
$ pulumi preview
```

We recommend that you use the `MONGODB_ATLAS_PUBLIC_API_KEY` and `MONGODB_ATLAS_PRIVATE_API_KEY` environment variables because they are compatible with other MongoDB tools, such as Atlas CLI.
You can still use `MONGODB_ATLAS_PUBLIC_KEY` and `MONGODB_ATLAS_PRIVATE_KEY` as alternative keys in your local environment. However, these environment variables are not guaranteed to work across all tools in the MongoDB ecosystem.
### AWS Secrets Manager
AWS Secrets Manager (AWS SM) helps to manage, retrieve, and rotate database credentials, API keys, and other secrets throughout their lifecycles. See [product page](https://aws.amazon.com/secrets-manager/) and [documentation](https://docs.aws.amazon.com/systems-manager/latest/userguide/what-is-systems-manager.html) for more details.

In order to enable the Pulumi MongoDB Atlas Provider with AWS SM, please follow the below steps:

1. Create Atlas API Keys and add them as one secret to AWS SM with a raw value. Take note of which AWS Region secret is being stored in. Public Key and Private Key each need to be entered as their own key value pair. See below example:
```
     {
      "public_key": "secret1",
      "private_key":"secret2"
     }
```
2. Create an AWS IAM Role to attach to the AWS STS (Security Token Service) generated short lived API keys. This is required since STS generated API Keys by default have restricted permissions and need to have their permissions elevated in order to authenticate with Pulumi. Take note of Role ARN and ensure IAM Role has permission for “sts:AssumeRole”. For example:
```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "Statement1",
            "Effect": "Allow",
            "Principal": {
                "AWS": "*"
            },
            "Action": "sts:AssumeRole"
        }
    ]
}
```
In addition, you are required to also attach the AWS Managed policy of `SecretsManagerReadWrite` to this IAM role.

Note: this policy may be overly broad for many use cases, feel free to adjust accordingly to your organization's needs.

3. In terminal, store as environmental variables AWS API Keys (while you can also hardcode in config files these will then be stored as plain text in .tfstate file and should be avoided if possible). For example:
```
export AWS_ACCESS_KEY_ID='<AWS_ACCESS_KEY_ID>'
export AWS_SECRET_ACCESS_KEY='<AWS_SECRET_ACCESS_KEY>'
```
4. In terminal, use the AWS CLI command: `aws sts assume-role --role-arn ROLE_ARN_FROM_ABOVE --role-session-name newSession`

Note: AWS STS secrets are short lived by default, use the ` --duration-seconds` flag to specify longer duration as needed

5. Store each of the 3 new created secrets from AWS STS as environment variables (hardcoding secrets into config file with additional risk is also supported). For example:
```
export AWS_ACCESS_KEY_ID='<AWS_ACCESS_KEY_ID>'
export AWS_SECRET_ACCESS_KEY='<AWS_SECRET_ACCESS_KEY>'
export AWS_SESSION_TOKEN="<AWS_SESSION_TOKEN>"
```

6. Add assumeRole block with `roleArn`, `secretName`, and AWS `region` where secret is stored as part of AWS SM. Each of these 3 fields are REQUIRED. For example:
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    mongodbatlas:awsAccessKeyId:
        value: <AWS_ACCESS_KEY_ID>
    mongodbatlas:awsSecretAccessKey:
        value: <AWS_SECRET_ACCESS_KEY>
    mongodbatlas:awsSessionToken:
        value: <AWS_SESSION_TOKEN>
    mongodbatlas:region:
        value: us-east-2
    mongodbatlas:secretName:
        value: mongodbsecret
    mongodbatlas:stsEndpoint:
        value: https://sts.us-east-2.amazonaws.com/

```

Note: `awsAccessKeyId`, `awsSecretAccessKey`, and `awsSessionToken` can also be passed in using environment variables i.e. awsAccessKeyId will accept AWS_ACCESS_KEY_ID and TF_VAR_AWS_ACCESS_KEY_ID as a default value in place of value in a pulumi file variable.

Note: Fully qualified `secretName` ARN as input is REQUIRED for cross-AWS account secrets. For more detatils see:
* <https://aws.amazon.com/blogs/security/how-to-access-secrets-across-aws-accounts-by-attaching-resource-based-policies/>
* <https://aws.amazon.com/premiumsupport/knowledge-center/secrets-manager-share-between-accounts/>

Note: `stsEndpoint` parameter is REQUIRED for cross-AWS region or cross-AWS account secrets.

7. In terminal, `pulumi up`
### Static Credentials

Static credentials can be provided by adding the following attributes in-line in the MongoDB Atlas provider configuration,
either directly or via input variable/local value:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    mongodbatlas:privateKey:
        value: atlas_private_api_key
    mongodbatlas:publicKey:
        value: atlas_public_api_key

```

> *IMPORTANT* Hard-coding your MongoDB Atlas programmatic API key pair into a Pulumi configuration is not recommended.
Consider the risks, especially the inadvertent submission of a configuration file containing secrets to a public repository.
## Configuration Reference

In addition to generic `provider` arguments
(e.g. `alias` and `version`), the MongoDB Atlas `provider` supports the following arguments:

* `publicKey` - (Optional) This is the public key of your MongoDB Atlas API key pair. It must be
  provided, but it can also be sourced from the `MONGODB_ATLAS_PUBLIC_KEY` or `MCLI_PUBLIC_API_KEY`
  environment variable.

* `privateKey` - (Optional) This is the private key of your MongoDB Atlas key pair. It must be
  provided, but it can also be sourced from the `MONGODB_ATLAS_PRIVATE_KEY` or `MCLI_PRIVATE_API_KEY`
  environment variable.

For more information on configuring and managing programmatic API Keys see the [MongoDB Atlas Documentation](https://docs.atlas.mongodb.com/tutorial/manage-programmatic-access/index.html).
## MongoDB Atlas Provider Versioning Policy

In order to promote stability, predictability, and transparency, the MongoDB Atlas Pulumi Provider will implement **semantic versioning** with a **scheduled release cadence**. Our goal is to deliver regular improvements to the provider without overburdening users with frequent breaking changes.

---
### Definition of Breaking Changes

Our definition of breaking changes aligns with the impact updates have on the customer:

Breaking changes are defined as any change that requires user intervention to address.
This may include:

- Modifying existing schema (e.g., removing or renaming fields, renaming resources)
- Changes to business logic (e.g., implicit default values or server-side behavior)
- Provider-level changes (e.g., changing retry behavior)

Final confirmation of a breaking change — possibly leading to an exemption — is subject to:

- MongoDB’s understanding of the adoption level of the feature
- Timing of the next planned major release
- The change's relation to a bug fix

---
### Versioning Strategy

We follow [semantic versioning](https://semver.org/) for all updates:

- **Major (X.0.0):** Introduces breaking changes (as defined by MongoDB)
- **Minor (X.Y.0):** Adds non-breaking changes and announces deprecations
- **Patch (X.Y.Z):** Includes bug fixes and documentation updates

We do not utilize pre-release versioning at this time.

---
### Release Cadence

To minimize unexpected changes, we follow a scheduled cadence:

- **Minor and patch** versions follow a **biweekly** release pattern
- **Major** versions are released **once per year**, with a maximum of **two per calendar year**
- The provider team may adjust the schedule based on need

**Off-cycle releases** may occur for critical security flaws or regressions.

---
### Deprecation Policy

We use a structured deprecation window to notify customers in advance:

- Breaking changes are **deprecated in a minor version** with:
  - Warnings in migration guides, changelogs, and resource usage
- Deprecated functionality is **removed in the next 1–2 major versions**, unless otherwise stated

---
### Customer Communication

We are committed to clear and proactive communication:

- **Each release** includes a changelog clearly labeling:
  - `breaking`, `deprecated`, `bug-fix`, `feature`, and `enhancement` changes
- **Major versions** include migration guides
- **Minor and patch versions** generally do not include migration guides, but may if warranted
- **GitHub tags** with `vX.Y.Z` format are provided for all releases

---