---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/qdrant/qdrant-cloud/1.3.1/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Qdrant-Cloud Provider
meta_desc: Provides an overview on how to configure the Pulumi Qdrant-Cloud provider.
layout: package
---

## Generate Provider

The Qdrant-Cloud provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider qdrant/qdrant-cloud
```
## Overview

The Pulumi Qdrant Cloud provider is a plugin for Pulumi that allows for the full lifecycle management of Qdrant Cloud resources.
Note that it's not intended to manage the content of the database itself.

Below is a sample that creates a [vector database] cluster and a token for accessing the cluster.

The ID and version of the cluster as well the URL (endpoint of the database cluster) are displayed (as output).

The access token created is scoped for the created cluster and is displayed as well.
Note that this token should be kept secret, with this token the database can be manipulated and viewed (CRUD operations).

To view the cluster itself, please visit (in a web-browser)
`<cluster_url>:6333/dashboard`
This will ask the token to grant you access.

The url and token can be used in client libraries as well

For more info, please visit <https://qdrant.tech/>

---
## Versioning and Compatibility

Please always use the latest version of the provider.

Versions below `1.1.0` are deprecated and should not be used.

---
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    qdrant-cloud:accountId:
        value: ""
    qdrant-cloud:apiKey:
        value: ""
    qdrant-cloud:apiUrl:
        value: ""

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as qdrant_cloud from "@pulumi/qdrant-cloud";

const example = new qdrant_cloud.index/accountsCluster.AccountsCluster("example", {
    name: "tf-example-cluster",
    cloudProvider: "gcp",
    cloudRegion: "us-east4",
    configuration: {
        numberOfNodes: 1,
        nodeConfiguration: {
            packageId: "7c939d96-d671-4051-aa16-3b8b7130fa42",
        },
    },
});
const example_key = new qdrant_cloud.index/accountsAuthKey.AccountsAuthKey("example-key", {clusterIds: [example.id]});
export const clusterId = example.id;
export const url = example.url;
export const token = example_key.token;
export const curlCommand = `curl \
    -X GET '${example.url}' \
    --header 'api-key: ${example_key.token}'`;
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    qdrant-cloud:accountId:
        value: ""
    qdrant-cloud:apiKey:
        value: ""
    qdrant-cloud:apiUrl:
        value: ""

```
```python
import pulumi
import pulumi_qdrant_cloud as qdrant_cloud

example = qdrant_cloud.index.accounts_cluster.AccountsCluster("example",
    name=tf-example-cluster,
    cloud_provider=gcp,
    cloud_region=us-east4,
    configuration={
        numberOfNodes: 1,
        nodeConfiguration: {
            packageId: 7c939d96-d671-4051-aa16-3b8b7130fa42,
        },
    })
example_key = qdrant_cloud.index.accounts_auth_key.AccountsAuthKey("example-key", cluster_ids=[example.id])
pulumi.export("clusterId", example["id"])
pulumi.export("url", example["url"])
pulumi.export("token", example_key["token"])
pulumi.export("curlCommand", f"""curl \
    -X GET '{example["url"]}' \
    --header 'api-key: {example_key["token"]}'""")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    qdrant-cloud:accountId:
        value: ""
    qdrant-cloud:apiKey:
        value: ""
    qdrant-cloud:apiUrl:
        value: ""

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using QdrantCloud = Pulumi.QdrantCloud;

return await Deployment.RunAsync(() =>
{
    var example = new QdrantCloud.Index.AccountsCluster.AccountsCluster("example", new()
    {
        Name = "tf-example-cluster",
        CloudProvider = "gcp",
        CloudRegion = "us-east4",
        Configuration =
        {
            { "numberOfNodes", 1 },
            { "nodeConfiguration",
            {
                { "packageId", "7c939d96-d671-4051-aa16-3b8b7130fa42" },
            } },
        },
    });

    var example_key = new QdrantCloud.Index.AccountsAuthKey.AccountsAuthKey("example-key", new()
    {
        ClusterIds = new[]
        {
            example.Id,
        },
    });

    return new Dictionary<string, object?>
    {
        ["clusterId"] = example.Id,
        ["url"] = example.Url,
        ["token"] = example_key.Token,
        ["curlCommand"] = @$"curl \
    -X GET '{example.Url}' \
    --header 'api-key: {example_key.Token}'",
    };
});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    qdrant-cloud:accountId:
        value: ""
    qdrant-cloud:apiKey:
        value: ""
    qdrant-cloud:apiUrl:
        value: ""

```
```go
package main

import (
	"fmt"

	qdrantcloud "github.com/pulumi/pulumi-qdrant-cloud/sdk/go/qdrant-cloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := index / accountscluster.NewAccountsCluster(ctx, "example", &index/accountscluster.AccountsClusterArgs{
			Name:          "tf-example-cluster",
			CloudProvider: "gcp",
			CloudRegion:   "us-east4",
			Configuration: map[string]interface{}{
				"numberOfNodes": 1,
				"nodeConfiguration": map[string]interface{}{
					"packageId": "7c939d96-d671-4051-aa16-3b8b7130fa42",
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = index / accountsauthkey.NewAccountsAuthKey(ctx, "example-key", &index/accountsauthkey.AccountsAuthKeyArgs{
			ClusterIds: []interface{}{
				example.Id,
			},
		})
		if err != nil {
			return err
		}
		ctx.Export("clusterId", example.Id)
		ctx.Export("url", example.Url)
		ctx.Export("token", example_key.Token)
		ctx.Export("curlCommand", pulumi.Sprintf("curl \\\n    -X GET '%v' \\\n    --header 'api-key: %v'", example.Url, example_key.Token))
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
    qdrant-cloud:accountId:
        value: ""
    qdrant-cloud:apiKey:
        value: ""
    qdrant-cloud:apiUrl:
        value: ""

```
```yaml
resources:
  example:
    type: qdrant-cloud:AccountsCluster
    properties:
      name: tf-example-cluster
      cloudProvider: gcp
      cloudRegion: us-east4
      configuration:
        numberOfNodes: 1
        nodeConfiguration:
          packageId: 7c939d96-d671-4051-aa16-3b8b7130fa42
  example-key:
    type: qdrant-cloud:AccountsAuthKey
    properties:
      clusterIds:
        - ${example.id}
outputs:
  clusterId: ${example.id}
  url: ${example.url}
  token: ${["example-key"].token}
  curlCommand: |-
    curl \
        -X GET '${example.url}' \
        --header 'api-key: ${["example-key"].token}'
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    qdrant-cloud:accountId:
        value: ""
    qdrant-cloud:apiKey:
        value: ""
    qdrant-cloud:apiUrl:
        value: ""

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.qdrantcloud.AccountsCluster;
import com.pulumi.qdrantcloud.AccountsClusterArgs;
import com.pulumi.qdrantcloud.AccountsAuthKey;
import com.pulumi.qdrantcloud.AccountsAuthKeyArgs;
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
        var example = new AccountsCluster("example", AccountsClusterArgs.builder()
            .name("tf-example-cluster")
            .cloudProvider("gcp")
            .cloudRegion("us-east4")
            .configuration(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
            .build());

        var example_key = new AccountsAuthKey("example-key", AccountsAuthKeyArgs.builder()
            .clusterIds(example.id())
            .build());

        ctx.export("clusterId", example.id());
        ctx.export("url", example.url());
        ctx.export("token", example_key.token());
        ctx.export("curlCommand", """
curl \
    -X GET '%s' \
    --header 'api-key: %s'", example.url(),example_key.token()));
    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference
### Required

- `apiKey` (String) The API Key for Qdrant Cloud API operations.

- `accountId` (String) Default Account Identifier for the Qdrant cloud
- `apiUrl` (String) The URL of the Qdrant Cloud API.