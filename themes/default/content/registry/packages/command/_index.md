---
title: Command
meta_desc: The Pulumi Command Provider enables you to execute commands and scripts either locally or remotely as part of the Pulumi resource model.
layout: package
---

The Pulumi Command Provider enables you to execute commands and scripts either locally or remotely as part of the Pulumi resource model.  The package also includes a resource for copying assets and archives to remote hosts via SSH.  Resources in the command package support running scripts on `create`, `update`, and `destroy` operations, supporting stateful command execution.

There are many scenarios where the Command package can be useful:

* Running a command locally after creating a resource, to register it with an external service
* Running a command locally before deleting a resource, to deregister it with an external service
* Running a command remotely on a remote host immediately after creating it
* Copying a file to a remote host after creating it (potentially as a script to be executed afterwards)
* As a simple alternative to some use cases for Dynamic Providers (especially in languages which do not yet support Dynamic Providers).

Some users may have experience with Terraform "provisioners", and the Command package offers support for similar scenarios.  However, the Command package is provided as independent resources which can be combined with other resources in many interesting ways. This has many strengths, but also some differences, such as the fact that a Command resource failing does not cause a resource it is operating on to fail.

You can use the Command package from a Pulumi program written in any Pulumi language: C#, Go, Java, JavaScript/TypeScript, Python, and YAML.
You'll need to [install and configure the Pulumi CLI](https://pulumi.com/docs/install/) if you haven't already.

## Examples

### A simple local resource (random)

The simplest use case for `local.Command` is to just run a command on `create`, which can return some value which will be stored in the state file, and will be persistent for the life of the stack (or until the resource is destroyed or replaced).  The example below uses this as an alternative to the `random` package to create some randomness which is stored in Pulumi state.

{{< chooser language "javascript,typescript,python,go,csharp,yaml,java" >}}

{{% choosable language javascript %}}

```javascript
"use strict";
const command = require("@pulumi/command");

const random = new command.local.Command("random", {
    create: "openssl rand -hex 16",
});

exports.output = random.stdout;
```

{{% /choosable %}}

{{% choosable language typescript %}}

```typescript
import { local } from "@pulumi/command";

const random = new local.Command("random", {
    create: "openssl rand -hex 16",
});

export const output = random.stdout;
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using Pulumi;
using Pulumi.Command.Local;

await Deployment.RunAsync(() =>
{
    var command = new Command("random", new CommandArgs
    {
        Create = "openssl rand -hex 16"
    });

    return new Dictionary<string, object?>
    {
        ["stdOut"] = command.Stdout
    };
});
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
from pulumi_command import local

random = local.Command("random",
    create="openssl rand -hex 16"
)

pulumi.export("random", random.stdout)
```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi-command/sdk/go/command/local"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		random, err := local.NewCommand(ctx, "random", &local.CommandArgs{
			Create: pulumi.String("openssl rand -hex 16"),
		})
		if err != nil {
			return err
		}

		ctx.Export("output", random.Stdout)
		return nil
	})
}
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: myproject
runtime: yaml
resources:
  random:
    type: command:local:Command
    properties:
      create: "openssl rand -hex 16"
outputs:
  rand: "${random.stdout}"
```

{{% /choosable %}}

{{% choosable language java %}}

```java
package myproject;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.command.local.Command;
import com.pulumi.command.local.CommandArgs;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        var random = new Command("random", CommandArgs.builder()
            .create("openssl rand -hex 16")
            .build());

        ctx.export("rand", random.stdout());
    }
}
```

{{% /choosable %}}

{{< /chooser >}}

### Remote provisioning of an EC2 instance

This example creates an EC2 instance, and then uses `remote.Command` and `remote.CopyToRemote` to run commands and copy files to the remote instance (via SSH). Similar things are possible with Azure, Google Cloud and other cloud provider virtual machines.  Support for Windows-based VMs is being tracked [here](https://github.com/pulumi/pulumi-command/issues/15).

Implicit and explicit (`dependsOn`) dependencies can be used to control the order that these `Command` and `CopyToRemote` resources are constructed relative to each other and to the cloud resources they depend on.  This ensures that the `create` operations run after all dependencies are created, and the `delete` operations run before all dependencies are deleted.

Because the `Command` and `CopyToRemote` resources replace on changes to their connection, if the EC2 instance is replaced, the commands will all re-run on the new instance (and the `delete` operations will run on the old instance).

Note also that `deleteBeforeReplace` can be composed with `Command` resources to ensure that the `delete` operation on an "old" instance is run before the `create` operation of the new instance, in case a scarce resource is managed by the command.  Similarly, other resource options can naturally be applied to `Command` resources, like `ignoreChanges`.

```typescript
import { interpolate, Config } from "@pulumi/pulumi";
import { local, remote, types } from "@pulumi/command";
import * as aws from "@pulumi/aws";
import * as fs from "fs";
import * as os from "os";
import * as path from "path";
import { size } from "./size";

