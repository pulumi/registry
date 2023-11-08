---
title: "GCP 7.0 Migration Guide"
h1: "Migrating from GCP 6.x to 7.x"
meta_desc: Practitioner level instructions for migrating from pulumi-gcp 6.x to 7.x.
layout: package
---

## Upstream Changes

The upstream target has been changed from [v4.84.0](https://github.com/pulumi/pulumi-gcp/pull/1220) to targeting [v5.0.0](https://github.com/hashicorp/terraform-provider-aws/releases/tag/v5.0.0). That means that the upstream [migration guide](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/version_5_upgrade) as well as the [v.5.0.0](https://github.com/hashicorp/terraform-provider-google/releases/tag/v5.0.0) `CHANGELOG`s are relevant.

## Renamed resources and functions

To improve resource naming uniformity, we renamed the serviceAccount module to serviceaccount (all lowercase). No action is required for Node.js, Python, and .NET users. Go, Java, and YAML users will have to adjust their import paths and resource types.

{{< chooser language "go,java,yaml" >}}

{{% choosable language go %}}

```diff
- import "github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/serviceAccount"
+ import "github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/serviceaccount"
```

{{% /choosable %}}

{{% choosable language java %}}

```diff
- import com.pulumi.gcp.serviceAccount.*;
+ import com.pulumi.gcp.serviceaccount.*;
```

{{% /choosable %}}

{{% choosable language yaml %}}

```diff
- type: gcp:serviceAccount:Account
+ type: gcp:serviceaccount:Account
```

{{% /choosable %}}

{{< /chooser >}}

### Resources

- gcp:serviceAccount:Account -> [gcp:serviceaccount:Account](https://www.pulumi.com/registry/packages/gcp/api-docs/serviceaccount/account/)
- gcp:serviceAccount:IAMBinding -> [gcp:serviceaccount:IAMBinding](https://www.pulumi.com/registry/packages/gcp/api-docs/serviceaccount/iambinding/)
- gcp:serviceAccount:IAMPolicy -> [gcp:serviceaccount:IAMPolicy](https://www.pulumi.com/registry/packages/gcp/api-docs/serviceaccount/iampolicy/)
- gcp:serviceAccount:Key -> [gcp:serviceaccount:Key](https://www.pulumi.com/registry/packages/gcp/api-docs/serviceaccount/key/)

### Functions

- gcp:serviceAccount:getAccount -> [gcp:serviceaccount:getAccount](https://www.pulumi.com/registry/packages/gcp/api-docs/serviceaccount/getaccount/)
- gcp:serviceAccount:getAccountAccessToken -> [gcp:serviceaccount:getAccountAccessToken](https://www.pulumi.com/registry/packages/gcp/api-docs/serviceaccount/getaccountaccesstoken/)
- gcp:serviceAccount:getAccountIdToken -> [gcp:serviceaccount:getAccountIdToken](https://www.pulumi.com/registry/packages/gcp/api-docs/serviceaccount/getaccountidtoken/)
- gcp:serviceAccount:getAccountJwt -> [gcp:serviceaccount:getAccountJwt](https://www.pulumi.com/registry/packages/gcp/api-docs/serviceaccount/getaccountjwt/)
- gcp:serviceAccount:getAccountKey -> [gcp:serviceaccount:getAccountKey](https://www.pulumi.com/registry/packages/gcp/api-docs/serviceaccount/getaccountkey/)
- gcp:serviceAccount:getIamPolicy -> [gcp:serviceaccount:getIamPolicy](https://www.pulumi.com/registry/packages/gcp/api-docs/serviceaccount/getiampolicy/)

## Labels Rework

Upstream changes for provider, resource, and data source labels are included in this release. Global Google Cloud Platform labels configured on the provider are supported through the new `defaultLabels` field. Default labels will be applied to all resources with a `labels` field.
Resources that previously contained a single `labels`  field will now contain three fields:

- The `labels`  field is now non-authoritative and only manages the label keys defined in your configuration for the resource.
- The output-only `pulumiLabels` field merges the labels defined in the resource's labels field and the global `defaultLabels`. If the same label key exists on both the resource level and provider level, the value on the resource will override the provider-level default.
- The output-only `effectiveLabels` will list all the labels present on the resource in GCP, including labels configured through Pulumi.

Note: `defaultLabels` do not support Secrets at this time. Any `defaultLabels` values marked as `Secret` will be written to the state file in plaintext.

Due to the two new output fields, a resource that has a `labels` field may show a diff when upgrading from pulumi-gcp v6 to v7. This diff reflects adding the new resource Output fields and is safe to do.

You can find more information about this change in the [upstream documentation](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/version_5_upgrade#provider-level-labels-rework).
