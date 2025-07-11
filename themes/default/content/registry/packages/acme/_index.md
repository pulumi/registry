---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumiverse/pulumi-acme/v0.11.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: ACME
meta_desc: Provides an overview of the ACME Provider for Pulumi.
layout: package
---

The Automated Certificate Management Environment (ACME) is an evolving standard
for the automation of a domain-validated certificate authority. Clients register
themselves on an authority using a private key and contact information, and
answer challenges for domains that they own by supplying response data issued by
the authority via either HTTP or DNS. Via this process, they prove that they own
the domains in question, and can then request certificates for them via the CA.
No part of this process requires user interaction, a traditional blocker in
obtaining a domain validated certificate.

Currently the major ACME CA is [Let's Encrypt][lets-encrypt], but the ACME
support in Pulumi can be configured to use any ACME CA, including an
internal one that is set up using [Boulder][boulder-gh], or another CA
that implements the ACME standard with Let's Encrypt's
[divergences][lets-encrypt-divergences].

[lets-encrypt]: https://letsencrypt.org
[boulder-gh]: https://github.com/letsencrypt/boulder
[lets-encrypt-divergences]: https://github.com/letsencrypt/boulder/blob/master/docs/acme-divergences.md

For more detail on the ACME process, see [here][lets-encrypt-how-it-works]. For
the ACME spec, click [here][about-acme]. Note that as mentioned in the last
paragraph, the ACME provider may [diverge][lets-encrypt-divergences] from the
current ACME spec to account for the real-world divergences that are made by
CAs such as Let's Encrypt.

[lets-encrypt-how-it-works]: https://letsencrypt.org/how-it-works/
[about-acme]: https://ietf-wg-acme.github.io/acme/draft-ietf-acme-acme.html

## Example

{{< chooser language "typescript,python,go,csharp" >}}
{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as acme from "@pulumiverse/acme";
import * as tls from "@pulumi/tls";

const privateKey = new tls.index.PrivateKey("privateKey", {algorithm: "RSA"});

const reg = new acme.index.Registration("reg", {
    accountKeyPem: privateKey.privateKeyPem,
    emailAddress: "nobody@example.com",
});
const certificate = new acme.index.Certificate("certificate", {
    accountKeyPem: reg.accountKeyPem,
    commonName: "www.example.com",
    subjectAlternativeNames: ["www2.example.com"],
    dnsChallenge: [{
        provider: "route53",
    }],
});
```
 
{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
import pulumiverse_acme as acme
import pulumi_tls as tls

private_key = tls.index.PrivateKey("privateKey", algorithm="RSA")

reg = acme.index.Registration("reg",
    account_key_pem=private_key.private_key_pem,
    email_address="nobody@example.com")

certificate = acme.index.Certificate("certificate",
    account_key_pem=reg.account_key_pem,
    common_name="www.example.com",
    subject_alternative_names=["www2.example.com"],
    dns_challenge=[{
        provider="route53",
    }])
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi-tls/sdk/v4/go/tls"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-acme/sdk/go/acme"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		privateKey, err := tls.NewPrivateKey(ctx, "privateKey", &tls.PrivateKeyArgs{
			Algorithm: "RSA",
		})
		if err != nil {
			return err
		}
		reg, err := acme.NewRegistration(ctx, "reg", &acme.RegistrationArgs{
			AccountKeyPem: privateKey.PrivateKeyPem,
			EmailAddress:  "nobody@example.com",
		})
		if err != nil {
			return err
		}
		_, err = acme.NewCertificate(ctx, "certificate", &acme.CertificateArgs{
			AccountKeyPem: reg.AccountKeyPem,
			CommonName:    "www.example.com",
			SubjectAlternativeNames: []string{
				"www2.example.com",
			},
			DnsChallenge: []map[string]interface{}{
				map[string]interface{}{
					"provider": "route53",
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Acme = Pulumiverse.Acme;
using Tls = Pulumi.Tls;

return await Deployment.RunAsync(() => 
{
    var privateKey = new Tls.Index.PrivateKey("privateKey", new()
    {
        Algorithm = "RSA",
    });

    var reg = new Acme.Index.Registration("reg", new()
    {
        AccountKeyPem = privateKey.PrivateKeyPem,
        EmailAddress = "nobody@example.com",
    });

    var certificate = new Acme.Index.Certificate("certificate", new()
    {
        AccountKeyPem = reg.AccountKeyPem,
        CommonName = "www.example.com",
        SubjectAlternativeNames = new[]
        {
            "www2.example.com",
        },
        DnsChallenge = new[]
        {
            
            {
                { "provider", "route53" },
            },
        },
    });

});
```

{{% /choosable %}}

{{< /chooser >}}
