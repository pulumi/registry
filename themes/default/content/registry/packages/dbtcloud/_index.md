---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-dbtcloud/v1.0.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: dbt Cloud Provider
meta_desc: Provides an overview on how to configure the Pulumi dbt Cloud provider.
layout: package
---
## Installation

The dbt Cloud provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/dbtcloud`](https://www.npmjs.com/package/@pulumi/dbtcloud)
* Python: [`pulumi-dbtcloud`](https://pypi.org/project/pulumi-dbtcloud/)
* Go: [`github.com/pulumi/pulumi-dbtcloud/sdk/go/dbtcloud`](https://github.com/pulumi/pulumi-dbtcloud)
* .NET: [`Pulumi.Dbtcloud`](https://www.nuget.org/packages/Pulumi.Dbtcloud)
* Java: [`com.pulumi/dbtcloud`](https://central.sonatype.com/artifact/com.pulumi/dbtcloud)
## Provider configuration

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    dbtcloud:accountId:
        value: 'TODO: var.dbt_cloud_account_id'
    dbtcloud:hostUrl:
        value: https://cloud.getdbt.com/api
    dbtcloud:token:
        value: 'TODO: var.dbt_cloud_token'

```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    dbtcloud:accountId:
        value: 'TODO: var.dbt_cloud_account_id'
    dbtcloud:hostUrl:
        value: https://cloud.getdbt.com/api
    dbtcloud:token:
        value: 'TODO: var.dbt_cloud_token'

```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    dbtcloud:accountId:
        value: 'TODO: var.dbt_cloud_account_id'
    dbtcloud:hostUrl:
        value: https://cloud.getdbt.com/api
    dbtcloud:token:
        value: 'TODO: var.dbt_cloud_token'

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    dbtcloud:accountId:
        value: 'TODO: var.dbt_cloud_account_id'
    dbtcloud:hostUrl:
        value: https://cloud.getdbt.com/api
    dbtcloud:token:
        value: 'TODO: var.dbt_cloud_token'

```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    dbtcloud:accountId:
        value: 'TODO: var.dbt_cloud_account_id'
    dbtcloud:hostUrl:
        value: https://cloud.getdbt.com/api
    dbtcloud:token:
        value: 'TODO: var.dbt_cloud_token'

```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    dbtcloud:accountId:
        value: 'TODO: var.dbt_cloud_account_id'
    dbtcloud:hostUrl:
        value: https://cloud.getdbt.com/api
    dbtcloud:token:
        value: 'TODO: var.dbt_cloud_token'

```

{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

- `accountId` (Number) Account identifier for your dbt Cloud implementation. Instead of setting the parameter, you can set the environment variable `DBT_CLOUD_ACCOUNT_ID`
- `hostUrl` (String) URL for your dbt Cloud deployment. Instead of setting the parameter, you can set the environment variable `DBT_CLOUD_HOST_URL` - Defaults to <https://cloud.getdbt.com/api>
- `token` (String, Sensitive) API token for your dbt Cloud. Instead of setting the parameter, you can set the environment variable `DBT_CLOUD_TOKEN`