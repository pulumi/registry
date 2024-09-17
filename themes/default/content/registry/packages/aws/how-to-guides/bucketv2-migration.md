---
title: "BucketV2 Migration Guide"
h1: "Migrating from aws.s3.Bucket to aws.s3.BucketV2"
meta_desc: Practitioner level instructions for migrating from aws.s3.Bucket to aws.s3.BucketV2 resources
layout: package
---

## Migrating from aws.s3.Bucket to aws.s3.BucketV2

In the upcoming AWS Classic major release (v7), `aws.s3.Bucket` will be discontinued in favor of `BucketV2`. Users of
`aws.s3.Bucket` resource are encouraged to migrate early. The migration is straightforward for simple use cases but
requires some care for advanced configurations. Specifically `BucketV2` inputs such as `policy` and `accelerationStatus`
are deprecated and using the requires migrating to side-by-side resources (see Removed inputs). This guide aims to cover
all relevant scenarios with precise migration instructions.

To migrate existing `aws.s3.Bucket` resources to `aws.s3.BucketV2`:

1. Edit your Pulumi program sources to replace removed inputs (see below) with equivalent side-by-side resources.

2. Perform `pulumi up` to replace (delete and re-create) the buckets in AWS.

If replacement is not acceptable, it is possible to perform a manual migration with `pulumi import` (see Avoiding
replacement).

## Migrating deprecated inputs

Inputs such as `policy` and `accelerationStatus` are deprecated on the `aws.s3.BucketV2` resource. Instead of using
these inputs, it is recommended to use side-by-side resources to achieve the same effect.

### policy input

