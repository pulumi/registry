---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/checkpointsw/checkpoint/2.10.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Checkpoint Provider
meta_desc: Provides an overview on how to configure the Pulumi Checkpoint provider.
layout: package
---

## Generate Provider

The Checkpoint provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider checkpointsw/checkpoint
```
## Overview

The Check Point provider can be used to automate security responses to threats, provision both physical and virtualized
next-generation firewalls and automate routine Security Management configuration tasks, saving time and reducing
configuration errors. With the Check Point provider, DevOps teams can automate their security and transform it into
DevSecOps workflows.
## Examples of usage
To use Check Point provider, copy and paste this code into your Pulumi configuration, update provider configuration and run `pulumi up`.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    checkpoint:context:
        value: web_api
    checkpoint:domain:
        value: MyDomain
    checkpoint:password:
        value: aaaa
    checkpoint:server:
        value: 192.0.2.1
    checkpoint:sessionFileName:
        value: mydomain.json
    checkpoint:sessionName:
        value: Pulumi session
    checkpoint:username:
        value: aa

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as checkpoint from "@pulumi/checkpoint";

// Create network object
const network = new checkpoint.ManagementNetwork("network", {
    name: "My network",
    subnet4: "192.0.2.0",
    maskLength4: 24,
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    checkpoint:context:
        value: web_api
    checkpoint:domain:
        value: MyDomain
    checkpoint:password:
        value: aaaa
    checkpoint:server:
        value: 192.0.2.1
    checkpoint:sessionFileName:
        value: mydomain.json
    checkpoint:sessionName:
        value: Pulumi session
    checkpoint:username:
        value: aa

```
```python
import pulumi
import pulumi_checkpoint as checkpoint

# Create network object
network = checkpoint.ManagementNetwork("network",
    name="My network",
    subnet4="192.0.2.0",
    mask_length4=24)
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    checkpoint:context:
        value: web_api
    checkpoint:domain:
        value: MyDomain
    checkpoint:password:
        value: aaaa
    checkpoint:server:
        value: 192.0.2.1
    checkpoint:sessionFileName:
        value: mydomain.json
    checkpoint:sessionName:
        value: Pulumi session
    checkpoint:username:
        value: aa

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Checkpoint = Pulumi.Checkpoint;

return await Deployment.RunAsync(() =>
{
    // Create network object
    var network = new Checkpoint.ManagementNetwork("network", new()
    {
        Name = "My network",
        Subnet4 = "192.0.2.0",
        MaskLength4 = 24,
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    checkpoint:context:
        value: web_api
    checkpoint:domain:
        value: MyDomain
    checkpoint:password:
        value: aaaa
    checkpoint:server:
        value: 192.0.2.1
    checkpoint:sessionFileName:
        value: mydomain.json
    checkpoint:sessionName:
        value: Pulumi session
    checkpoint:username:
        value: aa

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/checkpoint/v2/checkpoint"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create network object
		_, err := checkpoint.NewManagementNetwork(ctx, "network", &checkpoint.ManagementNetworkArgs{
			Name:        pulumi.String("My network"),
			Subnet4:     pulumi.String("192.0.2.0"),
			MaskLength4: pulumi.Float64(24),
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
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    checkpoint:context:
        value: web_api
    checkpoint:domain:
        value: MyDomain
    checkpoint:password:
        value: aaaa
    checkpoint:server:
        value: 192.0.2.1
    checkpoint:sessionFileName:
        value: mydomain.json
    checkpoint:sessionName:
        value: Pulumi session
    checkpoint:username:
        value: aa

```
```yaml
resources:
  # Create network object
  network:
    type: checkpoint:ManagementNetwork
    properties:
      name: My network
      subnet4: 192.0.2.0
      maskLength4: '24'
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    checkpoint:context:
        value: web_api
    checkpoint:domain:
        value: MyDomain
    checkpoint:password:
        value: aaaa
    checkpoint:server:
        value: 192.0.2.1
    checkpoint:sessionFileName:
        value: mydomain.json
    checkpoint:sessionName:
        value: Pulumi session
    checkpoint:username:
        value: aa

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.checkpoint.ManagementNetwork;
import com.pulumi.checkpoint.ManagementNetworkArgs;
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
        // Create network object
        var network = new ManagementNetwork("network", ManagementNetworkArgs.builder()
            .name("My network")
            .subnet4("192.0.2.0")
            .maskLength4("24")
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
config:
    checkpoint:context:
        value: gaia_api
    checkpoint:password:
        value: gaia_password
    checkpoint:server:
        value: 192.0.2.1
    checkpoint:username:
        value: gaia_user

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as checkpoint from "@pulumi/checkpoint";

// Set machine hostname
const hostname = new checkpoint.Hostname("hostname", {name: "pulumi_host"});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    checkpoint:context:
        value: gaia_api
    checkpoint:password:
        value: gaia_password
    checkpoint:server:
        value: 192.0.2.1
    checkpoint:username:
        value: gaia_user

```
```python
import pulumi
import pulumi_checkpoint as checkpoint

# Set machine hostname
hostname = checkpoint.Hostname("hostname", name="pulumi_host")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    checkpoint:context:
        value: gaia_api
    checkpoint:password:
        value: gaia_password
    checkpoint:server:
        value: 192.0.2.1
    checkpoint:username:
        value: gaia_user

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Checkpoint = Pulumi.Checkpoint;

return await Deployment.RunAsync(() =>
{
    // Set machine hostname
    var hostname = new Checkpoint.Hostname("hostname", new()
    {
        Name = "pulumi_host",
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    checkpoint:context:
        value: gaia_api
    checkpoint:password:
        value: gaia_password
    checkpoint:server:
        value: 192.0.2.1
    checkpoint:username:
        value: gaia_user

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/checkpoint/v2/checkpoint"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Set machine hostname
		_, err := checkpoint.NewHostname(ctx, "hostname", &checkpoint.HostnameArgs{
			Name: pulumi.String("pulumi_host"),
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
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    checkpoint:context:
        value: gaia_api
    checkpoint:password:
        value: gaia_password
    checkpoint:server:
        value: 192.0.2.1
    checkpoint:username:
        value: gaia_user

```
```yaml
resources:
  # Set machine hostname
  hostname:
    type: checkpoint:Hostname
    properties:
      name: pulumi_host
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    checkpoint:context:
        value: gaia_api
    checkpoint:password:
        value: gaia_password
    checkpoint:server:
        value: 192.0.2.1
    checkpoint:username:
        value: gaia_user

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.checkpoint.Hostname;
import com.pulumi.checkpoint.HostnameArgs;
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
        // Set machine hostname
        var hostname = new Hostname("hostname", HostnameArgs.builder()
            .name("pulumi_host")
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

The following configuration inputs are supported:

* `server` - (Optional) Check Point Management server IP. It must be provided, but can also be defined via
  the `CHECKPOINT_SERVER` environment variable.
* `username` - (Optional) Check Point Management admin name. It must be provided, but can also be defined via
  the `CHECKPOINT_USERNAME` environment variable.
* `password` - (Optional) Check Point Management admin password. It must be provided, but can also be defined via
  the `CHECKPOINT_PASSWORD` environment variable.
* `apiKey` - (Optional) Check Point Management admin API key. It must be provided, but can also be defined via
  the `CHECKPOINT_API_KEY` environment variable.
* `domain` - (Optional) Login to specific domain. Domain can be identified by name or UID. This can also be defined via
  the `CHECKPOINT_DOMAIN` environment variable.
* `context` - (Optional) Check Point access context - `webApi` or `gaiaApi`. This can also be defined via
  the `CHECKPOINT_CONTEXT` environment variable. Default value is `webApi`.
* `port` - (Optional) Port used for connection with the API server. This can also be defined via the `CHECKPOINT_PORT`
  environment variable. Default value is `443`.
* `proxyHost` - (Optional) Proxy host used for proxy connections. This can also be defined via
  the `CHECKPOINT_PROXY_HOST` environment variable.
* `proxyPort` - (Optional) Proxy port used for proxy connections. This can also be defined via
  the `CHECKPOINT_PROXY_PORT` environment variable.
* `sessionName` - (Optional) Session unique name. This can also be defined via
  the `CHECKPOINT_SESSION_NAME` environment variable.
* `sessionDescription` - (Optional) Session purpose description. This can also be defined via the `CHECKPOINT_SESSION_DESCRIPTION` environment variable.
* `sessionFileName` - (Optional) Session file name used to store the current session id. This can also be defined via
  the `CHECKPOINT_SESSION_FILE_NAME` environment variable. default value is `sid.json`.
* `sessionTimeout` - (Optional) Timeout in seconds for the session established in Check Point. This can also be defined via
  the `CHECKPOINT_SESSION_TIMEOUT` environment variable. The default for the value is `600`. The timeout can be `10` - `3600`.
* `timeout` - (Optional) Timeout in seconds for the Go SDK to complete a transaction. This can also be defined via
  the `CHECKPOINT_TIMEOUT` environment variable. Default value is `120` seconds.
* `cloudMgmtId` - (Optional) Smart-1 Cloud management UID. This can also be defined via
  the `CHECKPOINT_CLOUD_MGMT_ID` environment variable.
* `autoPublishBatchSize` - (Optional) Number of batch size to automatically run publish. This can also be defined via
  the `CHECKPOINT_AUTO_PUBLISH_BATCH_SIZE` environment variable.
* `ignoreServerCertificate` - (Optional) Indicates that the client should not check the server's certificate. This can also be defined via
  the `CHECKPOINT_IGNORE_SERVER_CERTIFICATE` environment variable.
## Authentication

Check Point Provider offers providing credentials for authentication. The following methods are supported:

- Static credentials
- Environment variables
### Static credentials

Usage with username and password:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    checkpoint:context:
        value: web_api
    checkpoint:domain:
        value: MyDomain
    checkpoint:password:
        value: aaaa
    checkpoint:server:
        value: 192.0.2.1
    checkpoint:sessionFileName:
        value: mydomain.json
    checkpoint:username:
        value: aa

```

