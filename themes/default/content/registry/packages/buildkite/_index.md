---
title: Buildkite
meta_desc: Provides an overview of the Buildkite Provider for Pulumi.
layout: overview
---

The Buildkite provider for Pulumi can be used to manage resources in
the [Buildkite](https://buildkite.com) CI/CD platform. The Buildkite
provider must be configured with credentials to deploy and update
resources.

## Example

{{< chooser language "typescript,python,go" >}}
{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as buildkite from "@grapl/pulumi-buildkite";

new buildkite.Pipeline(
    "testing-pipeline",
    {
        name: "testing-pipeline",
        repository: "https://github.com/grapl-security/pulumi-buildkite",
        description: ":pulumi::buildkite::nodejs: A pipeline created to test the pulumi-buildkite provider",
        defaultBranch: "main",
        steps: `
        steps:
          - label: "Hello World"
            command: echo 'Hello World'
        `
    }
);

```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
import pulumi_buildkite as buildkite

def main() -> None:
    buildkite.Pipeline(
        "testing-pipeline",
        name="testing-pipeline",
        repository="https://github.com/grapl-security/pulumi-buildkite",
        description=":pulumi::buildkite::python: A pipeline created to test the pulumi-buildkite provider",
        default_branch="main",
        steps="""
        steps:
          - label: "Hello World"
            command: echo 'Hello World'
        """
    )

if __name__ == "__main__":
    main()

```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
    "fmt"
    buildkite "github.com/grapl-security/pulumi-buildkite/sdk/go/buildkite"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {

        _, err := buildkite.NewPipeline(ctx, "testing-pipeline", &buildkite.PipelineArgs{
            Name:          pulumi.String("testing-pipeline"),
            Repository:    pulumi.String("https://github.com/grapl-security/pulumi-buildkite"),
            Description:   pulumi.String(":pulumi::buildkite::go: A pipeline created to test the pulumi-buildkite provider"),
            DefaultBranch: pulumi.String("main"),
            Steps: pulumi.String(`
        steps:
          - label: "Hello World"
            command: echo 'Hello World'
`),
        })
        if err != nil {
            return fmt.Errorf("error creating pipeline: %v", err)
        }

        return nil
    })
}
```

{{% /choosable %}}
{{< /chooser >}}
