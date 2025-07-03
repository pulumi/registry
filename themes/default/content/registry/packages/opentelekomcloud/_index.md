---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/opentelekomcloud/opentelekomcloud/1.36.42/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Opentelekomcloud Provider
meta_desc: Provides an overview on how to configure the Pulumi Opentelekomcloud provider.
layout: package
---

## Generate Provider

The Opentelekomcloud provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider opentelekomcloud/opentelekomcloud
```
## Overview

The Open Telekom Cloud provider is used to interact with the
many resources supported by OpenTelekomCloud. The provider needs to be configured
with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    opentelekomcloud:authUrl:
        value: https://iam.eu-de.otc.t-systems.com/v3
    opentelekomcloud:domainName:
        value: 'TODO: var.domain_name'
    opentelekomcloud:password:
        value: 'TODO: var.password'
    opentelekomcloud:tenantName:
        value: 'TODO: var.tenant_name'
    opentelekomcloud:userName:
        value: 'TODO: var.user_name'

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as opentelekomcloud from "@pulumi/opentelekomcloud";

// Create a web server
const test_server = new opentelekomcloud.ComputeInstanceV2("test-server", {});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    opentelekomcloud:authUrl:
        value: https://iam.eu-de.otc.t-systems.com/v3
    opentelekomcloud:domainName:
        value: 'TODO: var.domain_name'
    opentelekomcloud:password:
        value: 'TODO: var.password'
    opentelekomcloud:tenantName:
        value: 'TODO: var.tenant_name'
    opentelekomcloud:userName:
        value: 'TODO: var.user_name'

```
```python
import pulumi
import pulumi_opentelekomcloud as opentelekomcloud

# Create a web server
test_server = opentelekomcloud.ComputeInstanceV2("test-server")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    opentelekomcloud:authUrl:
        value: https://iam.eu-de.otc.t-systems.com/v3
    opentelekomcloud:domainName:
        value: 'TODO: var.domain_name'
    opentelekomcloud:password:
        value: 'TODO: var.password'
    opentelekomcloud:tenantName:
        value: 'TODO: var.tenant_name'
    opentelekomcloud:userName:
        value: 'TODO: var.user_name'

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Opentelekomcloud = Pulumi.Opentelekomcloud;

return await Deployment.RunAsync(() =>
{
    // Create a web server
    var test_server = new Opentelekomcloud.ComputeInstanceV2("test-server");

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    opentelekomcloud:authUrl:
        value: https://iam.eu-de.otc.t-systems.com/v3
    opentelekomcloud:domainName:
        value: 'TODO: var.domain_name'
    opentelekomcloud:password:
        value: 'TODO: var.password'
    opentelekomcloud:tenantName:
        value: 'TODO: var.tenant_name'
    opentelekomcloud:userName:
        value: 'TODO: var.user_name'

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/opentelekomcloud/opentelekomcloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a web server
		_, err := opentelekomcloud.NewComputeInstanceV2(ctx, "test-server", nil)
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
    opentelekomcloud:authUrl:
        value: https://iam.eu-de.otc.t-systems.com/v3
    opentelekomcloud:domainName:
        value: 'TODO: var.domain_name'
    opentelekomcloud:password:
        value: 'TODO: var.password'
    opentelekomcloud:tenantName:
        value: 'TODO: var.tenant_name'
    opentelekomcloud:userName:
        value: 'TODO: var.user_name'

```
```yaml
resources:
  # Create a web server
  test-server:
    type: opentelekomcloud:ComputeInstanceV2
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    opentelekomcloud:authUrl:
        value: https://iam.eu-de.otc.t-systems.com/v3
    opentelekomcloud:domainName:
        value: 'TODO: var.domain_name'
    opentelekomcloud:password:
        value: 'TODO: var.password'
    opentelekomcloud:tenantName:
        value: 'TODO: var.tenant_name'
    opentelekomcloud:userName:
        value: 'TODO: var.user_name'

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.opentelekomcloud.ComputeInstanceV2;
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
        // Create a web server
        var test_server = new ComputeInstanceV2("test-server");

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Authentication

This provider offers 5 means for authentication.

- User name + Password
- AK/SK
- Token
- Federated
- Assume Role
- OpenStack configuration file
### User name + Password

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    opentelekomcloud:authUrl:
        value: https://iam.eu-de.otc.t-systems.com/v3
    opentelekomcloud:domainName:
        value: 'TODO: var.domain_name'
    opentelekomcloud:password:
        value: 'TODO: var.password'
    opentelekomcloud:tenantName:
        value: 'TODO: var.tenant_name'
    opentelekomcloud:userName:
        value: 'TODO: var.user_name'

```
### AKSK

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    opentelekomcloud:accessKey:
        value: 'TODO: var.access_key'
    opentelekomcloud:authUrl:
        value: https://iam.eu-de.otc.t-systems.com/v3
    opentelekomcloud:domainName:
        value: 'TODO: var.domain_name'
    opentelekomcloud:secretKey:
        value: 'TODO: var.secret_key'
    opentelekomcloud:tenantName:
        value: 'TODO: var.tenant_name'

