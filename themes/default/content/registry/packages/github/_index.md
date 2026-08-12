---
title: GitHub
meta_desc: Install and configure the Pulumi GitHub provider — SDK installation for every language, CLI, personal access token and GitHub App credentials, and every configuration option.
layout: package
aliases:
    - /registry/packages/github/installation-configuration/
---

The Pulumi GitHub provider manages [GitHub](https://github.com/) resources — repositories, teams, organization membership, branch protection rules, actions secrets and more. It must be configured with credentials before it can be used.

{{< chooser language "typescript,python,go,csharp,java,yaml,hcl" / >}}

{{% choosable language typescript %}}

```typescript
import * as github from "@pulumi/github";

const repo = new github.Repository("example", {
    name: "pulumi-example",
    description: "Managed by Pulumi",
    visibility: "private",
});

export const repoUrl = repo.htmlUrl;
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_github as github

repo = github.Repository("example",
    name="pulumi-example",
    description="Managed by Pulumi",
    visibility="private")

pulumi.export("repoUrl", repo.html_url)
```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		repo, err := github.NewRepository(ctx, "example", &github.RepositoryArgs{
			Name:        pulumi.String("pulumi-example"),
			Description: pulumi.String("Managed by Pulumi"),
			Visibility:  pulumi.String("private"),
		})
		if err != nil {
			return err
		}

		ctx.Export("repoUrl", repo.HtmlUrl)
		return nil
	})
}
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using Pulumi;
using Github = Pulumi.Github;

return await Deployment.RunAsync(() =>
{
    var repo = new Github.Repository("example", new()
    {
        Name = "pulumi-example",
        Description = "Managed by Pulumi",
        Visibility = "private",
    });

    return new Dictionary<string, object?>
    {
        ["repoUrl"] = repo.HtmlUrl,
    };
});
```

{{% /choosable %}}

{{% choosable language java %}}

```java
import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.github.Repository;
import com.pulumi.github.RepositoryArgs;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    private static void stack(Context ctx) {
        var repo = new Repository("example", RepositoryArgs.builder()
            .name("pulumi-example")
            .description("Managed by Pulumi")
            .visibility("private")
            .build());

        ctx.export("repoUrl", repo.htmlUrl());
    }
}
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: github-example
runtime: yaml
resources:
  example:
    type: github:Repository
    properties:
      name: pulumi-example
      description: Managed by Pulumi
      visibility: private
outputs:
  repoUrl: ${example.htmlUrl}
```

{{% /choosable %}}

{{% choosable language hcl %}}

```hcl
resource "github_repository" "example" {
  name        = "pulumi-example"
  description = "Managed by Pulumi"
  visibility  = "private"
}

output "repo_url" {
  value = github_repository.example.html_url
}
```

{{% /choosable %}}

## Installation

The GitHub provider publishes an SDK for every language. Add it to an existing Pulumi project with your language's package manager; YAML and HCL projects use [`pulumi package add`](/docs/iac/cli/commands/pulumi_package_add/) instead:

{{< chooser language "typescript,python,go,csharp,java,yaml,hcl" / >}}

{{% choosable language typescript %}}

```bash
npm install @pulumi/github
```

{{% /choosable %}}

{{% choosable language python %}}

```bash
pip install pulumi-github
```

{{% /choosable %}}

{{% choosable language go %}}

```bash
go get github.com/pulumi/pulumi-github/sdk/v6
```

Import it in your program with `import "github.com/pulumi/pulumi-github/sdk/v6/go/github"`.

{{% /choosable %}}

{{% choosable language csharp %}}

```bash
dotnet add package Pulumi.Github
```

{{% /choosable %}}

{{% choosable language java %}}

Add the dependency to your build file. For Maven, in `pom.xml`:

```xml
<dependency>
    <groupId>com.pulumi</groupId>
    <artifactId>github</artifactId>
    <version>6.14.1</version>
