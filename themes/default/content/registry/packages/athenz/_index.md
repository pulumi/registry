---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/athenz/athenz/1.0.45/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Athenz Provider
meta_desc: Provides an overview on how to configure the Pulumi Athenz provider.
layout: package
---

## Generate Provider

The Athenz provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider athenz/athenz
```
## Background

As a part of simplifying and improving Athenz user experience, we would like to offer our customers Infrastructure As
Code capability.

Pulumi relies on plugins called "providers" to interact with cloud providers, SaaS providers, and other APIs.

Each provider adds a set of resource types and/or functions that Pulumi can manage. Every resource type is
implemented by a provider; without providers, Pulumi can't manage any kind of infrastructure.
## Concepts

Each provider adds a set of resource types and/or functions that Pulumi can manage.
Every resource type is implemented by a provider; without providers, Pulumi can't manage any kind of infrastructure.
## Data source

Data sources allow Pulumi use information defined outside of Pulumi, defined by another separate Pulumi
configuration, or modified by functions.
## Resources

Resources are the most important element in the Pulumi language. Each resource block describes one or more
infrastructure objects.
## Example Usage

Do not keep your authentication password in HCL for production environments, use Pulumi environment variables.

To install this provider, copy and paste this code into your Pulumi configuration. Make sure to update the version
parameter to appropriate value. Then, run `pulumi up`.

You can also obtain this snippet from
the provider home page by clicking the "USE PROVIDER"
button on top right hand side.
## Configuration Reference
### Required

- `zmsUrl` (String) Athenz API URL

- `key` (String) Athenz client private key
- `cert` (String) Athenz client x.509 certificate
- `cacert` (String) CA Certificate file path
- `disableResourceOwnership` (Bool) Disable resource ownership. Default is false
- `resourceOwner` (String) Resource owner. Default is "TF"
- `roleMetaResourceState` (Number) Bitmask of object state flags controlling role behavior when creating or destroying roleMeta resources. 0x01: create the role if not already present, 0x02: always delete the role when destroying the resource. Default value is 1. The value is used when the resourceState attribute at the athenz.RoleMeta level is set to -1
- `groupMetaResourceState` (Number) Bitmask of object state flags controlling group behavior when creating or destroying groupMeta resources. 0x01: create the group if not already present, 0x02: always delete the group when destroying the resource. Default value is 1. The value is used when the resourceState attribute at the athenz.GroupMeta level is set to -1