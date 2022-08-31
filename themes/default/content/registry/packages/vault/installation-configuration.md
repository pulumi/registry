---
title: HashiCorp Vault Setup
meta_desc: Provides an overview on how to configure the Pulumi Vault Provider.
layout: installation
---

The Pulumi Vault provider uses the Vault SDK to manage resources.

## Installation

The HashiCorp Vault provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/vault`](https://www.npmjs.com/package/@pulumi/vault)
* Python: [`pulumi-vault`](https://pypi.org/project/pulumi-vault/)
* Go: [`github.com/pulumi/pulumi-vault/sdk/v4/go/vault`](https://github.com/pulumi/pulumi-vault)
* .NET: [`Pulumi.Vault`](https://www.nuget.org/packages/Pulumi.Vault)

## Configuration

Pulumi relies on the Vault SDK to authenticate requests from your computer to HashiCorp Vault. Your credentials are never sent
to pulumi.com.

Once the credentials are obtained, there are two ways to communicate your configuration tokens to Pulumi:

1. Set the environment variables `VAULT_ADDR` and `VAULT_TOKEN`:

    ```bash
    $ export VAULT_ADDR=XXXXXX
    $ export VAULT_TOKEN=YYYYYY
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set vault:address XXXXXX
    $ pulumi config set vault:token YYYYYY --secret
    ```

The complete list of
configuration parameters is in the [HashiCorp Vault provider README](https://github.com/pulumi/pulumi-vault/blob/master/README.md).
Remember to pass `--secret` when setting `token` so that it is properly encrypted.