</dependency>
```

For Gradle, in `build.gradle`:

```groovy
implementation 'com.pulumi:github:6.14.1'
```

{{% /choosable %}}

{{% choosable language yaml %}}

```bash
pulumi package add github
```

Reference resources by their type token, for example `github:Repository`. Omit the version to add the latest, or append `@6.14.1` to pin one.

{{% /choosable %}}

{{% choosable language hcl %}}

```bash
pulumi package add github
```

This makes the `github_*` resource types available to your converted HCL. Omit the version to add the latest, or append `@6.14.1` to pin one.

{{% /choosable %}}

To install everything a project already declares — after cloning it from source control, for example — run [`pulumi install`](/docs/iac/cli/commands/pulumi_install/).

## Configuration

Provider settings can come from any of three places:

- A [Pulumi ESC](/docs/pulumi-cloud/esc/) environment, under `pulumiConfig`.
- [Stack configuration](/docs/iac/concepts/config/), set with `pulumi config set` and stored in `Pulumi.<stack-name>.yaml`.
- An explicit `github.Provider` instance, constructed in your program and passed to the resources that should use it.

See [Configuration](/docs/concepts/config/) for how these sources combine. Some options can also be read from environment variables, noted in their descriptions in the Reference.

The provider offers three ways to authenticate. Whichever you choose, set `owner` to the organization or user account you want to manage — without it, and with a token available, the provider falls back to the account that owns the token.

### Examples

Every example below is shown two ways: as a [Pulumi ESC](/docs/pulumi-cloud/esc/) environment, and as [stack configuration](/docs/iac/concepts/config/) — the `Pulumi.<stack-name>.yaml` file alongside your project. Pick whichever you use; the choice follows you down the page and across providers.

#### GitHub CLI

If you are already signed in with the [GitHub CLI](https://cli.github.com/), the provider picks up the token issued by [`gh auth login`](https://cli.github.com/manual/gh_auth_login) automatically — there are no credentials to configure, only the account to manage.

```bash
$ gh auth login
```

{{< chooser configsource "esc,stack" / >}}

{{% choosable configsource esc %}}

Set it with the CLI:

```bash
$ pulumi env set myorg/github-dev pulumiConfig.github:owner my-org
```

That produces the following environment definition in `environments/github-dev.yaml`:

```yaml
values:
  pulumiConfig:
    github:owner: my-org
```

{{% /choosable %}}

{{% choosable configsource stack %}}

Set it with the CLI:

```bash
$ pulumi config set github:owner my-org
```

That writes the following into `Pulumi.dev.yaml`:

```yaml
config:
  github:owner: my-org
```

{{% /choosable %}}

Set `GH_PATH` if the provider cannot locate the `gh` executable on its own — for example under a cygwin terminal.

#### Personal access token

The most common choice for CI. Create a [personal access token](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens) with the scopes your program needs.

{{< chooser configsource "esc,stack" / >}}

{{% choosable configsource esc %}}

Set it with the CLI:

```bash
$ pulumi env set myorg/github-dev pulumiConfig.github:owner my-org
$ pulumi env set --secret myorg/github-dev pulumiConfig.github:token ghp_xxxxxxxxxxxxxxxxxxxx
```

That produces the following environment definition in `environments/github-dev.yaml`:

```yaml
values:
  pulumiConfig:
    github:owner: my-org
    github:token:
      fn::secret: ghp_xxxxxxxxxxxxxxxxxxxx
```

{{% /choosable %}}

{{% choosable configsource stack %}}

Set it with the CLI:

```bash
$ pulumi config set github:owner my-org
$ pulumi config set --secret github:token ghp_xxxxxxxxxxxxxxxxxxxx
```

That writes the following into `Pulumi.dev.yaml`:

```yaml
config:
  github:owner: my-org
  github:token:
    secure: AAABAOWyM7t3q1u...KpR8Nx1yVCjKPB2nJWg==
```

{{% /choosable %}}

The token can also come from the `GITHUB_TOKEN` environment variable. With no token at all, the provider can only reach resources available anonymously.

#### GitHub App installation

Authenticate as a [GitHub App installation](https://docs.github.com/en/apps/creating-github-apps) rather than as a user. All three `appAuth` fields are required, and `owner` is mandatory here — leaving it out produces a `403 "Resource not accessible by integration"` error.

{{< chooser configsource "esc,stack" / >}}

{{% choosable configsource esc %}}

Set it with the CLI:

```bash
$ pulumi env set myorg/github-dev pulumiConfig.github:appAuth.id 123456
$ pulumi env set myorg/github-dev pulumiConfig.github:appAuth.installationId 78901234
$ pulumi env set --secret myorg/github-dev pulumiConfig.github:appAuth.pemFile "$(cat app.pem)"
```

That produces the following environment definition in `environments/github-dev.yaml`:

```yaml
values:
  pulumiConfig:
    github:owner: my-org
    github:appAuth:
      id: "123456"
      installationId: "78901234"
      pemFile:
        fn::secret: |
          -----BEGIN RSA PRIVATE KEY-----
          MIIEowIBAAKCAQEA...
          -----END RSA PRIVATE KEY-----
```

{{% /choosable %}}

{{% choosable configsource stack %}}

Set it with the CLI:

```bash
$ pulumi config set github:owner my-org
$ pulumi config set --path github:appAuth.id 123456
$ pulumi config set --path github:appAuth.installationId 78901234
$ pulumi config set --path --secret github:appAuth.pemFile "$(cat app.pem)"
```

That writes the following into `Pulumi.dev.yaml`:

```yaml
config:
  github:owner: my-org
  github:appAuth:
    id: "123456"
    installationId: "78901234"
    pemFile:
      secure: AAABAJ4nQm1xPz9...8sVdK2eLhTgYbXc5Rq==
