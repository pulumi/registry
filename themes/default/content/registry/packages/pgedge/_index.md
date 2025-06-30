---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pgEdge/pulumi-pgedge/v0.0.41/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: pgEdge
meta_desc: Provides an overview of the pgEdge Provider for Pulumi.
layout: package
---

# pgEdge Provider

The official Pulumi provider for [pgEdge Cloud](https://www.pgedge.com/cloud), designed to simplify the management of pgEdge Cloud resources for both **Developers** and **Enterprise** edition.

## Authentication

Before using the provider, you must create an API Client in [pgEdge Cloud](https://www.pgedge.com/cloud) and configure the following environment variables:

```sh
export PGEDGE_CLIENT_ID="your-client-id"
export PGEDGE_CLIENT_SECRET="your-client-secret"
```
These credentials authenticate the Pulumi provider with your pgEdge Cloud account

## Example Usage

### Developer Edition Configuration

For Developer Edition, pgEdge offers access to manage databases. Here's an example setup for Developer Edition:

{{< chooser language "typescript,go" >}}
{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as pgedge from "@pgEdge/pulumi-pgedge";

// Define a database
const defaultdb = new pgedge.Database("defaultdb", {
    name: "defaultdb",
    clusterId: "f12239ddq-df9d-4ded-adqwead9-3e2bvhe6d6ee",
    options: [
        "rest:enabled",
        "install:northwind"
    ]
});
```

{{% /choosable %}}
{{< /chooser >}}

## Enterprise Edition Configuration

Enterprise Edition users can manage Cloud Accounts, SSH keys, Backup Stores, and Clusters. Here's an Enterprise Edition example:

{{< chooser language "typescript,go" >}}
{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as pgedge from "@pgEdge/pulumi-pgedge";

// Create an SSH Key
const sshKey = new pgedge.SSHKey("exampleSSHKey", {
  name: "example",
  publicKey: "ssh-ed25519 AAAAC3NzaC1lZsdw877237ICXfT63i04t5fvvlGesddwed21VG7DkyxvyXbYQNhKP/rSeLY user@example.com",
});

// Create a Cloud Account
const cloudAccount = new pgedge.CloudAccount("exampleCloudAccount", {
  name: "my-aws-account",
  type: "aws",
  description: "My AWS Cloud Account",
  credentials: {
    role_arn: "arn:aws:iam::21112529deae39:role/pgedge-135232c",
  },
}, { dependsOn: sshKey });

// Create a Backup Store
const backupStore = new pgedge.BackupStore("exampleBackupStore", {
  name: "example",
  cloudAccountId: cloudAccount.id,
  region: "us-west-2",
}, { dependsOn: cloudAccount });

// Create a Cluster
const cluster = new pgedge.Cluster("exampleCluster", {
  name: "example",
  cloudAccountId: cloudAccount.id,
  regions: ["us-west-2", "us-east-1", "eu-central-1"],
  nodeLocation: "public",
  sshKeyId: sshKey.id,
  nodes: [
    {
      name: "n1",
      region: "us-west-2",
      instanceType: "r6g.medium",
      volumeSize: 100,
      volumeType: "gp2",
    },
    {
      name: "n2",
      region: "us-east-1",
      instanceType: "r6g.medium",
      volumeSize: 100,
      volumeType: "gp2",
    },
    {
      name: "n3",
      region: "eu-central-1",
      instanceType: "r6g.medium",
      volumeSize: 100,
      volumeType: "gp2",
    },
  ],
  networks: [
    {
      region: "us-west-2",
      cidr: "10.1.0.0/16",
      publicSubnets: ["10.1.0.0/24"],
      // privateSubnets: ["10.1.0.0/24"],
    },
    {
      region: "us-east-1",
      cidr: "10.2.0.0/16",
      publicSubnets: ["10.2.0.0/24"],
      // privateSubnets: ["10.2.0.0/24"],
    },
    {
      region: "eu-central-1",
      cidr: "10.3.0.0/16",
      publicSubnets: ["10.3.0.0/24"],
      // privateSubnets: ["10.3.0.0/24"],
    },
  ],
  backupStoreIds: [backupStore.id],
  firewallRules: [
    {
      name: "postgres",
      port: 5432,
      sources: ["123.456.789.0/32"],
    },
  ],
}, { dependsOn: backupStore });

// Create a Database
const database = new pgedge.Database("exampleDatabase", {
  name: "example",
  clusterId: cluster.id,
  options: [
    "autoddl:enabled",
    // "install:northwind",
    // "rest:enabled",
    // "cloudwatch_metrics:enabled",
  ],
  extensions: {
    autoManage: true,
    requesteds: [
      "postgis",
      "vector",
    ],
  },
  nodes:{
    n1: {
      name: "n1",
    },
    n2: {
      name: "n2",
    },
    n3: {
      name: "n3",
    },
  },
  backups: {
    provider: "pgbackrest",
    configs: [
      {
        id: "default",
        schedules: [
          {
            id: "daily-full-backup",
            cronExpression: "0 1 * * *",
            type: "full",
          },
          {
            id: "hourly-incr-backup",
            cronExpression: "0 * * * ?",
            type: "incr",
          },
        ]
      },
    ]
  },
}, { dependsOn: cluster });

// Export the resource IDs
export const sshKeyId = sshKey.id;
export const cloudAccountId = cloudAccount.id;
export const backupStoreId = backupStore.id;
export const clusterId = cluster.id;
export const databaseId = database.id;
```

{{% /choosable %}}
{{< /chooser >}}

> You could find more complete and detailed examples in the [pulumi-pgedge repository](https://github.com/pgEdge/pulumi-pgedge/tree/main/examples)
