---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/outscale/outscale/1.1.2/index.md
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
    outscale:region:
        value: eu-west-2
    outscale:secretKeyId:
        value: 'TODO: var.secret_key_id'
    outscale:x509CertPath:
        value: /tmp/client-certificate.pem
    outscale:x509KeyPath:
        value: /tmp/key.pem

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
    outscale:region:
        value: eu-west-2
    outscale:secretKeyId:
        value: mysecretkey

```
### Environment variables

You can provide your credentials with the `OUTSCALE_ACCESSKEYID` and `OUTSCALE_SECRETKEYID` environment variables:

Example:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:

```

Usage:

```console
$ export OUTSCALE_ACCESSKEYID="myaccesskey"
$ export OUTSCALE_SECRETKEYID="mysecretkey"
$ export OUTSCALE_REGION="cloudgouv-eu-west-1"
$ export OUTSCALE_X509CERT="~/certificate/certificate.crt"
$ export OUTSCALE_X509KEY="~/certificate/certificate.key"

$ pulumi preview
```
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

* `accessKeyId` - (Optional) The ID of the OUTSCALE access key. It must be provided, but it can also be sourced from the `OUTSCALE_ACCESSKEYID` environment variable.

* `secretKeyId` - (Optional) The OUTSCALE secret key. It must be provided, but it can also be sourced from the `OUTSCALE_SECRETKEYID` environment variable.

* `region` - (Optional) The Region that will be used as default value for all resources. It can also be sourced from the `OUTSCALE_REGION` environment variable. For more information on available Regions, see [Regions, Endpoints and Availability Zones Reference](https://docs.outscale.com/en/userguide/Regions-Endpoints-and-Availability-Zones-Reference.html).

* `endpoints` - (Optional) The shortened custom endpoint that will be used as default value for all resources. For more information on available endpoints, see [Regions, Endpoints and Availability Zones Reference](https://docs.outscale.com/en/userguide/Regions-Endpoints-and-Availability-Zones-Reference.html).

* `x509CertPath` - (Optional) The path to the x509 Client Certificate. It can also be sourced from the `OUTSCALE_X509CERT` environment variable. For more information on the use of those certificates, see [About API Access Rules](https://docs.outscale.com/en/userguide/About-API-Access-Rules.html).

* `x509KeyPath` - (Optional) The path to the private key of the x509 Client Certificate. It can also be sourced from the `OUTSCALE_X509KEY` environment variable. For more information on the use of those certificates, see [About API Access Rules](https://docs.outscale.com/en/userguide/About-API-Access-Rules.html).

* `configFile` - (Optional) The path to the OSC config file.

* `profile` - (Optional) The named profile you want to use.

* `OSC_PROFILE` - (Optional) The OSC profile name as set in the shared configuration and credentials files.

* `OSC_CONFIG_FILE` - (Optional) The path to the OSC config file.