```

{{% /choosable %}}

These can also come from the `GITHUB_APP_ID`, `GITHUB_APP_INSTALLATION_ID` and `GITHUB_APP_PEM_FILE` environment variables. Some API operations are unavailable to App installations — see GitHub's list of [supported endpoints](https://docs.github.com/en/rest/overview/endpoints-available-for-github-apps).

#### GitHub Enterprise

Point the provider at your own GitHub Enterprise instance. The value must end with a slash.

{{< chooser configsource "esc,stack" / >}}

{{% choosable configsource esc %}}

Set it with the CLI:

```bash
$ pulumi env set myorg/github-dev pulumiConfig.github:baseUrl https://github.example.com/api/v3/
$ pulumi env set --secret myorg/github-dev pulumiConfig.github:token ghp_xxxxxxxxxxxxxxxxxxxx
```

That produces the following environment definition in `environments/github-dev.yaml`:

```yaml
values:
  pulumiConfig:
    github:baseUrl: https://github.example.com/api/v3/
    github:owner: my-org
    github:token:
      fn::secret: ghp_xxxxxxxxxxxxxxxxxxxx
```

{{% /choosable %}}

{{% choosable configsource stack %}}

Set it with the CLI:

```bash
$ pulumi config set github:baseUrl https://github.example.com/api/v3/
$ pulumi config set github:owner my-org
$ pulumi config set --secret github:token ghp_xxxxxxxxxxxxxxxxxxxx
```

That writes the following into `Pulumi.dev.yaml`:

```yaml
config:
  github:baseUrl: https://github.example.com/api/v3/
  github:owner: my-org
  github:token:
    secure: AAABAOWyM7t3q1u...KpR8Nx1yVCjKPB2nJWg==
