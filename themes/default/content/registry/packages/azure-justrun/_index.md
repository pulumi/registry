---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-azure-justrun/v0.2.3/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Azure Justrun
meta_desc: Provides an overview of the Azure Justrun component.
layout: overview
---

This component makes it easy to deploy a simple web app to Azure using any of the supported Pulumi programming languages including markup languages like YAML and JSON.

## Examples

{{< chooser language "typescript,yaml" >}}
{{% choosable language typescript %}}

```typescript
import * as justrun from "@pulumi/pulumi-azure-justrun"
const args =  {
    filePath: "./www",
} as justrun.WebappArgs

const site = new staticwebsite.Website("website", args);
```

{{% /choosable %}}
{{% choosable language yaml %}}

```yaml
resources:
  web:
    type: "azure-justrun:index:webapp"
    properties:
      sitePath: "./www"
outputs:
  websiteURL: ${web.websiteURL}
```

{{% /choosable %}}
{{< /chooser >}}

## Input Properties

This component takes the following inputs.

- appSkuName - The name of the compute instance running the server. Also see appSkuTier
- appSkuTier - The tier of the compute instance running the server. Also see appSkuName
- filePath - The relative file path to the folder containing web files.
- containerPublicAccess - The public access level of the BlobContainer containg the website data.
- storageSkuName - The name of the SKU of the storage account created, if storageAccount is not provided
- storageAccount - The storage account to use. One will be created if not provided.
- resourceGroup - The resource group to use. One will be created if not provided.
- namePrefix - The name prefix given to child resources of this component. Should not contain dashes.

## Outputs

This component provides the following outputs.

- url - the URL to access the website
