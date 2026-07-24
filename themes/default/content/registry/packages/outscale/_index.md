---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/outscale/outscale/1.7.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Outscale Provider
meta_desc: Provides an overview on how to configure the Pulumi Outscale provider.
layout: package
---

## Generate Provider

The Outscale provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider outscale/outscale
```
## Overview

The OUTSCALE provider is used to manage OUTSCALE Cloud resources.Use the navigation to the left to read about the available resources. For more information on our resources, see the [User Guide](https://docs.outscale.com/en/userguide/Home.html).

The provider is based on our OUTSCALE API. For more information, see [APIs Reference](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) and the [API Documentation](https://docs.outscale.com/api).

The provider needs to be configured with the proper credentials before it can be used.

> **Note:**
To configure the provider, see our README and the Pulumi documentation. <br />
To configure a proxy, see our README.
## Example

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    outscale:accessKeyId:
        value: 'TODO: var.access_key_id'
    outscale:secretKeyId:
        value: 'TODO: var.secret_key_id'

```
## Authentication

OUTSCALE authentication is based on access keys composed of an **access key ID** and a **secret key**.
For more information on access keys, see [About Access Keys](https://docs.outscale.com/en/userguide/About-Access-Keys.html).
To retrieve your access keys, see [Getting Information About Your Access Keys](https://docs.outscale.com/en/userguide/Getting-Information-About-Your-Access-Keys.html).

The OUTSCALE provider supports different ways of providing credentials for authentication. The following methods are supported:

1. Static credentials
2. Environment variables
### Static credentials

!> Hard-coding credentials into any Pulumi configuration is not recommended, and risks secret leakage should this file ever be committed to a public version control system.

You can provide your credentials by specifying the `accessKeyId` and `secretKeyId` attributes in the provider configuration:

Example:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    outscale:accessKeyId:
        value: myaccesskey
    outscale:secretKeyId:
        value: mysecretkey

```
### Environment variables

You can provide your credentials with the `OSC_ACCESS_KEY` and `OSC_SECRET_KEY` environment variables.

Starting from version `v1.6.0`, the legacy `OUTSCALE_*` environment variables have been replaced by `OSC_*` variables.

Example:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:

```

Usage:

```console
$ export OSC_ACCESS_KEY="myaccesskey"
$ export OSC_SECRET_KEY="mysecretkey"
$ export OSC_REGION="cloudgouv-eu-west-1"
$ export OSC_X509_CLIENT_CERT="~/certificate/certificate.crt"
$ export OSC_X509_CLIENT_KEY="~/certificate/certificate.key"

$ pulumi preview
```

Migration from the legacy provider environment variables to the SDK environment variables:

|  Legacy variable supported up  | Variable to use in `v1.6.0+` |   Provider configuration   |

|      to `v1.5.0`       |                        |                            |
|------------------------|------------------------|----------------------------|
| `OUTSCALE_ACCESSKEYID` | `OSC_ACCESS_KEY`       | `accessKeyId`            |
| `OUTSCALE_SECRETKEYID` | `OSC_SECRET_KEY`       | `secretKeyId`            |
| `OUTSCALE_REGION`      | `OSC_REGION`           | `api.region`, `oks.region` |
| `OUTSCALE_X509CERT`    | `OSC_X509_CLIENT_CERT` | `api.x509_cert_path`       |
| `OUTSCALE_X509KEY`     | `OSC_X509_CLIENT_KEY`  | `api.x509_key_path`        |
| `OUTSCALE_OAPI_URL`    | `OSC_ENDPOINT_API`     | `api.endpoint`             |
## Configuration
### Set a profile using a configuration file

You can set a named profile by specifying the `profile` attribute in the provider configuration.

The location of the shared configuration and credential file can be specified using the `configFile` attribute:

Example:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    outscale:configFile:
        value: ./.osc/config.json
    outscale:profile:
        value: default

```
### Set a profile using environment variables

You can also set a named profile by specifying the `OSC_PROFILE` environment variable.

The locations of the shared configuration and credential file can be specified using the `OSC_CONFIG_FILE` environment variable:
## Configuration Reference Reference

In addition to generic provider arguments, the following arguments are supported in the OUTSCALE provider configuration:

* `configFile` - (Optional) The path to an OSC config file. It can also be sourced from the `OSC_CONFIG_FILE` environment variable.
* `profile` - (Optional) The named profile you want to use in the OSC config file. It can also be sourced from the `OSC_PROFILE` environment variable.
* `accessKeyId` - (Optional) The ID of the OUTSCALE access key. It must be provided, but it can also be sourced from the `OSC_ACCESS_KEY` environment variable.
* `secretKeyId` - (Optional) The OUTSCALE secret key. It must be provided, but it can also be sourced from the `OSC_SECRET_KEY` environment variable.
* `api` - (Optional) Configuration elements for OUTSCALE API operations.
  * `endpoint` - (Optional) The server endpoint to use for OUTSCALE API operations. The underlying HTTP client expects a server URL with scheme, and it may include a path prefix. Example: `https://api.eu-west-2.outscale.com/api/v1`. For backward compatibility, the provider also accepts `api.<region>.outscale.com` and `https://api.<region>.outscale.com` and normalizes them to the default `/api/v1` endpoint. It can also be sourced from the `OSC_ENDPOINT_API` environment variable. For more information on available endpoints, see [API Endpoints Reference > OUTSCALE API](https://docs.outscale.com/en/userguide/API-Endpoints-Reference.html#_outscale_api).
  * `region` - (Optional) The Region to use for OUTSCALE API operations. It can also be sourced from the `OSC_REGION` environment variable. For more information on available Regions, see [About Regions and Subregions](https://docs.outscale.com/en/userguide/About-Regions-and-Subregions.html).
  * `x509CertPath` - (Optional) The path to the x509 Client Certificate. It can also be sourced from the `OSC_X509_CLIENT_CERT` environment variable. For more information on the use of those certificates, see [About API Access Rules](https://docs.outscale.com/en/userguide/About-API-Access-Rules.html).
  * `x509KeyPath` - (Optional) The path to the private key of the x509 Client Certificate. It can also be sourced from the `OSC_X509_CLIENT_KEY` environment variable. For more information on the use of those certificates, see [About API Access Rules](https://docs.outscale.com/en/userguide/About-API-Access-Rules.html).
  * `insecure` - (Optional) Enables TLS insecure connection.
* `oks` - (Optional) Configuration elements for OKS API operations.
  * `endpoint` - (Optional) The server endpoint to use for OKS API operations. The underlying HTTP client expects a server URL with scheme, and it may include a path prefix. Example: `https://api.eu-west-2.oks.outscale.com/api/v2`. It can also be sourced from the `OSC_ENDPOINT_OKS` environment variable. For more information on available endpoints, see [API Endpoints Reference > OUTSCALE Kubernetes as a Service (OKS)](https://docs.outscale.com/en/userguide/API-Endpoints-Reference.html#_outscale_kubernetes_as_a_service_oks).
  * `region` - (Optional) The Region to use for OKS API operations. It can also be sourced from the `OSC_REGION` environment variable. For more information on available Regions, see [About Regions and Subregions](https://docs.outscale.com/en/userguide/About-Regions-and-Subregions.html).

The following top-level arguments are deprecated but still supported as a fallback:
* `endpoints` - (Optional, deprecated) The endpoints to use for OUTSCALE API and OKS API operations. For more information on available endpoints, see [Regions, Endpoints and Availability Zones Reference](https://docs.outscale.com/en/userguide/Regions-Endpoints-and-Availability-Zones-Reference.html).
  * `api` - (Optional, deprecated) For OUTSCALE API.
  * `oks` - (Optional, deprecated) For OKS API.
* `region` - (Optional, deprecated) The Region to use for OUTSCALE API and OKS API operations. It can also be sourced from the `OSC_REGION` environment variable. For more information on available Regions, see [About Regions and Subregions](https://docs.outscale.com/en/userguide/About-Regions-and-Subregions.html).
* `x509CertPath` - (Optional, deprecated) The path to the x509 Client Certificate. It can also be sourced from the `OSC_X509_CLIENT_CERT` environment variable. For more information on the use of those certificates, see [About API Access Rules](https://docs.outscale.com/en/userguide/About-API-Access-Rules.html).
* `x509KeyPath` - (Optional, deprecated) The path to the private key of the x509 Client Certificate. It can also be sourced from the `OSC_X509_CLIENT_KEY` environment variable. For more information on the use of those certificates, see [About API Access Rules](https://docs.outscale.com/en/userguide/About-API-Access-Rules.html).
* `insecure` - (Optional, deprecated) Enables TLS insecure connection.