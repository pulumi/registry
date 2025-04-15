---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/elastic-infra/ldap/2.0.1/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Ldap Provider
meta_desc: Provides an overview on how to configure the Pulumi Ldap provider.
layout: package
---

## Generate Provider

The Ldap provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider elastic-infra/ldap
```
## Overview

The LDAP provider provides resources to interact with a LDAP object.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    ldap:bindPassword:
        value: admin
    ldap:bindUser:
        value: cn=admin,dc=example,dc=com
    ldap:ldapHost:
        value: localhost
    ldap:ldapPort:
        value: 389

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
    ldap:bindPassword:
        value: admin
    ldap:bindUser:
        value: cn=admin,dc=example,dc=com
    ldap:ldapHost:
        value: localhost
    ldap:ldapPort:
        value: 389

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
    ldap:bindPassword:
        value: admin
    ldap:bindUser:
        value: cn=admin,dc=example,dc=com
    ldap:ldapHost:
        value: localhost
    ldap:ldapPort:
        value: 389

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
    ldap:bindPassword:
        value: admin
    ldap:bindUser:
        value: cn=admin,dc=example,dc=com
    ldap:ldapHost:
        value: localhost
    ldap:ldapPort:
        value: 389

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
    ldap:bindPassword:
        value: admin
    ldap:bindUser:
        value: cn=admin,dc=example,dc=com
    ldap:ldapHost:
        value: localhost
    ldap:ldapPort:
        value: 389

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
    ldap:bindPassword:
        value: admin
    ldap:bindUser:
        value: cn=admin,dc=example,dc=com
    ldap:ldapHost:
        value: localhost
    ldap:ldapPort:
        value: 389

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
### Required

- **bind_password** (String) Password to authenticate the Bind user.
- **bind_user** (String) Bind user to be used for authenticating on the LDAP server.
- **ldap_host** (String) The LDAP server to connect to.

- **ldap_port** (Number) The LDAP protocol port (default: 389).
- **start_tls** (Boolean) Upgrade TLS to secure the connection (default: false).
- **tls** (Boolean) Enable TLS encryption for LDAP (LDAPS) (default: false).
- **tls_insecure** (Boolean) Don't verify server TLS certificate (default: false).