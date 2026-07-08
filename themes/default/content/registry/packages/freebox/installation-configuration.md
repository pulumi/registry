---
# WARNING: this file was fetched from https://raw.githubusercontent.com/OlivierPaquien/pulumi-freebox/v0.3.11/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/OlivierPaquien/pulumi-freebox/blob/v0.3.11/docs/installation-configuration.md
title: Freebox Installation & Configuration
meta_desc: How to install and configure the Freebox Pulumi provider.
layout: package
---

## How installation works

This provider is distributed as a **plugin binary** (`pulumi-resource-freebox`), written in Go. Pulumi downloads and runs that binary regardless of the language you write your program in.

A **language SDK** (npm, PyPI, NuGet, Go module) is a separate generated package that gives you typed APIs in TypeScript, Python, C#, or Go. It is **not** the same thing as the plugin binary.

| Runtime | What you need |
|---------|---------------|
| **YAML** | Plugin binary only |
| **TypeScript / Python / C# / Go** | Plugin binary **+** language SDK |

## Plugin installation

Install the plugin from a [GitHub release](https://github.com/OlivierPaquien/pulumi-freebox/releases) or from a local build:

```bash
# From GitHub Releases
VERSION="$(gh release view --repo OlivierPaquien/pulumi-freebox --json tagName -q .tagName)"
VERSION="${VERSION#v}"
pulumi plugin install resource freebox "$VERSION" \
  --server github://api.github.com/OlivierPaquien/pulumi-freebox

# From a local build
VERSION="$(gh release view --repo OlivierPaquien/pulumi-freebox --json tagName -q .tagName)"
VERSION="${VERSION#v}"
go build -o bin/pulumi-resource-freebox .
pulumi plugin install resource freebox "$VERSION" --file ./bin/pulumi-resource-freebox
```

Pulumi installs the matching plugin automatically when you add a published language SDK to your project.

## Language SDKs

### YAML

No SDK required. Install the plugin, then use resource tokens directly (for example `freebox:fw:PortForwarding`). See the overview example in `_index.md`.

### TypeScript, Python, C#, Go

Install the SDK for your language from the public package registries:

| Language | Command |
|----------|---------|
| Node.js (TypeScript/JavaScript) | `npm install pulumi-freebox` |
| Python | `pip install pulumi-freebox` |
| .NET | `dotnet add package OlivierPaquien.Pulumi.Freebox` |
| Go | `go get github.com/OlivierPaquien/pulumi-freebox/sdk/go@sdk/go/vVERSION` |

Replace `VERSION` with a [release](https://github.com/OlivierPaquien/pulumi-freebox/releases) version (for example `sdk/go/v0.3.9`).

### Local SDK generation (alternative)

If you need an unreleased version or want SDK sources inside your project:

```bash
VERSION="$(gh release view --repo OlivierPaquien/pulumi-freebox --json tagName -q .tagName)"
VERSION="${VERSION#v}"
pulumi package add github.com/OlivierPaquien/pulumi-freebox@v$VERSION
```

This downloads the schema, installs the plugin if needed, and adds the SDK sources to your project. You can also point at a locally built plugin:

```bash
go build -o bin/pulumi-resource-freebox .
pulumi package add . --plugin bin/pulumi-resource-freebox
```

Alternatively, generate an SDK directory from `schema.json`:

```bash
pulumi package gen-sdk ./schema.json --language <nodejs|python|csharp|go>
```

## Freebox API authorization

Before using the provider, register an application with your Freebox and obtain an `app_id` and `token`. See the [Freebox API login documentation](https://dev.freebox.fr/sdk/login/).

Run the provider binary with the `authorize` subcommand to complete the pairing flow interactively:

```bash
go build -o bin/pulumi-resource-freebox .
./bin/pulumi-resource-freebox authorize
```

Follow the prompts, then approve the request on your Freebox (LCD or web UI). The command prints `FREEBOX_APP_ID` and `FREEBOX_TOKEN` for use in Pulumi config or environment variables.

Grant the application the permissions required by the resources you manage (for example firewall rules, DHCP, VPN, virtual machines).

## Configuration

The following configuration options are available:

| Option | Required | Description | Environment variable |
|--------|----------|-------------|----------------------|
| `freebox:appId` | Yes | Freebox API application ID | `FREEBOX_APP_ID` |
| `freebox:token` | Yes | Freebox API authentication token | `FREEBOX_TOKEN` |
| `freebox:endpoint` | No | Freebox URL (default: `http://mafreebox.freebox.fr`) | `FREEBOX_ENDPOINT` |
| `freebox:apiVersion` | No | API version to use (default: latest) | `FREEBOX_VERSION` |

### Set via Pulumi config

```bash
pulumi config set freebox:endpoint http://mafreebox.freebox.fr
pulumi config set freebox:appId your_app_id
pulumi config set freebox:token --secret your_token
# optional:
pulumi config set freebox:apiVersion latest
```

### Set via environment variables

```bash
export FREEBOX_ENDPOINT=http://mafreebox.freebox.fr
export FREEBOX_APP_ID=your_app_id
export FREEBOX_TOKEN=your_token
# optional:
export FREEBOX_VERSION=latest
```

## Debug logging

Pulumi does not always surface plugin logs on the console. The provider also writes debug output to a log file:

- **Default**: `/tmp/pulumi-freebox-provider.log`, or `$HOME/.pulumi/pulumi-freebox-provider.log` if `/tmp` is not writable. On Windows: `%TEMP%\pulumi-freebox-provider.log`.
- **Custom**: set `FREEBOX_DEBUG_LOG=/path/to/file` (when Pulumi passes environment variables to the plugin).

After rebuilding the plugin locally, reinstall it so Pulumi uses the new binary:

```bash
VERSION="$(gh release view --repo OlivierPaquien/pulumi-freebox --json tagName -q .tagName)"
VERSION="${VERSION#v}"
pulumi plugin install resource freebox "$VERSION" --file ./bin/pulumi-resource-freebox --reinstall
FREEBOX_DEBUG_LOG=$HOME/freebox-provider.log pulumi up
```

The log file contains a `[freebox] log file: ...` line at startup confirming the output path.
