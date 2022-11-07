---
title: Flux
meta_desc: Provides an overview of the Flux provider for Pulumi.
layout: overview
---

The Pulumi Flux provider lets you manage [Flux](https://fluxcd.io/) resources.

## Example

{{< chooser language "typescript,python,go,csharp" >}}
{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as tls from "@pulumi/tls";
import * as github from "@pulumi/github";
import * as k8s from "@pulumi/kubernetes";
import * as flux from "@worawat/flux";

// Required Github configuration
// export GITHUB_TOKEN=your-github-personal-access-token
// export GITHUB_OWNER=your-github-owner
const repoName = "pulumi-flux-sample";
const branch = "main";
const targetPath = "clusters/my-cluster";
const githubOwner = "oun";
const knownHosts =
  "github.com ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBEmKSENjQEezOmxkZMy7opKgwFB9nkt5YRrYMjNuG5N87uRgg6CLrbo5wAdT/y6v0mKV0U2w0WZ2YB/++Tpockg=";

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

new github.RepositoryDeployKey("key", {
  title: "fluxcd",
  repository: repo.name,
  key: key.publicKeyOpenssh,
  readOnly: true,
});

const fluxInstall = await flux.getFluxInstall({
  targetPath: targetPath,
});

const fluxSync = await flux.getFluxSync({
  targetPath: targetPath,
  url: `ssh://git@github.com/${githubOwner}/${repoName}.git`,
  branch: branch,
});

// Create kubernetes resource from generated manifests
const install = new k8s.yaml.ConfigGroup("flux-install", {
  yaml: fluxInstall.content,
});

const sync = new k8s.yaml.ConfigGroup("flux-sync", {
  yaml: fluxSync.content,
});

new k8s.core.v1.Secret(
  "flux",
  {
    metadata: {
      name: fluxSync.secret,
      namespace: fluxSync.namespace,
    },
    stringData: {
      identity: key.privateKeyPem,
      "identity.pub": key.publicKeyPem,
      known_hosts: knownHosts,
    },
  },
  { dependsOn: install }
);

// Commit files to Github
new github.RepositoryFile(
  "install",
  {
    repository: repo.name,
    file: fluxInstall.path,
    content: fluxInstall.content,
    branch: branch,
  },
  { dependsOn: install }
);

new github.RepositoryFile(
  "sync",
  {
    repository: repo.name,
    file: fluxSync.path,
    content: fluxSync.content,
    branch: branch,
  },
  { dependsOn: install }
);

new github.RepositoryFile(
  "kustomize",
  {
    repository: repo.name,
    file: fluxSync.kustomizePath,
    content: fluxSync.kustomizeContent,
    branch: branch,
  },
  { dependsOn: install }
);
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
import pulumi_flux as flux
import pulumi_github as github
import pulumi_tls as tls
import pulumi_kubernetes as k8s

# Require Github configurations
# export GITHUB_TOKEN=your-github-personal-access-token
# export GITHUB_OWNER=your-github-owner
config = pulumi.Config()
branch = config.require("branch")
target_path = config.require("target_path")
repository_name = config.require("repository")
github_owner = pulumi.Config("github").require("owner")

ssh_key = tls.PrivateKey("key", algorithm="ECDSA", ecdsa_curve="P256")

# Create Github repository
repository = github.Repository(
    "flux", name=repository_name, visibility="private", auto_init=True
)
github.BranchDefault("default", repository=repository.name, branch=branch)
github.RepositoryDeployKey(
    "flux-key",
    title="flux",
    repository=repository.name,
    key=ssh_key.public_key_openssh,
    read_only=True,
)

flux_install = flux.get_flux_install(target_path=target_path)
flux_sync = flux.get_flux_sync(
    target_path=target_path,
    url=f"ssh://git@github.com/{github_owner}/{repository_name}.git",
    branch=branch,
)

# Create kubernetes resource from generated manifests
install = k8s.yaml.ConfigGroup("flux-install", yaml=[flux_install.content])
sync = k8s.yaml.ConfigGroup("flux-sync", yaml=[flux_sync.content])

k8s.core.v1.Secret(
    "flux",
    metadata=k8s.meta.v1.ObjectMetaArgs(
        name=flux_sync.secret, namespace=flux_sync.namespace
    ),
    string_data={
        "identity": ssh_key.private_key_pem,
        "identity.pub": ssh_key.public_key_pem,
        "known_hosts": config.require("known_hosts"),
    },
    opts=pulumi.ResourceOptions(depends_on=install)
)

# Commit files to Github
github.RepositoryFile(
    "install",
    repository=repository.name,
    file=flux_install.path,
    content=flux_install.content,
    branch=branch,
    opts=pulumi.ResourceOptions(depends_on=install)
)
github.RepositoryFile(
    "sync",
    repository=repository.name,
    file=flux_sync.path,
    content=flux_sync.content,
    branch=branch,
    opts=pulumi.ResourceOptions(depends_on=install)
)
github.RepositoryFile(
    "kustomize",
    repository=repository.name,
    file=flux_sync.kustomize_path,
    content=flux_sync.kustomize_content,
    branch=branch,
    opts=pulumi.ResourceOptions(depends_on=install)
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
