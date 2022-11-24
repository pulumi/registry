---
title: AWSx (Pulumi Crosswalk for AWS)
meta_desc: Well-Architected Infrastructure as Code for AWS. The easiest way to AWS — from development to production.
layout: overview
---

{{% notes %}}
This AWSx documentation is generated using the v1.0.0-beta series of AWSx. If you have not upgraded to that version of
the provider then we suggest you continue to use our [@pulumi/awsx]({{<relref "/docs/reference/pkg/nodejs/pulumi/awsx">}}) NodeJS documentation.
{{% /notes %}}

The Amazon Web Services (AWS) [Crosswalk](../../../docs/guides/crosswalk/aws">}}) (AWSx) library uses automatic well-architected best practices to make common infrastructure-as-code tasks in [AWS](https://aws.amazon.com/) easier and more secure. It uses the AWS SDK to manage and provision resources.

The AWSx provider must be configured with credentials to deploy and update resources in AWS; see [Installation & Configuration]({{<relref "./installation-configuration">}}) for instructions.

**New to Pulumi and AWS?** [Get started with AWS using our AWSx tutorial]({{<relref "/docs/guides/crosswalk/aws">}}).

## Example

{{< chooser language "typescript,python" >}}

{{% choosable language typescript %}}

```typescript
import * as aws from "@pulumi/aws";
import * as awsx from "@pulumi/awsx";

const cluster = new aws.ecs.Cluster("default-cluster");

const lb = new awsx.lb.ApplicationLoadBalancer("nginx-lb");

const service = new awsx.ecs.FargateService("my-service", {
    cluster: cluster.arn,
    desiredCount: 2,
    taskDefinitionArgs: {
        container: {
            image: "nginx:latest",
            cpu: 512,
            memory: 128,
            essential: true,
            portMappings: [
                {
                    containerPort: 80,
                    targetGroup: lb.defaultTargetGroup,
                },
            ],
        },
    },
});
export const url = lb.loadBalancer.dnsName;
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_aws as aws
import pulumi_awsx as awsx

cluster = aws.ecs.Cluster("default-cluster")

lb = awsx.lb.ApplicationLoadBalancer("nginx-lb")

service = awsx.ecs.FargateService("my-service",
                                  cluster=cluster.arn,
                                  desired_count=2,
                                  task_definition_args=awsx.ecs.FargateServiceTaskDefinitionArgs(
                                      container=awsx.ecs.TaskDefinitionContainerDefinitionArgs(
                                          image="nginx:latest",
                                          cpu=512,
                                          memory=128,
                                          essential=True,
                                          port_mappings=[awsx.ecs.TaskDefinitionPortMappingArgs(
                                              target_group=lb.default_target_group
                                          )],
                                      )
                                  )
                                  )
pulumi.export("url", lb.load_balancer.dns_name)
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using Pulumi;
using Pulumi.Awsx.Ecr.Inputs;
using Pulumi.Awsx.Ecs.Inputs;
using Aws = Pulumi.Aws;
using Ecr = Pulumi.Awsx.Ecr;
using Ecs = Pulumi.Awsx.Ecs;
using Lb = Pulumi.Awsx.Lb;

await Deployment.RunAsync(() =>
{
    var cluster = new Aws.Ecs.Cluster("default-cluster");
    var lb = new Awsx.Lb.ApplicationLoadBalancer("nginx-lb");
    var service = new Awsx.Ecs.FargateService("my-service", new Awsx.Ecs.FargateServiceArgs
    {
        Cluster = cluster.Arn,
        DesiredCount = 2,
        TaskDefinitionArgs = new Awsx.Ecs.Inputs.FargateServiceTaskDefinitionArgs
        {
            Container = new Awsx.Ecs.Inputs.TaskDefinitionContainerDefinitionArgs
            {
                Image = "nginx:latest",
                Cpu = 512,
                Memory = 128,
                Essential = true,
                PortMappings =
                {
                    new Awsx.Ecs.Inputs.TaskDefinitionPortMappingArgs
                    {
                        TargetGroup = lb.DefaultTargetGroup,
                    }
                },
            }
        }
    });

    return new Dictionary<string, object?>
    {
        ["url"] = lb.LoadBalancer.Apply(lb => lb.DnsName)
    };
});
```

{{% /choosable %}}

{{% /chooser %}}
