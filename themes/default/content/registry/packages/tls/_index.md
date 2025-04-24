---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-tls/v5.2.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Tls Provider
meta_desc: Provides an overview on how to configure the Pulumi Tls provider.
layout: package
---
## Installation

The Tls provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/tls`](https://www.npmjs.com/package/@pulumi/tls)
* Python: [`pulumi-tls`](https://pypi.org/project/pulumi-tls/)
* Go: [`github.com/pulumi/pulumi-tls/sdk/v5/go/tls`](https://github.com/pulumi/pulumi-tls)
* .NET: [`Pulumi.Tls`](https://www.nuget.org/packages/Pulumi.Tls)
* Java: [`com.pulumi/tls`](https://central.sonatype.com/artifact/com.pulumi/tls)
## Overview

The TLS provider provides utilities for working with *Transport Layer Security*
keys and certificates. It provides resources that
allow private keys, certificates and certificate requests to be
created as part of a Pulumi deployment.

Another name for Transport Layer Security is *Secure Sockets Layer*,
or SSL. TLS and SSL are equivalent when considering the resources
managed by this provider.

This provider is not particularly useful on its own, but it can be
used to create certificates and credentials that can then be used
with other providers when creating resources that expose TLS
services or that themselves provision TLS certificates.

Use the navigation to the left to read about the available resources.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language python %}}
```python
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language go %}}
```go
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
resources:
  # This example creates a self-signed certificate,
  # and uses it to create an AWS IAM Server certificate.
  #
  # THIS IS NOT RECOMMENDED FOR PRODUCTION SERVICES.
  # See the detailed documentation of each resource for further
  # security considerations and other practical tradeoffs.
  example:
    type: tls:PrivateKey
    properties:
      algorithm: ECDSA
  exampleSelfSignedCert:
    type: tls:SelfSignedCert
    name: example
    properties:
      keyAlgorithm: ${example.algorithm}
      privateKeyPem: ${example.privateKeyPem}
      validityPeriodHours: 12 # Generate a new certificate if Pulumi is run within three
      #   # hours of the certificate's expiration time.
      earlyRenewalHours: 3 # Reasonable set of uses for a server SSL certificate.
      allowedUses:
        - key_encipherment
        - digital_signature
        - server_auth
      dnsNames:
        - example.com
        - example.net
      subject:
        commonName: example.com
        organization: ACME Examples, Inc
  # For example, this can be used to populate an AWS IAM server certificate.
  exampleServerCertificate:
    type: aws:iam:ServerCertificate
    name: example
    properties:
      name: example_self_signed_cert
      certificateBody: ${exampleSelfSignedCert.certPem}
      privateKey: ${example.privateKeyPem}
```
{{% /choosable %}}
{{% choosable language java %}}
```java
Example currently unavailable in this language
```
{{% /choosable %}}
{{< /chooser >}}
### Configuring Proxy

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as tls from "@pulumi/tls";

const test = tls.getCertificate({
    url: "https://example.com",
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python

```
```python
import pulumi
import pulumi_tls as tls

test = tls.get_certificate(url="https://example.com")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Tls = Pulumi.Tls;

return await Deployment.RunAsync(() =>
{
    var test = Tls.GetCertificate.Invoke(new()
    {
        Url = "https://example.com",
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go

```
```go
package main

import (
	"github.com/pulumi/pulumi-tls/sdk/v5/go/tls"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := tls.GetCertificate(ctx, &tls.GetCertificateArgs{
			Url: pulumi.StringRef("https://example.com"),
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml

```
```yaml
variables:
  test:
    fn::invoke:
      Function: tls:getCertificate
      Arguments:
        url: https://example.com
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.tls.TlsFunctions;
import com.pulumi.tls.inputs.GetCertificateArgs;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        final var test = TlsFunctions.getCertificate(GetCertificateArgs.builder()
            .url("https://example.com")
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as tls from "@pulumi/tls";

const test = tls.getCertificate({
    url: "https://example.com",
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python

```
```python
import pulumi
import pulumi_tls as tls

test = tls.get_certificate(url="https://example.com")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Tls = Pulumi.Tls;

return await Deployment.RunAsync(() =>
{
    var test = Tls.GetCertificate.Invoke(new()
    {
        Url = "https://example.com",
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go

```
```go
package main

import (
	"github.com/pulumi/pulumi-tls/sdk/v5/go/tls"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := tls.GetCertificate(ctx, &tls.GetCertificateArgs{
			Url: pulumi.StringRef("https://example.com"),
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml

```
```yaml
variables:
  test:
    fn::invoke:
      Function: tls:getCertificate
      Arguments:
        url: https://example.com
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.tls.TlsFunctions;
import com.pulumi.tls.inputs.GetCertificateArgs;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        final var test = TlsFunctions.getCertificate(GetCertificateArgs.builder()
            .url("https://example.com")
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

- `proxy` (Block List) Proxy used by resources and functions that connect to external endpoints. (see below for nested schema)

<a id="nestedblock--proxy"></a>
### Nested Schema for `proxy`

Optional:

- `fromEnv` (Boolean) When `true` the provider will discover the proxy configuration from environment variables. This is based upon [`http.ProxyFromEnvironment`](https://pkg.go.dev/net/http#ProxyFromEnvironment) and it supports the same environment variables (default: `true`).
- `password` (String, Sensitive) Password used for Basic authentication against the Proxy.
- `url` (String) URL used to connect to the Proxy. Accepted schemes are: `http`, `https`, `socks5`.
- `username` (String) Username (or Token) used for Basic authentication against the Proxy.
## Limitations
### `ECDSA` with `P224` elliptic curve

When using `ECDSA` with `P224`, all the (computed) attributes
that have to do with [OpenSSH](https://www.openssh.com/) will have a value of `""` (empty string).
This applies to different resources and functions offered by this provider,
like the `tls.PrivateKey` resource or the `tls.getPublicKey` function.

The attributes affected are:

* `.public_key_openssh`
* `.private_key_openssh`
* `.public_key_fingerprint_md5`
* `.public_key_fingerprint_sha256`

This is because the SSH ECC Algorithm Integration ([RFC 5656](https://datatracker.ietf.org/doc/html/rfc5656))
restricts support for elliptic curves to "nistp256", "nistp384" and "nistp521".