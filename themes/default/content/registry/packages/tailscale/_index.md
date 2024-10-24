---
title: Tailscale Provider
meta_desc: Provides an overview on how to configure the Pulumi Tailscale provider.
layout: package
---
## Installation

The tailscale provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/tailscale`](https://www.npmjs.com/package/@pulumi/tailscale)
* Python: [`pulumi-tailscale`](https://pypi.org/project/pulumi-tailscale/)
* Go: [`github.com/pulumi/pulumi-tailscale/sdk/go/tailscale`](https://github.com/pulumi/pulumi-tailscale)
* .NET: [`Pulumi.Tailscale`](https://www.nuget.org/packages/Pulumi.Tailscale)
* Java: [`com.pulumi/tailscale`](https://central.sonatype.com/artifact/com.pulumi/tailscale)
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    tailscale:apiKey:
        value: 12345
    tailscale:tailnet:
        value: example.com

```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    tailscale:apiKey:
        value: 12345
    tailscale:tailnet:
        value: example.com

```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    tailscale:apiKey:
        value: 12345
    tailscale:tailnet:
        value: example.com

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    tailscale:apiKey:
        value: 12345
    tailscale:tailnet:
        value: example.com

```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    tailscale:apiKey:
        value: 12345
    tailscale:tailnet:
        value: example.com

```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    tailscale:apiKey:
        value: 12345
    tailscale:tailnet:
        value: example.com

```

{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

- `apiKey` (String, Sensitive) The API key to use for authenticating requests to the API. Can be set via the TAILSCALE_API_KEY environment variable. Conflicts with 'oauth_client_id' and 'oauth_client_secret'.
- `baseUrl` (String) The base URL of the Tailscale API. Defaults to <https://api.tailscale.com>. Can be set via the TAILSCALE_BASE_URL environment variable.
- `oauthClientId` (String) The OAuth application's ID when using OAuth client credentials. Can be set via the TAILSCALE_OAUTH_CLIENT_ID environment variable. Both 'oauth_client_id' and 'oauth_client_secret' must be set. Conflicts with 'api_key'.
- `oauthClientSecret` (String, Sensitive) The OAuth application's secret when using OAuth client credentials. Can be set via the TAILSCALE_OAUTH_CLIENT_SECRET environment variable. Both 'oauth_client_id' and 'oauth_client_secret' must be set. Conflicts with 'api_key'.
- `scopes` (List of String) The OAuth 2.0 scopes to request when for the access token generated using the supplied OAuth client credentials. See <https://tailscale.com/kb/1215/oauth-clients/#scopes> for available scopes. Only valid when both 'oauth_client_id' and 'oauth_client_secret' are set.
- `tailnet` (String) The organization name of the Tailnet in which to perform actions. Can be set via the TAILSCALE_TAILNET environment variable. Default is the tailnet that owns API credentials passed to the provider.
- `userAgent` (String) User-Agent header for API requests.