---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-scm/v0.4.3/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi/pulumi-scm/blob/v0.4.3/docs/_index.md
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Strata Cloud Manager Provider
meta_desc: Provides an overview on how to configure the Pulumi Strata Cloud Manager provider.
layout: package
---

## Installation

The Strata Cloud Manager provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/scm`](https://www.npmjs.com/package/@pulumi/scm)
* Python: [`pulumi-scm`](https://pypi.org/project/pulumi-scm/)
* Go: [`github.com/pulumi/pulumi-scm/sdk/go/scm`](https://github.com/pulumi/pulumi-scm)
* .NET: [`Pulumi.Scm`](https://www.nuget.org/packages/Pulumi.Scm)
* Java: [`com.pulumi/scm`](https://central.sonatype.com/artifact/com.pulumi/scm)

## Overview

The `scm` provider provides resources and functions to manage and query Strata Cloud Manager.

This provider covers the following aspects of Strata Cloud Manager:
* NGFW & Prisma Access
## Beta Release Disclaimer

**This software is a pre-release version and is not ready for production use.**

* **No Warranty:** This software is provided "as is," without any warranty of any kind, either expressed or implied, including, but not limited to, the implied warranties of merchantability and fitness for a particular purpose.
* **Instability:** The beta software may contain defects, may not operate correctly, and may be substantially modified or withdrawn at any time.
* **Limitation of Liability:** In no event shall the authors or copyright holders be liable for any claim, damages, or other liability, whether in an action of contract, tort, or otherwise, arising from, out of, or in connection with the beta software or the use or other dealings in the beta software.
* **Feedback:** We encourage and appreciate your feedback and bug reports. However, you acknowledge that any feedback you provide is non-confidential.

By using this software, you agree to these terms.
## Release Notes
### v1.0.4
#### ENHANCEMENTS

* resource/scm_logical_router: Added additional tests and examples (resource, data-source)
* resource/scm_decryption_profile: Added additional tests and examples (resource, data-source)
* resource/scm_interface_management_profile: Added additional tests and examples (resource, data-source)
* resource/scm_radius_server_profile: Added additional tests and examples (resource, data-source)
* resource/scm_dns_security_profile: Added additional tests and examples (resource, data-source)
* resource/scm_ethernet_interface: Added additional tests
* resource/scm_vlan_interface: Added additional tests
* resource/scm_loopback_interface: Added additional tests
* resource/scm_tunnel_interface: Added additional tests
* resource/scm_zone: Added additional tests
#### BUG FIXES

* resource/scm_nat_rule: Fixed various errors in model definition
* resource/scm_radius_server_profile: Fixed missing `name` property in model definition (#35)
* resource/scm_vulnerability_protection_profile: Added support for missing category options (#42)
* resource/scm_aggregate_interface: Added correct api-path (#50) and fixed missing model properties
* resource/scm_logical_router: Fixed issue with model definition that caused failures during crud operations (#58)
* resource/scm_interface_management_profile: Fixed issue with model definition that caused payload unmarshal error (#59)
### v1.0.3
#### FEATURES

* Rule re-ordering: Support for placing a rule in a desired location while the resource is managed through pulumi
#### ENHANCEMENTS

* Improved Error Handling: Include API Errors by default in pulumi logs
* Doc Improvement: Adapted default folder for examples from `Shared` (Prisma Access)  to `All` (Global)
* resource/scm_ethernet_interface: Added additional tests and examples (resource, data-source)
* resource/scm_loopback_interface: Added additional tests and examples (resource, data-source)
* resource/scm_tunnel_interface: Added additional tests and examples (resource, data-source)
* resource/scm_layer2_subinterface: Added additional tests and examples (resource, data-source)
* resource/scm_layer3_subinterface: Added additional tests and examples (resource, data-source)
* resource/scm_vlan_interface: Added additional tests and examples (resource, data-source)
* resource/scm_folder: Added additional tests and examples (resource, data-source)
* resource/scm_snippet: Added additional tests and examples (resource, data-source)
* resource/scm_app_override_rule: Added additional tests and examples (resource, data-source)
* resource/scm_authentication_rule: Added additional tests and examples (resource, data-source)
* resource/scm_decryption_rule: Added additional tests and examples (resource, data-source)
* resource/scm_qos_rule: Added additional tests and examples (resource, data-source)
* resource/scm_nat_rule: Added additional tests and examples (resource, data-source)
* resource/scm_pdf_rule: Added additional tests and examples (resource, data-source)
* resource/scm_security_rule: Added additional tests and examples (resource, data-source)
* resource/scm_decryption_profile: Added additional tests and examples (resource, data-source)
* resource/scm_radius_server_profile: Fixed issue related to the All Caps property and the unintended underscores.
#### BUG FIXES

* resource/scm_service_connection: Fixed pushing a ikeGateway object for Service Connection (#21)
* resource/scm_addresses: Fixed referesh operation for scm.Address (#22)
* resource/scm_security_rule: Fixed Rule ordering (#27)
* resource/scm_ethernet: Fixed errors in model definition that resulted in resource failures (#36)
* resource/scm_radius_server_profile: Fixed nat rule bidirectional property (#36)
* resource/scm_external_dynamic_list: Added support for handling sestive values (#38)
* resource/scm_security_rule: Added Support for 'tenant_restrictions' field (#41)
* resource/scm_ethernet: Fixed errors in `layer3` model definition that resulted in resource failures (#51, #52)
* resource/scm_zone: Fixed errors in model definition that resulted in resource failures (#53)
* resource/scm_tunnel_interface: Fixed errors in model definition that resulted in resource failures (#55)
### v1.0.2
#### FEATURES

* Added support for recreating for out-of-band deleted resources
* **New Resource**: `scm.Folder`
* **New Resource**: `scm.Snippet`
* **New Resource**: `scm.Folder`
#### ENHANCEMENTS

* resource/scm_antispyware_profile: Added additional tests and examples (resource, data-source)
* resource/scm_vulnerability_protection_profile: Added additional tests and examples (resource, data-source)
* resource/scm_service_connection: Added additional tests and examples (resource, data-source)
* resource/scm_address_group: Added additional examples
* resource/service_group: Added additional examples
* resource/service: Added additional examples
#### BUG FIXES

* resource/scm_antispyware_profile: Fixed min length of name in model definition
* resource/scm_service_connection: Fixed incorrect folder reference in model definition
* resource/<multiple-resources>: Fixed issue with pattern matching where hyphen was not correctly escaped
### v1.0.1
#### ENHANCEMENTS

* Beta disclaimer added
### v1.0.0
#### FEATURES

* Initial Release - No backwards compatibility to v0.x.y
## Warranty

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

THIS SOFTWARE IS RELEASED AS A PROOF OF CONCEPT FOR EXPERIMENTAL PURPOSES ONLY. USE IT AT OWN RISK. THIS SOFTWARE IS NOT SUPPORTED.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    scm:authUrl:
        value: ""
    scm:clientId:
        value: ""
    scm:clientSecret:
        value: ""
    scm:host:
        value: ""
    scm:logging:
        value: ""
    scm:scope:
        value: ""

```
```typescript
import * as pulumi from "@pulumi/pulumi";

```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    scm:authUrl:
        value: ""
    scm:clientId:
        value: ""
    scm:clientSecret:
        value: ""
    scm:host:
        value: ""
    scm:logging:
        value: ""
    scm:scope:
        value: ""

