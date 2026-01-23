---
# WARNING: this file was fetched from https://raw.githubusercontent.com/byteplus-sdk/pulumi-bytepluscc/v0.0.12/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/byteplus-sdk/pulumi-bytepluscc/blob/v0.0.12/docs/installation-configuration.md
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Bytepluscc Provider
meta_desc: Provides an overview on how to configure the Pulumi Bytepluscc provider.
layout: package
---
## Installation

The Bytepluscc provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@byteplus/pulumi-bytepluscc`](https://www.npmjs.com/package/@byteplus/pulumi-bytepluscc)
* Python: [`pulumi-bytepluscc`](https://pypi.org/project/pulumi-bytepluscc/)
* Go: [`github.com/byteplus-sdk/pulumi-bytepluscc/sdk/go/bytepluscc`](https://github.com/byteplus-sdk/pulumi-bytepluscc)
* .NET: [`Byteplus.Pulumi.Bytepluscc`](https://www.nuget.org/packages/Byteplus.Pulumi.Bytepluscc)
* Java: [`com.byteplus/bytepluscc`](https://central.sonatype.com/artifact/com.byteplus/bytepluscc)


### Provider Binary

The provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource bytepluscc --server github://api.github.com/byteplus-sdk
```


## Overview

The Bytepluscc provider is used to interact with the
many resources supported by [Byteplus](https://www.byteplus.com). The provider needs to be configured
with the proper credentials before it can be used.

Use the navigation on the left to read about the available resources.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript

import * as pulumi from "@pulumi/pulumi";
import * as bp from  "@byteplus/pulumi-bytepluscc"


// ⚠️ Sensitive information should not be hardcoded in the source code.
//   This example is for demonstration purposes only.
//   It’s recommended to manage such data using environment variables,
//   secure configuration files, or a secret management service.
const provider = new bp.Provider("bp-provider",{
    accessKey:"*******************",
    secretKey:"*******************",
    region:"ap-southeast-1",
    endpoints:{
        cloudcontrolapi:"cloudcontrol.ap-southeast-1.byteplusapi.com"
    }
})

export const CrRepository = (provider:any) => {
    return new bp.cr.Repository("pulumi-node-crRepository",{
        namespace:"ccapi-test",
        registry:"ccapi-test",
        name:"pulumi-bp-nodejs-test",
        description:"pulumi test replace",
        accessLevel:"Public"
    },{provider})
}

```
{{% /choosable %}}
{{% choosable language python %}}
```python
Example currently unavailable in this language

```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: pulumi-dotnet-example
runtime: dotnet
```
```csharp
Example currently unavailable in this language


```
{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	myprovider "github.com/byteplus-sdk/pulumi-bytepluscc/sdk/go/bytepluscc"
	"github.com/byteplus-sdk/pulumi-bytepluscc/sdk/go/bytepluscc/clb"
)

func createTags(tagMap map[string]string) clb.ListenerTagArray {
	var tags []clb.ListenerTagInput
	for key, value := range tagMap {
		tags = append(tags, &clb.ListenerTagArgs{
			Key:   pulumi.String(key),
			Value: pulumi.String(value),
		})
	}
	return tags
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// ⚠️ Sensitive information should not be hardcoded in the source code.
		//   This example is for demonstration purposes only.
		//   It’s recommended to manage such data using environment variables,
		//   secure configuration files, or a secret management service.
		prov, err := myprovider.NewProvider(ctx, "pulumi", &myprovider.ProviderArgs{
			AccessKey: pulumi.String("******************"),
			SecretKey: pulumi.String("******************"),
			Region:    pulumi.String("ap-southeast-1"),
			Endpoints: &myprovider.ProviderEndpointsArgs{
				Cloudcontrolapi: pulumi.String("cloudcontrol.ap-southeast-1.byteplusapi.com"),
			},
		})
		if err != nil {
			return err
		}

		clbListener, err := clb.NewListener(ctx, "clb-listener-http", &clb.ListenerArgs{
			AclIds: pulumi.StringArray{
				pulumi.String("acl-14x5k4odrr3ls3tvs81th56kw"),
			},
			AclStatus:   pulumi.String("on"),
			AclType:     pulumi.String("black"),
			Description: pulumi.String("Clb listener http endpoint"),
			Enabled:     pulumi.String("on"),
			HealthCheck: &clb.ListenerHealthCheckArgs{
				Enabled:            pulumi.String("on"),
				HealthyThreshold:   pulumi.Int(3),
				Interval:           pulumi.Int(5),
				Timeout:            pulumi.Int(2),
				UnhealthyThreshold: pulumi.Int(3),
				Domain:             pulumi.String("www.yjtest.com"),
				Method:             pulumi.String("GET"),
				Uri:                pulumi.String("/"),
				HttpCode:           pulumi.String("http_2xx,http_3xx"),
				Port:               pulumi.Int(0),
			},
			ListenerName:        pulumi.String("clb-listener-http-pulumi"),
			LoadBalancerId:      pulumi.String("clb-14w0vmlz0qi2o3tvs80m4dfrj"),
			PersistenceType:     pulumi.String("server"),
			Cookie:              pulumi.String("cookie123"),
			Protocol:            pulumi.String("HTTP"),
			Port:                pulumi.Int(2005),
			Scheduler:           pulumi.String("wrr"),
			Bandwidth:           pulumi.Int(10),
			ServerGroupId:       pulumi.String("rsp-22qx2ligwir5sds3n8d5tizs"),
			ClientHeaderTimeout: pulumi.Int(60),
			ClientBodyTimeout:   pulumi.Int(40),
			KeepaliveTimeout:    pulumi.Int(90),
			ProxyConnectTimeout: pulumi.Int(10),
			ProxySendTimeout:    pulumi.Int(30),
			ProxyReadTimeout:    pulumi.Int(90),
			SendTimeout:         pulumi.Int(150),
			Tags: createTags(map[string]string{
				"key1": "2000000191",
				"key2": "2000000192",
				"key3": "",
			}),
		}, pulumi.Provider(prov))
		
		if err != nil {
			return err
		}
		return err
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
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language java %}}
```java
Example currently unavailable in this language
```
{{% /choosable %}}
{{< /chooser >}}
