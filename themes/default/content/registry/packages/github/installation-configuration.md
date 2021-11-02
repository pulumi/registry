---
title: GitHub Setup
meta_desc: Provides an overview on how to configure credentials for the Pulumi GitHub Provider.
layout: installation
---

The Pulumi GitHub provider uses the GitHub SDK to manage and provision resources.
Pulumi relies on the GitHub SDK to authenticate requests from your computer to GitHub. Your credentials are never sent to pulumi.com.
The Pulumi GitHub Provider needs to be configured with GitHub credentials
before it can be used to create resources.

## Installation

The GitHub provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/github`](https://www.npmjs.com/package/@pulumi/github)
* Python: [`pulumi-github`](https://pypi.org/project/pulumi-github/)
* Go: [`github.com/pulumi/pulumi-github/sdk/v4/go/github`](https://github.com/pulumi/pulumi-github)
* .NET: [`Pulumi.Github`](https://www.nuget.org/packages/Pulumi.Github)

## Configuring Credentials

Once the credetials are obtained, there are two ways to communicate your authorization tokens to Pulumi:

1. Set the environment variable `GITHUB_TOKEN`.

    ```bash
    $ export GITHUB_TOKEN=YYYYYYYYYYYYYY
    ```

2. Set the token value using `pulumi config`. This stores your token alongside your Pulumi stack for easy multi-user access.

    ```bash
    $ pulumi config set github:token XXXXXXXXXXXXXX --secret
    ```

    {{% notes type="info" %}}
Remember to pass `--secret` when setting `github:token` so that it is properly encrypted.
    {{% /notes %}}

3. (Optional) To target a specific GitHub organization or an individual user account, set the GitHub owner configuration value. If this is not provided, the owner of the GitHub Access Token will be the target of the provider.

    ```bash
    $ export GITHUB_OWNER=YYYYYYYYYYYYYY
    ```

    ```bash
    $ pulumi config set github:owner YYYYYYYYYYYYYY
    ```

The complete list of
configuration parameters is in the [GitHub provider README](https://github.com/pulumi/pulumi-github/blob/master/README.md).
