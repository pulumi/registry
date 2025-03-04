---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/dmacvicar/libvirt/0.8.3/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Libvirt Provider
meta_desc: Provides an overview on how to configure the Pulumi Libvirt provider.
layout: package
---

## Generate Provider

The Libvirt provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider dmacvicar/libvirt
```
## Overview

The Libvirt provider is used to interact with Linux
[libvirt](https://libvirt.org) hypervisors.

The provider needs to be configured with the proper connection information
before it can be used.

> **Note:** while libvirt can be used with several types of hypervisors, this
provider focuses on [KVM](http://libvirt.org/drvqemu.html). Other drivers may not be
working and haven't been tested.
## The connection URI

The provider understands [connection URIs](https://libvirt.org/uri.html). The supported transports are:

* `tcp` (non-encrypted connection)
* `unix` (UNIX domain socket)
* `tls` (See [here](https://libvirt.org/kbase/tlscerts.html) for information how to setup certificates)
* `ssh` (Secure shell)

Unlike the original libvirt, the `ssh` transport is not implemented using the ssh command and therefore does not require `nc` (netcat) on the server side.

Additionally, the `ssh` URI supports passwords using the `driver+ssh://[username:PASSWORD@][hostname][:port]/[path]?sshauth=ssh-password` syntax.

As the provider does not use libvirt on the client side, not all connection URI options are supported or apply.
## Configuration Reference

The following keys can be used to configure the provider.

* `uri` - (Required) The [connection URI](https://libvirt.org/uri.html) used
  to connect to the libvirt host.
## Environment variables

The libvirt connection URI can also be specified with the `LIBVIRT_DEFAULT_URI`
shell environment variable.