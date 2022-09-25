---
title: Google Cloud Static Website Installation
meta_desc: Information on how to install the Google Cloud Static Website component.
layout: installation
---

## Installation

The Pulumi Google Cloud Static Website component is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/google-cloud-static-website`](https://www.npmjs.com/package/@pulumi/google-cloud-static-website)
* Python: [`pulumi_google_cloud_static_website`](https://pypi.org/project/pulumi-google-cloud-static-website/)
* Go: [`github.com/pulumi/pulumi-google-cloud-static-website/sdk`](https://pkg.go.dev/github.com/pulumi/pulumi-google-cloud-static-website/sdk)
* .NET: [`Pulumi.GoogleCloudStaticWebsite`](https://www.nuget.org/packages/Pulumi.GoogleCloudStaticWebsite/)

If you are new to Pulumi and do not yet have Pulumi installed, see our [getting started](https://www.pulumi.com/docs/get-started/) guide to get up and running.

Install the SDK package corresponding to the language or runtime you are using. If using YAML or one of our other markup languages, this step can be skipped.

### NodeJS

```bash
npm install @pulumi/google-cloud-static-website
```

### Python

```bash
pip3 install pulumi_google_cloud_static_website
```

### Go

```bash
go get -t github.com/pulumi/pulumi-google-cloud-static-website/sdk
```

### .NET
```bash
dotnet add package Pulumi.GoogleCloudStaticWebsite
```

This package can then be consumed just like any other package by importing it into your project (e.g. `import * as website from '@pulumi/google-cloud-static-website'`)
