---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/prefecthq/prefect/2.87.1/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Prefect Provider
meta_desc: Provides an overview on how to configure the Pulumi Prefect provider.
layout: package
---

## Generate Provider

The Prefect provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider prefecthq/prefect
```
## Overview

Use the [Prefect](https://prefect.io) provider to configure your Prefect infrastructure.

The provider supports both [Prefect Cloud](https://app.prefect.cloud) and
[Prefect OSS](https://github.com/prefecthq/prefect).

The provider is tested against Prefect 3 and later. There is no compatibility guarantee for Prefect 2.

See getting started for more information on setting up the provider.

See troubleshoting for resources to address potential errors.
## Configuration Reference

- `accountId` (String) Default Prefect Cloud Account ID. Can also be set via the `PREFECT_CLOUD_ACCOUNT_ID` environment variable.
- `apiKey` (String, Sensitive) Prefect Cloud API key. Can also be set via the `PREFECT_API_KEY` environment variable.
- `basicAuthKey` (String, Sensitive) Prefect basic auth key. Can also be set via the `PREFECT_BASIC_AUTH_KEY` environment variable.
- `csrfEnabled` (Boolean) Enable CSRF protection for API requests. Defaults to false. If enabled, the provider will fetch a CSRF token from the Prefect API and include it in all requests. This should be enabled if your Prefect server instance has CSRF protection active. Can also be set via the `PREFECT_CSRF_ENABLED` environment variable.
- `endpoint` (String) The Prefect API URL. Can also be set via the `PREFECT_API_URL` environment variable. Defaults to `https://api.prefect.cloud` if not configured. Can optionally include the default account ID and workspace ID in the following format: `https://api.prefect.cloud/api/accounts/<accountID>/workspaces/<workspaceID>`. This is the same format used for the `PREFECT_API_URL` value in the Prefect CLI configuration file. The `accountId` and `workspaceId` attributes and their matching environment variables will take priority over any account and workspace ID values provided in the `endpoint` attribute.
- `profile` (String) Prefect profile name to use for authentication. If not specified, uses the active profile from `~/.prefect/profiles.toml`. This allows you to use a specific profile instead of the active one.
- `profileFile` (String) Path to the Prefect profiles file. If not specified, uses the default location `~/.prefect/profiles.toml`. This allows you to use a custom profiles file location.
- `workspaceId` (String) Default Prefect Cloud Workspace ID.