---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/vmware/vra/0.13.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Vra Provider
meta_desc: Provides an overview on how to configure the Pulumi Vra provider.
layout: package
---

## Generate Provider

The Vra provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider vmware/vra
```
## Overview

Use this Pulumi provider to interact with resources supported by [VMware Aria Automation](https://www.vmware.com/products/aria-automation.html) services, enabling you to deliver a self-service cloud consumption experience with VMware Cloud Foundation.

Please use the navigation to the left to read about available functions and resources.
## Basic Configuration of the Provider

With Pulumi 0.13 and later, the `pulumi` configuration block should be used in your configurations.

**Example**: Pulumi Configuration

In order to use the provider you must configure the provider to communicate with the VMware Aria Automation endpoint. The provider configuration requires the `url` and `refreshToken` or `accessToken`.

The provider also can accept both signed and self-signed server certificates. It is recommended that in production environments you only use certificates signed by a certificate authority. Setting the `insecure` parameter to `true` will direct the Pulumi to skip certificate verification. This is **not recommended** in production deployments. It is recommended that you use a trusted connection using certificates signed by a certificate authority.

**Example**: Configuration with Credentials

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    vra:insecure:
        value: false
    vra:refreshToken:
        value: 'TODO: var.vra_refresh_token'
    vra:url:
        value: 'TODO: var.vra_url'

```

**Example**: Setting Environment Variables

```shell
export VRA_URL="https://cloud.example.com"
export VRA_REFRESH_TOKEN="***********************"
```

Documentation about the provider resources and functions can be found within the sidebar, which has examples specific to their use. Additional examples on the use of the provider configuration, resources, and functions can be found in the `examples` directory of the [project](https://github.com/vmware/pulumi-provider-vra).

Note that in all of the examples you will need to update attributes - such as `url`, `refreshToken` or `accessToken`, and `insecure` - to match your environment.
## Configuration Reference

The following arguments are used to configure the Pulumi Provider for VMware Aria Automation:

- `url` - (Required) This is the URL to the VMware Aria Automation endpoint. Can also be specified with the `VRA_URL` environment variable.
- `organization` - (Optional) The name of the organization. Required when using VCF Automation, otherwise, this parameter is ignored. Can also be specified with the `VCFA_ORGANIZATION` environment variable.
- `accessToken` - (Optional) This is the access token used to create an API refresh token. Can also be specified with the `VRA_ACCESS_TOKEN` environment variable.
- `refreshToken` - (Optional) This is a refresh token used for API access that has been pre-generated. One of `accessToken` or `refreshToken` is required. Can also be specified with the `VRA_REFRESH_TOKEN` environment variable.
- `insecure` - (Optional) This specifies whether if the TLS certificates are validated. Can also be specified with the `VRA_INSECURE` environment variable.
- `reauthorizeTimeout` - (Optional) This specifies the timeout for how often to reauthorize the access token. Can also be specified with the `VRA_REAUTHORIZE_TIMEOUT` environment variable.
- `apiTimeout` - (Optional) This specifies the timeout in seconds for API operations. Can also be specified with the `VRA_API_TIMEOUT` environment variable.