---
title: Wavefront Setup
meta_desc: This page provides an overview on how to configure credentials for the Pulumi Wavefront Provider.
layout: installation
aliases:
    - "/docs/intro/cloud-providers/wavefront/setup/"
---

The Pulumi Wavefront provider uses the Wavefront SDK to manage and provision resources.

> Pulumi relies on the Wavefront SDK to authenticate requests from your computer to Wavefront. Your credentials are never sent
> to pulumi.com.

The Pulumi Wavefront Provider needs to be configured with Wavefront credentials
before it can be used to create resources.

### Configuring Credentials

Once obtained, there are two ways to communicate your authorization tokens to Pulumi:

1. Set the environment variables for `WAVEFRONT_TOKEN` and `WAVEFRONT_ADDRESS`:

    ```bash
    $ export WAVEFRONT_TOKEN=XXXXXXXXXXXXXX
    $ export WAVEFRONT_ADDRESS=YYYYYYYYYYYYYY
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set wavefront:token XXXXXXXXXXXXXX --secret
    $ pulumi config set wavefront:address YYYYYYYYYYYYYY
    ```

Remember to pass `--secret` when setting `wavefront:token` so that it is properly encrypted. A full set of configuration parameters
can be found listed on the [Project README](https://github.com/pulumi/pulumi-wavefront/blob/master/README.md).
