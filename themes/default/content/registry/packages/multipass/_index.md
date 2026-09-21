---
# WARNING: this file was fetched from https://raw.githubusercontent.com/incsteps/pulumi-provider-multipass/v0.3.2/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/incsteps/pulumi-provider-multipass/blob/v0.3.2/docs/_index.md
title: Multipass Provider
meta_desc: Provides an overview of the Pulumi Multipass provider, including installation and usage examples.
layout: overview
---

The Pulumi Multipass provider enables declarative management of [Canonical Multipass](https://multipass.run) virtual machines using Pulumi.

Multipass is a tool to launch and manage Ubuntu virtual machines on macOS, Linux, and Windows. 
This provider turns Multipass VMs into first-class Pulumi resources, allowing you to describe multi-node local topologies in code, track changes, and destroy environments without relying on cloud infrastructure.

## Installation

The Multipass provider is distributed as GitHub releases. Install the resource plugin using the Pulumi CLI:

```bash
pulumi plugin install resource multipass v0.3.1 --server github://api.github.com/incsteps/pulumi-provider-multipass
```

    Once the provider will be included as part of the Pulumi Community Registry, this step will be unnecessary.

Then add the SDK to your project:

```bash
pulumi package add multipass
```


### Usage

```typescript
import * as multipass from "@incsteps/pulumi-multipass";

const vm = new multipass.resources.Instance("dev", {
    name:   "dev",
    image:  "24.04",
    cpus:   2,
    memory: "4G",
    disk:   "20G",
});

export const ip = vm.ipv4;
```

