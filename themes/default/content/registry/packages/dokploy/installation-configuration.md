---
# WARNING: this file was fetched from https://raw.githubusercontent.com/dimeskigj/pulumi-dokploy/v0.3.1/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/dimeskigj/pulumi-dokploy/blob/v0.3.1/docs/installation-configuration.md
layout: package
title: Installation and configuration
meta_desc: Install and configure the community-maintained Pulumi Dokploy provider.
description: Install and configure the community-maintained Pulumi Dokploy provider.
---

# Installation and configuration

The provider's `pluginDownloadURL` metadata is
`github://api.github.com/dimeskigj/pulumi-dokploy`, which points Pulumi at the
GitHub release artifacts. Pulumi can acquire the provider plugin automatically
when a package is added or used; no separate plugin download is normally
required.

Install the SDK for your language. Replace `${DOKPLOY_VERSION}` with the
provider version you are using:

```bash
npm install @dimeskigj/pulumi-dokploy
pip install pulumi-dokploy
go get github.com/dimeskigj/pulumi-dokploy/sdk/go/dokploy
dotnet add package Dimeskigj.Pulumi.Dokploy
pulumi package add github.com/dimeskigj/pulumi-dokploy dokploy
```

For Java, add this dependency to your Maven project's `pom.xml`:

```xml
<dependency>
  <groupId>net.dimeski.pulumi</groupId>
  <artifactId>dokploy</artifactId>
  <version>${DOKPLOY_VERSION}</version>
</dependency>
```

The last command adds the provider to a YAML program. Add the XML dependency to
your Java project's `pom.xml`; its artifact coordinate is
`net.dimeski.pulumi:dokploy:${DOKPLOY_VERSION}`. This XML declares the Java SDK
dependency rather than installing anything by itself.
The other commands install the TypeScript,
Python, Go, and .NET SDKs respectively.

## Configuration

Pulumi configuration keys use the `dokploy` namespace:

```bash
pulumi config set dokploy:endpoint https://dokploy.example.com
pulumi config set --secret dokploy:apiKey your-api-key
```

`dokploy:endpoint` is the URL of the Dokploy instance. `dokploy:apiKey` is a
secret and should always be stored with `--secret`; Pulumi encrypts secret
configuration in the stack state. Do not commit API keys, put them in source
code, or print them in logs.

As an alternative, set `DOKPLOY_ENDPOINT` and `DOKPLOY_API_KEY` in the
environment running Pulumi. Environment variables are useful in CI when loaded
from the CI system's secret store. Keep both values out of shell history and
logs.

This package is community-maintained; neither Dokploy nor Pulumi maintains it.