const config = new Config();
const keyName = config.get("keyName") ?? new aws.ec2.KeyPair("key", { publicKey: config.require("publicKey") }).keyName;
const privateKeyBase64 = config.get("privateKeyBase64");
const privateKey = privateKeyBase64 ? Buffer.from(privateKeyBase64, 'base64').toString('ascii') : fs.readFileSync(path.join(os.homedir(), ".ssh", "id_rsa")).toString("utf8");

const secgrp = new aws.ec2.SecurityGroup("secgrp", {
    description: "Foo",
    ingress: [
        { protocol: "tcp", fromPort: 22, toPort: 22, cidrBlocks: ["0.0.0.0/0"] },
        { protocol: "tcp", fromPort: 80, toPort: 80, cidrBlocks: ["0.0.0.0/0"] },
    ],
});

const ami = aws.ec2.getAmiOutput({
    owners: ["amazon"],
    mostRecent: true,
    filters: [{
        name: "name",
        values: ["amzn2-ami-hvm-2.0.????????-x86_64-gp2"],
    }],
});

const server = new aws.ec2.Instance("server", {
    instanceType: size,
    ami: ami.id,
    keyName: keyName,
    vpcSecurityGroupIds: [secgrp.id],
}, { replaceOnChanges: ["instanceType"] });

// Now set up a connection to the instance and run some provisioning operations on the instance.

const connection: types.input.remote.ConnectionArgs = {
    host: server.publicIp,
    user: "ec2-user",
    privateKey: privateKey,
};

const hostname = new remote.Command("hostname", {
    connection,
    create: "hostname",
});

new remote.Command("remotePrivateIP", {
    connection,
    create: interpolate`echo ${server.privateIp} > private_ip.txt`,
    delete: `rm private_ip.txt`,
}, { deleteBeforeReplace: true });

new local.Command("localPrivateIP", {
    create: interpolate`echo ${server.privateIp} > private_ip.txt`,
    delete: `rm private_ip.txt`,
}, { deleteBeforeReplace: true });

const sizeFile = new remote.CopyToRemote("size", {
    connection,
    source: pulumi.asset.FileAsset("./size.ts"),
    remotePath: "size.ts",
})

const catSize = new remote.Command("checkSize", {
    connection,
    create: "cat size.ts",
}, { dependsOn: sizeFile })

export const confirmSize = catSize.stdout;
export const publicIp = server.publicIp;
export const publicHostName = server.publicDns;
export const hostnameStdout = hostname.stdout;
```

### Invoking a Lambda during Pulumi deployment

There may be cases where it is useful to run some code within an AWS Lambda or other serverless function during the deployment.  For example, this may allow running some code from within a VPC, or with a specific role, without needing to have persistent compute available (such as the EC2 example above).

Note that the Lambda function itself can be created within the same Pulumi program, and then invoked after creation.

The example below simply creates some random value within the Lambda, which is a very roundabout way of doing the same thing as the first "random" example above, but this pattern can be used for more complex scenarios where the Lambda does things a local script could not.

{{< chooser language "javascript,typescript,python,go,csharp,java,yaml" >}}

{{% choosable language "javascript" %}}

```javascript
"use strict";
const aws = require("@pulumi/aws");
const { local } = require("@pulumi/command");
const { getStack } = require("@pulumi/pulumi");

const f = new aws.lambda.CallbackFunction("f", {
    publish: true,
    callback: async (ev) => {
        return `Stack ${ev.stackName} is deployed!`;
    }
});

