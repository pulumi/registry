---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/incident-io/incident/6.11.0/index.md
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
## Supported Pulumi versions

From v6.0.0 this provider supports Pulumi 1.14 and above, and is tested
against the Pulumi versions HashiCorp still patch. It is also tested directly
against OpenTofu. Older Pulumi releases are no longer tested, and while the
provider may continue to work with them we won't be fixing issues that only
reproduce there. Pin to v5.x if you need to stay on an older CLI.
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
- `markImportedResourcesAsManaged` (Boolean) Whether importing a resource claims it as managed by Pulumi, which is what stops people editing it in the incident.io dashboard. Defaults to `true`. Pulumi runs imports during `plan` rather than apply, so this claim is a write to your account during an operation you may expect to be read-only: set this to `false` if plans must leave your account untouched. Creating or updating a resource claims it regardless of this setting, so a resource imported with this off is claimed by the first apply that changes it. It stays editable in the dashboard until then, and indefinitely if its configuration already matches the account and so never produces a change to apply.