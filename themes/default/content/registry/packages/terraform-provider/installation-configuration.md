---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-terraform-provider/v1.0.2/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi/pulumi-terraform-provider/blob/v1.0.2/docs/installation-configuration.md
title: Any Terraform Provider Installation
meta_desc: How to set up and start using any Terraform provider with Pulumi.
layout: package
---

## Installation

You don't need to explicitly install anything, beyond an up to date version of Pulumi.[^1]
If you want to explicitly install the provider, you can run:

```console
$ pulumi plugin install resource terraform-provider
[resource plugin terraform-provider-0.1.0] installing
```

This will install the `terraform-provider` binary, but will not install any usable
SDKs. These need to be generated for an existing Terraform provider.

[^1]: Up to date means >= `v3.130.0`.

## Usage

Using `terraform-provider` will require an actual Terraform provider to bridge. The
recommended way to add a new Terraform provider is by running `pulumi package add
terraform-provider <provider ...>` in an existing Pulumi project.

You can specify `<provider ...>` in a couple of different ways:

1. By specify a provider's source and version:

	```console
    $ pulumi package add terraform-provider author/name
    ```

    You *must* specify an author and a provider name. You *may* specify a registry source
    and a version constraint. If no registry source is specified, then
    [registry.opentofu.org](https://opentofu.org/registry/) is assumed. If no version is specified, then the latest
    version is assumed.

    The full format is:

    ```console
    $ pulumi package add terraform-provider [<registry>/]<author>/<name> [version]
    ```

1. By specifying a local path:

    ```console
    $ pulumi package add terraform-provider /path/to/my/terraform-provider-binary
    ```

    Local paths must start with `./` or `/` and must end with `terraform-provider-<name>`.

{{% notes type="warning" %}}

The generated SDK will only work as long as it can find the Terraform provider binary 
at the specified path.

{{% /notes %}}


Regardless of the way you run `pulumi package add`, it will generate a local Pulumi SDK
for your Terraform provider and walk you through linking it in your project. You can then
consume it like a normal Pulumi provider SDK.
