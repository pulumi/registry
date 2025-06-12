---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-okta/v4.20.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Okta Provider
meta_desc: Provides an overview on how to configure the Pulumi Okta provider.
layout: package
---
## Installation

The Okta provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/okta`](https://www.npmjs.com/package/@pulumi/okta)
* Python: [`pulumi-okta`](https://pypi.org/project/pulumi-okta/)
* Go: [`github.com/pulumi/pulumi-okta/sdk/v4/go/okta`](https://github.com/pulumi/pulumi-okta)
* .NET: [`Pulumi.Okta`](https://www.nuget.org/packages/Pulumi.Okta)
* Java: [`com.pulumi/okta`](https://central.sonatype.com/artifact/com.pulumi/okta)
## Overview

The Okta provider is used to interact with the resources supported by Okta. The provider needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources and functions.

In case the provider configuration is still using old `"oktadeveloper/okta"` source, please change it to `"okta/okta"`. Okta no longer supports `"oktadeveloper/okta"`.
## Example Usage



{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    okta:baseUrl:
        value: '[okta.com|oktapreview.com]'
    okta:clientId:
        value: '[APP CLIENT_ID]'
    okta:orgName:
        value: '[ORG NAME e.g. dev-123456]'
    okta:privateKey:
        value: '[PRIVATE KEY]'
    okta:privateKeyId:
        value: '[PRIVATE KEY ID - KID]'
    okta:scopes:
        value: '[COMMA,SEPARATED,SCOPE,VALUES]'

```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    okta:baseUrl:
        value: '[okta.com|oktapreview.com]'
    okta:clientId:
        value: '[APP CLIENT_ID]'
    okta:orgName:
        value: '[ORG NAME e.g. dev-123456]'
    okta:privateKey:
        value: '[PRIVATE KEY]'
    okta:privateKeyId:
        value: '[PRIVATE KEY ID - KID]'
    okta:scopes:
        value: '[COMMA,SEPARATED,SCOPE,VALUES]'

```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    okta:baseUrl:
        value: '[okta.com|oktapreview.com]'
    okta:clientId:
        value: '[APP CLIENT_ID]'
    okta:orgName:
        value: '[ORG NAME e.g. dev-123456]'
    okta:privateKey:
        value: '[PRIVATE KEY]'
    okta:privateKeyId:
        value: '[PRIVATE KEY ID - KID]'
    okta:scopes:
        value: '[COMMA,SEPARATED,SCOPE,VALUES]'

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    okta:baseUrl:
        value: '[okta.com|oktapreview.com]'
    okta:clientId:
        value: '[APP CLIENT_ID]'
    okta:orgName:
        value: '[ORG NAME e.g. dev-123456]'
    okta:privateKey:
        value: '[PRIVATE KEY]'
    okta:privateKeyId:
        value: '[PRIVATE KEY ID - KID]'
    okta:scopes:
        value: '[COMMA,SEPARATED,SCOPE,VALUES]'

```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    okta:baseUrl:
        value: '[okta.com|oktapreview.com]'
    okta:clientId:
        value: '[APP CLIENT_ID]'
    okta:orgName:
        value: '[ORG NAME e.g. dev-123456]'
    okta:privateKey:
        value: '[PRIVATE KEY]'
    okta:privateKeyId:
        value: '[PRIVATE KEY ID - KID]'
    okta:scopes:
        value: '[COMMA,SEPARATED,SCOPE,VALUES]'

```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    okta:baseUrl:
        value: '[okta.com|oktapreview.com]'
    okta:clientId:
        value: '[APP CLIENT_ID]'
    okta:orgName:
        value: '[ORG NAME e.g. dev-123456]'
    okta:privateKey:
        value: '[PRIVATE KEY]'
    okta:privateKeyId:
        value: '[PRIVATE KEY ID - KID]'
    okta:scopes:
        value: '[COMMA,SEPARATED,SCOPE,VALUES]'

```

{{% /choosable %}}
{{< /chooser >}}

For the resources and functions examples, please check the examples directory.
## Authentication

The Okta provider offers a flexible means of providing credentials for
authentication. The following methods are supported, in this order, and
explained below:

- Environment variables
- Provider Config
### Environment variables

You can provide your credentials via the `OKTA_ORG_NAME`, `OKTA_BASE_URL`,
`OKTA_ACCESS_TOKEN`, `OKTA_API_TOKEN`, `OKTA_API_CLIENT_ID`, `OKTA_API_SCOPES`,
`OKTA_API_PRIVATE_KEY_ID`, and `OKTA_API_PRIVATE_KEY` environment variables,
representing your Okta Organization Name, Okta Base URL (i.e. `"okta.com"` or
`"oktapreview.com"`), Okta Access Token, Okta API Token, Okta Client ID, Okta
API scopes and Okta API private key respectively.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs

```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python

```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go

