---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/philips-software/hsdp/0.71.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Hsdp Provider
meta_desc: Provides an overview on how to configure the Pulumi Hsdp provider.
layout: package
---

## Generate Provider

The Hsdp provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider philips-software/hsdp
```
## Overview

Use the HSDP provider to interact with the many resources supported by [HSDP](https://www.hsdp.io). This includes amongst others IAM entities, Container Host instances and Edge devices

Use the navigation to the left to read about the available resources.

To learn the basics of Pulumi, follow the hands-on get started tutorials on HashiCorp's Learn platform.
## Example usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    hsdp:environment:
        value: client-test
    hsdp:oauth2ClientId:
        value: 'TODO: var.oauth2_client_id'
    hsdp:oauth2Password:
        value: 'TODO: var.oauth2_password'
    hsdp:orgAdminPassword:
        value: 'TODO: var.org_admin_password'
    hsdp:orgAdminUsername:
        value: 'TODO: var.org_admin_username'
    hsdp:region:
        value: us-east

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as hsdp from "@pulumi/hsdp";

const hospitalA = new hsdp.IamOrg("hospital_a", {
    name: "HOSPITAL_A",
    description: "HOSPITAL A",
    parentOrgId: root.id,
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    hsdp:environment:
        value: client-test
    hsdp:oauth2ClientId:
        value: 'TODO: var.oauth2_client_id'
    hsdp:oauth2Password:
        value: 'TODO: var.oauth2_password'
    hsdp:orgAdminPassword:
        value: 'TODO: var.org_admin_password'
    hsdp:orgAdminUsername:
        value: 'TODO: var.org_admin_username'
    hsdp:region:
        value: us-east

```
```python
import pulumi
import pulumi_hsdp as hsdp

hospital_a = hsdp.IamOrg("hospital_a",
    name="HOSPITAL_A",
    description="HOSPITAL A",
    parent_org_id=root["id"])
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    hsdp:environment:
        value: client-test
    hsdp:oauth2ClientId:
        value: 'TODO: var.oauth2_client_id'
    hsdp:oauth2Password:
        value: 'TODO: var.oauth2_password'
    hsdp:orgAdminPassword:
        value: 'TODO: var.org_admin_password'
    hsdp:orgAdminUsername:
        value: 'TODO: var.org_admin_username'
    hsdp:region:
        value: us-east

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Hsdp = Pulumi.Hsdp;

return await Deployment.RunAsync(() =>
{
    var hospitalA = new Hsdp.IamOrg("hospital_a", new()
    {
        Name = "HOSPITAL_A",
        Description = "HOSPITAL A",
        ParentOrgId = root.Id,
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
    hsdp:environment:
        value: client-test
    hsdp:oauth2ClientId:
        value: 'TODO: var.oauth2_client_id'
    hsdp:oauth2Password:
        value: 'TODO: var.oauth2_password'
    hsdp:orgAdminPassword:
        value: 'TODO: var.org_admin_password'
    hsdp:orgAdminUsername:
        value: 'TODO: var.org_admin_username'
    hsdp:region:
        value: us-east

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/hsdp/hsdp"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hsdp.NewIamOrg(ctx, "hospital_a", &hsdp.IamOrgArgs{
			Name:        pulumi.String("HOSPITAL_A"),
			Description: pulumi.String("HOSPITAL A"),
			ParentOrgId: pulumi.Any(root.Id),
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
    hsdp:environment:
        value: client-test
    hsdp:oauth2ClientId:
        value: 'TODO: var.oauth2_client_id'
    hsdp:oauth2Password:
        value: 'TODO: var.oauth2_password'
    hsdp:orgAdminPassword:
        value: 'TODO: var.org_admin_password'
    hsdp:orgAdminUsername:
        value: 'TODO: var.org_admin_username'
    hsdp:region:
        value: us-east

```
```yaml
resources:
  hospitalA:
    type: hsdp:IamOrg
    name: hospital_a
    properties:
      name: HOSPITAL_A
      description: HOSPITAL A
      parentOrgId: ${root.id}
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    hsdp:environment:
        value: client-test
    hsdp:oauth2ClientId:
        value: 'TODO: var.oauth2_client_id'
    hsdp:oauth2Password:
        value: 'TODO: var.oauth2_password'
    hsdp:orgAdminPassword:
        value: 'TODO: var.org_admin_password'
    hsdp:orgAdminUsername:
        value: 'TODO: var.org_admin_username'
    hsdp:region:
        value: us-east

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.hsdp.IamOrg;
import com.pulumi.hsdp.IamOrgArgs;
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
        var hospitalA = new IamOrg("hospitalA", IamOrgArgs.builder()
            .name("HOSPITAL_A")
            .description("HOSPITAL A")
            .parentOrgId(root.id())
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Authentication

The HSDP provider can read credentials and settings from the Environment or as
arguments in its provider configuration. The following environment variables are recognized

|         Environment          |      Argument       | Required |   Default   |
|------------------------------|---------------------|----------|-------------|
| HSDP_REGION                  | region              | Optional | us-east     |
| HSDP_ENVIRONMENT             | environment         | Optional | client-test |
| HSDP_CARTEL_HOST             | cartelHost         | Optional |             |
| HSDP_CARTEL_SECRET           | cartelSecret       | Optional |             |
| HSDP_CARTEL_TOKEN            | cartelToken        | Optional |             |
| HSDP_IAM_SERVICE_ID          | serviceId          | Optional |             |
| HSDP_IAM_SERVICE_PRIVATE_KEY | servicePrivateKey | Optional |             |
| HSDP_IAM_ORG_ADMIN_USERNAME  | orgAdminUsername  | Optional |             |
| HSDP_IAM_ORG_ADMIN_PASSWORD  | orgAdminPassword  | Optional |             |
| HSDP_IAM_OAUTH2_CLIENT_ID    | oauth2ClientId    | Optional |             |
| HSDP_IAM_OAUTH2_PASSWORD     | oauth2Password     | Optional |             |
| HSDP_SHARED_KEY              | sharedKey          | Optional |             |
| HSDP_SECRET_KEY              | secretKey          | Optional |             |
| HSDP_UAA_USERNAME            | uaaUsername        | Optional |             |
| HSDP_UAA_PASSWORD            | uaaPassword        | Optional |             |
| HSDP_DEBUG_LOG               | debugLog           | Optional |             |
| HSDP_DEBUG_STDERR            | debugStderr        | Optional |             |
## Configuration Reference

In addition to generic provider arguments (e.g. alias and version), the following arguments are supported in the HSDP provider configuration:

* `region` - (Required) The HSDP region to use [`us-east`, `eu-west`, `sa1`, `ca1`, `apac3`, ...]. Default is `us-east`
* `environment` - (Optional) The HSDP environment to use within region [`client-test`, `prod`] . Default is `client-test`
* `credentials` - (Optional) Can point to a JSON file containing values for all fields here
* `iamUrl` - (Optional) IAM API endpoint. Auto-discovered from region and environment.
* `idmUrl` - (Optional) IDM API endpoint Auto-discovered from region and environment.
* `s3credsUrl` - (Optional) S3 Credentials API endpoint. Auto-discovered from region and environment.
* `notificationUrl` - (Optional) Notification service URL. Auto-discovered from region and environment.
* `oauth2ClientId` - (Optional) The OAuth2 client ID as provided by HSDP
* `oauth2Password` - (Optional) The OAuth2 password as provided by HSDP
* `serviceId` - (Optional) The service ID to use for IAM org admin operations (conflicts with: `orgAdminUsername`)
* `servicePrivateKey` - (Optional) The service private key to use for IAM org admin operations (conflicts with: `orgAdminPassword`)
* `orgAdminUsername` - (Optional) Your IAM admin username.
* `orgAdminPassword` - (Optional) Your IAM admin password.
* `uaaUsername` - (Optional) The HSDP CF UAA username.
* `uaaPassword` - (Optional) The HSDP CF UAA password.
* `uaaUrl` - (Optional) The URL of the UAA authentication service. Auto-discovered from region.
* `mdmUrl` - (Optional) The base URL of the MDM service. Auto-discovered from region and environment.
* `sharedKey` - (Optional) The shared key as provided by HSDP. Actions which require API signing will not work if this value is missing.
* `secretKey` - (Optional) The secret key as provided by HSDP. Actions which require API signing will not work if this value is missing.
* `cartelHost` - (Optional) The cartel host as provided by HSDP. Auto-discovered from region.
* `cartelToken` - (Optional) The cartel token as provided by HSDP.
* `cartelSecret` - (Optional) The cartel secret as provided by HSDP.
* `retryMax` - (Optional) Integer, when > 0 will use a retry-able HTTP client and retry requests when applicable.
* `debugLog` - (Optional) If set to a path, when debug is enabled outputs details to this file
* `debugStderr` - (Optional) If set to true sends debug logs to `stderr`