---
title: pgEdge
meta_desc: Provides an overview of the pgEdge Provider for Pulumi.
layout: package
---

The pgEdge provider for Pulumi can be used to provision any of the cloud resources available in [pgEdge](https://app.pgedge.com/).
The pgEdge provider must be configured with the a `Client ID` and `Client Secret` in order to use pgEdge resources.

## Example
{{< chooser language "typescript,go" >}}

{{% choosable language typescript %}}

```typescript
import * as pgedge from "@pgEdge/pulumi-pgedge";

var testCluster = new pgedge.Cluster("clusterCreate",{
  name: "n1",
  cloudAccountId: "", // cloud account id
  regions:["us-east-2"],
  firewallRules: [
    {
      port: 5432,
      sources: ["0.0.0.0/0"],
    }
  ],
  nodes: [
    {
      name: "n1",
      instanceType: "t4g.medium",
      volumeSize: 30,
      region: "us-east-2",
      availabilityZone: "us-east-2a",
      volumeType: "gp2",
    }
  ],
  networks: [
    {
      region: "us-east-2",
      cidr: "10.2.0.0/16",
      publicSubnets: ["10.2.1.0/24"],
    }
  ],
});

const testDatabase = new pgedge.Database("databaseCreate", {
  name: "testdb",
  clusterId: testCluster.id,
  options: ["install:northwind", "autoddl:enabled"],
});
```


{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"github.com/pgEdge/pulumi-pgedge/sdk/go/pgedge"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

	cluster, err := pgedge.NewCluster(ctx, "cluster", &pgedge.ClusterArgs{
		Name:           pulumi.String("n1"),
		CloudAccountId: pulumi.String(""), // cloud account id
		Regions: pulumi.StringArray{
			pulumi.String("us-east-2"),
		},
		FirewallRules: pgedge.ClusterFirewallRuleArray{
			&pgedge.ClusterFirewallRuleArgs{
				Port: pulumi.Int(5432),
				Sources: pulumi.ToStringArray([]string{
					"0.0.0.0/0",
				}),
			},
		},
		Nodes: pgedge.ClusterNodeArray{
			&pgedge.ClusterNodeArgs{
				InstanceType:     pulumi.String("t4g.small"),
				Name:             pulumi.String("n1"),
				VolumeSize:       pulumi.Int(30),
				Region:           pulumi.String("us-east-2"),
				AvailabilityZone: pulumi.String("us-east-2a"),
				VolumeType:       pulumi.String("gp2"),
			},
		},
		Networks: pgedge.ClusterNetworkArray{
			&pgedge.ClusterNetworkArgs{
				Region: pulumi.String("us-east-2"),
				Cidr:   pulumi.String("10.2.0.0/16"),
				PublicSubnets: pulumi.ToStringArray([]string{
					"10.2.1.0/24",
				}),
			},
		},
	})

	if err != nil {
		return err
	}

	_, err = pgedge.NewDatabase(ctx, "database", &pgedge.DatabaseArgs{
		ClusterId: cluster.ID(),
		Name:      pulumi.String("defaultdb"),
		Options:   pulumi.StringArray{pulumi.String("install:northwind"), pulumi.String("autoddl:enabled")},
	})

	if err != nil {
		return err
	}
	return nil
})
}
```

{{% /choosable %}}

{{< /chooser >}}

> You could find more complete and detailed examples in the [pulumi-pgedge repository](https://github.com/pgEdge/pulumi-pgedge/tree/main/examples)
