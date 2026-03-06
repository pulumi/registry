---
title: "6.0 Migration Guide"
h1: "Migrating from 5.x to 6.x"
meta_desc: Practitioner level instructions for migrating from pulumi-aws 5.x to 6.x.
layout: package
---

## Upstream Changes

The upstream target has been changed from [v4.67.0](https://github.com/pulumi/pulumi-aws/pull/2521) to targeting v5.11.0. That means that the upstream [migration guide](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/guides/version-5-upgrade) as well as the following `CHANGELOG`s are relevant:
- [v5.0.0](https://github.com/hashicorp/terraform-provider-aws/releases/tag/v5.0.0)
- [v5.0.1](https://github.com/hashicorp/terraform-provider-aws/releases/tag/v5.0.1)
- [v5.1.0](https://github.com/hashicorp/terraform-provider-aws/releases/tag/v5.1.0)
- [v5.2.0](https://github.com/hashicorp/terraform-provider-aws/releases/tag/v5.2.0)
- [v5.3.0](https://github.com/hashicorp/terraform-provider-aws/releases/tag/v5.3.0)
- [v5.4.0](https://github.com/hashicorp/terraform-provider-aws/releases/tag/v5.4.0)
- [v5.5.0](https://github.com/hashicorp/terraform-provider-aws/releases/tag/v5.5.0)
- [v5.6.0](https://github.com/hashicorp/terraform-provider-aws/releases/tag/v5.6.0)
- [v5.6.1](https://github.com/hashicorp/terraform-provider-aws/releases/tag/v5.6.1)
- [v5.6.2](https://github.com/hashicorp/terraform-provider-aws/releases/tag/v5.6.2)
- [v5.7.0](https://github.com/hashicorp/terraform-provider-aws/releases/tag/v5.7.0)
- [v5.8.0](https://github.com/hashicorp/terraform-provider-aws/releases/tag/v5.8.0)
- [v5.9.0](https://github.com/hashicorp/terraform-provider-aws/releases/tag/v5.9.0)
- [v5.10.0](https://github.com/hashicorp/terraform-provider-aws/releases/tag/v5.10.0)
- [v5.11.0](https://github.com/hashicorp/terraform-provider-aws/releases/tag/v5.11.0)

## Deprecated resources and functions

The resources and functions listed below were renamed in a previous version. With this release, the original names have been fully deprecated. Any existing references to the original name will need to be updated.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}

{{% choosable language typescript %}}

```diff
- const test = new aws.applicationloadbalancing.TargetGroup("test", {
+ const test = new aws.lb.TargetGroup("test", {
      port: 80,
      protocol: "HTTP",
      vpcId: main.id,
});
```

{{% /choosable %}}

{{% choosable language python %}}

```diff
-  test = aws.applicationloadbalancing.TargetGroup("test",
+  test = aws.lb.TargetGroup("test",
       port=80,
       protocol="HTTP",
       vpc_id=main.id)
```

{{% /choosable %}}

{{% choosable language go %}}

```diff
- "github.com/pulumi/pulumi-aws/sdk/v5/go/aws/applicationloadbalancing"
+ "github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lb"

- _, err = applicationloadbalancing.NewTargetGroup(ctx, "test", &lb.TargetGroupArgs{
+ _, err = lb.NewTargetGroup(ctx, "test", &lb.TargetGroupArgs{
      Port:     pulumi.Int(80),
      Protocol: pulumi.String("HTTP"),
      VpcId:    main.ID(),
  })
```

{{% /choosable %}}

{{% choosable language csharp %}}

```diff
- var test = new Aws.ApplicationLoadBalancing.TargetGroup("test", new()
+ var test = new Aws.LB.TargetGroup("test", new()
  {
      Port = 80,
      Protocol = "HTTP",
      VpcId = main.Id,
  });
```

{{% /choosable %}}

{{% choosable language java %}}

```diff
- import com.pulumi.aws.applicationLoadBalancing.TargetGroup;
+ import com.pulumi.aws.lb.TargetGroup;

    var test = new TargetGroup("test", TargetGroupArgs.builder()
        .port(80)
        .protocol("HTTP")
        .vpcId(main.id())
        .build());
```

{{% /choosable %}}

{{% choosable language yaml %}}

```diff
- type: aws:applicationLoadBalancing:TargetGroup
+ type: aws:lb:TargetGroup
    properties:
      port: 80
      protocol: HTTP
      vpcId: ${main.id}
```

{{% /choosable %}}

{{< /chooser >}}

### Resources

- aws:applicationloadbalancing:Listener -> [aws:alb:Listener](https://www.pulumi.com/registry/packages/aws/api-docs/alb/listener/)
- aws:applicationloadbalancing:ListenerCertificate -> [aws:applicationloadbalancing:ListenerCertificate](https://www.pulumi.com/registry/packages/aws/api-docs/alb/listenercertificate/)
- aws:alb:ListenerRule -> [aws:alb:ListenerRule](https://www.pulumi.com/registry/packages/aws/api-docs/alb/listenerrule/)
- aws:applicationloadbalancing:LoadBalancer -> [aws:alb:LoadBalancer](https://www.pulumi.com/registry/packages/aws/api-docs/alb/loadbalancer/)
- aws:applicationloadbalancing:TargetGroup -> [aws:alb:TargetGroup](https://www.pulumi.com/registry/packages/aws/api-docs/alb/targetgroup/)
- aws:applicationloadbalancing:TargetGroupAttachment -> [aws:applicationloadbalancing:TargetGroupAttachment](https://www.pulumi.com/registry/packages/aws/api-docs/alb/targetgroupattachment/)
- aws:ec2:TransitGatewayPeeringAttachmentAccepter -> [aws:ec2transitgateway:PeeringAttachment](https://www.pulumi.com/registry/packages/aws/api-docs/ec2transitgateway/peeringattachmentaccepter/)
- aws:elasticloadbalancing:AppCookieStickinessPolicy -> [aws:elb:AppCookieStickinessPolicy](https://www.pulumi.com/registry/packages/aws/api-docs/elb/appcookiestickinesspolicy/)
- aws:elasticloadbalancing:Attachment -> [aws:elb:Attachment](https://www.pulumi.com/registry/packages/aws/api-docs/elb/attachment/)
- aws:elasticloadbalancing:ListenerPolicy -> [aws:elb:ListenerPolicy](https://www.pulumi.com/registry/packages/aws/api-docs/elb/listenerpolicy/)
- aws:elasticloadbalancing:LoadBalancer -> [aws:elb:LoadBalancer](https://www.pulumi.com/registry/packages/aws/api-docs/elb/loadbalancer/)
- aws:elasticloadbalancing:LoadBalancerBackendServerPolicy -> [aws:elb:LoadBalancerBackendServerPolicy](https://www.pulumi.com/registry/packages/aws/api-docs/elb/loadbalancerbackendserverpolicy/)
- aws:elasticloadbalancing:LoadBalancerCookieStickinessPolicy -> [aws:elb:LoadBalancerCookieStickinessPolicy](https://www.pulumi.com/registry/packages/aws/api-docs/elb/loadbalancercookiestickinesspolicy/)
- aws:elasticloadbalancing:LoadBalancerPolicy -> [aws:elb:LoadBalancerPolicy](https://www.pulumi.com/registry/packages/aws/api-docs/elb/loadbalancerpolicy/)
- aws:elasticloadbalancing:SslNegotiationPolicy -> [aws:elb:SslNegotiationPolicy](https://www.pulumi.com/registry/packages/aws/api-docs/elb/sslnegotiationpolicy/)
- aws:elasticloadbalancingv2:Listener -> [aws:lb:Listener](https://www.pulumi.com/registry/packages/aws/api-docs/lb/listener/)
- aws:elasticloadbalancingv2:ListenerCertificate -> [aws:lb:ListenerCertificate](https://www.pulumi.com/registry/packages/aws/api-docs/lb/listenercertificate/)
- aws:elasticloadbalancingv2:ListenerRule -> [aws:lb:ListenerRule](https://www.pulumi.com/registry/packages/aws/api-docs/lb/listenerrule/)
- aws:elasticloadbalancingv2:LoadBalancer -> [aws:lb:LoadBalancer](https://www.pulumi.com/registry/packages/aws/api-docs/lb/loadbalancer/)
- aws:elasticloadbalancingv2:TargetGroup -> [aws:elbv2:TargetGroup](https://www.pulumi.com/registry/packages/aws/api-docs/lb/targetgroup/)
- aws:elasticloadbalancingv2:TargetGroupAttachment -> [aws:lb:TargetGroupAttachment](https://www.pulumi.com/registry/packages/aws/api-docs/lb/targetgroupattachment/)

### Functions

- aws:applicationloadbalancing:getListener -> [aws:alb:getListener](https://www.pulumi.com/registry/packages/aws/api-docs/alb/getlistener/)
- aws:applicationloadbalancing:getLoadBalancer -> [aws:alb:getLoadBalancer](https://www.pulumi.com/registry/packages/aws/api-docs/alb/getloadbalancer/)
- aws:applicationloadbalancing:getTargetGroup -> [aws:alb:getTargetGroup](https://www.pulumi.com/registry/packages/aws/api-docs/alb/gettargetgroup/)
- aws:cloudtrail:getFunction -> [aws:cloudfront:getFunction](https://www.pulumi.com/registry/packages/aws/api-docs/cloudfront/getfunction/)
- aws:elasticloadbalancing:getHostedZoneId -> [aws:elb:getHostedZoneId](https://www.pulumi.com/registry/packages/aws/api-docs/elb/gethostedzoneid)
- aws:elasticloadbalancing:getLoadBalancer -> [- aws:elb:getLoadBalancer](https://www.pulumi.com/registry/packages/aws/api-docs/elb/getloadbalancer)
- aws:elasticloadbalancing:getServiceAccount -> [- aws:elb:getServiceAccount](https://www.pulumi.com/registry/packages/aws/api-docs/elb/getserviceaccount)
- aws:elasticloadbalancingv2:getListener -> [aws:lb:getListener](https://www.pulumi.com/registry/packages/aws/api-docs/lb/getlistener)
- aws:elasticloadbalancingv2:getLoadBalancer -> [aws:lb:getLoadBalancer](https://www.pulumi.com/registry/packages/aws/api-docs/lb/getloadbalancer)
- aws:elasticloadbalancingv2:getTargetGroup -> [aws:lb:getTargetGroup](https://www.pulumi.com/registry/packages/aws/api-docs/lb/gettargetgroup)
- aws:iam:getOpenidConnectProvider -> [aws:iam:getOpenIdConnectProvider](https://www.pulumi.com/registry/packages/aws/api-docs/iam/getopenidconnectprovider/)
- aws:index:getAmi -> [aws:ec2:getAmi](https://www.pulumi.com/registry/packages/aws/api-docs/ec2/getami/)
- aws:index:getAmiIds -> [aws:ec2:getAmiIds](https://www.pulumi.com/registry/packages/aws/api-docs/ec2/getamiids/)
- aws:index:getAutoscalingGroups -> [aws:autoscaling:getAmiIds](https://www.pulumi.com/registry/packages/aws/api-docs/autoscaling/getamiids/)
- aws:index:getCanonicalUserId -> [aws:s3:getCanonicalUserId](https://www.pulumi.com/registry/packages/aws/api-docs/s3/getcanonicaluserid/)
- aws:index:getElasticIp -> [aws:ec2:getElasticIp](https://www.pulumi.com/registry/packages/aws/api-docs/ec2/getelasticip/)
- aws:index:getPrefixList -> [aws:ec2:getPrefixList](https://www.pulumi.com/registry/packages/aws/api-docs/ec2/getprefixlist/)

## Remove the deprecated AWS SDK v2 `aws.sdk` property (TypeScript only)

The [aws.sdk](https://github.com/pulumi/pulumi-aws/pull/2584)
property previously provided a direct way to access the AWS SDK v2 from runtime code.  The AWS SDK v2 has been deprecated by AWS, and as a result this property is being removed. The AWS SDK v3 is available to import and use directly via the various `@aws-sdk/client-*` libraries. See https://github.com/pulumi/pulumi-aws/pull/2584 for more details.

```patch
import * as aws from "@pulumi/aws";
+ import * as s3sdk from "@aws-sdk/client-s3";

const bucket = new aws.s3.Bucket("my-bucket");

 bucket.onObjectCreated("bucket-callback", async (event) => {
-    const s3 = new aws.sdk.S3({});
+    const s3 = new s3sdk.S3({});
     const recordFile = "lastPutFile.json";

     const records = event.Records || [];
     for (const record of records) {
         const key = record.s3.object.key;
         if (key !== recordFile) {
             // Construct an event arguments object.
             const args = {
                 key: record.s3.object.key,
                 size: record.s3.object.size,
                 eventTime: record.eventTime,
             };

             const res = await s3.putObject({
                 Bucket: bucket.id.get(),
                 Key: recordFile,
                 Body: JSON.stringify(args),
-            }).promise();
+            }));
         }
     }
 });
```

## WafV2 is now defined recursively

The `wafv2` module has been refactored to properly define recursive types resulting in a significant decrease in SDK size.
Any references to `RuleGroupRuleStatement` or `WebAclRuleStatement` properties on a wafv2 resource, will need to be updated to use the new recursive types in the `wafv2` module.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}

{{% choosable language typescript %}}

```diff
-aws.wafv2.inputs.RuleGroupRuleStatementAndStatementStatementGeoMatchStatement
+aws.wafv2.inputs.RuleGroupRuleStatement
```

{{% /choosable %}}

{{% choosable language python %}}

```diff
-pulumi_aws.wafv2.RuleGroupRuleStatementAndStatementStatementGeoMatchStatementArgs
+pulumi_aws.wafv2.RuleGroupRuleStatementArgs
```

{{% /choosable %}}

{{% choosable language go %}}

```diff
-wafv2.RuleGroupRuleStatementAndStatementStatementGeoMatchStatementArgs
+wafv2.RuleGroupRuleStatementArgs
```

{{% /choosable %}}

{{% choosable language csharp %}}

```diff
-Pulumi.Aws.WafV2.Inputs.RuleGroupRuleStatementAndStatementStatementGeoMatchStatementArgs
+Pulumi.Aws.WafV2.Inputs.RuleGroupRuleStatementArgs
```

{{% /choosable %}}

{{% choosable language java %}}

```diff
-com.pulumi.aws.wafv2.inputs.RuleGroupRuleStatementAndStatementStatementGeoMatchStatementArgs
+com.pulumi.aws.wafv2.inputs.RuleGroupRuleStatementArgs
```

{{% /choosable %}}

{{% choosable language yaml %}}

No changes are necessary.

{{% /choosable %}}

{{< /chooser >}}

## Unused Quicksight types have been removed

Unused types from the `quicksight` module have been
[removed](https://github.com/pulumi/pulumi-aws/pull/2609/commits/15f9e5eb3e2be831ed5318dd4a4e1c5faadd7bf6).
Specifically, types that begin with `AnalysisDefinition`, `DashboardDefinition`, or
`TemplateDefinition`. Since these were purely type definitions, you can replicate them by
copying them out of the v5 SDK. We have not removed any types used by quicksight resources
or functions.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}

{{% choosable language typescript %}}

```diff
-function newDataSetReference(): aws.inputs.quicksight.TemplateSourceEntitySourceAnalysisDataSetReference { ... }
+interface TemplateSourceEntitySourceAnalysisDataSetReference {
+    dataSetArn: pulumi.Input<string>,
+    dataSetPlaceholder: pulumi.Input<string>,
+}
+function newDataSetReference(): TemplateSourceEntitySourceAnalysisDataSetReference { ... }
```

{{% /choosable %}}

{{% choosable language python %}}

```diff
-args = aws.quicksight.TemplateSourceEntitySourceAnalysisDataSetReferenceArgs(data_set_arn, data_set_placeholder)
+@pulumi.input_type
+class TemplateSourceEntitySourceAnalysisDataSetReferenceArgs:
+    def __init__(__self__, *,
+                 data_set_arn: pulumi.Input[str],
+                 data_set_placeholder: pulumi.Input[str]):
+        """
+        :param pulumi.Input[str] data_set_arn: Dataset Amazon Resource Name (ARN).
+        :param pulumi.Input[str] data_set_placeholder: Dataset placeholder.
+        """
+        pulumi.set(__self__, "data_set_arn", data_set_arn)
+        pulumi.set(__self__, "data_set_placeholder", data_set_placeholder)
+
+    @property
+    @pulumi.getter(name="dataSetArn")
+    def data_set_arn(self) -> pulumi.Input[str]:
+        """
+        Dataset Amazon Resource Name (ARN).
+        """
+        return pulumi.get(self, "data_set_arn")
+
+    @data_set_arn.setter
+    def data_set_arn(self, value: pulumi.Input[str]):
+        pulumi.set(self, "data_set_arn", value)
+
+    @property
+    @pulumi.getter(name="dataSetPlaceholder")
+    def data_set_placeholder(self) -> pulumi.Input[str]:
+        """
+        Dataset placeholder.
+        """
+        return pulumi.get(self, "data_set_placeholder")
+
+    @data_set_placeholder.setter
+    def data_set_placeholder(self, value: pulumi.Input[str]):
+        pulumi.set(self, "data_set_placeholder", value)
+
+args = TemplateSourceEntitySourceAnalysisDataSetReferenceArgs(data_set_arn, data_set_placeholder)
```

{{% /choosable %}}

{{% choosable language go %}}

```diff
-var args quicksight.TemplateSourceEntitySourceAnalysisDataSetReferenceArgs
+type TemplateSourceEntitySourceAnalysisDataSetReferenceArgs struct {
+	// Dataset Amazon Resource Name (ARN).
+	DataSetArn pulumi.StringInput `pulumi:"dataSetArn"`
+	// Dataset placeholder.
+	DataSetPlaceholder pulumi.StringInput `pulumi:"dataSetPlaceholder"`
+}
+
+var args TemplateSourceEntitySourceAnalysisDataSetReferenceArgs
```

{{% /choosable %}}

{{% choosable language csharp %}}

```diff
-var args Pulumi.Aws.Quicksight.Inputs.TemplateSourceEntitySourceAnalysisDataSetReferenceArgs;
+public sealed class TemplateSourceEntitySourceAnalysisDataSetReferenceArgs : global::Pulumi.ResourceArgs
+{
+    /// <summary>
+    /// Dataset Amazon Resource Name (ARN).
+    /// </summary>
+    [Input("dataSetArn", required: true)]
+    public Input<string> DataSetArn { get; set; } = null!;
+
+    /// <summary>
+    /// Dataset placeholder.
+    /// </summary>
+    [Input("dataSetPlaceholder", required: true)]
+    public Input<string> DataSetPlaceholder { get; set; } = null!;
+
+    public TemplateSourceEntitySourceAnalysisDataSetReferenceArgs()
+    {
+    }
+    public static new TemplateSourceEntitySourceAnalysisDataSetReferenceArgs Empty => new +TemplateSourceEntitySourceAnalysisDataSetReferenceArgs();
+}
+var args TemplateSourceEntitySourceAnalysisDataSetReferenceArgs;
```

{{% /choosable %}}

{{% choosable language java %}}

```diff
-private com.pulumi.aws.quicksight.inputs.TemplateSourceEntitySourceAnalysisDataSetReferenceArgs;
+/*
+ * Copy in this file:
+ * https://github.com/pulumi/pulumi-aws/blob/v5.42.0/sdk/java/src/main/java/com/pulumi/aws/quicksight/inputs/TemplateSourceEntitySourceAnalysisDataSetReferenceArgs.java 
+ */
+ private TemplateSourceEntitySourceAnalysisDataSetReferenceArgs;
```

{{% /choosable %}}

{{% choosable language yaml %}}

Pulumi YAML is structurally typed, so you don't need to make any changes.

{{% /choosable %}}

{{< /chooser >}}

## Property name of `organizations/getOrganizationalUnits` resource changed from `childrens` to `children`

The name of property `aws:organizations/getOrganizationalUnits:getOrganizationalUnits` has
changed from `childrens` to `children` and the name of the associated type from `children`
to `child`. The name change is because `children` is itself already plural and it's
natural singular is `child`. See https://github.com/pulumi/pulumi-aws/pull/2634 for
details.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}

{{% choosable language typescript %}}

```diff
 const units = aws.organizations.getOrganizationalUnitsOutput({
     parentId: "SomeId",
 });
-export const children = units.apply(units => units.childrens);
+export const children = units.apply(units => units.children);
```

{{% /choosable %}}

{{% choosable language python %}}

```diff
 units = aws.organizations.get_organizational_units_output(parent_id="SomeId")
-pulumi.export("children", units.childrens)
+pulumi.export("children", units.children)
```

{{% /choosable %}}

{{% choosable language go %}}

```diff
 units := originizations.GetOrganizationalUnitsOutput(&originizations.GetOrganizationalUnitsOutputArgs{
     ParentId: pulumi.String("SomeId"),
 })
-pulumi.Export("children", units.Childrens)
+pulumi.Export("children", units.Children)
```

{{% /choosable %}}

{{% choosable language csharp %}}

```diff
 var units = Aws.Organizations.GetOrganizationalUnits.Invoke(new()
 {
     ParentId = "SomeId",
 });

 return new Dictionary<string, object?>
 {
-    ["children"] = units.Apply(getOrganizationalUnitsResult => getOrganizationalUnitsResult.Childrens),
+    ["children"] = units.Apply(getOrganizationalUnitsResult => getOrganizationalUnitsResult.Children),
 }
```

{{% /choosable %}}

{{% choosable language java %}}

```diff
 OrganizationsFunctions.getOrganizationalUnits(GetOrganizationalUnitsArgs.builder()
             .parentId("SomeId")
             .build())
-            .Childrens()
+            .Children()
```

{{% /choosable %}}

{{% choosable language yaml %}}

```diff
 variables:
   units:
     fn::aws:organizations:getOrganizationalUnits:
       parentId: "SomeId"
 outputs:
-    children: ${units.childrens}
+    children: ${units.children}
```

{{% /choosable %}}

{{< /chooser >}}

## Function signature change

For three functions, the signature has changed to accommodate a new argument. The new argument is optional, so passing an empty argument block is sufficient. Any reference to these three functions will need to be updated. The three impacted functions are:

- `aws:index/getBillingServiceAccount:getBillingServiceAccount`
- `aws:index/getCallerIdentity:getCallerIdentity`
- `aws:index/getPartition:getPartition`

```diff
-`(pulumi.InvokeOptions) -> T`
+ `(Args, pulumi.InvokeOptions) -> T`
```


{{< chooser language "typescript,python,go,csharp,java,yaml" >}}

{{% choosable language typescript %}}

```diff
-const result = getBillingServiceAccount(invokeOptions);
+const result = getBillingServiceAccount({ id }, invokeOptions);

-const result = getCallerIdentity(invokeOptions);
+const result = getCallerIdentity({ id }, invokeOptions);

+const result = getPartition(invokeOptions);
+const result = getPartition({ id }, invokeOptions);
```

{{% /choosable %}}

{{% choosable language python %}}

```diff
-result = get_billing_service_account(opts)
+result = get_billing_service_account(id, opts)

-result = get_caller_identity(opts)
+result = get_caller_identity(id, opts)

+result = get_partition(opts)
+result = get_partition(id, invokeOptions)
```

{{% /choosable %}}

{{% choosable language go %}}

```diff
-result, err := aws.GetBillingServiceAccount(ctx, opts)
+result, err := aws.GetBillingServiceAccount(ctx, &aws.GetBillingServiceAccountArgs{
+    Id: id,
+}, opts)

-result, err := aws.GetCallerIdentity(ctx, opts)
+result, err := aws.GetCallerIdentity(ctx, &aws.GetCallerIdentityArgs{
+    Id: id,
+}, opts)

-result, err := aws.GetPartition(ctx, opts)
+result, err := aws.GetPartition(ctx, &aws.GetPartitionArgs{
+    Id: id,
+}, opts)
```

{{% /choosable %}}

{{% choosable language csharp %}}

```diff
-var result = Aws.GetBillingServiceAccount.Invoke(opts);
+var result = Aws.GetBillingServiceAccount.Invoke(args, opts);

-var result = Aws.GetCallerIdentity.Invoke(opts);
+var result = Aws.GetCallerIdentity.Invoke(args, opts);

-var result = Aws.GetPartition.Invoke(opts);
+var result = Aws.GetPartition.Invoke(args, opts);
```

{{% /choosable %}}

{{% choosable language java %}}

```diff
-final var result AwsFunctions.getBillingServiceAccount(opts);
+final var result AwsFunctions.getBillingServiceAccount(args, opts);

-final var result AwsFunctions.getCallerIdentity(opts);
+final var result AwsFunctions.getCallerIdentity(args, opts);

-final var result AwsFunctions.getPartition(opts);
+final var result AwsFunctions.getPartition(args, opts);
```

{{% /choosable %}}

{{% choosable language yaml %}}

You may continue to use your existing code in Pulumi YAML.

{{% /choosable %}}

{{< /chooser >}}
