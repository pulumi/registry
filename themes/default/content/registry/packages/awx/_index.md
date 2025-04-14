---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/denouche/awx/0.29.1/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Awx Provider
meta_desc: Provides an overview on how to configure the Pulumi Awx provider.
layout: package
---

## Generate Provider

The Awx provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider denouche/awx
```
## Overview

Ansible Tower Provider for handle Tower Projects with [rest](https://docs.ansible.com/ansible-tower/latest/html/towerapi/api_ref.html)
## Example Usage

Using username and password:
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    awx:hostname:
        value: http://localhost:8078
    awx:password:
        value: changeme
    awx:username:
        value: test

```

Using token:
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    awx:hostname:
        value: http://localhost:8078
    awx:token:
        value: awxtoken

```

> ⚠️ Be careful, if you set both token and username/password the token will have the precedence.
## Configuration Reference

The following configuration inputs are supported:

* `hostname` - (Optional) The API endpoint for AWX. Defaults to `"http://localhost"`.
* `username` - (Optional) The username for API access. Defaults to `"admin"`.
* `password` - (Optional) The password for API access. Defaults to `"password"`.
* `token`    - (Optional) The AWX token for API access. Defaults to empty.
* `insecure` - (Optional) Whether to check the TLS certificate. Defaults to `false`.