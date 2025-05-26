---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/coder/coder/2.5.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Coder Provider
meta_desc: Provides an overview on how to configure the Pulumi Coder provider.
layout: package
---

## Generate Provider

The Coder provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider coder/coder
```
## Overview

Pulumi provider for managing Coder [templates](https://coder.com/docs/admin/templates), which are the underlying infrastructure for Coder [workspaces](https://coder.com/docs/user-guides/workspace-management).

> Requires Coder [v2.18.0](https://github.com/coder/coder/releases/tag/v2.18.0) or later.

!> `coderGitAuth` and owner related fields of `coder.getWorkspace` function have been removed. Follow the Version 2 Upgrade Guide to update your code.
## Example

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    gcp:region:
        value: us-central1

```
```typescript
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    gcp:region:
        value: us-central1

```
```python
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    gcp:region:
        value: us-central1

```
```csharp
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    gcp:region:
        value: us-central1

```
```go
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    gcp:region:
        value: us-central1

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
config:
    gcp:region:
        value: us-central1

```
```java
Example currently unavailable in this language
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

- `url` (String) The URL to access Coder.