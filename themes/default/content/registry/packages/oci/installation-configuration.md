---
title: Oracle Cloud OCI Setup
meta_desc: Provides an overview on how to configure credentials for the Pulumi Oracle Cloud OCI Provider.
layout: installation
---

The Pulumi OCI provider uses the Oracle Cloud SDK to manage and provision resources.

## Installation

The Oracle Cloud provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/oci`](https://www.npmjs.com/package/@pulumi/oci)
* Python: [`pulumi-oci`](https://pypi.org/project/pulumi-oci/)
* Go: [`github.com/pulumi/pulumi-oci/sdk/go/oci`](https://github.com/pulumi/pulumi-oci)
* .NET: [`Pulumi.Oci`](https://www.nuget.org/packages/Pulumi.Oci)

## Configuring Credentials

Pulumi relies on the Oracle Cloud SDK to authenticate requests from your computer to Oracle Cloud. Your credentials are never sent
to pulumi.com.
The Pulumi OCI Provider needs to be configured with Oracle Cloud credentials
before it can be used to create resources.

Once the credentials are obtained, there are two ways to communicate your authorization tokens to Pulumi:

1. Set the environment variable `TF_VAR_secret_key`:

    ```bash
    $ export TF_VAR_secret_key=XXXXXXXXXXXXXX
    $ export TF_VAR_region=YYYYYYYYYYYYYYY
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set oci:secretKey XXXXXXXXXXXXXX --secret
    $ pulumi config set oci:region YYYYYYYYYYYYYYY
    ```

Remember to pass `--secret` when setting `oci:secretKey` so that it is properly encrypted. The complete list of
configuration parameters is in the [Oci provider README](https://github.com/pulumi/pulumi-oci/blob/master/README.md).
