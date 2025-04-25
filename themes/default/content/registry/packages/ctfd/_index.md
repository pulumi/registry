---
# WARNING: this file was fetched from https://raw.githubusercontent.com/ctfer-io/pulumi-ctfd/v2.1.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: CTFd Provider
meta_desc: Provides an overview on how to setup the CTFd Provider for Pulumi.
layout: package
---

[CTFd](https://ctfd.io) is an open-source CTFd platform famous for its simplicity and extensibility.
It has been used for various Capture The Flag (CTF) events, originally for the [CSAW](https://www.csaw.io/ctf), later by others such the [BreizhCTF](https://www.breizhctf.com/).

One of [CTFer](https://ctfer.io) approach is seeing the CTFd resources as simply as objects with CRUD operations (e.g. challenges).
From this simplification arise a possibility: managing them as Terraform resources, which we did with the [Terraform Provider for CTFd](https://github.com/ctfer-io/terraform-provider-ctfd).
Nevertheless, due to the limitations of Terraform and the necessity to bring reproducibility to Capture The Flag events, as stated by multiple organizers, we had to go further and give it an additional predicate: managing them as code through programming language.

[Pulumi](https://www.pulumi.com/docs/get-started) is an Infrastructure-as-Code (IaC) and Free Open-Source Software (FOSS) tool that is a first-class opportunity for this.

To sum it up, you can manage a Capture The Flag event based upon CTFd in multiple programming languages.

## Use Cases

Users of the CTFd provider can:

- Drop unstable scripts and tooling to manage CTF resources in favor of this SDK.
- Share their configuration to public through a standardized API (common and encouraged after a Capture The Flag event), such that it could be replayed later for training and/or teaching.
- Plug in other providers (e.g. [Kubernetes](https://www.pulumi.com/registry/packages/kubernetes/)) to provision infrastucture for the platform and challenges.

## Provider

To use the CTFd Provider you will have to configure it and use it to create resources: it can't infer the CTFd url and credentials.

{{< chooser language "javascript,typescript,go,python,csharp" >}}

{{% choosable language javascript %}}

```javascript
import * as ctfd from '@ctfer-io/pulumi-ctfd';

// Create provider
let pv = new ctfd.Provider('ctfd-fine-grained', {
    url: 'https://my-ctf.lan',
    apiKey: 'ctfd_xxx', // please do not hardcode your credentials/api keys
});

// Create resources with the custom provider
let ch = new ctfd.ChallengeStandard('some-challenge', {
    name: 'My Challenge',
    category: 'misc',
    description: '...',
    value: 500,
}, { provider: pv });
```

{{% /choosable %}} {{% choosable language typescript %}}

```typescript
import * as ctfd from '@ctfer-io/pulumi-ctfd';

// Stack configuration, other resources, etc.
// ...

// Create provider
let pv = new ctfd.Provider('ctfd-fine-grained', {
    url: 'https://my-ctf.lan',
    apiKey: 'ctfd_xxx', // please do not hardcode your credentials/api keys
});

// Create resources with the custom provider
let ch = new ctfd.ChallengeStandard('some-challenge', {
    name: 'My Challenge',
    category: 'misc',
    description: '...',
    value: 500,
}, { provider: pv });

// Other resources, export, etc.
// ...
```

{{% /choosable %}} {{% choosable language go %}}

```go
package main

import (
    "github.com/ctfer-io/pulumi-ctfd/sdk/v2/go/ctfd"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
        // Stack configuration, other resources, etc.
        // ...

        // Create provider
        pv, err := ctfd.NewProvider(ctx, "ctfd-fine-grained", &ctfd.ProviderArgs{
            Url:    pulumi.String("https://my-ctf.lan"),
            ApiKey: pulumi.String("ctfd_xxx"), // please do not hardcode your credentials/api keys
        })
        if err != nil {
            return err
        }

        // Build resource options to use it
        opts := []pulumi.ResourceOption{
            pulumi.Provider(pv),
        }

        // Create resources with the custom provider
        _, err = ctfd.NewChallengeStandard(ctx, "some-challenge", &ctfd.ChallengeStandardArgs{
            Name:        pulumi.String("My Challenge"),
            Category:    pulumi.String("misc"),
            Description: pulumi.String("..."),
            Value:       pulumi.Int(500),
        }, opts...)
        if err != nil {
            return err
        }

        // Other resources, export, etc.
        // ...

        return nil
    })
}
```

{{% /choosable %}} {{% choosable language python %}}

```python
import ctfer-io_pulumi_ctfd as ctfd
import pulumi

# Stack configuration, other resources, etc.
# ...

# Create provider
pv = ctfd.Provider("ctfd-fine-grained", url="https://my-ctf.lan", api_key="ctfd_xxx") # please do not hardcode your credentials/api keys

# Create resources with the custom provider
ch = ctfd.ChallengeStandard("some-challenge", name="My Challenge", category="misc", description="...", value=500, opts=pulumi.ResourceOptions(provider=pv))

# Other resources, export, etc.
# ...
```

{{% /choosable %}} {{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;
using CTFerio.Ctfd;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            // Stack configuration, other resources, etc.
            // ...

            // Create provider
            var pv = new Ctfd.Provider("ctfd-fine-grained", new Ctfd.ProviderArgs{
                Url = "https://my-ctf.lan",
                ApiKey = "ctfd_xxx" // please do not hardcode your credentials/api keys
            });

            // Create resources with the custom provider
            var ch = new Ctfd.ChallengeStandard("my-challenge", new Ctfd.ChallengeArgs{
                Name = "My Challenge",
                Category = "misc",
                Description = "...",
                Value = 500
            }, new Pulumi.CustomResourceOptions { Provider = pv });

            // Other resources, export, etc.
            // ...
        });
}
```

{{% /choosable %}}

{{< /chooser >}}
