---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/supabase/supabase/1.8.2/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Supabase Provider
meta_desc: Provides an overview on how to configure the Pulumi Supabase provider.
layout: package
---

## Generate Provider

The Supabase provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider supabase/supabase
```
## Example Usage

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:

```
## Configuration Reference

- `accessToken` (String, Sensitive) Supabase access token. Can also be set via the `SUPABASE_ACCESS_TOKEN` environment variable. When both are specified, the provider configuration takes precedence over the environment variable. Generate a token from the [Supabase Dashboard](https://supabase.com/dashboard/account/tokens).
- `endpoint` (String) Supabase API endpoint. Can also be set via the `SUPABASE_API_ENDPOINT` environment variable. If neither is set, defaults to `https://api.supabase.com`. When both are specified, the provider configuration takes precedence over the environment variable.