---
title: onelogin Setup
meta_desc: This page provides an overview on how to configure credentials for the Pulumi onelogin Provider.
layout: installation
---

The Pulumi onelogin Provider needs to be configured with onelogin credentials
before it can be used to create resources.

### Configuring Credentials

You'll need admin access to a onelogin account where you can create API credentials.
Create a set of API credentials with manage all permission.
The following environment variables used to configure the onelogin provider:

* `onelogin:clientId` - your client id. It can also be sourced from the environment variable `ONELOGIN_CLIENT_ID`.
* `onelogin:clientSecret` - your client secret. It can also be sourced from the environment variable `ONELOGIN_CLIENT_SECRET`.
* `onelogin:region` - the region.
* `onelogin:url` - the api url for your region. It can also be sourced from the environment variable `ONELOGIN_OAPI_URL`.

Remember to pass `--secret` when setting `onelogin:clientSecret` so that it is properly encrypted. A
full set of configuration parameters can be found listed on the
[Project README](https://github.com/pulumi/pulumi-onelogin/blob/master/README.md).
