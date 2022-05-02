---
title: Random
meta_desc: Provides an overview of the Random Provider for Pulumi.
layout: overview
---

The Random provider for Pulumi can be used to help introduce random values when dealing with Pulumi resources.

## Example

{{< chooser language "javascript,typescript,python,go,csharp,yaml" >}}

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
{{% choosable language yaml %}}

```yaml
resources:
  username:
    type: random:RandomPet
```

{{% /choosable %}}

{{% choosable language java %}}

```java
import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.random.RandomPet;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    private static void stack(Context ctx) {
		final var username = new RandomPet("my-user-name");
	}
}
```

{{% /choosable %}}

{{< /chooser >}}
