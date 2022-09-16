---
title: Azure Static Website
meta_desc: Provides an overview of the Azure Static Website component.
layout: overview
---

This component makes it easy to deploy a static website to Azure using any of the supported Pulumi programming languages.

## Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as website from "@pulumi/azure-static-website";

const site = new website.Website("site", {
    sitePath: "./site",
});

export const { originURL } = site;
```

```yaml
name: my-website
runtime: yaml
description: A static website build with pulumi-azure-static-website.

resources:
  site:
    type: azure-static-website:index:Website
    properties:
      sitePath: ./site

outputs:
  originURL: ${site.originURL}
```

## Input Properties

This component takes the following inputs.

- sitePath (string) - the root directory containing the website's contents to be served (required)
- withCDN (boolean) - provision a CDN to serve content
- error404 (string) - the default 404 error page
- index.html (string) - the default document for the site. Defaults to index.html
- domainResourceGroup (string) - The name of the resource group your domain is attached to
- dnsZoneName (string) - The name of the DNS zone that will be used to serve the static website. This must be set in order for this component to make the site accessible from a custom domain. See [Azure docs](https://docs.microsoft.com/en-us/azure/dns/dns-zones-records) for more info.
- subdomain (string) - The subdomain used to access the static website. If not specified will configure with apex/root domain of the DNS zone specified

## Outputs

- originURL - the Storage URL for the site
- cdnURL - the CDN URL for the site
- customDomainURL - the custom domain URL where the static website can be accessed
- resourceGroupName - the name of the resource group that was provisioned to contain the needed static website resources


## Notes:

- If you would like to serve your site from a custom domain, you need to configure a DNS zone in Azure and set the name of the zone using the `dnsZoneName` input property.
- If a subdomain is not specified, the contents will be served from the apex. Serving the website over HTTPS is something that will need to be manually configured, as Azure will sign free certs for the subdomains, but not for the root domain.
- When destroying the site, you will need to manually delete the CNAME record that was provisioned for the domain (either using the console or CLI) in order for `pulumi destroy` to succeed.