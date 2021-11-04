---
title: GitLab Setup
meta_desc: Provides an overview on how to configure credentials for the Pulumi GitLab Provider.
layout: installation
---

The Pulumi GitLab provider uses the GitLab SDK to manage and provision resources.

## Installation

The GitLab provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/gitlab`](https://www.npmjs.com/package/@pulumi/gitlab)
* Python: [`pulumi-gitlab`](https://pypi.org/project/pulumi-gitlab/)
* Go: [`github.com/pulumi/pulumi-gitlab/sdk/v3/go/gitlab`](https://github.com/pulumi/pulumi-gitlab)
* .NET: [`Pulumi.Gitlab`](https://www.nuget.org/packages/Pulumi.Gitlab)

## Configuring Credentials

Pulumi relies on the GitLab SDK to authenticate requests from your computer to GitLab. Your credentials are never sent
to pulumi.com.
The Pulumi GitLab Provider needs to be configured with GitLab credentials
before it can be used to create resources.

Once the credetials are obtained, there are two ways to communicate your authorization tokens to Pulumi:

1. Set the environment variable `GITLAB_TOKEN`:

    ```bash
    $ export GITLAB_TOKEN=XXXXXXXXXXXXXX
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set gitlab:token XXXXXXXXXXXXXX --secret
    ```

Remember to pass `--secret` when setting `gitlab:token` so that it is properly encrypted. The complete list of
configuration parameters is in the [GitLab provider README](https://github.com/pulumi/pulumi-gitlab/blob/master/README.md).
