---
# WARNING: this file was fetched from https://raw.githubusercontent.com/koyeb/pulumi-koyeb/v0.1.7/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Koyeb Setup
meta_desc: Information on how to install the Koyeb provider.
layout: installation
---

## Installation

The Pulumi Koyeb provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@koyeb/pulumi-koyeb`](https://www.npmjs.com/package/@koyeb/pulumi-koyeb)
* Go: [`github.com/koyeb/pulumi-koyeb/sdk/go/koyeb`](https://pkg.go.dev/github.com/koyeb/pulumi-koyeb/sdk)
* Python: [`pulumi_koyeb`](https://pypi.org/project/pulumi-koyeb/)

## Setup

To provision resources with the Pulumi Koyeb provider, you need a Koyeb API access token to allow Pulumi to authenticate to Koyeb. Koyeb API access token credentials can be obtained from the [Koyeb Control Panel](https://app.koyeb.com/account/api).

### Set environment variables

Once the credentials are obtained, there are two ways to communicate your authorization tokens to Pulumi:

Set the environment variable `KOYEB_TOKEN`:

    ```bash
    $ export KOYEB_TOKEN=XXXXXXXXXXXXXX
    ```
