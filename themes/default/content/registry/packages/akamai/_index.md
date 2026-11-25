---
title: Akamai
meta_desc: Provides an overview of the Akamai Provider for Pulumi.
layout: package
---

The Akamai provider for Pulumi can be used to provision any of the cloud resources available in [Akamai](https://www.akamai.com/).
The Akamai provider must be configured with credentials to deploy and update resources in Akamai.

## Example

{{< chooser language "typescript,python,go,csharp,java,yaml,hcl" >}}

{{% choosable language typescript %}}

```typescript
import * as akamai from "@pulumi/akamai";

const contractId = akamai.getContractOutput({
    groupName: "example-group",
}).id;
const groupId = akamai.getGroupOutput({
    contractId: contractId,
    groupName: "example-group",
}).id;
const edgeHostname = new akamai.EdgeHostName("edgeHostname", {
    contractId: contractId,
    groupId: groupId,
    productId: "prd_Fresca",
    edgeHostname: "www.example.com.edgesuite.net",
    ipBehavior: "IPV4",
});
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi_akamai as akamai

contract_id = akamai.get_contract_output(group_name="example-group").id
group_id = akamai.get_group_output(contract_id=contract_id,
    group_name="example-group").id
edge_hostname = akamai.EdgeHostName("edgeHostname",
    contract_id=contract_id,
    group_id=group_id,
    product_id="prd_Fresca",
    edge_hostname="www.example.com.edgesuite.net",
    ip_behavior="IPV4")
```

{{% /choosable %}}

{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi-akamai/sdk/v12/go/akamai"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		contractId := akamai.GetContractOutput(ctx, akamai.GetContractOutputArgs{
			GroupName: pulumi.String("example-group"),
		}, nil).Id()
		groupId := akamai.GetGroupOutput(ctx, akamai.GetGroupOutputArgs{
			ContractId: contractId,
			GroupName:  pulumi.String("example-group"),
		}, nil).Id()
		_, err := akamai.NewEdgeHostName(ctx, "edgeHostname", &akamai.EdgeHostNameArgs{
			ContractId:   contractId,
			GroupId:      groupId,
			ProductId:    pulumi.String("prd_Fresca"),
			EdgeHostname: pulumi.String("www.example.com.edgesuite.net"),
			IpBehavior:   pulumi.String("IPV4"),
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
using Pulumi;
using Akamai = Pulumi.Akamai;

await Deployment.RunAsync(() =>
{
    var contractId = Akamai.GetContract.Invoke(new()
    {
        GroupName = "example-group",
    }).Apply(invoke => invoke.Id);

    var groupId = Akamai.GetGroup.Invoke(new()
    {
        ContractId = contractId,
        GroupName = "example-group",
    }).Apply(invoke => invoke.Id);

    var edgeHostname = new Akamai.EdgeHostName("edgeHostname", new()
    {
        ContractId = contractId,
        GroupId = groupId,
        ProductId = "prd_Fresca",
        EdgeHostname = "www.example.com.edgesuite.net",
        IpBehavior = "IPV4",
    });
});
```

{{% /choosable %}}

{{% choosable language java %}}

```java
import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.akamai.AkamaiFunctions;
import com.pulumi.akamai.inputs.GetContractArgs;
import com.pulumi.akamai.inputs.GetGroupArgs;
import com.pulumi.akamai.EdgeHostName;
import com.pulumi.akamai.EdgeHostNameArgs;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        final var contractId = AkamaiFunctions.getContract(GetContractArgs.builder()
            .groupName("example-group")
            .build()).applyValue(_invoke -> _invoke.id());

        final var groupId = AkamaiFunctions.getGroup(GetGroupArgs.builder()
            .contractId(contractId)
            .groupName("example-group")
            .build()).applyValue(_invoke -> _invoke.id());

        var edgeHostname = new EdgeHostName("edgeHostname", EdgeHostNameArgs.builder()
            .contractId(contractId)
            .groupId(groupId)
            .productId("prd_Fresca")
            .edgeHostname("www.example.com.edgesuite.net")
            .ipBehavior("IPV4")
            .build());
    }
}
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
variables:
  contractId:
    fn::invoke:
      function: akamai:getContract
      arguments:
        groupName: example-group
      return: id
  groupId:
    fn::invoke:
      function: akamai:getGroup
      arguments:
        contractId: ${contractId}
        groupName: example-group
      return: id
resources:
  edgeHostname:
    type: akamai:EdgeHostName
    properties:
      contractId: ${contractId}
      groupId: ${groupId}
      productId: prd_Fresca
      edgeHostname: www.example.com.edgesuite.net
      ipBehavior: IPV4
```

{{% /choosable %}}

{{% choosable language hcl %}}

```hcl
terraform {
  required_providers {
    akamai = {
      source = "pulumi/akamai"
    }
  }
}

data "akamai_contract" "contract" {
  group_name = "example-group"
}

data "akamai_group" "group" {
  contract_id = data.akamai_contract.contract.id
  group_name  = "example-group"
}

resource "akamai_edge_host_name" "edgeHostname" {
  contract_id   = data.akamai_contract.contract.id
  group_id      = data.akamai_group.group.id
  product_id    = "prd_Fresca"
  edge_hostname = "www.example.com.edgesuite.net"
  ip_behavior   = "IPV4"
}
```

{{% /choosable %}}

{{< /chooser >}}
