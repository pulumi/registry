---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-aws-static-website/v0.4.0/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: AWS Static Website Installation
meta_desc: Information on how to install the AWS Static Website component.
layout: installation
---

## Installation

The Pulumi AWS Static Website component is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/aws-static-website`](https://www.npmjs.com/package/@pulumi/aws-static-website)
* Python: [`pulumi_aws_static_website`](https://pypi.org/project/pulumi-aws-static-website/)
* Go: [`github.com/pulumi/pulumi-aws-static-website/sdk`](https://pkg.go.dev/github.com/pulumi/pulumi-aws-static-website/sdk)
* .NET: [`Pulumi.AwsStaticWebsite`](https://www.nuget.org/packages/Pulumi.AwsStaticWebsite/)

If you are new to Pulumi and do not yet have our Pulumi installed, see our [getting started](https://www.pulumi.com/docs/get-started/) guide to get up and running.

Install the SDK package corresponding to the language or runtime you are using. If using YAML or one of our other markup languages, this step can be skipped.

### NodeJS

```
npm install @pulumi/aws-static-website
```

### Python

```
pip3 install pulumi_aws_static_website
```

### Go

```
go get -t github.com/pulumi/pulumi-aws-static-website/sdk
```

This package can then be consumed just like any other package by importing it into your project (e.g. `import * as website from '@pulumi/aws-static-website'`)
