---
title: Honeycomb
meta_desc: Install and configure the Pulumi Honeycomb provider — SDK generation, API key examples for the v1 and v2 APIs, and every provider configuration option.
layout: package
---

[Honeycomb](https://honeycomb.io) provides observability for high-performance engineering teams, so they can quickly understand what their code does in the hands of real users in unpredictable and highly complex cloud environments. The Pulumi Honeycomb provider manages Honeycomb datasets, columns, queries, triggers, SLOs, markers and boards.

This provider is [dynamically bridged](/registry/packages/terraform-provider/) from the [honeycombio Terraform provider](https://github.com/honeycombio/terraform-provider-honeycombio) and is published by Honeycomb, not by Pulumi. You will need a Honeycomb account to use it; a free account is enough to get started.

{{< chooser language "typescript,python,go,csharp,java,yaml,hcl" / >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as honeycombio from "@pulumi/honeycombio";

const config = new pulumi.Config();
const dataset = config.require("dataset");

const hello = new honeycombio.Marker("hello", {
    message: "Hello world!",
    dataset: dataset,
});

export const markerId = hello.id;
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_honeycombio as honeycombio

config = pulumi.Config()
dataset = config.require("dataset")

hello = honeycombio.Marker("hello",
    message="Hello world!",
    dataset=dataset)

pulumi.export("markerId", hello.id)
```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi-terraform-provider/sdks/go/honeycombio/honeycombio"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		dataset := cfg.Require("dataset")

		hello, err := honeycombio.NewMarker(ctx, "hello", &honeycombio.MarkerArgs{
			Message: pulumi.String("Hello world!"),
			Dataset: pulumi.String(dataset),
		})
		if err != nil {
			return err
		}

		ctx.Export("markerId", hello.ID())
		return nil
	})
}
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using Pulumi;
using Honeycombio = Pulumi.Honeycombio;

return await Deployment.RunAsync(() =>
{
    var config = new Config();
    var dataset = config.Require("dataset");

    var hello = new Honeycombio.Marker("hello", new()
    {
        Message = "Hello world!",
        Dataset = dataset,
    });

    return new Dictionary<string, object?>
    {
        ["markerId"] = hello.Id,
    };
});
```

{{% /choosable %}}

{{% choosable language java %}}

```java
import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.honeycombio.Marker;
import com.pulumi.honeycombio.MarkerArgs;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    private static void stack(Context ctx) {
        final var config = ctx.config();
        final var dataset = config.require("dataset");

        var hello = new Marker("hello", MarkerArgs.builder()
            .message("Hello world!")
            .dataset(dataset)
            .build());

        ctx.export("markerId", hello.id());
    }
}
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: honeycomb-example
runtime: yaml
config:
  dataset:
    type: string
resources:
  hello:
    type: honeycombio:Marker
    properties:
      message: Hello world!
      dataset: ${dataset}
outputs:
  markerId: ${hello.id}
```

{{% /choosable %}}

{{% choosable language hcl %}}

```hcl
variable "dataset" {
  type = string
}

resource "honeycombio_marker" "hello" {
  message = "Hello world!"
  dataset = var.dataset
}

output "marker_id" {
  value = honeycombio_marker.hello.id
}
```

{{% /choosable %}}

## Installation

Honeycomb does not publish a pre-built Pulumi SDK for this provider, so there is nothing to install from npm, PyPI or NuGet. Use [`pulumi package add`](/docs/iac/cli/commands/pulumi_package_add/) instead: it downloads the provider plugin, generates a [local SDK](/docs/iac/guides/building-extending/packages/local-sdks/) from the bridged Terraform provider, and adds the package to your `Pulumi.yaml`. This requires Pulumi v3.147.0 or later:

{{< chooser language "typescript,python,go,csharp,java,yaml,hcl" / >}}

{{% choosable language typescript %}}

```bash
pulumi package add terraform-provider honeycombio/honeycombio
```

The CLI generates a Node.js SDK under `sdks/honeycombio`, adds it to your `package.json`, and installs it. Import it with `import * as honeycombio from "@pulumi/honeycombio";`.

{{% /choosable %}}

{{% choosable language python %}}

```bash
pulumi package add terraform-provider honeycombio/honeycombio
```

The CLI generates a Python SDK under `sdks/honeycombio`, adds it to your requirements, and installs it. Import it with `import pulumi_honeycombio as honeycombio`.

{{% /choosable %}}

{{% choosable language go %}}

```bash
pulumi package add terraform-provider honeycombio/honeycombio
```

The CLI generates a Go SDK under `sdks/honeycombio` and adds the `replace` directive to your `go.mod`. Import it with `import "github.com/pulumi/pulumi-terraform-provider/sdks/go/honeycombio/honeycombio"`.

{{% /choosable %}}

{{% choosable language csharp %}}

```bash
pulumi package add terraform-provider honeycombio/honeycombio
```

The CLI generates a .NET SDK under `sdks/honeycombio` and adds the project reference. Use it with `using Pulumi.Honeycombio;`.

{{% /choosable %}}

{{% choosable language java %}}

```bash
pulumi package add terraform-provider honeycombio/honeycombio
```

The CLI generates a Java SDK under `sdks/honeycombio` and prints the dependencies to add to your `pom.xml`. Import it with `import com.pulumi.honeycombio.*;`.

{{% /choosable %}}

{{% choosable language yaml %}}

```bash
pulumi package add terraform-provider honeycombio/honeycombio
```

The command adds the package to the `packages` block of your `Pulumi.yaml`. Reference resources by their type token, for example `honeycombio:Marker`.

{{% /choosable %}}

{{% choosable language hcl %}}

```bash
pulumi package add terraform-provider honeycombio/honeycombio
```

The command adds the package to the `packages` block of your `Pulumi.yaml`, making the `honeycombio_*` resource types available to your converted HCL.

{{% /choosable %}}

After adding a local SDK, run [`pulumi install`](/docs/iac/cli/commands/pulumi_install/) to complete the installation. Commit `Pulumi.yaml` to source control so that your teammates get the same SDK when they run it.

## Configuration

Provider settings can come from any of three places:

- A [Pulumi ESC](/docs/pulumi-cloud/esc/) environment, under `pulumiConfig`.
- [Stack configuration](/docs/iac/concepts/config/), set with `pulumi config set` and stored in `Pulumi.<stack-name>.yaml`.
- An explicit `honeycombio.Provider` instance, constructed in your program and passed to the resources that should use it.

See [Configuration](/docs/concepts/config/) for how these sources combine. Some options can also be read from environment variables, noted in their descriptions in the Reference.

The provider talks to both the Honeycomb v1 and v2 APIs, which take different key types. **At least one of `apiKey`, or the `apiKeyId` and `apiKeySecret` pair, must be configured.** A single provider instance can hold both. See [Best Practices for API Keys](https://docs.honeycomb.io/get-started/best-practices/api-keys/).

### Examples

Every example below is shown two ways: as a [Pulumi ESC](/docs/pulumi-cloud/esc/) environment, and as [stack configuration](/docs/iac/concepts/config/) — the `Pulumi.<stack-name>.yaml` file alongside your project. Pick whichever you use; the choice follows you down the page and across providers.

#### v1 APIs — Configuration key

Most resources and functions call the v1 APIs, which require a Configuration key. Its permissions are managed in *Environment settings*.

{{< chooser configsource "esc,stack" / >}}

{{% choosable configsource esc %}}

Set it with the CLI:

```bash
$ pulumi env set --secret myorg/honeycomb-dev pulumiConfig.honeycombio:apiKey hcaik_xxxxxxxxxxxxxxxxxxxx
```

That produces the following environment definition in `environments/honeycomb-dev.yaml`:

```yaml
values:
  pulumiConfig:
    honeycombio:apiKey:
      fn::secret: hcaik_xxxxxxxxxxxxxxxxxxxx
```

{{% /choosable %}}

{{% choosable configsource stack %}}

Set it with the CLI:

```bash
$ pulumi config set --secret honeycombio:apiKey hcaik_xxxxxxxxxxxxxxxxxxxx
```

That writes the following into `Pulumi.dev.yaml`:

```yaml
config:
  honeycombio:apiKey:
    secure: AAABAOWyM7t3q1u...KpR8Nx1yVCjKPB2nJWg==
```

{{% /choosable %}}

The key can also be supplied via the `HONEYCOMB_API_KEY` or `HONEYCOMBIO_APIKEY` environment variables; `HONEYCOMB_API_KEY` takes priority.

#### v2 APIs — Management key pair

Resources and functions that call the v2 APIs need a Management key, whose permissions are managed in *Team settings*. Each such resource documents the scope it requires.

{{< chooser configsource "esc,stack" / >}}

{{% choosable configsource esc %}}

Set it with the CLI:

```bash
$ pulumi env set myorg/honeycomb-dev pulumiConfig.honeycombio:apiKeyId hcaik_xxxxxxxxxxxxxxxxxxxx
$ pulumi env set --secret myorg/honeycomb-dev pulumiConfig.honeycombio:apiKeySecret xxxxxxxxxxxxxxxxxxxx
```

That produces the following environment definition in `environments/honeycomb-dev.yaml`:

```yaml
values:
  pulumiConfig:
    honeycombio:apiKeyId: hcaik_xxxxxxxxxxxxxxxxxxxx
    honeycombio:apiKeySecret:
      fn::secret: xxxxxxxxxxxxxxxxxxxx
```

{{% /choosable %}}

{{% choosable configsource stack %}}

Set it with the CLI:

```bash
$ pulumi config set honeycombio:apiKeyId hcaik_xxxxxxxxxxxxxxxxxxxx
$ pulumi config set --secret honeycombio:apiKeySecret xxxxxxxxxxxxxxxxxxxx
```

That writes the following into `Pulumi.dev.yaml`:

```yaml
config:
  honeycombio:apiKeyId: hcaik_xxxxxxxxxxxxxxxxxxxx
  honeycombio:apiKeySecret:
    secure: AAABAJ4nQm1xPz9...8sVdK2eLhTgYbXc5Rq==
```

{{% /choosable %}}

The pair can also be supplied via the `HONEYCOMB_KEY_ID` and `HONEYCOMB_KEY_SECRET` environment variables.

#### Honeycomb EU

Honeycomb EU customers must override the default API host.

{{< chooser configsource "esc,stack" / >}}

{{% choosable configsource esc %}}

Set it with the CLI:

```bash
$ pulumi env set myorg/honeycomb-dev pulumiConfig.honeycombio:apiUrl https://api.eu1.honeycomb.io
```

That produces the following environment definition in `environments/honeycomb-dev.yaml`:

```yaml
values:
  pulumiConfig:
    honeycombio:apiUrl: https://api.eu1.honeycomb.io
```

{{% /choosable %}}

{{% choosable configsource stack %}}

Set it with the CLI:

```bash
$ pulumi config set honeycombio:apiUrl https://api.eu1.honeycomb.io
```

That writes the following into `Pulumi.dev.yaml`:

```yaml
config:
  honeycombio:apiUrl: https://api.eu1.honeycomb.io
```

{{% /choosable %}}

This can also be set with the `HONEYCOMB_API_ENDPOINT` environment variable.

### Reference

Every configuration option accepted by the provider. Project any of them from an ESC environment under `pulumiConfig`, set them with `pulumi config set`, or pass them to an explicit [`honeycombio.Provider`](/registry/packages/honeycombio/api-docs/provider/). Nested objects are shown flattened, using [structured configuration](/docs/concepts/config#structured-configuration) paths.

`apiKey`, `apiKeyId` and `apiKeySecret` are marked **conditional** rather than required or optional: you need `apiKey`, *or* the `apiKeyId` and `apiKeySecret` pair, so no one of the three is required on its own.

<div class="api-params">
<div class="api-param">
<h3 class="api-param-name" id="apikey" data-link-title="apiKey">apiKey<a class="api-param-anchor" href="#apikey" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required is-unknown" title="Set either apiKey, or both apiKeyId and apiKeySecret">conditional</span>
<span class="api-param-badge api-param-secret">secret</span>
</div>
<div class="api-param-description">
<p>The Honeycomb API key to use. It can also be set via the <code>HONEYCOMB_API_KEY</code> or <code>HONEYCOMBIO_APIKEY</code> environment variables.</p>
<p class="api-param-meta">Config key: <code>honeycombio:apiKey</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> --secret myorg/myproject/dev pulumiConfig.honeycombio:apiKey hcaik_xxxxxxxxxxxxxxxxxxxx
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> --secret honeycombio:apiKey hcaik_xxxxxxxxxxxxxxxxxxxx
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="apikeyid" data-link-title="apiKeyId">apiKeyId<a class="api-param-anchor" href="#apikeyid" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required is-unknown" title="Set either apiKey, or both apiKeyId and apiKeySecret">conditional</span>
</div>
<div class="api-param-description">
<p>The ID portion of the Honeycomb Management API key to use. It can also be set via the <code>HONEYCOMB_KEY_ID</code> environment variable.</p>
<p class="api-param-meta">Config key: <code>honeycombio:apiKeyId</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.honeycombio:apiKeyId hcaik_xxxxxxxxxxxxxxxxxxxx
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> honeycombio:apiKeyId hcaik_xxxxxxxxxxxxxxxxxxxx
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="apikeysecret" data-link-title="apiKeySecret">apiKeySecret<a class="api-param-anchor" href="#apikeysecret" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required is-unknown" title="Set either apiKey, or both apiKeyId and apiKeySecret">conditional</span>
<span class="api-param-badge api-param-secret">secret</span>
</div>
<div class="api-param-description">
<p>The secret portion of the Honeycomb Management API key to use. It can also be set via the <code>HONEYCOMB_KEY_SECRET</code> environment variable.</p>
<p class="api-param-meta">Config key: <code>honeycombio:apiKeySecret</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> --secret myorg/myproject/dev pulumiConfig.honeycombio:apiKeySecret xxxxxxxxxxxxxxxxxxxx
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> --secret honeycombio:apiKeySecret xxxxxxxxxxxxxxxxxxxx
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="apiurl" data-link-title="apiUrl">apiUrl<a class="api-param-anchor" href="#apiurl" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>Override the URL of the Honeycomb API. Defaults to <code>https://api.honeycomb.io</code>. It can also be set via the <code>HONEYCOMB_API_ENDPOINT</code> environment variable.</p>
<p class="api-param-meta">Config key: <code>honeycombio:apiUrl</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.honeycombio:apiUrl https://api.eu1.honeycomb.io
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> honeycombio:apiUrl https://api.eu1.honeycomb.io
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="debug" data-link-title="debug">debug<a class="api-param-anchor" href="#debug" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">boolean</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>Enable the API client's debug logs. By default, a <code>TF_LOG</code> setting of debug or higher will enable this.</p>
<p class="api-param-meta">Config key: <code>honeycombio:debug</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.honeycombio:debug true
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> honeycombio:debug true
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="features-column-importonconflict" data-link-title="features.column.importOnConflict"><span class="api-param-nested-indicator">&#8627;</span>features.column.importOnConflict<a class="api-param-anchor" href="#features-column-importonconflict" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">boolean</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>This changes the creation behavior of the column resource to import an existing column if it already exists, rather than erroring out.</p>
<p class="api-param-meta">Config key: <code>honeycombio:features.column.importOnConflict</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.honeycombio:features.column.importOnConflict true
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> --path honeycombio:features.column.importOnConflict true
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="features-dataset-importonconflict" data-link-title="features.dataset.importOnConflict"><span class="api-param-nested-indicator">&#8627;</span>features.dataset.importOnConflict<a class="api-param-anchor" href="#features-dataset-importonconflict" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">boolean</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>This changes the creation behavior of the dataset resource to import an existing dataset if it already exists, rather than erroring out.</p>
<p class="api-param-meta">Config key: <code>honeycombio:features.dataset.importOnConflict</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.honeycombio:features.dataset.importOnConflict true
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> --path honeycombio:features.dataset.importOnConflict true
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="features-intelligence-enabled" data-link-title="features.intelligence.enabled"><span class="api-param-nested-indicator">&#8627;</span>features.intelligence.enabled<a class="api-param-anchor" href="#features-intelligence-enabled" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">boolean</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>Set to true to enable intelligence features such as autoInvestigate on triggers and burn alerts. Requires Honeycomb Intelligence to be enabled for your team in the Honeycomb UI.</p>
<p class="api-param-meta">Config key: <code>honeycombio:features.intelligence.enabled</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.honeycombio:features.intelligence.enabled true
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> --path honeycombio:features.intelligence.enabled true
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
</div>

## Datasets

Several resources in this provider accept a `dataset` or `datasets` argument naming the Honeycomb Dataset the resource belongs to. These include queries, triggers, SLOs, markers, columns and boards.

The argument is always a Dataset **slug** — not a Dataset name or ID. Slugs appear in the URL of the dataset in the Honeycomb UI, and in the `slug` field of the [Dataset API](https://api-docs.honeycomb.io/api/datasets/createdataset#datasets/createdataset/t=response&c=200&path=slug).

## Feature flags

The `features` configuration object changes the behavior of certain resources, letting you opt into non-default handling while preserving Pulumi-idiomatic defaults. The individual flags are listed in [Reference](#reference) above as `features.column.importOnConflict`, `features.dataset.importOnConflict` and `features.intelligence.enabled`.

{{% notes type="warning" %}}
`features.column.importOnConflict` is potentially dangerous: if the column type changes on update — switching from `string` to `boolean`, for example — it can cause data loss. Use it with caution.

[Honeycomb Intelligence](https://docs.honeycomb.io/security-compliance/honeycomb-intelligence) must be enabled for your team in the Honeycomb UI before setting `features.intelligence.enabled`. The flag tells Pulumi your team has it enabled and unlocks related attributes such as `autoInvestigate` on triggers and burn alerts.
{{% /notes %}}
