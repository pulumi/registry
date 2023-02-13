---
title: Pulumi Nuage Installation & Configuration
meta_desc: Information on how to install the Pulumi Nuage provider.
layout: installation
---

## Installation


To provision resources with the Pulumi Nuage provider,  you can set up the sdk of the language you prefer.

* JavaScript/TypeScript: [`@pulumi/nuage`](https://github.com/nuage-studio/pulumi-nuage/tree/master/sdk/nodejs)
* Python: [`pulumi_nuage`](https://github.com/nuage-studio/pulumi-nuage/tree/master/sdk/python)
* Go: [`nuage`](https://github.com/nuage-studio/pulumi-nuage/tree/master/sdk/go/nuage)
* .NET: [`Pulumi.Nuage`](https://github.com/nuage-studio/pulumi-nuage/tree/master/sdk/dotnet)

## Credentials

Pulumi Nuage components contains AWS Components from `pulumi-aws`and you can refer to corresponding docs for setting AWS credentials.  [correspnding docs](https://www.pulumi.com/registry/packages/aws/installation-configuration/)

## Setup Instructions for Python SDK

You can use the following command to add python library via Poetry:

```bash
poetry add "https://github.com/nuage-studio/pulumi-nuage.git#subdirectory=sdk/python"
```

You can check related [Poetry Docs](https://python-poetry.org/docs/dependency-specification/) for more detailed information about subdirectory dependency.

## Configuration Options

You don't need to set any config variable to use Pulumi Nuage.