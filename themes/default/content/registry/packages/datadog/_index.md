---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-datadog/v5.21.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi/pulumi-datadog/blob/v5.21.0/docs/_index.md
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Datadog Provider
meta_desc: Provides an overview on how to configure the Pulumi Datadog provider.
layout: package
---

## Installation

The Datadog provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/datadog`](https://www.npmjs.com/package/@pulumi/datadog)
* Python: [`pulumi-datadog`](https://pypi.org/project/pulumi-datadog/)
* Go: [`github.com/pulumi/pulumi-datadog/sdk/v5/go/datadog`](https://github.com/pulumi/pulumi-datadog)
* .NET: [`Pulumi.Datadog`](https://www.nuget.org/packages/Pulumi.Datadog)
* Java: [`com.pulumi/datadog`](https://central.sonatype.com/artifact/com.pulumi/datadog)

## Overview

The [Datadog](https://www.datadoghq.com) provider is used to interact with the resources supported by Datadog. The provider needs to be configured with the proper credentials before it can be used. 

Use the navigation to the left to read about the available resources.
## AWS Workload Identity Federation

Organizations configured for AWS Workload Identity Federation can authenticate
without Datadog API or application keys and without passing AWS credentials to
Pulumi:

The provider uses the AWS SDK default credential chain. This includes standard
sources such as environment credentials, shared configuration and credential
files, IAM Identity Center (SSO), web identity, ECS task roles, EC2 instance
roles, and assume-role profiles. For example, after logging in locally with
`aws sso login --profile sandbox`, run Pulumi with `AWS_PROFILE=sandbox`.

Explicit `awsAccessKeyId`, `awsSecretAccessKey`, and `awsSessionToken`
values remain supported but are not required. Workload Identity Federation must
be enabled for the Datadog organization, and its identity or intake mapping must
match the AWS identity resolved by the credential chain.
## Pulumi Dynamic Provider Credentials

For HCP Pulumi, the provider discovers Pulumi-issued workload identity tokens and exchanges them through Datadog Workload Identity Federation when authentication is not explicitly configured in the provider configuration. Use `workloadIdentityTokenTag` to select a distinct token for each provider configuration. Configure the WIF persona mapping and workspace audience as described in the Dynamic Provider Credentials guide.
## Authenticate with pup

For local development, the provider can use the OAuth session from the [pup](https://github.com/DataDog/pup) CLI, so you don't need to create or store an API key, application key, or access token.

**Prerequisites**: Install pup, make sure it's on your `PATH`, and run `pup auth login`.

To turn it on, set `usePup = true` in the provider configuration, or set the `DD_USE_PUP` environment variable to `true`. Also set `orgUuid`, which is required with pup, so every token must belong to that organization:

The provider gets a token by running `pup auth token`, and requests a new one when it expires, so long-running applies stay authenticated.

The provider uses pup only when no other credentials are configured. API and application keys, `bearerToken`, cloud provider authentication, and HCP Pulumi dynamic provider credentials all take precedence, whether they're set in the provider configuration or in environment variables.

- **Changing pup accounts during a run**: If pup is logged in as a different user partway through a run, the provider fails at the next token request instead of continuing as that user. Because `orgUuid` is required, a token for a different organization also stops the run.
- **Site matching**: Set `apiUrl` to the same Datadog site that pup is logged in to.
- **Parallel runs**: Each Pulumi run calls `pup auth token` on its own. If several runs try to refresh an expired session at the same time, the provider retries a failed `pup auth token` call up to three times before returning an error.
## Configuration Reference

- `apiKey` (String, Sensitive) Datadog API key. Required when `validate` is true unless authentication uses AWS WIF, a bearer token, or Pulumi Dynamic Provider Credentials. This can also be set via the `DD_API_KEY` environment variable.
- `apiUrl` (String) The API URL. This can also be set via the DD_HOST environment variable, and defaults to `https://api.datadoghq.com`. Note that this URL must not end with the `/api/` path. For example, `https://api.datadoghq.com/` is a correct value, while `https://api.datadoghq.com/api/` is not. And if you're working with "EU" version of Datadog, use `https://api.datadoghq.eu/`. Other Datadog region examples: `https://api.us5.datadoghq.com/`, `https://api.us3.datadoghq.com/` and `https://api.ddog-gov.com/`. See <https://docs.datadoghq.com/getting_started/site/> for all available regions.
- `appKey` (String, Sensitive) Datadog application key. Required when `validate` is true unless authentication uses AWS WIF, a bearer token, or Pulumi Dynamic Provider Credentials. This can also be set via the `DD_APP_KEY` environment variable.
- `awsAccessKeyId` (String, Sensitive) An optional explicit AWS access key ID used for cloud-provider-based authentication. This can also be set using the `AWS_ACCESS_KEY_ID` environment variable. When explicit credentials are omitted, the AWS SDK default credential chain is used.
- `awsSecretAccessKey` (String, Sensitive) An optional explicit AWS secret access key used for cloud-provider-based authentication. This can also be set using the `AWS_SECRET_ACCESS_KEY` environment variable. It must be provided together with `awsAccessKeyId`. Omit both fields to use the AWS SDK default credential chain.
- `awsSessionToken` (String, Sensitive) An optional AWS session token used with explicit temporary credentials for cloud-provider-based authentication. This can also be set using the `AWS_SESSION_TOKEN` environment variable.
- `bearerToken` (String, Sensitive) Datadog credential sent in the `Authorization: Bearer <token>` header. Accepts personal access tokens (`ddpat_*`) and service-account access tokens (`ddsat_*`). When set, the provider authenticates with `Authorization: Bearer <token>` instead of the `DD-API-KEY` / `DD-APPLICATION-KEY` headers. This can also be set via the `DD_BEARER_TOKEN` or `DATADOG_BEARER_TOKEN` environment variable.
- `cloudProviderRegion` (String) The cloud provider region specifier; used for cloud-provider-based authentication. For example, `us-east-1` for AWS. When omitted for AWS, the region is loaded from the AWS SDK default configuration chain.
- `cloudProviderType` (String) Specifies the cloud provider used for cloud-provider-based authentication, enabling keyless access without API or app keys. Only [`aws`] is supported. This can also be set using the `DD_CLOUD_PROVIDER_TYPE` environment variable. This feature is in Preview. If you'd like to enable it for your organization, contact [support](https://docs.datadoghq.com/help/).
- `defaultTags` (Block List, Max: 1) [Experimental - Action Connections, Logs Indexes, Logs Pipelines, Monitors, Security Monitoring Rules, and Service Level Objectives only] Configuration block containing settings to apply default resource tags across all resources. (see below for nested schema)
- `httpClientRetryBackoffBase` (Number) The HTTP request retry back off base. Defaults to 2.
- `httpClientRetryBackoffMultiplier` (Number) The HTTP request retry back off multiplier. Defaults to 2.
- `httpClientRetryEnabled` (String) Enables request retries on HTTP status codes 429 and 5xx. Valid values are [`true`, `false`]. Defaults to `true`.
- `httpClientRetryJitter` (Number) The maximum random delay added to each HTTP request retry. Defaults to 0 seconds.
- `httpClientRetryMaxRetries` (Number) The HTTP request maximum retry number. Defaults to 3.
- `httpClientRetryTimeout` (Number) The HTTP request retry timeout period. Defaults to 60 seconds.
- `ignoreTagKeys` (Set of String) [Experimental - Monitors and Service Level Objectives only] Tag keys whose drift Pulumi should ignore across all resources that support `ignoreTagKeys`. A resource's own `ignoreTagKeys` is merged with this list for that resource. Any `:value` suffix is ignored.
- `orgUuid` (String) The organization UUID; required for AWS WIF and for `usePup`. Optional for Pulumi Dynamic Provider Credentials, but when configured it must match the organization UUID in the token audience. This can also be set using the `DD_ORG_UUID` environment variable. See the [Datadog API documentation](https://docs.datadoghq.com/api/v1/organizations/) for more information.
- `usePup` (Boolean) Get a bearer token from the pup CLI (pup auth token) after you run pup auth login, so you don't need to store an API key, application key, or access token. Used only when api_key, app_key, bearer_token, cloud provider authentication, and HCP Pulumi dynamic provider credentials are all absent. The provider requests a new token when the current one expires, and fails if pup is then logged in as a different user. Requires org_uuid, and every token must belong to that organization. You can also set this with the `DD_USE_PUP` environment variable.
- `validate` (String) Enables credential validation during provider initialization. Valid values are [`true`, `false`]. Default is true. When false, initial credential validation is skipped.
- `workloadIdentityTokenTag` (String) Selects the HCP Pulumi workload identity token tag for this provider configuration, for example DATADOG_ORG_A selects TFC_WORKLOAD_IDENTITY_TOKEN_DATADOG_ORG_A. When set, that token must be present and have a valid Datadog audience; no other token or environment authentication is used as fallback. Explicit API/application keys, bearer_token, or cloudProviderType take precedence. When unset or empty, automatically tries the DATADOG tag and then the untagged token. Tags may contain only letters, numbers, and underscores; TYPE is reserved.

<a id="nestedblock--default_tags"></a>
### Nested Schema for `defaultTags`

Optional:

- `tags` (Map of String) [Experimental - Action Connections, Logs Indexes, Logs Pipelines, Monitors, Security Monitoring Rules, and Service Level Objectives only] Resource tags to be applied by default across all resources.