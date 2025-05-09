---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/ibm-cloud/ibm/1.78.1/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Ibm Provider
meta_desc: Provides an overview on how to configure the Pulumi Ibm provider.
layout: package
---

## Generate Provider

The Ibm provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider ibm-cloud/ibm
```
## Overview

The IBM Cloud provider is used to manage IBM Cloud resources. The provider must be configured with the proper credentials before it can be used.

Use the navigation menu on the left to read about the available functions and resources.
## Example usage of provider


{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    ibm:region:
        value: us-south

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as ibm from "@pulumi/ibm";

// Create a VPC
const testaccVpc = new ibm.IsVpc("testacc_vpc", {name: "test-vpc"});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    ibm:region:
        value: us-south

```
```python
import pulumi
import pulumi_ibm as ibm

# Create a VPC
testacc_vpc = ibm.IsVpc("testacc_vpc", name="test-vpc")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    ibm:region:
        value: us-south

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Ibm = Pulumi.Ibm;

return await Deployment.RunAsync(() =>
{
    // Create a VPC
    var testaccVpc = new Ibm.IsVpc("testacc_vpc", new()
    {
        Name = "test-vpc",
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
    ibm:region:
        value: us-south

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/ibm/ibm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a VPC
		_, err := ibm.NewIsVpc(ctx, "testacc_vpc", &ibm.IsVpcArgs{
			Name: pulumi.String("test-vpc"),
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
    ibm:region:
        value: us-south

```
```yaml
resources:
  # Create a VPC
  testaccVpc:
    type: ibm:IsVpc
    name: testacc_vpc
    properties:
      name: test-vpc
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    ibm:region:
        value: us-south

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.ibm.IsVpc;
import com.pulumi.ibm.IsVpcArgs;
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
        // Create a VPC
        var testaccVpc = new IsVpc("testaccVpc", IsVpcArgs.builder()
            .name("test-vpc")
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}


{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    ibm:region:
        value: us-south

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as ibm from "@pulumi/ibm";

// Create a VPC
const testaccVpc = new ibm.IsVpc("testacc_vpc", {name: "test-vpc"});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    ibm:region:
        value: us-south

```
```python
import pulumi
import pulumi_ibm as ibm

# Create a VPC
testacc_vpc = ibm.IsVpc("testacc_vpc", name="test-vpc")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    ibm:region:
        value: us-south

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Ibm = Pulumi.Ibm;

return await Deployment.RunAsync(() =>
{
    // Create a VPC
    var testaccVpc = new Ibm.IsVpc("testacc_vpc", new()
    {
        Name = "test-vpc",
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
    ibm:region:
        value: us-south

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/ibm/ibm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a VPC
		_, err := ibm.NewIsVpc(ctx, "testacc_vpc", &ibm.IsVpcArgs{
			Name: pulumi.String("test-vpc"),
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
    ibm:region:
        value: us-south

```
```yaml
resources:
  # Create a VPC
  testaccVpc:
    type: ibm:IsVpc
    name: testacc_vpc
    properties:
      name: test-vpc
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    ibm:region:
        value: us-south

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.ibm.IsVpc;
import com.pulumi.ibm.IsVpcArgs;
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
        // Create a VPC
        var testaccVpc = new IsVpc("testaccVpc", IsVpcArgs.builder()
            .name("test-vpc")
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

Visiblity support:
{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    ibm:visibility:
        value: private

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as ibm from "@pulumi/ibm";

// Create a VPC
const testaccVpc = new ibm.IsVpc("testacc_vpc", {name: "test-vpc"});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    ibm:visibility:
        value: private

```
```python
import pulumi
import pulumi_ibm as ibm

# Create a VPC
testacc_vpc = ibm.IsVpc("testacc_vpc", name="test-vpc")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    ibm:visibility:
        value: private

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Ibm = Pulumi.Ibm;

return await Deployment.RunAsync(() =>
{
    // Create a VPC
    var testaccVpc = new Ibm.IsVpc("testacc_vpc", new()
    {
        Name = "test-vpc",
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
    ibm:visibility:
        value: private

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/ibm/ibm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a VPC
		_, err := ibm.NewIsVpc(ctx, "testacc_vpc", &ibm.IsVpcArgs{
			Name: pulumi.String("test-vpc"),
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
    ibm:visibility:
        value: private

```
```yaml
resources:
  # Create a VPC
  testaccVpc:
    type: ibm:IsVpc
    name: testacc_vpc
    properties:
      name: test-vpc
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    ibm:visibility:
        value: private

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.ibm.IsVpc;
import com.pulumi.ibm.IsVpcArgs;
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
        // Create a VPC
        var testaccVpc = new IsVpc("testaccVpc", IsVpcArgs.builder()
            .name("test-vpc")
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

VPE endpoints support:
{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    ibm:privateEndpointType:
        value: vpe
    ibm:visibility:
        value: private

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as ibm from "@pulumi/ibm";

// List Cloud Logs alerts
const alerts = ibm.getLogsAlerts({
    instanceId: "logs_instance_guid",
    region: "logs_instance_region",
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    ibm:privateEndpointType:
        value: vpe
    ibm:visibility:
        value: private

```
```python
import pulumi
import pulumi_ibm as ibm

# List Cloud Logs alerts
alerts = ibm.get_logs_alerts(instance_id="logs_instance_guid",
    region="logs_instance_region")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    ibm:privateEndpointType:
        value: vpe
    ibm:visibility:
        value: private

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Ibm = Pulumi.Ibm;

return await Deployment.RunAsync(() =>
{
    // List Cloud Logs alerts
    var alerts = Ibm.GetLogsAlerts.Invoke(new()
    {
        InstanceId = "logs_instance_guid",
        Region = "logs_instance_region",
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
    ibm:privateEndpointType:
        value: vpe
    ibm:visibility:
        value: private

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/ibm/ibm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// List Cloud Logs alerts
		_, err := ibm.GetLogsAlerts(ctx, &ibm.GetLogsAlertsArgs{
			InstanceId: "logs_instance_guid",
			Region:     pulumi.StringRef("logs_instance_region"),
		}, nil)
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
    ibm:privateEndpointType:
        value: vpe
    ibm:visibility:
        value: private

```
```yaml
variables:
  # List Cloud Logs alerts
  alerts:
    fn::invoke:
      function: ibm:getLogsAlerts
      arguments:
        instanceId: logs_instance_guid
        region: logs_instance_region
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    ibm:privateEndpointType:
        value: vpe
    ibm:visibility:
        value: private

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.ibm.IbmFunctions;
import com.pulumi.ibm.inputs.GetLogsAlertsArgs;
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
        // List Cloud Logs alerts
        final var alerts = IbmFunctions.getLogsAlerts(GetLogsAlertsArgs.builder()
            .instanceId("logs_instance_guid")
            .region("logs_instance_region")
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Example usage of resources:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language python %}}
```python
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language go %}}
```go
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language java %}}
```java
Example currently unavailable in this language
```
{{% /choosable %}}
{{< /chooser >}}
## Authentication

The IBM Cloud provider offers a flexible means of providing credentials for authentication. The following methods are supported, in this order, and explained below:

- Static credentials
- Environment variables
### Static credentials

You can provide your static credentials by adding the `ibmcloudApiKey`, `iaasClassicUsername`, and `iaasClassicApiKey` arguments in the IBM Cloud provider configuration.

Usage:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    ibm:iaasClassicApiKey:
        value: ""
    ibm:iaasClassicUsername:
        value: ""
    ibm:ibmcloudApiKey:
        value: ""

```
### Environment variables

