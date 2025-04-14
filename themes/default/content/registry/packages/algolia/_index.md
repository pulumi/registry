---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/k-yomo/algolia/0.6.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Algolia Provider
meta_desc: Provides an overview on how to configure the Pulumi Algolia provider.
layout: package
---

## Generate Provider

The Algolia provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider k-yomo/algolia
```
## Overview

The Algolia provider is used to configure your Algolia settings.
## Setting API Key
In order to make requests against the Algolia API, you need to set an API key.
You set the API key to Pulumi using the environment variable `ALGOLIA_API_KEY`.
```sh
$ export ALGOLIA_API_KEY={{my-api-key}}
```
## Example Usage
Then typical provider configuration will look something like:
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    algolia:appId:
        value: my-app-id

```
## Configuration Reference
- `apiKey` (String) The API key to access algolia resources. Defaults to the env variable `ALGOLIA_API_KEY`.
- `appId` (String) The ID of the application. Defaults to the env variable `ALGOLIA_APP_ID`.