---
# WARNING: this file was fetched from https://raw.githubusercontent.com/dirien/pulumi-qovery/v0.41.0/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Qovery Setup
meta_desc: Information on how to install the Qovery provider.
layout: package
---

## Installation

The Pulumi Qovery provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@ediri/qovery`](https://www.npmjs.com/package/@ediri/qovery)
* Python: [`ediri_aqovery`](https://pypi.org/project/ediri_qovery/)
* Go: [`github.com/dirien/pulumi-qovery/sdk/go/qovery`](https://github.com/dirien/pulumi-qovery)
* .NET: [`ediri.Qovery`](https://www.nuget.org/packages/ediri.Qovery)

## Configuration Options

The following configuration points are supported for the `Qovery` provider:

* `token` (String) The Qovery API Token to use. This can also be specified with the `QOVERY_API_TOKEN` shell environment variable.
