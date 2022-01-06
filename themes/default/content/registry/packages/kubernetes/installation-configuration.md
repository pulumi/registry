---
title: Kubernetes Setup
meta_desc: Provides an overview on how to setup the Kubernetes Provider for Pulumi.
layout: installation
---

<!-- markdownlint-disable url -->
[pulumi-kubernetes-provider]: {{< relref "./" >}}
[client-go]: https://github.com/kubernetes/client-go
[gke-tutorial]: {{< relref "/registry/packages/kubernetes/how-to-guides/gke" >}}
[eks-tutorial]: {{< relref "/registry/packages/kubernetes/how-to-guides/eks" >}}
[aks-tutorial]: {{< relref "/registry/packages/kubernetes/how-to-guides/aks" >}}
[Heptio AWS quickstart]: https://aws.amazon.com/quickstart/architecture/heptio-kubernetes/
[provider-args]: {{< relref "/registry/packages/kubernetes/api-docs/provider" >}}
[provider-kubeconfig]: {{< relref "/registry/packages/kubernetes/api-docs/provider#inputs" >}}
[kubeconfig]: https://kubernetes.io/docs/tasks/access-application-cluster/configure-access-multiple-clusters/
[install]: {{< relref "/docs/get-started/install" >}}
[nodejs]: https://nodejs.org/en/
[npm]: https://www.npmjs.com/get-npm
[yarn]: https://yarnpkg.com/en/docs/install
<!-- markdownlint-enable url -->

Your Pulumi program is required to import the [pulumi/kubernetes][pulumi-kubernetes-provider] provider package to allow the Pulumi CLI to authenticate and interact with a running Kubernetes cluster.

By default, Pulumi will use a local [kubeconfig] if available, or one can be passed as a [provider argument][provider-kubeconfig] in the request.

With the `kubeconfig` available, Pulumi communicates with the API Server using the official Kubernetes [client-go] library, just like `kubectl` does.

## Installation

The Kubernetes provider is available as a package in all Pulumi languages:

