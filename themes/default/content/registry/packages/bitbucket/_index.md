---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/drfaust92/bitbucket/2.49.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Bitbucket Provider
meta_desc: Provides an overview on how to configure the Pulumi Bitbucket provider.
layout: package
---

## Generate Provider

The Bitbucket provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider drfaust92/bitbucket
```
## Overview

The Bitbucket provider allows you to manage resources including repositories,
webhooks, and default reviewers.

Use the navigation to the left to read about the available resources.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    bitbucket:password:
        value: idoillusions
    bitbucket:username:
        value: GobBluthe

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as bitbucket from "@pulumi/bitbucket";

const illusions = new bitbucket.Repository("illusions", {
    owner: "theleagueofmagicians",
    name: "illusions",
    scm: "hg",
    isPrivate: true,
});
const project = new bitbucket.Project("project", {
    owner: "theleagueofmagicians",
    name: "illusions-project",
    key: "ILLUSIONSPROJ",
    isPrivate: true,
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    bitbucket:password:
        value: idoillusions
    bitbucket:username:
        value: GobBluthe

```
```python
import pulumi
import pulumi_bitbucket as bitbucket

illusions = bitbucket.Repository("illusions",
    owner="theleagueofmagicians",
    name="illusions",
    scm="hg",
    is_private=True)
project = bitbucket.Project("project",
    owner="theleagueofmagicians",
    name="illusions-project",
    key="ILLUSIONSPROJ",
    is_private=True)
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    bitbucket:password:
        value: idoillusions
    bitbucket:username:
        value: GobBluthe

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Bitbucket = Pulumi.Bitbucket;

return await Deployment.RunAsync(() =>
{
    var illusions = new Bitbucket.Repository("illusions", new()
    {
        Owner = "theleagueofmagicians",
        Name = "illusions",
        Scm = "hg",
        IsPrivate = true,
    });

    var project = new Bitbucket.Project("project", new()
    {
        Owner = "theleagueofmagicians",
        Name = "illusions-project",
        Key = "ILLUSIONSPROJ",
        IsPrivate = true,
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
    bitbucket:password:
        value: idoillusions
    bitbucket:username:
        value: GobBluthe

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/bitbucket/v2/bitbucket"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := bitbucket.NewRepository(ctx, "illusions", &bitbucket.RepositoryArgs{
			Owner:     pulumi.String("theleagueofmagicians"),
			Name:      pulumi.String("illusions"),
			Scm:       pulumi.String("hg"),
			IsPrivate: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		_, err = bitbucket.NewProject(ctx, "project", &bitbucket.ProjectArgs{
			Owner:     pulumi.String("theleagueofmagicians"),
			Name:      pulumi.String("illusions-project"),
			Key:       pulumi.String("ILLUSIONSPROJ"),
			IsPrivate: pulumi.Bool(true),
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
    bitbucket:password:
        value: idoillusions
    bitbucket:username:
        value: GobBluthe

```
```yaml
resources:
  illusions:
    type: bitbucket:Repository
    properties:
      owner: theleagueofmagicians
      name: illusions
      scm: hg
      isPrivate: true
  project:
    type: bitbucket:Project
    properties:
      owner: theleagueofmagicians
      name: illusions-project
      key: ILLUSIONSPROJ
      isPrivate: true
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    bitbucket:password:
        value: idoillusions
    bitbucket:username:
        value: GobBluthe

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.bitbucket.Repository;
import com.pulumi.bitbucket.RepositoryArgs;
import com.pulumi.bitbucket.Project;
import com.pulumi.bitbucket.ProjectArgs;
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
        var illusions = new Repository("illusions", RepositoryArgs.builder()
            .owner("theleagueofmagicians")
            .name("illusions")
            .scm("hg")
            .isPrivate(true)
            .build());

        var project = new Project("project", ProjectArgs.builder()
            .owner("theleagueofmagicians")
            .name("illusions-project")
            .key("ILLUSIONSPROJ")
            .isPrivate(true)
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

The following configuration inputs are supported in the provider configuration:

* `username` - (Optional) Username to use for authentication via [Basic
  Auth](https://developer.atlassian.com/cloud/bitbucket/rest/intro/#basic-auth).
  You can also set this via the `BITBUCKET_USERNAME` environment variable.
  If configured, requires `password` to be configured as well.

* `password` - (Optional) Password to use for authentication via [Basic
  Auth](https://developer.atlassian.com/cloud/bitbucket/rest/intro/#basic-auth).
  Please note that this has to be an [App
  Password](https://support.atlassian.com/bitbucket-cloud/docs/app-passwords/)
  that has to be created in the [Account
  Settings](https://bitbucket.org/account/settings/app-passwords/). You can
  also set this via the `BITBUCKET_PASSWORD` environment variable. If
  configured, requires `username` to be configured as well.

* `oauthClientId` - (Optional) OAuth client ID to use for authentication via
  [Client Credentials
  Grant](https://developer.atlassian.com/cloud/bitbucket/rest/intro/#3--client-credentials-grant--4-4-).
  You can also set this via the `BITBUCKET_OAUTH_CLIENT_ID` environment
  variable. If configured, requires `oauthClientSecret` to be configured as
  well.

* `oauthClientSecret` - (Optional) OAuth client secret to use for authentication via
  [Client Credentials
  Grant](https://developer.atlassian.com/cloud/bitbucket/rest/intro/#3--client-credentials-grant--4-4-).
  You can also set this via the `BITBUCKET_OAUTH_CLIENT_SECRET` environment
  variable. If configured, requires `oauthClientId` to be configured as well.

* `oauthToken` - (Optional) An OAuth access token used for authentication via
  [OAuth](https://developer.atlassian.com/cloud/bitbucket/rest/intro/#oauth-2-0).
  You can also set this via the `BITBUCKET_OAUTH_TOKEN` environment variable.
## OAuth2 Scopes

To interacte with the Bitbucket API, an [App
Password](https://support.atlassian.com/bitbucket-cloud/docs/app-passwords/) or
[OAuth Client
Credentials](https://support.atlassian.com/bitbucket-cloud/docs/use-oauth-on-bitbucket-cloud/)
are required.

App passwords and OAuth client credentials are limited in scope, each API
requires certain scope to interact with, each resource doc will specify what
are the scopes required to use that resource.

See the [Bitbucket OAuth
Documentation](https://support.atlassian.com/bitbucket-cloud/docs/use-oauth-on-bitbucket-cloud/)
for more information on scopes.