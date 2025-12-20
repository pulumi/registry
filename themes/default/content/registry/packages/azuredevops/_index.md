---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-azuredevops/v3.11.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi/pulumi-azuredevops/blob/v3.11.0/docs/_index.md
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Azure DevOps Provider
meta_desc: Provides an overview on how to configure the Pulumi Azure DevOps provider.
layout: package
---

## Installation

The Azure DevOps provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/azuredevops`](https://www.npmjs.com/package/@pulumi/azuredevops)
* Python: [`pulumi-azuredevops`](https://pypi.org/project/pulumi-azuredevops/)
* Go: [`github.com/pulumi/pulumi-azuredevops/sdk/v3/go/azuredevops`](https://github.com/pulumi/pulumi-azuredevops)
* .NET: [`Pulumi.Azuredevops`](https://www.nuget.org/packages/Pulumi.Azuredevops)
* Java: [`com.pulumi/azuredevops`](https://central.sonatype.com/artifact/com.pulumi/azuredevops)

## Overview

The Azure DevOps provider can be used to configure Azure DevOps project in [Microsoft Azure](https://azure.microsoft.com/en-us/) using [Azure DevOps Service REST API](https://docs.microsoft.com/en-us/rest/api/azure/devops/?view=azure-devops-rest-7.0)

Use the navigation to the left to read about the available resources.

Interested in the provider's latest features, or want to make sure you're up to date? Check out the changelog for version information and release notes.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as azuredevops from "@pulumi/azuredevops";

const project = new azuredevops.Project("project", {
    name: "Project Name",
    description: "Project Description",
});
```

{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_azuredevops as azuredevops

project = azuredevops.Project("project",
    name="Project Name",
    description="Project Description")
```

{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using AzureDevOps = Pulumi.AzureDevOps;

return await Deployment.RunAsync(() =>
{
    var project = new AzureDevOps.Project("project", new()
    {
        Name = "Project Name",
        Description = "Project Description",
    });

});

```

{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-azuredevops/sdk/v3/go/azuredevops"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := azuredevops.NewProject(ctx, "project", &azuredevops.ProjectArgs{
			Name:        pulumi.String("Project Name"),
			Description: pulumi.String("Project Description"),
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
resources:
  project:
    type: azuredevops:Project
    properties:
      name: Project Name
      description: Project Description
```

{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.azuredevops.Project;
import com.pulumi.azuredevops.ProjectArgs;
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
        var project = new Project("project", ProjectArgs.builder()
            .name("Project Name")
            .description("Project Description")
            .build());

    }
}
```

{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

The following configuration inputs are supported in the provider configuration:

- `orgServiceUrl` - (Required) This is the Azure DevOps organization url. It can also be
  sourced from the `AZDO_ORG_SERVICE_URL` environment variable.

- `personalAccessToken` - This is the Azure DevOps organization personal access
  token. The account corresponding to the token will need "owner" privileges for this
  organization. It can also be sourced from the `AZDO_PERSONAL_ACCESS_TOKEN` environment variable.

- `clientId` - The client id used when authenticating to a service principal or the principal id when
  authenticating with a user specified managed service identity. It can also be sourced from
  the `ARM_CLIENT_ID` or `AZURE_CLIENT_ID` environment variable.

- `clientIdFilePath` - The path to a file containing a client id to authenticate. It can also be sourced from the `ARM_CLIENT_ID_FILE_PATH` environment variable.

- `tenantId` - The tenant id used when authenticating to a service principal.
  It can also be sourced from the `ARM_TENANT_ID` environment variable.

- `auxiliaryTenantIds` - List of auxiliary Tenant IDs required for multi-tenancy and cross-tenant scenarios. This can also be sourced from the `ARM_AUXILIARY_TENANT_IDS` environment variable.

- `clientSecret` - The client secret used to authenticate to a service principal.
  It can also be sourced from the `ARM_CLIENT_SECRET` environment variable.

- `clientSecretPath` - The path to a file containing a client secret to authenticate to a service principal.
  It can also be sourced from the `ARM_CLIENT_SECRET_PATH` or `ARM_CLIENT_SECRET_FILE_PATH` environment variable.

- `clientCertificatePath` - The path to a file containing a certificate to authenticate to a service
  principal, typically a .pfx file.
  It can also be sourced from the `ARM_CLIENT_CERTIFICATE_PATH` environment variable.

- `clientCertificate` - A base64 encoded certificate to authentiate to a service principal.
  It can also be sourced from the `ARM_CLIENT_CERTIFICATE` environment variable.

- `clientCertificatePassword` - This is the password associated with a certificate provided
  by `clientCertificatePath` or `clientCertificate`. It can also be sourced
  from the `ARM_CLIENT_CERTIFICATE_PASSWORD` environment variable.

- `oidcToken` - An OIDC token to authenticate to a service principal.
  It can also be sourced from the `ARM_OIDC_TOKEN` environment variable.

- `oidcTokenFilePath` - The path to a file containing nn OIDC token to authenticate to a service principal.
  It can also be sourced from the `AZDO_TOKEN_PATH`, or `AZURE_FEDERATED_TOKEN_FILE` environment variable.

- `oidcRequestToken` - The bearer token for the request to the OIDC provider. For use when authenticating as a Service Principal using OpenID Connect.
  It can also be sourced from the `ARM_OIDC_REQUEST_TOKEN`, `ACTIONS_ID_TOKEN_REQUEST_TOKEN`, `SYSTEM_ACCESSTOKEN` environment variables.

- `oidcRequestUrl` - The URL for the OIDC provider from which to request an ID token. For use when authenticating as a Service Principal using OpenID Connect.
  It can also be sourced from the `ARM_OIDC_REQUEST_URL`, `ACTIONS_ID_TOKEN_REQUEST_URL` or `SYSTEM_OIDCREQUESTURI` environment variables.

- `oidcAzureServiceConnectionId` - The Azure Pipelines Service Connection ID to use for authentication. This can also be sourced from the `ARM_ADO_PIPELINE_SERVICE_CONNECTION_ID`, `ARM_OIDC_AZURE_SERVICE_CONNECTION_ID`, or `AZURESUBSCRIPTION_SERVICE_CONNECTION_ID` environment variables.

- `useOidc` - Boolean, enables OIDC auth methods. It can also be sourced from the `ARM_USE_OIDC` environment variable.

- `useMsi` - Boolean, enables authentication with a Managed Service Identity in Azure. It can also be sourced from the `ARM_USE_MSI` environment variable.

- `useCli` - Should Azure CLI be used for authentication? This can also be sourced from the `ARM_USE_CLI` environment variable. Defaults to `true`.