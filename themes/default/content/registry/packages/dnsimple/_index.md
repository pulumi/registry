---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-dnsimple/v5.0.1/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi/pulumi-dnsimple/blob/v5.0.1/docs/_index.md
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: DNSimple Provider
meta_desc: Provides an overview on how to configure the Pulumi DNSimple provider.
layout: package
---

## Installation

The DNSimple provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/dnsimple`](https://www.npmjs.com/package/@pulumi/dnsimple)
* Python: [`pulumi-dnsimple`](https://pypi.org/project/pulumi-dnsimple/)
* Go: [`github.com/pulumi/pulumi-dnsimple/sdk/v5/go/dnsimple`](https://github.com/pulumi/pulumi-dnsimple)
* .NET: [`Pulumi.Dnsimple`](https://www.nuget.org/packages/Pulumi.Dnsimple)
* Java: [`com.pulumi/dnsimple`](https://central.sonatype.com/artifact/com.pulumi/dnsimple)

## Overview

---
## page_title: "Provider: DNSimple"
The DNSimple provider allows you to manage DNS records, domains, certificates, and other DNSimple resources using Pulumi.

This provider enables you to treat your DNS and domain infrastructure as code, making it easier to version, review, and manage your DNSimple resources alongside your other infrastructure.
## Authentication

The provider requires authentication credentials to interact with the DNSimple API. You can provide credentials in several ways:

1. **Provider configuration** (recommended for development)
2. **Environment variables** (recommended for CI/CD and production)
### Using Provider Configuration

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    dnsimple:account:
        value: 'TODO: var.dnsimple_account'
    dnsimple:token:
        value: 'TODO: var.dnsimple_token'

```
### Using Environment Variables

```bash
export DNSIMPLE_TOKEN="your-api-token"
export DNSIMPLE_ACCOUNT="your-account-id"
```

See the Configuration Reference section below for all configuration options.
## Example Usage

Configure the provider:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    dnsimple:account:
        value: 'TODO: var.dnsimple_account'
    dnsimple:token:
        value: 'TODO: var.dnsimple_token'

```

```typescript
import * as pulumi from "@pulumi/pulumi";

const config = new pulumi.Config();
// DNSimple API Token
const dnsimpleToken = config.require("dnsimpleToken");
// DNSimple Account ID
const dnsimpleAccount = config.require("dnsimpleAccount");
```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    dnsimple:account:
        value: 'TODO: var.dnsimple_account'
    dnsimple:token:
        value: 'TODO: var.dnsimple_token'

```

```python
import pulumi

config = pulumi.Config()
# DNSimple API Token
dnsimple_token = config.require("dnsimpleToken")
# DNSimple Account ID
dnsimple_account = config.require("dnsimpleAccount")
```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    dnsimple:account:
        value: 'TODO: var.dnsimple_account'
    dnsimple:token:
        value: 'TODO: var.dnsimple_token'

```

```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;

return await Deployment.RunAsync(() =>
{
    var config = new Config();
    // DNSimple API Token
    var dnsimpleToken = config.Require("dnsimpleToken");
    // DNSimple Account ID
    var dnsimpleAccount = config.Require("dnsimpleAccount");
});

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    dnsimple:account:
        value: 'TODO: var.dnsimple_account'
    dnsimple:token:
        value: 'TODO: var.dnsimple_token'

```

```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		// DNSimple API Token
		dnsimpleToken := cfg.Require("dnsimpleToken")
		// DNSimple Account ID
		dnsimpleAccount := cfg.Require("dnsimpleAccount")
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
config:
    dnsimple:account:
        value: 'TODO: var.dnsimple_account'
    dnsimple:token:
        value: 'TODO: var.dnsimple_token'

```

```yaml
configuration:
  dnsimpleToken:
    type: string
  dnsimpleAccount:
    type: string
```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    dnsimple:account:
        value: 'TODO: var.dnsimple_account'
    dnsimple:token:
        value: 'TODO: var.dnsimple_token'

```

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
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
        final var config = ctx.config();
        final var dnsimpleToken = config.get("dnsimpleToken");
        final var dnsimpleAccount = config.get("dnsimpleAccount");
    }
}
```

{{% /choosable %}}
{{< /chooser >}}

Now use the available resources to perform actions like managing DNS records or registering domains.

To manage your DNS records:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as dnsimple from "@pulumi/dnsimple";

// Create a zone
const example = new dnsimple.Zone("example", {name: "example.com"});
// Create DNS records
const www = new dnsimple.ZoneRecord("www", {
    zoneName: example.name,
    name: "www",
    value: "192.0.2.1",
    type: "A",
    ttl: 3600,
});
const apex = new dnsimple.ZoneRecord("apex", {
    zoneName: example.name,
    name: "",
    value: "192.0.2.1",
    type: "A",
    ttl: 3600,
});
```

