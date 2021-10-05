---
title: Equinix Metal Setup
meta_desc: This page provides an overview of how to setup the Equinix Metal SDK to manage and provision resources.
layout: installation
---

The Pulumi Equinix Metal provider uses the Equinix Metal SDK to manage and provision resources.

> Pulumi relies on the Equinix Metal SDK to authenticate requests from your computer to Equinix Metal. Your credentials are never sent
> to pulumi.com.

The Pulumi Equinix Metal Provider needs to be configured with Equinix Metal credentials
before it can be used to create resources.

### Configuring Credentials

Once obtained, there are two ways to communicate your authorization tokens to Pulumi:

1. Set the environment variable `PACKET_AUTH_TOKEN`:

    ```bash
    $ export PACKET_AUTH_TOKEN=XXXXXXXXXXXXXX
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set equinix-metal:authToken XXXXXXXXXXXXXX --secret
    ```

Remember to pass `--secret` when setting `authToken` so that it is properly encrypted.
