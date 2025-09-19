---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-signalfx/v7.18.1/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Signalfx Provider
meta_desc: Provides an overview on how to configure the Pulumi Signalfx provider.
layout: package
---

## Installation

The Signalfx provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/signalfx`](https://www.npmjs.com/package/@pulumi/signalfx)
* Python: [`pulumi-signalfx`](https://pypi.org/project/pulumi-signalfx/)
* Go: [`github.com/pulumi/pulumi-signalfx/sdk/v7/go/signalfx`](https://github.com/pulumi/pulumi-signalfx)
* .NET: [`Pulumi.Signalfx`](https://www.nuget.org/packages/Pulumi.Signalfx)
* Java: [`com.pulumi/signalfx`](https://central.sonatype.com/artifact/com.pulumi/signalfx)

## Overview

The [Splunk Observability Cloud](https://www.splunk.com/en_us/products/observability.html) provider, formerly known as SignalFx Pulumi provider, lets you interact with the resources supported by Splunk Observability Cloud. You must configure the provider with credentials before using it.

Use the navigation to learn more about the available resources.

> The next major release of the provider (v10) is will require a minimum Pulumi version of 1.11.0.
Please prepare to migrate to a newer version of Pulumi soon.
This will allow the provider to handle secrets more securely and take advantage of new features in the Pulumi Plugin Framework.
# Learn about Splunk Observability Cloud

To learn more about Splunk Observability Cloud and its features, see [the official documentation](https://docs.splunk.com/observability/en/).

You can use the SignalFlow programming language to create charts and detectors using `programText`. For more information about SignalFlow, see the [Splunk developer documentation](https://dev.splunk.com/observability/docs/signalflow/).
# Authentication

When authenticating to the Splunk Observability Cloud API you can use:

1. An Org token.
2. A Session token.
3. A Service account.

See [Authenticate API Requests](https://dev.splunk.com/observability/docs/apibasics/authentication_basics/) in the Splunk developer documentation.

Session tokens are short-lived and provide administrative permissions to edit integrations. They expire relatively quickly, but let you manipulate some sensitive resources. Resources that require session tokens are flagged in their documentation.

A Service account is term used when a user is created within organization that can login via Username and Password, this allows for a *Session Token* to be created by the pulumi provider and then used throughout the application.

ℹ️ **NOTE** Separate the less sensitive resources, such as dashboards, from the more sensitive ones, such as integrations, to avoid having to change tokens.
## Example

The following example shows how to configure the Splunk Observability Cloud provider for Pulumi:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    signalfx:authToken:
        value: 'TODO: "${var.signalfx_auth_token}"'

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as signalfx from "@pulumi/signalfx";

// Create a new detector
const _default = new signalfx.Detector("default", {});
// Create a new dashboard
const defaultDashboard = new signalfx.Dashboard("default", {});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    signalfx:authToken:
        value: 'TODO: "${var.signalfx_auth_token}"'

```
```python
import pulumi
import pulumi_signalfx as signalfx

# Create a new detector
default = signalfx.Detector("default")
# Create a new dashboard
default_dashboard = signalfx.Dashboard("default")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    signalfx:authToken:
        value: 'TODO: "${var.signalfx_auth_token}"'

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using SignalFx = Pulumi.SignalFx;

return await Deployment.RunAsync(() =>
{
    // Create a new detector
    var @default = new SignalFx.Detector("default");

    // Create a new dashboard
    var defaultDashboard = new SignalFx.Dashboard("default");

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    signalfx:authToken:
        value: 'TODO: "${var.signalfx_auth_token}"'

```
```go
package main

import (
	"github.com/pulumi/pulumi-signalfx/sdk/v7/go/signalfx"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a new detector
		_, err := signalfx.NewDetector(ctx, "default", nil)
		if err != nil {
			return err
		}
		// Create a new dashboard
		_, err = signalfx.NewDashboard(ctx, "default", nil)
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
    signalfx:authToken:
        value: 'TODO: "${var.signalfx_auth_token}"'

```
```yaml
resources:
  # Create a new detector
  default:
    type: signalfx:Detector
  # Create a new dashboard
  defaultDashboard:
    type: signalfx:Dashboard
    name: default
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    signalfx:authToken:
        value: 'TODO: "${var.signalfx_auth_token}"'

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.signalfx.Detector;
import com.pulumi.signalfx.Dashboard;
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
        // Create a new detector
        var default_ = new Detector("default");

        // Create a new dashboard
        var defaultDashboard = new Dashboard("defaultDashboard");

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    signalfx:email:
        value: service.account@example
    signalfx:organizationId:
        value: 'TODO: "${var.service_account_org_id}"'
    signalfx:password:
        value: 'TODO: "${var.service_account_password}"'

```
# Feature Previews

To allow for more experimental features to be added into the provider, a feature can be added behind a preview gate that defaults to being off and requires a user to opt into the change. Once a feature has been added into the provider, in can be set to globally available which will default to the feature being on by default.

There is an opportunity for the user to opt out of a globally available feature if an issue is experienced. If that is the case, please raise a support case with the provider configuration and any error messages.

The feature preview can be enabled by the following example:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    signalfx:featurePreview:
        value:
            feature-01: true
            feature-02: false

```

ℹ️ **NOTE** Preview features are a subject to change and/or removal in a future version of the provider.
## Configuration Reference

- `apiUrl` (String) API URL for your Splunk Observability Cloud org, may include a realm
- `authToken` (String) Splunk Observability Cloud auth token
- `customAppUrl` (String, Deprecated) Application URL for your Splunk Observability Cloud org, often customized for organizations using SSO
- `email` (String) Used to create a session token instead of an API token, it requires the account to be configured to login with Email and Password
- `featurePreview` (Map of Boolean) Allows for users to opt-in to new features that are considered experimental or not ready for general availability yet.
- `organizationId` (String) Required if the user is configured to be part of multiple organizations
- `password` (String, Sensitive) Used to create a session token instead of an API token, it requires the account to be configured to login with Email and Password
- `retryMaxAttempts` (Number) Max retries for a single HTTP call. Defaults to 4
- `retryWaitMaxSeconds` (Number) Maximum retry wait for a single HTTP call in seconds. Defaults to 30
- `retryWaitMinSeconds` (Number) Minimum retry wait for a single HTTP call in seconds. Defaults to 1
- `tags` (List of String) Allows for Tags to be added by default to resources that allow for tags to be included. If there is already tags configured, the global tags are added in prefix.
- `teams` (List of String) Allows for teams to be defined at a provider level, and apply to all applicable resources created.
- `timeoutSeconds` (Number) Timeout duration for a single HTTP call in seconds. Defaults to 120