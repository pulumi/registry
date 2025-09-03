---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-auth0/v3.26.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Auth0 Provider
meta_desc: Provides an overview on how to configure the Pulumi Auth0 provider.
layout: package
---

## Installation

The Auth0 provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/auth0`](https://www.npmjs.com/package/@pulumi/auth0)
* Python: [`pulumi-auth0`](https://pypi.org/project/pulumi-auth0/)
* Go: [`github.com/pulumi/pulumi-auth0/sdk/v3/go/auth0`](https://github.com/pulumi/pulumi-auth0)
* .NET: [`Pulumi.Auth0`](https://www.nuget.org/packages/Pulumi.Auth0)
* Java: [`com.pulumi/auth0`](https://central.sonatype.com/artifact/com.pulumi/auth0)

## Overview

The Auth0 provider is used to interact with the [Auth0 Management API](https://auth0.com/docs/api/management/v2) in
order to configure an Auth0 Tenant.

It provides resources that allow you to create and manage clients, resource servers, client grants, connections, email
providers and templates, rules and rule variables, users, roles, tenants, custom domains, and many more, as part of a
Pulumi deployment.

Use the navigation to the left to read about the available resources and functions.
## Example Usage

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    auth0:clientId:
        value: <client-id>
    auth0:clientSecret:
        value: <client-secret>
    auth0:debug:
        value: <debug>
    auth0:domain:
        value: <domain>

```

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    auth0:clientAssertionPrivateKey:
        value: 'TODO: file("<path-to-private-key>")'
    auth0:clientAssertionSigningAlg:
        value: <signing-algorithm>
    auth0:clientId:
        value: <client-id>
    auth0:debug:
        value: <debug>
    auth0:domain:
        value: <domain>

```

> Hard-coding credentials into any Pulumi configuration is not recommended, and risks secret leakage should this
file ever be committed to a public version control system. See Environment Variables for a
better alternative.
## Configuration Reference

- `apiToken` (String) Your Auth0 [management api access token](https://auth0.com/docs/security/tokens/access-tokens/management-api-access-tokens). It can also be sourced from the `AUTH0_API_TOKEN` environment variable. It can be used instead of `clientId` + `clientSecret`. If both are specified, `apiToken` will be used over `clientId` + `clientSecret` fields.
- `audience` (String) Your Auth0 audience when using a custom domain. It can also be sourced from the `AUTH0_AUDIENCE` environment variable.
- `cliLogin` (Boolean) While toggled on, the API token gets fetched from the keyring for the given domain
- `clientAssertionPrivateKey` (String) The private key used to sign the client assertion JWT. It can also be sourced from the `AUTH0_CLIENT_ASSERTION_PRIVATE_KEY` environment variable.
- `clientAssertionSigningAlg` (String) The algorithm used to sign the client assertion JWT. It can also be sourced from the `AUTH0_CLIENT_ASSERTION_SIGNING_ALG` environment variable.
- `clientId` (String) Your Auth0 client ID. It can also be sourced from the `AUTH0_CLIENT_ID` environment variable.
- `clientSecret` (String) Your Auth0 client secret. It can also be sourced from the `AUTH0_CLIENT_SECRET` environment variable.
- `customDomainHeader` (String) When specified, this header is added to requests targeting a set of pre-defined whitelisted URLs Global setting overrides all resource specific `customDomainHeader` value
- `debug` (Boolean) Enables HTTP request and response logging when TF_LOG=DEBUG is set. It can also be sourced from the `AUTH0_DEBUG` environment variable.
- `domain` (String) Your Auth0 domain name. It can also be sourced from the `AUTH0_DOMAIN` environment variable.
- `dynamicCredentials` (Boolean) Indicates whether credentials will be dynamically passed to the provider from other pulumi resources.
## Environment Variables

You can provide your credentials via the `AUTH0_DOMAIN`, `AUTH0_CLIENT_ID` and `AUTH0_CLIENT_SECRET` or `AUTH0_API_TOKEN`
or `AUTH0_DOMAIN`, `AUTH0_CLIENT_ID`, `AUTH0_CLIENT_ASSERTION_PRIVATE_KEY` and `AUTH0_CLIENT_ASSERTION_SIGNING_ALG` environment variables, respectively.

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:

```
### Example Usage with Client Secret

```shell
AUTH0_DOMAIN="<domain>" \
AUTH0_CLIENT_ID="<client-id>" \
AUTH0_CLIENT_SECRET="<client_secret>" \
pulumi preview
```
### Example Usage with Private JWT

```shell
AUTH0_DOMAIN="<domain>" \
AUTH0_CLIENT_ID="<client-id>" \
AUTH0_CLIENT_ASSERTION_PRIVATE_KEY="<private-key>" \
AUTH0_CLIENT_ASSERTION_SIGNING_ALG="<signing-algorithm>" \
pulumi preview
```
## Importing resources

To import Auth0 resources, you will need to know their ID. You can use
the [Auth0 API Explorer](https://auth0.com/docs/api/management/v2) to find your resource ID.