---
user-invocable: false
description: Setup checklist for first-party Pulumi provider packages
---

# First-Party Provider Setup Checklist

First-party Pulumi providers do **not** use the community package list (`community-packages/package-list.json`). Instead, they use `repository_dispatch` events to trigger updates in the registry.

## How It Works

1. The provider repository sends a `resource-provider` dispatch event to `pulumi/registry`
2. The `publish-provider-update.yml` workflow in `pulumi/registry` handles the event
3. The workflow updates the package YAML and publishes to the live registry

## Checklist for Provider Repository

The following must be configured in the **provider repository** (not in `pulumi/registry`):

- [ ] **Repository dispatch**: Configure CI/CD to send `repository_dispatch` events with type `resource-provider` to `pulumi/registry` on new releases
- [ ] **Event payload**: Include `provider_name` and `version` in the dispatch payload
- [ ] **GitHub token**: Ensure the CI/CD pipeline has a token with `repo` scope for `pulumi/registry` (typically via a shared bot token or GitHub App)
- [ ] **Schema file**: Ensure the schema JSON file is published and accessible at the URL specified in the package YAML

## Example Dispatch Event

```json
{
  "event_type": "resource-provider",
  "client_payload": {
    "provider_name": "example",
    "version": "v1.0.0"
  }
}
```

## What to Tell the User

Display this message after creating the package YAML:

```
First-party provider packages are updated automatically via repository_dispatch.

The provider repo ({repo_url}) needs to be configured to send
'resource-provider' dispatch events to pulumi/registry when new
versions are released. This triggers the publish-provider-update.yml
workflow which updates the package metadata and publishes to the
live registry.

Please coordinate with the provider team to set up the dispatch
event in their CI/CD pipeline.
```
