---
title: Pulumi Nuage
meta_desc: Provides an overview of the Nuage Provider for Pulumi.
layout: overview
---

Pulumi Nuage provider can be used to provision AWS cloud resources.

## Example

{{< chooser language "typescript,python" >}}
{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as nuage from "@pulumi/nuage";

const repository = new nuage.aws.Repository("repository-resourcename",{
    name:"repository",
    expireInDays:30
});

const lambdaContainer = new nuage.aws.ContainerFunction("myfunction",{
    name:"lambda-function",
    dockerfile:"./Dockerfile",
    repositoryUrl:repository.url
});
```
 
{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_nuage

repository = pulumi_nuage.aws.Repository(
    "repository-resourcename",
    name="repository",
    expire_in_days=30,
)

container_function = pulumi_nuage.aws.ContainerFunction("myfunction",
    name="lambda-function",
    repository_url=repository.url,
    dockerfile="./lambda/Dockerfile.lambda",
)
```

{{% /choosable %}}

{{< /chooser >}}
