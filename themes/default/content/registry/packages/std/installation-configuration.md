---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-std/v2.3.2/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi/pulumi-std/blob/v2.3.2/docs/installation-configuration.md
title: Installation & Configuration
meta_desc: Information on how to install the Pulumi Standard Library provider.
layout: installation
---

## Installation

The Pulumi Standard Library provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/std`](https://www.npmjs.com/package/@pulumi/std)
* Python: [`pulumi-std`](https://pypi.org/project/pulumi-std/)
* Go: [`github.com/pulumi/pulumi-std/sdk/v2/go/std`](https://pkg.go.dev/github.com/pulumi/pulumi-std/sdk/v2/go/std)
* .NET: [`Pulumi.Std`](https://www.nuget.org/packages/Pulumi.Std)
* Java: [`com.pulumi/std`](https://central.sonatype.com/artifact/com.pulumi/std)

## Setup

The Standard Library provider does not require any configuration or credentials. It provides pure functions for string manipulation, mathematical operations, and other standard library functionality that work identically across all Pulumi languages.

### Installation

To use the Standard Library provider with your Pulumi program, install the appropriate package for your language:

{{< chooser language "typescript,python,go,csharp,java" / >}}

{{% choosable language typescript %}}

```bash
npm install @pulumi/std
```

{{% /choosable %}}

{{% choosable language python %}}

```bash
pip install pulumi-std
```

{{% /choosable %}}

{{% choosable language go %}}

```bash
go get github.com/pulumi/pulumi-std/sdk/v2/go/std
```

{{% /choosable %}}

{{% choosable language csharp %}}

```bash
dotnet add package Pulumi.Std
```

{{% /choosable %}}

{{% choosable language java %}}

```xml
<dependency>
    <groupId>com.pulumi</groupId>
    <artifactId>std</artifactId>
</dependency>
```

{{% /choosable %}}

## Configuration Options

The Standard Library provider does not require any configuration options or environment variables. All functions are self-contained and require no external setup.
