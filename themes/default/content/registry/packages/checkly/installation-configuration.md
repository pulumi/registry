---
title: Checkly Installation & Configuration
meta_desc: Provides an overview on how to configure credentials for the Pulumi Checkly Provider.
layout: installation
---

## Installation

1. To use this package, please [install the Pulumi CLI first](https://www.pulumi.com/docs/get-started/install/).
1. This package is only available for JavaScript and TypeScript but support for other languages will be available soon.

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm` or `yarn`:

```bash
$ npm install @checkly/pulumi
$ yarn add @checkly/pulumi
```

### Python, Go & .NET

> *TBA*

## Authentication

The Pulumi Checkly Provider needs to be configured with a Checkly `API Key` and `Account ID` before it can be used to create resources.

> If you don't have an `API Key`, you can create one [here](https://app.checklyhq.com/settings/user/api-keys).

Once you generated the `API Key` there are two ways to communicate your authorization tokens to Pulumi:

1. Set the environment variables `CHECKLY_API_KEY` and `CHECKLY_ACCOUNT_ID`:

    ```bash
    $ export CHECKLY_API_KEY=cu_xxx
    $ export CHECKLY_ACCOUNT_ID=xxx
    ```

2. Set them using `pulumi config` command, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set checkly:apiKey cu_xxx --secret
    $ pulumi config set checkly:accountId xxx
    ```

> Remember to pass `--secret` when setting `checkly:apiKey` so it is properly encrypted.
