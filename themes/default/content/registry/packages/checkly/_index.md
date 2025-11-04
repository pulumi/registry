---
# WARNING: this file was fetched from https://raw.githubusercontent.com/checkly/pulumi-checkly/v2.7.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/checkly/pulumi-checkly/blob/v2.7.0/docs/_index.md
title: Checkly
meta_desc: Provides an overview of the Checkly provider for Pulumi.
layout: package
---
## Installation

The Checkly provider is available as a package in most Pulumi languages:

* JavaScript/TypeScript: [`@checkly/pulumi`](https://www.npmjs.com/package/@checkly/pulumi)
* Python: [`pulumi-checkly`](https://pypi.org/project/pulumi-checkly/)
* Go: [`github.com/checkly/pulumi-checkly/sdk/v2/go/checkly`](https://github.com/checkly/pulumi-checkly)
* .NET: [`Pulumi.Checkly`](https://www.nuget.org/packages/Pulumi.Checkly)

## Overview

The Checkly provider for Pulumi can be used to provision any of the monitoring resources available in [Checkly](https://www.checklyhq.com/).

## Authentication

The Checkly provider must be configured with an `API Key` and an `Account ID` in order to deploy Checkly resources. Sign up for a [Checkly](https://www.checklyhq.com) account and follow our [integration guide](https://www.checklyhq.com/docs/integrations/pulumi/) to create and configure your credentials.

## Example usage

{{< chooser language "javascript,typescript" >}}

{{% choosable language javascript %}}

```javascript
const checkly = require('@checkly/pulumi')

new checkly.Check('api-check', {
  activated: true,
  frequency: 10,
  type: 'API',
  request: {
    method: 'GET',
    url: "https://checklyhq.com",
  },
})

new checkly.Check('browser-check', {
  activated: true,
  frequency: 10,
  type: 'BROWSER',
  script: 'console.log("Hello World!")',
})
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as checkly from '@checkly/pulumi'

new checkly.Check('api-check', {
  activated: true,
  frequency: 10,
  type: 'API',
  request: {
    method: 'GET',
    url: 'https://checklyhq.com',
  }
})

new checkly.Check('browser-check', {
  activated: true,
  frequency: 10,
  type: 'BROWSER',
  script: 'console.log("Hello World!")',
})
```

{{% /choosable %}}
{{< /chooser >}}

> You can find more examples in the [pulumi-checkly repository](https://github.com/checkly/pulumi-checkly/tree/main/examples).
