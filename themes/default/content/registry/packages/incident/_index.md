---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/incident-io/incident/5.26.1/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Incident Provider
meta_desc: Provides an overview on how to configure the Pulumi Incident provider.
layout: package
---

## Generate Provider

The Incident provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider incident-io/incident
```
## Overview

This project is the official Pulumi provider for incident.io.

With this provider you manage configuration such as incident severities, roles,
custom fields and more inside of your incident.io account.

To view the full documentation of this provider, we recommend reading the
documentation on the Pulumi
Registry.
## Example Usage

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    incident:apiKey:
        value: <api-key>

```
## Configuration Reference

- `apiKey` (String, Sensitive) API key for incident.io (<https://app.incident.io/settings/api-keys)>. Sourced from the `INCIDENT_API_KEY` environment variable, if set.
- `endpoint` (String) URL of the incident.io API