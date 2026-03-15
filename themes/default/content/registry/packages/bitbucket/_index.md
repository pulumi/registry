---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/drfaust92/bitbucket/2.51.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Bitbucket Provider
meta_desc: Provides an overview on how to configure the Pulumi Bitbucket provider.
layout: package
---

## Generate Provider

The Bitbucket provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider drfaust92/bitbucket
```
## Overview

The Bitbucket provider allows you to manage resources including repositories,
webhooks, and default reviewers.

Use the navigation to the left to read about the available resources.
## Configuration Reference

The following configuration inputs are supported in the provider configuration:

* `username` - (Optional) Username to use for authentication via [Basic
  Auth](https://developer.atlassian.com/cloud/bitbucket/rest/intro/#basic-auth).
  When using an API token, this must be your **Atlassian account email address**.
  You can also set this via the `BITBUCKET_USERNAME` environment variable.
  If configured, requires `password` to be configured as well.

* `password` - (Optional) Password to use for authentication via [Basic
  Auth](https://developer.atlassian.com/cloud/bitbucket/rest/intro/#basic-auth).
  It is recommended to use an [API Token](https://support.atlassian.com/atlassian-account/docs/manage-api-tokens-for-your-atlassian-account/)
  created at [Atlassian Account Settings](https://id.atlassian.com/manage-profile/security/api-tokens)
  as your password, with your Atlassian account email as the username.
  **Note:** Bitbucket App Passwords are deprecated and will stop working on June 9, 2026.
  You can also set this via the `BITBUCKET_PASSWORD` environment variable. If
  configured, requires `username` to be configured as well.

* `oauthClientId` - (Optional) OAuth client ID to use for authentication via
  [Client Credentials
  Grant](https://developer.atlassian.com/cloud/bitbucket/rest/intro/#3--client-credentials-grant--4-4-).
  You can also set this via the `BITBUCKET_OAUTH_CLIENT_ID` environment
  variable. If configured, requires `oauthClientSecret` to be configured as
  well.

* `oauthClientSecret` - (Optional) OAuth client secret to use for authentication via
  [Client Credentials
  Grant](https://developer.atlassian.com/cloud/bitbucket/rest/intro/#3--client-credentials-grant--4-4-).
  You can also set this via the `BITBUCKET_OAUTH_CLIENT_SECRET` environment
  variable. If configured, requires `oauthClientId` to be configured as well.

* `oauthToken` - (Optional) An OAuth access token used for authentication via
  [OAuth](https://developer.atlassian.com/cloud/bitbucket/rest/intro/#oauth-2-0).
  You can also set this via the `BITBUCKET_OAUTH_TOKEN` environment variable.
## OAuth2 Scopes

To interact with the Bitbucket API, an [API Token](https://support.atlassian.com/atlassian-account/docs/manage-api-tokens-for-your-atlassian-account/)
or [OAuth Client Credentials](https://support.atlassian.com/bitbucket-cloud/docs/use-oauth-on-bitbucket-cloud/)
are required.

API tokens and OAuth client credentials are limited in scope, each API
requires certain scope to interact with, each resource doc will specify what
scopes are required to use that resource.

See the [Bitbucket OAuth
Documentation](https://support.atlassian.com/bitbucket-cloud/docs/use-oauth-on-bitbucket-cloud/)
for more information on scopes.