You can provide your credentials by exporting the `IC_API_KEY`, `IAAS_CLASSIC_USERNAME`, and `IAAS_CLASSIC_API_KEY` environment variables, representing your IBM Cloud platform API key, IBM Cloud Classic Infrastructure (SoftLayer) user name, and IBM Cloud infrastructure API key, respectively.

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:

```

Usage:

```shell
export IC_API_KEY="ibmcloud_api_key"
export IAAS_CLASSIC_USERNAME="iaas_classic_username"
export IAAS_CLASSIC_API_KEY="iaas_classic_api_key"
pulumi preview
```
***Note:***
1. Create or find your `ibmcloudApiKey` and `iaasClassicApiKey` [here](https://cloud.ibm.com/iam/apikeys).
* Select `My IBM Cloud API Keys` option from view dropdown for `ibmcloudApiKey`
* Select `Classic Infrastructure API Keys` option from view dropdown for `iaasClassicApiKey`
2. For `iaasClassicUsername`
* Go to [Users](https://cloud.ibm.com/iam/users)
* Click on user.
* Find user name in the `VPN password` section under `User Details` tab
## Argument reference

The following configuration inputs are supported in the provider configuration:

* `ibmcloudApiKey` - (optional) The IBM Cloud platform API key. You must either add it as a credential in the provider configuration or source it from the `IC_API_KEY` (higher precedence) or `IBMCLOUD_API_KEY` environment variable. The key is required to provision Cloud Foundry or IBM Cloud Container Service resources, such as any resource that begins with `ibm` or `ibmContainer`. `ibmcloudApiKey` will have higher precedence than `bluemixApiKey`.

* `bluemixApiKey` - (deprecated, optional) The IBM Cloud platform API key. You must either add it as a credential in the provider configuration or source it from the `BM_API_KEY` (higher precedence) or `BLUEMIX_API_KEY` environment variable. The key is required to provision Cloud Foundry or IBM Cloud Container Service resources, such as any resource that begins with `ibm` or `ibmContainer`.

* `ibmcloudTimeout` - (optional) The timeout, expressed in seconds, for interacting with IBM Cloud APIs. You can also source the timeout from the `IC_TIMEOUT` (higher precedence) or `IBMCLOUD_TIMEOUT` environment variable. The default value is `60`. `ibmcloudTimeout` will have higher precedence than `bluemixTimeout`.

* `bluemixTimeout` - (deprecated, optional) The timeout, expressed in seconds, for interacting with IBM Cloud APIs. You can also source the timeout from the `BM_TIMEOUT` (higher precedence) or `BLUEMIX_TIMEOUT` environment variable. The default value is `60`.

* `softlayerUsername` - (deprecated, optional) The IBM Cloud Classic Infrastructure (SoftLayer) user name. You must either add it as a credential in the provider configuration or source it from the `SL_USERNAME` (higher precedence) or `SOFTLAYER_USERNAME` environment variable. `iaasClassicUsername` will have higher precedence than `softlayerUsername`.

* `iaasClassicUsername` - (optional) The IBM Cloud Classic Infrastructure (SoftLayer) user name. You must either add it as a credential in the provider configuration or source it from the `IAAS_CLASSIC_USERNAME`  environment variable.

* `softlayerApiKey` - (deprecated, optional) The IBM Cloud Classic Infrastructure API key. You must either add it as a credential in the provider configuration or source it from the `SL_API_KEY` (higher precedence) or `SOFTLAYER_API_KEY` environment variable. The key is required to provision infrastructure resources, such as any resource that begins with `ibmCompute`. `iaasClassicApiKey` will have higher precedence than `softlayerApiKey`.

* `iaasClassicApiKey` - (optional) The IBM Cloud Classic Infrastructure API key. You must either add it as a credential in the provider configuration or source it from the `IAAS_CLASSIC_API_KEY` environment variable.

* `softlayerEndpointUrl` - (deprecated, optional) The IBM Cloud Classic Infrastructure endpoint url. You can also source it from the `SL_ENDPOINT_URL` (higher precedence) or `SOFTLAYER_ENDPOINT_URL` environment variable. `iaasClassicEndpointUrl` will have higher precedence than `softlayerEndpointUrl`.

* `iaasClassicEndpointUrl` - (optional) The IBM Cloud Classic Infrastructure endpoint url. You can also source it from the `IAAS_CLASSIC_ENDPOINT_URL` environment variable. The default value is `https://api.softlayer.com/rest/v3`.

