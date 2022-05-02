---
title: Docker
meta_desc: Provides an overview of the Docker Provider for Pulumi.
layout: overview
---

The Docker provider for Pulumi can be used to provision any of the resources available in [Docker](https://www.docker.com/).

## Example

{{< chooser language "javascript,typescript,python,go,csharp,yaml" >}}

{{% choosable language javascript %}}

```javascript
const docker = require("@pulumi/docker")

const image = new docker.RemoteImage("ubuntu", {
  name: "ubuntu:precise"
});

const container = new docker.Container("ubuntu", {
  image: image.latest,
});
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as docker from "@pulumi/docker";

const image = new docker.RemoteImage("ubuntu", {
  name: "ubuntu:precise"
});

const container = new docker.Container("ubuntu", {
  image: image.latest,
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_docker as docker

image = docker.RemoteImage("ubuntu",
  name='ubuntu:precise'
)

container = docker.Container("ubuntu",
  image=image.latest
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	do "github.com/pulumi/pulumi-docker/sdk/v3/go/docker"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		image, err := do.NewRemoteImage(ctx, "ubuntu", &docker.RemoteImageArgs{
			Name: pulumi.String("ubuntu:precise"),
		})
		if err != nil {
			return err
		}

		container, err := do.NewContainer(ctx, "ubuntu", &docker.ContainerArgs{
			Image: image.Latest(),
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
using Pulumi.Docker;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            var image = new Docker.RemoteImage("ubuntu", new Docker.RemoteImageArgs
            {
                Name = "ubuntu:precise",
            });

            var container = new Docker.Container("ubuntu", new Docker.ContainerArgs
            {
                Image = image.Latest,
            });
        });
}
```

{{% /choosable %}}
{{% choosable language yaml %}}

```yaml
resources:
  image:
    type: docker:RemoteImage
    properties:
      name: ubuntu:precise
  container:
    type: docker:Container
    properties:
      image: ${image.latest}
```

{{% /choosable %}}

{{< /chooser >}}
