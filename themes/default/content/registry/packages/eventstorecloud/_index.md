---
title: Event Store Cloud
meta_desc: Learn how you can use Event Store Cloud Provider for Pulumi to provision and manage Event Store Cloud resources.
layout: package
---

The Event Store Cloud (ESC) provider for Pulumi can provision many of the cloud resources available in [ESC](https://eventstore.com/cloud/). It uses the ESC API to manage and provision resources.

The ESC provider must be configured with credentials to deploy and update resources; see [Installation & Configuration]({{<relref "./installation-configuration">}}) for instructions.

## Example

```typescript
import * as esc from "@eventstore/pulumi-eventstorecloud";

const project = new esc.Project("sample-project", {
    name: "Improved Chicken Window"
});

const network = new esc.Network("sample-network", {
    name: "Chicken Window Net",
    projectId: project.id,
    resourceProvider: "aws",
    region: "eu-west1",
    cidrBlock: "172.21.0.0/16"
});
```
