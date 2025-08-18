---
# WARNING: this file was fetched from https://raw.githubusercontent.com/equinix/pulumi-equinix/v0.25.0/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Equinix Installation & Configuration
meta_desc: Information on how to install the Equinix provider.
layout: package
---

## Installation

The Pulumi Equinix provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@equinix-labs/pulumi-equinix`](https://www.npmjs.com/package/@equinix-labs/pulumi-equinix)
* Python: [`pulumi_equinix`](https://pypi.org/project/pulumi-equinix/)
* Go: [`github.com/equinix/pulumi-equinix/sdk/go/equinix`](https://pkg.go.dev/github.com/equinix/pulumi-equinix/sdk)
* .NET: [`Pulumi.Equinix`](https://www.nuget.org/packages/Pulumi.Equinix)
* Java: [`com.equinix.pulumi`](https://central.sonatype.com/namespace/com.equinix.pulumi)

### Provider Binary

The Equinix provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource equinix --server github://api.github.com/equinix
```

Replace the version string with your desired version.


## Credentials 

To provision resources with the Pulumi Equinix provider, you need to have Equinix credentials. 

For information about obtaining API key and secret required for Equinix Fabric and Network Edge refer to [Generating Client ID and Client Secret key](https://developer.equinix.com/dev-docs/fabric/getting-started/getting-access-token#generating-client-id-and-client-secret) from [Equinix Developer Platform portal](https://developer.equinix.com/).

Interacting with Equinix Metal requires an API auth token that can be generated at Project-level or User-level. User API keys can be obtained by creating them in the [Equinix Metal Portal](https://console.equinix.com/) or by using the [Create a User API Key](https://deploy.equinix.com/developers/api/metal/#operation/createAPIKey) endpoint. Project API keys can also be obtained by creating them in the [Equinix Metal Portal](https://console.equinix.com/) or by using the [Create a Project API Key](https://deploy.equinix.com/developers/api/metal/#operation/createProjectAPIKey) endpoint.

If you are only using Equinix Metal resources, you may omit the Client ID and Client Secret provider configuration parameters needed to access other Equinix resource types (Network Edge, Fabric, etc).

## Configuration

There are a few different ways you can configure your Equinix credentials to work with Pulumi.

### Set environment variables

Once you have your credentials, you can set environment variables to provision resources in Equinix:

{{< chooser os "linux,macos,windows" >}}
{{% choosable os linux %}}

```bash
$ export METAL_AUTH_TOKEN=<METAL_AUTH_TOKEN>
$ export EQUINIX_API_CLIENTID=<EQUINIX_API_CLIENTID>
$ export EQUINIX_API_CLIENTSECRET=<EQUINIX_API_CLIENTSECRET>
```

{{% /choosable %}}

{{% choosable os macos %}}

```bash
$ export METAL_AUTH_TOKEN=<METAL_AUTH_TOKEN>
$ export EQUINIX_API_CLIENTID=<EQUINIX_API_CLIENTID>
$ export EQUINIX_API_CLIENTSECRET=<EQUINIX_API_CLIENTSECRET>
```

{{% /choosable %}}

{{% choosable os windows %}}

```powershell
> $env:METAL_AUTH_TOKEN = "<METAL_AUTH_TOKEN>"
> $env:EQUINIX_API_CLIENTID = "<EQUINIX_API_CLIENTID>"
> $env:EQUINIX_API_CLIENTSECRET = "<EQUINIX_API_CLIENTSECRET>"
```

{{% /choosable %}}
{{< /chooser >}}

## Configuration Options

Use `pulumi config set equinix:<option>` or pass options to the [constructor of `new equinix.Provider`](/registry/packages/equinix/api-docs/provider).

| Option | Required/Optional | Description |
|-----|------|----|
| `authToken`| Optional | The Equinix Metal API auth key for API operations. |
| `clientId`| Optional | API Consumer Key available under My Apps section in Equinix developer portal. |
| `clientSecret`| Optional | API Consumer secret available under My Apps section in Equinix developer portal. |
| `token` | Optional | API token from the developer sandbox. Token's can be generated for the API Client using the OAuth2 Token features described in the [OAuth2 API](https://developer.equinix.com/catalog/accesstokenv1#operation/GetOAuth2AccessToken) documentation. The `client_id` and `client_secret` arguments will be ignored in the presence of a token argument. |
| `endpoint` | Optional | The Equinix API base URL to point out desired environmen. Defaults to `https://api.equinix.com`. |
| `maxRetries` | Optional | The maximum number of retries in case of network failure. |
| `maxRetryWaitSeconds` | Optional | The maximum time to wait in case of network failure. |
| `requestTimeout` | Optional | The duration of time, in seconds, that the Equinix Platform API Client should wait before canceling an API request. Defaults to 30. |
| `responseMaxPageSize` | Optional | The maximum number of records in a single response for REST queries that produce paginated responses. |