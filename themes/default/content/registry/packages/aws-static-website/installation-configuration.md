---
title: AWS Static Website Installation
meta_desc: Information on how to install the AWS Static Website component.
layout: installation
---

{{< aws-resource-note >}}

## Installation

The Pulumi AWS Static Website component is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/aws-static-website`](https://www.npmjs.com/package/@pulumi/aws-static-website)
* Python: [`pulumi_aws_static_website`](https://pypi.org/project/pulumi-aws-static-website/)
* Go: [`github.com/pulumi/pulumi-aws-static-website/sdk/go/aws-static-website`](https://pkg.go.dev/github.com/pulumi/pulumi-aws-static-website)
* .NET: [`Pulumi.AwsStaticWebsite`](https://www.nuget.org/packages/Pulumi.AwsStaticWebsite/)

If you are new to Pulumi, see our [getting started](https://www.pulumi.com/docs/get-started/) guide to get up and running.

Install the SDK package corresponding to the language or runtime you are using. If using YAML or one of our other markup languages, this step can be skipped.

### NodeJS

```
npm install @pulumi/aws-static-website
```

### Python

```
pip install pulumi-aws-static-website
```

### Go

```
go get -t github.com/pulumi/pulumi-aws-static-website/sdk/go/aws-static-website
```

This package can then be consumed just like any other package by importing it into your project (e.g. `import * as website from '@pulumi/aws-static-website'`)
