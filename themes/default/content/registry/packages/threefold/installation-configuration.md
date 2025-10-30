---
# WARNING: this file was fetched from https://raw.githubusercontent.com/threefoldtech/pulumi-threefold/v0.8.2/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/threefoldtech/pulumi-threefold/blob/v0.8.2/docs/installation-configuration.md
title: Threefold Grid Installation & Configuration
meta_desc: Information on how to install the Threefold provider.
layout: package
---

## Installation

The Pulumi Threefold provider is available as a package in those Pulumi languages:

* Go: [`github.com/threefoldtech/pulumi-threefold/sdk`](https://pkg.go.dev/github.com/threefoldtech/pulumi-threefold/sdk)
* JavaScript/TypeScript: [`@threefold/pulumi`](https://www.npmjs.com/package/@threefold/pulumi)
* Python: [`pulumi-threefold`](https://pypi.org/project/pulumi-threefold/)

## Setup

To provision resources with the Pulumi threefold provider, you need to provide the `mnemonic`.

## Configuration Options

Use `pulumi config set threefold:<option>`.

The following configuration points are available for the `threefold` provider:

* `threefold:mnemonic` (environment: `MNEMONIC`) -  This is the grid mnemonic. You can create a new account if you don't have [mnemonic](https://manual.grid.tf/documentation/dashboard/wallet_connector.html).

* `threefold:network` (environment: `NETWORK`) - specify which grid network (dev/qa/mainnet/testnet) to deploy on (default is dev).
