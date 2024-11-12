---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumiverse/pulumi-esxi-native/v1.0.0/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: ESXi Native Installation & Configuration
meta_desc: Information on how to install the Pulumi ESXi Native provider.
layout: package
---

## Installation

The Pulumi ESXi Native provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumiverse/esxi-native`](https://www.npmjs.com/package/@pulumiverse/esxi-native)
* Python: [`pulumiverse_esxi_native`](https://pypi.org/project/pulumiverse_esxi_native/)
* Go: [`github.com/pulumiverse/pulumi-esxi-native/sdk/go/esxi`](https://pkg.go.dev/github.com/pulumiverse/pulumi-esxi-native/sdk/go/esxi)
* .NET: [`Pulumiverse.EsxiNative`](https://www.nuget.org/packages/Pulumiverse.EsxiNative)

### Provider Binary

The ESXi Native provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource esxi-native <version> --server github://api.github.com/pulumiverse
```

Replace the `<version>` string with your desired version.

## Configuration

In order to use the provider, we need to provide SSH credentials to the ESXi Host

### Set configuration using `pulumi config`

Remember to pass `--secret` when setting `password` so that it is properly encrypted:

    ```bash
    $ pulumi config set esxi-native:username <username>
    $ pulumi config set esxi-native:password <password> --secret
    $ pulumi config set esxi-native:host <host IP or FQDN>
    ```

### Set configuration using environment variables

{{< chooser os "linux,macos,windows" >}}
{{% choosable os linux %}}

```bash
$ export ESXI_USERNAME=<YOUR_ESXI_USERNAME>
$ export ESXI_PASSWORD=<YOUR_ESXI_PASSWORD>
$ export ESXI_HOST=<YOUR_ESXI_HOST_IP>
```

{{% /choosable %}}

{{% choosable os macos %}}

```bash
$ export ESXI_USERNAME=<YOUR_ESXI_USERNAME>
$ export ESXI_PASSWORD=<YOUR_ESXI_PASSWORD>
$ export ESXI_HOST=<YOUR_ESXI_HOST_IP>
```

{{% /choosable %}}

{{% choosable os windows %}}

```powershell
> $env:ESXI_USERNAME = "<YOUR_ESXI_USERNAME>"
> $env:ESXI_PASSWORD = "<YOUR_ESXI_PASSWORD>"
> $env:ESXI_HOST = "<YOUR_ESXI_HOST>"
```

{{% /choosable %}}
{{< /chooser >}}

### Configuration options

Use `pulumi config set esxi-native:<option>` or pass options to the [constructor of `new esxi-native.Provider`](/registry/packages/esxi-native/api-docs/provider).

> Note: Each config can also be sourced from the environment variables given below

| Option     | Required? | Description                                    | Default | Env. Variable   |
|------------|-----------|------------------------------------------------|---------|-----------------|
| `username` | Required  | The ESXi Username                              |         | `ESXI_USERNAME` |
| `password` | Required  | The ESXi Password, has support for secrets too |         | `ESXI_PASSWORD` |
| `host`     | Required  | The ESXi Host Name where to connect            |         | `ESXI_HOST`     |
| `sshPort`  | Optional  | The ESXi Host SSH Port where to connect        | `22`    | `ESXI_SSH_PORT` |
| `sslPort`  | Optional  | The ESXi Host SSL Port where to connect        | `443`   | `ESXI_SSL_PORT` |
