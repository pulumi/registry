---
# WARNING: this file was fetched from https://raw.githubusercontent.com/incident-io/pulumi-incident/v0.1.2/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/incident-io/pulumi-incident/blob/v0.1.2/docs/installation-configuration.md
title: incident.io Installation & Configuration
meta_desc: How to install the incident.io Pulumi provider and authenticate it against your account.
layout: package
---

## Installation

The incident.io provider is available in three languages:

- JavaScript/TypeScript: [`@incident-io/pulumi`](https://www.npmjs.com/package/@incident-io/pulumi)
- Python: [`pulumi-incident`](https://pypi.org/project/pulumi-incident/)
- Go: [`github.com/incident-io/pulumi-incident/sdk/go/incident`](https://github.com/incident-io/pulumi-incident)

```bash
npm install @incident-io/pulumi
pip install pulumi-incident
go get github.com/incident-io/pulumi-incident/sdk/go/incident
```

## Authentication

The provider needs an incident.io API key. Create one in [Settings → API keys](https://app.incident.io/settings/api-keys), giving it the scopes for the resources you intend to manage.

### Stack configuration

Usually the better option, because the key is encrypted in your Pulumi state alongside the program that uses it, and different stacks can point at different incident.io accounts:

```bash
pulumi config set --secret incident:apiKey inc_...
```

### Environment variable

```bash
export INCIDENT_API_KEY=inc_...
```

Set this when you would rather the key came from your CI system's secret store than from stack config.

## Configuration reference

### `apiKey`

Your incident.io API key. Sourced from `INCIDENT_API_KEY` if not set in stack config. Sensitive.

### `endpoint`

The incident.io API URL. Defaults to `https://api.incident.io`, which is right for everyone on our standard deployment. Set it only if you are on a dedicated or self-hosted deployment. Reads from `INCIDENT_ENDPOINT`.

Note that `INCIDENT_ENDPOINT` takes precedence over the stack config value, so if both are set the environment variable wins.

### `markImportedResourcesAsManaged`

Whether importing a resource claims it as managed by code, which is what stops people editing it in the incident.io dashboard. Defaults to `true`.

The claim is a write to your incident.io account, so set this to `false` if you want imports to leave your account untouched.

Creating or updating a resource claims it regardless of this setting, so a resource imported with this off is claimed by the first update that changes it. Until then it stays editable in the dashboard, and indefinitely if its configuration already matches your account and so never produces a change to apply.

## Example

```yaml
# Pulumi.yaml
name: incident-config
runtime: nodejs

config:
  incident:apiKey:
    secret: true
```
