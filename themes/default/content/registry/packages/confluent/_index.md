---
title: Confluent Cloud (Deprecated)
meta_desc: Provides an overview of the Confluent Provider for Pulumi.
layout: package
---

{{% notes type="info" %}}
This provider has been deprecated as of July 2022. It is recommended to use the [Official Confluent Provider](/registry/packages/confluentcloud) as a replacement.
Unfortunately, there is no upgrade path from this provider to the Official Confluent provider, but you can take advantage of the [Pulumi Import](/docs/guides/adopting/import) to help achieve the migration.
{{% /notes %}}

The Confluent Cloud provider for Pulumi can be used to provision any of the cloud resources available in [Confluent Cloud](https://confluent.cloud/).
The Confluent Cloud provider must be configured with credentials to deploy and update resources in Confluent Cloud.

## Example

{{< chooser language "typescript,python,go,csharp" >}}

{{% choosable language typescript %}}

```typescript
import * as ccloud from "@pulumi/confluent";

const env = new ccloud.ConfluentEnvironment("ts-environment");
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_confluent as ccloud

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
using Pulumi;
using Pulumi.Confluent;

await Deployment.RunAsync(() =>
{
    var environment = new CCloud.ConfluentEnvironment("csharp-env");
});
```

{{% /choosable %}}

{{< /chooser >}}