```
### Temporary AKSK

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    opentelekomcloud:accessKey:
        value: 'TODO: var.access_key'
    opentelekomcloud:authUrl:
        value: https://iam.eu-de.otc.t-systems.com/v3
    opentelekomcloud:domainName:
        value: 'TODO: var.domain_name'
    opentelekomcloud:secretKey:
        value: 'TODO: var.secret_key'
    opentelekomcloud:securityToken:
        value: 'TODO: var.security_token'
    opentelekomcloud:tenantName:
        value: 'TODO: var.tenant_name'

```
### Token

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    opentelekomcloud:authUrl:
        value: https://iam.eu-de.otc.t-systems.com/v3
    opentelekomcloud:domainName:
        value: 'TODO: var.domain_name'
    opentelekomcloud:tenantName:
        value: 'TODO: var.tenant_name'
    opentelekomcloud:token:
        value: 'TODO: var.token'

```

> If token, AK/SK and password are set simultaneously, authentication will be done in the following order:
Token, AK/SK, and Password.
### Federated

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    opentelekomcloud:accessKey:
        value: 'TODO: var.access_key'
    opentelekomcloud:authUrl:
        value: https://iam.eu-de.otc.t-systems.com/v3
    opentelekomcloud:domainName:
        value: 'TODO: var.domain_name'
    opentelekomcloud:secretKey:
        value: 'TODO: var.secret_key'
    opentelekomcloud:securityToken:
        value: 'TODO: var.security_token'
    opentelekomcloud:tenantName:
        value: 'TODO: var.tenant_name'
    opentelekomcloud:token:
        value: 'TODO: var.token'

```
### Assume Role
#### User name + Password

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    opentelekomcloud:agencyDomainName:
        value: 'TODO: var.agency_domain_name'
    opentelekomcloud:agencyName:
        value: 'TODO: var.agency_name'
    opentelekomcloud:authUrl:
        value: https://iam.eu-de.otc.t-systems.com/v3
    opentelekomcloud:delegatedProject:
        value: 'TODO: var.delegated_project'
    opentelekomcloud:domainName:
        value: 'TODO: var.domain_name'
    opentelekomcloud:password:
        value: 'TODO: var.password'
    opentelekomcloud:userName:
        value: 'TODO: var.user_name'

