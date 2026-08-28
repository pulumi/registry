---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-snowflake/v2.20.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi/pulumi-snowflake/blob/v2.20.0/docs/_index.md
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

## Overview

> **Disclaimer** The project is in GA version, but some features are in preview. Such resources and functions are considered preview features in the provider, regardless of their state in Snowflake. We do not guarantee their stability. They will be reworked and marked as a stable feature in future releases. Breaking changes in these features are expected, even without bumping the major version. They are disabled by default. To use them, add the relevant feature name to `previewFeaturesEnabled` field in the provider configuration. The list of preview features is available below. Please always refer to the Getting Help section in our Github repo to best determine how to get help for your questions. You can also use `experimentalFeaturesEnabled` to alter the provider's behavior. **It's still considered a preview feature, even when applied to the stable resources.**

> **Sensitive values** Important: Do not include credentials, personal identifiers, or other regulated or sensitive information (e.g., GDPR, HIPAA, PCI-DSS data) in non-sensitive fields. Snowflake marks specific fields as sensitive—such as passwords, private keys, and tokens, meaning these fields will not appear in logs. Each sensitive field is properly marked in the documentation. All other fields are treated as non-sensitive by default. Some of them, like task's configuration, may contain sensitive information but are not marked as sensitive - you are responsible for safeguarding these fields according to your organization's security standards and regulatory requirements. Snowflake will not be liable for any exposure of data placed in non-sensitive fields. Read more in the Sensitive values limitations section.

This is a pulumi provider plugin for managing [Snowflake](https://www.snowflake.com/) accounts.
Coverage is focused on part of Snowflake related to access control.
## Example Provider Configuration

This is an example configuration of the provider. More examples are provided below.

{{< chooser language "typescript,python,go,csharp,java,yaml,hcl" >}}
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
    }
}
```

{{% /choosable %}}
{{% choosable language hcl %}}
```hcl
Example currently unavailable in this language
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

