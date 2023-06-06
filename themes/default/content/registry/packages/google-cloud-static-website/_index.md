---
title: Google Cloud Static Website
meta_desc: Provides an overview of the Google Cloud Static Website component.
layout: package
---

This component makes it easy to deploy a static website to Google Cloud using any of the supported Pulumi programming languages.

## Usage

{{< chooser language "typescript,yaml" >}}

{{% choosable language typescript %}}

```typescript
import { Website } from "@pulumi/google-cloud-static-website";

const site = new Website("site", {
    sitePath: "./site",
});

export const { originURL } = site;
```

{{% /choosable %}}
{{% choosable language yaml %}}

```yaml
name: my-website
runtime: yaml
description: A static website built with pulumi-google-cloud-static-website.

resources:
  site:
    type: google-cloud-static-website:index:Website
    properties:
      sitePath: ./site

outputs:
  originURL: ${site.originURL}
```

{{% /choosable %}}
{{< /chooser >}}

## Input Properties

This component takes the following inputs.

- sitePath (string) - the root directory containing contents of the built website contents
- withCDN (boolean) - whether to provision a Google Cloud CDN to serve website content
- error404 (string) - the default error page for the website. Defaults to error.html
- index.html (string) - the default document for the site. Defaults to index.html
- domain (string) - the domain of the website
- subdomain (string) - The subdomain used to access the static website. If not specified will configure with apex/root domain of the DNS zone specified

## Outputs

- originURL - the direct URL of the website (storage bucket endpoint)
- cdnURL - the CDN URL for the site
- customDomainURL - the custom domain URL where the static website can be accessed


## Notes:

- The SSL certs can take anywhere from 60 - 90 mins after the update has been completed in order to fully provision. Upon first provisioning, there may be a short period of time where the cert is invalid when accessing the website over HTTPS.