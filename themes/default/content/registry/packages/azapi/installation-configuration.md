---
title: AzAPI Setup
meta_desc: Information on how to install the AzAPI provider.
layout: installation
---

## Installation

The Pulumi AzAPI provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@ediri/azapi`](https://www.npmjs.com/package/@ediri/azapi)
* Python: [`ediri_azapi`](https://pypi.org/project/ediri_azapi/)
* Go: [`github.com/dirien/pulumi-azapi/sdk/go/azapi`](https://github.com/dirien/pulumi-azapi/sdk)
* .NET: [`ediri.Azapi`](https://www.nuget.org/packages/ediri.Azapi)

## Configuration Options

The following configuration points are supported for the `AzAPI` provider:

* `ClientId` - (Optional) The Client ID which should be used. This can also be sourced from the `ARM_CLIENT_ID` Environment Variable.

* `Environment` - (Optional) The Cloud Environment which should be used. Possible values are `public`, `usgovernment` and `china`. Defaults to `public`. This can also be sourced from the `ARM_ENVIRONMENT` Environment Variable.

* `SubscriptionId` - (Optional) The Subscription ID which should be used. This can also be sourced from the `ARM_SUBSCRIPTION_ID` Environment Variable.

* `TenantId` - (Optional) The Tenant ID should be used. This can also be sourced from the `ARM_TENANT_ID` Environment Variable.

---

It's possible to configure the behaviour of certain resources using the following properties:

* `DefaultTags` - (Optional) A mapping of tags which should be assigned to the azure resource as default tags. `tags` in each resource block can override the `default_tags`.

* `DefaultLocation` - (Optional) The default Azure Region where the azure resource should exist. `location` in each resource block can override the `default_location`. Changing this forces new resources to be created.

---

When authenticating as a Service Principal using a Client Certificate, the following fields can be set:

* `ClientCertificatePassword` - (Optional) The password associated with the Client Certificate. This can also be sourced from the `ARM_CLIENT_CERTIFICATE_PASSWORD` Environment Variable.

* `ClientCertificatePath` - (Optional) The path to the Client Certificate associated with the Service Principal which should be used. This can also be sourced from the `ARM_CLIENT_CERTIFICATE_PATH` Environment Variable.

---

When authenticating as a Service Principal using a Client Secret, the following fields can be set:

* `ClientSecret` - (Optional) The Client Secret which should be used. This can also be sourced from the `ARM_CLIENT_SECRET` Environment Variable.

---

When authenticating as a Service Principal using Open ID Connect, the following fields can be set:

* `OidcRequestToken` - (Optional) The bearer token for the request to the OIDC provider. This can also be sourced from the `ARM_OIDC_REQUEST_TOKEN` or `ACTIONS_ID_TOKEN_REQUEST_TOKEN` Environment Variables.

* `OidcRequestUrl` - (Optional) The URL for the OIDC provider from which to request an ID token. This can also be sourced from the `ARM_OIDC_REQUEST_URL` or `ACTIONS_ID_TOKEN_REQUEST_URL` Environment Variables.

* `OidcToken` - (Optional) The ID token when authenticating using OpenID Connect (OIDC). This can also be sourced from the `ARM_OIDC_TOKEN` environment Variable.

* `OidcTokenFilePath` - (Optional) The path to a file containing an ID token when authenticating using OpenID Connect (OIDC). This can also be sourced from the `ARM_OIDC_TOKEN_FILE_PATH` environment Variable.

* `UseOidc` - (Optional) Should OIDC be used for Authentication? This can also be sourced from the `ARM_USE_OIDC` Environment Variable. Defaults to `false`.

---

For some advanced scenarios, such as where more granular permissions are necessary - the following properties can be set:

* `DisableCorrelationRequestId` - (Optional) Disable sending the `x-ms-correlation-request-id` header. This can also be sourced from the `ARM_DISABLE_CORRELATION_REQUEST_ID` environment variable. Defaults to `false`.

* `DisableTerraformPartnerId` - (Optional) Disable sending the Terraform Partner ID if a custom `partner_id` isn't specified, which allows Microsoft to better understand the usage of Terraform. The Partner ID does not give HashiCorp any direct access to usage information. This can also be sourced from the `ARM_DISABLE_TERRAFORM_PARTNER_ID` environment variable. Defaults to `false`.

* `PartnerId` - (Optional) A GUID/UUID that is [registered](https://docs.microsoft.com/azure/marketplace/azure-partner-customer-usage-attribution#register-guids-and-offers) with Microsoft to facilitate partner resource usage attribution. This can also be sourced from the `ARM_PARTNER_ID` Environment Variable.

* `SkipProviderRegistration` - (Optional) Should the Provider skip registering the Resource Providers it supports? This can also be sourced from the `ARM_SKIP_PROVIDER_REGISTRATION` Environment Variable. Defaults to `false`.