```
```python
import pulumi

```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    scm:authUrl:
        value: ""
    scm:clientId:
        value: ""
    scm:clientSecret:
        value: ""
    scm:host:
        value: ""
    scm:logging:
        value: ""
    scm:scope:
        value: ""

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;

return await Deployment.RunAsync(() =>
{
});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    scm:authUrl:
        value: ""
    scm:clientId:
        value: ""
    scm:clientSecret:
        value: ""
    scm:host:
        value: ""
    scm:logging:
        value: ""
    scm:scope:
        value: ""

```
```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
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
    scm:authUrl:
        value: ""
    scm:clientId:
        value: ""
    scm:clientSecret:
        value: ""
    scm:host:
        value: ""
    scm:logging:
        value: ""
    scm:scope:
        value: ""

```
```yaml
{}
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    scm:authUrl:
        value: ""
    scm:clientId:
        value: ""
    scm:clientSecret:
        value: ""
    scm:host:
        value: ""
    scm:logging:
        value: ""
    scm:scope:
        value: ""

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
    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Provider Parameter Priority

There are multiple ways to specify the provider's parameters.  If overlapping values are configured for the provider, then this is the resolution order:

1. Statically configured in the provider configuration
2. Environment variable (where applicable)
3. Taken from the JSON config file
## Configuration Reference

- `authFile` (String) The file path to the JSON file with auth creds for SCM.
- `authUrl` (String) The URL to send auth credentials to which will return a JWT. Default: `https://auth.apps.paloaltonetworks.com/auth/v1/oauth2/access_token`. Environment variable: `SCM_AUTH_URL`. JSON config file variable: `authUrl`.
- `clientId` (String) The client ID for the connection. Environment variable: `SCM_CLIENT_ID`. JSON config file variable: `clientId`.
- `clientSecret` (String, Sensitive) The client secret for the connection. Environment variable: `SCM_CLIENT_SECRET`. JSON config file variable: `clientSecret`.
- `headers` (Map of String) Custom HTTP headers to be sent with all API commands. Environment variable: `SCM_HEADERS`. JSON config file variable: `headers`.
- `host` (String) The hostname of Strata Cloud Manager API. Default: `api.sase.paloaltonetworks.com`. Environment variable: `SCM_HOST`. JSON config file variable: `host`.
- `logging` (String) The logging level of the provider and the underlying communication. Default: `quiet`. Environment variable: `SCM_LOGGING`. JSON config file variable: `logging`.
- `port` (Number) The port number to use for API commands, if non-standard for the given protocol. Environment variable: `SCM_PORT`. JSON config file variable: `port`.
- `protocol` (String) The protocol to use for SCM. This should be 'http' or 'https'. Default: `https`. Environment variable: `SCM_PROTOCOL`. JSON config file variable: `protocol`.
- `scope` (String) The client scope. Environment variable: `SCM_SCOPE`. JSON config file variable: `scope`.