---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-venafi/v1.13.2/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi/pulumi-venafi/blob/v1.13.2/docs/_index.md
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Venafi Provider
meta_desc: Provides an overview on how to configure the Pulumi Venafi provider.
layout: package
---

## Installation

The Venafi provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/venafi`](https://www.npmjs.com/package/@pulumi/venafi)
* Python: [`pulumi-venafi`](https://pypi.org/project/pulumi-venafi/)
* Go: [`github.com/pulumi/pulumi-venafi/sdk/go/venafi`](https://github.com/pulumi/pulumi-venafi)
* .NET: [`Pulumi.Venafi`](https://www.nuget.org/packages/Pulumi.Venafi)
* Java: [`com.pulumi/venafi`](https://central.sonatype.com/artifact/com.pulumi/venafi)

## Overview

> We dropped support for RSA PKCS#1 formatted keys for TLS certificates in version 15.0 and also for EC Keys in version
0.15.4 (you can find out more about this transition in [here](https://github.com/Venafi/vcert/releases/tag/v4.17.0)).
For backward compatibility during Pulumi state refresh please update to version 0.15.5 or above.

> As a part for upgrading our provider to SDK version 2, we dropped support for Pulumi version 0.11 and below.

> With the introduction of version 0.18.0 the CyberArk
Certificate Manager Provider now incorporates a new feature related to certificate retirement. When an infrastructure is
decommissioned, the associated certificate will be automatically retired from the CyberArk Certificate Manager.

[CyberArk](https://www.cyberark.com) is the enterprise platform for Machine Identity Protection. The CyberArk Certificate
Manager Provider streamlines the process of acquiring SSL/TLS keys and certificates from CyberArk Machine Identity Service
giving assurance of compliance with Information Security policies. It provides resources that allow private keys and certificates
to be created as part of a Pulumi deployment.

Use the navigation to the left to read about the available resources.
## Example Usage for  CyberArk Certificate Manager, SaaS
You can sign up for a CyberArk Certificate Manager, SaaS account by visiting <https://www.cyberark.com/products/certificate-management/>.
Once registered, find your API key by clicking your name in the top right of the web interface.  You will also need to specify the `zone`
to use when requesting certificates. Zones define the machine identity policy that will be applied to certificate requests and the
certificate authority that will issue certificates. The zone is formed by combining the Application Name and Issuing
Template API Alias (e.g. "Business App\Enterprise CIT").
### US tenants

{{< chooser language "typescript,python,go,csharp,java,yaml,hcl" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    venafi:apiKey:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:zone:
        value: Business App\Enterprise CIT

```

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as venafi from "@pulumi/venafi";

// Generate a key pair and request a certificate
const webserver = new venafi.Certificate("webserver", {});
```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    venafi:apiKey:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:zone:
        value: Business App\Enterprise CIT

```

```python
import pulumi
import pulumi_venafi as venafi

# Generate a key pair and request a certificate
webserver = venafi.Certificate("webserver")
```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    venafi:apiKey:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:zone:
        value: Business App\Enterprise CIT

```

```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Venafi = Pulumi.Venafi;

return await Deployment.RunAsync(() =>
{
    // Generate a key pair and request a certificate
    var webserver = new Venafi.Certificate("webserver");

});

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    venafi:apiKey:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:zone:
        value: Business App\Enterprise CIT

```

```go
package main

import (
	"github.com/pulumi/pulumi-venafi/sdk/go/venafi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Generate a key pair and request a certificate
		_, err := venafi.NewCertificate(ctx, "webserver", nil)
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
config:
    venafi:apiKey:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:zone:
        value: Business App\Enterprise CIT

```

```yaml
resources:
  # Generate a key pair and request a certificate
  webserver:
    type: venafi:Certificate
```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    venafi:apiKey:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:zone:
        value: Business App\Enterprise CIT

```

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.venafi.Certificate;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        // Generate a key pair and request a certificate
        var webserver = new Certificate("webserver");

    }
}
```

{{% /choosable %}}
{{% choosable language hcl %}}
```hcl
pulumi {
  required_providers {
    venafi = {
      source = "pulumi/venafi"
    }
  }
}

