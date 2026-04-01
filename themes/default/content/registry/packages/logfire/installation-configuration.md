---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pydantic/pulumi-logfire/v0.1.7/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pydantic/pulumi-logfire/blob/v0.1.7/docs/installation-configuration.md
title: Installation & Configuration
meta_desc: Install and configure the Pulumi Logfire provider.
layout: package
---

# Installation

## Node.js (TypeScript/JavaScript)

Package page: [npm](https://www.npmjs.com/package/@pydantic/pulumi-logfire)

```bash
npm install @pydantic/pulumi-logfire
```

## Python

Package page: [PyPI](https://pypi.org/project/pydantic-pulumi-logfire/)

```bash
pip install pydantic-pulumi-logfire
```

Import it in code as `pulumi_logfire`.

## Go

Package page: [pkg.go.dev](https://pkg.go.dev/github.com/pydantic/pulumi-logfire/sdk/go/logfire)

```bash
go get github.com/pydantic/pulumi-logfire/sdk/go/logfire
```

## Provider Configuration

Set provider config with Pulumi config values (recommended):

```bash
pulumi config set --secret logfire:apiKey pylf_v2_us_...
# Self-hosted only:
# pulumi config set logfire:baseUrl https://<self-hosted-logfire>
```

You can also use environment variables:

- `LOGFIRE_API_KEY`
- `LOGFIRE_BASE_URL` (optional override; self-hosted customers should set this)

For Logfire SaaS, the provider infers `https://logfire-us.pydantic.dev` or `https://logfire-eu.pydantic.dev` from the API key region. If you set `logfire:baseUrl` or `LOGFIRE_BASE_URL`, that value is used instead.

## Configuration Reference

- `logfire:baseUrl` (string): Optional override for the Logfire API base URL. If omitted, the provider uses `LOGFIRE_BASE_URL` or infers the SaaS endpoint from the API key region. Self-hosted customers should set this explicitly.
- `logfire:apiKey` (secret string): Bearer token for the Logfire API.
