---
title: Google Cloud (GCP) Classic Installation & Configuration
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

{{% configure-gcp type="gcp" %}}

### Using a Service Account

If you are using Pulumi in a non-interactive setting (such as a CI/CD system) you will need to [configure and use a service account]({{< relref "service-account" >}}) instead.

## Configuration

There are a few different ways you can configure GCP credentials to work with Pulumi.

### Set configuration via pulumi config

You can set any configuration option in your Pulumi.yaml, for example:

```bash
$ pulumi config set gcp:project <your-gcp-project-id> # e.g. shinycorp-prod
$ pulumi config set gcp:region <your-region> # e.g us-west1
$ pulumi config set gcp:region <your-region> # e.g us-west1-a
```

### Set configuration via environment variables

We recommend using `pulumi config` for the options below, but you can also set some of them as environment variables instead.
For example:

* `GOOGLE_PROJECT` - The default project for new resources, if one is not specified when creating a resource
* `GOOGLE_REGION` - The default region for new resources, if one is not specified when creating a resource
* `GOOGLE_ZONE` - The default zone for new resources, if one is not specified when creating a resource.

## Configuration reference

Use `pulumi config set gcp:<option>` or pass options to the [constructor of `new gcp.Provider`]({{< relref "/registry/packages/gcp/api-docs/provider" >}}).

| Option                               | Required/Optional | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
|--------------------------------------|-------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `project`                            | Required          | The ID of the project to apply any resources to. This can also be specified using any of the following environment variables (listed in order of precedence): `GOOGLE_PROJECT`, `GOOGLE_CLOUD_PROJECT`, `GCLOUD_PROJECT`, `CLOUDSDK_CORE_PROJECT`.                                                                                                                                                                                                                                                                                                                                                                                                                          |
| `region`                             | Optional          | The region to operate under, if not specified by a given resource. This can also be specified using any of the following environment variables (listed in order of precedence): `GOOGLE_REGION`, `GCLOUD_REGION`, `CLOUDSDK_COMPUTE_REGION`.                                                                                                                                                                                                                                                                                                                                                                                                                                |
| `zone`                               | Optional          | The zone to operate under, if not specified by a given resource.  This can also be specified using any of the following environment variables (listed in order of precedence): `GOOGLE_ZONE`, `GCLOUD_ZONE`, `CLOUDSDK_COMPUTE_ZONE`.                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| `credentials`                        | Optional          | Contents of a file (or path to a file) that contains your [service account private key in JSON format]({{< relref "service-account" >}}). Credentials can also be specified using any of the following environment variables (listed in order of precedence): `GOOGLE_CREDENTIALS`, `GOOGLE_CLOUD_KEYFILE_JSON`, `GCLOUD_KEYFILE_JSON`. If no credentials are specified, the provider will fall back to using the Google Application Default Credentials. If you are running Pulumi from a GCE instance, see [Creating and Enabling Service Accounts for Instances](https://cloud.google.com/compute/docs/access/create-enable-service-accounts-for-instances) for details. |
| `accessToken`                        | Optional          | A temporary OAuth 2.0 access token obtained from the Google Authorization server, i.e. the `Authorization: Bearer` token used to authenticate HTTP requests to GCP APIs. Alternative to `credentials`. Ignores the `scopes` field.                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| `scopes`                             | Optional          | List of OAuth 2.0 [scopes](https://developers.google.com/identity/protocols/oauth2/scopes) requested when generating an access token using the service account key specified in `credentials`. Defaults: `https://www.googleapis.com/auth/cloud-platform` and `https://www.googleapis.com/auth/userinfo.email`                                                                                                                                                                                                                                                                                                                                                              |
| `impersonateServiceAccount`          | Optional          | Setting to impersonate a [Google service account](https://cloud.google.com/iam/docs/create-short-lived-credentials-direct) If you authenticate as a service account, Google Cloud derives your quota project and permissions from that service account rather than your primary authentication method. A valid primary authentication mechanism must be provided for the impersonation call, and your primary identity must have the `roles/iam.serviceAccountTokenCreator` role on the service account you are impersonating. This can also be specified by setting the `GOOGLE_IMPERSONATE_SERVICE_ACCOUNT` environment variable.                                         |
