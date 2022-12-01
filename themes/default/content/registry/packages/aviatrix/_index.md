---
title: Aviatrix
meta_desc: Provides an overview of the Aviatrix Provider for Pulumi.
layout: overview
---

The Aviatrix Resource Provider lets you manage [Aviatrix](https://aviatrix.com) resources.
The Aviatrix provider must be configured with a `Contoller IP` and needs the controller's  `Username` and `Password` in order to deploy Aviatrix resources.

## Disclaimer
This is an unofficial Aviatrix Pulumi provider and not supported by Aviatrix in any way. The material embodied in this software/code is provided to you "as-is" and without warranty of any kind, express, implied or otherwise, including without limitation, any warranty of fitness for a particular purpose. In no event shall the Aviatrix Inc. be liable to you or anyone else for any direct, special, incidental, indirect or consequential damages of any kind, or any damages whatsoever, including without limitation, loss of profit, loss of use, savings or revenue, or the claims of third parties, whether or not Aviatrix Inc. has been advised of the possibility of such loss, however caused and on any theory of liability, arising out of or in connection with the possession, use or performance of this software/code.

## Example

{{< chooser language "python" >}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_aviatrix as aviatrix
import os

vpc_tr1 = aviatrix.AviatrixVpc(
    "pulumi_transit1_vpc",
    cloud_type = 8,
    account_name = os.environ.get("AZURE_ACCT_NAME"),
    region = os.environ.get("AZURE_REGION"),
    name = "pulumi-tr1",
    cidr = "10.0.0.0/23",
    aviatrix_firenet_vpc = True
)
```

{{% /choosable %}}
{{< /chooser >}}

> You may find more complete and detailed examples in the pulumi-aviatrix [repository](https://github.com/astipkovits/pulumi-aviatrix/tree/main/examples)
