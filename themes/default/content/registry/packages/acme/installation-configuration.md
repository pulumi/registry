---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumiverse/pulumi-acme/v0.10.1/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Acme Installation & Configuration
meta_desc: Information on how to install the Acme provider.
layout: package
---

## Installation

The Pulumiverse Acme provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumiverse/acme`](https://www.npmjs.com/package/@pulumiverse/acme)
* Python: [`pulumiverse_acme`](https://pypi.org/project/pulumiverse-acme/)
* Go: [`github.com/pulumiverse/pulumi-acme/sdk/go/acme`](https://pkg.go.dev/github.com/pulumiverse/pulumi-acme/sdk)
* .NET: [`Pulumiverse.Acme`](https://www.nuget.org/packages/Pulumiverse.Acme)

## Setup

To provision resources with the Pulumi Acme provider, you need to have correct ACME endpoint configuration.

## Configuration Options

Use `pulumi config set acme:<option> (--secret)`.

| Option | Environment Variable | Required/Optional | Description | 
|-----|------|------|----|
| `serverUrl`|  | Required | ACME CA endpoint URL |
