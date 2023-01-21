---
title: Aviatrix Installation & Configuration
meta_desc: Provides an overview on how to configure credentials for the Pulumi Aviatrix Provider.
layout: installation
---

## Installation

1. To use this package, please [install the Pulumi CLI first](https://www.pulumi.com/docs/get-started/install/).
1. This package is only available for Python currently, but support for other languages might be available later.

### Python

To use from JavaScript or TypeScript in Node.js, install using either `npm` or `yarn`:

```bash
$ pip install pulumi-aviatrix
```

## Configuration

The following configuration points are available for the `aviatrix` provider:

Mandatory:
- `aviatrix:controllerIp` (environment: `AVIATRIX_CONTROLLER_IP`) - The Aviatrix controller's IP address
- `aviatrix:username` (environment: `AVIATRIX_USERNAME`) - The Aviatrix controller's username
- `aviatrix:password` (environment: `AVIATRIX_PASSWORD`) - The Aviatrix controller's password

Optional:
- `aviatrix:skip_version_validation` (environment: `AVIATRIX_SKIP_VERSION_VALIDATION`) - Valid values: true, false. Default: false. If set to true, it skips checking whether current Pulumi provider supports current Controller version.
- `aviatrix:verify_ssl_certificate` (environment: `AVIATRIX_VERIFY_SSL_CERTIFICATE`) - Valid values: true, false. Default: false. If set to true, the SSL certificate of the controller will be verified.

## Controller compatibility

This package has feature parity with the Aviatrix Terraform provider version 2.24.1. For controller version compatibility please visit the [compatibility page](https://registry.terraform.io/providers/AviatrixSystems/aviatrix/latest/docs/guides/release-compatibility).

## Reference

For detailed reference documentation, please visit the [Terraform registry](https://registry.terraform.io/providers/AviatrixSystems/aviatrix/latest/docs).
