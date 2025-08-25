---
# WARNING: this file was fetched from https://raw.githubusercontent.com/upstash/pulumi-upstash/v0.4.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Upstash
meta_desc: Provides an overview of the Upstash Provider for Pulumi.
layout: package
---

The Upstash provider for Pulumi can be used to provision Upstash cloud resources such as [Upstash Redis, QStash and Vector](https://upstash.com).

The Upstash provider must be configured with credentials to manage resources in Upstash. Necessary credentials - namely management api keys - can be obtained from [Upstash Console](https://console.upstash.com/account/api).

## Example

{{< chooser language "typescript,go,python" >}}
{{% choosable language typescript %}}

```typescript
import * as upstash from "@upstash/pulumi";
const createdDb = new upstash.RedisDatabase("mydb", {
    databaseName: "pulumi-ts-db",
    region: "global",
	primaryRegion: "us-east-1"
    tls: true,
    multizone: true
})
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/upstash/pulumi-upstash/sdk/go/upstash"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

        createdDb, err := upstash.NewRedisDatabase(ctx, "exampleDB", &upstash.RedisDatabaseArgs{
			DatabaseName: pulumi.String("pulumi-go-db"),
			Region:       pulumi.String("global"),
			PrimaryRegion:       pulumi.String("us-east-1"),
			Tls:          pulumi.Bool(true),
			Multizone:    pulumi.Bool(true),
		})
		if err != nil {
			return err
		}

		ctx.Export("redisDB", createdDb)

		return nil
	})
}
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import upstash_pulumi as upstash

example_db = upstash.RedisDatabase("exampleDB",
    database_name="Pulumi DB",
    region="global",
	primary_region="us-east-1"
    tls=True
)
```

{{% /choosable %}}
{{< /chooser >}}
