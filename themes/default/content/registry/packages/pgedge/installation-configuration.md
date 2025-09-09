---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pgEdge/pulumi-pgedge/v0.0.42/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: pgEdge Installation & Configuration
meta_desc: Provides an overview on how to configure credentials for the Pulumi pgEdge Provider.
layout: package
---

## Installation

1. To use this package, please [install the Pulumi CLI first](https://www.pulumi.com/docs/install/).

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm` or `yarn`:

```bash
$ npm install @pgedge/pulumi-pgedge
$ yarn add @pgedge/pulumi-pgedge
```

### Go

```bash
$ go get -u github.com/pgEdge/pulumi-pgedge/sdk
```

## Authentication

The Pulumi pgEdge Provider needs to be configured with a pgEdge `Client ID` and `Client Secret` before it can be used to create resources.

> If you don't have a `Client ID` or `Client Secret`, you can create one in [pgEdge Cloud](https://app.pgedge.com/) > Settings > API Clients.

Once you generated the `Client ID` and `Client Secret` set the environment variables as `PGEDGE_CLIENT_ID` and `PGEDGE_CLIENT_SECRET` in your shell:

    ```bash
    $ export PGEDGE_CLIENT_ID=xxx
    $ export PGEDGE_CLIENT_SECRET=xxx
    ```