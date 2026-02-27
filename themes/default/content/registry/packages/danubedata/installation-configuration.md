---
# WARNING: this file was fetched from https://raw.githubusercontent.com/AdrianSilaghi/pulumi-danubedata/v0.1.7/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/AdrianSilaghi/pulumi-danubedata/blob/v0.1.7/docs/installation-configuration.md
title: DanubeData Installation & Configuration
meta_desc: Information on how to install and configure the DanubeData provider for Pulumi.
layout: package
---

## Installation

The DanubeData provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@danubedata/pulumi`](https://www.npmjs.com/package/@danubedata/pulumi)
* Python: [`pulumi_danubedata`](https://pypi.org/project/pulumi_danubedata/)
* Go: [`github.com/AdrianSilaghi/pulumi-danubedata/sdk/go/danubedata`](https://github.com/AdrianSilaghi/pulumi-danubedata)
* .NET: [`DanubeData.DanubeData`](https://www.nuget.org/packages/DanubeData.DanubeData)

### Node.js (JavaScript/TypeScript)

```bash
npm install @danubedata/pulumi
```

### Python

```bash
pip install pulumi_danubedata
```

### Go

```bash
go get github.com/AdrianSilaghi/pulumi-danubedata/sdk/go/danubedata
```

### .NET

```bash
dotnet add package DanubeData.DanubeData
```

## Configuration

The DanubeData provider requires an API token for authentication. You can obtain an API token from your [DanubeData account settings](https://danubedata.ro/user/api-tokens).

### Setting the API Token

You can configure the provider using one of the following methods:

#### Environment Variable (Recommended)

Set the `DANUBEDATA_API_TOKEN` environment variable:

```bash
export DANUBEDATA_API_TOKEN="your-api-token-here"
```

#### Pulumi Configuration

```bash
pulumi config set danubedata:apiToken --secret "your-api-token-here"
```

#### Provider Configuration in Code

{{< chooser language "typescript,python,go,csharp" >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as danubedata from "@danubedata/pulumi";

const provider = new danubedata.Provider("danubedata", {
    apiToken: "your-api-token-here",
});

const bucket = new danubedata.StorageBucket("my-bucket", {
    region: "fsn1",
}, { provider });
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_danubedata as danubedata

provider = danubedata.Provider("danubedata",
    api_token="your-api-token-here")

bucket = danubedata.StorageBucket("my-bucket",
    region="fsn1",
    opts=pulumi.ResourceOptions(provider=provider))
```

{{% /choosable %}}

{{% choosable language go %}}

```go
provider, _ := danubedata.NewProvider(ctx, "danubedata", &danubedata.ProviderArgs{
    ApiToken: pulumi.String("your-api-token-here"),
})

bucket, _ := danubedata.NewStorageBucket(ctx, "my-bucket", &danubedata.StorageBucketArgs{
    Region: pulumi.String("fsn1"),
}, pulumi.Provider(provider))
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
var provider = new DanubeData.Provider("danubedata", new()
{
    ApiToken = "your-api-token-here",
});

var bucket = new DanubeData.StorageBucket("my-bucket", new()
{
    Region = "fsn1",
}, new CustomResourceOptions { Provider = provider });
```

{{% /choosable %}}

{{< /chooser >}}

## Configuration Options

| Property | Type | Required | Environment Variable | Description |
|----------|------|----------|---------------------|-------------|
| `apiToken` | `string` | Yes | `DANUBEDATA_API_TOKEN` | API token for DanubeData authentication |
| `baseUrl` | `string` | No | `DANUBEDATA_BASE_URL` | Base URL for the API (default: `https://danubedata.ro/api/v1`) |

## Regions

DanubeData currently supports the following datacenter region:

| Region | Location |
|--------|----------|
| `fsn1` | Falkenstein, Germany |
