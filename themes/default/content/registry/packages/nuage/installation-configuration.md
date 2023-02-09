---
title: Pulumi Nuage Installation & Configuration
meta_desc: Information on how to install the Pulumi Nuage provider.
layout: installation
---

## Installation

To use this resource, you need to install Pulumi plugin first (independent from the development language).

1. Download the latest release from <https://github.com/nuage-studio/pulumi-nuage/releases/>
2. Install the .tar.gz package using Pulumi

```bash
pulumi plugin install resource nuage 0.0.1 --file pulumi-resource-nuage-v0.0.1-darwin-amd64.tar.gz
```

You may also automate this process via Taskfile. Here is an example:

```yaml
version: "3"

tasks:
  pulumi_install:
    desc: Downloads binary
    vars:
      VERSION: 0.0.1
      FILE_NAME: "pulumi-resource-nuage-v0.0.1-darwin-amd64.tar.gz"
    cmds:
      - curl -L -o {{.FILE_NAME}} https://github.com/nuage-studio/pulumi-nuage/releases/download/{{.VERSION}}/{{.FILE_NAME}}
      - pulumi plugin install resource nuage {{.VERSION}} --file {{.FILE_NAME}}
      - rm {{.FILE_NAME}}
```

## Setup

To provision resources with the Pulumi Nuage provider,  you can set up the sdk of the language you prefer.

* JavaScript/TypeScript: [`@pulumi/nuage`](https://github.com/nuage-studio/pulumi-nuage/tree/master/sdk/nodejs)
* Python: [`pulumi_nuage`](https://github.com/nuage-studio/pulumi-nuage/tree/master/sdk/python)
* Go: [`nuage`](https://github.com/nuage-studio/pulumi-nuage/tree/master/sdk/go/nuage)
* .NET: [`Pulumi.Nuage`](https://github.com/nuage-studio/pulumi-nuage/tree/master/sdk/dotnet)



## Setup Instructions for Python SDK

You can use the following command to add python library via Poetry:

```bash
poetry add "https://github.com/nuage-studio/pulumi-nuage.git#subdirectory=sdk/python"
```

You can check related [Poetry Docs](https://python-poetry.org/docs/dependency-specification/) for more detailed information about subdirectory dependency.

## Configuration Options

You don't need to set any config variable to use Pulumi Nuage.