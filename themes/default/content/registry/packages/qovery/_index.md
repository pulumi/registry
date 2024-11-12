---
# WARNING: this file was fetched from https://raw.githubusercontent.com/dirien/pulumi-qovery/v0.41.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Qovery
meta_desc: Provides an overview of the Qovery Provider for Pulumi.
layout: package
---

The Qovery Resource Provider lets you manage [Qovery](https://qovery.com/) resources.


## Example

Some example code for `typescript`,`python` and `go`. Feel free to contribute examples for `csharp`, `javascript` or `java`

{{< chooser language "typescript,python,go" >}}


{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as qovery from "@ediri/qovery";

const config = new pulumi.Config();
const organizationId = config.get("organizationId") || "changeme";
const awsCredentials = new qovery.AwsCredentials("awsCredentials", {
    organizationId: organizationId,
    accessKeyId: config.requireSecret("accessKeyId"),
    secretAccessKey: config.requireSecret("secretAccessKey"),
});
const myCluster = new qovery.Cluster("myCluster", {
    organizationId: organizationId,
    name: "TS Demo Cluster",
    description: "Pulumi TS Demo Cluster",
    cloudProvider: "AWS",
    region: "eu-central-1",
    instanceType: "t3a.medium",
    minRunningNodes: 3,
    maxRunningNodes: 4,
    credentialsId: awsCredentials.id,
    state: "DEPLOYED",
});
const myProject = new qovery.Project("myProject", {
    organizationId: organizationId,
    name: "Demo Project",
    description: "Pulumi Demo Project",
});
const prodEnvironment = new qovery.Environment("prodEnvironment", {
    name: "Production",
    mode: "PRODUCTION",
    projectId: myProject.id,
    clusterId: myCluster.id,
});
const helloServerDemoApp = new qovery.Application("helloServerDemoApp", {
    name: "hello-server-demo-app",
    environmentId: prodEnvironment.id,
    gitRepository: {
        url: "https://github.com/dirien/hello-server.git",
        branch: "main",
        rootPath: "/golang",
    },
    cpu: 500,
    memory: 256,
    buildMode: "DOCKER",
    healthchecks: {
        livenessProbe: {
            type: {
                http: {
                    port: 8080,
                    path: "/",
                },
            },
            initialDelaySeconds: 30,
            periodSeconds: 10,
            timeoutSeconds: 10,
            successThreshold: 1,
            failureThreshold: 3,
        },
    },
    dockerfilePath: "Dockerfile",
    minRunningInstances: 1,
    maxRunningInstances: 1,
    ports: [{
        internalPort: 8080,
        externalPort: 443,
        protocol: "HTTP",
        publiclyAccessible: true,
        isDefault: true,
    }],
});
const helloServerDemoAppDeployment = new qovery.Deployment("helloServerDemoAppDeployment", {
    environmentId: prodEnvironment.id,
    desiredState: "RUNNING",
}, {
    dependsOn: [helloServerDemoApp],
});
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import ediri_qovery as qovery

config = pulumi.Config()
organization_id = config.get("organizationId")
access_key_id = config.require("accessKeyId")
secret_access_key= config.require("secretAccessKey")
if organization_id is None:
    organization_id = "e53e1c45-7449-4f44-8c59-ee0b2d3305de"
aws_credentials = qovery.AwsCredentials("awsCredentials",
    organization_id=organization_id,
    access_key_id=access_key_id,
    secret_access_key=secret_access_key)
my_cluster = qovery.Cluster("myCluster",
    organization_id=organization_id,
    name="Python Demo Cluster",
    description="Pulumi Python Demo Cluster",
    cloud_provider="AWS",
    region="eu-central-1",
    instance_type="t3a.medium",
    min_running_nodes=3,
    max_running_nodes=4,
    credentials_id=aws_credentials.id,
    state="DEPLOYED")
my_project = qovery.Project("myProject",
    organization_id=organization_id,
    name="Demo Project",
    description="Pulumi Demo Project")
prod_environment = qovery.Environment("prodEnvironment",
    name="Production",
    mode="PRODUCTION",
    project_id=my_project.id,
    cluster_id=my_cluster.id)
hello_server_demo_app = qovery.Application("helloServerDemoApp",
    name="hello-server-demo-app",
    environment_id=prod_environment.id,
    git_repository=qovery.ApplicationGitRepositoryArgs(
        url="https://github.com/dirien/hello-server.git",
        branch="main",
        root_path="/golang",
    ),
    cpu=500,
    memory=256,
    build_mode="DOCKER",
    healthchecks=qovery.ApplicationHealthchecksArgs(
        liveness_probe=qovery.ApplicationHealthchecksLivenessProbeArgs(
            type=qovery.ApplicationHealthchecksLivenessProbeTypeArgs(
                http=qovery.ApplicationHealthchecksLivenessProbeTypeHttpArgs(
                    port=8080,
                    path="/",
                ),
            ),
            initial_delay_seconds=30,
            period_seconds=10,
            timeout_seconds=10,
            success_threshold=1,
            failure_threshold=3,
        ),
    ),
    dockerfile_path="Dockerfile",
    min_running_instances=1,
    max_running_instances=1,
    ports=[qovery.ApplicationPortArgs(
        internal_port=8080,
        external_port=443,
        protocol="HTTP",
        publicly_accessible=True,
        is_default=True,
    )])
hello_server_demo_app_deployment = qovery.Deployment("helloServerDemoAppDeployment",
    environment_id=prod_environment.id,
    desired_state="RUNNING",
    opts=pulumi.ResourceOptions(depends_on=[hello_server_demo_app]))
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"github.com/dirien/pulumi-qovery/sdk/go/qovery"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		organizationId := "fff"
		if param := cfg.Get("organizationId"); param != "" {
			organizationId = param
		}
		awsCredentials, err := qovery.NewAwsCredentials(ctx, "awsCredentials", &qovery.AwsCredentialsArgs{
			OrganizationId:  pulumi.String(organizationId),
			AccessKeyId:     cfg.GetSecret("accessKeyId"),
			SecretAccessKey: cfg.GetSecret("secretAccessKey"),
		})
		if err != nil {
			return err
		}
		myCluster, err := qovery.NewCluster(ctx, "myCluster", &qovery.ClusterArgs{
			OrganizationId:  pulumi.String(organizationId),
			Name:            pulumi.String("Go Demo Cluster"),
			Description:     pulumi.String("Pulumi Go Demo Cluster"),
			CloudProvider:   pulumi.String("AWS"),
			Region:          pulumi.String("eu-central-1"),
			InstanceType:    pulumi.String("t3a.medium"),
			MinRunningNodes: pulumi.Int(3),
			MaxRunningNodes: pulumi.Int(4),
			CredentialsId:   awsCredentials.ID(),
			State:           pulumi.String("DEPLOYED"),
		})
		if err != nil {
			return err
		}
		myProject, err := qovery.NewProject(ctx, "myProject", &qovery.ProjectArgs{
			OrganizationId: pulumi.String(organizationId),
			Name:           pulumi.String("Demo Project"),
			Description:    pulumi.String("Pulumi Demo Project"),
		})
		if err != nil {
			return err
		}
		prodEnvironment, err := qovery.NewEnvironment(ctx, "prodEnvironment", &qovery.EnvironmentArgs{
			Name:      pulumi.String("Production"),
			Mode:      pulumi.String("PRODUCTION"),
			ProjectId: myProject.ID(),
			ClusterId: myCluster.ID(),
		})
		if err != nil {
			return err
		}
		helloServerDemoApp, err := qovery.NewApplication(ctx, "helloServerDemoApp", &qovery.ApplicationArgs{
			Name:          pulumi.String("hello-server-demo-app"),
			EnvironmentId: prodEnvironment.ID(),
			GitRepository: &qovery.ApplicationGitRepositoryArgs{
				Url:      pulumi.String("https://github.com/dirien/hello-server.git"),
				Branch:   pulumi.String("main"),
				RootPath: pulumi.String("/golang"),
			},
			Cpu:       pulumi.Int(500),
			Memory:    pulumi.Int(256),
			BuildMode: pulumi.String("DOCKER"),
			Healthchecks: &qovery.ApplicationHealthchecksArgs{
				LivenessProbe: &qovery.ApplicationHealthchecksLivenessProbeArgs{
					Type: &qovery.ApplicationHealthchecksLivenessProbeTypeArgs{
						Http: &qovery.ApplicationHealthchecksLivenessProbeTypeHttpArgs{
							Port: pulumi.Int(8080),
							Path: pulumi.String("/"),
						},
					},
					InitialDelaySeconds: pulumi.Int(30),
					PeriodSeconds:       pulumi.Int(10),
					TimeoutSeconds:      pulumi.Int(10),
					SuccessThreshold:    pulumi.Int(1),
					FailureThreshold:    pulumi.Int(3),
				},
			},
			DockerfilePath:      pulumi.String("Dockerfile"),
			MinRunningInstances: pulumi.Int(1),
			MaxRunningInstances: pulumi.Int(1),
			Ports: qovery.ApplicationPortArray{
				&qovery.ApplicationPortArgs{
					InternalPort:       pulumi.Int(8080),
					ExternalPort:       pulumi.Int(443),
					Protocol:           pulumi.String("HTTP"),
					PubliclyAccessible: pulumi.Bool(true),
					IsDefault:          pulumi.Bool(true),
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = qovery.NewDeployment(ctx, "helloServerDemoAppDeployment", &qovery.DeploymentArgs{
			EnvironmentId: prodEnvironment.ID(),
			DesiredState:  pulumi.String("RUNNING"),
		}, pulumi.DependsOn([]pulumi.Resource{
			helloServerDemoApp,
		}))
		if err != nil {
			return err
		}
		return nil
	})
}
```

{{% /choosable %}}

{{< /chooser >}}
