---
title: Metabase
meta_desc: Use Pulumi's Metabase Component to quickly provision and get started with Metabase.
layout: overview
---

You can use the Pulumi Metabase Component to help you get started quickly running Metabase in the cloud.
Currently this Component only supports AWS. The code below will show you examples of each resources
supported in this Component, but please refer to the API Docs for more detailed descriptions and
information about each resource.

## Quick Start

The following steps will get you up and running with Metabase on AWS with very little effort. Once you
have completed the steps you will have an RDS Database, Metabase running in a Fargate Task, and a Load
Balancer.

### Configure Environment

Before you get started using Pulumi, let's run through a few quick steps to ensure your environment is set up correctly.

### Install Pulumi

{{< install-pulumi >}}
{{% notes "info" %}}
All Windows examples in this tutorial assume you are running in PowerShell.
{{% /notes %}}
{{< /install-pulumi >}}

Next, install the required language runtime, if you have not already.

### Install Language Runtime

#### Choose Your Language

{{< chooser language "typescript,python,go,csharp,yaml" / >}}

{{% choosable language "typescript" %}}
{{< install-node >}}
{{% /choosable %}}

{{% choosable language python %}}
{{< install-python >}}
{{% /choosable %}}

{{% choosable language go %}}
{{< install-go >}}
{{% /choosable %}}

{{% choosable language "csharp,fsharp,visualbasic" %}}
{{< install-dotnet >}}
{{% /choosable %}}

{{% choosable language "yaml" %}}
{{< install-yaml >}}
{{% /choosable %}}

### Configure Pulumi to access your AWS account

Pulumi requires cloud credentials to manage and provision resources. You must use an IAM user account that has **Programmatic access** with rights to deploy and manage resources handled through Pulumi.

If you have previously <a href="https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-install.html" target="_blank">installed</a> and <a href="https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html" target="_blank">configured</a> the AWS CLI, Pulumi will respect and use your configuration settings.

If you do not have the AWS CLI installed or plan on using Pulumi from within a CI/CD pipeline, <a href="https://docs.aws.amazon.com/general/latest/gr/aws-sec-cred-types.html#access-keys-and-secret-access-keys" target="_blank">retrieve your access key ID and secret access key</a> and then set the `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` environment variables on your workstation.

{{< chooser os "linux,macos,windows" >}}
{{% choosable os linux %}}

```bash
$ export AWS_ACCESS_KEY_ID=<YOUR_ACCESS_KEY_ID>
$ export AWS_SECRET_ACCESS_KEY=<YOUR_SECRET_ACCESS_KEY>
```

{{% /choosable %}}

{{% choosable os macos %}}

```bash
$ export AWS_ACCESS_KEY_ID=<YOUR_ACCESS_KEY_ID>
$ export AWS_SECRET_ACCESS_KEY=<YOUR_SECRET_ACCESS_KEY>
```

{{% /choosable %}}

{{% choosable os windows %}}

```powershell
> $env:AWS_ACCESS_KEY_ID = "<YOUR_ACCESS_KEY_ID>"
> $env:AWS_SECRET_ACCESS_KEY = "<YOUR_SECRET_ACCESS_KEY>"
```

{{% /choosable %}}
{{< /chooser >}}

For additional information on setting and using AWS credentials, see [AWS Setup]({{< relref "/registry/packages/aws/installation-configuration" >}}).

Next, you'll create a new Pulumi project.

### Create New Project

Now that you have set up your environment by installing Pulumi, installing your preferred language runtime,
and configuring your AWS credentials, let's create your first Pulumi program.

{{< chooser language "typescript,python,go,csharp,yaml" / >}}
{{% choosable language typescript %}}

```bash
$ mkdir metabase-quickstart && cd metabase-quickstart
$ pulumi new aws-typescript
```

{{% /choosable %}}
{{% choosable language python %}}

```bash
$ mkdir metabase-quickstart && cd metabase-quickstart
$ pulumi new aws-python
```

{{% /choosable %}}
{{% choosable language go %}}

```bash
# from within your $GOPATH
$ mkdir metabase-quickstart && cd metabase-quickstart
$ pulumi new aws-go
```

{{% /choosable %}}
{{% choosable language csharp %}}

```bash
$ mkdir metabase-quickstart && cd metabase-quickstart
$ pulumi new aws-csharp
```

{{% /choosable %}}
{{% choosable language yaml %}}

```bash
$ mkdir metabase-quickstart && cd metabase-quickstart
$ pulumi new aws-yaml
```

