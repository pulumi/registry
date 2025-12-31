---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-sdwan/v0.5.2/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi/pulumi-sdwan/blob/v0.5.2/docs/_index.md
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Cisco Catalyst SD-WAN Provider
meta_desc: Provides an overview on how to configure the Pulumi Cisco Catalyst SD-WAN provider.
layout: package
---

## Installation

The Cisco Catalyst SD-WAN provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/sdwan`](https://www.npmjs.com/package/@pulumi/sdwan)
* Python: [`pulumi-sdwan`](https://pypi.org/project/pulumi-sdwan/)
* Go: [`github.com/pulumi/pulumi-sdwan/sdk/go/sdwan`](https://github.com/pulumi/pulumi-sdwan)
* .NET: [`Pulumi.Sdwan`](https://www.nuget.org/packages/Pulumi.Sdwan)
* Java: [`com.pulumi/sdwan`](https://central.sonatype.com/artifact/com.pulumi/sdwan)

## Overview

The SDWAN provider provides resources to interact with a Cisco Catalyst SD-WAN environment. It communicates with the SD-WAN Manager via the REST API.

All resources and functions have been tested with the following releases.

|    Platform     | Version |
|-----------------|---------|
| Catalyst SD-WAN |   20.09 |
| Catalyst SD-WAN |   20.12 |
## Example Usage

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    sdwan:password:
        value: password
    sdwan:url:
        value: https://10.1.1.1
    sdwan:username:
        value: admin

```
## Configuration Reference

- `insecure` (Boolean) Allow insecure HTTPS client. This can also be set as the `SDWAN_INSECURE` environment variable. Defaults to `true`.
- `password` (String, Sensitive) Password for the SD-WAN Manager account. This can also be set as the `SDWAN_PASSWORD` environment variable.
- `retries` (Number) Number of retries for REST API calls. This can also be set as the `SDWAN_RETRIES` environment variable. Defaults to `3`.
- `taskTimeout` (Number) Timeout in seconds for asynchronous tasks. This can also be set as the `SDWAN_TASK_TIMEOUT` environment variable. Defaults to `1500`.
- `url` (String) URL of the Cisco SD-WAN Manager device. This can also be set as the `SDWAN_URL` environment variable.
- `username` (String) Username for the SD-WAN Manager account. This can also be set as the `SDWAN_USERNAME` environment variable.