---
title: Google Cloud (GCP) Classic Setup
meta_desc: How to set up credentials to use the Pulumi GCP Provider and choose configuration options to tailor the provider to suit your use case.
layout: installation
---

## Installation

The Google Cloud (GCP) Classic provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/gcp`](https://www.npmjs.com/package/@pulumi/gcp)
* Python: [`pulumi-gcp`](https://pypi.org/project/pulumi-gcp/)
* Go: [`github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp`](https://github.com/pulumi/pulumi-gcp)
* .NET: [`Pulumi.Gcp`](https://www.nuget.org/packages/Pulumi.Gcp)
* Java: [`com.pulumi.gcp`](https://search.maven.org/search?q=com.pulumi.gcp)

## Credentials

To provision resources with the Pulumi Google Cloud Provider, you need to have Google credentials.

### Default auth credentials

{{% configure-gcp %}}

### Using a Service Account

Using a Google service account allows you to use Pulumi in a non-interactive setting (for example CI/CD systems, where a person can not complete the normal `gcloud auth application-default login` flow). A service account can also be used when developing locally to ensure a specific set of scoped credentials not tied to a user account are used. This can be useful even when developing locally to give you more control over the account role used for deployment.

To use a service account with Pulumi you will need to provide the Google Cloud Platform Provider with your Google service account private key. You can
create and download credentials using the [Google Cloud Platform Credentials] https://console.cloud.google.com/apis/credentials

In order to create new credentials to use with Pulumi, go to the `APIs and Services` section of of the Google Cloud Platform Console
and select the `Credentials` sub-menu. From here, select the `Create credentials` drop-down menu and click `Service account key`
to create a new key for a service account.

![Create new credentials](/images/docs/gcp_configure/gcp_create_credentials.png)

On the next screen, select `JSON` as the key type and select the service account to which this key will be associated.

![Create new credentials](/images/docs/gcp_configure/gcp_create_service_account_key.png)

Pressing the `Create` button will download a JSON file. This file contains your
new credentials.

> Your credentials are only used to authenticate with Google Cloud APIs on your behalf. Your credentials are never sent to pulumi.com.

To communicate your credentials to the Pulumi Google Cloud Platform Provider,
export the contents of your credentials file to the `GOOGLE_CREDENTIALS`
environment variable:

Linux and Mac OS X

```bash
export GOOGLE_CREDENTIALS=$(cat credentials.json)
```

Windows Powershell

```bash
$env:GOOGLE_CREDENTIALS=cat credentials.json
```

## Configuration

There are a few different ways you can configure GCP credentials to work with Pulumi.

### Set configuration via pulumi config

You can set any configuration in your Pulumi.yaml, for example:

```bash
$ pulumi config set gcp:project <your-gcp-project-id> # e.g. shinycorp-prod
$ pulumi config set gcp:region <your-region> # e.g us-west1
$ pulumi config set gcp:region <your-region> # e.g us-west1-a
```

See a full config list below.

### Set configuration via environment variables

We recommend using `pulumi config` for the options below, but you can also set some of them as environment variables instead.
For example:

* `GOOGLE_PROJECT` - The default project for new resources, if one is not specified when creating a resource
* `GOOGLE_REGION` - The default region for new resources, if one is not specified when creating a resource
* `GOOGLE_ZONE` - The default zone for new resources, if one is not specified when creating a resource.

## Configuration reference

Use `pulumi config set gcp:<option>` or pass options to the [constructor of `new gcp.Provider`]({{< relref "/registry/packages/gcp/api-docs/provider" >}}).

| Option        | Required/Optional | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
|---------------|-------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `project`     | Required          | The ID of the project to apply any resources to. This can also be specified using any of the following environment variables (listed in order of precedence): `GOOGLE_PROJECT`, `GOOGLE_CLOUD_PROJECT`, `GCLOUD_PROJECT`, `CLOUDSDK_CORE_PROJECT`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| `credentials` | Optional          | Contents of a file that contains your service account private key in JSON format. You can download your existing Google Cloud service account file from the Google Cloud Console, or you can create a new one from the same page. Credentials can also be specified using any of the following environment variables (listed in order of precedence): `GOOGLE_CREDENTIALS`, `GOOGLE_CLOUD_KEYFILE_JSON`, `GCLOUD_KEYFILE_JSON`. The `GOOGLE_APPLICATION_CREDENTIALS` environment variable can also contain the path of a file to obtain credentials from. If no credentials are specified, the provider will fall back to using the Google Application Default Credentials. If you are running Pulumi from a GCE instance, see [Creating and Enabling Service Accounts for Instances](https://cloud.google.com/compute/docs/access/create-enable-service-accounts-for-instances) for details. |
| `region`      | Optional          | The region to operate under, if not specified by a given resource. This can also be specified using any of the following environment variables (listed in order of precedence): `GOOGLE_REGION`, `GCLOUD_REGION`, `CLOUDSDK_COMPUTE_REGION`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| `zone`        | Optional          | The zone to operate under, if not specified by a given resource.  This can also be specified using any of the following environment variables (listed in order of precedence): `GOOGLE_ZONE`, `GCLOUD_ZONE`, `CLOUDSDK_COMPUTE_ZONE`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |

