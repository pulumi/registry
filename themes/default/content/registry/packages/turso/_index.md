---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/celest-dev/turso/0.2.3/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Turso Provider
meta_desc: Provides an overview on how to configure the Pulumi Turso provider.
layout: package
---

## Generate Provider

The Turso provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider celest-dev/turso
```
## Example Usage

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    turso:apiToken:
        value: <API_TOKEN>

```
## Configuration Reference
### Required

- `organization` (String) The name of the Turso organization

- `apiToken` (String, Sensitive) The API token to authenticate with Turso API. If not provided, the TURSO_API_TOKEN environment variable will be used. Finally, `turso auth token` is used to get the token.