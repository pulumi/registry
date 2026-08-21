index_md_content = """---
title: Multipass Provider
meta_desc: Provides an overview of the Pulumi Multipass provider, including installation and usage examples.
layout: overview
---

The Pulumi Multipass provider enables declarative management of [Canonical Multipass](https://multipass.run) virtual machines using Pulumi.

Multipass is a tool to launch and manage Ubuntu virtual machines on macOS, Linux, and Windows. This provider turns Multipass VMs into first-class Pulumi resources, allowing you to describe multi-node local topologies in code, track changes, take snapshots, and destroy environments without relying on cloud infrastructure.

## Installation

The Multipass provider is distributed as GitHub releases. Install the resource plugin using the Pulumi CLI:

```bash
pulumi plugin install resource multipass v0.1.0 \\
  --server github://[api.github.com/incsteps/pulumi-provider-multipass](https://api.github.com/incsteps/pulumi-provider-multipass)
```

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

