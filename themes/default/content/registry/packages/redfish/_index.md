---
# WARNING: this file was fetched from https://d2dcabd1cce183.cloudfront.net/docs/registry.opentofu.org/dell/redfish/1.5.0/index.md
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
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as redfish from "@pulumi/redfish";

const rr: redfish.UserAccount[] = [];
for (const range = {value: 0}; range.value < rack1; range.value++) {
    rr.push(new redfish.UserAccount(`rr-${range.value}`, {
        redfishServers: [{
            user: range.value.user,
            password: range.value.password,
            endpoint: range.value.endpoint,
            sslInsecure: range.value.sslInsecure,
        }],
        userId: "4",
        username: "test",
        password: "Test@123",
        roleId: "Operator",
        enabled: true,
    }));
}
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_redfish as redfish

rr = []
for range in [{"value": i} for i in range(0, rack1)]:
    rr.append(redfish.UserAccount(f"rr-{range['value']}",
        redfish_servers=[{
            "user": range["value"]["user"],
            "password": range["value"]["password"],
            "endpoint": range["value"]["endpoint"],
            "ssl_insecure": range["value"]["sslInsecure"],
        }],
        user_id="4",
        username="test",
        password="Test@123",
        role_id="Operator",
        enabled=True))
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Redfish = Pulumi.Redfish;

return await Deployment.RunAsync(() =>
{
    var rr = new List<Redfish.UserAccount>();
    for (var rangeIndex = 0; rangeIndex < rack1; rangeIndex++)
    {
        var range = new { Value = rangeIndex };
        rr.Add(new Redfish.UserAccount($"rr-{range.Value}", new()
        {
            RedfishServers = new[]
            {
                new Redfish.Inputs.UserAccountRedfishServerArgs
                {
                    User = range.Value.User,
                    Password = range.Value.Password,
                    Endpoint = range.Value.Endpoint,
                    SslInsecure = range.Value.SslInsecure,
                },
            },
            UserId = "4",
            Username = "test",
            Password = "Test@123",
            RoleId = "Operator",
            Enabled = true,
        }));
    }
});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/redfish/redfish"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		var rr []*redfish.UserAccount
		for index := 0; index < rack1; index++ {
			key0 := index
			val0 := index
			__res, err := redfish.NewUserAccount(ctx, fmt.Sprintf("rr-%v", key0), &redfish.UserAccountArgs{
				RedfishServers: redfish.UserAccountRedfishServerArray{
					&redfish.UserAccountRedfishServerArgs{
						User:        pulumi.Any(val0),
						Password:    pulumi.Any(val0),
						Endpoint:    pulumi.Any(val0),
						SslInsecure: pulumi.Any(val0),
					},
				},
				UserId:   pulumi.String("4"),
				Username: pulumi.String("test"),
				Password: pulumi.String("Test@123"),
				RoleId:   pulumi.String("Operator"),
				Enabled:  pulumi.Bool(true),
			})
			if err != nil {
				return err
			}
			rr = append(rr, __res)
		}
		return nil
	})
}
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
resources:
  rr:
    type: redfish:UserAccount
    properties:
      redfishServers:
        - user: ${range.value.user}
          password: ${range.value.password}
          endpoint: ${range.value.endpoint}
          sslInsecure: ${range.value.sslInsecure}
      userId: '4'
      username: test
      password: Test@123
      roleId: Operator
      enabled: true
    options: {}
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.redfish.UserAccount;
import com.pulumi.redfish.UserAccountArgs;
import com.pulumi.redfish.inputs.UserAccountRedfishServerArgs;
import com.pulumi.codegen.internal.KeyedValue;
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
        for (var i = 0; i < rack1; i++) {
            new UserAccount("rr-" + i, UserAccountArgs.builder()
                .redfishServers(UserAccountRedfishServerArgs.builder()
                    .user(range.value().user())
                    .password(range.value().password())
                    .endpoint(range.value().endpoint())
                    .sslInsecure(range.value().sslInsecure())
                    .build())
                .userId("4")
                .username("test")
                .password("Test@123")
                .roleId("Operator")
                .enabled(true)
                .build());


}
    }
}
```
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
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs

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

variables.tf
{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";

const config = new pulumi.Config();
const rack1 = config.requireObject<Record<string, {endpoint?: string, password?: string, sslInsecure?: boolean, user?: string}>>("rack1");
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi

config = pulumi.Config()
rack1 = config.require_object("rack1")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;

return await Deployment.RunAsync(() =>
{
    var config = new Config();
    var rack1 = config.RequireObject<Dictionary<string, Rack1>>("rack1");
});

public class Rack1
{
    public string endpoint { get; set; }
    public string password { get; set; }
    public bool sslInsecure { get; set; }
    public string user { get; set; }
}

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		rack1 := cfg.Require("rack1")
		return nil
	})
}
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
configuration:
  # /*
  # Copyright (c) 2022-2024 Dell Inc., or its subsidiaries. All Rights Reserved.

  # Licensed under the Mozilla Public License Version 2.0 (the "License");
  # you may not use this file except in compliance with the License.
  # You may obtain a copy of the License at

  #     http://mozilla.org/MPL/2.0/


  # Unless required by applicable law or agreed to in writing, software
  # distributed under the License is distributed on an "AS IS" BASIS,
  # WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  # See the License for the specific language governing permissions and
  # limitations under the License.
  # */
  rack1:
    type: map(object({endpoint = union(none, string), password = union(none, string), sslInsecure = union(bool, none), user = union(none, string)}))
```
{{% /choosable %}}
{{% choosable language java %}}
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
        final var config = ctx.config();
        final var rack1 = config.get("rack1");
    }
}
```
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