# Generate a key pair and request a certificate
resource "venafi_certificate" "webserver" {
}
```

{{% /choosable %}}
{{< /chooser >}}
### EU tenants

{{< chooser language "typescript,python,go,csharp,java,yaml,hcl" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    venafi:apiKey:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:url:
        value: https://api.venafi.eu
    venafi:zone:
        value: Business App\Enterprise CIT

```

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as venafi from "@pulumi/venafi";

// Generate a key pair and request a certificate
const webserver = new venafi.Certificate("webserver", {});
```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    venafi:apiKey:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:url:
        value: https://api.venafi.eu
    venafi:zone:
        value: Business App\Enterprise CIT

```

```python
import pulumi
import pulumi_venafi as venafi

# Generate a key pair and request a certificate
webserver = venafi.Certificate("webserver")
```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    venafi:apiKey:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:url:
        value: https://api.venafi.eu
    venafi:zone:
        value: Business App\Enterprise CIT

```

```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Venafi = Pulumi.Venafi;

return await Deployment.RunAsync(() =>
{
    // Generate a key pair and request a certificate
    var webserver = new Venafi.Certificate("webserver");

});

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    venafi:apiKey:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:url:
        value: https://api.venafi.eu
    venafi:zone:
        value: Business App\Enterprise CIT

```

```go
package main

import (
	"github.com/pulumi/pulumi-venafi/sdk/go/venafi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Generate a key pair and request a certificate
		_, err := venafi.NewCertificate(ctx, "webserver", nil)
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
config:
    venafi:apiKey:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:url:
        value: https://api.venafi.eu
    venafi:zone:
        value: Business App\Enterprise CIT

```

```yaml
resources:
  # Generate a key pair and request a certificate
  webserver:
    type: venafi:Certificate
```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    venafi:apiKey:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:url:
        value: https://api.venafi.eu
    venafi:zone:
        value: Business App\Enterprise CIT

```

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.venafi.Certificate;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        // Generate a key pair and request a certificate
        var webserver = new Certificate("webserver");

    }
}
```

{{% /choosable %}}
{{% choosable language hcl %}}
```hcl
pulumi {
  required_providers {
    venafi = {
      source = "pulumi/venafi"
    }
  }
}

# Generate a key pair and request a certificate
resource "venafi_certificate" "webserver" {
}
```

{{% /choosable %}}
{{< /chooser >}}
### AU tenants

{{< chooser language "typescript,python,go,csharp,java,yaml,hcl" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    venafi:apiKey:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:url:
        value: https://api.au.venafi.cloud
    venafi:zone:
        value: Business App\Enterprise CIT

```

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as venafi from "@pulumi/venafi";

// Generate a key pair and request a certificate
const webserver = new venafi.Certificate("webserver", {});
```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    venafi:apiKey:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:url:
        value: https://api.au.venafi.cloud
    venafi:zone:
        value: Business App\Enterprise CIT

```

```python
import pulumi
import pulumi_venafi as venafi

# Generate a key pair and request a certificate
webserver = venafi.Certificate("webserver")
```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    venafi:apiKey:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:url:
        value: https://api.au.venafi.cloud
    venafi:zone:
        value: Business App\Enterprise CIT

```

```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Venafi = Pulumi.Venafi;

return await Deployment.RunAsync(() =>
{
    // Generate a key pair and request a certificate
    var webserver = new Venafi.Certificate("webserver");

});

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    venafi:apiKey:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:url:
        value: https://api.au.venafi.cloud
    venafi:zone:
        value: Business App\Enterprise CIT

```

```go
package main

