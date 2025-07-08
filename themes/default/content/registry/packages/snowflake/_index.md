---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-snowflake/v2.2.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Snowflake Provider
meta_desc: Provides an overview on how to configure the Pulumi Snowflake provider.
layout: package
---

## Installation

The Snowflake provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/snowflake`](https://www.npmjs.com/package/@pulumi/snowflake)
* Python: [`pulumi-snowflake`](https://pypi.org/project/pulumi-snowflake/)
* Go: [`github.com/pulumi/pulumi-snowflake/sdk/v2/go/snowflake`](https://github.com/pulumi/pulumi-snowflake)
* .NET: [`Pulumi.Snowflake`](https://www.nuget.org/packages/Pulumi.Snowflake)
* Java: [`com.pulumi/snowflake`](https://central.sonatype.com/artifact/com.pulumi/snowflake)

## Support

For official support and urgent, production-impacting issues, please [contact Snowflake Support](https://community.snowflake.com/s/article/How-To-Submit-a-Support-Case-in-Snowflake-Lodge).

> **Keep in mind** that the official support starts with the v2.0.0 version for stable resources only. All previous versions and preview resources are not officially supported. Also, consult supported architectures.

Please follow creating issues guidelines, FAQ, and known issues before submitting an issue on GitHub or directly to Snowflake Support.
> **Disclaimer** The project is in GA version, but some features are in preview. Such resources and functions are considered preview features in the provider, regardless of their state in Snowflake. We do not guarantee their stability. They will be reworked and marked as a stable feature in future releases. Breaking changes in these features are expected, even without bumping the major version. They are disabled by default. To use them, add the relevant feature name to `previewFeaturesEnabled` field in the provider configuration. The list of preview features is available below. Please always refer to the Getting Help section in our Github repo to best determine how to get help for your questions.

!> **Sensitive values** Important: Do not include credentials, personal identifiers, or other regulated or sensitive information (e.g., GDPR, HIPAA, PCI-DSS data) in non-sensitive fields. Snowflake marks specific fields as sensitive—such as passwords, private keys, and tokens, meaning these fields will not appear in logs. Each sensitive field is properly marked in the documentation. All other fields are treated as non-sensitive by default. Some of them, like task's configuration, may contain sensitive information but are not marked as sensitive - you are responsible for safeguarding these fields according to your organization's security standards and regulatory requirements. Snowflake will not be liable for any exposure of data placed in non-sensitive fields. Read more in the Sensitive values limitations section.

This is a pulumi provider plugin for managing [Snowflake](https://www.snowflake.com/) accounts.
Coverage is focused on part of Snowflake related to access control.
## Supported architectures

We have compiled a list to clarify which binaries are officially supported and which are provided additionally but not officially supported.
The lists are based on what the underlying [gosnowflake driver](https://github.com/snowflakedb/gosnowflake) supports and what HashiCorp recommends for Pulumi providers.

The provider officially supports the binaries built for the following OSes and architectures:
- Windows: amd64
- Linux: amd64 and arm64
- Darwin: amd64 and arm64

Currently, we also provide the binaries for the following OSes and architectures, but they are not officially supported, and we do not prioritize fixes for them:
- Windows: arm64 and 386
- Linux: 386
- Darwin: 386
- Freebsd: any architecture
## Example Provider Configuration

This is an example configuration of the provider. More examples are provided below.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    snowflake:accountName:
        value: '...'
    snowflake:host:
        value: '...'
    snowflake:organizationName:
        value: '...'
    snowflake:params:
        value:
            query_tag: '...'
    snowflake:password:
        value: '...'
    snowflake:role:
        value: '...'
    snowflake:user:
        value: '...'
    snowflake:warehouse:
        value: '...'

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
    snowflake:accountName:
        value: '...'
    snowflake:host:
        value: '...'
    snowflake:organizationName:
        value: '...'
    snowflake:params:
        value:
            query_tag: '...'
    snowflake:password:
        value: '...'
    snowflake:role:
        value: '...'
    snowflake:user:
        value: '...'
    snowflake:warehouse:
        value: '...'

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
    snowflake:accountName:
        value: '...'
    snowflake:host:
        value: '...'
    snowflake:organizationName:
        value: '...'
    snowflake:params:
        value:
            query_tag: '...'
    snowflake:password:
        value: '...'
    snowflake:role:
        value: '...'
    snowflake:user:
        value: '...'
    snowflake:warehouse:
        value: '...'

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
    snowflake:accountName:
        value: '...'
    snowflake:host:
        value: '...'
    snowflake:organizationName:
        value: '...'
    snowflake:params:
        value:
            query_tag: '...'
    snowflake:password:
        value: '...'
    snowflake:role:
        value: '...'
    snowflake:user:
        value: '...'
    snowflake:warehouse:
        value: '...'

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
    snowflake:accountName:
        value: '...'
    snowflake:host:
        value: '...'
    snowflake:organizationName:
        value: '...'
    snowflake:params:
        value:
            query_tag: '...'
    snowflake:password:
        value: '...'
    snowflake:role:
        value: '...'
    snowflake:user:
        value: '...'
    snowflake:warehouse:
        value: '...'

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
    snowflake:accountName:
        value: '...'
    snowflake:host:
        value: '...'
    snowflake:organizationName:
        value: '...'
    snowflake:params:
        value:
            query_tag: '...'
    snowflake:password:
        value: '...'
    snowflake:role:
        value: '...'
    snowflake:user:
        value: '...'
    snowflake:warehouse:
        value: '...'

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

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    snowflake:accountName:
        value: '...'
    snowflake:authenticator:
        value: SNOWFLAKE_JWT
    snowflake:organizationName:
        value: '...'
    snowflake:privateKey:
        value: '-----BEGIN ENCRYPTED PRIVATE KEY-----...'
    snowflake:privateKeyPassphrase:
        value: passphrase
    snowflake:user:
        value: '...'

```

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    snowflake:profile:
        value: securityadmin

