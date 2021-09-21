---
title: Confluent Cloud
meta_desc: This page provides an overview of the Confluent Cloud Provider for Pulumi.
layout: overview
---

The Confluent Cloud provider for Pulumi can be used to provision any of the cloud resources available in [Confluent Cloud](https://confluent.cloud/).
The Confluent Cloud provider must be configured with credentials to deploy and update resources in Confluent Cloud.

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const ccloud = require("@pulumi/confluent")

const env = new ccloud.ConfluentEnvironment("ts-environment");
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as confluent from "@pulumi/confluent";

const env = new ccloud.ConfluentEnvironment("ts-environment");
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_confluent as confluent

environment = ccloud.ConfluentEnvironment("py-env")
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
  "github.com/pulumi/pulumi-confluent/sdk/go/confluent"
  "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
  pulumi.Run(func(ctx *pulumi.Context) error {

    env, err := confluent.NewConfluentEnvironment(ctx, "py-env", nil)
    if err != nil {
      return err
    }

    return nil
  })
}

```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;
using Pulumi.Confluent;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            var environment = new CCloud.ConfluentEnvironment("csharp-env");
        });
}
```

{{% /choosable %}}

{{< /chooser >}}

## Libraries

The following packages are available in packager managers:

* JavaScript/TypeScript: [`@pulumi/confluent`](https://www.npmjs.com/package/@pulumi/confluent)
* Python: [`pulumi-confluent`](https://pypi.org/project/pulumi-confluent/)
* Go: [`github.com/pulumi/pulumi-confluent/sdk/go/confluent`](https://github.com/pulumi/pulumi-confluent)
* .NET: [`Pulumi.Confluent`](https://www.nuget.org/packages/Pulumi.Confluent)
