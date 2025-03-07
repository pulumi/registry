---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/constellix/constellix/0.4.6/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Constellix Provider
meta_desc: Provides an overview on how to configure the Pulumi Constellix provider.
layout: package
---

## Generate Provider

The Constellix provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider constellix/constellix
```
## Constellix Provider
Constellix is a leading DNS service provider with a feature rich DNS services which includes, various kinds of dns records such as Aname record, Cname record, HTTPredirection, MX record, a granular filtering method using Geofilter and Geoproximity. The Constellix provider is used to manage various DNS Objects supported by Constellix DNS platform. The provider needs to be configured with the proper credentials before it can be used.
## Authentication
The Provider supports authentication with Constellix using API-key and SECRET-key.

1. Authentication with user-id and password.example:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    constellix:apikey:
        value: apikey
    constellix:insecure:
        value: true
    constellix:proxyUrl:
        value: https://proxy_server:proxy_port
    constellix:secretkey:
        value: secretkey

```
## Example Usage
{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    constellix:apikey:
        value: apikey
    constellix:insecure:
        value: true
    constellix:proxyUrl:
        value: https://proxy_server:proxy_port
    constellix:secretkey:
        value: secretkey

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as constellix from "@pulumi/constellix";

const domain1 = new constellix.Domain("domain1", {
    name: "domain1.com",
    soa: {
        primary_nameserver: "ns41.constellix.com.",
        ttl: "1800",
        refresh: "48100",
        retry: "7200",
        expire: "1209",
        negcache: "8000",
    },
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    constellix:apikey:
        value: apikey
    constellix:insecure:
        value: true
    constellix:proxyUrl:
        value: https://proxy_server:proxy_port
    constellix:secretkey:
        value: secretkey

```
```python
import pulumi
import pulumi_constellix as constellix

domain1 = constellix.Domain("domain1",
    name="domain1.com",
    soa={
        "primary_nameserver": "ns41.constellix.com.",
        "ttl": "1800",
        "refresh": "48100",
        "retry": "7200",
        "expire": "1209",
        "negcache": "8000",
    })
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    constellix:apikey:
        value: apikey
    constellix:insecure:
        value: true
    constellix:proxyUrl:
        value: https://proxy_server:proxy_port
    constellix:secretkey:
        value: secretkey

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Constellix = Pulumi.Constellix;

return await Deployment.RunAsync(() =>
{
    var domain1 = new Constellix.Domain("domain1", new()
    {
        Name = "domain1.com",
        Soa =
        {
            { "primary_nameserver", "ns41.constellix.com." },
            { "ttl", "1800" },
            { "refresh", "48100" },
            { "retry", "7200" },
            { "expire", "1209" },
            { "negcache", "8000" },
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
config:
    constellix:apikey:
        value: apikey
    constellix:insecure:
        value: true
    constellix:proxyUrl:
        value: https://proxy_server:proxy_port
    constellix:secretkey:
        value: secretkey

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/constellix/constellix"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := constellix.NewDomain(ctx, "domain1", &constellix.DomainArgs{
			Name: pulumi.String("domain1.com"),
			Soa: pulumi.StringMap{
				"primary_nameserver": pulumi.String("ns41.constellix.com."),
				"ttl":                pulumi.String("1800"),
				"refresh":            pulumi.String("48100"),
				"retry":              pulumi.String("7200"),
				"expire":             pulumi.String("1209"),
				"negcache":           pulumi.String("8000"),
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
config:
    constellix:apikey:
        value: apikey
    constellix:insecure:
        value: true
    constellix:proxyUrl:
        value: https://proxy_server:proxy_port
    constellix:secretkey:
        value: secretkey

```
```yaml
resources:
  domain1:
    type: constellix:Domain
    properties:
      name: domain1.com
      soa:
        primary_nameserver: ns41.constellix.com.
        ttl: 1800
        refresh: 48100
        retry: 7200
        expire: 1209
        negcache: 8000
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    constellix:apikey:
        value: apikey
    constellix:insecure:
        value: true
    constellix:proxyUrl:
        value: https://proxy_server:proxy_port
    constellix:secretkey:
        value: secretkey

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.constellix.Domain;
import com.pulumi.constellix.DomainArgs;
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
        var domain1 = new Domain("domain1", DomainArgs.builder()
            .name("domain1.com")
            .soa(Map.ofEntries(
                Map.entry("primary_nameserver", "ns41.constellix.com."),
                Map.entry("ttl", 1800),
                Map.entry("refresh", 48100),
                Map.entry("retry", 7200),
                Map.entry("expire", 1209),
                Map.entry("negcache", 8000)
            ))
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference
Following arguments are supported with Constellix pulumi provider.

* `apikey` - (Required) API key of a user which has the access to perform CRUD operations on all the DNS objects of Constellix platform.
* `secretkey` - (Required) Secret key of a user which has the access to perform CRUD operations on all the DNS objects of Constellix platform.
* `insecure` - (Optional) This determines whether to use insecure HTTP connection or not. Default value is `true`.
* `proxyUrl` - (Optional) A proxy server URL when configured, all the requests to Constellix platform will be passed through the proxy-server configured.