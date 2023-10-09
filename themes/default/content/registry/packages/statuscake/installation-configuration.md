---
title: Statuscake Installation & Configuration
meta_desc: Information on how to install the Statuscake provider.
layout: installation
---

## Installation

The Pulumi Statuscake provider is available as a package in all Pulumi languages:

- JavaScript/TypeScript: [`@pulumiserve/statuscake`](https://www.npmjs.com/package/@pulumiverse/statuscake)
- Python: [`pulumiverse_statuscake`](https://pypi.org/project/pulumiverse-statuscake/)
- Go: [`github.com/pulumiverse/pulumi-statuscake/sdk`](https://pkg.go.dev/github.com/pulumiverse/pulumi-statuscake/sdk)
- .NET: [`Pulumiverse.Statuscake`](https://www.nuget.org/packages/Pulumiverse.Statuscake)

## Setup

To provision resources with the Pulumi Statuscake provider, you need to have a Statuscake API token.
Statuscake maintains documentation on how to [create API tokens](https://www.statuscake.com/blog/how-to-use-the-statuscake-api/).

## Configuration Options

Use `pulumi config set statuscake:<option>`.

| Option     | Required/Optional | Description          |
| ---------- | ----------------- | -------------------- |
| `apiToken` | Required          | Statuscake API token |
