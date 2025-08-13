---
# WARNING: this file was fetched from https://raw.githubusercontent.com/formalco/pulumi-formal/v1.0.1/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Formal Provider
meta_desc: Provides an overview on how to configure the Pulumi Formal provider.
layout: package
---

## Installation

The Formal provider is available as a package in the following Pulumi languages:

* JavaScript/TypeScript: [`@formalco/pulumi`](https://www.npmjs.com/package/@formalco/pulumi)
* Python: [`pulumi-formal`](https://pypi.org/project/pulumi-formal/)
* Go: [`github.com/formalco/pulumi-formal/sdk/go/formal`](https://pkg.go.dev/github.com/formalco/pulumi-formal/sdk/go/formal)

## Overview

Use the Formal Pulumi Provider to interact with the
many resources supported by Formal.

Use the navigation to the left to read about the available resources.

## Authentication and Configuration

Configuration for the Formal Provider is derived from the API tokens you can generate via the Formal Console.

### Provider Configuration

!> **Warning:** Hard-coded credentials are not recommended in any Pulumi
configuration and risks secret leakage should this file ever be committed to a
public version control system.

Credentials can be provided by adding an `apiKey`.

Usage:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    formal:apiKey:
        value: '<apiKey>'
```

Credentials can also be provided by using the `FORMAL_API_KEY` environment variable.

For example:

Usage:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:

```

```bash
export FORMAL_API_KEY="some_api_key" pulumi up
```

## Examples

See various examples on how to deploy Formal resources in the [`examples/`](https://github.com/formalco/pulumi-formal/tree/main/examples) folder.