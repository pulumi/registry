---
title: Chronosphere Installation & Configuration
meta_desc: How to install and configure the Chronosphere provider.
layout: installation
---

## Installation

The Chronosphere provider is available in the following languages:

* JavaScript/TypeScript: [`@pulumi-chronosphere/pulumi-chronosphere`](https://www.npmjs.com/package/@pulumi-chronosphere/pulumi-chronosphere)
* Python: [`pulumi_chronosphere`](https://pypi.org/project/pulumi-chronosphere/)
* Go: [`github.com/chronosphereio/pulumi-chronosphere/sdk/go/chronosphere`](https://pkg.go.dev/github.com/chronosphereio/pulumi-chronosphere/sdk)
* .NET: [`Chronosphere.Pulumi.Chronosphere`](https://www.nuget.org/packages/Chronosphere.Pulumi.Chronosphere)

## Configuration

| Configuration Key | Description | Environment Variables |
|-------------------|-------------|-----------------------|
| `org`               | The organization name for Chronosphere. | `CHRONOSPHERE_ORG`, `CHRONOSPHERE_ORG_NAME` |
| `api_token`         | The API token for accessing Chronosphere. | `CHRONOSPHERE_API_TOKEN` |
