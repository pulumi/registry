---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/paloaltonetworks/prismacloud/1.7.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Prismacloud Provider
meta_desc: Provides an overview on how to configure the Pulumi Prismacloud provider.
layout: package
---

## Generate Provider

The Prismacloud provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider paloaltonetworks/prismacloud
```
## Overview

This provider is for the Palo Alto Networks Prisma Cloud platform.
## Example Provider Usage

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    prismacloud:jsonConfigFile:
        value: .prismacloud_auth.json

```
## Configuration Reference

There are multiple ways to specify provider config, and they may all be combined if desired.  The params are taken from the following locations, in order of preference:

1) Any param specified explicitly in the provider configuration
2) From the param's environment variable, where applicable.
3) From the JSON config file, if specified.

The following configuration inputs are supported:

* `url` - (Env: `PRISMACLOUD_URL`) The API URL without the leading protocol.
* `username` - (Env: `PRISMACLOUD_USERNAME`) Access key ID.
* `password` - (Env: `PRISMACLOUD_PASSWORD`) Secret key.
* `customerName` - (Env: `PRISMACLOUD_CUSTOMER_NAME`) Customer name.
* `protocol` - (Env: `PRISMACLOUD_PROTOCOL`) The protocol.  Valid values are `https` or `http`.
* `port` - (Env: `PRISMACLOUD_PORT`, int) If the port is non-standard for the protocol, the port number to use.
* `timeout` - The default timeout (in seconds) for all communications with Prisma Cloud (default: `180`).
* `skipSslCertVerification` - (Env: `PRISMACLOUD_SKIP_SSL_CERT_VERIFICATION`, bool) Skip SSL certificate verification.
* `logging` - Map of logging options for the API connection.  Valid values are `quiet` (disable logging), `action`, `path`, `send`, and `receive`.
* `disableReconnect` - (bool) Prisma Cloud invalidates authenticated sessions after 10minutes.  By default the provider will silently get a new JSON web token and continue deploying the plan.  If you do not want the provider to fetch a new JSON web token, set this to `true`.
* `jsonWebToken` - (Env: `PRISMACLOUD_JSON_WEB_TOKEN`) A JSON web token.  These are only valid for 10 minutes once issued.  If this is specified but not the `username` / `password` then the provider will not have a way to reauthenticate once the JSON web token expires.
* `jsonConfigFile` - (Env: `PRISMACLOUD_JSON_CONFIG_FILE`) Retrieve the provider configuration from this JSON file.  When retrieving params from the JSON configuration file, the param names are the same as the provider params, except that underscores in provider params become hyphens in the JSON config file.  For example, the provider param `jsonWebToken` is `json-web-token` in the config file.
* `maxRetries` - (Optional) Maximum number of times an API call is retried when requests are throttled (default: `5`).
* `retryMaxDelay` - (Optional) Maximum time the API calls are retried when creating or updating resources (default: `30`).
* `retryType` - (Optional) Specifies the type of backoff strategy for handling retries, allowing users to customize the delay between retry attempts. Valid values are `exponentialBackoff` and `linearBackoff` (default: `exponentialBackoff`).
## Support

This template/solution are released under an as-is, best effort, support
policy. These scripts should be seen as community supported and Palo Alto
Networks will contribute our expertise as and when possible. We do not
provide technical support or help in using or troubleshooting the components
of the project through our normal support options such as Palo Alto Networks
support teams, or ASC (Authorized Support Centers) partners and backline
support options. The underlying product used (the VM-Series firewall) by the
scripts or templates are still supported, but the support is only for the
product functionality and not for help in deploying or using the template or
script itself. Unless explicitly tagged, all projects or work posted in our
GitHub repository (at <https://github.com/PaloAltoNetworks>) or sites other
than our official Downloads page on <https://support.paloaltonetworks.com>
are provided under the best effort policy.