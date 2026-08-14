---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/stripe/stripe/0.3.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Stripe Provider
meta_desc: Provides an overview on how to configure the Pulumi Stripe provider.
layout: package
---

## Generate Provider

The Stripe provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider stripe/stripe
```
## Overview

Manage Stripe resources with Pulumi.
## Configuration Reference

- `apiKey` (String, Sensitive) Stripe API key. Can also be set via the STRIPE_API_KEY environment variable.
- `stripeAccount` (String) Connected account context for Connect-scoped requests. Can also be set via the STRIPE_ACCOUNT environment variable.