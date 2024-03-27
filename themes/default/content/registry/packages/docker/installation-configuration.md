---
title: Docker Installation & Configuration
meta_desc: Provides an overview on how to configure the Pulumi Docker Provider.
layout: package
---

The Pulumi Docker provider uses the Docker SDK to manage resources.

## Installation

The Docker provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/docker`](https://www.npmjs.com/package/@pulumi/docker)
* Python: [`pulumi-docker`](https://pypi.org/project/pulumi-docker/)
* Go: [`github.com/pulumi/pulumi-docker/sdk/v4/go/docker`](https://github.com/pulumi/pulumi-docker)
* .NET: [`Pulumi.Docker`](https://www.nuget.org/packages/Pulumi.Docker)
* Java: [`com.pulumi/docker`](https://central.sonatype.com/artifact/com.pulumi/docker)

## Configuring The Provider

Pulumi relies on the Docker SDK to authenticate requests from your computer to Docker. Your credentials are never sent
to pulumi.com.

Once the credentials are obtained, there are two ways to communicate your configuration tokens to Pulumi:

1. Set the environment variable `DOCKER_HOST`:

    ```bash
    $ export DOCKER_HOST=tcp://127.0.0.1:2376/
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set docker:host tcp://127.0.0.1:2376/
    ```

Alternative for Windows users:
    
Use the following PowerShell command to add reference the pipe used by Docker Engine on your machine. 

```powershell
pulumi config set docker:host "npipe:////.//pipe//docker_engine"
```
