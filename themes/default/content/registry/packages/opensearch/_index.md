---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/opensearch-project/opensearch/2.3.1/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Opensearch Provider
meta_desc: Provides an overview on how to configure the Pulumi Opensearch provider.
layout: package
---

## Generate Provider

The Opensearch provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider opensearch-project/opensearch
```
## Overview

The provider is used to interact with the resources supported by OpenSearch.
The provider needs to be configured with an endpoint URL before it can be
used.

AWS OpenSearch Service domains and OpenSearch clusters deployed on Kubernetes
and other infrastructure are supported.

Use the navigation to the left to read about the available resources.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    opensearch:url:
        value: http://127.0.0.1:9200

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as opensearch from "@pulumi/opensearch";

// Create an index template
const template1 = new opensearch.IndexTemplate("template_1", {
    name: "template_1",
    body: "{  \\\"index_patterns\\\": [    \\\"your-pattern-here-*\\\"  ],  \\\"template\\\": {    \\\"settings\\\": {      \\\"index\\\": {        \\\"number_of_shards\\\": \\\"1\\\"      }    },    \\\"mappings\\\": {      \\\"_source\\\": {        \\\"enabled\\\": false      },      \\\"properties\\\": {        \\\"host_name\\\": {          \\\"type\\\": \\\"keyword\\\"        },        \\\"created_at\\\": {          \\\"type\\\": \\\"date\\\",          \\\"format\\\": \\\"EEE MMM dd HH:mm:ss Z YYYY\\\"        }      }    }  }}\n",
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    opensearch:url:
        value: http://127.0.0.1:9200

```
```python
import pulumi
import pulumi_opensearch as opensearch

# Create an index template
template1 = opensearch.IndexTemplate("template_1",
    name="template_1",
    body="{  \\\"index_patterns\\\": [    \\\"your-pattern-here-*\\\"  ],  \\\"template\\\": {    \\\"settings\\\": {      \\\"index\\\": {        \\\"number_of_shards\\\": \\\"1\\\"      }    },    \\\"mappings\\\": {      \\\"_source\\\": {        \\\"enabled\\\": false      },      \\\"properties\\\": {        \\\"host_name\\\": {          \\\"type\\\": \\\"keyword\\\"        },        \\\"created_at\\\": {          \\\"type\\\": \\\"date\\\",          \\\"format\\\": \\\"EEE MMM dd HH:mm:ss Z YYYY\\\"        }      }    }  }}\n")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    opensearch:url:
        value: http://127.0.0.1:9200

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Opensearch = Pulumi.Opensearch;

return await Deployment.RunAsync(() =>
{
    // Create an index template
    var template1 = new Opensearch.IndexTemplate("template_1", new()
    {
        Name = "template_1",
        Body = @"{  \""index_patterns\"": [    \""your-pattern-here-*\""  ],  \""template\"": {    \""settings\"": {      \""index\"": {        \""number_of_shards\"": \""1\""      }    },    \""mappings\"": {      \""_source\"": {        \""enabled\"": false      },      \""properties\"": {        \""host_name\"": {          \""type\"": \""keyword\""        },        \""created_at\"": {          \""type\"": \""date\"",          \""format\"": \""EEE MMM dd HH:mm:ss Z YYYY\""        }      }    }  }}
",
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
    opensearch:url:
        value: http://127.0.0.1:9200

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/opensearch/v2/opensearch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create an index template
		_, err := opensearch.NewIndexTemplate(ctx, "template_1", &opensearch.IndexTemplateArgs{
			Name: pulumi.String("template_1"),
			Body: pulumi.String("{  \\\"index_patterns\\\": [    \\\"your-pattern-here-*\\\"  ],  \\\"template\\\": {    \\\"settings\\\": {      \\\"index\\\": {        \\\"number_of_shards\\\": \\\"1\\\"      }    },    \\\"mappings\\\": {      \\\"_source\\\": {        \\\"enabled\\\": false      },      \\\"properties\\\": {        \\\"host_name\\\": {          \\\"type\\\": \\\"keyword\\\"        },        \\\"created_at\\\": {          \\\"type\\\": \\\"date\\\",          \\\"format\\\": \\\"EEE MMM dd HH:mm:ss Z YYYY\\\"        }      }    }  }}\n"),
		})
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
    opensearch:url:
        value: http://127.0.0.1:9200

```
```yaml
resources:
  # Create an index template
  template1:
    type: opensearch:IndexTemplate
    name: template_1
    properties:
      name: template_1
      body: |
        {  \"index_patterns\": [    \"your-pattern-here-*\"  ],  \"template\": {    \"settings\": {      \"index\": {        \"number_of_shards\": \"1\"      }    },    \"mappings\": {      \"_source\": {        \"enabled\": false      },      \"properties\": {        \"host_name\": {          \"type\": \"keyword\"        },        \"created_at\": {          \"type\": \"date\",          \"format\": \"EEE MMM dd HH:mm:ss Z YYYY\"        }      }    }  }}
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    opensearch:url:
        value: http://127.0.0.1:9200

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.opensearch.IndexTemplate;
import com.pulumi.opensearch.IndexTemplateArgs;
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
        // Create an index template
        var template1 = new IndexTemplate("template1", IndexTemplateArgs.builder()
            .name("template_1")
            .body("""
{  \"index_patterns\": [    \"your-pattern-here-*\"  ],  \"template\": {    \"settings\": {      \"index\": {        \"number_of_shards\": \"1\"      }    },    \"mappings\": {      \"_source\": {        \"enabled\": false      },      \"properties\": {        \"host_name\": {          \"type\": \"keyword\"        },        \"created_at\": {          \"type\": \"date\",          \"format\": \"EEE MMM dd HH:mm:ss Z YYYY\"        }      }    }  }}
            """)
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference
### Required

- `url` (String) OpenSearch URL

- `awsAccessKey` (String) The access key for use with AWS OpenSearch Service domains
- `awsAssumeRoleArn` (String) Amazon Resource Name of an IAM Role to assume prior to making AWS API calls.
- `awsAssumeRoleExternalId` (String) External ID configured in the IAM policy of the IAM Role to assume prior to making AWS API calls.
- `awsProfile` (String) The AWS profile for use with AWS OpenSearch Service domains
- `awsRegion` (String) The AWS region for use in signing of AWS OpenSearch requests. Must be specified in order to use AWS URL signing with AWS OpenSearch endpoint exposed on a custom DNS domain.
- `awsSecretKey` (String) The secret key for use with AWS OpenSearch Service domains
- `awsSignatureService` (String) AWS service name used in the credential scope of signed requests to OpenSearch.
- `awsToken` (String) The session token for use with AWS OpenSearch Service domains
- `cacertFile` (String) A Custom CA certificate
- `clientCertPath` (String) A X509 certificate to connect to OpenSearch
- `clientKeyPath` (String) A X509 key to connect to OpenSearch
- `healthcheck` (Boolean) Set the client healthcheck option for the OpenSearch client. Healthchecking is designed for direct access to the cluster.
- `hostOverride` (String) If provided, sets the 'Host' header of requests and the 'ServerName' for certificate validation to this value. See the documentation on connecting to OpenSearch via an SSH tunnel.
- `insecure` (Boolean) Disable SSL verification of API calls
- `opensearchVersion` (String) OpenSearch Version
- `password` (String) Password to use to connect to OpenSearch using basic auth
- `proxy` (String) Proxy URL to use for requests to OpenSearch.
- `signAwsRequests` (Boolean) Enable signing of AWS OpenSearch requests. The `url` must refer to AWS ES domain (`*.<region>.es.amazonaws.com`), or `awsRegion` must be specified explicitly.
- `sniff` (Boolean) Set the node sniffing option for the OpenSearch client. Client won't work with sniffing if nodes are not routable.
- `token` (String) A bearer token or ApiKey for an Authorization header, e.g. Active Directory API key.
- `tokenName` (String) The type of token, usually ApiKey or Bearer
- `username` (String) Username to use to connect to OpenSearch using basic auth
- `versionPingTimeout` (Number) Version ping timeout in seconds
## Authentication
### AWS authentication

The provider is flexible in the means of providing credentials for authentication with AWS OpenSearch domains. The following methods are supported, in this order, and explained below:

- Static credentials
- Assume role configuration
- Environment variables
- Shared credentials file

If a [custom domain](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/customendpoint.html) is being used (instead of the default, of the form `https://search-mydomain-1a2a3a4a5a6a7a8a9a0a9a8a7a.us-east-1.es.amazonaws.com`), please make sure to set `awsRegion` in the provider configuration.
#### Static credentials
##### IAM user management

