---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-kafka/v3.12.3/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi/pulumi-kafka/blob/v3.12.3/docs/_index.md
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

## Overview

The Kafka provider is used to interact with [Apache Kafka](https://kafka.apache.org/) clusters. The provider allows you to manage Kafka topics, ACLs, quotas, and SCRAM credentials. It supports various authentication methods including TLS, SASL/PLAIN, SASL/SCRAM, AWS IAM, and OAuth.
## Documentation

- Quick Start Guide - Get started quickly with common scenarios
- Authentication Guide - Detailed authentication configuration
- AWS MSK Integration - Complete MSK setup guide
- Migration Guide - Migrate between setups and versions
- Troubleshooting Guide - Common issues and solutions
## Example Usage
### Basic Configuration with TLS

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    kafka:bootstrapServers:
        value:
            - localhost:9092
    kafka:caCert:
        value: 'TODO: file("../secrets/ca.crt")'
    kafka:clientCert:
        value: 'TODO: file("../secrets/client-cert.pem")'
    kafka:clientKey:
        value: 'TODO: file("../secrets/client-key.pem")'
    kafka:tlsEnabled:
        value: true

```

```typescript
import * as pulumi from "@pulumi/pulumi";

```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    kafka:bootstrapServers:
        value:
            - localhost:9092
    kafka:caCert:
        value: 'TODO: file("../secrets/ca.crt")'
    kafka:clientCert:
        value: 'TODO: file("../secrets/client-cert.pem")'
    kafka:clientKey:
        value: 'TODO: file("../secrets/client-key.pem")'
    kafka:tlsEnabled:
        value: true

```

```python
import pulumi

```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    kafka:bootstrapServers:
        value:
            - localhost:9092
    kafka:caCert:
        value: 'TODO: file("../secrets/ca.crt")'
    kafka:clientCert:
        value: 'TODO: file("../secrets/client-cert.pem")'
    kafka:clientKey:
        value: 'TODO: file("../secrets/client-key.pem")'
    kafka:tlsEnabled:
        value: true

```

```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;

return await Deployment.RunAsync(() =>
{
});

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    kafka:bootstrapServers:
        value:
            - localhost:9092
    kafka:caCert:
        value: 'TODO: file("../secrets/ca.crt")'
    kafka:clientCert:
        value: 'TODO: file("../secrets/client-cert.pem")'
    kafka:clientKey:
        value: 'TODO: file("../secrets/client-key.pem")'
    kafka:tlsEnabled:
        value: true

```

```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		return nil
	})
}
```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    kafka:bootstrapServers:
        value:
            - localhost:9092
    kafka:caCert:
        value: 'TODO: file("../secrets/ca.crt")'
    kafka:clientCert:
        value: 'TODO: file("../secrets/client-cert.pem")'
    kafka:clientKey:
        value: 'TODO: file("../secrets/client-key.pem")'
    kafka:tlsEnabled:
        value: true

```

```yaml
{}
```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    kafka:bootstrapServers:
        value:
            - localhost:9092
    kafka:caCert:
        value: 'TODO: file("../secrets/ca.crt")'
    kafka:clientCert:
        value: 'TODO: file("../secrets/client-cert.pem")'
    kafka:clientKey:
        value: 'TODO: file("../secrets/client-key.pem")'
    kafka:tlsEnabled:
        value: true

```

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
    }
}
```

{{% /choosable %}}
{{< /chooser >}}
### SASL/PLAIN Authentication

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    kafka:bootstrapServers:
        value:
            - localhost:9092
    kafka:saslMechanism:
        value: plain
    kafka:saslPassword:
        value: 'TODO: var.kafka_password'
    kafka:saslUsername:
        value: pulumi
    kafka:tlsEnabled:
        value: true

```
### SASL/SCRAM Authentication

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    kafka:bootstrapServers:
        value:
            - localhost:9092
    kafka:saslMechanism:
        value: scram-sha256
    kafka:saslPassword:
        value: 'TODO: var.kafka_password'
    kafka:saslUsername:
        value: pulumi
    kafka:tlsEnabled:
        value: true

```
### AWS MSK with IAM Authentication (Using AssumeRole)

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    kafka:bootstrapServers:
        value:
            - b-1.msk-cluster.xxx.kafka.us-east-1.amazonaws.com:9098
    kafka:saslAwsRegion:
        value: us-east-1
    kafka:saslAwsRoleArn:
        value: arn:aws:iam::123456789012:role/kafka-pulumi-role
    kafka:saslMechanism:
        value: aws-iam
    kafka:tlsEnabled:
        value: true

```
### AWS MSK with IAM Authentication (Using AWS Profile)

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    kafka:bootstrapServers:
        value:
            - b-1.msk-cluster.xxx.kafka.us-east-1.amazonaws.com:9098
    kafka:saslAwsProfile:
        value: production
    kafka:saslAwsRegion:
        value: us-east-1
    kafka:saslMechanism:
        value: aws-iam
    kafka:tlsEnabled:
        value: true

```
### AWS MSK with IAM Authentication (Using Static Credentials)

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    kafka:bootstrapServers:
        value:
            - b-1.msk-cluster.xxx.kafka.us-east-1.amazonaws.com:9098
    kafka:saslAwsAccessKey:
        value: 'TODO: data.vault_aws_access_credentials.creds.access_key'
    kafka:saslAwsRegion:
        value: us-east-1
    kafka:saslAwsSecretKey:
        value: 'TODO: data.vault_aws_access_credentials.creds.secret_key'
    kafka:saslAwsToken:
        value: 'TODO: data.vault_aws_access_credentials.creds.security_token'
    kafka:saslMechanism:
        value: aws-iam
    kafka:tlsEnabled:
        value: true

```

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as vault from "@pulumi/vault";

const creds = vault.aws.getAccessCredentials({
    backend: "aws",
    type: "sts",
    role: "kafka-access-role",
});
```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    kafka:bootstrapServers:
        value:
            - b-1.msk-cluster.xxx.kafka.us-east-1.amazonaws.com:9098
    kafka:saslAwsAccessKey:
        value: 'TODO: data.vault_aws_access_credentials.creds.access_key'
    kafka:saslAwsRegion:
        value: us-east-1
    kafka:saslAwsSecretKey:
        value: 'TODO: data.vault_aws_access_credentials.creds.secret_key'
    kafka:saslAwsToken:
        value: 'TODO: data.vault_aws_access_credentials.creds.security_token'
    kafka:saslMechanism:
        value: aws-iam
    kafka:tlsEnabled:
        value: true

```

```python
import pulumi
import pulumi_vault as vault

creds = vault.aws.get_access_credentials(backend="aws",
    type="sts",
    role="kafka-access-role")
```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    kafka:bootstrapServers:
        value:
            - b-1.msk-cluster.xxx.kafka.us-east-1.amazonaws.com:9098
    kafka:saslAwsAccessKey:
        value: 'TODO: data.vault_aws_access_credentials.creds.access_key'
    kafka:saslAwsRegion:
        value: us-east-1
    kafka:saslAwsSecretKey:
        value: 'TODO: data.vault_aws_access_credentials.creds.secret_key'
    kafka:saslAwsToken:
        value: 'TODO: data.vault_aws_access_credentials.creds.security_token'
    kafka:saslMechanism:
        value: aws-iam
    kafka:tlsEnabled:
        value: true

```

```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Vault = Pulumi.Vault;

return await Deployment.RunAsync(() =>
{
    var creds = Vault.Aws.GetAccessCredentials.Invoke(new()
    {
        Backend = "aws",
        Type = "sts",
        Role = "kafka-access-role",
    });

});

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    kafka:bootstrapServers:
        value:
            - b-1.msk-cluster.xxx.kafka.us-east-1.amazonaws.com:9098
    kafka:saslAwsAccessKey:
        value: 'TODO: data.vault_aws_access_credentials.creds.access_key'
    kafka:saslAwsRegion:
        value: us-east-1
    kafka:saslAwsSecretKey:
        value: 'TODO: data.vault_aws_access_credentials.creds.secret_key'
    kafka:saslAwsToken:
        value: 'TODO: data.vault_aws_access_credentials.creds.security_token'
    kafka:saslMechanism:
        value: aws-iam
    kafka:tlsEnabled:
        value: true

```

