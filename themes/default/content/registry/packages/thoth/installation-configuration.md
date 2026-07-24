---
# WARNING: this file was fetched from https://raw.githubusercontent.com/atensecurity/pulumi-thoth/v0.1.15/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/atensecurity/pulumi-thoth/blob/v0.1.15/docs/installation-configuration.md
title: Thoth Installation & Configuration
meta_desc: Install and configure the Thoth Pulumi provider for runtime authorization of AI agents.
layout: package
---

# Installation & Configuration

## Published SDKs

- Node.js: [`@atensec/pulumi-thoth`](https://www.npmjs.com/package/@atensec/pulumi-thoth)
- Python: [`pulumi-thoth`](https://pypi.org/project/pulumi-thoth/)
- .NET: [`AtenSecurity.Pulumi.Thoth`](https://www.nuget.org/packages/AtenSecurity.Pulumi.Thoth)

This package currently supports Node.js, Python, and .NET SDKs.

## Install

### Node.js

```bash
npm install @pulumi/pulumi @atensec/pulumi-thoth
```

### Python

```bash
pip install pulumi pulumi-thoth
```

### .NET

```bash
dotnet add package AtenSecurity.Pulumi.Thoth
```

## Install provider plugin

Pulumi installs plugins automatically during `pulumi up`, but you can install
the provider plugin manually:

```bash
pulumi plugin install resource thoth 0.1.15 --server github://api.github.com/atensecurity/pulumi-thoth
```

## Configure provider credentials

Set provider configuration with `pulumi config`:

```bash
pulumi config set thoth:tenantId <tenant-id>
pulumi config set --secret thoth:orgApiKey <org-api-key>
```

Alternative auth inputs are also supported:

- `thoth:orgApiKeyFile`
- `thoth:adminBearerToken`
- `thoth:adminBearerTokenFile`

Optional endpoint controls:

- `thoth:apexDomain` (defaults to `atensecurity.com`)
- `thoth:apiBaseUrl` (explicit GovAPI URL override)

Environment variable fallbacks:

- `THOTH_TENANT_ID`
- `THOTH_API_KEY`
