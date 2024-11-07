---
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Civo Provider
meta_desc: Provides an overview on how to configure the Pulumi Civo provider.
layout: package
---
## Installation

The civo provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/civo`](https://www.npmjs.com/package/@pulumi/civo)
* Python: [`pulumi-civo`](https://pypi.org/project/pulumi-civo/)
* Go: [`github.com/pulumi/pulumi-civo/sdk/v2/go/civo`](https://github.com/pulumi/pulumi-civo)
* .NET: [`Pulumi.Civo`](https://www.nuget.org/packages/Pulumi.Civo)
* Java: [`com.pulumi/civo`](https://central.sonatype.com/artifact/com.pulumi/civo)
## Overview

The Civo provider is used to interact with the resources supported by Civo.

Use the navigation to the left to read about the available resources.
## Configuration

The provider will use the credentials of the [Civo CLI](https://github.com/civo/cli) (stored in ` ~/.civo.json`) if no other credentials have been set up. The provider will use credentials in the following order:

1. Environment variable (`CIVO_TOKEN`).
2. Token provided via a credentials file (See `credentialsFile` input below)
3. [CLI](https://github.com/civo/cli) configuration (`~/.civo.json`)

That means that if the `CIVO_TOKEN` variable is set, all other credentials will be ignored, and if the `credentialsFile` is set, that will be used over the CLI credentials.
### Obtaining a token

First you will need to create a [Civo Account](https://dashboard.civo.com/signup) and then you can do the following:

* If you can want to configure your credentials with the CLI, instructions are [here](https://www.civo.com/docs/overview/civo-cli#add-an-api-key-to-civo-cli)
* To fetch an API key go to the [security section](https://dashboard.civo.com/security) on the dashboard
### Using the CIVO_TOKEN variable

To use the Civo token, export the variable containing your token:

```bash
export CIVO_TOKEN=<your token>
```
### Using a credentials file

The format of the credentials file is as follows:

```json
{
	"apikeys": {
		"tf_key": "write-your-token-here"
	},
	"meta": {
		"current_apikey": "tf_key"
	}
}
```

you will then need to configure the `credentialsFile` input to the correct location, for example:

`credentialsFile = "/secure/path/civo.json"`
### Using the CLI

If you install the CLI and [configure a token](https://www.civo.com/docs/overview/civo-cli#add-an-api-key-to-civo-cli), there is nothing else you need to do if those are the credentials you wish to use, ideal for local usage.
## Example Usage
### Simplest usage

In this example the provider will look for credentials set by the CLI and use LON1 region as default, or use the environment variable `CIVO_TOKEN` if set.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    civo:region:
        value: LON1

```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    civo:region:
        value: LON1

```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    civo:region:
        value: LON1

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    civo:region:
        value: LON1

```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    civo:region:
        value: LON1

```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    civo:region:
        value: LON1

```

{{% /choosable %}}
{{< /chooser >}}
### Example with credentials file

In this example we are providing a specific version of the pulumi provider and setting a credentials file to use. The credentials file will be used if the environment variable `CIVO_TOKEN` is not set.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    civo:credentialsFile:
        value: /secure/path/civo.json
    civo:region:
        value: LON1

```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    civo:credentialsFile:
        value: /secure/path/civo.json
    civo:region:
        value: LON1

```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    civo:credentialsFile:
        value: /secure/path/civo.json
    civo:region:
        value: LON1

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    civo:credentialsFile:
        value: /secure/path/civo.json
    civo:region:
        value: LON1

```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    civo:credentialsFile:
        value: /secure/path/civo.json
    civo:region:
        value: LON1

```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    civo:credentialsFile:
        value: /secure/path/civo.json
    civo:region:
        value: LON1

```

{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

- `apiEndpoint` (String) The Base URL to use for CIVO API.
- `region` (String) This sets the default region for all resources. If no default region is set, you will need to specify individually in every resource.
  <a id="credentialsFile"></a>
- `credentialsFile` (string) specify a location for a file containing your civo credentials token
- `token` (String) (**Deprecated**) for legacy reasons the user can still specify the token as an input, but in order to avoid storing that in pulumi state we have deprecated this and will be remove in future versions - don't use it.