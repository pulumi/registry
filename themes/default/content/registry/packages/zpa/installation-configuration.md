---
# WARNING: this file was fetched from https://raw.githubusercontent.com/zscaler/pulumi-zpa/v1.0.2/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Zscaler Private Access (ZPA) Installation & Configuration
meta_desc: Information on how to install the ZPA provider.
layout: package
---

## Installation

The Pulumi ZPA provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@bdzscaler/pulumi-zpa`](https://www.npmjs.com/package/@bdzscaler/pulumi-zpa)
* Python: [`zscaler_pulumi_zpa`](https://pypi.org/project/zscaler-pulumi-zpa/)
* Go: [`github.com/zscaler/pulumi-zpa/sdk/go/zpa`](https://pkg.go.dev/github.com/zscaler/pulumi-zpa/sdk)
* .NET: [`zscaler.PulumiPackage.Zpa`](https://www.nuget.org/packages/zscaler.Zpa)

### Provider Binary

The ZPA provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource zpa <version> --server github://api.github.com/zscaler
```

Replace the version string with your desired version.

## Setup

To provision resources with the Pulumi ZPA provider, you need to have ZPA credentials. Zscaler maintains documentation on how to create API keys [here](https://help.zscaler.com/zpa/about-api-keys)

## Zscaler OneAPI New Framework

The ZPA Terraform Provider now offers support for [OneAPI](https://help.zscaler.com/oneapi/understanding-oneapi) Oauth2 authentication through [Zidentity](https://help.zscaler.com/zidentity/what-zidentity).

**NOTE** As of version v4.0.0, this Terraform provider offers backwards compatibility to the Zscaler legacy API framework. This is the recommended authentication method for organizations whose tenants are still not migrated to [Zidentity](https://help.zscaler.com/zidentity/what-zidentity). 

**NOTE** Notice that OneAPI and Zidentity is not currently supported for the following clouds: `GOV` and `GOVUS`. Refer to the [Legacy API Framework](#legacy-api-framework) for more information on how authenticate to these environments

## Authentication - OneAPI New Framework

As of version v4.0.0, this provider supports authentication via the new Zscaler API framework [OneAPI](https://help.zscaler.com/oneapi/understanding-oneapi)

Zscaler OneAPI uses the OAuth 2.0 authorization framework to provide secure access to Zscaler Private Access (ZPA) APIs. OAuth 2.0 allows third-party applications to obtain controlled access to protected resources using access tokens. OneAPI uses the Client Credentials OAuth flow, in which client applications can exchange their credentials with the authorization server for an access token and obtain access to the API resources, without any user authentication involved in the process.

* [ZPA API](https://help.zscaler.com/oneapi/understanding-oneapi#:~:text=Workload%20Groups-,ZPA%20API,-Zscaler%20Private%20Access)

### OneAPI (API Client Scope)

OneAPI Resources are automatically created within the ZIdentity Admin UI based on the RBAC Roles
applicable to APIs within the various products. For example, in ZPA, navigate to `Administration -> Role
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
| `customer_id`   | _(String)_ A string that contains the ZPA customer ID which identifies the tenant                   | `ZPA_CUSTOMER_ID`    |
| `microtenant_id`| _(String)_ A string that contains the ZPA microtenant ID which identifies the tenant                | `ZPA_MICROTENANT_ID`    |
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
$ export ZPA_CUSTOMER_ID=<ZPA_CUSTOMER_ID>
$ export ZPA_MICROTENANT_ID=<ZPA_MICROTENANT_ID>
```

{{% /choosable %}}

{{% choosable os macos %}}

```bash
$ export ZSCALER_CLIENT_ID=<ZSCALER_CLIENT_ID>
$ export ZSCALER_CLIENT_SECRET=<ZSCALER_CLIENT_SECRET>
$ export ZSCALER_VANITY_DOMAIN=<ZSCALER_VANITY_DOMAIN>
$ export ZSCALER_CLOUD=<ZSCALER_CLOUD>
$ export ZPA_CUSTOMER_ID=<ZPA_CUSTOMER_ID>
$ export ZPA_MICROTENANT_ID=<ZPA_MICROTENANT_ID>
```

{{% /choosable %}}

{{% choosable os windows %}}

> $env:ZSCALER_CLIENT_ID = "<ZSCALER_CLIENT_ID>"
> $env:ZSCALER_CLIENT_SECRET = "<ZSCALER_CLIENT_SECRET>"
> $env:ZSCALER_VANITY_DOMAIN = "<ZSCALER_VANITY_DOMAIN>"
> $env:ZSCALER_CLOUD = "<ZSCALER_CLOUD>"
> $env:ZPA_CUSTOMER_ID = "<ZPA_CUSTOMER_ID>"
> $env:ZPA_MICROTENANT_ID = "<ZPA_MICROTENANT_ID>"

{{% /choosable %}}
{{< /chooser >}}

## Examples Usage - Private Key Authentication

{{< chooser os "linux,macos,windows" >}}{{% choosable os linux %}}

```bash
$ export ZSCALER_CLIENT_ID=<ZSCALER_CLIENT_ID>
$ export ZSCALER_PRIVATE_KEY=<ZSCALER_PRIVATE_KEY>
$ export ZSCALER_VANITY_DOMAIN=<ZSCALER_VANITY_DOMAIN>
$ export ZSCALER_CLOUD=<ZSCALER_CLOUD>
$ export ZPA_CUSTOMER_ID=<ZPA_CUSTOMER_ID>
$ export ZPA_MICROTENANT_ID=<ZPA_MICROTENANT_ID>
```

{{% /choosable %}}

{{% choosable os macos %}}

```bash
$ export ZSCALER_CLIENT_ID=<ZSCALER_CLIENT_ID>
$ export ZSCALER_PRIVATE_KEY=<ZSCALER_PRIVATE_KEY>
$ export ZSCALER_VANITY_DOMAIN=<ZSCALER_VANITY_DOMAIN>
$ export ZSCALER_CLOUD=<ZSCALER_CLOUD>
$ export ZPA_CUSTOMER_ID=<ZPA_CUSTOMER_ID>
$ export ZPA_MICROTENANT_ID=<ZPA_MICROTENANT_ID>
```

{{% /choosable %}}

{{% choosable os windows %}}

> $env:ZSCALER_CLIENT_ID = "<ZSCALER_CLIENT_ID>"
> $env:ZSCALER_PRIVATE_KEY=<ZSCALER_PRIVATE_KEY>
> $env:ZSCALER_VANITY_DOMAIN=<ZSCALER_VANITY_DOMAIN>
> $env:ZSCALER_CLOUD = "<ZSCALER_CLOUD>"
> $env:ZPA_CUSTOMER_ID = "<ZPA_CUSTOMER_ID>"
> $env:ZPA_MICROTENANT_ID = "<ZPA_MICROTENANT_ID>"

{{% /choosable %}}
{{< /chooser >}}

**NOTE**: The `ZSCALER_CLOUD` is optional and only required when authenticating to other environments i.e `beta`


## Legacy API Framework

### ZPA native authentication

* As of version v0.1.0, this provider offers backwards compatibility to the Zscaler legacy API framework. This is the recommended authentication method for organizations whose tenants are still not migrated to [Zidentity](https://help.zscaler.com/zidentity/what-zidentity).


Once you have provisioned these credentials, you can set environment variables to provision resources in ZPA:

{{< chooser os "linux,macos,windows" >}}
{{% choosable os linux %}}

```bash
$ export ZPA_CLIENT_ID=<ZPA_CLIENT_ID>
$ export ZPA_CLIENT_SECRET=<ZPA_CLIENT_SECRET>
$ export ZPA_CUSTOMER_ID=<ZPA_CUSTOMER_ID>
$ export ZPA_CLOUD=<ZPA_CLOUD>
$ export ZSCALER_USE_LEGACY_CLIENT=true
```

{{% /choosable %}}

{{% choosable os macos %}}

```bash
$ export ZPA_CLIENT_ID=<ZPA_CLIENT_ID>
$ export ZPA_CLIENT_SECRET=<ZPA_CLIENT_SECRET>
$ export ZPA_CUSTOMER_ID=<ZPA_CUSTOMER_ID>
$ export ZPA_CLOUD=<ZPA_CLOUD>
$ export ZSCALER_USE_LEGACY_CLIENT=true
```

{{% /choosable %}}

{{% choosable os windows %}}

```powershell
> $env:ZPA_CLIENT_ID = "<ZPA_CLIENT_ID>"
> $env:ZPA_CLIENT_SECRET = "<ZPA_CLIENT_SECRET>"
> $env:ZPA_CUSTOMER_ID = "<ZPA_CUSTOMER_ID>"
> $env:ZPA_CLOUD = "<ZPA_CLOUD>"
> $env:ZSCALER_USE_LEGACY_CLIENT = true
```

{{% /choosable %}}
{{< /chooser >}}

## Configuration Options

Use `pulumi config set zpa:<option>` or pass options to the [constructor of `new zpa.Provider`]({{< relref "/registry/packages/zpa/api-docs/provider" >}}).

| Option | Required/Optional | Description |
|-----|------|----|
| `zpa_client_id`| Required | [ZPA Client ID](https://help.zscaler.com/zpa/about-api-keys) |
| `zpa_client_secret`| Required | [ZPA Client Secret](https://help.zscaler.com/zpa/about-api-keys) |
| `zpa_customer_id` | Required | [ZPA Customer ID](https://help.zscaler.com/zpa/about-api-keys) |
| `zpa_cloud` | Optional | [ZPA Cloud Name](https://registry.terraform.io/providers/zscaler/zpa/latest/docs) |
| `use_legacy_client` | Optional | Enable Legacy Mode

~> **NOTE** The environment variable `ZPA_CLOUD` is only required if your are interacting with any of the below environments listed below:

- `BETA`
- `GOV`
- `GOVUS`
- `PREVIEW`
- `ZPATWO`

### Zscaler Private Access Microtenant

A Microtenant is a delegated administrator responsibility that is assigned to an admin by an admin with Microtenant administrator privileges. Microtenants are defined by an authentication domain and assigned to admins based on country, department, and company for role-based administration control. A Microtenant is created within a tenant and is used when departments or subsidiaries within an organization want to manage their configurations independently.[Read More](https://help.zscaler.com/zpa/about-microtenants)

To manage a microtenant using the ZPA Terraform provider, the administrator for the parent or default tenant, must first provision the microtenant using the resource `zpa_microtenant_controller`. The resource will output the administrator credentials for the new microtenant, which can then be provided to the microtenant administrator.

The microtenant administrator can then create his own microtenant API credentials required to authenticate via API to the ZPA platform. From that point, the administrator can then individually manage his own resources in an isolated manner.
When authenticating to microtenant via API using the ZPA Terraform provider, the administrator must provide the following environment variable credentials: `ZPA_CLIENT_ID`, `ZPA_CLIENT_SECRET`, `ZPA_CUSTOMER_ID`, `ZPA_CLOUD`, `ZPA_MICROTENANT_ID`