```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml

```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java

```

{{% /choosable %}}
{{< /chooser >}}

Usage:

```sh
# Change place holder values denoted by brackets to real values, including the
# brackets.

$ export OKTA_ORG_NAME="[ORG NAME e.g. dev-123456]"
$ export OKTA_BASE_URL="[okta.com|oktapreview.com]"
$ export OKTA_API_CLIENT_ID="[APP CLIENT_ID]"
$ export OKTA_API_PRIVATE_KEY_ID="[PRIVATE KEY ID - KID]"
$ export OKTA_API_PRIVATE_KEY="[PRIVATE KEY]"
$ export OKTA_API_SCOPES="[COMMA,SEPARATED,SCOPE,VALUES]"

$ pulumi preview
```
## Configuration Reference

Note: `apiToken` is mutually exclusive of the set `accessToken`, `clientId`,
`privateKeyId`, `privateKey`, and `scopes`. `clientId`, `privateKeyId`,
`privateKey`, and `scopes` are for [OAuth 2.0
client](https://developer.okta.com/docs/reference/api/apps/#add-oauth-2-0-client-application)
authentication for application operations. `accessToken` is used in situations
where the caller has already performed the OAuth 2.0 client authentication
process. Okta recommend using OAuth2 for authorizing your Pulumi modules. `apiToken` is utilized for
Okta's [SSWS Authorization
Scheme](https://developer.okta.com/docs/reference/core-okta-api/#authentication)
and applies to org level operations. This is a legacy authorization scheme.

In addition to generic `provider`
arguments (e.g.
`alias` and `version`), the following arguments are supported in the Okta
provider configuration:

- `orgName` - (Optional) This is the org name of your Okta account, for example `dev-123456.oktapreview.com` would have an org name of `dev-123456`. It must be provided, but it can also be sourced from the `OKTA_ORG_NAME` environment variable.

- `baseUrl` - (Optional) This is the domain of your Okta account, for example `dev-123456.oktapreview.com` would have a base url of `oktapreview.com`. It must be provided, but it can also be sourced from the `OKTA_BASE_URL` environment variable.

- `httpProxy` - (Optional) This is a custom URL endpoint that can be used for unit testing or local caching proxies. Can also be sourced from the `OKTA_HTTP_PROXY` environment variable.

- `accessToken` - (Optional) This is an OAuth 2.0 access token to interact with your Okta org. It can be sourced from the `OKTA_ACCESS_TOKEN` environment variable. `accessToken` conflicts with `apiToken`, `clientId`, `scopes` and `privateKey`.

- `apiToken` - (Optional) This is the API token to interact with your Okta org. It can also be sourced from the `OKTA_API_TOKEN` environment variable. `apiToken` conflicts with `accessToken`, `clientId`, `scopes` and `privateKey`.

- `clientId` - (Optional) This is the client ID for obtaining the API token. It can also be sourced from the `OKTA_API_CLIENT_ID` environment variable. `clientId` conflicts with `accessToken` and `apiToken`.

- `scopes` - (Optional) These are scopes for obtaining the API token in form of a comma separated list. It can also be sourced from the `OKTA_API_SCOPES` environment variable. `scopes` conflicts with `accessToken` and `apiToken`.

- `privateKey` - (Optional) This is the private key for obtaining the API token (can be represented by a filepath, or the key itself). It can also be sourced from the `OKTA_API_PRIVATE_KEY` environment variable. `privateKey` conflicts with `accessToken` and `apiToken`. The format of the PK is PKCS#1 unencrypted (header starts with `-----BEGIN RSA PRIVATE KEY-----` or PKCS#8 unencrypted (header starts with `-----BEGIN PRIVATE KEY-----`).

- `privateKeyId` - (Optional) This is the private key ID (kid) for obtaining the API token. It can also be sourced from `OKTA_API_PRIVATE_KEY_ID` environmental variable. `privateKeyId` conflicts with `apiToken`.

- `backoff` - (Optional) Whether to use exponential back off strategy for rate limits, the default is `true`.

- `minWaitSeconds` - (Optional) Minimum seconds to wait when rate limit is hit, the default is `30`.

- `maxWaitSeconds` - (Optional) Maximum seconds to wait when rate limit is hit, the default is `300`.

- `maxRetries` - (Optional) Maximum number of retries to attempt before returning an error, the default is `5`.

- `requestTimeout` - (Optional) Timeout for single request (in seconds) which is made to Okta, the default is `0` (means no limit is set). The maximum value can be `300`.

- `maxApiCapacity` - (Optional, experimental) sets what percentage of capacity the provider can use of the total
  rate limit capacity while making calls to the Okta management API endpoints. Okta API operates in one minute buckets.
  See Okta Management API Rate Limits: <https://developer.okta.com/docs/reference/rl-global-mgmt>. Can be set to a value between 1 and 100.