---
# WARNING: this file was fetched from https://raw.githubusercontent.com/kurrent-io/pulumi-kurrentcloud/v0.3.1/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/kurrent-io/pulumi-kurrentcloud/blob/v0.3.1/docs/installation-configuration.md
title: Kurrent Cloud Installation & Configuration
meta_desc: How to set up credentials to use the Kurrent Cloud provider for Pulumi.
layout: package
---

## Installation

The Kurrent Cloud provider is available as a package in all Pulumi languages:

- JavaScript/TypeScript: [`@kurrent/pulumi-kurrentcloud`](https://www.npmjs.com/package/@kurrent/pulumi-kurrentcloud)
- Python: [`pulumi_kurrentcloud`](https://pypi.org/project/pulumi-kurrentcloud/)
- Go: [`github.com/kurrent-io/pulumi-kurrentcloud/sdk/go/kurrentcloud`](https://github.com/kurrent-io/pulumi-kurrentcloud)
- .NET: [`Kurrent.Pulumi.KurrentCloud`](https://www.nuget.org/packages/Kurrent.Pulumi.KurrentCloud)

## Setup

### Configure the provider

The Pulumi provider needs credentials to authenticate requests from your computer to Kurrent Cloud. Your credentials are never sent
to pulumi.com. The provider needs to be configured with Kurrent Cloud credentials before it can be used to create resources.

The provider supports two authentication methods:

- **Service Account credentials (recommended for automation)** - configure `clientId` and `clientSecret` with the credentials of a Kurrent Cloud Service Account. The provider then obtains short-lived access tokens via the OAuth2 client-credentials grant, and `token` is not used.
- **Refresh token** - configure `token` with a personal refresh token. This is the legacy method; the token is bound to a user account.

In both cases you also need the organization id: go to the list of organizations you have access to in the Kurrent Cloud console, choose the organization that you will be provisioning resources for, and find the organization id in the settings.

**Service Account credentials** can be provided via environment variables:

```bash
$ export ESC_CLIENT_ID=<YOUR_CLIENT_ID>
$ export ESC_CLIENT_SECRET=<YOUR_CLIENT_SECRET>
$ export ESC_ORG_ID=<YOUR_ORGANIZATION_ID>
```

or via stack configuration:

```bash
pulumi config set kurrentcloud:clientId <YOUR_CLIENT_ID>
pulumi config set kurrentcloud:clientSecret <YOUR_CLIENT_SECRET> --secret
pulumi config set kurrentcloud:organizationId <YOUR_ORGANIZATION_ID> --secret
```

**Refresh token** authentication works the same way. First, you need an access token for your user, which you can obtain from the Kurrent Cloud console.

- `<YOUR_ACCESS_TOKEN>`: your access token
- `<YOUR_ORGANIZATION_ID>`: the Kurrent Cloud organization id

Once the credentials are obtained, there are two ways to communicate your authorization tokens to Pulumi:

1. Set the environment variables `ESC_TOKEN` and `ESC_ORG_ID`:

    ```bash
    $ export ESC_TOKEN=<YOUR_ACCESS_TOKEN>
    $ export ESC_ORG_ID=<YOUR_ORGANIZATION_ID>
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    pulumi config set kurrentcloud:token <YOUR_ACCESS_TOKEN> --secret
    pulumi config set kurrentcloud:organizationId <YOUR_ORGANIZATION_ID> --secret
    ```

{{% notes "info" %}}
Required options can be omitted if you configure them using environment variables.
{{% /notes %}}

| Option           | Required/Optional | Description                                                                                       |
| ---------------- | ----------------- | ------------------------------------------------------------------------------------------------- |
| `token`          | Required unless Service Account credentials are set | Refresh token. You can retrieve this from the ‘Access Tokens’ section of the Kurrent Cloud console. |
| `organizationId` | Required          | The organization id. You can find it in the organization settings page of the Kurrent Cloud console. |
| `clientId`       | Optional          | Service Account client id. Must be set together with `clientSecret`.                              |
| `clientSecret`   | Optional          | Service Account client secret. When both `clientId` and `clientSecret` are set, Service Account authentication is used and takes priority over `token`. |
| `url`            | Optional          | The URL of the Kurrent Cloud API. This defaults to the public cloud instance of Kurrent Cloud.    |
| `tokenStore`     | Optional          | The location on the local filesystem of the token cache. This is shared with the Kurrent CLI. Only used with refresh token authentication. |
| `identityKitUrl` | Optional          | The base URL of the token endpoint used for Service Account authentication. You would normally not need to set it. |