const invoke = new local.Command("execf", {
    create: `aws lambda invoke --function-name "$FN" --payload '{"stackName": "${getStack()}"}' --cli-binary-format raw-in-base64-out out.txt >/dev/null && cat out.txt | tr -d '"'  && rm out.txt`,
    environment: {
        FN: f.qualifiedArn,
        AWS_REGION: aws.config.region,
        AWS_PAGER: "",
    },
}, { dependsOn: f })

exports.output = invoke.stdout;
```

{{% /choosable %}}

{{% choosable language "typescript" %}}

```typescript
import * as aws from "@pulumi/aws";
import { local } from "@pulumi/command";
import { getStack } from "@pulumi/pulumi";

const f = new aws.lambda.CallbackFunction("f", {
    publish: true,
    callback: async (ev: any) => {
        return `Stack ${ev.stackName} is deployed!`;
    }
});

const invoke = new local.Command("execf", {
    create: `aws lambda invoke --function-name "$FN" --payload '{"stackName": "${getStack()}"}' --cli-binary-format raw-in-base64-out out.txt >/dev/null && cat out.txt | tr -d '"'  && rm out.txt`,
    environment: {
        FN: f.qualifiedArn,
        AWS_REGION: aws.config.region!,
        AWS_PAGER: "",
    },
}, { dependsOn: f })

export const output = invoke.stdout;
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import json
import pulumi_aws as aws
import pulumi_command as command

lambda_role = aws.iam.Role("lambdaRole", assume_role_policy=json.dumps({
    "Version": "2012-10-17",
    "Statement": [{
        "Action": "sts:AssumeRole",
        "Effect": "Allow",
        "Principal": {
            "Service": "lambda.amazonaws.com",
        },
    }],
}))

lambda_function = aws.lambda_.Function("lambdaFunction",
    name="f",
    publish=True,
    role=lambda_role.arn,
    handler="index.handler",
    runtime=aws.lambda_.Runtime.NODE_JS20D_X,
    code=pulumi.FileArchive("./handler"))

aws_config = pulumi.Config("aws")
aws_region = aws_config.require("region")

invoke_command = command.local.Command("invokeCommand",
    create=f"aws lambda invoke --function-name \"$FN\" --payload '{{\"stackName\": \"{pulumi.get_stack()}\"}}' --cli-binary-format raw-in-base64-out out.txt >/dev/null && cat out.txt | tr -d '\"'  && rm out.txt",
    environment={
        "FN": lambda_function.arn,
        "AWS_REGION": aws_region,
        "AWS_PAGER": "",
    },
    opts = pulumi.ResourceOptions(depends_on=[lambda_function]))

pulumi.export("output", invoke_command.stdout)
```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi-command/sdk/go/command/local"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		awsConfig := config.New(ctx, "aws")
		awsRegion := awsConfig.Require("region")

		tmpJSON, err := json.Marshal(map[string]interface{}{
			"Version": "2012-10-17",
			"Statement": []map[string]interface{}{
				{
					"Action": "sts:AssumeRole",
					"Effect": "Allow",
					"Principal": map[string]interface{}{
						"Service": "lambda.amazonaws.com",
					},
				},
			},
		})
		if err != nil {
			return err
		}
		lambdaRole, err := iam.NewRole(ctx, "lambdaRole", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(string(tmpJSON)),
		})
		if err != nil {
			return err
		}

        lambdaFunction, err := lambda.NewFunction(ctx, "lambdaFunction", &lambda.FunctionArgs{
			Name:    pulumi.String("f"),
			Publish: pulumi.Bool(true),
			Role:    lambdaRole.Arn,
			Handler: pulumi.String("index.handler"),
			Runtime: pulumi.String(lambda.RuntimeNodeJS20dX),
			Code:    pulumi.NewFileArchive("./handler"),
		})
		if err != nil {
			return err
		}

        invokeCommand, err := local.NewCommand(ctx, "invokeCommand", &local.CommandArgs{
			Create: pulumi.String(fmt.Sprintf("aws lambda invoke --function-name \"$FN\" --payload '{\"stackName\": \"%v\"}' --cli-binary-format raw-in-base64-out out.txt >/dev/null && cat out.txt | tr -d '\"'  && rm out.txt", ctx.Stack())),
			Environment: pulumi.StringMap{
				"FN":         lambdaFunction.Arn,
				"AWS_REGION": pulumi.String(awsRegion),
				"AWS_PAGER":  pulumi.String(""),
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			lambdaFunction,
		}))
		if err != nil {
			return err
		}
		ctx.Export("output", invokeCommand.Stdout)
		return nil
	})
}
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;
using Command = Pulumi.Command;

return await Deployment.RunAsync(() =>
{
    var awsConfig = new Config("aws");

    var lambdaRole = new Aws.Iam.Role("lambdaRole", new()
    {
        AssumeRolePolicy = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Version"] = "2012-10-17",
            ["Statement"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["Action"] = "sts:AssumeRole",
                    ["Effect"] = "Allow",
                    ["Principal"] = new Dictionary<string, object?>
                    {
                        ["Service"] = "lambda.amazonaws.com",
                    },
                },
            },
        }),
    });

    var lambdaFunction = new Aws.Lambda.Function("lambdaFunction", new()
    {
        Name = "f",
        Publish = true,
        Role = lambdaRole.Arn,
        Handler = "index.handler",
        Runtime = Aws.Lambda.Runtime.NodeJS20dX,
        Code = new FileArchive("./handler"),
    });

    var invokeCommand = new Command.Local.Command("invokeCommand", new()
    {
        Create = $"aws lambda invoke --function-name \"$FN\" --payload '{{\"stackName\": \"{Deployment.Instance.StackName}\"}}' --cli-binary-format raw-in-base64-out out.txt >/dev/null && cat out.txt | tr -d '\"'  && rm out.txt",
        Environment =
        {
            { "FN", lambdaFunction.Arn },
            { "AWS_REGION", awsConfig.Require("region") },
            { "AWS_PAGER", "" },
        },
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            lambdaFunction,
        },
    });

    return new Dictionary<string, object?>
    {
        ["output"] = invokeCommand.Stdout,
    };
});
```

