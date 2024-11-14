---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-snowflake/v0.61.1/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Snowflake Provider
meta_desc: Provides an overview on how to configure the Pulumi Snowflake provider.
layout: package
---
## Installation

The snowflake provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/snowflake`](https://www.npmjs.com/package/@pulumi/snowflake)
* Python: [`pulumi-snowflake`](https://pypi.org/project/pulumi-snowflake/)
* Go: [`github.com/pulumi/pulumi-snowflake/sdk/go/snowflake`](https://github.com/pulumi/pulumi-snowflake)
* .NET: [`Pulumi.Snowflake`](https://www.nuget.org/packages/Pulumi.Snowflake)
* Java: [`com.pulumi/snowflake`](https://central.sonatype.com/artifact/com.pulumi/snowflake)
## Overview

This is a pulumi provider plugin for managing [Snowflake](https://www.snowflake.com/) accounts.
Coverage is focused on part of Snowflake related to access control.
## Example Provider Configuration

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    snowflake:account:
        value: '...'
    snowflake:authenticator:
        value: '...'
    snowflake:host:
        value: '...'
    snowflake:oauthAccessToken:
        value: '...'
    snowflake:oauthClientId:
        value: '...'
    snowflake:oauthClientSecret:
        value: '...'
    snowflake:oauthEndpoint:
        value: '...'
    snowflake:oauthRedirectUrl:
        value: '...'
    snowflake:oauthRefreshToken:
        value: '...'
    snowflake:password:
        value: '...'
    snowflake:privateKey:
        value: '...'
    snowflake:privateKeyPassphrase:
        value: '...'
    snowflake:privateKeyPath:
        value: '...'
    snowflake:region:
        value: '...'
    snowflake:role:
        value: '...'
    snowflake:sessionParams:
        value:
            query_tag: '...'
    snowflake:username:
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
    snowflake:account:
        value: '...'
    snowflake:authenticator:
        value: '...'
    snowflake:host:
        value: '...'
    snowflake:oauthAccessToken:
        value: '...'
    snowflake:oauthClientId:
        value: '...'
    snowflake:oauthClientSecret:
        value: '...'
    snowflake:oauthEndpoint:
        value: '...'
    snowflake:oauthRedirectUrl:
        value: '...'
    snowflake:oauthRefreshToken:
        value: '...'
    snowflake:password:
        value: '...'
    snowflake:privateKey:
        value: '...'
    snowflake:privateKeyPassphrase:
        value: '...'
    snowflake:privateKeyPath:
        value: '...'
    snowflake:region:
        value: '...'
    snowflake:role:
        value: '...'
    snowflake:sessionParams:
        value:
            query_tag: '...'
    snowflake:username:
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
    snowflake:account:
        value: '...'
    snowflake:authenticator:
        value: '...'
    snowflake:host:
        value: '...'
    snowflake:oauthAccessToken:
        value: '...'
    snowflake:oauthClientId:
        value: '...'
    snowflake:oauthClientSecret:
        value: '...'
    snowflake:oauthEndpoint:
        value: '...'
    snowflake:oauthRedirectUrl:
        value: '...'
    snowflake:oauthRefreshToken:
        value: '...'
    snowflake:password:
        value: '...'
    snowflake:privateKey:
        value: '...'
    snowflake:privateKeyPassphrase:
        value: '...'
    snowflake:privateKeyPath:
        value: '...'
    snowflake:region:
        value: '...'
    snowflake:role:
        value: '...'
    snowflake:sessionParams:
        value:
            query_tag: '...'
    snowflake:username:
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
    snowflake:account:
        value: '...'
    snowflake:authenticator:
        value: '...'
    snowflake:host:
        value: '...'
    snowflake:oauthAccessToken:
        value: '...'
    snowflake:oauthClientId:
        value: '...'
    snowflake:oauthClientSecret:
        value: '...'
    snowflake:oauthEndpoint:
        value: '...'
    snowflake:oauthRedirectUrl:
        value: '...'
    snowflake:oauthRefreshToken:
        value: '...'
    snowflake:password:
        value: '...'
    snowflake:privateKey:
        value: '...'
    snowflake:privateKeyPassphrase:
        value: '...'
    snowflake:privateKeyPath:
        value: '...'
    snowflake:region:
        value: '...'
    snowflake:role:
        value: '...'
    snowflake:sessionParams:
        value:
            query_tag: '...'
    snowflake:username:
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
    snowflake:account:
        value: '...'
    snowflake:authenticator:
        value: '...'
    snowflake:host:
        value: '...'
    snowflake:oauthAccessToken:
        value: '...'
    snowflake:oauthClientId:
        value: '...'
    snowflake:oauthClientSecret:
        value: '...'
    snowflake:oauthEndpoint:
        value: '...'
    snowflake:oauthRedirectUrl:
        value: '...'
    snowflake:oauthRefreshToken:
        value: '...'
    snowflake:password:
        value: '...'
    snowflake:privateKey:
        value: '...'
    snowflake:privateKeyPassphrase:
        value: '...'
    snowflake:privateKeyPath:
        value: '...'
    snowflake:region:
        value: '...'
    snowflake:role:
        value: '...'
    snowflake:sessionParams:
        value:
            query_tag: '...'
    snowflake:username:
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
    snowflake:account:
        value: '...'
    snowflake:authenticator:
        value: '...'
    snowflake:host:
        value: '...'
    snowflake:oauthAccessToken:
        value: '...'
    snowflake:oauthClientId:
        value: '...'
    snowflake:oauthClientSecret:
        value: '...'
    snowflake:oauthEndpoint:
        value: '...'
    snowflake:oauthRedirectUrl:
        value: '...'
    snowflake:oauthRefreshToken:
        value: '...'
    snowflake:password:
        value: '...'
    snowflake:privateKey:
        value: '...'
    snowflake:privateKeyPassphrase:
        value: '...'
    snowflake:privateKeyPath:
        value: '...'
    snowflake:region:
        value: '...'
    snowflake:role:
        value: '...'
    snowflake:sessionParams:
        value:
            query_tag: '...'
    snowflake:username:
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

- `account` (String) Specifies your Snowflake account identifier assigned, by Snowflake. For information about account identifiers, see the [Snowflake documentation](https://docs.snowflake.com/en/user-guide/admin-account-identifier.html). Can also be sourced from the `SNOWFLAKE_ACCOUNT` environment variable. Required unless using `profile`.
- `authenticator` (String) Specifies the [authentication type](https://pkg.go.dev/github.com/snowflakedb/gosnowflake#AuthType) to use when connecting to Snowflake. Valid values include: Snowflake, OAuth, ExternalBrowser, Okta, JWT, TokenAccessor, UsernamePasswordMFA. Can also be sourced from the `SNOWFLAKE_AUTHENTICATOR` environment variable. It has to be set explicitly to JWT for private key authentication.
- `browserAuth` (Boolean, Deprecated) Required when `oauthRefreshToken` is used. Can also be sourced from `SNOWFLAKE_USE_BROWSER_AUTH` environment variable.
- `clientIp` (String) IP address for network checks. Can also be sourced from the `SNOWFLAKE_CLIENT_IP` environment variable.
- `clientRequestMfaToken` (Boolean) When true the MFA token is cached in the credential manager. True by default in Windows/OSX. False for Linux. Can also be sourced from the `SNOWFLAKE_CLIENT_REQUEST_MFA_TOKEN` environment variable.
- `clientStoreTemporaryCredential` (Boolean) When true the ID token is cached in the credential manager. True by default in Windows/OSX. False for Linux. Can also be sourced from the `SNOWFLAKE_CLIENT_STORE_TEMPORARY_CREDENTIAL` environment variable.
- `clientTimeout` (Number) The timeout in seconds for the client to complete the authentication. Default is 900 seconds. Can also be sourced from the `SNOWFLAKE_CLIENT_TIMEOUT` environment variable.
- `disableQueryContextCache` (Boolean) Should HTAP query context cache be disabled. Can also be sourced from the `SNOWFLAKE_DISABLE_QUERY_CONTEXT_CACHE` environment variable.
- `disableTelemetry` (Boolean) Indicates whether to disable telemetry. Can also be sourced from the `SNOWFLAKE_DISABLE_TELEMETRY` environment variable.
- `externalBrowserTimeout` (Number) The timeout in seconds for the external browser to complete the authentication. Default is 120 seconds. Can also be sourced from the `SNOWFLAKE_EXTERNAL_BROWSER_TIMEOUT` environment variable.
- `host` (String) Supports passing in a custom host value to the snowflake go driver for use with privatelink. Can also be sourced from the `SNOWFLAKE_HOST` environment variable.
- `insecureMode` (Boolean) If true, bypass the Online Certificate Status Protocol (OCSP) certificate revocation check. IMPORTANT: Change the default value for testing or emergency situations only. Can also be sourced from the `SNOWFLAKE_INSECURE_MODE` environment variable.
- `jwtClientTimeout` (Number) The timeout in seconds for the JWT client to complete the authentication. Default is 10 seconds. Can also be sourced from the `SNOWFLAKE_JWT_CLIENT_TIMEOUT` environment variable.
- `jwtExpireTimeout` (Number) JWT expire after timeout in seconds. Can also be sourced from the `SNOWFLAKE_JWT_EXPIRE_TIMEOUT` environment variable.
- `keepSessionAlive` (Boolean) Enables the session to persist even after the connection is closed. Can also be sourced from the `SNOWFLAKE_KEEP_SESSION_ALIVE` environment variable.
- `loginTimeout` (Number) Login retry timeout EXCLUDING network roundtrip and read out http response. Can also be sourced from the `SNOWFLAKE_LOGIN_TIMEOUT` environment variable.
- `oauthAccessToken` (String, Sensitive, Deprecated) Token for use with OAuth. Generating the token is left to other tools. Cannot be used with `browserAuth`, `privateKeyPath`, `oauthRefreshToken` or `password`. Can also be sourced from `SNOWFLAKE_OAUTH_ACCESS_TOKEN` environment variable.
- `oauthClientId` (String, Sensitive, Deprecated) Required when `oauthRefreshToken` is used. Can also be sourced from `SNOWFLAKE_OAUTH_CLIENT_ID` environment variable.
- `oauthClientSecret` (String, Sensitive, Deprecated) Required when `oauthRefreshToken` is used. Can also be sourced from `SNOWFLAKE_OAUTH_CLIENT_SECRET` environment variable.
- `oauthEndpoint` (String, Sensitive, Deprecated) Required when `oauthRefreshToken` is used. Can also be sourced from `SNOWFLAKE_OAUTH_ENDPOINT` environment variable.
- `oauthRedirectUrl` (String, Sensitive, Deprecated) Required when `oauthRefreshToken` is used. Can also be sourced from `SNOWFLAKE_OAUTH_REDIRECT_URL` environment variable.
- `oauthRefreshToken` (String, Sensitive, Deprecated) Token for use with OAuth. Setup and generation of the token is left to other tools. Should be used in conjunction with `oauthClientId`, `oauthClientSecret`, `oauthEndpoint`, `oauthRedirectUrl`. Cannot be used with `browserAuth`, `privateKeyPath`, `oauthAccessToken` or `password`. Can also be sourced from `SNOWFLAKE_OAUTH_REFRESH_TOKEN` environment variable.
- `ocspFailOpen` (Boolean) True represents OCSP fail open mode. False represents OCSP fail closed mode. Fail open true by default. Can also be sourced from the `SNOWFLAKE_OCSP_FAIL_OPEN` environment variable.
- `oktaUrl` (String) The URL of the Okta server. e.g. <https://example.okta.com>. Can also be sourced from the `SNOWFLAKE_OKTA_URL` environment variable.
- `params` (Map of String) Sets other connection (i.e. session) parameters. [Parameters](https://docs.snowflake.com/en/sql-reference/parameters)
- `passcode` (String) Specifies the passcode provided by Duo when using multi-factor authentication (MFA) for login. Can also be sourced from the `SNOWFLAKE_PASSCODE` environment variable.
- `passcodeInPassword` (Boolean) False by default. Set to true if the MFA passcode is embedded in the login password. Appends the MFA passcode to the end of the password. Can also be sourced from the `SNOWFLAKE_PASSCODE_IN_PASSWORD` environment variable.
- `password` (String, Sensitive) Password for username+password auth. Cannot be used with `browserAuth` or `privateKeyPath`. Can also be sourced from the `SNOWFLAKE_PASSWORD` environment variable.
- `port` (Number) Support custom port values to snowflake go driver for use with privatelink. Can also be sourced from the `SNOWFLAKE_PORT` environment variable.
- `privateKey` (String, Sensitive) Private Key for username+private-key auth. Cannot be used with `browserAuth` or `password`. Can also be sourced from `SNOWFLAKE_PRIVATE_KEY` environment variable.
- `privateKeyPassphrase` (String, Sensitive) Supports the encryption ciphers aes-128-cbc, aes-128-gcm, aes-192-cbc, aes-192-gcm, aes-256-cbc, aes-256-gcm, and des-ede3-cbc. Can also be sourced from `SNOWFLAKE_PRIVATE_KEY_PASSPHRASE` environment variable.
- `privateKeyPath` (String, Sensitive, Deprecated) Path to a private key for using keypair authentication. Cannot be used with `browserAuth`, `oauthAccessToken` or `password`. Can also be sourced from `SNOWFLAKE_PRIVATE_KEY_PATH` environment variable.
- `profile` (String) Sets the profile to read from ~/.snowflake/config file. Can also be sourced from the `SNOWFLAKE_PROFILE` environment variable.
- `protocol` (String) Either http or https, defaults to https. Can also be sourced from the `SNOWFLAKE_PROTOCOL` environment variable.
- `region` (String, Deprecated) Snowflake region, such as "eu-central-1", with this parameter. However, since this parameter is deprecated, it is best to specify the region as part of the account parameter. For details, see the description of the account parameter. [Snowflake region](https://docs.snowflake.com/en/user-guide/intro-regions.html) to use.  Required if using the [legacy format for the `account` identifier](https://docs.snowflake.com/en/user-guide/admin-account-identifier.html#format-2-legacy-account-locator-in-a-region) in the form of `<cloud_region_id>.<cloud>`. Can also be sourced from the `SNOWFLAKE_REGION` environment variable.
- `requestTimeout` (Number) request retry timeout EXCLUDING network roundtrip and read out http response. Can also be sourced from the `SNOWFLAKE_REQUEST_TIMEOUT` environment variable.
- `role` (String) Specifies the role to use by default for accessing Snowflake objects in the client session. Can also be sourced from the `SNOWFLAKE_ROLE` environment variable. .
- `sessionParams` (Map of String, Deprecated) Sets session parameters. [Parameters](https://docs.snowflake.com/en/sql-reference/parameters)
- `token` (String, Sensitive) Token to use for OAuth and other forms of token based auth. Can also be sourced from the `SNOWFLAKE_TOKEN` environment variable.
- `tokenAccessor` (Block List, Max: 1) (see below for nested schema)
- `user` (String) Username. Can also be sourced from the `SNOWFLAKE_USER` environment variable. Required unless using `profile`.
- `username` (String, Deprecated) Username for username+password authentication. Can also be sourced from the `SNOWFLAKE_USERNAME` environment variable. Required unless using `profile`.
- `validateDefaultParameters` (Boolean) True by default. If false, disables the validation checks for Database, Schema, Warehouse and Role at the time a connection is established. Can also be sourced from the `SNOWFLAKE_VALIDATE_DEFAULT_PARAMETERS` environment variable.
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

In all cases account and username are required.
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
export SNOWFLAKE_PRIVATE_KEY_PATH="~/.ssh/snowflake_key"
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
export SNOWFLAKE_PRIVATE_KEY_PATH="~/.ssh/snowflake_key.p8"
export SNOWFLAKE_PRIVATE_KEY_PASSPHRASE="..."
```
### OAuth Access Token