```
## Configuration Reference

**Warning: these values are passed directly to the gosnowflake library, which may not work exactly the way you expect. See the [gosnowflake docs](https://godoc.org/github.com/snowflakedb/gosnowflake#hdr-Connection_Parameters) for more.**

> **Note**: In Go Snowflake driver 1.12.1 ([release notes](https://docs.snowflake.com/en/release-notes/clients-drivers/golang-2024#version-1-12-1-december-05-2024)), configuration field `InsecureMode` has been deprecated in favor of `DisableOCSPChecks`. This field is not available in the provider yet. Please use `InsecureMode` instead, which has the same behavior. We are planning to support this new field and deprecate the old one.

> **Note** If a field has a default value, it is shown next to the type in the schema. Most of the values in provider schema can be sourced from environment value (check field descriptions), but If a specified environment variable is not found, then the driver's default value is used instead.
## Configuration Reference

- `accountName` (String) Specifies your Snowflake account name assigned by Snowflake. For information about account identifiers, see the [Snowflake documentation](https://docs.snowflake.com/en/user-guide/admin-account-identifier#account-name). Required unless using `profile`. Can also be sourced from the `SNOWFLAKE_ACCOUNT_NAME` environment variable.
- `authenticator` (String) Specifies the [authentication type](https://pkg.go.dev/github.com/snowflakedb/gosnowflake#AuthType) to use when connecting to Snowflake. Valid options are: `SNOWFLAKE` | `OAUTH` | `EXTERNALBROWSER` | `OKTA` | `SNOWFLAKE_JWT` | `TOKENACCESSOR` | `USERNAMEPASSWORDMFA`. Can also be sourced from the `SNOWFLAKE_AUTHENTICATOR` environment variable.
- `clientIp` (String) IP address for network checks. Can also be sourced from the `SNOWFLAKE_CLIENT_IP` environment variable.
- `clientRequestMfaToken` (String) When true the MFA token is cached in the credential manager. True by default in Windows/OSX. False for Linux. Can also be sourced from the `SNOWFLAKE_CLIENT_REQUEST_MFA_TOKEN` environment variable.
- `clientStoreTemporaryCredential` (String) When true the ID token is cached in the credential manager. True by default in Windows/OSX. False for Linux. Can also be sourced from the `SNOWFLAKE_CLIENT_STORE_TEMPORARY_CREDENTIAL` environment variable.
- `clientTimeout` (Number) The timeout in seconds for the client to complete the authentication. Can also be sourced from the `SNOWFLAKE_CLIENT_TIMEOUT` environment variable.
- `disableConsoleLogin` (String) Indicates whether console login should be disabled in the driver. Can also be sourced from the `SNOWFLAKE_DISABLE_CONSOLE_LOGIN` environment variable.
- `disableQueryContextCache` (Boolean) Disables HTAP query context cache in the driver. Can also be sourced from the `SNOWFLAKE_DISABLE_QUERY_CONTEXT_CACHE` environment variable.
- `disableTelemetry` (Boolean) Disables telemetry in the driver. Can also be sourced from the `DISABLE_TELEMETRY` environment variable.
- `driverTracing` (String) Specifies the logging level to be used by the driver. Valid options are: `trace` | `debug` | `info` | `print` | `warning` | `error` | `fatal` | `panic`. Can also be sourced from the `SNOWFLAKE_DRIVER_TRACING` environment variable.
- `externalBrowserTimeout` (Number) The timeout in seconds for the external browser to complete the authentication. Can also be sourced from the `SNOWFLAKE_EXTERNAL_BROWSER_TIMEOUT` environment variable.
- `host` (String) Specifies a custom host value used by the driver for privatelink connections. Can also be sourced from the `SNOWFLAKE_HOST` environment variable.
- `includeRetryReason` (String) Should retried request contain retry reason. Can also be sourced from the `SNOWFLAKE_INCLUDE_RETRY_REASON` environment variable.
- `insecureMode` (Boolean) If true, bypass the Online Certificate Status Protocol (OCSP) certificate revocation check. IMPORTANT: Change the default value for testing or emergency situations only. Can also be sourced from the `SNOWFLAKE_INSECURE_MODE` environment variable.
- `jwtClientTimeout` (Number) The timeout in seconds for the JWT client to complete the authentication. Can also be sourced from the `SNOWFLAKE_JWT_CLIENT_TIMEOUT` environment variable.
- `jwtExpireTimeout` (Number) JWT expire after timeout in seconds. Can also be sourced from the `SNOWFLAKE_JWT_EXPIRE_TIMEOUT` environment variable.
- `keepSessionAlive` (Boolean) Enables the session to persist even after the connection is closed. Can also be sourced from the `SNOWFLAKE_KEEP_SESSION_ALIVE` environment variable.
- `loginTimeout` (Number) Login retry timeout in seconds EXCLUDING network roundtrip and read out http response. Can also be sourced from the `SNOWFLAKE_LOGIN_TIMEOUT` environment variable.
- `maxRetryCount` (Number) Specifies how many times non-periodic HTTP request can be retried by the driver. Can also be sourced from the `SNOWFLAKE_MAX_RETRY_COUNT` environment variable.
- `ocspFailOpen` (String) True represents OCSP fail open mode. False represents OCSP fail closed mode. Fail open true by default. Can also be sourced from the `SNOWFLAKE_OCSP_FAIL_OPEN` environment variable.
- `oktaUrl` (String) The URL of the Okta server. e.g. <https://example.okta.com>. Okta URL host needs to to have a suffix `okta.com`. Read more in Snowflake [docs](https://docs.snowflake.com/en/user-guide/oauth-okta). Can also be sourced from the `SNOWFLAKE_OKTA_URL` environment variable.
- `organizationName` (String) Specifies your Snowflake organization name assigned by Snowflake. For information about account identifiers, see the [Snowflake documentation](https://docs.snowflake.com/en/user-guide/admin-account-identifier#organization-name). Required unless using `profile`. Can also be sourced from the `SNOWFLAKE_ORGANIZATION_NAME` environment variable.
- `params` (Map of String) Sets other connection (i.e. session) parameters. [Parameters](https://docs.snowflake.com/en/sql-reference/parameters). This field can not be set with environmental variables.
- `passcode` (String, Sensitive) Specifies the passcode provided by Duo when using multi-factor authentication (MFA) for login. Can also be sourced from the `SNOWFLAKE_PASSCODE` environment variable.
- `passcodeInPassword` (Boolean) False by default. Set to true if the MFA passcode is embedded to the configured password. Can also be sourced from the `SNOWFLAKE_PASSCODE_IN_PASSWORD` environment variable.
- `password` (String, Sensitive) Password for user + password auth. Cannot be used with `privateKey` and `privateKeyPassphrase`. Can also be sourced from the `SNOWFLAKE_PASSWORD` environment variable.
- `port` (Number) Specifies a custom port value used by the driver for privatelink connections. Can also be sourced from the `SNOWFLAKE_PORT` environment variable.
- `previewFeaturesEnabled` (Set of String) A list of preview features that are handled by the provider. See preview features list. Preview features may have breaking changes in future releases, even without raising the major version. This field can not be set with environmental variables. Valid options are: `snowflakeCurrentAccountFunction` | `snowflakeAccountAuthenticationPolicyAttachmentResource` | `snowflakeAccountPasswordPolicyAttachmentResource` | `snowflakeAlertResource` | `snowflakeAlertsFunction` | `snowflakeApiIntegrationResource` | `snowflakeAuthenticationPolicyResource` | `snowflakeCortexSearchServiceResource` | `snowflakeCortexSearchServicesFunction` | `snowflakeDatabaseFunction` | `snowflakeDatabaseRoleFunction` | `snowflakeDynamicTableResource` | `snowflakeDynamicTablesFunction` | `snowflakeExternalFunctionResource` | `snowflakeExternalFunctionsFunction` | `snowflakeExternalTableResource` | `snowflakeExternalTablesFunction` | `snowflakeExternalVolumeResource` | `snowflakeFailoverGroupResource` | `snowflakeFailoverGroupsFunction` | `snowflakeFileFormatResource` | `snowflakeFileFormatsFunction` | `snowflakeFunctionJavaResource` | `snowflakeFunctionJavascriptResource` | `snowflakeFunctionPythonResource` | `snowflakeFunctionScalaResource` | `snowflakeFunctionSqlResource` | `snowflakeFunctionsFunction` | `snowflakeManagedAccountResource` | `snowflakeMaterializedViewResource` | `snowflakeMaterializedViewsFunction` | `snowflakeNetworkPolicyAttachmentResource` | `snowflakeNetworkRuleResource` | `snowflakeEmailNotificationIntegrationResource` | `snowflakeNotificationIntegrationResource` | `snowflakeObjectParameterResource` | `snowflakePasswordPolicyResource` | `snowflakePipeResource` | `snowflakePipesFunction` | `snowflakeCurrentRoleFunction` | `snowflakeSequenceResource` | `snowflakeSequencesFunction` | `snowflakeShareResource` | `snowflakeSharesFunction` | `snowflakeParametersFunction` | `snowflakeProcedureJavaResource` | `snowflakeProcedureJavascriptResource` | `snowflakeProcedurePythonResource` | `snowflakeProcedureScalaResource` | `snowflakeProcedureSqlResource` | `snowflakeProceduresFunction` | `snowflakeStageResource` | `snowflakeStagesFunction` | `snowflakeStorageIntegrationResource` | `snowflakeStorageIntegrationsFunction` | `snowflakeSystemGenerateScimAccessTokenFunction` | `snowflakeSystemGetAwsSnsIamPolicyFunction` | `snowflakeSystemGetPrivatelinkConfigFunction` | `snowflakeSystemGetSnowflakePlatformInfoFunction` | `snowflakeTableColumnMaskingPolicyApplicationResource` | `snowflakeTableConstraintResource` | `snowflakeTableResource` | `snowflakeTablesFunction` | `snowflakeUserAuthenticationPolicyAttachmentResource` | `snowflakeUserPublicKeysResource` | `snowflakeUserPasswordPolicyAttachmentResource`.
- `privateKey` (String, Sensitive) Private Key for username+private-key auth. Cannot be used with `password`. Can also be sourced from the `SNOWFLAKE_PRIVATE_KEY` environment variable.
- `privateKeyPassphrase` (String, Sensitive) Supports the encryption ciphers aes-128-cbc, aes-128-gcm, aes-192-cbc, aes-192-gcm, aes-256-cbc, aes-256-gcm, and des-ede3-cbc. Can also be sourced from the `SNOWFLAKE_PRIVATE_KEY_PASSPHRASE` environment variable.
- `profile` (String) Sets the profile to read from ~/.snowflake/config file. Can also be sourced from the `SNOWFLAKE_PROFILE` environment variable.
- `protocol` (String) A protocol used in the connection. Valid options are: `http` | `https`. Can also be sourced from the `SNOWFLAKE_PROTOCOL` environment variable.
- `requestTimeout` (Number) request retry timeout in seconds EXCLUDING network roundtrip and read out http response. Can also be sourced from the `SNOWFLAKE_REQUEST_TIMEOUT` environment variable.
- `role` (String) Specifies the role to use by default for accessing Snowflake objects in the client session. Can also be sourced from the `SNOWFLAKE_ROLE` environment variable.
- `skipTomlFilePermissionVerification` (Boolean) False by default. Skips TOML configuration file permission verification. This flag has no effect on Windows systems, as the permissions are not checked on this platform. Instead of skipping the permissions verification, we recommend setting the proper privileges - see the section below. Can also be sourced from the `SNOWFLAKE_SKIP_TOML_FILE_PERMISSION_VERIFICATION` environment variable.
- `tmpDirectoryPath` (String) Sets temporary directory used by the driver for operations like encrypting, compressing etc. Can also be sourced from the `SNOWFLAKE_TMP_DIRECTORY_PATH` environment variable.
- `token` (String, Sensitive) Token to use for OAuth and other forms of token based auth. Can also be sourced from the `SNOWFLAKE_TOKEN` environment variable.
- `tokenAccessor` (Block List, Max: 1) (see below for nested schema)
- `useLegacyTomlFile` (Boolean) False by default. When this is set to true, the provider expects the legacy TOML format. Otherwise, it expects the new format. See more in the section below Can also be sourced from the `SNOWFLAKE_USE_LEGACY_TOML_FILE` environment variable.
- `user` (String) Username. Required unless using `profile`. Can also be sourced from the `SNOWFLAKE_USER` environment variable.
- `validateDefaultParameters` (String) True by default. If false, disables the validation checks for Database, Schema, Warehouse and Role at the time a connection is established. Can also be sourced from the `SNOWFLAKE_VALIDATE_DEFAULT_PARAMETERS` environment variable.
- `warehouse` (String) Specifies the virtual warehouse to use by default for queries, loading, etc. in the client session. Can also be sourced from the `SNOWFLAKE_WAREHOUSE` environment variable.

<a id="nestedblock--token_accessor"></a>
### Nested Schema for `tokenAccessor`

Required:

- `clientId` (String, Sensitive) The client ID for the OAuth provider when using a refresh token to renew access token. Can also be sourced from the `SNOWFLAKE_TOKEN_ACCESSOR_CLIENT_ID` environment variable.
- `clientSecret` (String, Sensitive) The client secret for the OAuth provider when using a refresh token to renew access token. Can also be sourced from the `SNOWFLAKE_TOKEN_ACCESSOR_CLIENT_SECRET` environment variable.
- `redirectUri` (String, Sensitive) The redirect URI for the OAuth provider when using a refresh token to renew access token. Can also be sourced from the `SNOWFLAKE_TOKEN_ACCESSOR_REDIRECT_URI` environment variable.
- `refreshToken` (String, Sensitive) The refresh token for the OAuth provider when using a refresh token to renew access token. Can also be sourced from the `SNOWFLAKE_TOKEN_ACCESSOR_REFRESH_TOKEN` environment variable.
- `tokenEndpoint` (String, Sensitive) The token endpoint for the OAuth provider e.g. https://{yourDomain}/oauth/token when using a refresh token to renew access token. Can also be sourced from the `SNOWFLAKE_TOKEN_ACCESSOR_TOKEN_ENDPOINT` environment variable.
## Authentication

The Snowflake provider support multiple ways to authenticate:

* Password
* OAuth Access Token
* OAuth Refresh Token
* Browser Auth
* Private Key
* Config File

In all cases `organizationName`, `accountName` and `user` are required.

> **Note** Storing the credentials and other secret values safely is on the users' side. Read more in Authentication Methods guide.
### Keypair Authentication Environment Variables

You should generate the public and private keys and set up environment variables.

```shell

