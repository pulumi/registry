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

Wrap them in the site's language chooser so a reader sees only the language they use. The chooser is a pair of Hugo shortcodes — `chooser` with angle-bracket delimiters, `choosable` with percent delimiters — and the language keys are `typescript`, `python`, `go`, `csharp`, `java`, `yaml` and `hcl`. List in the `chooser` tag only the languages your package actually supports:

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

```bash
pulumi package add your-package
```

{{% /choosable %}}
{{< /chooser >}}
````

Notes:

- **Mind the delimiters.** `{{< chooser >}}` and `{{< /chooser >}}` use angle brackets; `{{% choosable %}}` and `{{% /choosable %}}` use percent signs. Mixing them silently breaks the rendered page, and CI checks for it (`make lint-markdown` runs a malformed-delimiter check over published content).
- **YAML and HCL use `pulumi package add`.** They consume the package directly rather than through a per-language SDK.
- **If you publish no SDKs at all**, `pulumi package add` *is* your installation section — one command, and no chooser needed. Many bridged providers currently title this section "Generate Provider"; `## Installation` is the preferred heading.
- **Java has no one-line install command**, so give the Maven and Gradle dependency coordinates, as above.
- **Links to package feeds** (npm, PyPI, NuGet, pkg.go.dev, Maven Central) are accepted, and many existing packages use them instead. Prefer commands; add links alongside them if you like.
- **If your provider ships a plugin binary** users must install separately, give the `pulumi plugin install resource <name> <version> --server <url>` command here too.

### `## Example Usage`

Provide a **complete, minimal Pulumi program** — one that declares a single resource from your provider — **in every language you support**. It needs the imports and enough surrounding code to actually run. A bare resource declaration on its own is not enough.

Alongside the program, give the **full configuration** needed to run it, as `pulumi config set` commands:

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