If your AWS OpenSearch domain uses IAM user management, static credentials can be provided by adding an `awsAccessKey` and `awsSecretKey` in-line in the provider configuration. If applicable, you may also specify a `awsToken` value.

Example usage:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    opensearch:awsAccessKey:
        value: anaccesskey
    opensearch:awsSecretKey:
        value: asecretkey
    opensearch:awsToken:
        value: ""
    opensearch:url:
        value: https://search-foo-bar-pqrhr4w3u4dzervg41frow4mmy.us-east-1.es.amazonaws.com

```
##### HTTP basic auth

If your AWS OpenSearch domain uses an internal user database, static credentials can be provided by adding a `username` and `password` in-line in the provider configuration. Note: You will need to explicitly disabled request signing.

Example usage:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    opensearch:password:
        value: apassword
    opensearch:signAwsRequests:
        value: false
    opensearch:url:
        value: https://search-foo-bar-pqrhr4w3u4dzervg41frow4mmy.us-east-1.es.amazonaws.com
    opensearch:username:
        value: ausername

```
#### Assume role configuration

You can instruct the provider to assume a role in AWS before interacting with the cluster by setting the `awsAssumeRoleArn` variable.
When necessary, use the awsAssumeRoleExternalId to pass the extenral ID configured in the policy of the role for the provider to assume the role.

