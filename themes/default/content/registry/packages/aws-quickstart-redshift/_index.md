---
title: AWS QuickStart Redshift
meta_desc: Use Pulumi's Component for creating an AWS QuickStart Redshift Cluster using infrastructure as code.
layout: overview
---

Easily create AWS Redshift Cluster based on the AWS QuickStart Redshift guide as a package available in all Pulumi languages.

Example:

{{< chooser language "typescript" >}}

{{% choosable language typescript %}}

```typescript
const cluster = new redshift.Cluster("demo-cluster", {
    vpcID: "<vpcID>",
    subnetIDs: "<privateSubnetIDList>",
    dbMasterPassword: pulumi.secret("Password1!"),
    dbMasterUsername: "test-username",
    dbName: "test_database",
    dbNodeType: "dc2.large",
    dbClusterIdentifier: "demo-cluster",
    enableEventSubscription: false,
})
```

{{% /choosable %}}

{{< /chooser >}}

## SDK packages

The AWS API Gateway provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/aws-quickstart-redshift`](https://www.npmjs.com/package/@pulumi/aws-quickstart-redshift)
* Python: [`pulumi-aws-quickstart-redshift`](https://pypi.org/project/pulumi-aws-quickstart-redshift/)
* Go: [`github.com/pulumi/pulumi-aws-quickstart-redshift/sdk/go/aws`](https://github.com/pulumi/pulumi-aws-quickstart-redshift)
* .NET: [`Pulumi.AwsQuickStartRedshift`](https://www.nuget.org/packages/Pulumi.AwsQuickStartRedshift)
