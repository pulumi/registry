---
title: AWSx
meta_desc: Well-Architected Infrastructure as Code for AWS. The easiest way to AWS — from development to production.
layout: overview
---

The Amazon Web Services (AWS) Crosswalk (AWSx) provider for Pulumi can provision many of the cloud resources available in [AWS](https://aws.amazon.com/). It uses the AWS SDK to manage and provision resources.

The AWSx provider must be configured with credentials to deploy and update resources in AWS; see [Installation & Configuration]({{<relref "./installation-configuration">}}) for instructions.

**New to Pulumi and AWS?** [Get started with AWS using our AWSx tutorial]({{<relref "docs/guides/crosswalk/aws/">}}).

## Example

{{< chooser language "typescript,python" >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as awsx from "@pulumi/awsx";

const vpc = awsx.ec2.Vpc.getDefault();

const cluster = new awsx.ecs.Cluster("cluster", {
    capacityProviders: ["FARGATE_SPOT"],
    defaultCapacityProviderStrategies: [
        {
            capacityProvider: "FARGATE_SPOT",
            weight: 1,
        },
    ],
});

// // Create a load balancer on port 80 and spin up two instances of Nginx.
const lb = new awsx.lb.ApplicationLoadBalancer("nginx-lb", {
    subnetIds: vpc.publicSubnetIds,
});

const service = new awsx.ecs.FargateService("my-service", {
    cluster: cluster.cluster.arn,
    taskDefinitionArgs: {
        container: {
            image: "nginx:latest",
            cpu: 512,
            memory: 128,
            essential: true,
            portMappings: [{ targetGroup: lb.defaultTargetGroup }],
        },
    },
    networkConfiguration: {
        subnets: vpc.publicSubnetIds,
        assignPublicIp: true,
    },
});

// Export the load balancer's address so that it's easy to access.
export const url = lb.loadBalancer.dnsName;
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_awsx as awsx

vpc = awsx.ec2.Vpc.get_default()

cluster = awsx.ecs.Cluster("cluster",
                           capacity_providers=["FARGATE_SPOT"],
                           default_capacity_provider_strategies=[awsx.ecs.DefaultCapacityProviderStrategyArgs(
                               capacity_provider="FARGATE_SPOT",
                               weight=1
                           )])

lb = awsx.lb.ApplicationLoadBalancer("nginx-lb", subnet_ids=vpc.public_subnet_ids)

service = awsx.ecs.FargateService("my-service",
                                  cluster=cluster.cluster.arn,
                                  task_definition_args=awsx.ecs.TaskDefinitionArgs(
                                      container=awsx.ecs.ContainerArgs(
                                          image="nginx:latest",
                                          cpu=512,
                                          memory=128,
                                          essential=True,
                                          port_mappings=[awsx.ecs.PortMappingArgs(
                                              target_group=lb.default_target_group
                                          )],
                                  )),
                                  network_configuration=awsx.ecs.NetworkConfigArgs(
                                      subnets=vpc.public_subnet_ids,
                                      assign_public_ip=True,
                                  ))

pulumi.export("url", lb.load_balancer.dns_name)
```

{{% /choosable %}}

{{% /chooser %}}