{{% /choosable %}}

The [`pulumi new`]({{< relref "/docs/reference/cli/pulumi_new" >}}) command creates a new Pulumi project with some basic scaffolding based on the cloud and language specified.

{{< cli-note >}}

After logging in, the CLI will proceed with walking you through creating a new project.

First, you will be asked for a project name and description. Hit `ENTER` to accept the default values or specify new values.

Next, you will be asked for the name of a stack. Hit `ENTER` to accept the default value of `dev`.

Finally, you will be prompted for some configuration values for the stack. For AWS projects, you will be prompted for the AWS region. You can accept the default value or choose another value like `us-west-2`.

> What are [projects]({{< relref "/docs/intro/concepts/project" >}}) and [stacks]({{< relref "/docs/intro/concepts/stack" >}})? Pulumi projects and stacks let you organize Pulumi code. Consider a Pulumi _project_ to be analogous to a GitHub repo---a single place for code---and a _stack_ to be an instance of that code with a separate configuration. For instance, _Project Foo_ may have multiple stacks for different development environments (Dev, Test, or Prod), or perhaps for different cloud configurations (geographic region for example). See [Organizing Projects and Stacks]({{< relref "/docs/guides/organizing-projects-stacks" >}}) for some best practices on organizing your Pulumi projects and stacks.

{{% choosable language "typescript" %}}

After some dependency installations from `npm`, your project and stack will be ready.

#### Install Metabase Component

Next you will need to install the Metabase Component so you can use it in your program.

##### Yarn
```bash
$ yarn add @pulumi/metabase
```

##### NPM
```bash
$ npm install @pulumi/metabase
```
{{% /choosable %}}

{{% choosable language python %}}

After the command completes, the project and stack will be ready.

#### Install Metabase Component

Next you will need to install the Metabase Component so you can use it in your program.

```bash
$ pip3 install pulumi_metabase
```
{{% /choosable %}}

{{% choosable language go %}}

After the command completes, the project and stack will be ready.

#### Install Metabase Component

Next you will need to install the Metabase Component so you can use it in your program.

```bash
$ go get -u github.com/pulumi/pulumi-metabase/sdk
```

{{% /choosable %}}

{{% choosable language "csharp,fsharp,visualbasic" %}}

After the command completes, the project and stack will be ready.

#### Install Metabase Component

Next you will need to install the Metabase Component so you can use it in your program.

```bash
$ dotnet add package Pulumi.Metabase
```

{{% /choosable %}}

### Update Code

Now that you have all your dependencies installed and your project configured, you can now add the code that will
provision your Metabase Service.

{{< chooser language "typescript,python,go,csharp,yaml" / >}}
{{% choosable language typescript %}}

Replace your `index.ts` with the following:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as metabase from "@pulumi/metabase";

const metabaseService = new metabase.Metabase("metabaseService", {});
export const url = metabaseService.dnsName;
```

{{% /choosable %}}
{{% choosable language python %}}

Replace your `__main__.py` with the following:

```python
import pulumi
import pulumi_metabase as metabase

metabase_service = metabase.Metabase("metabaseService")
pulumi.export("url", metabase_service.dns_name)

```

{{% /choosable %}}
{{% choosable language go %}}

Replace your `main.go` with the following:

```go
package main

import (
	"github.com/pulumi/pulumi-metabase/sdk/go/metabase"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		metabaseService, err := metabase.NewMetabase(ctx, "metabaseService", nil)
		if err != nil {
			return err
		}
		ctx.Export("url", metabaseService.DnsName)
		return nil
	})
}
```

{{% /choosable %}}
{{% choosable language csharp %}}

Replace your `Program.cs` with the following:

```csharp
using Pulumi;
using Metabase = Pulumi.Metabase;

await Deployment.RunAsync(() =>
{
   var metabaseService = new Metabase.Metabase("metabaseService");

   return new Dictionary<string, object?>
   {
      ["url"] = metabaseService.DnsName
   };
});
```

{{% /choosable %}}
{{% choosable language yaml %}}

Replace your `Pulumi.yaml` with the following:

```yaml
name: metabase-yaml
runtime: yaml
resources:
    metabaseService:
        type: "metabase:index:Metabase"
outputs:
    url: ${metabaseService.dnsName}
