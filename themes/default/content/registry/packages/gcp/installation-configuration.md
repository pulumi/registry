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

## Configuration

To provision resources with the Pulumi Google Cloud Provider, you need to have Google credentials.

{{% configure-gcp %}}

{{% notes "info" %}}
If you are using Pulumi in an non-interactive setting (such as a CI/CD system) you will need to [configure and use a service account]({{< relref "service-account" >}}) instead.
{{% /notes %}}

### Configuration Options

Use `pulumi config set gcp:<option>` or pass options to the [constructor of `new gcp.Provider`]({{< relref "/registry/packages/gcp/api-docs/provider" >}}).

| Option        | Required/Optional | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
|---------------|-------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `project`     | Required          | The ID of the project to apply any resources to. This can also be specified using any of the following environment variables (listed in order of precedence): `GOOGLE_PROJECT`, `GOOGLE_CLOUD_PROJECT`, `GCLOUD_PROJECT`, `CLOUDSDK_CORE_PROJECT`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| `credentials` | Optional          | Contents of a file that contains your service account private key in JSON format. You can download your existing Google Cloud service account file from the Google Cloud Console, or you can create a new one from the same page. Credentials can also be specified using any of the following environment variables (listed in order of precedence): `GOOGLE_CREDENTIALS`, `GOOGLE_CLOUD_KEYFILE_JSON`, `GCLOUD_KEYFILE_JSON`. The `GOOGLE_APPLICATION_CREDENTIALS` environment variable can also contain the path of a file to obtain credentials from. If no credentials are specified, the provider will fall back to using the Google Application Default Credentials. If you are running Pulumi from a GCE instance, see [Creating and Enabling Service Accounts for Instances](https://cloud.google.com/compute/docs/access/create-enable-service-accounts-for-instances) for details. |
| `region`      | Optional          | The region to operate under, if not specified by a given resource. This can also be specified using any of the following environment variables (listed in order of precedence): `GOOGLE_REGION`, `GCLOUD_REGION`, `CLOUDSDK_COMPUTE_REGION`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| `zone`        | Optional          | The zone to operate under, if not specified by a given resource.  This can also be specified using any of the following environment variables (listed in order of precedence): `GOOGLE_ZONE`, `GCLOUD_ZONE`, `CLOUDSDK_COMPUTE_ZONE`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |

### Use environment variables

We recommend using `pulumi config` for the options above, but you can also set many of the options above as environment variables instead.

* `GOOGLE_PROJECT` - The default project for new resources, if one is not specified when creating a resource
* `GOOGLE_REGION` - The default region for new resources, if one is not specified when creating a resource
* `GOOGLE_ZONE` - The default zone for new resources, if one is not specified when creating a resource.
