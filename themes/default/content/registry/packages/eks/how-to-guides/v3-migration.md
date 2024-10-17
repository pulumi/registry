---
title: Amazon EKS v3 Migration Guide
meta_desc: How to upgrade from v2 to v3 of the EKS provider
layout: package
---

In 2024 AWS announced the deprecation of two features used by default in Pulumi EKS: the aws-auth ConfigMap and the AL2 operating system. Pulumi EKS v3 addresses these deprecations, enhances the maintainability of the provider, and aligns it with EKS best practices.
There are some breaking changes which are detailed in this document.

## VPC CNI Component changes

The VPC CNI cluster component is now configured as an EKS addon as mentioned in the “New Features” section above. This brings the following changes:

- Removed `enableIpv6` input property. The component automatically configures the IP version now, depending on whether the cluster is running in IPv4 or IPv6 mode.
- Removed `image`, `initImage`, `nodeAgentImage` input properties. The component now automatically selects an image registry in the cluster’s region to pull the images from.

During the update, you'll observe two key changes:
1. The removal of the previous `kubectl`-based `VpcCni` component.
2. The creation of the new EKS addon-based `VpcCniAddon` component.

It's important to note that deleting the `VpcCni` component will not modify or delete any existing on-cluster resources. When creating the `VpcCniAddon` component, it will safely adopt the existing VPC CNI resources and update them to the current recommended version.

## Node Group Updates

### `NodeGroup` component deprecation

