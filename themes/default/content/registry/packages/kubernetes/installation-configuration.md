---
title: Kubernetes Installation & Configuration
meta_desc: Provides an overview on how to setup the Kubernetes Provider for Pulumi.
layout: package
---

<!-- markdownlint-disable url -->
[pulumi-kubernetes-provider]: ./
[client-go]: https://github.com/kubernetes/client-go
[provider-args]: /registry/packages/kubernetes/api-docs/provider
[provider-kubeconfig]: /registry/packages/kubernetes/api-docs/provider#inputs
[ssa-guide]: /registry/packages/kubernetes/how-to-guides/managing-resources-with-server-side-apply
[kubeconfig]: https://kubernetes.io/docs/tasks/access-application-cluster/configure-access-multiple-clusters/
[install]: /docs/install/
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
- Java: [`com.pulumi/kubernetes`](https://central.sonatype.com/artifact/com.pulumi/kubernetes)

## Pre-Requisites

If you do not have a cluster set up and running yet, you'll need to do the
following steps.

1. [Install the Pulumi CLI][install].
1. Install a package manager for your Pulumi program language runtime, such as [npm] or [Yarn] for [Node.js][nodejs], or PyPI for Python.
1. Provision a Kubernetes cluster. For a new managed Kubernetes cluster, check out the [guides.](/registry/packages/kubernetes/how-to-guides)
1. Download [`kubectl`](https://kubernetes.io/docs/tasks/tools/install-kubectl/).

## Setup

By default, Pulumi will look for a kubeconfig file in the following locations,
just like `kubectl`:

- The environment variable: `$KUBECONFIG`,
- Or in current user's default kubeconfig directory: `~/.kube/config`

If the kubeconfig file is not in either of these locations, Pulumi will **not** find it, and it will
fail to authenticate against the cluster. Set one of these locations to a valid kubeconfig file, if you have not done so
already.

Verify the cluster is configured and up by running `kubectl get pods`.

Then follow the [getting started guide](/docs/get-started/kubernetes) or explore some [code examples](https://github.com/pulumi/examples#kubernetes).

Note: Pulumi **never** sends **any** authentication [secrets](/docs/concepts/secrets/) or credentials to the Pulumi Cloud.

### Contexts

The kubeconfig file defines some number of _contexts_. Each context is a name that is associated
with a _cluster_, _namespace_, and a "_user_" (a local-only name that's associated with a credential
that allows access to the cluster).

Your Kubernetes implementation may have already written out an appropriate kubeconfig. 
For example, `minikube start` does this, unless you specified `--keep-context=true`.
Moreover, multiple sources of kubeconfig are merged and the result may surprise you.
Therefore, check your current kubeconfig using:

```shell
$ kubectl \
    config \
        view
```

To create a context, for example, you can run the `kubectl set-context` command as follows:

```shell
$ kubectl config \
    set-context my-context \
    --cluster=my-cluster \
    --user=my-user
```

Mind that e.g., `minikube` uses the same name for the context and the cluster in it.

If you have completed your kubeconfig configuration, and are using the default kubeconfig file, you will be able to set the
configuration variable `kubernetes:context` in the Kubernetes provider to the given context name:

```shell
$ pulumi stack init new-kube-stack
$ pulumi config set kubernetes:context my-context
```

If you don't want to select a context, you can always make it the operating system user-wide default:

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

Each Kubernetes resource managed by Pulumi will have a link in the [Pulumi Cloud](https://app.pulumi.com")
to view the resource in the cluster. These links are local, and require the client run `kubectl proxy` beforehand to access the resource.

To learn more about `kubectl proxy` check out the [reference docs](https://kubernetes.io/docs/concepts/cluster-administration/proxies/).

## Configuration

The Kubernetes provider accepts the following configuration settings. These can be provided to the default Kubernetes provider via `pulumi config set kubernetes:<option>`, or passed to the constructor of `new kubernetes.Provider` to construct a [specific instance](https://www.pulumi.com/docs/concepts/resources/providers/#explicit-provider-configuration) of the Kubernetes provider.

See the [Provider configuration][provider-args] docs for a complete list of options.

## Server-Side Apply

[Server-Side Apply](https://kubernetes.io/docs/reference/using-api/server-side-apply/) (SSA) is a resource management strategy that was introduced in Kubernetes `v1.13`. Clients using SSA can safely share the management of Kubernetes resources by making the API Server responsible for computing diffs and resolving conflicts.

The [v4 release](https://github.com/pulumi/pulumi-kubernetes/releases/tag/v4.0.1) of the Pulumi Kubernetes provider added support for managing resources using SSA by default. Using SSA provides the following benefits:

1. Kubernetes resources may be safely managed by more than one controller.
2. It is now possible to "Upsert" resources; create the resource if it does not exist, or apply the configuration to an existing resource.
3. It is now possible to patch resources with the Patch resource types in the SDK. Each resource type in the SDK has a corresponding Patch resource.
4. The `last-applied-configuration` annotation is no longer used.

See the [SSA how-to guide][ssa-guide] for more information about using SSA.

## Annotations

A few Pulumi-specific annotations can be applied to Kubernetes resources to control aspects of how Pulumi deploys and manages them:

### pulumi.com/skipAwait

Controls Pulumi's behavior while waiting for resources to become ready.
When set to "true" Pulumi will create the resource but will not wait for it to become ready.

{{% notes type="warning" %}}
A small number of resources (Deployments, DaemonSets, StatefulSets, Pods and Namespaces) currently respect the `skipAwait` annotation during deletion and do not wait for deletion to succeed.
However, using `skipAwait` during deletion is not recommended when server-side apply is enabled because it can lead to race conditions during replacement.
The current behavior is considered buggy and these resource may change in the future to no longer respect the `skipAwait` annotation.

The `pulumi.com/deletionPropagationPolicy` annotation, described below, is almost always the preferred way to delete something quickly and safely.
{{% /notes %}}

{{% notes type="info" %}}
Pulumi does not have a concept of "readiness" for all resources by default, and in many cases it will assume a resource is immediately ready even without a `skipAwait` annotation.
This can cause problems if dependent resources do depend on readiness.

You can use the `waitFor` annotation (described below), or you can try running your program with the environment variable `PULUMI_K8S_AWAIT_ALL=true`, to have Pulumi wait for arbitrary resources to become ready.
{{% /notes %}}

### pulumi.com/deletionPropagationPolicy

By default Pulumi uses [foreground](https://kubernetes.io/docs/concepts/architecture/garbage-collection/#foreground-deletion) cascading deletion, which deletes the resource _and_ all of its dependents.
This ensures all dependent resources are cleaned up, but it can be slow.

The `pulumi.com/deletionPropagationPolicy` annotation configures alternative deletion propagation policies:
1. "background": delete the owner resource and let dependent resources be asynchronously garbage collected.
   This is faster than "foreground" deletion propagation, but dependent resources can remain temporarily or even indefinitely if they are not finalized.
2. "orphan": delete the owner resource and leave dependent resources untouched.
   This can be useful if you want to keep resources around for migration or debugging purposes.
3. "foreground": the default behavior of deleting the resource and all of its dependents.
   This is slower but guarantees all dependents have been cleaned up if it succeeds.

### pulumi.com/waitFor

(New in v4.17.0.)

Defines custom criteria for Pulumi to use when waiting for the resource to become ready.
It accepts three possible forms:

  1. A `kubectl`
     [JSONPath](https://kubernetes.io/docs/reference/kubectl/jsonpath):
     a string prefixed with "jsonpath=" followed by a path expression and an
     optional value.

     If a value is provided, the resource is considered ready when the
     JSONPath expression evaluates to the same value. For example this
     resource expects its "phase" field to have a value of "Running":

         "pulumi.com/waitFor": "jsonpath={.phase}=Running"

     If a value is not provided, the resource will be considered ready when
     any value exists at the given path. This resource will wait until it has
     a webhook configured with a CA bundle:

         "pulumi.com/waitFor": "jsonpath={.webhooks[*].clientConfig.caBundle}"

  2. A string prefixed with "condition=" followed by the type of the
     condition and an optional status. This matches the behavior of
     `kubectl --for=condition=...` and will wait until the resource has a
     matching condition. The expected status defaults to "True" if not
     specified. For example:

         "pulumi.com/waitFor": "condition=Synced"

         "pulumi.com/waitFor": "condition=Reconciling=False"

  3. A string containing a JSON array of one or more "jsonpath=" or
     "condition=" expressions.

         "pulumi.com/waitFor": '["jsonpath={.foo}", "condition=Bar"]'

     The resource will be considered ready when all of the criteria are
     simultaneously met.

  This annotation has no effect if the `pulumi.com/skipAwait` annotation is also present with a value of "true" or "ready".

### pulumi.com/patchForce

(Server-Side Apply option.)

Force override any conflicts for the specified resource.

### pulumi.com/patchFieldManager

(Server-Side Apply option.)

Specify the `FieldManager` name to use for the Server-Side Apply operation.

### pulumi.com/timeoutSeconds

Specifies the number of seconds that the Pulumi Kubernetes provider will wait for the resource to become "ready".
Consider using [custom timeouts](https://www.pulumi.com/docs/concepts/options/customtimeouts/) instead.

### pulumi.com/replaceUnready

If the resource failed to become ready in the previous Pulumi update, replace the resource rather than continuing to wait for it to become ready. Only `batch/v1/Job` currently supports this annotation.

### Others

In addition, the Pulumi provider may write the following annotations onto resources it manages:

- `app.kubernetes.io/managed-by`: Indicates the controller managing a Kubernetes resource. This annotation is not set when using the Server-Side Apply mode.
- `pulumi.com/autonamed`: Indicates that the Pulumi Kubernetes provider decided to autoname the resource (instead of using an explicitly provided `metadata.name`).
