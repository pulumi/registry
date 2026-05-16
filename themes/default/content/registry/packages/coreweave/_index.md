---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-coreweave/v0.2.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi/pulumi-coreweave/blob/v0.2.0/docs/_index.md
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: CoreWeave Provider
meta_desc: Provides an overview on how to configure the Pulumi CoreWeave provider.
layout: package
---

## Installation

The CoreWeave provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/coreweave`](https://www.npmjs.com/package/@pulumi/coreweave)
* Python: [`pulumi-coreweave`](https://pypi.org/project/pulumi-coreweave/)
* Go: [`github.com/pulumi/pulumi-coreweave/sdk/v4/go/coreweave`](https://github.com/pulumi/pulumi-coreweave)

* .NET: [`Pulumi.CoreWeave`](https://www.nuget.org/packages/Pulumi.CoreWeave)


## Overview

The "CoreWeave" provider allows the use of CoreWeave cloud resources within Pulumi

## Configuration

The CoreWeave provider has the following configuration options:

| Option | Required | Description |
|--------|----------|-------------|
| `coreweave:token` | Optional | CoreWeave API token in the form `CW-SECRET-<secret>`. Can also be set via the `COREWEAVE_API_TOKEN` environment variable, which takes precedence. |
| `coreweave:endpoint` | Optional | CoreWeave API endpoint. Can also be set via the `COREWEAVE_API_ENDPOINT` environment variable, which takes precedence. Defaults to `https://api.coreweave.com/`. |
| `coreweave:httpTimeout` | Optional | Timeout duration for the HTTP client (e.g. `30s`). Can also be set via the `COREWEAVE_HTTP_TIMEOUT` environment variable, which takes precedence. Defaults to `10s`. |
| `coreweave:s3Endpoint` | Optional | CoreWeave S3 endpoint, used for CoreWeave Object Storage. Can also be set via the `COREWEAVE_S3_ENDPOINT` environment variable, which takes precedence. Defaults to `https://cwobject.com`. |