{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_dnsimple as dnsimple

# Create a zone
example = dnsimple.Zone("example", name="example.com")
# Create DNS records
www = dnsimple.ZoneRecord("www",
    zone_name=example.name,
    name="www",
    value="192.0.2.1",
    type="A",
    ttl=3600)
apex = dnsimple.ZoneRecord("apex",
    zone_name=example.name,
    name="",
    value="192.0.2.1",
    type="A",
    ttl=3600)
```

{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using DNSimple = Pulumi.DNSimple;

return await Deployment.RunAsync(() =>
{
    // Create a zone
    var example = new DNSimple.Zone("example", new()
    {
        Name = "example.com",
    });

    // Create DNS records
    var www = new DNSimple.ZoneRecord("www", new()
    {
        ZoneName = example.Name,
        Name = "www",
        Value = "192.0.2.1",
        Type = "A",
        Ttl = 3600,
    });

    var apex = new DNSimple.ZoneRecord("apex", new()
    {
        ZoneName = example.Name,
        Name = "",
        Value = "192.0.2.1",
        Type = "A",
        Ttl = 3600,
    });

});

```

{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-dnsimple/sdk/v5/go/dnsimple"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a zone
		example, err := dnsimple.NewZone(ctx, "example", &dnsimple.ZoneArgs{
			Name: pulumi.String("example.com"),
		})
		if err != nil {
			return err
		}
		// Create DNS records
		_, err = dnsimple.NewZoneRecord(ctx, "www", &dnsimple.ZoneRecordArgs{
			ZoneName: example.Name,
			Name:     pulumi.String("www"),
			Value:    pulumi.String("192.0.2.1"),
			Type:     pulumi.String("A"),
			Ttl:      pulumi.Int(3600),
		})
		if err != nil {
			return err
		}
		_, err = dnsimple.NewZoneRecord(ctx, "apex", &dnsimple.ZoneRecordArgs{
			ZoneName: example.Name,
			Name:     pulumi.String(""),
			Value:    pulumi.String("192.0.2.1"),
			Type:     pulumi.String("A"),
			Ttl:      pulumi.Int(3600),
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
  # Create a zone
  example:
    type: dnsimple:Zone
    properties:
      name: example.com
  # Create DNS records
  www:
    type: dnsimple:ZoneRecord
    properties:
      zoneName: ${example.name}
      name: www
      value: 192.0.2.1
      type: A
      ttl: 3600
  apex:
    type: dnsimple:ZoneRecord
    properties:
      zoneName: ${example.name}
      name: ""
      value: 192.0.2.1
      type: A
      ttl: 3600
```

{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.dnsimple.Zone;
import com.pulumi.dnsimple.ZoneArgs;
import com.pulumi.dnsimple.ZoneRecord;
import com.pulumi.dnsimple.ZoneRecordArgs;
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
        // Create a zone
        var example = new Zone("example", ZoneArgs.builder()
            .name("example.com")
            .build());

        // Create DNS records
        var www = new ZoneRecord("www", ZoneRecordArgs.builder()
            .zoneName(example.name())
            .name("www")
            .value("192.0.2.1")
            .type("A")
            .ttl(3600)
            .build());

        var apex = new ZoneRecord("apex", ZoneRecordArgs.builder()
            .zoneName(example.name())
            .name("")
            .value("192.0.2.1")
            .type("A")
            .ttl(3600)
            .build());

    }
}
```

{{% /choosable %}}
{{< /chooser >}}

To register a domain:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as dnsimple from "@pulumi/dnsimple";

// Create a contact for domain registration
const registrant = new dnsimple.Contact("registrant", {
    label: "Main Contact",
    firstName: "John",
    lastName: "Doe",
    organizationName: "Example Inc",
    address1: "123 Main Street",
    city: "San Francisco",
    stateProvince: "California",
    postalCode: "94105",
    country: "US",
    phone: "+1.4155551234",
    email: "john@example.com",
});
// Register a domain
const exampleCom = new dnsimple.RegisteredDomain("example_com", {
    name: "example.com",
    contactId: registrant.id,
    autoRenewEnabled: true,
    whoisPrivacyEnabled: true,
    transferLockEnabled: true,
});
```

{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_dnsimple as dnsimple

# Create a contact for domain registration
registrant = dnsimple.Contact("registrant",
    label="Main Contact",
    first_name="John",
    last_name="Doe",
    organization_name="Example Inc",
    address1="123 Main Street",
    city="San Francisco",
    state_province="California",
    postal_code="94105",
    country="US",
    phone="+1.4155551234",
    email="john@example.com")
# Register a domain
example_com = dnsimple.RegisteredDomain("example_com",
    name="example.com",
    contact_id=registrant.id,
    auto_renew_enabled=True,
    whois_privacy_enabled=True,
    transfer_lock_enabled=True)
```

{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using DNSimple = Pulumi.DNSimple;

return await Deployment.RunAsync(() =>
{
    // Create a contact for domain registration
    var registrant = new DNSimple.Contact("registrant", new()
    {
        Label = "Main Contact",
        FirstName = "John",
        LastName = "Doe",
        OrganizationName = "Example Inc",
        Address1 = "123 Main Street",
        City = "San Francisco",
        StateProvince = "California",
        PostalCode = "94105",
        Country = "US",
        Phone = "+1.4155551234",
        Email = "john@example.com",
    });

    // Register a domain
    var exampleCom = new DNSimple.RegisteredDomain("example_com", new()
    {
        Name = "example.com",
        ContactId = registrant.Id,
        AutoRenewEnabled = true,
        WhoisPrivacyEnabled = true,
        TransferLockEnabled = true,
    });

});

```

{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-dnsimple/sdk/v5/go/dnsimple"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a contact for domain registration
		registrant, err := dnsimple.NewContact(ctx, "registrant", &dnsimple.ContactArgs{
			Label:            pulumi.String("Main Contact"),
			FirstName:        pulumi.String("John"),
			LastName:         pulumi.String("Doe"),
			OrganizationName: pulumi.String("Example Inc"),
			Address1:         pulumi.String("123 Main Street"),
			City:             pulumi.String("San Francisco"),
			StateProvince:    pulumi.String("California"),
			PostalCode:       pulumi.String("94105"),
			Country:          pulumi.String("US"),
			Phone:            pulumi.String("+1.4155551234"),
			Email:            pulumi.String("john@example.com"),
		})
		if err != nil {
			return err
		}
		// Register a domain
		_, err = dnsimple.NewRegisteredDomain(ctx, "example_com", &dnsimple.RegisteredDomainArgs{
			Name:                pulumi.String("example.com"),
			ContactId:           registrant.ID(),
			AutoRenewEnabled:    pulumi.Bool(true),
			WhoisPrivacyEnabled: pulumi.Bool(true),
			TransferLockEnabled: pulumi.Bool(true),
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
  # Create a contact for domain registration
  registrant:
    type: dnsimple:Contact
    properties:
      label: Main Contact
      firstName: John
      lastName: Doe
      organizationName: Example Inc
      address1: 123 Main Street
      city: San Francisco
      stateProvince: California
      postalCode: '94105'
      country: US
      phone: '+1.4155551234'
      email: john@example.com
  # Register a domain
  exampleCom:
    type: dnsimple:RegisteredDomain
    name: example_com
    properties:
      name: example.com
      contactId: ${registrant.id}
      autoRenewEnabled: true
      whoisPrivacyEnabled: true
      transferLockEnabled: true
```

{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.dnsimple.Contact;
import com.pulumi.dnsimple.ContactArgs;
import com.pulumi.dnsimple.RegisteredDomain;
import com.pulumi.dnsimple.RegisteredDomainArgs;
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
        // Create a contact for domain registration
        var registrant = new Contact("registrant", ContactArgs.builder()
            .label("Main Contact")
            .firstName("John")
            .lastName("Doe")
            .organizationName("Example Inc")
            .address1("123 Main Street")
            .city("San Francisco")
            .stateProvince("California")
            .postalCode("94105")
            .country("US")
            .phone("+1.4155551234")
            .email("john@example.com")
            .build());

        // Register a domain
        var exampleCom = new RegisteredDomain("exampleCom", RegisteredDomainArgs.builder()
            .name("example.com")
            .contactId(registrant.id())
            .autoRenewEnabled(true)
            .whoisPrivacyEnabled(true)
            .transferLockEnabled(true)
            .build());

    }
}
```

{{% /choosable %}}
{{< /chooser >}}

For more elaborate use cases, and to learn more about the capabilities offered by the DNSimple Pulumi provider, view the individual resource and function pages.
## Configuration Reference

The following configuration inputs are supported in the provider configuration:

- **`token`** (Required) - The DNSimple [API v2 token](https://support.dnsimple.com/articles/api-access-token/). Can be provided via the `DNSIMPLE_TOKEN` environment variable. You can use either a User or Account token, but an Account token is recommended for better security and access control.

- **`account`** (Required) - The ID of the account associated with the token. Can be provided via the `DNSIMPLE_ACCOUNT` environment variable.

- **`sandbox`** (Optional) - Set to `true` to connect to the API [sandbox environment](https://developer.dnsimple.com/sandbox/) for testing. Can be provided via the `DNSIMPLE_SANDBOX` environment variable. Defaults to `false`.

- **`prefetch`** (Optional) - Set to `true` to enable prefetching zone records when dealing with large configurations. This is useful when you are dealing with API rate limitations given your number of zones and zone records. Can be provided via the `DNSIMPLE_PREFETCH` environment variable. Defaults to `false`.

- **`userAgent`** (Optional) - Custom string to append to the user agent used for sending HTTP requests to the API. Useful for identifying your automation or integration.