* `softlayerTimeout` - (optional) The timeout, expressed in seconds, for the IBM Cloud Classic Infrastructure APIs. You can also source the timeout from the `SL_TIMEOUT` (higher precedence) or `SOFTLAYER_TIMEOUT` environment variable. `iaasClassicTimeout` will have higher precedence than `softlayerTimeout`.

* `iaasClassicTimeout` - (optional) The timeout, expressed in seconds, for the IBM Cloud Clasic Infrastructure APIs. You can also source the timeout from the `IAAS_CLASSIC_TIMEOUT` environment variable. The default value is `60`.

* `region` - (optional) The IBM Cloud region. You can also source it from the `IC_REGION` (higher precedence) or `IBMCLOUD_REGION` `BM_REGION` `BLUEMIX_REGION` environment variable. The default value is `us-south`.

* `resourceGroup` - (optional) The Resource Group ID. You can also source it from the `IC_RESOURCE_GROUP` (higher precedence) or `IBMCLOUD_RESOURCE_GROUP` `BM_RESOURCE_GROUP` `BLUEMIX_RESOURCE_GROUP` environment variable.

* `maxRetries` - (Optional) This is the maximum number of times an IBM Cloud infrastructure API call is retried, in the case where requests are getting network related timeout and rate limit exceeded error code. You can also source it from the `MAX_RETRIES` environment variable. The default value is `10`.

