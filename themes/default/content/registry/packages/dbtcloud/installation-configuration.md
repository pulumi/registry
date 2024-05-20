---
title: dbt Cloud Installation & Configuration
meta_desc: Information on how to install the dbt Cloud provider.
layout: package
---

## Installation

This package is available for several languages/platforms:

- JavaScript/TypeScript: [`@pulumi/dbtcloud`](https://www.npmjs.com/package/@pulumi/dbtcloud)
- Python: [`pulumi-dbtcloud`](https://pypi.org/project/pulumi-dbtcloud/)
- Go: [`github.com/pulumi/pulumi-dbtcloud/sdk/go/dbtcloud`](https://pkg.go.dev/github.com/pulumi/pulumi-dbtcloud/sdk/go/dbtcloud)
- .NET: [`Pulumi.DbtCloud`](https://www.nuget.org/packages/Pulumi.DbtCloud)

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @pulumi/dbtcloud
```

or `yarn`:

```bash
yarn add @pulumi/dbtcloud
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi-dbtcloud
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/pulumi/pulumi-dbtcloud/sdk/go/dbtcloud
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Pulumi.DbtCloud
```

## Configuration Options

Use `pulumi config set dbtcloud:<option> (--secret)`.

| Option     | Environment Variable   | Required/Optional | Default                                                      | Description                             |
|------------|------------------------|-------------------|--------------------------------------------------------------|-----------------------------------------|
| `token`    | `DBT_CLOUD_TOKEN`      | Required          |                                                              | The API token for your dbt Cloud user   |
| `accountId`| `DBT_CLOUD_ACCOUNT_ID` | Required          |                                                              | The ID for your dbt Cloud account       |
| `hostUrl`  | `DBT_CLOUD_HOST_URL`   | Optional          | [https://cloud.getdbt.com/api](https://cloud.getdbt.com/api) | The host URL for your dbt Cloud account |
