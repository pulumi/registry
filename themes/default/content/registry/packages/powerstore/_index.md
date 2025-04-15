---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/dell/powerstore/1.2.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Powerstore Provider
meta_desc: Provides an overview on how to configure the Pulumi Powerstore provider.
layout: package
---

## Generate Provider

The Powerstore provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider dell/powerstore
```
## Overview

Provider for PowerStore
## Example Usage

provider.tf
{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    powerstore:endpoint:
        value: 'TODO: var.endpoint'
    powerstore:insecure:
        value: true
    powerstore:password:
        value: 'TODO: var.password'
    powerstore:timeout:
        value: 'TODO: var.timeout'
    powerstore:username:
        value: 'TODO: var.username'

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
    powerstore:endpoint:
        value: 'TODO: var.endpoint'
    powerstore:insecure:
        value: true
    powerstore:password:
        value: 'TODO: var.password'
    powerstore:timeout:
        value: 'TODO: var.timeout'
    powerstore:username:
        value: 'TODO: var.username'

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
    powerstore:endpoint:
        value: 'TODO: var.endpoint'
    powerstore:insecure:
        value: true
    powerstore:password:
        value: 'TODO: var.password'
    powerstore:timeout:
        value: 'TODO: var.timeout'
    powerstore:username:
        value: 'TODO: var.username'

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
    powerstore:endpoint:
        value: 'TODO: var.endpoint'
    powerstore:insecure:
        value: true
    powerstore:password:
        value: 'TODO: var.password'
    powerstore:timeout:
        value: 'TODO: var.timeout'
    powerstore:username:
        value: 'TODO: var.username'

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
    powerstore:endpoint:
        value: 'TODO: var.endpoint'
    powerstore:insecure:
        value: true
    powerstore:password:
        value: 'TODO: var.password'
    powerstore:timeout:
        value: 'TODO: var.timeout'
    powerstore:username:
        value: 'TODO: var.username'

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
    powerstore:endpoint:
        value: 'TODO: var.endpoint'
    powerstore:insecure:
        value: true
    powerstore:password:
        value: 'TODO: var.password'
    powerstore:timeout:
        value: 'TODO: var.timeout'
    powerstore:username:
        value: 'TODO: var.username'

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

variables.tf
{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";

const config = new pulumi.Config();
// Stores the username of PowerStore host.
const username = config.require("username");
// Stores the password of PowerStore host.
const password = config.require("password");
// Stores the timeout of PowerStore host.
const timeout = config.require("timeout");
// Stores the endpoint of PowerStore host. eg: https://10.1.1.1/api/rest
const endpoint = config.require("endpoint");
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi

config = pulumi.Config()
# Stores the username of PowerStore host.
username = config.require("username")
# Stores the password of PowerStore host.
password = config.require("password")
# Stores the timeout of PowerStore host.
timeout = config.require("timeout")
# Stores the endpoint of PowerStore host. eg: https://10.1.1.1/api/rest
endpoint = config.require("endpoint")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;

return await Deployment.RunAsync(() =>
{
    var config = new Config();
    // Stores the username of PowerStore host.
    var username = config.Require("username");
    // Stores the password of PowerStore host.
    var password = config.Require("password");
    // Stores the timeout of PowerStore host.
    var timeout = config.Require("timeout");
    // Stores the endpoint of PowerStore host. eg: https://10.1.1.1/api/rest
    var endpoint = config.Require("endpoint");
});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		// Stores the username of PowerStore host.
		username := cfg.Require("username")
		// Stores the password of PowerStore host.
		password := cfg.Require("password")
		// Stores the timeout of PowerStore host.
		timeout := cfg.Require("timeout")
		// Stores the endpoint of PowerStore host. eg: https://10.1.1.1/api/rest
		endpoint := cfg.Require("endpoint")
		return nil
	})
}
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
configuration:
  # /*
  # Copyright (c) 2024 Dell Inc., or its subsidiaries. All Rights Reserved.

  # Licensed under the Mozilla Public License Version 2.0 (the "License");
  # you may not use this file except in compliance with the License.
  # You may obtain a copy of the License at

  #     http://mozilla.org/MPL/2.0/


  # Unless required by applicable law or agreed to in writing, software
  # distributed under the License is distributed on an "AS IS" BASIS,
  # WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  # See the License for the specific language governing permissions and
  # limitations under the License.
  # */
  username:
    type: string
  password:
    type: string
  timeout:
    type: string
  endpoint:
    type: string
```
{{% /choosable %}}
{{% choosable language java %}}
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
        final var config = ctx.config();
        final var username = config.get("username");
        final var password = config.get("password");
        final var timeout = config.get("timeout");
        final var endpoint = config.get("endpoint");
    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference
### Required

- `endpoint` (String) IP or FQDN of the PowerStore host
- `password` (String, Sensitive) The password of the PowerStore host.
- `username` (String) The username of the PowerStore host.

- `insecure` (Boolean) Boolean variable to specify whether to validate SSL certificate or not.
- `timeout` (Number) The default timeout value for the Powerstore host.