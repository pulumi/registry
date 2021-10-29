---
title: vSphere Setup
meta_desc: This page provides an overview on how to setup use Pulumi and the vSphere SDK to
           manage and provision resources.
layout: installation
---

The Pulumi vSphere provider uses the vSphere SDK to manage and provision resources.
Pulumi relies on the vSphere SDK to authenticate requests from your computer to vSphere. Your credentials are never sent
to pulumi.com.
The Pulumi vSphere Provider needs to be configured with vSphere credentials
before it can be used to create resources.

## Installation

The vSphere provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/vsphere`](https://www.npmjs.com/package/@pulumi/vsphere)
* Python: [`pulumi-vsphere`](https://pypi.org/project/pulumi-vsphere/)
* Go: [`github.com/pulumi/pulumi-vsphere/sdk/v3/go/vsphere`](https://github.com/pulumi/pulumi-vsphere)
* .NET: [`Pulumi.Vsphere`](https://www.nuget.org/packages/Pulumi.Vsphere)

## Configuring Credentials

Once obtained, there are two ways to communicate your authorization tokens to Pulumi:

1. Set the environment variables `VSPHERE_USER`, `VSPHERE_PASSWORD` and `VSPHERE_SERVER`:

    ```bash
    $ export VSPHERE_USER=XXXXXXXXXXXX
    $ export VSPHERE_PASSWORD=YYYYYYYYYYYY
    $ export VSPHERE_SERVER=ZZZZZZZZZZZZ
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set vsphere:user XXXXXXXXXXXX
    $ pulumi config set vsphere:password YYYYYYYYYYYY --secret
    $ pulumi config set vsphere:vsphereServer ZZZZZZZZZZZZ
    ```

Remember to pass `--secret` when setting `password` so that it is properly encrypted.
