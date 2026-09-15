---
# WARNING: this file was fetched from https://raw.githubusercontent.com/jschady/pulumi-filescom/v0.1.1/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/jschady/pulumi-filescom/blob/v0.1.1/docs/installation-configuration.md
title: Files.com Installation & Configuration
meta_desc: Provides an overview on how to configure the Pulumi Files.com provider.
layout: package
---

## Packages

The Files.com provider ships as one package per language.

| Language | Package |
| --- | --- |
| JavaScript and TypeScript | [`pulumi-filescom`](https://www.npmjs.com/package/pulumi-filescom) |
| Python | [`pulumi-filescom`](https://pypi.org/project/pulumi-filescom/) |
| Go | [`github.com/jschady/pulumi-filescom/sdk/go/filescom`](https://pkg.go.dev/github.com/jschady/pulumi-filescom/sdk/go/filescom) |
| .NET | [`Jschady.Filescom`](https://www.nuget.org/packages/Jschady.Filescom) |

## Installation

Run the command for the language your program uses.

### Node.js

```bash
npm install pulumi-filescom
```

### Python

```bash
pip install pulumi-filescom
```

### Go

```bash
go get github.com/jschady/pulumi-filescom/sdk/go/filescom
```

### .NET

```bash
dotnet add package Jschady.Filescom
```

## API key

The provider authenticates with one API key. Create the key in the Files.com web application, under
**Settings**, then **API Keys**.

**Warning:** If your program source holds the API key, everyone who reads it can read your account.
Keep the key in the environment or in the encrypted stack configuration.

Set the key in the environment:

```bash
export FILES_API_KEY=your-api-key
```

Or set it in the stack configuration, where Pulumi encrypts it:

```bash
pulumi config set --secret filescom:apiKey your-api-key
```

The provider reads the environment variable when the stack configuration carries no `apiKey`.

## Provider configuration

| Property | Environment variable | What it is |
| --- | --- | --- |
| `apiKey` | `FILES_API_KEY` | The API key the provider authenticates with. |
| `endpointOverride` | none | The endpoint to call. Set it when your site disables global acceleration, or to reach a mock server. |
| `environment` | none | Files.com documents no text for this property. |
| `featureFlags` | none | Files.com documents no text for this property. |

## Explicit provider

The program below builds one provider from the stack configuration and creates a group with it. Use
this shape when one stack must reach two Files.com accounts.

{{< chooser language "typescript,python,csharp,go" >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as filescom from "pulumi-filescom";

const config = new pulumi.Config();

const files = new filescom.Provider("files", {
    apiKey: config.requireSecret("filesApiKey"),
});

const engineers = new filescom.Group("engineers", {
    name: "engineers",
}, { provider: files });
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_filescom as filescom

config = pulumi.Config()

files = filescom.Provider("files", api_key=config.require_secret("filesApiKey"))

engineers = filescom.Group("engineers",
    name="engineers",
    opts=pulumi.ResourceOptions(provider=files))
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using Pulumi;
using Filescom = Jschady.Filescom;

return await Deployment.RunAsync(() =>
{
    var config = new Config();

    var files = new Filescom.Provider("files", new()
    {
        ApiKey = config.RequireSecret("filesApiKey"),
    });

    var engineers = new Filescom.Group("engineers", new()
    {
        Name = "engineers",
    }, new CustomResourceOptions { Provider = files });
});
```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"github.com/jschady/pulumi-filescom/sdk/go/filescom"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		files, err := filescom.NewProvider(ctx, "files", &filescom.ProviderArgs{
			ApiKey: config.RequireSecret(ctx, "filesApiKey"),
		})
		if err != nil {
			return err
		}
		_, err = filescom.NewGroup(ctx, "engineers", &filescom.GroupArgs{
			Name: pulumi.String("engineers"),
		}, pulumi.Provider(files))
		return err
	})
}
```

{{% /choosable %}}

{{< /chooser >}}
