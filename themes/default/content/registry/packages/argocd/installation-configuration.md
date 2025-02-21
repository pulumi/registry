---
# WARNING: this file was fetched from https://raw.githubusercontent.com/three141/pulumi-argocd/v1.0.1/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: ArgoCD Provider Installation & Configuration
meta_desc: Provides an overview on how to configure credentials for the ArgoCD provider for Pulumi.
layout: package
---

The Argo CD Resource Provider lets you manage [Argo CD](https://argoproj.github.io/cd/) resources.

## Installation

The ArgoCD provider for Pulumi is available as a package in most Pulumi languages:

- JavaScript/TypeScript: [`@three14/pulumi-argocd`](https://www.npmjs.com/package/@three14/pulumi-argocd)
- Python: [`pulumi-argocd`](https://pypi.org/project/pulumi-argocd/)
- Go: [`github.com/Three14/pulumi-argocd/sdk/go/argocd`](https://pkg.go.dev/github.com/Three141/pulumi-argocd/sdk/go/argocd)
- .NET: [`Three14.Argocd`](https://www.nuget.org/packages/Three14.Argocd)

### Provider Binary

The ArgoCD provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```sh
pulumi plugin install resource argocd <version> --server 'https://github.com/Three141/pulumi-argocd/releases/download/v${VERSION}/'
```

Replace the version string with your desired version.

## Configuration

You must configure the ArgoCD provider for Pulumi with your ArgoCD credentials before the provider can be used to access and manage your ArgoCD instance. You can use one of the following:

The following configuration points are available for the `argocd` provider:

- `argocd:auth_token` (String, Sensitive) ArgoCD authentication token, takes precedence over `username`/`password`. Can be set through the `ARGOCD_AUTH_TOKEN` environment variable.
- `argocd:cert_file` (String) Additional root CA certificates file to add to the client TLS connection pool.
- `argocd:client_cert_file` (String) Client certificate.
- `argocd:client_cert_key` (String) Client certificate key.
- `argocd:config_path` (String) Override the default config path of `$HOME/.config/argocd/config`. Only relevant when `argocd:use_local_config`. Can be set through the `ARGOCD_CONFIG_PATH` environment variable.
- `argocd:context` (String) Context to choose when using a local ArgoCD config file. Only relevant when `use_local_config`. Can be set through `ARGOCD_CONTEXT` environment variable.
- `argocd:core` (Boolean) Configure direct access using Kubernetes API server.

  **Warning**: this feature works by starting a local ArgoCD API server that talks directly to the Kubernetes API using the **current context in the default kubeconfig** (`~/.kube/config`). This behavior cannot be overridden using either environment variables or the `kubernetes` block in the provider configuration at present).

  If the server fails to start (e.g. your kubeconfig is misconfigured) then the provider will fail as a result of the `argocd` module forcing it to exit and no logs will be available to help you debug this. The error message will be similar to
  > `The plugin encountered an error, and failed to respond to the plugin.(*GRPCProvider).ReadResource call. The plugin logs may contain more details.`

  To debug this, you will need to login via the ArgoCD CLI using `argocd login --core` and then running an operation. E.g. `argocd app list`.
- `argocd:grpc_web` (Boolean) Whether to use gRPC web proxy client. Useful if Argo CD server is behind proxy which does not support HTTP2.
- `argocd:grpc_web_root_path` (String) Use the gRPC web proxy client and set the web root, e.g. `argo-cd`. Useful if the Argo CD server is behind a proxy at a non-root path.
- `argocd:headers` (Set of String) Additional headers to add to each request to the ArgoCD server.
- `argocd:insecure` (Boolean) Whether to skip TLS server certificate. Can be set through the `ARGOCD_INSECURE` environment variable.
- `argocd:password` (String, Sensitive) Authentication password. Can be set through the `ARGOCD_AUTH_PASSWORD` environment variable.
- `argocd:plain_text` (Boolean) Whether to initiate an unencrypted connection to ArgoCD server.
- `argocd:port_forward` (Boolean) Connect to a random argocd-server port using port forwarding.
- `argocd:port_forward_with_namespace` (String) Namespace name which should be used for port forwarding.
- `argocd:server_addr` (String) ArgoCD server address with port. Can be set through the `ARGOCD_SERVER` environment variable.
- `argocd:use_local_config` (Boolean) Use the authentication settings found in the local config file. Useful when you have previously logged in using SSO. Conflicts with `auth_token`, `username` and `password`.
- `argocd:user_agent` (String) User-Agent request header override.
- `argocd:username` (String) Authentication username. Can be set through the `ARGOCD_AUTH_USERNAME` environment variable.
