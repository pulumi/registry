---
title: DigitalOcean Setup
meta_desc: Provides an overview on how to configure credentials for the Pulumi DigitalOcean Provider.
layout: installation
---

The Pulumi DigitalOcean provider uses the DigitalOcean SDK to manage and provision resources.

> Pulumi relies on the DigitalOcean SDK to authenticate requests from your computer to DigitalOcean. Your credentials are never sent
> to pulumi.com.

The Pulumi DigitalOcean Provider needs to be configured with DigitalOcean credentials
before it can be used to create resources.

## Installation

The DigitalOcean provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/digitalocean`](https://www.npmjs.com/package/@pulumi/digitalocean)
* Python: [`pulumi-digitalocean`](https://pypi.org/project/pulumi-digitalocean/)
* Go: [`github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean`](https://github.com/pulumi/pulumi-digitalocean)
* .NET: [`Pulumi.DigitalOcean`](https://www.nuget.org/packages/Pulumi.DigitalOcean)

## Configuring Credentials

Once the credetials are obtained, there are two ways to communicate your authorization tokens to Pulumi:

1. Set the environment variable `DIGITALOCEAN_TOKEN`:

    ```bash
    $ export DIGITALOCEAN_TOKEN=XXXXXXXXXXXXXX
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set digitalocean:token XXXXXXXXXXXXXX --secret
    ```

Remember to pass `--secret` when setting `token` so that it is properly encrypted.