If you have an OAuth access token, export these credentials as environment variables:

```shell
export SNOWFLAKE_USER='...'
export SNOWFLAKE_OAUTH_ACCESS_TOKEN='...'
```

Note that once this access token expires, you'll need to request a new one through an external application.
### OAuth Refresh Token

If you have an OAuth Refresh token, export these credentials as environment variables:

```shell
export SNOWFLAKE_OAUTH_REFRESH_TOKEN='...'
export SNOWFLAKE_OAUTH_CLIENT_ID='...'
export SNOWFLAKE_OAUTH_CLIENT_SECRET='...'
export SNOWFLAKE_OAUTH_ENDPOINT='...'
export SNOWFLAKE_OAUTH_REDIRECT_URL='https://localhost.com'
```

Note because access token have a short life; typically 10 minutes, by passing refresh token new access token will be generated.
### Username and Password Environment Variables

If you choose to use Username and Password Authentication, export these credentials:

```shell
export SNOWFLAKE_USER='...'
export SNOWFLAKE_PASSWORD='...'
```
### Config File

If you choose to use a config file, the optional `profile` attribute specifies the profile to use from the config file. If no profile is specified, the default profile is used. The Snowflake config file lives at `~/.snowflake/config` and uses [TOML](https://toml.io/) format. You can override this location by setting the `SNOWFLAKE_CONFIG_PATH` environment variable. If no username and account are specified, the provider will fall back to reading the config file.

```shell
[default]
account='TESTACCOUNT'
user='TEST_USER'
password='hunter2'
role='ACCOUNTADMIN'

[securityadmin]
account='TESTACCOUNT'
user='TEST_USER'
password='hunter2'
role='SECURITYADMIN'
```
## Order Precedence

The Snowflake provider will use the following order of precedence when determining which credentials to use:
1) Provider Configuration
2) Environment Variables
3) Config File
## Currently deprecated resources

- snowflake.DatabaseOld
- snowflake.OauthIntegration
- snowflake.Role - use snowflake.AccountRole instead
- snowflake.SamlIntegration - use snowflake.Saml2Integration instead
## Currently deprecated functions

- snowflake.Role - use snowflake.getRoles instead