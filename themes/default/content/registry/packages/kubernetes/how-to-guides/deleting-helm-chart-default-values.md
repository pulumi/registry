---
title: "Deleting Helm Chart Default Values"
meta_desc: How to delete a default value from a Helm chart in a Pulumi program by setting it to null.
layout: package
---

> **Note:** Requires pulumi-kubernetes v4.30.0 or later. Earlier versions strip `null` values during Helm value merging, so the patterns below have no effect on those versions.

Helm charts commonly provide default values in `values.yaml`, and many chart templates guard fields with truthy checks like `{{- if .Values.foo }}`. The standard way to remove one of those defaults is to set the value to `null` — the equivalent of `helm install --set foo=null` on the Helm CLI.

## TypeScript: set the key to `null` in the inline `values` map

In TypeScript this is straightforward — set the key to `null` in the `values` map and it flows through to Helm:

```typescript
import * as k8s from "@pulumi/kubernetes";

const nginx = new k8s.helm.v3.Release("nginx", {
    chart: "nginx",
    repositoryOpts: {
        repo: "https://charts.bitnami.com/bitnami",
    },
    values: {
        containerPorts: {
            http: null,
        },
    },
});
```

## Other SDKs: use `valueYamlFiles`

Inline `null` in the `values` map only works from TypeScript — other Pulumi SDKs strip null map values before they reach the provider. The `valueYamlFiles` path avoids this because the file content is shipped as a Pulumi [Asset](https://www.pulumi.com/docs/concepts/assets-archives/#assets) and parsed on the provider side. Put the explicit `null` in a yaml file and reference it:

```yaml
# overrides.yaml
containerPorts:
  http: null
```

{{< chooser language "python,go,csharp,java,yaml" >}}

{{% choosable language python %}}

```python
import pulumi
from pulumi_kubernetes.helm.v3 import Release, ReleaseArgs, RepositoryOptsArgs

nginx = Release(
    "nginx",
    ReleaseArgs(
        chart="nginx",
        repository_opts=RepositoryOptsArgs(
            repo="https://charts.bitnami.com/bitnami",
        ),
        value_yaml_files=[pulumi.FileAsset("./overrides.yaml")],
    ),
)
```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := helmv3.NewRelease(ctx, "nginx", &helmv3.ReleaseArgs{
			Chart: pulumi.String("nginx"),
			RepositoryOpts: helmv3.RepositoryOptsArgs{
				Repo: pulumi.String("https://charts.bitnami.com/bitnami"),
			},
			ValueYamlFiles: pulumi.AssetOrArchiveArray{
				pulumi.NewFileAsset("./overrides.yaml"),
			},
		})
		return err
	})
}
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using Pulumi;
using Pulumi.Kubernetes.Helm.V3;
using Pulumi.Kubernetes.Types.Inputs.Helm.V3;

return await Deployment.RunAsync(() =>
{
    new Release("nginx", new ReleaseArgs
    {
        Chart = "nginx",
        RepositoryOpts = new RepositoryOptsArgs
        {
            Repo = "https://charts.bitnami.com/bitnami",
        },
        ValueYamlFiles = { new FileAsset("./overrides.yaml") },
    });
});
```

{{% /choosable %}}

{{% choosable language java %}}

```java
package generated_program;

import com.pulumi.Pulumi;
import com.pulumi.asset.FileAsset;
import com.pulumi.kubernetes.helm.v3.Release;
import com.pulumi.kubernetes.helm.v3.ReleaseArgs;
import com.pulumi.kubernetes.helm.v3.inputs.RepositoryOptsArgs;

public class App {
    public static void main(String[] args) {
        Pulumi.run(ctx -> {
            new Release("nginx", ReleaseArgs.builder()
                    .chart("nginx")
                    .repositoryOpts(RepositoryOptsArgs.builder()
                            .repo("https://charts.bitnami.com/bitnami")
                            .build())
                    .valueYamlFiles(new FileAsset("./overrides.yaml"))
                    .build());
        });
    }
}
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: example
runtime: yaml
resources:
  nginx:
    type: kubernetes:helm.sh/v3:Release
    properties:
      chart: nginx
      repositoryOpts:
        repo: https://charts.bitnami.com/bitnami
      valueYamlFiles:
        - fn::fileAsset: overrides.yaml
```

{{% /choosable %}}

{{< /chooser >}}

On the next `pulumi up`, the `http` default from the chart's `values.yaml` is dropped and the `containerPorts` block is rendered without it.
