---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-cloudngfwaws/v0.1.3/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Palo Alto Networks Cloud NGFW for AWS Provider
meta_desc: Provides an overview on how to configure the Pulumi Palo Alto Networks Cloud NGFW for AWS provider.
layout: package
---
## Installation

The Palo Alto Networks Cloud NGFW for AWS provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/cloudngfwaws`](https://www.npmjs.com/package/@pulumi/cloudngfwaws)
* Python: [`pulumi-cloudngfwaws`](https://pypi.org/project/pulumi-cloudngfwaws/)
* Go: [`github.com/pulumi/pulumi-cloudngfwaws/sdk/go/cloudngfwaws`](https://github.com/pulumi/pulumi-cloudngfwaws)
* .NET: [`Pulumi.Cloudngfwaws`](https://www.nuget.org/packages/Pulumi.Cloudngfwaws)
* Java: [`com.pulumi/cloudngfwaws`](https://central.sonatype.com/artifact/com.pulumi/cloudngfwaws)
## Overview

The cloudngfwaws provider provides resources to manage your Palo Alto Networks Cloud Next-Gen Firewalls for AWS.

Palo Alto Networks Cloud NGFW combines best-in-class network security with cloud ease of use and delivers the leading ML-powered NGFW as a cloud-native service on AWS.

Refer to the changelog to see what's new.
## Enable API Access

Before this provider can be used, you must first enable API access.  The steps to do this can be found [here](https://pan.dev/cloudngfw/aws/api/).
## Authentication Workflow

This provider first authenticates against AWS, performing STS assume role. After that is successful, it then retrieves the JWTs for firewall and rulestack administration.

The AWS access key and secret key can be statically specified in the provider configuration or they will be picked up from the shared credentials file.
## AWS Config

The initial portion of the authentication workflow uses the AWS Go SDK, and as such control over the AWS portion of the authentication can be controlled using AWS environment variables (e.g. - `AWS_PROFILE` to specific a specific profile if multiple are present in your shared credentials file).  Please refer to the AWS documentation for more control on the AWS authentication.
## Example Usage



{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    cloudngfwaws:jsonConfigFile:
        value: ~/.cloudngfwaws_creds.json

```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    cloudngfwaws:jsonConfigFile:
        value: ~/.cloudngfwaws_creds.json

```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    cloudngfwaws:jsonConfigFile:
        value: ~/.cloudngfwaws_creds.json

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    cloudngfwaws:jsonConfigFile:
        value: ~/.cloudngfwaws_creds.json

```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    cloudngfwaws:jsonConfigFile:
        value: ~/.cloudngfwaws_creds.json

```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    cloudngfwaws:jsonConfigFile:
        value: ~/.cloudngfwaws_creds.json

```

{{% /choosable %}}
{{< /chooser >}}

Example JSON config file:

```shell
{
  "host": "api.us-east-1.aws.cloudngfw.paloaltonetworks.com",
  "region": "us-east-1",
  "arn": "arn:aws:iam::123456789:role/MyRole"
}
```
## Provider Parameter Priority

There are multiple ways to specify the provider's parameters.  If overlapping values are configured for the provider, then this is the resolution order:

1. Statically configured in the provider configuration
2. Environment variable (where applicable)
3. Taken from the JSON config file
## Configuration Reference

- `accessKey` (String) (Used for the initial `sts assume role`) AWS access key. Environment variable: `CLOUDNGFWAWS_ACCESS_KEY`. JSON conf file variable: `access-key`.
- `accountAdminArn` (String) The ARN allowing account admin permissions. Environment variable: `CLOUDNGFWAWS_ACCT_ADMIN_ARN`. JSON conf file variable: `account-admin-arn`.
- `arn` (String) The ARN allowing firewall, rulestack, and global rulestack admin permissions. Global rulestack admin permissions can be enabled only if the AWS account is onboarded by AWS Firewall Manager. Use 'lfa_arn' and 'lra_arn' if you want to enable only firewall and rulestack admin permissions. Environment variable: `CLOUDNGFWAWS_ARN`. JSON conf file variable: `arn`.
- `graArn` (String) The ARN allowing global rulestack admin permissions. Global rulestack admin permissions can be enabled only if the AWS account is onboarded by AWS Firewall Manager. 'gra_arn' is preferentially used over the `arn` param if both are specified. Environment variable: `CLOUDNGFWAWS_GRA_ARN`. JSON conf file variable: `gra-arn`.
- `headers` (Map of String) Additional HTTP headers to send with API calls. Environment variable: `CLOUDNGFWAWS_HEADERS`. JSON conf file variable: `headers`.
- `host` (String) The hostname of the API (default: `api.us-east-1.aws.cloudngfw.paloaltonetworks.com`). Environment variable: `CLOUDNGFWAWS_HOST`. JSON conf file variable: `host`.
- `jsonConfigFile` (String) Retrieve provider configuration from this JSON file.
- `lfaArn` (String) The ARN allowing firewall admin permissions. This is preferentially used over the `arn` param if both are specified. Environment variable: `CLOUDNGFWAWS_LFA_ARN`. JSON conf file variable: `lfa-arn`.
- `logging` (List of String) The logging options for the provider. Environment variable: `CLOUDNGFWAWS_LOGGING`. JSON conf file variable: `logging`.
- `lraArn` (String) The ARN allowing rulestack admin permissions. This is preferentially used over the `arn` param if both are specified. Environment variable: `CLOUDNGFWAWS_LRA_ARN`. JSON conf file variable: `lra-arn`.
- `mpRegion` (String) AWS management plane region. Environment variable: `CLOUDNGFWAWS_MP_REGION`. JSON conf file variable: `mpRegion`.
- `mpRegionHost` (String) AWS management plane MP region host Environment variable: `CLOUDNGFWAWS_MP_REGION_HOST`. JSON conf file variable: `mpRegionHost`.
- `profile` (String) (Used for the initial `sts assume role`) AWS PROFILE. Environment variable: `CLOUDNGFWAWS_PROFILE`. JSON conf file variable: `profile`.
- `protocol` (String) The protocol (defaults to `https`). Environment variable: `CLOUDNGFWAWS_PROTOCOL`. JSON conf file variable: `protocol`. Valid values are `https` or `http`.
- `region` (String) AWS region. Environment variable: `CLOUDNGFWAWS_REGION`. JSON conf file variable: `region`.
- `secretKey` (String) (Used for the initial `sts assume role`) AWS secret key. Environment variable: `CLOUDNGFWAWS_SECRET_KEY`. JSON conf file variable: `secret-key`.
- `skipVerifyCertificate` (Boolean) Skip verifying the SSL certificate. Environment variable: `CLOUDNGFWAWS_SKIP_VERIFY_CERTIFICATE`. JSON conf file variable: `skip-verify-certificate`.
- `syncMode` (Boolean) Enable synchronous mode while creating resources Environment variable: `CLOUDNGFWAWS_SYNC_MODE`. JSON conf file variable: `syncMode`.
- `timeout` (Number) The timeout for any single API call (default: `30`). Environment variable: `CLOUDNGFWAWS_TIMEOUT`. JSON conf file variable: `timeout`.
- `resourceTimeout` (Number) The timeout for pulumi resource create/update/delete operations (default: `7200s`). Environment variable: `CLOUDNGFWAWS_RESOURCE_TIMEOUT`. JSON conf file variable: `resourceTimeout`.