---
title: Self Signed Certificate
meta_desc: Use Pulumi's Self Signed Certificate Package to quickly create a self signed certificate.
layout: overview
---

Pulumi's Self Signed Certificate Package makes it simple for you to quickly create a self signed
certificate. The guide below will quickly you through how to provision a self signed certificate
with the minimal set of inputs. Please refer to the API Docs for more detailed information on this
Packages's usage.

## Quick Start

The following steps will get you started with a self signed certificate.

### Configure Environment

Before you get started using Pulumi, let's run through a few quick steps to ensure your environment is set up correctly.

### Install Pulumi

{{< install-pulumi >}}
{{% notes "info" %}}
All Windows examples in this tutorial assume you are running in PowerShell.
{{% /notes %}}
{{< /install-pulumi >}}

Next, install the required language runtime, if you have not already.

### Install Language Runtime

#### Choose Your Language

{{< chooser language "typescript,python,go,csharp,yaml" / >}}

{{% choosable language "typescript" %}}
{{< install-node >}}
{{% /choosable %}}

{{% choosable language python %}}
{{< install-python >}}
{{% /choosable %}}

{{% choosable language go %}}
{{< install-go >}}
{{% /choosable %}}

{{% choosable language "csharp,fsharp,visualbasic" %}}
{{< install-dotnet >}}
{{% /choosable %}}

{{% choosable language "yaml" %}}
{{< install-yaml >}}
{{% /choosable %}}

### Create New Project

Now that you have set up your environment by installing Pulumi and installing your preferred language runtime,
let's create your Pulumi program.

{{< chooser language "typescript,python,go,csharp,yaml" / >}}
{{% choosable language typescript %}}

```bash
$ mkdir tls-self-signed-cert-quickstart && cd tls-self-signed-cert-quickstart
$ pulumi new typescript
```

{{% /choosable %}}
{{% choosable language python %}}

```bash
$ mkdir tls-self-signed-cert-quickstart && cd tls-self-signed-cert-quickstart
$ pulumi new python
```

{{% /choosable %}}
{{% choosable language go %}}

```bash
# from within your $GOPATH
$ mkdir tls-self-signed-cert-quickstart && cd tls-self-signed-cert-quickstart
$ pulumi new go
```

{{% /choosable %}}
{{% choosable language csharp %}}

```bash
$ mkdir tls-self-signed-cert-quickstart && cd tls-self-signed-cert-quickstart
$ pulumi new csharp
```

{{% /choosable %}}
{{% choosable language yaml %}}

```bash
$ mkdir tls-self-signed-cert-quickstart && cd tls-self-signed-cert-quickstart
$ pulumi new yaml
```

{{% /choosable %}}

The [`pulumi new`]({{< relref "/docs/reference/cli/pulumi_new" >}}) command creates a new Pulumi project with some basic scaffolding based on the cloud and language specified.

{{< cli-note >}}

After logging in, the CLI will proceed with walking you through creating a new project.

First, you will be asked for a project name and description. Hit `ENTER` to accept the default values or specify new values.

Next, you will be asked for the name of a stack. Hit `ENTER` to accept the default value of `dev`.

> What are [projects]({{< relref "/docs/intro/concepts/project" >}}) and [stacks]({{< relref "/docs/intro/concepts/stack" >}})? Pulumi projects and stacks let you organize Pulumi code. Consider a Pulumi _project_ to be analogous to a GitHub repo---a single place for code---and a _stack_ to be an instance of that code with a separate configuration. For instance, _Project Foo_ may have multiple stacks for different development environments (Dev, Test, or Prod), or perhaps for different cloud configurations (geographic region for example). See [Organizing Projects and Stacks]({{< relref "/docs/guides/organizing-projects-stacks" >}}) for some best practices on organizing your Pulumi projects and stacks.

{{% choosable language "typescript" %}}

After some dependency installations from `npm`, your project and stack will be ready.

#### Install the Self Signed Certificate Package

Next you will need to install the Self Signed Certificate Package so you can use it in your program.

##### Yarn

```bash
$ yarn add @pulumi/tls-self-signed-cert
```

##### NPM

```bash
$ npm install @pulumi/tls-self-signed-cert
```

{{% /choosable %}}

{{% choosable language python %}}

After the command completes, the project and stack will be ready.

#### Install the Self Signed Certificate Package

Next you will need to install the Self Signed Certificate Package so you can use it in your program.

```bash
$ pip3 install pulumi_tls_self_signed_cert
```

{{% /choosable %}}

{{% choosable language go %}}

After the command completes, the project and stack will be ready.

#### Install the Self Signed Certificate Package

Next you will need to install the Self Signed Certificate Package so you can use it in your program.

