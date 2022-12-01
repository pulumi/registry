---
title: Buildkite Installation & Configuration
meta_desc: Information on how to install the Buildkite provider.
layout: installation
---

## Installation

The Pulumi Buildkite provider is available as a package in the following Pulumi languages:

* JavaScript/TypeScript: [`@grapl/pulumi-buildkite`](https://www.npmjs.com/package/@grapl/pulumi-buildkite)
* Python: [`pulumi_buildkite`](https://pypi.org/project/pulumi-buildkite/)
* Go: [`github.com/graphl-security/pulumi-buildkite/sdk`](https://pkg.go.dev/github.com/grapl-security/pulumi-buildkite/sdk)

## Setup

To provision resources with the Pulumi Buildkite provider, you need to have
Buildkite credentials.

### Set environment variables

Once you have provisioned these credentials, you can set environment
variables to provision resources in BUILDKITE:

{{< chooser os "linux,macos,windows" >}}
{{% choosable os linux %}}

```bash
$ export BUILDKITE_API_TOKEN=<BUILDKITE_API_TOKEN>
$ export BUILDKITE_ORGANIZATION=<BUILDKITE_ORGANIZATION>
```

{{% /choosable %}}

{{% choosable os macos %}}

```bash
$ export BUILDKITE_API_TOKEN=<BUILDKITE_API_TOKEN>
$ export BUILDKITE_ORGANIZATION=<BUILDKITE_ORGANIZATION>
```

{{% /choosable %}}

{{% choosable os windows %}}

```powershell
> $env:BUILDKITE_API_TOKEN = "<BUILDKITE_API_TOKEN>"
> $env:BUILDKITE_ORGANIZATION = "<BUILDKITE_ORGANIZATION>"
```

{{% /choosable %}}
{{< /chooser >}}

## Configuration Options

Use `pulumi config set buildkite:<option>` or pass options to the [constructor of `new buildkite.Provider`]({{< relref "/registry/packages/buildkite/api-docs/provider" >}}).

| Option          | Required/Optional | Description                                                                                                       |
|-----------------|-------------------|-------------------------------------------------------------------------------------------------------------------|
| `api_token`     | Required          | A Buildkite API Access Token. Can be configured from the environment variable `BUILDKITE_API_TOKEN`. Must have GraphQL access, as well as the `write_pipelines` and `read_pipelines` scopes. |
| `organization`  | Required          | The Buildkite organization slug. Can be configured from the environment variable `BUILDKITE_ORGANIZATION`. |
