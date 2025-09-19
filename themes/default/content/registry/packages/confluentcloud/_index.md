---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-confluentcloud/v2.39.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Confluent Provider
meta_desc: Provides an overview on how to configure the Pulumi Confluent provider.
layout: package
---

## Installation

The Confluent provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/confluentcloud`](https://www.npmjs.com/package/@pulumi/confluentcloud)
* Python: [`pulumi-confluentcloud`](https://pypi.org/project/pulumi-confluentcloud/)
* Go: [`github.com/pulumi/pulumi-confluentcloud/sdk/v2/go/confluentcloud`](https://github.com/pulumi/pulumi-confluentcloud)
* .NET: [`Pulumi.Confluentcloud`](https://www.nuget.org/packages/Pulumi.Confluentcloud)
* Java: [`com.pulumi/confluentcloud`](https://central.sonatype.com/artifact/com.pulumi/confluentcloud)

## Overview

Simplify Apache Kafka Pulumi deployment with the Confluent Pulumi Provider. Manage Environments, Kafka Clusters, Kafka Topics, Kafka ACLs, Service Accounts, and more in Confluent.

Use the Confluent provider to deploy and manage [Confluent Cloud](https://www.confluent.io/confluent-cloud/) infrastructure. You must provide appropriate credentials to use the provider.
## Example Usage


##### Option #1: Manage multiple clusters in the same Pulumi Stack
{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    confluent:cloudApiKey:
        value: 'TODO: var.confluent_cloud_api_key'
    confluent:cloudApiSecret:
        value: 'TODO: var.confluent_cloud_api_secret'

```
```typescript
import * as pulumi from "@pulumi/pulumi";

```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    confluent:cloudApiKey:
        value: 'TODO: var.confluent_cloud_api_key'
    confluent:cloudApiSecret:
        value: 'TODO: var.confluent_cloud_api_secret'

```
```python
import pulumi

```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    confluent:cloudApiKey:
        value: 'TODO: var.confluent_cloud_api_key'
    confluent:cloudApiSecret:
        value: 'TODO: var.confluent_cloud_api_secret'

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;

return await Deployment.RunAsync(() =>
{
});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    confluent:cloudApiKey:
        value: 'TODO: var.confluent_cloud_api_key'
    confluent:cloudApiSecret:
        value: 'TODO: var.confluent_cloud_api_secret'

```
```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
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
    confluent:cloudApiKey:
        value: 'TODO: var.confluent_cloud_api_key'
    confluent:cloudApiSecret:
        value: 'TODO: var.confluent_cloud_api_secret'

```
```yaml
{}
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    confluent:cloudApiKey:
        value: 'TODO: var.confluent_cloud_api_key'
    confluent:cloudApiSecret:
        value: 'TODO: var.confluent_cloud_api_secret'

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
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
    }
}
```
{{% /choosable %}}
{{< /chooser >}}
#### Option #2: Manage a single Kafka cluster in the same Pulumi Stack
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    confluent:kafkaApiKey:
        value: 'TODO: var.kafka_api_key'
    confluent:kafkaApiSecret:
        value: 'TODO: var.kafka_api_secret'
    confluent:kafkaId:
        value: 'TODO: var.kafka_id'
    confluent:kafkaRestEndpoint:
        value: 'TODO: var.kafka_rest_endpoint'

```
#### Option #3: Manage a single Schema Registry cluster in the same Pulumi Stack
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    confluent:schemaRegistryApiKey:
        value: 'TODO: var.schema_registry_api_key'
    confluent:schemaRegistryApiSecret:
        value: 'TODO: var.schema_registry_api_secret'
    confluent:schemaRegistryId:
        value: 'TODO: var.schema_registry_id'
    confluent:schemaRegistryRestEndpoint:
        value: 'TODO: var.schema_registry_rest_endpoint'

```
## Enable Confluent Cloud Access

