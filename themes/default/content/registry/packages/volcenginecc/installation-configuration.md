---
# WARNING: this file was fetched from https://raw.githubusercontent.com/volcengine/pulumi-volcenginecc/v0.0.5/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/volcengine/pulumi-volcenginecc/blob/v0.0.5/docs/installation-configuration.md
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Volcenginecc Provider
meta_desc: Provides an overview on how to configure the Pulumi Volcenginecc provider.
layout: package
---
## Installation

The Volcenginecc provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@volcengine/pulumi-volcenginecc`](https://www.npmjs.com/package/@volcengine/pulumi-volcenginecc)
* Python: [`pulumi-volcenginecc`](https://pypi.org/project/pulumi-volcenginecc/)
* Go: [`github.com/volcengine/pulumi-volcenginecc/sdk/go/volcenginecc`](https://github.com/volcengine/pulumi-volcenginecc)
* .NET: [`Volcengine.Pulumi.Volcenginecc`](https://www.nuget.org/packages/Volcengine.Pulumi.Volcenginecc)
* Java: [`com.volcengine/volcenginecc`](https://central.sonatype.com/artifact/com.volcengine/volcenginecc)


### Provider Binary

The provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource volcenginecc --server github://api.github.com/volcengine
```


## Overview

The Volcenginecc provider is used to interact with the
many resources supported by [Volcengine](https://www.volcengine.com). The provider needs to be configured
with the proper credentials before it can be used.

Use the navigation on the left to read about the available resources.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: pulumi-typescript-example
description: A minimal TypeScript Pulumi program
runtime:
  name: nodejs
  options:
    packagemanager: npm
config:
  pulumi:tags:
    value:
      pulumi:template: typescript

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as ve from "@volcengine/pulumi-volcenginecc"
import {ProviderEndpoints} from "../pulumi-volcenginecc/sdk/nodejs/types/input";


// ⚠️ Sensitive information should not be hardcoded in the source code.
//   This example is for demonstration purposes only.
//   It’s recommended to manage such data using environment variables,
//   secure configuration files, or a secret management service.
const provider = new ve.Provider("volcengine-provider", {
    region: "cn-beijing",   // REGION
    accessKey: "",   // ACCESSKEY
    secretKey: "", // SECRET_ACCESS_KEY
    endpoints: {
        cloudcontrolapi: "cloudcontrol.cn-beijing.volcengineapi.com"
    },
});

const user = new ve.iam.User("pulumi-node-user", {
    userName: "pulumi-nodejs-user-1",           // Required
    description: "pulumi user description update",     // Optional
    displayName: "pulumi-nodejs-display",       // Optional
    groups: ["ccapi-test"],                     // Optional
    policies: [                                 // Optional
        { policyName: "ReadOnlyAccess", policyType: "System" },
        { policyName: "TOSReadOnlyAccess", policyType: "System" },
        { policyName: "VPCFullAccess", policyType: "System" },
        { policyName: "IAMFullAccess", policyType: "System" },

    ],
    tags: [                                     // Optional
        { key: "env", value: "dev" },
        { key: "team", value: "backend" },
    ],
}, { provider });
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
name: pulumi-python-demo
description: pulumi python demo
runtime:
  name: python
  options:
    toolchain: pip
    virtualenv: venv
config:
  pulumi:tags:
    value:
      pulumi:template: python

```
```python
"""A Python Pulumi program"""

import pulumi
import pulumi_volcenginecc as ve
# ⚠️ Sensitive information should not be hardcoded in the source code.
#   This example is for demonstration purposes only.
#   It’s recommended to manage such data using environment variables,
#   secure configuration files, or a secret management service.
provider = ve.Provider("volcengine",
                       access_key="",
                       secret_key="",
                       region="cn-beijing",
                       endpoints={
                           "cloudcontrol": "cloudcontrol.cn-beijing.volcengineapi.com"
                       })
args = ve.iam.UserArgs(
    user_name="pulumi-python-user-1",
    description="pulumi user description update",
    display_name="pulumi-python-user-diplay",
    groups=["ccapi-test"],
    policies=[
        ve.iam.UserPolicyArgs(policy_name="ReadOnlyAccess", policy_type="System"),
        ve.iam.UserPolicyArgs(policy_name="TOSReadOnlyAccess", policy_type="System"),
        ve.iam.UserPolicyArgs(policy_name="ECSFullAccess", policy_type="System"),
        ve.iam.UserPolicyArgs(policy_name="VPCFullAccess", policy_type="System"),
    ],
    tags=[
        ve.iam.UserTagArgs(key="env", value="dev"),
        ve.iam.UserTagArgs(key="team", value="backend")
    ]
)

user = ve.iam.User("pulumi-python-user-1",
                   args, opts=pulumi.ResourceOptions(provider=provider))

```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: pulumi-dotnet-example
runtime: dotnet
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Pulumi.Volcenginecc;
using Pulumi.Volcenginecc.Inputs;

return await Deployment.RunAsync(() => 
{
    // ⚠️ Sensitive information should not be hardcoded in the source code.
    //   This example is for demonstration purposes only.
    //   It’s recommended to manage such data using environment variables,
    //   secure configuration files, or a secret management service.
    var provider = new Provider("volcengine", new ProviderArgs
    {
        AccessKey = "",
        SecretKey = "",
        Region = "cn-beijing",
        Endpoints =new ProviderEndpointsArgs
        {
            Cloudcontrolapi = "cloudcontrol.cn-beijing.volcengineapi.com",
        }
        
    });
     var user = new Pulumi.Volcenginecc.Iam.User("pulumi-csharp-user-1", new Pulumi.Volcenginecc.Iam.UserArgs
        {
            UserName = "pulumi-dotnet-user-1",
            Description = "pulumi user description ",
            DisplayName = "pulumi-java-user-diplay",
            Groups = { "ccapi-test" },

            Policies =
            {
                new Pulumi.Volcenginecc.Iam.Inputs.UserPolicyArgs { PolicyName = "TOSReadOnlyAccess", PolicyType = "System" },
                new Pulumi.Volcenginecc.Iam.Inputs.UserPolicyArgs { PolicyName = "ReadOnlyAccess",    PolicyType = "System" },
                new Pulumi.Volcenginecc.Iam.Inputs.UserPolicyArgs { PolicyName = "IAMFullAccess",     PolicyType = "System" },
                new Pulumi.Volcenginecc.Iam.Inputs.UserPolicyArgs { PolicyName = "ECSFullAccess",     PolicyType = "System" },

            },

            Tags =
            {
                new Pulumi.Volcenginecc.Iam.Inputs.UserTagArgs { Key = "key1",    Value = "Value1" },
                new Pulumi.Volcenginecc.Iam.Inputs.UserTagArgs { Key = "TagXXX",  Value = "ValueXXX" },
            },
        }, new CustomResourceOptions
        {
            Provider = provider
        });
});



```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: pulumi-iam-user
description: pulumi-iam-user
runtime: go
config:
  pulumi:tags:
    value:
      pulumi:template: go
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	myprovider "github.com/volcengine/pulumi-volcenginecc/sdk/go/volcenginecc"
	"github.com/volcengine/pulumi-volcenginecc/sdk/go/volcenginecc/iam"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// ⚠️ Sensitive information should not be hardcoded in the source code.
		//   This example is for demonstration purposes only.
		//   It’s recommended to manage such data using environment variables,
		//   secure configuration files, or a secret management service.
		prov, err := myprovider.NewProvider(ctx, "pulumi", &myprovider.ProviderArgs{
			AccessKey: pulumi.String(""),
			SecretKey: pulumi.String("=="),
			Region:    pulumi.String("cn-beijing"),
			Endpoints: &myprovider.ProviderEndpointsArgs{
				Cloudcontrolapi: pulumi.String("cloudcontrol.cn-beijing.volcengineapi.com"),
			},
		})
		if err != nil {
			return err
		}

		
		user, err := iam.NewUser(ctx, "pulumi-test-user", &iam.UserArgs{
			UserName:    pulumi.String("pulumi-user-2"),
			DisplayName: pulumi.StringPtr("pulumi-user-display"),
			Description: pulumi.StringPtr("pulumi user description update"),
			Groups:      pulumi.ToStringArray([]string{"ccapi-test"}),
			Policies: iam.UserPolicyArray{
				iam.UserPolicyArgs{
					PolicyName: pulumi.StringPtr("TOSReadOnlyAccess"),
					PolicyType: pulumi.StringPtr("System"),
				},
				iam.UserPolicyArgs{
					PolicyName: pulumi.StringPtr("ReadOnlyAccess"),
					PolicyType: pulumi.StringPtr("System"),
				},
				iam.UserPolicyArgs{
					PolicyName: pulumi.StringPtr("ECSFullAccess"),
					PolicyType: pulumi.StringPtr("System"),
				},
				iam.UserPolicyArgs{
					PolicyName: pulumi.StringPtr("IAMFullAccess"),
					PolicyType: pulumi.StringPtr("System"),
				},
			},
			Tags: iam.UserTagArray{
				iam.UserTagArgs{
					Key:   pulumi.StringPtr("key1"),
					Value: pulumi.StringPtr("Value1"),
				},
				iam.UserTagArgs{
					Key:   pulumi.StringPtr("TagKey1"),
					Value: pulumi.StringPtr("TagValue1"),
				},
				iam.UserTagArgs{
					Key:   pulumi.StringPtr("TagXXX"),
					Value: pulumi.StringPtr("ValueXXX"),
				},
			},
		}, pulumi.Provider(prov))
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
```yaml
# Pulumi.yaml provider configuration file
name: pulumi-java-iam-user
description: pulumi java iam user
runtime: java
config:
  pulumi:tags:
    value:
      pulumi:template: java

```
```java
package myproject;

import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.resources.CustomResourceOptions;
import com.volcengine.volcenginecc.Provider;
import com.volcengine.volcenginecc.ProviderArgs;
import com.volcengine.volcenginecc.iam.User;
import com.volcengine.volcenginecc.iam.UserArgs;
import com.volcengine.volcenginecc.iam.inputs.UserPolicyArgs;
import com.volcengine.volcenginecc.iam.inputs.UserTagArgs;
import com.volcengine.volcenginecc.inputs.ProviderEndpointsArgs;

public class App {
    public static void main(String[] args) {
        Pulumi.run(ctx -> {
            // ⚠️ Sensitive information should not be hardcoded in the source code.
            //   This example is for demonstration purposes only.
            //   It’s recommended to manage such data using environment variables,
            //   secure configuration files, or a secret management service.
            Provider volcenginecc = new Provider("volcenginecc", ProviderArgs.builder()
                    .accessKey("")
                    .secretKey("")
                    .region("cn-beijing")
                    .endpoints(ProviderEndpointsArgs.builder()
                            .cloudcontrolapi("cloudcontrol.cn-beijing.volcengineapi.com")
                            .build())
                    .build());
         
            User user = new User("pulumi-java-user-1", UserArgs.builder()
                    .userName("pulumi-java-user-1")
                    .description("pulumi user description update")
                    .displayName("pulumi-java-user-diplay")
                    .groups("ccapi-test")
                    .policies(
                            UserPolicyArgs.builder().policyName("TOSReadOnlyAccess").policyType("System").build(),
                            UserPolicyArgs.builder().policyName("ReadOnlyAccess").policyType("System").build(),
                            UserPolicyArgs.builder().policyName("IAMFullAccess").policyType("System").build(),
                            UserPolicyArgs.builder().policyName("VPCFullAccess").policyType("System").build()
                    )
                    .tags(
                            UserTagArgs.builder().key("key1").value("Value1").build(),
                            UserTagArgs.builder().key("TagKey1").value("TagValue1").build(),
                            UserTagArgs.builder().key("TagXXX").value("ValueXXX").build()
                    )
                    .build(),
                    CustomResourceOptions.builder().provider(volcenginecc).build());
        });
    }
}

```
{{% /choosable %}}
{{< /chooser >}}
