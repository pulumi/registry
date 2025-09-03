---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/vantage-sh/vantage/0.1.63-beta/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Vantage Provider
meta_desc: Provides an overview on how to configure the Pulumi Vantage provider.
layout: package
---

## Generate Provider

The Vantage provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider vantage-sh/vantage
```
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    vantage:apiToken:
        value: 'TODO: var.api_token'

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as vantage from "@pulumi/vantage";

const aws = new vantage.Folder("aws", {title: "AWS Costs"});
const awsCostReport = new vantage.CostReport("aws", {
    folderToken: aws.token,
    filter: "costs.provider = 'aws'",
    title: "AWS Costs",
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    vantage:apiToken:
        value: 'TODO: var.api_token'

```
```python
import pulumi
import pulumi_vantage as vantage

aws = vantage.Folder("aws", title="AWS Costs")
aws_cost_report = vantage.CostReport("aws",
    folder_token=aws.token,
    filter="costs.provider = 'aws'",
    title="AWS Costs")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    vantage:apiToken:
        value: 'TODO: var.api_token'

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Vantage = Pulumi.Vantage;

return await Deployment.RunAsync(() =>
{
    var aws = new Vantage.Folder("aws", new()
    {
        Title = "AWS Costs",
    });

    var awsCostReport = new Vantage.CostReport("aws", new()
    {
        FolderToken = aws.Token,
        Filter = "costs.provider = 'aws'",
        Title = "AWS Costs",
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
    vantage:apiToken:
        value: 'TODO: var.api_token'

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/vantage/vantage"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		aws, err := vantage.NewFolder(ctx, "aws", &vantage.FolderArgs{
			Title: pulumi.String("AWS Costs"),
		})
		if err != nil {
			return err
		}
		_, err = vantage.NewCostReport(ctx, "aws", &vantage.CostReportArgs{
			FolderToken: aws.Token,
			Filter:      pulumi.String("costs.provider = 'aws'"),
			Title:       pulumi.String("AWS Costs"),
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
    vantage:apiToken:
        value: 'TODO: var.api_token'

```
```yaml
resources:
  aws:
    type: vantage:Folder
    properties:
      title: AWS Costs
  awsCostReport:
    type: vantage:CostReport
    name: aws
    properties:
      folderToken: ${aws.token}
      filter: costs.provider = 'aws'
      title: AWS Costs
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    vantage:apiToken:
        value: 'TODO: var.api_token'

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.vantage.Folder;
import com.pulumi.vantage.FolderArgs;
import com.pulumi.vantage.CostReport;
import com.pulumi.vantage.CostReportArgs;
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
        var aws = new Folder("aws", FolderArgs.builder()
            .title("AWS Costs")
            .build());

        var awsCostReport = new CostReport("awsCostReport", CostReportArgs.builder()
            .folderToken(aws.token())
            .filter("costs.provider = 'aws'")
            .title("AWS Costs")
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

- `apiToken` (String, Sensitive)
- `host` (String)