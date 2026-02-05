---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-mongodbatlas/v4.3.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi/pulumi-mongodbatlas/blob/v4.3.0/docs/_index.md
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Mongodbatlas Provider
meta_desc: Provides an overview on how to configure the Pulumi Mongodbatlas provider.
layout: package
---

## Installation

The Mongodbatlas provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/mongodbatlas`](https://www.npmjs.com/package/@pulumi/mongodbatlas)
* Python: [`pulumi-mongodbatlas`](https://pypi.org/project/pulumi-mongodbatlas/)
* Go: [`github.com/pulumi/pulumi-mongodbatlas/sdk/v4/go/mongodbatlas`](https://github.com/pulumi/pulumi-mongodbatlas)
* .NET: [`Pulumi.Mongodbatlas`](https://www.nuget.org/packages/Pulumi.Mongodbatlas)
* Java: [`com.pulumi/mongodbatlas`](https://central.sonatype.com/artifact/com.pulumi/mongodbatlas)

## Overview

The MongoDB Atlas provider is used to interact with the resources supported by [MongoDB Atlas](https://www.mongodb.com/cloud/atlas). The provider needs to be configured with proper credentials before it can be used.
## Example Usage

This example shows how to set up the MongoDB Atlas provider and create a cluster:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    mongodbatlas:clientId:
        value: 'TODO: var.mongodbatlas_client_id'
    mongodbatlas:clientSecret:
        value: 'TODO: var.mongodbatlas_client_secret'

```

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as mongodbatlas from "@pulumi/mongodbatlas";

// Create a project
const _this = new mongodbatlas.Project("this", {
    name: "my-project",
    orgId: orgId,
});
// Create a cluster
const thisAdvancedCluster = new mongodbatlas.AdvancedCluster("this", {
    projectId: _this.id,
    name: "my-cluster",
    clusterType: "REPLICASET",
    replicationSpecs: [{
        regionConfigs: [{
            regionName: "US_EAST_1",
            priority: 7,
            providerName: "AWS",
            electableSpecs: {
                instanceSize: "M10",
                nodeCount: 3,
            },
        }],
    }],
});
```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    mongodbatlas:clientId:
        value: 'TODO: var.mongodbatlas_client_id'
    mongodbatlas:clientSecret:
        value: 'TODO: var.mongodbatlas_client_secret'

```

```python
import pulumi
import pulumi_mongodbatlas as mongodbatlas

# Create a project
this = mongodbatlas.Project("this",
    name="my-project",
    org_id=org_id)
# Create a cluster
this_advanced_cluster = mongodbatlas.AdvancedCluster("this",
    project_id=this.id,
    name="my-cluster",
    cluster_type="REPLICASET",
    replication_specs=[{
        "region_configs": [{
            "region_name": "US_EAST_1",
            "priority": 7,
            "provider_name": "AWS",
            "electable_specs": {
                "instance_size": "M10",
                "node_count": 3,
            },
        }],
    }])
```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    mongodbatlas:clientId:
        value: 'TODO: var.mongodbatlas_client_id'
    mongodbatlas:clientSecret:
        value: 'TODO: var.mongodbatlas_client_secret'

```

```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Mongodbatlas = Pulumi.Mongodbatlas;

