---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumiverse/pulumi-buildkite/v3.2.0/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumiverse/pulumi-buildkite/blob/v3.2.0/docs/installation-configuration.md
title: Buildkite Installation & Configuration
meta_desc: Information on how to install the Buildkite provider.
layout: package
---

## Installation

The Buildkite provider is available as a package in all Pulumi languages:

<!-- x-release-please-start-major -->
* JavaScript/TypeScript: [`@pulumiverse/buildkite`](https://www.npmjs.com/package/@pulumiverse/buildkite)
* Python: [`pulumiverse-buildkite`](https://pypi.org/project/pulumiverse-buildkite/)
* Go: [`https://pkg.go.dev/github.com/pulumiverse/pulumi-buildkite/sdk/v2/go/buildkite`](https://pkg.go.dev/github.com/pulumiverse/pulumi-buildkite/sdk/v2/go/buildkite)
* .NET: [`Pulumiverse.Buildkite`](https://www.nuget.org/packages/Pulumiverse.Buildkite)
<!-- x-release-please-end -->

## Setup

To provision resources with the Pulumi Buildkite provider, you need to have Buildkite credentials.

## Configuration Options

Pulumi relies on the Buildkite API to authenticate requests from your computer to Buildkite. Your credentials are never sent to pulumi.com.
The Pulumi Buildkite Provider needs to be configured with a Buildkite token before it can be used to create resources.

Use `pulumi config set buildkite:<option>` or pass options to the [constructor of `new buildkite.Provider`](/registry/packages/buildkite/api-docs/provider).

| Option          | Required/Optional | Description                                                                                                       |
|-----------------|-------------------|-------------------------------------------------------------------------------------------------------------------|
| `apiToken`      | Required          | A Buildkite API Access Token. Can be configured from the environment variable `BUILDKITE_API_TOKEN`. Must have GraphQL access, as well as the `write_pipelines` and `read_pipelines` scopes. |
| `archivePipelineOnDelete` | Optional | Enable this to archive pipelines when destroying the resource. This is opposed to completely deleting pipelines. |
| `organization`  | Required          | The Buildkite organization slug. Can be configured from the environment variable `BUILDKITE_ORGANIZATION`. |
| `graphqlUrl`    | Optional          | The Buildkite GraphQL URL. Can be configured from the environment variable `BUILDKITE_GRAPHQL_URL`. |
| `restUrl`       | Optional          | The Buildkite REST URL. Can be configured from the environment variable `BUILDKITE_REST_URL`. |
| `timeouts`      | Optional          | A Timeout object (see below) |

### `Timeout` Object

| Field    | Description     |
|----------|-----------------|
| `create` | A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). |
| `delete` | Value as above. Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs. |
| `read`   | Value as above. Read operations occur during any refresh or planning operation when refresh is enabled. |
| `update` | Value as above. |
