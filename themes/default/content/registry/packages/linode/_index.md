---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-linode/v5.2.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Linode Provider
meta_desc: Provides an overview on how to configure the Pulumi Linode provider.
layout: package
---

## Installation

The Linode provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/linode`](https://www.npmjs.com/package/@pulumi/linode)
* Python: [`pulumi-linode`](https://pypi.org/project/pulumi-linode/)
* Go: [`github.com/pulumi/pulumi-linode/sdk/v5/go/linode`](https://github.com/pulumi/pulumi-linode)
* .NET: [`Pulumi.Linode`](https://www.nuget.org/packages/Pulumi.Linode)
* Java: [`com.pulumi/linode`](https://central.sonatype.com/artifact/com.pulumi/linode)

## Overview

The Linode provider exposes resources and functions to interact with [Linode](https://www.linode.com/) services.
The provider needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available functions.
## Example Usage



{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as linode from "@pulumi/linode";

// Create a Linode
const foobar = new linode.Instance("foobar", {});
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
import pulumi_linode as linode

# Create a Linode
foobar = linode.Instance("foobar")
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
using Linode = Pulumi.Linode;

return await Deployment.RunAsync(() =>
{
    // Create a Linode
    var foobar = new Linode.Instance("foobar");

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
	"github.com/pulumi/pulumi-linode/sdk/v5/go/linode"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a Linode
		_, err := linode.NewInstance(ctx, "foobar", nil)
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
  # Create a Linode
  foobar:
    type: linode:Instance
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
import com.pulumi.linode.Instance;
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
        // Create a Linode
        var foobar = new Instance("foobar");

    }
}
```
{{% /choosable %}}
{{< /chooser >}}



{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as linode from "@pulumi/linode";

// Create a Linode
const foobar = new linode.Instance("foobar", {});
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
import pulumi_linode as linode

# Create a Linode
foobar = linode.Instance("foobar")
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
using Linode = Pulumi.Linode;

return await Deployment.RunAsync(() =>
{
    // Create a Linode
    var foobar = new Linode.Instance("foobar");

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
	"github.com/pulumi/pulumi-linode/sdk/v5/go/linode"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a Linode
		_, err := linode.NewInstance(ctx, "foobar", nil)
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
  # Create a Linode
  foobar:
    type: linode:Instance
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
import com.pulumi.linode.Instance;
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
        // Create a Linode
        var foobar = new Instance("foobar");

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

The following keys can be used to configure the provider.
### Basic Configuration

This section outlines commonly used provider configuration options.

* `configPath` - (Optional) The path to the Linode config file to use. (default `~/.config/linode`)

* `configProfile` - (Optional) The Linode config profile to use. (default `default`)

* `token` - (Optional) This is your [Linode APIv4 Token](https://techdocs.akamai.com/linode-api/reference/get-started#personal-access-tokens).

  The Linode Token can also be specified using the `LINODE_TOKEN` shell environment variable. (e.g. `export LINODE_TOKEN=mytoken`)

  Specifying a token through the `token` field or through the `LINODE_TOKEN` shell environment variable will override the token loaded through a config.

  Configs are not required if a `token` is defined.

* `url` - (Optional) The HTTP(S) API address of the Linode API to use.

  The Linode API URL can also be specified using the `LINODE_URL` environment variable.

  Overrides the Linode Config `apiUrl` field.

* `apiVersion` (Optional) The version of the Linode API to use. (default `v4`)

  The Linode API version can also be specified using the `LINODE_API_VERSION` environment variable.

* `apiCaPath` (Optional) The path to a CA file to trust when making API requests.

  The Linode API CA file path can also be specified using the `LINODE_CA` environment variable.

* `objAccessKey` - (Optional) The access key to be used in linode.ObjectStorageBucket and linode_object_storage_object.

  The Object Access Key can also be specified using the `LINODE_OBJ_ACCESS_KEY` shell environment variable.

* `objSecretKey` - (Optional) The secret key to be used in linode.ObjectStorageBucket and linode_object_storage_object.

  The Object Secret Key can also be specified using the `LINODE_OBJ_SECRET_KEY` shell environment variable.

* `objUseTempKeys` - (Optional) If true, temporary object keys will be created implicitly at apply-time for the linode.ObjectStorageBucket and linode.ObjectStorageObject resource to use.

* `objBucketForceDelete` - (Optional) If true, all objects and versions will purged from a linode.ObjectStorageBucket before it is destroyed.

* `skipInstanceReadyPoll` - (Optional) Skip waiting for a linode.Instance resource to be running.

* `skipInstanceDeletePoll` - (Optional) Skip waiting for a linode.Instance resource to finish deleting.

* `skipImplicitReboots` - (Optional) If true, Linode Instances will not be rebooted on config and interface changes. (default `false`)
### Advanced Configuration

This section outlines less frequently used provider configuration options.

* `uaPrefix` - (Optional) An HTTP User-Agent Prefix to prepend in API requests.

  The User-Agent Prefix can also be specified using the `LINODE_UA_PREFIX` environment variable.

* `minRetryDelayMs` - (Optional) Minimum delay in milliseconds before retrying a request. (default `100`)

* `maxRetryDelayMs` - (Optional) Maximum delay in milliseconds before retrying a request. (default `2000`)

* `eventPollMs` - (Optional) The rate in milliseconds to poll for Linode events. (default `4000`)

  The event polling rate can also be configured using the `LINODE_EVENT_POLL_MS` environment variable.

* `lkeEventPollMs` - (Optional) The rate in milliseconds to poll for LKE events. (default `3000`)

* `lkeNodeReadyPollMs` - (Optional) The rate in milliseconds to poll for an LKE node to be ready. (default `3000`)

* `disableInternalCache` - (Optional) If true, the internal caching system that backs certain Linode API requests will be disabled. (default `false`)
## Early Access

Some resources are made available before the feature reaches general availability. These resources are subject to change, and may not be available to all customers in all regions. Early access features can be accessed by configuring the provider to use a different version of the API.
### Configuring the Target API Version

The `apiVersion` can be set on the provider configuration like so:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    linode:apiVersion:
        value: v4beta

```

Additionally, the version can be set with the `LINODE_API_VERSION` environment variable.
## Rate Limiting

The Linode API may apply rate limiting when you update the state for a large inventory:

```
Error: Error getting Linode DomainRecord ID 123456: [002] unexpected end of JSON input



Error: Error finding the specified Linode DomainRecord: [002] unexpected end of JSON input
```

If this affects you, run Pulumi with --parallelism=1
## Using Configuration Files

Configuration files can be used to specify Linode client configuration options across various Linode integrations.

For example:

`~/.config/linode`

```ini
[default]
token = mylinodetoken
```

`Pulumi.yaml`

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:

```

Specifying the `token` provider options or defining `LINODE_TOKEN` in the environment will override any tokens loaded from a configuration file.

Profiles can also be defined for multitenant use-cases. Every profile will inherit fields from the `default` profile.

For example:

`~/.config/linode`

```ini
[default]
token = alinodetoken

[foo]
token = anotherlinodetoken

[bar]
token = yetanotherlinodetoken
```

`Pulumi.yaml`

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    linode:configProfile:
        value: bar

```

Configuration Profiles also expose additional client configuration fields such as `apiUrl` and `apiVersion`.

For example:

`~/.config/linode`

```ini
[default]
token = mylinodetoken

[stable]
api_version = v4

[beta]
api_version = v4beta

[alpha]
api_version = v4beta
api_url = https://my.alpha.endpoint.com
```

`Pulumi.yaml`

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    linode:configProfile:
        value: beta

```