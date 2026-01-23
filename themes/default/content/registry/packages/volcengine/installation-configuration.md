---
# WARNING: this file was fetched from https://raw.githubusercontent.com/volcengine/pulumi-volcengine/v0.0.43/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/volcengine/pulumi-volcengine/blob/v0.0.43/docs/installation-configuration.md
title: Volcengine Installation & Configuration
meta_desc: Provides an overview of how to install and configure the Volcengine Provider
layout: package
---

The Pulumi Volcengine Provider uses Volcengine SDK to manage and provision resources.

## Installation

The Volcengine provider is available as a package in the following Pulumi languages:

- JavaScript/TypeScript: [@volcengine/pulumi](https://www.npmjs.com/package/@volcengine/pulumi)
- Python: [pulumi-volcengine](https://pypi.org/project/pulumi-volcengine/)
- Go: [github.com/volcengine/pulumi-volcengine/sdk/go](https://pkg.go.dev/github.com/volcengine/pulumi-volcengine/sdk)
- .Net: [Pulumi.Volcengine](https://www.nuget.org/packages/Pulumi.Volcengine)

### Provider Binary

The provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource volcengine --server github://api.github.com/volcengine
```

## Setup

The following configuration are available for the `volcengine` provider:
- `accessKey`: (Required) The API Secret ID.
- `secretKey`: (Required) The API Secret Key.
- `region`: (Required) The region in which to deploy resources.

## Configuration Options

Using `pulumi config set volcengine:<option> (--secret)`.

```bash
pulumi config set volcengine:accessKey <your_secret_id> --secret
pulumi config set volcengine:secretKey <your_secret_key> --secret
pulumi config set volcengine:region cn-beijing
```

You can also configure these value by environment variables:

```bash
export VOLCENGINE_ACCESS_KEY=<your_secret_id>
export VOLCENGINE_SECRET_KEY=<your_secret_key>
export VOLCENGINE_REGION=cn-beijing
```