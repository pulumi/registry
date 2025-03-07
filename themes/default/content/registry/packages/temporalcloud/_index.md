---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/registry/refs/heads/iwahbe/preview-wave-2/schemas/temporalcloud.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Temporalcloud Provider
meta_desc: Provides an overview on how to configure the Pulumi Temporalcloud provider.
layout: package
---

## Generate Provider

The Temporalcloud provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider temporalio/temporalcloud
```
## Overview

Use the `temporalcloud` provider to interact with resources supported by [Temporal Cloud](https://temporal.io/cloud).

Use the navigation to the left to learn about the available resources supported by this provider.

> This provider is in Public Preview, is under active development, and is subject to change. We reserve the right to make breaking changes during this pre-GA period, though we will do our best to maintain compatibility wherever possible.
## Provider Configuration

Credentials for Temporal Cloud can be provided by adding an `apiKey` property or by setting the environment variable `TEMPORAL_CLOUD_API_KEY`.
You can generate an API key for Temporal Cloud by following [these instructions](https://docs.temporal.io/cloud/api-keys).

!> Hard-coded credentials are not recommended in any Pulumi configuration and should not be committed
in version control. We recommend passing credentials to this provider via environment variables.
## Example Usage

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    temporalcloud:allowInsecure:
        value: false
    temporalcloud:allowedAccountId:
        value: my-temporalcloud-account-id
    temporalcloud:apiKey:
        value: my-temporalcloud-api-key
    temporalcloud:endpoint:
        value: saas-api.tmprl.cloud:443

```
## Configuration Reference

- `allowInsecure` (Boolean) If set to True, it allows for an insecure connection to the Temporal Cloud API. This should never be set to 'true' in production and defaults to false.
- `allowedAccountId` (String) The ID of the account to operate on. Prevents accidental mutation of accounts other than that provided.
- `apiKey` (String, Sensitive) The API key for Temporal Cloud. See [this documentation](https://docs.temporal.io/cloud/api-keys) for information on how to obtain an API key.
- `endpoint` (String) The endpoint for the Temporal Cloud API. Defaults to `saas-api.tmprl.cloud:443`.