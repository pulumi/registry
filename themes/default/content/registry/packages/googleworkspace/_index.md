---
# WARNING: this file was fetched from https://ds7waww6k7rw2.cloudfront.net/docs/registry.opentofu.org/hashicorp/googleworkspace/0.7.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Googleworkspace Provider
meta_desc: Provides an overview on how to configure the Pulumi Googleworkspace provider.
layout: package
---

## Generate Provider

The Googleworkspace provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider hashicorp/googleworkspace
```
## Overview

The Google Workspace provider provides resources to interact with Google Workspace.
## Authorization
This provider uses [Admin SDK](https://developers.google.com/admin-sdk) API methods to manage resources on a Workspace customer domain. There are multiple ways to set up proper authorization for a service account:

* Enable domain-wide delegation to impersonate a user that has super administrator privileges. You cannot directly grant super administrator privileges to service accounts.

* Assign [specific administrator roles](https://support.google.com/a/answer/9807615?hl=en&ref_topic=9832445) directly to the service account.
## Authentication
### Using Domain-Wide Delegation
#### Creating a Service Account and Credentials

Pulumi uses a GCP service account to manage resources created by the provider. To create the service account and generate a service account key:

1. Follow the instructions in the [create service account and credentials documentation](https://developers.google.com/admin-sdk/directory/v1/guides/delegation#create_the_service_account_and_credentials).
2. Save the json file containing your service account key credentials locally and set the `GOOGLEWORKSPACE_CREDENTIALS` environment variable to the path of that file. Pulumi will now use that key for authentication.
   * Other environment variable options include: `GOOGLEWORKSPACE_CLOUD_KEYFILE_JSON` or `GOOGLE_CREDENTIALS` (>= 0.2.0)
#### Configuring the Service Account's granted scopes in Google Workspace

To access user data on a Google Workspace domain, the service account that you created needs to be granted access
by a super administrator for the domain. Follow the instructions in the
[delegate domain-wide authority documentation](https://developers.google.com/admin-sdk/directory/v1/guides/delegation#delegate_domain-wide_authority_to_your_service_account).

* Note: The Oauth scopes granted to your service account must match, or be a superset, of the `oauthScopes` granted to
  the `googleworkspace` provider.

More information about scopes is below.
#### Configuring scopes requested by the provider

The provider can be configured with an `oauthScopes` field containing a list of requested scopes. If `oauthScopes` is *not* set in the provider configuration, the provider will fall back to a default list of scopes, which are all the scopes needed to manage resources that can be provisioned by the provider (see default scopes for v0.6.0 here). If default scopes are used, the service account needs to be granted all of those scopes in Google Workspace.

The scopes declared in the provider's configuration need to match, or be a subset of, the scopes granted to the service account. If a provider is configured with scopes the service account isn't granted to use, the provider will receive a `401 Unauthorized` response when it requests an access token.

->It's recommended to include `oathScopes` in your provider configuration to make the requested scopes explicit and easier to debug issues.
#### Impersonating a Google Workspace User

Only users with access to the Admin APIs can access the Admin SDK Directory API, therefore your service account needs to impersonate one of those users to access the Admin SDK Directory API. This user's email
must be set in the environment variable `GOOGLEWORKSPACE_IMPERSONATED_USER_EMAIL` or in the `impersonatedUserEmail` attribute in the provider. Additionally, the user must have logged in at least once and accepted the Google Workspace Terms of Service.
### Using Specific Administrator Roles

You do not need to set up domain-wide delegation if you are granting more specific administrator roles to the service account. If the Pulumi pipeline execution environment provides an appropriate token as Application Default Credentials (ADC), you can use the provider without any further setup.

To do this it's recommended that you create a custom admin role(s) with the Admin API privileges you need, as [pre-built administrator roles](https://support.google.com/a/answer/2405986) might not cover your use case.

> Some resources controlled via this provider can only be managed by a Super Admin, so user impersonation is necessary in those cases. For instance, [Admin Roles can only be managed via the console or API by a Super Admin](https://support.google.com/a/answer/2406043?hl=en). Also, access to some API endpoints may not be possible to grant as privileges in custom roles.

When using gcloud locally, you can provide the required scopes for ADC login by adding the `--scopes` parameter to [`gcloud auth application-default login`](https://cloud.google.com/sdk/gcloud/reference/auth/application-default/login). For example, you can [provide additional scopes](https://cloud.google.com/sdk/gcloud/reference/beta/compute/instances/set-scopes) on Compute Engine. You can do this to configure access for both service accounts and end users.

The approach outlined above does not work on Cloud Build because it does not (yet) support specifying additional scopes for service account tokens accessible during builds. Other pipeline setups use the `googleServiceAccountAccessToken` to impersonate a service account. This allows them to use a single identity regardless of who is initiating the execution. For these cases, set the `accessToken` parameter to the appropriate credentials.

You can also provide an exported service account key in the `credentials` parameter without specifying an `impersonatedUserEmail`.
## Configuration Reference

- `accessToken` (String) A temporary [OAuth 2.0 access token] obtained from the Google Authorization server, i.e. the `Authorization: Bearer` token used to authenticate HTTP requests to Google Admin SDK APIs. This is an alternative to `credentials`, and ignores the `oauthScopes` field. If both are specified, `accessToken` will be used over the `credentials` field.
- `credentials` (String) Either the path to or the contents of a service account key file in JSON format you can manage key files using the Cloud Console).  If not provided, the application default credentials will be used.
- `customerId` (String) The customer id provided with your Google Workspace subscription. It is found in the admin console under Account Settings.
- `impersonatedUserEmail` (String) The impersonated user's email with access to the Admin APIs can access the Admin SDK Directory API. `impersonatedUserEmail` is required for all services except group and user management.
- `oauthScopes` (List of String) The list of the scopes required for your application (for a list of possible scopes, see [Authorize requests](https://developers.google.com/admin-sdk/directory/v1/guides/authorizing))
- `serviceAccount` (String) The service account used to create the provided `accessToken` if authenticating using the `accessToken` method and needing to impersonate a user. This service account will require the GCP role `Service Account Token Creator` if needing to impersonate a user.