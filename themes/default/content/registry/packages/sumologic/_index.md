---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-sumologic/v1.0.11/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Sumologic Provider
meta_desc: Provides an overview on how to configure the Pulumi Sumologic provider.
layout: package
---
## Installation

The Sumologic provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/sumologic`](https://www.npmjs.com/package/@pulumi/sumologic)
* Python: [`pulumi-sumologic`](https://pypi.org/project/pulumi-sumologic/)
* Go: [`github.com/pulumi/pulumi-sumologic/sdk/go/sumologic`](https://github.com/pulumi/pulumi-sumologic)
* .NET: [`Pulumi.Sumologic`](https://www.nuget.org/packages/Pulumi.Sumologic)
* Java: [`com.pulumi/sumologic`](https://central.sonatype.com/artifact/com.pulumi/sumologic)
## Overview

This provider is used to manage resources supported by Sumo Logic. The provider needs to be configured with the proper credentials before it can be used.
## Example Usage
### Configure the Sumo Logic Provider
{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    sumologic:accessId:
        value: 'TODO: "${var.sumologic_access_id}"'
    sumologic:accessKey:
        value: 'TODO: "${var.sumologic_access_key}"'
    sumologic:environment:
        value: us2

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as sumologic from "@pulumi/sumologic";

const config = new pulumi.Config();
// Sumo Logic Access ID
const sumologicAccessId = config.require("sumologicAccessId");
// Sumo Logic Access Key
const sumologicAccessKey = config.require("sumologicAccessKey");
// Create a collector
const collector = new sumologic.Collector("collector", {name: "MyCollector"});
// Create a HTTP source
const httpSource = new sumologic.HttpSource("http_source", {
    name: "http-source",
    category: "my/source/category",
    collectorId: collector.id,
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    sumologic:accessId:
        value: 'TODO: "${var.sumologic_access_id}"'
    sumologic:accessKey:
        value: 'TODO: "${var.sumologic_access_key}"'
    sumologic:environment:
        value: us2

```
```python
import pulumi
import pulumi_sumologic as sumologic

config = pulumi.Config()
# Sumo Logic Access ID
sumologic_access_id = config.require("sumologicAccessId")
# Sumo Logic Access Key
sumologic_access_key = config.require("sumologicAccessKey")
# Create a collector
collector = sumologic.Collector("collector", name="MyCollector")
# Create a HTTP source
http_source = sumologic.HttpSource("http_source",
    name="http-source",
    category="my/source/category",
    collector_id=collector.id)
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    sumologic:accessId:
        value: 'TODO: "${var.sumologic_access_id}"'
    sumologic:accessKey:
        value: 'TODO: "${var.sumologic_access_key}"'
    sumologic:environment:
        value: us2

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using SumoLogic = Pulumi.SumoLogic;

return await Deployment.RunAsync(() =>
{
    var config = new Config();
    // Sumo Logic Access ID
    var sumologicAccessId = config.Require("sumologicAccessId");
    // Sumo Logic Access Key
    var sumologicAccessKey = config.Require("sumologicAccessKey");
    // Create a collector
    var collector = new SumoLogic.Collector("collector", new()
    {
        Name = "MyCollector",
    });

    // Create a HTTP source
    var httpSource = new SumoLogic.HttpSource("http_source", new()
    {
        Name = "http-source",
        Category = "my/source/category",
        CollectorId = collector.Id,
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
    sumologic:accessId:
        value: 'TODO: "${var.sumologic_access_id}"'
    sumologic:accessKey:
        value: 'TODO: "${var.sumologic_access_key}"'
    sumologic:environment:
        value: us2

```
```go
package main

import (
	"github.com/pulumi/pulumi-sumologic/sdk/go/sumologic"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		// Sumo Logic Access ID
		sumologicAccessId := cfg.Require("sumologicAccessId")
		// Sumo Logic Access Key
		sumologicAccessKey := cfg.Require("sumologicAccessKey")
		// Create a collector
		collector, err := sumologic.NewCollector(ctx, "collector", &sumologic.CollectorArgs{
			Name: pulumi.String("MyCollector"),
		})
		if err != nil {
			return err
		}
		// Create a HTTP source
		_, err = sumologic.NewHttpSource(ctx, "http_source", &sumologic.HttpSourceArgs{
			Name:        pulumi.String("http-source"),
			Category:    pulumi.String("my/source/category"),
			CollectorId: collector.ID(),
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
    sumologic:accessId:
        value: 'TODO: "${var.sumologic_access_id}"'
    sumologic:accessKey:
        value: 'TODO: "${var.sumologic_access_key}"'
    sumologic:environment:
        value: us2

```
```yaml
configuration:
  # Setup authentication variables. See "Authentication" section for more details.
  sumologicAccessId:
    type: string
  sumologicAccessKey:
    type: string
resources:
  # Create a collector
  collector:
    type: sumologic:Collector
    properties:
      name: MyCollector
  # Create a HTTP source
  httpSource:
    type: sumologic:HttpSource
    name: http_source
    properties:
      name: http-source
      category: my/source/category
      collectorId: ${collector.id}
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    sumologic:accessId:
        value: 'TODO: "${var.sumologic_access_id}"'
    sumologic:accessKey:
        value: 'TODO: "${var.sumologic_access_key}"'
    sumologic:environment:
        value: us2

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.sumologic.Collector;
import com.pulumi.sumologic.CollectorArgs;
import com.pulumi.sumologic.HttpSource;
import com.pulumi.sumologic.HttpSourceArgs;
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
        final var config = ctx.config();
        final var sumologicAccessId = config.get("sumologicAccessId");
        final var sumologicAccessKey = config.get("sumologicAccessKey");
        // Create a collector
        var collector = new Collector("collector", CollectorArgs.builder()
            .name("MyCollector")
            .build());

        // Create a HTTP source
        var httpSource = new HttpSource("httpSource", HttpSourceArgs.builder()
            .name("http-source")
            .category("my/source/category")
            .collectorId(collector.id())
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
### Configure the Sumo Logic Provider in Admin Mode
{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    sumologic:accessId:
        value: 'TODO: "${var.sumologic_access_id}"'
    sumologic:accessKey:
        value: 'TODO: "${var.sumologic_access_key}"'
    sumologic:adminMode:
        value: true
    sumologic:environment:
        value: us2

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as sumologic from "@pulumi/sumologic";

const config = new pulumi.Config();
// Sumo Logic Access ID
const sumologicAccessId = config.require("sumologicAccessId");
// Sumo Logic Access Key
const sumologicAccessKey = config.require("sumologicAccessKey");
// Look up the Admin Recommended Folder
const folder = sumologic.getAdminRecommendedFolder({});
// Create a folder underneath the Admin Recommended Folder (which requires Admin Mode)
const test = new sumologic.Folder("test", {
    name: "test",
    description: "A test folder",
    parentId: folder.then(folder => folder.id),
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    sumologic:accessId:
        value: 'TODO: "${var.sumologic_access_id}"'
    sumologic:accessKey:
        value: 'TODO: "${var.sumologic_access_key}"'
    sumologic:adminMode:
        value: true
    sumologic:environment:
        value: us2

```
```python
import pulumi
import pulumi_sumologic as sumologic

config = pulumi.Config()
# Sumo Logic Access ID
sumologic_access_id = config.require("sumologicAccessId")
# Sumo Logic Access Key
sumologic_access_key = config.require("sumologicAccessKey")
# Look up the Admin Recommended Folder
folder = sumologic.get_admin_recommended_folder()
# Create a folder underneath the Admin Recommended Folder (which requires Admin Mode)
test = sumologic.Folder("test",
    name="test",
    description="A test folder",
    parent_id=folder.id)
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    sumologic:accessId:
        value: 'TODO: "${var.sumologic_access_id}"'
    sumologic:accessKey:
        value: 'TODO: "${var.sumologic_access_key}"'
    sumologic:adminMode:
        value: true
    sumologic:environment:
        value: us2

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using SumoLogic = Pulumi.SumoLogic;

return await Deployment.RunAsync(() =>
{
    var config = new Config();
    // Sumo Logic Access ID
    var sumologicAccessId = config.Require("sumologicAccessId");
    // Sumo Logic Access Key
    var sumologicAccessKey = config.Require("sumologicAccessKey");
    // Look up the Admin Recommended Folder
    var folder = SumoLogic.GetAdminRecommendedFolder.Invoke();

    // Create a folder underneath the Admin Recommended Folder (which requires Admin Mode)
    var test = new SumoLogic.Folder("test", new()
    {
        Name = "test",
        Description = "A test folder",
        ParentId = folder.Apply(getAdminRecommendedFolderResult => getAdminRecommendedFolderResult.Id),
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
    sumologic:accessId:
        value: 'TODO: "${var.sumologic_access_id}"'
    sumologic:accessKey:
        value: 'TODO: "${var.sumologic_access_key}"'
    sumologic:adminMode:
        value: true
    sumologic:environment:
        value: us2

```
```go
package main

import (
	"github.com/pulumi/pulumi-sumologic/sdk/go/sumologic"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		// Sumo Logic Access ID
		sumologicAccessId := cfg.Require("sumologicAccessId")
		// Sumo Logic Access Key
		sumologicAccessKey := cfg.Require("sumologicAccessKey")
		// Look up the Admin Recommended Folder
		folder, err := sumologic.GetAdminRecommendedFolder(ctx, &sumologic.GetAdminRecommendedFolderArgs{}, nil)
		if err != nil {
			return err
		}
		// Create a folder underneath the Admin Recommended Folder (which requires Admin Mode)
		_, err = sumologic.NewFolder(ctx, "test", &sumologic.FolderArgs{
			Name:        pulumi.String("test"),
			Description: pulumi.String("A test folder"),
			ParentId:    pulumi.String(folder.Id),
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
    sumologic:accessId:
        value: 'TODO: "${var.sumologic_access_id}"'
    sumologic:accessKey:
        value: 'TODO: "${var.sumologic_access_key}"'
    sumologic:adminMode:
        value: true
    sumologic:environment:
        value: us2

```
```yaml
configuration:
  sumologicAccessId:
    type: string
  sumologicAccessKey:
    type: string
resources:
  # Create a folder underneath the Admin Recommended Folder (which requires Admin Mode)
  test:
    type: sumologic:Folder
    properties:
      name: test
      description: A test folder
      parentId: ${folder.id}
variables:
  # Look up the Admin Recommended Folder
  folder:
    fn::invoke:
      Function: sumologic:getAdminRecommendedFolder
      Arguments: {}
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    sumologic:accessId:
        value: 'TODO: "${var.sumologic_access_id}"'
    sumologic:accessKey:
        value: 'TODO: "${var.sumologic_access_key}"'
    sumologic:adminMode:
        value: true
    sumologic:environment:
        value: us2

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.sumologic.SumologicFunctions;
import com.pulumi.sumologic.inputs.GetAdminRecommendedFolderArgs;
import com.pulumi.sumologic.Folder;
import com.pulumi.sumologic.FolderArgs;
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
        final var config = ctx.config();
        final var sumologicAccessId = config.get("sumologicAccessId");
        final var sumologicAccessKey = config.get("sumologicAccessKey");
        // Look up the Admin Recommended Folder
        final var folder = SumologicFunctions.getAdminRecommendedFolder();

        // Create a folder underneath the Admin Recommended Folder (which requires Admin Mode)
        var test = new Folder("test", FolderArgs.builder()
            .name("test")
            .description("A test folder")
            .parentId(folder.applyValue(getAdminRecommendedFolderResult -> getAdminRecommendedFolderResult.id()))
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Authentication
The Sumo Logic Provider offers a flexible means of providing credentials for authentication. The following methods are supported and explained below:

- Static credentials
- Environment variables
### Static credentials
Static credentials can be provided by adding an `accessId` and `accessKey` in-line in the Sumo Logic provider configuration:

Usage:
{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    sumologic:accessId:
        value: your-access-id
    sumologic:accessKey:
        value: your-access-key
    sumologic:environment:
        value: us2

```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    sumologic:accessId:
        value: your-access-id
    sumologic:accessKey:
        value: your-access-key
    sumologic:environment:
        value: us2

