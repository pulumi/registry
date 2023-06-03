---
title: DNSimple
meta_desc: Provides an overview of the DNSimple Provider for Pulumi.
layout: package
---

The DNSimple provider for Pulumi can be used to provision any of the cloud resources available in [DNSimple](https://dnsimple.com/).
The DNSimple provider must be configured with credentials to deploy and update resources in DNSimple.

## Example

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as dnsimple from "@pulumi/dnsimple";

const record = new dnsimple.Record("test", {
  name: "test",
  domain: "mydomain.dev",
  type: dnsimple.RecordTypes.CNAME,
  value: "api.devflix.watch.herokudns.com"
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_dnsimple as dnsimple

record = dnsimple.Record("test",
  name="test",
  domain="mydomain.dev",
  type="CNAME",
  value="api.devflix.watch.herokudns.com"
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	dnsimple "github.com/pulumi/pulumi-dnsimple/sdk/v3/go/dnsimple"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		record, err := dnsimple.NewRecord(ctx, "test", &dnsimple.RecordArgs{
			Name:   pulumi.String("test"),
			Domain: pulumi.String("mydomain.dev"),
			Type:   pulumi.String("CNAME"),
			Value:  pulumi.String("api.devflix.watch.herokudns.com"),
		})
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
using Pulumi.Dnsimple;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            var record = new Record("test", new RecordArgs
            {
                Name = "test",
                Domain = "mydomain.dev",
                Type = "CNAME",
                Value = "api.devflix.watch.herokudns.com",
            });
        });
}
```

{{% /choosable %}}

{{% choosable language java %}}

```java
import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.dnsimple.Record;
import com.pulumi.dnsimple.RecordArgs;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        var record = new Record("record", RecordArgs.builder()
                .name("test")
                .domain("mydomain.dev")
                .type("CNAME")
                .value("api.devflix.watch.herokudns.com")
                .build());

    }
}
```

{{% /choosable %}}
{{% choosable language yaml %}}

```yaml
resources:
  record:
    type: dnsimple:Record
    properties:
      name: test
      domain: mydomain.dev
      type: CNAME
      value: api.devflix.watch.herokudns.com
```

{{% /choosable %}}

{{< /chooser >}}
