---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/juju/juju/1.0.0-beta3/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Juju Provider
meta_desc: Provides an overview on how to configure the Pulumi Juju provider.
layout: package
---

## Generate Provider

The Juju provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider juju/juju
```
## Overview

The provider can be used to interact with [Juju](https://juju.is "Juju | An open source application orchestration engine") - an open source orchestration engine by Canonical.
Additionally, the provider supports interactions with [JAAS](https://documentation.ubuntu.com/jaas/ "JAAS | An enterprise gateway into your Juju estate") - an orchestrator of Juju controllers.

The provider only interacts with a single controller at a time.

Today this provider allows you to manage the following via resources:

* Applications and deploy charms
* Credentials for existing clouds
* Integrations
* Machines
* Models
* Model ssh keys
* Offers
* Users

and refer to the following via functions:

* Applications
* Machines
* Models
* Offers

Work is ongoing to include support for more of the juju CLIs capabilities within this provider.
## Prerequisites

* [Juju](https://juju.is "Juju | An open source application orchestration engine") `2.9.49+`
## Authentication

There are 3 ways to define credentials for authentication with the Juju controller you wish to target.
They are displayed in the order in which the provider looks for credentials.
### Static credentials

Define the Juju controller credentials in the provider definition in your pulumi preview.

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    juju:caCertificate:
        value: 'TODO: file("~/ca-cert.pem")'
    juju:controllerAddresses:
        value: 10.225.205.241:17070,10.225.205.242:17070,[fd42:791:fa5e:6834:216:3eff:fe7a:8e6a]:17070
    juju:password:
        value: password1
    juju:username:
        value: jujuuser

```
### Client credentials

Note: Authentication with client credentials is only supported when communicating with JAAS.

Define the client credentials in the provider definition in your pulumi preview.

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    juju:caCertificate:
        value: 'TODO: file("~/ca-cert.pem")'
    juju:clientId:
        value: jujuclientid
    juju:clientSecret:
        value: jujuclientsecret
    juju:controllerAddresses:
        value: 10.225.205.241:17070,10.225.205.242:17070,[fd42:791:fa5e:6834:216:3eff:fe7a:8e6a]:17070