import (
	"github.com/pulumi/pulumi-venafi/sdk/go/venafi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Generate a key pair and request a certificate
		_, err := venafi.NewCertificate(ctx, "webserver", nil)
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
config:
    venafi:apiKey:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:url:
        value: https://api.au.venafi.cloud
    venafi:zone:
        value: Business App\Enterprise CIT

```

```yaml
resources:
  # Generate a key pair and request a certificate
  webserver:
    type: venafi:Certificate
```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    venafi:apiKey:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:url:
        value: https://api.au.venafi.cloud
    venafi:zone:
        value: Business App\Enterprise CIT

```

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.venafi.Certificate;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        // Generate a key pair and request a certificate
        var webserver = new Certificate("webserver");

    }
}
```

{{% /choosable %}}
{{% choosable language hcl %}}
```hcl
pulumi {
  required_providers {
    venafi = {
      source = "pulumi/venafi"
    }
  }
}

# Generate a key pair and request a certificate
resource "venafi_certificate" "webserver" {
}
```

{{% /choosable %}}
{{< /chooser >}}
### UK tenants

{{< chooser language "typescript,python,go,csharp,java,yaml,hcl" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    venafi:apiKey:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:url:
        value: https://api.uk.venafi.cloud
    venafi:zone:
        value: Business App\Enterprise CIT

```

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as venafi from "@pulumi/venafi";

// Generate a key pair and request a certificate
const webserver = new venafi.Certificate("webserver", {});
```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    venafi:apiKey:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:url:
        value: https://api.uk.venafi.cloud
    venafi:zone:
        value: Business App\Enterprise CIT

```

```python
import pulumi
import pulumi_venafi as venafi

# Generate a key pair and request a certificate
webserver = venafi.Certificate("webserver")
```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    venafi:apiKey:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:url:
        value: https://api.uk.venafi.cloud
    venafi:zone:
        value: Business App\Enterprise CIT

```

```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Venafi = Pulumi.Venafi;

return await Deployment.RunAsync(() =>
{
    // Generate a key pair and request a certificate
    var webserver = new Venafi.Certificate("webserver");

});

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    venafi:apiKey:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:url:
        value: https://api.uk.venafi.cloud
    venafi:zone:
        value: Business App\Enterprise CIT

```

```go
package main

import (
	"github.com/pulumi/pulumi-venafi/sdk/go/venafi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Generate a key pair and request a certificate
		_, err := venafi.NewCertificate(ctx, "webserver", nil)
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
config:
    venafi:apiKey:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:url:
        value: https://api.uk.venafi.cloud
    venafi:zone:
        value: Business App\Enterprise CIT

```

```yaml
resources:
  # Generate a key pair and request a certificate
  webserver:
    type: venafi:Certificate
```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    venafi:apiKey:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:url:
        value: https://api.uk.venafi.cloud
    venafi:zone:
        value: Business App\Enterprise CIT

```

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.venafi.Certificate;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        // Generate a key pair and request a certificate
        var webserver = new Certificate("webserver");

    }
}
```

{{% /choosable %}}
{{% choosable language hcl %}}
```hcl
pulumi {
  required_providers {
    venafi = {
      source = "pulumi/venafi"
    }
  }
}

# Generate a key pair and request a certificate
resource "venafi_certificate" "webserver" {
}
```

{{% /choosable %}}
{{< /chooser >}}
### SG tenants

{{< chooser language "typescript,python,go,csharp,java,yaml,hcl" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    venafi:apiKey:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:url:
        value: https://api.sg.venafi.cloud
    venafi:zone:
        value: Business App\Enterprise CIT

```

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as venafi from "@pulumi/venafi";

// Generate a key pair and request a certificate
const webserver = new venafi.Certificate("webserver", {});
```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    venafi:apiKey:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:url:
        value: https://api.sg.venafi.cloud
    venafi:zone:
        value: Business App\Enterprise CIT

```

```python
import pulumi
import pulumi_venafi as venafi

# Generate a key pair and request a certificate
webserver = venafi.Certificate("webserver")
```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    venafi:apiKey:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:url:
        value: https://api.sg.venafi.cloud
    venafi:zone:
        value: Business App\Enterprise CIT

```

```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Venafi = Pulumi.Venafi;

