---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-rancher2/v9.0.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Rancher2 Provider
meta_desc: Provides an overview on how to configure the Pulumi Rancher2 provider.
layout: package
---
## Installation

The Rancher2 provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/rancher2`](https://www.npmjs.com/package/@pulumi/rancher2)
* Python: [`pulumi-rancher2`](https://pypi.org/project/pulumi-rancher2/)
* Go: [`github.com/pulumi/pulumi-rancher2/sdk/v9/go/rancher2`](https://github.com/pulumi/pulumi-rancher2)
* .NET: [`Pulumi.Rancher2`](https://www.nuget.org/packages/Pulumi.Rancher2)
* Java: [`com.pulumi/rancher2`](https://central.sonatype.com/artifact/com.pulumi/rancher2)
## Overview

The Rancher2 provider is used to interact with the
resources supported by Rancher v2.

The provider can be configured in 2 modes:
- Admin: this is the default mode, intended to manage rancher2 resources. It should be configured with the `apiUrl` of the Rancher server and API credentials, `tokenKey` or `accessKey` and `secretKey`.
- Bootstrap: this mode is intended to bootstrap a rancher2 system. It is enabled if `bootstrap = true`. In this mode, `tokenKey` or `accessKey` and `secretKey` can not be provided. More info at rancher2.Bootstrap resource
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    rancher2:accessKey:
        value: 'TODO: var.rancher2_access_key'
    rancher2:apiUrl:
        value: https://rancher.my-domain.com
    rancher2:secretKey:
        value: 'TODO: var.rancher2_secret_key'

```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    rancher2:accessKey:
        value: 'TODO: var.rancher2_access_key'
    rancher2:apiUrl:
        value: https://rancher.my-domain.com
    rancher2:secretKey:
        value: 'TODO: var.rancher2_secret_key'

```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    rancher2:accessKey:
        value: 'TODO: var.rancher2_access_key'
    rancher2:apiUrl:
        value: https://rancher.my-domain.com
    rancher2:secretKey:
        value: 'TODO: var.rancher2_secret_key'

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    rancher2:accessKey:
        value: 'TODO: var.rancher2_access_key'
    rancher2:apiUrl:
        value: https://rancher.my-domain.com
    rancher2:secretKey:
        value: 'TODO: var.rancher2_secret_key'

```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    rancher2:accessKey:
        value: 'TODO: var.rancher2_access_key'
    rancher2:apiUrl:
        value: https://rancher.my-domain.com
    rancher2:secretKey:
        value: 'TODO: var.rancher2_secret_key'

```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    rancher2:accessKey:
        value: 'TODO: var.rancher2_access_key'
    rancher2:apiUrl:
        value: https://rancher.my-domain.com
    rancher2:secretKey:
        value: 'TODO: var.rancher2_secret_key'

```

{{% /choosable %}}
{{< /chooser >}}

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    rancher2:apiUrl:
        value: https://rancher.my-domain.com
    rancher2:bootstrap:
        value: true

```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    rancher2:apiUrl:
        value: https://rancher.my-domain.com
    rancher2:bootstrap:
        value: true

```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    rancher2:apiUrl:
        value: https://rancher.my-domain.com
    rancher2:bootstrap:
        value: true

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    rancher2:apiUrl:
        value: https://rancher.my-domain.com
    rancher2:bootstrap:
        value: true

```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    rancher2:apiUrl:
        value: https://rancher.my-domain.com
    rancher2:bootstrap:
        value: true

```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    rancher2:apiUrl:
        value: https://rancher.my-domain.com
    rancher2:bootstrap:
        value: true

