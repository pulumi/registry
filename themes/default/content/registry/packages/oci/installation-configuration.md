---
title: Oracle Cloud Infrastructure Setup
meta_desc: Provides an overview on how to configure credentials for the Oracle Cloud Infrastructure Provider for Pulumi.
layout: installation
---

The Oracle Cloud Infrastructure (OCI) provider for Pulumi uses the OCI SDK to manage and provision resources.

## Installation

The OCI provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/oci`](https://www.npmjs.com/package/@pulumi/oci)
* Python: [`pulumi-oci`](https://pypi.org/project/pulumi-oci/)
* Go: [`github.com/pulumi/pulumi-oci/sdk/go/oci`](https://github.com/pulumi/pulumi-oci)
* .NET: [`Pulumi.Oci`](https://www.nuget.org/packages/Pulumi.Oci)

## Configuring Credentials

Pulumi relies on the OCI SDK to authenticate requests from your computer to OCI. Your credentials are never sent to Pulumi.

Use the [Required Keys and OCIDs](https://docs.oracle.com/en-us/iaas/Content/API/Concepts/apisigningkey.htm#Required_Keys_and_OCIDs) chapter of the OCI Developer Guide to learn:

* [How to Generate an API Signing Key](https://docs.oracle.com/en-us/iaas/Content/API/Concepts/apisigningkey.htm#two)
* [How to Get the Key's Fingerprint](https://docs.oracle.com/en-us/iaas/Content/API/Concepts/apisigningkey.htm#four)
* [Where to Get the Tenancy's OCID and User's OCID](https://docs.oracle.com/en-us/iaas/Content/API/Concepts/apisigningkey.htm#five)

Once you have generated an API signing key pair and have your tenancy and user OCIDs
you can make them available using either environment variables or Pulumi configuration.

### Environment Variables

```bash
export TF_VAR_tenancy_ocid="ocid1.tenancy.oc1..<unique_ID>"
export TF_VAR_user_ocid="ocid1.user.oc1..<unique_ID>"
export TF_VAR_fingerprint="<key_fingerprint>"
export TF_VAR_region="us-ashburn-1"
export TF_VAR_private_key_file="/path/to/oci_api_key.pem"
```

### Pulumi Configuration

If you prefer to store the credentials alongside your Pulumi stack for multi-user access:

```bash
pulumi config set oci:tenancyOcid "ocid1.tenancy.oc1..<unique_ID>" --secret
pulumi config set oci:userOcid "ocid1.user.oc1..<unique_ID>" --secret
pulumi config set oci:fingerprint "<key_fingerprint>" --secret
pulumi config set oci:region "us-ashburn-1"
# Set the private key from standard input to retain the format
cat "~/.oci/oci_api_key.pem" | pulumi config set oci:privateKey --secret
```

Remember to pass `--secret` when setting `oci:secretKey` so that it is properly encrypted.

## Configuration options

The following configuration options are available for the `oci` provider:

| Option                   | Environment variable          | Description                                                                                                                 |
| ------------------------ | ----------------------------- | --------------------------------------------------------------------------------------------------------------------------- |
| `oci:tenancyOcid`        | `TF_VAR_tenancy_ocid`         | OCID of your tenancy.                                                                                                       |
| `oci:userOcid`           | `TF_VAR_user_ocid`            | OCID of the user calling the API.                                                                                           |
| `oci:privateKey`         | `TF_VAR_private_key`          | The contents of the private key file. Required if `privateKeyPath` is not defined and takes precedence if both are defined. |
| `oci:privateKeyPath`     | `TF_VAR_private_key_path`     | The path (including filename) of the private key stored on your computer. Required if `privateKey` is not defined.          |
| `oci:privateKeyPassword` | `TF_VAR_private_key_password` | Passphrase used for the key, if it is encrypted.                                                                            |
| `oci:fingerprint`        | `TF_VAR_fingerprint`          | Fingerprint for the key pair being used.                                                                                    |
| `oci:region`             | `TF_VAR_region`               | An OCI region.                                                                                                              |
| `oci:configFileProfile`  | `TF_VAR_config_file_profile`  | The custom profile to use instead of the `DEFAULT` profile in `.oci/config`.             |
