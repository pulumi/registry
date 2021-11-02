---
title: HashiCorp Nomad Setup
meta_desc: Provides an overview on how to configure the Pulumi Nomad Provider.
layout: installation
---

The [Pulumi Nomad provider]({{< relref "./" >}}) uses the Nomad SDK to manage resources.
Pulumi relies on the Nomad SDK to authenticate requests from your computer to HashiCorp Nomad. Your credentials are never sent
to pulumi.com.

## Installation

The HashiCorp Nomad provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/nomad`](https://www.npmjs.com/package/@pulumi/nomad)
* Python: [`pulumi-nomad`](https://pypi.org/project/pulumi-nomad/)
* Go: [`github.com/pulumi/pulumi-nomad/sdk/go/nomad`](https://github.com/pulumi/pulumi-nomad)
* .NET: [`Pulumi.Nomad`](https://www.nuget.org/packages/Pulumi.Nomad)

## Configuring The Provider

Once the credetials are obtained, there are two ways to communicate your configuration tokens to Pulumi:

1. Set the environment variables `NOMAD_ADDR` and `NOMAD_TOKEN`:

    ```bash
    $ export NOMAD_ADDR=XXXXXX
    $ export NOMAD_TOKEN=YYYYYY
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set nomad:address  XXXXXX
    $ pulumi config set nomad:secretId YYYYYY --secret
    ```

The complete list of
configuration parameters is in the  [HashiCorp Nomad provider README](https://github.com/pulumi/pulumi-nomad/blob/master/README.md).
Remember to pass `--secret` when setting `secretId` so that it is properly encrypted.
