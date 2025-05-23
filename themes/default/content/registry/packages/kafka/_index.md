---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-kafka/v3.10.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Kafka Provider
meta_desc: Provides an overview on how to configure the Pulumi Kafka provider.
layout: package
---
## Installation

The Kafka provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/kafka`](https://www.npmjs.com/package/@pulumi/kafka)
* Python: [`pulumi-kafka`](https://pypi.org/project/pulumi-kafka/)
* Go: [`github.com/pulumi/pulumi-kafka/sdk/v3/go/kafka`](https://github.com/pulumi/pulumi-kafka)
* .NET: [`Pulumi.Kafka`](https://www.nuget.org/packages/Pulumi.Kafka)
* Java: [`com.pulumi/kafka`](https://central.sonatype.com/artifact/com.pulumi/kafka)
## Configuration Reference
### Required

- `bootstrapServers` (List of String) A list of kafka brokers

- `caCert` (String) CA certificate file to validate the server's certificate.
- `caCertFile` (String, Deprecated) Path to a CA certificate file to validate the server's certificate.
- `clientCert` (String) The client certificate.
- `clientCertFile` (String, Deprecated) Path to a file containing the client certificate.
- `clientKey` (String) The private key that the certificate was issued for.
- `clientKeyFile` (String, Deprecated) Path to a file containing the private key that the certificate was issued for.
- `clientKeyPassphrase` (String) The passphrase for the private key that the certificate was issued for.
- `kafkaVersion` (String) The version of Kafka protocol to use in `$MAJOR.$MINOR.$PATCH` format. Some features may not be available on older versions. Default is 2.7.0.
- `saslAwsCredsDebug` (Boolean) Set this to true to turn AWS credentials debug.
- `saslAwsProfile` (String) AWS profile name to use
- `saslAwsRegion` (String) AWS region where MSK is deployed.
- `saslAwsRoleArn` (String) Arn of an AWS IAM role to assume
- `saslMechanism` (String) SASL mechanism, can be plain, scram-sha512, scram-sha256, aws-iam
- `saslPassword` (String) Password for SASL authentication.
- `saslTokenUrl` (String) The url to retrieve oauth2 tokens from, when using sasl mechanism oauthbearer
- `saslUsername` (String) Username for SASL authentication.
- `skipTlsVerify` (Boolean) Set this to true only if the target Kafka server is an insecure development instance.
- `timeout` (Number) Timeout in seconds
- `tlsEnabled` (Boolean) Enable communication with the Kafka Cluster over TLS.