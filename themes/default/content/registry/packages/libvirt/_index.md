---
title: Libvirt Provider
meta_desc: Provides an overview on how to configure the Pulumi Libvirt provider.
layout: package
---
## Installation

The libvirt provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/libvirt`](https://www.npmjs.com/package/@pulumi/libvirt)
* Python: [`pulumi-libvirt`](https://pypi.org/project/pulumi-libvirt/)
* Go: [`github.com/pulumi/pulumi-libvirt/sdk/go/libvirt`](https://github.com/pulumi/pulumi-libvirt)
* .NET: [`Pulumi.Libvirt`](https://www.nuget.org/packages/Pulumi.Libvirt)
* Java: [`com.pulumi/libvirt`](https://central.sonatype.com/artifact/com.pulumi/libvirt)
## Overview

The Libvirt provider is used to interact with Linux
[libvirt](https://libvirt.org) hypervisors.

The provider needs to be configured with the proper connection information
before it can be used.

> **Note:** while libvirt can be used with several types of hypervisors, this
provider focuses on [KVM](http://libvirt.org/drvqemu.html). Other drivers may not be
working and haven't been tested.
## The connection URI

The provider understands [connection URIs](https://libvirt.org/uri.html). The supported transports are:

* `tcp` (non-encrypted connection)
* `unix` (UNIX domain socket)
* `tls` (See [here](https://libvirt.org/kbase/tlscerts.html) for information how to setup certificates)
* `ssh` (Secure shell)

Unlike the original libvirt, the `ssh` transport is not implemented using the ssh command and therefore does not require `nc` (netcat) on the server side.

Additionally, the `ssh` URI supports passwords using the `driver+ssh://[username:PASSWORD@][hostname][:port]/[path]?sshauth=ssh-password` syntax.

As the provider does not use libvirt on the client side, not all connection URI options are supported or apply.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    libvirt:uri:
        value: qemu:///system

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as libvirt from "@pulumi/libvirt";

// Create a new domain
const test1 = new libvirt.Domain("test1", {});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    libvirt:uri:
        value: qemu:///system

```
```python
import pulumi
import pulumi_libvirt as libvirt

# Create a new domain
test1 = libvirt.Domain("test1")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    libvirt:uri:
        value: qemu:///system

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Libvirt = Pulumi.Libvirt;

return await Deployment.RunAsync(() =>
{
    // Create a new domain
    var test1 = new Libvirt.Domain("test1");

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    libvirt:uri:
        value: qemu:///system

```
```go
package main

import (
	"github.com/pulumi/pulumi-libvirt/sdk/go/libvirt"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a new domain
		_, err := libvirt.NewDomain(ctx, "test1", nil)
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
config:
    libvirt:uri:
        value: qemu:///system

```
```yaml
resources:
  # Create a new domain
  test1:
    type: libvirt:Domain
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    libvirt:uri:
        value: qemu:///system

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.libvirt.Domain;
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
        // Create a new domain
        var test1 = new Domain("test1");

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

The following keys can be used to configure the provider.

* `uri` - (Required) The [connection URI](https://libvirt.org/uri.html) used
  to connect to the libvirt host.
## Environment variables

The libvirt connection URI can also be specified with the `LIBVIRT_DEFAULT_URI`
shell environment variable.

```
$ export LIBVIRT_DEFAULT_URI="qemu+ssh://root@192.168.1.100/system"
$ pulumi preview
```