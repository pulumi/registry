---
title: Akamai Setup
meta_desc: Provides an overview on how to configure credentials for the Pulumi Akamai Provider.
layout: installation
---

The Pulumi Akamai provider uses the Akamai SDK to manage and provision resources.

## Installation

The Akamai provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/akamai`](https://www.npmjs.com/package/@pulumi/akamai)
* Python: [`pulumi-akamai`](https://pypi.org/project/pulumi-akamai/)
* Go: [`github.com/pulumi/pulumi-akamai/sdk/v2/go/akamai`](https://github.com/pulumi/pulumi-akamai)
* .NET: [`Pulumi.Akamai`](https://www.nuget.org/packages/Pulumi.Akamai)

## Configuring Credentials

Pulumi relies on the Akamai SDK to authenticate requests from your computer to Akamai. Your credentials are never sent
to pulumi.com. The Pulumi Akamai Provider needs to be configured with Akamai credentials
before it can be used to create resources. Once the credentials are obtained, there are two ways to communicate your authorization tokens to Pulumi:

1. Create environment variables in the format:

    `AKAMAI{_SECTION_NAME}_*`

    For example, if you specify akamai:propertySection papi you would set the following ENV variables:

    `AKAMAI_PAPI_HOST`  
    `AKAMAI_PAPI_ACCESS_TOKEN`  
    `AKAMAI_PAPI_CLIENT_TOKEN`  
    `AKAMAI_PAPI_CLIENT_SECRET`

    If the section name is `default`, you can omit it, instead using:

    `AKAMAI_HOST`  
    `AKAMAI_ACCESS_TOKEN`  
    `AKAMAI_CLIENT_TOKEN`  
    `AKAMAI_CLIENT_SECRET`

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set akamai:edgerc XXXXXXXXXXXXXX
    $ pulumi config set akamai:propertySection YYYYYYYYYYYYYY
    ```

A full set of configuration parameters
is listed in the [Akamai Provider README](https://github.com/pulumi/pulumi-akamai/blob/master/README.md).
