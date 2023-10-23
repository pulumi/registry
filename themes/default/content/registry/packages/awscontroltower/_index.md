---
title: AWS Control Tower
meta_desc: Provides an overview of the AWS Control Tower package for Pulumi
layout: overview
---

The AWS Control Tower provider for Pulumi can be used to provision AWS account with [AWS Control Tower](https://aws.amazon.com/controltower/).

The AWS Control Tower provider must be configured with the required permissions to manage AWS accounts in AWS Control Tower.

## Example

{{< chooser language "typescript,python" >}}
{{% choosable language typescript %}}

```typescript

import * as controltower from "@lbrlabs/pulumi-awscontroltower";

const account = new controltower.ControlTowerAwsAccount("account", {
  organizationalUnit: "Production",
  email: "mail@example.com",
  name: "Some User",
  organizationalUnitIdOnDelete: "ou-48hfnvbc-ufo",
  closeAccountOnDelete: true,
  sso: {
    firstName: "Some",
    lastName: "User",
    email: "mail@example.com",
  },
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import lbrlabs_pulumi_aws_controltower as controltower

account = controltower.ControlTowerAwsAccount(
    "account",
    organizational_unit="Production",
    email="mail@example.com",
    name="Some User",
    organizational_unit_id_on_delete="ou-48hfnvbc-ufo",
    close_account_on_delete=True,
    sso={
        "firstName": "Some",
        "lastName": "User",
        "email": "mail@example.com"
    }
```

{{% /choosable %}}
{{< /chooser >}}