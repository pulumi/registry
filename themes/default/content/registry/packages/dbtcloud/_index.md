---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-dbtcloud/v1.4.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi/pulumi-dbtcloud/blob/v1.4.0/docs/_index.md
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: dbt Cloud Provider
meta_desc: Provides an overview on how to configure the Pulumi dbt Cloud provider.
layout: package
---

## Installation

The dbt Cloud provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/dbtcloud`](https://www.npmjs.com/package/@pulumi/dbtcloud)
* Python: [`pulumi-dbtcloud`](https://pypi.org/project/pulumi-dbtcloud/)
* Go: [`github.com/pulumi/pulumi-dbtcloud/sdk/go/dbtcloud`](https://github.com/pulumi/pulumi-dbtcloud)
* .NET: [`Pulumi.Dbtcloud`](https://www.nuget.org/packages/Pulumi.Dbtcloud)
* Java: [`com.pulumi/dbtcloud`](https://central.sonatype.com/artifact/com.pulumi/dbtcloud)

## Provider configuration

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    dbtcloud:accountId:
        value: 'TODO: var.dbt_cloud_account_id'
    dbtcloud:disableRetry:
        value: false
    dbtcloud:hostUrl:
        value: https://cloud.getdbt.com/api
    dbtcloud:maxRetries:
        value: 10
    dbtcloud:retriableStatusCodes:
        value:
            - "429"
            - "500"
            - "502"
            - "503"
            - "504"
    dbtcloud:retryIntervalSeconds:
        value: 10
    dbtcloud:skipCredentialsValidation:
        value: false
    dbtcloud:timeoutSeconds:
        value: 60
    dbtcloud:token:
        value: 'TODO: var.dbt_cloud_token'

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
config:
    dbtcloud:accountId:
        value: 'TODO: var.dbt_cloud_account_id'
    dbtcloud:disableRetry:
        value: false
    dbtcloud:hostUrl:
        value: https://cloud.getdbt.com/api
    dbtcloud:maxRetries:
        value: 10
    dbtcloud:retriableStatusCodes:
        value:
            - "429"
            - "500"
            - "502"
            - "503"
            - "504"
    dbtcloud:retryIntervalSeconds:
        value: 10
    dbtcloud:skipCredentialsValidation:
        value: false
    dbtcloud:timeoutSeconds:
        value: 60
    dbtcloud:token:
        value: 'TODO: var.dbt_cloud_token'

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
config:
    dbtcloud:accountId:
        value: 'TODO: var.dbt_cloud_account_id'
    dbtcloud:disableRetry:
        value: false
    dbtcloud:hostUrl:
        value: https://cloud.getdbt.com/api
    dbtcloud:maxRetries:
        value: 10
    dbtcloud:retriableStatusCodes:
        value:
            - "429"
            - "500"
            - "502"
            - "503"
            - "504"
    dbtcloud:retryIntervalSeconds:
        value: 10
    dbtcloud:skipCredentialsValidation:
        value: false
    dbtcloud:timeoutSeconds:
        value: 60
    dbtcloud:token:
        value: 'TODO: var.dbt_cloud_token'

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
config:
    dbtcloud:accountId:
        value: 'TODO: var.dbt_cloud_account_id'
    dbtcloud:disableRetry:
        value: false
    dbtcloud:hostUrl:
        value: https://cloud.getdbt.com/api
    dbtcloud:maxRetries:
        value: 10
    dbtcloud:retriableStatusCodes:
        value:
            - "429"
            - "500"
            - "502"
            - "503"
            - "504"
    dbtcloud:retryIntervalSeconds:
        value: 10
    dbtcloud:skipCredentialsValidation:
        value: false
    dbtcloud:timeoutSeconds:
        value: 60
    dbtcloud:token:
        value: 'TODO: var.dbt_cloud_token'

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
config:
    dbtcloud:accountId:
        value: 'TODO: var.dbt_cloud_account_id'
    dbtcloud:disableRetry:
        value: false
    dbtcloud:hostUrl:
        value: https://cloud.getdbt.com/api
    dbtcloud:maxRetries:
        value: 10
    dbtcloud:retriableStatusCodes:
        value:
            - "429"
            - "500"
            - "502"
            - "503"
            - "504"
    dbtcloud:retryIntervalSeconds:
        value: 10
    dbtcloud:skipCredentialsValidation:
        value: false
    dbtcloud:timeoutSeconds:
        value: 60
    dbtcloud:token:
        value: 'TODO: var.dbt_cloud_token'

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
config:
    dbtcloud:accountId:
        value: 'TODO: var.dbt_cloud_account_id'
    dbtcloud:disableRetry:
        value: false
    dbtcloud:hostUrl:
        value: https://cloud.getdbt.com/api
    dbtcloud:maxRetries:
        value: 10
    dbtcloud:retriableStatusCodes:
        value:
            - "429"
            - "500"
            - "502"
            - "503"
            - "504"
    dbtcloud:retryIntervalSeconds:
        value: 10
    dbtcloud:skipCredentialsValidation:
        value: false
    dbtcloud:timeoutSeconds:
        value: 60
    dbtcloud:token:
        value: 'TODO: var.dbt_cloud_token'

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
## Configuration Reference

- `accountId` (Number) Account identifier for your dbt Cloud implementation. Instead of setting the parameter, you can set the environment variable `DBT_CLOUD_ACCOUNT_ID`
- `disableRetry` (Boolean) If set to true, the provider will not retry requests that fail due to rate limiting. Defaults to false.
- `hostUrl` (String) URL for your dbt Cloud deployment. Instead of setting the parameter, you can set the environment variable `DBT_CLOUD_HOST_URL` - Defaults to <https://cloud.getdbt.com/api>
- `maxRetries` (Number) The maximum number of retries to attempt for requests that fail due to rate limiting. Defaults to 3 retries.
- `retriableStatusCodes` (List of String) List of HTTP status codes that should be retried when encountered. Defaults to [429, 500, 502, 503, 504].
- `retryIntervalSeconds` (Number) The number of seconds to wait before retrying a request that failed due to rate limiting. Defaults to 10 seconds.
- `skipCredentialsValidation` (Boolean) If set to true, the provider will not validate credentials during initialization. This can be useful for testing and for dbt Cloud API implementations that do not have standard authentication available. Defaults to false.
- `timeoutSeconds` (Number) The timeout duration in seconds for HTTP requests to the dbt Cloud API. Defaults to 30 seconds.
- `token` (String, Sensitive) API token for your dbt Cloud. Instead of setting the parameter, you can set the environment variable `DBT_CLOUD_TOKEN`