```bash
$ go get -u github.com/pulumi/pulumi-tls-self-signed-cert/sdk
```

{{% /choosable %}}

{{% choosable language "csharp,fsharp,visualbasic" %}}

After the command completes, the project and stack will be ready.

#### Install the Self Signed Certificate Package

Next you will need to install the Self Signed Certificate Package so you can use it in your program.

```bash
$ dotnet add package Pulumi.TlsSelfSignedCert
```

{{% /choosable %}}

### Update Code

Now that you have all your dependencies installed and your project configured, you can now add the code that will
provision your self signed certificate.

{{< chooser language "typescript,python,go,csharp,yaml" / >}}
{{% choosable language typescript %}}

Replace your `index.ts` with the following:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as tls_self_signed_cert from "@pulumi/tls-self-signed-cert";

const cert = new tls_self_signed_cert.SelfSignedCertificate("cert", {
    dnsName: "cert.example.com",
    validityPeriodHours: 807660,
    localValidityPeriodHours: 17520,
    subject: {
        commonName: "example-cert",
        organization: "example-cert LLC",
    },
});
export const pem = cert.pem;
export const privateKey = cert.privateKey;
export const caCert = cert.caCert;
```

{{% /choosable %}}
{{% choosable language python %}}

Replace your `__main__.py` with the following:

```python
import pulumi
import pulumi_tls_self_signed_cert as tls_self_signed_cert

cert = tls_self_signed_cert.SelfSignedCertificate("cert",
    dns_name="cert.example.com",
    validity_period_hours=807660,
    local_validity_period_hours=17520,
    subject=%!v(PANIC=Format method: interface conversion: interface {} is json.RawMessage, not python.PackageInfo))
pulumi.export("pem", cert.pem)
pulumi.export("privateKey", cert.private_key)
pulumi.export("caCert", cert.ca_cert)

```

{{% /choosable %}}
{{% choosable language go %}}

Replace your `main.go` with the following:

```go
package main

import (
	selfSignedCert "github.com/pulumi/pulumi-tls-self-signed-cert/sdk/go/tls-self-signed-cert"
	"github.com/pulumi/pulumi-tls/sdk/v4/go/tls"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cert, err := selfSignedCert.NewSelfSignedCertificate(ctx, "cert", &selfSignedCert.SelfSignedCertificateArgs{
			DnsName:                  pulumi.String("cert.example.com"),
			ValidityPeriodHours:      pulumi.Int(807660),
			LocalValidityPeriodHours: pulumi.Int(17520),
			Subject: tls.SelfSignedCertSubjectArgs{
				CommonName:   pulumi.String("example-cert"),
				Organization: pulumi.String("example-cert LLC"),
			},
		})
		if err != nil {
			return err
		}

		ctx.Export("pem", cert.Pem)
		ctx.Export("privateKey", cert.PrivateKey)
		ctx.Export("caCert", cert.CaCert)
	})
}
```

{{% /choosable %}}
{{% choosable language csharp %}}

Replace your `Program.cs` with the following:

```csharp
using System.Collections.Generic;
using Pulumi;
using TlsSelfSignedCert = Pulumi.TlsSelfSignedCert;

return await Deployment.RunAsync(() =>
{
    var cert = new TlsSelfSignedCert.SelfSignedCertificate("cert", new()
    {
        DnsName = "cert.example.com",
        ValidityPeriodHours = 807660,
        LocalValidityPeriodHours = 17520,
        Subject = %!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference),
    });

    return new Dictionary<string, object?>
    {
        ["pem"] = cert.Pem,
        ["privateKey"] = cert.PrivateKey,
        ["caCert"] = cert.CaCert,
    };
});
```

{{% /choosable %}}
{{% choosable language yaml %}}

Replace your `Pulumi.yaml` with the following:

```yaml
name: tls-self-signed-cert
runtime: yaml
resources:
    cert:
        type: "tls-self-signed-cert:index:SelfSignedCertificate"
        properties:
            dnsName: "cert.example.com"
            validityPeriodHours: 807660
            localValidityPeriodHours: 17520
            subject:
                commonName: "example-cert"
                organization: "example-cert LLC"
outputs:
    pem: ${cert.pem}
    privateKey: ${cert.privateKey}
    caCert: ${cert.caCert}
```

{{% /choosable %}}

### Deploy

Once you have updated your code you are ready to create your Self Signed Certificate. To do so, run the the following command:

```bash
$ pulumi up
```

First Pulumi will perform a preview showing you exactly what will be created. Once the preview is complete Pulumi will ask you if you want to continue.
Select `yes` to proceed to actually provisioning the service.

All the different resources need to create a self signed certificate will be created.

### (Optional) Destroy

You can destroy all the resources by running `pulumi destroy`.
