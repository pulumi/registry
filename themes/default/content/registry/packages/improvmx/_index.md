---
# WARNING: this file was fetched from https://raw.githubusercontent.com/lokkju/pulumi-improvmx/v0.2.7/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/lokkju/pulumi-improvmx/blob/v0.2.7/docs/_index.md
title: ImprovMX
meta_desc: Manage ImprovMX email forwarding resources with Pulumi.
layout: package
---

The ImprovMX provider lets you manage [ImprovMX](https://improvmx.com/) email forwarding resources as infrastructure as code.

## Resources

- **Domain** — Register and configure domains for email forwarding
- **EmailAlias** — Create email aliases with forwarding rules (including catch-all `*`)
- **SmtpCredential** — Manage SMTP credentials for sending email

## Example

{{< chooser language "typescript,python,csharp,go" >}}

{{% choosable language typescript %}}

```typescript
import * as improvmx from "pulumi-improvmx";

const domain = new improvmx.Domain("my-domain", {
    domain: "example.com",
});

const alias = new improvmx.EmailAlias("catch-all", {
    domain: domain.domain,
    alias: "*",
    forward: "me@gmail.com",
});
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi_improvmx as improvmx

domain = improvmx.Domain("my-domain", domain="example.com")

alias = improvmx.EmailAlias(
    "catch-all",
    domain=domain.domain,
    alias="*",
    forward="me@gmail.com",
)
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using Pulumi;
using PulumiImprovmx = Pulumi.Improvmx;

var domain = new PulumiImprovmx.Domain("my-domain", new PulumiImprovmx.DomainArgs
{
    DomainName = "example.com",
});

var alias = new PulumiImprovmx.EmailAlias("catch-all", new PulumiImprovmx.EmailAliasArgs
{
    DomainName = domain.DomainName,
    Alias = "*",
    Forward = "me@gmail.com",
});
```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"github.com/lokkju/pulumi-improvmx/sdk/go/improvmx"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		domain, err := improvmx.NewDomain(ctx, "my-domain", &improvmx.DomainArgs{
			Domain: pulumi.String("example.com"),
		})
		if err != nil {
			return err
		}

		_, err = improvmx.NewEmailAlias(ctx, "catch-all", &improvmx.EmailAliasArgs{
			Domain:  domain.Domain,
			Alias:   pulumi.String("*"),
			Forward: pulumi.String("me@gmail.com"),
		})
		return err
	})
}
```

{{% /choosable %}}

{{< /chooser >}}
