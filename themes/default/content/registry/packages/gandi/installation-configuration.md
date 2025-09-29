---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumiverse/pulumi-gandi/v2.0.0/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumiverse/pulumi-gandi/blob/v2.0.0/docs/installation-configuration.md
title: Gandi Installation & Configuration
meta_desc: Provides an overview on how to configure credentials for the Pulumi Gandi Provider.
layout: installation
---

## Installation

1. To use this package, please [install the Pulumi CLI first](https://www.pulumi.com/docs/get-started/install/).
1. This package is only available for JavaScript and TypeScript but support for other languages will be available soon.

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm` or `yarn`:

```bash
$ npm install @pulumiverse/gandi
$ yarn add @pulumiverse/gandi
```

### Python

```bash
$ pip install pulumiverse_gandi
```
### Go & .NET

> *TBA*

## Authentication

The Pulumi Gandi Provider needs to be configured with a Gandi key.

Once you generated the key there are two ways to communicate your authorization tokens to Pulumi:

1. Set the environment variables `GANDI_KEY`

    ```bash
    $ export GANDI_KEY=xxx
    ```

2. Set them using `pulumi config` command, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set gandi:key xxx --secret
    ```

> Remember to pass `--secret` when setting `gandi:key` so it is properly encrypted.
