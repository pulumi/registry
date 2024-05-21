---
title: Cisco ISE
meta_desc: Provides an overview of the Cisco ISE Provider for Pulumi.
layout: overview
---

The Cisco ISE provider for Pulumi can be used to provision the resources available in Cisco ISE.
The Cisco ISE provider must be configured with credentials to deploy and update resources in ISE.

## Example

{{< chooser language "typescript,python,go,csharp,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as ise from "@pulumi/ise";

const example = new ise.system.Repository("example", {
    name: "repo1",
    protocol: "SFTP",
    path: "/dir",
    serverName: "server1",
    userName: "user9",
    password: "cisco123",
    enablePki: false,
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_ise as ise

example = ise.system.Repository("example",
    name="repo1",
    protocol="SFTP",
    path="/dir",
    server_name="server1",
    user_name="user9",
    password="cisco123",
    enable_pki=False)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi-ise/sdk/go/ise/system"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := system.NewRepository(ctx, "example", &system.RepositoryArgs{
			Name:       pulumi.String("repo1"),
			Protocol:   pulumi.String("SFTP"),
			Path:       pulumi.String("/dir"),
			ServerName: pulumi.String("server1"),
			UserName:   pulumi.String("user9"),
			Password:   pulumi.String("cisco123"),
			EnablePki:  pulumi.Bool(false),
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
using Pulumi;
using Ise = Pulumi.Ise;

return await Deployment.RunAsync(() => 
{
    var example = new Ise.System.Repository("example", new()
    {
        Name = "repo1",
        Protocol = "SFTP",
        Path = "/dir",
        ServerName = "server1",
        UserName = "user9",
        Password = "cisco123",
        EnablePki = false,
    });

});
```

{{% /choosable %}}
{{% choosable language yaml %}}

```yaml
resources:
  example:
    type: ise:system:Repository
    properties:
      name: repo1
      protocol: SFTP
      path: /dir
      serverName: server1
      userName: user9
      password: cisco123
      enablePki: false
```

{{% /choosable %}}

{{< /chooser >}}
