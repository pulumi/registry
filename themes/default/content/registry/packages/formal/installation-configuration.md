---
# WARNING: this file was fetched from https://raw.githubusercontent.com/formalco/pulumi-formal/v1.0.2/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Formal Installation & Configuration
meta_desc: Provides an overview on how to configure credentials for the Formal provider for Pulumi.
layout: package
---
## Installation

The Formal provider is available as a package in the following Pulumi languages:

* JavaScript/TypeScript: [`@formalco/pulumi`](https://www.npmjs.com/package/@formalco/pulumi)
* Python: [`pulumi-formal`](https://pypi.org/project/pulumi-formal/)
* Go: [`github.com/formalco/pulumi-formal/sdk/go/formal`](https://pkg.go.dev/github.com/formalco/pulumi-formal/sdk/go/formal)

## Authentication

The Formal provider must be configured with an `API Key` in order to deploy Formal resources. See the Formal documentation on API keys [here.](https://docs.joinformal.com/tools/api-keys)

## Example configuration

Configure your Formal API key (with `--secret`):
```
pulumi config set formal:apiKey FORMAL_API_KEY --secret
```

You should now be able to deploy Formal resources.

## Configuration options

The following configuration options are available for the Formal provider:

- `formal:apiKey` (environment: `FORMAL_API_KEY`) - The API key used to access the formal control plane.