{{% /choosable %}}

{{% choosable language java %}}

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.lambda.Function;
import com.pulumi.aws.lambda.FunctionArgs;
import com.pulumi.command.local.Command;
import com.pulumi.command.local.CommandArgs;
import static com.pulumi.codegen.internal.Serialization.*;
import com.pulumi.resources.CustomResourceOptions;
import com.pulumi.asset.FileArchive;
import java.util.Map;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        var awsConfig = ctx.config("aws");
        var awsRegion = awsConfig.require("region");

        var lambdaRole = new Role("lambdaRole", RoleArgs.builder()
                .assumeRolePolicy(serializeJson(
                        jsonObject(
                                jsonProperty("Version", "2012-10-17"),
                                jsonProperty("Statement", jsonArray(jsonObject(
                                        jsonProperty("Action", "sts:AssumeRole"),
                                        jsonProperty("Effect", "Allow"),
                                        jsonProperty("Principal", jsonObject(
                                                jsonProperty("Service", "lambda.amazonaws.com")))))))))
                .build());

        var lambdaFunction = new Function("lambdaFunction", FunctionArgs.builder()
                .name("f")
                .publish(true)
                .role(lambdaRole.arn())
                .handler("index.handler")
                .runtime("nodejs20.x")
                .code(new FileArchive("./handler"))
                .build());

        // Work around the lack of Output.all for Maps in Java. We cannot use a plain Map because
        // `lambdaFunction.arn()` is an Output<String>.
        var invokeEnv = Output.tuple(
                Output.of("FN"), lambdaFunction.arn(),
                Output.of("AWS_REGION"), Output.of(awsRegion),
                Output.of("AWS_PAGER"), Output.of("")
        ).applyValue(t -> Map.of(t.t1, t.t2, t.t3, t.t4, t.t5, t.t6));

        var invokeCommand = new Command("invokeCommand", CommandArgs.builder()
                .create(String.format(
                        "aws lambda invoke --function-name \"$FN\" --payload '{\"stackName\": \"%s\"}' --cli-binary-format raw-in-base64-out out.txt >/dev/null && cat out.txt | tr -d '\"'  && rm out.txt",
                        ctx.stackName()))
                .environment(invokeEnv)
                .build(),
                CustomResourceOptions.builder()
                        .dependsOn(lambdaFunction)
                        .build());

        ctx.export("output", invokeCommand.stdout());
    }
}
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: myproject
runtime: yaml
resources:
  lambdaRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy:
        fn::toJSON:
          Version: "2012-10-17"
          Statement:
            - Action: sts:AssumeRole
              Effect: Allow
              Principal:
                Service: lambda.amazonaws.com

  lambdaFunction:
    type: aws:lambda:Function
    properties:
      name: f
      publish: true
      role: ${lambdaRole.arn}
      handler: index.handler
      runtime: "nodejs20.x"
      code:
        fn::fileArchive: ./handler

  invokeCommand:
    type: command:local:Command
    properties:
      create: 'aws lambda invoke --function-name "$FN" --payload ''{"stackName": "${pulumi.stack}"}'' --cli-binary-format raw-in-base64-out out.txt >/dev/null && cat out.txt | tr -d ''"''  && rm out.txt'
      environment:
        FN: ${lambdaFunction.arn}
        AWS_REGION: ${aws:region}
        AWS_PAGER: ""
    options:
      dependsOn:
        - ${lambdaFunction}

