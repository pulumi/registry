---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pierskarsenbarg/pulumi-nutanix/v0.11.0/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pierskarsenbarg/pulumi-nutanix/blob/v0.11.0/docs/installation-configuration.md
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

To provision resources with the Pulumi Nutanix provider, you need to authenticate using the configuration options as specified below.

## Configuration Options

Use `pulumi config set nutanix:<option> --secret`.

| Option     | Required/Optional | Description                                                                                                                                         |
| ---------- | ----------------- | --------------------------------------------------------------------------------------------------------------------------------------------------- |
| `username` | Required          | This is the username for the Prism Elements or Prism Central instance. This can also be specified with the `NUTANIX_USERNAME` environment variable. |
| `password` | Required          | This is the password for the Prism Elements or Prism Central instance. This can also be specified with the `NUTANIX_PASSWORD` environment variable. |
| `endpoint` | Required          | This is the endpoint for the Prism Elements or Prism Central instance. This can also be specified with the NUTANIX_ENDPOINT environment variable.   |
| `insecure` | Optional          | This specifies whether to allow verify ssl certificates. This can also be specified with `NUTANIX_INSECURE`. Defaults to `false`.                     |
| `port`     | Optional          | This is the port for the Prism Elements or Prism Central instance. This can also be specified with the `NUTANIX_PORT` environment variable. Defaults to `9440`. |

{{% notes type="warning" %}}
You should use the `--secret` flag to encrypt the config values using your secret provider. For more information on this, view the [Pulumi Configuration Secrets](https://www.pulumi.com/docs/intro/concepts/secrets/#secrets) section in the Pulumi docs.
{{% /notes %}}
