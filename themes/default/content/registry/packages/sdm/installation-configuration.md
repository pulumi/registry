---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pierskarsenbarg/pulumi-sdm/v1.33.0/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pierskarsenbarg/pulumi-sdm/blob/v1.33.0/docs/installation-configuration.md
title: StrongDM Installation & Configuration
meta_desc: Information on how to install the StrongDM provider.
layout: package
---

## Installation

The Pulumi StrongDM provider is available as a package from the following repositories:

* JavaScript/TypeScript: [`@pierskarsenbarg/sdm`](https://www.npmjs.com/package/@pierskarsenbarg/sdm)
* Python: [`pierskarsenbarg-pulumi-sdm`](https://pypi.org/project/pierskarsenbarg-pulumi-sdm/)
* Go: [`github.com/pierskarsenbarg/pulumi-sdm/sdk/go/sdm`](https://pkg.go.dev/github.com/pierskarsenbarg/pulumi-sdm/sdk)
* .NET: [`PiersKarsenbarg.Sdm`](https://www.nuget.org/packages/PiersKarsenbarg.Sdm)

## Setup

To provision resources with the Pulumi StrongDM provider, you need to provide the `apiAccessKey` and `apiSecretKey`. 

## Configuration Options

Use `pulumi config set sdm:<option> --secret`.

| Option     | Required/Optional | Description                                                                                     |
|------------|-------------------|-------------------------------------------------------------------------------------------------|
| `apiAccessKey` | Required          | This is the user id that should be used to make the connection (environment: `SDM_API_ACCESS_KEY`).      |
| `apiSecretKey` | Required          | This is the password that should be used to make the connection (environment: `SDM_API_SECRET_KEY`). |

{{% notes type="warning" %}}
You should use the `--secret` flag to encrypt the config values using your secret provider. For more information on this, view the [Pulumi Configuration Secrets](https://www.pulumi.com/docs/intro/concepts/secrets/#secrets) section in the Pulumi docs.
{{% /notes %}}
