---
title: Google Cloud Native Installation & Configuration
meta_desc: How to set up credentials to use the Pulumi Google Cloud Native Provider and choose configuration options to tailor the provider to suit your use case.
layout: package
---

## Installation

The Google Cloud Native provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/google-native`](https://www.npmjs.com/package/@pulumi/google-native)
* Python: [`pulumi-google-native`](https://pypi.org/project/pulumi-google-native/)
* Go: [`github.com/pulumi/pulumi-google-native/sdk/go/google`](https://github.com/pulumi/pulumi-google-native)
* .NET: [`Pulumi.GoogleNative`](https://www.nuget.org/packages/Pulumi.GoogleNative)

## Authentication Methods

To provision resources with the Pulumi Google Cloud Native Provider, you need to have Google credentials.

Pulumi can authenticate to Google Cloud via several methods:

- Google Cloud CLI
- OpenID Connect (OIDC)
- Service account

## Configuration

There are a few different ways you can configure your Google Cloud credentials to work with Pulumi.

### Authenticate using the CLI

{{% configure-gcp type="google-native" %}}

### Authenticate using a service account

If you are using Pulumi in an non-interactive setting (such as a CI/CD system) you will need to [configure and use a service account](/registry/packages/google-native/api-docs/iam/v1/serviceaccount/) instead.

### Authenticate with dynamically generated credentials

In addition to configuring the GCP provider locally, you also have the option to centralize your configurations using [Pulumi ESC (Environments, Secrets, and Configuration)](/docs/pulumi-cloud/esc/). Using this service will enable you to run Pulumi CLI commands with dynamically generated credentials, removing the need to configure and manage your credentials locally.

To do this, you will need to complete the following steps:

#### Configure OIDC between Pulumi and GCP

Refer to the [Configuring OpenID Connect for GCP Guide](/docs/pulumi-cloud/oidc/gcp/) for the step-by-step process on how to do this. Once you have completed these steps, you can define and expose environment variables as shown below:

```yaml
values:
  gcp:
    login:
      fn::open::gcp-login:
        project: <your-project-id>
        oidc:
          workloadPoolId: <your-pool-id>
          providerId: <your-provider-id>
          serviceAccount: <your-service-account>
  pulumiConfig:
    gcp:accessToken: ${gcp.login.accessToken}
  environmentVariables:
    GOOGLE_PROJECT: ${gcp.login.project}
    GOOGLE_REGION: <your-region>
    GOOGLE_ZONE: <your-zone>
```

{{< notes type="warning" >}}
Your GCP access token must always be defined under the `pulumiConfig` section. The deployment will fail if it is defined as an environment variable in the `environmentVariables` section.
{{< /notes >}}

To [expose these values to Pulumi IaC](/docs/pulumi-cloud/esc/environments/#using-environments-with-pulumi-iac), you will need to add any desired values underneath the `pulumiConfig` key. Further, if your workflow does not require the exposure of environment variables, you can also define those variables under the `pulumiConfig` block as shown below:

```yaml
values:
  gcp:
    login:
      fn::open::gcp-login:
        project: <your-project-id>
        oidc:
          workloadPoolId: <your-pool-id>
          providerId: <your-provider-id>
          serviceAccount: <your-service-account>
  pulumiConfig:
    project:environment: 'dev'
    gcp:accessToken: ${gcp.login.accessToken}
    gcp:project: ${gcp.login.project}
    gcp:region: <your-region>
    gcp:zone: <your-zone>
```

This will ensure that those values are scoped only to your `pulumi` run.

{{< notes type="info" >}}
The configuration values under `pulumiConfig` can also be referenced directly from within your Pulumi program code. This is done using the same method to reference values from your project's stack settings file. You can see examples of how to do this in the [Accessing Configuration from Code](/docs/concepts/config/#code) section of the Pulumi documentation.
{{< /notes >}}

#### Import your environment

The last step is to update your project's stack settings file (`Pulumi.<stack-name>.yaml`) to import your ESC environment as shown below:

```yaml
environment:
  - <your-environment-name>
```

Make sure to replace `<your-environment-name>` with the name of the ESC environment you created in the previous steps.

You can test that your configuration is working by running the `pulumi preview` command. This will validate that your GCP resources can be deployed using the dynamically generated credentials in your environment file.

{{< notes type="info" >}}
Make sure that your local environment does not have GCP credentials configured before running this command. You can logout by running the `gcloud auth revoke` command.
{{< /notes >}}

To learn more about projecting environment variables in Pulumi ESC, refer to the [relevant Pulumi ESC documentation](/docs/pulumi-cloud/esc/environments/#projecting-environment-variables).

### Configuration options

Use `pulumi config set google-native:<option>` or pass options to the [constructor of `new Provider`](/registry/packages/google-native/api-docs/provider).

| Option  Required? | Description |
| - | - | - |
| `project` | Optional | The default project for new resources, if one is not specified when creating a resource. This can also be specified using any of the following environment variables (listed in order of precedence): `GOOGLE_PROJECT`, `GOOGLE_CLOUD_PROJECT`, `GCLOUD_PROJECT`, `CLOUDSDK_CORE_PROJECT`. |
| `region` | Optional | The region to operate under, if not specified by a given resource. This can also be specified using any of the following environment variables (listed in order of precedence): `GOOGLE_REGION`, `GCLOUD_REGION`, `CLOUDSDK_COMPUTE_REGION`. |
| `zone` | Optional | The zone to operate under, if not specified by a given resource.  This can also be specified using any of the following environment variables (listed in order of precedence): `GOOGLE_ZONE`, `GCLOUD_ZONE`, `CLOUDSDK_COMPUTE_ZONE`. |
