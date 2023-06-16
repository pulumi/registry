---
title: Buildkite Installation & Configuration
meta_desc: Information on how to install the Buildkite provider.
layout: package
---

## Installation

The Buildkite provider is available as a package in all Pulumi languages:

<!-- x-release-please-start-major -->
* JavaScript/TypeScript: [`@pulumiverse/pulumi-buildkite`](https://www.npmjs.com/package/@pulumiverse/pulumi-buildkite)
* Python: [`pulumiverse-buildkite`](https://pypi.org/project/pulumiverse-buildkite/)
* Go: [`github.com/pulumiverse/pulumi-buildkite/sdk/v2/go/buildkite`](https://github.com/pulumiverse/pulumi-buildkite)
* .NET: [`Pulumiverse.Buildkite`](https://www.nuget.org/packages/Pulumiverse.Buildkite)
<!-- x-release-please-end -->

## Setup

To provision resources with the Pulumi Buildkite provider, you need to have Buildkite credentials.

## Configuration Options

Pulumi relies on the Buildkite API to authenticate requests from your computer to Buildkite. Your credentials are never sent to pulumi.com.
The Pulumi Buildkite Provider needs to be configured with a Buildkite token before it can be used to create resources.

Use `pulumi config set buildkite:<option>` or pass options to the [constructor of `new buildkite.Provider`]({{< relref "/registry/packages/buildkite/api-docs/provider" >}}).

| Option          | Required/Optional | Description                                                                                                       |
|-----------------|-------------------|-------------------------------------------------------------------------------------------------------------------|
| `api_token`     | Required          | A Buildkite API Access Token. Can be configured from the environment variable `BUILDKITE_API_TOKEN`. Must have GraphQL access, as well as the `write_pipelines` and `read_pipelines` scopes. |
| `organization`  | Required          | The Buildkite organization slug. Can be configured from the environment variable `BUILDKITE_ORGANIZATION`. |
| `graphql_url`  | Optional          | The Buildkite GraphQL URL. Can be configured from the environment variable `BUILDKITE_GRAPHQL_URL`. |
| `rest_url`  | Optional          | The Buildkite REST URL. Can be configured from the environment variable `BUILDKITE_REST_URL`. |
