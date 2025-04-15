---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/ferlab-ste-justine/etcd/0.10.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Etcd Provider
meta_desc: Provides an overview on how to configure the Pulumi Etcd provider.
layout: package
---

## Generate Provider

The Etcd provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider ferlab-ste-justine/etcd
```
## Example Usage

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    etcd:caCert:
        value: 'TODO: "${path.module}/certs/ca.pem"'
    etcd:cert:
        value: 'TODO: "${path.module}/certs/root.pem"'
    etcd:endpoints:
        value: 127.0.0.1:32379
    etcd:key:
        value: 'TODO: "${path.module}/certs/root.key"'

```
## Configuration Reference

- `caCert` (String) File that contains the CA certificate that signed the etcd servers' certificates. Can alternatively be set with the ETCDCTL_CACERT environment variable. Can also be omitted.
- `cert` (String) File that contains the client certificate used to authentify the user. Can alternatively be set with the ETCDCTL_CERT environment variable. Can be omitted if password authentication is used.
- `connectionTimeout` (String) Timeout to establish the etcd servers connection as a duration. Defaults to 10s.
- `endpoints` (String) Endpoints of the etcd servers. The entry of each server should follow the ip:port format and be coma separated. Can alternatively be set with the ETCDCTL_ENDPOINTS environment variable.
- `key` (String) File that contains the client encryption key used to authentify the user. Can alternatively be set with the ETCDCTL_KEY environment variable. Can be omitted if password authentication is used.
- `password` (String, Sensitive) Password of the etcd user that will be used to access etcd. Can alternatively be set with the ETCDCTL_PASSWORD environment variable. Can also be omitted if tls certificate authentication will be used instead.
- `requestTimeout` (String) Timeout for individual requests the provider makes on the etcd servers as a duration. Defaults to 10s.
- `retries` (Number) Number of times operations that result in retriable errors should be re-attempted. Defaults to 10.
- `retryInterval` (String) Duration to wait after a failing etcd request before retrying. Defaults to 100ms.
- `skipTls` (Boolean) If set to true, connection to the etcd cluster will be attempted in plaintext without encryption. Default to false
- `username` (String) Name of the etcd user that will be used to access etcd. Can alternatively be set with the ETCDCTL_USERNAME environment variable. Can also be omitted if tls certificate authentication will be used instead as the username will be infered from the certificate.