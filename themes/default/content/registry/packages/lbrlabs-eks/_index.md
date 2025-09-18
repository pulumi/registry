---
# WARNING: this file was fetched from https://raw.githubusercontent.com/lbrlabs/pulumi-lbrlabs-eks/v1.5.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: LBr Labs EKS
meta_desc: Provides an overview of the batteries include EKS provider for Pulumi by LBr Labs.
layout: package
---

The LBr Labs EKS component for Pulumi can be used to provision a batteries include EKS cluster.

It provides resources to create a functioning EKS cluster with several Kubernetes addons and operators which will allow you to be quickly successful with Amazon EKS such as:

- an EKS cluster with [CloudTrail logging enabled](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html), and [secret encryption enabled](https://docs.aws.amazon.com/eks/latest/userguide/enable-kms.html)
- A node group designed to run "system" workloads with a taint
- The [AWS EBS CSI Addon](https://docs.aws.amazon.com/eks/latest/userguide/ebs-csi.html) with an [IAM role using IAM Roles for Service accounts](https://docs.aws.amazon.com/eks/latest/userguide/iam-roles-for-service-accounts.html)
- The [AWS VPC CBI Addon](https://docs.aws.amazon.com/eks/latest/userguide/managing-vpc-cni.html) with an [IAM role using IAM Roles for Service accounts](https://docs.aws.amazon.com/eks/latest/userguide/iam-roles-for-service-accounts.html)
- The CoreDNS Addon
- An internal and external [NGINX ingress controller](https://github.com/kubernetes/ingress-nginx) with relevant ingress classes
- [External DNS](https://github.com/kubernetes-sigs/external-dns), with an IAM role that can modify Route53 domains using IAM Roles for service accounts
- [Cert Manager](https://cert-manager.io/), with an IAM role that can modify Route53 domains using IAM Roles for service accounts

In addition to the cluster control plane and system worker nodes, it also provides an SDK to create attached worker nodes that can be used to schedule Kubernetes workloads.

## Example

{{< chooser language "typescript,python,csharp,go,yaml" >}}
{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as awsx from "@pulumi/awsx";
import * as kubernetes from "@pulumi/kubernetes";
import * as lbrlabs_eks from "@lbrlabs/pulumi-eks";

const vpc = new awsx.ec2.Vpc("vpc", {
  cidrBlock: "172.16.0.0/16",
  subnetSpecs: [
    {
      type: awsx.ec2.SubnetType.Public,
      tags: {
        "kubernetes.io/role/elb": "1",
      },
    },
    {
      type: awsx.ec2.SubnetType.Private,
      tags: {
        "kubernetes.io/role/internal-elb": "1",
      },
    },
  ],
});
const cluster = new lbrlabs_eks.Cluster("cluster", {
  clusterSubnetIds: vpc.privateSubnetIds,
  letsEncryptEmail: "mail@lbrlabs.com",
  systemNodeSubnetIds: vpc.publicSubnetIds,
  systemNodeInstanceTypes: ["t3.large"],
  systemNodeDesiredCount: 4,
});

const workloadNodes = new lbrlabs_eks.AttachedNodeGroup("workload", {
  clusterName: cluster.controlPlane.name,
  subnetIds: vpc.privateSubnetIds,
  scalingConfig: {
    desiredSize: 3,
    maxSize: 10,
    minSize: 1,
  },
});

const roleMapping = new lbrlabs_eks.IamRoleMapping("roleMapping", {
  roleArn: workloadNodes.nodeRole.arn,
  username: "system:node:{{EC2PrivateDNSName}}",
  grups: ["system:bootstrappers", "system:nodes"],
})

const provider = new kubernetes.Provider("provider", {
  kubeconfig: cluster.kubeconfig,
});

const wordpress = new kubernetes.helm.v3.Release(
  "wordpress",
  {
    chart: "wordpress",
    repositoryOpts: {
      repo: "https://charts.bitnami.com/bitnami",
    },
    values: {
      wordpressUsername: "lbrlabs",
      wordpressPassword: "correct-horse-battery-stable",
      wordpressEmail: "mail@lbrlabs.com",
      ingress: {
        enabled: true,
        ingressClassName: "external",
        hostname: "wordpress.aws.briggs.work",
        tls: true,
        annotations: {
          "cert-manager.io/cluster-issuer": "letsencrypt-prod",
          "nginx.ingress.kubernetes.io/force-ssl-redirect": "true",
        },
      },
    },
  },
  {
    provider: provider,
  }
);
export const kubeconfig = cluster.kubeconfig;
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
import lbrlabs_pulumi_eks as lbrlabs_eks
import pulumi_aws as aws
import pulumi_awsx as awsx
import pulumi_kubernetes as kubernetes

vpc = awsx.ec2.Vpc(
    "vpc",
    cidr_block="172.16.0.0/16",
    subnet_specs=[
        awsx.ec2.SubnetSpecArgs(
            type=awsx.ec2.SubnetType.PUBLIC,
            tags={
                "kubernetes.io/role/elb": "1",
            },
        ),
        awsx.ec2.SubnetSpecArgs(
            type=awsx.ec2.SubnetType.PRIVATE,
            tags={
                "kubernetes.io/role/internal-elb": "1",
            },
        ),
    ],
)
cluster = lbrlabs_eks.Cluster(
    "cluster",
    cluster_subnet_ids=vpc.public_subnet_ids,
    system_node_subnet_ids=vpc.private_subnet_ids,
    system_node_instance_types=["t3.large"],
    system_node_desired_count=4,
    lets_encrypt_email="mail@lbrlabs.com",
)

workload_nodes = lbrlabs_eks.AttachedNodeGroup(
    "workloadNodes",
    cluster_name=cluster.control_plane.name,
    subnet_ids=vpc.private_subnet_ids,
    scaling_config=aws.eks.NodeGroupScalingConfigArgs(
        desired_size=4,
        max_size=10,
        min_size=1,
    ),
)

roleMapping = lbrlabs_eks.IamRoleMapping(
    "workload",
    role_arn=workload_nodes.node_role.arn,
    username="system:node:{{EC2PrivateDNSName}}",
    groups=["system:bootstrappers", "system:nodes"],
)


provider = kubernetes.Provider("provider", kubeconfig=cluster.kubeconfig)
wordpress = kubernetes.helm.v3.Release(
    "wordpress",
    chart="wordpress",
    repository_opts=kubernetes.helm.v3.RepositoryOptsArgs(
        repo="https://charts.bitnami.com/bitnami",
    ),
    values={
        "wordpressUsername": "lbrlabs",
        "wordpressPassword": "correct-horse-battery-stable",
        "wordpressEmail": "mail@lbrlabs.com",
        "ingress": {
            "enabled": True,
            "ingressClassName": "external",
            "hostname": "wordpress.aws.briggs.work",
            "tls": True,
            "annotations": {
                "cert-manager.io/cluster-issuer": "letsencrypt-prod",
                "nginx.ingress.kubernetes.io/force-ssl-redirect": "true",
            },
        },
    },
    opts=pulumi.ResourceOptions(provider=provider),
)
pulumi.export("kubeconfig", cluster.kubeconfig)

```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Awsx = Pulumi.Awsx;
using Aws = Pulumi.Aws;
using Kubernetes = Pulumi.Kubernetes;
using LbrlabsEks = Lbrlabs.PulumiPackage.LbrlabsEks;

return await Deployment.RunAsync(() =>
{
    var vpc = new Awsx.Ec2.Vpc("vpc", new()
    {
        CidrBlock = "172.16.0.0/16",
        SubnetSpecs = new List<Awsx.Ec2.Inputs.SubnetSpecArgs>
        {
            new Awsx.Ec2.Inputs.SubnetSpecArgs
            {
                Type = Awsx.Ec2.SubnetType.Public,
                Tags = new InputMap<string>
                {
                    { "kubernetes.io/role/elb", "1" },
                },
            },
            new Awsx.Ec2.Inputs.SubnetSpecArgs
            {
                Type = Awsx.Ec2.SubnetType.Private,
                Tags = new InputMap<string>
                {
                    { "kubernetes.io/role/internal-elb", "1" },
                },
            }
        }
    });

    var cluster = new LbrlabsEks.Cluster("cluster", new()
    {
        ClusterSubnetIds = vpc.PublicSubnetIds,
        SystemNodeSubnetIds = vpc.PrivateSubnetIds,
        LetsEncryptEmail = "mail@lbrlabs.com",
        SystemNodeInstanceTypes = new[]
        {
            "t3.large",
        },
        SystemNodeDesiredCount = 4,
    });

    var workloadNodes = new LbrlabsEks.AttachedNodeGroup("workloadNodes", new()
    {
        ClusterName = cluster.ControlPlane.Apply(controlPlane => controlPlane.Name),
        SubnetIds = vpc.PrivateSubnetIds,
        // ScalingConfig = new Aws.Eks.Inputs.NodeGroupScalingConfigArgs
        // {
        //     DesiredSize = 4,
        //     MaxSize = 4,
        //     MinSize = 4,
        // },
    });

    var roleMapping = new LbrlabsEks.IamRoleMapping("workload", new()
    {
        RoleArn = workloadNodes.NodeRole.Arn,
        Username = "system:node:{{EC2PrivateDNSName}}",
        Groups = new[]
        {
            "system:bootstrappers",
            "system:nodes",
        },
    });

    var provider = new Kubernetes.Provider("provider", new()
    {
        KubeConfig = cluster.Kubeconfig,
    });

    var wordpress = new Kubernetes.Helm.V3.Release("wordpress", new()
    {
        Chart = "wordpress",
        RepositoryOpts = new Kubernetes.Types.Inputs.Helm.V3.RepositoryOptsArgs
        {
            Repo = "https://charts.bitnami.com/bitnami",
        },
        Values = new Dictionary<string, object>
        {
            ["ingress"] = new Dictionary<string, object>
            {
                ["annotations"] = new Dictionary<string, object>
                {
                    ["cert-manager.io/cluster-issuer"] = "letsencrypt-prod",
                    ["nginx.ingress.kubernetes.io/force-ssl-redirect"] = "true",
                },
                ["enabled"] = true,
                ["ingressClassName"] = "external",
                ["hostname"] = "wordpress.aws.briggs.work",
                ["tls"] = true,
            },
            ["wordpressUsername"] = "lbrlabs",
            ["wordpressPassword"] = "correct-horse-battery-stable",
            ["wordpressEmail"] = "mail@lbrlabs.com",
        }
    }, new CustomResourceOptions
    {
        Provider = provider,
    });

    return new Dictionary<string, object?>
    {
        ["kubeconfig"] = cluster.Kubeconfig,
    };
});


});
```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"fmt"

	lbrlabs "github.com/lbrlabs/pulumi-lbrlabs-eks/sdk/go/eks"
	"github.com/pulumi/pulumi-awsx/sdk/go/awsx/ec2"
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes"
	helm "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/helm/v3"
	eks "github.com/pulumi/pulumi-aws/sdk/v6/go/aws/eks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create an AWS resource (S3 Bucket)

		cidrBlock := "172.16.0.0/16"
		vpc, err := ec2.NewVpc(ctx, "vpc", &ec2.VpcArgs{
			CidrBlock: &cidrBlock,
			SubnetSpecs: []ec2.SubnetSpecArgs{
				{
					Type: ec2.SubnetTypePublic,
				},
				{
					Type: ec2.SubnetTypePrivate,
				},
			},
		})
		if err != nil {
			return fmt.Errorf("error creating VPC")
		}

		cluster, err := lbrlabs.NewCluster(ctx, "cluster", &lbrlabs.ClusterArgs{
			ClusterSubnetIds:    vpc.PublicSubnetIds,
			SystemNodeSubnetIds: vpc.PrivateSubnetIds,
			SystemNodeInstanceTypes: pulumi.StringArray{
				pulumi.String("t3.large"),
			},
			SystemNodeDesiredCount: pulumi.Float64Ptr(4),
			LetsEncryptEmail:       pulumi.String("mail@lbrlabs.com"),
		})

		workloadNodes, err := lbrlabs.NewAttachedNodeGroup(ctx, "workloadNodes", &lbrlabs.AttachedNodeGroupArgs{
			ClusterName: cluster.ControlPlane.Name(),
			SubnetIds:   vpc.PrivateSubnetIds,
			ScalingConfig: &eks.NodeGroupScalingConfigArgs{
				DesiredSize: pulumi.Int(4),
				MaxSize:     pulumi.Int(10),
				MinSize:     pulumi.Int(1),
			},
		})

		_, err = lbrlabs.NewRoleMapping(ctx, "roleMapping", &lbrlabs.RoleMappingArgs{
			RoleArn:  workloadNodes.NodeRole.Arn(),
			Username: pulumi.String("system:node:{{EC2PrivateDNSName}}"),
			Groups: pulumi.StringArray{
				pulumi.String("system:bootstrappers"),
				pulumi.String("system:nodes"),
			},
		})

		provider, err := kubernetes.NewProvider(ctx, "provider", &kubernetes.ProviderArgs{
			Kubeconfig: cluster.Kubeconfig,
		})

		_, err = helm.NewRelease(ctx, "wordpress", &helm.ReleaseArgs{
			Chart: pulumi.String("wordpress"),
			RepositoryOpts: &helm.RepositoryOptsArgs{
				Repo: pulumi.String("https://charts.bitnami.com/bitnami"),
			},
			Values: pulumi.Map{
				"wordpressUsername": pulumi.String("user"),
				"wordpressPassword": pulumi.String("correct-horse-battery-stable"),
				"ingress": pulumi.Map{
					"enabled":          pulumi.Bool(true),
					"ingressClassName": pulumi.String("external"),
					"hostname":         pulumi.String("wordpress.aws.briggs.work"),
					"tls":              pulumi.Bool(true),
					"annotations": pulumi.Map{
						"cert-manager.io/cluster-issuer":                 pulumi.String("letsencrypt-prod"),
						"nginx.ingress.kubernetes.io/force-ssl-redirect": pulumi.String("true"),
					},
				},
			},
		}, pulumi.Provider(provider))

		ctx.Export("vpcId", vpc.VpcId)

		return nil
	})
}

```
{{% /choosable %}}

{{% choosable language yaml %}}
name: lbrlabs-eks-yaml
runtime: yaml
description: A minimal Pulumi YAML program
resources:
  vpc:
    type: awsx:ec2:Vpc
    properties:
      cidrBlock: "172.16.0.0/16"
      subnetSpecs:
        - type: Public
          tags:
            "kubernetes.io/role/elb": "1"
        - type: Private
          tags:
            "kubernetes.io/role/internal-elb": "1"
  cluster:
    type: lbrlabs-eks:index:Cluster
    properties:
      clusterSubnetIds: ${vpc.publicSubnetIds}
      systemNodeSubnetIds: ${vpc.privateSubnetIds}
      systemNodeInstanceTypes: [ "t3.large" ]
      systemNodeDesiredCount: 4
      letsEncryptEmail: "mail@lbrlabs.com"
  workloadNodes:
    type: lbrlabs-eks:index:AttachedNodeGroup
    properties:
      clusterName: ${cluster.controlPlane.name}
      subnetIds: ${vpc.privateSubnetIds}
      scalingConfig:
        desiredSize: 3
        maxSize: 10
        minSize: 1
  rolemapping:
    type: lbrlabs-eks:index:IamRoleMapping
    properties:
      roleArn: ${workloadNodes.nodeRole.arn}
      username: system:node:{{EC2PrivateDNSName}}
      groups:
        - system:bootstrappers
        - system:nodes
  provider:
    type: pulumi:providers:kubernetes
    properties:
      kubeconfig: ${cluster.kubeconfig}
  wordpress:
    type: kubernetes:helm.sh/v3:Release
    properties:
      chart: wordpress
      repositoryOpts:
        repo: https://charts.bitnami.com/bitnami
      values:
        wordpressUsername: lbrlabs
        wordpressPassword: correct-horse-battery-stable
        wordpressEmail: mail@lbrlabs.com
        ingress:
          enabled: true
          ingressClassName: external
          hostname: wordpress.aws.briggs.work # this domain must in exist in route53
          tls: true
          annotations:
            "cert-manager.io/cluster-issuer": "letsencrypt-prod"
            "nginx.ingress.kubernetes.io/force-ssl-redirect": "true"
    options:
      provider: ${provider}
outputs:
  kubeconfig: ${cluster.kubeconfig}
{{% /choosable %}}
{{< /chooser >}}
