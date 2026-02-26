---
# WARNING: this file was fetched from https://raw.githubusercontent.com/AdrianSilaghi/pulumi-danubedata/v0.1.7/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/AdrianSilaghi/pulumi-danubedata/blob/v0.1.7/docs/_index.md
title: DanubeData
meta_desc: Provides an overview of the DanubeData Provider for Pulumi.
layout: package
---

The DanubeData provider for Pulumi can be used to provision and manage [DanubeData](https://danubedata.ro) cloud infrastructure resources. DanubeData is a managed infrastructure provider offering VPS instances, cache services, databases, object storage, and serverless containers.

## Example

{{< chooser language "typescript,python,go,csharp" >}}

{{% choosable language typescript %}}

```typescript
import * as danubedata from "@danubedata/pulumi";

// Create a storage bucket
const bucket = new danubedata.StorageBucket("my-bucket", {
    region: "fsn1",
    versioningEnabled: true,
});

// Create a Redis cache instance
const cache = new danubedata.Cache("my-cache", {
    name: "app-cache",
    cacheProvider: "redis",
    resourceProfile: "small",
    datacenter: "fsn1",
});

// Create a PostgreSQL database
const database = new danubedata.Database("my-database", {
    name: "app-db",
    databaseProvider: "postgresql",
    providerVersion: "16",
    resourceProfile: "small",
    datacenter: "fsn1",
});

// Export endpoints
export const bucketEndpoint = bucket.endpointUrl;
export const cacheEndpoint = cache.endpoint;
export const databaseEndpoint = database.endpoint;
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi_danubedata as danubedata

# Create a storage bucket
bucket = danubedata.StorageBucket("my-bucket",
    region="fsn1",
    versioning_enabled=True)

# Create a Redis cache instance
cache = danubedata.Cache("my-cache",
    name="app-cache",
    cache_provider="redis",
    resource_profile="small",
    datacenter="fsn1")

# Create a PostgreSQL database
database = danubedata.Database("my-database",
    name="app-db",
    database_provider="postgresql",
    provider_version="16",
    resource_profile="small",
    datacenter="fsn1")

# Export endpoints
pulumi.export("bucket_endpoint", bucket.endpoint_url)
pulumi.export("cache_endpoint", cache.endpoint)
pulumi.export("database_endpoint", database.endpoint)
```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"github.com/AdrianSilaghi/pulumi-danubedata/sdk/go/danubedata"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a storage bucket
		bucket, err := danubedata.NewStorageBucket(ctx, "my-bucket", &danubedata.StorageBucketArgs{
			Region:            pulumi.String("fsn1"),
			VersioningEnabled: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}

		// Create a Redis cache instance
		cache, err := danubedata.NewCache(ctx, "my-cache", &danubedata.CacheArgs{
			Name:            pulumi.String("app-cache"),
			CacheProvider:   pulumi.String("redis"),
			ResourceProfile: pulumi.String("small"),
			Datacenter:      pulumi.String("fsn1"),
		})
		if err != nil {
			return err
		}

		// Export endpoints
		ctx.Export("bucketEndpoint", bucket.EndpointUrl)
		ctx.Export("cacheEndpoint", cache.Endpoint)
		return nil
	})
}
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using Pulumi;
using DanubeData = DanubeData.DanubeData;

return await Deployment.RunAsync(() =>
{
    // Create a storage bucket
    var bucket = new DanubeData.StorageBucket("my-bucket", new()
    {
        Region = "fsn1",
        VersioningEnabled = true,
    });

    // Create a Redis cache instance
    var cache = new DanubeData.Cache("my-cache", new()
    {
        Name = "app-cache",
        CacheProvider = "redis",
        ResourceProfile = "small",
        Datacenter = "fsn1",
    });

    // Export endpoints
    return new Dictionary<string, object?>
    {
        ["bucketEndpoint"] = bucket.EndpointUrl,
        ["cacheEndpoint"] = cache.Endpoint,
    };
});
```

{{% /choosable %}}

{{< /chooser >}}

## Resources

The DanubeData provider supports the following resources:

### Compute
- **Vps** - KubeVirt-based virtual machines with IPv4/IPv6 support
- **Serverless** - Knative-based serverless containers with scale-to-zero

### Data Services
- **Cache** - In-memory data stores (Redis, Valkey, Dragonfly)
- **Database** - Managed databases (MySQL, PostgreSQL, MariaDB)

### Storage
- **StorageBucket** - S3-compatible object storage
- **StorageAccessKey** - Access keys for object storage

### Security
- **SshKey** - SSH public keys for VPS access
- **Firewall** - Network firewall rules

### Backup
- **VpsSnapshot** - Point-in-time snapshots of VPS instances
