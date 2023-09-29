---
title: OVH
meta_desc: Provides an overview of the OVH Provider for Pulumi.
layout: overview
---

The `OVH` provider for Pulumi can be used to provision any of the resources available in [OVHcloud](https://www.ovhcloud.com/fr/).
The `OVH` provider must be configured with credentials to deploy and update resources in `OVH`.

## Information

Note that the [lbrlabs Pulumi OVH provider](https://github.com/lbrlabs/pulumi-ovh) is replaced by this official one. See the
[Migration Guide](./how-to-guides/migration-from-lbrlabs-package.md) for more information.

## Example

{{< chooser language "go,typescript,python,csharp" >}}
{{% choosable language go %}}

```golang
// Deploy a new Kubernetes cluster
myKube, err := cloudproject.NewKube(ctx, "my_desired_cluster", &cloudproject.KubeArgs{
    ServiceName: pulumi.String("xxxxxxxxxxxxx-xxxx-xxxx-xxxxxxxxx"),
    Name:        pulumi.String("my_desired_cluster"),
    Region:      pulumi.String("GRA5"),
})
if err != nil {
    return err
}

// Export kubeconfig file to a secret
ctx.Export("kubeconfig", pulumi.ToSecret(myKube.Kubeconfig))
```

{{% /choosable %}}

{{% choosable language python %}}

```python
"""Get a Kubernetes cluster version"""
import pulumi
import pulumi_ovh as ovh

config = pulumi.Config();
service_name = config.require('serviceName')

print(service_name);

my_kube_cluster = ovh.cloudproject.get_kube(service_name=service_name,
    kube_id="xxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx");
pulumi.export("version", my_kube_cluster.version)
```

{{% /choosable %}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";

import * as ovh from "@ovh-devrelteam/pulumi-ovh"

let config = new pulumi.Config();
let serviceName = config.require("serviceName")

console.log(serviceName)

// Get a Kubernetes cluster version
let myKubeCluster = ovh.cloudproject.getKube({
    serviceName: serviceName,
    kubeId: "xxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxx"
})

export const version = myKubeCluster.then(myKubeCluster => myKubeCluster.version);
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Ovh = Pulumi.Ovh;
using System;

return await Deployment.RunAsync(() =>
{

    var config = new Pulumi.Config();
    var serviceName = config.Require("serviceName");

    System.Console.WriteLine(serviceName);

    var myKubeCluster = Ovh.CloudProject.GetKube.Invoke(new()
    {
        ServiceName = serviceName,
        KubeId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxx",
    });

    return new Dictionary<string, object?>
    {
        ["version"] = myKubeCluster.Apply(getKubeResult => getKubeResult.Version),
    };
});
```

{{% /choosable %}}

{{< /chooser >}}

## Migrating from the LbrLabs OVH Package

The types in each of the Pulumi OVH SDKs have been kept as compatible as possible to the LbrLabs community package.
This should make the migration from the community package to this official OVH package almost a drop in replacement.

{{< chooser language "typescript,python,go,csharp" >}}

{{% choosable language typescript %}}

Replace the NPM package in your project first:

```sh
$ npm uninstall --save @lbrlabs/pulumi-ovh
$ npm install --save @ovh-devrelteam/pulumi-ovh
```

After switching the package, update the import statements in your Pulumi infrastructure code:

```diff
- import * as ovh from "@lbrlabs/pulumi-ovh";
+ import * as ovh from "@ovh-devrelteam/pulumi-ovh";
```

{{% /choosable %}}

{{% choosable language python %}}

Replace the Python package in your project first:

```sh
$ pip uninstall lbrlabs-pulumi-ovh
$ pip install pulumi-ovh
```

After switching the package, update the import statements in your Pulumi infrastructure code:

```diff
-  import lbrlabs_pulumi_ovh as ovh
+  import pulumi_ovh as ovh
```

{{% /choosable %}}

{{% choosable language go %}}

Replace the Go package in your project first:

```sh
$ go get github.com/ovh/pulumi-ovh/sdk/go/...
```

After switching the package, update the import statements in your Pulumi infrastructure code:

```diff
- "github.com/lbrlabs/pulumi-ovh/sdk/go/ovh/CloudProject"
+ "github.com/ovh/pulumi-ovh/sdk/go/ovh/CloudProject"
```

{{% /choosable %}}

{{% choosable language csharp %}}

Replace the Nuget package in your project first:

```sh
$ dotnet remove package Lbrlabs.PulumiPackage.Ovh
$ dotnet add package Pulumi.Ovh
```

After switching the package, update the import statements in your Pulumi infrastructure code:

```diff
- using Ovh = Lbrlabs.PulumiPackage.Ovh;
+ using Ovh = Pulumi.Ovh;
```

{{% /choosable %}}

{{< /chooser >}}

Your project has now been updated to use the new SDK. When running `pulumi preview` with this new setup,
you should see no changes to your OVH resources reported.

If you do encounter something unexpected, please file an issue in [the Github repository](https://github.com/ovh/pulumi-ovh/issues).

## Provider configuration

The LbrLabs OVH package marked the consumer key as a Pulumi secret. This was incorrect. This OVH package has it marked
as a regular configuration item, but marks the application secret correctly as a Pulumi secret.

If you have an existing setup where the provider configuration was passed using stack configuration, you have to
change the stack configuration. Make sure you have valid credentials at hand. Then run the following commands:

```sh
pulumi config rm ovh:consumerKey
pulumi config set ovh:consumerKey <your-consumer-key>
```

The consumer key is removed from the stack config as the old encrypted value, and re-added as a non-encrypted value.

```sh
pulumi config rm ovh:applicationSecret
pulumi config set ovh:applicationSecret --secret
value: <paste-your-application-secret>
```

The application secret was unencrypted before. This means that the value could have been save in Pulumi state unencrypted.
By removing it from config and re-adding it as a Pulumi secret config value, Pulumi can now track the value wherever it
is used and it will never be saved in clear text.

Because the application secret was processed unencrypted before by the LbrLabs OVH package,
it is advised to rotate your application credentials and add the new set as an encrypted value.
In this case, it is possible that Pulumi reports a change to the state, without any clear change to a resource.
The clear text value is replaced by the encrypted value. Please accept this Pulumi update to increase
the security posture of your setup.
