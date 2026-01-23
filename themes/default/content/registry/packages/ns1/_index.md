---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-ns1/v3.8.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi/pulumi-ns1/blob/v3.8.0/docs/_index.md
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Ns1 Provider
meta_desc: Provides an overview on how to configure the Pulumi Ns1 provider.
layout: package
---

## Installation

The Ns1 provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/ns1`](https://www.npmjs.com/package/@pulumi/ns1)
* Python: [`pulumi-ns1`](https://pypi.org/project/pulumi-ns1/)
* Go: [`github.com/pulumi/pulumi-ns1/sdk/v3/go/ns1`](https://github.com/pulumi/pulumi-ns1)
* .NET: [`Pulumi.Ns1`](https://www.nuget.org/packages/Pulumi.Ns1)
* Java: [`com.pulumi/ns1`](https://central.sonatype.com/artifact/com.pulumi/ns1)

## Overview

The NS1 provider exposes resources to interact with the NS1 REST API. The
provider needs to be configured with the proper credentials before it can be
used. Note also that for a given resource to function, the API key used must
have the corresponding permissions set.

Use the navigation to the left to read about the available resources.
## Example Usage

Additional usage examples can be found in <https://github.com/ns1-pulumi/pulumi-provider-ns1/tree/master/examples>

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    ns1:apikey:
        value: 'TODO: var.ns1_apikey'
    ns1:rateLimitParallelism:
        value: 60

```

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as ns1 from "@pulumi/ns1";

// Create a new zone
const foobar = new ns1.Zone("foobar", {});
```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    ns1:apikey:
        value: 'TODO: var.ns1_apikey'
    ns1:rateLimitParallelism:
        value: 60

```

```python
import pulumi
import pulumi_ns1 as ns1

# Create a new zone
foobar = ns1.Zone("foobar")
```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    ns1:apikey:
        value: 'TODO: var.ns1_apikey'
    ns1:rateLimitParallelism:
        value: 60

```

```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Ns1 = Pulumi.Ns1;

return await Deployment.RunAsync(() =>
{
    // Create a new zone
    var foobar = new Ns1.Zone("foobar");

});

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    ns1:apikey:
        value: 'TODO: var.ns1_apikey'
    ns1:rateLimitParallelism:
        value: 60

```

```go
package main

import (
	"github.com/pulumi/pulumi-ns1/sdk/v3/go/ns1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a new zone
		_, err := ns1.NewZone(ctx, "foobar", nil)
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
    ns1:apikey:
        value: 'TODO: var.ns1_apikey'
    ns1:rateLimitParallelism:
        value: 60

```

```yaml
resources:
  # Create a new zone
  foobar:
    type: ns1:Zone
```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    ns1:apikey:
        value: 'TODO: var.ns1_apikey'
    ns1:rateLimitParallelism:
        value: 60

```

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.ns1.Zone;
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
        // Create a new zone
        var foobar = new Zone("foobar");

    }
}
```

{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

The following configuration inputs are supported:

* `apikey` - (Required) NS1 API token. It must be provided, but it can also
  be sourced from the `NS1_APIKEY` environment variable.
* `version` - (Optional, but recommended if you don't like surprises) Require a specific version of the NS1 provider. Run `pulumi up` to get your current version.
* `retryMax` - (Optional, introduced in v1.13.2) Sets the number of retries for 50x-series errors. The default is 3. A negative value such as -1 disables this feature.
* `endpoint` - (Optional) NS1 API endpoint. Normally not set unless testing or using non-standard proxies.
* `ignoreSsl` - (Optional) This normally does not need to be set.
* `enableDdi` - (Deprecated) Enable the DDI-compatible permissions schema. No longer in use.
* `userAgent` - (Optional, introduced in v1.13.4) Sets the User-Agent header in the NS1 API.
* `rateLimitParallelism` - (Optional) Integer for alternative rate limit and parallelism strategy.
  NS1 uses a token-based method for rate limiting API requests. Full details can be found at <https://help.ns1.com/hc/en-us/articles/360020250573-About-API-rate-limiting>.

  By default, the NS1 provider uses the "sleep" strategy of the underlying [NS1 Go SDK](https://github.com/ns1/ns1-go) for handling the NS1 API rate limit:
  an operation waits after every API request for a time equal to the rate limit period of that request type divided by the corresponding tokens remaining.

  Furthermore, the default behaviour of Pulumi uses ten concurrent operations.
  This means that the provider will burst through available request tokens, gradually slowing until it reaches an equilibrium point where the ten operations wait long enough between requests to replenish ten tokens.
  However, if there are other concurrent uses of the API this can lead to the tokens being entirely depleted when a Pulumi operation makes a new request.
  This results in a 429 response which will cause the entire Pulumi process to fail.

  If you encounter this scenario, or believe you are likely to, then you can set the `rateLimitParallelism` to enable an alternative rate limiting strategy.
  Here the Pulumi operations will burst through all available tokens until they reach a point where the remaining limit is less, or equal, to the value set;
  after this point an operation will wait for the time it would take to replenish an equal number of tokens.

  Setting this to a value of 60 represents a good balance between optimising for performance and reducing the risk of a 429 response.
  If you still encounter issues then you can increase this value: we would recommend you do so in increments of 20.

  Note: We recommend that you NOT set the Pulumi command line `-parallelism=n` option when you run `pulumi up`.
  The default value of ten is sufficient - increasing it will lead to a greater risk of encountering a 429 response.
## Environment Variables

The provider honors the following environment variables for its configuration
variables if they are not specified in the configuration:

* `NS1_APIKEY` - (string) Explained above.
* `NS1_ENDPOINT` - (string) Explained above.
* `NS1_RETRY_MAX` - (integer) Explained above.
* `NS1_IGNORE_SSL` - (boolean) If set, follows the convention of
  [strconv.ParseBool](https://golang.org/pkg/strconv/#ParseBool).
* `NS1_RATE_LIMIT_PARALLELISM` - (integer) Explained above.
* `NS1_TF_USER_AGENT` - (string) Sets the User-Agent header in the NS1 API.