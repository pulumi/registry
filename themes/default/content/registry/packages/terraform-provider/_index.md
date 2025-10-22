---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-terraform-provider/v1.0.1/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi/pulumi-terraform-provider/blob/v1.0.1/docs/_index.md
title: Any Terraform Provider
meta_desc: Learn how to use any Terraform provider in Pulumi.
layout: package
---

Pulumi's `terraform-provider` can be used to generate full Pulumi SDKs for *any* Terraform provider.

{{% notes type="warning" %}}

This provider is in Public Beta. We are still making breaking changes to nail down the
final design.

{{% /notes %}}

Any feedback is welcome! Please file suggestions and bugs at [pulumi/pulumi-terraform-provider](https://github.com/pulumi/pulumi-terraform-provider/issues).

## Example

As an example, let's depend directly on Hashicorp's [terraform-provider-random](https://github.com/hashicorp/terraform-provider-random).

{{% notes type="info" %}}

You will need to be using a version of Pulumi >= 3.147.0.

{{% /notes %}}

In an existing project, run:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}

```console
$ pulumi package add terraform-provider hashicorp/random
Successfully generated a Nodejs SDK for the random package at /Projects/pulumi/langs/typescript/sdks/random

To use this SDK in your Nodejs project, run the following command:

  npm add random@file:sdks/random

You can then import the SDK in your TypeScript code with:

  import * as random from "random";

```

{{% /choosable %}}
{{% choosable language python %}}

```console
$ pulumi package add terraform-provider hashicorp/random
Successfully generated a Python SDK for the random package at /Projects/pulumi/langs/python/sdks/random

To use this SDK in your Python project, run the following command:

  echo sdks/random >> requirements.txt

  pulumi install

You can then import the SDK in your Python code with:

  import pulumi_random as random

```

{{% /choosable %}}
{{% choosable language go %}}

```console
$ pulumi package add terraform-provider hashicorp/random
Successfully generated a Go SDK for the random package at /Projects/pulumi/langs/golang/sdks/random

To use this SDK in your Go project, run the following command:

   go mod edit -replace github.com/pulumi/pulumi-terraform-provider/sdks/go/random/v3=./sdks/random

You can then use the SDK in your Go code with:

  import "github.com/pulumi/pulumi-terraform-provider/sdks/go/random/v3/random"

```

{{% /choosable %}}
{{% choosable language csharp %}}

```console
$ pulumi package add terraform-provider hashicorp/random
Successfully generated a .NET SDK for the random package at /Projects/pulumi/langs/csharp/sdks/random

To use this SDK in your .NET project, run the following command:

  dotnet add reference sdks/random

You also need to add the following to your .csproj file of the program:

  <DefaultItemExcludes>$(DefaultItemExcludes);sdks/**/*.cs</DefaultItemExcludes>

You can then use the SDK in your .NET code with:

  using Pulumi.Random;

```

{{% /choosable %}}
{{% choosable language java %}}

```console
$ pulumi package add terraform-provider hashicorp/random
Successfully generated a Java SDK for the random package at /Projects/pulumi/langs/java/sdks/random

To use this SDK in your Java project, complete the following steps:

1. Copy the contents of the generated SDK to your Java project:
     cp -r /Projects/pulumi/langs/java/sdks/random/src/* /Projects/pulumi/langs/java/src

2. Add the SDK's dependencies to your Java project's build configuration.
   If you are using Maven, add the following dependencies to your pom.xml:

     <dependencies>
         <dependency>
             <groupId>com.google.code.findbugs</groupId>
             <artifactId>jsr305</artifactId>
             <version>3.0.2</version>
         </dependency>
         <dependency>
             <groupId>com.google.code.gson</groupId>
             <artifactId>gson</artifactId>
             <version>2.8.9</version>
         </dependency>
     </dependencies>
```

{{% /choosable %}}
{{% choosable language yaml %}}

```console
$ pulumi package add terraform-provider hashicorp/random
Downloading provider: terraform-provider
```

Note that the YAML language does not need an SDK.

{{% /choosable %}}
{{< /chooser >}}

By following the instructions given, we have generated a Pulumi SDK for the provider and
linked it into our project. We can now import the SDK we generate and consume a resource from it:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}

```ts
import * as random from "random";

new random.Pet("hi");
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_random as random

random.Pet("hi")
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi-terraform-provider/sdks/go/random/v3/random"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := random.NewPet(ctx, "pet", &random.PetArgs{})
		return err
	})
}
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using Pulumi;
using Pulumi.Random;

class MyStack : Stack
{
    public MyStack()
    {
        new Pet("pet");
    }
}
```

{{% /choosable %}}
{{% choosable language java %}}

```java
package myproject;

import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.random.Pet;

public class App {
    public static void main(String[] args) {
        Pulumi.run(ctx -> {
                new Pet("pet");
        });
    }
}
```

{{% /choosable %}}
{{% choosable language yaml %}}

```yaml
name: Pulumi Program
runtime: yaml
resources:
  pet:
    type: random:Pet
    properties:
      length: 2
```

{{% /choosable %}}
{{< /chooser >}}
