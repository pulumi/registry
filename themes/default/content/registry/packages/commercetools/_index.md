---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/labd/commercetools/1.21.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Commercetools Provider
meta_desc: Provides an overview on how to configure the Pulumi Commercetools provider.
layout: package
---

## Generate Provider

The Commercetools provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider labd/commercetools
```
## Commercial support
Need support implementing this pulumi module in your organization? We are
able to offer support. Please contact us at
opensource@labdigital.nl!
## Installation
Pulumi automatically downloads providers from the pulumi registry. Add the
following to your pulumi project

Packages of the releases are available at the GitHub Repo.
See the pulumi documentation
for more information about installing third-party providers.
## Using the provider
The provider attempts to read the required values from environment variables:
- `CTP_CLIENT_ID`
- `CTP_CLIENT_SECRET`
- `CTP_PROJECT_KEY`
- `CTP_SCOPES`
- `CTP_API_URL`
- `CTP_AUTH_URL`

Alternatively, you can set it up directly in the pulumi file:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    commercetools:apiUrl:
        value: <api url>
    commercetools:clientId:
        value: <your client id>
    commercetools:clientSecret:
        value: <your client secret>
    commercetools:projectKey:
        value: <your project key>
    commercetools:scopes:
        value: <space seperated list of scopes>
    commercetools:tokenUrl:
        value: <token url>

```
## Configuration Reference

- `apiUrl` (String) The API URL of the commercetools platform. <https://docs.commercetools.com/api/general-concepts#hosts>
- `clientId` (String, Sensitive) The OAuth Client ID for a commercetools platform project. <https://docs.commercetools.com/api/authorization>
- `clientSecret` (String, Sensitive) The OAuth Client Secret for a commercetools platform project. <https://docs.commercetools.com/api/authorization>
- `projectKey` (String, Sensitive) The project key of commercetools platform project. <https://docs.commercetools.com/getting-started>
- `scopes` (String) A list as string of OAuth scopes assigned to a project key, to access resources in a commercetools platform project. <https://docs.commercetools.com/api/authorization>
- `tokenUrl` (String) The authentication URL of the commercetools platform. <https://docs.commercetools.com/api/authorization>
## Using with docker

The included `Dockerfile` bundles the official  `pulumi/pulumi:light` docker image with
our `pulumi-provider-commercetools`.

To build the docker image file locally, use:
```sh
docker build . -t pulumi-with-provider-commercetools:latest
```
Then you can run a pulumi command on files in the current directory with:
```sh
docker run -v "${pwd}:/config" pulumi-with-provider-commercetools:latest <CMD>
```
## Authors
This project is developed by [Lab Digital](https://www.labdigital.nl). We
welcome additional contributors. Please see our
GitHub repository
for more information.