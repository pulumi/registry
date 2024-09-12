---
title: "GCP 8.0 Migration Guide"
h1: "Migrating from GCP 7.x to 8.x"
meta_desc: Practitioner level instructions for migrating from pulumi-gcp 7.x to 8.x.
layout: package
---

## Upstream Changes
The v8 release of the pulumi-gcp provider is targeting upstream [v6.0.1](https://github.com/hashicorp/terraform-provider-google/releases/tag/v6.1.0). 
Please refer to the [upstream migration guide](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/version_6_upgrade) as well as the [v.6.0.0 CHANGELOG](https://github.com/hashicorp/terraform-provider-google/releases/tag/v6.0.0).

Of note: Multiple resources now have a new field `deletionProtection` which is set to true by default and will prevent these resources from getting deleted with `pulumi destroy`. When upgrading, it is recommended you add this field and set it to the desired behavior.

## Breaking Changes

### cloudrunv2.Service Ports

The following change is additionally breaking in Pulumi:

The `template.containers[].ports` field in `cloudrunv2.Service` has been changed from a list item to an object item. 
You may need to make a small adjustment to your code to reflect this change.

### Default label

This version of the provider adds a new global default label, `goog-pulumi-provisioned`, to each resource. 
It will be written to the `pulumiLabels` and `effectiveLabels` output fields.
Note that this label is not managed in the resource’s `labels` field. 
It is set by the provider, to help differentiate pulumi-provisioned resources from others in the GCP console.
To opt out of this default behavior, a new config value on the provider is available, called `addPulumiAttributionLabel`. 

Set this to false as follows:

`pulumi config set gcp:addPulumiAttributionLabel false`
