---
# WARNING: this file was fetched from https://raw.githubusercontent.com/datarobot-community/pulumi-datarobot/v0.10.8/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: DataRobot
meta_desc: Provides an overview of the DataRobot Provider for Pulumi.
layout: package
---

The `DataRobot` provider for Pulumi can be used to provision any of the resources available with [DataRobot](https://www.datarobot.com/).

## Example

{{< chooser language "python,yaml,typescript" >}}

{{% choosable language python %}}

```python
import pulumi_datarobot as datarobot
import pulumi

use_case = datarobot.UseCase("example fom python",
                            name="example use case",
                            description="pulumi"
)

playground = datarobot.Playground("playground",
                                name="example playground",
                                use_case_id=use_case.id,
)
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: yaml
runtime: yaml
description: Example using Pulumi YAML 
config: {'pulumi:tags': {value: {'pulumi:template': yaml}}}
resources:
  datarobotUseCase:
    type: datarobot:UseCase
    properties:
      name: Pulumi YAML Example
      description: Example using Pulumi YAML
```

{{% /choosable %}}

{{% choosable language typescript %}}

```typescript
import * as datarobot from "@datarobot/pulumi-datarobot";

const useCase = new datarobot.UseCase("example from Typescript", {
    name: "example from TypeScript",
    description: "An example use case from Typescript",
});
```

{{% /choosable %}}

{{< /chooser >}}
