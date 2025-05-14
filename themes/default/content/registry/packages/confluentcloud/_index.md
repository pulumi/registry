---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-confluentcloud/v2.28.0/docs/_index.md
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

{{% /choosable %}}
{{< /chooser >}}
#### Option #2: Manage a single Kafka cluster in the same Pulumi Stack
{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
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

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
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

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
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

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
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

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
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

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
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

{{% /choosable %}}
{{< /chooser >}}
#### Option #3: Manage a single Schema Registry cluster in the same Pulumi Stack
{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
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

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
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

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
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

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
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

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
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

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
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

{{% /choosable %}}
{{< /chooser >}}
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