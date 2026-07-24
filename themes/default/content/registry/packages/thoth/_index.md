---
# WARNING: this file was fetched from https://raw.githubusercontent.com/atensecurity/pulumi-thoth/v0.1.15/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/atensecurity/pulumi-thoth/blob/v0.1.15/docs/_index.md
title: Thoth
meta_desc: Manage Aten Security Thoth resources with Pulumi - runtime authorization and enforcement for AI agents.
layout: package
---

# Thoth Provider

The Thoth Pulumi provider lets you manage Aten Security Thoth governance
resources from Pulumi programs.

This package is published as:

- **Node.js:** `@atensec/pulumi-thoth`
- **Python:** `pulumi-thoth`
- **.NET:** `AtenSecurity.Pulumi.Thoth`

For installation and configuration instructions, see
[`installation-configuration`](./installation-configuration).

## Minimal examples

### Node.js

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as thoth from "@atensec/pulumi-thoth";

const cfg = new pulumi.Config();
const provider = new thoth.Provider("thoth", {
  tenantId: cfg.require("tenantId"),
});

new thoth.governance.GovernanceSettings(
  "baseline-governance",
  {
    complianceProfile: "soc2",
    shadowLow: "allow",
    shadowMedium: "step_up",
    shadowHigh: "block",
    shadowCritical: "block",
  },
  { provider }
);
```

### Python

```python
import pulumi
import pulumi_thoth as thoth

cfg = pulumi.Config()
provider = thoth.Provider("thoth", tenant_id=cfg.require("tenantId"))

thoth.governance.GovernanceSettings(
    "baseline-governance",
    compliance_profile="soc2",
    shadow_low="allow",
    shadow_medium="step_up",
    shadow_high="block",
    shadow_critical="block",
    opts=pulumi.ResourceOptions(provider=provider),
)
```

### C#

```csharp
using AtenSecurity.Pulumi.Thoth;
using AtenSecurity.Pulumi.Thoth.Governance;
using Pulumi;

var config = new Config();
var provider = new Provider("thoth", new ProviderArgs
{
    TenantId = config.Require("tenantId"),
});

_ = new GovernanceSettings("baseline-governance", new GovernanceSettingsArgs
{
    ComplianceProfile = "soc2",
    ShadowLow = "allow",
    ShadowMedium = "step_up",
    ShadowHigh = "block",
    ShadowCritical = "block",
}, new CustomResourceOptions { Provider = provider });
```