Usage with API key:
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    checkpoint:apiKey:
        value: tBdloE9eOYzzSQicNxS7mA==
    checkpoint:context:
        value: web_api
    checkpoint:domain:
        value: MyDomain
    checkpoint:server:
        value: 192.0.2.1
    checkpoint:sessionFileName:
        value: mydomain.json

```

Usage for Smart-1 Cloud:
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    checkpoint:apiKey:
        value: tBdloE9eOYzzSQicNxS7mA==
    checkpoint:cloudMgmtId:
        value: de9a9b08-c7c7-436e-a64a-a54136301701
    checkpoint:context:
        value: web_api
    checkpoint:server:
        value: chkp-vmnc6s4y.maas.checkpoint.com

```

Or for GAIA API:
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    checkpoint:context:
        value: gaia_api
    checkpoint:password:
        value: gaia_password
    checkpoint:server:
        value: 192.0.2.1
    checkpoint:username:
        value: gaia_user

```
### Environment variables

You can provide your credentials via environment variables. Note that setting your Check Point credentials using static
credentials will override the environment variables.

Usage:

```bash
$ export CHECKPOINT_SERVER="192.0.2.1"
$ export CHECKPOINT_USERNAME="aa"
$ export CHECKPOINT_PASSWORD="aaaa"
$ export CHECKPOINT_CONTEXT="web_api"
$ export CHECKPOINT_DOMAIN="MyDomain"
$ export CHECKPOINT_TIMEOUT=10
$ export CHECKPOINT_PORT=443
$ export CHECKPOINT_SESSION_NAME="Pulumi session name"
$ export CHECKPOINT_SESSION_DESCRIPTION="Pulumi session description"
$ export CHECKPOINT_SESSION_FILE_NAME="mydomain.json"
$ export CHECKPOINT_SESSION_TIMEOUT=600
$ export CHECKPOINT_PROXY_HOST="1.2.3.4"
$ export CHECKPOINT_PROXY_PORT="123"
$ export CHECKPOINT_CLOUD_MGMT_ID="de9a9b08-c7c7-436e-a64a-a54136301701"
$ export CHECKPOINT_AUTO_PUBLISH_BATCH_SIZE=100
$ export CHECKPOINT_IGNORE_SERVER_CERTIFICATE=false
```

Usage with api key:

```bash
$ export CHECKPOINT_SERVER="192.0.2.1"
$ export CHECKPOINT_API_KEY="tBdloE9eOYzzSQicNxS7mA=="
$ export CHECKPOINT_CONTEXT="web_api"
$ export CHECKPOINT_DOMAIN="MyDomain"
$ export CHECKPOINT_TIMEOUT=10
$ export CHECKPOINT_PORT=443
$ export CHECKPOINT_SESSION_NAME="Pulumi session name"
$ export CHECKPOINT_SESSION_DESCRIPTION="Pulumi session description"
$ export CHECKPOINT_SESSION_FILE_NAME="mydomain.json"
$ export CHECKPOINT_SESSION_TIMEOUT=600
$ export CHECKPOINT_PROXY_HOST="1.2.3.4"
$ export CHECKPOINT_PROXY_PORT="123"
$ export CHECKPOINT_CLOUD_MGMT_ID="de9a9b08-c7c7-436e-a64a-a54136301701"
$ export CHECKPOINT_AUTO_PUBLISH_BATCH_SIZE=100
$ export CHECKPOINT_IGNORE_SERVER_CERTIFICATE=false
```

Then configure the Check Point Provider as following:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as checkpoint from "@pulumi/checkpoint";

// Create network object
const network = new checkpoint.ManagementNetwork("network", {
    name: "My network",
    subnet4: "192.0.2.0",
    maskLength4: 24,
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
import pulumi_checkpoint as checkpoint

# Create network object
network = checkpoint.ManagementNetwork("network",
    name="My network",
    subnet4="192.0.2.0",
    mask_length4=24)
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
using Checkpoint = Pulumi.Checkpoint;

return await Deployment.RunAsync(() =>
{
    // Create network object
    var network = new Checkpoint.ManagementNetwork("network", new()
    {
        Name = "My network",
        Subnet4 = "192.0.2.0",
        MaskLength4 = 24,
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
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/checkpoint/v2/checkpoint"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create network object
		_, err := checkpoint.NewManagementNetwork(ctx, "network", &checkpoint.ManagementNetworkArgs{
			Name:        pulumi.String("My network"),
			Subnet4:     pulumi.String("192.0.2.0"),
			MaskLength4: pulumi.Float64(24),
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
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml

```
```yaml
resources:
  # Create network object
  network:
    type: checkpoint:ManagementNetwork
    properties:
      name: My network
      subnet4: 192.0.2.0
      maskLength4: '24'
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
import com.pulumi.checkpoint.ManagementNetwork;
import com.pulumi.checkpoint.ManagementNetworkArgs;
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
        // Create network object
        var network = new ManagementNetwork("network", ManagementNetworkArgs.builder()
            .name("My network")
            .subnet4("192.0.2.0")
            .maskLength4("24")
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

Or for GAIA API:

```bash
$ export CHECKPOINT_SERVER="192.0.2.1"
$ export CHECKPOINT_USERNAME="gaia_user"
$ export CHECKPOINT_PASSWORD="gaia_password"
$ export CHECKPOINT_CONTEXT="gaia_api"
$ export CHECKPOINT_TIMEOUT=10
$ export CHECKPOINT_PORT=443
```

Then configure the Check Point Provider as following:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as checkpoint from "@pulumi/checkpoint";

// Set machine hostname
const hostname = new checkpoint.Hostname("hostname", {name: "pulumi_host"});
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
import pulumi_checkpoint as checkpoint

# Set machine hostname
hostname = checkpoint.Hostname("hostname", name="pulumi_host")
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
using Checkpoint = Pulumi.Checkpoint;

return await Deployment.RunAsync(() =>
{
    // Set machine hostname
    var hostname = new Checkpoint.Hostname("hostname", new()
    {
        Name = "pulumi_host",
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
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/checkpoint/v2/checkpoint"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Set machine hostname
		_, err := checkpoint.NewHostname(ctx, "hostname", &checkpoint.HostnameArgs{
			Name: pulumi.String("pulumi_host"),
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
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml

```
```yaml
resources:
  # Set machine hostname
  hostname:
    type: checkpoint:Hostname
    properties:
      name: pulumi_host
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
import com.pulumi.checkpoint.Hostname;
import com.pulumi.checkpoint.HostnameArgs;
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
        // Set machine hostname
        var hostname = new Hostname("hostname", HostnameArgs.builder()
            .name("pulumi_host")
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Post Apply / Destroy scripts

There are actions that can run out-of-band Pulumi using dedicated scripts for publish, install-policy and more.

In order to use post apply or post destroy commands, the authentication method must be via environment variables.
### Publish

Please use the following script for Publish:

```bash
$ cd $GOPATH/src/github.com/pulumi-providers/pulumi-provider-checkpoint/commands/publish
$ go build publish.go
$ mv publish $GOPATH/src/github.com/pulumi-providers/pulumi-provider-checkpoint
$ pulumi up && publish
```
### Install Policy

The following configuration inputs are supported:

* `policy-package` - (Required) The name of the Policy Package to be installed.
* `target` - (Optional) On what targets to execute this command. Targets may be identified by their name or object unique identifier. Multiple targets can be added.
* `access` - (Optional) Set to be true in order to install the Access Control policy. By default, the value is true if Access Control policy is enabled on the input policy package, otherwise false.
* `desktop-security` - (Optional) Set to be true in order to install the Desktop Security policy. By default, the value is true if desktop security policy is enabled on the input policy package, otherwise false.
* `qos` - (Optional) Set to be true in order to install the QoS policy. By default, the value is true if Quality-of-Service policy is enabled on the input policy package, otherwise false.
* `threat-prevention` - (Optional) Set to be true in order to install the Threat Prevention policy. By default, the value is true if Threat Prevention policy is enabled on the input policy package, otherwise false.
* `install-on-all-cluster-members-or-fail` - (Optional) Relevant for the gateway clusters. If true, the policy is installed on all the cluster members. If the installation on a cluster member fails, don't install on that cluster.
* `prepare-only` - (Optional) If true, prepares the policy for the installation, but doesn't install it on an installation target.
* `revision` - (Optional) The UID of the revision of the policy to install.
* `ignore-warnings` - (Optional) Install policy ignoring policy mismatch warnings.

Please use the following script for Install Policy:

```bash
$ cd $GOPATH/src/github.com/pulumi-providers/pulumi-provider-checkpoint/commands/install_policy
$ go build install_policy.go
$ mv install_policy $GOPATH/src/github.com/pulumi-providers/pulumi-provider-checkpoint
$ pulumi up && install_policy -policy-package <package name> -target <target name or uid>
```
### Logout

Please use the following script for Logout:

```bash
$ cd $GOPATH/src/github.com/pulumi-providers/pulumi-provider-checkpoint/commands/logout
$ go build logout.go
$ mv logout $GOPATH/src/github.com/pulumi-providers/pulumi-provider-checkpoint/logout_from_session
$ pulumi up && publish && logout_from_session
```
#### Example of usage

Run pulumi then Publish & Install Policy & Logout from session

```bash
$ pulumi up && publish && install_policy -policy-package "standard" -target "corporate-gateway" && logout_from_session
```
### Discard

Please use the following script for Discard:

```bash
$ cd $GOPATH/src/github.com/pulumi-providers/pulumi-provider-checkpoint/commands/discard
$ go build discard.go
$ mv discard $GOPATH/src/github.com/pulumi-providers/pulumi-provider-checkpoint
$ discard
```
### Approve Session

Please use the following script for Approve Session:

```bash
$ cd $GOPATH/src/github.com/pulumi-providers/pulumi-provider-checkpoint/commands/approve_session
$ go build approve_session.go
$ mv approve_session $GOPATH/src/github.com/pulumi-providers/pulumi-provider-checkpoint
$ approve_session "SESSION_UID"
```
### Reject Session

Please use the following script for Reject Session:

```bash
$ cd $GOPATH/src/github.com/pulumi-providers/pulumi-provider-checkpoint/commands/reject_session
$ go build reject_session.go
$ mv reject_session $GOPATH/src/github.com/pulumi-providers/pulumi-provider-checkpoint
$ reject_session "SESSION_UID" "REJECT_REASON"
```
### Submit Session

Please use the following script for Submit Session:

```bash
$ cd $GOPATH/src/github.com/pulumi-providers/pulumi-provider-checkpoint/commands/submit_session
$ go build submit_session.go
$ mv submit_session $GOPATH/src/github.com/pulumi-providers/pulumi-provider-checkpoint
$ submit_session "SESSION_UID"
```

if no `sessionUid` is provided it will submit the current session.
### Verify Policy

The following configuration inputs are supported:

* `policy-package` - (Required) Policy package identified by the name or UID to be verified.

Please use the following script for Verify Policy:

```bash
$ cd $GOPATH/src/github.com/pulumi-providers/pulumi-provider-checkpoint/commands/verify_policy
$ go build verify_policy.go
$ mv verify_policy $GOPATH/src/github.com/pulumi-providers/pulumi-provider-checkpoint
$ pulumi up && verify_policy -policy-package <package name>
```
## Compatibility with Management
Check Point Provider supports Management server from version R80 and above.
However, some Pulumi resources or specific fields in Pulumi resource might not be available because they are not supported in your Management API version.
<br>You can check the Management API [versions list](https://sc1.checkpoint.com/documents/latest/APIs/index.html#api_versions) to see what is supported by your Management server.
## Compatibility with CME
Check Point Provider supports configuring objects in CME configuration file starting from Security Management/Multi-Domain Security Management Server version R81.10 and higher.

The table below shows the compatibility between the Pulumi Release version and the CME API version:

| Pulumi Release version | CME API version |    CME Take    |
|---------------------------|-----------------|----------------|
| v2.9.0                    | v1.2.2          | 289 and higher |
| v2.8.0                    | v1.2            | 279 and higher |
| v2.7.0                    | v1.1            | 255 and higher |

<br>
> **Note:** When you install or upgrade the Pulumi Release version, make sure to also upgrade CME to the corresponding CME Take to properly configure CME resources.

For details about upgrading CME, please refer to the documentation [here](https://sc1.checkpoint.com/documents/IaaS/WebAdminGuides/EN/CP_CME/Content/Topics-CME/Installing_and_Updating_CME.htm?tocpath=_____4).
## Import Resources

In order to import resource, use the `pulumi import` command with object unique identifier.

Example:

For existing Host object with UID `9423d36f-2d66-4754-b9e2-e7f4493756d4`

Use the following resource configuration block:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as checkpoint from "@pulumi/checkpoint";

const host = new checkpoint.ManagementHost("host", {
    name: "myhost",
    ipv4Address: "1.1.1.1",
});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_checkpoint as checkpoint

host = checkpoint.ManagementHost("host",
    name="myhost",
    ipv4_address="1.1.1.1")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Checkpoint = Pulumi.Checkpoint;

return await Deployment.RunAsync(() =>
{
    var host = new Checkpoint.ManagementHost("host", new()
    {
        Name = "myhost",
        Ipv4Address = "1.1.1.1",
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/checkpoint/v2/checkpoint"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := checkpoint.NewManagementHost(ctx, "host", &checkpoint.ManagementHostArgs{
			Name:        pulumi.String("myhost"),
			Ipv4Address: pulumi.String("1.1.1.1"),
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
  host:
    type: checkpoint:ManagementHost
    properties:
      name: myhost
      ipv4Address: 1.1.1.1
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.checkpoint.ManagementHost;
import com.pulumi.checkpoint.ManagementHostArgs;
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
        var host = new ManagementHost("host", ManagementHostArgs.builder()
            .name("myhost")
            .ipv4Address("1.1.1.1")
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

Run the following command:

```bash
$ pulumi import checkpoint_management_host.host 9423d36f-2d66-4754-b9e2-e7f4493756d4
```

For more information about `pulumi import` command, please
refer here.
## Tips & Best Practices

This section describes best practices for working with the Check Point provider.

* Use one or more dedicated users for provider operations to make sure minimum permissions are granted.
* Keep on object name uniqueness in your environment.
* Use object name when reference to an object (avoid use of object UID).
* Use post apply scripts (e.g. publish, install policy, logout) to run actions after apply your changes. Pulumi runs in parallel and because of that we can't predict the order of when changes will execute, running post apply scripts will ensure to run last after all changes submitted successfully.
* Create implicit / explicit dependencies between resources or modules. Pulumi uses this dependency information to determine the correct order in which to create the different resources. To do so, it creates a dependency graph of all of the resources defined by the configuration. For more information, please refer here.
* Keep on unique `sessionFileName` when configure more than one provider for authentication purposes.
* Resources and Functions that start with `checkpoint_management_*` using Management API and require set context to `webApi`. For GAIA API resources set context to `gaiaApi`.
* When configure provider context to `gaiaApi` you can run only GAIA resources. Management resources will not be supported.
* Provider state policy is to capture all resource attributes into Pulumi state. All attributes defined in the resource schema are recorded and kept up-to-date in the state. For more information, please refer here.
### Publish best options and practices
#### Trigger field
From version 1.2 the provider was enhanced where a `triggers` field for resource `install-policy`, `publish` and `logout` was added for re-execution if there are any changes in the configuration files.
#### Avoid large bulk publishes
From version 2.5.0 the provider was enhanced with support to auto publish mode using `autoPublishBatchSize` or via the `CHECKPOINT_AUTO_PUBLISH_BATCH_SIZE` environment variable to configure the number of batch size to automatically run publish.
<br>Note: To make sure all changes are published need to do publish explicitly at the end of the execution.
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    checkpoint:apiKey:
        value: admin_api_key
    checkpoint:autoPublishBatchSize:
        value: "100"
    checkpoint:context:
        value: web_api
    checkpoint:server:
        value: chkp-mgmt-srv.local

```
#### Control publish post destroy
From version 2.6.0 the provider was enhanced where a new flag was added `runPublishOnDestroy` to `checkpoint.ManagementPublish` which indicates whether to run publish on destroy.
{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as checkpoint from "@pulumi/checkpoint";

// Make the publish resources dependent of the module and trigger it if there is a change in the configuration files
const publish = new checkpoint.ManagementPublish("publish", {
    triggers: publishTriggers,
    runPublishOnDestroy: true,
}, {
    dependsOn: [policy],
});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_checkpoint as checkpoint

# Make the publish resources dependent of the module and trigger it if there is a change in the configuration files
publish = checkpoint.ManagementPublish("publish",
    triggers=publish_triggers,
    run_publish_on_destroy=True,
    opts = pulumi.ResourceOptions(depends_on=[policy]))
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Checkpoint = Pulumi.Checkpoint;

return await Deployment.RunAsync(() =>
{
    // Make the publish resources dependent of the module and trigger it if there is a change in the configuration files
    var publish = new Checkpoint.ManagementPublish("publish", new()
    {
        Triggers = publishTriggers,
        RunPublishOnDestroy = true,
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            policy,
        },
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/checkpoint/v2/checkpoint"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Make the publish resources dependent of the module and trigger it if there is a change in the configuration files
		_, err := checkpoint.NewManagementPublish(ctx, "publish", &checkpoint.ManagementPublishArgs{
			Triggers:            pulumi.Any(publishTriggers),
			RunPublishOnDestroy: pulumi.Bool(true),
		}, pulumi.DependsOn([]pulumi.Resource{
			policy,
		}))
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
  # Make the publish resources dependent of the module and trigger it if there is a change in the configuration files
  publish:
    type: checkpoint:ManagementPublish
    properties:
      triggers: ${publishTriggers}
      runPublishOnDestroy: true
    options:
      dependsOn:
        - ${policy}
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.checkpoint.ManagementPublish;
import com.pulumi.checkpoint.ManagementPublishArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        // Make the publish resources dependent of the module and trigger it if there is a change in the configuration files
        var publish = new ManagementPublish("publish", ManagementPublishArgs.builder()
            .triggers(publishTriggers)
            .runPublishOnDestroy(true)
            .build(), CustomResourceOptions.builder()
                .dependsOn(policy)
                .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}