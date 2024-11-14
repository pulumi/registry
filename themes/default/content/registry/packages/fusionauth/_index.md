---
# WARNING: this file was fetched from https://raw.githubusercontent.com/theogravity/pulumi-fusionauth/v5.1.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: FusionAuth
meta_desc: Provides configuration management for FusionAuth Provider
layout: package
---

FusionAuth for Pulumi can be used to configure FusionAuth instances.

This is bridged using the [gpsinsight Terraform FusionAuth package](https://github.com/gpsinsight/terraform-provider-fusionauth).

You can look at the [Terraform FusionAuth docs by gpsinsight](https://registry.terraform.io/providers/gpsinsight/fusionauth/latest/docs)
to get an idea of what is available in this package.

## Example

{{< chooser language "typescript" >}}

{{% choosable language typescript %}}

```typescript
import { Provider } from 'pulumi-fusionauth';

const fusionAuthProvider = new Provider('fusion-auth', {
  host: process.env.FUSION_AUTH_HOST_URL,
  apiKey: process.env.FUSION_AUTH_API_KEY,
});

// Create a new signing key
const clientDataSigningKey = new FusionAuthKey(
  'sample-jwt-key',
  {
    algorithm: 'RS256',
    name: 'Sample jwt key',
    length: 2048,
  },
  { provider: fusionAuthProvider },
);

export const appClientDataJwtKeyId = clientDataSigningKey.id;
```


{{% /choosable %}}

{{< /chooser >}}
