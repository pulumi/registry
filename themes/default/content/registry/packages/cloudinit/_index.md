---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-cloudinit/v1.4.12/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Cloudinit Provider
meta_desc: Provides an overview on how to configure the Pulumi Cloudinit provider.
layout: package
---
## Installation

The cloudinit provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/cloudinit`](https://www.npmjs.com/package/@pulumi/cloudinit)
* Python: [`pulumi-cloudinit`](https://pypi.org/project/pulumi-cloudinit/)
* Go: [`github.com/pulumi/pulumi-cloudinit/sdk/go/cloudinit`](https://github.com/pulumi/pulumi-cloudinit)
* .NET: [`Pulumi.Cloudinit`](https://www.nuget.org/packages/Pulumi.Cloudinit)
* Java: [`com.pulumi/cloudinit`](https://central.sonatype.com/artifact/com.pulumi/cloudinit)
## Overview

The cloud-init Pulumi provider exposes the `cloudinit.Config` function which renders a [multipart MIME configuration](https://cloudinit.readthedocs.io/en/latest/explanation/format.html#mime-multi-part-archive) for use with [cloud-init](https://cloudinit.readthedocs.io/en/latest/).

This provider requires no configuration. For information on the resources it provides, see the navigation bar.