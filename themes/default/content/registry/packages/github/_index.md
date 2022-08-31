---
title: GitHub
meta_desc: Provides an overview of the GitHub Provider for Pulumi.
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
{{% choosable language java %}}

```java
import com.pulumi.Pulumi;
import com.pulumi.github.Repository;
import com.pulumi.github.RepositoryArgs;

public class App {
    public static void main(String[] args) {
        Pulumi.run(ctx -> {
            var repository = new Repository("demo-repo", RepositoryArgs.builder()
                    .description("Generated from automated test")
                    .visibility("private")
                    .build());
        });
    }
}
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using Pulumi;
using Pulumi.Github;

await Deployment.RunAsync(() =>
{
  var repo = new Repository("test", new RepositoryArgs
  {
      Description = "Generated from automated test",
      Private = true,
  });
});
```

{{% /choosable %}}

{{< /chooser >}}
