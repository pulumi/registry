---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/jdamata/sonarqube/0.16.15/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Sonarqube Provider
meta_desc: Provides an overview on how to configure the Pulumi Sonarqube provider.
layout: package
---

## Generate Provider

The Sonarqube provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider jdamata/sonarqube
```
## Overview

The sonarqube provider is used to configure sonarqube. The provider needs to be configured with a url, and either with user and password or token.
## Example: Authenticate with username and password

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    sonarqube:host:
        value: http://127.0.0.1:9000
    sonarqube:pass:
        value: admin
    sonarqube:user:
        value: admin

```
```typescript
import * as pulumi from "@pulumi/pulumi";

```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    sonarqube:host:
        value: http://127.0.0.1:9000
    sonarqube:pass:
        value: admin
    sonarqube:user:
        value: admin

```
```python
import pulumi

```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    sonarqube:host:
        value: http://127.0.0.1:9000
    sonarqube:pass:
        value: admin
    sonarqube:user:
        value: admin

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;

return await Deployment.RunAsync(() =>
{
});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    sonarqube:host:
        value: http://127.0.0.1:9000
    sonarqube:pass:
        value: admin
    sonarqube:user:
        value: admin

```
```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
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
    sonarqube:host:
        value: http://127.0.0.1:9000
    sonarqube:pass:
        value: admin
    sonarqube:user:
        value: admin

```
```yaml
{}
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    sonarqube:host:
        value: http://127.0.0.1:9000
    sonarqube:pass:
        value: admin
    sonarqube:user:
        value: admin

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
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
    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Example: Authenticate with token

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    sonarqube:host:
        value: http://127.0.0.1:9000
    sonarqube:token:
        value: d4at55a6f7r199bd707h39625685510880gbf7ff

```
```typescript
import * as pulumi from "@pulumi/pulumi";

```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    sonarqube:host:
        value: http://127.0.0.1:9000
    sonarqube:token:
        value: d4at55a6f7r199bd707h39625685510880gbf7ff

```
```python
import pulumi

```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    sonarqube:host:
        value: http://127.0.0.1:9000
    sonarqube:token:
        value: d4at55a6f7r199bd707h39625685510880gbf7ff

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;

return await Deployment.RunAsync(() =>
{
});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    sonarqube:host:
        value: http://127.0.0.1:9000
    sonarqube:token:
        value: d4at55a6f7r199bd707h39625685510880gbf7ff

```
```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
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
    sonarqube:host:
        value: http://127.0.0.1:9000
    sonarqube:token:
        value: d4at55a6f7r199bd707h39625685510880gbf7ff

```
```yaml
{}
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    sonarqube:host:
        value: http://127.0.0.1:9000
    sonarqube:token:
        value: d4at55a6f7r199bd707h39625685510880gbf7ff

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
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
    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

The following configuration inputs are supported:

- `user` - (Optional) Sonarqube user. This can also be set via the `SONARQUBE_USER` environment variable.
- `pass` - (Optional) Sonarqube pass. This can also be set via the `SONARQUBE_PASS` environment variable.
- `token` - (Optional) Sonarqube token. This can also be set via the `SONARQUBE_TOKEN` environment variable.
- `host` - (Required) Sonarqube url. This can be also be set via the `SONARQUBE_HOST` environment variable.
- `installedVersion` - (Optional) The version of the Sonarqube server. When specified, the provider will avoid requesting this from the
  server during the initialization process. This can be helpful when using the same Pulumi code to install Sonarqube and configure it.
- `tlsInsecureSkipVerify` - (Optional) Allows ignoring insecure certificates when set to true. Defaults to false. Disabling TLS verification
  is dangerous and should only be done for local testing.
- `anonymizeUserOnDelete` - (Optional) Allows anonymizing users on destroy. Requires Sonarqube version >= `9.7`. This can be helpful
  to comply with regulations like [GDPR](https://en.wikipedia.org/wiki/General_Data_Protection_Regulation).