outputs:
  output: ${invokeCommand.stdout}
```

{{% /choosable %}}

{{< /chooser >}}

### Using `local.Command` with CURL to manage external REST API

This example uses `local.Command` to create a simple resource provider for managing GitHub labels, by invoking `curl` commands on `create` and `delete` commands against the GitHub REST API.  A similar approach could be applied to build other simple providers against any REST API directly from within Pulumi programs in any language.  This approach is somewhat limited by the fact that `local.Command` does not yet support `diff`/`update`/`read`.  Support for those may be [added in the future](https://github.com/pulumi/pulumi-command/issues/20).

This example also shows how `local.Command` can be used as an implementation detail inside a nicer abstraction, like the `GitHubLabel` component defined below.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as random from "@pulumi/random";
import { local } from "@pulumi/command";

interface LabelArgs {
    owner: pulumi.Input<string>;
    repo: pulumi.Input<string>;
    name: pulumi.Input<string>;
    githubToken: pulumi.Input<string>;
}

class GitHubLabel extends pulumi.ComponentResource {
    public url: pulumi.Output<string>;

    constructor(name: string, args: LabelArgs, opts?: pulumi.ComponentResourceOptions) {
        super("example:github:Label", name, args, opts);

        const label = new local.Command("label", {
            create: "./create_label.sh",
            delete: "./delete_label.sh",
            environment: {
                OWNER: args.owner,
                REPO: args.repo,
                NAME: args.name,
                GITHUB_TOKEN: args.githubToken,
            }
        }, { parent: this });

        const response = label.stdout.apply(JSON.parse);
        this.url = response.apply((x: any) => x.url as string);
    }
}

const config = new pulumi.Config();
const rand = new random.RandomString("s", { length: 10, special: false });

const label = new GitHubLabel("l", {
    owner: "pulumi",
    repo: "pulumi-command",
    name: rand.result,
    githubToken: config.requireSecret("githubToken"),
});

export const labelUrl = label.url;
```

```sh
# create_label.sh
curl \
  -s \
  -X POST \
  -H "authorization: Bearer $GITHUB_TOKEN" \
  -H "Accept: application/vnd.github.v3+json" \
  https://api.github.com/repos/$OWNER/$REPO/labels \
  -d "{\"name\":\"$NAME\"}"
```

```sh
# delete_label.sh
curl \
  -s \
  -X DELETE \
  -H "authorization: Bearer $GITHUB_TOKEN" \
  -H "Accept: application/vnd.github.v3+json" \
  https://api.github.com/repos/$OWNER/$REPO/labels/$NAME
```

### Graceful cleanup of workloads in a Kubernetes cluster

There are cases where it's important to run some cleanup operation before destroying a resource, in case destroying the resource does not properly handle orderly cleanup.  For example, destroying an EKS Cluster will not ensure that all Kubernetes object finalizers are run, which may lead to leaking external resources managed by those Kubernetes resources.  This example shows how we can use a `delete`-only `Command` to ensure some cleanup is run within a cluster before destroying it.

{{< chooser language "javascript,typescript,python,go,csharp,java,yaml" >}}

{{% choosable language "javascript" %}}

```javascript
"use strict";
const command = require("@pulumi/command");
const eks = require("@pulumi/eks");

const cluster = new eks.Cluster("cluster", {});
const cleanupKubernetesNamespaces = new command.local.Command("cleanupKubernetesNamespaces", {
    "delete": "kubectl --kubeconfig <(echo \"$KUBECONFIG_DATA\") delete namespace nginx\n",
    interpreter: [
        "/bin/bash",
        "-c",
    ],
    environment: {
        KUBECONFIG_DATA: cluster.kubeconfigJson,
    },
});
```

