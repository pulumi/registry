---
# WARNING: this file was fetched from https://d3b41xu1gacj9g.cloudfront.net/docs/registry.opentofu.org/dell/redfish/1.5.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Redfish Provider
meta_desc: Provides an overview on how to configure the Pulumi Redfish provider.
layout: package
---

## Generate Provider

The Redfish provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider dell/redfish
```
## Overview

Pulumi Provider Redfish
# Information about Redfish Pulumi Provider
This guide will explain different parts of the provider and will give an overview about how the provider is built.
## 1. Provider's way of operation
When you think of Pulumi, normally operators tend to think that the way a provider connects with a cloud provider is via a single endpoint. Well, actually that's the way it works. Cloud providers provide an endpoint and operators point to that endpoint when configuring pulumi.
```
  +-----------------+
  | Cloud provider  |
  +-------+---------+
          ^
          |
          |
+---------+----------+
| Pulumi provider |
+--------------------+
```

With the **Redfish Pulumi Provider**, that operating model has been changed because of the way the provider interacts with the infrastructure (Redfish endpoints).
In a regular scenario (for instance a datacenter), operators don't just have one endpoint, but a bunch of them. Each redfish endpoint corresponds to each physical server.
```
+------------------+     +------------------+      ...N        +------------------+
| PowerEdge Server |     | PowerEdge Server |  +-----------+   | PowerEdge Server |
| with Redfish API |     | with Redfish API |                  | with Redfish API |
+-------+----------+     +--------+---------+                  +---------+--------+
        ^                         ^                   ^                  ^
        |                         |                   |                  |
        +-------------------------+----------+--------+------------------+
                                             |
                                             |
                                  +----------+---------+
                                  | Pulumi Provider |
                                  +--------------------+

```
## How we overcome this

Normally the provider is initialized in the provider configuration, giving it your cloud credentials to deal with the infrastructure. Something like this:
```
provider "aws" {
	region     = "eu-west-1"
	access_key = "myaccesskey"
	secret_key = "mysecretkey"
}
```
When that is done, then operators would start writing the resources they want to deploy in those regions.

With this **pulumi redfish provider** a different approach had to be followed since there are multiple endpoints. What has been done (and kudos to Kyriakos Oikonomakos from Pulumi for proposing this) was to initialize the client at the resource level. This allows operators to manage different servers from one central point. Take a look into this example:

users.tf
{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}

{{% /choosable %}}
{{% choosable language python %}}

{{% /choosable %}}
{{% choosable language csharp %}}

{{% /choosable %}}
{{% choosable language go %}}

{{% /choosable %}}
{{% choosable language yaml %}}

{{% /choosable %}}
{{% choosable language java %}}

{{% /choosable %}}
{{< /chooser >}}

pulumi.tfvars
```
rack1 = {
  "my-server-1" = {
    user         = "admin"
    password     = "passw0rd"
    endpoint     = "https://my-server-1.myawesomecompany.org"
    ssl_insecure = true
  },
  "my-server-2" = {
    user         = "admin"
    password     = "passw0rd"
    endpoint     = "https://my-server-2.myawesomecompany.org"
    ssl_insecure = true
  },
  "my-server-3" = {
    user         = "admin"
    password     = "passw0rd"
    endpoint     = "https:://[2001:db8:a::123]"
    ssl_insecure = true
  },
}
```

By doing this, operators create two users on two different servers using this provider and the Redfish API.
*Remember, in every CRUD operation, the client must be initialized.*
## Overwriting client credentials
There might be scenarios where operators have the same credentials for all machines they want to manage. In that case they don't need to repeatedly write the *user* and *password* for all servers. They can write their credentials at the provider configuration level.
```
provider "redfish" {
    user = "root"
    password = "passw0rd"
}
```

After the user specifies their credentials, they will next need to define the infrastructure. Instead of defining credentials for each endpoint they need only provide the *endpoint* and *ssl_insecure* values:

```
rack1 = {
    "my-server-1" = {
        endpoint = "https://my-server-1.myawesomecompany.org"
        ssl_insecure = true
    },
    "my-server-2" = {
        endpoint = "https://my-server-2.myawesomecompany.org"
        ssl_insecure = true
    },
}
```

Pulumi will always use the most specific client values. In the case client credentials are defined at both the provider configuration and resource level, **the credentials defined at the resource level** will be used.
## Example Usage

provider.tf
{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}

{{% /choosable %}}
{{% choosable language python %}}

{{% /choosable %}}
{{% choosable language csharp %}}

{{% /choosable %}}
{{% choosable language go %}}

{{% /choosable %}}
{{% choosable language yaml %}}

{{% /choosable %}}
{{% choosable language java %}}

{{% /choosable %}}
{{< /chooser >}}

variables.tf
{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}

{{% /choosable %}}
{{% choosable language python %}}

{{% /choosable %}}
{{% choosable language csharp %}}

{{% /choosable %}}
{{% choosable language go %}}

{{% /choosable %}}
{{% choosable language yaml %}}

{{% /choosable %}}
{{% choosable language java %}}

{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

- `password` (String, Sensitive) This field is the password related to the user given
- `redfishServers` (Attributes Map) Map of server BMCs with their alias keys and respective user credentials. This is required when resource/function's `redfishAlias` is not null (see below for nested schema)
- `user` (String) This field is the user to login against the redfish API

<a id="nestedatt--redfish_servers"></a>
### Nested Schema for `redfishServers`

Required:

- `endpoint` (String) Server BMC IP address or hostname

Optional:

- `password` (String, Sensitive) User password for login
- `sslInsecure` (Boolean) This field indicates whether the SSL/TLS certificate must be verified or not
- `user` (String) User name for login