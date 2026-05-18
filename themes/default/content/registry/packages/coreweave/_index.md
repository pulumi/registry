---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-coreweave/v1.0.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi/pulumi-coreweave/blob/v1.0.0/docs/_index.md
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: CoreWeave Provider
meta_desc: Provides an overview on how to configure the Pulumi CoreWeave provider.
layout: package
---

## Installation

The CoreWeave provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/coreweave`](https://www.npmjs.com/package/@pulumi/coreweave)
* Python: [`pulumi-coreweave`](https://pypi.org/project/pulumi-coreweave/)
* Go: [`github.com/pulumi/pulumi-coreweave/sdk/go/coreweave`](https://github.com/pulumi/pulumi-coreweave)
* .NET: [`Pulumi.Coreweave`](https://www.nuget.org/packages/Pulumi.Coreweave)
* Java: [`com.pulumi/coreweave`](https://central.sonatype.com/artifact/com.pulumi/coreweave)

## Example Usage

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    coreweave:token:
        value: CW-SECRET-XXXXXXXXXXXXX

```
## Configuration Reference

- `endpoint` (String) CoreWeave API Endpoint. This can also be set via the COREWEAVE_API_ENDPOINT environment variable, which takes precedence. Defaults to `https://api.coreweave.com/`
- `httpTimeout` (String) Timeout duration for the HTTP client to use. This can also be set via the COREWEAVE_HTTP_TIMEOUT environment variable, which takes precedence. If unset, defaults to 10 seconds
- `s3Endpoint` (String) CoreWeave S3 Endpoint, used for CoreWeave Object Storage. This can also be set via the COREWEAVE_S3_ENDPOINT environment variable, which takes precedence. Defaults to `https://cwobject.com`
- `token` (String, Sensitive) CoreWeave API Token in the form `CW-SECRET-<secret>`. This can also be set via the COREWEAVE_API_TOKEN environment variable, which takes precedence.