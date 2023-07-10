---
title: "v4 Provider Migration Guide"
meta_desc: Learn how to migrate pulumi-kubernetes from v3 to v4
layout: how-to-guide
---

### Breaking changes

#### Make all resource output properties required
In previous versions of pulumi-kubernetes, resource output properties were required types in some SDKs, and optional in others. This inconsistency caused problems for multi-language features and documentation. Go, Python, and Java SDKs have been updated to represent all resource outputs as required. This matches the previous representation in the NodeJS and .NET SDKs. This does not change the resources on the cluster, but will require users to update any explicit references to the underlying types in their Pulumi code.

Examples:

{{< chooser language "python,go,java" >}}

{{% choosable language go %}}

```diff
-	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
+	Metadata metav1.ObjectMetaOutput `pulumi:"metadata"`
```

{{% /choosable %}}

{{% choosable language python %}}

```diff
-    def metadata(self) -> pulumi.Output[Optional['_meta.v1.outputs.ObjectMeta']]:
+    def metadata(self) -> pulumi.Output['_meta.v1.outputs.ObjectMeta']:
```

{{% /choosable %}}

{{% choosable language java %}}

```diff
-    private Output</* @Nullable */ ObjectMeta> metadata;
+    private Output<ObjectMeta> metadata;
-    public Output<Optional<ObjectMeta>> metadata() {
-        return Codegen.optional(this.metadata);
+    public Output<ObjectMeta> metadata() {
+        return this.metadata;
```

{{% /choosable %}}

{{< /chooser >}}

#### Drop support for Kubernetes clusters older than v1.13
pulumi-kubernetes maintains backwards compatibility with a broad range of Kubernetes versions, including those past End of Life (EOL). This update drops support for clusters older than v1.13, which was released in December 2018, and has been EOL since October 2019. The older cluster versions do not support Server-side Apply (SSA), which is now the default for pulumi-kubernetes.

#### Remove deprecated helm/v2 SDKs
The helm/v2 SDK was deprecated in v3.19.2 (May 2022), and is now removed. The helm/v3 SDK provides equivalent support, but uses the helm client library rather than shelling out to the helm CLI, which required the user to manage that dependency separately, and led to a variety of compatibility bugs. The v2:Chart and v3:Chart resources are aliased, so users can change any v2.Chart resources to v3.Chart in their Pulumi code without a disruptive change on the next update.

{{< chooser language "typescript,python,go,csharp" >}}

{{% choosable language typescript %}}

```diff
 import * as k8s from "@pulumi/kubernetes";
 
-const nginxIngress = new k8s.helm.v2.Chart("nginx-ingress", {
+const nginxIngress = new k8s.helm.v3.Chart("nginx-ingress", {
```

{{% /choosable %}}

{{% choosable language python %}}

```diff
-from pulumi_kubernetes.helm.v2 import Chart, LocalChartOpts
+from pulumi_kubernetes.helm.v3 import Chart, LocalChartOpts
```

{{% /choosable %}}

{{% choosable language go %}}

```diff
 import (
-	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/helm/v2"
+	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/helm/v3"
 )
```

{{% /choosable %}}

{{% choosable language csharp %}}

```diff
-using Pulumi.Kubernetes.Helm.V2;
+using Pulumi.Kubernetes.Helm.V3;
```

{{% /choosable %}}

{{< /chooser >}}

### Remove deprecated enableDryRun Provider option
Initial preview versions of Server-side Apply (SSA) preview were feature-gated using the `enableDryRun` provider option. This option was deprecated in v3.20.0 (July 2022), and is now removed. SSA mode can be disabled by using the `enableServerSideApply` provider option.

### Remove deprecated enableReplaceCRD Provider option
v3.12.1 (December 2021) changed the way updates were applied to CustomResourceDefinition (CRD) resources. This behavior was initially feature-gated using the `enableReplaceCRD` Provider option. This behavior became the default in v3.19.2 (May 2022), and the deprecated option is now removed.

## Notable changes

### Server-side Apply by default
Server-side Apply (SSA) mode is now default in the v4 provider. SSA mode enables several useful new features:
Kubernetes resources can be safely managed by more than one controller. Any conflicting fields will produce a detailed error message which avoids accidentally overriding changes made by other controllers.
Resources can be "upserted"; in other words, create the resource if it does not exist, or apply the specified configuration to an existing resource.
Patch resources add a new way to manage shared Kubernetes resources. Patch resources can be used to make targeted changes, and do not require the entire resource to be specified. This is a convenient way to update shared resources; a common example is resources that are provisioned by the cloud provider rather than as part of the Pulumi stack. Deleting a Patch resource undoes the patched change, but does not remove the shared Kubernetes resource.
See [the SSA how-to guide](https://www.pulumi.com/registry/packages/kubernetes/how-to-guides/managing-resources-with-server-side-apply/) for more details and usage examples for SSA mode.

{{% notes type="info" %}}
Users may opt into the previous default Client-side Apply (CSA) mode by setting the `enableServerSideApply` provider configuration to "false".
{{% /notes %}}

### Improved diff experience
The process for calculating what changes a provider needs to make to achieve the desired state is called "diffing". This is a complex process that accounts for the current and previous state of the Pulumi program (the "inputs"), the stack’s state/checkpoint, and the actual current state of any resources managed by the provider. A major challenge of this process is detecting "drift", where resources have been changed by something other than the Pulumi stack so that these changes are not reflected in Pulumi’s view of the world. The experimental Server-side Apply (SSA) support available in v3.x versions does detect resource drift, but the implementation was slow for larger stacks. The v4 provider now detects resource drift in both Client-side Apply and Server-side Apply by running a refresh operation. This brings the behavior in line with other providers, and offers consistent performance that scales well to larger stacks.

As part of this work, we removed the use of the `kubectl.kubernetes.io/last-applied-configuration` annotation, which was previously used by the provider to track input values for Kubernetes resources. This change addresses many longstanding issues caused by the use of this annotation.
