---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/chainguard-dev/cosign/0.0.31/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Cosign Provider
meta_desc: Provides an overview on how to configure the Pulumi Cosign provider.
layout: package
---

## Generate Provider

The Cosign provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider chainguard-dev/cosign
```
## Usage
### Verify

This provides a `cosign.getVerify` function, which can be used with any
containerized infrastructure rules to enforce deploy-time policy checking:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:

```

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language python %}}
```python
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language go %}}
```go
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language yaml %}}
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
## Configuration Reference

- `defaultAttestationEntryType` (String) Default Rekor entry type to use for attestations. Valid values are 'intoto' (default) or 'dsse'.
## Sign

This provider also exposes `cosign.Sign` and `cosign.Attest` resources that will
sign and attest a provided OCI digest, which is intended to compose with
OCI providers such as `ko`,
`apko`, and
`oci`.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as cosign from "@pulumi/cosign";

const example = new cosign.Sign("example", {image: image_build.imageRef});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_cosign as cosign

example = cosign.Sign("example", image=image_build["imageRef"])
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Cosign = Pulumi.Cosign;

return await Deployment.RunAsync(() =>
{
    var example = new Cosign.Sign("example", new()
    {
        Image = image_build.ImageRef,
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/cosign/cosign"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cosign.NewSign(ctx, "example", &cosign.SignArgs{
			Image: pulumi.Any(image_build.ImageRef),
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
  example:
    type: cosign:Sign
    properties:
      image: ${["image-build"].imageRef}
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.cosign.Sign;
import com.pulumi.cosign.SignArgs;
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
        var example = new Sign("example", SignArgs.builder()
            .image(image_build.imageRef())
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Attest

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as cosign from "@pulumi/cosign";

const example = new cosign.Attest("example", {
    image: exampleCosignSign.signedRef,
    predicates: [
        {
            type: "https://example.com/my/predicate/type",
            json: JSON.stringify({}),
        },
        {
            type: "https://example.com/my/predicate/too-big-for-pulumi.tfstate",
            files: {
                path: "/tmp/giant-file.json",
                sha256: "74af7407b59f9021f76a6f9ee66149c5df1ef6442617a805a7860ce18074158d",
            },
        },
    ],
});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import json
import pulumi_cosign as cosign

example = cosign.Attest("example",
    image=example_cosign_sign["signedRef"],
    predicates=[
        {
            "type": "https://example.com/my/predicate/type",
            "json": json.dumps({}),
        },
        {
            "type": "https://example.com/my/predicate/too-big-for-pulumi.tfstate",
            "files": {
                "path": "/tmp/giant-file.json",
                "sha256": "74af7407b59f9021f76a6f9ee66149c5df1ef6442617a805a7860ce18074158d",
            },
        },
    ])
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Cosign = Pulumi.Cosign;

return await Deployment.RunAsync(() =>
{
    var example = new Cosign.Attest("example", new()
    {
        Image = exampleCosignSign.SignedRef,
        Predicates = new[]
        {
            new Cosign.Inputs.AttestPredicateArgs
            {
                Type = "https://example.com/my/predicate/type",
                Json = JsonSerializer.Serialize(new Dictionary<string, object?>
                {
                }),
            },
            new Cosign.Inputs.AttestPredicateArgs
            {
                Type = "https://example.com/my/predicate/too-big-for-pulumi.tfstate",
                Files =
                {
                    { "path", "/tmp/giant-file.json" },
                    { "sha256", "74af7407b59f9021f76a6f9ee66149c5df1ef6442617a805a7860ce18074158d" },
                },
            },
        },
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/cosign/cosign"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = cosign.NewAttest(ctx, "example", &cosign.AttestArgs{
			Image: pulumi.Any(exampleCosignSign.SignedRef),
			Predicates: cosign.AttestPredicateArray{
				&cosign.AttestPredicateArgs{
					Type: pulumi.String("https://example.com/my/predicate/type"),
					Json: pulumi.String(json0),
				},
				&cosign.AttestPredicateArgs{
					Type: pulumi.String("https://example.com/my/predicate/too-big-for-pulumi.tfstate"),
					Files: cosign.AttestPredicateFileArray{
						Path:   "/tmp/giant-file.json",
						Sha256: "74af7407b59f9021f76a6f9ee66149c5df1ef6442617a805a7860ce18074158d",
					},
				},
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
resources:
  example:
    type: cosign:Attest
    properties:
      image: ${exampleCosignSign.signedRef}
      predicates:
        - type: https://example.com/my/predicate/type
          json:
            fn::toJSON: {}
        - type: https://example.com/my/predicate/too-big-for-pulumi.tfstate
          files:
            path: /tmp/giant-file.json
            sha256: 74af7407b59f9021f76a6f9ee66149c5df1ef6442617a805a7860ce18074158d
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.cosign.Attest;
import com.pulumi.cosign.AttestArgs;
import com.pulumi.cosign.inputs.AttestPredicateArgs;
import static com.pulumi.codegen.internal.Serialization.*;
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
        var example = new Attest("example", AttestArgs.builder()
            .image(exampleCosignSign.signedRef())
            .predicates(
                AttestPredicateArgs.builder()
                    .type("https://example.com/my/predicate/type")
                    .json(serializeJson(
                        jsonObject(

                        )))
                    .build(),
                AttestPredicateArgs.builder()
                    .type("https://example.com/my/predicate/too-big-for-pulumi.tfstate")
                    .files(AttestPredicateFileArgs.builder()
                        .path("/tmp/giant-file.json")
                        .sha256("74af7407b59f9021f76a6f9ee66149c5df1ef6442617a805a7860ce18074158d")
                        .build())
                    .build())
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Environment Variables

|          Name           |              Default               |                                          Description                                           |
|-------------------------|------------------------------------|------------------------------------------------------------------------------------------------|
| TF_COSIGN_LOCAL         |                                    | Enables interative signing with Cosign locally (browser based flow only).                      |
| TF_COSIGN_CONNECTOR_ID  |                                    | Configures Sigstore Connector ID for local based signing.                                      |
| TF_COSIGN_ISSUER        | <https://oauth2.sigstore.dev/auth> | Configures Sigstore token issuer.                                                              |
| TF_COSIGN_CLIENT_SECRET |                                    | Configures Sigstore OAuth Client Secret.                                                       |
| TF_COSIGN_REDIRECT_URL  |                                    | Configures Sigstore interactive OAuth redirect URL. If empty, a random localhost port is used. |
| TF_COSIGN_DISABLE       |                                    | Disables all cosign operations.                                                                |