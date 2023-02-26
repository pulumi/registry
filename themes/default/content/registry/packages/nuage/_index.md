---
title: Pulumi Nuage
meta_desc: Provides an overview of the Nuage Provider for Pulumi.
layout: overview
---

The Pulumi Nuage Provider is a valuable tool for provisioning cloud resources on Amazon Web Services (AWS). This provider boasts well-designed, pre-packaged infrastructure as code components that streamline the deployment process.

With Pulumi Nuage, users have access to various components such as `ContainerFunction`, `Repository`, and `ServerlessDatabase`. Each of these components represents a carefully curated bundle of AWS resources, offering a convenient and efficient solution for your cloud resource needs.

## Example

### ContainerFunction

The ContainerFunction component is a powerful tool for managing and executing Serverless Lambda Functions on AWS. This component provides a comprehensive set of features to streamline the deployment process, including an optimized Lambda architecture with the capability to build Docker containers, function URLs, and CloudWatch keep-warm rules. The ContainerFunction component efficiently manages the build and deployment of Docker builds, untags automatically generated random image names during the build process and retaining only human-readable names for easy identification. Furthermore, this component creates a log group with a retention policy and the necessary policies to ensure secure and efficient operation.

Here are the summary of features:

* Lambda Functions
* Function URLs
* CloudWatch keep-warm rules
* Management of build and deployment of Docker builds
* Removal of docker build artifacts (randomly generated image names that pollute your local docker)
* Creation of a log group with a retention policy
* Creating Role to run Lambda and Write Logs
* Automated X-Ray tracing

### Repository

The Repository component is a convenient tool for managing Amazon Elastic Container Registry (ECR) repositories. It offers the ability to create an ECR repository with a lifecycle policy on demand. The output of this repository can be easily utilized as an input for the ContainerFunction component, providing a seamless integration of these two components.


{{< chooser language "typescript,python" >}}
{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as nuage from "@pulumi/nuage";

const repository = new nuage.aws.Repository("foo",{
    name:"repository",
    expireInDays:30
});

const lambdaContainer = new nuage.aws.ContainerFunction("foo",{
    name:"lambda-function",
    description:"Nuage AWS ContainerFunction resource.",
    repositoryUrl:repository.url,
    architecture:"x86_64",
    memorySize:512,
    timeout:30,
    environment:{bar:"baz"},
    keepWarm:true,
    url:true,
    logRetentionInDays:90
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_nuage as nuage

repository = nuage.aws.Repository(
    "foo",
    name="repository",
    expire_in_days=30,
)

container_function = nuage.aws.ContainerFunction("foo",
    name="lambda-function",
    description="Nuage AWS ContainerFunction resource.",
    repository_url=repository.url,    
    architecture="X86_64",
    memory_size=512,
    timeout=30,
    environment={"bar":"baz"},
    keep_warm=True,
    url=True,
    log_retention_in_days=90
)
```

{{% /choosable %}}

{{< /chooser >}}


### ServerlessDatabase

The ServerlessDatabase component is an effective solution for creating serverless databases using Amazon Relational Database Service (RDS) Aurora databases. It streamlines the process by automatically creating the necessary components, including a subnet group, security group, security group rules, and an RDS cluster. Additionally, this component generates a random password securely managed by Pulumi for the DB credentials.

With the ServerlessDatabase component, you can create fully configured databases in either MySQL or PostgreSQL with ease. The end result is a comprehensive serverless database resource with all the necessary configurations in place, providing a convenient and efficient solution for your serverless database needs.

{{< chooser language "typescript,python" >}}
{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as nuage from "@pulumi/nuage";

const db = new nuage.aws.Repository("foo",{
    vpcId:my_vpc.id,
    vpcSubnets:my_vpc.private_subnet_ids,
    databaseType:"mysql",
    databaseName:"bar",
    masterUserName:"root",
    ipWhitelist:["0.0.0.0/0"],
    skipFinalSnapshot:true
});
```
 
{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_nuage as nuage

db = nuage.aws.ServerlessDatabase(
    "foo",
    vpc_id=my_vpc.id,
    vpc_subnets=my_vpc.private_subnet_ids,
    database_type="mysql",
    database_name="bar",
    master_username="root",
    ip_whitelist=["0.0.0.0/0"],
    skip_final_snapshot=True
)
```

{{% /choosable %}}

{{< /chooser >}}
