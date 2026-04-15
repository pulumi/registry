---
# WARNING: this file was fetched from https://raw.githubusercontent.com/ionos-cloud/pulumi-ionoscloud/v0.3.0/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/ionos-cloud/pulumi-ionoscloud/blob/v0.3.0/docs/installation-configuration.md
title: IONOS Cloud Installation & Configuration
meta_desc: Information on how to install the IONOS Cloud provider.
layout: installation
---

## Installation

The Pulumi `IONOS Cloud` provider enables you to manage IONOS Cloud resources using Pulumi in your preferred language. Follow the steps below to install the provider and set up your environment.

### 1. Install the Pulumi CLI

If you haven't already, [install the Pulumi CLI](https://www.pulumi.com/docs/get-started/install/).

### 2. Install the IONOS Cloud Provider Package

Depending on your language, install the provider package:

- **JavaScript/TypeScript**:  
  [`@ionos-cloud/sdk-pulumi`](https://www.npmjs.com/package/@ionos-cloud/sdk-pulumi)
  ```bash
  npm install @ionos-cloud/sdk-pulumi
  # or
  yarn add @ionos-cloud/sdk-pulumi
  ```

- **Python**:  
  [`pulumi_ionoscloud`](https://pypi.org/project/pulumi-ionoscloud)
  ```bash
  pip install pulumi_ionoscloud
  ```

- **Go**:  
  [`github.com/ionos-cloud/pulumi-ionoscloud/sdk/go/ionoscloud`](https://pkg.go.dev/github.com/ionos-cloud/pulumi-ionoscloud/sdk)
  ```bash
  go get github.com/ionos-cloud/pulumi-ionoscloud/sdk
  ```

- **.NET**:  
  [`Ionoscloud.pulumi.ionoscloud`](https://www.nuget.org/packages/Ionoscloud.pulumi.ionoscloud/)
  ```bash
  dotnet add package Ionoscloud.Pulumi.Ionoscloud
  ```

### 3. Install the Provider Binary

The IONOS Cloud provider binary is a third-party binary required by Pulumi to interact with IONOS Cloud. Install it using the Pulumi plugin command:

```bash
pulumi plugin install resource ionoscloud vX.Y.Z --server github://api.github.com/ionos-cloud/pulumi-ionoscloud
```

Replace `vX.Y.Z` with the desired version.  
You can check for the latest version [here](https://github.com/ionos-cloud/pulumi-ionoscloud/releases).

---

## Configuration

Before you can provision resources, you must configure the provider with your IONOS Cloud credentials.

### Required Credentials

You need either:

- **Username and Password** or **API Token**

See [Token Management](https://docs.ionos.com/cloud/set-up-ionos-cloud/management/identity-access-management/token-manager) for details on generating a token.

### Setting Credentials

You can provide credentials using environment variables (recommended for security):

| Environment Variable    | Description                                                                                                                                                              |
|-------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `IONOS_USERNAME`        | Username for IONOS Cloud API authentication                                                                                                                              |
| `IONOS_PASSWORD`        | Password for IONOS Cloud API authentication                                                                                                                              |
| `IONOS_TOKEN`           | API token for authentication (alternative to username/password)                                                                    |
| `IONOS_API_URL`         | API endpoint URL (default: `api.ionos.com`). Override only if necessary.                                                          |
| `IONOS_LOG_LEVEL`       | Log level: `Off`, `Debug`, or `Trace` (default: `Off`)                                                                           |
| `IONOS_PINNED_CERT`     | SHA-256 public fingerprint for certificate pinning                                                                                |
| `IONOS_CONTRACT_NUMBER` | Contract number (for reseller accounts only)                                                                                      |
| `IONOS_S3_ACCESS_KEY`   | Access key for IONOS Object Storage API                                                                                           |
| `IONOS_S3_SECRET_KEY`   | Secret key for IONOS Object Storage API                                                                                           |
| `IONOS_S3_REGION`       | Region for Object Storage (default: `eu-central-3`). Required if using `IONOS_API_URL_OBJECT_STORAGE`.                            |

Example (Linux/macOS):
```bash
export IONOS_USERNAME="your-username"
export IONOS_PASSWORD="your-password"
# or
export IONOS_TOKEN="your-api-token"
```

---

## Certificate Pinning

To enable certificate pinning, set the `IONOS_PINNED_CERT` environment variable to the SHA-256 public fingerprint of the certificate.  
You can obtain the fingerprint from your browser's certificate details.

---

## Debugging

Enable logging for troubleshooting by setting `IONOS_LOG_LEVEL`:

```bash
export IONOS_LOG_LEVEL=debug
```

Allowed values: `Off`, `Debug`, `Trace`.  
**Note:** Use `trace` only for debugging. Disable it in your production environments because it can log sensitive data. It logs the full request and response without encryption, even for an HTTPS call. Verbose request and response logging can also significantly impact your application’s performance.

---

## Troubleshooting

- Ensure your credentials are correct and exported in your environment.
- If you encounter plugin errors, verify the provider binary version matches your Pulumi project requirements.
- For more details, refer to the [Pulumi IONOS Cloud documentation](https://github.com/ionos-cloud/pulumi-ionoscloud) or open an issue on GitHub.
