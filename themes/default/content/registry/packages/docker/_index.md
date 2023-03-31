---
title: Docker
meta_desc: Provides an overview of the Docker Provider for Pulumi.
layout: overview
---

The Pulumi Docker Provider can be used to provision and interact with any of the resources available in [Docker](https://www.docker.com/) including containers, images, networks, volumes and more.

## Example

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}


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

await Deployment.RunAsync(() =>
{
  var image = new Docker.RemoteImage("ubuntu", new Docker.RemoteImageArgs
  {
    Name = "ubuntu:precise",
  });

  var container = new Docker.Container("ubuntu", new Docker.ContainerArgs
  {
    Image = image.Latest,
  });
});
```

{{% /choosable %}}
{{% choosable language java %}}

```java
import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.docker.RemoteImage;
import com.pulumi.docker.RemoteImageArgs;
import com.pulumi.docker.Container;
import com.pulumi.docker.ContainerArgs;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        var ubuntuRemoteImage = new RemoteImage("ubuntuRemoteImage", RemoteImageArgs.builder()
                .name("ubuntu:precise")
                .build());

        var ubuntuContainer = new Container("ubuntuContainer", ContainerArgs.builder()
                .image(ubuntuRemoteImage.imageId())
                .build());

    }
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
