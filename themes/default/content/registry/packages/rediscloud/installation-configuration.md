---
# WARNING: this file was fetched from https://raw.githubusercontent.com/RedisLabs/pulumi-rediscloud/v1.3.5/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Redis Cloud Installation & Configuration
meta_desc: Information on how to install the Redis Cloud provider.
layout: package
---

## Installation

The Pulumi Redis Cloud provider is available as a package in the following Pulumi languages:

- JavaScript/TypeScript: [`@rediscloud/pulumi-rediscloud`](https://www.npmjs.com/package/@rediscloud/pulumi-rediscloud)
- Python: [`pulumi_rediscloud`](https://pypi.org/project/pulumi-rediscloud/)
- Go: [`github.com/RedisLabs/pulumi-rediscloud/sdk/go/rediscloud`](https://pkg.go.dev/github.com/RedisLabs/pulumi-rediscloud/sdk)

### Provider Binary

The Redis Cloud provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource rediscloud <version> --server github://api.github.com/RedisLabs
```

Replace the version string with your desired version.

## Setup

To provision resources with the Pulumi Redis Cloud provider, you need to have Redis Cloud credentials. Redis maintains documentation on how to create API keys [here](https://docs.redis.com/latest/rc/security/access-control/access-management/)

### Set environment variables

Once you have provisioned these credentials, you can set environment variables to provision resources in Redis Cloud:

{{< chooser os "linux,macos,windows" >}}
{{% choosable os linux %}}

```bash
$ export REDISCLOUD_ACCESS_KEY=<REDISCLOUD_ACCESS_KEY>
$ export REDISCLOUD_SECRET_KEY=<REDISCLOUD_SECRET_KEY>
```

{{% /choosable %}}

{{% choosable os macos %}}

```bash
$ export REDISCLOUD_ACCESS_KEY=<REDISCLOUD_ACCESS_KEY>
$ export REDISCLOUD_SECRET_KEY=<REDISCLOUD_SECRET_KEY>
```

{{% /choosable %}}

{{% choosable os windows %}}

```powershell
> $env:REDISCLOUD_ACCESS_KEY = "<REDISCLOUD_ACCESS_KEY>"
> $env:REDISCLOUD_SECRET_KEY = "<REDISCLOUD_SECRET_KEY>"
```

{{% /choosable %}}
{{< /chooser >}}

## Configuration Options

Use `pulumi config set rediscloud:<option>` or pass options to the [constructor of `new rediscloud.Provider`]({{< relref "/registry/packages/rediscloud/api-docs/provider" >}}).

| Option      | Environment Variables   | Required/Optional | Description            |
| ----------- | ----------------------- | ----------------- | ---------------------- |
| `apiKey`    | `REDISCLOUD_ACCESS_KEY` | Required          | Redis Cloud access     |
| `secretKey` | `REDISCLOUD_SECRET_KEY` | Required          | Redis Cloud secret key |
