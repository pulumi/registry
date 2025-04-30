---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-consul/v3.12.5/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Consul Provider
meta_desc: Provides an overview on how to configure the Pulumi Consul provider.
layout: package
---
## Installation

The Consul provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/consul`](https://www.npmjs.com/package/@pulumi/consul)
* Python: [`pulumi-consul`](https://pypi.org/project/pulumi-consul/)
* Go: [`github.com/pulumi/pulumi-consul/sdk/v3/go/consul`](https://github.com/pulumi/pulumi-consul)
* .NET: [`Pulumi.Consul`](https://www.nuget.org/packages/Pulumi.Consul)
* Java: [`com.pulumi/consul`](https://central.sonatype.com/artifact/com.pulumi/consul)
## Overview

[Consul](https://www.consul.io) is a service networking platform which provides
service discovery, service mesh, and application configuration capabilities.
The Consul provider exposes resources used to interact with a
Consul cluster. Configuration of the provider is optional, as it provides
reasonable defaults for all arguments.

Use the navigation to the left to read about the available resources.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    consul:address:
        value: demo.consul.io:80
    consul:datacenter:
        value: nyc1

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as consul from "@pulumi/consul";

// Access a key in Consul
const app = consul.getKeys({
    keys: [{
        name: "ami",
        path: "service/app/launch_ami",
        "default": "ami-1234",
    }],
});
// Use our variable from Consul
const appInstance = new aws.ec2.Instance("app", {ami: app.then(app => app["var"]?.ami)});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    consul:address:
        value: demo.consul.io:80
    consul:datacenter:
        value: nyc1

```
```python
import pulumi
import pulumi_aws as aws
import pulumi_consul as consul

# Access a key in Consul
app = consul.get_keys(keys=[{
    "name": "ami",
    "path": "service/app/launch_ami",
    "default": "ami-1234",
}])
# Use our variable from Consul
app_instance = aws.ec2.Instance("app", ami=app.var["ami"])
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    consul:address:
        value: demo.consul.io:80
    consul:datacenter:
        value: nyc1

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;
using Consul = Pulumi.Consul;

return await Deployment.RunAsync(() =>
{
    // Access a key in Consul
    var app = Consul.GetKeys.Invoke(new()
    {
        Keys = new[]
        {
            new Consul.Inputs.GetKeysKeyInputArgs
            {
                Name = "ami",
                Path = "service/app/launch_ami",
                Default = "ami-1234",
            },
        },
    });

    // Use our variable from Consul
    var appInstance = new Aws.Ec2.Instance("app", new()
    {
        Ami = app.Apply(getKeysResult => getKeysResult.Var?.Ami),
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
    consul:address:
        value: demo.consul.io:80
    consul:datacenter:
        value: nyc1

```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-consul/sdk/v3/go/consul"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Access a key in Consul
		app, err := consul.LookupKeys(ctx, &consul.LookupKeysArgs{
			Keys: []consul.GetKeysKey{
				{
					Name:    "ami",
					Path:    "service/app/launch_ami",
					Default: pulumi.StringRef("ami-1234"),
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		// Use our variable from Consul
		_, err = ec2.NewInstance(ctx, "app", &ec2.InstanceArgs{
			Ami: pulumi.String(app.Var.Ami),
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
    consul:address:
        value: demo.consul.io:80
    consul:datacenter:
        value: nyc1

```
```yaml
resources:
  # Use our variable from Consul
  appInstance:
    type: aws:ec2:Instance
    name: app
    properties:
      ami: ${app.var.ami}
variables:
  # Access a key in Consul
  app:
    fn::invoke:
      Function: consul:getKeys
      Arguments:
        keys:
          - name: ami
            path: service/app/launch_ami
            default: ami-1234
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    consul:address:
        value: demo.consul.io:80
    consul:datacenter:
        value: nyc1

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.consul.ConsulFunctions;
import com.pulumi.consul.inputs.GetKeysArgs;
import com.pulumi.aws.ec2.Instance;
import com.pulumi.aws.ec2.InstanceArgs;
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
        // Access a key in Consul
        final var app = ConsulFunctions.getKeys(GetKeysArgs.builder()
            .keys(GetKeysKeyArgs.builder()
                .name("ami")
                .path("service/app/launch_ami")
                .default_("ami-1234")
                .build())
            .build());

        // Use our variable from Consul
        var appInstance = new Instance("appInstance", InstanceArgs.builder()
            .ami(app.applyValue(getKeysResult -> getKeysResult.var().ami()))
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

- `address` (String) The HTTP(S) API address of the agent to use. Defaults to "127.0.0.1:8500".
- `authJwt` (Block List, Max: 1) Authenticates to Consul using a JWT authentication method. (see below for nested schema)
- `caFile` (String) A path to a PEM-encoded certificate authority used to verify the remote agent's certificate.
- `caPath` (String) A path to a directory of PEM-encoded certificate authority files to use to check the authenticity of client and server connections. Can also be specified with the `CONSUL_CAPATH` environment variable.
- `caPem` (String) PEM-encoded certificate authority used to verify the remote agent's certificate.
- `certFile` (String) A path to a PEM-encoded certificate provided to the remote agent; requires use of `keyFile` or `keyPem`.
- `certPem` (String) PEM-encoded certificate provided to the remote agent; requires use of `keyFile` or `keyPem`.
- `datacenter` (String) The datacenter to use. Defaults to that of the agent.
- `header` (Block List) A configuration block, described below, that provides additional headers to be sent along with all requests to the Consul server. This block can be specified multiple times. (see below for nested schema)
- `httpAuth` (String) HTTP Basic Authentication credentials to be used when communicating with Consul, in the format of either `user` or `user:pass`. This may also be specified using the `CONSUL_HTTP_AUTH` environment variable.
- `insecureHttps` (Boolean) Boolean value to disable SSL certificate verification; setting this value to true is not recommended for production use. Only use this with scheme set to "https".
- `keyFile` (String) A path to a PEM-encoded private key, required if `certFile` or `certPem` is specified.
- `keyPem` (String) PEM-encoded private key, required if `certFile` or `certPem` is specified.
- `namespace` (String)
- `scheme` (String) The URL scheme of the agent to use ("http" or "https"). Defaults to "http".
- `token` (String, Sensitive) The ACL token to use by default when making requests to the agent. Can also be specified with `CONSUL_HTTP_TOKEN` or `CONSUL_TOKEN` as an environment variable.

<a id="nestedblock--auth_jwt"></a>
### Nested Schema for `authJwt`

Required:

- `authMethod` (String) The name of the auth method to use for login.

Optional:

- `bearerToken` (String) The bearer token to present to the auth method during login for authentication purposes. For the Kubernetes auth method this is a [Service Account Token (JWT)](https://kubernetes.io/docs/reference/access-authn-authz/authentication/#service-account-tokens).
- `meta` (Map of String) Specifies arbitrary KV metadata linked to the token. Can be useful to track origins.
- `usePulumiCloudWorkloadIdentity` (Boolean) Whether to use a Pulumi Workload Identity token. The token will be read from the `TFC_WORKLOAD_IDENTITY_TOKEN` environment variable.

<a id="nestedblock--header"></a>
### Nested Schema for `header`

Required:

- `name` (String) The name of the header.
- `value` (String) The value of the header.
## Environment Variables

All environment variables listed in the [Consul environment variables](https://www.consul.io/docs/commands/index.html#environment-variables)
documentation are supported by the Pulumi provider.