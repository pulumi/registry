---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-vault/v7.3.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: HashiCorp Vault Provider
meta_desc: Provides an overview on how to configure the Pulumi HashiCorp Vault provider.
layout: package
---

## Installation

The HashiCorp Vault provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/vault`](https://www.npmjs.com/package/@pulumi/vault)
* Python: [`pulumi-vault`](https://pypi.org/project/pulumi-vault/)
* Go: [`github.com/pulumi/pulumi-vault/sdk/v7/go/vault`](https://github.com/pulumi/pulumi-vault)
* .NET: [`Pulumi.Vault`](https://www.nuget.org/packages/Pulumi.Vault)
* Java: [`com.pulumi/vault`](https://central.sonatype.com/artifact/com.pulumi/vault)

## Overview

The Vault provider allows Pulumi to read from, write to, and configure
[HashiCorp Vault](https://vaultproject.io/).
## Using Vault credentials in Pulumi configuration

> **Important** It is important to ensure that the Vault token
has a long enough `time-to-live` to allow for all Vault resources to
be successfully provisioned. In the case where the `TTL` is insufficient,
you may encounter unexpected permission denied errors.
See [Vault Token TTLs](https://vaultproject.io/docs/concepts/tokens#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
for more details.

Most Pulumi providers require credentials to interact with a third-party
service that they wrap. This provider allows such credentials to be obtained
from Vault, which means that operators or systems running Pulumi need
only access to a suitably-privileged Vault token in order to temporarily
lease the credentials for other providers.

To reduce the exposure of secrets, the provider requests a Vault token
with a relatively-short TTL (20 minutes, by default) which in turn means
that where possible Vault will revoke any issued credentials after that
time, but in particular it is unable to retract any static secrets such as
those stored in Vault's "generic" secret backend.

The requested token TTL can be controlled by the `maxLeaseTtlSeconds`
provider argument described below.

Except as otherwise noted, the resources that read secrets from Vault
are designed such that they require only the *read* capability on the relevant
resources.
## Provider Arguments

The provider configuration block accepts the following arguments.
In most cases it is recommended to set them via the indicated environment
variables in order to keep credential information out of the configuration.

* `address` - (Required) Origin URL of the Vault server. This is a URL
  with a scheme, a hostname and a port but with no path. May be set
  via the `VAULT_ADDR` environment variable.

* `addAddressToEnv` - (Optional) If `true` the environment variable
  `VAULT_ADDR` in the Pulumi process environment will be set to the
  value of the `address` argument from this provider. By default, this is false.

* `token` - (Optional) Vault token that will be used by Pulumi to
  authenticate. May be set via the `VAULT_TOKEN` environment variable.
  If none is otherwise supplied, Pulumi will attempt to read it from
  `~/.vault-token` (where the vault command stores its current token).
  Pulumi will issue itself a new token that is a child of the one given,
  with a short TTL to limit the exposure of any requested secrets, unless
  `skipChildToken` is set to `true` (see below). Note that
  the given token must have the update capability on the auth/token/create
  path in Vault in order to create child tokens.  A token is required for
  the provider.  A token can explicitly set via token argument, alternatively
  a token can be dynamically set via an `auth_login*` block.

* `tokenName` - (Optional) Token name, that will be used by Pulumi when
  creating the child token (`displayName`). This is useful to provide a reference of the
  Pulumi run traceable in vault audit log, e.g. commit hash or id of the CI/CD
  execution job. May be set via the `VAULT_TOKEN_NAME` environment variable.
  Default value will be `pulumi` if not set or empty.

* `caCertFile` - (Optional) Path to a file on local disk that will be
  used to validate the certificate presented by the Vault server.
  May be set via the `VAULT_CACERT` environment variable.

* `caCertDir` - (Optional) Path to a directory on local disk that
  contains one or more certificate files that will be used to validate
  the certificate presented by the Vault server. May be set via the
  `VAULT_CAPATH` environment variable.

* `authLoginUserpass` - (Optional) Utilizes the `userpass` authentication engine. *See usage details below.*

* `authLoginAws` - (Optional) Utilizes the `aws` authentication engine. *See usage details below.*

* `authLoginCert` - (Optional) Utilizes the `cert` authentication engine. *See usage details below.*

* `authLoginGcp` - (Optional) Utilizes the `gcp` authentication engine. *See usage details below.*

* `authLoginKerberos` - (Optional) Utilizes the `kerberos` authentication engine. *See usage details below.*

* `authLoginRadius` - (Optional) Utilizes the `radius` authentication engine. *See usage details below.*

* `authLoginOci` - (Optional) Utilizes the `oci` authentication engine. *See usage details below.*

* `authLoginOidc` - (Optional) Utilizes the `oidc` authentication engine. *See usage details below.*

* `authLoginJwt` - (Optional) Utilizes the `jwt` authentication engine. *See usage details below.*

* `authLoginAzure` - (Optional) Utilizes the `azure` authentication engine. *See usage details below.*

* `authLoginTokenFile` - (Optional) Utilizes a local file containing a Vault token. *See usage details below.*
* `authLogin` - (Optional) A configuration block, described below, that
  attempts to authenticate using the `auth/<method>/login` path to
  acquire a token which Pulumi will use. Pulumi still issues itself
  a limited child token using auth/token/create in order to enforce a short
  TTL and limit exposure. *See usage details below.*

* `skipTlsVerify` - (Optional) Set this to `true` to disable verification
  of the Vault server's TLS certificate. This is strongly discouraged except
  in prototype or development environments, since it exposes the possibility
  that Pulumi can be tricked into writing secrets to a server controlled
  by an intruder. May be set via the `VAULT_SKIP_VERIFY` environment variable.

* `tlsServerName` - (Optional) Name to use as the SNI host when connecting
  via TLS. May be set via the `VAULT_TLS_SERVER_NAME` environment variable.

* `skipChildToken` - (Optional) Set this to `true` to disable
  creation of an intermediate ephemeral Vault token for Pulumi to
  use. Enabling this is strongly discouraged since it increases
  the potential for a renewable Vault token being exposed in clear text.
  Only change this setting when the provided token cannot be permitted to
  create child tokens and there is no risk of exposure from the output of
  Pulumi. May be set via the `TERRAFORM_VAULT_SKIP_CHILD_TOKEN` environment
  variable. **Note**: Setting to `true` will cause `tokenName`
  and `maxLeaseTtlSeconds` to be ignored.
  Please see Using Vault credentials in Pulumi configuration
  before enabling this setting.

* `maxLeaseTtlSeconds` - (Optional) Used as the duration for the
  intermediate Vault token Pulumi issues itself, which in turn limits
  the duration of secret leases issued by Vault. Defaults to 20 minutes
  and may be set via the `TERRAFORM_VAULT_MAX_TTL` environment variable.
  See the section above on *Using Vault credentials in Pulumi configuration*
  for the implications of this setting.

* `maxRetries` - (Optional) Used as the maximum number of retries when a 5xx
  error code is encountered. Defaults to `2` retries and may be set via the
  `VAULT_MAX_RETRIES` environment variable.

* `maxRetriesCcc` - (Optional) Maximum number of retries for *Client Controlled Consistency*
  related operations. Defaults to `10` retries and may also be set via the
  `VAULT_MAX_RETRIES_CCC` environment variable. See
  [Vault Eventual Consistency](https://www.vaultproject.io/docs/enterprise/consistency#vault-eventual-consistency)
  for more information.*As of Vault Enterprise 1.10 changing this parameter should no longer be required
  See [Vault Eventual Consistency - Vault 1.10 Mitigations](https://www.vaultproject.io/docs/enterprise/consistency#vault-1-10-mitigations)
  for more information.*

* `namespace` - (Optional) Set the namespace to use. May be set via the
  `VAULT_NAMESPACE` environment variable.
  See [namespaces](https://www.vaultproject.io/docs/enterprise/namespaces) for more info.
  *Available only for Vault Enterprise*.

* `useRootNamespace` - (Optional) Authenticate to the root Vault namespace. Conflicts with `namespace`.

* `skipGetVaultVersion` - (Optional) Skip the dynamic fetching of the Vault server version.
  Set to `true` when the */sys/seal-status* API endpoint is not available. See vaultVersionOverride
  for related info

* `vaultVersionOverride` - (Optional) Override the target Vault server semantic version.
  Normally the version is dynamically set from the */sys/seal-status* API endpoint. In the case where this endpoint
  is not available an override can be specified here.

> Setting the `vaultVersionOverride` determines Vault server's API compatability, so
it's important that the value specified here matches the target server. It is recommended to
only ever use this option in the case where the server version cannot be dynamically determined.

* `headers` - (Optional) A configuration block, described below, that provides headers
  to be sent along with all requests to the Vault server.  This block can be specified
  multiple times.

The `headers` configuration block accepts the following arguments:

* `name` - (Required) The name of the header.

* `value` - (Required) The value of the header.
## Vault Authentication Configuration Options

The Vault provider supports the following Vault authentication engines.
### Userpass

Provides support for authenticating to Vault using the Username & Password authentication engine.

*For more details see:
[Userpass Auth Method (HTTP API)](https://www.vaultproject.io/api-docs/auth/userpass#userpass-auth-method-http-api)*

The `authLoginUserpass` configuration block accepts the following arguments:

* `namespace` - (Optional) The path to the namespace that has the mounted auth method.
  This defaults to the root namespace. Cannot contain any leading or trailing slashes.
  *Available only for Vault Enterprise*.

* `useRootNamespace` - (Optional) Authenticate to the root Vault namespace. Conflicts with `namespace`.

* `mount` - (Optional) The name of the authentication engine mount.Default: `userpass`

* `username` - (Required) The username to log into Vault with.
  Can be specified with the `TERRAFORM_VAULT_USERNAME` environment variable.

* `password` - (Optional) The password to log into Vault with.
  Can be specified with the `TERRAFORM_VAULT_PASSWORD` environment variable. *Cannot be specified with `passwordFile`*.

* `passwordFile` - (Optional) A file containing the password to log into Vault with.
  Can be specified with the `TERRAFORM_VAULT_PASSWORD_FILE` environment variable. *Cannot be specified with `password`*
### AWS

Provides support for authenticating to Vault using the AWS authentication engine.

*For more details see:
[AWS Auth Method (API)](https://www.vaultproject.io/api-docs/auth/aws#aws-auth-method-api)*

The `authLoginAws` configuration block accepts the following arguments:

* `namespace` - (Optional) The path to the namespace that has the mounted auth method.
  This defaults to the root namespace. Cannot contain any leading or trailing slashes.
  *Available only for Vault Enterprise*.

* `useRootNamespace` - (Optional) Authenticate to the root Vault namespace. Conflicts with `namespace`.

* `mount` - (Optional) The name of the authentication engine mount.Default: `aws`

* `role` - (Required) The name of the role against which the login is being attempted.

* `awsAccessKeyId` - (Optional) The AWS access key ID.*Can be specified with the `AWS_ACCESS_KEY_ID` environment variable.*

* `awsSecretAccessKey` - (Optional) The AWS secret access key.*Can be specified with the `AWS_SECRET_ACCESS_KEY` environment variable.*

* `awsSessionToken` - (Optional) The AWS session token.*Can be specified with the `AWS_SESSION_TOKEN` environment variable.*

* `awsProfile` - (Optional) The name of the AWS profile.*Can be specified with the `AWS_PROFILE` environment variable.*

* `awsSharedCredentialsFile` - (Optional) Path to the AWS shared credentials file.*Can be specified with the `AWS_SHARED_CREDENTIALS_FILE` environment variable.*

* `awsWebIdentityTokenFile` - (Optional) Path to the file containing an OAuth 2.0 access token or OpenID
  Connect ID token.*Can be specified with the `AWS_WEB_IDENTITY_TOKEN_FILE` environment variable.*

* `awsRegion` - (Optional) The AWS region.*Can be specified with the `AWS_REGION` or `AWS_DEFAULT_REGION` environment variables.*

* `awsRoleArn` - (Optional) The ARN of the AWS Role to assume. *Used during STS AssumeRole**Can be specified with the `AWS_ROLE_ARN` environment variable.*

* `awsRoleSessionName` - (Optional) Specifies the name to attach to the AWS role session. *Used during STS AssumeRole**Can be specified with the `AWS_ROLE_SESSION_NAME` environment variable.*

* `awsStsEndpoint` - (Optional) The STS endpoint URL.

* `awsIamEndpoint` - (Optional) The IAM endpoint URL.

* `headerValue` - (Optional) The Vault header value to include in the STS signing request.
### TLS Certificate

Provides support for authenticating to Vault using the TLS Certificate authentication engine.

*For more details see:
[TLS Certificate Auth Method (API)](https://www.vaultproject.io/api-docs/auth/cert#tls-certificate-auth-method-api)*

The `authLoginCert` configuration block accepts the following arguments:

* `namespace` - (Optional) The path to the namespace that has the mounted auth method.
  This defaults to the root namespace. Cannot contain any leading or trailing slashes.
  *Available only for Vault Enterprise*.

* `useRootNamespace` - (Optional) Authenticate to the root Vault namespace. Conflicts with `namespace`.

* `mount` - (Optional) The name of the authentication engine mount.
  Default: `cert`

* `name` - (Optional) Authenticate against only the named certificate role.

* `certFile` - (Required) Path to a file on local disk that contains the
  PEM-encoded certificate to present to the server.

* `keyFile` - (Required) Path to a file on local disk that contains the
  PEM-encoded private key for which the authentication certificate was issued.

*This login configuration honors the top-level TLS configuration parameters:
ca_cert_file, ca_cert_dir, skip_tls_verify,
tls_server_name*
### GCP

Provides support for authenticating to Vault using the Google Cloud Auth engine.

*For more details see:
[Google Cloud Auth Method (API)](https://www.vaultproject.io/api-docs/auth/gcp#google-cloud-auth-method-api)*

The `authLoginGcp` configuration block accepts the following arguments:

* `namespace` - (Optional) The path to the namespace that has the mounted auth method.
  This defaults to the root namespace. Cannot contain any leading or trailing slashes.
  *Available only for Vault Enterprise*.

* `useRootNamespace` - (Optional) Authenticate to the root Vault namespace. Conflicts with `namespace`.

* `mount` - (Optional) The name of the authentication engine mount.Default: `gcp`

* `role` - (Required) The name of the role against which the login is being attempted.

* `jwt` - (Optional) The signed JSON Web Token against which the login is being attempted.

* `credentials` - (Optional) Path to the Google Cloud credentials to use when getting the signed
  JWT token from the IAM service.*conflicts with `jwt`*

* `serviceAccount` - (Optional) Name of the service account to issue the JWT token for.*requires `credentials`*

*This login configuration will attempt to get a signed JWT token if `jwt` is not specified.
It supports both the IAM and GCE meta-data services as the token source.*
### Kerberos

Provides support for authenticating to Vault using the Kerberos Auth engine.

*For more details see:
[Kerberos Auth Method (API)](https://www.vaultproject.io/api-docs/auth/kerberos#kerberos-auth-method-api)*

The `authLoginKerberos` configuration block accepts the following arguments:

* `namespace` - (Optional) The path to the namespace that has the mounted auth method.
  This defaults to the root namespace. Cannot contain any leading or trailing slashes.
  *Available only for Vault Enterprise*.

* `useRootNamespace` - (Optional) Authenticate to the root Vault namespace. Conflicts with `namespace`.

* `mount` - (Optional) The name of the authentication engine mount.Default: `kerberos`

* `token` - (Optional) Simple and Protected GSSAPI Negotiation Mechanism (SPNEGO) token.
  Can be specified with the `KRB_SPNEGO_TOKEN` environment variable.

* `username` - (Conflicts with `token`)  The username to login into Kerberos with.

* `service` - (Conflicts with `token`) The service principle name.

* `realm` - (Conflicts with `token`) The Kerberos server's authoritative authentication domain.

* `krb5confPath` - (Conflicts with `token`) A valid Kerberos configuration file e.g. /etc/krb5.conf.
  Can be specified with the `KRB5_CONFIG` environment variable.

* `keytabPath` - (Conflicts with `token`)  The Kerberos keytab file containing the entry of the login entity.
  Can be specified with the `KRB_KEYTAB` environment variable.

* `disableFastNegotiation` - (Conflicts with `token`) Disable the Kerberos FAST negotiation.

* `removeInstanceName` - (Conflicts with `token`) Strip the host from the username found in the keytab.

*This login configuration will attempt to get a SPNEGO init token from the `service` domain if `token` is not specified.
The following fields are required when token is not specified:
`username`, `service`, `realm`, `krb5confPath`, `keytabPath`*
### Radius

Provides support for authenticating to Vault using the Radius Auth engine.

*For more details see:
[Radius Auth Method (API)](https://www.vaultproject.io/api-docs/auth/radius#radius-auth-method-api)*

The `authLoginRadius` configuration block accepts the following arguments:

* `namespace` - (Optional) The path to the namespace that has the mounted auth method.
  This defaults to the root namespace. Cannot contain any leading or trailing slashes.
  *Available only for Vault Enterprise*.

* `useRootNamespace` - (Optional) Authenticate to the root Vault namespace. Conflicts with `namespace`.

* `mount` - (Optional) The name of the authentication engine mount.Default: `radius`

* `username` - (Required) The username to Radius username to login into Vault with.

* `password` - (Required) The password for the Radius `username` to login into Vault with.
### OCI

Provides support for authenticating to Vault using the OCI (Oracle Cloud Infrastructure) Auth engine.

*For more details see:
[OCI Auth Method (API)](https://www.vaultproject.io/api-docs/auth/oci#oci-auth-method-api)*

The `authLoginOci` configuration block accepts the following arguments:

* `namespace` - (Optional) The path to the namespace that has the mounted auth method.
  This defaults to the root namespace. Cannot contain any leading or trailing slashes.
  *Available only for Vault Enterprise*.

* `useRootNamespace` - (Optional) Authenticate to the root Vault namespace. Conflicts with `namespace`.

* `mount` - (Optional) The name of the authentication engine mount.Default: `oci`

* `role` - (Required) The name of the role against which the login is being attempted.

* `authType` - (Required) The OCI authentication type to use. Valid choices are: *apikeys*, *instance*
### OIDC

Provides support for authenticating to Vault using the OIDC Auth engine.

> Use of this login method requires access to a web browser on the host machine in
order to complete the authorization flow.

*For more details see the OIDC specific documentation here:
[OIDC/JWT Auth Method (API)](https://www.vaultproject.io/api-docs/auth/jwt#jwt-oidc-auth-method-api)*

The `authLoginOidc` configuration block accepts the following arguments:

* `namespace` - (Optional) The path to the namespace that has the mounted auth method.
  This defaults to the root namespace. Cannot contain any leading or trailing slashes.
  *Available only for Vault Enterprise*.

* `useRootNamespace` - (Optional) Authenticate to the root Vault namespace. Conflicts with `namespace`.

* `mount` - (Optional) The name of the authentication engine mount.Default: `oidc`

* `role` - (Required) The name of the role against which the login is being attempted.

* `callbackListenerAddress` - (Optional) The callback listener's address. *Must be a valid URI without the path.*

* `callbackAddress` - (Optional)  The callback address. *Must be a valid URI without the path.*
### JWT

Provides support for authenticating to Vault using the JWT Auth engine.

*For more details see the JWT specific documentation here:
[OIDC/JWT Auth Method (API)](https://www.vaultproject.io/api-docs/auth/jwt#jwt-oidc-auth-method-api)*

The `authLoginJwt` configuration block accepts the following arguments:

* `namespace` - (Optional) The path to the namespace that has the mounted auth method.
  This defaults to the root namespace. Cannot contain any leading or trailing slashes.
  *Available only for Vault Enterprise*.

* `useRootNamespace` - (Optional) Authenticate to the root Vault namespace. Conflicts with `namespace`.

* `mount` - (Optional) The name of the authentication engine mount.Default: `jwt`

* `role` - (Required) The name of the role against which the login is being attempted.

* `jwt` - (Required) The signed JSON Web Token against which the login is being attempted.*Can be specified with the `TERRAFORM_VAULT_AUTH_JWT` environment variable.*
### Azure

Provides support for authenticating to Vault using the Azure Auth engine.

*For more details see the Azure specific documentation here:
[Azure Auth Method (API)](https://www.vaultproject.io/api-docs/auth/azure#azure-auth-method-api)*

The `authLoginAzure` configuration block accepts the following arguments:

* `namespace` - (Optional) The path to the namespace that has the mounted auth method.
  This defaults to the root namespace. Cannot contain any leading or trailing slashes.
  *Available only for Vault Enterprise*.

* `useRootNamespace` - (Optional) Authenticate to the root Vault namespace. Conflicts with `namespace`.

* `mount` - (Optional) The name of the authentication engine mount.Default: `azure`

* `role` - (Required) The name of the role against which the login is being attempted.

* `jwt` - (Optional) The signed JSON Web Token against which the login is being attempted.
  If not provided a token will be created from Azure's managed identities for Azure resources API.
  *Can be specified with the `TERRAFORM_VAULT_AZURE_AUTH_JWT` environment variable.*

* `subscriptionId` - (Required) The subscription ID for the machine that generated the MSI token.
  This information can be obtained through instance metadata.

* `resourceGroupName` - (Required) The resource group for the machine that generated the MSI token.
  This information can be obtained through instance metadata.

* `vmName` - (Optional) The virtual machine name for the machine that generated the MSI token.
  This information can be obtained through instance metadata.

* `vmssName` - (Optional) The virtual machine scale set name for the machine that generated
  the MSI token. This information can be obtained through instance metadata.

* `tenantId` - (Optional) Provides the tenant ID to use in a multi-tenant authentication scenario.

* `clientId` - (Optional) The identity's client ID.

* `scope` - (Optional) The scopes to include in the token request. Defaults to `https://management.azure.com/`
### Token File