cd ~/.ssh
openssl genrsa -out snowflake_key 4096
openssl rsa -in snowflake_key -pubout -out snowflake_key.pub
```

To export the variables into your provider:

```shell
export SNOWFLAKE_USER="..."
export SNOWFLAKE_PRIVATE_KEY=$(cat ~/.ssh/snowflake_key.p8)
```
### Keypair Authentication Passphrase

If your private key requires a passphrase, then this can be supplied via the
environment variable `SNOWFLAKE_PRIVATE_KEY_PASSPHRASE`.

Only the ciphers aes-128-cbc, aes-128-gcm, aes-192-cbc, aes-192-gcm,
aes-256-cbc, aes-256-gcm, and des-ede3-cbc are supported on the private key

```shell
cd ~/.ssh
openssl genrsa -out snowflake_key 4096
openssl rsa -in snowflake_key -pubout -out snowflake_key.pub
openssl pkcs8 -topk8 -inform pem -in snowflake_key -outform PEM -v2 aes-256-cbc -out snowflake_key.p8
```

To export the variables into your provider:

```shell
export SNOWFLAKE_USER="..."
export SNOWFLAKE_PRIVATE_KEY=$(cat ~/.ssh/snowflake_key.p8)
export SNOWFLAKE_PRIVATE_KEY_PASSPHRASE="..."
```
### OAuth Access Token

If you have an OAuth access token, export these credentials as environment variables:

```shell
export SNOWFLAKE_USER='...'
export SNOWFLAKE_TOKEN='...'
```

Note that once this access token expires, you'll need to request a new one through an external application.
### OAuth Refresh Token

If you have an OAuth Refresh token, export these credentials as environment variables:

```shell
export SNOWFLAKE_TOKEN_ACCESSOR_REFRESH_TOKEN='...'
export SNOWFLAKE_TOKEN_ACCESSOR_CLIENT_ID='...'
export SNOWFLAKE_TOKEN_ACCESSOR_CLIENT_SECRET='...'
export SNOWFLAKE_TOKEN_ACCESSOR_TOKEN_ENDPOINT='...'
export SNOWFLAKE_TOKEN_ACCESSOR_REDIRECT_URI='https://localhost.com'
```

Note because access token have a short life; typically 10 minutes, by passing refresh token new access token will be generated.
### Username and Password Environment Variables

If you choose to use Username and Password Authentication, export these credentials:

```shell
export SNOWFLAKE_USER='...'
export SNOWFLAKE_PASSWORD='...'
```
## Order Precedence

Currently, the provider can be configured in three ways:
1. In a Pulumi file located in the Pulumi module with other resources.
2. In environmental variables (envs). This is mainly used to provide sensitive values.
3. In a TOML file (default in ~/.snowflake/config).
### Pulumi file located in the Pulumi module with other resources
One of the methods of configuring the provider is in the Pulumi module. Read more in the Pulumi docs.

Example content of the Pulumi file configuration:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    snowflake:accountName:
        value: '...'
    snowflake:organizationName:
        value: '...'
    snowflake:password:
        value: '...'
    snowflake:username:
        value: '...'

```
### Environmental variables
The second method is to use environmental variables. This is mainly used to provide sensitive values.

