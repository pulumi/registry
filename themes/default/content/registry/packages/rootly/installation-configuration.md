---
# WARNING: this file was fetched from https://raw.githubusercontent.com/rootlyhq/pulumi-rootly/v1.5.0/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Rootly Installation & Configuration
meta_desc: Provides an overview on how to configure credentials for the Pulumi Rootly Provider.
layout: package
---

## Installation

1. To use this package, please [install the Pulumi CLI first](https://www.pulumi.com/docs/get-started/install/).
2. This package is only available for JavaScript and TypeScript but support for other languages will be available soon.

### Node.js (JavaScript/TypeScript)

1. To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    ```bash
    npm install @rootly/pulumi
    ```

   or `yarn`:

    ```bash
    yarn add @rootly/pulumi
    ```

1. Install the provider binary plugin. This is only needed due to an outstanding bug in with Pulumi registry

```bash
pulumi plugin install resource rootly v0.0.2 --server https://github.com/rootlyhq/pulumi-rootly/releases/download/v0.0.2
```

## Authentication

The Pulumi Rootly Provider needs to be configured with a Rootly API token.

Once you generated the API token there are two ways to communicate your authorization token to Pulumi:

1. Set the environment variables `ROOTLY_API_TOKEN`:

    ```bash
    $ export ROOTLY_API_TOKEN=cu_xxx
    ```

2. Set them using `pulumi config` command, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set rootly:apiToken cu_xxx --secret
    ```

> Remember to pass `--secret` when setting `rootly:apiToken` so it is properly encrypted.
