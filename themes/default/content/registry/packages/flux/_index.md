---
title: Flux
meta_desc: Provides an overview of the Flux provider for Pulumi.
layout: package
---

The Pulumi Flux provider lets you manage [Flux](https://fluxcd.io/) resources.

## Example

{{< chooser language "typescript,python,go,csharp" >}}
{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as tls from "@pulumi/tls";
import * as github from "@pulumi/github";
import * as flux from "@worawat/flux";

// Require Github configurations
// export GITHUB_TOKEN=your-github-personal-access-token
// export GITHUB_OWNER=your-github-owner

const repoName = "your-github-repo";
const branch = "main";
const path = "clusters/dev";
const githubOwner = "your-github-user";

// Generate ssh keys
const key = new tls.PrivateKey("key", {
  algorithm: "ECDSA",
  ecdsaCurve: "P256",
});

// Create Github repository
const repo = new github.Repository("repo", {
  name: repoName,
  visibility: "private",
  autoInit: true,
});

new github.BranchDefault("default", {
  repository: repo.name,
  branch: branch,
});

// Add generated public key to Github deploy key
const deployKey = new github.RepositoryDeployKey("key", {
  title: "fluxcd",
  repository: repo.name,
  key: key.publicKeyOpenssh,
  readOnly: false,
});

const provider = new flux.Provider("flux", {
  kubernetes: {
    configPath: "~/.kube/config",
  },
  git: {
    url: `ssh://git@github.com/${githubOwner}/${repoName}.git`,
    branch: branch,
    ssh: {
      username: "git",
      privateKey: key.privateKeyPem,
    },
  },
});

const resource = new flux.FluxBootstrapGit(
  "flux",
  {
    path: path,
  },
  {
    provider: provider,
    dependsOn: deployKey,
  }
);
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
import pulumi_flux as flux
import pulumi_github as github
import pulumi_tls as tls

# Require Github configurations
# export GITHUB_TOKEN=your-github-personal-access-token
# export GITHUB_OWNER=your-github-owner

branch = "main"
target_path = "clusters/dev"
repo_name = "your-github-repo"
github_owner = "your-github-user"

ssh_key = tls.PrivateKey("key", algorithm="ECDSA", ecdsa_curve="P256")

repository = github.Repository(
    "flux", name=repo_name, visibility="private", auto_init=True
)

github.BranchDefault("default", repository=repository.name, branch=branch)

deploy_key = github.RepositoryDeployKey(
    "flux-key",
    title="flux",
    repository=repository.name,
    key=ssh_key.public_key_openssh,
    read_only=False,
)

provider = flux.Provider(
    "flux",
    kubernetes=flux.ProviderKubernetesArgs(config_path="~/.kube/config"),
    git=flux.ProviderGitArgs(
        url=f"ssh://git@github.com/{github_owner}/{repo_name}.git",
        branch=branch,
        ssh=flux.ProviderGitSshArgs(username="git", private_key=ssh_key.private_key_pem)
    )
)

resource = flux.FluxBootstrapGit(
    "flux",
    path=target_path,
    opts=pulumi.ResourceOptions(provider=provider, depends_on=deploy_key)
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
 "fmt"
 flux "github.com/oun/pulumi-flux/sdk/go/flux"
 "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using Pulumi;
using Pulumi.Flux;

```

{{% /choosable %}}
{{< /chooser >}}