Confluent Cloud requires API keys to manage access and authentication to different parts of the service. An API key consists of a key and a secret. You can create and manage API keys by using either the [Confluent Cloud CLI](https://docs.confluent.io/ccloud-cli/current/index.html) or the [Confluent Cloud Console](https://confluent.cloud/). Learn more about Confluent Cloud API Key access [here](https://docs.confluent.io/cloud/current/client-apps/api-keys.html#ccloud-api-keys).
## Provider Authentication

Confluent Pulumi provider allows authentication by using environment variables or static credentials.
### Environment Variables

Run the following commands to set the `CONFLUENT_CLOUD_API_KEY` and `CONFLUENT_CLOUD_API_SECRET` environment variables:

```shell
# Option #1: Manage multiple clusters in the same Pulumi Stack
$ export CONFLUENT_CLOUD_API_KEY="<cloud_api_key>"
$ export CONFLUENT_CLOUD_API_SECRET="<cloud_api_secret>"

# Option #2: Manage a single Kafka cluster in the same Pulumi Stack
$ export CONFLUENT_CLOUD_API_KEY="<cloud_api_key>"
$ export CONFLUENT_CLOUD_API_SECRET="<cloud_api_secret>"
$ export KAFKA_ID="<kafka_id>"
$ export KAFKA_REST_ENDPOINT="<kafka_rest_endpoint>"
$ export KAFKA_API_KEY="<kafka_api_key>"
$ export KAFKA_API_SECRET="<kafka_api_secret>"

# Option #2: Manage a single Schema Registry cluster in the same Pulumi Stack
$ export SCHEMA_REGISTRY_ID="<schema_registry_id>"
$ export SCHEMA_REGISTRY_REST_ENDPOINT="<schema_registry_rest_endpoint>"
$ export SCHEMA_REGISTRY_API_KEY="<schema_registry_api_key>"
$ export SCHEMA_REGISTRY_API_SECRET="<schema_registry_api_secret>"
```

> **Note:** Quotation marks are required around the API key and secret strings.
### OAuth Credentials

> **Note:** Authentication using the `oauth` credentials block is available in **General Availability** for all customers. However, it's recommended to try OAuth features in non-production environments before rolling out to production.

Confluent Pulumi provider allows authentication by using OAuth credentials. You can use the `oauth` block to configure the provider with OAuth credentials.

```shell
# Option #1: Provide Identity Provider client id and secret, token retrieval URL, and the established Identity Pool ID
provider "confluent" {
  oauth {
    oauth_external_token_url = var.oauth_external_token_url            # the URL to retrieve the token from your Identity Provider, such as "https://mycompany.okta.com/oauth2/abc123/v1/token"
    oauth_external_client_id  = var.oauth_external_client_id           # the client id of your Identity Provider authorization server
    oauth_external_client_secret = var.oauth_external_client_secret    # the client secret of your Identity Provider authorization server
    oauth_identity_pool_id = var.oauth_identity_pool_id                # the established Identity Pool ID on Confluent Cloud based on your Identity Provider
    oauth_external_token_scope = var.oauth_external_token_scope        # the application client scope, might be required by your Identity Provider, such as Microsoft Azure Entra ID requires "api://<client_id>/.default" scope
  }
}
# Token refresh capability is supported by Confluent Provider for Option #1.

# Option #2: Provide a static token from the Identity Provider the established Identity Pool ID
provider "confluent" {
  oauth {
    oauth_external_access_token = var.oauth_external_access_token    # the static access token from your Identity Provider, please ensure it is not expired
    oauth_identity_pool_id = var.oauth_identity_pool_id              # the established Identity Pool ID on Confluent Cloud based on your Identity Provider
  }
}
# Token refresh capability is NOT supported by Confluent Provider for Option #2.
```
Complete examples (with Okta and Microsoft Azure Entra ID as identity provider) for using OAuth credentials with the Confluent Pulumi Provider can be found here.

> **Note:** You still need `cloudApiKey` and `cloudApiSecret` to manage below Confluent Cloud resources/data-sources as they are not supported with OAuth credentials yet:
* `confluentcloud.ApiKey` resource.
* `confluentcloud.CatalogIntegration` resource and data-source.
* `confluentcloud.CustomConnectorPlugin` resource.
* `confluentcloud.FlinkArtifact` resource.
* `confluentcloud.TableflowTopic` resource and data-source.

> **Note:** An Identity Provider must be set up first on Confluent Cloud before using the OAuth credentials for Pulumi Provider. You can find more information about Identity Provider setting up [here](https://docs.confluent.io/cloud/current/security/authenticate/workload-identities/identity-providers/oauth/identity-providers.htmll).

> **Note:** After Identity Provider is set up, an Identity Pool must be added and assigned proper RBAC roles to manage Confluent Cloud resources/data-sources with corresponding scope, more details can be found [here](https://docs.confluent.io/cloud/current/security/authenticate/workload-identities/identity-providers/oauth/identity-pools.html).

> **Note:** `oauthExternalTokenScope` could be optional or required based on your Identity Provider. For example, Microsoft Azure Entra ID requires `api://<client_id>/.default` scope to retrieve the token, while Okta does not require any scope.

> **Note:** To switch your Pulumi configuration from API key/secret authentication to OAuth, update your provider configuration by removing any references to variables such as `cloudApiKey`, `flinkApiKey`, `kafkaApiKey`, `schemaRegistryApiKey`, and similar variables. Also, remove any `credentials` blocks from resources like `confluentcloud.KafkaTopic`, `confluentcloud.Schema`, and `confluentcloud.FlinkStatement`. Instead, specify your authentication details within the `oauth {}` block. After making these changes, apply your configuration to start using OAuth.

!> **Warning:** Without proper Identity Provider setup, Identity Pool creation and RBAC roles assignment, the OAuth credentials will not work with Confluent Pulumi Provider.