```bash
export SNOWFLAKE_USER="..."
export SNOWFLAKE_PRIVATE_KEY=$(cat ~/.ssh/snowflake_key.p8)
```
### TOML file
The third method is to use a TOML configuration file (default location in ~/.snowflake/config). Notice the use of different profiles. The profile name needs to be specified in the Pulumi configuration file in `profile` field. When this is not specified, `default` profile is loaded.
When a `default` profile is not present in the TOML file, it is treated as "empty", without failing.

Read [TOML](https://toml.io/en/) specification for more details on the syntax.

Example content of the Pulumi file configuration:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    snowflake:profile:
        value: default

```

Example content of the TOML file configuration is listed below. Note that this example follows a new TOML format, for the legacy format see examples section.

```toml
[default]
organization_name='organization_name'
account_name='account_name'
user='user'
password='password'
role='ACCOUNTADMIN'

[secondary_test_account]
organization_name='organization_name'
account_name='account2_name'
user='user'
password='password'
role='ACCOUNTADMIN'
```
#### TOML file limitations
To ensure a better security of the provider, the following limitations are introduced:

> **Note**: TOML file size is limited to 10MB.

> **Note**: Only TOML file with restricted privileges can be read. Any privileges for group or others cannot be set (the maximum valid privilege is `700`). You can set the expected privileges like `chmod 0600 ~/.snowflake/config`. This is checked only on non-Windows platforms. If you are using the provider on Windows, please make sure that your configuration file has not too permissive privileges.
### Source Hierarchy
Not all fields must be configured in one source; users can choose which fields are configured in which source.
Provider uses an established hierarchy of sources. The current behavior is that for each field:
1. Check if it is present in the provider configuration. If yes, use this value. If not, go to step 2.
2. Check if it is present in the environment variables. If yes, use this value. If not, go to step 3.
3. Check if it is present in the TOML config file (specifically, use the profile name configured in one of the steps above). If yes, use this value. If not, the value is considered empty.

> **Note** Currently `privateKey` and `privateKeyPassphrase` are coupled and must be set in one source (both on Pulumi side or both in TOML config, see <https://github.com/snowflakedb/pulumi-provider-snowflake/issues/3332)>. This will be fixed in the future.

> **Note** Currently both legacy and new formats are supported. The new format can be enabled with setting `useLegacyTomlFile = false` in the provider configuration. We encourage using the new format for now, as it will be a default one in v2 version of the provider. The differences between these formats are:
- The keys in the provider contain an underscore (`_`) as a separator, but the TOML schema has fields without any separator.
- The field `driverTracing` in the provider is related to `tracing` in the TOML schema.
### Examples

An example new TOML file contents:

```toml
[example]
account_name = 'account_name'
organization_name = 'organization_name'
user = 'user'
password = 'password'
warehouse = 'SNOWFLAKE'
role = 'ACCOUNTADMIN'
client_ip = '1.2.3.4'
protocol = 'https'
port = 443
okta_url = 'https://example.com'
client_timeout = 10
jwt_client_timeout = 20
login_timeout = 30
request_timeout = 40
jwt_expire_timeout = 50
external_browser_timeout = 60
max_retry_count = 1
authenticator = 'snowflake'
insecure_mode = true
ocsp_fail_open = true
keep_session_alive = true
disable_telemetry = true
validate_default_parameters = true
client_request_mfa_token = true
client_store_temporary_credential = true
driver_tracing = 'info'
tmp_dir_path = '/tmp/pulumi-provider/'
disable_query_context_cache = true
include_retry_reason = true
disable_console_login = true

[example.params]
param_key = 'param_value'
```

An example legacy TOML file contents:

```toml
[example]
accountname = 'account_name'
organizationname = 'organization_name'
user = 'user'
password = 'password'
warehouse = 'SNOWFLAKE'
role = 'ACCOUNTADMIN'
clientip = '1.2.3.4'
protocol = 'https'
port = 443
oktaurl = 'https://example.com'
clienttimeout = 10
jwtclienttimeout = 20
logintimeout = 30
requesttimeout = 40
jwtexpiretimeout = 50
externalbrowsertimeout = 60
maxretrycount = 1
authenticator = 'snowflake'
insecuremode = true
ocspfailopen = true
keepsessionalive = true
disabletelemetry = true
validatedefaultparameters = true
clientrequestmfatoken = true
clientstoretemporarycredential = true
tracing = 'info'
tmpdirpath = '/tmp/pulumi-provider/'
disablequerycontextcache = true
includeretryreason = true
disableconsolelogin = true

[example.params]
param_key = 'param_value'
```

An example pulumi configuration file equivalent:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    snowflake:accountName:
        value: account_name
    snowflake:authenticator:
        value: snowflake
    snowflake:clientIp:
        value: 1.2.3.4
    snowflake:clientRequestMfaToken:
        value: true
    snowflake:clientStoreTemporaryCredential:
        value: true
    snowflake:clientTimeout:
        value: 40
    snowflake:disableConsoleLogin:
        value: true
    snowflake:disableQueryContextCache:
        value: true
    snowflake:disableTelemetry:
        value: true
    snowflake:driverTracing:
        value: info
    snowflake:externalBrowserTimeout:
        value: 60
    snowflake:includeRetryReason:
        value: true
    snowflake:insecureMode:
        value: true
    snowflake:jwtClientTimeout:
        value: 50
    snowflake:jwtExpireTimeout:
        value: 30
    snowflake:keepSessionAlive:
        value: true
    snowflake:loginTimeout:
        value: 10
    snowflake:maxRetryCount:
        value: 3
    snowflake:ocspFailOpen:
        value: true
    snowflake:oktaUrl:
        value: https://example.com
    snowflake:organizationName:
        value: organization_name
    snowflake:params:
        value:
            param_key: param_value
    snowflake:password:
        value: password
    snowflake:port:
        value: "443"
    snowflake:protocol:
        value: https
    snowflake:requestTimeout:
        value: 20
    snowflake:role:
        value: ACCOUNTADMIN
    snowflake:tmpDirectoryPath:
        value: /tmp/pulumi-provider/
    snowflake:user:
        value: user
    snowflake:validateDefaultParameters:
        value: true
    snowflake:warehouse:
        value: SNOWFLAKE

```

<!-- Section of deprecated resources -->

<!-- Section of deprecated functions -->
## Sensitive values limitations

The provider marks fields containing access credentials and other such information as sensitive. This means that the values of these fields will not be logged.

There are some limitations to this mechanism:
- Sensitive values are stored as plaintext in the state file. This is a limitation of Pulumi itself (reference). You should take care to secure access to the state file.
- In Plugin SDK there is no possibility to mark sensitive values conditionally (reference). This means it is not possible to mark sensitive values based on other fields, like marking `body` based on the value of `secure` field in views, functions, and procedures. As a result, this field is not marked as sensitive. For such cases, we add disclaimers in the resource documentation.
- In Plugin SDK, there is no possibility to mark sensitive values in nested fields (reference). This means the nested fields, like these in `showOutput` and `describeOutput` cannot be sensitive.
  As a result, such nested fields are not marked as sensitive. For such cases, we add disclaimers in the resource documentation. Additionally, some fields are missing from `showOutput` and `describeOutput`. However, these fields are present in the resource's root, so they can still be referenced.
  The alternative solution we considered was setting the whole `showOutput` and `describeOutput` as sensitive. However, this solution could reduce the provider functionality and would require changes in user's configurations.

As a general rule, please ensure that no personal data, sensitive data, export-controlled data, or other regulated data is entered as metadata when using the provider. If you use one of these fields, they may be present in logs, so ensure that the provider logs are properly restricted. For more information, see Sensitive values limitations and [Metadata fields in Snowflake](https://docs.snowflake.com/en/sql-reference/metadata).

Read more about sensitive values in the Pulumi documentation.

We are planning to research migration to Plugin Framework and we will investigate if the limitations coming from Plugin SDK can be addressed.
## Features
### Operation Timeouts
By default, Pulumi sets resource operation timeouts to 20 minutes (reference). Now, the provider enables configuration of these values by users in `timeouts` block in each resource.
The default timeouts are in general aligned with the Pulumi defaults. If a resource has different timeouts, it is specified in the resource documentation.

Data sources will be supported in the future.
Read more in following official documentation).

You can specify the timeouts like the following:
{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as snowflake from "@pulumi/snowflake";

const test = new snowflake.Execute("test", {
    execute: "CREATE DATABASE ABC",
    revert: "DROP DATABASE ABC",
    query: "SHOW DATABASES LIKE '%ABC%'",
});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_snowflake as snowflake

test = snowflake.Execute("test",
    execute="CREATE DATABASE ABC",
    revert="DROP DATABASE ABC",
    query="SHOW DATABASES LIKE '%ABC%'")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Snowflake = Pulumi.Snowflake;

return await Deployment.RunAsync(() =>
{
    var test = new Snowflake.Execute("test", new()
    {
        ExecuteSQL = "CREATE DATABASE ABC",
        Revert = "DROP DATABASE ABC",
        Query = "SHOW DATABASES LIKE '%ABC%'",
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-snowflake/sdk/v2/go/snowflake"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := snowflake.NewExecute(ctx, "test", &snowflake.ExecuteArgs{
			Execute: pulumi.String("CREATE DATABASE ABC"),
			Revert:  pulumi.String("DROP DATABASE ABC"),
			Query:   pulumi.String("SHOW DATABASES LIKE '%ABC%'"),
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
  test:
    type: snowflake:Execute
    properties:
      execute: CREATE DATABASE ABC
      revert: DROP DATABASE ABC
      query: SHOW DATABASES LIKE '%ABC%'
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.snowflake.Execute;
import com.pulumi.snowflake.ExecuteArgs;
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
        var test = new Execute("test", ExecuteArgs.builder()
            .execute("CREATE DATABASE ABC")
            .revert("DROP DATABASE ABC")
            .query("SHOW DATABASES LIKE '%ABC%'")
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

> Note: Timeouts can be also set at driver's level (see [driver documentation](https://pkg.go.dev/github.com/snowflakedb/gosnowflake)). These timeouts are independent. We recommend tweaking the timeouts on Pulumi level first.
## A list of preview and stable resources and functions

The provider supports a number of Snowflake features. Within the provider, some features are stable, while others are in preview
(stability of the feature in the provider is not connected to the stability of the feature in Snowflake).

Preview features are **experimental** and may introduce **breaking changes**, even between non-major versions of the provider.
Eventually, every preview resource will be promoted to stable, but the timeline for each feature is not defined (you can find more details on the current/future plans in our roadmap).
New resources will be introduced as preview ones and promoted over time to stable as we gain more confidence in their stability.

Preview features are disabled by default and should be used with caution.
To use them, add the relevant feature name to the `previewFeaturesEnabled` field in the provider configuration.

<!-- Section of stable resources -->
## Currently stable resources

- snowflake.Account
- snowflake.AccountParameter
- snowflake.AccountRole
- snowflake.ApiAuthenticationIntegrationWithAuthorizationCodeGrant
- snowflake.ApiAuthenticationIntegrationWithClientCredentials
- snowflake.ApiAuthenticationIntegrationWithJwtBearer
- snowflake.Database
- snowflake.DatabaseRole
- snowflake.Execute
- snowflake.ExternalOauthIntegration
- snowflake.GrantAccountRole
- snowflake.GrantApplicationRole
- snowflake.GrantDatabaseRole
- snowflake.GrantOwnership
- snowflake.GrantPrivilegesToAccountRole
- snowflake.GrantPrivilegesToDatabaseRole
- snowflake.GrantPrivilegesToShare
- snowflake.LegacyServiceUser
- snowflake.MaskingPolicy
- snowflake.NetworkPolicy
- snowflake.OauthIntegrationForCustomClients
- snowflake.OauthIntegrationForPartnerApplications
- snowflake.PrimaryConnection
- snowflake.ResourceMonitor
- snowflake.RowAccessPolicy
- snowflake.Saml2Integration
- snowflake.Schema
- snowflake.ScimIntegration
- snowflake.SecondaryConnection
- snowflake.SecondaryDatabase
- snowflake.SecretWithAuthorizationCodeGrant
- snowflake.SecretWithBasicAuthentication
- snowflake.SecretWithClientCredentials
- snowflake.SecretWithGenericString
- snowflake.ServiceUser
- snowflake.SharedDatabase
- snowflake.StreamOnDirectoryTable
- snowflake.StreamOnExternalTable
- snowflake.StreamOnTable
- snowflake.StreamOnView
- snowflake.Streamlit
- snowflake.Tag
- snowflake.TagAssociation
- snowflake.Task
- snowflake.User
- snowflake.View
- snowflake.Warehouse

<!-- Section of stable functions -->
## Currently stable functions

- snowflake.getAccountRoles
- snowflake.getAccounts
- snowflake.getConnections
- snowflake.getDatabaseRoles
- snowflake.getDatabases
- snowflake.getGrants
- snowflake.getMaskingPolicies
- snowflake.getNetworkPolicies
- snowflake.getResourceMonitors
- snowflake.getRowAccessPolicies
- snowflake.getSchemas
- snowflake.getSecrets
- snowflake.getSecurityIntegrations
- snowflake.getStreamlits
- snowflake.getStreams
- snowflake.getTags
- snowflake.getTasks
- snowflake.getUsers
- snowflake.getViews
- snowflake.getWarehouses

<!-- Section of preview resources -->
## Currently preview resources

- snowflake.AccountAuthenticationPolicyAttachment
- snowflake.AccountPasswordPolicyAttachment
- snowflake.Alert
- snowflake.ApiIntegration
- snowflake.AuthenticationPolicy
- snowflake.CortexSearchService
- snowflake.DynamicTable
- snowflake.EmailNotificationIntegration
- snowflake.ExternalFunction
- snowflake.ExternalTable
- snowflake.ExternalVolume
- snowflake.FailoverGroup
- snowflake.FileFormat
- snowflake.FunctionJava
- snowflake.FunctionJavascript
- snowflake.FunctionPython
- snowflake.FunctionScala
- snowflake.FunctionSql
- snowflake.ManagedAccount
- snowflake.MaterializedView
- snowflake.NetworkPolicyAttachment
- snowflake.NetworkRule
- snowflake.NotificationIntegration
- snowflake.ObjectParameter
- snowflake.PasswordPolicy
- snowflake.Pipe
- snowflake.ProcedureJava
- snowflake.ProcedureJavascript
- snowflake.ProcedurePython
- snowflake.ProcedureScala
- snowflake.ProcedureSql
- snowflake.Sequence
- snowflake.Share
- snowflake.Stage
- snowflake.StorageIntegration
- snowflake.Table
- snowflake.TableColumnMaskingPolicyApplication
- snowflake.TableConstraint
- snowflake.UserAuthenticationPolicyAttachment
- snowflake.UserPasswordPolicyAttachment
- snowflake.UserPublicKeys

<!-- Section of preview functions -->
## Currently preview functions

- snowflake.getAlerts
- snowflake.getCortexSearchServices
- snowflake.getCurrentAccount
- snowflake.getCurrentRole
- snowflake.Database
- snowflake.DatabaseRole
- snowflake.getDynamicTables
- snowflake.getExternalFunctions
- snowflake.getExternalTables
- snowflake.getFailoverGroups
- snowflake.getFileFormats
- snowflake.getFunctions
- snowflake.getMaterializedViews
- snowflake.getParameters
- snowflake.getPipes
- snowflake.getProcedures
- snowflake.getSequences
- snowflake.getShares
- snowflake.getStages
- snowflake.getStorageIntegrations
- snowflake.getSystemGenerateScimAccessToken
- snowflake.getSystemGetAwsSnsIamPolicy
- snowflake.getSystemGetPrivateLinkConfig
- snowflake.getSystemGetSnowflakePlatformInfo
- snowflake.getTables