---
title: AWS Control Tower Installation & Configuration
meta_desc: Information on how to install the AWS Control Tower provider.
layout: package
---

## Installation

The Pulumi AWS Control Tower provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@lbrlabs/pulumi-awscontroltwer`](https://www.npmjs.com/package/@lbrlabs/pulumi-awscontroltwer)
* Python: [`lbrlabs_pulumi_aws_controltwer`](https://pypi.org/project/lbrlabs-pulumi-aws-controltower/)
* Go: [`github.com/lbrlabs/pulumi-awscontroltower/sdk/go/controltwer`](https://pkg.go.dev/github.com/lbrlabs/pulumi-awscontroltower/sdk)
* .NET: [`Lbrlabs.PulumiPackage.Scaleway`](https://www.nuget.org/packages/Lbrlabs.PulumiPackage.Awscontroltower)

### Provider Binary

The AWS control tower provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource awscontroltower --server github://api.github.com/lbrlabs
```

Replace the version string with your desired version.

## Setup

The AWS control tower uses the AWS SDK to connect to AWS. Many of the authentication options you're familiar with using [Pulumi AWS)(https://www.pulumi.com/registry/packages/aws/) apply here.

{{% /choosable %}}

{{% choosable os macos %}}

```bash
$ export AWS_ACCESS_KEY_ID=<AWS_ACCESS_KEY_ID>
$ export AWS_SECRET_ACCESS_KEY=<AWS_SECRET_ACCESS_KEY>
```

{{% /choosable %}}

{{% choosable os windows %}}

```powershell
> $env:AWS_ACCESS_KEY_ID = "<AWS_ACCESS_KEY_ID>"
> $env:AWS_SECRET_ACCESS_KEY = "<AWS_SECRET_ACCESS_KEY>"
```

{{% /choosable %}}
{{< /chooser >}}

## Configuration Options

Use `pulumi config set awscontroltower:<option>` or pass options to the [constructor of `new awscontroltower.Provider`]({{< relref "/registry/packages/awscontroltower/api-docs/provider" >}}).