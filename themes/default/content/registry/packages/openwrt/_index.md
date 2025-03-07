---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/joneshf/openwrt/0.0.20/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Openwrt Provider
meta_desc: Provides an overview on how to configure the Pulumi Openwrt provider.
layout: package
---

## Generate Provider

The Openwrt provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider joneshf/openwrt
```
## Overview

Interfaces with an OpenWrt device through LuCI RPC. See <https://github.com/openwrt/luci/wiki/JsonRpcHowTo#basics> for setup instructions.
## Example Usage

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:

```
## Configuration Reference

- `hostname` (String) The hostname to use. Defaults to "192.168.1.1".
- `password` (String, Sensitive) The password to use. Defaults to "".
- `port` (Number) The port to use. Defaults to 80.
- `scheme` (String) The URI scheme to use. Defaults to "http".
- `username` (String) The username to use. Defaults to "root".