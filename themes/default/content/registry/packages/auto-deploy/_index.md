---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-auto-deploy/v0.0.4/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Pulumi Auto Deploy (preview)
meta_desc: The Pulumi Auto Deploy Provider enables you to automatically deploy dependent stacks using Pulumi Deployments.
layout: package
---

A Pulumi Component for configuring automated updates of dependent stacks using [Pulumi Deployments](https://www.pulumi.com/docs/pulumi-cloud/deployments/). It lets you simply express dependencies between stacks, and takes care of creating and updating the necessary Deployment Webhooks under the hood. Each stack that you configure must have [Deployment Settings](https://www.pulumi.com/docs/pulumi-cloud/deployments/reference/#deployment-settings).


You can use the Pulumi Auto Deploy package from a Pulumi program written in any Pulumi language: C#, Go, JavaScript/TypeScript, YAML, and Python.
You'll need to [install and configure the Pulumi CLI](https://pulumi.com/docs/install/) if you haven't already.

> **NOTE**: The Auto Deploy package is in preview.  The API design may change ahead of general availability based on [user feedback](https://github.com/pulumi/pulumi-auto-deploy/issues).

## Example

{{< chooser language "typescript,python,go,csharp,yaml" >}}

{{% choosable language typescript %}}

```ts
import * as autodeploy from "@pulumi/auto-deploy";
import * as pulumi from "@pulumi/pulumi";

/**
 *
 * The following example configures automatic deployment of stacks with the following dependency graph:
    a
    ├── b
    │   ├── d
    │   ├── e
    │   └── f
    └── c
 * Whenever a node in the graph is updated, 
 * all downstream nodes will be automatically updated via a webhook triggering Pulumi Deployments.
 */

const organization = pulumi.getOrganization();
const project = "dependency-example"

export const f = new autodeploy.AutoDeployer("auto-deployer-f", {
    organization,
    project,
    stack: "f",
    downstreamRefs: [],
});

export const e = new autodeploy.AutoDeployer("auto-deployer-e", {
    organization,
    project,
    stack: "e",
    downstreamRefs: [],
});

export const d = new autodeploy.AutoDeployer("auto-deployer-d", {
    organization,
    project,
    stack: "d",
    downstreamRefs: [],
});

export const c = new autodeploy.AutoDeployer("auto-deployer-c", {
    organization,
    project,
    stack: "c",
    downstreamRefs: [],
});

export const b = new autodeploy.AutoDeployer("auto-deployer-b", {
    organization,
    project,
    stack: "b",
    downstreamRefs: [d.ref, e.ref, f.ref],

});

export const a = new autodeploy.AutoDeployer("auto-deployer-a", {
    organization,
    project,
    stack: "a",
    downstreamRefs: [b.ref, c.ref],
});
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using AutoDeploy = Pulumi.AutoDeploy;

/**
 *
 * The following example configures automatic deployment of stacks with the following dependency graph:
    a
    ├── b
    │   ├── d
    │   ├── e
    │   └── f
    └── c
 * Whenever a node in the graph is updated, 
 * all downstream nodes will be automatically updated via a webhook triggering Pulumi Deployments.
 */

return await Deployment.RunAsync(() => 
{
    var projectVar = "dependency-example";

    var organization = "pulumi";

    var f = new AutoDeploy.AutoDeployer("f", new()
    {
        Organization = organization,
        Project = projectVar,
        Stack = "f",
        DownstreamRefs = new[] {},
    });

    var e = new AutoDeploy.AutoDeployer("e", new()
    {
        Organization = organization,
        Project = projectVar,
        Stack = "e",
        DownstreamRefs = new[] {},
    });

    var d = new AutoDeploy.AutoDeployer("d", new()
    {
        Organization = organization,
        Project = projectVar,
        Stack = "d",
        DownstreamRefs = new[] {},
    });

    var c = new AutoDeploy.AutoDeployer("c", new()
    {
        Organization = organization,
        Project = projectVar,
        Stack = "c",
        DownstreamRefs = new[] {},
    });

    var b = new AutoDeploy.AutoDeployer("b", new()
    {
        Organization = organization,
        Project = projectVar,
        Stack = "b",
        DownstreamRefs = new[]
        {
            d.Ref,
            e.Ref,
            f.Ref,
        },
    });

    var a = new AutoDeploy.AutoDeployer("a", new()
    {
        Organization = organization,
        Project = projectVar,
        Stack = "a",
        DownstreamRefs = new[]
        {
            b.Ref,
            c.Ref,
        },
    });

});
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_auto_deploy as auto_deploy

'''

The following example configures automatic deployment of stacks with the following dependency graph:
    a
    ├── b
    │   ├── d
    │   ├── e
    │   └── f
    └── c
Whenever a node in the graph is updated, 
all downstream nodes will be automatically updated via a webhook triggering Pulumi Deployments.
'''

project_var = "dependency-example"
organization = pulumi.get_organization()
f = auto_deploy.AutoDeployer("f",
    organization=organization,
    project=project_var,
    stack="f",
    downstream_refs=[])
e = auto_deploy.AutoDeployer("e",
    organization=organization,
    project=project_var,
    stack="e",
    downstream_refs=[])
d = auto_deploy.AutoDeployer("d",
    organization=organization,
    project=project_var,
    stack="d",
    downstream_refs=[])
c = auto_deploy.AutoDeployer("c",
    organization=organization,
    project=project_var,
    stack="c",
    downstream_refs=[])
b = auto_deploy.AutoDeployer("b",
    organization=organization,
    project=project_var,
    stack="b",
    downstream_refs=[
        d.ref,
        e.ref,
        f.ref,
    ])
a = auto_deploy.AutoDeployer("a",
    organization=organization,
    project=project_var,
    stack="a",
    downstream_refs=[
        b.ref,
        c.ref,
    ])
```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi-auto-deploy/sdk/go/autodeploy"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	/**
	*
	* The following example configures automatic deployment of stacks with the following dependency graph:
	   a
	   ├── b
	   │   ├── d
	   │   ├── e
	   │   └── f
	   └── c
	* Whenever a node in the graph is updated,
	* all downstream nodes will be automatically updated via a webhook triggering Pulumi Deployments.
	*/

	pulumi.Run(func(ctx *pulumi.Context) error {
		projectVar := "dependency-example"
		organization := "pulumi"
		f, err := autodeploy.NewAutoDeployer(ctx, "f", &autodeploy.AutoDeployerArgs{
			Organization:   pulumi.String(organization),
			Project:        pulumi.String(projectVar),
			Stack:          pulumi.String("f"),
			DownstreamRefs: pulumi.StringArray{},
		})
		if err != nil {
			return err
		}
		e, err := autodeploy.NewAutoDeployer(ctx, "e", &autodeploy.AutoDeployerArgs{
			Organization:   pulumi.String(organization),
			Project:        pulumi.String(projectVar),
			Stack:          pulumi.String("e"),
			DownstreamRefs: pulumi.StringArray{},
		})
		if err != nil {
			return err
		}
		d, err := autodeploy.NewAutoDeployer(ctx, "d", &autodeploy.AutoDeployerArgs{
			Organization:   pulumi.String(organization),
			Project:        pulumi.String(projectVar),
			Stack:          pulumi.String("d"),
			DownstreamRefs: pulumi.StringArray{},
		})
		if err != nil {
			return err
		}
		c, err := autodeploy.NewAutoDeployer(ctx, "c", &autodeploy.AutoDeployerArgs{
			Organization:   pulumi.String(organization),
			Project:        pulumi.String(projectVar),
			Stack:          pulumi.String("c"),
			DownstreamRefs: pulumi.StringArray{},
		})
		if err != nil {
			return err
		}
		b, err := autodeploy.NewAutoDeployer(ctx, "b", &autodeploy.AutoDeployerArgs{
			Organization: pulumi.String(organization),
			Project:      pulumi.String(projectVar),
			Stack:        pulumi.String("b"),
			DownstreamRefs: pulumi.StringArray{
				d.Ref,
				e.Ref,
				f.Ref,
			},
		})
		if err != nil {
			return err
		}
		_, err = autodeploy.NewAutoDeployer(ctx, "a", &autodeploy.AutoDeployerArgs{
			Organization: pulumi.String(organization),
			Project:      pulumi.String(projectVar),
			Stack:        pulumi.String("a"),
			DownstreamRefs: pulumi.StringArray{
				b.Ref,
				c.Ref,
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

{{% choosable language yaml %}}

```yaml
name: auto-deploy-demo
runtime: yaml
description: A simple auto-deploy example
variables:
  project: dependency-example
  # TODO: update once https://github.com/pulumi/pulumi-yaml/issues/461 is fixed
  organization: pulumi
resources:
  f:
    type: auto-deploy:AutoDeployer
    properties:
      organization: ${organization}
      project: ${project}
      stack: f
      downstreamRefs: []
  e:
    type: auto-deploy:AutoDeployer
    properties:
      organization: ${organization}
      project: ${project}
      stack: e
      downstreamRefs: []
  d:
    type: auto-deploy:AutoDeployer
    properties:
      organization: ${organization}
      project: ${project}
      stack: d
      downstreamRefs: []
  c:
    type: auto-deploy:AutoDeployer
    properties:
      organization: ${organization}
      project: ${project}
      stack: c
      downstreamRefs: []
  b:
    type: auto-deploy:AutoDeployer
    properties:
      organization: ${organization}
      project: ${project}
      stack: b
      downstreamRefs:
        - ${d.ref}
        - ${e.ref}
        - ${f.ref}
  a:
    type: auto-deploy:AutoDeployer
    properties:
      organization: ${organization}
      project: ${project}
      stack: a
      downstreamRefs:
        - ${b.ref}
        - ${c.ref}
```

{{% /choosable %}}

{{< /chooser >}}