```

{{% /choosable %}}

### Reference

Every configuration option accepted by the provider. Project any of them from an ESC environment under `pulumiConfig`, set them with `pulumi config set`, or pass them to an explicit [`github.Provider`](/registry/packages/github/api-docs/provider/). Nested objects are shown flattened, using [structured configuration](/docs/concepts/config#structured-configuration) paths.

{{% notes type="warning" %}}
`organization` is deprecated — use `owner` instead. For backwards compatibility, when more than one of `organization`, `GITHUB_ORGANIZATION`, `GITHUB_OWNER` and `owner` is set, the **first** in that list wins. That `GITHUB_OWNER` outranks `owner` is a known bug that may be fixed in a future major release; set only one of the four to stay compatible.
{{% /notes %}}

<div class="api-params">
<div class="api-param">
<h3 class="api-param-name" id="appauth-id" data-link-title="appAuth.id"><span class="api-param-nested-indicator">&#8627;</span>appAuth.id<a class="api-param-anchor" href="#appauth-id" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required is-required">required</span>
</div>
<div class="api-param-description">
<p>The GitHub App ID.</p>
<p class="api-param-meta">Config key: <code>github:appAuth.id</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.github:appAuth.id 123456
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> --path github:appAuth.id 123456
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="appauth-installationid" data-link-title="appAuth.installationId"><span class="api-param-nested-indicator">&#8627;</span>appAuth.installationId<a class="api-param-anchor" href="#appauth-installationid" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required is-required">required</span>
</div>
<div class="api-param-description">
<p>The GitHub App installation instance ID.</p>
<p class="api-param-meta">Config key: <code>github:appAuth.installationId</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.github:appAuth.installationId 78901234
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> --path github:appAuth.installationId 78901234
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="appauth-pemfile" data-link-title="appAuth.pemFile"><span class="api-param-nested-indicator">&#8627;</span>appAuth.pemFile<a class="api-param-anchor" href="#appauth-pemfile" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required is-required">required</span>
<span class="api-param-badge api-param-secret">secret</span>
</div>
<div class="api-param-description">
<p>The GitHub App PEM file contents.</p>
<p class="api-param-meta">Config key: <code>github:appAuth.pemFile</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> --secret myorg/myproject/dev pulumiConfig.github:appAuth.pemFile $(cat app.pem)
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> --path --secret github:appAuth.pemFile $(cat app.pem)
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="baseurl" data-link-title="baseUrl">baseUrl<a class="api-param-anchor" href="#baseurl" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required">optional</span>
<span class="api-param-badge api-param-in">GITHUB_BASE_URL</span>
</div>
<div class="api-param-description">
<p>The GitHub Base API URL</p>
<p class="api-param-meta">Config key: <code>github:baseUrl</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.github:baseUrl https://github.example.com/api/v3/
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> github:baseUrl https://github.example.com/api/v3/
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="insecure" data-link-title="insecure">insecure<a class="api-param-anchor" href="#insecure" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">boolean</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>Enable <code>insecure</code> mode for testing purposes</p>
<p class="api-param-meta">Config key: <code>github:insecure</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.github:insecure true
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> github:insecure true
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="maxperpage" data-link-title="maxPerPage">maxPerPage<a class="api-param-anchor" href="#maxperpage" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">integer</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>Number of items per page for paginationDefaults to 100</p>
<p class="api-param-meta">Config key: <code>github:maxPerPage</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.github:maxPerPage 3
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> github:maxPerPage 3
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="maxretries" data-link-title="maxRetries">maxRetries<a class="api-param-anchor" href="#maxretries" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">integer</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>Number of times to retry a request after receiving an error status codeDefaults to 3</p>
<p class="api-param-meta">Config key: <code>github:maxRetries</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.github:maxRetries 3
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> github:maxRetries 3
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="organization" data-link-title="organization">organization<a class="api-param-anchor" href="#organization" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>The GitHub organization name to manage. Use this field instead of <code>owner</code> when managing organization accounts.</p>
<p class="api-param-meta">Config key: <code>github:organization</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.github:organization my-org
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> github:organization my-org
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="owner" data-link-title="owner">owner<a class="api-param-anchor" href="#owner" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>The GitHub owner name to manage. Use this field instead of <code>organization</code> when managing individual accounts.</p>
<p class="api-param-meta">Config key: <code>github:owner</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.github:owner my-org
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> github:owner my-org
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="parallelrequests" data-link-title="parallelRequests">parallelRequests<a class="api-param-anchor" href="#parallelrequests" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">boolean</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>Allow the provider to make parallel API calls to GitHub. You may want to set it to true when you have a private Github Enterprise without strict rate limits. While it is possible to enable this setting on github.com, github.com's best practices recommend using serialization to avoid hitting abuse rate limitsDefaults to false if not set</p>
<p class="api-param-meta">Config key: <code>github:parallelRequests</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.github:parallelRequests true
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> github:parallelRequests true
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="readdelayms" data-link-title="readDelayMs">readDelayMs<a class="api-param-anchor" href="#readdelayms" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">integer</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>Amount of time in milliseconds to sleep in between non-write requests to GitHub API. Defaults to 0ms if not set.</p>
<p class="api-param-meta">Config key: <code>github:readDelayMs</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.github:readDelayMs 3
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> github:readDelayMs 3
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="retrydelayms" data-link-title="retryDelayMs">retryDelayMs<a class="api-param-anchor" href="#retrydelayms" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">integer</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>Amount of time in milliseconds to sleep in between requests to GitHub API after an error response. Defaults to 1000ms or 1s if not set, the maxRetries must be set to greater than zero.</p>
<p class="api-param-meta">Config key: <code>github:retryDelayMs</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.github:retryDelayMs 3
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> github:retryDelayMs 3
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="retryableerrors" data-link-title="retryableErrors">retryableErrors<a class="api-param-anchor" href="#retryableerrors" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">list&lt;integer&gt;</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>Allow the provider to retry after receiving an error status code, the maxRetries should be set for this to workDefaults to [500, 502, 503, 504]</p>
<p class="api-param-meta">Config key: <code>github:retryableErrors</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.github:retryableErrors &lt;value&gt;
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> github:retryableErrors &lt;value&gt;
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="token" data-link-title="token">token<a class="api-param-anchor" href="#token" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required">optional</span>
<span class="api-param-badge api-param-secret">secret</span>
<span class="api-param-badge api-param-in">GITHUB_TOKEN</span>
</div>
<div class="api-param-description">
<p>The OAuth token used to connect to GitHub. Anonymous mode is enabled if both <code>token</code> and <code>appAuth</code> are not set.</p>
<p class="api-param-meta">Config key: <code>github:token</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> --secret myorg/myproject/dev pulumiConfig.github:token ghp_xxxxxxxxxxxxxxxxxxxx
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> --secret github:token ghp_xxxxxxxxxxxxxxxxxxxx
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="writedelayms" data-link-title="writeDelayMs">writeDelayMs<a class="api-param-anchor" href="#writedelayms" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">integer</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>Amount of time in milliseconds to sleep in between writes to GitHub API. Defaults to 1000ms or 1s if not set.</p>
<p class="api-param-meta">Config key: <code>github:writeDelayMs</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.github:writeDelayMs 3
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> github:writeDelayMs 3
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
</div>
