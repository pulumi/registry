---
# WARNING: this file was fetched from https://raw.githubusercontent.com/ovh/pulumi-ovh/v2.4.0/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: OVH Installation & Configuration
meta_desc: Information on how to install the OVH provider.
layout: installation
---

## Information

Note that the [lbrlabs Pulumi OVH provider](https://github.com/lbrlabs/pulumi-ovh) is replaced by this official one.

## Installation

The Pulumi `OVH` provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@ovhcloud/pulumi-ovh`](https://www.npmjs.com/package/@ovhcloud/pulumi-ovh)
* Python: [`pulumi_ovh`](https://pypi.org/project/pulumi-ovh/)
* Go: [`github.com/ovh/pulumi-ovh/sdk/v2/go/ovh`](https://pkg.go.dev/github.com/ovh/pulumi-ovh/sdk/v2)
* .NET: [`Pulumi.Ovh`](https://www.nuget.org/packages/Pulumi.Ovh)
* Java: [`com.ovhcloud.pulumi.ovh`](https://s01.oss.sonatype.org/#nexus-search;quick~com.ovhcloud.pulumi.ovh)

### Provider Binary

The OVH provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource ovh vX.Y.Z --server github://api.github.com/ovh/pulumi-ovh
```

Replace the version string with your desired version.

## Setup

To provision resources with the Pulumi OVH provider, you need to have OVH credentials.
Your OVH credentials are never sent to pulumi.com. Pulumi uses the OVH API and the credentials in your environment to authenticate requests from your computer to OVH.

The provider needs to be configured with the proper credentials before it can be used. Requests to OVHcloud APIs require a set of secrets keys and the definition of the API endpoint. See [First Steps with the API](https://docs.ovh.com/gb/en/customer/first-steps-with-ovh-api/) (or the French version, [Premiers pas avec les API OVHcloud](https://docs.ovh.com/fr/api/api-premiers-pas/)) for a detailed explanation.

Three forms of authentication are supported by the provider:

* OAuth2, using scoped service accounts, and compatible with OVHcloud IAM
* Short-lived access token received from [OVH API](https://support.us.ovhcloud.com/hc/en-us/articles/19901571606547-Using-Service-Accounts-to-Connect-to-OVHcloud-APIs) (for example with the help of Hashicorp Vault OAuth2 secret engine configured to work with OVH auth API).
* Application key & application secret & consumer key

### Get your credentials

The "OVH provider" needs to be configured with a set of credentials, which can be set using
[Pulumi stack configuration](https://www.pulumi.com/docs/concepts/config/) or environment variables.

#### OAuth2

First, you need to generate a pair of valid `client_id` and `client_secret`: you can proceed by [following this documentation](https://help.ovhcloud.com/csm/en-manage-service-account?id=kb_article_view&sysparm_article=KB0059343).

Once you have retrieved your `client_id` and `client_secret`, these parameters can be configured as shown hereafter:

* `ovh:endpoint` (environment variable: `OVH_ENDPOINT`)
* `ovh:clientId` (environment variable: `OVH_CLIENT_ID`)
* `ovh:clientSecret` (secret) (environment variable: `OVH_CLIENT_SECRET`)

#### Access Token

The provider will look for the token either at `OVH_ACCESS_TOKEN` environment variable, or get it via `ovh:access_token` argument in the provider's stanza.

Similarly to OAuth2 method, the endpoint must be configured (either via `ovh:endpoint` argument, or with `OVH_ENDPOINT` environment variable).

#### Application Key/Application Secret

In order to retrieve this necessary information, please follow [First steps with the OVHcloud APIs](https://docs.ovh.com/gb/en/customer/first-steps-with-ovh-api/) tutorial.

Concretely, you have to generate these credentials via the [OVH token generation page](https://api.ovh.com/createToken/?GET=/*&POST=/*&PUT=/*&DELETE=/*) with the following rights:

* GET `/`
* POST `/*`
* PUT `/*`
* DELETE `/*`

Once you have retrieved your `application_key`, `application_secret` and `consumer_key`, these parameters can be configured as shown hereafter:

* `ovh:endpoint` (environment variable: `OVH_ENDPOINT`)
* `ovh:applicationKey` (environment variable: `OVH_APPLICATION_KEY`)
* `ovh:applicationSecret` (secret) (environment variable: `OVH_APPLICATION_SECRET`)
* `ovh:consumerKey` (environment variable: `OVH_CONSUMER_KEY`)

### Endpoints

Depending on the API you want to use, you may set the `endpoint` to:

* `ovh-eu` for OVHcloud Europe API
* `ovh-us` for OVHcloud US API
* `ovh-ca` for OVHcloud Canada API
* `soyoustart-eu` for So you Start Europe API
* `soyoustart-ca` for So you Start Canada API
* `kimsufi-eu` for Kimsufi Europe API
* `kimsufi-ca` for Kimsufi Canada API
* Or any arbitrary URL to use in a test for example
