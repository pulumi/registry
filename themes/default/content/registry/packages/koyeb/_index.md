---
# WARNING: this file was fetched from https://raw.githubusercontent.com/koyeb/pulumi-koyeb/v0.1.11/docs/_index.md
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

const koyebApp = new koyeb.App("sample-app", {
  name: "sample-app",
});

const koyebService = new koyeb.Service("sample-service", {
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
		app, err := koyeb.App(ctx, "sample-app", &koyeb.AppArgs{
			Name: pulumi.String("sample-app"),
		})

		if err != nil {
			return err
		}

		_, err = koyeb.NewKoyebService(ctx, "sample-service", &koyeb.ServiceArgs{
			AppName: app.Name,
			Definition: &koyeb.ServiceDefinitionArgs{
				Name: pulumi.String("sample-service"),
				Regions: pulumi.StringArray{
					pulumi.String("fra"),
				},
				Git: &koyeb.ServiceDefinitionGitArgs{
					Repository: pulumi.String("github.com/koyeb/example-golang"),
					Branch:     pulumi.String("main"),
				},
				InstanceTypes: &koyeb.ServiceDefinitionInstanceTypesArgs{
					Type: pulumi.String("micro"),
				},
				Routes: koyeb.ServiceDefinitionRouteArray{
					&koyeb.ServiceDefinitionRouteArgs{
						Path: pulumi.String("/"),
						Port: pulumi.Int(8080),
					},
				},
				Ports: koyeb.ServiceDefinitionPortArray{
					&koyeb.ServiceDefinitionPortArgs{
						Port:     pulumi.Int(8080),
						Protocol: pulumi.String("http"),
					},
				},
				Scalings: &koyeb.ServiceDefinitionScalingsArgs{
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

app = koyeb.App("sample-app", name="sample-app")

koyeb.Service("sample-service", app_name=app.name,  definition=koyeb.ServiceDefinitionArgs(
    name="sample-service",
    regions=["fra"],
    git=koyeb.ServiceDefinitionGitArgs(
        repository="github.com/koyeb/example-golang",
        branch="main",
    ),
    instance_types=koyeb.ServiceDefinitionInstanceTypesArgs(
        type="micro",
    ),
    routes=[koyeb.ServiceDefinitionRouteArgs(
        path="/",
        port=8080,
    )],
    ports=[koyeb.ServiceDefinitionPortArgs(
        port=8080,
        protocol="http",
    )],
    scalings=koyeb.ServiceDefinitionScalingsArgs(
        min=1,
        max=1,
    ),
))
```

{{% /choosable %}}
{{< /chooser >}}