return await Deployment.RunAsync(() =>
{
    // Create a project
    var @this = new Mongodbatlas.Project("this", new()
    {
        Name = "my-project",
        OrgId = orgId,
    });

    // Create a cluster
    var thisAdvancedCluster = new Mongodbatlas.AdvancedCluster("this", new()
    {
        ProjectId = @this.Id,
        Name = "my-cluster",
        ClusterType = "REPLICASET",
        ReplicationSpecs = new[]
        {
            new Mongodbatlas.Inputs.AdvancedClusterReplicationSpecArgs
            {
                RegionConfigs = new[]
                {
                    new Mongodbatlas.Inputs.AdvancedClusterReplicationSpecRegionConfigArgs
                    {
                        RegionName = "US_EAST_1",
                        Priority = 7,
                        ProviderName = "AWS",
                        ElectableSpecs = new Mongodbatlas.Inputs.AdvancedClusterReplicationSpecRegionConfigElectableSpecsArgs
                        {
                            InstanceSize = "M10",
                            NodeCount = 3,
                        },
                    },
                },
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
    mongodbatlas:clientId:
        value: 'TODO: var.mongodbatlas_client_id'
    mongodbatlas:clientSecret:
        value: 'TODO: var.mongodbatlas_client_secret'

```

```go
package main

import (
	"github.com/pulumi/pulumi-mongodbatlas/sdk/v4/go/mongodbatlas"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a project
		this, err := mongodbatlas.NewProject(ctx, "this", &mongodbatlas.ProjectArgs{
			Name:  pulumi.String("my-project"),
			OrgId: pulumi.Any(orgId),
		})
		if err != nil {
			return err
		}
		// Create a cluster
		_, err = mongodbatlas.NewAdvancedCluster(ctx, "this", &mongodbatlas.AdvancedClusterArgs{
			ProjectId:   this.ID(),
			Name:        pulumi.String("my-cluster"),
			ClusterType: pulumi.String("REPLICASET"),
			ReplicationSpecs: mongodbatlas.AdvancedClusterReplicationSpecArray{
				&mongodbatlas.AdvancedClusterReplicationSpecArgs{
					RegionConfigs: mongodbatlas.AdvancedClusterReplicationSpecRegionConfigArray{
						&mongodbatlas.AdvancedClusterReplicationSpecRegionConfigArgs{
							RegionName:   pulumi.String("US_EAST_1"),
							Priority:     pulumi.Int(7),
							ProviderName: pulumi.String("AWS"),
							ElectableSpecs: &mongodbatlas.AdvancedClusterReplicationSpecRegionConfigElectableSpecsArgs{
								InstanceSize: pulumi.String("M10"),
								NodeCount:    pulumi.Int(3),
							},
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
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    mongodbatlas:clientId:
        value: 'TODO: var.mongodbatlas_client_id'
    mongodbatlas:clientSecret:
        value: 'TODO: var.mongodbatlas_client_secret'

```

```yaml
resources:
  # Create a project
  this:
    type: mongodbatlas:Project
    properties:
      name: my-project
      orgId: ${orgId}
  # Create a cluster
  thisAdvancedCluster:
    type: mongodbatlas:AdvancedCluster
    name: this
    properties:
      projectId: ${this.id}
      name: my-cluster
      clusterType: REPLICASET
      replicationSpecs:
        - regionConfigs:
            - regionName: US_EAST_1
              priority: 7
              providerName: AWS
              electableSpecs:
                instanceSize: M10
                nodeCount: 3
```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    mongodbatlas:clientId:
        value: 'TODO: var.mongodbatlas_client_id'
    mongodbatlas:clientSecret:
        value: 'TODO: var.mongodbatlas_client_secret'

```

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.mongodbatlas.Project;
import com.pulumi.mongodbatlas.ProjectArgs;
import com.pulumi.mongodbatlas.AdvancedCluster;
import com.pulumi.mongodbatlas.AdvancedClusterArgs;
import com.pulumi.mongodbatlas.inputs.AdvancedClusterReplicationSpecArgs;
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
        // Create a project
        var this_ = new Project("this", ProjectArgs.builder()
            .name("my-project")
            .orgId(orgId)
            .build());

        // Create a cluster
        var thisAdvancedCluster = new AdvancedCluster("thisAdvancedCluster", AdvancedClusterArgs.builder()
            .projectId(this_.id())
            .name("my-cluster")
            .clusterType("REPLICASET")
            .replicationSpecs(AdvancedClusterReplicationSpecArgs.builder()
                .regionConfigs(AdvancedClusterReplicationSpecRegionConfigArgs.builder()
                    .regionName("US_EAST_1")
                    .priority(7)
                    .providerName("AWS")
                    .electableSpecs(AdvancedClusterReplicationSpecRegionConfigElectableSpecsArgs.builder()
                        .instanceSize("M10")
                        .nodeCount(3)
                        .build())
                    .build())
                .build())
            .build());

    }
}
```

{{% /choosable %}}
{{< /chooser >}}
## Authentication

The MongoDB Atlas provider uses Service Account (SA) as the recommended authentication method.

For detailed authentication configuration, see:
- Service Account (SA)
- Programmatic Access Key (PAK)
- AWS Secrets Manager integration
## MongoDB Atlas for Government

MongoDB Atlas for Government is a dedicated deployment option for government agencies and contractors requiring FedRAMP compliance.
For more details on configuration, see the Provider Configuration Guide.