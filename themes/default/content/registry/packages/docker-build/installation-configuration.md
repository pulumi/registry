---
title: Docker-Build Installation & Configuration
meta_desc: Provides an overview on how to configure the Pulumi Docker-Build Provider.
layout: package
---

The Pulumi Docker-build provider builds modern Docker images with [buildx](https://docs.docker.com/reference/cli/docker/buildx/) and [BuildKit](https://docs.docker.com/build/buildkit/).

## Installation

The Docker-Build provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/docker-build`](https://www.npmjs.com/package/@pulumi/docker-build)
* Python: [`pulumi-docker-build`](https://pypi.org/project/pulumi-docker-build/)
* Go: [`github.com/pulumi/pulumi-docker-build/sdk/go/dockerbuild`](https://github.com/pulumi/pulumi-docker-build)
* .NET: [`Pulumi.DockerBuild`](https://www.nuget.org/packages/Pulumi.DockerBuild)
* Java: [`com.pulumi/docker-build`](https://central.sonatype.com/artifact/com.pulumi/docker-build)

## Configuring The Provider

### Host

The `DOCKER_HOST` environment variable can be used to specify a custom build daemon's location.

```bash
$ export DOCKER_HOST=tcp://127.0.0.1:2376/
```

This can also be specified in your stack's configuration:

```bash
$ pulumi config set docker-build:host tcp://127.0.0.1:2376/
```
