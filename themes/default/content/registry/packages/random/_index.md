---
# WARNING: this file was fetched from https://d3uflda8hlgvg8.cloudfront.net/docs/registry.opentofu.org/hashicorp/random/3.6.3/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Random Provider
meta_desc: Provides an overview on how to configure the Pulumi Random provider.
layout: package
---

## Generate Provider

The Random provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider hashicorp/random
```
## Overview

The "random" provider allows the use of randomness within Pulumi
configurations. This is a *logical provider*, which means that it works
entirely within Pulumi's logic, and doesn't interact with any other
services.

Unconstrained randomness within a Pulumi configuration would not be very
useful, since Pulumi's goal is to converge on a fixed configuration by
applying a diff. Because of this, the "random" provider provides an idea of
*managed randomness*: it provides resources that generate random values during
their creation and then hold those values steady until the inputs are changed.

Even with these resources, it is advisable to keep the use of randomness within
Pulumi configuration to a minimum, and retain it for special cases only;
Pulumi works best when the configuration is well-defined, since its behavior
can then be more readily predicted.

Unless otherwise stated within the documentation of a specific resource, this
provider's results are **not** sufficiently random for cryptographic use.

For more information on the specific resources available, see the links in the
navigation bar. Read on for information on the general patterns that apply
to this provider's resources.
## Resource "Keepers"

As noted above, the random resources generate randomness only when they are
created; the results produced are stored in the Pulumi state and re-used
until the inputs change, prompting the resource to be recreated.

The resources all provide a map argument called `keepers` that can be populated
with arbitrary key/value pairs that should be selected such that they remain
the same until new random values are desired.

For example:

Resource "keepers" are optional. The other arguments to each resource must
*also* remain constant in order to retain a random result.

`keepers` are *not* treated as sensitive attributes; a value used for `keepers` will be displayed in Pulumi UI output as plaintext.

To force a random result to be replaced, the `taint` command can be used to
produce a new result on the next run.