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
