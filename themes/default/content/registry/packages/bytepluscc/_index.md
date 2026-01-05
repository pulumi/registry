---
# WARNING: this file was fetched from https://raw.githubusercontent.com/byteplus-sdk/pulumi-bytepluscc/v0.0.7/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/byteplus-sdk/pulumi-bytepluscc/blob/v0.0.7/docs/_index.md
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Bytepluscc Provider
meta_desc: Provides an overview on how to configure the Pulumi Bytepluscc provider.
layout: package
---

## Installation

The Bytepluscc provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@byteplus/pulumi-bytepluscc`](https://www.npmjs.com/package/@byteplus/pulumi-bytepluscc)
* Python: [`pulumi-bytepluscc`](https://pypi.org/project/pulumi-bytepluscc/)
* Go: [`github.com/byteplus-sdk/pulumi-bytepluscc/sdk/go/bytepluscc`](https://github.com/byteplus-sdk/pulumi-bytepluscc)
* .NET: [`Byteplus.Pulumi.Bytepluscc`](https://www.nuget.org/packages/Byteplus.Pulumi.Bytepluscc)
* Java: [`com.byteplus/bytepluscc`](https://central.sonatype.com/artifact/com.byteplus/bytepluscc)

## Overview

The Byteplus Cloud Control Provider enables interaction with various Byteplus-supported resources through the Cloud Control API. Prior to usage, you must configure the provider with appropriate credentials.

## Authentication

The bytepluscc provider accepts several ways to enter credentials for authentication.
The following methods are supported, in this order, and explained below:

- Static credentials
- Environment variables

### Static credentials

Static credentials can be provided by adding `accessKey`, `secretKey` and `region` in-line in the
bytepluscc provider configuration:

Usage:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    byteplus:accessKey:
        value: 'TODO: var.access_key'
    byteplus:region:
        value: 'TODO: var.region'
    byteplus:secretKey:
        value: 'TODO: var.secret_key'

```

### Environment variables

You can provide your credentials via `BYTEPLUS_ACCESS_KEY`, `BYTEPLUS_SECRET_KEY` environment variables. The Region can be set using the `BYTEPLUS_REGION` environment variables.

Usage:
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:

```

```shell
$ export BYTEPLUS_ACCESS_KEY="<Your-Access-Key-ID>"
$ export BYTEPLUS_SECRET_KEY="<Your-Access-Key-Secret>"
$ export BYTEPLUS_REGION="cn-beijing"
```

## Configuration Reference

In addition to generic `provider` arguments
(e.g. `alias` and `version`), the following arguments are supported in the Bytepluscc
provider configuration:

### Optional

- `accessKey` (String) The Access Key for Byteplus Provider. It must be provided, but it can also be sourced from the `BYTEPLUS_ACCESS_KEY` environment variable
- `secretKey` (String) he Secret Key for Byteplus Provider. It must be provided, but it can also be sourced from the `BYTEPLUS_SECRET_KEY` environment variable
- `assumeRole` (Attributes) An `assume_role` block (documented below). Only one `assume_role` block may be in the configuration. (see [below for nested schema](#nestedatt--assume_role))
- `customerHeaders` (String) CUSTOMER HEADERS for Byteplus Provider. The customer_headers field uses commas (,) to separate multiple headers, and colons (:) to separate each header key from its corresponding value.
- `disableSsl` (Boolean) Disable SSL for Byteplus Provider
- `endpoints` (Attributes) An `endpoints` block (documented below). Only one `endpoints` block may be in the configuration. (see [below for nested schema](#nestedatt--endpoints))
- `proxyUrl` (String) PROXY URL for Byteplus Provider
- `region` (String) The Region for Byteplus Provider. It must be provided, but it can also be sourced from the `BYTEPLUS_REGION` environment variable


<a id="nestedatt--assume_role"></a>

### Nested Schema for `assume_role`

Required:

- `assumeRoleTrn` (String) he TRN of the role to assume.
- `durationSeconds` (Number) The duration of the session when making the AssumeRole call. Its value ranges from 900 to 43200(seconds), and default is 3600 seconds.
  Optional:

- `policy` (String) A more restrictive policy when making the AssumeRole call

<a id="nestedatt--endpoints"></a>

### Nested Schema for `endpoints`

Optional:

- `cloudcontrolapi` (String) Use this to override the default Cloud Control API service endpoint URL
- `sts` (String) Use this to override the default STS service endpoint URL