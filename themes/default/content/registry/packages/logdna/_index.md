---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/logdna/logdna/1.16.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Logdna Provider
meta_desc: Provides an overview on how to configure the Pulumi Logdna provider.
layout: package
---

## Generate Provider

The Logdna provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider logdna/logdna
```
## Overview

[![Public Beta](https://img.shields.io/badge/-Public%20Beta-404346?style=flat)](#)

[LogDNA](https://logdna.com) is a centralized log management platform. The LogDNA Pulumi Provider allows organizations to manage certain LogDNA resources (alerts, views, etc) programmatically via Pulumi.
## Example Usage
{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    logdna:servicekey:
        value: xxxxxxxxxxxxxxxxxxxxxxxx
    logdna:url:
        value: https://api.logdna.com

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as logdna from "@pulumi/logdna";

const http500 = new logdna.View("http500", {
    name: "HTTP 500s",
    query: "response:500",
    emailChannels: [{
        emails: ["you@yourdomain.com"],
        operator: "presence",
        terminal: "true",
        triggerinterval: "15m",
        triggerlimit: 15,
    }],
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    logdna:servicekey:
        value: xxxxxxxxxxxxxxxxxxxxxxxx
    logdna:url:
        value: https://api.logdna.com

```
```python
import pulumi
import pulumi_logdna as logdna

http500 = logdna.View("http500",
    name="HTTP 500s",
    query="response:500",
    email_channels=[{
        "emails": ["you@yourdomain.com"],
        "operator": "presence",
        "terminal": "true",
        "triggerinterval": "15m",
        "triggerlimit": 15,
    }])
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    logdna:servicekey:
        value: xxxxxxxxxxxxxxxxxxxxxxxx
    logdna:url:
        value: https://api.logdna.com

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Logdna = Pulumi.Logdna;

return await Deployment.RunAsync(() =>
{
    var http500 = new Logdna.View("http500", new()
    {
        Name = "HTTP 500s",
        Query = "response:500",
        EmailChannels = new[]
        {
            new Logdna.Inputs.ViewEmailChannelArgs
            {
                Emails = new[]
                {
                    "you@yourdomain.com",
                },
                Operator = "presence",
                Terminal = "true",
                Triggerinterval = "15m",
                Triggerlimit = 15,
            },
        },
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
    logdna:servicekey:
        value: xxxxxxxxxxxxxxxxxxxxxxxx
    logdna:url:
        value: https://api.logdna.com

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/logdna/logdna"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := logdna.NewView(ctx, "http500", &logdna.ViewArgs{
			Name:  pulumi.String("HTTP 500s"),
			Query: pulumi.String("response:500"),
			EmailChannels: logdna.ViewEmailChannelArray{
				&logdna.ViewEmailChannelArgs{
					Emails: pulumi.StringArray{
						pulumi.String("you@yourdomain.com"),
					},
					Operator:        pulumi.String("presence"),
					Terminal:        pulumi.String("true"),
					Triggerinterval: pulumi.String("15m"),
					Triggerlimit:    pulumi.Float64(15),
				},
			},
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
    logdna:servicekey:
        value: xxxxxxxxxxxxxxxxxxxxxxxx
    logdna:url:
        value: https://api.logdna.com

```
```yaml
resources:
  http500:
    type: logdna:View
    properties:
      name: HTTP 500s
      query: response:500
      emailChannels:
        - emails:
            - you@yourdomain.com
          operator: presence
          terminal: 'true'
          triggerinterval: 15m
          triggerlimit: 15
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    logdna:servicekey:
        value: xxxxxxxxxxxxxxxxxxxxxxxx
    logdna:url:
        value: https://api.logdna.com

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.logdna.View;
import com.pulumi.logdna.ViewArgs;
import com.pulumi.logdna.inputs.ViewEmailChannelArgs;
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
        var http500 = new View("http500", ViewArgs.builder()
            .name("HTTP 500s")
            .query("response:500")
            .emailChannels(ViewEmailChannelArgs.builder()
                .emails("you@yourdomain.com")
                .operator("presence")
                .terminal("true")
                .triggerinterval("15m")
                .triggerlimit(15)
                .build())
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Pre-requirements and considerations
Before using Pulumi for creating resources in LogDNA, review the following notes:
- Verify Pulumi is installed. The minimum supported version is 0.12.0 and can be checked by running `pulumi version`.
- The configurations seen in the examples will go into a Pulumi configuration file such as `main.tf`.
- Have the service key for your Organization available. To obtain the service key for your LogDNA Organization, go to the LogDNA dashboard and navigate to **Settings > Organization > API Keys** or follow this link [here](https://app.logdna.com/manage/api-keys).
- Authentication is handled via the `servicekey` parameter and can be set in the `provider` configuration section in the `.tf` file.
- When using the LogDNA Pulumi provider, be aware that there is a rate limit of 50 requests per minute.
- If you do not provide a specific a `url` in the provider configuration, the URL defaults to `https://api.logdna.com` (recommended).
- If you want to create an Alert that uses PagerDuty to notify you, you will need to provide LogDNA with the [PagerDuty API key](https://support.pagerduty.com/docs/generating-api-keys#events-api-keys). To ensure that the LogDNA Dashboard properly displays the PagerDuty alert notification channel, we recommend that you first link the PagerDuty service to LogDNA via the [Dashboard UI](https://docs.logdna.com/docs/pagerduty-alert-integration) before using this plugin to create a PagerDuty Alert. You may choose to create such resources first and then link PagerDuty, but be aware that they will not work as intended until the connection is reconciled.
## Configuration Reference

The following configuration inputs are supported by the `provider` section of the `.tf` file:

- `servicekey`: **string *(Required)*** LogDNA Account Service Key. This can be generated or retrieved from Settings > Organization > API Keys.
- `url`: **string** *(Optional; Default: api.logdna.com)* The LogDNA region URL. If you’re configuring an IBM Log Analysis with LogDNA or IBM Cloud Activity Tracker with LogDNA, you’ll need to ensure `url` is set to the [correct endpoint depending on the IBM region](https://cloud.ibm.com/docs/log-analysis?topic=log-analysis-endpoints#endpoints_api).