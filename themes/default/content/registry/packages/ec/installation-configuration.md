---
title: ElasticCloud Installation & Configuration
meta_desc: Provides an overview on how to configure the Pulumi ElasticCloud Provider.
layout: package
---

The Pulumi ElasticCloud provider uses the ElasticCloud SDK to manage resources.

## Installation

The ElasticCloud provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/ec`](https://www.npmjs.com/package/@pulumi/ec)
* Python: [`pulumi-ec`](https://pypi.org/project/pulumi-ec/)
* Go: [pulumi/pulumi](https://github.com/pulumi/pulumi-ec)
* .NET: [`Pulumi.ElasticCloud`](https://www.nuget.org/packages/Pulumi.ElasticCloud)
* Java: [`com.pulumi/ec`](https://central.sonatype.com/artifact/com.pulumi/ec)

### Configuring The Provider

Pulumi relies on the ElasticCloud SDK to authenticate requests from your computer to ElasticCloud. Your credentials are never sent
to pulumi.com. Once the credentials are obtained, there are two ways to communicate your configuration parameters to Pulumi:

1. Set the environment variable `EC_API_KEY`:

    ```bash
    $ export EC_API_KEY=XXXXXX
    ```

2. If you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set ec:apikey XXXXXX --secret
    ```

 The complete list of
configuration parameters is in the [ElasticCloud provider README](https://github.com/pulumi/pulumi-ec/blob/master/README.md).
