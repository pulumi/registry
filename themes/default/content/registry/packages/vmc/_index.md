---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/vmware/vmc/1.15.7/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Vmc Provider
meta_desc: Provides an overview on how to configure the Pulumi Vmc provider.
layout: package
---

## Generate Provider

The Vmc provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider vmware/vmc
```
## Overview

<img src="https://raw.githubusercontent.com/vmware/pulumi-provider-vmc/main/docs/images/icon-color.png" alt="VMware Cloud on AWS" width="150">
The Pulumi Provider for [VMware Cloud on AWS](https://techdocs.broadcom.com/us/en/vmware-cis/cloud/vmware-cloud-on-aws/SaaS.html) is a plugin for Pulumi that allows you to
interact with VMware Cloud on AWS.
##

In order to use the provider you need to obtain the authentication token from the Cloud Service Provider by providing
the org scoped API token. The provider client uses Cloud Service Provider (CSP) API to exchange this org scoped API
token/OAuth App ID and Secret for user access token.

Refer to the provider documentation for information on all of the resources
and functions supported by this provider. Each includes a detailed
description of the purpose and how to use it.
## Configuration Reference

The following arguments are used to configure the provider:

* `apiToken` - (Required, in conflict with `clientId` and `clientSecret`)
  API token is used to authenticate when calling VMware Cloud Services APIs.
  This token is scoped within the organization.
* `clientId` - (Required in pair with `clientSecret`, in conflict with `apiToken`)
  ID of OAuth App associated with the organization. The combination with
  "clientSecret" is used to authenticate when calling VMware Cloud Services
  APIs.
* `clientSecret` - (Required in pair with `clientId`, in conflict with `apiToken`)
  Secret of OAuth App associated with the organization. The combination with
  "clientId" is used to authenticate when calling VMware Cloud Services APIs.
* `orgId` - (Required) Organization Identifier.
* `vmcUrl` - (Optional) VMware Cloud on AWS URL. Default: <https://vmc.vmware.com>
* `cspUrl` - (Optional) Cloud Service Provider URL. Default: <https://console.cloud.vmware.com>