```
### Environment variables

Define the Juju controller credentials in the provider definition via environment variables. These can be set up as follows:

```shell
export CONTROLLER=$(juju whoami | yq .Controller)
export JUJU_CONTROLLER_ADDRESSES=$(juju show-controller | yq .$CONTROLLER.details.api-endpoints | yq -r '. | join(",")')
export JUJU_USERNAME="$(cat ~/.local/share/juju/accounts.yaml | yq .controllers.$CONTROLLER.user|tr -d '"')"
export JUJU_PASSWORD="$(cat ~/.local/share/juju/accounts.yaml | yq .controllers.$CONTROLLER.password|tr -d '"')"
export JUJU_CA_CERT="$(juju show-controller $(echo $CONTROLLER|tr -d '"') | yq '.[$CONTROLLER]'.details.\"ca-cert\"|tr -d '"'|sed 's/\\n/\n/g')"
```
### Populated by the provider via the juju CLI client.

This is the most straightforward solution. Remember that it will use the configuration used by the Juju CLI client at that moment. The fields are populated using the
output from running the command `juju show-controller` with the `--show-password` flag.
## Example Usage


{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as juju from "@pulumi/juju";

const development = new juju.Model("development", {
    name: "development",
    clouds: [{
        name: "aws",
        region: "eu-west-1",
    }],
});
const wordpress = new juju.Application("wordpress", {
    name: "wordpress",
    modelUuid: development.uuid,
    charms: [{
        name: "wordpress",
    }],
    units: 3,
});
const percona_cluster = new juju.Application("percona-cluster", {
    name: "percona-cluster",
    modelUuid: development.uuid,
    charms: [{
        name: "percona-cluster",
    }],
    units: 3,
});
const wpToPercona = new juju.Integration("wp_to_percona", {
    modelUuid: development.uuid,
    applications: [
        {
            name: wordpress.name,
            endpoint: "db",
        },
        {
            name: percona_cluster.name,
            endpoint: "server",
        },
    ],
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python

```
```python
import pulumi
import pulumi_juju as juju

development = juju.Model("development",
    name="development",
    clouds=[{
        "name": "aws",
        "region": "eu-west-1",
    }])
wordpress = juju.Application("wordpress",
    name="wordpress",
    model_uuid=development.uuid,
    charms=[{
        "name": "wordpress",
    }],
    units=3)
percona_cluster = juju.Application("percona-cluster",
    name="percona-cluster",
    model_uuid=development.uuid,
    charms=[{
        "name": "percona-cluster",
    }],
    units=3)
wp_to_percona = juju.Integration("wp_to_percona",
    model_uuid=development.uuid,
    applications=[
        {
            "name": wordpress.name,
            "endpoint": "db",
        },
        {
            "name": percona_cluster.name,
            "endpoint": "server",
        },
    ])
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Juju = Pulumi.Juju;

return await Deployment.RunAsync(() =>
{
    var development = new Juju.Model("development", new()
    {
        Name = "development",
        Clouds = new[]
        {
            new Juju.Inputs.ModelCloudArgs
            {
                Name = "aws",
                Region = "eu-west-1",
            },
        },
    });

    var wordpress = new Juju.Application("wordpress", new()
    {
        Name = "wordpress",
        ModelUuid = development.Uuid,
        Charms = new[]
        {
            new Juju.Inputs.ApplicationCharmArgs
            {
                Name = "wordpress",
            },
        },
        Units = 3,
    });

    var percona_cluster = new Juju.Application("percona-cluster", new()
    {
        Name = "percona-cluster",
        ModelUuid = development.Uuid,
        Charms = new[]
        {
            new Juju.Inputs.ApplicationCharmArgs
            {
                Name = "percona-cluster",
            },
        },
        Units = 3,
    });

    var wpToPercona = new Juju.Integration("wp_to_percona", new()
    {
        ModelUuid = development.Uuid,
        Applications = new[]
        {
            new Juju.Inputs.IntegrationApplicationArgs
            {
                Name = wordpress.Name,
                Endpoint = "db",
            },
            new Juju.Inputs.IntegrationApplicationArgs
            {
                Name = percona_cluster.Name,
                Endpoint = "server",
            },
        },
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/juju/juju"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		development, err := juju.NewModel(ctx, "development", &juju.ModelArgs{
			Name: pulumi.String("development"),
			Clouds: juju.ModelCloudArray{
				&juju.ModelCloudArgs{
					Name:   pulumi.String("aws"),
					Region: pulumi.String("eu-west-1"),
				},
			},
		})
		if err != nil {
			return err
		}
		wordpress, err := juju.NewApplication(ctx, "wordpress", &juju.ApplicationArgs{
			Name:      pulumi.String("wordpress"),
			ModelUuid: development.Uuid,
			Charms: juju.ApplicationCharmArray{
				&juju.ApplicationCharmArgs{
					Name: pulumi.String("wordpress"),
				},
			},
			Units: pulumi.Float64(3),
		})
		if err != nil {
			return err
		}
		percona_cluster, err := juju.NewApplication(ctx, "percona-cluster", &juju.ApplicationArgs{
			Name:      pulumi.String("percona-cluster"),
			ModelUuid: development.Uuid,
			Charms: juju.ApplicationCharmArray{
				&juju.ApplicationCharmArgs{
					Name: pulumi.String("percona-cluster"),
				},
			},
			Units: pulumi.Float64(3),
		})
		if err != nil {
			return err
		}
		_, err = juju.NewIntegration(ctx, "wp_to_percona", &juju.IntegrationArgs{
			ModelUuid: development.Uuid,
			Applications: juju.IntegrationApplicationArray{
				&juju.IntegrationApplicationArgs{
					Name:     wordpress.Name,
					Endpoint: pulumi.String("db"),
				},
				&juju.IntegrationApplicationArgs{
					Name:     percona_cluster.Name,
					Endpoint: pulumi.String("server"),
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
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml

```
```yaml
resources:
  development:
    type: juju:Model
    properties:
      name: development
      clouds:
        - name: aws
          region: eu-west-1
  wordpress:
    type: juju:Application
    properties:
      name: wordpress
      modelUuid: ${development.uuid}
      charms:
        - name: wordpress
      units: 3
  percona-cluster:
    type: juju:Application
    properties:
      name: percona-cluster
      modelUuid: ${development.uuid}
      charms:
        - name: percona-cluster
      units: 3
  wpToPercona:
    type: juju:Integration
    name: wp_to_percona
    properties:
      modelUuid: ${development.uuid}
      applications:
        - name: ${wordpress.name}
          endpoint: db
        - name: ${["percona-cluster"].name}
          endpoint: server
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.juju.Model;
import com.pulumi.juju.ModelArgs;
import com.pulumi.juju.inputs.ModelCloudArgs;
import com.pulumi.juju.Application;
import com.pulumi.juju.ApplicationArgs;
import com.pulumi.juju.inputs.ApplicationCharmArgs;
import com.pulumi.juju.Integration;
import com.pulumi.juju.IntegrationArgs;
import com.pulumi.juju.inputs.IntegrationApplicationArgs;
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
        var development = new Model("development", ModelArgs.builder()
            .name("development")
            .clouds(ModelCloudArgs.builder()
                .name("aws")
                .region("eu-west-1")
                .build())
            .build());

        var wordpress = new Application("wordpress", ApplicationArgs.builder()
            .name("wordpress")
            .modelUuid(development.uuid())
            .charms(ApplicationCharmArgs.builder()
                .name("wordpress")
                .build())
            .units(3)
            .build());

        var percona_cluster = new Application("percona-cluster", ApplicationArgs.builder()
            .name("percona-cluster")
            .modelUuid(development.uuid())
            .charms(ApplicationCharmArgs.builder()
                .name("percona-cluster")
                .build())
            .units(3)
            .build());

        var wpToPercona = new Integration("wpToPercona", IntegrationArgs.builder()
            .modelUuid(development.uuid())
            .applications(
                IntegrationApplicationArgs.builder()
                    .name(wordpress.name())
                    .endpoint("db")
                    .build(),
                IntegrationApplicationArgs.builder()
                    .name(percona_cluster.name())
                    .endpoint("server")
                    .build())
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}


{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    juju:caCertificate:
        value: 'TODO: file("~/ca-cert.pem")'
    juju:controllerAddresses:
        value: 10.225.205.241:17070,10.225.205.242:17070
    juju:password:
        value: password1
    juju:username:
        value: jujuuser

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as juju from "@pulumi/juju";

const development = new juju.Model("development", {
    name: "development",
    clouds: [{
        name: "aws",
        region: "eu-west-1",
    }],
});
const wordpress = new juju.Application("wordpress", {
    name: "wordpress",
    modelUuid: development.uuid,
    charms: [{
        name: "wordpress",
    }],
    units: 3,
});
const percona_cluster = new juju.Application("percona-cluster", {
    name: "percona-cluster",
    modelUuid: development.uuid,
    charms: [{
        name: "percona-cluster",
    }],
    units: 3,
});
const wpToPercona = new juju.Integration("wp_to_percona", {
    modelUuid: development.uuid,
    applications: [
        {
            name: wordpress.name,
            endpoint: "db",
        },
        {
            name: percona_cluster.name,
            endpoint: "server",
        },
    ],
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    juju:caCertificate:
        value: 'TODO: file("~/ca-cert.pem")'
    juju:controllerAddresses:
        value: 10.225.205.241:17070,10.225.205.242:17070
    juju:password:
        value: password1
    juju:username:
        value: jujuuser

```
```python
import pulumi
import pulumi_juju as juju

development = juju.Model("development",
    name="development",
    clouds=[{
        "name": "aws",
        "region": "eu-west-1",
    }])
wordpress = juju.Application("wordpress",
    name="wordpress",
    model_uuid=development.uuid,
    charms=[{
        "name": "wordpress",
    }],
    units=3)
percona_cluster = juju.Application("percona-cluster",
    name="percona-cluster",
    model_uuid=development.uuid,
    charms=[{
        "name": "percona-cluster",
    }],
    units=3)
wp_to_percona = juju.Integration("wp_to_percona",
    model_uuid=development.uuid,
    applications=[
        {
            "name": wordpress.name,
            "endpoint": "db",
        },
        {
            "name": percona_cluster.name,
            "endpoint": "server",
        },
    ])
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    juju:caCertificate:
        value: 'TODO: file("~/ca-cert.pem")'
    juju:controllerAddresses:
        value: 10.225.205.241:17070,10.225.205.242:17070
    juju:password:
        value: password1
    juju:username:
        value: jujuuser

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Juju = Pulumi.Juju;

return await Deployment.RunAsync(() =>
{
    var development = new Juju.Model("development", new()
    {
        Name = "development",
        Clouds = new[]
        {
            new Juju.Inputs.ModelCloudArgs
            {
                Name = "aws",
                Region = "eu-west-1",
            },
        },
    });

    var wordpress = new Juju.Application("wordpress", new()
    {
        Name = "wordpress",
        ModelUuid = development.Uuid,
        Charms = new[]
        {
            new Juju.Inputs.ApplicationCharmArgs
            {
                Name = "wordpress",
            },
        },
        Units = 3,
    });

    var percona_cluster = new Juju.Application("percona-cluster", new()
    {
        Name = "percona-cluster",
        ModelUuid = development.Uuid,
        Charms = new[]
        {
            new Juju.Inputs.ApplicationCharmArgs
            {
                Name = "percona-cluster",
            },
        },
        Units = 3,
    });

    var wpToPercona = new Juju.Integration("wp_to_percona", new()
    {
        ModelUuid = development.Uuid,
        Applications = new[]
        {
            new Juju.Inputs.IntegrationApplicationArgs
            {
                Name = wordpress.Name,
                Endpoint = "db",
            },
            new Juju.Inputs.IntegrationApplicationArgs
            {
                Name = percona_cluster.Name,
                Endpoint = "server",
            },
        },
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    juju:caCertificate:
        value: 'TODO: file("~/ca-cert.pem")'
    juju:controllerAddresses:
        value: 10.225.205.241:17070,10.225.205.242:17070
    juju:password:
        value: password1
    juju:username:
        value: jujuuser

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/juju/juju"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		development, err := juju.NewModel(ctx, "development", &juju.ModelArgs{
			Name: pulumi.String("development"),
			Clouds: juju.ModelCloudArray{
				&juju.ModelCloudArgs{
					Name:   pulumi.String("aws"),
					Region: pulumi.String("eu-west-1"),
				},
			},
		})
		if err != nil {
			return err
		}
		wordpress, err := juju.NewApplication(ctx, "wordpress", &juju.ApplicationArgs{
			Name:      pulumi.String("wordpress"),
			ModelUuid: development.Uuid,
			Charms: juju.ApplicationCharmArray{
				&juju.ApplicationCharmArgs{
					Name: pulumi.String("wordpress"),
				},
			},
			Units: pulumi.Float64(3),
		})
		if err != nil {
			return err
		}
		percona_cluster, err := juju.NewApplication(ctx, "percona-cluster", &juju.ApplicationArgs{
			Name:      pulumi.String("percona-cluster"),
			ModelUuid: development.Uuid,
			Charms: juju.ApplicationCharmArray{
				&juju.ApplicationCharmArgs{
					Name: pulumi.String("percona-cluster"),
				},
			},
			Units: pulumi.Float64(3),
		})
		if err != nil {
			return err
		}
		_, err = juju.NewIntegration(ctx, "wp_to_percona", &juju.IntegrationArgs{
			ModelUuid: development.Uuid,
			Applications: juju.IntegrationApplicationArray{
				&juju.IntegrationApplicationArgs{
					Name:     wordpress.Name,
					Endpoint: pulumi.String("db"),
				},
				&juju.IntegrationApplicationArgs{
					Name:     percona_cluster.Name,
					Endpoint: pulumi.String("server"),
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
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    juju:caCertificate:
        value: 'TODO: file("~/ca-cert.pem")'
    juju:controllerAddresses:
        value: 10.225.205.241:17070,10.225.205.242:17070
    juju:password:
        value: password1
    juju:username:
        value: jujuuser

```
```yaml
resources:
  development:
    type: juju:Model
    properties:
      name: development
      clouds:
        - name: aws
          region: eu-west-1
  wordpress:
    type: juju:Application
    properties:
      name: wordpress
      modelUuid: ${development.uuid}
      charms:
        - name: wordpress
      units: 3
  percona-cluster:
    type: juju:Application
    properties:
      name: percona-cluster
      modelUuid: ${development.uuid}
      charms:
        - name: percona-cluster
      units: 3
  wpToPercona:
    type: juju:Integration
    name: wp_to_percona
    properties:
      modelUuid: ${development.uuid}
      applications:
        - name: ${wordpress.name}
          endpoint: db
        - name: ${["percona-cluster"].name}
          endpoint: server
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    juju:caCertificate:
        value: 'TODO: file("~/ca-cert.pem")'
    juju:controllerAddresses:
        value: 10.225.205.241:17070,10.225.205.242:17070
    juju:password:
        value: password1
    juju:username:
        value: jujuuser

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.juju.Model;
import com.pulumi.juju.ModelArgs;
import com.pulumi.juju.inputs.ModelCloudArgs;
import com.pulumi.juju.Application;
import com.pulumi.juju.ApplicationArgs;
import com.pulumi.juju.inputs.ApplicationCharmArgs;
import com.pulumi.juju.Integration;
import com.pulumi.juju.IntegrationArgs;
import com.pulumi.juju.inputs.IntegrationApplicationArgs;
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
        var development = new Model("development", ModelArgs.builder()
            .name("development")
            .clouds(ModelCloudArgs.builder()
                .name("aws")
                .region("eu-west-1")
                .build())
            .build());

        var wordpress = new Application("wordpress", ApplicationArgs.builder()
            .name("wordpress")
            .modelUuid(development.uuid())
            .charms(ApplicationCharmArgs.builder()
                .name("wordpress")
                .build())
            .units(3)
            .build());

        var percona_cluster = new Application("percona-cluster", ApplicationArgs.builder()
            .name("percona-cluster")
            .modelUuid(development.uuid())
            .charms(ApplicationCharmArgs.builder()
                .name("percona-cluster")
                .build())
            .units(3)
            .build());

        var wpToPercona = new Integration("wpToPercona", IntegrationArgs.builder()
            .modelUuid(development.uuid())
            .applications(
                IntegrationApplicationArgs.builder()
                    .name(wordpress.name())
                    .endpoint("db")
                    .build(),
                IntegrationApplicationArgs.builder()
                    .name(percona_cluster.name())
                    .endpoint("server")
                    .build())
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

Pulumi 0.12 and later with client credentials:
{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    juju:caCertificate:
        value: 'TODO: file("~/ca-cert.pem")'
    juju:clientId:
        value: jujuclientid
    juju:clientSecret:
        value: jujuclientsecret
    juju:controllerAddresses:
        value: 10.225.205.241:17070,10.225.205.242:17070

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as juju from "@pulumi/juju";

const development = new juju.Model("development", {
    name: "development",
    clouds: [{
        name: "aws",
        region: "eu-west-1",
    }],
});
const wordpress = new juju.Application("wordpress", {
    name: "wordpress",
    modelUuid: development.uuid,
    charms: [{
        name: "wordpress",
    }],
    units: 3,
});
const percona_cluster = new juju.Application("percona-cluster", {
    name: "percona-cluster",
    modelUuid: development.uuid,
    charms: [{
        name: "percona-cluster",
    }],
    units: 3,
});
const wpToPercona = new juju.Integration("wp_to_percona", {
    modelUuid: development.uuid,
    applications: [
        {
            name: wordpress.name,
            endpoint: "db",
        },
        {
            name: percona_cluster.name,
            endpoint: "server",
        },
    ],
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    juju:caCertificate:
        value: 'TODO: file("~/ca-cert.pem")'
    juju:clientId:
        value: jujuclientid
    juju:clientSecret:
        value: jujuclientsecret
    juju:controllerAddresses:
        value: 10.225.205.241:17070,10.225.205.242:17070

```
```python
import pulumi
import pulumi_juju as juju

development = juju.Model("development",
    name="development",
    clouds=[{
        "name": "aws",
        "region": "eu-west-1",
    }])
wordpress = juju.Application("wordpress",
    name="wordpress",
    model_uuid=development.uuid,
    charms=[{
        "name": "wordpress",
    }],
    units=3)
percona_cluster = juju.Application("percona-cluster",
    name="percona-cluster",
    model_uuid=development.uuid,
    charms=[{
        "name": "percona-cluster",
    }],
    units=3)
wp_to_percona = juju.Integration("wp_to_percona",
    model_uuid=development.uuid,
    applications=[
        {
            "name": wordpress.name,
            "endpoint": "db",
        },
        {
            "name": percona_cluster.name,
            "endpoint": "server",
        },
    ])
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    juju:caCertificate:
        value: 'TODO: file("~/ca-cert.pem")'
    juju:clientId:
        value: jujuclientid
    juju:clientSecret:
        value: jujuclientsecret
    juju:controllerAddresses:
        value: 10.225.205.241:17070,10.225.205.242:17070

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Juju = Pulumi.Juju;

return await Deployment.RunAsync(() =>
{
    var development = new Juju.Model("development", new()
    {
        Name = "development",
        Clouds = new[]
        {
            new Juju.Inputs.ModelCloudArgs
            {
                Name = "aws",
                Region = "eu-west-1",
            },
        },
    });

    var wordpress = new Juju.Application("wordpress", new()
    {
        Name = "wordpress",
        ModelUuid = development.Uuid,
        Charms = new[]
        {
            new Juju.Inputs.ApplicationCharmArgs
            {
                Name = "wordpress",
            },
        },
        Units = 3,
    });

    var percona_cluster = new Juju.Application("percona-cluster", new()
    {
        Name = "percona-cluster",
        ModelUuid = development.Uuid,
        Charms = new[]
        {
            new Juju.Inputs.ApplicationCharmArgs
            {
                Name = "percona-cluster",
            },
        },
        Units = 3,
    });

    var wpToPercona = new Juju.Integration("wp_to_percona", new()
    {
        ModelUuid = development.Uuid,
        Applications = new[]
        {
            new Juju.Inputs.IntegrationApplicationArgs
            {
                Name = wordpress.Name,
                Endpoint = "db",
            },
            new Juju.Inputs.IntegrationApplicationArgs
            {
                Name = percona_cluster.Name,
                Endpoint = "server",
            },
        },
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    juju:caCertificate:
        value: 'TODO: file("~/ca-cert.pem")'
    juju:clientId:
        value: jujuclientid
    juju:clientSecret:
        value: jujuclientsecret
    juju:controllerAddresses:
        value: 10.225.205.241:17070,10.225.205.242:17070

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/juju/juju"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		development, err := juju.NewModel(ctx, "development", &juju.ModelArgs{
			Name: pulumi.String("development"),
			Clouds: juju.ModelCloudArray{
				&juju.ModelCloudArgs{
					Name:   pulumi.String("aws"),
					Region: pulumi.String("eu-west-1"),
				},
			},
		})
		if err != nil {
			return err
		}
		wordpress, err := juju.NewApplication(ctx, "wordpress", &juju.ApplicationArgs{
			Name:      pulumi.String("wordpress"),
			ModelUuid: development.Uuid,
			Charms: juju.ApplicationCharmArray{
				&juju.ApplicationCharmArgs{
					Name: pulumi.String("wordpress"),
				},
			},
			Units: pulumi.Float64(3),
		})
		if err != nil {
			return err
		}
		percona_cluster, err := juju.NewApplication(ctx, "percona-cluster", &juju.ApplicationArgs{
			Name:      pulumi.String("percona-cluster"),
			ModelUuid: development.Uuid,
			Charms: juju.ApplicationCharmArray{
				&juju.ApplicationCharmArgs{
					Name: pulumi.String("percona-cluster"),
				},
			},
			Units: pulumi.Float64(3),
		})
		if err != nil {
			return err
		}
		_, err = juju.NewIntegration(ctx, "wp_to_percona", &juju.IntegrationArgs{
			ModelUuid: development.Uuid,
			Applications: juju.IntegrationApplicationArray{
				&juju.IntegrationApplicationArgs{
					Name:     wordpress.Name,
					Endpoint: pulumi.String("db"),
				},
				&juju.IntegrationApplicationArgs{
					Name:     percona_cluster.Name,
					Endpoint: pulumi.String("server"),
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
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    juju:caCertificate:
        value: 'TODO: file("~/ca-cert.pem")'
    juju:clientId:
        value: jujuclientid
    juju:clientSecret:
        value: jujuclientsecret
    juju:controllerAddresses:
        value: 10.225.205.241:17070,10.225.205.242:17070

```
```yaml
resources:
  development:
    type: juju:Model
    properties:
      name: development
      clouds:
        - name: aws
          region: eu-west-1
  wordpress:
    type: juju:Application
    properties:
      name: wordpress
      modelUuid: ${development.uuid}
      charms:
        - name: wordpress
      units: 3
  percona-cluster:
    type: juju:Application
    properties:
      name: percona-cluster
      modelUuid: ${development.uuid}
      charms:
        - name: percona-cluster
      units: 3
  wpToPercona:
    type: juju:Integration
    name: wp_to_percona
    properties:
      modelUuid: ${development.uuid}
      applications:
        - name: ${wordpress.name}
          endpoint: db
        - name: ${["percona-cluster"].name}
          endpoint: server
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    juju:caCertificate:
        value: 'TODO: file("~/ca-cert.pem")'
    juju:clientId:
        value: jujuclientid
    juju:clientSecret:
        value: jujuclientsecret
    juju:controllerAddresses:
        value: 10.225.205.241:17070,10.225.205.242:17070

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.juju.Model;
import com.pulumi.juju.ModelArgs;
import com.pulumi.juju.inputs.ModelCloudArgs;
import com.pulumi.juju.Application;
import com.pulumi.juju.ApplicationArgs;
import com.pulumi.juju.inputs.ApplicationCharmArgs;
import com.pulumi.juju.Integration;
import com.pulumi.juju.IntegrationArgs;
import com.pulumi.juju.inputs.IntegrationApplicationArgs;
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
        var development = new Model("development", ModelArgs.builder()
            .name("development")
            .clouds(ModelCloudArgs.builder()
                .name("aws")
                .region("eu-west-1")
                .build())
            .build());

        var wordpress = new Application("wordpress", ApplicationArgs.builder()
            .name("wordpress")
            .modelUuid(development.uuid())
            .charms(ApplicationCharmArgs.builder()
                .name("wordpress")
                .build())
            .units(3)
            .build());

        var percona_cluster = new Application("percona-cluster", ApplicationArgs.builder()
            .name("percona-cluster")
            .modelUuid(development.uuid())
            .charms(ApplicationCharmArgs.builder()
                .name("percona-cluster")
                .build())
            .units(3)
            .build());

        var wpToPercona = new Integration("wpToPercona", IntegrationArgs.builder()
            .modelUuid(development.uuid())
            .applications(
                IntegrationApplicationArgs.builder()
                    .name(wordpress.name())
                    .endpoint("db")
                    .build(),
                IntegrationApplicationArgs.builder()
                    .name(percona_cluster.name())
                    .endpoint("server")
                    .build())
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

- `caCertificate` (String) If the controller was deployed with a self-signed certificate: This is the certificate to use for identification. This can also be set by the `JUJU_CA_CERT` environment variable
- `clientId` (String) If using JAAS: This is the client ID (OAuth2.0, created by the external identity provider) to be used. This can also be set by the `JUJU_CLIENT_ID` environment variable
- `clientSecret` (String, Sensitive) If using JAAS: This is the client secret (OAuth2.0, created by the external identity provider) to be used. This can also be set by the `JUJU_CLIENT_SECRET` environment variable
- `controllerAddresses` (String) This is the controller addresses to connect to, defaults to localhost:17070, multiple addresses can be provided in this format: <host>:<port>,<host>:<port>,.... This can also be set by the `JUJU_CONTROLLER_ADDRESSES` environment variable.
- `password` (String, Sensitive) This is the password of the username to be used. This can also be set by the `JUJU_PASSWORD` environment variable
- `skipFailedDeletion` (Boolean) Whether to issue a warning instead of an error and continue if a resource deletion fails. This can also be set by the `JUJU_SKIP_FAILED_DELETION` environment variable. Defaults to false.
- `username` (String) This is the username registered with the controller to be used. This can also be set by the `JUJU_USERNAME` environment variable