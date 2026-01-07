---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-databricks/v1.80.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi/pulumi-databricks/blob/v1.80.0/docs/_index.md
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Databricks Provider
meta_desc: Provides an overview on how to configure the Pulumi Databricks provider.
layout: package
---

## Installation

The Databricks provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/databricks`](https://www.npmjs.com/package/@pulumi/databricks)
* Python: [`pulumi-databricks`](https://pypi.org/project/pulumi-databricks/)
* Go: [`github.com/pulumi/pulumi-databricks/sdk/go/databricks`](https://github.com/pulumi/pulumi-databricks)
* .NET: [`Pulumi.Databricks`](https://www.nuget.org/packages/Pulumi.Databricks)
* Java: [`com.pulumi/databricks`](https://central.sonatype.com/artifact/com.pulumi/databricks)

## Overview

Use the Databricks Pulumi provider to interact with almost all of [Databricks](http://databricks.com/) resources.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs

```

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as databricks from "@pulumi/databricks";
import * as std from "@pulumi/std";

function notImplemented(message: string) {
    throw new Error(message);
}

const me = databricks.getCurrentUser({});
const latest = databricks.getSparkVersion({});
const smallest = databricks.getNodeType({
    localDisk: true,
});
const _this = new databricks.Notebook("this", {
    path: me.then(me => `${me.home}/Pulumi`),
    language: "PYTHON",
    contentBase64: std.abspath({
        input: notImplemented("path.module"),
    }).then(invoke => std.base64encode({
        input: `# created from ${invoke.result}
display(spark.range(10))
`,
    })).then(invoke => invoke.result),
});
const thisJob = new databricks.Job("this", {
    name: me.then(me => `Pulumi Demo (${me.alphanumeric})`),
    tasks: [{
        taskKey: "task1",
        notebookTask: {
            notebookPath: _this.path,
        },
        newCluster: {
            numWorkers: 1,
            sparkVersion: latest.then(latest => latest.id),
            nodeTypeId: smallest.then(smallest => smallest.id),
        },
    }],
});
export const notebookUrl = _this.url;
export const jobUrl = thisJob.url;
```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python

```

```python
import pulumi
import pulumi_databricks as databricks
import pulumi_std as std


def not_implemented(msg):
    raise NotImplementedError(msg)

me = databricks.get_current_user()
latest = databricks.get_spark_version()
smallest = databricks.get_node_type(local_disk=True)
this = databricks.Notebook("this",
    path=f"{me.home}/Pulumi",
    language="PYTHON",
    content_base64=std.base64encode(input=f"""# created from {std.abspath(input=not_implemented("path.module")).result}
display(spark.range(10))
""").result)
this_job = databricks.Job("this",
    name=f"Pulumi Demo ({me.alphanumeric})",
    tasks=[{
        "task_key": "task1",
        "notebook_task": {
            "notebook_path": this.path,
        },
        "new_cluster": {
            "num_workers": 1,
            "spark_version": latest.id,
            "node_type_id": smallest.id,
        },
    }])
pulumi.export("notebookUrl", this.url)
pulumi.export("jobUrl", this_job.url)
```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet

```

```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Databricks = Pulumi.Databricks;
using Std = Pulumi.Std;


object NotImplemented(string errorMessage)
{
    throw new System.NotImplementedException(errorMessage);
}

return await Deployment.RunAsync(() =>
{
    var me = Databricks.GetCurrentUser.Invoke();

    var latest = Databricks.GetSparkVersion.Invoke();

    var smallest = Databricks.GetNodeType.Invoke(new()
    {
        LocalDisk = true,
    });

    var @this = new Databricks.Notebook("this", new()
    {
        Path = $"{me.Apply(getCurrentUserResult => getCurrentUserResult.Home)}/Pulumi",
        Language = "PYTHON",
        ContentBase64 = Std.Abspath.Invoke(new()
        {
            Input = NotImplemented("path.module"),
        }).Apply(invoke => Std.Base64encode.Invoke(new()
        {
            Input = @$"# created from {invoke.Result}
display(spark.range(10))
",
        })).Apply(invoke => invoke.Result),
    });

    var thisJob = new Databricks.Job("this", new()
    {
        Name = $"Pulumi Demo ({me.Apply(getCurrentUserResult => getCurrentUserResult.Alphanumeric)})",
        Tasks = new[]
        {
            new Databricks.Inputs.JobTaskArgs
            {
                TaskKey = "task1",
                NotebookTask = new Databricks.Inputs.JobTaskNotebookTaskArgs
                {
                    NotebookPath = @this.Path,
                },
                NewCluster = new Databricks.Inputs.JobTaskNewClusterArgs
                {
                    NumWorkers = 1,
                    SparkVersion = latest.Apply(getSparkVersionResult => getSparkVersionResult.Id),
                    NodeTypeId = smallest.Apply(getNodeTypeResult => getNodeTypeResult.Id),
                },
            },
        },
    });

    return new Dictionary<string, object?>
    {
        ["notebookUrl"] = @this.Url,
        ["jobUrl"] = thisJob.Url,
    };
});

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go

