---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/aptible/aptible/0.9.18/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Aptible Provider
meta_desc: Provides an overview on how to configure the Pulumi Aptible provider.
layout: package
---

## Generate Provider

The Aptible provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider aptible/aptible
```
## Feature Support

For a detailed list of what features the provider supports see the
[Feature Support Matrix](https://www.aptible.com/docs/reference/interface-feature)
in the Aptible Documentation.
## Example Usage
### Authentication and Authorization

Authorization and Authentication is controlled using the same mechanism
that the [CLI](https://www.aptible.com/docs/reference/aptible-cli/overview) uses.
Therefore, you should log into the account you want to use Pulumi with using
the `aptible login` CLI command before running any Pulumi commands.

As another option the environment variables `APTIBLE_USERNAME` and
`APTIBLE_PASSWORD` can be set for the provider to use. In this case it is
strongly recommended that a robot account be used, especially as MFA needs to
be disabled for truly automated runs.
### Determining the Environment ID

Each resource managed via Pulumi requires an Environment ID specifying which
[Environment](https://www.aptible.com/docs/core-concepts/architecture/environments)
the resource should be created in. Currently the Aptible Deploy Pulumi
provider does not manage Environments, so you will need the Environment ID for
a pre-existing Environment. The easiest way to determine the Environment ID is
by using the Environment function. For example, if you have an Environment
with the handle "techco-test-environment" you can create the function:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aptible from "@pulumi/aptible";

const techco_test_environment = aptible.getEnvironment({
    handle: "techco-test-environment",
});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_aptible as aptible

techco_test_environment = aptible.get_environment(handle="techco-test-environment")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aptible = Pulumi.Aptible;

return await Deployment.RunAsync(() =>
{
    var techco_test_environment = Aptible.GetEnvironment.Invoke(new()
    {
        Handle = "techco-test-environment",
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/aptible/aptible"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aptible.LookupEnvironment(ctx, &aptible.LookupEnvironmentArgs{
			Handle: "techco-test-environment",
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
variables:
  techco-test-environment:
    fn::invoke:
      function: aptible:getEnvironment
      arguments:
        handle: techco-test-environment
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aptible.AptibleFunctions;
import com.pulumi.aptible.inputs.GetEnvironmentArgs;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        final var techco-test-environment = AptibleFunctions.getEnvironment(GetEnvironmentArgs.builder()
            .handle("techco-test-environment")
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

Once defined, you can use this function in your resource definitions.
For example, when defining an App:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language python %}}
```python
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language go %}}
```go
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language java %}}
```java
Example currently unavailable in this language
```
{{% /choosable %}}
{{< /chooser >}}

You can also specify the environment as a resource. You will need to pass in an
`orgId` (ex: `dashboard.aptible.com/organizations/<ORG_ID>/members`) and
`stackId` (ex: `dashboard.aptible.com/stack/<STACK_ID>/accounts`) which you can
get from the Aptible dashboard (on the settings/members panel (`orgId`) or on
the stack view pages (`stackId`)):

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aptible from "@pulumi/aptible";

const techco_test_environment = new aptible.Environment("techco-test-environment", {
    orgId: "some-uuid-that-represents-your-org",
    stackId: "your_stack_id",
    handle: "techco-test-environment",
});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_aptible as aptible

techco_test_environment = aptible.Environment("techco-test-environment",
    org_id="some-uuid-that-represents-your-org",
    stack_id="your_stack_id",
    handle="techco-test-environment")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aptible = Pulumi.Aptible;

return await Deployment.RunAsync(() =>
{
    var techco_test_environment = new Aptible.Environment("techco-test-environment", new()
    {
        OrgId = "some-uuid-that-represents-your-org",
        StackId = "your_stack_id",
        Handle = "techco-test-environment",
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/aptible/aptible"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aptible.NewEnvironment(ctx, "techco-test-environment", &aptible.EnvironmentArgs{
			OrgId:   pulumi.String("some-uuid-that-represents-your-org"),
			StackId: pulumi.Float64("your_stack_id"),
			Handle:  pulumi.String("techco-test-environment"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
resources:
  techco-test-environment:
    type: aptible:Environment
    properties:
      orgId: some-uuid-that-represents-your-org
      stackId: your_stack_id
      handle: techco-test-environment
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aptible.Environment;
import com.pulumi.aptible.EnvironmentArgs;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        var techco_test_environment = new Environment("techco-test-environment", EnvironmentArgs.builder()
            .orgId("some-uuid-that-represents-your-org")
            .stackId("your_stack_id")
            .handle("techco-test-environment")
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
### Apps

[Apps](https://www.aptible.com/docs/core-concepts/apps) can be
created using the `pulumiAptibleApp` resource.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language python %}}
```python
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language go %}}
```go
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language java %}}
```java
Example currently unavailable in this language
```
{{% /choosable %}}
{{< /chooser >}}
#### Configuring and Deploying Apps

!> Currently the only supported deployment method via Pulumi is of
Docker images hosted in a Docker image registry.

Apps configurations can be managed via the nested `config` element.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aptible from "@pulumi/aptible";

const APP = new aptible.App("APP", {
    envId: ENVIRONMENT_ID,
    handle: "APP_HANDLE",
    config: {
        KEY: "value",
    },
});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_aptible as aptible

app = aptible.App("APP",
    env_id=environmen_t__id,
    handle="APP_HANDLE",
    config={
        "KEY": "value",
    })
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aptible = Pulumi.Aptible;

return await Deployment.RunAsync(() =>
{
    var APP = new Aptible.App("APP", new()
    {
        EnvId = ENVIRONMENT_ID,
        Handle = "APP_HANDLE",
        Config =
        {
            { "KEY", "value" },
        },
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/aptible/aptible"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aptible.NewApp(ctx, "APP", &aptible.AppArgs{
			EnvId:  pulumi.Any(ENVIRONMENT_ID),
			Handle: pulumi.String("APP_HANDLE"),
			Config: pulumi.StringMap{
				"KEY": pulumi.String("value"),
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
{{% choosable language yaml %}}
```yaml
resources:
  APP:
    type: aptible:App
    properties:
      envId: ${ENVIRONMENT_ID}
      handle: APP_HANDLE
      config:
        KEY: value
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aptible.App;
import com.pulumi.aptible.AppArgs;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        var aPP = new App("APP", AppArgs.builder()
            .envId(ENVIRONMENT_ID)
            .handle("APP_HANDLE")
            .config(Map.of("KEY", "value"))
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

If you specify a Docker image as the `APTIBLE_DOCKER_IMAGE`
configuration value, that Docker image will be deployed to the App.
Authentication for Docker images located in
private repositories can be provided using the
`APTIBLE_PRIVATE_REGISTRY_USERNAME` and
`APTIBLE_PRIVATE_REGISTRY_PASSWORD` configuration values.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aptible from "@pulumi/aptible";

const APP = new aptible.App("APP", {
    envId: ENVIRONMENT_ID,
    handle: "APP_HANDLE",
    config: {
        KEY: "value",
        APTIBLE_DOCKER_IMAGE: "quay.io/aptible/deploy-demo-app",
        APTIBLE_PRIVATE_REGISTRY_USERNAME: "registry_username",
        APTIBLE_PRIVATE_REGISTRY_PASSWORD: "registry_password",
    },
});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_aptible as aptible

app = aptible.App("APP",
    env_id=environmen_t__id,
    handle="APP_HANDLE",
    config={
        "KEY": "value",
        "APTIBLE_DOCKER_IMAGE": "quay.io/aptible/deploy-demo-app",
        "APTIBLE_PRIVATE_REGISTRY_USERNAME": "registry_username",
        "APTIBLE_PRIVATE_REGISTRY_PASSWORD": "registry_password",
    })
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aptible = Pulumi.Aptible;

return await Deployment.RunAsync(() =>
{
    var APP = new Aptible.App("APP", new()
    {
        EnvId = ENVIRONMENT_ID,
        Handle = "APP_HANDLE",
        Config =
        {
            { "KEY", "value" },
            { "APTIBLE_DOCKER_IMAGE", "quay.io/aptible/deploy-demo-app" },
            { "APTIBLE_PRIVATE_REGISTRY_USERNAME", "registry_username" },
            { "APTIBLE_PRIVATE_REGISTRY_PASSWORD", "registry_password" },
        },
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/aptible/aptible"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aptible.NewApp(ctx, "APP", &aptible.AppArgs{
			EnvId:  pulumi.Any(ENVIRONMENT_ID),
			Handle: pulumi.String("APP_HANDLE"),
			Config: pulumi.StringMap{
				"KEY":                               pulumi.String("value"),
				"APTIBLE_DOCKER_IMAGE":              pulumi.String("quay.io/aptible/deploy-demo-app"),
				"APTIBLE_PRIVATE_REGISTRY_USERNAME": pulumi.String("registry_username"),
				"APTIBLE_PRIVATE_REGISTRY_PASSWORD": pulumi.String("registry_password"),
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
{{% choosable language yaml %}}
```yaml
resources:
  APP:
    type: aptible:App
    properties:
      envId: ${ENVIRONMENT_ID}
      handle: APP_HANDLE
      config:
        KEY: value
        APTIBLE_DOCKER_IMAGE: quay.io/aptible/deploy-demo-app
        APTIBLE_PRIVATE_REGISTRY_USERNAME: registry_username
        APTIBLE_PRIVATE_REGISTRY_PASSWORD: registry_password
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aptible.App;
import com.pulumi.aptible.AppArgs;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        var aPP = new App("APP", AppArgs.builder()
            .envId(ENVIRONMENT_ID)
            .handle("APP_HANDLE")
            .config(Map.ofEntries(
                Map.entry("KEY", "value"),
                Map.entry("APTIBLE_DOCKER_IMAGE", "quay.io/aptible/deploy-demo-app"),
                Map.entry("APTIBLE_PRIVATE_REGISTRY_USERNAME", "registry_username"),
                Map.entry("APTIBLE_PRIVATE_REGISTRY_PASSWORD", "registry_password")
            ))
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
#### Scaling Services

Each App is comprised of one or more
[Services](https://www.aptible.com/docs/core-concepts/apps/deploying-apps/services).
These Services must be defined in the
[Procfile](https://www.aptible.com/docs/how-to-guides/app-guides/define-services#explicit-services-procfiles)
for your App.

Services can be scaled independently both in terms of the number of running
[containers](https://www.aptible.com/docs/core-concepts/architecture/containers/overview)
and size of the running Containers. This is done using the nested `service`
element for the App resource:

> The `processType` in the `service` element maps directly to the
Service name used in the Procfile. If you are not using a Procfile,
you will have a single Service with the `processType` of `cmd`

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aptible from "@pulumi/aptible";

const APP = new aptible.App("APP", {
    envId: ENVIRONMENT_ID,
    handle: "APP_HANDLE",
    services: [
        {
            processType: "SERVICE_NAME1",
            containerCount: 1,
            containerMemoryLimit: 1024,
        },
        {
            processType: "SERVICE_NAME2",
            containerCount: 2,
            containerMemoryLimit: 2048,
        },
    ],
});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_aptible as aptible

app = aptible.App("APP",
    env_id=environmen_t__id,
    handle="APP_HANDLE",
    services=[
        {
            "process_type": "SERVICE_NAME1",
            "container_count": 1,
            "container_memory_limit": 1024,
        },
        {
            "process_type": "SERVICE_NAME2",
            "container_count": 2,
            "container_memory_limit": 2048,
        },
    ])
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aptible = Pulumi.Aptible;

return await Deployment.RunAsync(() =>
{
    var APP = new Aptible.App("APP", new()
    {
        EnvId = ENVIRONMENT_ID,
        Handle = "APP_HANDLE",
        Services = new[]
        {
            new Aptible.Inputs.AppServiceArgs
            {
                ProcessType = "SERVICE_NAME1",
                ContainerCount = 1,
                ContainerMemoryLimit = 1024,
            },
            new Aptible.Inputs.AppServiceArgs
            {
                ProcessType = "SERVICE_NAME2",
                ContainerCount = 2,
                ContainerMemoryLimit = 2048,
            },
        },
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/aptible/aptible"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aptible.NewApp(ctx, "APP", &aptible.AppArgs{
			EnvId:  pulumi.Any(ENVIRONMENT_ID),
			Handle: pulumi.String("APP_HANDLE"),
			Services: aptible.AppServiceArray{
				&aptible.AppServiceArgs{
					ProcessType:          pulumi.String("SERVICE_NAME1"),
					ContainerCount:       pulumi.Float64(1),
					ContainerMemoryLimit: pulumi.Float64(1024),
				},
				&aptible.AppServiceArgs{
					ProcessType:          pulumi.String("SERVICE_NAME2"),
					ContainerCount:       pulumi.Float64(2),
					ContainerMemoryLimit: pulumi.Float64(2048),
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
{{% choosable language yaml %}}
```yaml
resources:
  APP:
    type: aptible:App
    properties:
      envId: ${ENVIRONMENT_ID}
      handle: APP_HANDLE
      services:
        - processType: SERVICE_NAME1
          containerCount: 1
          containerMemoryLimit: 1024
        - processType: SERVICE_NAME2
          containerCount: 2
          containerMemoryLimit: 2048
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aptible.App;
import com.pulumi.aptible.AppArgs;
import com.pulumi.aptible.inputs.AppServiceArgs;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        var aPP = new App("APP", AppArgs.builder()
            .envId(ENVIRONMENT_ID)
            .handle("APP_HANDLE")
            .services(
                AppServiceArgs.builder()
                    .processType("SERVICE_NAME1")
                    .containerCount(1)
                    .containerMemoryLimit(1024)
                    .build(),
                AppServiceArgs.builder()
                    .processType("SERVICE_NAME2")
                    .containerCount(2)
                    .containerMemoryLimit(2048)
                    .build())
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
#### Autoscaling

Services can be configured to [Autoscale](https://www.aptible.com/docs/core-concepts/scaling/app-scaling)
either vertically or horizontally. The configuration is per service

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aptible from "@pulumi/aptible";

const APP = new aptible.App("APP", {
    envId: ENVIRONMENT_ID,
    handle: "APP_HANDLE",
    services: [
        {
            processType: "SERVICE_NAME1",
            containerCount: 1,
            containerMemoryLimit: 1024,
            autoscalingPolicies: [{
                autoscalingType: "horizontal",
                minContainers: 2,
                maxContainers: 5,
                minCpuThreshold: 0.4,
                maxCpuThreshold: 0.8,
            }],
        },
        {
            processType: "SERVICE_NAME2",
            containerCount: 2,
            containerMemoryLimit: 2048,
            autoscalingPolicies: [{
                autoscalingType: "vertical",
                memScaleUpThreshold: 0.8,
            }],
        },
    ],
});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_aptible as aptible

app = aptible.App("APP",
    env_id=environmen_t__id,
    handle="APP_HANDLE",
    services=[
        {
            "process_type": "SERVICE_NAME1",
            "container_count": 1,
            "container_memory_limit": 1024,
            "autoscaling_policies": [{
                "autoscaling_type": "horizontal",
                "min_containers": 2,
                "max_containers": 5,
                "min_cpu_threshold": 0.4,
                "max_cpu_threshold": 0.8,
            }],
        },
        {
            "process_type": "SERVICE_NAME2",
            "container_count": 2,
            "container_memory_limit": 2048,
            "autoscaling_policies": [{
                "autoscaling_type": "vertical",
                "mem_scale_up_threshold": 0.8,
            }],
        },
    ])
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aptible = Pulumi.Aptible;

return await Deployment.RunAsync(() =>
{
    var APP = new Aptible.App("APP", new()
    {
        EnvId = ENVIRONMENT_ID,
        Handle = "APP_HANDLE",
        Services = new[]
        {
            new Aptible.Inputs.AppServiceArgs
            {
                ProcessType = "SERVICE_NAME1",
                ContainerCount = 1,
                ContainerMemoryLimit = 1024,
                AutoscalingPolicies = new[]
                {
                    new Aptible.Inputs.AppServiceAutoscalingPolicyArgs
                    {
                        AutoscalingType = "horizontal",
                        MinContainers = 2,
                        MaxContainers = 5,
                        MinCpuThreshold = 0.4,
                        MaxCpuThreshold = 0.8,
                    },
                },
            },
            new Aptible.Inputs.AppServiceArgs
            {
                ProcessType = "SERVICE_NAME2",
                ContainerCount = 2,
                ContainerMemoryLimit = 2048,
                AutoscalingPolicies = new[]
                {
                    new Aptible.Inputs.AppServiceAutoscalingPolicyArgs
                    {
                        AutoscalingType = "vertical",
                        MemScaleUpThreshold = 0.8,
                    },
                },
            },
        },
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/aptible/aptible"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aptible.NewApp(ctx, "APP", &aptible.AppArgs{
			EnvId:  pulumi.Any(ENVIRONMENT_ID),
			Handle: pulumi.String("APP_HANDLE"),
			Services: aptible.AppServiceArray{
				&aptible.AppServiceArgs{
					ProcessType:          pulumi.String("SERVICE_NAME1"),
					ContainerCount:       pulumi.Float64(1),
					ContainerMemoryLimit: pulumi.Float64(1024),
					AutoscalingPolicies: aptible.AppServiceAutoscalingPolicyArray{
						&aptible.AppServiceAutoscalingPolicyArgs{
							AutoscalingType: pulumi.String("horizontal"),
							MinContainers:   pulumi.Float64(2),
							MaxContainers:   pulumi.Float64(5),
							MinCpuThreshold: pulumi.Float64(0.4),
							MaxCpuThreshold: pulumi.Float64(0.8),
						},
					},
				},
				&aptible.AppServiceArgs{
					ProcessType:          pulumi.String("SERVICE_NAME2"),
					ContainerCount:       pulumi.Float64(2),
					ContainerMemoryLimit: pulumi.Float64(2048),
					AutoscalingPolicies: aptible.AppServiceAutoscalingPolicyArray{
						&aptible.AppServiceAutoscalingPolicyArgs{
							AutoscalingType:     pulumi.String("vertical"),
							MemScaleUpThreshold: pulumi.Float64(0.8),
						},
					},
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
{{% choosable language yaml %}}
```yaml
resources:
  APP:
    type: aptible:App
    properties:
      envId: ${ENVIRONMENT_ID}
      handle: APP_HANDLE
      services:
        - processType: SERVICE_NAME1
          containerCount: 1
          containerMemoryLimit: 1024
          autoscalingPolicies:
            - autoscalingType: horizontal
              minContainers: 2
              maxContainers: 5
              minCpuThreshold: 0.4
              maxCpuThreshold: 0.8
        - processType: SERVICE_NAME2
          containerCount: 2
          containerMemoryLimit: 2048
          autoscalingPolicies:
            - autoscalingType: vertical
              memScaleUpThreshold: 0.8
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aptible.App;
import com.pulumi.aptible.AppArgs;
import com.pulumi.aptible.inputs.AppServiceArgs;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        var aPP = new App("APP", AppArgs.builder()
            .envId(ENVIRONMENT_ID)
            .handle("APP_HANDLE")
            .services(
                AppServiceArgs.builder()
                    .processType("SERVICE_NAME1")
                    .containerCount(1)
                    .containerMemoryLimit(1024)
                    .autoscalingPolicies(AppServiceAutoscalingPolicyArgs.builder()
                        .autoscalingType("horizontal")
                        .minContainers(2)
                        .maxContainers(5)
                        .minCpuThreshold(0.4)
                        .maxCpuThreshold(0.8)
                        .build())
                    .build(),
                AppServiceArgs.builder()
                    .processType("SERVICE_NAME2")
                    .containerCount(2)
                    .containerMemoryLimit(2048)
                    .autoscalingPolicies(AppServiceAutoscalingPolicyArgs.builder()
                        .autoscalingType("vertical")
                        .memScaleUpThreshold(0.8)
                        .build())
                    .build())
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
### Endpoints

Endpoints for [Apps](https://www.aptible.com/docs/core-concepts/apps) and
[Databases](https://www.aptible.com/docs/core-concepts/managed-databases) can be
managed using the `pulumiAptibleEndpoint` resource.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aptible from "@pulumi/aptible";

const EXAMPLE = new aptible.Endpoint("EXAMPLE", {
    envId: ENVIRONMENT_ID,
    processType: "SERVICE_NAME",
    resourceId: APP.appId,
    defaultDomain: true,
    endpointType: "https",
    internal: false,
    platform: "alb",
    containerPort: 5000,
});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_aptible as aptible

example = aptible.Endpoint("EXAMPLE",
    env_id=environmen_t__id,
    process_type="SERVICE_NAME",
    resource_id=app["appId"],
    default_domain=True,
    endpoint_type="https",
    internal=False,
    platform="alb",
    container_port=5000)
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aptible = Pulumi.Aptible;

return await Deployment.RunAsync(() =>
{
    var EXAMPLE = new Aptible.Endpoint("EXAMPLE", new()
    {
        EnvId = ENVIRONMENT_ID,
        ProcessType = "SERVICE_NAME",
        ResourceId = APP.AppId,
        DefaultDomain = true,
        EndpointType = "https",
        Internal = false,
        Platform = "alb",
        ContainerPort = 5000,
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/aptible/aptible"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aptible.NewEndpoint(ctx, "EXAMPLE", &aptible.EndpointArgs{
			EnvId:         pulumi.Any(ENVIRONMENT_ID),
			ProcessType:   pulumi.String("SERVICE_NAME"),
			ResourceId:    pulumi.Any(APP.AppId),
			DefaultDomain: pulumi.Bool(true),
			EndpointType:  pulumi.String("https"),
			Internal:      pulumi.Bool(false),
			Platform:      pulumi.String("alb"),
			ContainerPort: pulumi.Float64(5000),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
resources:
  EXAMPLE:
    type: aptible:Endpoint
    properties:
      envId: ${ENVIRONMENT_ID}
      processType: SERVICE_NAME
      resourceId: ${APP.appId}
      defaultDomain: true
      endpointType: https
      internal: false
      platform: alb
      containerPort: 5000
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aptible.Endpoint;
import com.pulumi.aptible.EndpointArgs;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        var eXAMPLE = new Endpoint("EXAMPLE", EndpointArgs.builder()
            .envId(ENVIRONMENT_ID)
            .processType("SERVICE_NAME")
            .resourceId(APP.appId())
            .defaultDomain(true)
            .endpointType("https")
            .internal(false)
            .platform("alb")
            .containerPort(5000)
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
### Databases

[Databases](https://www.aptible.com/docs/core-concepts/managed-databases) can be
managed using the `pulumiAptibleDatabase` resource.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aptible from "@pulumi/aptible";

const DATABASE = new aptible.Database("DATABASE", {
    envId: ENVIRONMENT_ID,
    handle: "DATABASE_HANDLE",
    databaseType: "redis",
    containerSize: 512,
    diskSize: 10,
});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_aptible as aptible

database = aptible.Database("DATABASE",
    env_id=environmen_t__id,
    handle="DATABASE_HANDLE",
    database_type="redis",
    container_size=512,
    disk_size=10)
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aptible = Pulumi.Aptible;

return await Deployment.RunAsync(() =>
{
    var DATABASE = new Aptible.Database("DATABASE", new()
    {
        EnvId = ENVIRONMENT_ID,
        Handle = "DATABASE_HANDLE",
        DatabaseType = "redis",
        ContainerSize = 512,
        DiskSize = 10,
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/aptible/aptible"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aptible.NewDatabase(ctx, "DATABASE", &aptible.DatabaseArgs{
			EnvId:         pulumi.Any(ENVIRONMENT_ID),
			Handle:        pulumi.String("DATABASE_HANDLE"),
			DatabaseType:  pulumi.String("redis"),
			ContainerSize: pulumi.Float64(512),
			DiskSize:      pulumi.Float64(10),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
resources:
  DATABASE:
    type: aptible:Database
    properties:
      envId: ${ENVIRONMENT_ID}
      handle: DATABASE_HANDLE
      databaseType: redis
      containerSize: 512
      diskSize: 10
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aptible.Database;
import com.pulumi.aptible.DatabaseArgs;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        var dATABASE = new Database("DATABASE", DatabaseArgs.builder()
            .envId(ENVIRONMENT_ID)
            .handle("DATABASE_HANDLE")
            .databaseType("redis")
            .containerSize(512)
            .diskSize(10)
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
#### Replication

Database [Replicas and
Clusters](https://www.aptible.com/docs/core-concepts/managed-databases/managing-databases/replication-clustering)
can be created using the `pulumiAptibleReplica` resource.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aptible from "@pulumi/aptible";

const REPLICA_HANDLE = new aptible.Replica("REPLICA_HANDLE", {
    envId: ENVIRONMENT_ID,
    primaryDatabaseId: DATABASE.databaseId,
    handle: "REPLICA_HANDLE",
    diskSize: 30,
});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_aptible as aptible

replic_a__handle = aptible.Replica("REPLICA_HANDLE",
    env_id=environmen_t__id,
    primary_database_id=database["databaseId"],
    handle="REPLICA_HANDLE",
    disk_size=30)
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aptible = Pulumi.Aptible;

return await Deployment.RunAsync(() =>
{
    var REPLICA_HANDLE = new Aptible.Replica("REPLICA_HANDLE", new()
    {
        EnvId = ENVIRONMENT_ID,
        PrimaryDatabaseId = DATABASE.DatabaseId,
        Handle = "REPLICA_HANDLE",
        DiskSize = 30,
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/aptible/aptible"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aptible.NewReplica(ctx, "REPLICA_HANDLE", &aptible.ReplicaArgs{
			EnvId:             pulumi.Any(ENVIRONMENT_ID),
			PrimaryDatabaseId: pulumi.Any(DATABASE.DatabaseId),
			Handle:            pulumi.String("REPLICA_HANDLE"),
			DiskSize:          pulumi.Float64(30),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
resources:
  REPLICA_HANDLE:
    type: aptible:Replica
    properties:
      envId: ${ENVIRONMENT_ID}
      primaryDatabaseId: ${DATABASE.databaseId}
      handle: REPLICA_HANDLE
      diskSize: 30
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aptible.Replica;
import com.pulumi.aptible.ReplicaArgs;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        var rEPLICA_HANDLE = new Replica("REPLICA_HANDLE", ReplicaArgs.builder()
            .envId(ENVIRONMENT_ID)
            .primaryDatabaseId(DATABASE.databaseId())
            .handle("REPLICA_HANDLE")
            .diskSize(30)
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

There are currently no arguments to provide directly to the provider