---
# WARNING: this file was fetched from https://raw.githubusercontent.com/kurrent-io/pulumi-kurrentcloud/v0.3.1/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/kurrent-io/pulumi-kurrentcloud/blob/v0.3.1/docs/_index.md
title: Kurrent Cloud
meta_desc: Use the Kurrent Cloud provider for Pulumi to manage Kurrent Cloud projects, networks, peerings, managed clusters, backups and integrations.
layout: package
---

The Kurrent Cloud provider for Pulumi lets you manage [Kurrent Cloud](https://www.kurrent.io/kurrent-cloud) resources, including projects, networks, peerings, managed KurrentDB clusters, read-only replica sets, scheduled backups and integrations, as part of your Pulumi programs.

## Installation

{{< chooser language "typescript,python,go,csharp,yaml" >}}
{{% choosable language typescript %}}
```bash
npm install @kurrent/pulumi-kurrentcloud
```
{{% /choosable %}}
{{% choosable language python %}}
```bash
pip install pulumi_kurrentcloud
```
{{% /choosable %}}
{{% choosable language go %}}
```bash
go get github.com/kurrent-io/pulumi-kurrentcloud/sdk/go/kurrentcloud
```
{{% /choosable %}}
{{% choosable language csharp %}}
```bash
dotnet add package Kurrent.Pulumi.KurrentCloud
```
{{% /choosable %}}
{{% choosable language yaml %}}
```bash
pulumi package add kurrentcloud
```
{{% /choosable %}}
{{< /chooser >}}

Each SDK downloads the provider plugin the first time a program runs. To install it ahead of time, for example in a CI image:

```bash
# The latest release
pulumi plugin install resource kurrentcloud --server github://api.github.com/kurrent-io
# A specific release
pulumi plugin install resource kurrentcloud v0.3.1 --server github://api.github.com/kurrent-io
```

## Example Usage

Each program creates a Kurrent Cloud project. Configure the provider first:

```bash
pulumi config set kurrentcloud:organizationId <YOUR_ORGANIZATION_ID>
pulumi config set --secret kurrentcloud:token <YOUR_ACCESS_TOKEN>
```

{{< chooser language "typescript,python,go,csharp,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as kurrent from "@kurrent/pulumi-kurrentcloud";

const project = new kurrent.Project("sample-project", {
    name: "sample-project",
});

export const projectId = project.id;
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_kurrentcloud as kurrent

project = kurrent.Project("sample-project", name="sample-project")

pulumi.export("project_id", project.id)
```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/kurrent-io/pulumi-kurrentcloud/sdk/go/kurrentcloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		project, err := kurrentcloud.NewProject(ctx, "sample-project", &kurrentcloud.ProjectArgs{
			Name: pulumi.String("sample-project"),
		})
		if err != nil {
			return err
		}
		ctx.Export("projectId", project.ID())
		return nil
	})
}
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using Pulumi;
using KurrentCloud = Kurrent.Pulumi.KurrentCloud;

return await Deployment.RunAsync(() =>
{
    var project = new KurrentCloud.Project("sample-project", new()
    {
        Name = "sample-project",
    });

    return new Dictionary<string, object?>
    {
        ["projectId"] = project.Id,
    };
});
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
name: sample
runtime: yaml
resources:
  sample-project:
    type: kurrentcloud:index:Project
    properties:
      name: sample-project
outputs:
  projectId: ${sample-project.id}
```
{{% /choosable %}}
{{< /chooser >}}

## Configuration

The provider needs an organization ID and one set of credentials. The organization ID is on the organization's settings page in the Kurrent Cloud console. For credentials, use either:

- **A Service Account** (recommended for automation): set `clientId` and `clientSecret` to the Service Account's credentials. When both are set, the provider uses them and ignores `token`.
- **An access token**: create one in the *Access Tokens* section of the console and set `token`.

Set them with `pulumi config` (`pulumi config set --secret kurrentcloud:clientSecret <YOUR_CLIENT_SECRET>`), or with environment variables:

```bash
export ESC_ORG_ID=<YOUR_ORGANIZATION_ID>
# Service Account
export ESC_CLIENT_ID=<YOUR_CLIENT_ID>
export ESC_CLIENT_SECRET=<YOUR_CLIENT_SECRET>
# or an access token
export ESC_TOKEN=<YOUR_ACCESS_TOKEN>
```

| Name | Required | Secret | Description |
| --- | --- | --- | --- |
| `token` | Unless `clientId` and `clientSecret` are set | Yes | Access token from the Kurrent Cloud console. Environment variable: `ESC_TOKEN`. |
| `organizationId` | Yes | No | ID of the Kurrent Cloud organization to manage. Environment variable: `ESC_ORG_ID`. |
| `url` | No | No | URL of the Kurrent Cloud API. Defaults to `https://api.eventstore.cloud`. Environment variable: `ESC_URL`. |
| `tokenStore` | No | No | Local directory where access tokens are cached, shared with the Kurrent Cloud CLI. Defaults to `~/.esctf/tokens`. Environment variable: `ESC_TOKEN_STORE`. |
| `identityProviderUrl` | No | No | Identity provider used to exchange the token. Leave unset unless Kurrent support asks you to change it. Environment variable: `ESC_IDENTITY_PROVIDER_URL`. |
| `clientId` | For a Service Account | No | Service Account client ID. Set together with `clientSecret`. Environment variable: `ESC_CLIENT_ID`. |
| `clientSecret` | For a Service Account | Yes | Service Account client secret. When `clientId` and `clientSecret` are both set, they take priority over `token`. Environment variable: `ESC_CLIENT_SECRET`. |
| `identityKitUrl` | No | No | Token endpoint for Service Account login. Defaults to `https://auth.kurrent.io`. Environment variable: `ESC_IDENTITY_KIT_URL`. |

## Migrating from `eventstorecloud`

This provider was previously published as `eventstorecloud`. Every resource declares an alias to its old type, so existing stacks move to `kurrentcloud` without replacing any resources. See the [migration guide](https://github.com/kurrent-io/pulumi-kurrentcloud/blob/main/MIGRATION.md).