```
#### User ID + Password + TOTP
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    opentelekomcloud:agencyDomainName:
        value: 'TODO: var.agency_domain_name'
    opentelekomcloud:agencyName:
        value: 'TODO: var.agency_name'
    opentelekomcloud:authUrl:
        value: https://iam.eu-de.otc.t-systems.com/v3
    opentelekomcloud:delegatedProject:
        value: 'TODO: var.delegated_project'
    opentelekomcloud:domainName:
        value: 'TODO: var.domain_name'
    opentelekomcloud:passcode:
        value: 'TODO: var.passcode'
    opentelekomcloud:password:
        value: 'TODO: var.password'
    opentelekomcloud:userId:
        value: 'TODO: var.user_id'

```
#### AK/SK

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    opentelekomcloud:accessKey:
        value: 'TODO: var.access_key'
    opentelekomcloud:agencyDomainName:
        value: 'TODO: var.agency_domain_name'
    opentelekomcloud:agencyName:
        value: 'TODO: var.agency_name'
    opentelekomcloud:authUrl:
        value: https://iam.eu-de.otc.t-systems.com/v3
    opentelekomcloud:delegatedProject:
        value: 'TODO: var.delegated_project'
    opentelekomcloud:domainName:
        value: 'TODO: var.domain_name'
    opentelekomcloud:secretKey:
        value: 'TODO: var.secret_key'

```
### Temporary AKSK

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    opentelekomcloud:accessKey:
        value: 'TODO: var.access_key'
    opentelekomcloud:agencyDomainName:
        value: 'TODO: var.agency_domain_name'
    opentelekomcloud:agencyName:
        value: 'TODO: var.agency_name'
    opentelekomcloud:authUrl:
        value: https://iam.eu-de.otc.t-systems.com/v3
    opentelekomcloud:delegatedProject:
        value: 'TODO: var.delegated_project'
    opentelekomcloud:domainName:
        value: 'TODO: var.domain_name'
    opentelekomcloud:secretKey:
        value: 'TODO: var.secret_key'
    opentelekomcloud:securityToken:
        value: 'TODO: var.security_token'

```
#### Token

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    opentelekomcloud:agencyDomainName:
        value: 'TODO: var.agency_domain_name'
    opentelekomcloud:agencyName:
        value: 'TODO: var.agency_name'
    opentelekomcloud:authUrl:
        value: https://iam.eu-de.otc.t-systems.com/v3
    opentelekomcloud:delegatedProject:
        value: 'TODO: var.delegated_project'
    opentelekomcloud:token:
        value: 'TODO: var.token'

```

`token` specified is not the normal token, but must have the authority of `Agent Operator`.
### OpenStack configuration file

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    opentelekomcloud:cloud:
        value: 'TODO: var.cloud_name'

```

`cloud` should be the name of cloud in `clouds.yaml`

