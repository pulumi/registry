---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-aws-static-website/v0.4.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: AWS Static Website
meta_desc: Provides an overview of the AWS Static Website component.
layout: overview
---

This component makes it easy to deploy a static website to s3 along with an optional CloudFront distribution using any of the supported Pulumi programming languages including markup languages like YAML and JSON.

## Examples

{{< chooser language "typescript,yaml" >}}
{{% choosable language typescript %}}

```typescript
const args =  {
    sitePath: "../website/build",
} as staticwebsite.WebsiteArgs

const site = new staticwebsite.Website("website", args);
```

{{% /choosable %}}
{{% choosable language yaml %}}

```yaml
resources:
  web:
    type: "aws-static-website:index:Website"
    properties:
      sitePath: "../website/build"
outputs:
  websiteURL: ${web.websiteURL}
```

{{% /choosable %}}
{{< /chooser >}}

## Input Properties

This component takes the following inputs.

- sitePath - the root directory containing the website's contents to be served (required)
- withCDN - provision a CloudFront CDN to serve content
- targetDomain - the domain used to serve the content. A Route53 hosted zone must exist for this domain is this option is specified
- index.html - the default document for the site. Defaults to index.html
- error404 - the default 404 error page
- certificateARN - the ARN of the ACM certificate to use for serving HTTPS. If one is not provided, a certificate will be created during the provisioning process
- cacheTTL - TTL inseconds for cached objects
- withLogs - provision a bucket to house access logs
- priceClass - the price class to use for the CDN. Defaults to `100` if not specified.

## Outputs

This component provides the following outputs. Some may not be available depending on the given input configuration (e.g. if `withCDN` was not specified there will be no output for `cdnDomainName` and `cdnURL`)

- bucketName - the name of the s3 bucket containing the website's contents
- bucketWebsiteURL - the website URL for the s3 bucket
- cdnDomainName - the CDN domain name
- cdnURL - the CDN's endpoint URL
- logsBucketName - the name of the s3 bucket containing the access logs
- websiteURL - the URL to access the website

## Notes:

- If specifying a target domain and provisioning a CloudFront distribution, it is assumed there is a hosted zone configured in route53 for that target domain.
- If you already have an ACM certificate provisioned for your domain, then you can simply pass the ARN as one of the input properties. If not we will attempt to provision one for you based on the target domain provided.
