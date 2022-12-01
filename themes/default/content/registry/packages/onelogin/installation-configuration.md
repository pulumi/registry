---
title: onelogin Installation & Configuration
meta_desc: Provides an overview on how to configure credentials for the Pulumi onelogin Provider.
layout: installation
---

## Installation

The OneLogin provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/onelogin`](https://www.npmjs.com/package/@pulumi/onelogin)
* Python: [`pulumi-onelogin`](https://pypi.org/project/pulumi-onelogin/)
* Go: [`github.com/pulumi/pulumi-onelogin/sdk/go/onelogin`](https://github.com/pulumi/pulumi-onelogin)
* .NET: [`Pulumi.Onelogin`](https://www.nuget.org/packages/Pulumi.Onelogin)

## Configuring Credentials

The Pulumi OneLogin Provider needs to be configured with OneLogin credentials
before it can be used to create resources.

You'll need admin access to a OneLogin account where you can create API credentials.
Create a set of API credentials with manage all permission.
The following environment variables used to configure the OneLogin provider:

* `onelogin:clientId` - your client id. It can also be sourced from the environment variable `ONELOGIN_CLIENT_ID`.
* `onelogin:clientSecret` - your client secret. It can also be sourced from the environment variable `ONELOGIN_CLIENT_SECRET`.
* `onelogin:region` - the region.
* `onelogin:url` - the api url for your region. It can also be sourced from the environment variable `ONELOGIN_OAPI_URL`.

Remember to pass `--secret` when setting `onelogin:clientSecret` so that it is properly encrypted. The complete list of
configuration parameters is in the
[OneLogin provider README](https://github.com/pulumi/pulumi-onelogin/blob/master/README.md).
