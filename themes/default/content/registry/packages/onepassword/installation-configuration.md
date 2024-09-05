---
title: 1Password Provider Installation & Configuration
meta_desc: Provides an overview on how to configure credentials for the 1Password provider for Pulumi.
layout: package
---

The 1Password provider for Pulumi allows you to access and manage items in your 1Password vaults using Pulumi.

## Installation

The 1Password provider for Pulumi is available as a package in most Pulumi languages:

- JavaScript/TypeScript: [@1password/pulumi-onepassword](https://www.npmjs.com/package/@1password/pulumi-onepassword)
- Python: [pulumi_onepassword](https://pypi.org/project/pulumi-onepassword/)
- Go: [github.com/1Password/pulumi-onepassword/sdk/go/onepassword](https://pkg.go.dev/github.com/1Password/pulumi-onepassword/sdk/go/onepassword)
- .NET: _coming soon_

### Provider Binary

The 1Password provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```sh
pulumi plugin install resource onepassword <version> --server github://api.github.com/1Password/pulumi-onepassword
```

Replace the version string with your desired version.

## Configuration

You must configure the 1Password provider for Pulumi with your 1Password credentials before the provider can be used to access and manage items in your 1Password vaults. You can use one of the following:

- [Service account](https://developer.1password.com/docs/service-accounts/get-started)
  - `pulumi-onepassword:service_account_token` (environment: `OP_SERVICE_ACCOUNT_TOKEN`) - the 1Password service account token. Requires [1Password CLI](https://developer.1password.com/docs/cli/get-started/).
- [Connect server](https://developer.1password.com/docs/connect/get-started)
  - `pulumi-onepassword:url` (environment: `OP_CONNECT_HOST`) - the URL where your 1Password Connect API can be found.
  - `pulumi-onepassword:token` (environment: `OP_CONNECT_TOKEN`) - the token for your Connect API.
- Personal account
  - `pulumi-onepassword:account` (environment: `OP_ACCOUNT`) - the [sign-in address](https://support.1password.com/1password-glossary/#sign-in-address) or [unique identifier](https://developer.1password.com/docs/cli/reference/#unique-identifiers-ids) for your 1Password account. Requires [1Password CLI](https://developer.1password.com/docs/cli/get-started/) and biometric unlock.

After you find your credentials, there are two ways to provide them to Pulumi:

1. Set the environment variables for the preferred method. For example, to set the environment variable for a service account token:

   ```sh
   $ export OP_SERVICE_ACCOUNT_TOKEN=XXXXXXXXXXXXXX
   ```

2. If you prefer to store your credentials alongside your Pulumi stack for easy multi-user access, use configuration to set them.

   ```sh
   $ pulumi config set pulumi-onepassword:service_account_token --secret
   Value: <paste token here>
   ```

Make sure to pass `--secret` when setting any sensitive data (in this example `pulumi-onepassword:service_account_token`) so that it's properly encrypted. The complete list of configuration parameters is in the [1Password provider for Pulumi README](https://github.com/1Password/pulumi-onepassword/blob/main/README.md#configuration).
