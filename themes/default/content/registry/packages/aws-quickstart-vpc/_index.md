---
title: AWS QuickStart VPC
meta_desc: Use Pulumi's Component for creating an AWS QuickStart VPC using infrastructure as code.
layout: overview
---

Easily create AWS VPC based on the AWS QuickStart VPC guide as a package available in all Pulumi languages.

Example:

{{< chooser language "typescript" >}}

{{% choosable language typescript %}}

```typescript
const myVpc = new aws.Vpc("test", {
    cidrBlock: "10.0.0.0/16",
    availabilityZoneConfig: [{
        availabilityZone: "us-west-2a",
        publicSubnetCidr: "10.0.128.0/20",
        privateSubnetACidr: "10.0.32.0/19",
    }, {
        availabilityZone: "us-west-2b",
        privateSubnetACidr: "10.0.64.0/19",
    }]
})

export const vpcID = myVpc.vpcID;
export const publicSubnetIDs = myVpc.publicSubnetIDs;
export const privateSubnetIDs = myVpc.privateSubnetIDs;
export const natgatewayIPs = myVpc.natGatewayIPs;
```

{{% /choosable %}}

{{< /chooser >}}

## SDK packages

The AWS API Gateway provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/aws-quickstart-vpc`](https://www.npmjs.com/package/@pulumi/aws-quickstart-vpc)
* Python: [`pulumi-aws-quickstart-vpc`](https://pypi.org/project/pulumi-aws-quickstart-vpc/)
* Go: [`github.com/pulumi/pulumi-aws-quickstart-vpc/sdk/go/aws`](https://github.com/pulumi/pulumi-aws-quickstart-vpc)
* .NET: [`Pulumi.AwsQuickStartVpc`](https://www.nuget.org/packages/Pulumi.AwsQuickStartVpc)
