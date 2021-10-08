---
title: Google Native Setup
meta_desc: How to set up credentials to use the Pulumi Google Native Provider and choose configuration options to tailor the provider to suit your use case.
layout: installation
---

To provision resources with the Pulumi Google Cloud Provider, you need to have Google credentials.

{{% configure-gcp %}}

{{% notes "info" %}}
If you are using Pulumi in an non-interactive setting (such as a CI/CD system) you will need to [configure and use a service account]({{< relref "service-account" >}}) instead.
{{% /notes %}}
