---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/datadrivers/nexus/2.5.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Nexus Provider
meta_desc: Provides an overview on how to configure the Pulumi Nexus provider.
layout: package
---

## Generate Provider

The Nexus provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider datadrivers/nexus
```
## Overview

The Nexus provider allows Pulumi to read from, write to, and configure [Sonatype Nexus Repository Manager](https://www.sonatype.com/product-nexus-repository).

> **Note** This provider hat been implemented and tested with Sonatype Nexus Repository Manager OSS `3.70.1-02`.
## Usage
### Provider config

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    nexus:insecure:
        value: true
    nexus:password:
        value: admin123
    nexus:url:
        value: https://127.0.0.1:8080
    nexus:username:
        value: admin

```
## Configuration Reference

- `clientCertPath` (String) Path to a client PEM certificate to load for mTLS. Reading environment variable NEXUS_CLIENT_CERT_PATH. Default:``
- `clientKeyPath` (String) Path to a client PEM key to load for mTLS. Reading environment variable NEXUS_CLIENT_KEY_PATH. Default:``
- `insecure` (Boolean) Boolean to specify wether insecure SSL connections are allowed or not. Reading environment variable NEXUS_INSECURE_SKIP_VERIFY. Default:`true`
- `password` (String) Password of user to connect to API. Reading environment variable NEXUS_PASSWORD. Default:`admin123`
- `rootCaPath` (String) Path to a root CA certificate to load for mTLS. Reading environment variable NEXUS_ROOT_CA_PATH. Default:``
- `timeout` (Number) Timeout in seconds to connect to API. Reading environment variable NEXUS_TIMEOUT. Default:`30`
- `url` (String) URL of Nexus to reach API. Reading environment variable NEXUS_URL. Default:`http://127.0.0.1:8080`
- `username` (String) Username used to connect to API. Reading environment variable NEXUS_USERNAME. Default:`admin`
## Author

[Datadrivers GmbH](https://www.datadrivers.de)