```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    sumologic:accessId:
        value: your-access-id
    sumologic:accessKey:
        value: your-access-key
    sumologic:environment:
        value: us2

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    sumologic:accessId:
        value: your-access-id
    sumologic:accessKey:
        value: your-access-key
    sumologic:environment:
        value: us2

```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    sumologic:accessId:
        value: your-access-id
    sumologic:accessKey:
        value: your-access-key
    sumologic:environment:
        value: us2

```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    sumologic:accessId:
        value: your-access-id
    sumologic:accessKey:
        value: your-access-key
    sumologic:environment:
        value: us2

```

{{% /choosable %}}
{{< /chooser >}}
### Environment variables
You can provide your credentials via the `SUMOLOGIC_ACCESSID` and `SUMOLOGIC_ACCESSKEY` environment variables, representing your Sumo Logic Access ID and Sumo Logic Access Key, respectively.

Usage:
{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs

```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python

```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go

```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml

```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java

```

{{% /choosable %}}
{{< /chooser >}}

```bash
$ export SUMOLOGIC_ACCESSID="your-access-id"
$ export SUMOLOGIC_ACCESSKEY="your-access-key"
$ export SUMOLOGIC_ENVIRONMENT=us2
$ pulumi preview
```
## Configuration Reference
- `accessId` - (Required) This is the Sumo Logic Access ID. It must be provided, but it can also be source from the SUMOLOGIC_ACCESSID environment variable.
- `accessKey` - (Required) This is the Sumo Logic Access Key. It must be provided, but it can also be sourced from the SUMOLOGIC_ACCESSKEY variable.
- `environment` - (Required) This is the API endpoint to use. See the [Sumo Logic documentation](https://help.sumologic.com/APIs/General_API_Information/Sumo_Logic_Endpoints_and_Firewall_Security) for details on which environment you should use. It must be provided, but it can be sourced from the SUMOLOGIC_ENVIRONMENT variable.
## Common Source Properties

The following properties are common to ALL sources and can be used to configure each source.

- `collectorId` - (Required) The ID of the collector to attach this source to.
- `name` - (Required) The name of the source. This is required, and has to be unique in the scope of the collector. Changing this will force recreation the source.
- `description` - (Optional) Description of the source.
- `category` - (Optional) The source category this source logs to.
- `hostName` - (Optional) The source host this source logs to.
- `timezone` - (Optional) The timezone assigned to the source. The value follows the [tzdata](https://en.wikipedia.org/wiki/Tz_database) naming convention.
- `automaticDateParsing` - (Optional) Determines if timestamp information is parsed or not. Type true to enable automatic parsing of dates (the default setting); type false to disable. If disabled, no timestamp information is parsed at all.
- `multilineProcessingEnabled` - (Optional) Type true to enable; type false to disable. The default setting is true. Consider setting to false to avoid unnecessary processing if you are collecting single message per line files (for example, Linux system.log). If you're working with multiline messages (for example, log4J or exception stack traces), keep this setting enabled.
- `useAutolineMatching` - (Optional) Type true to enable if you'd like message boundaries to be inferred automatically; type false to prevent message boundaries from being automatically inferred (equivalent to the Infer Boundaries option in the UI). The default setting is true.
- `manualPrefixRegexp` - (Optional) When using useAutolineMatching=false, type a regular expression that matches the first line of the message to manually create the boundary. Note that any special characters in the regex, such as backslashes or double quotes, must be escaped.
- `forceTimezone` - (Optional) Type true to force the source to use a specific time zone, otherwise type false to use the time zone found in the logs. The default setting is false.
- `defaultDateFormats` - (Optional) Define the format for the timestamps present in your log messages. You can specify a locator regex to identify where timestamps appear in log lines. Requires 'automatic_date_parsing' set to True.
  + `format` - (Required) The timestamp format supplied as a Java SimpleDateFormat, or "epoch" if the timestamp is in epoch format.
  + `locator` - (Optional) Regular expression to locate the timestamp within the messages.

- `filters` - (Optional) If you'd like to add a filter to the source.
  + `filterType` - (Required) The type of filter to apply. (Exclude, Include, Mask, or Hash)
  + `name` - (Required) The Name for the filter.
  + `regexp` - (Required) Regular expression to match within the messages. When used with Incude/Exclude the expression must match the entire message. When used with Mask/Hash rules the expression must contain an unnamed capture group to hash/mask.
  + `mask` - (Optional) When applying a Mask rule, replaces the detected expression with this string.

- `hashAlgorithm` - (Optional) Define the hash algorithm used for Hash type filters. Available values are "MD5" and "SHA-256". The default value will be "MD5".
- `cutoffTimestamp` - (Optional) Only collect data more recent than this timestamp, specified as milliseconds since epoch (13 digit). This maps to the `Collection should begin` field on the UI. Example: using `1663786159000` will set the cutoff timestamp to `Wednesday, September 21, 2022 6:49:19 PM GMT`
- `cutoffRelativeTime` - (Optional) Can be specified instead of cutoffTimestamp to provide a relative offset with respect to the current time.This maps to the `Collection should begin` field on the UI. Example: use -1h, -1d, or -1w to collect data that's less than one hour, one day, or one week old, respectively.
- `fields` - (Optional) Map containing key/value pairs.
## Configuring SNS Subscription
This is supported in the following resources.
- `sumologic.CloudfrontSource`
- `sumologic.CloudtrailSource`
- `sumologic.ElbSource`
- `sumologic.S3AuditSource`
- `sumologic.S3Source`

Steps to configure SNS subscription and sync the state in terrform:
- Step 1: Create the source via pulumi.
- Step 2: Setup [SNS subscription](https://help.sumologic.com/03Send-Data/Sources/02Sources-for-Hosted-Collectors/Amazon-Web-Services/AWS_Sources#set-up-sns-in-aws-highly-recommended) outside of pulumi on Sumologic UI.
- Step 3: Run `pulumi preview --refresh` to review the changes and verify the state with the SNS subscription information. Make sure only `snsTopicOrSubscriptionArn` is updated. If SNS has been successfully configured and has received a subscription confirmation request `isSuccess` parameter will be true.
- Step 4: Apply the changes with `pulumi up --refresh`.