---
title: Defang (Deprecated)
meta_desc: The Defang provider has been split into cloud-specific packages for AWS, GCP, and Azure.
layout: package
---

{{% notes type="info" %}}
This package has been split into three cloud-specific providers. Choose the one for your target cloud:

- [`defang-aws`](/registry/packages/defang-aws/): deploy containerized services to AWS.
- [`defang-gcp`](/registry/packages/defang-gcp/): deploy containerized services to GCP.
- [`defang-azure`](/registry/packages/defang-azure/): deploy containerized services to Azure.
{{% /notes %}}

The original `defang` provider deployed Docker Compose projects to any supported cloud through a single package with a `providerID` input. It has been replaced by one provider package per cloud, each with the same Compose-shaped inputs, so you pick the cloud you target at install time rather than at configuration time.

To migrate, install the package for your cloud (for example `@defang-io/pulumi-defang-aws`) and construct a `Project` resource with your services. See the per-cloud package pages linked above for installation and examples in every supported language.