```

```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-databricks/sdk/go/databricks"
	"github.com/pulumi/pulumi-std/sdk/go/std"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func notImplemented(message string) pulumi.AnyOutput {
	panic(message)
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		me, err := databricks.GetCurrentUser(ctx, map[string]interface{}{}, nil)
		if err != nil {
			return err
		}
		latest, err := databricks.GetSparkVersion(ctx, &databricks.GetSparkVersionArgs{}, nil)
		if err != nil {
			return err
		}
		smallest, err := databricks.GetNodeType(ctx, &databricks.GetNodeTypeArgs{
			LocalDisk: pulumi.BoolRef(true),
		}, nil)
		if err != nil {
			return err
		}
		invokeBase64encode, err := std.Base64encode(ctx, &std.Base64encodeArgs{
			Input: fmt.Sprintf("# created from %v\ndisplay(spark.range(10))\n", std.Abspath(ctx, &std.AbspathArgs{
				Input: notImplemented("path.module"),
			}, nil).Result),
		}, nil)
		if err != nil {
			return err
		}
		this, err := databricks.NewNotebook(ctx, "this", &databricks.NotebookArgs{
			Path:          pulumi.Sprintf("%v/Pulumi", me.Home),
			Language:      pulumi.String("PYTHON"),
			ContentBase64: pulumi.String(invokeBase64encode.Result),
		})
		if err != nil {
			return err
		}
		thisJob, err := databricks.NewJob(ctx, "this", &databricks.JobArgs{
			Name: pulumi.Sprintf("Pulumi Demo (%v)", me.Alphanumeric),
			Tasks: databricks.JobTaskArray{
				&databricks.JobTaskArgs{
					TaskKey: pulumi.String("task1"),
					NotebookTask: &databricks.JobTaskNotebookTaskArgs{
						NotebookPath: this.Path,
					},
					NewCluster: &databricks.JobTaskNewClusterArgs{
						NumWorkers:   pulumi.Int(1),
						SparkVersion: pulumi.String(latest.Id),
						NodeTypeId:   pulumi.String(smallest.Id),
					},
				},
			},
		})
		if err != nil {
			return err
		}
		ctx.Export("notebookUrl", this.Url)
		ctx.Export("jobUrl", thisJob.Url)
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

```

```yaml
Example currently unavailable in this language
```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java

```

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.databricks.DatabricksFunctions;
import com.pulumi.databricks.inputs.GetSparkVersionArgs;
import com.pulumi.databricks.inputs.GetNodeTypeArgs;
import com.pulumi.databricks.Notebook;
import com.pulumi.databricks.NotebookArgs;
import com.pulumi.std.StdFunctions;
import com.pulumi.std.inputs.AbspathArgs;
import com.pulumi.std.inputs.Base64encodeArgs;
import com.pulumi.databricks.Job;
import com.pulumi.databricks.JobArgs;
import com.pulumi.databricks.inputs.JobTaskArgs;
import com.pulumi.databricks.inputs.JobTaskNotebookTaskArgs;
import com.pulumi.databricks.inputs.JobTaskNewClusterArgs;
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
        final var me = DatabricksFunctions.getCurrentUser(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference);

        final var latest = DatabricksFunctions.getSparkVersion(GetSparkVersionArgs.builder()
            .build());

        final var smallest = DatabricksFunctions.getNodeType(GetNodeTypeArgs.builder()
            .localDisk(true)
            .build());

        var this_ = new Notebook("this", NotebookArgs.builder()
            .path(String.format("%s/Pulumi", me.home()))
            .language("PYTHON")
            .contentBase64(StdFunctions.base64encode(Base64encodeArgs.builder()
                .input("""
# created from %s
display(spark.range(10))
", StdFunctions.abspath(AbspathArgs.builder()
                    .input("TODO: call notImplemented")
                    .build()).result()))
                .build()).result())
            .build());

        var thisJob = new Job("thisJob", JobArgs.builder()
            .name(String.format("Pulumi Demo (%s)", me.alphanumeric()))
            .tasks(JobTaskArgs.builder()
                .taskKey("task1")
                .notebookTask(JobTaskNotebookTaskArgs.builder()
                    .notebookPath(this_.path())
                    .build())
                .newCluster(JobTaskNewClusterArgs.builder()
                    .numWorkers(1)
                    .sparkVersion(latest.id())
                    .nodeTypeId(smallest.id())
                    .build())
                .build())
            .build());

        ctx.export("notebookUrl", this_.url());
        ctx.export("jobUrl", thisJob.url());
    }
}
```

{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

Most provider arguments can be configured either directly in the `provider "databricks"` block or by setting an environment variable, listed for each argument below.

The provider configuration supports the following arguments:

* `host` - (optional, environment variable `DATABRICKS_HOST`) The host of the Databricks account or workspace. See `host` argument for more information.
* `accountId` - (required for account-level operations, environment variable `DATABRICKS_ACCOUNT_ID`) Account ID found in the top right corner of [Accounts Console](https://accounts.cloud.databricks.com/). **Note: do NOT set this variable when using a workspace-level provider. If set, you may see `...invalid Databricks Account configuration` errors**.
* `azureWorkspaceResourceId` - (optional, environment variable `DATABRICKS_AZURE_RESOURCE_ID`) `id` attribute of azurermDatabricksWorkspace resource. Combination of subscription id, resource group name, and workspace name. Required when authenticating using Azure MSI.

The following arguments control the provider authentication:

* `profile` - (optional, environment variable `DATABRICKS_CONFIG_PROFILE`) Connection profile specified within ~/.databrickscfg. Please check [connection profiles section](https://docs.databricks.com/aws/en/dev-tools/cli/profiles) for more details. If unspecified, the `DEFAULT` profile is used.
* `clientId` - (optional, environment variable `DATABRICKS_CLIENT_ID`) The `applicationId` of the Service Principal.
* `clientSecret` - (optional, environment variable `DATABRICKS_CLIENT_SECRET`) Secret of the service principal.
* `token` - (optional, environment variable `DATABRICKS_TOKEN`) The API token to authenticate into the workspace.
* `configFile` - (optional, environment variable `DATABRICKS_CONFIG_FILE`) Location of the Databricks CLI credentials file created by `databricks configure --token` command (~/.databrickscfg by default). Check [Databricks CLI documentation](https://docs.databricks.com/dev-tools/cli/index.html#set-up-authentication) for more details. The provider uses configuration file credentials when you don't specify host/token/azure attributes. This field defaults to `~/.databrickscfg`.
* `azureClientId` - (optional, environment variable `ARM_CLIENT_ID`) This is the Azure Enterprise Application (Service principal) client id. This service principal requires contributor access to your Azure Databricks deployment.
* `azureTenantId` - (optional, environment variable `ARM_TENANT_ID`) This is the Azure Active Directory Tenant id in which the Enterprise Application (Service Principal) resides.
* `azureEnvironment` - (optional, environment variable `ARM_ENVIRONMENT`) This is the Azure Environment which defaults to the `public` cloud. Other options are `german`, `china` and `usgovernment`.
* `azureUseMsi` - (optional, environment variable `ARM_USE_MSI`) Use Azure Managed Service Identity authentication.
* `googleCredentials` - (optional, environment variable `GOOGLE_CREDENTIALS`) A GCP Service Account Credentials JSON or the path to the file containing these credentials.
* `googleServiceAccount` - (optional, environment variable `DATABRICKS_GOOGLE_SERVICE_ACCOUNT`) The Google Cloud Platform (GCP) service account e-mail used for impersonation. Default Application Credentials must be configured, and the principal must be able to impersonate this service account.
* `authType` - (optional, environment variable `DATABRICKS_AUTH_TYPE`) enforce specific auth type to be used in very rare cases, where a single Pulumi state manages Databricks workspaces on more than one cloud and `more than one authorization method configured` error is a false positive. Valid values are `pat`, `basic` (deprecated), `oauth-m2m`, `databricks-cli`, `azure-client-secret`, `azure-msi`, `azure-cli`, `github-oidc-azure`,`env-oidc`, `file-oidc`,  `github-oidc`, `google-credentials`, and `google-id`.

The provider supports additional configuration parameters not related to authentication. They could be used when debugging problems, or do an additional tuning of provider's behavior:

* `httpTimeoutSeconds` - (optional) the amount of time Pulumi waits for a response from Databricks REST API. Default is *60*.
* `rateLimit` - (optional, environment variable `DATABRICKS_RATE_LIMIT`) defines maximum number of requests per second made to Databricks REST API by Pulumi. Default is *15*.
* `debugTruncateBytes` - (optional, environment variable `DATABRICKS_DEBUG_TRUNCATE_BYTES`) Applicable only when `TF_LOG=DEBUG` is set. Truncate JSON fields in HTTP requests and responses above this limit. Default is *96*.
* `debugHeaders` - (optional, environment variable `DATABRICKS_DEBUG_HEADERS`) Applicable only when `TF_LOG=DEBUG` is set. Debug HTTP headers of requests made by the provider. Default is *false*. We recommend turning this flag on only under exceptional circumstances, when troubleshooting authentication issues. Turning this flag on will log first `debugTruncateBytes` of any HTTP header value in cleartext.
* `skipVerify` - skips SSL certificate verification for HTTP calls. *Use at your own risk.* Default is *false* (don't skip verification).

!> **Warning** Sensitive credentials are printed to the log when `debugHeaders` is `true`. Use it for troubleshooting purposes only.
### `host` argument

The `host` argument configures the endpoint that the Pulumi Provider for Databricks interacts with. This must be configured according to the following table:

|              Environment               |                   `host`                    |
|----------------------------------------|---------------------------------------------|
| Databricks Account on AWS              | `https://accounts.cloud.databricks.com`     |
| Databricks Account on AWS GovCloud     | `https://accounts.cloud.databricks.us`      |
| Databricks Account on AWS GovCloud DOD | `https://accounts-dod.cloud.databricks.mil` |
| Azure Databricks Account               | `https://accounts.azuredatabricks.net`      |
| Azure Databricks Account (US Gov)      | `https://accounts.azuredatabricks.us`       |
| Azure Databricks Account (China)       | `https://accounts.azuredatabricks.cn`       |
| Databricks Account on GCP              | `https://accounts.gcp.databricks.com`       |
| Databricks Workspace (any cloud)       | `https://<workspace hostname>`              |
## Authentication

There are currently a number of supported methods to [authenticate](https://docs.databricks.com/dev-tools/api/latest/authentication.html) into the Databricks platform to create resources:

* (recommended for CI/CD) OpenID Connect
* (recommended for local development) Databricks CLI
* AWS, Azure and GCP via Databricks-managed Service Principals
* GCP via Google Cloud CLI
* Azure Active Directory Tokens via Azure CLI, Azure-managed Service Principals, or Managed Service Identities
* PAT Tokens
### Authenticating with GitHub OpenID Connect (OIDC)
The arguments `host` and `clientId` are used for the authentication which maps to the `github-oidc` authentication type.

These can be declared in the provider configuration or set in the environment variables `DATABRICKS_HOST` and `DATABRICKS_CLIENT_ID` respectively. Example:

Workspace level provider:

Configure the account-level provider as follows. Make sure to configure the account host as described above.
### Authenticating with Databricks CLI

The provider can authenticate using the Databricks CLI. After logging in with the `databricks auth login` command to your account or workspace, you only need to specify the name of the profile in your provider configuration. Pulumi will automatically read and reuse the cached OAuth token to interact with the Databricks REST API. See [the user-to-machine authentication guide](https://docs.databricks.com/aws/en/dev-tools/cli/authentication#oauth-user-to-machine-u2m-authentication) for more details.

You can specify a [CLI connection profile](https://docs.databricks.com/aws/en/dev-tools/cli/profiles) through `profile` parameter or `DATABRICKS_CONFIG_PROFILE` environment variable:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    databricks:profile:
        value: ML_WORKSPACE

```

You can specify non-standard location of configuration file through `configFile` parameter or `DATABRICKS_CONFIG_FILE` environment variable:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    databricks:configFile:
        value: /opt/databricks/cli-config

```
### Authenticating with Databricks-managed Service Principal

You can use the `clientId` + `clientSecret` attributes to authenticate with a Databricks-managed service principal at both the account and workspace levels in all supported clouds. The `clientId` is the `applicationId` of the Service Principal and `clientSecret` is its secret. You can generate the secret from Databricks Accounts Console (see [instruction](https://docs.databricks.com/dev-tools/authentication-oauth.html#step-2-create-an-oauth-secret-for-a-service-principal)) or by using the Pulumi resource databricks_service_principal_secret.

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    databricks:clientId:
        value: 'TODO: var.client_id'
    databricks:clientSecret:
        value: 'TODO: var.client_secret'
    databricks:host:
        value: https://abc-cdef-ghi.cloud.databricks.com

```

To create resources at both the account and workspace levels, you can create two providers as shown below:

Next, you can specify the corresponding provider when creating the resource. For example, you can use the workspace provider to create a workspace group

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as databricks from "@pulumi/databricks";

const clusterAdmin = new databricks.Group("cluster_admin", {
    displayName: "cluster_admin",
    allowClusterCreate: true,
    allowInstancePoolCreate: false,
});
```

{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_databricks as databricks

cluster_admin = databricks.Group("cluster_admin",
    display_name="cluster_admin",
    allow_cluster_create=True,
    allow_instance_pool_create=False)
```

{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Databricks = Pulumi.Databricks;

return await Deployment.RunAsync(() =>
{
    var clusterAdmin = new Databricks.Group("cluster_admin", new()
    {
        DisplayName = "cluster_admin",
        AllowClusterCreate = true,
        AllowInstancePoolCreate = false,
    });

});

```

{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-databricks/sdk/go/databricks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := databricks.NewGroup(ctx, "cluster_admin", &databricks.GroupArgs{
			DisplayName:             pulumi.String("cluster_admin"),
			AllowClusterCreate:      pulumi.Bool(true),
			AllowInstancePoolCreate: pulumi.Bool(false),
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
  clusterAdmin:
    type: databricks:Group
    name: cluster_admin
    properties:
      displayName: cluster_admin
      allowClusterCreate: true
      allowInstancePoolCreate: false
```

{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.databricks.Group;
import com.pulumi.databricks.GroupArgs;
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
        var clusterAdmin = new Group("clusterAdmin", GroupArgs.builder()
            .displayName("cluster_admin")
            .allowClusterCreate(true)
            .allowInstancePoolCreate(false)
            .build());

    }
}
```

{{% /choosable %}}
{{< /chooser >}}
### Authenticating with hostname and token

> Databricks strongly recommends using OAuth instead of PATs for user account client authentication and authorization due to the improved security OAuth has

You can use `host` and `token` parameters to supply credentials to the workspace. When environment variables are preferred, then you can specify `DATABRICKS_HOST` and `DATABRICKS_TOKEN` instead. Environment variables are the second most recommended way of configuring this provider.

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    databricks:host:
        value: https://<workspace hostname>
    databricks:token:
        value: dapitokenhere

```
### Authenticating with Workload Identity Federation (WIF)

Workload Identity Federation can be used to authenticate Databricks from automated workflows. This is done through the tokens issued by the automation environment. For more details on environment variables regarding the specific environments, please see: <https://docs.databricks.com/aws/en/dev-tools/auth/oauth-federation-provider>.

To create resources at both the account and workspace levels, you can create two providers as shown below:

Workspace level provider:

Account level provider:

Note: `authType` for Github Actions would be "github-oidc". For more details, please see the document linked above.
## Special configurations for Azure

The below Azure authentication options are supported at both the account and workspace levels. The provider works with [Azure CLI authentication](https://docs.microsoft.com/en-us/cli/azure/authenticate-azure-cli?view=azure-cli-latest) to facilitate local development workflows, though for automated scenarios, managed identity or service principal auth is recommended (and specification of `azureUseMsi`, `azureClientId`, `azureClientSecret` and `azureTenantId` parameters).
### Authenticating with Azure MSI

Since v0.3.8, it's possible to leverage Azure Managed Service Identity authentication, which is using the same environment variables as `azurerm` provider. Both `SystemAssigned` and `UserAssigned` identities work, as long as they have `Contributor` role on subscription level and created the workspace resource, or directly added to workspace through databricks_service_principal.

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    databricks:azureUseMsi:
        value: true
    databricks:azureWorkspaceResourceId:
        value: 'TODO: azurerm_databricks_workspace.this.id'
    databricks:host:
        value: 'TODO: data.azurerm_databricks_workspace.this.workspace_url'

```
### Authenticating with Azure-managed Service Principal using GitHub OpenID Connect (OIDC)

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    azure:clientId:
        value: 'TODO: var.client_id'
    azure:subscriptionId:
        value: 'TODO: var.subscription_id'
    azure:tenantId:
        value: 'TODO: var.tenant_id'
    azure:useOidc:
        value: true
    databricks:authType:
        value: github-oidc-azure
    databricks:azureClientId:
        value: 'TODO: var.client_id'
    databricks:azureTenantId:
        value: 'TODO: var.tenant_id'
    databricks:azureWorkspaceResourceId:
        value: 'TODO: azurerm_databricks_workspace.this.id'
    databricks:host:
        value: 'TODO: azurerm_databricks_workspace.this.workspace_url'

```

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as azure from "@pulumi/azure";
import * as databricks from "@pulumi/databricks";

const _this = new azure.databricks.Workspace("this", {
    location: "centralus",
    name: "my-workspace-name",
    resourceGroupName: resourceGroup,
    sku: "premium",
});
const my_user = new databricks.User("my-user", {userName: "test-user@databricks.com"});
```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    azure:clientId:
        value: 'TODO: var.client_id'
    azure:subscriptionId:
        value: 'TODO: var.subscription_id'
    azure:tenantId:
        value: 'TODO: var.tenant_id'
    azure:useOidc:
        value: true
    databricks:authType:
        value: github-oidc-azure
    databricks:azureClientId:
        value: 'TODO: var.client_id'
    databricks:azureTenantId:
        value: 'TODO: var.tenant_id'
    databricks:azureWorkspaceResourceId:
        value: 'TODO: azurerm_databricks_workspace.this.id'
    databricks:host:
        value: 'TODO: azurerm_databricks_workspace.this.workspace_url'

```

```python
import pulumi
import pulumi_azure as azure
import pulumi_databricks as databricks

this = azure.databricks.Workspace("this",
    location="centralus",
    name="my-workspace-name",
    resource_group_name=resource_group,
    sku="premium")
my_user = databricks.User("my-user", user_name="test-user@databricks.com")
```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    azure:clientId:
        value: 'TODO: var.client_id'
    azure:subscriptionId:
        value: 'TODO: var.subscription_id'
    azure:tenantId:
        value: 'TODO: var.tenant_id'
    azure:useOidc:
        value: true
    databricks:authType:
        value: github-oidc-azure
    databricks:azureClientId:
        value: 'TODO: var.client_id'
    databricks:azureTenantId:
        value: 'TODO: var.tenant_id'
    databricks:azureWorkspaceResourceId:
        value: 'TODO: azurerm_databricks_workspace.this.id'
    databricks:host:
        value: 'TODO: azurerm_databricks_workspace.this.workspace_url'

```

```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Azure = Pulumi.Azure;
using Databricks = Pulumi.Databricks;

return await Deployment.RunAsync(() =>
{
    var @this = new Azure.DataBricks.Workspace("this", new()
    {
        Location = "centralus",
        Name = "my-workspace-name",
        ResourceGroupName = resourceGroup,
        Sku = "premium",
    });

    var my_user = new Databricks.User("my-user", new()
    {
        UserName = "test-user@databricks.com",
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
    azure:clientId:
        value: 'TODO: var.client_id'
    azure:subscriptionId:
        value: 'TODO: var.subscription_id'
    azure:tenantId:
        value: 'TODO: var.tenant_id'
    azure:useOidc:
        value: true
    databricks:authType:
        value: github-oidc-azure
    databricks:azureClientId:
        value: 'TODO: var.client_id'
    databricks:azureTenantId:
        value: 'TODO: var.tenant_id'
    databricks:azureWorkspaceResourceId:
        value: 'TODO: azurerm_databricks_workspace.this.id'
    databricks:host:
        value: 'TODO: azurerm_databricks_workspace.this.workspace_url'

```

```go
package main

import (
	"github.com/pulumi/pulumi-azure/sdk/v6/go/azure/databricks"
	godatabricks "github.com/pulumi/pulumi-databricks/sdk/go/databricks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := databricks.NewWorkspace(ctx, "this", &databricks.WorkspaceArgs{
			Location:          pulumi.String("centralus"),
			Name:              pulumi.String("my-workspace-name"),
			ResourceGroupName: pulumi.Any(resourceGroup),
			Sku:               pulumi.String("premium"),
		})
		if err != nil {
			return err
		}
		_, err = godatabricks.NewUser(ctx, "my-user", &godatabricks.UserArgs{
			UserName: pulumi.String("test-user@databricks.com"),
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
    azure:clientId:
        value: 'TODO: var.client_id'
    azure:subscriptionId:
        value: 'TODO: var.subscription_id'
    azure:tenantId:
        value: 'TODO: var.tenant_id'
    azure:useOidc:
        value: true
    databricks:authType:
        value: github-oidc-azure
    databricks:azureClientId:
        value: 'TODO: var.client_id'
    databricks:azureTenantId:
        value: 'TODO: var.tenant_id'
    databricks:azureWorkspaceResourceId:
        value: 'TODO: azurerm_databricks_workspace.this.id'
    databricks:host:
        value: 'TODO: azurerm_databricks_workspace.this.workspace_url'

```

```yaml
resources:
  this:
    type: azure:databricks:Workspace
    properties:
      location: centralus
      name: my-workspace-name
      resourceGroupName: ${resourceGroup}
      sku: premium
  my-user:
    type: databricks:User
    properties:
      userName: test-user@databricks.com
```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    azure:clientId:
        value: 'TODO: var.client_id'
    azure:subscriptionId:
        value: 'TODO: var.subscription_id'
    azure:tenantId:
        value: 'TODO: var.tenant_id'
    azure:useOidc:
        value: true
    databricks:authType:
        value: github-oidc-azure
    databricks:azureClientId:
        value: 'TODO: var.client_id'
    databricks:azureTenantId:
        value: 'TODO: var.tenant_id'
    databricks:azureWorkspaceResourceId:
        value: 'TODO: azurerm_databricks_workspace.this.id'
    databricks:host:
        value: 'TODO: azurerm_databricks_workspace.this.workspace_url'

```

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.azure.databricks.Workspace;
import com.pulumi.azure.databricks.WorkspaceArgs;
import com.pulumi.databricks.User;
import com.pulumi.databricks.UserArgs;
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
        var this_ = new Workspace("this", WorkspaceArgs.builder()
            .location("centralus")
            .name("my-workspace-name")
            .resourceGroupName(resourceGroup)
            .sku("premium")
            .build());

        var my_user = new User("my-user", UserArgs.builder()
            .userName("test-user@databricks.com")
            .build());

    }
}
```

{{% /choosable %}}
{{< /chooser >}}

Follow the [Configuring OpenID Connect in Azure](https://docs.github.com/en/actions/security-for-github-actions/security-hardening-your-deployments/configuring-openid-connect-in-azure). You can then use the Azure service principal to authenticate in databricks.

There are `ARM_*` environment variables provide a way to share authentication configuration using the `databricks` provider alongside the `azurerm` provider.

When a workspace is created using a service principal account, that service principal account is automatically added to the workspace as a member of the admins group. To add a new service principal account to an existing workspace, create a databricks_service_principal.
### Authenticating with Azure CLI

It's possible to use [Azure CLI](https://docs.microsoft.com/cli/azure/) authentication, where the provider would rely on access token cached by `az login` command so that local development scenarios are possible. Technically, the provider will call `az account get-access-token` each time before an access token is about to expire.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    databricks:host:
        value: 'TODO: azurerm_databricks_workspace.this.workspace_url'

```

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as azure from "@pulumi/azure";
import * as databricks from "@pulumi/databricks";

const _this = new azure.databricks.Workspace("this", {
    location: "centralus",
    name: "my-workspace-name",
    resourceGroupName: resourceGroup,
    sku: "premium",
});
const my_user = new databricks.User("my-user", {
    userName: "test-user@databricks.com",
    displayName: "Test User",
});
```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    databricks:host:
        value: 'TODO: azurerm_databricks_workspace.this.workspace_url'

```

```python
import pulumi
import pulumi_azure as azure
import pulumi_databricks as databricks

this = azure.databricks.Workspace("this",
    location="centralus",
    name="my-workspace-name",
    resource_group_name=resource_group,
    sku="premium")
my_user = databricks.User("my-user",
    user_name="test-user@databricks.com",
    display_name="Test User")
```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    databricks:host:
        value: 'TODO: azurerm_databricks_workspace.this.workspace_url'

```

```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Azure = Pulumi.Azure;
using Databricks = Pulumi.Databricks;

return await Deployment.RunAsync(() =>
{
    var @this = new Azure.DataBricks.Workspace("this", new()
    {
        Location = "centralus",
        Name = "my-workspace-name",
        ResourceGroupName = resourceGroup,
        Sku = "premium",
    });

    var my_user = new Databricks.User("my-user", new()
    {
        UserName = "test-user@databricks.com",
        DisplayName = "Test User",
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
    databricks:host:
        value: 'TODO: azurerm_databricks_workspace.this.workspace_url'

```

```go
package main

import (
	"github.com/pulumi/pulumi-azure/sdk/v6/go/azure/databricks"
	godatabricks "github.com/pulumi/pulumi-databricks/sdk/go/databricks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := databricks.NewWorkspace(ctx, "this", &databricks.WorkspaceArgs{
			Location:          pulumi.String("centralus"),
			Name:              pulumi.String("my-workspace-name"),
			ResourceGroupName: pulumi.Any(resourceGroup),
			Sku:               pulumi.String("premium"),
		})
		if err != nil {
			return err
		}
		_, err = godatabricks.NewUser(ctx, "my-user", &godatabricks.UserArgs{
			UserName:    pulumi.String("test-user@databricks.com"),
			DisplayName: pulumi.String("Test User"),
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
    databricks:host:
        value: 'TODO: azurerm_databricks_workspace.this.workspace_url'

```

```yaml
resources:
  this:
    type: azure:databricks:Workspace
    properties:
      location: centralus
      name: my-workspace-name
      resourceGroupName: ${resourceGroup}
      sku: premium
  my-user:
    type: databricks:User
    properties:
      userName: test-user@databricks.com
      displayName: Test User
```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    databricks:host:
        value: 'TODO: azurerm_databricks_workspace.this.workspace_url'

```

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.azure.databricks.Workspace;
import com.pulumi.azure.databricks.WorkspaceArgs;
import com.pulumi.databricks.User;
import com.pulumi.databricks.UserArgs;
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
        var this_ = new Workspace("this", WorkspaceArgs.builder()
            .location("centralus")
            .name("my-workspace-name")
            .resourceGroupName(resourceGroup)
            .sku("premium")
            .build());

        var my_user = new User("my-user", UserArgs.builder()
            .userName("test-user@databricks.com")
            .displayName("Test User")
            .build());

    }
}
```

{{% /choosable %}}
{{< /chooser >}}
### Authenticating with Azure-managed Service Principal using Client Secret

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    azure:clientId:
        value: 'TODO: var.client_id'
    azure:clientSecret:
        value: 'TODO: var.client_secret'
    azure:subscriptionId:
        value: 'TODO: var.subscription_id'
    azure:tenantId:
        value: 'TODO: var.tenant_id'
    databricks:azureClientId:
        value: 'TODO: var.client_id'
    databricks:azureClientSecret:
        value: 'TODO: var.client_secret'
    databricks:azureTenantId:
        value: 'TODO: var.tenant_id'
    databricks:azureWorkspaceResourceId:
        value: 'TODO: azurerm_databricks_workspace.this.id'
    databricks:host:
        value: 'TODO: azurerm_databricks_workspace.this.workspace_url'

```

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as azure from "@pulumi/azure";
import * as databricks from "@pulumi/databricks";

const _this = new azure.databricks.Workspace("this", {
    location: "centralus",
    name: "my-workspace-name",
    resourceGroupName: resourceGroup,
    sku: "premium",
});
const my_user = new databricks.User("my-user", {userName: "test-user@databricks.com"});
```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    azure:clientId:
        value: 'TODO: var.client_id'
    azure:clientSecret:
        value: 'TODO: var.client_secret'
    azure:subscriptionId:
        value: 'TODO: var.subscription_id'
    azure:tenantId:
        value: 'TODO: var.tenant_id'
    databricks:azureClientId:
        value: 'TODO: var.client_id'
    databricks:azureClientSecret:
        value: 'TODO: var.client_secret'
    databricks:azureTenantId:
        value: 'TODO: var.tenant_id'
    databricks:azureWorkspaceResourceId:
        value: 'TODO: azurerm_databricks_workspace.this.id'
    databricks:host:
        value: 'TODO: azurerm_databricks_workspace.this.workspace_url'

```

```python
import pulumi
import pulumi_azure as azure
import pulumi_databricks as databricks

this = azure.databricks.Workspace("this",
    location="centralus",
    name="my-workspace-name",
    resource_group_name=resource_group,
    sku="premium")
my_user = databricks.User("my-user", user_name="test-user@databricks.com")
```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    azure:clientId:
        value: 'TODO: var.client_id'
    azure:clientSecret:
        value: 'TODO: var.client_secret'
    azure:subscriptionId:
        value: 'TODO: var.subscription_id'
    azure:tenantId:
        value: 'TODO: var.tenant_id'
    databricks:azureClientId:
        value: 'TODO: var.client_id'
    databricks:azureClientSecret:
        value: 'TODO: var.client_secret'
    databricks:azureTenantId:
        value: 'TODO: var.tenant_id'
    databricks:azureWorkspaceResourceId:
        value: 'TODO: azurerm_databricks_workspace.this.id'
    databricks:host:
        value: 'TODO: azurerm_databricks_workspace.this.workspace_url'

```

```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Azure = Pulumi.Azure;
using Databricks = Pulumi.Databricks;

return await Deployment.RunAsync(() =>
{
    var @this = new Azure.DataBricks.Workspace("this", new()
    {
        Location = "centralus",
        Name = "my-workspace-name",
        ResourceGroupName = resourceGroup,
        Sku = "premium",
    });

    var my_user = new Databricks.User("my-user", new()
    {
        UserName = "test-user@databricks.com",
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
    azure:clientId:
        value: 'TODO: var.client_id'
    azure:clientSecret:
        value: 'TODO: var.client_secret'
    azure:subscriptionId:
        value: 'TODO: var.subscription_id'
    azure:tenantId:
        value: 'TODO: var.tenant_id'
    databricks:azureClientId:
        value: 'TODO: var.client_id'
    databricks:azureClientSecret:
        value: 'TODO: var.client_secret'
    databricks:azureTenantId:
        value: 'TODO: var.tenant_id'
    databricks:azureWorkspaceResourceId:
        value: 'TODO: azurerm_databricks_workspace.this.id'
    databricks:host:
        value: 'TODO: azurerm_databricks_workspace.this.workspace_url'

```

```go
package main

import (
	"github.com/pulumi/pulumi-azure/sdk/v6/go/azure/databricks"
	godatabricks "github.com/pulumi/pulumi-databricks/sdk/go/databricks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := databricks.NewWorkspace(ctx, "this", &databricks.WorkspaceArgs{
			Location:          pulumi.String("centralus"),
			Name:              pulumi.String("my-workspace-name"),
			ResourceGroupName: pulumi.Any(resourceGroup),
			Sku:               pulumi.String("premium"),
		})
		if err != nil {
			return err
		}
		_, err = godatabricks.NewUser(ctx, "my-user", &godatabricks.UserArgs{
			UserName: pulumi.String("test-user@databricks.com"),
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
    azure:clientId:
        value: 'TODO: var.client_id'
    azure:clientSecret:
        value: 'TODO: var.client_secret'
    azure:subscriptionId:
        value: 'TODO: var.subscription_id'
    azure:tenantId:
        value: 'TODO: var.tenant_id'
    databricks:azureClientId:
        value: 'TODO: var.client_id'
    databricks:azureClientSecret:
        value: 'TODO: var.client_secret'
    databricks:azureTenantId:
        value: 'TODO: var.tenant_id'
    databricks:azureWorkspaceResourceId:
        value: 'TODO: azurerm_databricks_workspace.this.id'
    databricks:host:
        value: 'TODO: azurerm_databricks_workspace.this.workspace_url'

```

```yaml
resources:
  this:
    type: azure:databricks:Workspace
    properties:
      location: centralus
      name: my-workspace-name
      resourceGroupName: ${resourceGroup}
      sku: premium
  my-user:
    type: databricks:User
    properties:
      userName: test-user@databricks.com
```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    azure:clientId:
        value: 'TODO: var.client_id'
    azure:clientSecret:
        value: 'TODO: var.client_secret'
    azure:subscriptionId:
        value: 'TODO: var.subscription_id'
    azure:tenantId:
        value: 'TODO: var.tenant_id'
    databricks:azureClientId:
        value: 'TODO: var.client_id'
    databricks:azureClientSecret:
        value: 'TODO: var.client_secret'
    databricks:azureTenantId:
        value: 'TODO: var.tenant_id'
    databricks:azureWorkspaceResourceId:
        value: 'TODO: azurerm_databricks_workspace.this.id'
    databricks:host:
        value: 'TODO: azurerm_databricks_workspace.this.workspace_url'

```

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.azure.databricks.Workspace;
import com.pulumi.azure.databricks.WorkspaceArgs;
import com.pulumi.databricks.User;
import com.pulumi.databricks.UserArgs;
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
        var this_ = new Workspace("this", WorkspaceArgs.builder()
            .location("centralus")
            .name("my-workspace-name")
            .resourceGroupName(resourceGroup)
            .sku("premium")
            .build());

        var my_user = new User("my-user", UserArgs.builder()
            .userName("test-user@databricks.com")
            .build());

    }
}
```

{{% /choosable %}}
{{< /chooser >}}

There are `ARM_*` environment variables provide a way to share authentication configuration using the `databricks` provider alongside the `azurerm` provider.

When a workspace is created using a service principal account, that service principal account is automatically added to the workspace as a member of the admins group. To add a new service principal account to an existing workspace, create a databricks_service_principal.
## Special configurations for GCP

The provider works with [Google Cloud CLI authentication](https://cloud.google.com/sdk/docs/authorizing) to facilitate local development workflows. For automated scenarios, a service principal auth is necessary using `googleServiceAccount` parameter with [impersonation](https://cloud.google.com/docs/authentication#service-accounts) and Application Default Credentials. Alternatively, you could provide the service account key directly by passing it to `googleCredentials` parameter (or `GOOGLE_CREDENTIALS` environment variable)
## Special configuration for Unity Catalog

Except for metastore, metastore assignment and storage credential objects, Unity Catalog APIs are accessible via **workspace-level APIs**. This design may change in the future.

If you are configuring a new Databricks account for the first time, please create at least one workspace with an identity (user or service principal) that you intend to use for Unity Catalog rollout. You can then configure the provider using that identity and workspace to provision the required Unity Catalog resources.
## Special considerations for Unity Catalog Resources

When performing a single pulumi up to update both the owner and other fields for Unity Catalog resources, the process first updates the owner, followed by the other fields using the new owner's permissions. If your principal is not the owner (specifically, the newly updated owner), you will not have the authority to modify those fields. In cases where you wish to change the owner to another individual and also update other fields, we recommend initially updating the fields using your principal, which should have owner permissions, and then updating the owner in a separate step.