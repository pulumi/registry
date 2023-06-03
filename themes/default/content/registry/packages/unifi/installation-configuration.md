---
title: Unifi Installation & Configuration
meta_desc: Information on how to install the Unifi provider.
layout: package
---

## Installation

The Pulumiverse Unifi provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumiverse/unifi`](https://www.npmjs.com/package/@pulumiverse/unifi)
* Python: [`pulumiverse_unifi`](https://pypi.org/project/pulumiverse-unifi/)
* Go: [`github.com/pulumiverse/pulumi-unifi/sdk/go/unifi`](https://pkg.go.dev/github.com/pulumiverse/pulumi-unifi/sdk)
* .NET: [`Pulumiverse.Unifi`](https://www.nuget.org/packages/Pulumiverse.Unifi)

## Setup

To provision resources with the Pulumi Unifi provider, you need to have correct endpoint configuration towards your Unifi controller.

## Configuration Options

Use `pulumi config set unifi:<option> (--secret)`.

| Option | Environment Variable | Required/Optional | Description | 
|-----|------|------|----|
| `username`| `UNIFI_USERNAME` | Required | Unifi user name |
| `password`| `UNIFI_PASSWORD` | Required (Secret) | Unifi Password |
| `api_url` | `UNIFI_API` | Required | Unifi user name |
| `allowInsecure` | `UNIFI_INSECURE` | Optional | Do not check HTTPS certificate |
