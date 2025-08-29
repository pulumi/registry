---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/dell/powerscale/1.8.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Powerscale Provider
meta_desc: Provides an overview on how to configure the Pulumi Powerscale provider.
layout: package
---

## Generate Provider

The Powerscale provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider dell/powerscale
```
## Overview

The Pulumi provider for Dell PowerScale can be used to interact with a Dell PowerScale array in order to manage the array resources.
## Example Usage
The following abridged example demonstrates the usage of the provider to create groupnet, subnet, network pool, ads provider, access zone, quota, snapshot, snapshot schedule, user, user group, filesystem ,nfs export and smb share.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    powerscale:endpoint:
        value: 'TODO: var.endpoint'
    powerscale:insecure:
        value: 'TODO: var.insecure'
    powerscale:password:
        value: 'TODO: var.password'
    powerscale:username:
        value: 'TODO: var.username'

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
    powerscale:endpoint:
        value: 'TODO: var.endpoint'
    powerscale:insecure:
        value: 'TODO: var.insecure'
    powerscale:password:
        value: 'TODO: var.password'
    powerscale:username:
        value: 'TODO: var.username'

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
    powerscale:endpoint:
        value: 'TODO: var.endpoint'
    powerscale:insecure:
        value: 'TODO: var.insecure'
    powerscale:password:
        value: 'TODO: var.password'
    powerscale:username:
        value: 'TODO: var.username'

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
    powerscale:endpoint:
        value: 'TODO: var.endpoint'
    powerscale:insecure:
        value: 'TODO: var.insecure'
    powerscale:password:
        value: 'TODO: var.password'
    powerscale:username:
        value: 'TODO: var.username'

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
    powerscale:endpoint:
        value: 'TODO: var.endpoint'
    powerscale:insecure:
        value: 'TODO: var.insecure'
    powerscale:password:
        value: 'TODO: var.password'
    powerscale:username:
        value: 'TODO: var.username'

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
    powerscale:endpoint:
        value: 'TODO: var.endpoint'
    powerscale:insecure:
        value: 'TODO: var.insecure'
    powerscale:password:
        value: 'TODO: var.password'
    powerscale:username:
        value: 'TODO: var.username'

```
```java
Example currently unavailable in this language
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

- `authType` (Number) what should be the auth type, 0 for basic and 1 for session-based. This can also be set using the environment variable POWERSCALE_AUTH_TYPE
- `endpoint` (String) The API endpoint, ex. https://172.17.177.230:8080. This can also be set using the environment variable POWERSCALE_ENDPOINT
- `insecure` (Boolean) whether to skip SSL validation. This can also be set using the environment variable POWERSCALE_INSECURE
- `password` (String, Sensitive) The password. This can also be set using the environment variable POWERSCALE_PASSWORD
- `timeout` (Number) specifies a time limit for requests. This can also be set using the environment variable POWERSCALE_TIMEOUT
- `username` (String) The username. This can also be set using the environment variable POWERSCALE_USERNAME
## Best Practices
1. The parent resource attributes of a certain resource (e.g. groupnet field of subnet resource) can only be designated
   at creation. Once designated, they cannot be modified except for parent resource renaming.
2. The name of a resource is modifiable, but it is necessary to make sure its name referenced in the child resources
   is also updated (can be done manually or use reference resource_id.name).
3. Resources with child resources cannot be deleted independently. Use pulumi destroy to delete all resources directly
   or delete all the child resources at the same time (depends_on is recommended to manage resources, serving as a
   precheck for delete operations).