Example usage:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    opensearch:awsAssumeRoleArn:
        value: arn:aws:iam::012345678901:role/rolename
    opensearch:awsAssumeRoleExternalId:
        value: SecretID
    opensearch:url:
        value: https://search-foo-bar-pqrhr4w3u4dzervg41frow4mmy.us-east-1.es.amazonaws.com

```
#### Environment variables

You can provide your credentials via the `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY`, environment variables, representing your AWS Access Key and AWS Secret Key. If applicable, the `AWS_SESSION_TOKEN` environment variables is also supported.

Example usage:

```shell
$ export AWS_ACCESS_KEY_ID="anaccesskey"
$ export AWS_SECRET_ACCESS_KEY="asecretkey"
$ pulumi preview
```
#### AWS profile

You can specify a named profile that will be used for credentials (either static, or sts assumed role creds).  eg:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    opensearch:awsProfile:
        value: profilename
    opensearch:url:
        value: https://search-foo-bar-pqrhr4w3u4dzervg41frow4mmy.us-east-1.es.amazonaws.com

```
#### Shared Credentials file

You can use an AWS credentials file to specify your credentials. The default location is `$HOME/.aws/credentials` on Linux and macOS, or `%USERPROFILE%\.aws\credentials` for Windows users.

Please refer to the official [userguide](https://docs.aws.amazon.com/cli/latest/userguide/cli-config-files.html) for instructions on how to create the credentials file.
### Connecting to a cluster via an SSH Tunnel

If you need to connect to a cluster via an SSH tunnel (for example, to an AWS VPC Cluster), set the following configuration options in your provider:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    opensearch:hostOverride:
        value: vpc-<******>.us-east-1.es.amazonaws.com
    opensearch:url:
        value: https://localhost:9999

```

The `hostOverride` flag will set the `Host` header of requests to the cluster and the `ServerName` used for certificate validation. It is recommended to set this flag instead of `insecure = true`, which causes certificate validation to be skipped. Note that if both `hostOverride` and `insecure = true` are set, certificate validation will be skipped and the `Host` header will be overridden.
