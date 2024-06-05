---
title: Nutanix Installation & Configuration
meta_desc: Information on how to install the Nutanix provider.
layout: package
---

## Installation

The Pulumi Nutanix provider is available as a package from the following repositories:

- JavaScript/TypeScript: [`@pierskarsenbarg/nutanix`](https://www.npmjs.com/package/@pierskarsenbarg/nutanix)
- Python: [`pierskarsenbarg-pulumi-nutanix`](https://pypi.org/project/pulumi-nutanix/)
- Go: [`github.com/pierskarsenbarg/pulumi-nutanix/sdk/go/nutanix`](https://pkg.go.dev/github.com/pierskarsenbarg/pulumi-nutanix/sdk)
- .NET: [`PiersKarsenbarg.Nutanix`](https://www.nuget.org/packages/PiersKarsenbarg.Nutanix)

### Provider Binary

The Nutanix provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource nutanix <version> --server github://api.github.com/pierskarsenbarg
```

Replace the version string with your desired version.

## Setup

To provision resources with the Pulumi Nutanix provider, you need to provide the `apiKey` and `region`.

## Configuration Options

Use `pulumi config set nutanix:<option> --secret`.

| Option   | Required/Optional | Description                                                                                      |
| -------- | ----------------- | ------------------------------------------------------------------------------------------------ |
| `apiKey` | Required          | This is the user id that should be used to make the connection (environment: `NUTANIX_API_KEY`). |
| `region` | Required          | This is the password that should be used to make the connection (environment: `NUTANIX_REGION`). |

{{% notes type="warning" %}}
You should use the `--secret` flag to encrypt the config values using your secret provider. For more information on this, view the [Pulumi Configuration Secrets](https://www.pulumi.com/docs/intro/concepts/secrets/#secrets) section in the Pulumi docs.
{{% /notes %}}
