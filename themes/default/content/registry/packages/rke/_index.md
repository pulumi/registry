---
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Rke Provider
meta_desc: Provides an overview on how to configure the Pulumi Rke provider.
layout: package
---
## Installation

The rke provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/rke`](https://www.npmjs.com/package/@pulumi/rke)
* Python: [`pulumi-rke`](https://pypi.org/project/pulumi-rke/)
* Go: [`github.com/pulumi/pulumi-rke/sdk/v3/go/rke`](https://github.com/pulumi/pulumi-rke)
* .NET: [`Pulumi.Rke`](https://www.nuget.org/packages/Pulumi.Rke)
* Java: [`com.pulumi/rke`](https://central.sonatype.com/artifact/com.pulumi/rke)
## Overview

The RKE provider is used to interact with Rancher Kubernetes Engine kubernetes clusters.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    rke:debug:
        value: true
    rke:logFile:
        value: <rke_log_file>

```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    rke:debug:
        value: true
    rke:logFile:
        value: <rke_log_file>

```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    rke:debug:
        value: true
    rke:logFile:
        value: <rke_log_file>

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    rke:debug:
        value: true
    rke:logFile:
        value: <rke_log_file>

```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    rke:debug:
        value: true
    rke:logFile:
        value: <rke_log_file>

```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    rke:debug:
        value: true
    rke:logFile:
        value: <rke_log_file>

```

{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

The following configuration inputs are supported:

* `debug` - (Optional) Enable RKE debug logs. It can also be sourced from the `RKE_DEBUG` environment variable. Default `false` (bool)
* `logFile` - (Optional) Save RKE logs to a file. It can also be sourced from the `RKE_LOG_FILE` environment variable (string)