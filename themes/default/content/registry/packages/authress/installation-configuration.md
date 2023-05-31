---
title: Authress Installation & Configuration
meta_desc: Information on how to install the Authress provider.
layout: installation
---

## Installation

The Pulumi Authress provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/authress`](https://www.npmjs.com/package/@pulumi/authress)
* Python: [`pulumi_authress`](https://pypi.org/project/pulumi-authress/)
* Go: [`github.com/authress/pulumi-authress/sdk/go/authress`](https://pkg.go.dev/github.com/authress/pulumi-authress/sdk)
* .NET: [`Pulumi.Authress`](https://www.nuget.org/packages/Pulumi.Authress)

## Setup

To provision resources with the Pulumi Authress provider, you need to provide the `customDomain` and `accessKey`. 

## Configuration Options

Use `pulumi config set authress:<option>`.

| Option    | Required/Optional | Description                                                       |
|-----------|-------------------|-------------------------------------------------------------------|
| `customDomain`     | Required          | Authress Custom Domain. [Configured a custom domain for Account](https://authress.io/app/#/settings?focus=domain)         |
| `accessKey`  | Required          | Authress API access key, automatically injected in via your [CI/CD Provider](https://authress.io/knowledge-base/docs/category/cicd).|