Provides support for "authenticating" to Vault using a local file containing a Vault token.

> Using `authLoginTokenFile` is not recommended, since it relies on a Vault token that is persisted to disk.
Please ensure you have processes in place that will remove the token file between Pulumi executions.

The `authLoginTokenFile` configuration block accepts the following arguments:

* `namespace` - (Optional) The path to the namespace that has the mounted auth method.
  This defaults to the root namespace. Cannot contain any leading or trailing slashes.
  *Available only for Vault Enterprise*.

* `useRootNamespace` - (Optional) Authenticate to the root Vault namespace. Conflicts with `namespace`.

* `filename` - (Required) The filename containing a Vault token. The file must contain a single Vault token
  and be user readable e.g. perms=`0600`. May be set via the `TERRAFORM_VAULT_TOKEN_FILENAME`
  environment variable.
### Generic

Provides support for path based authentication to Vault.

> It is recommended to use one of the authentication engine specific configurations above.
This configuration can be used for custom authentication engines, or in the case where an official authentication
engine is not yet supported by the provider

The path-based `authLogin` configuration block accepts the following arguments:

* `path` - (Required) The login path of the auth backend. For example, login with
  approle by setting this path to `auth/approle/login`. Additionally, some mounts use parameters
  in the URL, like with `userpass`: `auth/userpass/login/:username`.

