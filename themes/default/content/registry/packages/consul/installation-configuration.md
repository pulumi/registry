---
title: HashiCorp Consul Setup
meta_desc: This page provides an overview on how to configure the Pulumi Consul Provider.
layout: installation
---

The Pulumi Consul provider uses the Consul SDK to manage resources.
Pulumi relies on the Consul SDK to authenticate requests from your computer to HashiCorp Consul. Your credentials are never sent
to pulumi.com.

## Installation

The HashiCorp Consul provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/consul`](https://www.npmjs.com/package/@pulumi/consul)
* Python: [`pulumi-consul`](https://pypi.org/project/pulumi-consul/)
* Go: [`github.com/pulumi/pulumi-consul/sdk/v3/go/consul`](https://github.com/pulumi/pulumi-consul)
* .NET: [`Pulumi.Consul`](https://www.nuget.org/packages/Pulumi.Consul)

## Configuring The Provider

Once obtained, there are two ways to communicate your configuration parameters to Pulumi:

1. Set the environment variable `CONSUL_HTTP_ADDR`:

    ```bash
    $ export CONSUL_HTTP_ADDR=XXXXXX
    ```

2. If you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set consul:address XXXXXX
    ```

A full set of configuration parameters can be found listed on the [Project README](https://github.com/pulumi/pulumi-consul/blob/master/README.md).
