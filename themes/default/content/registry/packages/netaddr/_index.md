---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/ferlab-ste-justine/netaddr/0.5.1/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Netaddr Provider
meta_desc: Provides an overview on how to configure the Pulumi Netaddr provider.
layout: package
---

## Generate Provider

The Netaddr provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider ferlab-ste-justine/netaddr
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
- `connectionTimeout` (Number) Timeout to establish the etcd servers connection in seconds. Defaults to 10.
- `endpoints` (String) Endpoints of the etcd servers. The entry of each server should follow the ip:port format and be coma separated. Can alternatively be set with the ETCDCTL_ENDPOINTS environment variable.
- `key` (String) File that contains the client encryption key used to authentify the user. Can alternatively be set with the ETCDCTL_KEY environment variable. Can be omitted if password authentication is used.
- `password` (String, Sensitive) Password of the etcd user that will be used to access etcd. Can alternatively be set with the ETCDCTL_PASSWORD environment variable. Can also be omitted if tls certificate authentication will be used instead.
- `requestTimeout` (Number) Timeout for individual requests the provider makes on the etcd servers in seconds. Defaults to 10.
- `retries` (Number) Number of times operations that result in retriable errors should be re-attempted. Defaults to 10.
- `strict` (Boolean) Whether the provider should trigger a failure if resources are already existing during their creation, already absent during their deletion or otherwise absent during reads. Setting this value to false is convenient, but it might not alert you of bad failure situations (like resource name duplicates or the etcd state being tampered outside of pulumi) so we recommend using this setting only to recover for failure situations that are well understood like Pulumi having failed to persist its state in a previous apply.
- `username` (String) Name of the etcd user that will be used to access etcd. Can alternatively be set with the ETCDCTL_USERNAME environment variable. Can also be omitted if tls certificate authentication will be used instead as the username will be infered from the certificate.