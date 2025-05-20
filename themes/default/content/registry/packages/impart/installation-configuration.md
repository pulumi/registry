---
# WARNING: this file was fetched from https://raw.githubusercontent.com/impart-security/pulumi-impart/v0.10.0/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Impart Installation & Configuration
meta_desc: Information on how to install the Impart provider.
layout: package
---

## Installation

The Pulumi Impart provider is available as a package from the following repositories:

- JavaScript/TypeScript: [`@impart-security/pulumi-impart`](https://www.npmjs.com/package/@impart-security/pulumi-impart)
- Go: [`github.com/impart-security/pulumi-impart/sdk/go/impart`](https://github.com/impart-security/pulumi-impart)

## Setup

To provision resources with the Pulumi Impart provider, you need to provide the `token`.
- Login to [https://console.impartsecurity.net](https://console.impartsecurity.net). Under the `Manage` section click `Settings => Access Tokens => New API access token`. Select the scope `read:org_api_access_tokens` along with scopes for resources that will be managed by Pulumi.

## Configuration Options

Use `pulumi config set impart:<option> --secret`.

| Option | Required? | Description |
| - | - | - |
| `token` | Required | This is the Impart API access token (environment: `IMPART_TOKEN`). |
| `endpoint` | Optional | The API URL used for Impart service. The default is `https://api.impartsecurity.net.` (environment: `IMPART_ENDPOINT`). |

{{% notes type="warning" %}}
You should use the `--secret` flag to encrypt the config values using your secret provider. For more information on this, view the [Pulumi Configuration Secrets](https://www.pulumi.com/docs/concepts/secrets/#secrets) section in the Pulumi docs.
{{% /notes %}}
