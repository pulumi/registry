# The Overview page (`docs/_index.md`)

`docs/_index.md` is the only documentation page the Pulumi Registry requires of a package. It renders as the package's index page, and it is where a reader decides whether your package does what they need and how to start using it.

You author this file in **your provider's repository**, at `docs/_index.md`. `resourcedocsgen` fetches it from your release tag and publishes it; you never commit it to `pulumi/registry` yourself.

This document is the standard we hold Overview pages to. It applies to every package, whether you publish per-language SDKs or your users install with `pulumi package add`.

---

## Front matter

The file **must begin** with a YAML front-matter block. Generation fails with `expected file ... to start with YAML front-matter` if it does not — this is the single most common reason a package's docs fail to publish.

```yaml
---
title: Logfire
meta_desc: Use the Pulumi Logfire provider to manage projects, alerts, channels, dashboards, and API tokens.
layout: package
---
```

| Key | Notes |
|---|---|
| `title` | The package display name. Should match `displayName` in your `schema.json`. Rendered as the page's heading. |
| `meta_desc` | One sentence, used as the page's meta description. Include the package name. |
| `layout` | Use `package`. |

`resourcedocsgen` prepends a `# WARNING:` comment and an `edit_url:` key when it publishes; you do not write those.

---

## Page structure

Use `##` (h2) for every top-level section, and `###` for anything nested beneath one. **Do not use `#`** — the page heading comes from front-matter `title`, and a second h1 in the body competes with it.

The sections below are the expected shape of the page, in order.

### The overview itself — no heading

Open with prose directly beneath the front matter, with **no heading above it**. The page already carries the package name as its title; an `## Overview` heading underneath it is redundant.

State what the provider lets the reader do, and link the product or service:

```markdown
The Logfire provider for Pulumi lets you manage [Logfire](https://pydantic.dev/logfire) resources — projects, alerts, channels, dashboards, and API tokens — as part of your Pulumi programs.
```

Keep it to a short paragraph. The detail goes in the sections below.

### `## Installation`

Give the installation command for **each language the package supports**. A command is more useful than a link, because a reader can copy it and run it.

Wrap them in the site's language chooser so a reader sees only the language they use. The chooser is a pair of Hugo shortcodes — `chooser` with angle-bracket delimiters, `choosable` with percent delimiters — and the language keys are `typescript`, `python`, `go`, `csharp`, `java`, `yaml` and `hcl`. List all seven — a language you publish no SDK for still gets a tab, carrying `pulumi package add` (see the notes below):

````markdown
## Installation

{{< chooser language "typescript,python,go,csharp,java,yaml,hcl" >}}
{{% choosable language typescript %}}

```bash
npm install @your-org/your-package
```

{{% /choosable %}}
{{% choosable language python %}}

```bash
pip install your_org_your_package
```

{{% /choosable %}}
{{% choosable language go %}}

```bash
go get github.com/your-org/pulumi-your-package/sdk/go/yourpackage
```

{{% /choosable %}}
{{% choosable language csharp %}}

```bash
dotnet add package YourOrg.YourPackage
```

{{% /choosable %}}
{{% choosable language java %}}

Maven:

```xml
<dependency>
    <groupId>com.yourorg</groupId>
    <artifactId>your-package</artifactId>
    <version>1.2.3</version>
</dependency>
```

Gradle:

```groovy
implementation 'com.yourorg:your-package:1.2.3'
```

{{% /choosable %}}
{{% choosable language yaml %}}

```bash
pulumi package add your-package
```

{{% /choosable %}}
{{% choosable language hcl %}}

Declare the provider in your program, then run `pulumi install`:

```hcl
terraform {
  required_providers {
    your-package = {
      source  = "pulumi/your-package"
      version = "1.2.3"
    }
  }
}
```

```bash
pulumi install
```

{{% /choosable %}}
{{< /chooser >}}
````

`resourcedocsgen gen-install --schemaFile <your schema.json>` (experimental) generates this whole block from your schema. It cannot tell which SDKs you actually publish — that is not in the schema — so pass `--languages` to correct it, or `--languages none` if you publish none. Languages you leave out still get a tab showing `pulumi package add`.

Notes:

