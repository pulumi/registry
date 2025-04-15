---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/rollbar/rollbar/1.15.1/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Rollbar Provider
meta_desc: Provides an overview on how to configure the Pulumi Rollbar provider.
layout: package
---

## Generate Provider

The Rollbar provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider rollbar/rollbar
```
## Overview

The Rollbar provider is used to interact with [Rollbar](https://rollbar.com)
resources.

The provider allows you to manage your Rollbar account's projects, members, and
teams easily. It needs to be configured with the proper credentials before it
can be used.
## Example Usage

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    rollbar:apiKey:
        value: 'TODO: var.rollbar_token'

```
## Configuration Reference

The following configuration inputs are supported:

* `apiKey` - (Required) Rollbar API authentication token. Value will be
  sourced from environment variable `ROLLBAR_API_KEY` if set.
* `projectApiKey` - (Optional) Rollbar API authentication token (project level).
  Value will be sourced from environment variable `ROLLBAR_PROJECT_API_KEY` if set.
* `apiUrl` - (Optional) Base URL for the Rollbar API.  Defaults to
  <https://api.rollbar.com>.  Value will be sourced from environment variable
  `ROLLBAR_API_URL` if set.
## Functions

* `rollbar.Project` - A Rollbar project
* `rollbar.getProjects` - List all Rollbar
  projects
* `rollbar.ProjectAccessToken`
  - An access token belonging to a Rollbar project
* `rollbar.getProjectAccessTokens`
  - List all access tokens belonging to a Rollbar project
* `rollbar.Team` - A Rollbar team
## Resources

* `rollbar.Project` - A Rollbar project
* `rollbar.ProjectAccessToken` - A
  Rollbar project access token
* `rollbar.Notification` - A Rollbar notification
  channel rule
* `rollbar.Team` - A Rollbar team
* `rollbar.User` - A Rollbar user