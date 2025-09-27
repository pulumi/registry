---
# WARNING: this file was fetched from https://raw.githubusercontent.com/runpod/pulumi-runpod-native/v1.9.99/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/runpod/pulumi-runpod-native/blob/v1.9.99/docs/installation-configuration.md
title: RunPod Setup
meta_desc: Information on how to install the RunPod Provider for Pulumi.
layout: package
---

## Installation

The RunPod provider is available as a package in these Pulumi languages:

* JavaScript/TypeScript: [`@runpod-infra/pulumi`](https://www.npmjs.com/package/@runpod-infra/pulumi)
* Python: [`runpodinfra`](https://pypi.org/project/runpodinfra/)
* Go: [`github.com/runpod/pulumi-runpod-native/tree/main/sdk/go/runpod`](https://www.github.com/runpod/pulumi-runpod-native)
* Pulumi YAML

## Pulumi guide
Create an empty directory and navigate inside

```bash
mkdir -p empty
cd empty
```

Create a new pulumi stack:
```
    pulumi new
```

A dropdown box will appear. Please select a minimal project from in there. For example, if you prefer using Python, you would select the following:
```bash
    python                             A minimal Python Pulumi program
```

Then populate the entrypoint file (__main__.py/main.go/index.ts) with the your data. Please use the guide below to understand what parameters you can pass. For Python, please remember to activate the virtual environment.

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/runpod/pulumi-runpod-native/sdk/go/runpod
```

We advise you to pin to a specific version.

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @runpod-infra/pulumi
```

or `yarn`:

```bash
yarn add @runpod-infra/pulumi
```

### Python

Please activate the virtual environment created when you ran the ```pulumi new``` command:

```bash
    source venv/bin/activate
```

Then install the official pulumi package from our pip repository:

```bash
pip install runpodinfra
```

## Configuration

To begin with, please set your RunPod API key using Pulumi.

```bash
pulumi config set runpod:token --secret
```
