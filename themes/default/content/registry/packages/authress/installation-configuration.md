---
title: Authress Installation & Configuration
meta_desc: Information on how to install the Authress provider.
layout: package
---

## Installation

The Pulumi Authress provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/authress`](https://www.npmjs.com/package/@pulumi/authress)
* Python: [`pulumi_authress`](https://pypi.org/project/pulumi-authress/)
* Go: [`github.com/authress/pulumi-authress/sdk/go/authress`](https://pkg.go.dev/github.com/authress/pulumi-authress/sdk)
* .NET: [`Pulumi.Authress`](https://www.nuget.org/packages/Pulumi.Authress)

## Configuration Options

To provision resources with the Pulumi Authress provider, you need to provide the `customDomain` and `accessKey`. 

Pulumi relies on the Authress SDK to authenticate requests from your computer to Authress. Your credentials are never sent to pulumi.com. The Pulumi Authress Provider needs to be configured with Authress credentials before it can be used to create resources. Once the credentials are obtained, there are two ways to communicate your authorization tokens to Pulumi:

1. CI / CD. Automatically inject in the credentials from your CI/CD platform such as GitHub or GitLab.

See [setting up GitHub/GitLab CI/CD credentials for Pulumi](https://authress.io/knowledge-base/docs/category/cicd).

2. Directly specify the credentials using the `pulumi config set`


Use `pulumi config set authress:<option>`.

| Option    | Required | Description                                                       |
|-----------|-------------------|-------------------------------------------------------------------|
| `customDomain`     | Required          | Authress Custom Domain. [Configure a custom domain for your Authress account](https://authress.io/app/#/settings?focus=domain)         |
| `accessKey`        | Required (Automatically Injected)          | Authress API access key, automatically injected in via your [CI/CD Provider](https://authress.io/knowledge-base/docs/category/cicd).|

When using on your local machine, set the accessKey using `pulumi config set --secret authress:accessKey ACCESS_KEY`. `--secret` is required to keep these credentials a secret. To get an access key Create an [Authress Service Client](https://authress.io/app/#/settings?focus=clients).
