---
# WARNING: this file was fetched from https://raw.githubusercontent.com/zscaler/pulumi-zia/v1.0.1/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: ZIA Setup
meta_desc: Information on how to install the ZIA provider.
layout: installation
---

## Installation

The Pulumi ZIA provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@bdzscaler/pulumi-zia`](https://www.npmjs.com/package/@zscaler/pulumi-zia)
* Python: [`zscaler_pulumi_zia`](https://pypi.org/project/zscaler-pulumi-zia/)
* Go: [`github.com/zscaler/pulumi-zia/sdk/go/zia`](https://pkg.go.dev/github.com/zscaler/pulumi-zia/sdk)
* .NET: [`zscaler.PulumiPackage.Zia`](https://www.nuget.org/packages/zscaler.PulumiPackage.Zia)

### Provider Binary

The ZIA provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource zia <version> --server github://api.github.com/zscaler
```

Replace the version string with your desired version.

## Setup

To provision resources with the Pulumi ZIA provider, you need to have ZIA credentials. Zscaler maintains documentation on how to create API keys [here](https://help.zscaler.com/zia/getting-started-zia-api)

## Zscaler OneAPI New Framework

The ZIA Pulumi Provider now offers support for [OneAPI](https://help.zscaler.com/oneapi/understanding-oneapi) Oauth2 authentication through [Zidentity](https://help.zscaler.com/zidentity/what-zidentity).

**NOTE** As of version v1.0.0, this Pulumi provider offers backwards compatibility to the Zscaler legacy API framework. This is the recommended authentication method for organizations whose tenants are still not migrated to [Zidentity](https://help.zscaler.com/zidentity/what-zidentity). 

**NOTE** Notice that OneAPI and Zidentity is not currently supported for the following clouds: `GOV` and `GOVUS`. Refer to the [Legacy API Framework](#legacy-api-framework) for more information on how authenticate to these environments

## Authentication - OneAPI New Framework

As of version v1.0.0, this provider supports authentication via the new Zscaler API framework [OneAPI](https://help.zscaler.com/oneapi/understanding-oneapi)

Zscaler OneAPI uses the OAuth 2.0 authorization framework to provide secure access to Zscaler Internet Access (ZIA) APIs. OAuth 2.0 allows third-party applications to obtain controlled access to protected resources using access tokens. OneAPI uses the Client Credentials OAuth flow, in which client applications can exchange their credentials with the authorization server for an access token and obtain access to the API resources, without any user authentication involved in the process.

* [ZIA API](https://help.zscaler.com/oneapi/understanding-oneapi#:~:text=managed%20using%20OneAPI.-,ZIA%20API,-Zscaler%20Internet%20Access)

### OneAPI (API Client Scope)

OneAPI Resources are automatically created within the ZIdentity Admin UI based on the RBAC Roles
applicable to APIs within the various products. For example, in ZIA, navigate to `Administration -> Role
Management` and select `Add API Role`.

Once this role has been saved, return to the ZIdentity Admin UI and from the Integration menu
select API Resources. Click the `View` icon to the right of Zscaler APIs and under the ZIA
dropdown you will see the newly created Role. In the event a newly created role is not seen in the
ZIdentity Admin UI a `Sync Now` button is provided in the API Resources menu which will initiate an
on-demand sync of newly created roles.

### Default Environment variables

You can provide credentials via the `ZSCALER_CLIENT_ID`, `ZSCALER_CLIENT_SECRET`, `ZSCALER_VANITY_DOMAIN`, `ZSCALER_CLOUD` environment variables, representing your Zidentity OneAPI credentials `clientId`, `clientSecret`, `vanityDomain` and `zscaler_cloud` respectively.

| Argument        | Description                                                                                         | Environment Variable     |
|-----------------|-----------------------------------------------------------------------------------------------------|--------------------------|
| `client_id`     | _(String)_ Zscaler API Client ID, used with `clientSecret` or `PrivateKey` OAuth auth mode.         | `ZSCALER_CLIENT_ID`      |
| `client_secret` | _(String)_ Secret key associated with the API Client ID for authentication.                         | `ZSCALER_CLIENT_SECRET`  |
| `privateKey`    | _(String)_ A string Private key value.                                                              | `ZSCALER_PRIVATE_KEY`    |
| `vanity_domain` | _(String)_ Refers to the domain name used by your organization.                                     | `ZSCALER_VANITY_DOMAIN`  |
| `zscaler_cloud`         | _(String)_ The name of the Zidentity cloud, e.g., beta.                                             | `ZSCALER_CLOUD`          |

### Alternative OneAPI Cloud Environments

OneAPI supports authentication and can interact with alternative Zscaler enviornments i.e `beta`, `alpha` etc. To authenticate to these environments you must provide the following values:

| Argument         | Description                                                                                         |   | Environment Variable     |
|------------------|-----------------------------------------------------------------------------------------------------|---|--------------------------|
| `vanity_domain`   | _(String)_ Refers to the domain name used by your organization |   | `ZSCALER_VANITY_DOMAIN`  |
| `zscaler_cloud`          | _(String)_ The name of the Zidentity cloud i.e beta      |   | `ZSCALER_CLOUD`          |

For example: Authenticating to Zscaler Beta environment:

```sh
export ZSCALER_VANITY_DOMAIN="acme"
export ZSCALER_CLOUD="beta"
```

## Examples Usage - Client Secret Authentication

{{< chooser os "linux,macos,windows" >}}{{% choosable os linux %}}

```bash
$ export ZSCALER_CLIENT_ID=<ZSCALER_CLIENT_ID>
$ export ZSCALER_CLIENT_SECRET=<ZSCALER_CLIENT_SECRET>
$ export ZSCALER_VANITY_DOMAIN=<ZSCALER_VANITY_DOMAIN>
$ export ZSCALER_CLOUD=<ZSCALER_CLOUD>
```

{{% /choosable %}}

{{% choosable os macos %}}

```bash
$ export ZSCALER_CLIENT_ID=<ZSCALER_CLIENT_ID>
$ export ZSCALER_CLIENT_SECRET=<ZSCALER_CLIENT_SECRET>
$ export ZSCALER_VANITY_DOMAIN=<ZSCALER_VANITY_DOMAIN>
$ export ZSCALER_CLOUD=<ZSCALER_CLOUD>
```

{{% /choosable %}}

{{% choosable os windows %}}

> $env:ZSCALER_CLIENT_ID = "<ZSCALER_CLIENT_ID>"
> $env:ZSCALER_CLIENT_SECRET = "<ZSCALER_CLIENT_SECRET>"
> $env:ZSCALER_VANITY_DOMAIN = "<ZSCALER_VANITY_DOMAIN>"
> $env:ZSCALER_CLOUD = "<ZSCALER_CLOUD>"

{{% /choosable %}}
{{< /chooser >}}

## Examples Usage - Private Key Authentication

{{< chooser os "linux,macos,windows" >}}{{% choosable os linux %}}

```bash
$ export ZSCALER_CLIENT_ID=<ZSCALER_CLIENT_ID>
$ export ZSCALER_PRIVATE_KEY=<ZSCALER_PRIVATE_KEY>
$ export ZSCALER_VANITY_DOMAIN=<ZSCALER_VANITY_DOMAIN>
$ export ZSCALER_CLOUD=<ZSCALER_CLOUD>
```

{{% /choosable %}}

{{% choosable os macos %}}

```bash
$ export ZSCALER_CLIENT_ID=<ZSCALER_CLIENT_ID>
$ export ZSCALER_PRIVATE_KEY=<ZSCALER_PRIVATE_KEY>
$ export ZSCALER_VANITY_DOMAIN=<ZSCALER_VANITY_DOMAIN>
$ export ZSCALER_CLOUD=<ZSCALER_CLOUD>
```

{{% /choosable %}}

{{% choosable os windows %}}

> $env:ZSCALER_CLIENT_ID = "<ZSCALER_CLIENT_ID>"
> $env:ZSCALER_PRIVATE_KEY=<ZSCALER_PRIVATE_KEY>
> $env:ZSCALER_VANITY_DOMAIN=<ZSCALER_VANITY_DOMAIN>
> $env:ZSCALER_CLOUD = "<ZSCALER_CLOUD>"

{{% /choosable %}}
{{< /chooser >}}

**NOTE**: The `ZSCALER_CLOUD` is optional and only required when authenticating to other environments i.e `beta`



## Legacy API Framework

### ZIA native authentication

* As of version v1.0.0, this provider offers backwards compatibility to the Zscaler legacy API framework. This is the recommended authentication method for organizations whose tenants are still not migrated to [Zidentity](https://help.zscaler.com/zidentity/what-zidentity).


Once you have provisioned these credentials, you can set environment variables to provision resources in ZIA:

{{< chooser os "linux,macos,windows" >}}
{{% choosable os linux %}}

```bash
$ export ZIA_USERNAME=<ZIA_USERNAME>
$ export ZIA_PASSWORD=<ZIA_PASSWORD>
$ export ZIA_API_KEY=<ZIA_API_KEY>
$ export ZIA_CLOUD=<ZIA_CLOUD>
$ export ZSCALER_USE_LEGACY_CLIENT=true
```

{{% /choosable %}}

{{% choosable os macos %}}

```bash
$ export ZIA_USERNAME=<ZIA_USERNAME>
$ export ZIA_PASSWORD=<ZIA_PASSWORD>
$ export ZIA_API_KEY=<ZIA_API_KEY>
$ export ZIA_CLOUD=<ZIA_CLOUD>
$ export ZSCALER_USE_LEGACY_CLIENT=true
```

{{% /choosable %}}

{{% choosable os windows %}}

```powershell
> $env:ZIA_USERNAME = "<ZIA_USERNAME>"
> $env:ZIA_PASSWORD = "<ZIA_PASSWORD>"
> $env:ZIA_API_KEY = "<ZIA_API_KEY>"
> $env:ZIA_CLOUD = "<ZIA_CLOUD>"
> $env:ZSCALER_USE_LEGACY_CLIENT = true
```

{{% /choosable %}}
{{< /chooser >}}

## Configuration Options

Once you have provisioned these credentials, you can set environment variables to provision resources in ZIA:

{{< chooser os "linux,macos,windows" >}}
{{% choosable os linux %}}

```bash
$ export ZIA_USERNAME=<ZIA_USERNAME>
$ export ZIA_PASSWORD=<ZIA_PASSWORD>
$ export ZIA_API_KEY=<ZIA_API_KEY>
$ export ZIA_CLOUD=<ZIA_CLOUD>
$ export ZSCALER_USE_LEGACY_CLIENT=true
```

{{% /choosable %}}

{{% choosable os macos %}}

```bash
$ export ZIA_USERNAME=<ZIA_USERNAME>
$ export ZIA_PASSWORD=<ZIA_PASSWORD>
$ export ZIA_API_KEY=<ZIA_API_KEY>
$ export ZIA_CLOUD=<ZIA_CLOUD>
$ export ZSCALER_USE_LEGACY_CLIENT=true
```

{{% /choosable %}}

{{% choosable os windows %}}

```powershell
> $env:ZIA_USERNAME = "<ZIA_USERNAME>"
> $env:ZIA_PASSWORD = "<ZIA_PASSWORD>"
> $env:ZIA_API_KEY = "<ZIA_API_KEY>"
> $env:ZIA_CLOUD = "<ZIA_CLOUD>"
> $env:ZSCALER_USE_LEGACY_CLIENT = true
```

{{% /choosable %}}
{{< /chooser >}}

## Configuration Options

Use `pulumi config set zia:<option>` or pass options to the [constructor of `new zia.Provider`]({{< relref "/registry/packages/zia/api-docs/provider" >}}).

| Option | Required/Optional | Description |
|-----|------|----|
| `username`| Required | [ZIA API Username](https://help.zscaler.com/zia/getting-started-zia-api) |
| `password`| Required | [ZIA API Password](https://help.zscaler.com/zia/getting-started-zia-api) |
| `api_key` | Required | [ZIA API Key](https://help.zscaler.com/zia/getting-started-zia-api) |
| `zia_cloud` | Optional | [ZIA Cloud](https://registry.terraform.io/providers/zscaler/zia/latest/docs) |

Currently the following cloud names are supported:
  * `zscaler`
  * `zscalerone`
  * `zscalertwo`
  * `zscalerthree`
  * `zscloud`
  * `zscalerbeta`
  * `zscalergov`
  * `zscalerten`
  * `zspreview`

### Zscaler Sandbox Authentication

As of version v1.0.0, the ZIA Pulumi provider the legacy sandbox authentication environment variables `ZIA_CLOUD` and `ZIA_SANDBOX_TOKEN` are no longer supported.

Authentication to the Zscaler Sandbox service requires the following new environment variables the `ZSCALER_SANDBOX_CLOUD` and `ZSCALER_SANDBOX_TOKEN` or authentication attributes `sandbox_token` and `sandbox_cloud`. For details on how obtain the API Token visit the Zscaler help portal [About Sandbox API Token](https://help.zscaler.com/zia/about-sandbox-api-token)

Currently the following cloud names are supported:
  * `zscaler`
  * `zscalerone`
  * `zscalertwo`
  * `zscalerthree`
  * `zscloud`
  * `zscalerbeta`
  * `zscalergov`
  * `zscalerten`
  * `zspreview`