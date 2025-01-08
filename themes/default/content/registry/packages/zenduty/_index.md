---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/zenduty/zenduty/0.2.6/index.md
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
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}

{{% /choosable %}}
{{% choosable language python %}}

{{% /choosable %}}
{{% choosable language csharp %}}

{{% /choosable %}}
{{% choosable language go %}}

{{% /choosable %}}
{{% choosable language yaml %}}

{{% /choosable %}}
{{% choosable language java %}}

{{% /choosable %}}
{{< /chooser >}}
# Configuration options

The zenduty provider offers two means of providing credentials for authentication.

- Static credentials
- Environment variables
### Static credentials

!> **Warning:** Hard-coding credentials into any Pulumi configuration is not
recommended, and risks secret leakage should this file ever be committed to a
public version control system.

Static credentials can be provided by adding `token` in-line in the Zenduty provider configuration.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}

{{% /choosable %}}
{{% choosable language python %}}

{{% /choosable %}}
{{% choosable language csharp %}}

{{% /choosable %}}
{{% choosable language go %}}

{{% /choosable %}}
{{% choosable language yaml %}}

{{% /choosable %}}
{{% choosable language java %}}

{{% /choosable %}}
{{< /chooser >}}
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