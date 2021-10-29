---
title: Random
meta_desc: This page provides an overview of the Random Provider for Pulumi.
layout: overview
---

The Random provider for Pulumi can be used to help introduce random values when dealing with Pulumi resources.

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const random = require("@pulumi/random")

const username = new random.RandomPet("my-user-name");
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as random from "@pulumi/random";

const username = new random.RandomPet("my-user-name");
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_random as random

username = random.RandomPet("my-user-name")
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	random "github.com/pulumi/pulumi-random/sdk/v4/go/random"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		username, err := random.NewRandomPet(ctx, "my-user-name", &random.RandomPetArgs{})
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
using Pulumi.Random;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            var username = new RandomPet("my-user-name", new RandomPetArgs{});
        });
}
```

{{% /choosable %}}

{{< /chooser >}}
