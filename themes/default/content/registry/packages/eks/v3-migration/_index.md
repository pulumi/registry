---
title: Amazon EKS v3 Migration Guide
meta_desc: How to upgrade from v2 to v3 of the EKS provider
layout: package
---

In 2024 AWS announced the deprecation of two features used by default in Pulumi EKS: the aws-auth ConfigMap and the AL2 operating system. Pulumi EKS v3 addresses these deprecations, enhances the maintainability of the provider, and aligns it with EKS best practices.
There are some breaking changes which are detailed in this document.

## VPC CNI Component changes

The VPC CNI cluster component is now configured as an EKS addon as mentioned in the “New Features” section above. This brings the following changes:

- Removed `enableIpv6` input property. The component automatically configures the IP version now depending on whether the cluster is running in IPv4 or IPv6 mode.  
- Removed `image`, `initImage`, `nodeAgentImage` input properties. The component now automatically selects an image registry in the cluster’s region to pull the images from.

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
    securityGroupId: oldNG.nodeSecurityGroup.id,
    referencedSecurityGroupId: newNG.nodeSecurityGroup.id,
    fromPort: 0,
    ipProtocol: "-1",
    toPort: 0,
});
const newToOld = new aws.vpc.SecurityGroupIngressRule("newToOld", {
    securityGroupId: newNG.nodeSecurityGroup.id,
    referencedSecurityGroupId: oldNG.nodeSecurityGroup.id,
    fromPort: 0,
    ipProtocol: "-1",
    toPort: 0,
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
                                          security_group_id=old_ng.node_security_group.id,
                                          referenced_security_group_id=new_ng.node_security_group.id,
                                          from_port=0,
                                          ip_protocol='-1',
                                          to_port=0)
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
	SecurityGroupId:             oldNg.NodeSecurityGroup.ID(),
	SourceSecurityGroupId:       newNg.NodeSecurityGroup.ID(),
	FromPort:                    pulumi.Int(0),
	ToPort:                      pulumi.Int(0),
	Protocol:                    pulumi.String("-1"),
})
if err != nil {
	return err
}

_, err = ec2.NewSecurityGroupIngress(ctx, "newToOld", &ec2.SecurityGroupIngressArgs{
	SecurityGroupId:             newNg.NodeSecurityGroup.ID(),
	SourceSecurityGroupId:       oldNg.NodeSecurityGroup.ID(),
	FromPort:                    pulumi.Int(0),
	ToPort:                      pulumi.Int(0),
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
    SecurityGroupId = oldNG.NodeSecurityGroup.Id,
    ReferencedSecurityGroupId = newNG.NodeSecurityGroup.Id,
    FromPort = 0,
    IpProtocol = "-1",
    ToPort = 0,
});

var newToOld = new SecurityGroupIngressRule("newToOld", new()
{
    SecurityGroupId = newNG.NodeSecurityGroup.Id,
    ReferencedSecurityGroupId = oldNG.NodeSecurityGroup.Id,
    FromPort = 0,
    IpProtocol = "-1",
    ToPort = 0,
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
    .fromPort(0)
    .ipProtocol("-1")
    .toPort(0)
    .build()
);

var newToOld = new SecurityGroupIngressRule("newToOld", SecurityGroupIngressRuleArgs.builder()
    .securityGroupId(newNG.nodeSecurityGroupId())
    .referencedSecurityGroupId(oldNG.nodeSecurityGroupId())
    .fromPort(0)
    .ipProtocol("-1")
    .toPort(0)
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
A step-by-step migration guide can be found [here](https://github.com/pulumi/pulumi-eks/blob/master/docs/authentication-mode-migration.md).

We currently recommend users create all new clusters with the `API` authentication mode.
More details about this can be found [here](https://docs.aws.amazon.com/eks/latest/userguide/grant-k8s-access.html).

## Nodejs SDK changes

The Nodejs SDK is updated to use state of the art Pulumi tooling, improving stability, documentation and security. The update requires the following changes to programs:

- Properties of the components are now outputs instead of plain types. Notable cases where you need to change your program are:
  - The [`Cluster::getKubeConfig`](https://www.pulumi.com/registry/packages/eks/api-docs/cluster/#method_GetKubeconfig) method now returns an output.  
  - Creating an IRSA based IAM role now requires you to use `apply` for accessing the OIDC provider ARN and URL. An example of how this works can be found [here](https://github.com/pulumi/pulumi-eks/blob/release-3.x.x/examples/oidc-iam-sa/index.ts).  
  - `k8s.Provider` will be deleted if not referenced (no impact, but it will appear in the diff)
  - `clusterOidcProvider` is an output now. `getKubeConfig` returns an output now  
- The deprecated input property `deployDashboard` of the `Cluster` component has been removed from the Nodejs SDK. This has already been removed from the other SDKs in the past. If you’d like to continue using it, you can adopt the existing code into your own program from [here](https://github.com/pulumi/pulumi-eks/blob/bcc170e72b802a78e7f0a99bc92316a5f8a62b0e/nodejs/eks/dashboard.ts).
- The `createManagedNodeGroup` function will now create an Pulumi EKS `ManagedNodeGroup` instead of creating the underlying `aws.eks.NodeGroup` resource directly. During the upgrade to Pulumi EKS v3 you'll see the additional wrapper component being created.

## Miscellaneous changes

### NodeGroup & NodeGroupV2 accept inputs for its properties
The `NodeGroup` and `NodeGroupV2` components now accept inputs for the following input properties:
- `kubeletExtraArgs`
- `bootstrapExtraArgs`
- `labels`
- `taints`
- `nodeAssociatePublicIpAddress`

If you are using Go you will need to adjust your program to handle those types being inputs.

### Default Security Groups can now be disabled
If you do not need the default cluster and node security groups you can disable those now
with the `skipDefaultSecurityGroups` flag. Those security groups will not be created when setting that flag to true.

Because of this change, the `clusterSecurityGroup`, `nodeSecurityGroup` and `clusterIngressRule` properties are optional now. If you're using those outputs you'll need to update your code accordingly.

### Cluster does not create extraneous node IAM role if `skipDefaultNodeGroup` is set to `true`

Previously the Cluster component created a default node IAM role even if `skipDefaultNodeGroup` was set to `true`. This role gets correctly omitted now if you are specifying `skipDefaultNodeGroup`.
