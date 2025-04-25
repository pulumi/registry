---
# WARNING: this file was fetched from https://raw.githubusercontent.com/ctfer-io/pulumi-ctfd/v2.1.0/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: CTFd Installation & Configuration
meta_desc: Provides an overview on how to setup the CTFd Provider for Pulumi.
layout: package
---

## Installation

The CTFd provider is available as a package in all Pulumi languages:

- JavaScript/TypeScript: [`@ctfer-io/pulumi-ctfd`](https://www.npmjs.com/package/@ctfer-io/pulumi-ctfd)
- Python: [`ctfer-io_pulumi-ctfd`](https://pypi.org/project/ctfer-io_pulumi-ctfd/)
- Golang: [`github.com/ctfer-io/pulumi-ctfd/sdk/v2/go/ctfd`](https://github.com/ctfer-io/pulumi-ctfd)
- .NET: [`CTFerio.Ctfd`](https://www.nuget.org/packages/CTFerio.Ctfd)

## Provider Binary

The CTFd provider binary is a third party binary. It can be installed using the pulumi plugin command.

```bash
pulumi plugin install resource ctfd <version> --server github://api.github.com/ctfer-io
```

Replace the version string with your desired version.

## Configuration

Pulumi relies on the CTFd REST JSON API to authenticate requests from the host machine to CTFd. Your credentials are never sent to pulumi.com.
The Pulumi CTFd Provider needs to be configured with CTFd credentials before it can be used to manage resources.

Your can either configure it using:

1. [environment variables](#environment-variables)
2. [provider configuration](#provider-configuration)

### Environment variables

Using the following you can configure the CTFd. For description of each variable, refer to the [provider configuration](#provider-configuration).

```bash
export CTFD_URL="https://my-ctf.lan"
export CTFD_API_KEY="ctfd_xxx"
# additional configuration
```

We recommend using an API key rather than a nonce/session combo for security purposes: the API key is natively supported thus enable better logging, authentication & authorization.
Moreover, it is possible to rotate the keys and revoke them on the fly using the API, while a session/nonce is not build in this way.

### Provider Configuration

You can also configure the provider using the stack configuration.
For instance, you can set the url using `pulumi config set url https://my-ctf.lan` then use it with, for instance, the following Go code.

```go
package main

import (
    "github.com/ctfer-io/pulumi-ctfd/sdk/v2/go/ctfd"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
        conf := config.New(ctx, "prod")

        // Create provider
        _, err := ctfd.NewProvider(ctx, "ctfd-fine-grained", &ctfd.ProviderArgs{
            Url:    pulumi.String(conf.Get("url")),
            ApiKey: conf.GetSecret("api_key"),
        })
        if err != nil {
            return err
        }

        // ...

        return nil
    })
}
```
