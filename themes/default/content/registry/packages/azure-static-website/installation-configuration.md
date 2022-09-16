---
title: Azure Static Website Installation
meta_desc: Information on how to install the Azure Static Website component.
layout: installation
---

## Installation

The Pulumi Azure Static Website component is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/azure-static-website`](https://www.npmjs.com/package/@pulumi/azure-static-website)
* Python: [`pulumi_azure_static_website`](https://pypi.org/project/pulumi-azure-static-website/)
* Go: [`github.com/pulumi/pulumi-azure-static-website/sdk`](https://pkg.go.dev/github.com/pulumi/pulumi-azure-static-website/sdk)
* .NET: [`Pulumi.AzureStaticWebsite`](https://www.nuget.org/packages/Pulumi.AzureStaticWebsite/)

If you are new to Pulumi and do not yet have Pulumi installed, see our [getting started](https://www.pulumi.com/docs/get-started/) guide to get up and running.

Install the SDK package corresponding to the language or runtime you are using. If using YAML or one of our other markup languages, this step can be skipped.

### NodeJS

```bash
npm install @pulumi/azure-static-website
```

### Python

```bash
pip3 install pulumi_azure_static_website
```

### Go

```bash
go get -t github.com/pulumi/pulumi-azure-static-website/sdk
```

### .NET
```bash
dotnet add package Pulumi.AzureStaticWebsite
```

This package can then be consumed just like any other package by importing it into your project (e.g. `import * as website from '@pulumi/azure-static-website'`)
