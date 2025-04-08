---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumiverse/pulumi-heroku/v1.0.4/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Heroku Provider
meta_desc: Provides an overview on how to configure the Pulumi Heroku provider.
layout: package
---

## Installation

The Heroku provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/heroku`](https://www.npmjs.com/package/@pulumi/heroku)
* Python: [`pulumi-heroku`](https://pypi.org/project/pulumi-heroku/)
* Go: [`github.com/pulumiverse/pulumi-heroku/sdk/go/heroku`](https://github.com/pulumi/pulumi-heroku)
* .NET: [`Pulumi.Heroku`](https://www.nuget.org/packages/Pulumi.Heroku)

## Overview

The Heroku provider is used to interact with the resources provided by Heroku
Platform API and needs to be configured with credentials before it can be used.
## Background

[Heroku](https://www.heroku.com) is a fully-managed platform that gives you the
simplest path to delivering apps quickly:

* Using Pulumi with Heroku
* [Platform API reference](https://devcenter.heroku.com/articles/platform-api-reference)
* [Command Line Interface (CLI)](https://devcenter.heroku.com/articles/heroku-cli)
## Guides

* Upgrading
* Secure Practices
## Example Usage

```hcl-pulumi
# Configure the Heroku provider
provider "heroku" {
  api_key = var.heroku_api_key
}

# Create a new application
resource "heroku_app" "default" {
  # ...
}
```
## Authentication

The Heroku provider offers a flexible means of providing credentials for
authentication. The following methods are supported, listed in order of
precedence, and explained below:

* Static credentials
* Environment variables
* Netrc
### Generating tokens

All authentication tokens must be generated with one of these methods:

* [Heroku Dashboard](https://dashboard.heroku.com) → Account Settings → Applications → [Authorizations](https://dashboard.heroku.com/account/applications)
* `heroku auth` command of the [Heroku CLI](https://devcenter.heroku.com/articles/heroku-cli)
* [Heroku Platform APIs: OAuth](https://devcenter.heroku.com/articles/platform-api-reference#oauth-authorization)

🔐  See Secure Practices for help creating a safe API token.

⛔️  Direct username-password authentication is [no longer supported by Heroku API](https://devcenter.heroku.com/changelog-items/2516).
### Static credentials

Credentials can be provided statically by adding `apiKey` property
to the Heroku provider configuration:

```hcl-pulumi
variable "heroku_api_key" {
  type      = string
  sensitive = true
}

provider "heroku" {
  api_key = var.heroku_api_key
}
```
### Environment variables

When the Heroku provider configuration does not contain an `apiKey`
argument, the missing credentials will be sourced from the environment via the
`HEROKU_API_KEY` environment variable:

```hcl-pulumi
provider "heroku" {}
```

```shell
$ export HEROKU_API_KEY="<heroku_auth_token>"
$ pulumi preview
Refreshing Pulumi state in-memory prior to plan...
```
### Netrc

Credentials can instead be sourced from the [`.netrc`](https://ec.haxx.se/usingcurl-netrc.html)
file in your home directory:

```hcl-pulumi
provider "heroku" {}
```

```shell
$ cat ~/.netrc
...
machine api.heroku.com
  login <ignored, can be any value>
  password <heroku_auth_token>
...
```

The directory containing the `.netrc` file can be overridden by the `NETRC` environment variable as described [here](https://www.gnu.org/software/inetutils/manual/html_node/The-_002enetrc-file.html).
## Configuration Reference

The following configuration inputs are supported:

* `apiKey` - (Required) Heroku API token. It must be provided, but it can also
  be sourced from other locations. See also Secure Practices.

* `email` - (Ignored) This field originally supported username-password authentication,
  but has since [been deprecated](https://devcenter.heroku.com/changelog-items/2516).
  Instead, simply set an auth token in the `apiKey` property.

* `headers` - (Optional) Additional Headers to be sent to Heroku, as a string-encoded JSON object,
  for example: `{"X-Custom-Header":"yes","X-Custom-Header-Too":"no"}`. If not provided, it will be
  sourced from the `HEROKU_HEADERS` environment variable (if set).

* `customizations` - (Optional) Various attributes altering the behavior of certain resources.
  Only a single `customizations` block may be specified, and it supports the following arguments:

  * `setAppAllConfigVarsInState` - (Optional) Controls whether the `heroku_app.all_config_vars` attribute
    is set in the state file. Normally a snapshot of all config vars is stored in state, even though they are
    not managed by Pulumi, such as secrets set via `heroku config` CLI, web dashboard, or add-ons like
    Postgres' `DATABASE_URL`. Set to `false` to only track managed config vars in the state. Defaults to `true`.
    See also Secure Practices.

  * `setAddonConfigVarsInState` - (Optional) Controls whether the `heroku_addon.config_var_values` attribute
    is set in the state file. The attribute stores each addon's config vars in Pulumi state. This means
    sensitive add-on config vars, such as Postgres' `DATABASE_URL`, are always accessible in the state.
    Set to `false` to prevent capturing these values. Defaults to `true`.
    See also Secure Practices.

* `delays` - (Optional) Delays help mitigate issues that can arise due to
  Heroku's eventually consistent data model. Only a single `delays` block may be
  specified, and it supports the following arguments:

  * `postAppCreateDelay` - (Optional) The number of seconds to wait after an
    app is created. Default is to wait 5 seconds.

  * `postSpaceCreateDelay` - (Optional) The number of seconds to wait after a
    private space is created. Default is to wait 5 seconds.

  * `postDomainCreateDelay` - (Optional) The number of seconds to wait after
    a domain is created. Default is to wait 5 seconds.

* `timeouts` - (Optional) Define a max duration the provider will wait for certain resources
  to be properly modified before proceeding with further action(s). Only a single `timeouts` block may be specified,
  and it supports the following arguments:

  * `addonCreateTimeout` - (Optional) The number of minutes for the provider to wait for an addon to be
    created/provisioned. Defaults to 20 minutes. Minimum required value is 10 minutes.
