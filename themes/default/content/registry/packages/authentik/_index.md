---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/goauthentik/authentik/2025.8.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Authentik Provider
meta_desc: Provides an overview on how to configure the Pulumi Authentik provider.
layout: package
---

## Generate Provider

The Authentik provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider goauthentik/authentik
```
## Overview

The authentik provider provides resources to interact with the authentik API.
## Example Usage

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    authentik:token:
        value: foo-bar
    authentik:url:
        value: https://authentik.company

```
### Configure provider with environment variables
It is optionally possible to configure the provider by passing environment variables to pulumi
```bash
export AUTHENTIK_URL=https://...
export AUTHENTIK_TOKEN=<secret_token>
export AUTHENTIK_INSECURE=false
```
## Configuration Reference
### Required

- `token` (String, Sensitive) The authentik API token, can optionally be passed as `AUTHENTIK_TOKEN` environmental variable
- `url` (String) The authentik API endpoint, can optionally be passed as `AUTHENTIK_URL` environmental variable

- `headers` (Map of String, Sensitive) Optional HTTP headers sent with every request
- `insecure` (Boolean) Whether to skip TLS verification, can optionally be passed as `AUTHENTIK_INSECURE` environmental variable