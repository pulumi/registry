---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/hashicorp/local/2.5.3-alpha1/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Local Provider
meta_desc: Provides an overview on how to configure the Pulumi Local provider.
layout: package
---

## Generate Provider

The Local provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider hashicorp/local
```
## Overview

The Local provider is used to manage local resources, such as files.

Use the navigation to the left to read about the available resources.

> **Note** Pulumi primarily deals with remote resources which are able
to outlive a single Pulumi run, and so local resources can sometimes violate
its assumptions. The resources here are best used with care, since depending
on local state can make it hard to apply the same Pulumi configuration on
many different local systems where the local resources may not be universally
available. See specific notes in each resource for more information.