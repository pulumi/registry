---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi-labs/pulumi-nvidia-aicr/v0.1.6/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi-labs/pulumi-nvidia-aicr/blob/v0.1.6/docs/installation-configuration.md
title: NVIDIA AI Cluster Runtime (AICR) Installation & Configuration
meta_desc: How to install the Pulumi NVIDIA AICR provider and its language SDKs.
layout: package
---

## Installation

The NVIDIA AICR provider is available as a package in all Pulumi-supported languages:

* JavaScript/TypeScript: [`@pulumi-labs/nvidia-aicr`](https://www.npmjs.com/package/@pulumi-labs/nvidia-aicr)
* Python: [`pulumi-labs-nvidia-aicr`](https://pypi.org/project/pulumi-labs-nvidia-aicr/)
* Go: [`github.com/pulumi-labs/pulumi-nvidia-aicr/sdk/go/nvidiaaicr`](https://github.com/pulumi-labs/pulumi-nvidia-aicr)
* .NET: [`Pulumi.Labs.NvidiaAicr`](https://www.nuget.org/packages/Pulumi.Labs.NvidiaAicr)
* Java: [`com.pulumi.labs:nvidia-aicr`](https://central.sonatype.com/artifact/com.pulumi.labs/nvidia-aicr)

In normal usage, installing one of the language SDKs above is enough. The first `pulumi up` automatically downloads the matching provider binary from the URL embedded in the SDK, so no separate step is required.

### Provider Binary (optional)

If you need to install the provider binary explicitly (for example, in an air-gapped environment or to prewarm a CI cache), use the `pulumi plugin` command:

```bash
pulumi plugin install resource nvidia-aicr <version> --server github://api.github.com/pulumi-labs/pulumi-nvidia-aicr
```

Replace `<version>` with your desired release (see the [releases page](https://github.com/pulumi-labs/pulumi-nvidia-aicr/releases)).

## Configuration

The NVIDIA AICR provider is a Pulumi *component provider* and does not require any provider-level configuration. The deployed software stack is configured per-resource via the inputs to `ClusterStack` (`accelerator`, `service`, `intent`, `platform`, `kubeconfig`, and optional `componentOverrides`). See the API Reference (left navigation) for the full input surface.

`ClusterStack` requires only a Kubernetes `kubeconfig`, which can be:

* A literal kubeconfig string passed via the `kubeconfig` input.
* A path to a kubeconfig file on disk via the `kubeconfigPath` input.
* An `Output` from another Pulumi resource (e.g. `EksCluster.kubeconfig` from `pulumi-eks`).

For runnable per-cloud examples (EKS, AKS, GKE, OKE, CoreWeave, kind), see the [`examples/`](https://github.com/pulumi-labs/pulumi-nvidia-aicr/tree/main/examples) directory in the provider repository.
