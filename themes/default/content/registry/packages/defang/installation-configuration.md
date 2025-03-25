---
# WARNING: this file was fetched from https://raw.githubusercontent.com/DefangLabs/pulumi-defang/v1.0.0/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Defang Provider for Pulumi Installation & Configuration
meta_desc: Provides an overview on how to configure the Pulumi Defang Provider.
layout: package
---
## Installation

The Pulumi Provider for [Defang](https://defang.io) — Take your app from Docker Compose to a secure and scalable cloud deployment with Pulumi.

The Defang Pulumi Provider is available in most pulumi languages. 

* JavaScript/TypeScript: [`@defang-io/pulumi-defang`](https://www.npmjs.com/package/@defang-io/pulumi-defang)
* Python: [`pulumi-defang`](https://pypi.org/project/pulumi-defang/)
* Go: [`github.com/DefangLabs/pulumi-defang/sdk/v1/go/defang`](https://github.com/DefangLabs/pulumi-defang)
* .NET: [`DefangLabs.Defang`](https://www.nuget.org/packages/DefangLabs.Defang/)
* Java: Coming soon

### Installing the Pulumi Plugin directly
```
pulumi plugin install resource defang --server github://api.github.com/DefangLabs
```

## Authentication

### Authenticating with Defang

Sign up for [Defang](https://defang.io) with your Github account.

#### Authenticating in Github Actions workflows

When run in a Github Actions workflow, the Defang Pulumi Provider will automatically use environment varialbes Github providew to authenticate your Github user with Defang if you give your workflow the [appropriate permissions](https://docs.github.com/en/actions/security-for-github-actions/security-hardening-your-deployments/about-security-hardening-with-openid-connect#adding-permissions-settings). Defang use the `ACTIONS_ID_TOKEN_REQUEST_URL` and `ACTIONS_ID_TOKEN_REQUEST_TOKEN` env vars.

#### Authenticating with `defang token`

You can run `defang token --expires 30d` out of band with a reasonable duration and you can store the result in `DEFANG_ACCESS_TOKEN`.

### Authenticating with your cloud provider

You will also need to authenticate with your cloud provider.

* For AWS, there are many ways to authenticate
    - Use the [`aws-actions/configure-aws-credentials`](https://github.com/aws-actions/configure-aws-credentials) Github Action
    - Use AWS Access Keys by setting the `AWS_ACCESS_KEY_ID`, and `AWS_ACCESS_KEY_SECRET` env vars.
* For Digital Ocean, you will need to set the following env vars:
    - `DIGITALOCEAN_TOKEN`
    - `SPACES_ACCESS_KEY_ID`
    - `SPACES_SECRET_ACCESS_KEY`
* For Google Cloud, you may wish to use the [`google-github-actions/auth`](https://github.com/google-github-actions/auth) Github Action

## Using Pulumi Cloud

Defang runs the Pulumi CLI in your cloud account. You can use [Pulumi Cloud](https://www.pulumi.com/product/pulumi-cloud/) to manage the Pulumi resources which Defang creates by setting the following environment variables:

* `DEFANG_PULUMI_BACKEND=pulumi-cloud`
* `PULUMI_ACCESS_TOKEN`

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/defang/).
