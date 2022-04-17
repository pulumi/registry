---
title: HCP
meta_desc: Provides an overview of the HashiCorp Cloud Platform Provider for Pulumi.
layout: overview
---

The HashiCorp Cloud Platform provider (HCP) for Pulumi can be used to
provision any of the cloud resources available in [HashiCorp Cloud
Platform](https://www.hashicorp.com/cloud-platform). The HCP provider
must be configured with credentials to deploy and update resources.

## Example

{{< chooser language "typescript,python,go" >}}
{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as hcp from "@grapl/pulumi-hcp";

const hvn = new hcp.Hvn(
    "my-hvn",
    {
        hvnId: "my-hvn",
        cloudProvider: "aws",
        region: "us-east-1"
    }
);

new hcp.VaultCluster(
    "my-vault-cluster",
    {
        hvnId: hvn.hvnId,
        clusterId: "my-vault-cluster",
        tier: "dev"
    }
);
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
import pulumi_hcp as hcp

def main() -> None:
    hvn = hcp.Hvn(
        "my-hvn",
        hvn_id="my-hvn",
        cloud_provider="aws",
        region="us-east-1",
    )

    hcp.VaultCluster(
        "my-vault-cluster",
        hvn_id=hvn.hvn_id,
        cluster_id="my-vault-cluster",
        tier="dev",
    )


if __name__ == "__main__":
    main()
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
    "fmt"
    hcp "github.com/grapl-security/pulumi-hcp/sdk/go/hcp"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {

        hvn, err := hcp.NewHvn(ctx, "my-hvn", &hcp.HvnArgs{
            HvnId:         pulumi.String("my-hvn"),
            CloudProvider: pulumi.String("aws"),
            Region:        pulumi.String("us-east-1"),
        })
        if err != nil {
            return fmt.Errorf("error creating HVN: %v", err)
        }

        _, err = hcp.NewVaultCluster(ctx, "my-vault-cluster", &hcp.VaultClusterArgs{
            HvnId:     hvn.HvnId,
            ClusterId: pulumi.String("my-vault-cluster"),
            Tier:      pulumi.String("dev"),
        })
        if err != nil {
            return fmt.Errorf("error creating Vault cluster: %v", err)
        }

        return nil
    })
}
```

{{% /choosable %}}
{{< /chooser >}}