To specify a
[bucket policy](https://docs.aws.amazon.com/AmazonS3/latest/dev/example-bucket-policies.html),
instead of:

``` typescript
new aws.s3.Bucket("my-bucket", {
  bucket: "my-bucket-26224916",
  policy: JSON.stringify({
    "Version": "2012-10-17",
    "Id": "PutObjPolicy",
    "Statement": [
      {
        "Sid": "DenyObjectsThatAreNotSSEKMS",
        "Principal": "*",
        "Effect": "Deny",
        "Action": "s3:PutObject",
        "Resource": "arn:aws:s3:::my-bucket-26224916/*",
        "Condition": {
          "Null": {
            "s3:x-amz-server-side-encryption-aws-kms-key-id": "true"
          }
        }
      }
    ]
  }),
});
```

Use the `aws.s3.BucketPolicy` resource:

``` typescript
const myBucket = new aws.s3.BucketV2("my-new-bucket", {});

new aws.s3.BucketPolicy("my-bucket-policy", {
  bucket: myBucket.bucket,
  policy: myBucket.bucket.apply(bucket => JSON.stringify({
    "Version": "2012-10-17",
    "Id": "PutObjPolicy",
    "Statement": [
      {
        "Sid": "DenyObjectsThatAreNotSSEKMS",
        "Principal": "*",
        "Effect": "Deny",
        "Action": "s3:PutObject",
        "Resource": `arn:aws:s3:::${bucket}/*`,
        "Condition": {
          "Null": {
            "s3:x-amz-server-side-encryption-aws-kms-key-id": "true"
          }
        }
      }
    ]
  })),
});
```

As a bonus, the policy can now more easily refer to the concrete name of the bucket.

### serverSideEncryptionConfiguration input

To specify [server-side encryption configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html),
instead of:

``` typescript
new aws.s3.Bucket("my-bucket", {
  serverSideEncryptionConfiguration: {
    rule: {
      applyServerSideEncryptionByDefault: {
        sseAlgorithm: "aws:kms",
      },
    },
  },
});
```

Use the `aws.s3.BucketServerSideEncryptionConfiguration` resource:

``` typescript
const myBucket = new aws.s3.BucketV2("my-new-bucket", {});

new aws.s3.BucketServerSideEncryptionConfigurationV2("my-new-bucket-sse-config", {
  bucket: myBucket.bucket,
  rules: [{
    applyServerSideEncryptionByDefault: {
      sseAlgorithm: "aws:kms"
    },
  }],
});
```

### acceleration input

To enable acceleration, instead of:

``` typescript
new aws.s3.Bucket("my-bucket", {
  accelerationStatus: "Enabled",
});
```

Use the `aws.s3.BucketAccelerationConfiguration` resource:

``` typescript
const myBucket = new aws.s3.BucketV2("my-new-bucket", {});

new aws.s3.BucketAccelerateConfigurationV2("my-bucket-acceleration", {
  bucket: myBucket.bucket,
  status: "Enabled",
});
```

### corsRules input

To configure [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html),
instead of:

``` typescript
new aws.s3.Bucket("my-bucket", {
  corsRules: [
    {
      allowedHeaders: ["*"],
      allowedMethods: ["GET"],
      allowedOrigins: ["*"],
      exposeHeaders: ["ETag"],
      maxAgeSeconds: 3000,
    },
  ],
});
```

Use the `aws.s3.BucketCorsConfiguration` resource:

``` typescript
const myBucket = new aws.s3.BucketV2("my-new-bucket", {});

new aws.s3.BucketCorsConfigurationV2("my-bucket-cors", {
  bucket: myBucket.bucket,
  corsRules: [
    {
      allowedHeaders: ["*"],
      allowedMethods: ["GET"],
      allowedOrigins: ["*"],
      exposeHeaders: ["ETag"],
      maxAgeSeconds: 3000,
    },
  ]
});
```

### acl and grants inputs

The `aws.s3.BucketAcl` resource is now required to configure either
[canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)
or [ACL policy grant](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#sample-acl),
and it replaces the use of `acl` and `grants` inputs. To illustrate migrating the uses of `grants`, instead of:

``` typescript
const currentUser = aws.s3.getCanonicalUserIdOutput({});

new aws.s3.Bucket("my-bucket", {
  grants: [
    {
      permissions: ["READ"],
      type: "CanonicalUser",
      id: currentUser.id,
    },
  ],
});
```

Use the `aws.s3.BucketAcl` resource like this:

``` typescript
const myBucket = new aws.s3.BucketV2("my-new-bucket", {});

const myBucketOwnershipControls = new aws.s3.BucketOwnershipControls("my-new-bucket-oc", {
  bucket: myBucket.id,
  rule: {
    objectOwnership: "BucketOwnerPreferred",
  },
});

new aws.s3.BucketAclV2("grant", {
  bucket: myBucket.bucket,
  accessControlPolicy: {
    owner: {id: currentUser.id},
    grants: [
      {
        permission: "READ",
        grantee: {
          type: "CanonicalUser",
          id: currentUser.id,
        },
      },
    ],
  },
}, {
  dependsOn: [myBucketOwnershipControls]
});
```

Note that `dependsOn` on `BucketOwnershipControls` is required for Pulumi to correctly order the operations for this
infrastructure. Consult the documentation on `aws.s3.BucketAclV2` for more fully worked examples of configuring ACL.

### hostedZoneId input

This will only be available as an output. If your program specified this as an input, simply remove it. The input was
ignored by prior versions of the provider and was exposed in error.

### other inputs

All other removed inputs are treated similarly. Consult the documentation for the matching resource replacing the input
for more information:

| Input                    | Resource                                   |
|--------------------------|--------------------------------------------|
| lifecycleRules           | aws.s3.BucketLifecycleConfigurationV2      |
| loggings                 | aws.s3.BucketLoggingV2                     |
| objectLockConfiguration  | aws.s3.BucketObjectLockConfigurationV2     |
| replicationConfiguration | aws.s3.BucketReplicationConfig             |
| requestPayer             | aws.s3.BucketRequestPaymentConfigurationV2 |
| versionings              | aws.s3.BucketVersioningV2                  |
| website                  | aws.s3.BucketWebsiteConfigurationV2        |
| websiteDomain            | aws.s3.BucketWebsiteConfigurationV2        |
| websiteEndpoint          | aws.s3.BucketWebsiteConfigurationV2        |


## Avoiding replacement

In situations when replacing the bucket in the AWS account is not acceptable, it is possible to perform a manual
migration that changes Pulumi program and state to track buckets using the new resource without executing any changes
against the actual cloud account. While the details will vary depending on your use case, this procedure generally
involves the following steps:

1. Find URNs for legacy Bucket Pulumi resources using `pulumi stack export`
2. Determine the actual bucket name(s)
3. Determine which side-by-side resources will be needed for each bucket
4. Construct an `pulumi-import.json` file listing the buckets and their side-by-side resources
5. Run `pulumi import --file import-file.json` using the [Bulk Importing](https://www.pulumi.com/tutorials/importing/bulk-importing/) feature
6. Add the suggested code into your Pulumi program source
7. Remove the legacy Bucket code from your Pulumi program source
8. Remove the legacy Bucket resources from state using `pulumi state delete $bucketURN`
9. Run `pulumi preview` to confirm a no-change plan

Consider a concrete example, suppose you have provisioned a bucket with a `serverSideEncryptionConfiguration` as
follows:

```typescript
import * as aws from "@pulumi/aws";

new aws.s3.Bucket("my-bucket", {
  serverSideEncryptionConfiguration: {
    rule: {
      applyServerSideEncryptionByDefault: {
        sseAlgorithm: "aws:kms",
      },
    },
  },
});
```

1. Scanning through the state in `pulumi stack export`, observe and note its URN is
   `"urn:pulumi:dev::bucket-playground::aws:s3/bucket:Bucket::my-bucket"`

2. The state file should also include the actual cloud name for the bucket such as `"bucket": "my-bucket-cd24744"`

3. This bucket will require a `BucketServerSideEncryptionConfiguration` side-by-side resource

4. The import file should therefore look like this:

    ```json
    {
      "resources": [
        {
          "type": "aws:s3/bucketV2:BucketV2",
          "name": "my-bucket",
          "id": "my-bucket-cd24744"
        },
        {
          "type": "aws:s3/bucketServerSideEncryptionConfigurationV2:BucketServerSideEncryptionConfigurationV2",
          "name": "my-bucket-encryption-configuration",
          "id": "my-bucket-cd24744"
        }
      ]
    }
    ```

5. `pulumi import --file import-file.json` will suggest new code to include in your program, for example:

    ```typescript
    const my_bucket = new aws.s3.BucketV2("my-bucket", {
        bucket: "my-bucket-cd24744",
    }, {
        protect: true,
    });

    const my_bucket_encryption_configuration = new aws.s3.BucketServerSideEncryptionConfigurationV2("my-bucket-encryption-configuration", {
        bucket: "my-bucket-cd24744",
        rules: [{
            applyServerSideEncryptionByDefault: {
                sseAlgorithm: "aws:kms",
            },
        }],
    }, {
        protect: true,
    });
    ```

6. Run `pulumi state delete "urn:pulumi:dev::bucket-playground::aws:s3/bucket:Bucket::my-bucket"` to remove the old
   bucket from the state

7. Delete the code for the old bucket from the sources.

8. `pulumi preview` should result in `Resources: N unchanged` to confirm everything went well.
