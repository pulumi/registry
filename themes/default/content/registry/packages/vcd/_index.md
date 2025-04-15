---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/vmware/vcd/3.14.1/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Vcd Provider
meta_desc: Provides an overview on how to configure the Pulumi Vcd provider.
layout: package
---

## Generate Provider

The Vcd provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider vmware/vcd
```
## Overview

The VMware Cloud Director provider is used to interact with the resources supported by VMware Cloud Director. The provider needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources. Please refer to
CHANGELOG.md
to track feature additions.

> **NOTE:** The VMware Cloud Director Provider documentation pages include *v2.x+* or *v3.x+* labels in resource and/or field
descriptions. These labels are designed to show at which provider version a certain feature was introduced.
When upgrading the provider please check for such labels for the resources you are using.
## Supported VCD Versions

The following Cloud Director versions are supported by this provider:

* 10.4
* 10.5
* 10.6

Also Cloud Director Service (CDS) is supported.
## Connecting as Org Admin

The most common - tenant - use case when you set user to organization administrator and when all resources are in a single organization.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    vcd:allowUnverifiedSsl:
        value: 'TODO: var.vcd_allow_unverified_ssl'
    vcd:authType:
        value: integrated
    vcd:maxRetryTimeout:
        value: 'TODO: var.vcd_max_retry_timeout'
    vcd:org:
        value: 'TODO: var.vcd_org'
    vcd:password:
        value: 'TODO: var.vcd_pass'
    vcd:url:
        value: 'TODO: var.vcd_url'
    vcd:user:
        value: 'TODO: var.vcd_user'
    vcd:vdc:
        value: 'TODO: var.vcd_vdc'

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as vcd from "@pulumi/vcd";

// Create a new network in organization and VDC defined above
const net = new vcd.NetworkRouted("net", {});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    vcd:allowUnverifiedSsl:
        value: 'TODO: var.vcd_allow_unverified_ssl'
    vcd:authType:
        value: integrated
    vcd:maxRetryTimeout:
        value: 'TODO: var.vcd_max_retry_timeout'
    vcd:org:
        value: 'TODO: var.vcd_org'
    vcd:password:
        value: 'TODO: var.vcd_pass'
    vcd:url:
        value: 'TODO: var.vcd_url'
    vcd:user:
        value: 'TODO: var.vcd_user'
    vcd:vdc:
        value: 'TODO: var.vcd_vdc'

```
```python
import pulumi
import pulumi_vcd as vcd

# Create a new network in organization and VDC defined above
net = vcd.NetworkRouted("net")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    vcd:allowUnverifiedSsl:
        value: 'TODO: var.vcd_allow_unverified_ssl'
    vcd:authType:
        value: integrated
    vcd:maxRetryTimeout:
        value: 'TODO: var.vcd_max_retry_timeout'
    vcd:org:
        value: 'TODO: var.vcd_org'
    vcd:password:
        value: 'TODO: var.vcd_pass'
    vcd:url:
        value: 'TODO: var.vcd_url'
    vcd:user:
        value: 'TODO: var.vcd_user'
    vcd:vdc:
        value: 'TODO: var.vcd_vdc'

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Vcd = Pulumi.Vcd;

return await Deployment.RunAsync(() =>
{
    // Create a new network in organization and VDC defined above
    var net = new Vcd.NetworkRouted("net");

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    vcd:allowUnverifiedSsl:
        value: 'TODO: var.vcd_allow_unverified_ssl'
    vcd:authType:
        value: integrated
    vcd:maxRetryTimeout:
        value: 'TODO: var.vcd_max_retry_timeout'
    vcd:org:
        value: 'TODO: var.vcd_org'
    vcd:password:
        value: 'TODO: var.vcd_pass'
    vcd:url:
        value: 'TODO: var.vcd_url'
    vcd:user:
        value: 'TODO: var.vcd_user'
    vcd:vdc:
        value: 'TODO: var.vcd_vdc'

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/vcd/v3/vcd"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a new network in organization and VDC defined above
		_, err := vcd.NewNetworkRouted(ctx, "net", nil)
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
    vcd:allowUnverifiedSsl:
        value: 'TODO: var.vcd_allow_unverified_ssl'
    vcd:authType:
        value: integrated
    vcd:maxRetryTimeout:
        value: 'TODO: var.vcd_max_retry_timeout'
    vcd:org:
        value: 'TODO: var.vcd_org'
    vcd:password:
        value: 'TODO: var.vcd_pass'
    vcd:url:
        value: 'TODO: var.vcd_url'
    vcd:user:
        value: 'TODO: var.vcd_user'
    vcd:vdc:
        value: 'TODO: var.vcd_vdc'

```
```yaml
resources:
  # Create a new network in organization and VDC defined above
  net:
    type: vcd:NetworkRouted
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    vcd:allowUnverifiedSsl:
        value: 'TODO: var.vcd_allow_unverified_ssl'
    vcd:authType:
        value: integrated
    vcd:maxRetryTimeout:
        value: 'TODO: var.vcd_max_retry_timeout'
    vcd:org:
        value: 'TODO: var.vcd_org'
    vcd:password:
        value: 'TODO: var.vcd_pass'
    vcd:url:
        value: 'TODO: var.vcd_url'
    vcd:user:
        value: 'TODO: var.vcd_user'
    vcd:vdc:
        value: 'TODO: var.vcd_vdc'

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.vcd.NetworkRouted;
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
        // Create a new network in organization and VDC defined above
        var net = new NetworkRouted("net");

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Connecting as Sys Admin

When you want to manage resources across different organizations from a single configuration.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    vcd:allowUnverifiedSsl:
        value: 'TODO: var.vcd_allow_unverified_ssl'
    vcd:authType:
        value: integrated
    vcd:maxRetryTimeout:
        value: 'TODO: var.vcd_max_retry_timeout'
    vcd:org:
        value: System
    vcd:password:
        value: 'TODO: var.vcd_pass'
    vcd:url:
        value: 'TODO: var.vcd_url'
    vcd:user:
        value: administrator

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as vcd from "@pulumi/vcd";

// Create a new network in some organization and VDC
const net1 = new vcd.NetworkRouted("net1", {
    org: "Org1",
    vdc: "Org1VDC",
});
// Create a new network in a different organization and VDC
const net2 = new vcd.NetworkRouted("net2", {
    org: "Org2",
    vdc: "Org2VDC",
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    vcd:allowUnverifiedSsl:
        value: 'TODO: var.vcd_allow_unverified_ssl'
    vcd:authType:
        value: integrated
    vcd:maxRetryTimeout:
        value: 'TODO: var.vcd_max_retry_timeout'
    vcd:org:
        value: System
    vcd:password:
        value: 'TODO: var.vcd_pass'
    vcd:url:
        value: 'TODO: var.vcd_url'
    vcd:user:
        value: administrator

```
```python
import pulumi
import pulumi_vcd as vcd

# Create a new network in some organization and VDC
net1 = vcd.NetworkRouted("net1",
    org="Org1",
    vdc="Org1VDC")
# Create a new network in a different organization and VDC
net2 = vcd.NetworkRouted("net2",
    org="Org2",
    vdc="Org2VDC")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    vcd:allowUnverifiedSsl:
        value: 'TODO: var.vcd_allow_unverified_ssl'
    vcd:authType:
        value: integrated
    vcd:maxRetryTimeout:
        value: 'TODO: var.vcd_max_retry_timeout'
    vcd:org:
        value: System
    vcd:password:
        value: 'TODO: var.vcd_pass'
    vcd:url:
        value: 'TODO: var.vcd_url'
    vcd:user:
        value: administrator

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Vcd = Pulumi.Vcd;

return await Deployment.RunAsync(() =>
{
    // Create a new network in some organization and VDC
    var net1 = new Vcd.NetworkRouted("net1", new()
    {
        Org = "Org1",
        Vdc = "Org1VDC",
    });

    // Create a new network in a different organization and VDC
    var net2 = new Vcd.NetworkRouted("net2", new()
    {
        Org = "Org2",
        Vdc = "Org2VDC",
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    vcd:allowUnverifiedSsl:
        value: 'TODO: var.vcd_allow_unverified_ssl'
    vcd:authType:
        value: integrated
    vcd:maxRetryTimeout:
        value: 'TODO: var.vcd_max_retry_timeout'
    vcd:org:
        value: System
    vcd:password:
        value: 'TODO: var.vcd_pass'
    vcd:url:
        value: 'TODO: var.vcd_url'
    vcd:user:
        value: administrator

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/vcd/v3/vcd"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a new network in some organization and VDC
		_, err := vcd.NewNetworkRouted(ctx, "net1", &vcd.NetworkRoutedArgs{
			Org: pulumi.String("Org1"),
			Vdc: pulumi.String("Org1VDC"),
		})
		if err != nil {
			return err
		}
		// Create a new network in a different organization and VDC
		_, err = vcd.NewNetworkRouted(ctx, "net2", &vcd.NetworkRoutedArgs{
			Org: pulumi.String("Org2"),
			Vdc: pulumi.String("Org2VDC"),
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
config:
    vcd:allowUnverifiedSsl:
        value: 'TODO: var.vcd_allow_unverified_ssl'
    vcd:authType:
        value: integrated
    vcd:maxRetryTimeout:
        value: 'TODO: var.vcd_max_retry_timeout'
    vcd:org:
        value: System
    vcd:password:
        value: 'TODO: var.vcd_pass'
    vcd:url:
        value: 'TODO: var.vcd_url'
    vcd:user:
        value: administrator

```
```yaml
resources:
  # Create a new network in some organization and VDC
  net1:
    type: vcd:NetworkRouted
    properties:
      org: Org1
      vdc: Org1VDC
  # Create a new network in a different organization and VDC
  net2:
    type: vcd:NetworkRouted
    properties:
      org: Org2
      vdc: Org2VDC
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    vcd:allowUnverifiedSsl:
        value: 'TODO: var.vcd_allow_unverified_ssl'
    vcd:authType:
        value: integrated
    vcd:maxRetryTimeout:
        value: 'TODO: var.vcd_max_retry_timeout'
    vcd:org:
        value: System
    vcd:password:
        value: 'TODO: var.vcd_pass'
    vcd:url:
        value: 'TODO: var.vcd_url'
    vcd:user:
        value: administrator

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.vcd.NetworkRouted;
import com.pulumi.vcd.NetworkRoutedArgs;
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
        // Create a new network in some organization and VDC
        var net1 = new NetworkRouted("net1", NetworkRoutedArgs.builder()
            .org("Org1")
            .vdc("Org1VDC")
            .build());

        // Create a new network in a different organization and VDC
        var net2 = new NetworkRouted("net2", NetworkRoutedArgs.builder()
            .org("Org2")
            .vdc("Org2VDC")
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Connecting as Sys Admin with Default Org and VDC

When you want to manage resources across different organizations but set a default one.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    vcd:allowUnverifiedSsl:
        value: 'TODO: var.vcd_allow_unverified_ssl'
    vcd:authType:
        value: integrated
    vcd:maxRetryTimeout:
        value: 'TODO: var.vcd_max_retry_timeout'
    vcd:org:
        value: 'TODO: var.vcd_org'
    vcd:password:
        value: 'TODO: var.vcd_pass'
    vcd:sysorg:
        value: System
    vcd:url:
        value: 'TODO: var.vcd_url'
    vcd:user:
        value: administrator
    vcd:vdc:
        value: 'TODO: var.vcd_vdc'

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as vcd from "@pulumi/vcd";

// Create a new network in the default organization and VDC
const net1 = new vcd.NetworkRouted("net1", {});
// Create a new network in a specific organization and VDC
const net2 = new vcd.NetworkRouted("net2", {
    org: "OrgZ",
    vdc: "OrgZVDC",
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    vcd:allowUnverifiedSsl:
        value: 'TODO: var.vcd_allow_unverified_ssl'
    vcd:authType:
        value: integrated
    vcd:maxRetryTimeout:
        value: 'TODO: var.vcd_max_retry_timeout'
    vcd:org:
        value: 'TODO: var.vcd_org'
    vcd:password:
        value: 'TODO: var.vcd_pass'
    vcd:sysorg:
        value: System
    vcd:url:
        value: 'TODO: var.vcd_url'
    vcd:user:
        value: administrator
    vcd:vdc:
        value: 'TODO: var.vcd_vdc'

```
```python
import pulumi
import pulumi_vcd as vcd

# Create a new network in the default organization and VDC
net1 = vcd.NetworkRouted("net1")
# Create a new network in a specific organization and VDC
net2 = vcd.NetworkRouted("net2",
    org="OrgZ",
    vdc="OrgZVDC")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    vcd:allowUnverifiedSsl:
        value: 'TODO: var.vcd_allow_unverified_ssl'
    vcd:authType:
        value: integrated
    vcd:maxRetryTimeout:
        value: 'TODO: var.vcd_max_retry_timeout'
    vcd:org:
        value: 'TODO: var.vcd_org'
    vcd:password:
        value: 'TODO: var.vcd_pass'
    vcd:sysorg:
        value: System
    vcd:url:
        value: 'TODO: var.vcd_url'
    vcd:user:
        value: administrator
    vcd:vdc:
        value: 'TODO: var.vcd_vdc'

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Vcd = Pulumi.Vcd;

return await Deployment.RunAsync(() =>
{
    // Create a new network in the default organization and VDC
    var net1 = new Vcd.NetworkRouted("net1");

    // Create a new network in a specific organization and VDC
    var net2 = new Vcd.NetworkRouted("net2", new()
    {
        Org = "OrgZ",
        Vdc = "OrgZVDC",
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    vcd:allowUnverifiedSsl:
        value: 'TODO: var.vcd_allow_unverified_ssl'
    vcd:authType:
        value: integrated
    vcd:maxRetryTimeout:
        value: 'TODO: var.vcd_max_retry_timeout'
    vcd:org:
        value: 'TODO: var.vcd_org'
    vcd:password:
        value: 'TODO: var.vcd_pass'
    vcd:sysorg:
        value: System
    vcd:url:
        value: 'TODO: var.vcd_url'
    vcd:user:
        value: administrator
    vcd:vdc:
        value: 'TODO: var.vcd_vdc'

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/vcd/v3/vcd"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a new network in the default organization and VDC
		_, err := vcd.NewNetworkRouted(ctx, "net1", nil)
		if err != nil {
			return err
		}
		// Create a new network in a specific organization and VDC
		_, err = vcd.NewNetworkRouted(ctx, "net2", &vcd.NetworkRoutedArgs{
			Org: pulumi.String("OrgZ"),
			Vdc: pulumi.String("OrgZVDC"),
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
config:
    vcd:allowUnverifiedSsl:
        value: 'TODO: var.vcd_allow_unverified_ssl'
    vcd:authType:
        value: integrated
    vcd:maxRetryTimeout:
        value: 'TODO: var.vcd_max_retry_timeout'
    vcd:org:
        value: 'TODO: var.vcd_org'
    vcd:password:
        value: 'TODO: var.vcd_pass'
    vcd:sysorg:
        value: System
    vcd:url:
        value: 'TODO: var.vcd_url'
    vcd:user:
        value: administrator
    vcd:vdc:
        value: 'TODO: var.vcd_vdc'

```
```yaml
resources:
  # Create a new network in the default organization and VDC
  net1:
    type: vcd:NetworkRouted
  # Create a new network in a specific organization and VDC
  net2:
    type: vcd:NetworkRouted
    properties:
      org: OrgZ
      vdc: OrgZVDC
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    vcd:allowUnverifiedSsl:
        value: 'TODO: var.vcd_allow_unverified_ssl'
    vcd:authType:
        value: integrated
    vcd:maxRetryTimeout:
        value: 'TODO: var.vcd_max_retry_timeout'
    vcd:org:
        value: 'TODO: var.vcd_org'
    vcd:password:
        value: 'TODO: var.vcd_pass'
    vcd:sysorg:
        value: System
    vcd:url:
        value: 'TODO: var.vcd_url'
    vcd:user:
        value: administrator
    vcd:vdc:
        value: 'TODO: var.vcd_vdc'

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.vcd.NetworkRouted;
import com.pulumi.vcd.NetworkRoutedArgs;
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
        // Create a new network in the default organization and VDC
        var net1 = new NetworkRouted("net1");

        // Create a new network in a specific organization and VDC
        var net2 = new NetworkRouted("net2", NetworkRoutedArgs.builder()
            .org("OrgZ")
            .vdc("OrgZVDC")
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Connecting with authorization or bearer token

You can connect using an authorization token instead of username and password.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    vcd:allowUnverifiedSsl:
        value: 'TODO: var.vcd_allow_unverified_ssl'
    vcd:authType:
        value: token
    vcd:maxRetryTimeout:
        value: 'TODO: var.vcd_max_retry_timeout'
    vcd:org:
        value: 'TODO: var.vcd_org'
    vcd:password:
        value: none
    vcd:sysorg:
        value: System
    vcd:token:
        value: 'TODO: var.token'
    vcd:url:
        value: 'TODO: var.vcd_url'
    vcd:user:
        value: none
    vcd:vdc:
        value: 'TODO: var.vcd_vdc'

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as vcd from "@pulumi/vcd";

// Create a new network in the default organization and VDC
const net1 = new vcd.NetworkRouted("net1", {});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    vcd:allowUnverifiedSsl:
        value: 'TODO: var.vcd_allow_unverified_ssl'
    vcd:authType:
        value: token
    vcd:maxRetryTimeout:
        value: 'TODO: var.vcd_max_retry_timeout'
    vcd:org:
        value: 'TODO: var.vcd_org'
    vcd:password:
        value: none
    vcd:sysorg:
        value: System
    vcd:token:
        value: 'TODO: var.token'
    vcd:url:
        value: 'TODO: var.vcd_url'
    vcd:user:
        value: none
    vcd:vdc:
        value: 'TODO: var.vcd_vdc'

```
```python
import pulumi
import pulumi_vcd as vcd

# Create a new network in the default organization and VDC
net1 = vcd.NetworkRouted("net1")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    vcd:allowUnverifiedSsl:
        value: 'TODO: var.vcd_allow_unverified_ssl'
    vcd:authType:
        value: token
    vcd:maxRetryTimeout:
        value: 'TODO: var.vcd_max_retry_timeout'
    vcd:org:
        value: 'TODO: var.vcd_org'
    vcd:password:
        value: none
    vcd:sysorg:
        value: System
    vcd:token:
        value: 'TODO: var.token'
    vcd:url:
        value: 'TODO: var.vcd_url'
    vcd:user:
        value: none
    vcd:vdc:
        value: 'TODO: var.vcd_vdc'

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Vcd = Pulumi.Vcd;

return await Deployment.RunAsync(() =>
{
    // Create a new network in the default organization and VDC
    var net1 = new Vcd.NetworkRouted("net1");

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    vcd:allowUnverifiedSsl:
        value: 'TODO: var.vcd_allow_unverified_ssl'
    vcd:authType:
        value: token
    vcd:maxRetryTimeout:
        value: 'TODO: var.vcd_max_retry_timeout'
    vcd:org:
        value: 'TODO: var.vcd_org'
    vcd:password:
        value: none
    vcd:sysorg:
        value: System
    vcd:token:
        value: 'TODO: var.token'
    vcd:url:
        value: 'TODO: var.vcd_url'
    vcd:user:
        value: none
    vcd:vdc:
        value: 'TODO: var.vcd_vdc'

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/vcd/v3/vcd"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a new network in the default organization and VDC
		_, err := vcd.NewNetworkRouted(ctx, "net1", nil)
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
    vcd:allowUnverifiedSsl:
        value: 'TODO: var.vcd_allow_unverified_ssl'
    vcd:authType:
        value: token
    vcd:maxRetryTimeout:
        value: 'TODO: var.vcd_max_retry_timeout'
    vcd:org:
        value: 'TODO: var.vcd_org'
    vcd:password:
        value: none
    vcd:sysorg:
        value: System
    vcd:token:
        value: 'TODO: var.token'
    vcd:url:
        value: 'TODO: var.vcd_url'
    vcd:user:
        value: none
    vcd:vdc:
        value: 'TODO: var.vcd_vdc'

```
```yaml
resources:
  # Create a new network in the default organization and VDC
  net1:
    type: vcd:NetworkRouted
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    vcd:allowUnverifiedSsl:
        value: 'TODO: var.vcd_allow_unverified_ssl'
    vcd:authType:
        value: token
    vcd:maxRetryTimeout:
        value: 'TODO: var.vcd_max_retry_timeout'
    vcd:org:
        value: 'TODO: var.vcd_org'
    vcd:password:
        value: none
    vcd:sysorg:
        value: System
    vcd:token:
        value: 'TODO: var.token'
    vcd:url:
        value: 'TODO: var.vcd_url'
    vcd:user:
        value: none
    vcd:vdc:
        value: 'TODO: var.vcd_vdc'

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.vcd.NetworkRouted;
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
        // Create a new network in the default organization and VDC
        var net1 = new NetworkRouted("net1");

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

When using a token, the fields `user` and `password` will be ignored, but they need to be in the script.
## Connecting with an API token/API token file

With VCD 10.3.1+, you can connect using an API token, as defined in the [documentation](https://docs.vmware.com/en/VMware-Cloud-Director/10.3/VMware-Cloud-Director-Service-Provider-Admin-Portal-Guide/GUID-A1B3B2FA-7B2C-4EE1-9D1B-188BE703EEDE.html).
The API token is not a bearer token, but one will be created and automatically used by the Pulumi provider when an API
token is supplied. You can create an API token file by utilizing the [`vcd.ApiToken`](https://www.terraform.io/providers/vmware/vcd/latest/docs/resource/api_token) resource.
#### Example usage (API token)

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    vcd:allowUnverifiedSsl:
        value: 'TODO: var.vcd_allow_unverified_ssl'
    vcd:apiToken:
        value: 'TODO: var.api_token'
    vcd:authType:
        value: api_token
    vcd:maxRetryTimeout:
        value: 'TODO: var.vcd_max_retry_timeout'
    vcd:org:
        value: 'TODO: var.vcd_org'
    vcd:password:
        value: none
    vcd:sysorg:
        value: System
    vcd:url:
        value: 'TODO: var.vcd_url'
    vcd:user:
        value: none
    vcd:vdc:
        value: 'TODO: var.vcd_vdc'

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as vcd from "@pulumi/vcd";

// Create a new network in the default organization and VDC
const net1 = new vcd.NetworkRouted("net1", {});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    vcd:allowUnverifiedSsl:
        value: 'TODO: var.vcd_allow_unverified_ssl'
    vcd:apiToken:
        value: 'TODO: var.api_token'
    vcd:authType:
        value: api_token
    vcd:maxRetryTimeout:
        value: 'TODO: var.vcd_max_retry_timeout'
    vcd:org:
        value: 'TODO: var.vcd_org'
    vcd:password:
        value: none
    vcd:sysorg:
        value: System
    vcd:url:
        value: 'TODO: var.vcd_url'
    vcd:user:
        value: none
    vcd:vdc:
        value: 'TODO: var.vcd_vdc'

```
```python
import pulumi
import pulumi_vcd as vcd

# Create a new network in the default organization and VDC
net1 = vcd.NetworkRouted("net1")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    vcd:allowUnverifiedSsl:
        value: 'TODO: var.vcd_allow_unverified_ssl'
    vcd:apiToken:
        value: 'TODO: var.api_token'
    vcd:authType:
        value: api_token
    vcd:maxRetryTimeout:
        value: 'TODO: var.vcd_max_retry_timeout'
    vcd:org:
        value: 'TODO: var.vcd_org'
    vcd:password:
        value: none
    vcd:sysorg:
        value: System
    vcd:url:
        value: 'TODO: var.vcd_url'
    vcd:user:
        value: none
    vcd:vdc:
        value: 'TODO: var.vcd_vdc'

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Vcd = Pulumi.Vcd;

return await Deployment.RunAsync(() =>
{
    // Create a new network in the default organization and VDC
    var net1 = new Vcd.NetworkRouted("net1");

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    vcd:allowUnverifiedSsl:
        value: 'TODO: var.vcd_allow_unverified_ssl'
    vcd:apiToken:
        value: 'TODO: var.api_token'
    vcd:authType:
        value: api_token
    vcd:maxRetryTimeout:
        value: 'TODO: var.vcd_max_retry_timeout'
    vcd:org:
        value: 'TODO: var.vcd_org'
    vcd:password:
        value: none
    vcd:sysorg:
        value: System
    vcd:url:
        value: 'TODO: var.vcd_url'
    vcd:user:
        value: none
    vcd:vdc:
        value: 'TODO: var.vcd_vdc'

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/vcd/v3/vcd"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a new network in the default organization and VDC
		_, err := vcd.NewNetworkRouted(ctx, "net1", nil)
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
    vcd:allowUnverifiedSsl:
        value: 'TODO: var.vcd_allow_unverified_ssl'
    vcd:apiToken:
        value: 'TODO: var.api_token'
    vcd:authType:
        value: api_token
    vcd:maxRetryTimeout:
        value: 'TODO: var.vcd_max_retry_timeout'
    vcd:org:
        value: 'TODO: var.vcd_org'
    vcd:password:
        value: none
    vcd:sysorg:
        value: System
    vcd:url:
        value: 'TODO: var.vcd_url'
    vcd:user:
        value: none
    vcd:vdc:
        value: 'TODO: var.vcd_vdc'

```
```yaml
resources:
  # Create a new network in the default organization and VDC
  net1:
    type: vcd:NetworkRouted
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    vcd:allowUnverifiedSsl:
        value: 'TODO: var.vcd_allow_unverified_ssl'
    vcd:apiToken:
        value: 'TODO: var.api_token'
    vcd:authType:
        value: api_token
    vcd:maxRetryTimeout:
        value: 'TODO: var.vcd_max_retry_timeout'
    vcd:org:
        value: 'TODO: var.vcd_org'
    vcd:password:
        value: none
    vcd:sysorg:
        value: System
    vcd:url:
        value: 'TODO: var.vcd_url'
    vcd:user:
        value: none
    vcd:vdc:
        value: 'TODO: var.vcd_vdc'

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.vcd.NetworkRouted;
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
        // Create a new network in the default organization and VDC
        var net1 = new NetworkRouted("net1");

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
#### Example usage (API token file)

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    vcd:allowUnverifiedSsl:
        value: 'TODO: var.vcd_allow_unverified_ssl'
    vcd:apiTokenFile:
        value: token.json
    vcd:authType:
        value: api_token_file
    vcd:maxRetryTimeout:
        value: 'TODO: var.vcd_max_retry_timeout'
    vcd:org:
        value: 'TODO: var.vcd_org'
    vcd:password:
        value: none
    vcd:sysorg:
        value: System
    vcd:url:
        value: 'TODO: var.vcd_url'
    vcd:user:
        value: none
    vcd:vdc:
        value: 'TODO: var.vcd_vdc'

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as vcd from "@pulumi/vcd";

// Create a new network in the default organization and VDC
const net1 = new vcd.NetworkRouted("net1", {});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    vcd:allowUnverifiedSsl:
        value: 'TODO: var.vcd_allow_unverified_ssl'
    vcd:apiTokenFile:
        value: token.json
    vcd:authType:
        value: api_token_file
    vcd:maxRetryTimeout:
        value: 'TODO: var.vcd_max_retry_timeout'
    vcd:org:
        value: 'TODO: var.vcd_org'
    vcd:password:
        value: none
    vcd:sysorg:
        value: System
    vcd:url:
        value: 'TODO: var.vcd_url'
    vcd:user:
        value: none
    vcd:vdc:
        value: 'TODO: var.vcd_vdc'

```
```python
import pulumi
import pulumi_vcd as vcd

# Create a new network in the default organization and VDC
net1 = vcd.NetworkRouted("net1")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    vcd:allowUnverifiedSsl:
        value: 'TODO: var.vcd_allow_unverified_ssl'
    vcd:apiTokenFile:
        value: token.json
    vcd:authType:
        value: api_token_file
    vcd:maxRetryTimeout:
        value: 'TODO: var.vcd_max_retry_timeout'
    vcd:org:
        value: 'TODO: var.vcd_org'
    vcd:password:
        value: none
    vcd:sysorg:
        value: System
    vcd:url:
        value: 'TODO: var.vcd_url'
    vcd:user:
        value: none
    vcd:vdc:
        value: 'TODO: var.vcd_vdc'

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Vcd = Pulumi.Vcd;

return await Deployment.RunAsync(() =>
{
    // Create a new network in the default organization and VDC
    var net1 = new Vcd.NetworkRouted("net1");

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    vcd:allowUnverifiedSsl:
        value: 'TODO: var.vcd_allow_unverified_ssl'
    vcd:apiTokenFile:
        value: token.json
    vcd:authType:
        value: api_token_file
    vcd:maxRetryTimeout:
        value: 'TODO: var.vcd_max_retry_timeout'
    vcd:org:
        value: 'TODO: var.vcd_org'
    vcd:password:
        value: none
    vcd:sysorg:
        value: System
    vcd:url:
        value: 'TODO: var.vcd_url'
    vcd:user:
        value: none
    vcd:vdc:
        value: 'TODO: var.vcd_vdc'

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/vcd/v3/vcd"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a new network in the default organization and VDC
		_, err := vcd.NewNetworkRouted(ctx, "net1", nil)
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
    vcd:allowUnverifiedSsl:
        value: 'TODO: var.vcd_allow_unverified_ssl'
    vcd:apiTokenFile:
        value: token.json
    vcd:authType:
        value: api_token_file
    vcd:maxRetryTimeout:
        value: 'TODO: var.vcd_max_retry_timeout'
    vcd:org:
        value: 'TODO: var.vcd_org'
    vcd:password:
        value: none
    vcd:sysorg:
        value: System
    vcd:url:
        value: 'TODO: var.vcd_url'
    vcd:user:
        value: none
    vcd:vdc:
        value: 'TODO: var.vcd_vdc'

```
```yaml
resources:
  # Create a new network in the default organization and VDC
  net1:
    type: vcd:NetworkRouted
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    vcd:allowUnverifiedSsl:
        value: 'TODO: var.vcd_allow_unverified_ssl'
    vcd:apiTokenFile:
        value: token.json
    vcd:authType:
        value: api_token_file
    vcd:maxRetryTimeout:
        value: 'TODO: var.vcd_max_retry_timeout'
    vcd:org:
        value: 'TODO: var.vcd_org'
    vcd:password:
        value: none
    vcd:sysorg:
        value: System
    vcd:url:
        value: 'TODO: var.vcd_url'
    vcd:user:
        value: none
    vcd:vdc:
        value: 'TODO: var.vcd_vdc'

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.vcd.NetworkRouted;
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
        // Create a new network in the default organization and VDC
        var net1 = new NetworkRouted("net1");

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

The file containing the API token needs to be readable and writable, in `json` format with the API key. e.g:
```json
{"refresh_token":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"}
```

Note that when connecting with API tokens you can't create or modify users, roles, global roles, or rights bundles.
## Connecting with a Service Account API token

With VCD 10.4.0+, similar to API token file, you can connect using a service account API token, as
defined in the
[documentation](https://blogs.vmware.com/cloudprovider/2022/07/cloud-director-service-accounts.html).
Because a new API token is provided on every authentication request,
the user is required to provide a readable+writable file in `json`
format with the current API key. e.g:
```json
{"refresh_token":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"}
```
Note that the file will be rewritten at every usage, and the updated file will have additional fields, such as
```json
{
  "token_type": "Service Account",
  "refresh_token": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
  "updated_by": "pulumi-provider-vcd/v3.9.0 (darwin/arm64; isProvider:true)",
  "updated_on": "2023-04-18T14:33:07+02:00"
 }
```

The API token file is **sensitive data** and it's up to the user to secure it.

> **NOTE:** The service account needs to be in `Active Stage` and
it's up to the user to provide the initial API token. A service account
can be created using the [`serviceAccount`](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/service_account) resource,
also it can be done using a sample shell script for creating, authorizing
and activating a VCD Service Account can be found in the
[repository](https://github.com/vmware/pulumi-provider-vcd/blob/main/scripts/create_service_account.sh)

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    vcd:allowUnverifiedSsl:
        value: 'TODO: var.vcd_allow_unverified_ssl'
    vcd:authType:
        value: service_account_token_file
    vcd:maxRetryTimeout:
        value: 'TODO: var.vcd_max_retry_timeout'
    vcd:org:
        value: 'TODO: var.vcd_org'
    vcd:serviceAccountTokenFile:
        value: token.json
    vcd:url:
        value: 'TODO: var.vcd_url'
    vcd:vdc:
        value: 'TODO: var.vcd_vdc'

```
## Shell script to obtain a bearer token
To obtain a bearer token you can use this sample shell script:

```sh
#!/bin/bash
user=$1
password=$2
org=$3
IP=$4

if [ -z "$IP" ]
then
    echo "Syntax $0 user password organization hostname_or_IP_address"
    exit 1
fi

options=""
os=$(uname -s)
is_linux=$(echo "$os" | grep -i linux)
if [ -n "$is_linux" ]
then
  options="-w 0"
fi

auth=$(echo -n "$user@$org:$password" |base64 $options)

curl -I -k --header "Accept: application/*;version=32.0" \
    --header "Authorization: Basic $auth" \
    --request POST https://$IP/api/sessions
```

If successful, the output of this command will include lines like the following:

```
X-VCLOUD-AUTHORIZATION: 08a321735de84f1d9ec80c3b3e18fa8b
X-VMWARE-VCLOUD-ACCESS-TOKEN: eyJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJhZG1pbmlzdHJhdG9yIiwiaXNzIjoiYTkzYzlkYjktNzQ3MS0zMTkyLThkMDktYThmN2VlZGE4NWY5QGY5MDZlODE1LTM0NjgtNGQ0ZS04MmJlLTcyYzFjMmVkMTBiMyIsImV4cCI6MTYwNzUxMjgyOCwidmVyc2lvbiI6InZjbG91ZF8xLjAiLCJqdGkiOiJjY2IwZjIwN2JjY2Y0NmYwYmEwNTcyNzgxZDQyNDg2MyJ9.SMjp5wsSd7CXGMdlj-weeCRdr5AazA74pwwx2w3Eqh3RdzyiEMvQfWQAuPAQjM1oOsEUnFOg2u0gYsnIyQg_p7kzXKPQwPNz3BPi0tm2DxxQtQVhOBRXCqUJ9OmRlMVu7FZZ6gKD4GhpbTkZyKMN_IgOFkkt8iXs1-weNZw5TmyVHeWiJdV0JFM45CV47jQNdQMy4OSsU-CqE2VVLOK83oJhRnlnc3O4OAAIfuVZ4SLWqgi1lIoc2vbZv0HYeWO7L_2pGfmja8CVzVhPrgIGEoDhXnvO29z1ToEXRnyMKh9cisiRkhUISpsh4aHRGUUzaZYeOejVX3PAO9aCX3iYWA

The string after `X-VCLOUD-AUTHORIZATION:` is the old (deprecated) token.
The string after `X-VMWARE-VCLOUD-ACCESS-TOKEN` is the bearer token
```

Either token will grant the same abilities as the account used to run the above script. Note, however, that the deprecated
token may not work in recent VCD versions.

Using a token produced by an org admin to run a task that requires a system administrator will fail.
## Connecting with SAML user using Microsoft Active Directory Federation Services (ADFS) and setting custom Relaying Party Trust Identifier

Take special attention to `user`, `useSamlAdfs` and `samlRptId` fields.

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    vcd:allowUnverifiedSsl:
        value: 'TODO: var.vcd_allow_unverified_ssl'
    vcd:authType:
        value: saml_adfs
    vcd:maxRetryTimeout:
        value: 'TODO: var.vcd_max_retry_timeout'
    vcd:org:
        value: 'TODO: var.vcd_org'
    vcd:password:
        value: 'TODO: var.vcd_pass'
    vcd:samlAdfsCookie:
        value: sso-preferred=yes; sso_redirect_org={{.Org}}
    vcd:samlAdfsRptId:
        value: my-custom-rpt-id
    vcd:sysorg:
        value: my-org
    vcd:url:
        value: 'TODO: var.vcd_url'
    vcd:user:
        value: test@contoso.com
    vcd:vdc:
        value: 'TODO: var.vcd_vdc'

```
## Configuration Reference

The following arguments are used to configure the VMware Cloud Director Provider:

* `user` - (Required) This is the username for Cloud Director API operations. Can also be specified
  with the `VCD_USER` environment variable. *v2.0+* `user` may be "administrator" (set `org` or
  `sysorg` to "System" in this case).
  *v2.9+* When using with SAML and ADFS - username format must be in Active Directory format -
  `user@contoso.com` or `contoso.com\user` in combination with `useSamlAdfs` option.

* `password` - (Required) This is the password for Cloud Director API operations. Can
  also be specified with the `VCD_PASSWORD` environment variable.

* `authType` - (Optional) `integrated`, `token`, `apiToken`, `serviceAccountTokenFile` or `samlAdfs`.
  Default is `integrated`. Can also be set with `VCD_AUTH_TYPE` environment variable.
  * `integrated` - VCD local users and LDAP users (provided LDAP is configured for Organization).
  * `samlAdfs` allows to use SAML login flow with Active Directory Federation
    Services (ADFS) using "/adfs/services/trust/13/usernamemixed" endpoint. Please note that
    credentials for ADFS should be formatted as `user@contoso.com` or `contoso.com\user`.
    `samlAdfsRptId` can be used to specify a different RPT ID.
  * `token` allows to specify token in `token` field.
  * `apiToken` allows to specify an API token.
  * `apiTokenFile` allows to specify a file containing an API token.
  * `serviceAccountTokenFile` allows to specify a file containing a service account's token.

* `token` - (Optional; *v2.6+*) This is the bearer token that can be used instead of username
  and password (in combination with field `auth_type=token`). When this is set, username and
  password will be ignored, but should be left in configuration either empty or with any custom
  values. A token can be specified with the `VCD_TOKEN` environment variable.
  Both a (deprecated) authorization token or a bearer token (*v3.1+*) can be used in this field.

* `apiToken` - (Optional; *v3.5+*) This is the API token that a System or organization administrator can create and
  distribute to users. It is used instead of username and password (in combination with `auth_type=api_token`). When
  this field is filled, username and password are ignored. An API token can also be specified with the `VCD_API_TOKEN`
  environment variable. This token requires at least VCD 10.3.1. There are restrictions to its use, as defined in
  [the documentation](https://docs.vmware.com/en/VMware-Cloud-Director/10.3/VMware-Cloud-Director-Service-Provider-Admin-Portal-Guide/GUID-A1B3B2FA-7B2C-4EE1-9D1B-188BE703EEDE.html)

* `apiTokenFile` - (Optional; *v3.10+*)) Same as `apiToken`, only provided
  as a JSON file. Can also be specified with the `VCD_API_TOKEN_FILE` environment variable.

* `serviceAccountTokenFile` - (Optional; *v3.9+, VCD 10.4+*) This is the file that contains a Service Account API token. The
  path to the file could be provided as absolute or relative to the working directory. It is used instead of username
  and password (in combination with `auth_type=service_account_token_file`. The file can also be specified with the
  `VCD_SA_TOKEN_FILE` environment variable. There are restrictions to its use, as defined in
  [the documentation](https://docs.vmware.com/en/VMware-Cloud-Director/10.4/VMware-Cloud-Director-Service-Provider-Admin-Portal-Guide/GUID-8CD3C8BE-3187-4769-B960-3E3315492C16.html)

* `allowServiceAccountTokenFile` - (Optional; *v3.9+, VCD 10.4+*) When using `auth_type=service_account_token_file`,
  if set to `true`, will suppress a warning to the user about the service account token file containing *sensitive information*.
  Can also be set with `VCD_ALLOW_SA_TOKEN_FILE`.

* `samlAdfsRptId` - (Optional) When using `auth_type=saml_adfs` VCD SAML entity ID will be used
  as Relaying Party Trust Identifier (RPT ID) by default. If a different RPT ID is needed - one can
  set it using this field. It can also be set with `VCD_SAML_ADFS_RPT_ID` environment variable.

* `samlAdfsCookie` - (Optional; *v3.14+*) An additional cookie that can be injected when looking
  up ADFS server from VCD. Example `sso-preferred=yes; sso_redirect_org={{.Org}}`. `{{.Org}}` will be
  replaced with actual Org during runtime.

* `org` - (Required) This is the Cloud Director Org on which to run API
  operations. Can also be specified with the `VCD_ORG` environment
  variable.*v2.0+* `org` may be set to "System" when connection as Sys Admin is desired
  (set `user` to "administrator" in this case).Note: `org` value is case sensitive.

* `sysorg` - (Optional; *v2.0+*) - Organization for user authentication. Can also be
  specified with the `VCD_SYS_ORG` environment variable. Set `sysorg` to "System" and
  `user` to "administrator" to free up `org` argument for setting a default organization
  for resources to use.

* `url` - (Required) This is the URL for the Cloud Director API endpoint. e.g.
  <https://server.domain.com/api>. Can also be specified with the `VCD_URL` environment variable.

* `vdc` - (Optional) This is the virtual datacenter within Cloud Director to run
  API operations against. If not set the plugin will select the first virtual
  datacenter available to your Org. Can also be specified with the `VCD_VDC` environment
  variable.

* `maxRetryTimeout` - (Optional) This provides you with the ability to specify the maximum
  amount of time (in seconds) you are prepared to wait for interactions on resources managed
  by Cloud Director to be successful. If a resource action fails, the action will be retried
  (as long as it is still within the `maxRetryTimeout` value) to try and ensure success.
  Defaults to 60 seconds if not set.
  Can also be specified with the `VCD_MAX_RETRY_TIMEOUT` environment variable.

* `maxRetryTimeout` - (Deprecated) Use `maxRetryTimeout` instead.

* `allowUnverifiedSsl` - (Optional) Boolean that can be set to true to
  disable SSL certificate verification. This should be used with care as it
  could allow an attacker to intercept your auth token. If omitted, default
  value is false. Can also be specified with the
  `VCD_ALLOW_UNVERIFIED_SSL` environment variable.

* `logging` - (Optional; *v2.0+*) Boolean that enables API calls logging from upstream library `go-vcloud-director`.
  The logging file will record all API requests and responses, plus some debug information that is part of this
  provider. Logging can also be activated using the `VCD_API_LOGGING` environment variable.

* `loggingFile` - (Optional; *v2.0+*) The name of the log file (when `logging` is enabled). By default is
  `go-vcloud-director` and it can also be changed using the `VCD_API_LOGGING_FILE` environment variable.

* `importSeparator` - (Optional; *v2.5+*) The string to be used as separator with `pulumi import`. By default
  it is a dot (`.`).

* `ignoreMetadataChanges` - (Optional; Experimental; *v3.10+*) Use one or more of these blocks to ignore specific metadata entries from being changed by this Pulumi provider
  after creation or when they were created outside Pulumi.
  See "Ignore Metadata Changes" for more details.
## Ignore metadata changes

=> This is an **EXPERIMENTAL FEATURE** that may change in a future release.

One or more `ignoreMetadataChanges` blocks can be optionally set in the provider configuration, which will allow to ignore specific `metadataEntry`
items during all Pulumi operations. This is useful, for example, to avoid removing metadata entries that were created
by an external actor, or after they were created by Pulumi.

> Note that this feature is only considered when using the `metadataEntry` argument in the resources and functions that support
it. In other words, to ignore metadata when the deprecated `metadata` argument is used, please use the native Pulumi `lifecycle.ignore_changes` block.

> Be aware that setting a `metadataEntry` in the Pulumi configuration that matches any `ignoreMetadataChanges` can produce inconsistent
results, as the metadata will be stored in state but nothing will be done in VCD. Using `ignoreMetadataChanges` with matching metadata entries
in the code is NOT recommended. In the event that it contains such conflict, though, the ensuing action can be controlled with
`conflictAction`, which can be `error`, `warn` or `none`.

The available sub-attributes for `ignoreMetadataChanges` are:

* `resourceType` - (Optional) Specifies the resource type which metadata needs to be ignored. If set, the resource type must be one of:
  *"vcd.Catalog"*, *"vcd.CatalogItem"*, *"vcd.CatalogMedia"*, *"vcd.CatalogVappTemplate"*, *"vcd.IndependentDisk"*, *"vcd.NetworkDirect"*,
  *"vcd.NetworkIsolated"*, *"vcd.NetworkIsolatedV2"*, *"vcd.NetworkRouted"*, *"vcd.NetworkRoutedV2"*, *"vcd.Org"*, *"vcd.OrgVdc"*, *"vcd.ProviderVdc"*,
  *"vcd.Rde" (v3.11+)*, *"vcd.getStorageProfile"*, *"vcd.Vapp"*, *"vcd.VappVm"* or *"vcd.Vm"*, which are the resources compatible with `metadataEntry`.
* `resourceName`- (Optional) Specifies the name of the entity in VCD which metadata needs to be ignored. This attribute can be used with
  any kind of `resourceType`, except for *vcd_storage_profile* which **cannot be filtered by name**.
* `keyRegex`- (Optional) A regular expression that can filter out metadata keys that match. Either `keyRegex` or `valueRegex` are required on each block.
* `valueRegex`- (Optional) A regular expression that can filter out metadata values that match. Either `keyRegex` or `valueRegex` are required on each block.
* `conflictAction` - (Optional) Defines what to do if a conflict exists between a `metadataEntry` that is managed
  by Pulumi, and it matches the criteria defined in the `ignoreMetadataChanges` block, as the metadata will be stored in state but nothing will be done in VCD.
  If the value is `error`, when this happens, any read operation (like a Plan or Refresh) will fail. When the value is `warn`, it will just give a warning but the operation will continue,
  and with the `none` value nothing will be shown. Defaults to `error`.

> The `conflictAction` mechanism will be evaluated on every read, including `pulumi destroy`, as it will trigger a refresh before deleting
resources. To avoid this situation, we can use the `-refresh=false` option.

Note that these attributes **are evaluated as a logical `and`**. This means that the snippet below would ignore all metadata entries
that belong to the specific Organization named "client1" **and** which keys match the regular expression `[Ee]nvironment`:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:

```

We can have more than one block, to ignore more entries:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as vcd from "@pulumi/vcd";

const myOrg = new vcd.Org("my_org", {
    name: "MyOrg",
    metadataEntries: [{
        key: "OneKey",
        value: "OneValue",
        type: "MetadataStringValue",
        userAccess: "READWRITE",
        isSystem: false,
    }],
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
import pulumi_vcd as vcd

my_org = vcd.Org("my_org",
    name="MyOrg",
    metadata_entries=[{
        "key": "OneKey",
        "value": "OneValue",
        "type": "MetadataStringValue",
        "user_access": "READWRITE",
        "is_system": False,
    }])
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
using Vcd = Pulumi.Vcd;

return await Deployment.RunAsync(() =>
{
    var myOrg = new Vcd.Org("my_org", new()
    {
        Name = "MyOrg",
        MetadataEntries = new[]
        {
            new Vcd.Inputs.OrgMetadataEntryArgs
            {
                Key = "OneKey",
                Value = "OneValue",
                Type = "MetadataStringValue",
                UserAccess = "READWRITE",
                IsSystem = false,
            },
        },
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
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/vcd/v3/vcd"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := vcd.NewOrg(ctx, "my_org", &vcd.OrgArgs{
			Name: pulumi.String("MyOrg"),
			MetadataEntries: vcd.OrgMetadataEntryArray{
				&vcd.OrgMetadataEntryArgs{
					Key:        pulumi.String("OneKey"),
					Value:      pulumi.String("OneValue"),
					Type:       pulumi.String("MetadataStringValue"),
					UserAccess: pulumi.String("READWRITE"),
					IsSystem:   pulumi.Bool(false),
				},
			},
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
  myOrg:
    type: vcd:Org
    name: my_org
    properties:
      name: MyOrg
      metadataEntries:
        - key: OneKey
          value: OneValue
          type: MetadataStringValue
          userAccess: READWRITE
          isSystem: false
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
import com.pulumi.vcd.Org;
import com.pulumi.vcd.OrgArgs;
import com.pulumi.vcd.inputs.OrgMetadataEntryArgs;
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
        var myOrg = new Org("myOrg", OrgArgs.builder()
            .name("MyOrg")
            .metadataEntries(OrgMetadataEntryArgs.builder()
                .key("OneKey")
                .value("OneValue")
                .type("MetadataStringValue")
                .userAccess("READWRITE")
                .isSystem(false)
                .build())
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

Note that this argument **does not affect metadata of the [function filters](https://www.terraform.io/providers/vmware/vcd/latest/docs/guides/data_source_filters)**.
## Connection Cache (*2.0+*)

Cloud Director connection calls can be expensive, and if a definition file contains several resources, it may trigger
multiple connections. There is a cache engine, disabled by default, which can be activated by the `VCD_CACHE`
environment variable. When enabled, the provider will not reconnect, but reuse an active connection for up to 20
minutes, and then connect again.