return await Deployment.RunAsync(() =>
{
    // Generate a key pair and request a certificate
    var webserver = new Venafi.Certificate("webserver");

});

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    venafi:apiKey:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:url:
        value: https://api.sg.venafi.cloud
    venafi:zone:
        value: Business App\Enterprise CIT

```

```go
package main

import (
	"github.com/pulumi/pulumi-venafi/sdk/go/venafi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Generate a key pair and request a certificate
		_, err := venafi.NewCertificate(ctx, "webserver", nil)
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
config:
    venafi:apiKey:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:url:
        value: https://api.sg.venafi.cloud
    venafi:zone:
        value: Business App\Enterprise CIT

```

```yaml
resources:
  # Generate a key pair and request a certificate
  webserver:
    type: venafi:Certificate
```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    venafi:apiKey:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:url:
        value: https://api.sg.venafi.cloud
    venafi:zone:
        value: Business App\Enterprise CIT

```

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.venafi.Certificate;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        // Generate a key pair and request a certificate
        var webserver = new Certificate("webserver");

    }
}
```

{{% /choosable %}}
{{% choosable language hcl %}}
```hcl
pulumi {
  required_providers {
    venafi = {
      source = "pulumi/venafi"
    }
  }
}

# Generate a key pair and request a certificate
resource "venafi_certificate" "webserver" {
}
```

{{% /choosable %}}
{{< /chooser >}}
### CA tenants

{{< chooser language "typescript,python,go,csharp,java,yaml,hcl" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    venafi:apiKey:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:url:
        value: https://api.ca.venafi.cloud
    venafi:zone:
        value: Business App\Enterprise CIT

```

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as venafi from "@pulumi/venafi";

// Generate a key pair and request a certificate
const webserver = new venafi.Certificate("webserver", {});
```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    venafi:apiKey:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:url:
        value: https://api.ca.venafi.cloud
    venafi:zone:
        value: Business App\Enterprise CIT

```

```python
import pulumi
import pulumi_venafi as venafi

# Generate a key pair and request a certificate
webserver = venafi.Certificate("webserver")
```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    venafi:apiKey:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:url:
        value: https://api.ca.venafi.cloud
    venafi:zone:
        value: Business App\Enterprise CIT

```

```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Venafi = Pulumi.Venafi;

return await Deployment.RunAsync(() =>
{
    // Generate a key pair and request a certificate
    var webserver = new Venafi.Certificate("webserver");

});

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    venafi:apiKey:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:url:
        value: https://api.ca.venafi.cloud
    venafi:zone:
        value: Business App\Enterprise CIT

```

```go
package main

import (
	"github.com/pulumi/pulumi-venafi/sdk/go/venafi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Generate a key pair and request a certificate
		_, err := venafi.NewCertificate(ctx, "webserver", nil)
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
config:
    venafi:apiKey:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:url:
        value: https://api.ca.venafi.cloud
    venafi:zone:
        value: Business App\Enterprise CIT

```

```yaml
resources:
  # Generate a key pair and request a certificate
  webserver:
    type: venafi:Certificate
```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    venafi:apiKey:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:url:
        value: https://api.ca.venafi.cloud
    venafi:zone:
        value: Business App\Enterprise CIT

```

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.venafi.Certificate;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        // Generate a key pair and request a certificate
        var webserver = new Certificate("webserver");

    }
}
```

{{% /choosable %}}
{{% choosable language hcl %}}
```hcl
pulumi {
  required_providers {
    venafi = {
      source = "pulumi/venafi"
    }
  }
}

