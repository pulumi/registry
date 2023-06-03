---
title: Auth0 Installation & Configuration
meta_desc: Provides an overview on how to configure credentials for the Pulumi Auth0 Provider.
layout: package
---

The Pulumi Auth0 provider uses the Auth0 SDK to manage and provision resources.

## Installation

The Auth0 provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/auth0`](https://www.npmjs.com/package/@pulumi/auth0)
* Python: [`pulumi-auth0`](https://pypi.org/project/pulumi-auth0/)
* Go: [`github.com/pulumi/pulumi-auth0/sdk/v2/go/auth0`](https://github.com/pulumi/pulumi-auth0)
* .NET: [`Pulumi.Auth0`](https://www.nuget.org/packages/Pulumi.Auth0)

## Configuring Credentials

Pulumi relies on the Auth0 SDK to authenticate requests from your computer to Auth0. Your credentials are never sent
to pulumi.com. The Pulumi Auth0 Provider needs to be configured with Auth0 credentials
before it can be used to create resources. Once the credentials are obtained, there are two ways to communicate your authorization tokens to Pulumi:

1. Set the environment variables `AUTH0_DOMAIN`, `AUTH0_CLIENT_ID` and `AUTH0_CLIENT_SECRET`:

    ```bash
    $ export AUTH0_DOMAIN=XXXXXXXXXXXXXX
    $ export AUTH0_CLIENT_ID=YYYYYYYYYYYYYY
    $ export AUTH0_CLIENT_SECRET=ZZZZZZZZZZZZZZ
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set auth0:domain XXXXXXXXXXXXXX
    $ pulumi config set auth0:client_id YYYYYYYYYYYYYY --secret
    $ pulumi config set auth0:client_secret ZZZZZZZZZZZZZZ --secret
    ```

Remember to pass `--secret` when setting `auth0:client_id` and `auth0:client_secret` so that it is properly encrypted. The complete list of
configuration parameters is in the [Auth0 Provider README](https://github.com/pulumi/pulumi-auth0/blob/master/README.md).
