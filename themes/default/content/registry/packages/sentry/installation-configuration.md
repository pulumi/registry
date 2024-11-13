---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumiverse/pulumi-sentry/v0.0.8/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Sentry Installation & Configuration
meta_desc: Information on how to install the Sentry provider.
layout: installation
---

## Installation

The Pulumi Sentry provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumiserve/sentry`](https://www.npmjs.com/package/@pulumiverse/sentry)
* Python: [`pulumiverse_sentry`](https://pypi.org/project/pulumiverse-sentry/)
* Go: [`github.com/pulumiverse/pulumi-sentry/sdk/go/sentry`](https://pkg.go.dev/github.com/pulumiverse/pulumi-sentry/sdk)
* .NET: [`Pulumiverse.Sentry`](https://www.nuget.org/packages/Pulumiverse.Sentry)

## Setup

To provision resources with the Pulumi Sentry provider, you need to have a Sentry Auth token. 
Sentry maintains documentation on how to [create auth tokens](https://docs.sentry.io/api/auth/#auth-tokens).

## Configuration Options

Use `pulumi config set sentry:<option>`.

| Option | Required/Optional | Description |
|-----|------|----|
| `token`| Required | Sentry access token |
| `base_url`| Optional | Sentry base URL when self-hosting Sentry |
