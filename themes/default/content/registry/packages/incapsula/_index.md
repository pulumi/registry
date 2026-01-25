---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/imperva/incapsula/3.37.1/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Incapsula Provider
meta_desc: Provides an overview on how to configure the Pulumi Incapsula provider.
layout: package
---

## Generate Provider

The Incapsula provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider imperva/incapsula
```
## Overview

The Incapsula provider is used to interact with resources supported by Imperva. The provider needs to be configured with the proper credentials before it can be used.

The current API that the Incapsula provider is calling requires sequential execution. You can either use `dependsOn` or specify the `parallelism` flag. Imperva recommends the latter and setting the value to `1`. Example call: `pulumi up -parallelism=1`.

Use the navigation to the left to read about the available resources.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    incapsula:apiId:
        value: 'TODO: "${var.incapsula_api_id}"'
    incapsula:apiKey:
        value: 'TODO: "${var.incapsula_api_key}"'

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as incapsula from "@pulumi/incapsula";

// Create a site
const example_site = new incapsula.Site("example-site", {domain: "examplesite.com"});
// Create a ACL security rule
const example_global_blacklist_ip_rule = new incapsula.index.AclSecurityRule("example-global-blacklist-ip-rule", {
    ruleId: "api.acl.blacklisted_ips",
    siteId: example_site.siteId,
    ips: "192.168.1.1,192.168.1.2",
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    incapsula:apiId:
        value: 'TODO: "${var.incapsula_api_id}"'
    incapsula:apiKey:
        value: 'TODO: "${var.incapsula_api_key}"'

```
```python
import pulumi
import pulumi_incapsula as incapsula

# Create a site
example_site = incapsula.Site("example-site", domain="examplesite.com")
# Create a ACL security rule
example_global_blacklist_ip_rule = incapsula.index.AclSecurityRule("example-global-blacklist-ip-rule",
    rule_id=api.acl.blacklisted_ips,
    site_id=example_site.site_id,
    ips=192.168.1.1,192.168.1.2)
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    incapsula:apiId:
        value: 'TODO: "${var.incapsula_api_id}"'
    incapsula:apiKey:
        value: 'TODO: "${var.incapsula_api_key}"'

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Incapsula = Pulumi.Incapsula;

return await Deployment.RunAsync(() =>
{
    // Create a site
    var example_site = new Incapsula.Site("example-site", new()
    {
        Domain = "examplesite.com",
    });

    // Create a ACL security rule
    var example_global_blacklist_ip_rule = new Incapsula.Index.AclSecurityRule("example-global-blacklist-ip-rule", new()
    {
        RuleId = "api.acl.blacklisted_ips",
        SiteId = example_site.SiteId,
        Ips = "192.168.1.1,192.168.1.2",
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    incapsula:apiId:
        value: 'TODO: "${var.incapsula_api_id}"'
    incapsula:apiKey:
        value: 'TODO: "${var.incapsula_api_key}"'

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/incapsula/v3/incapsula"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a site
		example_site, err := incapsula.NewSite(ctx, "example-site", &incapsula.SiteArgs{
			Domain: pulumi.String("examplesite.com"),
		})
		if err != nil {
			return err
		}
		// Create a ACL security rule
		_, err = incapsula.NewAclSecurityRule(ctx, "example-global-blacklist-ip-rule", &incapsula.AclSecurityRuleArgs{
			RuleId: "api.acl.blacklisted_ips",
			SiteId: example_site.SiteId,
			Ips:    "192.168.1.1,192.168.1.2",
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    incapsula:apiId:
        value: 'TODO: "${var.incapsula_api_id}"'
    incapsula:apiKey:
        value: 'TODO: "${var.incapsula_api_key}"'

```
```yaml
resources:
  # Create a site
  example-site:
    type: incapsula:Site
    properties:
      domain: examplesite.com
  # Create a ACL security rule
  example-global-blacklist-ip-rule:
    type: incapsula:AclSecurityRule
    properties:
      ruleId: api.acl.blacklisted_ips
      siteId: ${["example-site"].siteId}
      ips: 192.168.1.1,192.168.1.2
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    incapsula:apiId:
        value: 'TODO: "${var.incapsula_api_id}"'
    incapsula:apiKey:
        value: 'TODO: "${var.incapsula_api_key}"'

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.incapsula.Site;
import com.pulumi.incapsula.SiteArgs;
import com.pulumi.incapsula.AclSecurityRule;
import com.pulumi.incapsula.AclSecurityRuleArgs;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        // Create a site
        var example_site = new Site("example-site", SiteArgs.builder()
            .domain("examplesite.com")
            .build());

        // Create a ACL security rule
        var example_global_blacklist_ip_rule = new AclSecurityRule("example-global-blacklist-ip-rule", AclSecurityRuleArgs.builder()
            .ruleId("api.acl.blacklisted_ips")
            .siteId(example_site.siteId())
            .ips("192.168.1.1,192.168.1.2")
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

The following configuration inputs are supported:

* `apiId` - (Required) The Incapsula API id associated with the account. This can also be
  specified with the `INCAPSULA_API_ID` shell environment variable.
* `apiKey` - (Required) The Incapsula API key. This can also be specified with the
  `INCAPSULA_API_KEY` shell environment variable.