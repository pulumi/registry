# 6.x to 7.x Migration Guide

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

gcp:serviceAccount:Account -> [gcp:serviceaccount:Account](https://www.pulumi.com/registry/packages/gcp/api-docs/serviceaccount/account/)
gcp:serviceAccount:IAMBinding -> [gcp:serviceaccount:IAMBinding](https://www.pulumi.com/registry/packages/gcp/api-docs/serviceaccount/iambinding/)
gcp:serviceAccount:IAMPolicy -> [gcp:serviceaccount:IAMPolicy](https://www.pulumi.com/registry/packages/gcp/api-docs/serviceaccount/iampolicy/)
gcp:serviceAccount:Key -> [gcp:serviceaccount:Key](https://www.pulumi.com/registry/packages/gcp/api-docs/serviceaccount/key/)

### Functions

gcp:serviceAccount:getAccount -> [gcp:serviceaccount:getAccount](https://www.pulumi.com/registry/packages/gcp/api-docs/serviceaccount/getaccount/)
gcp:serviceAccount:getAccountAccessToken -> [gcp:serviceaccount:getAccountAccessToken](https://www.pulumi.com/registry/packages/gcp/api-docs/serviceaccount/getaccountaccesstoken/)
gcp:serviceAccount:getAccountIdToken -> [gcp:serviceaccount:getAccountIdToken](https://www.pulumi.com/registry/packages/gcp/api-docs/serviceaccount/getaccountidtoken/)
gcp:serviceAccount:getAccountJwt -> [gcp:serviceaccount:getAccountJwt](https://www.pulumi.com/registry/packages/gcp/api-docs/serviceaccount/getaccountjwt/)
gcp:serviceAccount:getAccountKey -> [gcp:serviceaccount:getAccountKey](https://www.pulumi.com/registry/packages/gcp/api-docs/serviceaccount/getaccountkey/)
gcp:serviceAccount:getIamPolicy -> [gcp:serviceaccount:getIamPolicy](https://www.pulumi.com/registry/packages/gcp/api-docs/serviceaccount/getiampolicy/)

## Labels Rework

Upstream changes for provider, resource, and data source labels are included in this release. You can find more information about this change in the [upstream documentation](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/version_5_upgrade#provider-level-labels-rework)

<!-- TODO -->

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}

{{% chooseable language typescript %}}

```typescript

```

{{% /choosable %}}

{{% chooseable language python %}}

```python

```

{{% /choosable %}}

{{% choosable language go %}}

```go

```

{{% /choosable %}}

{{% chooseable language chsarp %}}

```csharp

```

{{% /choosable %}}

{{% choosable language java %}}

```java

```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml

```

{{% /choosable %}}

{{< /chooser >}}
