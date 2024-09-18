---
title: Aiven Provider
meta_desc: Provides an overview on how to configure the Pulumi Aiven provider.
layout: package
---
## Installation

The aiven provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/aiven`](https://www.npmjs.com/package/@pulumi/aiven)
* Python: [`pulumi-aiven`](https://pypi.org/project/pulumi-aiven/)
* Go: [`github.com/pulumi/pulumi-aiven/sdk/v6/go/aiven`](https://github.com/pulumi/pulumi-aiven)
* .NET: [`Pulumi.Aiven`](https://www.nuget.org/packages/Pulumi.Aiven)
* Java: [`com.pulumi/aiven`](https://central.sonatype.com/artifact/com.pulumi/aiven)

The Pulumi provider for [Aiven](https://aiven.io/), the trusted open source data platform for everyone.
## Authentication
Sign up for Aiven and [create a personal token](https://aiven.io/docs/platform/howto/create_authentication_token).

You can also create an [application user](https://aiven.io/docs/platform/howto/manage-application-users) and use its token for accessing the Aiven Provider.
## Example usage


{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    aiven:apiToken:
        value: 'TODO: var.aiven_api_token'

```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    aiven:apiToken:
        value: 'TODO: var.aiven_api_token'

```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    aiven:apiToken:
        value: 'TODO: var.aiven_api_token'

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    aiven:apiToken:
        value: 'TODO: var.aiven_api_token'

```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    aiven:apiToken:
        value: 'TODO: var.aiven_api_token'

```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    aiven:apiToken:
        value: 'TODO: var.aiven_api_token'

```

{{% /choosable %}}
{{< /chooser >}}
## Environment variables

* For authentication, you can set the `AIVEN_TOKEN` to your token value.
* To use beta resources, set `PROVIDER_AIVEN_ENABLE_BETA` to any value.
* To allow IP filters to be purged, set `AIVEN_ALLOW_IP_FILTER_PURGE` to any value. This feature prevents accidental purging of IP filters, which can cause you to lose access to services.
## Resource options
The list of options in this document is not comprehensive. However, most map directly to the [Aiven REST API](https://api.aiven.io/doc/) properties.