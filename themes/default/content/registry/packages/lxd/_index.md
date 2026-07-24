---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/terraform-lxd/lxd/3.0.2/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Lxd Provider
meta_desc: Provides an overview on how to configure the Pulumi Lxd provider.
layout: package
---

## Generate Provider

The Lxd provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider terraform-lxd/lxd
```
## Overview

The LXD provider allows infrastructure as code tools to manage resources on [LXD](https://documentation.ubuntu.com/lxd/latest/) servers, such as instances, networks, storage, and more.

LXD is a modern, secure, and powerful system container and virtual machine manager.
If you are new to LXD, see the [Getting started guide](https://documentation.ubuntu.com/lxd/latest/tutorial/first_steps/) in the official documentation.

Minimum supported LXD version is **5.0**.
## Getting Started
### Prerequisites

1. A running LXD server. See [How to install LXD](https://documentation.ubuntu.com/lxd/latest/installing/).
2. An installed infrastructure as code tool, such as Pulumi.
3. Authentication credentials for connecting to your LXD server (see Authentication below).
### Minimal Example

The following configuration launches an Ubuntu container on a local LXD server:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as lxd from "@pulumi/lxd";

const myContainer = new lxd.Instance("my_container", {
    name: "my-container",
    image: "ubuntu-daily:24.04",
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python

```
```python
import pulumi
import pulumi_lxd as lxd

my_container = lxd.Instance("my_container",
    name="my-container",
    image="ubuntu-daily:24.04")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Lxd = Pulumi.Lxd;

return await Deployment.RunAsync(() =>
{
    var myContainer = new Lxd.Instance("my_container", new()
    {
        Name = "my-container",
        Image = "ubuntu-daily:24.04",
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/lxd/v3/lxd"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lxd.NewInstance(ctx, "my_container", &lxd.InstanceArgs{
			Name:  pulumi.String("my-container"),
			Image: pulumi.String("ubuntu-daily:24.04"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml

```
```yaml
resources:
  myContainer:
    type: lxd:Instance
    name: my_container
    properties:
      name: my-container
      image: ubuntu-daily:24.04
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.lxd.Instance;
import com.pulumi.lxd.InstanceArgs;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        var myContainer = new Instance("myContainer", InstanceArgs.builder()
            .name("my-container")
            .image("ubuntu-daily:24.04")
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

Save this as `main.tf`, then:

```shell
pulumi up
pulumi up
```
## Provider Configuration

The provider connects to the LXD daemon via local Unix socket or HTTPS.

All LXD remotes used by the provider must be explicitly defined in the provider configuration.
The LXD built-in image remotes (such as `ubuntu:` and `images:`) are predefined and do not need to be manually configured.
For more information on image remotes, see [Remote image servers](https://documentation.ubuntu.com/lxd/latest/reference/remote_image_servers/).
### Authentication

LXD supports multiple authentication methods. See [Remote API authentication](https://documentation.ubuntu.com/lxd/latest/authentication/) in the LXD documentation for background.

This provider supports the following methods:

- **Bearer token** - For remote servers that support API extension `authBearer`. LXD bearer tokens also embed the server certificate fingerprint, so `serverCertificateFingerprint` does not need to be set separately.
- **Mutual TLS (mTLS)** - Client certificate authentication. Requires a client certificate that is already trusted by the server, or a trust token to bootstrap trust on the first connection.
- **Unix socket** - For local connections. Requires access to the local LXD unix socket.
#### Handling Sensitive Information

> Avoid passing sensitive values with `-var` on the command line, as they may be stored in shell history.
Prefer interactive input or a local `pulumi.tfvars` file instead.

When providing sensitive values, such as tokens or certificates, through Pulumi configuration, use variables marked as sensitive and ephemeral.
This prevents values from being shown in interactive prompts or Pulumi output, and avoids storing them in pulumi preview or state files.

Alternatively, the provider can source sensitive values from local files using the `*_file` variants (e.g. `bearerTokenFile`, `clientCertificateFile`, `clientKeyFile`).
#### Unix Socket

Connect to a local LXD server via Unix socket.
Setting the remote address to `unix://` instructs the provider to search for a local LXD Unix socket in the standard locations.

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:

```
#### Bearer Token Authentication

Authenticate with an LXD server using a bearer token.
See [Bearer token authentication](https://documentation.ubuntu.com/lxd/latest/authentication/#bearer-token-authentication) for background and setup instructions.
#### Mutual TLS Authentication

Provide the client certificate and key. The client certificate must already be [trusted by the LXD server](https://documentation.ubuntu.com/lxd/latest/authentication/#tls-client-certificates).

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:

```

If the server certificate is self-signed or not otherwise trusted by the client, set `serverCertificateFingerprint` so the provider can verify the server identity. Retrieve the fingerprint with `lxc info` or by calling the LXD `/1.0` API endpoint.
##### Bootstrap mTLS Using a Trust Token

For a first-time connection, a [trust token](https://documentation.ubuntu.com/lxd/latest/howto/server_expose/#authenticate-with-the-lxd-server) can bootstrap trust. The token allows the server to add the client certificate to its trust store automatically, after which subsequent connections use mTLS.

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:

```
### Multiple Remotes

When defining multiple remotes, set `defaultRemote` to specify which remote is used when one is not specified in a resource:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    lxd:defaultRemote:
        value: lxd-server-1

```

When only one remote is defined, it is automatically used as the default remote.
## Configuration Reference
### Provider Arguments

* `remote` - **Required** - Defines a LXD or simplestreams remote the provider can use. At least one remote must be defined. See the `remote` block reference below.

* `defaultRemote` - *Optional* - Name of the default LXD remote to use when no remote is specified in a resource. Required when two or more remotes are defined.
### `remote` Block

* `name` - **Required** - The name of the remote.

* `address` - **Required** - The remote address. Must start with `https://` for HTTPS connections or `unix://` for Unix socket connections.

* `protocol` - *Optional* - The protocol of remote server (`lxd` or `simplestreams`). Defaults to `lxd`.

* `bearerToken` - *Optional* - Bearer token for authentication.

* `bearerTokenFile` - *Optional* - Path to a file containing the bearer token.

* `clientCertificate` - *Optional* - PEM-encoded client certificate for mTLS authentication. Must be provided together with `clientKey` or `clientKeyFile`.

* `clientCertificateFile` - *Optional* - Path to the PEM-encoded client certificate file. Must be provided together with `clientKey` or `clientKeyFile`.

* `clientKey` - *Optional* - PEM-encoded private key for mTLS authentication. Must be provided together with `clientCertificate` or `clientCertificateFile`.

* `clientKeyFile` - *Optional* - Path to the PEM-encoded private key file. Must be provided together with `clientCertificate` or `clientCertificateFile`.

* `serverCertificateFingerprint` - *Optional* - SHA-256 fingerprint of the remote server's TLS certificate. Used to pin and verify the server certificate.

* `trustToken` - *Optional* - Trust token for adding the client certificate to the server's trust store on first connection. Used together with `clientCertificate`/`clientCertificateFile` and `clientKey`/`clientKeyFile`.