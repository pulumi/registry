---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumiverse/pulumi-vercel/v2.8.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Vercel
meta_desc: Provides an overview of the Vercel Provider for Pulumi.
layout: package
---

The `Vercel` provider for Pulumi can be used to provision any of the resources available with `Vercel`.

## Example

{{< chooser language "javascript,typescript,python" >}}

{{% choosable language javascript %}}

```javascript
"use strict";
const vercel = require("@pulumiverse/vercel");

const project = new vercel.Project("pulumi-vercel", {
  name: "pulumi-vercel",
  rootDirectory: "src",
});

const deployment = new vercel.Deployment("deployment", {
  projectId: project.id,
  production: true,
  files: vercel.getProjectDirectoryOutput({ path: "./src" }).files,
});
```

{{% /choosable %}}

{{% choosable language typescript %}}

```typescript
import * as vercel from "@pulumiverse/vercel";

const project = new vercel.Project("pulumi-vercel", {
  name: "pulumi-vercel",
  rootDirectory: "src",
});

const deployment = new vercel.Deployment("deployment", {
  projectId: project.id,
  production: true,
  files: vercel.getProjectDirectoryOutput({ path: "./src" }).files,
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumiverse_vercel as vercel
import pulumi

project = vercel.Project("pulumi-vercel",
                         name="pulumi-vercel",
                         rootDirectory="src")

deployment = vercel.Deployment("deployment",
                         name="pulumi-vercel",
                         project_id=project.id,
                         production="true")
```

{{% /choosable %}}

{{< /chooser >}}
