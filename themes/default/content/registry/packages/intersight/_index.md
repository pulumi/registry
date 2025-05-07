---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/ciscodevnet/intersight/1.0.65/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Intersight Provider
meta_desc: Provides an overview on how to configure the Pulumi Intersight provider.
layout: package
---

## Generate Provider

The Intersight provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider ciscodevnet/intersight
```
## Overview

The Cisco Intersight Pulumi Provider is used to interact with the
many resources supported by Cisco Intersight. You must have a valid Intersight account in order to use the provider.

Use the navigation to the left to read about the available resources.
## Installing Pulumi
* Follow the steps here for installing pulumi.
## Using the Cisco Intersight Pulumi Provider
### Prerequisites
The system must have:
* Pulumi
* Cisco Intersight Pulumi Provider
* An active Cisco Intersight Account.

A `apikey`, `secretkey` and Intersight `endpoint` URL are required to start using the provider.
The following code must be included in a `.tf` file in the working directory, to establish connection between
the provider and your Intersight account.
```hcl-pulumi
provider "intersight" {
  apikey    = "<api key>"
  secretkey = "<secret key as a string or in a file>"
  endpoint = "https://intersight.com"
}
```
### Writing a managed source / resource
The resource objects’ name are in the format `intersight_<model_name_in_snake_case>`. E.g., `intersight.SolPolicy`
is the resource object for **SOL Policy**. Each resource is assigned a name which can be later used for
tracking and referencing.The following is an example of a resource for server profile and NTP policy attached to the server profile:
```hcl-pulumi
resource "intersight_server_profile" "server1" {
  name = "demo_server_profile"
  organization {
    moid   = "changeMe"
    object_type = "organization.Organization"
  }
  policy_bucket {
    object_type = "ntp.Policy"
    moid = intersight_ntp_policy.ntp1.moid
  }
}

resource "intersight_ntp_policy" "ntp1" {
  name    = "demo_ntp"
  enabled = true
  ntp_servers = [
    "10.10.10.10",
    "10.10.10.11",
    "10.10.10.12",
    "10.10.10.13"
  ]
  organization {
    moid   = "changeMe"
    object_type = "organization.Organization"
  }
}
```
The first resource is for creation of `server profile`. It is named `server1`. This name will not be reflected anywhere
in the Cisco Intersight. It is for reference among the .tf files. The server profile is attached to the NTP policy created
later. This is done by referencing to the `ntp1` policy in `policyBucket`. A resource can point or reference to
another resource in the format `<resource>.<resource_name>.<property_name>` .
### Referencing Intersight MOs using function
The function objects’ name are in the format `intersight_<model_name_in_snake_case>`. E.g., `intersight.SolPolicy`
is the function object for SOL Policy. Each function is assigned a name which can be later used for tracking and
referencing.
The following is an example of accessing a pre-existing server profile and attaching a NTP policy to the server profile.
```hcl-pulumi
resource "intersight_server_profile" "server1" {
  name = "demo_server_profile"
  organization {
    moid   = "changeMe"
    object_type = "organization.Organization"
  }
  policy_bucket {
    object_type = "ntp.Policy"
    moid = data.intersight_ntp_policy.ntp1.results.0.moid
  }
}

data "intersight_ntp_policy" "ntp1" {
  name    = "demo_ntp"
}
```
An NTP policy with the given constraints i.e. name as `demoNtp` must exist in the Intersight account. Data source
is used to access this object. The function is named `ntp1`. The NTP policy is attached to the server profile with the name
`server1`. Data sources are referenced in the format `data.<data_source>.<data_source_name>.results.<index>.<property_name>`.