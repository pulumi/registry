---
# WARNING: this file was fetched from https://raw.githubusercontent.com/EventStore/pulumi-eventstorecloud/v0.2.20/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Event Store Cloud Installation & Configuration
meta_desc: How to set up credentials to use the Event Store Cloud provider for Pulumi.
layout: package
---

## Installation

The Event Store Cloud provider is available as a package in all Pulumi languages:

- JavaScript/TypeScript: [`@eventstore/pulumi-eventstorecloud`](https://www.npmjs.com/package/@eventstore/pulumi-eventstorecloud)
- Go: [`github.com/EventStore/pulumi-eventstorecloud/sdk/go/eventstorecloud`](https://github.com/EventStore/pulumi-eventstorecloud)
- .NET: [`Pulumi.EventStoreCloud`](https://www.nuget.org/packages/Pulumi.EventStoreCloud)

## Setup

### Download the provider

For projects that use .NET and Go Pulumi SDK you have to install the provider before trying to update the stack.

Use the following command to add the plugin to your environment:

```
pulumi plugin install resource eventstorecloud --server github://api.github.com/EventStore
```

### Configure the provider

The Pulumi provider needs credentials to authenticate requests from your computer to Event Store Cloud. Your credentials are never sent
to pulumi.com. The provider needs to be configured with Event Store Cloud credentials before it can be used to create resources.

First, you need an [access token](https://developers.eventstore.com/cloud/automation/#obtaining-the-access-token) for your user.

Then, go to the list of organizations you have access to in Event Store Cloud console, choose the organization that you will be provisioning resources for, and look the organization id in the settings.

- `<YOUR_ACCESS_TOKEN>`: your access token
- `<YOUR_ORGANIZATION_ID>`: the Event Store Cloud organization id

Once the credentials are obtained, there are two ways to communicate your authorization tokens to Pulumi:

1. Set the environment variables `ESC_TOKEN` and `ESC_ORG_ID`:

    ```bash
    $ export ESC_TOKEN=<YOUR_ACCESS_TOKEN>
    $ export ESC_ORG_ID=<YOUR_ORGANIZATION_ID>
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    pulumi config set eventstorecloud:token <YOUR_ACCESS_TOKEN> --secret
    pulumi config set eventstorecloud:organizationId <YOUR_ORGANIZATION_ID> --secret
    ```

{{% notes "info" %}}
Required options can be omitted if you configure them using environment variables.
{{% /notes %}}

| Option           | Required/Optional | Description                                                                                              |
| ---------------- | ----------------- | -------------------------------------------------------------------------------------------------------- |
| `token`          | Required          | Access token. You can retrieve this from the ‘Access Tokens’ section of the Event Store Cloud console.   |
| `organizationId` | Required          | The organization id. You can find it in the organization settings page of the Event Store Cloud console. |
| `url`            | Optional          | The URL of the Event Store Cloud API. This defaults to the public cloud instance of Event Store Cloud.   |
| `tokenStore`     | Optional          | The location on the local filesystem of the token cache. This is shared with the Event Store Cloud CLI.  |