* `functionNamespace` - (Optional) Your Cloud Functions namespace is composed from your IBM Cloud org and space like \<org\>_\<space\>. This attribute is required only when creating a Cloud Functions resource. It must be provided when you are creating such resources in IBM Cloud. You can also source it from the FUNCTION_NAMESPACE environment variable.

* `riaasEndpoint` - (deprected, Optional) The next generation infrastructure service API endpoint . It can also be sourced from the `RIAAS_ENDPOINT`. Default value: `us-south.iaas.cloud.ibm.com`.

* `generation` - (deprected, Optional) The generation is deprecated by default the provider targets to the IBM Cloud VPC infrastructure.

* `zone` - (optional) The IBM Cloud zone for a region. You can also source it from the `IC_ZONE` (higher precedence) or `IBMCLOUD_ZONE` environment variable. This value is required for power resources if the region supports multi-zone. For region `eu-de` it supports two zones `eu-de-1` and `eu-de-2`. Set the region and zone for the Power Virtual Server.

* `visibility` - (Optional) The visibility to IBM Cloud endpoint - `public`, `private`, `public-and-private`. Default value: `public`. Allowable values are `public`, `private`, `public-and-private`.
  * If visibility is set to `public`, use the regional public endpoint or global public endpoint. The regional public endpoints has higher precedence.
  * If visibility is set to `private`, use the regional private endpoint or global private endpoint. The regional private endpoint is given higher precedence.  In order to use the private endpoint from an IBM Cloud resource (such as, a classic VM instance), one must have VRF-enabled account.  If the Cloud service does not support private endpoint, the pulumi resource or function will log an error.
  * If visibility is set to `public-and-private`, use regional private endpoints or global private endpoint. If service doesn't support regional or global private endpoints it will use the regional or global public endpoint.
  * This can also be sourced from the `IC_VISIBILITY` (higher precedence) or `IBMCLOUD_VISIBILITY` environment variable.
* `privateEndpointType` - (Optional) Private Endpoint type used by the service endpoints. Allowable values are `vpe`.
  By default provider targets to cse endpoints when the `visibility` is set to `private`. If you want to target to vpe private endpoints, set `privateEndpointType` to `vpe`.
  * This can also be sourced from the `IC_PRIVATE_ENDPOINT_TYPE` (higher precedence) or `IBMCLOUD_PRIVATE_ENDPOINT_TYPE` environment variable.

***Note***
The CloudFoundry endpoint has been updated in this release of IBM Cloud Pulumi provider v0.17.4.  If you are using an earlier version of IBM Cloud Pulumi provider, export the `IBMCLOUD_UAA_ENDPOINT` to the new authentication endpoint, as illustrated below

```shell
export IBMCLOUD_UAA_ENDPOINT="https://iam.cloud.ibm.com/cloudfoundry/login/<region>/"
```
## References

* IBM Cloud Pulumi Docs