The `NodeGroup` component uses the deprecated AWS Launch Configuration (see [AWS docs](https://docs.aws.amazon.com/autoscaling/ec2/userguide/launch-configurations.html)). Launch Configurations do not support new instance types released after December 31, 2022 and starting on October 1, 2024, new AWS accounts will not be able to create launch configurations.
Its successor, the `NodeGroupV2` component is functionally equivalent and easier to operate because it does not use CloudFormation templates under the hood like `NodeGroup` did.

The default node group of the `Cluster` component has been updated to use the `NodeGroupV2` component as well.
Updates to the default node group will be done by first creating the new replacement nodes and then shutting down the old ones which will move pods to the new nodes. If you need to perform the update gracefully, please have a look at [Gracefully upgrading node groups](#gracefully-upgrading-node-groups).

### Default OS of node groups updated to AL2023

AWS recently deprecated the Amazon Linux 2 (AL2) operating system. It will reach end of life in June 2025, after which it will receive no more security and maintenance updates.  
Until now, this was the OS used by node groups created with the Pulumi EKS provider.
To align the provider with EKS best practices we’ve updated the default operating system to the new AWS-recommended option, Amazon Linux 2023 (AL2023).

You can either upgrade the OS right away, or intentionally configure AL2 as the desired operating system by using the `operatingSystem` parameter that’s available on all node groups.

Have a look at [Gracefully upgrading node groups](#gracefully-upgrading-node-groups) for more information around the upgrade procedure.

### Gracefully upgrading node groups

The `ManagedNodeGroup` component gracefully handles updates by default. EKS will:
- boot the updated replacement nodes  
- cordon the old nodes to ensure no new pods get launched onto them  
- drain the old nodes one-by-one
- shut down the empty old nodes

The detailed update procedure can be seen in the [AWS docs](https://docs.aws.amazon.com/eks/latest/userguide/managed-node-update-behavior.html).

For self-managed node groups (i.e., the `NodeGroup` and `NodeGroupV2` components) you have two options:

1. Update the node group in place. Pulumi does this by first creating the new replacement nodes and then shutting down the old ones which will move pods to the new nodes forcibly. This is the default behavior when node groups are updated.  
2. Create a new node group and move your Pods to that group. Migrating to a new node group is more graceful than simply updating the node group in place. This is because the migration process taints the old node group as `NoSchedule` and drains the nodes gradually.

The second option involves the following steps:

* Create the replacement node group side-by-side with the existing node group. When doing this you need to make sure that the two node groups are allowed to communicate with each other. You can achieve this in the following way:


{{< chooser language "typescript,python,go,csharp,java" >}}

{{% choosable language typescript %}}

```typescript
const oldNG = new eks.NodeGroupV2("old", {
  cluster: cluster,
  operatingSystem: eks.OperatingSystem.AL2,
  instanceProfile: instanceProfile,
});
const export oldAsgName = oldNG.autoScalingGroup.name;

const newNG = new eks.NodeGroupV2("new", {
  cluster: cluster,
  operatingSystem: eks.OperatingSystem.AL2023,
  instanceProfile: instanceProfile,
});

// Allow all traffic between the old & new NG
const oldToNew = new aws.vpc.SecurityGroupIngressRule("oldToNew", {
    securityGroupId: oldNG.nodeSecurityGroupId,
    referencedSecurityGroupId: newNG.nodeSecurityGroupId,
    ipProtocol: "-1",
});
const newToOld = new aws.vpc.SecurityGroupIngressRule("newToOld", {
    securityGroupId: newNG.nodeSecurityGroupId,
    referencedSecurityGroupId: oldNG.nodeSecurityGroupId,
    ipProtocol: "-1",
});
```

{{% /choosable %}}

{{% choosable language python %}}

```python
old_ng = eks.NodeGroupV2("old",
                         cluster=cluster,
                         operating_system=eks.OperatingSystem.AL2,
                         instance_profile=instance_profile)

# Exporting the auto scaling group name
pulumi.export('old_asg_name', old_ng.auto_scaling_group.name)

new_ng = eks.NodeGroupV2("new",
                         cluster=cluster,
                         operating_system=eks.OperatingSystem.AL2023,
                         instance_profile=instance_profile)

# Allow all traffic between the old & new NG
old_to_new = aws.ec2.SecurityGroupIngress('oldToNew',
                                          security_group_id=old_ng.node_security_group_id,
                                          referenced_security_group_id=new_ng.node_security_group_id,
                                          ip_protocol='-1')
old_to_new = aws.ec2.SecurityGroupIngress('newToOld',
                                          security_group_id=new_ng.node_security_group_id,
                                          referenced_security_group_id=old_ng.node_security_group_id,
                                          ip_protocol='-1')
```
{{% /choosable %}}

{{% choosable language go %}}

```go
oldNg, err := eks.NewNodeGroupV2(ctx, "old", &eks.NodeGroupV2Args{
    Cluster:          cluster,
    OperatingSystem:  eks.OperatingSystemAL2,
    InstanceProfile:  instanceProfile,
})
if err != nil {
    return err
}

ctx.Export("old_asg_name", oldNg.AutoScalingGroup.ApplyT(func(asg *aws.autoscaling.Group) string {
    return asg.Name
}).(pulumi.StringOutput))

newNg, err := eks.NewNodeGroupV2(ctx, "new", &eks.NodeGroupV2Args{
    Cluster:          cluster,
    OperatingSystem:  eks.OperatingSystemAL2023,
    InstanceProfile:  instanceProfile,
})
if err != nil {
    return err
}

_, err = ec2.NewSecurityGroupIngress(ctx, "oldToNew", &ec2.SecurityGroupIngressArgs{
    SecurityGroupId:             oldNg.NodeSecurityGroupId,
    SourceSecurityGroupId:       newNg.NodeSecurityGroupId,
    Protocol:                    pulumi.String("-1"),
})
if err != nil {
    return err
}

_, err = ec2.NewSecurityGroupIngress(ctx, "newToOld", &ec2.SecurityGroupIngressArgs{
    SecurityGroupId:             newNg.NodeSecurityGroupId,
    SourceSecurityGroupId:       oldNg.NodeSecurityGroupId,
    Protocol:                    pulumi.String("-1"),
})
if err != nil {
    return err
}
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
var oldNG = new NodeGroupV2("old", new()
{
    Cluster = cluster,
    OperatingSystem = Pulumi.Eks.OperatingSystem.AL2,
    InstanceProfile = instanceProfile,
});
var oldAsgName = oldNG.AutoScalingGroup.Apply(asg => asg.Name);

var newNG = new NodeGroupV2("new", new()
{
    Cluster = cluster,
    OperatingSystem = Pulumi.Eks.OperatingSystem.AL2023,
    InstanceProfile = instanceProfile,
});

var oldToNew = new SecurityGroupIngressRule("oldToNew", new()
{
    SecurityGroupId = oldNG.NodeSecurityGroupId,
    ReferencedSecurityGroupId = newNG.NodeSecurityGroupId,
    IpProtocol = "-1",
});

var newToOld = new SecurityGroupIngressRule("newToOld", new()
{
    SecurityGroupId = newNG.NodeSecurityGroupId,
    ReferencedSecurityGroupId = oldNG.NodeSecurityGroupId,
    IpProtocol = "-1",
});
```

{{% /choosable %}}

{{% choosable language java %}}

```java
var oldNG = new NodeGroupV2("old", NodeGroupV2Args.builder()
    .cluster(cluster.id())
    .operatingSystem("AL2")
    .instanceProfile(instanceProfile.arn())
    .build()
);

Output<String> oldAsgName = oldNG.getAutoScalingGroup().applyValue(asg -> asg.name());

var newNG = new NodeGroupV2("new", NodeGroupV2Args.builder()
    .cluster(cluster.id())
    .operatingSystem("AL2023")
    .instanceProfile(instanceProfile.arn())
    .build()
);

var oldToNew = new SecurityGroupIngressRule("oldToNew", SecurityGroupIngressRuleArgs.builder()
    .securityGroupId(oldNG.nodeSecurityGroupId())
    .referencedSecurityGroupId(newNG.nodeSecurityGroupId())
    .ipProtocol("-1")
    .build()
);

var newToOld = new SecurityGroupIngressRule("newToOld", SecurityGroupIngressRuleArgs.builder()
    .securityGroupId(newNG.nodeSecurityGroupId())
    .referencedSecurityGroupId(oldNG.nodeSecurityGroupId())
    .ipProtocol("-1")
    .build()
);
```

{{% /choosable %}}

{{< /chooser >}}

* Find the nodes of the old node group. First take a note of the name of the auto scaling group associated with that node group and then run the following AWS CLI command, replacing `$ASG_GROUP_NAME` with the actual name of the auto scaling group:

```bash
aws ec2 describe-instances --filter "Name=tag:aws:autoscaling:groupName,Values=$ASG_GROUP_NAME" \
    | jq -r '.Reservations[].Instances[].PrivateDnsName'
```

* Drain each of the nodes of the old node group one by one. This will mark the nodes as unschedulable and gracefully move pods to other nodes. For more information have a look at this article in the [kubernetes documentation](https://kubernetes.io/docs/tasks/administer-cluster/safely-drain-node/).

```bash
kubectl drain $NODE_NAME --ignore-daemonsets --delete-emptydir-data
```

* The old nodes are now empty and can be safely shut down. Remove them from your pulumi program and run `pulumi up`

## aws-auth ConfigMap Deprecation

AWS introduced a new method for granting IAM principals access to Kubernetes resources called [Access Entries](https://docs.aws.amazon.com/eks/latest/userguide/grant-k8s-access.html#authentication-modes). In contrast to the existing approach using the aws-auth ConfigMap, this solely relies on AWS resources for managing Kubernetes auth.  
Recently the aws-auth ConfigMap has been deprecated in favor of Access Entries.  
You can start using Access Entries with your cluster by changing the `authenticationMode` to `API`.  
A step-by-step migration guide can be found [here](https://www.pulumi.com/registry/packages/eks/how-to-guides/access-entries-migration/).

We currently recommend users create all new clusters with the `API` authentication mode.
More details about this can be found [here](https://docs.aws.amazon.com/eks/latest/userguide/grant-k8s-access.html).

## Node.js SDK changes

The Node.js SDK is updated to use state of the art Pulumi tooling, improving stability, documentation and security. The update requires the following changes to programs:
- Properties of the components are now outputs instead of plain types. Accessing those optional resource outputs now requires using `apply`. More details about this in the next section. 
- The [`Cluster.getKubeConfig`](https://www.pulumi.com/registry/packages/eks/api-docs/cluster/#method_GetKubeconfig) method now returns an output. 
- The `cluster.provider` will be deleted if not referenced (no impact, but it will appear in the diff)
- The deprecated input property `deployDashboard` of the `Cluster` component has been removed from the Node.js SDK. This has already been removed from the other SDKs in the past. If you’d like to continue using it, you can adopt the existing code into your own program from [here](https://github.com/pulumi/pulumi-eks/blob/bcc170e72b802a78e7f0a99bc92316a5f8a62b0e/nodejs/eks/dashboard.ts).
- The `createManagedNodeGroup` function will now create an Pulumi EKS `ManagedNodeGroup` instead of creating the underlying `aws.eks.NodeGroup` resource directly. During the upgrade to Pulumi EKS v3 you'll see the additional wrapper component being created.
- The capitalization of the `AuthenticationMode` and `AccessEntryType` enum values has been aligned with other providers (e.g. `AuthenticationMode.API` => `AuthenticationMode.Api`). The previous values have been marked as deprecated and now point to their new counterparts.

### Properties of the components are now outputs instead of plain types
All properties of components in the Node.js SDK are now wrapped in `Output`. This aligns it with other language SDKs but introduces breaking changes for optional resource properties.
Accessing those optional resource outputs now requires using `apply`.

We added convenient access properties for frequently used outputs:
- Creating an IRSA based IAM role now requires using `apply` for accessing the cluster's OIDC provider ARN and URL. Alternatively, use the new top level properties `oidcProviderArn` and `oidcIssuer`. An example can be found [here](https://github.com/pulumi/pulumi-eks/blob/release-3.x.x/examples/oidc-iam-sa/index.ts).
- Accessing a cluster's `clusterSecurityGroup`, `nodeSecurityGroup`, and `eksClusterIngressRule` requires using `apply`. Alternatively, use the new top-level properties. The `NodeGroup` and `NodeGroupV2` components were updated to also take those new top-level properties as inputs.
  - Use the following top-level properties of the `Cluster` component:
    - `clusterSecurityGroupId` instead of `cluster.clusterSecurityGroup.id`
    - `nodeSecurityGroupId` instead of `cluster.nodeSecurityGroup.id`
    - `clusterIngressRuleId` instead of `cluster.eksClusterIngressRule.id`
  - When passing these to the `NodeGroup` and `NodeGroupV2` components use the following inputs:
    - `nodeSecurityGroupId` instead of `nodeSecurityGroup`
    - `clusterIngressRuleId` instead of `clusterIngressRule`

#### Cluster
| New Top-Level Property | <div style="width:9rem">Type</div> | Description | Default Value |
|------------------------|------|-------------|---------------|
| `clusterSecurityGroupId` | `Output<string>` | ID of the cluster security group | EKS-created security group ID |
| `nodeSecurityGroupId` | `Output<string>` | ID of the node security group | EKS-created security group ID |
| `clusterIngressRuleId` | `Output<string>` | ID of the cluster ingress rule. This is the rule that gives the nodes API server access | `""` |
| `defaultNodeGroupAsgName` | `Output<string>` | Name of the default node group's auto scaling group | `""` |
| `fargateProfileId` | `Output<string>` | ID of the Fargate profile if enabled | `""` |
| `fargateProfileStatus` | `Output<string>` | Status of the Fargate profile if enabled | `""` |
| `oidcProviderArn` | `Output<string>` | ARN of the OIDC provider if enabled | `""` |
| `oidcProviderUrl` | `Output<string>` | URL of the OIDC provider | n/a |
| `oidcIssuer` | `Output<string>` | URL of the OIDC provider without 'https://'. | n/a |

#### NodeGroup & NodeGroupV2
| New Top-Level Property | <div style="width:9rem">Type</div> | Description | Default Value |
|------------------------|------|-------------|---------------|
| `nodeSecurityGroupId`  | `Output<string>` | ID of the node security group | n/a |

## Miscellaneous changes

### NodeGroup & NodeGroupV2 changes

#### Input property types
The `NodeGroup` and `NodeGroupV2` components now accept inputs for the following input properties:
- `kubeletExtraArgs`
- `bootstrapExtraArgs`
- `labels`
- `taints`
- `nodeAssociatePublicIpAddress`

If you are using Go you will need to adjust your program to handle those types being inputs.

#### Security group and ingress rule inputs
We have updated the `NodeGroup` and `NodeGroupV2` components to allow passing in `nodeSecurityGroup` and `clusterIngressRule` by their IDs. You can now use `nodeSecurityGroupId` and `clusterIngressRuleId` as input properties.

Because of this change, the `nodeSecurityGroup` output property is now optional. To work around this, we have added a required `nodeSecurityGroupId` output property you can use instead. Use `nodegroup.nodeSecurityGroupId` instead of `nodegroup.nodeSecurityGroup.id`.

### Default Security Groups can now be disabled
If you do not need the default cluster and node security groups you can disable those now
with the `skipDefaultSecurityGroups` flag. Those security groups will not be created when setting this flag to true.

Because of this change, the `clusterSecurityGroup`, `nodeSecurityGroup` and `clusterIngressRule` properties are optional now. If you're using those outputs you'll need to update your code accordingly.

### Cluster does not create extraneous node IAM role if `skipDefaultNodeGroup` is set to `true`

Previously the Cluster component created a default node IAM role even if `skipDefaultNodeGroup` was set to `true`. This role gets correctly omitted now if you are specifying `skipDefaultNodeGroup`.
