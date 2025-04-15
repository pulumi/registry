---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/rancher/rke/1.7.5/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Rke Provider
meta_desc: Provides an overview on how to configure the Pulumi Rke provider.
layout: package
---

## Generate Provider

The Rke provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider rancher/rke
```
~> **NOTE:** This provider was previously published as @pulumi/rke.
However, that package is no longer being updated.Going forward, it is available as a [Local Package](https://www.pulumi.com/blog/any-terraform-provider/) instead.
Please see the [provider's repository](https://github.com/pulumi/pulumi-rke) for details.

## Overview

The RKE provider is used to interact with Rancher Kubernetes Engine kubernetes clusters.
## Example Usage

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    rke:debug:
        value: true
    rke:logFile:
        value: <rke_log_file>

```
## Configuration Reference

The following configuration inputs are supported:

* `debug` - (Optional) Enable RKE debug logs. It can also be sourced from the `RKE_DEBUG` environment variable. Default `false` (bool)
* `logFile` - (Optional) Save RKE logs to a file. It can also be sourced from the `RKE_LOG_FILE` environment variable (string)