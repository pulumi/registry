---
# WARNING: this file was fetched from https://raw.githubusercontent.com/ionos-cloud/pulumi-ionoscloud/v0.2.1/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Ionoscloud Installation & Configuration
meta_desc: Information on how to install the Ionoscloud provider.
layout: installation
---

## Installation

The Pulumi `Ionoscloud` provider is available as a package in most Pulumi languages:

* JavaScript/TypeScript: [`@ovhcloud/pulumi-ovh`](https://www.npmjs.com/package/@ionos-cloud/sdk-pulumi)
* Python: [`pulumi_ionoscloud`](https://pypi.org/project/pulumi-ionoscloud)
* Go: [`github.com/ionosc-cloud/pulumi-ionoscloud/sdk/go/ionoscloud`](https://pkg.go.dev/github.com/ionos-cloud/pulumi-ionoscloud/sdk)
* .NET: [`Ionoscloud.Pulumi.Ionoscloud`](https://www.nuget.org/packages/Ionoscloud.pulumi.ionoscloud/)

### Provider Binary

The Ionoscloud provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource ionoscloud vX.Y.Z --server github://api.github.com/ionos-cloud/pulumi-ionoscloud
```

Replace the version string with your desired version.

## Setup

To provision resources with the Pulumi Ionoscloud provider, you need to have Ionoscloud credentials.
Your Ionoscloud credentials are never sent to pulumi.com. Pulumi uses the Ionoscloud API and the credentials in your environment to authenticate requests from your computer to Ionoscloud.

The provider needs to be configured with the proper credentials before it can be used. Requests to the Ionoscloud APIs require generating a token. See [Token Management](https://docs.ionos.com/cloud/set-up-ionos-cloud/management/token-management) for a detailed explanation.

## Environment Variables

| Environment Variable    | Description                                                                                                                                                              |
|-------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `IONOS_USERNAME`        | Specify the username used to login, to authenticate against the IONOS Cloud API                                                                                          |
| `IONOS_PASSWORD`        | Specify the password used to login, to authenticate against the IONOS Cloud API                                                                                          |
| `IONOS_TOKEN`           | Specify the token used to login, if a token is being used instead of username and password                                                                               |
| `IONOS_API_URL`         | Specify the API URL. It will overwrite the API endpoint default value `api.ionos.com`. It is not necessary to override this value unless you have special routing config |
| `IONOS_LOG_LEVEL`       | Specify the Log Level used to log messages. Possible values: Off, Debug, Trace                                                                                           |
| `IONOS_PINNED_CERT`     | Specify the SHA-256 public fingerprint here, enables certificate pinning                                                                                                 |
| `IONOS_CONTRACT_NUMBER` | Specify the contract number on which you wish to provision. Only valid for reseller accounts, for other types of accounts the header will be ignored                     |
| `IONOS_S3_ACCESS_KEY`   | Specify the access key used to authenticate against the IONOS Object Storage API                                                                                         |
| `IONOS_S3_SECRET_KEY`   | Specify the secret key used to authenticate against the IONOS Object Storage API                                                                                         |
| `IONOS_S3_REGION`       | Region for IONOS Object Storage operations. Default value: eu-central-3. **If you use IONOS_API_URL_OBJECT_STORAGE, `IONOS_S3_REGION` is mandatory**    
