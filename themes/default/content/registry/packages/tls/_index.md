---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-tls/v5.2.2/docs/_index.md
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
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as tls from "@pulumi/tls";

// This example creates a self-signed certificate,
// and uses it to create an AWS IAM Server certificate.
//
// THIS IS NOT RECOMMENDED FOR PRODUCTION SERVICES.
// See the detailed documentation of each resource for further
// security considerations and other practical tradeoffs.
const example = new tls.PrivateKey("example", {algorithm: "ECDSA"});
const exampleSelfSignedCert = new tls.SelfSignedCert("example", {
    keyAlgorithm: example.algorithm,
    privateKeyPem: example.privateKeyPem,
    validityPeriodHours: 12,
    earlyRenewalHours: 3,
    allowedUses: [
        "key_encipherment",
        "digital_signature",
        "server_auth",
    ],
    dnsNames: [
        "example.com",
        "example.net",
    ],
    subject: {
        commonName: "example.com",
        organization: "ACME Examples, Inc",
    },
});
// For example, this can be used to populate an AWS IAM server certificate.
const exampleServerCertificate = new aws.iam.ServerCertificate("example", {
    name: "example_self_signed_cert",
    certificateBody: exampleSelfSignedCert.certPem,
    privateKey: example.privateKeyPem,
});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_aws as aws
import pulumi_tls as tls

# This example creates a self-signed certificate,
# and uses it to create an AWS IAM Server certificate.
#
# THIS IS NOT RECOMMENDED FOR PRODUCTION SERVICES.
# See the detailed documentation of each resource for further
# security considerations and other practical tradeoffs.
example = tls.PrivateKey("example", algorithm="ECDSA")
example_self_signed_cert = tls.SelfSignedCert("example",
    key_algorithm=example.algorithm,
    private_key_pem=example.private_key_pem,
    validity_period_hours=12,
    early_renewal_hours=3,
    allowed_uses=[
        "key_encipherment",
        "digital_signature",
        "server_auth",
    ],
    dns_names=[
        "example.com",
        "example.net",
    ],
    subject={
        "common_name": "example.com",
        "organization": "ACME Examples, Inc",
    })
# For example, this can be used to populate an AWS IAM server certificate.
example_server_certificate = aws.iam.ServerCertificate("example",
    name="example_self_signed_cert",
    certificate_body=example_self_signed_cert.cert_pem,
    private_key=example.private_key_pem)
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;
using Tls = Pulumi.Tls;

return await Deployment.RunAsync(() =>
{
    // This example creates a self-signed certificate,
    // and uses it to create an AWS IAM Server certificate.
    //
    // THIS IS NOT RECOMMENDED FOR PRODUCTION SERVICES.
    // See the detailed documentation of each resource for further
    // security considerations and other practical tradeoffs.
    var example = new Tls.PrivateKey("example", new()
    {
        Algorithm = "ECDSA",
    });

    var exampleSelfSignedCert = new Tls.SelfSignedCert("example", new()
    {
        KeyAlgorithm = example.Algorithm,
        PrivateKeyPem = example.PrivateKeyPem,
        ValidityPeriodHours = 12,
        EarlyRenewalHours = 3,
        AllowedUses = new[]
        {
            "key_encipherment",
            "digital_signature",
            "server_auth",
        },
        DnsNames = new[]
        {
            "example.com",
            "example.net",
        },
        Subject = new Tls.Inputs.SelfSignedCertSubjectArgs
        {
            CommonName = "example.com",
            Organization = "ACME Examples, Inc",
        },
    });

    // For example, this can be used to populate an AWS IAM server certificate.
    var exampleServerCertificate = new Aws.Iam.ServerCertificate("example", new()
    {
        Name = "example_self_signed_cert",
        CertificateBody = exampleSelfSignedCert.CertPem,
        PrivateKey = example.PrivateKeyPem,
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/iam"
	"github.com/pulumi/pulumi-tls/sdk/v5/go/tls"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// This example creates a self-signed certificate,
		// and uses it to create an AWS IAM Server certificate.
		//
		// THIS IS NOT RECOMMENDED FOR PRODUCTION SERVICES.
		// See the detailed documentation of each resource for further
		// security considerations and other practical tradeoffs.
		example, err := tls.NewPrivateKey(ctx, "example", &tls.PrivateKeyArgs{
			Algorithm: pulumi.String("ECDSA"),
		})
		if err != nil {
			return err
		}
		exampleSelfSignedCert, err := tls.NewSelfSignedCert(ctx, "example", &tls.SelfSignedCertArgs{
			KeyAlgorithm:        example.Algorithm,
			PrivateKeyPem:       example.PrivateKeyPem,
			ValidityPeriodHours: pulumi.Int(12),
			EarlyRenewalHours:   pulumi.Int(3),
			AllowedUses: pulumi.StringArray{
				pulumi.String("key_encipherment"),
				pulumi.String("digital_signature"),
				pulumi.String("server_auth"),
			},
			DnsNames: pulumi.StringArray{
				pulumi.String("example.com"),
				pulumi.String("example.net"),
			},
			Subject: &tls.SelfSignedCertSubjectArgs{
				CommonName:   pulumi.String("example.com"),
				Organization: pulumi.String("ACME Examples, Inc"),
			},
		})
		if err != nil {
			return err
		}
		// For example, this can be used to populate an AWS IAM server certificate.
		_, err = iam.NewServerCertificate(ctx, "example", &iam.ServerCertificateArgs{
			Name:            pulumi.String("example_self_signed_cert"),
			CertificateBody: exampleSelfSignedCert.CertPem,
			PrivateKey:      example.PrivateKeyPem,
		})
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
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.tls.PrivateKey;
import com.pulumi.tls.PrivateKeyArgs;
import com.pulumi.tls.SelfSignedCert;
import com.pulumi.tls.SelfSignedCertArgs;
import com.pulumi.tls.inputs.SelfSignedCertSubjectArgs;
import com.pulumi.aws.iam.ServerCertificate;
import com.pulumi.aws.iam.ServerCertificateArgs;
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
        // This example creates a self-signed certificate,
        // and uses it to create an AWS IAM Server certificate.
        //
        // THIS IS NOT RECOMMENDED FOR PRODUCTION SERVICES.
        // See the detailed documentation of each resource for further
        // security considerations and other practical tradeoffs.
        var example = new PrivateKey("example", PrivateKeyArgs.builder()
            .algorithm("ECDSA")
            .build());

        var exampleSelfSignedCert = new SelfSignedCert("exampleSelfSignedCert", SelfSignedCertArgs.builder()
            .keyAlgorithm(example.algorithm())
            .privateKeyPem(example.privateKeyPem())
            .validityPeriodHours(12)
            .earlyRenewalHours(3)
            .allowedUses(
                "key_encipherment",
                "digital_signature",
                "server_auth")
            .dnsNames(
                "example.com",
                "example.net")
            .subject(SelfSignedCertSubjectArgs.builder()
                .commonName("example.com")
                .organization("ACME Examples, Inc")
                .build())
            .build());

        // For example, this can be used to populate an AWS IAM server certificate.
        var exampleServerCertificate = new ServerCertificate("exampleServerCertificate", ServerCertificateArgs.builder()
            .name("example_self_signed_cert")
            .certificateBody(exampleSelfSignedCert.certPem())
            .privateKey(example.privateKeyPem())
            .build());

    }
}
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
      function: tls:getCertificate
      arguments:
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
      function: tls:getCertificate
      arguments:
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