- **Mind the delimiters.** `{{< chooser >}}` and `{{< /chooser >}}` use angle brackets; `{{% choosable %}}` and `{{% /choosable %}}` use percent signs. Mixing them silently breaks the rendered page, and CI checks for it (`make lint-markdown` runs a malformed-delimiter check over published content).
- **Give every language a tab, not only the ones you publish an SDK for.** A reader who picks C# and finds an empty panel has been told nothing. For a language with no published SDK, `pulumi package add <name>` generates one locally: it takes the SDK language from the `runtime` in the reader's `Pulumi.yaml`, records the package there, and prints the import line — so the command reads the same for every language. Run outside a project it needs `--language nodejs|python|go|dotnet|java`. (`pulumi package gen-sdk` is the lower-level form, writing an SDK to `./sdk` without recording it; show `add` unless you mean the difference.)
- **YAML uses `pulumi package add`** for the same reason: it consumes the package directly rather than through a per-language SDK.
- **HCL does not use `pulumi package add` at all.** An HCL program names the provider in its own `required_providers` block and `pulumi install` fetches it. A source prefixed with `pulumi/` resolves to the native Pulumi provider and must be pinned to an **exact** semver version, not a constraint; any other source is bridged from its Terraform provider, so a package that wraps one (`pulumi package add terraform-provider <ns>/<name>` elsewhere) uses that same upstream address here. See the [Pulumi HCL reference](https://www.pulumi.com/docs/iac/languages-sdks/hcl/hcl-language-reference/).
- **If you publish no SDKs at all**, every tab still earns its place: the SDK languages and YAML share one `pulumi package add` command, and HCL differs. Many bridged providers currently title this section "Generate Provider"; `## Installation` is the preferred heading.
- **Local SDK generation needs a `version` in your schema.** Without one, `pulumi package add` fails outright with `version must be provided when package supports packing`.
- **Java has no one-line install command**, so give the Maven and Gradle dependency coordinates, as above.
- **Links to package feeds** (npm, PyPI, NuGet, pkg.go.dev, Maven Central) are accepted, and many existing packages use them instead. Prefer commands; add links alongside them if you like.
- **You almost certainly do not need a `pulumi plugin install` command.** If your schema sets `pluginDownloadURL`, that value is compiled into every SDK you publish and the engine downloads the plugin binary on first use. Give the `pulumi plugin install resource <name> <version> --server <url>` command only if you know a reader must run it by hand.

### `## Example Usage`

Provide a **complete, minimal Pulumi program** — one that declares a single resource from your provider — **in every language you support**. It needs the imports and enough surrounding code to actually run. A bare resource declaration on its own is not enough.

Wrap the programs in the same language chooser the Installation section uses, so a reader sees only their own language rather than scrolling past five copies of the same program. Use the same language keys, and list only the languages you support:

````markdown
## Example Usage

{{< chooser language "typescript,python,go" >}}
{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as yourpackage from "@your-org/your-package";

const example = new yourpackage.Widget("example", {size: "small"});

export const widgetId = example.id;
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
import your_org_your_package as yourpackage

example = yourpackage.Widget("example", size="small")

pulumi.export("widgetId", example.id)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"github.com/your-org/pulumi-your-package/sdk/go/yourpackage"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := yourpackage.NewWidget(ctx, "example", &yourpackage.WidgetArgs{
			Size: pulumi.String("small"),
		})
		if err != nil {
			return err
		}
		ctx.Export("widgetId", example.ID())
		return nil
	})
}
```

{{% /choosable %}}
{{< /chooser >}}
````

The delimiter rule from the Installation section applies here too: `{{< chooser >}}` takes angle brackets, `{{% choosable %}}` takes percent signs.

Alongside the programs, give the **full configuration** needed to run them, as `pulumi config set` commands:

```bash
pulumi config set --secret your-package:apiToken <your-token>
pulumi config set your-package:region us-east-1
```

Optionally, show the same configuration as a [Pulumi ESC](https://www.pulumi.com/docs/esc/) environment. If you do, link to the ESC documentation so a reader unfamiliar with it can follow along.

### `## Configuration`

This section has one required part and one optional part.

#### Configuration parameters reference (required)

Document every configuration parameter the provider accepts. For each one give:

- **Name** — the bare option name, e.g. `apiToken`. Do not prefix it with the provider: the whole table is about your provider, so repeating `your-package:` on every row is noise. The prefix belongs only in `pulumi config set` commands, where it is required to disambiguate.
- **Required?** — whether the provider works without it
- **Secret?** — whether it should be set with `pulumi config set --secret`
- **Description** — what it does and what a valid value looks like

A bullet list is the preferred shape. It is what most existing packages use, and it keeps long descriptions readable where a table column squeezes them:

```markdown
- `apiToken` (Required, Secret) — The API token used to authenticate. May also be set with the `YOURPACKAGE_API_TOKEN` environment variable.
- `region` (Optional) — The region to operate in, e.g. `us-east-1`. Defaults to `us-east-1`.
```

A table carrying the same four fields is equally acceptable, and gives requiredness and secrecy their own columns to scan:

```markdown
| Name | Required | Secret | Description |
|---|---|---|---|
| `apiToken` | Yes | Yes | The API token used to authenticate. May also be set with the `YOURPACKAGE_API_TOKEN` environment variable. |
| `region` | No | No | The region to operate in, e.g. `us-east-1`. Defaults to `us-east-1`. |
```

`resourcedocsgen gen-config --schemaFile <your schema.json>` (experimental) generates either form from your schema's `config` block, as a starting point to edit. It emits the list by default; `--style table` selects the second.

Two things routinely get missed here, because they are not derivable from your schema. Put them in the parameter descriptions:

- **Environment variables.** A parameter's environment-variable fallback often is not in the Pulumi (or, for a bridged provider, the Terraform) schema at all — it is read by the vendor SDK at a layer beneath it. If a parameter can be supplied by an environment variable, say so explicitly and name the variable.
- **Mutually exclusive options.** If setting one parameter forbids or overrides another, or if a group of parameters must be supplied together, say so in the descriptions of every parameter involved. Nothing else in the docs will surface that constraint.

#### Configuration examples (optional)

Worked examples as `pulumi config set` commands, with explanatory text and links where a reader needs background. Add these if there is **more than one way to authenticate** — CLI login, a service principal, OIDC. Show each one as its own example instead of describing them in prose.

### Optional sections

Anything else useful goes after the sections above. What existing packages add, most common first: `## Authentication` (when there is enough of it to want its own section instead of living under Configuration), `## Environment Variables`, `## Resources`, `## Requirements` or `## Prerequisites`, `## Troubleshooting`, `## Migration` notes for a major version bump, and links to further reading.

---

## When the page gets long

If the installation and configuration material outgrows the Overview — several authentication methods, a long configuration table — you may move it into a second file, `docs/installation-configuration.md`, which renders as a separate Installation & Configuration page. The large cloud providers do this: the [AWS install page](https://www.pulumi.com/registry/packages/aws/installation-configuration/) runs to several thousand words covering shared credentials files, EC2 instance metadata, OIDC and ESC.

This is about volume, not an extra requirement. Most packages do not need it, and a single Overview page is a complete submission.
