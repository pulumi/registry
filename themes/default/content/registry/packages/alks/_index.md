---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/cox-automotive/alks/2.9.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Alks Provider
meta_desc: Provides an overview on how to configure the Pulumi Alks provider.
layout: package
---

## Generate Provider

The Alks provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider cox-automotive/alks
```
## Overview

The ALKS Pulumi Provider is used to interact with several resources supported by ALKS. The provider needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    alks:url:
        value: https://alks.foo.com/rest

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as alks from "@pulumi/alks";

// Create an ALKS IAM role
const testRole = new alks.Iamrole("test_role", {
    name: "My_Test_Role",
    assumeRolePolicy: JSON.stringify({
        Version: "2012-10-17",
        Statement: [{
            Action: "sts:AssumeRole",
            Effect: "Allow",
            Principal: {
                Service: "ec2.amazonaws.com",
            },
            Sid: "",
        }],
    }),
    includeDefaultPolicies: false,
    enableAlksAccess: false,
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    alks:url:
        value: https://alks.foo.com/rest

```
```python
import pulumi
import json
import pulumi_alks as alks

# Create an ALKS IAM role
test_role = alks.Iamrole("test_role",
    name="My_Test_Role",
    assume_role_policy=json.dumps({
        "Version": "2012-10-17",
        "Statement": [{
            "Action": "sts:AssumeRole",
            "Effect": "Allow",
            "Principal": {
                "Service": "ec2.amazonaws.com",
            },
            "Sid": "",
        }],
    }),
    include_default_policies=False,
    enable_alks_access=False)
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    alks:url:
        value: https://alks.foo.com/rest

```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Alks = Pulumi.Alks;

