---
title: Checkly Setup
meta_desc: Provides an overview on how to configure credentials for the Pulumi Checkly Provider.
layout: installation
---

## Installation

The Checkly provider is available as JavaScript/TypeScript package (support for other languages is coming soon):

* JavaScript/TypeScript: [`@checkly/pulumi`](https://www.npmjs.com/package/@checkly/pulumi)

## Configuring Credentials

Pulumi relies on the Checkly GO SDK to authenticate requests from your computer to Checkly. Your credentials are never sent
to pulumi.com. The Pulumi Checkly Provider needs to be configured with Checkly `API Key` and `Account ID` before it can be used to create resources.

> If you don't have and `API Key`, you can create one [here](https://app.checklyhq.com/settings/user/api-keys).

Once you generated the `API Key` there are two ways to communicate your authorization keys to Pulumi:

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