# Generate a key pair and request a certificate
resource "venafi_certificate" "webserver" {
}
```

{{% /choosable %}}
{{< /chooser >}}
## Example Usage for CyberArk Certificate Manager, Self-Hosted

Your CyberArk administrator can provide you with the URL for the CyberArk Certificate Manager, Self-Hosted REST API and grant you
permission to use it.  At the same time they'll provide you with the Distinguished Name of a policy folder to specify
for the `zone`. Policy folders define the machine identity policy applied  to certificate requests and the certificate
authority that will issue certificates. You may also need to ask them for a root CA certificate for your `trustBundle`
if the CyberArk Certificate Manager, Self-Hosted URL is secured by a certificate your Pulumi computer does not already trust.

Obtain the required `accessToken` for CyberArk Certificate Manager, Self-Hosted using the [VCert CLI](https://github.com/Venafi/vcert/blob/master/README-CLI-PLATFORM.md#obtaining-an-authorization-token)
(`getcred action` with `--client-id "pulumi-pulumi-by-venafi"` and `--scope "certificate:manage"`) or the
CyberArk Certificate Manager, Self-Hosted Platform's Authorize REST API method. The *configuration:manage* scope is required to set certificate policy using the
`venafi.Policy` resource.

{{< chooser language "typescript,python,go,csharp,java,yaml,hcl" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    venafi:accessToken:
        value: p0WTt3sDPbzm2BDIkoJROQ==
    venafi:trustBundle:
        value: 'TODO: "${file("/opt/venafi/bundle.pem")}"'
    venafi:url:
        value: https://tpp.venafi.example
    venafi:zone:
        value: DevOps\Pulumi

```

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as venafi from "@pulumi/venafi";

// Generate a key pair and request a certificate
const webserver = new venafi.Certificate("webserver", {});
```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    venafi:accessToken:
        value: p0WTt3sDPbzm2BDIkoJROQ==
    venafi:trustBundle:
        value: 'TODO: "${file("/opt/venafi/bundle.pem")}"'
    venafi:url:
        value: https://tpp.venafi.example
    venafi:zone:
        value: DevOps\Pulumi

```

```python
import pulumi
import pulumi_venafi as venafi

# Generate a key pair and request a certificate
webserver = venafi.Certificate("webserver")
```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    venafi:accessToken:
        value: p0WTt3sDPbzm2BDIkoJROQ==
    venafi:trustBundle:
        value: 'TODO: "${file("/opt/venafi/bundle.pem")}"'
    venafi:url:
        value: https://tpp.venafi.example
    venafi:zone:
        value: DevOps\Pulumi

```

```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Venafi = Pulumi.Venafi;

return await Deployment.RunAsync(() =>
{
    // Generate a key pair and request a certificate
    var webserver = new Venafi.Certificate("webserver");

});

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    venafi:accessToken:
        value: p0WTt3sDPbzm2BDIkoJROQ==
    venafi:trustBundle:
        value: 'TODO: "${file("/opt/venafi/bundle.pem")}"'
    venafi:url:
        value: https://tpp.venafi.example
    venafi:zone:
        value: DevOps\Pulumi

```

```go
package main

import (
	"github.com/pulumi/pulumi-venafi/sdk/go/venafi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Generate a key pair and request a certificate
		_, err := venafi.NewCertificate(ctx, "webserver", nil)
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
config:
    venafi:accessToken:
        value: p0WTt3sDPbzm2BDIkoJROQ==
    venafi:trustBundle:
        value: 'TODO: "${file("/opt/venafi/bundle.pem")}"'
    venafi:url:
        value: https://tpp.venafi.example
    venafi:zone:
        value: DevOps\Pulumi

```

```yaml
resources:
  # Generate a key pair and request a certificate
  webserver:
    type: venafi:Certificate
```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    venafi:accessToken:
        value: p0WTt3sDPbzm2BDIkoJROQ==
    venafi:trustBundle:
        value: 'TODO: "${file("/opt/venafi/bundle.pem")}"'
    venafi:url:
        value: https://tpp.venafi.example
    venafi:zone:
        value: DevOps\Pulumi

```

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.venafi.Certificate;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        // Generate a key pair and request a certificate
        var webserver = new Certificate("webserver");

    }
}
```

{{% /choosable %}}
{{% choosable language hcl %}}
```hcl
pulumi {
  required_providers {
    venafi = {
      source = "pulumi/venafi"
    }
  }
}

