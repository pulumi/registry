---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-tailscale/v0.26.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi/pulumi-tailscale/blob/v0.26.0/docs/_index.md
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Tailscale Provider
meta_desc: Provides an overview on how to configure the Pulumi Tailscale provider.
layout: package
---

## Installation

The Tailscale provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/tailscale`](https://www.npmjs.com/package/@pulumi/tailscale)
* Python: [`pulumi-tailscale`](https://pypi.org/project/pulumi-tailscale/)
* Go: [`github.com/pulumi/pulumi-tailscale/sdk/go/tailscale`](https://github.com/pulumi/pulumi-tailscale)
* .NET: [`Pulumi.Tailscale`](https://www.nuget.org/packages/Pulumi.Tailscale)
* Java: [`com.pulumi/tailscale`](https://central.sonatype.com/artifact/com.pulumi/tailscale)

## Overview

This provider is used to interact with resources supported by the [Tailscale API](https://tailscale.com/api).

Use the navigation to the left to read about the available resources and functions.
## Authentication

There are several ways to authenticate the Tailscale provider with the Tailscale API.

Using a [trust credential](https://tailscale.com/kb/1623/trust-credentials) (an OAuth client or federated identity) is
recommended whenever possible as trust credentials can have granular access scopes applied to them whereas API keys cannot.

Available authentication methods are detailed below.
### OAuth clients

[OAuth clients](https://tailscale.com/kb/1215/oauth-clients) can be used for authentication by setting the `oauthClientId`
and `oauthClientSecret` arguments in the provider configuration to the client ID and client secret of a configured OAuth client
respectively:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    tailscale:oauthClientId:
        value: my_client_id
    tailscale:oauthClientSecret:
        value: my_client_secret
    tailscale:tailnet:
        value: example.com

```

See argument reference for more details.
### Federated identities

[Workload identity federation](https://tailscale.com/kb/1581/workload-identity-federation) can be used for authentication
by setting the `oauthClientId` and `identityToken` in the provider configuration to the client ID of a configured
federated identity and a JWT identity token from a compatible issuer respectively:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    tailscale:identityToken:
        value: my_identity_token
    tailscale:oauthClientId:
        value: my_client_id
    tailscale:tailnet:
        value: example.com

```

See argument reference for more details.
### API keys

[API keys](https://tailscale.com/kb/1101/api#authentication) can be used for authentication by setting the `apiKey`
argument in the provider configuration:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    tailscale:apiKey:
        value: my_api_key
    tailscale:tailnet:
        value: example.com

```

See argument reference for more details.
## Configuration Reference

- `apiKey` (String, Sensitive) The API key to use for authenticating requests to the API. Can be set via the TAILSCALE_API_KEY environment variable. Conflicts with 'oauth_client_id' and 'oauth_client_secret'.
- `baseUrl` (String) The base URL of the Tailscale API. Defaults to <https://api.tailscale.com>. Can be set via the TAILSCALE_BASE_URL environment variable.
- `identityToken` (String, Sensitive) The jwt identity token to exchange for a Tailscale API token when using a federated identity. Can be set via the TAILSCALE_IDENTITY_TOKEN environment variable. Conflicts with 'api_key' and 'oauth_client_secret'.
- `oauthClientId` (String) The OAuth application or federated identity's ID when using OAuth client credentials or workload identity federation. Can be set via the TAILSCALE_OAUTH_CLIENT_ID environment variable. Either 'oauth_client_secret' or 'identity_token' must be set alongside 'oauth_client_id'. Conflicts with 'api_key'.
- `oauthClientSecret` (String, Sensitive) The OAuth application's secret when using OAuth client credentials. Can be set via the TAILSCALE_OAUTH_CLIENT_SECRET environment variable. Conflicts with 'api_key' and 'identity_token'.
- `scopes` (List of String) The OAuth 2.0 scopes to request when generating the access token using the supplied OAuth client credentials. See <https://tailscale.com/kb/1623/trust-credentials#scopes> for available scopes. Only valid when both 'oauth_client_id' and 'oauth_client_secret', or both are set.
- `tailnet` (String) The tailnet ID. Tailnets created before Oct 2025 can still use the legacy ID, but the Tailnet ID is the preferred identifier. Can be set via the TAILSCALE_TAILNET environment variable. Default is the tailnet that owns API credentials passed to the provider.
- `userAgent` (String) User-Agent header for API requests.