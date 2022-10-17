---
title: Google Native Setup
meta_desc: How to set up credentials to use the Pulumi Google Native Provider and choose configuration options to tailor the provider to suit your use case.
layout: installation
---

## Installation

The Google Native provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/google-native`](https://www.npmjs.com/package/@pulumi/google-native)
* Python: [`pulumi-google-native`](https://pypi.org/project/pulumi-google-native/)
* Go: [`github.com/pulumi/pulumi-google-native/sdk/go/google`](https://github.com/pulumi/pulumi-google-native)
* .NET: [`Pulumi.GoogleNative`](https://www.nuget.org/packages/Pulumi.GoogleNative)

## Configuration

To provision resources with the Pulumi Google Cloud Provider, you need to have Google credentials.

{{% configure-gcp type="google-native" %}}

{{% notes "info" %}}
If you are using Pulumi in an non-interactive setting (such as a CI/CD system) you will need to [configure and use a service account]({{< relref "service-account" >}}) instead.
{{% /notes %}}

### Configuration Options

Use `pulumi config set google-native:<option>` or pass options to the [constructor of `new Provider`]({{< relref "/registry/packages/google-native/api-docs/provider" >}}).

| Option        | Required/Optional | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
|---------------|-------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `project`     | Optional          | The default project for new resources, if one is not specified when creating a resource. This can also be specified using any of the following environment variables (listed in order of precedence): `GOOGLE_PROJECT`, `GOOGLE_CLOUD_PROJECT`, `GCLOUD_PROJECT`, `CLOUDSDK_CORE_PROJECT`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| `region`      | Optional          | The region to operate under, if not specified by a given resource. This can also be specified using any of the following environment variables (listed in order of precedence): `GOOGLE_REGION`, `GCLOUD_REGION`, `CLOUDSDK_COMPUTE_REGION`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| `zone`        | Optional          | The zone to operate under, if not specified by a given resource.  This can also be specified using any of the following environment variables (listed in order of precedence): `GOOGLE_ZONE`, `GCLOUD_ZONE`, `CLOUDSDK_COMPUTE_ZONE`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |

### Use environment variables

We recommend using `pulumi config` for the options above, but you can also set many of the options above as environment variables instead.

* `GOOGLE_PROJECT` - The default project for new resources, if one is not specified when creating a resource
* `GOOGLE_REGION` - The default region for new resources, if one is not specified when creating a resource
* `GOOGLE_ZONE` - The default zone for new resources, if one is not specified when creating a resource.
