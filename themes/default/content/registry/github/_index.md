---
title: GitHub
meta_desc: This page provides an overview of the GitHub Provider for Pulumi.
menu:
  intro:
    parent: cloud-providers
    identifier: clouds-github
    weight: 2
layout: overview
---

The GitHub provider for Pulumi can be used to provision any of the cloud resources available in [GitHub](https://github.com/).
The GitHub provider must be configured with credentials to deploy and update resources in GitHub.

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const github = require("@pulumi/github")

const repo = new github.Repository("demo-repo", {
  description: "Generated from automated test",
  private: true,
});
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as github from "@pulumi/github";

const repo = new github.Repository("demo-repo", {
  description: "Generated from automated test",
  private: true,
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_github as github

repo = github.Repository("demo-repo",
  description="Generated from automated test",
  private="true",
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	github "github.com/pulumi/pulumi-github/sdk/v4/go/github"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		repository, err := github.NewRepository(ctx, "demo-repo", &github.RepositoryArgs{
			Description: pulumi.String("Generated from automated test"),
			Private:     pulumi.Bool(true),
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
using Pulumi.Github;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            var repo = new Repository("test", new RepositoryArgs
            {
                Description = "Generated from automated test",
                Private = true,
            });
        });
}
```

{{% /choosable %}}

{{< /chooser >}}

## Libraries

The following packages are available in packager managers:

* JavaScript/TypeScript: [`@pulumi/github`](https://www.npmjs.com/package/@pulumi/github)
* Python: [`pulumi-github`](https://pypi.org/project/pulumi-github/)
* Go: [`github.com/pulumi/pulumi-github/sdk/v4/go/github`](https://github.com/pulumi/pulumi-github)
* .NET: [`Pulumi.Github`](https://www.nuget.org/packages/Pulumi.Github)

The GitHub provider is open source and available in the [pulumi/pulumi-github](https://github.com/pulumi/pulumi-github) repo.
