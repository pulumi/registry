---
title: AWSx (Pulumi Crosswalk for AWS)
meta_desc: Well-Architected Infrastructure as Code for AWS. The easiest way to AWS — from development to production.
layout: overview
---

{{% notes %}}
This AWSx documentation is generated using the v1.0.0 series of AWSx. If you have not upgraded to that version of
the provider then we suggest you continue to use our [@pulumi/awsx]({{<relref "/docs/reference/pkg/nodejs/pulumi/awsx">}}) NodeJS documentation.
{{% /notes %}}

The Amazon Web Services (AWS) [Crosswalk](../../../docs/guides/crosswalk/aws/) (AWSx) library uses automatic well-architected best practices to make common infrastructure-as-code tasks in [AWS](https://aws.amazon.com/) easier and more secure. It uses the AWS SDK to manage and provision resources.

The AWSx provider must be configured with credentials to deploy and update resources in AWS; see [Installation & Configuration]({{<relref "./installation-configuration">}}) for instructions.

**New to Pulumi and AWS?** [Get started with AWS using our AWSx tutorial]({{<relref "/docs/guides/crosswalk/aws">}}).

## Example

{{< chooser language "typescript,python,csharp,java,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as awsx from "@pulumi/awsx";

// Allocate a new VPC with the default settings:
const vpc = new awsx.ec2.Vpc("custom");

// Export a few resulting fields to make them easy to use:
export const vpcId = vpc.vpcId;
export const privateSubnetIds = vpc.privateSubnetIds;
export const publicSubnetIds = vpc.publicSubnetIds;
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_awsx as awsx

vpc = awsx.ec2.Vpc("custom")

pulumi.export("vpcId", vpc.vpc_id)
pulumi.export("publicSubnetIds", vpc.public_subnet_ids)
pulumi.export("privateSubnetIds", vpc.private_subnet_ids)
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using Pulumi;
using Ec2 = Pulumi.Awsx.Ec2;

class MyStack : Stack
{
    public MyStack()
    {
        var vpc = new Ec2.Vpc("custom");

        this.VpcId = vpc.VpcId;
    }


    [Output] public Output<string> VpcId { get; set; }
}

class Program
{
    static Task<int> Main(string[] args) => Deployment.RunAsync<MyStack>();
}
```

{{% /choosable %}}

{{% choosable language java %}}

```java
package myproject;

import com.pulumi.Pulumi;
import com.pulumi.awsx.ec2.Vpc;

public class App {
    public static void main(String[] args) {
        Pulumi.run(ctx -> {
            var vpc = new Vpc("custom");

            ctx.export("vpcId", vpc.vpcId());
            ctx.export("privateSubnetIds", vpc.privateSubnetIds());
            ctx.export("publicSubnetIds", vpc.publicSubnetIds());
        });
    }
}
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: awsx-networking-yaml
runtime: yaml
description: A minimal AWS Pulumi YAML program

resources:
  # Allocate a new VPC with the default settings:
  vpc:
    type: awsx:ec2:Vpc

outputs:
  vpcId: ${vpc.VpcId}
  privateSubnetIds: ${vpc.privateSubnetIds}
  publicSubnetIds: ${vpc.publicSubnetIds}

```

{{% /choosable %}}

{{% /chooser %}}
