---
# WARNING: this file was fetched from https://raw.githubusercontent.com/theogravity/pulumi-fusionauth/v5.1.0/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: FusionAuth Installation & Configuration
meta_desc: Provides an overview on how to setup the FusionAuth Provider for Pulumi.
layout: package
---

## Installation

The FusionAuth provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`pulumi-fusionauth`](https://www.npmjs.com/package/pulumi-fusionauth)
* Python: [`theogravity-pulumi-fusionauth`](https://pypi.org/project/theogravity-pulumi-fusionauth/)
* Go: [`https://github.com/theogravity/pulumi-fusionauth/tree/main/sdk/go/fusionauth`](https://github.com/theogravity/pulumi-fusionauth/tree/main/sdk/go/fusionauth)
* .NET: [`theogravity.Fusionauth`](https://www.nuget.org/packages/theogravity.Fusionauth)

## Configuring Credentials

1. Export required variables:

```bash
$ export FUSION_AUTH_HOST_URL=<URL to your FusionAuth Instance>
$ export FUSION_AUTH_API_KEY=<API Key for FusionAuth APIs>
```

You can get / generate the API key by going to the FusionAuth UI:

1. `Settings` > `API keys` > add

2. Create an instance of the provider

   ```typescript
    import { Provider } from 'pulumi-fusionauth';

    const fusionAuthProvider = new Provider('fusion-auth', {
      host: process.env.FUSION_AUTH_HOST_URL,
      apiKey: process.env.FUSION_AUTH_API_KEY,
    });
   ```

3. Create the FusionAuth resources and assign it the provider

   ```typescript
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
   ```
