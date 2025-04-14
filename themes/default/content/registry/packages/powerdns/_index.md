---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/pan-net/powerdns/1.5.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Powerdns Provider
meta_desc: Provides an overview on how to configure the Pulumi Powerdns provider.
layout: package
---

## Generate Provider

The Powerdns provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider pan-net/powerdns
```
## Overview

The PowerDNS provider is used manipulate DNS records supported by PowerDNS server. The provider needs to be configured
with the proper credentials before it can be used. It supports both the [legacy API](https://doc.powerdns.com/3/httpapi/api_spec/) and the new [version 1 API](https://doc.powerdns.com/md/httpapi/api_spec/), however resources may need to be configured differently.

NOTE: if you're using the sqlite3 PowerDNS backend, you might face a problem (as described in #75) with pulumi's
default behavior to run mulitple operations in parallel. Using `-parallelism=1` can help solve the limitations of
the sqlite3 PowerDNS Backend. The MySQL Backend has been verified to work with parallelism, however.

Use the navigation to the left to read about the available resources.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    powerdns:apiKey:
        value: 'TODO: "${var.pdns_api_key}"'
    powerdns:serverUrl:
        value: 'TODO: "${var.pdns_server_url}"'

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as powerdns from "@pulumi/powerdns";

// Create a record
const www = new powerdns.Record("www", {});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    powerdns:apiKey:
        value: 'TODO: "${var.pdns_api_key}"'
    powerdns:serverUrl:
        value: 'TODO: "${var.pdns_server_url}"'

```
```python
import pulumi
import pulumi_powerdns as powerdns

# Create a record
www = powerdns.Record("www")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    powerdns:apiKey:
        value: 'TODO: "${var.pdns_api_key}"'
    powerdns:serverUrl:
        value: 'TODO: "${var.pdns_server_url}"'

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Powerdns = Pulumi.Powerdns;

return await Deployment.RunAsync(() =>
{
    // Create a record
    var www = new Powerdns.Record("www");

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    powerdns:apiKey:
        value: 'TODO: "${var.pdns_api_key}"'
    powerdns:serverUrl:
        value: 'TODO: "${var.pdns_server_url}"'

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/powerdns/powerdns"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a record
		_, err := powerdns.NewRecord(ctx, "www", nil)
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
    powerdns:apiKey:
        value: 'TODO: "${var.pdns_api_key}"'
    powerdns:serverUrl:
        value: 'TODO: "${var.pdns_server_url}"'

```
```yaml
resources:
  # Create a record
  www:
    type: powerdns:Record
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    powerdns:apiKey:
        value: 'TODO: "${var.pdns_api_key}"'
    powerdns:serverUrl:
        value: 'TODO: "${var.pdns_server_url}"'

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.powerdns.Record;
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
        // Create a record
        var www = new Record("www");

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

The following configuration inputs are supported:

* `apiKey` - (Required) The PowerDNS API key. This can also be specified with `PDNS_API_KEY` environment variable.
* `serverUrl` - (Required) The address of PowerDNS server. This can also be specified with `PDNS_SERVER_URL` environment variable. When no schema is provided, the default is `https`.
* `caCertificate` - (Optional) A valid path of a Root CA Certificate in PEM format *or* the content of a Root CA certificate in PEM format. This can also be specified with `PDNS_CACERT` environment variable.
* `insecureHttps` - (Optional) Set this to `true` to disable verification of the PowerDNS server's TLS certificate. This can also be specified with the `PDNS_INSECURE_HTTPS` environment variable.
* `cacheRequests` - (Optional) Set this to `true` to enable cache of the PowerDNS REST API requests. This can also be specified with the `PDNS_CACHE_REQUESTS` environment variable. `WARNING! Enabling this option can lead to the use of stale records when you use other automation to populate the DNS zone records at the same time.`
* `cacheMemSize` - (Optional) Memory size in MB for a cache of the PowerDNS REST API requests. This can also be specified with the `PDNS_CACHE_MEM_SIZE` environment variable.
* `cacheTtl` - (Optional) TTL in seconds for a cache of the PowerDNS REST API requests. This can also be specified with the `PDNS_CACHE_TTL` environment variable.