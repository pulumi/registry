---
# WARNING: this file was fetched from https://raw.githubusercontent.com/EventStore/pulumi-eventstorecloud/v0.2.20/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Event Store Cloud
meta_desc: Learn how you can use Event Store Cloud Provider for Pulumi to provision and manage Event Store Cloud resources.
layout: package
---

The Event Store Cloud provider for Pulumi can provision many of the cloud resources available in [Event Store Cloud](https://eventstore.com/cloud/). It uses the Event Store Cloud API to manage and provision resources.

The Event Store Cloud provider must be configured with credentials to deploy and update resources; see [Installation & Configuration](./installation-configuration) for instructions.

## Example

```typescript
import * as eventstore from "@eventstore/pulumi-eventstorecloud";

const project = new eventstore.Project("sample-project", {
    name: "Improved Chicken Window",
});

const network = new eventstore.Network("sample-network", {
    name: "Chicken Window Net",
    projectId: project.id,
    resourceProvider: "aws",
    region: "eu-west1",
    cidrBlock: "172.21.0.0/16",
});
```