* `namespace` - (Optional) The path to the namespace that has the mounted auth method.
  This defaults to the root namespace. Cannot contain any leading or trailing slashes.
  *Available only for Vault Enterprise*.

* `useRootNamespace` - (Optional) Authenticate to the root Vault namespace. Conflicts with `namespace`.

* `method` - (Optional) When configured, will enable auth method specific operations.
  For example, when set to `aws`, the provider will automatically sign login requests
  for AWS authentication. Valid values include: `aws`.

* `parameters` - (Optional) A map of key-value parameters to send when authenticating
  against the auth backend. Refer to [Vault API documentation](https://www.vaultproject.io/api-docs/auth) for a particular auth method
  to see what can go here.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as vault from "@pulumi/vault";

const example = new vault.generic.Secret("example", {
    path: "secret/foo",
    dataJson: JSON.stringify({
        foo: "bar",
        pizza: "cheese",
    }),
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
import json
import pulumi_vault as vault

example = vault.generic.Secret("example",
    path="secret/foo",
    data_json=json.dumps({
        "foo": "bar",
        "pizza": "cheese",
    }))
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
using System.Text.Json;
using Pulumi;
using Vault = Pulumi.Vault;

return await Deployment.RunAsync(() =>
{
    var example = new Vault.Generic.Secret("example", new()
    {
        Path = "secret/foo",
        DataJson = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["foo"] = "bar",
            ["pizza"] = "cheese",
        }),
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
	"encoding/json"

	"github.com/pulumi/pulumi-vault/sdk/v7/go/vault/generic"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"foo":   "bar",
			"pizza": "cheese",
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = generic.NewSecret(ctx, "example", &generic.SecretArgs{
			Path:     pulumi.String("secret/foo"),
			DataJson: pulumi.String(json0),
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
  example:
    type: vault:generic:Secret
    properties:
      path: secret/foo
      dataJson:
        fn::toJSON:
          foo: bar
          pizza: cheese
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
import com.pulumi.vault.generic.Secret;
import com.pulumi.vault.generic.SecretArgs;
import static com.pulumi.codegen.internal.Serialization.*;
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
        var example = new Secret("example", SecretArgs.builder()
            .path("secret/foo")
            .dataJson(serializeJson(
                jsonObject(
                    jsonProperty("foo", "bar"),
                    jsonProperty("pizza", "cheese")
                )))
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
### Example `authLogin` Usage
With the `userpass` backend:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs

```
```typescript
import * as pulumi from "@pulumi/pulumi";

const config = new pulumi.Config();
const loginUsername = config.requireObject<any>("loginUsername");
const loginPassword = config.requireObject<any>("loginPassword");
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

config = pulumi.Config()
login_username = config.require_object("loginUsername")
login_password = config.require_object("loginPassword")
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

return await Deployment.RunAsync(() =>
{
    var config = new Config();
    var loginUsername = config.RequireObject<dynamic>("loginUsername");
    var loginPassword = config.RequireObject<dynamic>("loginPassword");
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
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		loginUsername := cfg.RequireObject("loginUsername")
		loginPassword := cfg.RequireObject("loginPassword")
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
configuration:
  loginUsername:
    type: dynamic
  loginPassword:
    type: dynamic
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
        final var loginUsername = config.get("loginUsername");
        final var loginPassword = config.get("loginPassword");
    }
}
```
{{% /choosable %}}
{{< /chooser >}}

Or, using `approle`:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs

```
```typescript
import * as pulumi from "@pulumi/pulumi";

const config = new pulumi.Config();
const loginApproleRoleId = config.requireObject<any>("loginApproleRoleId");
const loginApproleSecretId = config.requireObject<any>("loginApproleSecretId");
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

config = pulumi.Config()
login_approle_role_id = config.require_object("loginApproleRoleId")
login_approle_secret_id = config.require_object("loginApproleSecretId")
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

return await Deployment.RunAsync(() =>
{
    var config = new Config();
    var loginApproleRoleId = config.RequireObject<dynamic>("loginApproleRoleId");
    var loginApproleSecretId = config.RequireObject<dynamic>("loginApproleSecretId");
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
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		loginApproleRoleId := cfg.RequireObject("loginApproleRoleId")
		loginApproleSecretId := cfg.RequireObject("loginApproleSecretId")
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
configuration:
  loginApproleRoleId:
    type: dynamic
  loginApproleSecretId:
    type: dynamic
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
        final var loginApproleRoleId = config.get("loginApproleRoleId");
        final var loginApproleSecretId = config.get("loginApproleSecretId");
    }
}
```
{{% /choosable %}}
{{< /chooser >}}
### Example `authLogin` With AWS Signing

Sign AWS metadata for instance profile login requests:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    vault:address:
        value: http://127.0.0.1:8200

```

If the Vault server's AWS auth method requires the `X-Vault-AWS-IAM-Server-ID` header to be set by clients, specify the server ID in `headerValue` within the `parameters` block:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    vault:address:
        value: http://127.0.0.1:8200

```