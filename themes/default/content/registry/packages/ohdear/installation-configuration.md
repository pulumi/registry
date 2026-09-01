---
# WARNING: this file was fetched from https://raw.githubusercontent.com/matthiashamacher/pulumi-ohdear/v1.0.0/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/matthiashamacher/pulumi-ohdear/blob/v1.0.0/docs/installation-configuration.md
title: Oh Dear Installation & Configuration
meta_desc: How to install the Pulumi Oh Dear provider and authenticate it with an Oh Dear API token.
layout: installation
---

## Installation

The Oh Dear provider is available for Node.js/TypeScript and Python.

```bash
npm install @matthiashamacher/ohdear
```

```bash
pip install matthiashamacher_ohdear
```

The provider is a native (Go) plugin. Pulumi downloads it automatically from the
package's `pluginDownloadURL`; to install it by hand:

```bash
pulumi plugin install resource ohdear <version> \
  --server github://api.github.com/matthiashamacher/pulumi-ohdear
```

## Authentication

Every Oh Dear API call uses one token, sent as `Authorization: Bearer <token>`.
Create a token under **Account → API tokens** in Oh Dear
(<https://ohdear.app/user/api-tokens>).

Provide it in one of two ways:

### Environment variable

```bash
export OHDEAR_API_TOKEN="your-api-token"
```

### Pulumi configuration

```bash
pulumi config set --secret ohdear:apiToken "your-api-token"
```

Explicit config wins over the environment variable. If neither is set, the
provider fails at startup with a clear error.

## Configuration reference

| Key | Environment variable | Required | Description |
| --- | --- | --- | --- |
| `ohdear:apiToken` | `OHDEAR_API_TOKEN` | yes | Oh Dear API token. Stored as a secret; sent as a Bearer token on every request. |

## Finding IDs

Most resources need a numeric `teamId`. List your teams and their IDs with:

```bash
curl https://ohdear.app/api/me \
  -H "Authorization: Bearer $OHDEAR_API_TOKEN" \
  -H 'Accept: application/json'
```
