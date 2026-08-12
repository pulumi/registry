---
title: Kubernetes
meta_desc: Install and configure the Pulumi Kubernetes provider — SDK installation for every language, kubeconfig and context examples, and every provider configuration option.
layout: package
aliases:
    - /registry/packages/kubernetes/installation-configuration/
---

[Kubernetes](https://kubernetes.io) is an open source project for running and managing containerized applications on a cluster of machines. The Pulumi Kubernetes provider can provision any resource available in the Kubernetes API, in the programming language of your choice — no YAML templates, DSLs or tar archives.

Unlike a cloud provider, this provider has no credentials of its own. It connects to a cluster using a `kubeconfig`, exactly as `kubectl` does. New to Pulumi and Kubernetes? Start with the [Get Started guide](/docs/get-started/kubernetes).

{{< chooser language "typescript,python,go,csharp,java,yaml,hcl" / >}}

{{% choosable language typescript %}}

```typescript
import * as k8s from "@pulumi/kubernetes";

const nginx = new k8s.apps.v1.Deployment("nginx", {
    spec: {
        selector: { matchLabels: { app: "nginx" } },
        replicas: 2,
        template: {
            metadata: { labels: { app: "nginx" } },
            spec: { containers: [{ name: "nginx", image: "nginx:1.27" }] },
        },
    },
});

export const name = nginx.metadata.name;
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_kubernetes as k8s

nginx = k8s.apps.v1.Deployment("nginx",
    spec={
        "selector": {"match_labels": {"app": "nginx"}},
        "replicas": 2,
        "template": {
            "metadata": {"labels": {"app": "nginx"}},
            "spec": {"containers": [{"name": "nginx", "image": "nginx:1.27"}]},
        },
    })

pulumi.export("name", nginx.metadata["name"])
```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/apps/v1"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		labels := pulumi.StringMap{"app": pulumi.String("nginx")}

		nginx, err := appsv1.NewDeployment(ctx, "nginx", &appsv1.DeploymentArgs{
			Spec: &appsv1.DeploymentSpecArgs{
				Selector: &metav1.LabelSelectorArgs{MatchLabels: labels},
				Replicas: pulumi.Int(2),
				Template: &corev1.PodTemplateSpecArgs{
					Metadata: &metav1.ObjectMetaArgs{Labels: labels},
					Spec: &corev1.PodSpecArgs{
						Containers: corev1.ContainerArray{
							&corev1.ContainerArgs{
								Name:  pulumi.String("nginx"),
								Image: pulumi.String("nginx:1.27"),
							},
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}

		ctx.Export("name", nginx.Metadata.Name())
		return nil
	})
}
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using Pulumi;
using Kubernetes = Pulumi.Kubernetes;

return await Deployment.RunAsync(() =>
{
    var labels = new InputMap<string> { { "app", "nginx" } };

    var nginx = new Kubernetes.Apps.V1.Deployment("nginx", new()
    {
        Spec = new Kubernetes.Types.Inputs.Apps.V1.DeploymentSpecArgs
        {
            Selector = new Kubernetes.Types.Inputs.Meta.V1.LabelSelectorArgs { MatchLabels = labels },
            Replicas = 2,
            Template = new Kubernetes.Types.Inputs.Core.V1.PodTemplateSpecArgs
            {
                Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs { Labels = labels },
                Spec = new Kubernetes.Types.Inputs.Core.V1.PodSpecArgs
                {
                    Containers = new[]
                    {
                        new Kubernetes.Types.Inputs.Core.V1.ContainerArgs
                        {
                            Name = "nginx",
                            Image = "nginx:1.27",
                        },
                    },
                },
            },
        },
    });

    return new Dictionary<string, object?>
    {
        ["name"] = nginx.Metadata.Apply(m => m.Name),
    };
});
```

{{% /choosable %}}

{{% choosable language java %}}

```java
import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.kubernetes.apps_v1.Deployment;
import com.pulumi.kubernetes.apps_v1.DeploymentArgs;
import com.pulumi.kubernetes.apps_v1.inputs.DeploymentSpecArgs;
import com.pulumi.kubernetes.core_v1.inputs.ContainerArgs;
import com.pulumi.kubernetes.core_v1.inputs.PodSpecArgs;
import com.pulumi.kubernetes.core_v1.inputs.PodTemplateSpecArgs;
import com.pulumi.kubernetes.meta_v1.inputs.LabelSelectorArgs;
import com.pulumi.kubernetes.meta_v1.inputs.ObjectMetaArgs;
import java.util.Map;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    private static void stack(Context ctx) {
        var labels = Map.of("app", "nginx");

        var nginx = new Deployment("nginx", DeploymentArgs.builder()
            .spec(DeploymentSpecArgs.builder()
                .selector(LabelSelectorArgs.builder().matchLabels(labels).build())
                .replicas(2)
                .template(PodTemplateSpecArgs.builder()
                    .metadata(ObjectMetaArgs.builder().labels(labels).build())
                    .spec(PodSpecArgs.builder()
                        .containers(ContainerArgs.builder()
                            .name("nginx")
                            .image("nginx:1.27")
                            .build())
                        .build())
                    .build())
                .build())
            .build());

        ctx.export("name", nginx.metadata().applyValue(m -> m.name()));
    }
}
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: kubernetes-example
runtime: yaml
resources:
  nginx:
    type: kubernetes:apps/v1:Deployment
    properties:
      spec:
        selector:
          matchLabels:
            app: nginx
        replicas: 2
        template:
          metadata:
            labels:
              app: nginx
          spec:
            containers:
              - name: nginx
                image: nginx:1.27
outputs:
  name: ${nginx.metadata.name}
```

{{% /choosable %}}

{{% choosable language hcl %}}

```hcl
resource "kubernetes_deployment" "nginx" {
  spec {
    replicas = 2

    selector {
      match_labels = {
        app = "nginx"
      }
    }

    template {
      metadata {
        labels = {
          app = "nginx"
        }
      }

      spec {
        container {
          name  = "nginx"
          image = "nginx:1.27"
        }
      }
    }
  }
}
```

{{% /choosable %}}

## Installation

The Kubernetes provider publishes an SDK for every language. Add it to an existing Pulumi project with your language's package manager; YAML and HCL projects use [`pulumi package add`](/docs/iac/cli/commands/pulumi_package_add/) instead:

{{< chooser language "typescript,python,go,csharp,java,yaml,hcl" / >}}

{{% choosable language typescript %}}

```bash
npm install @pulumi/kubernetes
```

{{% /choosable %}}

{{% choosable language python %}}

```bash
pip install pulumi-kubernetes
```

{{% /choosable %}}

{{% choosable language go %}}

```bash
go get github.com/pulumi/pulumi-kubernetes/sdk/v4
```

Import it in your program with `import "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"`.

{{% /choosable %}}

{{% choosable language csharp %}}

```bash
dotnet add package Pulumi.Kubernetes
```

{{% /choosable %}}

{{% choosable language java %}}

Add the dependency to your build file. For Maven, in `pom.xml`:

```xml
<dependency>
    <groupId>com.pulumi</groupId>
    <artifactId>kubernetes</artifactId>
    <version>4.33.0</version>
</dependency>
```

For Gradle, in `build.gradle`:

```groovy
implementation 'com.pulumi:kubernetes:4.33.0'
```

{{% /choosable %}}

{{% choosable language yaml %}}

```bash
pulumi package add kubernetes
```

Reference resources by their type token, for example `kubernetes:apps/v1:Deployment`. Omit the version to add the latest, or append `@4.33.0` to pin one.

{{% /choosable %}}

{{% choosable language hcl %}}

```bash
pulumi package add kubernetes
```

This makes the `kubernetes_*` resource types available to your converted HCL. Omit the version to add the latest, or append `@4.33.0` to pin one.

{{% /choosable %}}

To install everything a project already declares — after cloning it from source control, for example — run [`pulumi install`](/docs/iac/cli/commands/pulumi_install/).

## Configuration

Provider settings can come from any of three places:

- A [Pulumi ESC](/docs/pulumi-cloud/esc/) environment, under `pulumiConfig`.
- [Stack configuration](/docs/iac/concepts/config/), set with `pulumi config set` and stored in `Pulumi.<stack-name>.yaml`.
- An explicit `kubernetes.Provider` instance, constructed in your program and passed to the resources that should use it.

See [Configuration](/docs/concepts/config/) for how these sources combine.

This provider has no credentials of its own. It reads a `kubeconfig` from `$KUBECONFIG` or `~/.kube/config`, just like `kubectl`. If the file is in neither location the provider will fail to reach the cluster. Verify your setup with `kubectl get pods` before running Pulumi.

{{% notes type="info" %}}
Pulumi **never** sends authentication [secrets](/docs/concepts/secrets/) or credentials to Pulumi Cloud.
{{% /notes %}}

### Examples

Every example below is shown two ways: as a [Pulumi ESC](/docs/pulumi-cloud/esc/) environment, and as [stack configuration](/docs/iac/concepts/config/) — the `Pulumi.<stack-name>.yaml` file alongside your project. Pick whichever you use; the choice follows you down the page and across providers.

#### Default kubeconfig

Nothing to configure. The provider finds `$KUBECONFIG` or `~/.kube/config` and uses its current context.

```bash
$ kubectl config view
$ pulumi up
```

{{% notes type="warning" %}}
Relying on the ambient current context is a bad idea for a stack you intend to share — it makes deployments depend on information Pulumi doesn't know about, which is how you get unrepeatable results. Set `kubernetes:context` explicitly instead.
{{% /notes %}}

#### An explicit context

A kubeconfig defines any number of *contexts*, each pairing a cluster, a namespace and a user. Multiple kubeconfig sources are merged, and the result can surprise you — check with `kubectl config view` first.

```bash
$ kubectl config set-context my-context --cluster=my-cluster --user=my-user
```

{{< chooser configsource "esc,stack" / >}}

{{% choosable configsource esc %}}

Set it with the CLI:

```bash
$ pulumi env set myorg/k8s-dev pulumiConfig.kubernetes:context my-context
$ pulumi env set myorg/k8s-dev pulumiConfig.kubernetes:namespace my-namespace
```

That produces the following environment definition in `environments/k8s-dev.yaml`:

```yaml
values:
  pulumiConfig:
    kubernetes:context: my-context
    kubernetes:namespace: my-namespace
```

{{% /choosable %}}

{{% choosable configsource stack %}}

Set it with the CLI:

```bash
$ pulumi config set kubernetes:context my-context
$ pulumi config set kubernetes:namespace my-namespace
```

That writes the following into `Pulumi.dev.yaml`:

```yaml
config:
  kubernetes:context: my-context
  kubernetes:namespace: my-namespace
```

{{% /choosable %}}

#### An explicit kubeconfig

Pass a cluster's kubeconfig directly — the usual pattern when the same program also creates the cluster, or when running somewhere with no kubeconfig on disk.

{{< chooser configsource "esc,stack" / >}}

{{% choosable configsource esc %}}

Set it with the CLI:

```bash
$ pulumi env set --secret myorg/k8s-dev pulumiConfig.kubernetes:kubeconfig "$(cat ~/.kube/config)"
```

That produces the following environment definition in `environments/k8s-dev.yaml`:

```yaml
values:
  pulumiConfig:
    kubernetes:kubeconfig:
      fn::secret: |
        apiVersion: v1
        clusters:
          - cluster:
              server: https://my-cluster.example.com
            name: my-cluster
```

{{% /choosable %}}

{{% choosable configsource stack %}}

Set it with the CLI:

```bash
$ pulumi config set --secret kubernetes:kubeconfig "$(cat ~/.kube/config)"
```

That writes the following into `Pulumi.dev.yaml`:

```yaml
config:
  kubernetes:kubeconfig:
    secure: AAABAOWyM7t3q1u...KpR8Nx1yVCjKPB2nJWg==
```

{{% /choosable %}}

#### Render to YAML instead of applying

Set `renderYamlToDirectory` and the provider writes manifests to disk rather than talking to a cluster — useful for GitOps flows and for inspecting exactly what Pulumi would apply.

{{< chooser configsource "esc,stack" / >}}

{{% choosable configsource esc %}}

Set it with the CLI:

```bash
$ pulumi env set myorg/k8s-dev pulumiConfig.kubernetes:renderYamlToDirectory ./rendered
```

That produces the following environment definition in `environments/k8s-dev.yaml`:

```yaml
values:
  pulumiConfig:
    kubernetes:renderYamlToDirectory: ./rendered
```

{{% /choosable %}}

{{% choosable configsource stack %}}

Set it with the CLI:

```bash
$ pulumi config set kubernetes:renderYamlToDirectory ./rendered
```

That writes the following into `Pulumi.dev.yaml`:

```yaml
config:
  kubernetes:renderYamlToDirectory: ./rendered
```

{{% /choosable %}}

### Reference

Every configuration option accepted by the provider. Project any of them from an ESC environment under `pulumiConfig`, set them with `pulumi config set`, or pass them to an explicit [`kubernetes.Provider`](/registry/packages/kubernetes/api-docs/provider/).

<div class="api-params">
<div class="api-param">
<h3 class="api-param-name" id="alwaysrender" data-link-title="alwaysRender">alwaysRender<a class="api-param-anchor" href="#alwaysrender" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">boolean</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>If present and set to true, all resources will be rendered to the directory specified by renderYamlToDirectory on every update, even if the resource has not changed. This is useful for tools like ArgoCD Config Management Plugin that require all manifests to be regenerated on each run. Only valid when renderYamlToDirectory is set.</p>
<p class="api-param-meta">Config key: <code>kubernetes:alwaysRender</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.kubernetes:alwaysRender true
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> kubernetes:alwaysRender true
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="cluster" data-link-title="cluster">cluster<a class="api-param-anchor" href="#cluster" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>If present, the name of the kubeconfig cluster to use.</p>
<p class="api-param-meta">Config key: <code>kubernetes:cluster</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.kubernetes:cluster my-cluster
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> kubernetes:cluster my-cluster
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="clusteridentifier" data-link-title="clusterIdentifier">clusterIdentifier<a class="api-param-anchor" href="#clusteridentifier" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>If present, this value will control the provider's replacement behavior. In particular, the provider will _only_ be replaced when <code>clusterIdentifier</code> changes; all other changes to provider configuration will be treated as updates. Kubernetes does not yet offer an API for cluster identification, so Pulumi uses heuristics to decide when a provider resource should be replaced or updated. These heuristics can sometimes lead to destructive replace operations when an update would be more appropriate, or vice versa. Use <code>clusterIdentifier</code> for more fine-grained control of the provider resource's lifecycle.</p>
<p class="api-param-meta">Config key: <code>kubernetes:clusterIdentifier</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.kubernetes:clusterIdentifier my-cluster
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> kubernetes:clusterIdentifier my-cluster
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="context" data-link-title="context">context<a class="api-param-anchor" href="#context" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>If present, the name of the kubeconfig context to use.</p>
<p class="api-param-meta">Config key: <code>kubernetes:context</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.kubernetes:context my-cluster-context
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> kubernetes:context my-cluster-context
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="deleteunreachable" data-link-title="deleteUnreachable">deleteUnreachable<a class="api-param-anchor" href="#deleteunreachable" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">boolean</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>If present and set to true, the provider will delete resources associated with an unreachable Kubernetes cluster from Pulumi state</p>
<p class="api-param-meta">Config key: <code>kubernetes:deleteUnreachable</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.kubernetes:deleteUnreachable true
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> kubernetes:deleteUnreachable true
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="enableconfigmapmutable" data-link-title="enableConfigMapMutable">enableConfigMapMutable<a class="api-param-anchor" href="#enableconfigmapmutable" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">boolean</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>If present and set to true, allow ConfigMaps to be mutated. Disabled by default for backward-compatibility. This config can be specified in the following ways using this precedence: 1. This <code>enableConfigMapMutable</code> parameter. 2. The <code>PULUMI_K8S_ENABLE_CONFIGMAP_MUTABLE</code> environment variable.</p>
<p class="api-param-meta">Config key: <code>kubernetes:enableConfigMapMutable</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.kubernetes:enableConfigMapMutable true
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> kubernetes:enableConfigMapMutable true
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="enablepatchforce" data-link-title="enablePatchForce">enablePatchForce<a class="api-param-anchor" href="#enablepatchforce" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">boolean</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>If present and set to true, enable patch force on all Server-Side Apply operations, overriding any field conflicts. See https://github.com/pulumi/pulumi-kubernetes/issues/2280 for additional details. This config can be specified in the following ways using this precedence: 1. The <code>pulumi.com/patchForce</code> annotation on the resource. 2. This <code>enablePatchForce</code> parameter. 3. The <code>PULUMI_K8S_ENABLE_PATCH_FORCE</code> environment variable.</p>
<p class="api-param-meta">Config key: <code>kubernetes:enablePatchForce</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.kubernetes:enablePatchForce true
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> kubernetes:enablePatchForce true
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="enablereplacecrd" data-link-title="enableReplaceCRD">enableReplaceCRD<a class="api-param-anchor" href="#enablereplacecrd" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">boolean</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>Obsolete. This option has no effect.</p>
<p class="api-param-meta">Config key: <code>kubernetes:enableReplaceCRD</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.kubernetes:enableReplaceCRD true
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> kubernetes:enableReplaceCRD true
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="enablesecretmutable" data-link-title="enableSecretMutable">enableSecretMutable<a class="api-param-anchor" href="#enablesecretmutable" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">boolean</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>If present and set to true, allow Secrets to be mutated. Disabled by default for backward compatibility. This config can be specified in the following ways using this precedence: 1. This <code>enableSecretMutable</code> parameter. 2. The <code>PULUMI_K8S_ENABLE_SECRET_MUTABLE</code> environment variable.</p>
<p class="api-param-meta">Config key: <code>kubernetes:enableSecretMutable</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.kubernetes:enableSecretMutable true
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> kubernetes:enableSecretMutable true
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="enableserversideapply" data-link-title="enableServerSideApply">enableServerSideApply<a class="api-param-anchor" href="#enableserversideapply" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">boolean</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>If present and set to false, disable Server-Side Apply mode. See https://github.com/pulumi/pulumi-kubernetes/issues/2011 for additional details.</p>
<p class="api-param-meta">Config key: <code>kubernetes:enableServerSideApply</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.kubernetes:enableServerSideApply true
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> kubernetes:enableServerSideApply true
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="kubeconfig" data-link-title="kubeconfig">kubeconfig<a class="api-param-anchor" href="#kubeconfig" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>The contents of a kubeconfig file or the path to a kubeconfig file. If this is set, this config will be used instead of $KUBECONFIG.</p>
<p class="api-param-meta">Config key: <code>kubernetes:kubeconfig</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.kubernetes:kubeconfig $(cat ~/.kube/config)
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> kubernetes:kubeconfig $(cat ~/.kube/config)
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="namespace" data-link-title="namespace">namespace<a class="api-param-anchor" href="#namespace" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>If present, the default namespace to use. This flag is ignored for cluster-scoped resources. A namespace can be specified in multiple places, and the precedence is as follows: 1. <code>.metadata.namespace</code> set on the resource. 2. This <code>namespace</code> parameter. 3. <code>namespace</code> set for the active context in the kubeconfig.</p>
<p class="api-param-meta">Config key: <code>kubernetes:namespace</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.kubernetes:namespace my-namespace
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> kubernetes:namespace my-namespace
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="renderyamltodirectory" data-link-title="renderYamlToDirectory">renderYamlToDirectory<a class="api-param-anchor" href="#renderyamltodirectory" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>BETA FEATURE - If present, render resource manifests to this directory. In this mode, resources will not be created on a Kubernetes cluster, but the rendered manifests will be kept in sync with changes to the Pulumi program. This feature is in developer preview, and is disabled by default. Render mode attempts to connect to the cluster identified by your kubeconfig to determine whether custom resources are namespaced or cluster-scoped. When no cluster is reachable, rendering proceeds anyway. Affected resources are written without a namespace scope, falling back to kubectl's default namespace behavior on apply, and a warning naming each unresolved kind is emitted. Note that some computed Outputs such as status fields will not be populated since the resources are not created on a Kubernetes cluster. These Output values will remain undefined, and may result in an error if they are referenced by other resources. Also note that any secret values used in these resources will be rendered in plaintext to the resulting YAML.</p>
<p class="api-param-meta">Config key: <code>kubernetes:renderYamlToDirectory</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.kubernetes:renderYamlToDirectory ./rendered
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> kubernetes:renderYamlToDirectory ./rendered
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="skipupdateunreachable" data-link-title="skipUpdateUnreachable">skipUpdateUnreachable<a class="api-param-anchor" href="#skipupdateunreachable" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">boolean</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>If present and set to true, the provider will skip resources update associated with an unreachable Kubernetes cluster from Pulumi state</p>
<p class="api-param-meta">Config key: <code>kubernetes:skipUpdateUnreachable</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.kubernetes:skipUpdateUnreachable true
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> kubernetes:skipUpdateUnreachable true
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="strictmode" data-link-title="strictMode">strictMode<a class="api-param-anchor" href="#strictmode" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">boolean</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>If present and set to true, the provider will use strict configuration mode. Recommended for production stacks. In this mode, the default Kubernetes provider is disabled, and the <code>kubeconfig</code> and <code>context</code> settings are required for Provider configuration. These settings unambiguously ensure that every Kubernetes resource is associated with a particular cluster.</p>
<p class="api-param-meta">Config key: <code>kubernetes:strictMode</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.kubernetes:strictMode true
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> kubernetes:strictMode true
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="suppressdeprecationwarnings" data-link-title="suppressDeprecationWarnings">suppressDeprecationWarnings<a class="api-param-anchor" href="#suppressdeprecationwarnings" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">boolean</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>If present and set to true, suppress apiVersion deprecation warnings from the CLI. This config can be specified in the following ways, using this precedence: 1. This <code>suppressDeprecationWarnings</code> parameter. 2. The <code>PULUMI_K8S_SUPPRESS_DEPRECATION_WARNINGS</code> environment variable.</p>
<p class="api-param-meta">Config key: <code>kubernetes:suppressDeprecationWarnings</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.kubernetes:suppressDeprecationWarnings true
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> kubernetes:suppressDeprecationWarnings true
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="suppresshelmhookwarnings" data-link-title="suppressHelmHookWarnings">suppressHelmHookWarnings<a class="api-param-anchor" href="#suppresshelmhookwarnings" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">boolean</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>If present and set to true, suppress unsupported Helm hook warnings from the CLI. This config can be specified in the following ways, using this precedence: 1. This <code>suppressHelmHookWarnings</code> parameter. 2. The <code>PULUMI_K8S_SUPPRESS_HELM_HOOK_WARNINGS</code> environment variable.</p>
<p class="api-param-meta">Config key: <code>kubernetes:suppressHelmHookWarnings</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.kubernetes:suppressHelmHookWarnings true
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> kubernetes:suppressHelmHookWarnings true
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="upsertexistingobjects" data-link-title="upsertExistingObjects">upsertExistingObjects<a class="api-param-anchor" href="#upsertexistingobjects" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">boolean</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>If present and set to true, allow Pulumi to create resources that already exist in the cluster by updating them instead of returning an error. By default, Pulumi will error if a resource already exists in the cluster to prevent accidental data loss. When a Pulumi resource is renamed without using aliases, the engine plans a create followed by a delete targeting the same cluster object. With server-side apply, the create silently updates the existing object, and the subsequent delete removes it — resulting in unexpected resource deletion. Enabling this option restores the previous upsert behavior for users who intentionally adopt existing cluster resources into Pulumi. This config can be specified in the following ways using this precedence: 1. This <code>upsertExistingObjects</code> parameter. 2. The <code>PULUMI_K8S_UPSERT_EXISTING_OBJECTS</code> environment variable.</p>
<p class="api-param-meta">Config key: <code>kubernetes:upsertExistingObjects</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.kubernetes:upsertExistingObjects true
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> kubernetes:upsertExistingObjects true
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
</div>

## Server-Side Apply

[Server-Side Apply](https://kubernetes.io/docs/reference/using-api/server-side-apply/) (SSA) makes the API server responsible for computing diffs and resolving conflicts, so several clients can safely share management of a resource. The [v4 release](https://github.com/pulumi/pulumi-kubernetes/releases/tag/v4.0.1) of this provider enabled SSA by default, which means:

1. Kubernetes resources may be safely managed by more than one controller.
2. Resources can be "upserted" — created if absent, otherwise updated in place.
3. Resources can be patched with the `Patch` resource types; every resource type in the SDK has a corresponding one.
4. The `last-applied-configuration` annotation is no longer used.

See the [Server-Side Apply how-to guide](/registry/packages/kubernetes/how-to-guides/managing-resources-with-server-side-apply/) for details.

## Resource annotations

These annotations are set on individual Kubernetes resources — they are not provider configuration, and so do not appear in [Reference](#reference) above. They control how Pulumi deploys and manages the resource they are attached to.

### pulumi.com/skipAwait

Controls Pulumi's behavior while waiting for a resource to become ready. Set to `"true"`, Pulumi creates the resource but does not wait for readiness.

{{% notes type="warning" %}}
A small number of resources (Deployments, DaemonSets, StatefulSets, Pods and Namespaces) currently respect `skipAwait` during deletion and do not wait for deletion to succeed. Using `skipAwait` during deletion is not recommended when server-side apply is enabled, because it can cause race conditions during replacement. This behavior is considered buggy and may change.

`pulumi.com/deletionPropagationPolicy`, below, is almost always the better way to delete something quickly and safely.
{{% /notes %}}

{{% notes type="info" %}}
Pulumi has no concept of readiness for every resource, and will often assume a resource is immediately ready even without `skipAwait` — a problem if dependents genuinely need it to be ready. Use the `waitFor` annotation below, or run with `PULUMI_K8S_AWAIT_ALL=true` to wait on arbitrary resources.
{{% /notes %}}

### pulumi.com/deletionPropagationPolicy

(New in v4.12.0.) By default Pulumi uses [foreground](https://kubernetes.io/docs/concepts/architecture/garbage-collection/#foreground-deletion) cascading deletion, removing the resource *and* all of its dependents. That is thorough but slow. This annotation selects an alternative:

- `background` — delete the owner and leave dependents to asynchronous garbage collection. Faster, but dependents can linger, temporarily or indefinitely, if they are not finalized.
- `orphan` — delete the owner and leave dependents untouched. Useful for migration or debugging.
- `foreground` — the default: delete the resource and all dependents, guaranteeing cleanup if it succeeds.

### pulumi.com/waitFor

(New in v4.18.0.) Defines custom readiness criteria. It accepts three forms:

1. A `kubectl` [JSONPath](https://kubernetes.io/docs/reference/kubectl/jsonpath) — a string prefixed with `jsonpath=` followed by a path expression and an optional value. With a value, the resource is ready when the expression evaluates to it; without one, when any value exists at that path.

    ```yaml
    "pulumi.com/waitFor": "jsonpath={.phase}=Running"
    "pulumi.com/waitFor": "jsonpath={.webhooks[*].clientConfig.caBundle}"
    ```

2. A string prefixed with `condition=` followed by a condition type and optional status, matching `kubectl --for=condition=...`. The expected status defaults to `True`.

    ```yaml
    "pulumi.com/waitFor": "condition=Synced"
    "pulumi.com/waitFor": "condition=Reconciling=False"
    ```

3. A JSON array of several such expressions; the resource is ready when all are met simultaneously.

    ```yaml
    "pulumi.com/waitFor": '["jsonpath={.foo}", "condition=Bar"]'
    ```

This annotation has no effect if `pulumi.com/skipAwait` is also present with a value of `"true"` or `"ready"`.

### pulumi.com/patchForce

(Server-Side Apply option.) Force override any conflicts for the specified resource.

### pulumi.com/patchFieldManager

(Server-Side Apply option.) Specify the `FieldManager` name to use for the Server-Side Apply operation.

### pulumi.com/timeoutSeconds

How many seconds the provider will wait for the resource to become ready. Consider [custom timeouts](/docs/concepts/options/customtimeouts/) instead.

### pulumi.com/replaceUnready

If the resource failed to become ready in the previous update, replace it rather than continuing to wait. Only `batch/v1/Job` supports this today.

### Annotations Pulumi writes

The provider may also write these onto resources it manages:

- `app.kubernetes.io/managed-by` — indicates the controller managing a resource. Not set in Server-Side Apply mode.
- `pulumi.com/autonamed` — indicates the provider auto-named the resource rather than using an explicit `metadata.name`.

## Related packages

To provision the clusters themselves, or to work with Kubernetes-adjacent tooling:

- [AWS](/registry/packages/aws/) · [Azure Native](/registry/packages/azure-native/) · [Google Cloud](/registry/packages/gcp/) · [DigitalOcean](/registry/packages/digitalocean/)
- [EKS](/registry/packages/eks/) — managed Amazon EKS clusters
- [AWSx](/registry/packages/awsx/) — AWS extension components

The [Pulumi Kubernetes Operator](https://github.com/pulumi/pulumi-kubernetes-operator) makes a Pulumi [Stack](/docs/concepts/stack/) a first-class Kubernetes API resource, so you can drive stack updates from inside the cluster as part of a [continuous delivery](/docs/iac/using-pulumi/continuous-delivery/) pipeline.

## kubectl proxy

Every Kubernetes resource managed by Pulumi gets a link in the [Pulumi Cloud](https://app.pulumi.com) console to view it in the cluster. Those links are local: run [`kubectl proxy`](https://kubernetes.io/docs/concepts/cluster-administration/proxies/) first for them to resolve.
