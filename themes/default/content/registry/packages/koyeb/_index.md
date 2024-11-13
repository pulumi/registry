---
# WARNING: this file was fetched from https://raw.githubusercontent.com/koyeb/pulumi-koyeb/v0.1.7/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Koyeb
meta_desc: Provides an overview of the Koyeb Provider for Pulumi.
layout: overview
---

The Koyeb provider for Pulumi can be used to provision any of the cloud resources available in [Koyeb](https://www.koyeb.com).

To be able to manage Koyeb resources using Pulumi, you first need a Koyeb API access token to allow Pulumi to authenticate to Koyeb. Koyeb API access token credentials can be obtained from the [Koyeb Control Panel](https://app.koyeb.com/account/api).

## Example

{{< chooser language "typescript,go,python" >}}
{{% choosable language typescript %}}

```typescript
import * as koyeb from "@koyeb/pulumi-koyeb";

const koyebApp = new koyeb.KoyebApp("sample-app", {
  name: "sample-app",
});

const koyebService = new koyeb.KoyebService("sample-service", {
  appName: koyebApp.name,
  definition: {
    name: "sample-service",
    regions: ["fra"],
    git: {
      repository: "github.com/koyeb/example-golang",
      branch: "main",
    },
    instanceTypes: {
      type: "micro",
    },
    routes: [
      {
        path: "/",
        port: 8080,
      },
    ],
    ports: [
      {
        port: 8080,
        protocol: "http",
      },
    ],
    scalings: {
      min: 1,
      max: 1,
    },
  },
});
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/koyeb/pulumi-koyeb/sdk/go/koyeb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		app, err := koyeb.NewKoyebApp(ctx, "sample-app", &koyeb.KoyebAppArgs{
			Name: pulumi.String("sample-app"),
		})

		if err != nil {
			return err
		}

		_, err = koyeb.NewKoyebService(ctx, "sample-service", &koyeb.KoyebServiceArgs{
			AppName: app.Name,
			Definition: &koyeb.KoyebServiceDefinitionArgs{
				Name: pulumi.String("sample-service"),
				Regions: pulumi.StringArray{
					pulumi.String("fra"),
				},
				Git: &koyeb.KoyebServiceDefinitionGitArgs{
					Repository: pulumi.String("github.com/koyeb/example-golang"),
					Branch:     pulumi.String("main"),
				},
				InstanceTypes: &koyeb.KoyebServiceDefinitionInstanceTypesArgs{
					Type: pulumi.String("micro"),
				},
				Routes: koyeb.KoyebServiceDefinitionRouteArray{
					&koyeb.KoyebServiceDefinitionRouteArgs{
						Path: pulumi.String("/"),
						Port: pulumi.Int(8080),
					},
				},
				Ports: koyeb.KoyebServiceDefinitionPortArray{
					&koyeb.KoyebServiceDefinitionPortArgs{
						Port:     pulumi.Int(8080),
						Protocol: pulumi.String("http"),
					},
				},
				Scalings: &koyeb.KoyebServiceDefinitionScalingsArgs{
					Min: pulumi.Int(1),
					Max: pulumi.Int(1),
				},
			},
		})

		if err != nil {
			return err
		}

		return nil
	})
}
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_koyeb as koyeb

app = koyeb.KoyebApp("sample-app", name="sample-app")

koyeb.KoyebService("sample-service", app_name=app.name,  definition=koyeb.KoyebServiceDefinitionArgs(
    name="sample-service",
    regions=["fra"],
    git=koyeb.KoyebServiceDefinitionGitArgs(
        repository="github.com/koyeb/example-golang",
        branch="main",
    ),
    instance_types=koyeb.KoyebServiceDefinitionInstanceTypesArgs(
        type="micro",
    ),
    routes=[koyeb.KoyebServiceDefinitionRouteArgs(
        path="/",
        port=8080,
    )],
    ports=[koyeb.KoyebServiceDefinitionPortArgs(
        port=8080,
        protocol="http",
    )],
    scalings=koyeb.KoyebServiceDefinitionScalingsArgs(
        min=1,
        max=1,
    ),
))
```

{{% /choosable %}}
{{< /chooser >}}