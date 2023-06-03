---
title: Keycloak Installation & Configuration
meta_desc: Provides an overview on how to configure credentials for the Pulumi Keycloak Provider.
layout: package
---

The Pulumi Keycloak provider uses the Keycloak SDK to manage and provision resources.

## Installation

The Keycloak provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/keycloak`](https://www.npmjs.com/package/@pulumi/keycloak)
* Python: [`pulumi-keycloak`](https://pypi.org/project/pulumi-keycloak/)
* Go: [`github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak`](https://github.com/pulumi/pulumi-keycloak)
* .NET: [`Pulumi.Keycloak`](https://www.nuget.org/packages/Pulumi.Keycloak)

## Configuring Credentials

Pulumi relies on the Keycloak SDK to authenticate requests from your computer to Keycloak. Your credentials are never sent
to pulumi.com.
The Pulumi Keycloak Provider needs to be configured with Keycloak credentials
before it can be used to create resources.

Once the credentials are obtained, there are two ways to communicate your authorization tokens to Pulumi:

1. Set the environment variables `KEYCLOAK_URL`, `KEYCLOAK_CLIENT_ID` and `KEYCLOAK_CLIENT_SECRET`:

    ```bash
    $ export KEYCLOAK_URL=XXXXXXXXXXXXXX
    $ export KEYCLOAK_CLIENT_ID=YYYYYYYYYYYYYY
    $ export KEYCLOAK_CLIENT_SECRET=ZZZZZZZZZZZZZZ
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set keycloak:url XXXXXXXXXXXXXX
    $ pulumi config set keycloak:clientId YYYYYYYYYYYYYY --secret
    $ pulumi config set keycloak:clientSecret ZZZZZZZZZZZZZZ --secret
    ```

Remember to pass `--secret` when setting `keycloak:clientId` and `keycloak:clientSecret` so that they are properly encrypted. The complete list of
configuration parameters is in the [Keycloak provider README](https://github.com/pulumi/pulumi-keycloak/blob/master/README.md).
