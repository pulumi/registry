---
# WARNING: this file was fetched from https://ds7waww6k7rw2.cloudfront.net/docs/registry.opentofu.org/zenduty/zenduty/0.2.6/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Zenduty Provider
meta_desc: Provides an overview on how to configure the Pulumi Zenduty provider.
layout: package
---

## Generate Provider

The Zenduty provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider zenduty/zenduty
```
## Overview

The Zenduty provider is used to interact with the zenduty service. The provider needs to be configured with the proper credentials before it can be used.
# Configuration options

The zenduty provider offers two means of providing credentials for authentication.

- Static credentials
- Environment variables
### Static credentials

!> **Warning:** Hard-coding credentials into any Pulumi configuration is not
recommended, and risks secret leakage should this file ever be committed to a
public version control system.

Static credentials can be provided by adding `token` in-line in the Zenduty provider configuration.
### Environment Variables

You can provide your credentials via the `ZENDUTY_API_KEY` environment variables.

Usage:

```sh
$ export ZENDUTY_API_KEY="your-api-key"
$ pulumi preview
```
## Configuration Reference
### Required

- **token** (String) Your Zenduty API key

- **base_url** (String) The base url of the Zenduty