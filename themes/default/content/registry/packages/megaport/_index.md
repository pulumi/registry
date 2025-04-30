---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/megaport/megaport/1.3.6/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Megaport Provider
meta_desc: Provides an overview on how to configure the Pulumi Megaport provider.
layout: package
---

## Generate Provider

The Megaport provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider megaport/megaport
```
## Overview

The `pulumi-provider-megaport` or Megaport Pulumi Provider lets you create and manage
Megaport's product and services using the [Megaport API](https://dev.megaport.com).

This provides an opportunity for true multi-cloud hybrid environments supported by Megaport's Software
Defined Network (SDN). Using the Pulumi provider, you can create and manage Ports,
Virtual Cross Connects (VXCs), Megaport Cloud Routers (MCRs), Megaport Virtual Edges (MVEs), and Partner VXCs.

The Megaport Pulumi Provider is released as a tool for use with the Megaport API.

**Important:** The usage of the Megaport Pulumi Provider constitutes your acceptance of the terms available
in the Megaport [Acceptable Use Policy](https://www.megaport.com/legal/acceptable-use-policy/) and
[Global Services Agreement](https://www.megaport.com/legal/global-services-agreement/).
## Configuring the provider

First you'll need to create an API key, docs for which can be found [here](https://docs.megaport.com/api/api-key/).

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    megaport:acceptPurchaseTerms:
        value: true
    megaport:accessKey:
        value: access_key
    megaport:environment:
        value: staging
    megaport:secretKey:
        value: secret_Key
    megaport:waitTime:
        value: 20

```