# Generate a key pair and request a certificate
resource "venafi_certificate" "webserver" {
}
```

{{% /choosable %}}
{{< /chooser >}}
## Example Usage for Palo Alto Networks Next-Gen Trust Security (NGTS)

{{< chooser language "typescript,python,go,csharp,java,yaml,hcl" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    venafi:clientId:
        value: your-service-account@panserviceaccount.com
    venafi:clientSecret:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:tokenUrl:
        value: https://auth.apps.paloaltonetworks.com/auth/v1/oauth2/access_token
    venafi:tsgId:
        value: "1234567890"
    venafi:url:
        value: https://api.strata.paloaltonetworks.com/ngts
    venafi:zone:
        value: Built-In CA Open Policy

```

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as venafi from "@pulumi/venafi";

// Generate a key pair and request a certificate
const webserver = new venafi.Certificate("webserver", {});
```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    venafi:clientId:
        value: your-service-account@panserviceaccount.com
    venafi:clientSecret:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:tokenUrl:
        value: https://auth.apps.paloaltonetworks.com/auth/v1/oauth2/access_token
    venafi:tsgId:
        value: "1234567890"
    venafi:url:
        value: https://api.strata.paloaltonetworks.com/ngts
    venafi:zone:
        value: Built-In CA Open Policy

```

```python
import pulumi
import pulumi_venafi as venafi

# Generate a key pair and request a certificate
webserver = venafi.Certificate("webserver")
```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    venafi:clientId:
        value: your-service-account@panserviceaccount.com
    venafi:clientSecret:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:tokenUrl:
        value: https://auth.apps.paloaltonetworks.com/auth/v1/oauth2/access_token
    venafi:tsgId:
        value: "1234567890"
    venafi:url:
        value: https://api.strata.paloaltonetworks.com/ngts
    venafi:zone:
        value: Built-In CA Open Policy

```

```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Venafi = Pulumi.Venafi;

return await Deployment.RunAsync(() =>
{
    // Generate a key pair and request a certificate
    var webserver = new Venafi.Certificate("webserver");

});

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    venafi:clientId:
        value: your-service-account@panserviceaccount.com
    venafi:clientSecret:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:tokenUrl:
        value: https://auth.apps.paloaltonetworks.com/auth/v1/oauth2/access_token
    venafi:tsgId:
        value: "1234567890"
    venafi:url:
        value: https://api.strata.paloaltonetworks.com/ngts
    venafi:zone:
        value: Built-In CA Open Policy

```

```go
package main

import (
	"github.com/pulumi/pulumi-venafi/sdk/go/venafi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Generate a key pair and request a certificate
		_, err := venafi.NewCertificate(ctx, "webserver", nil)
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
config:
    venafi:clientId:
        value: your-service-account@panserviceaccount.com
    venafi:clientSecret:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:tokenUrl:
        value: https://auth.apps.paloaltonetworks.com/auth/v1/oauth2/access_token
    venafi:tsgId:
        value: "1234567890"
    venafi:url:
        value: https://api.strata.paloaltonetworks.com/ngts
    venafi:zone:
        value: Built-In CA Open Policy

```

```yaml
resources:
  # Generate a key pair and request a certificate
  webserver:
    type: venafi:Certificate
```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    venafi:clientId:
        value: your-service-account@panserviceaccount.com
    venafi:clientSecret:
        value: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    venafi:tokenUrl:
        value: https://auth.apps.paloaltonetworks.com/auth/v1/oauth2/access_token
    venafi:tsgId:
        value: "1234567890"
    venafi:url:
        value: https://api.strata.paloaltonetworks.com/ngts
    venafi:zone:
        value: Built-In CA Open Policy

```

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.venafi.Certificate;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        // Generate a key pair and request a certificate
        var webserver = new Certificate("webserver");

    }
}
```

{{% /choosable %}}
{{% choosable language hcl %}}
```hcl
pulumi {
  required_providers {
    venafi = {
      source = "pulumi/venafi"
    }
  }
}

