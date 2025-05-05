---
# WARNING: this file was fetched from https://raw.githubusercontent.com/crowdstrike/pulumi-crowdstrike/v0.0.14/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: CrowdStrike
meta_desc: Provides an overview of the CrowdStrike Provider for Pulumi.
layout: overview
---

The CrowdStrike provider for Pulumi can be used to provision any of the [supported cloud resources](#supported_resources) available in CrowdStrike.

The CrowdStrike provider must be configured with credentials to deploy and update resources in CrowdStrike.

## Example

### Python

```python
import pulumi
import crowdstrike_pulumi

host_group = crowdstrike_pulumi.HostGroup(
    resource_name="hostgroup_1",
    description="A host group created using Pulumi",
    type="dynamic",
    assignment_rule="tags:'SensorGroupingTags/production'+os_version:'Amazon Linux 2'",
    
)
```

### Typescript

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as crowdstrike from "@crowdstrike/pulumi"

//
const hostGroup = new crowdstrike.HostGroup("hostgroup_2", {
    "description": "A host group created using Pulumi",
    "type": "dynamic",
    "assignmentRule": "tags:'SensorGroupingTags/production'+os_version:'Amazon Linux 2'"
}
)

```

### Go

```go
package main

import (
 "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
 "github.com/crowdstrike/pulumi-crowdstrike/sdk/go/crowdstrike"
)

func main() {
 pulumi.Run(func(ctx *pulumi.Context) error {
  hostGroup, err := crowdstrike.NewHostGroup(ctx, "hostgroup_3", &crowdstrike.HostGroupArgs{
   Description:    pulumi.String("A host group created using Pulumi"),
   Type:           pulumi.String("dynamic"),
   AssignmentRule: pulumi.String("tags:'tags:'SensorGroupingTags/production'+os_version:'Amazon Linux 2'"),
  })
  if err != nil {
   return err
  }
  ctx.Export("hostgroup_3", hostGroup)
  return nil
 })
}

```

### CSharp

```csharp
using System;
using Pulumi;
using CrowdStrike.Crowdstrike;

class MyStack : Stack
{
    public MyStack()
    {
        var hostGroup = new HostGroup("hostgroup_pulumi_csharp_published", new HostGroupArgs
        {
            Name = "hostgroup_pulumi_dotnet",
            Type = "dynamic",
            Description = "Test pulumi hostgroup",
            AssignmentRule = "tags:'SensorGroupingTags/cloud-lab'+os_version:'Amazon Linux 2'"
        });
    }
}

```

## <a name="supported_resources"></a> Supported Resources and Required Scopes
>
> [!NOTE]
> See a resource's documentation for the specific scopes required for that resource.

| Scope                   | Permission      |
|-------------------------|-----------------|
| Device Control Policies | *READ*, *WRITE* |
| Prevention Policies     | *READ*, *WRITE* |
| Response Policies       | *READ*, *WRITE* |
| Firewall Management     | *READ*, *WRITE* |
| Host Groups             | *READ*, *WRITE* |
| Sensor Update Policies  | *READ*, *WRITE* |
| Falcon FileVantage      | *READ*, *WRITE* |
