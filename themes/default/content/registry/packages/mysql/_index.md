---
title: MySQL
meta_desc: Provides an overview of the MySQL Provider for Pulumi.
layout: overview
---

The MySQL provider for Pulumi can be used to provision any of the resources available for MySQL.
The MySQL provider must be configured with credentials to deploy and update resources in MySQL.

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const mysql = require("@pulumi/mysql")

const myDb = new mysql.Database("my-database");
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as mysql from "@pulumi/mysql";

const myDb = new mysql.Database("my-database");
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_mysql as mysql

my_db = mysql.Database("my-database")
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	mysql "github.com/pulumi/pulumi-mysql/sdk/v3/go/mysql"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		myDb, err := mysql.NewDatabase(ctx, "my-database", &mysql.DatabaseArgs{})
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
using Pulumi.Mysql;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            mydatabase = new Database("my-database", new DatabaseArgs{});
        });
}
```

{{% /choosable %}}

{{< /chooser >}}
