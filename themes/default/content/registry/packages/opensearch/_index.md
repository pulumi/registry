---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/ferlab-ste-justine/opensearch/0.1.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Opensearch Provider
meta_desc: Provides an overview on how to configure the Pulumi Opensearch provider.
layout: package
---

## Generate Provider

The Opensearch provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider ferlab-ste-justine/opensearch
```
## Example Usage

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    opensearch:caCert:
        value: ../shared/opensearch-ca.pem
    opensearch:clientCert:
        value: ../shared/opensearch-admin.pem
    opensearch:clientKey:
        value: ../shared/opensearch-admin.key
    opensearch:endpoints:
        value: https://192.168.122.165:9200,https://192.168.122.166:9200

```
## Configuration Reference

- **ca_cert** (String) File that contains the CA certificate that signed the opensearch servers' certificates. Can alternatively be set with the OPENSEARCH_CACERT environment variable. Can also be omitted.
- **client_cert** (String) File that contains the client certificate used to authentify the user. Can alternatively be set with the OPENSEARCH_CERT environment variable. Can be omitted if password authentication is used.
- **client_key** (String) File that contains the client encryption key used to authentify the user. Can alternatively be set with the OPENSEARCH_KEY environment variable. Can be omitted if password authentication is used.
- **connection_timeout** (String) Timeout to establish the opensearch servers connection in golang duration format. Defaults to 10 seconds.
- **endpoints** (String) Endpoints of the opensearch servers. The entry of each server should follow the http|https://ip:port format and be coma separated. Can alternatively be set with the OPENSEARCH_ENDPOINTS environment variable.
- **password** (String, Sensitive) Password of the opensearch user that will be used to access opensearch. Can alternatively be set with the OPENSEARCH_PASSWORD environment variable. Can also be omitted if tls certificate authentication will be used instead.
- **request_timeout** (String) Timeout for individual requests the provider makes on the opensearch servers in golang duration format. Defaults to 10 seconds.
- **retries** (Number) Number of times operations that result in retriable errors should be re-attempted. Defaults to 10.
- **username** (String) Name of the opensearch user that will be used to access opensearch. Can alternatively be set with the OPENSEARCH_USERNAME environment variable. Can also be omitted if tls certificate authentication will be used instead as the username will be infered from the certificate.