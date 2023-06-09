---
title: Astra DB Installation & Configuration
meta_desc: Information on how to install the Astra DB provider.
layout: installation
---

## Installation

The Pulumi Astra DB provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumiverse/astra`](https://www.npmjs.com/package/@pulumiverse/astra)
* Python: [`pulumiverse_astra`](https://pypi.org/project/pulumiverse-astra/)
* Go: [`github.com/pulumiverse/pulumi-astra/sdk/go/astra`](https://pkg.go.dev/github.com/pulumiverse/pulumi-astra/sdk)
* .NET: [`Pulumiverse.Astra`](https://www.nuget.org/packages/Pulumiverse.Astra)

## Setup

To provision resources with the Pulumi Astra DB provider, you need to have an Astra DB token. Astra DB maintains documentation on how to [create API tokens](https://docs.datastax.com/en/astra/docs/manage/org/managing-org.html#_manage_application_tokens)

## Configuration Options

Use `pulumi config set astra:<option>`.

| Option | Required/Optional | Description |
|-----|------|----|
| `token`| Required | Astra access token |