```go
package main

import (
	"github.com/pulumi/pulumi-vault/sdk/v7/go/vault/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aws.GetAccessCredentials(ctx, &aws.GetAccessCredentialsArgs{
			Backend: "aws",
			Type:    pulumi.StringRef("sts"),
			Role:    "kafka-access-role",
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}
```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    kafka:bootstrapServers:
        value:
            - b-1.msk-cluster.xxx.kafka.us-east-1.amazonaws.com:9098
    kafka:saslAwsAccessKey:
        value: 'TODO: data.vault_aws_access_credentials.creds.access_key'
    kafka:saslAwsRegion:
        value: us-east-1
    kafka:saslAwsSecretKey:
        value: 'TODO: data.vault_aws_access_credentials.creds.secret_key'
    kafka:saslAwsToken:
        value: 'TODO: data.vault_aws_access_credentials.creds.security_token'
    kafka:saslMechanism:
        value: aws-iam
    kafka:tlsEnabled:
        value: true

```

```yaml
variables:
  creds:
    fn::invoke:
      function: vault:aws:getAccessCredentials
      arguments:
        backend: aws
        type: sts
        role: kafka-access-role
```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    kafka:bootstrapServers:
        value:
            - b-1.msk-cluster.xxx.kafka.us-east-1.amazonaws.com:9098
    kafka:saslAwsAccessKey:
        value: 'TODO: data.vault_aws_access_credentials.creds.access_key'
    kafka:saslAwsRegion:
        value: us-east-1
    kafka:saslAwsSecretKey:
        value: 'TODO: data.vault_aws_access_credentials.creds.secret_key'
    kafka:saslAwsToken:
        value: 'TODO: data.vault_aws_access_credentials.creds.security_token'
    kafka:saslMechanism:
        value: aws-iam
    kafka:tlsEnabled:
        value: true

```

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.vault.aws.AwsFunctions;
import com.pulumi.vault.aws.inputs.GetAccessCredentialsArgs;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        final var creds = AwsFunctions.getAccessCredentials(GetAccessCredentialsArgs.builder()
            .backend("aws")
            .type("sts")
            .role("kafka-access-role")
            .build());

    }
}
```

{{% /choosable %}}
{{< /chooser >}}
### OAuth2 Authentication

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    kafka:bootstrapServers:
        value:
            - localhost:9092
    kafka:saslMechanism:
        value: oauthbearer
    kafka:saslOauthScopes:
        value:
            - kafka:read
            - kafka:write
    kafka:saslTokenUrl:
        value: https://oauth.example.com/oauth2/token
    kafka:tlsEnabled:
        value: true

```
### Redpanda Compatibility

When using with [Redpanda](https://redpanda.com/), you may need to force a specific Kafka protocol version:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    kafka:bootstrapServers:
        value:
            - localhost:9092
    kafka:kafkaVersion:
        value: 2.1.0

```
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
- `saslAwsAccessKey` (String) The AWS access key.
- `saslAwsContainerAuthorizationTokenFile` (String) Path to a file containing the AWS pod identity authorization token
- `saslAwsContainerCredentialsFullUri` (String) URI to retrieve AWS credentials from
- `saslAwsCredsDebug` (Boolean) Set this to true to turn AWS credentials debug.
- `saslAwsExternalId` (String) External ID of the AWS IAM role to assume
- `saslAwsProfile` (String) AWS profile name to use
- `saslAwsRegion` (String) AWS region where MSK is deployed.
- `saslAwsRoleArn` (String) Arn of an AWS IAM role to assume
- `saslAwsSecretKey` (String) The AWS secret key.
- `saslAwsSharedConfigFiles` (List of String) List of paths to AWS shared config files.
- `saslAwsToken` (String) The AWS session token. Only required if you are using temporary security credentials.
- `saslMechanism` (String) SASL mechanism, can be plain, scram-sha512, scram-sha256, aws-iam
- `saslOauthScopes` (List of String) OAuth scopes to request when using the oauthbearer mechanism
- `saslPassword` (String) Password for SASL authentication.
- `saslTokenUrl` (String) The url to retrieve oauth2 tokens from, when using sasl mechanism oauthbearer
- `saslUsername` (String) Username for SASL authentication.
- `skipTlsVerify` (Boolean) Set this to true only if the target Kafka server is an insecure development instance.
- `timeout` (Number) Timeout in seconds
- `tlsEnabled` (Boolean) Enable communication with the Kafka Cluster over TLS.