See [OpenStack configuration documentation](https://docs.openstack.org/python-openstackclient/latest/configuration/index.html) for details.
## Configuration Reference

The following configuration inputs are supported:

* `accessKey` - (Optional) The access key of the OpenTelekomCloud cloud to use.
  If omitted, the `OS_ACCESS_KEY` environment variable is used.

* `secretKey` - (Optional) The secret key of the OpenTelekomCloud cloud to use.
  If omitted, the `OS_SECRET_KEY` environment variable is used.

* `authUrl` - (Optional; required if `cloud` is not specified) The Identity
  authentication URL. If omitted, the `OS_AUTH_URL` environment variable is used.

* `cloud` - (Optional; required if `authUrl` is not specified) An entry in a
  `clouds.yaml` file. See the OpenStack `os-client-config`
  [documentation](https://docs.openstack.org/os-client-config/latest/user/configuration.html)
  for more information about `clouds.yaml` files. If omitted, the `OS_CLOUD`
  environment variable is used.

* `userName` - (Optional) The Username to login with. If omitted, the
  `OS_USERNAME` environment variable is used.

* `userId` - (Optional) The ID of the user to login with. Required when TOTP is used (`passcode` is not empty).
  If `userId` is set, `userName` is ignored.

* `tenantName` - (Optional) The Name of the Tenant (Identity v2) or Project
  (Identity v3) to login with. If omitted, the `OS_TENANT_NAME` or
  `OS_PROJECT_NAME` environment variable are used.

* `region` - (Optional) The name of the region to be used. Required for some resources
  (e.g. `s3Bucket`) in case no tenant name provided and no region is defined in the
  resource. If omitted, the `OS_REGION` or `OS_REGION_NAME` environment variables are used.

* `password` - (Optional) The Password to login with. If omitted, the
  `OS_PASSWORD` environment variable is used.

* `token` - (Optional; Required if not using `userName` and `password`)
  A token is an expiring, temporary means of access issued via the Keystone
  service. By specifying a token, you do not have to specify a username/password
  combination, since the token was already created by a username/password out of
  band of Pulumi. If omitted, the `OS_AUTH_TOKEN` or `OS_TOKEN` environment
  variable is used.

* `securityToken` - (Optional) Security token required to authenticate with temporary AK/SK.

* `passcode` - (Optional) One-time password provided by your authentication app.

->
Please note that MFA requires `userId` to be used. Setting `userName` won't work.

* `domainName` - (Optional) The Name of the Domain to scope to (Identity v3).
  If omitted, the following environment variables are checked (in this order):
  `OS_USER_DOMAIN_NAME`, `OS_PROJECT_DOMAIN_NAME`, `OS_DOMAIN_NAME`,
  `DEFAULT_DOMAIN`.

* `insecure` - (Optional) Trust self-signed SSL certificates. If omitted, the
  `OS_INSECURE` environment variable is used.

* `cacertFile` - (Optional) Specify a custom CA certificate when communicating
  over SSL. You can specify either a path to the file or the contents of the
  certificate. If omitted, the `OS_CACERT` environment variable is used.

* `cert` - (Optional) Specify client certificate file for SSL client
  authentication. You can specify either a path to the file or the contents of
  the certificate. If omitted the `OS_CERT` environment variable is used.

* `key` - (Optional) Specify client private key file for SSL client
  authentication. You can specify either a path to the file or the contents of
  the key. If omitted the `OS_KEY` environment variable is used.

* `endpointType` - (Optional) Specify which type of endpoint to use from the
  service catalog. It can be set using the `OS_ENDPOINT_TYPE` environment
  variable. If not set, public endpoints is used.

* `swauth` - (Optional) Set to `true` to authenticate against Swauth, a
  Swift-native authentication system. If omitted, the `OS_SWAUTH` environment
  variable is used. You must also set `username` to the Swauth/Swift username
  such as `username:project`. Set the `password` to the Swauth/Swift key.
  Finally, set `authUrl` as the location of the Swift service.

> This will only work when used with the OpenTelekomCloud Object Storage resources.

* `agencyName` - (Optional) if authorized by assume role, it must be set. The
  name of agency.

* `agencyDomainName` - (Optional) if authorized by assume role, it must be set.
  The name of domain who created the agency (Identity v3).

* `delegatedProject` - (Optional) The name of delegated project (Identity v3).

* `maxRetries` - (Optional) Maximum number of retries of HTTP requests failed
  due to connection issues.

* `allowReauth` - (Optional) If set to false, authorization won't be performed automatically if the
  initial auth token get expired. It can be set using the `OS_ALLOW_REAUTH` environment variable.
  Default: `true`.

* `maxBackoffRetries` - (Optional) Maximum number of retries of HTTP requests failed
  due to reaching the rate limit. It can be set using the `OS_MAX_BACKOFF_RETRIES` environment
  variable. If not set, default value is used.
  Default: `5`

* `backoffRetryTimeout` - (Optional) Timeout in seconds for backoff retry due to reaching the rate limit.
  It can be set using the `OS_BACKOFF_RETRY_TIMEOUT` environment
  variable. If not set, default value is used.
  Default: `60` seconds.
## Creating an issue

Issues
can be used to keep track of bugs, enhancements, or other requests.
See the github help [here](https://help.github.com/articles/creating-an-issue/)