```

{{% /choosable %}}

### Deploy

Once you have updated your code you are ready to deploy your Metabase Component. To do so, just run the the following command:

```bash
$ pulumi up
```

First Pulumi will perform a preview showing you exactly what will be created. Once the preview is complete Pulumi will ask you if you want to continue.
Select `yes` to proceed to actually provisioning the service.

All the resources will take a few minutes to fully provision. Once the update has completed it is likely it will take a few more minutes for the Metabase
task to finishing provisioning and start accepting traffic.

### (Optional) Destroy

You can destroy all the resources by running `pulumi destroy`. This will ultimately delete your Metabase Service's database so you will lose all stored
data and have to start from scratch if you provision a new Metabase service.

## Full Examples

Below you will find complete examples (all arguments supplied) for the Metabase Components.

{{< chooser language "typescript,python,go,csharp,yaml" / >}}
{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as metabase from "@pulumi/metabase";

const metabaseService = new metabase.Metabase("metabaseService", {
    vpcId: "vpc-123",
    networking: {
        ecsSubnetIds: [
            "subnet-123",
            "subnet-456",
        ],
        dbSubnetIds: [
            "subnet-789",
            "subnet-abc",
        ],
        lbSubnetIds: [
            "subnet-def",
            "subnet-ghi",
        ],
    },
    domain: {
        hostedZoneName: "example.com",
        domainName: "metabase.example.com",
    },
});
export const url = metabaseService.dnsName;
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
import pulumi_metabase as metabase

metabase_service = metabase.Metabase("metabaseService",
    vpc_id="vpc-123",
    networking=metabase.NetworkingArgs(
        ecs_subnet_ids=[
            "subnet-123",
            "subnet-456",
        ],
        db_subnet_ids=[
            "subnet-789",
            "subnet-abc",
        ],
        lb_subnet_ids=[
            "subnet-def",
            "subnet-ghi",
        ],
    ),
    domain=metabase.CustomDomainArgs(
        hosted_zone_name="example.com",
        domain_name="metabase.example.com",
    ))
pulumi.export("url", metabase_service.dns_name)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi-metabase/sdk/go/metabase"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		metabaseService, err := metabase.NewMetabase(ctx, "metabaseService", &metabase.MetabaseArgs{
			VpcId: pulumi.String("vpc-123"),
			Networking: &metabase.NetworkingArgs{
				EcsSubnetIds: pulumi.StringArray{
					pulumi.String("subnet-123"),
					pulumi.String("subnet-456"),
				},
				DbSubnetIds: pulumi.StringArray{
					pulumi.String("subnet-789"),
					pulumi.String("subnet-abc"),
				},
				LbSubnetIds: pulumi.StringArray{
					pulumi.String("subnet-def"),
					pulumi.String("subnet-ghi"),
				},
			},
			Domain: &metabase.CustomDomainArgs{
				HostedZoneName: pulumi.String("example.com"),
				DomainName:     pulumi.String("metabase.example.com"),
			},
		})
		if err != nil {
			return err
		}
		ctx.Export("url", metabaseService.DnsName)
		return nil
	})
}
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using Pulumi;
using Metabase = Pulumi.Metabase;

await Deployment.RunAsync(() =>
{
   var metabaseService = new Metabase.Metabase("metabaseService", new Metabase.MetabaseArgs
    {
        VpcId = "vpc-123",
        Networking = new Metabase.Inputs.NetworkingArgs
        {
            EcsSubnetIds =
            {
                "subnet-123",
                "subnet-456",
            },
            DbSubnetIds =
            {
                "subnet-789",
                "subnet-abc",
            },
            LbSubnetIds =
            {
                "subnet-def",
                "subnet-ghi",
            },
        },
        Domain = new Metabase.Inputs.CustomDomainArgs
        {
            HostedZoneName = "example.com",
            DomainName = "metabase.example.com",
        },
    });

   return new Dictionary<string, object?>
   {
      ["url"] = metabaseService.DnsName
   };
});
```

{{% /choosable %}}
{{% choosable language yaml %}}

```yaml
name: metabase-yaml
runtime: yaml
resources:
    metabaseService:
        type: "metabase:index:Metabase"
        properties:
            vpcId: "vpc-123"
            networking:
                ecsSubnetIds: [ "subnet-123", "subnet-456" ]
                dbSubnetIds: [ "subnet-789", "subnet-abc" ]
                lbSubnetIds: [ "subnet-def", "subnet-ghi" ]
            domain:
                hostedZoneName: "example.com"
                domainName: "metabase.example.com"
outputs:
    url: ${metabaseService.dnsName}
```

{{% /choosable %}}
