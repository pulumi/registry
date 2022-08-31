---
title: PagerDuty Setup
meta_desc: Provides an overview on how to configure credentials for the Pulumi PagerDuty Provider.
layout: installation
---

The Pulumi PagerDuty provider uses the PagerDuty SDK to manage and provision resources.

## Installation

The PagerDuty provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/pagerduty`](https://www.npmjs.com/package/@pulumi/pagerduty)
* Python: [`pulumi-pagerduty`](https://pypi.org/project/pulumi-pagerduty/)
* Go: [`github.com/pulumi/pulumi-pagerduty/sdk/v2/go/pagerduty`](https://github.com/pulumi/pulumi-pagerduty)
* .NET: [`Pulumi.Pagerduty`](https://www.nuget.org/packages/Pulumi.Pagerduty)

## Configuring Credentials

Pulumi relies on the PagerDuty SDK to authenticate requests from your computer to PagerDuty. Your credentials are never sent
to pulumi.com.
The Pulumi PagerDuty Provider needs to be configured with PagerDuty credentials
before it can be used to create resources.

Once the credentials are obtained, there are two ways to communicate your authorization tokens to Pulumi:

1. Set the environment variable `PAGERDUTY_TOKEN`:

    ```bash
    $ export PAGERDUTY_TOKEN=XXXXXXXXXXXXXX
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set pagerduty:token XXXXXXXXXXXXXX --secret
    ```

Remember to pass `--secret` when setting `pagerduty:token` so that it is properly encrypted. The complete list of
configuration parameters is in the [PagerDuty provider README](https://github.com/pulumi/pulumi-pagerduty/blob/master/README.md).
