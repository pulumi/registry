---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-snowflake/v1.3.0/docs/_index.md
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
* Go: [`github.com/pulumi/pulumi-snowflake/sdk/go/snowflake`](https://github.com/pulumi/pulumi-snowflake)
* .NET: [`Pulumi.Snowflake`](https://www.nuget.org/packages/Pulumi.Snowflake)
* Java: [`com.pulumi/snowflake`](https://central.sonatype.com/artifact/com.pulumi/snowflake)
## Overview

This is a pulumi provider plugin for managing [Snowflake](https://www.snowflake.com/) accounts.
Coverage is focused on part of Snowflake related to access control.
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

{{% /choosable %}}
{{< /chooser >}}

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
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

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
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

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
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

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
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

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
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

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
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

{{% /choosable %}}
{{< /chooser >}}

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    snowflake:profile:
        value: securityadmin

```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    snowflake:profile:
        value: securityadmin

```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    snowflake:profile:
        value: securityadmin

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    snowflake:profile:
        value: securityadmin

```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    snowflake:profile:
        value: securityadmin

```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    snowflake:profile:
        value: securityadmin

```

{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

**Warning: these values are passed directly to the gosnowflake library, which may not work exactly the way you expect. See the [gosnowflake docs](https://godoc.org/github.com/snowflakedb/gosnowflake#hdr-Connection_Parameters) for more.**
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
- `passcode` (String) Specifies the passcode provided by Duo when using multi-factor authentication (MFA) for login. Can also be sourced from the `SNOWFLAKE_PASSCODE` environment variable.
- `passcodeInPassword` (Boolean) False by default. Set to true if the MFA passcode is embedded to the configured password. Can also be sourced from the `SNOWFLAKE_PASSCODE_IN_PASSWORD` environment variable.
- `password` (String, Sensitive) Password for user + password auth. Cannot be used with `privateKey` and `privateKeyPassphrase`. Can also be sourced from the `SNOWFLAKE_PASSWORD` environment variable.
- `port` (Number) Specifies a custom port value used by the driver for privatelink connections. Can also be sourced from the `SNOWFLAKE_PORT` environment variable.
- `privateKey` (String, Sensitive) Private Key for username+private-key auth. Cannot be used with `password`. Can also be sourced from the `SNOWFLAKE_PRIVATE_KEY` environment variable.
- `privateKeyPassphrase` (String, Sensitive) Supports the encryption ciphers aes-128-cbc, aes-128-gcm, aes-192-cbc, aes-192-gcm, aes-256-cbc, aes-256-gcm, and des-ede3-cbc. Can also be sourced from the `SNOWFLAKE_PRIVATE_KEY_PASSPHRASE` environment variable.
- `profile` (String) Sets the profile to read from ~/.snowflake/config file. Can also be sourced from the `SNOWFLAKE_PROFILE` environment variable.
- `protocol` (String) A protocol used in the connection. Valid options are: `http` | `https`. Can also be sourced from the `SNOWFLAKE_PROTOCOL` environment variable.
- `requestTimeout` (Number) request retry timeout in seconds EXCLUDING network roundtrip and read out http response. Can also be sourced from the `SNOWFLAKE_REQUEST_TIMEOUT` environment variable.
- `role` (String) Specifies the role to use by default for accessing Snowflake objects in the client session. Can also be sourced from the `SNOWFLAKE_ROLE` environment variable.
- `tmpDirectoryPath` (String) Sets temporary directory used by the driver for operations like encrypting, compressing etc. Can also be sourced from the `SNOWFLAKE_TMP_DIRECTORY_PATH` environment variable.
- `token` (String, Sensitive) Token to use for OAuth and other forms of token based auth. Can also be sourced from the `SNOWFLAKE_TOKEN` environment variable.
- `tokenAccessor` (Block List, Max: 1) (see below for nested schema)
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
export SNOWFLAKE_PRIVATE_KEY="~/.ssh/snowflake_key"
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
export SNOWFLAKE_PRIVATE_KEY="~/.ssh/snowflake_key.p8"
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

Example content of the Pulumi file configuration:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
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

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
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

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
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

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
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

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
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

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
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

{{% /choosable %}}
{{< /chooser >}}

2. In environmental variables (envs). This is mainly used to provide sensitive values.

```bash
export SNOWFLAKE_USER="..."
export SNOWFLAKE_PRIVATE_KEY="~/.ssh/snowflake_key"
```

3. In a TOML file (default in ~/.snowflake/config). Notice the use of different profiles. The profile name needs to be specified in the Pulumi configuration file in `profile` field. When this is not specified, `default` profile is loaded.
   When a `default` profile is not present in the TOML file, it is treated as "empty", without failing.

Example content of the Pulumi file configuration:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    snowflake:profile:
        value: default

```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    snowflake:profile:
        value: default

```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    snowflake:profile:
        value: default

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    snowflake:profile:
        value: default

```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    snowflake:profile:
        value: default

```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    snowflake:profile:
        value: default

```

{{% /choosable %}}
{{< /chooser >}}

Example content of the TOML file configuration:

```toml
[default]
organizationname='organization_name'
accountname='account_name'
user='user'
password='password'
role='ACCOUNTADMIN'

[secondary_test_account]
organizationname='organization_name'
accountname='account2_name'
user='user'
password='password'
role='ACCOUNTADMIN'
```

Not all fields must be configured in one source; users can choose which fields are configured in which source.
Provider uses an established hierarchy of sources. The current behavior is that for each field:
1. Check if it is present in the provider configuration. If yes, use this value. If not, go to step 2.
2. Check if it is present in the environment variables. If yes, use this value. If not, go to step 3.
3. Check if it is present in the TOML config file (specifically, use the profile name configured in one of the steps above). If yes, use this value. If not, the value is considered empty.

An example TOML file contents:

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

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
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

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
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

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
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

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
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

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
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

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
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

{{% /choosable %}}
{{< /chooser >}}

<!-- Section of deprecated resources -->

<!-- Section of deprecated functions -->