- JavaScript/TypeScript: [`@pulumi/kubernetes`](https://www.npmjs.com/package/@pulumi/kubernetes)
- Python: [`pulumi-kubernetes`](https://pypi.org/project/pulumi-kubernetes/)
- Go: [`github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes`](https://github.com/pulumi/pulumi-kubernetes)
- .NET: [`Pulumi.Kubernetes`](https://www.nuget.org/packages/Pulumi.Kubernetes)

## Pre-Requisites

If you do not have a cluster set up and running yet, you'll need to do the
following steps.

1. [Install the Pulumi CLI][install].
1. Install a package manager for your Pulumi program language runtime, such as [npm] or [Yarn] for [Node.js][nodejs], or PyPI for Python.
1. Provision a Kubernetes cluster. For a new managed Kubernetes cluster, check out the [guides.]({{< relref "/registry/packages/kubernetes/how-to-guides" >}})
1. Download [`kubectl`](https://kubernetes.io/docs/tasks/tools/install-kubectl/) and verify the cluster is up and running.

## Setup

By default, Pulumi will look for a kubeconfig file in the following locations,
just like `kubectl`:

- The environment variable: `$KUBECONFIG`,
- Or in current user's default kubeconfig directory: `~/.kube/config`

If the kubeconfig file is not in either of these locations, Pulumi will **not** find it, and it will
fail to authenticate against the cluster. Set one of these locations to a valid kubeconfig file, if you have not done so
already.

Once the cluster is accessible, setup is complete and you can proceed to the
desired tutorials.

> Note: Pulumi **never** sends **any** authentication secrets or credentials to the Pulumi service. See the [FAQ]({{< relref "faq#does-the-pulumi-service-see-my-credentials-in-the-kubeconfig-file">}}) for more detail.

### Contexts

The kubeconfig file defines some number of _contexts_. Each context is a name that is associated
with a _cluster_, _namespace_, and a "_user_" (a local-only name that's associated with a credential
that allows access to the cluster).

To create a context, for example, you can run the `kubectl set-context` command as follows:

```shell
$ kubectl config \
    set-context my-context \
    --cluster=my-cluster \
    --user=my-user
```

If you have done this and are using the default context file, you will be able to set the
configuration variable `kubernetes:context` in the Pulumi config system to the given context name:

```shell
$ pulumi stack init new-kube-stack
$ pulumi config set kubernetes:context my-context
```

If you don't want to select a context, you can always make it the default:

```shell
$ kubectl config \
    use-context my-context
```

> **Note:** Depending on a default context is a bad idea if you're going to share your stack with
> others; it makes your stack dependent on ambient information not known to Pulumi, an anti-pattern
> that leads to unrepeatable deployments.

Additionally, the Kubernetes provider accepts many [configuration settings][provider-args].

These can be provided to the default Kubernetes provider via `pulumi config set kubernetes:<option>`, or passed
to the constructor of a `new kubernetes.Provider` to construct a specific instance of the Kubernetes provider for your requests.

### kubectl proxy

Each Kubernetes resource managed by Pulumi will have a link in the corresponding [Pulumi Console](https://app.pulumi.com")
to view the resource in the cluster. These links are local, and require the client run `kubectl proxy` beforehand to access the resource.

To learn more about `kubectl proxy` check out the [reference docs](https://kubernetes.io/docs/concepts/cluster-administration/proxies/).

## Configuration

The Kubernetes provider accepts the following configuration settings.  These can be provided to the default Kubernetes provider via `pulumi config set kubernetes:<option>`, or passed to the constructor of `new kubernetes.Provider` to construct a specific instance of the Kubernetes provider.

| Option                        | Required/Optional | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
|-------------------------------|-------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `cluster`                     | Optional          | If present, the name of the kubeconfig cluster to use.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| `context`                     | Optional          | If present, the name of the kubeconfig context to use.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| `enableDryRun`                | Optional          | BETA FEATURE - If present and set to true, enable server-side diff calculations. This feature is in developer preview, and is disabled by default. This config can be specified in the following ways, using this precedence: (1) this `enableDryRun` parameter or (2) the `PULUMI_K8S_ENABLE_DRY_RUN` environment variable.                                                                                                                                                                                                                                                                                                                                                |
| `kubeconfig`                  | Optional          | The contents of a kubeconfig file. If this is set, this config will be used instead of `$KUBECONFIG`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| `namespace`                   | Optional          | If present, the default namespace to use. This flag is ignored for cluster-scoped resources.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| `renderYamlToDirectory`       | Optional          | BETA FEATURE - If present, render resource manifests to this directory. In this mode, resources will not be created on a Kubernetes cluster, but the rendered manifests will be kept in sync with changes to the Pulumi program. This feature is in developer preview, and is disabled by default. Note that some computed Outputs such as status fields will not be populated since the resources are not created on a Kubernetes cluster. These Output values will remain undefined, and may result in an error if they are referenced by other resources. Also note that any secret values used in these resources will be rendered in plain text to the resulting YAML. |
| `suppressDeprecationWarnings` | Optional          | If present and set to true, suppress `apiVersion` deprecation warnings from the CLI. This config can be specified in the following ways, using this precedence: (1) this `suppressDeprecationWarnings` parameter or (2) the `PULUMI_K8S_SUPPRESS_DEPRECATION_WARNINGS` environment variable.                                                                                                                                                                                                                                                                                                                                                                                |

## Annotations

A few Pulumi-specific annotations can be applied to Kubernetes resources managed by Pulumi to control aspects of how Pulumi deploys and manages the Kubernetes resource:

- `pulumi.com/skipAwait`: Disables Pulumi's default await logic that waits for a Kubernetes resource to become "ready" before marking the resource as having created or updated succesfully.
- `pulumi.com/timeoutSeconds`: Specifies the number of seconds that the Pulumi Kubernetes provider will wait for the resource to become "ready".

In addition, the Pulumi provider may write the following annotations onto resources it manages:

- `pulumi.com/autonamed`: Indicates that the Pulumi Kubernetes provider decided to autoname the resource (instead of using an explicitly provided `metadata.name`).
