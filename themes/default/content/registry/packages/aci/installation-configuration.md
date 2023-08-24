---
title: ACI Installation & Configuration
meta_desc: Information on how to install the ACI provider.
layout: installation
---

## Installation

The Pulumiverse ACI provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@netascode/aci`](https://www.npmjs.com/package/@netascode/aci)
* Python: [`pulumi_aci`](https://pypi.org/project/pulumi-aci/)
* Go: [`github.com/netascode/pulumi-aci/sdk/go/aci`](https://pkg.go.dev/github.com/netascode/pulumi-aci/sdk)
* .NET: [`Pulumi.Aci`](https://www.nuget.org/packages/Pulumi.Aci)

## Setup

To provision resources with the Pulumi ACI provider, you need to have correct endpoint configuration towards your APIC controller.

## Configuration Options

Use `pulumi config set aci:<option> (--secret)`.

| Option | Environment Variable | Required/Optional | Description |
|-----|------|------|----|
| `url` | `ACI_URL` | Required | URL of the Cisco APIC web interface |
| `username`| `ACI_USERNAME` | Required | Username for the APIC Account |
| `password`| `ACI_PASSWORD` | Required (Secret) | Password for the APIC Account |
| `insecure` | `ACI_INSECURE` | Optional | Allow insecure HTTPS client |
| `logging` | `ACI_LOGGING` | Optional | Enable debug logging |
| `retries` | `ACI_RETRIES` | Optional | Number of retries for REST API calls |