- `account` (String) Specifies the Snowflake account identifier. Can be provided in the `org-name` format (e.g. `"myorg-myaccount"`) or as an account locator (e.g. `"xy12345"`). Use as a fallback when `accountName` and `organizationName` are not set. If both `accountName` and `organizationName` are set, they take precedence. Requires the `PROVIDER_CONFIGURATION_ACCOUNT_FALLBACK` experiment to be enabled. Can also be sourced from the `SNOWFLAKE_ACCOUNT` environment variable; without the experiment, the variable's value is ignored with a warning instead of resulting in an error.
- `accountName` (String) Specifies your Snowflake account name assigned by Snowflake. For information about account identifiers, see the [Snowflake documentation](https://docs.snowflake.com/en/user-guide/admin-account-identifier#account-name). Required unless using `profile`. Can also be sourced from the `SNOWFLAKE_ACCOUNT_NAME` environment variable.
- `authenticator` (String) Specifies the [authentication type](https://pkg.go.dev/github.com/snowflakedb/gosnowflake#AuthType) to use when connecting to Snowflake. Valid options are: `SNOWFLAKE` | `OAUTH` | `EXTERNALBROWSER` | `OKTA` | `SNOWFLAKE_JWT` | `TOKENACCESSOR` | `USERNAMEPASSWORDMFA` | `PROGRAMMATIC_ACCESS_TOKEN` | `OAUTH_CLIENT_CREDENTIALS` | `OAUTH_AUTHORIZATION_CODE` | `WORKLOAD_IDENTITY`. Can also be sourced from the `SNOWFLAKE_AUTHENTICATOR` environment variable.
- `certRevocationCheckMode` (String) Specifies the certificate revocation check mode. Valid options are: `DISABLED` | `ADVISORY` | `ENABLED`. The value is case-insensitive. Can also be sourced from the `SNOWFLAKE_CERT_REVOCATION_CHECK_MODE` environment variable.
- `clientIp` (String, Deprecated) This field is deprecated. It will be removed in the next major release. The driver was accepting this value in the previous versions but it had no impact. Setting this field causes no action on the provider side. Can also be sourced from the `SNOWFLAKE_CLIENT_IP` environment variable.
- `clientRequestMfaToken` (String) When true the MFA token is cached in the credential manager. True by default in Windows/OSX. False for Linux. Can also be sourced from the `SNOWFLAKE_CLIENT_REQUEST_MFA_TOKEN` environment variable.
- `clientStoreTemporaryCredential` (String) When true the ID token is cached in the credential manager. True by default in Windows/OSX. False for Linux. Can also be sourced from the `SNOWFLAKE_CLIENT_STORE_TEMPORARY_CREDENTIAL` environment variable.
- `clientTimeout` (Number) The timeout in seconds for the client to complete the authentication. Can also be sourced from the `SNOWFLAKE_CLIENT_TIMEOUT` environment variable.
- `crlAllowCertificatesWithoutCrlUrl` (String) Allow certificates (not short-lived) without CRL DP included to be treated as correct ones. Can also be sourced from the `SNOWFLAKE_CRL_ALLOW_CERTIFICATES_WITHOUT_CRL_URL` environment variable.
- `crlHttpClientTimeout` (Number) Timeout in seconds for HTTP client used to download CRL. Can also be sourced from the `SNOWFLAKE_CRL_HTTP_CLIENT_TIMEOUT` environment variable.
- `crlInMemoryCacheDisabled` (Boolean) False by default. When set to true, the CRL in-memory cache is disabled. Can also be sourced from the `SNOWFLAKE_CRL_IN_MEMORY_CACHE_DISABLED` environment variable.
- `crlOnDiskCacheDisabled` (Boolean) False by default. When set to true, the CRL on-disk cache is disabled. Can also be sourced from the `SNOWFLAKE_CRL_ON_DISK_CACHE_DISABLED` environment variable.
- `disableConsoleLogin` (String) Indicates whether console login should be disabled in the driver. Can also be sourced from the `SNOWFLAKE_DISABLE_CONSOLE_LOGIN` environment variable.
- `disableOcspChecks` (Boolean) False by default. When set to true, the driver doesn't check certificate revocation status. Can also be sourced from the `SNOWFLAKE_DISABLE_OCSP_CHECKS` environment variable.
- `disableQueryContextCache` (Boolean) Disables HTAP query context cache in the driver. Can also be sourced from the `SNOWFLAKE_DISABLE_QUERY_CONTEXT_CACHE` environment variable.
- `disableSamlUrlCheck` (String) Indicates whether the SAML URL check should be disabled. Can also be sourced from the `SNOWFLAKE_DISABLE_SAML_URL_CHECK` environment variable.
- `disableTelemetry` (Boolean, Deprecated) This field is deprecated. It will be removed in the next major release. Use `params` to set `CLIENT_TELEMETRY_ENABLED` session parameter instead. Setting this field adds `CLIENT_TELEMETRY_ENABLED` with value `false` to `params`. Disables telemetry in the driver. Can also be sourced from the `DISABLE_TELEMETRY` environment variable.
- `driverTracing` (String) Specifies the logging level to be used by the driver. Valid options are (case-insensitive): `TRACE` | `DEBUG` | `INFO` | `WARN` | `ERROR` | `FATAL` | `OFF`. The following values are deprecated and will be removed in v3: `WARNING` (uses `WARN` instead), `PRINT` (uses `INFO` instead), `PANIC` (uses `FATAL` instead). Can also be sourced from the `SNOWFLAKE_DRIVER_TRACING` environment variable.
- `enableSingleUseRefreshTokens` (Boolean) Enables single use refresh tokens for Snowflake IdP. Can also be sourced from the `SNOWFLAKE_ENABLE_SINGLE_USE_REFRESH_TOKENS` environment variable.
- `experimentalFeaturesEnabled` (Set of String) A list of experimental features. Similarly to preview features, they are not yet stable features of the provider. Enabling given experiment is still considered a preview feature, even when applied to the stable resource. These switches offer experiments altering the provider behavior. If the given experiment is successful, it can be considered an addition in the future provider versions. This field can not be set with environmental variables. Check more details in the experimental features section. Active experiments are: `WAREHOUSE_SHOW_IMPROVED_PERFORMANCE` | `GRANTS_STRICT_PRIVILEGE_MANAGEMENT` | `PARAMETERS_IGNORE_VALUE_CHANGES_IF_NOT_ON_OBJECT_LEVEL` | `PARAMETERS_REDUCED_OUTPUT` | `USER_ENABLE_DEFAULT_WORKLOAD_IDENTITY` | `GRANTS_IMPORT_VALIDATION` | `TAGS_ALLOW_EMPTY_ALLOWED_VALUES` | `IMPORT_BOOLEAN_DEFAULT` | `GRANTS_SAFE_DESTROY` | `TAG_ASSOCIATION_SAFE_DESTROY` | `GRANT_ACCOUNT_ROLE_SHOW_CACHING` | `ACCOUNT_ROLE_SHOW_CACHING` | `GRANTS_SHOW_CACHING` | `GRANT_ACCOUNT_ROLE_SAFE_PUBLIC_ROLE` | `HIERARCHY_RENAMES` | `INHERITED_GRANTS` | `OBJECT_PARAMETER_UNSET_ON_DELETE` | `AUTHENTICATOR_EXPLICIT_ONLY` | `PROVIDER_CONFIGURATION_ACCOUNT_FALLBACK`.
- `externalBrowserTimeout` (Number) The timeout in seconds for the external browser to complete the authentication. Can also be sourced from the `SNOWFLAKE_EXTERNAL_BROWSER_TIMEOUT` environment variable.
- `host` (String) Specifies a custom host value used by the driver for privatelink connections. Can also be sourced from the `SNOWFLAKE_HOST` environment variable.
- `includeRetryReason` (String) Should retried request contain retry reason. Can also be sourced from the `SNOWFLAKE_INCLUDE_RETRY_REASON` environment variable.
- `insecureMode` (Boolean, Deprecated) This field is deprecated. It will be removed in the next major release. Use `disableOcspChecks` instead. Setting this field sets `disableOcspChecks` in the underlying driver. If true, bypass the Online Certificate Status Protocol (OCSP) certificate revocation check. IMPORTANT: Change the default value for testing or emergency situations only. Can also be sourced from the `SNOWFLAKE_INSECURE_MODE` environment variable.
- `jwtClientTimeout` (Number) The timeout in seconds for the JWT client to complete the authentication. Can also be sourced from the `SNOWFLAKE_JWT_CLIENT_TIMEOUT` environment variable.
- `jwtExpireTimeout` (Number) JWT expire after timeout in seconds. Can also be sourced from the `SNOWFLAKE_JWT_EXPIRE_TIMEOUT` environment variable.
- `keepSessionAlive` (Boolean) Enables the session to persist even after the connection is closed. Can also be sourced from the `SNOWFLAKE_KEEP_SESSION_ALIVE` environment variable.
- `logQueryParameters` (Boolean) When set to true, the parameters will be logged. Requires logQueryText to be enabled first. Be aware that it may include sensitive information. Default value is false. Can also be sourced from the `SNOWFLAKE_LOG_QUERY_PARAMETERS` environment variable.
- `logQueryText` (Boolean) When set to true, the full query text will be logged. Be aware that it may include sensitive information. Default value is false. Can also be sourced from the `SNOWFLAKE_LOG_QUERY_TEXT` environment variable.
- `loginTimeout` (Number) Login retry timeout in seconds EXCLUDING network roundtrip and read out http response. Can also be sourced from the `SNOWFLAKE_LOGIN_TIMEOUT` environment variable.
- `maxRetryCount` (Number) Specifies how many times non-periodic HTTP request can be retried by the driver. Can also be sourced from the `SNOWFLAKE_MAX_RETRY_COUNT` environment variable.
- `noProxy` (String) A comma-separated list of hostnames, domains, and IP addresses to exclude from proxying. See more in the proxy section below. Can also be sourced from the `SNOWFLAKE_NO_PROXY` environment variable.
- `oauthAuthorizationUrl` (String, Sensitive) Authorization URL of OAuth2 external IdP. See [Snowflake OAuth documentation](https://docs.snowflake.com/en/user-guide/oauth). Can also be sourced from the `SNOWFLAKE_OAUTH_AUTHORIZATION_URL` environment variable.
- `oauthClientId` (String, Sensitive) Client id for OAuth2 external IdP. See [Snowflake OAuth documentation](https://docs.snowflake.com/en/user-guide/oauth). Can also be sourced from the `SNOWFLAKE_OAUTH_CLIENT_ID` environment variable.
- `oauthClientSecret` (String, Sensitive) Client secret for OAuth2 external IdP. See [Snowflake OAuth documentation](https://docs.snowflake.com/en/user-guide/oauth). Can also be sourced from the `SNOWFLAKE_OAUTH_CLIENT_SECRET` environment variable.
- `oauthRedirectUri` (String, Sensitive) Redirect URI registered in IdP. See [Snowflake OAuth documentation](https://docs.snowflake.com/en/user-guide/oauth). Can also be sourced from the `SNOWFLAKE_OAUTH_REDIRECT_URI` environment variable.
- `oauthScope` (String) Comma separated list of scopes. If empty it is derived from role. See [Snowflake OAuth documentation](https://docs.snowflake.com/en/user-guide/oauth). Can also be sourced from the `SNOWFLAKE_OAUTH_SCOPE` environment variable.
- `oauthTokenRequestUrl` (String, Sensitive) Token request URL of OAuth2 external IdP. See [Snowflake OAuth documentation](https://docs.snowflake.com/en/user-guide/oauth). Can also be sourced from the `SNOWFLAKE_OAUTH_TOKEN_REQUEST_URL` environment variable.
- `ocspFailOpen` (String) True represents OCSP fail open mode. False represents OCSP fail closed mode. Fail open true by default. Can also be sourced from the `SNOWFLAKE_OCSP_FAIL_OPEN` environment variable.
- `oktaUrl` (String) The URL of the Okta server. e.g. <https://example.okta.com>. Okta URL host needs to to have a suffix `okta.com`. Read more in Snowflake [docs](https://docs.snowflake.com/en/user-guide/oauth-okta). Can also be sourced from the `SNOWFLAKE_OKTA_URL` environment variable.
- `organizationName` (String) Specifies your Snowflake organization name assigned by Snowflake. For information about account identifiers, see the [Snowflake documentation](https://docs.snowflake.com/en/user-guide/admin-account-identifier#organization-name). Required unless using `profile`. Can also be sourced from the `SNOWFLAKE_ORGANIZATION_NAME` environment variable.
- `params` (Map of String) Sets other connection (i.e. session) parameters. [Parameters](https://docs.snowflake.com/en/sql-reference/parameters). This field can not be set with environmental variables.
- `passcode` (String, Sensitive) Specifies the passcode provided by Duo when using multi-factor authentication (MFA) for login. Can also be sourced from the `SNOWFLAKE_PASSCODE` environment variable.
- `passcodeInPassword` (Boolean) False by default. Set to true if the MFA passcode is embedded to the configured password. Can also be sourced from the `SNOWFLAKE_PASSCODE_IN_PASSWORD` environment variable.
- `password` (String, Sensitive) Password for user + password or [token](https://docs.snowflake.com/en/user-guide/programmatic-access-tokens#generating-a-programmatic-access-token) for [PAT auth](https://docs.snowflake.com/en/user-guide/programmatic-access-tokens). Cannot be used with `privateKey` and `privateKeyPassphrase`. Can also be sourced from the `SNOWFLAKE_PASSWORD` environment variable.
- `port` (Number) Specifies a custom port value used by the driver for privatelink connections. Can also be sourced from the `SNOWFLAKE_PORT` environment variable.
- `previewFeaturesEnabled` (Set of String) A list of preview features that are handled by the provider. See preview features list. Preview features may have breaking changes in future releases, even without raising the major version. This field can not be set with environmental variables. Preview features that can be enabled are: <code>snowflake_account_authentication_policy_attachment_resource</code> | <code>snowflake_account_password_policy_attachment_resource</code> | <code>snowflake_alert_resource</code> | <code>snowflake_alerts_datasource</code> | <code>snowflake_api_integrations_datasource</code> | <code>snowflake_api_integration_resource</code> | <code>snowflake_api_integration_amazon_api_gateway_resource</code> | <code>snowflake_api_integration_azure_api_management_resource</code> | <code>snowflake_api_integration_external_mcp_dynamic_client_resource</code> | <code>snowflake_api_integration_external_mcp_oauth2_resource</code> | <code>snowflake_api_integration_git_repository_github_app_resource</code> | <code>snowflake_api_integration_git_repository_oauth2_resource</code> | <code>snowflake_api_integration_git_repository_private_link_resource</code> | <code>snowflake_api_integration_git_repository_token_resource</code> | <code>snowflake_api_integration_google_cloud_api_gateway_resource</code> | <code>snowflake_cortex_agent_resource</code> | <code>snowflake_cortex_agents_datasource</code> | <code>snowflake_cortex_search_service_resource</code> | <code>snowflake_cortex_search_services_datasource</code> | <code>snowflake_current_account_datasource</code> | <code>snowflake_database_datasource</code> | <code>snowflake_database_role_datasource</code> | <code>snowflake_dynamic_table_resource</code> | <code>snowflake_dynamic_tables_datasource</code> | <code>snowflake_external_access_integration_resource</code> | <code>snowflake_external_access_integrations_datasource</code> | <code>snowflake_external_function_resource</code> | <code>snowflake_external_functions_datasource</code> | <code>snowflake_external_table_resource</code> | <code>snowflake_external_tables_datasource</code> | <code>snowflake_failover_group_resource</code> | <code>snowflake_failover_groups_datasource</code> | <code>snowflake_file_format_resource</code> | <code>snowflake_file_format_avro_resource</code> | <code>snowflake_file_format_csv_resource</code> | <code>snowflake_file_format_json_resource</code> | <code>snowflake_file_format_orc_resource</code> | <code>snowflake_file_format_parquet_resource</code> | <code>snowflake_file_format_xml_resource</code> | <code>snowflake_file_formats_datasource</code> | <code>snowflake_function_java_resource</code> | <code>snowflake_function_javascript_resource</code> | <code>snowflake_function_python_resource</code> | <code>snowflake_function_scala_resource</code> | <code>snowflake_function_sql_resource</code> | <code>snowflake_functions_datasource</code> | <code>snowflake_hybrid_table_resource</code> | <code>snowflake_iceberg_table_resource</code> | <code>snowflake_iceberg_table_from_aws_glue_resource</code> | <code>snowflake_iceberg_table_from_delta_files_resource</code> | <code>snowflake_iceberg_table_from_files_resource</code> | <code>snowflake_iceberg_table_from_rest_resource</code> | <code>snowflake_iceberg_tables_datasource</code> | <code>snowflake_job_service_resource</code> | <code>snowflake_listings_datasource</code> | <code>snowflake_managed_account_resource</code> | <code>snowflake_materialized_view_resource</code> | <code>snowflake_materialized_views_datasource</code> | <code>snowflake_mcp_server_resource</code> | <code>snowflake_mcp_servers_datasource</code> | <code>snowflake_network_policy_attachment_resource</code> | <code>snowflake_notebook_resource</code> | <code>snowflake_notebooks_datasource</code> | <code>snowflake_email_notification_integration_resource</code> | <code>snowflake_notification_integration_resource</code> | <code>snowflake_object_parameter_resource</code> | <code>snowflake_pipe_resource</code> | <code>snowflake_pipes_datasource</code> | <code>snowflake_postgres_instance_resource</code> | <code>snowflake_current_role_datasource</code> | <code>snowflake_semantic_view_resource</code> | <code>snowflake_semantic_views_datasource</code> | <code>snowflake_sequence_resource</code> | <code>snowflake_sequences_datasource</code> | <code>snowflake_share_resource</code> | <code>snowflake_shares_datasource</code> | <code>snowflake_parameters_datasource</code> | <code>snowflake_procedure_java_resource</code> | <code>snowflake_procedure_javascript_resource</code> | <code>snowflake_procedure_python_resource</code> | <code>snowflake_procedure_scala_resource</code> | <code>snowflake_procedure_sql_resource</code> | <code>snowflake_procedures_datasource</code> | <code>snowflake_stage_resource</code> | <code>snowflake_stages_datasource</code> | <code>snowflake_storage_integration_resource</code> | <code>snowflake_storage_lifecycle_policy_resource</code> | <code>snowflake_storage_lifecycle_policies_datasource</code> | <code>snowflake_system_generate_scim_access_token_datasource</code> | <code>snowflake_system_get_aws_sns_iam_policy_datasource</code> | <code>snowflake_system_get_privatelink_config_datasource</code> | <code>snowflake_system_get_snowflake_platform_info_datasource</code> | <code>snowflake_table_column_masking_policy_application_resource</code> | <code>snowflake_table_constraint_resource</code> | <code>snowflake_table_storage_lifecycle_policy_attachment_resource</code> | <code>snowflake_table_resource</code> | <code>snowflake_tables_datasource</code> | <code>snowflake_user_authentication_policy_attachment_resource</code> | <code>snowflake_user_password_policy_attachment_resource</code> | <code>snowflake_user_public_keys_resource</code> | <code>snowflake_warehouse_adaptive_resource</code> | <code>snowflake_warehouse_interactive_resource</code>. Promoted features that are stable and are enabled by default are: <code>snowflake_account_session_policy_attachment_resource</code> | <code>snowflake_authentication_policy_resource</code> | <code>snowflake_authentication_policies_datasource</code> | <code>snowflake_catalog_integration_aws_glue_resource</code> | <code>snowflake_catalog_integration_object_storage_resource</code> | <code>snowflake_catalog_integration_open_catalog_resource</code> | <code>snowflake_catalog_integration_iceberg_rest_resource</code> | <code>snowflake_catalog_integrations_datasource</code> | <code>snowflake_compute_pool_resource</code> | <code>snowflake_compute_pools_datasource</code> | <code>snowflake_current_account_resource</code> | <code>snowflake_current_organization_account_resource</code> | <code>snowflake_stage_external_azure_resource</code> | <code>snowflake_stage_external_gcs_resource</code> | <code>snowflake_stage_external_s3_resource</code> | <code>snowflake_stage_external_s3_compatible_resource</code> | <code>snowflake_external_volume_resource</code> | <code>snowflake_external_volumes_datasource</code> | <code>snowflake_git_repository_resource</code> | <code>snowflake_git_repositories_datasource</code> | <code>snowflake_image_repository_resource</code> | <code>snowflake_image_repositories_datasource</code> | <code>snowflake_stage_internal_resource</code> | <code>snowflake_listing_resource</code> | <code>snowflake_network_rule_resource</code> | <code>snowflake_network_rules_datasource</code> | <code>snowflake_password_policies_datasource</code> | <code>snowflake_password_policy_resource</code> | <code>snowflake_service_resource</code> | <code>snowflake_services_datasource</code> | <code>snowflake_session_policies_datasource</code> | <code>snowflake_session_policy_resource</code> | <code>snowflake_storage_integration_aws_resource</code> | <code>snowflake_storage_integration_azure_resource</code> | <code>snowflake_storage_integration_gcs_resource</code> | <code>snowflake_storage_integrations_datasource</code> | <code>snowflake_user_programmatic_access_token_resource</code> | <code>snowflake_user_programmatic_access_tokens_datasource</code> | <code>snowflake_user_session_policy_attachment_resource</code>. Promoted features can be safely removed from this field. They will be removed in the next major version.
- `privateKey` (String, Sensitive) Private Key for username+private-key auth. Must be PEM-encoded with literal newlines (escaped `\n` sequences are not supported). See the authentication methods guide. Cannot be used with `password`. Can also be sourced from the `SNOWFLAKE_PRIVATE_KEY` environment variable.
- `privateKeyPassphrase` (String, Sensitive) Supports the encryption ciphers aes-128-cbc, aes-128-gcm, aes-192-cbc, aes-192-gcm, aes-256-cbc, aes-256-gcm, and des-ede3-cbc. Can also be sourced from the `SNOWFLAKE_PRIVATE_KEY_PASSPHRASE` environment variable.
- `profile` (String) Sets the profile to read from ~/.snowflake/config file. Can also be sourced from the `SNOWFLAKE_PROFILE` environment variable.
- `protocol` (String) A protocol used in the connection. Valid options are: `http` | `https`. Can also be sourced from the `SNOWFLAKE_PROTOCOL` environment variable.
- `proxyHost` (String) The host of the proxy to use for the connection. See more in the proxy section below. Can also be sourced from the `SNOWFLAKE_PROXY_HOST` environment variable.
- `proxyPassword` (String, Sensitive) The password of the proxy to use for the connection. See more in the proxy section below. Can also be sourced from the `SNOWFLAKE_PROXY_PASSWORD` environment variable.
- `proxyPort` (Number) The port of the proxy to use for the connection. See more in the proxy section below. Can also be sourced from the `SNOWFLAKE_PROXY_PORT` environment variable.
- `proxyProtocol` (String) The protocol of the proxy to use for the connection. Valid options are: `http` | `https`. The value is case-insensitive. See more in the proxy section below. Can also be sourced from the `SNOWFLAKE_PROXY_PROTOCOL` environment variable.
- `proxyUser` (String) The user of the proxy to use for the connection. See more in the proxy section below. Can also be sourced from the `SNOWFLAKE_PROXY_USER` environment variable.
- `requestTimeout` (Number) request retry timeout in seconds EXCLUDING network roundtrip and read out http response. Can also be sourced from the `SNOWFLAKE_REQUEST_TIMEOUT` environment variable.
- `role` (String) Specifies the role to use by default for accessing Snowflake objects in the client session. Can also be sourced from the `SNOWFLAKE_ROLE` environment variable.
- `skipTomlFilePermissionVerification` (Boolean, Deprecated) This field is deprecated. It will be removed in the next major release. False by default. Skips TOML configuration file permission verification. This flag has no effect on Windows systems, as the permissions are not checked on this platform. Instead of skipping the permissions verification, we recommend setting the proper privileges - see the section below. Can also be sourced from the `SNOWFLAKE_SKIP_TOML_FILE_PERMISSION_VERIFICATION` environment variable.
- `tfcWorkloadIdentityTokenTag` (String) Tag suffix used to read the Pulumi Cloud/Enterprise workload identity token from the `TFC_WORKLOAD_IDENTITY_TOKEN_<TAG>` environment variable (the tag is upper-cased). Requires `authenticator` to be `WORKLOAD_IDENTITY` and `workloadIdentityProvider` to be `OIDC`. Takes precedence over `token` and every other token source. Can also be sourced from the `SNOWFLAKE_TFC_WORKLOAD_IDENTITY_TOKEN_TAG` environment variable.
- `tmpDirectoryPath` (String) Sets temporary directory used by the driver for operations like encrypting, compressing etc. Can also be sourced from the `SNOWFLAKE_TMP_DIRECTORY_PATH` environment variable.
- `token` (String, Sensitive) Token to use for OAuth and other forms of token based auth. When this field is set here, or in the TOML file, the provider sets the `authenticator` to `OAUTH`. Optionally, set the `authenticator` field to the authenticator you want to use. Can also be sourced from the `SNOWFLAKE_TOKEN` environment variable.
- `tokenAccessor` (Block List, Max: 1) If you are using the OAuth authentication flows, use the dedicated `authenticator` and `oauth...` fields instead. See our authentication methods guide for more information. (see below for nested schema)
- `useLegacyTomlFile` (Boolean) False by default. When this is set to true, the provider expects the legacy TOML format. Otherwise, it expects the new format. See more in the section below Can also be sourced from the `SNOWFLAKE_USE_LEGACY_TOML_FILE` environment variable.
- `user` (String) Username. Required unless using `profile`. Can also be sourced from the `SNOWFLAKE_USER` environment variable.
- `validateDefaultParameters` (String) True by default. If false, disables the validation checks for Database, Schema, Warehouse and Role at the time a connection is established. Can also be sourced from the `SNOWFLAKE_VALIDATE_DEFAULT_PARAMETERS` environment variable.
- `warehouse` (String) Specifies the virtual warehouse to use by default for queries, loading, etc. in the client session. Can also be sourced from the `SNOWFLAKE_WAREHOUSE` environment variable.
- `workloadIdentityEntraResource` (String) The resource to use for WIF authentication on Azure environment. Can also be sourced from the `SNOWFLAKE_WORKLOAD_IDENTITY_ENTRA_RESOURCE` environment variable.
- `workloadIdentityProvider` (String) The workload identity provider to use for WIF authentication. Can also be sourced from the `SNOWFLAKE_WORKLOAD_IDENTITY_PROVIDER` environment variable.

<a id="nestedblock--token_accessor"></a>
### Nested Schema for `tokenAccessor`

Required:

- `clientId` (String, Sensitive) The client ID for the OAuth provider when using a refresh token to renew access token. Can also be sourced from the `SNOWFLAKE_TOKEN_ACCESSOR_CLIENT_ID` environment variable.
- `clientSecret` (String, Sensitive) The client secret for the OAuth provider when using a refresh token to renew access token. Can also be sourced from the `SNOWFLAKE_TOKEN_ACCESSOR_CLIENT_SECRET` environment variable.
- `redirectUri` (String, Sensitive) The redirect URI for the OAuth provider when using a refresh token to renew access token. Can also be sourced from the `SNOWFLAKE_TOKEN_ACCESSOR_REDIRECT_URI` environment variable.
- `refreshToken` (String, Sensitive) The refresh token for the OAuth provider when using a refresh token to renew access token. Can also be sourced from the `SNOWFLAKE_TOKEN_ACCESSOR_REFRESH_TOKEN` environment variable.
- `tokenEndpoint` (String, Sensitive) The token endpoint for the OAuth provider e.g. https://{yourDomain}/oauth/token when using a refresh token to renew access token. Can also be sourced from the `SNOWFLAKE_TOKEN_ACCESSOR_TOKEN_ENDPOINT` environment variable.
## Authentication

The Snowflake provider supports multiple ways to authenticate:

* Password
* PAT (Personal Access Token)
* OAuth Access Token
* OAuth Refresh Token
* Browser Auth
* Private Key
* Config File
* Oauth with Client Credentials
* Oauth with Authorization Code
* Workload Identity Federation (WIF)

In all cases `organizationName`, and `accountName` are required. In all cases except for Oauth with Client Credentials, `user` is required.

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

> **Note** See the authentication methods guide for more details about the private key format.
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
### Oauth with Client Credentials Environment Variables

If you choose to use Oauth with Client Credentials Authentication, export these credentials:

```shell
export SNOWFLAKE_OAUTH_CLIENT_ID='...'
export SNOWFLAKE_OAUTH_CLIENT_SECRET='...'
export SNOWFLAKE_OAUTH_TOKEN_REQUEST_URL='...'
```
### Oauth with Authorization Code Environment Variables

If you choose to use Oauth with Authorization Code Authentication, export these credentials:

```shell
export SNOWFLAKE_OAUTH_CLIENT_ID='...'
export SNOWFLAKE_OAUTH_CLIENT_SECRET='...'
export SNOWFLAKE_OAUTH_AUTHORIZATION_URL='...'
export SNOWFLAKE_OAUTH_TOKEN_REQUEST_URL='...'
export SNOWFLAKE_OAUTH_REDIRECT_URI='...'
export SNOWFLAKE_OAUTH_SCOPE='...'
```
### Workload Identity Federation (WIF) Authentication

If you choose to use Workload Identity Federation (WIF) Authentication, export these credentials:

```shell
export SNOWFLAKE_WORKLOAD_IDENTITY_PROVIDER='...'
export SNOWFLAKE_WORKLOAD_IDENTITY_ENTRA_RESOURCE='...'
```
## Order Precedence

Currently, the provider can be configured in three ways:
1. In a Pulumi file located in the Pulumi module with other resources.
2. In environmental variables (envs). This is mainly used to provide sensitive values.
3. In a TOML file (default in `~/.snowflake/config`).
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

> **Note** This configuration file is distinct from the ones used to configure [Snowflake CLI](https://docs.snowflake.com/en/developer-guide/snowflake-cli/connecting/configure-cli) or [SnowSQL](https://docs.snowflake.com/en/user-guide/snowsql-config).

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

> **Note** TOML file size is limited to 10MB.

> **Note** Only TOML file with restricted privileges can be read. Any privileges for group or others cannot be set (the maximum valid privilege is `700`). You can set the expected privileges like `chmod 0600 ~/.snowflake/config`. This is checked only on non-Windows platforms. If you are using the provider on Windows, please make sure that your configuration file has not too permissive privileges.
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
oauth_client_id = 'oauth_client_id'
oauth_client_secret = 'oauth_client_secret'
oauth_token_request_url = 'oauth_token_request_url'
oauth_authorization_url = 'oauth_authorization_url'
oauth_redirect_uri = 'oauth_redirect_uri'
oauth_scope = 'oauth_scope'
workload_identity_provider = 'azure'
workload_identity_entra_resource = 'workload_identity_entra_resource'
enable_single_use_refresh_tokens = true
log_query_text = false
log_query_parameters = false
proxy_host = 'proxy.example.com'
proxy_port = 443
proxy_user = 'username'
proxy_password = 'proxy_password'
proxy_protocol = 'https'
no_proxy = 'localhost,snowflake.computing.com'
disable_ocsp_checks = true
cert_revocation_check_mode = 'ADVISORY'
crl_allow_certificates_without_crl_url = true
crl_in_memory_cache_disabled = false
crl_on_disk_cache_disabled = true
crl_http_client_timeout = 30
disable_saml_url_check = true

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
oauthclientid = 'oauth_client_id'
oauthclientsecret = 'oauth_client_secret'
oauthtokenrequesturl = 'oauth_token_request_url'
oauthauthorizationurl = 'oauth_authorization_url'
oauthredirecturi = 'oauth_redirect_uri'
oauthscope = 'oauth_scope'
workloadidentityprovider = 'azure'
workloadidentityentraresource = 'workload_identity_entra_resource'
enablesingleuserefreshtokens = true
logquerytext = false
logqueryparameters = false
proxyhost = 'proxy.example.com'
proxyport = 443
proxyuser = 'username'
proxypassword = '****'
proxyprotocol = 'https'
noproxy = 'localhost,snowflake.computing.com'
disableocspchecks = true
certrevocationcheckmode = 'ADVISORY'
crlallowcertificateswithoutcrlurl = true
crlinmemorycachedisabled = false
crlondiskcachedisabled = true
crlhttpclienttimeout = 30
disablesamlurlcheck = true

[example.params]
param_key = 'param_value'
```

An example pulumi configuration file equivalent:

{{< chooser language "typescript,python,go,csharp,java,yaml,hcl" >}}
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
    snowflake:certRevocationCheckMode:
        value: ADVISORY
    snowflake:clientIp:
        value: 1.2.3.4
    snowflake:clientRequestMfaToken:
        value: true
    snowflake:clientStoreTemporaryCredential:
        value: true
    snowflake:clientTimeout:
        value: 40
    snowflake:crlAllowCertificatesWithoutCrlUrl:
        value: true
    snowflake:crlHttpClientTimeout:
        value: 30
    snowflake:crlInMemoryCacheDisabled:
        value: false
    snowflake:crlOnDiskCacheDisabled:
        value: true
    snowflake:disableConsoleLogin:
        value: true
    snowflake:disableOcspChecks:
        value: true
    snowflake:disableQueryContextCache:
        value: true
    snowflake:disableSamlUrlCheck:
        value: true
    snowflake:disableTelemetry:
        value: true
    snowflake:driverTracing:
        value: info
    snowflake:enableSingleUseRefreshTokens:
        value: true
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
    snowflake:logQueryParameters:
        value: false
    snowflake:logQueryText:
        value: false
    snowflake:loginTimeout:
        value: 10
    snowflake:maxRetryCount:
        value: 3
    snowflake:noProxy:
        value: localhost,snowflake.computing.com
    snowflake:oauthAuthorizationUrl:
        value: 'TODO: var.oauth_authorization_url'
    snowflake:oauthClientId:
        value: 'TODO: var.oauth_client_id'
    snowflake:oauthClientSecret:
        value: 'TODO: var.oauth_client_secret'
    snowflake:oauthRedirectUri:
        value: 'TODO: var.oauth_redirect_uri'
    snowflake:oauthScope:
        value: session:role:PUBLIC
    snowflake:oauthTokenRequestUrl:
        value: 'TODO: var.oauth_token_request_url'
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
    snowflake:proxyHost:
        value: proxy.example.com
    snowflake:proxyPassword:
        value: 'TODO: var.proxy_password'
    snowflake:proxyPort:
        value: 443
    snowflake:proxyProtocol:
        value: https
    snowflake:proxyUser:
        value: username
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
    snowflake:workloadIdentityEntraResource:
        value: workload_identity_entra_resource
    snowflake:workloadIdentityProvider:
        value: azure

```

```typescript
import * as pulumi from "@pulumi/pulumi";

const config = new pulumi.Config();
const proxyPassword = config.require("proxyPassword");
const oauthClientId = config.require("oauthClientId");
const oauthClientSecret = config.require("oauthClientSecret");
const oauthTokenRequestUrl = config.require("oauthTokenRequestUrl");
const oauthAuthorizationUrl = config.require("oauthAuthorizationUrl");
const oauthRedirectUri = config.require("oauthRedirectUri");
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
    snowflake:certRevocationCheckMode:
        value: ADVISORY
    snowflake:clientIp:
        value: 1.2.3.4
    snowflake:clientRequestMfaToken:
        value: true
    snowflake:clientStoreTemporaryCredential:
        value: true
    snowflake:clientTimeout:
        value: 40
    snowflake:crlAllowCertificatesWithoutCrlUrl:
        value: true
    snowflake:crlHttpClientTimeout:
        value: 30
    snowflake:crlInMemoryCacheDisabled:
        value: false
    snowflake:crlOnDiskCacheDisabled:
        value: true
    snowflake:disableConsoleLogin:
        value: true
    snowflake:disableOcspChecks:
        value: true
    snowflake:disableQueryContextCache:
        value: true
    snowflake:disableSamlUrlCheck:
        value: true
    snowflake:disableTelemetry:
        value: true
    snowflake:driverTracing:
        value: info
    snowflake:enableSingleUseRefreshTokens:
        value: true
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
    snowflake:logQueryParameters:
        value: false
    snowflake:logQueryText:
        value: false
    snowflake:loginTimeout:
        value: 10
    snowflake:maxRetryCount:
        value: 3
    snowflake:noProxy:
        value: localhost,snowflake.computing.com
    snowflake:oauthAuthorizationUrl:
        value: 'TODO: var.oauth_authorization_url'
    snowflake:oauthClientId:
        value: 'TODO: var.oauth_client_id'
    snowflake:oauthClientSecret:
        value: 'TODO: var.oauth_client_secret'
    snowflake:oauthRedirectUri:
        value: 'TODO: var.oauth_redirect_uri'
    snowflake:oauthScope:
        value: session:role:PUBLIC
    snowflake:oauthTokenRequestUrl:
        value: 'TODO: var.oauth_token_request_url'
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
    snowflake:proxyHost:
        value: proxy.example.com
    snowflake:proxyPassword:
        value: 'TODO: var.proxy_password'
    snowflake:proxyPort:
        value: 443
    snowflake:proxyProtocol:
        value: https
    snowflake:proxyUser:
        value: username
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
    snowflake:workloadIdentityEntraResource:
        value: workload_identity_entra_resource
    snowflake:workloadIdentityProvider:
        value: azure

```

```python
import pulumi

config = pulumi.Config()
proxy_password = config.require("proxyPassword")
oauth_client_id = config.require("oauthClientId")
oauth_client_secret = config.require("oauthClientSecret")
oauth_token_request_url = config.require("oauthTokenRequestUrl")
oauth_authorization_url = config.require("oauthAuthorizationUrl")
oauth_redirect_uri = config.require("oauthRedirectUri")
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
    snowflake:certRevocationCheckMode:
        value: ADVISORY
    snowflake:clientIp:
        value: 1.2.3.4
    snowflake:clientRequestMfaToken:
        value: true
    snowflake:clientStoreTemporaryCredential:
        value: true
    snowflake:clientTimeout:
        value: 40
    snowflake:crlAllowCertificatesWithoutCrlUrl:
        value: true
    snowflake:crlHttpClientTimeout:
        value: 30
    snowflake:crlInMemoryCacheDisabled:
        value: false
    snowflake:crlOnDiskCacheDisabled:
        value: true
    snowflake:disableConsoleLogin:
        value: true
    snowflake:disableOcspChecks:
        value: true
    snowflake:disableQueryContextCache:
        value: true
    snowflake:disableSamlUrlCheck:
        value: true
    snowflake:disableTelemetry:
        value: true
    snowflake:driverTracing:
        value: info
    snowflake:enableSingleUseRefreshTokens:
        value: true
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
    snowflake:logQueryParameters:
        value: false
    snowflake:logQueryText:
        value: false
    snowflake:loginTimeout:
        value: 10
    snowflake:maxRetryCount:
        value: 3
    snowflake:noProxy:
        value: localhost,snowflake.computing.com
    snowflake:oauthAuthorizationUrl:
        value: 'TODO: var.oauth_authorization_url'
    snowflake:oauthClientId:
        value: 'TODO: var.oauth_client_id'
    snowflake:oauthClientSecret:
        value: 'TODO: var.oauth_client_secret'
    snowflake:oauthRedirectUri:
        value: 'TODO: var.oauth_redirect_uri'
    snowflake:oauthScope:
        value: session:role:PUBLIC
    snowflake:oauthTokenRequestUrl:
        value: 'TODO: var.oauth_token_request_url'
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
    snowflake:proxyHost:
        value: proxy.example.com
    snowflake:proxyPassword:
        value: 'TODO: var.proxy_password'
    snowflake:proxyPort:
        value: 443
    snowflake:proxyProtocol:
        value: https
    snowflake:proxyUser:
        value: username
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
    snowflake:workloadIdentityEntraResource:
        value: workload_identity_entra_resource
    snowflake:workloadIdentityProvider:
        value: azure

```

```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;

return await Deployment.RunAsync(() =>
{
    var config = new Config();
    var proxyPassword = config.Require("proxyPassword");
    var oauthClientId = config.Require("oauthClientId");
    var oauthClientSecret = config.Require("oauthClientSecret");
    var oauthTokenRequestUrl = config.Require("oauthTokenRequestUrl");
    var oauthAuthorizationUrl = config.Require("oauthAuthorizationUrl");
    var oauthRedirectUri = config.Require("oauthRedirectUri");
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
        value: account_name
    snowflake:authenticator:
        value: snowflake
    snowflake:certRevocationCheckMode:
        value: ADVISORY
    snowflake:clientIp:
        value: 1.2.3.4
    snowflake:clientRequestMfaToken:
        value: true
    snowflake:clientStoreTemporaryCredential:
        value: true
    snowflake:clientTimeout:
        value: 40
    snowflake:crlAllowCertificatesWithoutCrlUrl:
        value: true
    snowflake:crlHttpClientTimeout:
        value: 30
    snowflake:crlInMemoryCacheDisabled:
        value: false
    snowflake:crlOnDiskCacheDisabled:
        value: true
    snowflake:disableConsoleLogin:
        value: true
    snowflake:disableOcspChecks:
        value: true
    snowflake:disableQueryContextCache:
        value: true
    snowflake:disableSamlUrlCheck:
        value: true
    snowflake:disableTelemetry:
        value: true
    snowflake:driverTracing:
        value: info
    snowflake:enableSingleUseRefreshTokens:
        value: true
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
    snowflake:logQueryParameters:
        value: false
    snowflake:logQueryText:
        value: false
    snowflake:loginTimeout:
        value: 10
    snowflake:maxRetryCount:
        value: 3
    snowflake:noProxy:
        value: localhost,snowflake.computing.com
    snowflake:oauthAuthorizationUrl:
        value: 'TODO: var.oauth_authorization_url'
    snowflake:oauthClientId:
        value: 'TODO: var.oauth_client_id'
    snowflake:oauthClientSecret:
        value: 'TODO: var.oauth_client_secret'
    snowflake:oauthRedirectUri:
        value: 'TODO: var.oauth_redirect_uri'
    snowflake:oauthScope:
        value: session:role:PUBLIC
    snowflake:oauthTokenRequestUrl:
        value: 'TODO: var.oauth_token_request_url'
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
    snowflake:proxyHost:
        value: proxy.example.com
    snowflake:proxyPassword:
        value: 'TODO: var.proxy_password'
    snowflake:proxyPort:
        value: 443
    snowflake:proxyProtocol:
        value: https
    snowflake:proxyUser:
        value: username
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
    snowflake:workloadIdentityEntraResource:
        value: workload_identity_entra_resource
    snowflake:workloadIdentityProvider:
        value: azure

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
		proxyPassword := cfg.Require("proxyPassword")
		oauthClientId := cfg.Require("oauthClientId")
		oauthClientSecret := cfg.Require("oauthClientSecret")
		oauthTokenRequestUrl := cfg.Require("oauthTokenRequestUrl")
		oauthAuthorizationUrl := cfg.Require("oauthAuthorizationUrl")
		oauthRedirectUri := cfg.Require("oauthRedirectUri")
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
        value: account_name
    snowflake:authenticator:
        value: snowflake
    snowflake:certRevocationCheckMode:
        value: ADVISORY
    snowflake:clientIp:
        value: 1.2.3.4
    snowflake:clientRequestMfaToken:
        value: true
    snowflake:clientStoreTemporaryCredential:
        value: true
    snowflake:clientTimeout:
        value: 40
    snowflake:crlAllowCertificatesWithoutCrlUrl:
        value: true
    snowflake:crlHttpClientTimeout:
        value: 30
    snowflake:crlInMemoryCacheDisabled:
        value: false
    snowflake:crlOnDiskCacheDisabled:
        value: true
    snowflake:disableConsoleLogin:
        value: true
    snowflake:disableOcspChecks:
        value: true
    snowflake:disableQueryContextCache:
        value: true
    snowflake:disableSamlUrlCheck:
        value: true
    snowflake:disableTelemetry:
        value: true
    snowflake:driverTracing:
        value: info
    snowflake:enableSingleUseRefreshTokens:
        value: true
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
    snowflake:logQueryParameters:
        value: false
    snowflake:logQueryText:
        value: false
    snowflake:loginTimeout:
        value: 10
    snowflake:maxRetryCount:
        value: 3
    snowflake:noProxy:
        value: localhost,snowflake.computing.com
    snowflake:oauthAuthorizationUrl:
        value: 'TODO: var.oauth_authorization_url'
    snowflake:oauthClientId:
        value: 'TODO: var.oauth_client_id'
    snowflake:oauthClientSecret:
        value: 'TODO: var.oauth_client_secret'
    snowflake:oauthRedirectUri:
        value: 'TODO: var.oauth_redirect_uri'
    snowflake:oauthScope:
        value: session:role:PUBLIC
    snowflake:oauthTokenRequestUrl:
        value: 'TODO: var.oauth_token_request_url'
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
    snowflake:proxyHost:
        value: proxy.example.com
    snowflake:proxyPassword:
        value: 'TODO: var.proxy_password'
    snowflake:proxyPort:
        value: 443
    snowflake:proxyProtocol:
        value: https
    snowflake:proxyUser:
        value: username
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
    snowflake:workloadIdentityEntraResource:
        value: workload_identity_entra_resource
    snowflake:workloadIdentityProvider:
        value: azure

```

```yaml
configuration:
  # Password for the proxy.
  proxyPassword:
    type: string
  # Client ID from the Okta application.
  oauthClientId:
    type: string
  # Client Secret from the Okta application.
  oauthClientSecret:
    type: string
  # Client Token Request URL from the Okta API Authorization Server.
  oauthTokenRequestUrl:
    type: string
  # Authorization URL for the Oauth flow.
  oauthAuthorizationUrl:
    type: string
  # Redirect URI for the Oauth flow.
  oauthRedirectUri:
    type: string
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
    snowflake:certRevocationCheckMode:
        value: ADVISORY
    snowflake:clientIp:
        value: 1.2.3.4
    snowflake:clientRequestMfaToken:
        value: true
    snowflake:clientStoreTemporaryCredential:
        value: true
    snowflake:clientTimeout:
        value: 40
    snowflake:crlAllowCertificatesWithoutCrlUrl:
        value: true
    snowflake:crlHttpClientTimeout:
        value: 30
    snowflake:crlInMemoryCacheDisabled:
        value: false
    snowflake:crlOnDiskCacheDisabled:
        value: true
    snowflake:disableConsoleLogin:
        value: true
    snowflake:disableOcspChecks:
        value: true
    snowflake:disableQueryContextCache:
        value: true
    snowflake:disableSamlUrlCheck:
        value: true
    snowflake:disableTelemetry:
        value: true
    snowflake:driverTracing:
        value: info
    snowflake:enableSingleUseRefreshTokens:
        value: true
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
    snowflake:logQueryParameters:
        value: false
    snowflake:logQueryText:
        value: false
    snowflake:loginTimeout:
        value: 10
    snowflake:maxRetryCount:
        value: 3
    snowflake:noProxy:
        value: localhost,snowflake.computing.com
    snowflake:oauthAuthorizationUrl:
        value: 'TODO: var.oauth_authorization_url'
    snowflake:oauthClientId:
        value: 'TODO: var.oauth_client_id'
    snowflake:oauthClientSecret:
        value: 'TODO: var.oauth_client_secret'
    snowflake:oauthRedirectUri:
        value: 'TODO: var.oauth_redirect_uri'
    snowflake:oauthScope:
        value: session:role:PUBLIC
    snowflake:oauthTokenRequestUrl:
        value: 'TODO: var.oauth_token_request_url'
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
    snowflake:proxyHost:
        value: proxy.example.com
    snowflake:proxyPassword:
        value: 'TODO: var.proxy_password'
    snowflake:proxyPort:
        value: 443
    snowflake:proxyProtocol:
        value: https
    snowflake:proxyUser:
        value: username
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
    snowflake:workloadIdentityEntraResource:
        value: workload_identity_entra_resource
    snowflake:workloadIdentityProvider:
        value: azure

```

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
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
        final var config = ctx.config();
        final var proxyPassword = config.require("proxyPassword");
        final var oauthClientId = config.require("oauthClientId");
        final var oauthClientSecret = config.require("oauthClientSecret");
        final var oauthTokenRequestUrl = config.require("oauthTokenRequestUrl");
        final var oauthAuthorizationUrl = config.require("oauthAuthorizationUrl");
        final var oauthRedirectUri = config.require("oauthRedirectUri");
    }
}
```

{{% /choosable %}}
{{% choosable language hcl %}}
```hcl
# Password for the proxy.
variable "proxyPassword" {
  type = string
}
# Client ID from the Okta application.
variable "oauthClientId" {
  type = string
}
# Client Secret from the Okta application.
variable "oauthClientSecret" {
  type = string
}
# Client Token Request URL from the Okta API Authorization Server.
variable "oauthTokenRequestUrl" {
  type = string
}
# Authorization URL for the Oauth flow.
variable "oauthAuthorizationUrl" {
  type = string
}
# Redirect URI for the Oauth flow.
variable "oauthRedirectUri" {
  type = string
}
```

{{% /choosable %}}
{{< /chooser >}}
## Proxy

Pulumi is plugin-based. It means that every plugin (provider) is responsible for making its own network requests.
Not all providers follow the same standardized ways, so familiarize yourself with proxy setting for each of the providers used within your module.

A few important pointers for setting the proxy connection:
- As far as we are aware, there are no official Pulumi docs regarding proxy, but there are some discussions on the official HashiCorp discussion forum (e.g. this one).
- Pulumi relies on Go default proxy setting (so it supports `HTTPS_PROXY`, `HTTP_PROXY`, `NO_PROXY`).
- The official Go driver for Snowflake, which is used in this provider, also supports the default Go environment variables (`HTTPS_PROXY`, `HTTP_PROXY`, `NO_PROXY`). Documented [here](https://pkg.go.dev/github.com/snowflakedb/gosnowflake#hdr-Proxy).
- The provider offers a separate config (through the provider configuration, dedicated environment variables, and the TOML file).
- The order of precedence is as follows:
  1. Provider configuration (following its own precedence).
  2. Standard environment variables (`HTTPS_PROXY`, `HTTP_PROXY`, `NO_PROXY`).

References:
- Pulumi discussion group example
- [Go driver documentation](https://pkg.go.dev/github.com/snowflakedb/gosnowflake#hdr-Proxy)
- [Go documentation](https://go.dev/src/vendor/golang.org/x/net/http/httpproxy/proxy.go)
## Sensitive values limitations

The provider marks fields containing access credentials and other such information as sensitive. This means that the values of these fields will not be logged.

There are some limitations to this mechanism:
- Sensitive values are stored as plaintext in the state file. This is a limitation of Pulumi itself (reference). You should take care to secure access to the state file.
- In Plugin SDK there is no possibility to mark sensitive values conditionally (reference). This means it is not possible to mark sensitive values based on other fields, like marking `body` based on the value of `secure` field in views, functions, and procedures. As a result, this field is not marked as sensitive. For such cases, we add disclaimers in the resource documentation.
- In Plugin SDK, there is no possibility to mark sensitive values in nested fields (reference). This means the nested fields, like these in `showOutput` and `describeOutput` cannot be sensitive.
  As a result, such nested fields are not marked as sensitive. For such cases, we add disclaimers in the resource documentation. Additionally, some fields are missing from `showOutput` and `describeOutput`. However, these fields are present in the resource's root, so they can still be referenced.
  The alternative solution we considered was setting the whole `showOutput` and `describeOutput` as sensitive. However, this solution could reduce the provider functionality and would require changes in user's configurations.
- Sensitive values cannot be used as `forEach` keys (reference). This means that if a resource attribute is marked as sensitive (e.g. `name` in `snowflake.User`), it cannot be used directly in a `forEach` expression. For example, iterating over a set of user names to create role grants will fail with `Sensitive values, or values derived from sensitive values, cannot be used as forEach arguments`. The workaround is to wrap the sensitive value with the `nonsensitive` function when you are certain the value is not actually sensitive in your context (e.g. a username that is not a secret):
  {{< chooser language "typescript,python,go,csharp,java,yaml,hcl" >}}
  {{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as snowflake from "@pulumi/snowflake";
import * as std from "@pulumi/std";

function notImplemented(message: string) {
    throw new Error(message);
}

export = async () => {
    const example = new snowflake.User("example", {name: "my_user"});
    const exampleRole = new snowflake.index.Role("example", {name: "my_role"});
    const exampleGrantAccountRole: snowflake.GrantAccountRole[] = [];
    for (const range of std.toset({
        input: [example.name].map(u => (notImplemented("nonsensitive(u)"))),
    }).result.map((v, k) => ({key: k, value: v}))) {
        exampleGrantAccountRole.push(new snowflake.GrantAccountRole(`example-${range.key}`, {
            roleName: exampleRole.name,
            userName: range.value,
        }));
    }
}
```

{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
from typing import Any
import pulumi_snowflake as snowflake
import pulumi_std as std


def not_implemented(msg):
    raise NotImplementedError(msg)

example = snowflake.User("example", name="my_user")
example_role = snowflake.Role("example", name=my_role)
example_grant_account_role: list[snowflake.GrantAccountRole] = []
for example_grant_account_role_range in [{"key": k, "value": v} for [k, v] in enumerate(std.toset(input=[not_implemented(nonsensitive(u)) for u in [example.name]]).result)]:
    example_grant_account_role.append(snowflake.GrantAccountRole(f"example-{example_grant_account_role_range['key']}",
        role_name=example_role["name"],
        user_name=example_grant_account_role_range["value"]))
```

{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Pulumi;
using Snowflake = Pulumi.Snowflake;
using Std = Pulumi.Std;


object NotImplemented(string errorMessage)
{
    throw new System.NotImplementedException(errorMessage);
}

return await Deployment.RunAsync(async() =>
{
    var example = new Snowflake.User("example", new()
    {
        Name = "my_user",
    });

    var exampleRole = new Snowflake.Role("example", new()
    {
        Name = "my_role",
    });

    var exampleGrantAccountRole = new List<Snowflake.GrantAccountRole>();
    foreach (var range in )
    {
        exampleGrantAccountRole.Add(new Snowflake.GrantAccountRole($"example-{range.Key}", new()
        {
            RoleName = exampleRole.Name,
            UserName = range.Value,
        }));
    }
});

```

{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-snowflake/sdk/v2/go/snowflake"
	"github.com/pulumi/pulumi-std/sdk/go/std"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func notImplemented(message string) pulumi.AnyOutput {
	panic(message)
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := snowflake.NewUser(ctx, "example", &snowflake.UserArgs{
			Name: pulumi.String("my_user"),
		})
		if err != nil {
			return err
		}
		exampleRole, err := snowflake.NewRole(ctx, "example", &snowflake.RoleArgs{
			Name: "my_role",
		})
		if err != nil {
			return err
		}
		var exampleGrantAccountRole []*snowflake.GrantAccountRole
		for key0, val0 := range []interface{}(std.Toset(ctx, &std.TosetArgs{
			Input: "TODO: For expression",
		}, nil).Result) {
			__res, err := snowflake.NewGrantAccountRole(ctx, fmt.Sprintf("example-%v", key0), &snowflake.GrantAccountRoleArgs{
				RoleName: exampleRole.Name,
				UserName: pulumi.Any(val0),
			})
			if err != nil {
				return err
			}
			exampleGrantAccountRole = append(exampleGrantAccountRole, __res)
		}
		return nil
	})
}
```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
resources:
  example:
    type: snowflake:User
    properties:
      name: my_user
  exampleRole:
    type: snowflake:Role
    name: example
    properties:
      name: my_role
  exampleGrantAccountRole:
    type: snowflake:GrantAccountRole
    name: example
    properties:
      roleName: ${exampleRole.name}
      userName: ${range.value}
    options: {}
```

{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.snowflake.User;
import com.pulumi.snowflake.UserArgs;
import com.pulumi.snowflake.Role;
import com.pulumi.snowflake.RoleArgs;
import com.pulumi.snowflake.GrantAccountRole;
import com.pulumi.snowflake.GrantAccountRoleArgs;
import com.pulumi.codegen.internal.KeyedValue;
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
        var example = new User("example", UserArgs.builder()
            .name("my_user")
            .build());

        var exampleRole = new Role("exampleRole", RoleArgs.builder()
            .name("my_role")
            .build());

        for (var range : KeyedValue.of(com.pulumi.std.StdFunctions(com.pulumi.std.inputs.TosetArgs.builder()
            .input("TODO: ForExpression")
            .build()).result())) {
            new GrantAccountRole("exampleGrantAccountRole-" + range.key(), GrantAccountRoleArgs.builder()
                .roleName(exampleRole.name())
                .userName(range.value())
                .build());
        }

    }
}
```

{{% /choosable %}}
{{% choosable language hcl %}}
```hcl
pulumi {
  required_providers {
    snowflake = {
      source = "pulumi/snowflake"
    }
    std = {
      source = "pulumi/std"
    }
  }
}

resource "snowflake_user" "example" {
  name = "my_user"
}
resource "snowflake_role" "example" {
  name = "my_role"
}
resource "snowflake_grantaccountrole" "example" {
  for_each  = toset([for u in [snowflake_user.example.name] : nonsensitive(u)])
  role_name = snowflake_role.example.name
  user_name = each.value
}
```

{{% /choosable %}}
{{< /chooser >}}
``
Note: use `nonsensitive` only when you are confident the value does not need to be protected. Misuse can inadvertently expose secrets in logs or plan output.

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
{{< chooser language "typescript,python,go,csharp,java,yaml,hcl" >}}
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
        var test = new Execute("test", ExecuteArgs.builder()
            .execute("CREATE DATABASE ABC")
            .revert("DROP DATABASE ABC")
            .query("SHOW DATABASES LIKE '%ABC%'")
            .build());

    }
}
```

{{% /choosable %}}
{{% choosable language hcl %}}
```hcl
pulumi {
  required_providers {
    snowflake = {
      source = "pulumi/snowflake"
    }
  }
}

resource "snowflake_execute" "test" {
  execute = "CREATE DATABASE ABC"
  revert  = "DROP DATABASE ABC"
  query   = "SHOW DATABASES LIKE '%ABC%'"
}
```

{{% /choosable %}}
{{< /chooser >}}

> Note: Timeouts can be also set at driver's level (see [driver documentation](https://pkg.go.dev/github.com/snowflakedb/gosnowflake)). These timeouts are independent. We recommend tweaking the timeouts on Pulumi level first.
## General provider rules

> Note: This section is in a `work in progress` state and will be updated over time.

In this section, we describe general rules that apply to multiple resources and functions in the provider.
This may help you understand the provider behavior when you are getting started with it.

However, getting familiar with existing guides (`Guides` section on the left),
resource-specific documentation, and Snowflake-specific documentation for a given object is still recommended.

Here's a list of general rules:
- All fields representing object identifiers (e.g., allowedNetworkRuleList in `snowflake.NetworkPolicy`) or parts of them (e.g., `database`, `schema`, and `name` in the snowflake.NetworkRule resource) are case-sensitive. This is true for all stable resources (there may be some exceptions in the preview ones; especially older ones). However, you can make all identifiers case-insensitive by enabling [QUOTED_IDENTIFIERS_IGNORE_CASE](https://docs.snowflake.com/en/sql-reference/parameters#quoted-identifiers-ignore-case), but be aware with the issues you may have when using it.
## A list of preview and stable resources and functions

The provider supports a number of Snowflake features. Within the provider, some features are stable, while others are in preview
(stability of the feature in the provider is not connected to the stability of the feature in Snowflake).

Preview features are **experimental** and may introduce **breaking changes**, even between non-major versions of the provider.
Eventually, every preview resource will be promoted to stable, but the timeline for each feature is not defined (you can find more details on the current/future plans in our roadmap).
New resources will be introduced as preview ones and promoted over time to stable as we gain more confidence in their stability.

Preview features are disabled by default and should be used with caution.
To use them, add the relevant feature name to the `previewFeaturesEnabled` field in the provider configuration.

<!-- Section of stable resources -->
### Currently stable resources

- snowflake.Account
- snowflake.AccountParameter
- snowflake.AccountRole
- snowflake.AccountSessionPolicyAttachment
- snowflake.ApiAuthenticationIntegrationWithAuthorizationCodeGrant
- snowflake.ApiAuthenticationIntegrationWithClientCredentials
- snowflake.ApiAuthenticationIntegrationWithJwtBearer
- snowflake.AuthenticationPolicy
- snowflake.CatalogIntegrationAwsGlue
- snowflake.CatalogIntegrationIcebergRest
- snowflake.CatalogIntegrationObjectStorage
- snowflake.CatalogIntegrationOpenCatalog
- snowflake.ComputePool
- snowflake.CurrentAccount
- snowflake.CurrentOrganizationAccount
- snowflake.Database
- snowflake.DatabaseRole
- snowflake.Execute
- snowflake.ExternalOauthIntegration
- snowflake.ExternalVolume
- snowflake.GitRepository
- snowflake.GrantAccountRole
- snowflake.GrantApplicationRole
- snowflake.GrantDatabaseRole
- snowflake.GrantOwnership
- snowflake.GrantPrivilegesToAccountRole
- snowflake.GrantPrivilegesToDatabaseRole
- snowflake.GrantPrivilegesToShare
- snowflake.ImageRepository
- snowflake.LegacyServiceUser
- snowflake.Listing
- snowflake.MaskingPolicy
- snowflake.NetworkPolicy
- snowflake.NetworkRule
- snowflake.OauthIntegrationForCustomClients
- snowflake.OauthIntegrationForPartnerApplications
- snowflake.PasswordPolicy
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
- snowflake.Service
- snowflake.ServiceUser
- snowflake.SessionPolicy
- snowflake.SharedDatabase
- snowflake.StageExternalAzure
- snowflake.StageExternalGcs
- snowflake.StageExternalS3
- snowflake.StageExternalS3Compatible
- snowflake.StageInternal
- snowflake.StorageIntegrationAws
- snowflake.StorageIntegrationAzure
- snowflake.StorageIntegrationGcs
- snowflake.StreamOnDirectoryTable
- snowflake.StreamOnExternalTable
- snowflake.StreamOnTable
- snowflake.StreamOnView
- snowflake.Streamlit
- snowflake.Tag
- snowflake.TagAssociation
- snowflake.Task
- snowflake.User
- snowflake.UserProgrammaticAccessToken
- snowflake.UserSessionPolicyAttachment
- snowflake.View
- snowflake.Warehouse

<!-- Section of stable functions -->
### Currently stable functions

- snowflake.getAccountRoles
- snowflake.getAccounts
- snowflake.getAuthenticationPolicies
- snowflake.getCatalogIntegrations
- snowflake.getComputePools
- snowflake.getConnections
- snowflake.getDatabaseRoles
- snowflake.getDatabases
- snowflake.getExternalVolumes
- snowflake.getGitRepositories
- snowflake.getGrants
- snowflake.getImageRepositories
- snowflake.getMaskingPolicies
- snowflake.getNetworkPolicies
- snowflake.getNetworkRules
- snowflake.getPasswordPolicies
- snowflake.getResourceMonitors
- snowflake.getRowAccessPolicies
- snowflake.getSchemas
- snowflake.getSecrets
- snowflake.getSecurityIntegrations
- snowflake.getServices
- snowflake.getSessionPolicies
- snowflake.getStorageIntegrations
- snowflake.getStreamlits
- snowflake.getStreams
- snowflake.getTags
- snowflake.getTasks
- snowflake.getUserProgrammaticAccessTokens
- snowflake.getUsers
- snowflake.getViews
- snowflake.getWarehouses

<!-- Section of preview resources -->
### Currently preview resources

- snowflake.AccountAuthenticationPolicyAttachment
- snowflake.AccountPasswordPolicyAttachment
- snowflake.Alert
- snowflake.ApiIntegration
- snowflake.ApiIntegrationAmazonApiGateway
- snowflake.ApiIntegrationAzureApiManagement
- snowflake.ApiIntegrationExternalMcpDynamicClient
- snowflake.ApiIntegrationExternalMcpOauth2
- snowflake.ApiIntegrationGitRepositoryGithubApp
- snowflake.ApiIntegrationGitRepositoryOauth2
- snowflake.ApiIntegrationGitRepositoryPrivateLink
- snowflake.ApiIntegrationGitRepositoryToken
- snowflake.ApiIntegrationGoogleCloudApiGateway
- snowflake.CortexAgent
- snowflake.CortexSearchService
- snowflake.DynamicTable
- snowflake.EmailNotificationIntegration
- snowflake.ExternalAccessIntegration
- snowflake.ExternalFunction
- snowflake.ExternalTable
- snowflake.FailoverGroup
- snowflake.FileFormat
- snowflake.FileFormatAvro
- snowflake.FileFormatCsv
- snowflake.FileFormatJson
- snowflake.FileFormatOrc
- snowflake.FileFormatParquet
- snowflake.FileFormatXml
- snowflake.FunctionJava
- snowflake.FunctionJavascript
- snowflake.FunctionPython
- snowflake.FunctionScala
- snowflake.FunctionSql
- snowflake.HybridTable
- snowflake.IcebergTable
- snowflake.IcebergTableFromAwsGlue
- snowflake.IcebergTableFromDeltaFiles
- snowflake.IcebergTableFromFiles
- snowflake.IcebergTableFromRest
- snowflake.JobService
- snowflake.ManagedAccount
- snowflake.MaterializedView
- snowflake.McpServer
- snowflake.NetworkPolicyAttachment
- snowflake.Notebook
- snowflake.NotificationIntegration
- snowflake.ObjectParameter
- snowflake.Pipe
- snowflake.PostgresInstance
- snowflake.ProcedureJava
- snowflake.ProcedureJavascript
- snowflake.ProcedurePython
- snowflake.ProcedureScala
- snowflake.ProcedureSql
- snowflake.SemanticView
- snowflake.Sequence
- snowflake.Share
- snowflake.Stage
- snowflake.StorageIntegration
- snowflake.StorageLifecyclePolicy
- snowflake.Table
- snowflake.TableColumnMaskingPolicyApplication
- snowflake.TableConstraint
- snowflake.TableStorageLifecyclePolicyAttachment
- snowflake.UserAuthenticationPolicyAttachment
- snowflake.UserPasswordPolicyAttachment
- snowflake.UserPublicKeys
- snowflake.WarehouseAdaptive
- snowflake.WarehouseInteractive

<!-- Section of preview functions -->
### Currently preview functions

- snowflake.getAlerts
- snowflake.getApiIntegrations
- snowflake.getCortexAgents
- snowflake.getCortexSearchServices
- snowflake.CurrentAccount
- snowflake.getCurrentRole
- snowflake.Database
- snowflake.DatabaseRole
- snowflake.getDynamicTables
- snowflake.getExternalAccessIntegrations
- snowflake.getExternalFunctions
- snowflake.getExternalTables
- snowflake.getFailoverGroups
- snowflake.getFileFormats
- snowflake.getFunctions
- snowflake.getIcebergTables
- snowflake.getListings
- snowflake.getMaterializedViews
- snowflake.getMcpServers
- snowflake.getNotebooks
- snowflake.getParameters
- snowflake.getPipes
- snowflake.getProcedures
- snowflake.getSemanticViews
- snowflake.getSequences
- snowflake.getShares
- snowflake.getStages
- snowflake.getStorageLifecyclePolicies
- snowflake.getSystemGenerateScimAccessToken
- snowflake.getSystemGetAwsSnsIamPolicy
- snowflake.getSystemGetPrivateLinkConfig
- snowflake.getSystemGetSnowflakePlatformInfo
- snowflake.getTables

<!-- Section of deprecated resources -->
### Currently deprecated resources

- snowflake.ApiIntegration - use snowflake_api_integration_amazon_api_gateway, snowflake_api_integration_azure_api_management, snowflake_api_integration_google_cloud_api_gateway, snowflake_api_integration_git_repository_github_app, snowflake_api_integration_git_repository_oauth2, snowflake_api_integration_git_repository_token, snowflake_api_integration_git_repository_private_link, snowflake_api_integration_external_mcp_oauth2, snowflake.ApiIntegrationExternalMcpDynamicClient instead
- snowflake.FileFormat - use snowflake_file_format_csv, snowflake_file_format_json, snowflake_file_format_avro, snowflake_file_format_orc, snowflake_file_format_parquet, snowflake.FileFormatXml instead
- snowflake.Stage - use snowflake_stage_internal, snowflake_stage_external_s3, snowflake_stage_external_s3_compatible, snowflake_stage_external_gcs, snowflake.StageExternalAzure instead
- snowflake.StorageIntegration - use snowflake_storage_integration_aws, snowflake_storage_integration_azure, snowflake.StorageIntegrationGcs instead

<!-- Section of deprecated functions -->
## Experimental features

Experiments alter the provider behavior.
Similarly to preview features, they are not yet stable features of the provider.
Enabling the given experiment is still considered a preview feature, even when applied to the stable resource.
If the given experiment is successful, it can be considered an addition in the future provider versions.
### Active experiments

The following experiments are currently active. Depending on the feedback, we may decide to include them as default behavior/stable feature of the provider in the future.

To share feedback please reach out to us through your Snowflake account manager.
#### WAREHOUSE_SHOW_IMPROVED_PERFORMANCE
It's meant to improve the performance for accounts with many warehouses.

When enabled, it uses a slightly different SHOW query to read warehouse details (`SHOW WAREHOUSES LIKE '<identifier>' STARTS WITH '<identifier>' LIMIT 1`).

This feature is enabled by default on the Snowflake side.
#### GRANTS_STRICT_PRIVILEGE_MANAGEMENT
The new `strictPrivilegeManagement` flag was added to the `snowflake.GrantPrivilegesToAccountRole` resource.

It has similar behavior to the `enableMultipleGrants` flag present in the old grant resources, and it makes the resource able to detect external changes for privileges other than those present in the configuration which can make the `snowflake.GrantPrivilegesToAccountRole` resource a central point of knowledge privilege management for a given object and role.

Read more in our strict privilege management guide.

This feature works independently of the `GRANTS_IMPORT_VALIDATION` flag.
#### PARAMETERS_IGNORE_VALUE_CHANGES_IF_NOT_ON_OBJECT_LEVEL
Currently, not setting the parameter value on the object level can unnecessarily react to external changes to this parameter's value on the higher levels (e.g. not setting `dataRetentionTimeInDays` on `snowflake.Schema` can result in non-empty plan when the parameter value changes on the database/account level).

When enabled, the provider ignores changes to the parameter value happening on the higher hierarchy levels.
#### PARAMETERS_REDUCED_OUTPUT
Currently, the `parameters` field in various resources contains a verbatim output for the `SHOW PARAMETERS IN <object>` command. One of the fields contained in the output is the `description`. It does not change and is repeated for all objects containing the given parameter. It leads to an excessive output (check e.g., #3118).

To mitigate the problem, we are adding this option to reduce the output to only `value` and `level` fields, which should significantly reduce the state size. **Note**: it's also affecting the `parameters` output for functions.

We considered the option to remove the `parameters` output completely, however, we plan to change the external change logic detection to use it (to make it consistent with other attributes using `showOutput` and because we won't be able to implement the current logic when switching to the Pulumi Plugin Framework) and it still allows referencing the parameter value/level from other parts of the configuration.
#### USER_ENABLE_DEFAULT_WORKLOAD_IDENTITY
The new `defaultWorkloadIdentityFederation` field was added to the `snowflake.LegacyServiceUser` and `snowflake.ServiceUser` resources. This field allows for managing WIFs. Due to feature complexity, it requires enabling this experiment.

Read more in our migration guide.
#### GRANTS_IMPORT_VALIDATION
Enables import validation for the `snowflake.GrantPrivilegesToAccountRole` resource.

When enabled, importing a grant resource with a fixed set of privileges (`privileges` field) will validate that the specified privileges actually exist in Snowflake with the correct `withGrantOption` setting, and error immediately if they don't match.

This feature works independently of the `GRANTS_STRICT_PRIVILEGE_MANAGEMENT` flag.
#### TAGS_ALLOW_EMPTY_ALLOWED_VALUES
Enables behavior changes for the `allowedValues` field in the `snowflake.Tag` resource.

When enabled, the three possible states in Snowflake for allowed values will be supported: `nil` (any value is allowed; whenever `allowedValues` are empty), `empty` (no value is allowed; handled by the `noAllowedValues` field), and `set` (all values defined in `allowedValues` are allowed).

Otherwise, the `noAllowedValues` field will be ignored (explicit changes will cause updates, but without any effect) and the `allowedValues` field will follow the old behavior: `nil` (any value is allowed; only available whenever tag resource is created without `allowedValues`), `empty` (no value is allowed; always set when updating from filled `allowedValues` set to empty one or completely removed from config), `set` (all values defined in `allowedValues` are allowed).
#### IMPORT_BOOLEAN_DEFAULT
Changes import behavior for boolean fields using the special `"default"` value.

When enabled, boolean fields using the special `"default"` value are set to `"default"` during import instead of the actual Snowflake value (e.g., `"false"`). This prevents unavoidable diffs on every plan after import.

Note: this is supported on all stage resources (`snowflake.StageExternalS3`, `snowflake.StageExternalAzure`, `snowflake.StageExternalGcs`, `snowflake.StageExternalS3Compatible`, and `snowflake.StageInternal`) and stream resources (`snowflake.StreamOnTable` and `snowflake.StreamOnView`).
#### GRANTS_SAFE_DESTROY
When enabled, grant destroy operations silently succeed when the underlying Snowflake object (or its dependencies) no longer exists.

Currently supported by: `snowflake.GrantPrivilegesToAccountRole`, `snowflake.GrantPrivilegesToDatabaseRole`, `snowflake.GrantPrivilegesToShare`, `snowflake.GrantAccountRole`, `snowflake.GrantDatabaseRole`, `snowflake.GrantApplicationRole`, `snowflake.GrantOwnership`.

This prevents errors when, for example, a warehouse or role is deleted externally and the corresponding grant resource is later removed from the Pulumi configuration.

Without this experiment, destroying such resources fails with `does not exist or not authorized`.
#### TAG_ASSOCIATION_SAFE_DESTROY
When enabled, tag association destroy operations silently succeed when the tagged object (or its parent hierarchy) no longer exists.

Currently supported by: `snowflake.TagAssociation`.

This prevents errors when, for example, a table or schema is deleted externally and the corresponding tag association resource is later removed from the Pulumi configuration.

Without this experiment, destroying such resources fails with `does not exist or not authorized`.
#### GRANT_ACCOUNT_ROLE_SHOW_CACHING
Enables per-plan in-memory caching of `SHOW GRANTS OF ROLE` results for the `snowflake.GrantAccountRole` resource.

Without caching, every resource instance issues an independent `SHOW GRANTS OF ROLE <name>` call during Read. In configurations with many grants sharing the same role, this results in N identical round-trips returning the same full result set — only 1 is needed.

When enabled, the first Read for a given role fetches and caches the result; subsequent Reads in the same plan reuse it. The cache is invalidated on Create and Delete so mutations within a single apply remain visible to subsequent Reads.

Additionally, the trailing Read at the end of Create is skipped (this resource has no computed or server-default fields to populate), removing a redundant `SHOW GRANTS OF ROLE` call per grant during apply.

Intended for large configurations (thousands of `snowflake.GrantAccountRole` resources) where plan and apply time is dominated by redundant `SHOW GRANTS OF ROLE` calls.
#### ACCOUNT_ROLE_SHOW_CACHING
When enabled, the result of looking up an account role by identifier (`SHOW ROLES LIKE '<name>'`, via the underlying `ShowByID`/`ShowByIDSafely` calls) is cached in memory for the duration of a single plan or apply cycle.

Currently supported by: `snowflake.AccountRole`, `snowflake.GrantApplicationRole`, `snowflake.GrantPrivilegesToAccountRole`.

Without caching, every lookup of a given role — whether it's the role's own `snowflake.AccountRole` Read, or an existence check performed by a grant resource before granting to/from it — issues an independent round trip, even when many resource instances reference the same role. When enabled, the first lookup for a given role fetches and caches the result; subsequent lookups in the same plan reuse it. The cache is invalidated on `snowflake.AccountRole` Update (rename or comment change) and Delete, since only that resource can change what a cached lookup would return.

This is a separate flag from `GRANT_ACCOUNT_ROLE_SHOW_CACHING`: enabling this does not enable caching for `snowflake.GrantAccountRole`'s `SHOW GRANTS OF ROLE` calls, and vice versa. Both can be enabled together.

Intended for large configurations (thousands of role or grant resources) where plan and apply time is dominated by redundant role lookups.
#### GRANTS_SHOW_CACHING
When enabled, `SHOW GRANTS` results are cached in memory for the duration of a single plan or apply cycle, so multiple resource instances resolving to the same underlying SHOW statement share one round-trip instead of each issuing their own.

Currently supported by: `snowflake.GrantPrivilegesToAccountRole`, `snowflake.GrantOwnership`.

Without caching, every resource instance issues an independent `SHOW GRANTS ON <object>` / `SHOW FUTURE GRANTS IN <container>` call during Read. In configurations with many grants resolving to the same underlying SHOW statement (e.g. many privilege grants on the same schema, or many future-grant roles on the same database), this results in N identical round-trips returning the same full result set — only 1 is needed.

The first Read for a given SHOW statement fetches and caches the result; subsequent Reads in the same plan reuse it. The cache is invalidated on Create, Update, and Delete of the resources listed above so mutations within a single apply remain visible to subsequent Reads.

This is a separate flag from `GRANT_ACCOUNT_ROLE_SHOW_CACHING`: enabling this does not enable caching for `snowflake.GrantAccountRole`, and vice versa. Both can be enabled together.

Intended for large configurations (thousands of grant resources) where plan and apply time is dominated by redundant `SHOW GRANTS` calls.
#### GRANT_ACCOUNT_ROLE_SAFE_PUBLIC_ROLE
When enabled, `snowflake.GrantAccountRole` treats granting the PUBLIC role as a silent no-op instead of producing an error.

Snowflake implicitly grants PUBLIC to every role and user (see [Snowflake documentation](https://docs.snowflake.com/en/user-guide/security-access-control-overview#system-defined-roles)), so an explicit `GRANT ROLE PUBLIC` is always a no-op at the SQL level. However, the provider's Read function cannot find the explicit grant via `SHOW GRANTS` and clears the state, causing an inconsistent-result error.

With this experiment, Create, Read, and Delete all treat PUBLIC role grants as permanent fixtures that require no actual SQL.
#### HIERARCHY_RENAMES
When enabled, allows in-place handling of hierarchy renames and moves for supported resources.

Currently supported by: `snowflake.Schema`, `snowflake.Table`.

Without this experiment, changing the `database` field on `snowflake.Schema` or the `database`/`schema` fields on `snowflake.Table` forces resource recreation. With this experiment, the provider detects whether the parent was renamed or the object should be moved, and handles it without recreation.

For more information, see the object renaming guide.
#### INHERITED_GRANTS
Enables the `inherited` block in the `onAccountObject`, `onSchema`, and `onSchemaObject` blocks of the `snowflake.GrantPrivilegesToAccountRole` resource, and in the `onSchema` and `onSchemaObject` blocks of the `snowflake.GrantPrivilegesToDatabaseRole` resource.

Without this experiment, using an `inherited` block results in an error.
#### OBJECT_PARAMETER_UNSET_ON_DELETE
Changes the delete behavior of the `snowflake.ObjectParameter` resource to use `ALTER <OBJECT_TYPE> <identifier> UNSET <PARAMETER>` instead of resetting the parameter to its default value.

Without this experiment, deleting the resource fetches the parameter's default value and explicitly sets it back, which is fragile and doesn't truly remove the object-level override.

When enabled, the parameter is properly unset, allowing the inherited value from the higher hierarchy level to take effect.
#### AUTHENTICATOR_EXPLICIT_ONLY
Removes implicit authenticator derivation from other provider configuration fields.

Without this experiment, the provider automatically sets the authenticator to `OAUTH` when the `token` or `tokenAccessor` field is configured, even if `authenticator` is not explicitly set. This implicit behavior can be confusing and will be removed in v3.

When enabled, the `authenticator` field must be set explicitly in the provider configuration or TOML profile. The `SNOWFLAKE` default (when no authenticator is configured anywhere) is preserved.
#### PROVIDER_CONFIGURATION_ACCOUNT_FALLBACK
Re-introduces the `account` field as a fallback for `organizationName` and `accountName` in both the provider configuration and TOML profiles.

When enabled, you can set `account` instead of setting `organizationName` and `accountName` separately. The field accepts both the `org-name` format (e.g. `"myorg-myaccount"`) and an account locator (e.g. `"xy12345"`). If both `organizationName` and `accountName` are set, they take precedence over `account`. The `SNOWFLAKE_ACCOUNT` environment variable is used as the `account` value only when this experiment is enabled.

Without this experiment, setting the `account` field in the provider configuration or in a TOML profile results in an error directing you to enable this experiment. A value coming from the `SNOWFLAKE_ACCOUNT` environment variable is ignored with a warning instead, because this experiment will be enabled by default in v3 and the variable will be taken into account from that version on.