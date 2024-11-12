---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumiverse/pulumi-aquasec/v0.8.29/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Aquasec
meta_desc: Provides an overview of the Aquasec Provider for Pulumi.
layout: package
---

The Aquasec provider for Pulumi can be used to provision any of the cloud resources available in [Aquasec](https://www.aquasec.com/aqua-cloud-native-security-platform/).
The Aquasec provider must be configured with credentials and url to deploy and update resources in Aquasec.

## Example

{{< chooser language "typescript,python,go,csharp" >}}
{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aquasec from "@pulumiverse/aquasec";

const registry = "Docker Hub";
const repository = "golang";
const tag = "1.19";

let golangImage = new aquasec.Image("image", {
  registry: registry,
  repository: repository,
  tag: tag,
})

export const architecture = golangImage.architecture
export const criticalVulnerabilities = golangImage.criticalVulnerabilities
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumiverse_aquasec as aquasec

image = aquasec.Image("image",
    registry="Docker Hub",
    repository="golang",
    tag=["1.19"],
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-aquasec/sdk/go/aquasec"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		image, err := aquasec.NewImage(ctx, "image", &aquasec.ImageArgs{
			Registry:   pulumi.String("Docker Hub"),
			Repository: pulumi.String("golang"),
			Tag:        pulumi.String("1.19"),
		})
		if err != nil {
			return err
		}
		ctx.Export("architecture", image.Architecture)
		ctx.Export("criticalVulnerabilities", image.CriticalVulnerabilities)
		return nil
	})
}
```

{{% /choosable %}}

{{< /chooser >}}