# Generate a key pair and request a certificate
resource "venafi_certificate" "webserver" {
}
```

{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

The following configuration inputs are supported:

* `accessToken` - (Optional, string) Authentication token for the 'pulumi-pulumi-by-venafi' API Application.
  Applies only to CyberArk Certificate Manager, Self-Hosted.
* `apiKey` - (Optional, string) REST API key for authentication. Applies only to CyberArk Certificate Manager, SaaS.
* `clientId` - (Optional, string) Application that will be using the token. Defaults to `pulumi-pulumi-by-venafi`.
* `clientSecret` - (Optional, string) Client Secret for CyberArk Certificate Manager, Self-Hosted or NGTS authentication.
* `devMode` - (Optional, boolean) When "true" will test the provider without connecting to CyberArk Certificate Manager, Self-Hosted
  or CyberArk Certificate Manager, SaaS
* `externalJwt` - (Optional, string) JWT of the Identity Provider associated to a service account for authentication.
  Applies only to CyberArk Certificate Manager, SaaS.
* `p12CertFilename` - (Optional, string) Filename of PKCS#12 keystore containing a client certificate, private key,
* `p12CertData` - (Optional, string) Base64 encoded PKCS#12 keystore containing a client certificate, private key,
  and chain certificates to authenticate to CyberArk Certificate Manager, Self-Hosted.
* `p12CertPassword` - (Optional, string) Password for the PKCS#12 keystore declared in `p12CertFilename` or in `p12CertData`. Applies
  only to CyberArk Certificate Manager, Self-Hosted.
* `skipRetirement` - (Optional, boolean) If it's specified with value `true` then the certificate retirement on the
  related CyberArk Certificate Manager, Self-Hosted or CyberArk Certificate Manager, SaaS will be skipped. A value of `false` is equivalent to omit this argument.
* `tokenUrl` - (Optional, string) - URL to request access tokens for CyberArk Certificate Manager, SaaS or NGTS.
* `tppPassword` **[DEPRECATED]** - (Optional, string) WebSDK account password for authentication (applies only to
  CyberArk Certificate Manager, Self-Hosted).
* `tppUsername` **[DEPRECATED]** - (Optional, string) WebSDK account username for authentication (applies only to
  CyberArk Certificate Manager, Self-Hosted).
* `trustBundle` - (Optional, string) PEM trust bundle for CyberArk Certificate Manager, Self-Hosted server certificate (e.g. "${file("bundle.pem")}").
* `tsgId` - (Optional, string) The Palo Alto Networks NGTS TSG ID used when issuing a token. Only used when platform is NGTS.
* `url` - (Optional, string) CyberArk Certificate Manager, Self-Hosted URL (e.g. "https://tpp.cyberark.example") or NGTS URL (e.g. "https://api.strata.paloaltonetworks.com/ngts").
* `zone` - (**Required**, string) Application Name and Issuing Template API Alias (e.g. "Business App\Enterprise CIT")
  for CyberArk Certificate Manager, SaaS, policy folder for CyberArk Certificate Manager, Self-Hosted, or Certificate Issuing Template name for NGTS.
## Environment Variables

The following environment variables can also be used to specify provider
argument values:

* `VENAFI_API` - for `apiKey` argument
* `VENAFI_CLIENT_ID` - for `clientId` argument
* `VENAFI_CLIENT_SECRET` - for `clientSecret` argument
* `VENAFI_DEVMODE` - for `devMode` argument
* `VENAFI_EXTERNAL_JWT` - for `externalJwt` argument
* `VENAFI_PASS` - for `tppPassword` argument
* `VENAFI_P12_CERTIFICATE` - for `p12CertFilename` argument
* `VENAFI_P12_PASSWORD` - for `p12Password` argument
* `VENAFI_SKIP_RETIREMENT` - for `skipRetirement` argument
* `VENAFI_TOKEN` - for `accessToken` argument
* `VENAFI_TOKEN_URL` - for `tokenUrl` argument
* `VENAFI_TSG_ID` - for `tsgId` argument
* `VENAFI_URL` - for `url` argument
* `VENAFI_USER` - for `tppUsername` argument
* `VENAFI_ZONE` - for `zone` argument