{{% /choosable %}}

{{% choosable language "typescript" %}}

```typescript
import * as command from "@pulumi/command";
import * as eks from "@pulumi/eks";

const cluster = new eks.Cluster("cluster", {});
const cleanupKubernetesNamespaces = new command.local.Command("cleanupKubernetesNamespaces", {
    "delete": "kubectl --kubeconfig <(echo \"$KUBECONFIG_DATA\") delete namespace nginx\n",
    interpreter: [
        "/bin/bash",
        "-c",
    ],
    environment: {
        KUBECONFIG_DATA: cluster.kubeconfigJson,
    },
});
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_command as command
import pulumi_eks as eks

cluster = eks.Cluster("cluster")
cleanup_kubernetes_namespaces = command.local.Command("cleanupKubernetesNamespaces",
    delete="kubectl --kubeconfig <(echo \"$KUBECONFIG_DATA\") delete namespace nginx\n",
    interpreter=[
        "/bin/bash",
        "-c",
    ],
    environment={
        "KUBECONFIG_DATA": cluster.kubeconfig_json,
    })
```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi-command/sdk/go/command/local"
	"github.com/pulumi/pulumi-eks/sdk/v2/go/eks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cluster, err := eks.NewCluster(ctx, "cluster", nil)
		if err != nil {
			return err
		}
		_, err = local.NewCommand(ctx, "cleanupKubernetesNamespaces", &local.CommandArgs{
			Delete: pulumi.String("kubectl --kubeconfig <(echo \"$KUBECONFIG_DATA\") delete namespace nginx\n"),
			Interpreter: pulumi.StringArray{
				pulumi.String("/bin/bash"),
				pulumi.String("-c"),
			},
			Environment: pulumi.StringMap{
				"KUBECONFIG_DATA": cluster.KubeconfigJson,
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using Pulumi;
using Command = Pulumi.Command;
using Eks = Pulumi.Eks;

return await Deployment.RunAsync(() =>
{
    var cluster = new Eks.Cluster("cluster");

    var cleanupKubernetesNamespaces = new Command.Local.Command("cleanupKubernetesNamespaces", new()
    {
        Delete = @"kubectl --kubeconfig <(echo ""$KUBECONFIG_DATA"") delete namespace nginx",
        Interpreter = new[]
        {
            "/bin/bash", "-c",
        },
        Environment =
        {
            { "KUBECONFIG_DATA", cluster.KubeconfigJson },
        },
    });
});
```

{{% /choosable %}}

{{% choosable language java %}}

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.eks.Cluster;
import com.pulumi.command.local.Command;
import com.pulumi.command.local.CommandArgs;
import java.util.Map;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        var cluster = new Cluster("cluster");

        // Work around the lack of Output.all for Maps in Java.
        var envMap = Output.tuple(Output.of("KUBECONFIG"), cluster.kubeconfigJson())
            .applyValue(t -> Map.of(t.t1, t.t2));

        var cleanupKubernetesNamespaces = new Command("cleanupKubernetesNamespaces", CommandArgs.builder()
            .delete("""
kubectl --kubeconfig <(echo "$KUBECONFIG_DATA") delete namespace nginx
            """)
            .interpreter("/bin/bash", "-c")
            .environment(envMap)
            .build());
    }
}
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: myproject
runtime: yaml
resources:
  cluster:
    type: eks:Cluster

  cleanupKubernetesNamespaces:
    # We could also use `RemoteCommand` to run this from
    # within a node in the cluster.
    type: command:local:Command
    properties:
      # This will run before the cluster is destroyed.
      # Everything else will need to depend on this resource
      # to ensure this cleanup doesn't happen too early.
      delete: |
        kubectl --kubeconfig <(echo "$KUBECONFIG_DATA") delete namespace nginx
      # Process substitution "<()" doesn't work in the default interpreter sh.
      interpreter: ["/bin/bash", "-c"]
      environment:
        KUBECONFIG_DATA: "${cluster.kubeconfigJson}"
```

{{% /choosable %}}

{{< /chooser >}}
