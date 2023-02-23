---
title: Cloud Toolkit AWS
meta_desc: Cloud Toolkit AWS bridges Infrastructure as Code and accessibility, allowing users to manage their infrastructure by leveraging real-world usage and patterns.
layout: overview
---

The Cloud Toolkit AWS provider for Pulumi provision well-architected solutions in [AWS](https://aws.amazon.com/). It uses the AWS SDK to manage and provision resources.

The AWS provider must be configured with credentials to deploy and update resources in AWS; see [Installation & Configuration](./installation-configuration/) for instructions.

## Example

{{< chooser language "typescript,javascript,python" >}}

{{% choosable language typescript %}}

```typescript
import * as ct from require("@cloudtoolkit/aws");

const bucket = new ct.storage.Bucket("mybucket");
```

{{% /choosable %}}

{{% choosable language javascript %}}

```typescript
const ct = require("@cloudtoolkit/aws");

const bucket = new ct.storage.Bucket("mybucket");
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import cloud_toolkit_aws as ct

bucket = ct.storage.Bucket("bucket")
```

{{% /choosable %}}

{{< /chooser >}}
