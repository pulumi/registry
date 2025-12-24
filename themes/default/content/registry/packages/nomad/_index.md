---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-nomad/v2.5.4/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi/pulumi-nomad/blob/v2.5.4/docs/_index.md
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Nomad Provider
meta_desc: Provides an overview on how to configure the Pulumi Nomad provider.
layout: package
---

## Installation

The Nomad provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/nomad`](https://www.npmjs.com/package/@pulumi/nomad)
* Python: [`pulumi-nomad`](https://pypi.org/project/pulumi-nomad/)
* Go: [`github.com/pulumi/pulumi-nomad/sdk/v2/go/nomad`](https://github.com/pulumi/pulumi-nomad)
* .NET: [`Pulumi.Nomad`](https://www.nuget.org/packages/Pulumi.Nomad)
* Java: [`com.pulumi/nomad`](https://central.sonatype.com/artifact/com.pulumi/nomad)

## Overview

[HashiCorp Nomad](https://www.nomadproject.io) is an application scheduler. The
Nomad provider exposes resources to interact with a Nomad cluster.

Use the navigation to the left to read about the available resources.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    nomad:address:
        value: http://nomad.mycompany.com:4646
    nomad:region:
        value: us-east-2

```

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as nomad from "@pulumi/nomad";
import * as std from "@pulumi/std";

function notImplemented(message: string) {
    throw new Error(message);
}

// Register a job
const monitoring = new nomad.Job("monitoring", {jobspec: std.file({
    input: `${notImplemented("path.module")}/jobspec.hcl`,
}).then(invoke => invoke.result)});
```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    nomad:address:
        value: http://nomad.mycompany.com:4646
    nomad:region:
        value: us-east-2

```

```python
import pulumi
import pulumi_nomad as nomad
import pulumi_std as std


def not_implemented(msg):
    raise NotImplementedError(msg)

# Register a job
monitoring = nomad.Job("monitoring", jobspec=std.file(input=f"{not_implemented('path.module')}/jobspec.hcl").result)
```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    nomad:address:
        value: http://nomad.mycompany.com:4646
    nomad:region:
        value: us-east-2

```

```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Nomad = Pulumi.Nomad;
using Std = Pulumi.Std;


object NotImplemented(string errorMessage)
{
    throw new System.NotImplementedException(errorMessage);
}

return await Deployment.RunAsync(() =>
{
    // Register a job
    var monitoring = new Nomad.Job("monitoring", new()
    {
        Jobspec = Std.File.Invoke(new()
        {
            Input = $"{NotImplemented("path.module")}/jobspec.hcl",
        }).Apply(invoke => invoke.Result),
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
    nomad:address:
        value: http://nomad.mycompany.com:4646
    nomad:region:
        value: us-east-2

```

```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-nomad/sdk/v2/go/nomad"
	"github.com/pulumi/pulumi-std/sdk/go/std"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func notImplemented(message string) pulumi.AnyOutput {
	panic(message)
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		invokeFile, err := std.File(ctx, &std.FileArgs{
			Input: fmt.Sprintf("%v/jobspec.hcl", notImplemented("path.module")),
		}, nil)
		if err != nil {
			return err
		}
		// Register a job
		_, err = nomad.NewJob(ctx, "monitoring", &nomad.JobArgs{
			Jobspec: pulumi.String(invokeFile.Result),
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
    nomad:address:
        value: http://nomad.mycompany.com:4646
    nomad:region:
        value: us-east-2

```

```yaml
Example currently unavailable in this language
```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    nomad:address:
        value: http://nomad.mycompany.com:4646
    nomad:region:
        value: us-east-2

```

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.nomad.Job;
import com.pulumi.nomad.JobArgs;
import com.pulumi.std.StdFunctions;
import com.pulumi.std.inputs.FileArgs;
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
        // Register a job
        var monitoring = new Job("monitoring", JobArgs.builder()
            .jobspec(StdFunctions.file(FileArgs.builder()
                .input(String.format("%s/jobspec.hcl", "TODO: call notImplemented"))
                .build()).result())
            .build());

    }
}
```

{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

The following configuration inputs are supported:

- `address` `(string: "http://127.0.0.1:4646")` - The HTTP(S) API address of the
  Nomad agent. This must include the leading protocol (e.g. `https://`). This
  can also be specified as the `NOMAD_ADDR` environment variable.

- `region` `(string: "")` - The Nomad region to target. This can also be
  specified as the `NOMAD_REGION` environment variable.

- `httpAuth` `(string: "")` - HTTP Basic Authentication credentials to be used
  when communicating with Nomad, in the format of either `user` or `user:pass`.
  This can also be specified using the `NOMAD_HTTP_AUTH` environment variable.

- `caFile` `(string: "")` - A local file path to a PEM-encoded certificate
  authority used to verify the remote agent's certificate. This can also be
  specified as the `NOMAD_CACERT` environment variable.

- `caPem` `(string: "")` - PEM-encoded certificate authority used to verify
  the remote agent's certificate.

- `certFile` `(string: "")` - A local file path to a PEM-encoded certificate
  provided to the remote agent. If this is specified, `keyFile` or `keyPem`
  is also required. This can also be specified as the `NOMAD_CLIENT_CERT`
  environment variable.

- `certPem` `(string: "")` - PEM-encoded certificate provided to the remote
  agent. If this is specified, `keyFile` or `keyPem` is also required.

- `keyFile` `(string: "")` - A local file path to a PEM-encoded private key.
  This is required if `certFile` or `certPem` is specified. This can also be
  specified via the `NOMAD_CLIENT_KEY` environment variable.

- `keyPem` `(string: "")` - PEM-encoded private key. This is required if
  `certFile` or `certPem` is specified.

- `skipVerify` `(boolean: false)` - Set this to true if you want to skip TLS verification on the client side.
  This can also be specified via the `NOMAD_SKIP_VERIFY` environment variable.

- `headers` - (Optional) A configuration block, described below, that provides headers
  to be sent along with all requests to Nomad.  This block can be specified
  multiple times.

- `secretId` `(string: "")` - The Secret ID of an ACL token to make requests with,
  for ACL-enabled clusters. This can also be specified via the `NOMAD_TOKEN`
  environment variable.

- `ignoreEnvVars` `(map[string]bool: {})` - A map of environment variables
  that are ignored by the provider when configuring the Nomad API client.
  Supported keys are: `NOMAD_NAMESPACE` and `NOMAD_REGION`.

- `authJwt` - Authenticates to Nomad using a JWT authentication method.

The `headers` nested type accepts the following arguments:
* `name` - (Required) The name of the header.
* `value` - (Required) The value of the header.

The `authJwt` configuration block accepts the following arguments:
* `authMethod` - (Required) The name of the auth method.
* `loginToken` - (Required) The value of the jwt token.

An example using the `authJwt` configuration block :
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:

```