```

{{% /choosable %}}
{{< /chooser >}}

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as rancher2 from "@pulumi/rancher2";

// Create a new rancher2_bootstrap using bootstrap provider config
const admin = new rancher2.Bootstrap("admin", {
    password: "blahblah",
    telemetry: true,
});
// Create a new rancher2 resource using admin provider config
const foo = new rancher2.Catalog("foo", {
    name: "test",
    url: "http://foo.com:8080",
});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_rancher2 as rancher2

# Create a new rancher2_bootstrap using bootstrap provider config
admin = rancher2.Bootstrap("admin",
    password="blahblah",
    telemetry=True)
# Create a new rancher2 resource using admin provider config
foo = rancher2.Catalog("foo",
    name="test",
    url="http://foo.com:8080")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Rancher2 = Pulumi.Rancher2;

return await Deployment.RunAsync(() =>
{
    // Create a new rancher2_bootstrap using bootstrap provider config
    var admin = new Rancher2.Bootstrap("admin", new()
    {
        Password = "blahblah",
        Telemetry = true,
    });

    // Create a new rancher2 resource using admin provider config
    var foo = new Rancher2.Catalog("foo", new()
    {
        Name = "test",
        Url = "http://foo.com:8080",
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-rancher2/sdk/v8/go/rancher2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a new rancher2_bootstrap using bootstrap provider config
		_, err := rancher2.NewBootstrap(ctx, "admin", &rancher2.BootstrapArgs{
			Password:  pulumi.String("blahblah"),
			Telemetry: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		// Create a new rancher2 resource using admin provider config
		_, err = rancher2.NewCatalog(ctx, "foo", &rancher2.CatalogArgs{
			Name: pulumi.String("test"),
			Url:  pulumi.String("http://foo.com:8080"),
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
resources:
  # Create a new rancher2_bootstrap using bootstrap provider config
  admin:
    type: rancher2:Bootstrap
    properties:
      password: blahblah
      telemetry: true
  # Create a new rancher2 resource using admin provider config
  foo:
    type: rancher2:Catalog
    properties:
      name: test
      url: http://foo.com:8080
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.rancher2.Bootstrap;
import com.pulumi.rancher2.BootstrapArgs;
import com.pulumi.rancher2.Catalog;
import com.pulumi.rancher2.CatalogArgs;
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
        // Create a new rancher2_bootstrap using bootstrap provider config
        var admin = new Bootstrap("admin", BootstrapArgs.builder()
            .password("blahblah")
            .telemetry(true)
            .build());

        // Create a new rancher2 resource using admin provider config
        var foo = new Catalog("foo", CatalogArgs.builder()
            .name("test")
            .url("http://foo.com:8080")
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

The following configuration inputs are supported:

* `apiUrl` - (Required) Rancher API url. It must be provided, but it can also be sourced from the `RANCHER_URL` environment variable.
* `accessKey` - (Optional/Sensitive) Rancher API access key to connect to rancher. It can also be sourced from the `RANCHER_ACCESS_KEY` environment variable.
* `secretKey` - (Optional/Sensitive) Rancher API secret key to connect to rancher. It can also be sourced from the `RANCHER_SECRET_KEY` environment variable.
* `tokenKey` - (Optional/Sensitive) Rancher API token key to connect to rancher. It can also be sourced from the `RANCHER_TOKEN_KEY` environment variable. Could be used instead `accessKey` and `secretKey`.
* `caCerts` - CA certificates used to sign Rancher server tls certificates. Mandatory if self signed tls and insecure option false. It can also be sourced from the `RANCHER_CA_CERTS` environment variable.
* `insecure` - (Optional) Allow insecure connection to Rancher. Mandatory if self signed tls and not caCerts provided. It can also be sourced from the `RANCHER_INSECURE` environment variable.
* `bootstrap` - (Optional) Enable bootstrap mode to manage `rancher2.Bootstrap` resource. It can also be sourced from the `RANCHER_BOOTSTRAP` environment variable. Default: `false`
* `retries` - (Deprecated) Use timeout instead
* `timeout` - (Optional) Timeout duration to retry for Rancher connectivity and resource operations. Default: `"120s"`