return await Deployment.RunAsync(() =>
{
    // Create an ALKS IAM role
    var testRole = new Alks.Iamrole("test_role", new()
    {
        Name = "My_Test_Role",
        AssumeRolePolicy = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Version"] = "2012-10-17",
            ["Statement"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["Action"] = "sts:AssumeRole",
                    ["Effect"] = "Allow",
                    ["Principal"] = new Dictionary<string, object?>
                    {
                        ["Service"] = "ec2.amazonaws.com",
                    },
                    ["Sid"] = "",
                },
            },
        }),
        IncludeDefaultPolicies = false,
        EnableAlksAccess = false,
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    alks:url:
        value: https://alks.foo.com/rest

```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/alks/v2/alks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Version": "2012-10-17",
			"Statement": []map[string]interface{}{
				map[string]interface{}{
					"Action": "sts:AssumeRole",
					"Effect": "Allow",
					"Principal": map[string]interface{}{
						"Service": "ec2.amazonaws.com",
					},
					"Sid": "",
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		// Create an ALKS IAM role
		_, err = alks.NewIamrole(ctx, "test_role", &alks.IamroleArgs{
			Name:                   pulumi.String("My_Test_Role"),
			AssumeRolePolicy:       pulumi.String(json0),
			IncludeDefaultPolicies: pulumi.Bool(false),
			EnableAlksAccess:       pulumi.Bool(false),
		})
		if err != nil {
			return err
		}
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
config:
    alks:url:
        value: https://alks.foo.com/rest

```
```yaml
resources:
  # Create an ALKS IAM role
  testRole:
    type: alks:Iamrole
    name: test_role
    properties:
      name: My_Test_Role
      assumeRolePolicy:
        fn::toJSON:
          Version: 2012-10-17
          Statement:
            - Action: sts:AssumeRole
              Effect: Allow
              Principal:
                Service: ec2.amazonaws.com
              Sid: ""
      includeDefaultPolicies: false
      enableAlksAccess: false
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    alks:url:
        value: https://alks.foo.com/rest

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.alks.Iamrole;
import com.pulumi.alks.IamroleArgs;
import static com.pulumi.codegen.internal.Serialization.*;
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
        // Create an ALKS IAM role
        var testRole = new Iamrole("testRole", IamroleArgs.builder()
            .name("My_Test_Role")
            .assumeRolePolicy(serializeJson(
                jsonObject(
                    jsonProperty("Version", "2012-10-17"),
                    jsonProperty("Statement", jsonArray(jsonObject(
                        jsonProperty("Action", "sts:AssumeRole"),
                        jsonProperty("Effect", "Allow"),
                        jsonProperty("Principal", jsonObject(
                            jsonProperty("Service", "ec2.amazonaws.com")
                        )),
                        jsonProperty("Sid", "")
                    )))
                )))
            .includeDefaultPolicies(false)
            .enableAlksAccess(false)
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Authentication

The ALKS Pulumi Provider offers a flexible means of providing credentials for authentication. The following methods are supported, in this order, and explained below:
### Static credentials

!> **Warning:** Hard-coding credentials into any Pulumi configuration is not recommended, and risks secret leakage should this file ever be committed to a public version control system.

Static credentials can be provided via an `accessKey`, `secretKey` and `token` in-line in the ALKS provider configuration. This method is generally not recommended, since the credentials could accidentally be committed or shared.

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    alks:accessKey:
        value: accesskey
    alks:secretKey:
        value: secretkey
    alks:token:
        value: sessiontoken
    alks:url:
        value: https://alks.foo.com/rest

```
### Environment variables
You can provide your credentials via the `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY` and `AWS_SESSION_TOKEN` environment variables. The `ALKS_` names (for example `ALKS_ACCESS_KEY_ID`) are also read, but only when the matching `AWS_` name is not set: `AWS_*` takes precedence. If you need the ALKS provider to use different credentials from other AWS tooling in the same shell (for example PowerUser for the AWS provider and IAMAdmin for ALKS), use the `ALKSTFP_` names (for example `ALKSTFP_ACCESS_KEY_ID`), which take precedence over `AWS_*` for this provider only. See Environment variables for every variable.

> **Changed in 2.9.0:** earlier versions of this page said `ALKS_` credentials take priority over `AWS_`. The provider has always used `AWS_` first; the page was wrong, and the behaviour is unchanged.

Pulumi file:
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    alks:url:
        value: https://alks.foo.com/rest

```

Terminal:
### Shared credentials/configuration file
You can use an AWS credentials file to specify your credentials. The default location is `$HOME/.aws/credentials` on Linux and OSX, or `"%USERPROFILE%\.aws\credentials"` for Windows users. If we fail to detect credentials inline, or in the environment, Pulumi will check this location last. You can optionally specify a different location in the configuration via the `sharedCredentialsFile` attribute, or via the environment with the `AWS_SHARED_CREDENTIALS_FILE` variable. This method also supports a profile configuration and matching `AWS_PROFILE` environment variable.

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    alks:profile:
        value: foo
    alks:sharedCredentialsFile:
        value: /Users/my_user/.aws/credentials
    alks:url:
        value: https://alks.foo.com/rest

```
### Machine Identities
You can use a role created with ALKS with the `enableAlksAccess` flag set to `true` to authenticate requests against ALKS.

In order to do this, ALKS must be called from within AWS using STS credentials from an instance profile associated with the role with `enableAlksAccess` set. This also works from Lambda functions in the same way.

The STS credentials are used and provided in the same way that the AWS CLI uses the credentials, so there is nothing special you have to do to use Machine Identities.

Your ALKS provider configuration can look just like this:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    alks:url:
        value: https://alks.foo.com/rest

```

Since Machine Identities work with Instance Profile Metadata directly, it can be helpful to assume another role or cross account trust. For example:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    alks:url:
        value: https://alks.foo.com/rest

```
### Tags
You can specify tags to add to all of your roles created with ALKS by using the `defaultTags` block in the provider configuration.  You can also choose to ignore existing tags on a resource by including tag keys or key prefixes in the `ignoreTags` block.  These ignored tags will not show up on Pulumi Plans or Applys, and will not be removed from the resource by Pulumi.

You may also specify tags on individual roles using the `tags` block.

Your ALKS configuration could look like this:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    alks:url:
        value: https://alks.foo.com/rest

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as alks from "@pulumi/alks";

const testRole = new alks.Iamrole("test_role", {
    name: "My_Test_Role",
    type: "Amazon EC2",
    includeDefaultPolicies: false,
    enableAlksAccess: false,
    tags: {
        roleSpecificTagKey: "value",
    },
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    alks:url:
        value: https://alks.foo.com/rest

```
```python
import pulumi
import pulumi_alks as alks

test_role = alks.Iamrole("test_role",
    name="My_Test_Role",
    type="Amazon EC2",
    include_default_policies=False,
    enable_alks_access=False,
    tags={
        "roleSpecificTagKey": "value",
    })
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    alks:url:
        value: https://alks.foo.com/rest

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Alks = Pulumi.Alks;

return await Deployment.RunAsync(() =>
{
    var testRole = new Alks.Iamrole("test_role", new()
    {
        Name = "My_Test_Role",
        Type = "Amazon EC2",
        IncludeDefaultPolicies = false,
        EnableAlksAccess = false,
        Tags =
        {
            { "roleSpecificTagKey", "value" },
        },
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    alks:url:
        value: https://alks.foo.com/rest

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/alks/v2/alks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := alks.NewIamrole(ctx, "test_role", &alks.IamroleArgs{
			Name:                   pulumi.String("My_Test_Role"),
			Type:                   pulumi.String("Amazon EC2"),
			IncludeDefaultPolicies: pulumi.Bool(false),
			EnableAlksAccess:       pulumi.Bool(false),
			Tags: pulumi.StringMap{
				"roleSpecificTagKey": pulumi.String("value"),
			},
		})
		if err != nil {
			return err
		}
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
config:
    alks:url:
        value: https://alks.foo.com/rest

```
```yaml
resources:
  testRole:
    type: alks:Iamrole
    name: test_role
    properties:
      name: My_Test_Role
      type: Amazon EC2
      includeDefaultPolicies: false
      enableAlksAccess: false
      tags:
        roleSpecificTagKey: value
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    alks:url:
        value: https://alks.foo.com/rest

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.alks.Iamrole;
import com.pulumi.alks.IamroleArgs;
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
        var testRole = new Iamrole("testRole", IamroleArgs.builder()
            .name("My_Test_Role")
            .type("Amazon EC2")
            .includeDefaultPolicies(false)
            .enableAlksAccess(false)
            .tags(Map.of("roleSpecificTagKey", "value"))
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

Note: Role specific tag values will overwrite default values if the key appears in both places.
### Multiple Provider Configuration

You can configure multiple ALKS providers to each have their own account context.

The initial provider must have credentials set in a default way (static, shared credentials file, environment variables, etc) before the second provider can determine whether your account/role combination are allowed.

The second (or so) provider can then be used to generate resources for multiple accounts in one plan / apply.

Note: This only works for accounts you have access to!

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    alks:url:
        value: https://alks.coxautoinc.com/rest

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as alks from "@pulumi/alks";

// CREATE IAM ROLE -- PROVIDER 1
const testRole = new alks.Iamrole("test_role", {
    name: "TEST-DELETE",
    type: "AWS CodeBuild",
    includeDefaultPolicies: false,
    enableAlksAccess: true,
});
// CREATE IAM ROLE -- PROVIDER 2
const testRoleNonprod = new alks.Iamrole("test_role_nonprod", {
    name: "TEST-DELETE",
    type: "AWS CodeBuild",
    includeDefaultPolicies: false,
    enableAlksAccess: true,
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    alks:url:
        value: https://alks.coxautoinc.com/rest

```
```python
import pulumi
import pulumi_alks as alks

# CREATE IAM ROLE -- PROVIDER 1
test_role = alks.Iamrole("test_role",
    name="TEST-DELETE",
    type="AWS CodeBuild",
    include_default_policies=False,
    enable_alks_access=True)
# CREATE IAM ROLE -- PROVIDER 2
test_role_nonprod = alks.Iamrole("test_role_nonprod",
    name="TEST-DELETE",
    type="AWS CodeBuild",
    include_default_policies=False,
    enable_alks_access=True)
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    alks:url:
        value: https://alks.coxautoinc.com/rest

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Alks = Pulumi.Alks;

return await Deployment.RunAsync(() =>
{
    // CREATE IAM ROLE -- PROVIDER 1
    var testRole = new Alks.Iamrole("test_role", new()
    {
        Name = "TEST-DELETE",
        Type = "AWS CodeBuild",
        IncludeDefaultPolicies = false,
        EnableAlksAccess = true,
    });

    // CREATE IAM ROLE -- PROVIDER 2
    var testRoleNonprod = new Alks.Iamrole("test_role_nonprod", new()
    {
        Name = "TEST-DELETE",
        Type = "AWS CodeBuild",
        IncludeDefaultPolicies = false,
        EnableAlksAccess = true,
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    alks:url:
        value: https://alks.coxautoinc.com/rest

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/alks/v2/alks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// CREATE IAM ROLE -- PROVIDER 1
		_, err := alks.NewIamrole(ctx, "test_role", &alks.IamroleArgs{
			Name:                   pulumi.String("TEST-DELETE"),
			Type:                   pulumi.String("AWS CodeBuild"),
			IncludeDefaultPolicies: pulumi.Bool(false),
			EnableAlksAccess:       pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		// CREATE IAM ROLE -- PROVIDER 2
		_, err = alks.NewIamrole(ctx, "test_role_nonprod", &alks.IamroleArgs{
			Name:                   pulumi.String("TEST-DELETE"),
			Type:                   pulumi.String("AWS CodeBuild"),
			IncludeDefaultPolicies: pulumi.Bool(false),
			EnableAlksAccess:       pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
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
config:
    alks:url:
        value: https://alks.coxautoinc.com/rest

```
```yaml
resources:
  # CREATE IAM ROLE -- PROVIDER 1
  testRole:
    type: alks:Iamrole
    name: test_role
    properties:
      name: TEST-DELETE
      type: AWS CodeBuild
      includeDefaultPolicies: false
      enableAlksAccess: true
  # CREATE IAM ROLE -- PROVIDER 2
  testRoleNonprod:
    type: alks:Iamrole
    name: test_role_nonprod
    properties:
      name: TEST-DELETE
      type: AWS CodeBuild
      includeDefaultPolicies: false
      enableAlksAccess: true
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    alks:url:
        value: https://alks.coxautoinc.com/rest

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.alks.Iamrole;
import com.pulumi.alks.IamroleArgs;
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
        // CREATE IAM ROLE -- PROVIDER 1
        var testRole = new Iamrole("testRole", IamroleArgs.builder()
            .name("TEST-DELETE")
            .type("AWS CodeBuild")
            .includeDefaultPolicies(false)
            .enableAlksAccess(true)
            .build());

        // CREATE IAM ROLE -- PROVIDER 2
        var testRoleNonprod = new Iamrole("testRoleNonprod", IamroleArgs.builder()
            .name("TEST-DELETE")
            .type("AWS CodeBuild")
            .includeDefaultPolicies(false)
            .enableAlksAccess(true)
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

In addition to generic `provider` arguments (e.g. `alias` and `version`), the following arguments are supported in the ALKS provider configuration:

* `url` - (Required) The URL to your ALKS server. Also read from ENV.ALKSTFP_URL or ENV.ALKS_URL
* `accessKey` - (Optional) The access key from a valid STS session. Also read from ENV.ALKSTFP_ACCESS_KEY_ID, ENV.AWS_ACCESS_KEY_ID and ENV.ALKS_ACCESS_KEY_ID, in that order.
* `secretKey` - (Optional) The secret key from a valid STS session. Also read from ENV.ALKSTFP_SECRET_ACCESS_KEY, ENV.AWS_SECRET_ACCESS_KEY and ENV.ALKS_SECRET_ACCESS_KEY, in that order.
* `token` - (Optional) The session token from a valid STS session. Also read from ENV.ALKSTFP_SESSION_TOKEN, ENV.AWS_SESSION_TOKEN and ENV.ALKS_SESSION_TOKEN, in that order.
* `sharedCredentialsFile` - (Optional) The the path to the shared credentials file. Also read from ENV.ALKSTFP_SHARED_CREDENTIALS_FILE, ENV.AWS_SHARED_CREDENTIALS_FILE and ENV.ALKS_SHARED_CREDENTIALS_FILE, in that order.
* `profile` - (Optional) This is the AWS profile name as set in the shared credentials file. Also read from ENV.ALKSTFP_PROFILE, ENV.AWS_PROFILE and ENV.ALKS_PROFILE, in that order.
* `proxyMode` - (Optional) How ALKS requests choose a proxy: `auto` (default), `environment` or `none`. See Proxies. Applies to ALKS requests only, never to AWS calls. Added in 2.9.0.
* `proxyUrl` - (Optional) Send every ALKS request through this proxy, for example `http://host.docker.internal:9000`. `NO_PROXY` does not apply to it. Cannot be combined with `proxyMode = "none"`. The URL has no path, query or fragment; percent-encode `/`, `?`, `#` and `@` in a proxy password. Added in 2.9.0.
* `assumeRole` - (Optional) This is the role information to assume before making calling ALKS. This feature works the same as the assumeRole feature of the AWS Pulumi Provider.
  * `roleArn` - (Required) The Role ARN to assume for calling the ALKS API.
  * `sessionName` - (Optional) The session name to provide to AWS when creating STS credentials. Please see the AWS SDK documentation for more information.
  * `externalId` - (Optional) The external identifier to provide to AWS when creating STS credentials. Please see the AWS SDK documentation for more information.
  * `policy` - (Optional) This specifies additional policy restrictions to apply to the resulting STS credentials beyond any existing inline or managed policies. Please see the AWS SDK documentation for more information.
* `defaultTags` - (Optional) This block can hold a block of tags to add to all roles created by this provider
  * `tags` - (Optional) Block of key value pairs to add to all roles
* `ignoreTags` - (Optional) Can contain a list of tag keys or key prefixes to exclude from `pulumi preview` and `pulumi up`.  This is for tags added outside of the alks provider that are managed externally
  * `keys` - (Optional) List of keys to ignore
  * `keyPrefixes` - (Optional) List of key prefixes to ignore. Any key starting with a string in this list will be ignored.

---
### Environment variables reference

Every provider argument can also come from the environment. **A value in the provider configuration wins for that argument**, and then that argument's environment variables are not read. Otherwise the first variable that is set wins, left to right. This is decided argument by argument, not for the identity as a whole: for example, a `profile` in the provider configuration does not stop `AWS_ACCESS_KEY_ID` from being used, because access keys take precedence over a profile.

|         Argument          |                                    Environment variables, in                                     |

|                           |                                         precedence order                                         |
|---------------------------|--------------------------------------------------------------------------------------------------|
| `url`                     | `ALKSTFP_URL`, `ALKS_URL`                                                                        |
| `accessKey`              | `ALKSTFP_ACCESS_KEY_ID`, `AWS_ACCESS_KEY_ID`, `ALKS_ACCESS_KEY_ID`                               |
| `secretKey`              | `ALKSTFP_SECRET_ACCESS_KEY`, `AWS_SECRET_ACCESS_KEY`, `ALKS_SECRET_ACCESS_KEY`                   |
| `token`                   | `ALKSTFP_SESSION_TOKEN`, `AWS_SESSION_TOKEN`, `ALKS_SESSION_TOKEN`                               |
| `profile`                 | `ALKSTFP_PROFILE`, `AWS_PROFILE`, `ALKS_PROFILE`                                                 |
| `sharedCredentialsFile` | `ALKSTFP_SHARED_CREDENTIALS_FILE`, `AWS_SHARED_CREDENTIALS_FILE`, `ALKS_SHARED_CREDENTIALS_FILE` |
| `account`                 | `ALKSTFP_ACCOUNT`, `ALKS_ACCOUNT`, `Account`                                                     |
| `role`                    | `ALKSTFP_ROLE`, `ALKS_ROLE`, `Role`                                                              |
| `proxyMode`              | `ALKSTFP_PROXY_MODE`, `ALKS_PROXY_MODE`                                                          |
| `proxyUrl`               | `ALKSTFP_PROXY_URL`, `ALKS_PROXY_URL`                                                            |

**Conflicts are errors.** If two of the provider's own names for the same argument are set (`ALKSTFP_*`, `ALKS_*`, or the legacy `Account`/`Role`), the provider stops with an error naming them, even if their values are equal. Unset one, or set the argument in the provider configuration. `AWS_*` names are not part of this check, so `AWS_ACCESS_KEY_ID` and `ALKS_ACCESS_KEY_ID` together is fine; `AWS_*` wins.

**Newly read in 2.9.0:** `ALKS_PROFILE`, `ALKS_SHARED_CREDENTIALS_FILE`, `ALKS_ACCOUNT` and `ALKS_ROLE` were ignored before 2.9.0. If your shell already exports them for other ALKS tooling, the provider now uses them. `ALKS_ACCOUNT` with `ALKS_ROLE` switches it to that account and role. Unset them, or set the argument in the provider configuration, to keep the earlier result.

`Account` and `Role` are the original names and keep working. On macOS and Linux they are case-sensitive (`ACCOUNT` is a different variable).

**`ALKSTFP_*` gives this provider its own identity.**
- **Credentials are a set.** If any of `ALKSTFP_ACCESS_KEY_ID`, `ALKSTFP_SECRET_ACCESS_KEY` or `ALKSTFP_SESSION_TOKEN` is set, all three come from `ALKSTFP_`. A missing one stays empty; it is never filled from `AWS_*` or `ALKS_*`, so a key pair from one session is never combined with another session's token.
- **Not mixed with the provider configuration.** Setting `accessKey`, `secretKey` or `token` in the provider configuration while an `ALKSTFP_` credential variable is set is an error. Put all the credentials in one place.
- **An incomplete set is an error.** If an `ALKSTFP_` credential is set but the access key or secret key is still missing, and no profile is set, the provider stops with an error. It does not fall back to `AWS_*` keys, `AWS_PROFILE` or the default profile. While an `ALKSTFP_` credential is set, the profile comes only from `ALKSTFP_PROFILE` (or the provider configuration), never from `AWS_PROFILE` or `ALKS_PROFILE`.
- **The profile wins.** If `ALKSTFP_PROFILE` is set, `AWS_*` and `ALKS_*` keys exported for other tooling are ignored, so the profile is used.
- **Scope.** These rules apply only when the `ALKSTFP_` variables are used. Otherwise the behaviour is as in earlier versions, where access keys, including `AWS_*`, take precedence over a profile, as in the AWS provider. A value in the provider configuration wins for its own argument.
- **`sharedCredentialsFile` is honoured throughout.** When a credentials file is given in the provider configuration or through `ALKSTFP_`/`ALKS_`, profiles are read only from that file, and a file that cannot be read is an error. Profiles are never silently read from the default `~/.aws/credentials`.

`accessKey`, `secretKey`, `token` and `proxyUrl` are marked sensitive, so Pulumi does not show them.
### Proxies

From 2.9.0 the provider finds a proxy for its **ALKS** requests the way a browser does. It never changes how it calls AWS. With the default, `proxyMode = "auto"`, each ALKS request uses:

1. `HTTPS_PROXY`, unless `NO_PROXY` matches the ALKS host. This is the same as earlier versions.
2. Otherwise the operating system's proxy settings: on **macOS** the system proxy and PAC file; on **Windows** the WinHTTP / Internet Options settings, PAC file and automatic detection. Linux has no OS step.
3. Otherwise a direct connection.

Rules:
- Only the first entry of a PAC answer is used. If that proxy does not answer, the request fails with an error naming it.
- `NO_PROXY` also exempts hosts from an OS proxy.
- `localhost` is never proxied.

This means Zscaler laptops that only route proxy-aware applications reach ALKS with nothing to configure.

|   `proxyMode`   |                             Behaviour                             |
|------------------|-------------------------------------------------------------------|
| `auto` (default) | As above                                                          |
| `environment`    | `HTTPS_PROXY`/`NO_PROXY` only: exactly the behaviour before 2.9.0 |
| `none`           | Always direct, ignoring every proxy setting                       |

`proxyUrl` forces one proxy for every ALKS request.

**In a container** the host's proxy settings are not visible, so `auto` falls back to `HTTPS_PROXY`, or direct. With Docker Desktop, the Zscaler local proxy is reachable as `host.docker.internal`:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    alks:proxyUrl:
        value: http://host.docker.internal:9000
    alks:url:
        value: https://alks.example.com/rest

```

`127.0.0.1:9000` does not work inside a container: it is the container itself. Setting `HTTPS_PROXY` instead also sends AWS calls through the proxy. `proxyUrl`/`ALKS_PROXY_URL` affect ALKS only.

**Debugging:** with `TF_LOG=DEBUG`, the log has one line per client with the proxy decision, for example `alks-go: proxy mode=auto source=pac proxy=http://127.0.0.1:9000`. It never contains credentials.
### Supported Versions

|    Pulumi 0.10.x    |    Pulumi 0.11.x     |  Pulumi 0.12.x - Pulumi  | Pulumi 1.0.x |

|                        |                         |     0.15.x      |                 |
|------------------------|-------------------------|-----------------|-----------------|
| ALKS TFP 0.9.0 < 1.3.0 | ALKS TFP 1.3.0 < 1.5.15 | ALKS TFP 1.3.0+ | ALKS TFP 2.0.0+ |

For questions, please reach out to the [ALKS team](https://github.com/orgs/Cox-Automotive/teams/cai-internal-tools).