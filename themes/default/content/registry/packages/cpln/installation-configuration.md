---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumiverse/pulumi-cpln/v0.0.68/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Control Plane (cpln) Installation & Configuration
meta_desc: Information on how to install the Control Plane (cpln) provider.
layout: package
---

## Installation

The Pulumi Control Plane (cpln) provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumiserve/cpln`](https://www.npmjs.com/package/@pulumiverse/cpln)
* Python: [`pulumiverse-cpln`](https://pypi.org/project/pulumiverse-cpln/)
* Go: [`github.com/pulumiverse/pulumi-cpln/sdk/go/cpln`](https://pkg.go.dev/github.com/pulumiverse/pulumi-cpln/sdk)
* .NET: [`Pulumiverse.cpln`](https://www.nuget.org/packages/Pulumiverse.cpln)

## Setup

To provision resources with the Pulumi Control Plane (cpln) provider, you need to authenticate with one of the available options.
1. CLI
    - [Install the CLI](https://docs.controlplane.com/reference/cli#installation) and execute the command `cpln login`. After a successful login, the Pulumi provider will use the `default` profile to authenticate. To use a different profile, set the `profile` variable when initializing the provider or set the `CPLN_PROFILE` environment variable.
2. Token
    - The `token` variable can be set when initializing the provider or by setting the `CPLN_TOKEN` environment variable.
    - The value of `token` can be either:
        - The output of running the command `cpln profile token PROFILE_NAME`, or
        - In the case of a [Service Account](https://docs.controlplane.com/reference/serviceaccount), the value of one of it's [keys](https://docs.controlplane.com/reference/serviceaccount#keys)
3. Refresh Token
    - The `refreshToken` variable is used when the provider is required to create an org or update the `authConfig` property using the `Org` resource. The `refreshToken` variable can be set when initializing the provider or by setting the `CPLN_REFRESH_TOKEN` environment variable.
    - When creating an org, the `refreshToken` **must** belong to a user that has the org_creator role for the associated account.
    - When updating the org `authConfig` property, the `refreshToken` **must** belong to a user that was authenticated using SAML.
    - The `refreshToken` can be obtained by following these steps:
        - Using the CLI, authenticate with a user account by executing `cpln login`.
        - Browser to the path `~/.config/cpln/profiles`. This path will contain JSON files corresponding to the name of the profile (i.e., `default.json`).
        - The contents of the JSON file will contain a key named `refreshToken`. Use the value of this key for the `refreshToken` variable.

To perform automated tasks using Pulumi, the preferred method is to use a `Service Account` and one of it's `keys` as the `token` value.

## Configuration Options

Use `pulumi config set cpln:<option>`.

| Option | Required/Optional | Description |
|-----|------|----|
| `org`| Required | The Control Plane org that this provider will perform actions against. Can be specified with the `CPLN_ORG` environment variable. |
| `endpoint`| Optional | The Control Plane Data Service API endpoint. Default is: `https://api.cpln.io`. Can be specified with the `CPLN_ENDPOINT` environment variable. |
| `profile`| Optional | The user/service account profile that this provider will use to authenticate to the data service. Can be specified with the `CPLN_PROFILE` environment variable. |
| `token`| Optional | A generated token that can be used to authenticate to the data service API. Can be specified with the `CPLN_TOKEN` environment variable |
| `refreshToken`| Optional | A generated token that can be used to authenticate to the data service API. Can be specified with the `CPLN_REFRESH_TOKEN` environment variable. Used when the provider is required to create an org or update the `authConfig` property. |
