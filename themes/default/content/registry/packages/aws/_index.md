---
title: AWS
meta_desc: Install and configure the Pulumi AWS provider — SDK installation for every language, credentials examples for every authentication method, and all 50 configuration options.
layout: package
aliases:
    - /registry/packages/aws/installation-configuration/
---

The Pulumi AWS provider can provision most of the cloud resources available in [Amazon Web Services](https://aws.amazon.com/). It uses the AWS SDK to manage and provision resources, and must be configured with credentials before it can deploy or update anything.

New to Pulumi and AWS? [Get started with AWS using our tutorial](/docs/get-started/aws).

{{< chooser language "typescript,python,go,csharp,java,yaml,hcl" / >}}

{{% choosable language typescript %}}

```typescript
import * as aws from "@pulumi/aws";

const bucket = new aws.s3.Bucket("my-bucket");

export const bucketName = bucket.bucket;
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_aws as aws

bucket = aws.s3.Bucket("my-bucket")

pulumi.export("bucketName", bucket.bucket)
```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		bucket, err := s3.NewBucket(ctx, "my-bucket", &s3.BucketArgs{})
		if err != nil {
			return err
		}

		ctx.Export("bucketName", bucket.Bucket)
		return nil
	})
}
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() =>
{
    var bucket = new Aws.S3.Bucket("my-bucket");

    return new Dictionary<string, object?>
    {
        ["bucketName"] = bucket.BucketName,
    };
});
```

{{% /choosable %}}

{{% choosable language java %}}

```java
import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.aws.s3.Bucket;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    private static void stack(Context ctx) {
        var bucket = new Bucket("my-bucket");

        ctx.export("bucketName", bucket.bucket());
    }
}
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: aws-example
runtime: yaml
resources:
  my-bucket:
    type: aws:s3:Bucket
outputs:
  bucketName: ${my-bucket.bucket}
```

{{% /choosable %}}

{{% choosable language hcl %}}

```hcl
resource "aws_s3_bucket" "my_bucket" {}

output "bucket_name" {
  value = aws_s3_bucket.my_bucket.bucket
}
```

{{% /choosable %}}

## Installation

The AWS provider publishes an SDK for every language. Add it to an existing Pulumi project with your language's package manager; YAML and HCL projects use [`pulumi package add`](/docs/iac/cli/commands/pulumi_package_add/) instead:

{{< chooser language "typescript,python,go,csharp,java,yaml,hcl" / >}}

{{% choosable language typescript %}}

```bash
npm install @pulumi/aws
```

{{% /choosable %}}

{{% choosable language python %}}

```bash
pip install pulumi-aws
```

{{% /choosable %}}

{{% choosable language go %}}

```bash
go get github.com/pulumi/pulumi-aws/sdk/v7
```

Import it in your program with `import "github.com/pulumi/pulumi-aws/sdk/v7/go/aws"`.

{{% /choosable %}}

{{% choosable language csharp %}}

```bash
dotnet add package Pulumi.Aws
```

{{% /choosable %}}

{{% choosable language java %}}

Add the dependency to your build file. For Maven, in `pom.xml`:

```xml
<dependency>
    <groupId>com.pulumi</groupId>
    <artifactId>aws</artifactId>
    <version>7.40.0</version>
</dependency>
```

For Gradle, in `build.gradle`:

```groovy
implementation 'com.pulumi:aws:7.40.0'
```

{{% /choosable %}}

{{% choosable language yaml %}}

```bash
pulumi package add aws
```

Reference resources by their type token, for example `aws:s3:Bucket`. Omit the version to add the latest, or append `@7.40.0` to pin one.

{{% /choosable %}}

{{% choosable language hcl %}}

```bash
pulumi package add aws
```

This makes the `aws_*` resource types available to your converted HCL. Omit the version to add the latest, or append `@7.40.0` to pin one.

{{% /choosable %}}

To install everything a project already declares — after cloning it from source control, for example — run [`pulumi install`](/docs/iac/cli/commands/pulumi_install/).

## Configuration

Provider settings can come from any of three places:

- A [Pulumi ESC](/docs/pulumi-cloud/esc/) environment, under `pulumiConfig`.
- [Stack configuration](/docs/iac/concepts/config/), set with `pulumi config set` and stored in `Pulumi.<stack-name>.yaml`.
- An explicit `aws.Provider` instance, constructed in your program and passed to the resources that should use it.

See [Configuration](/docs/concepts/config/) for how these sources combine. Some options can also be read from environment variables, noted in their descriptions in the [configuration reference](/registry/packages/aws/configuration-reference/).

The examples below cover the common credential setups. For every option the provider accepts — all 50 of them, each with a `pulumi env set` and `pulumi config set` example — see the [AWS configuration reference](/registry/packages/aws/configuration-reference/).

Whichever method you choose, you will need an [IAM identity](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users_create.html#id_users_create_console) with sufficient permissions for the resources your program manages, and you must set `aws:region`. Your AWS credentials are never sent to Pulumi Cloud — the provider uses the AWS SDK and the credentials in your environment to authenticate requests from your machine to AWS.

### Examples

Every example below is shown two ways: as a [Pulumi ESC](/docs/pulumi-cloud/esc/) environment, and as [stack configuration](/docs/iac/concepts/config/) — the `Pulumi.<stack-name>.yaml` file alongside your project. Pick whichever you use; the choice follows you down the page and across providers.

#### Shared credentials file

Recommended for local development. A plaintext `~/.aws/credentials` file holds one or more named profiles; this is Amazon's recommended way to manage multiple roles. Create it with [`aws configure`](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html), or by hand:

```ini
[default]
aws_access_key_id = <YOUR_DEFAULT_ACCESS_KEY_ID>
aws_secret_access_key = <YOUR_DEFAULT_SECRET_ACCESS_KEY>

[prod-account]
aws_access_key_id = <YOUR_PROD_ACCESS_KEY_ID>
aws_secret_access_key = <YOUR_PROD_SECRET_ACCESS_KEY>
```

Then point Pulumi at a profile.

{{< chooser configsource "esc,stack" / >}}

{{% choosable configsource esc %}}

Set it with the CLI:

```bash
$ pulumi env set myorg/aws-dev pulumiConfig.aws:region us-west-2
$ pulumi env set myorg/aws-dev pulumiConfig.aws:profile prod-account
```

That produces the following environment definition in `environments/aws-dev.yaml`:

```yaml
values:
  pulumiConfig:
    aws:region: us-west-2
    aws:profile: prod-account
```

{{% /choosable %}}

{{% choosable configsource stack %}}

Set it with the CLI:

```bash
$ pulumi config set aws:region us-west-2
$ pulumi config set aws:profile prod-account
```

That writes the following into `Pulumi.dev.yaml`:

```yaml
config:
  aws:region: us-west-2
  aws:profile: prod-account
```

{{% /choosable %}}

Omit `aws:profile` to use the `[default]` profile. The profile can also come from the `AWS_PROFILE` environment variable.

#### Static credentials

Useful for CI systems without an AWS CLI, or to [temporarily override](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-quickstart.html#cli-configure-quickstart-precedence) your credentials file.

{{< chooser configsource "esc,stack" / >}}

{{% choosable configsource esc %}}

Set it with the CLI:

```bash
$ pulumi env set myorg/aws-dev pulumiConfig.aws:region us-west-2
$ pulumi env set --secret myorg/aws-dev pulumiConfig.aws:accessKey AKIAIOSFODNN7EXAMPLE
$ pulumi env set --secret myorg/aws-dev pulumiConfig.aws:secretKey wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
```

That produces the following environment definition in `environments/aws-dev.yaml`:

```yaml
values:
  pulumiConfig:
    aws:region: us-west-2
    aws:accessKey:
      fn::secret: AKIAIOSFODNN7EXAMPLE
    aws:secretKey:
      fn::secret: wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
```

{{% /choosable %}}

{{% choosable configsource stack %}}

Set it with the CLI:

```bash
$ pulumi config set aws:region us-west-2
$ pulumi config set --secret aws:accessKey AKIAIOSFODNN7EXAMPLE
$ pulumi config set --secret aws:secretKey wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
```

That writes the following into `Pulumi.dev.yaml`:

```yaml
config:
  aws:region: us-west-2
  aws:accessKey:
    secure: AAABAOWyM7t3q1u...KpR8Nx1yVCjKPB2nJWg==
  aws:secretKey:
    secure: AAABAJ4nQm1xPz9...8sVdK2eLhTgYbXc5Rq==
```

{{% /choosable %}}

The same values can come from the `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY` and `AWS_REGION` environment variables instead. If you are using [temporary security credentials](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_use-resources.html), also set `aws:token` (or `AWS_SESSION_TOKEN`).

#### Assume a role

Assume an IAM role after authenticating — the usual pattern for cross-account deployments. `assumeRoles` is a list, so a chain of roles can be assumed in order.

{{< chooser configsource "esc,stack" / >}}

{{% choosable configsource esc %}}

Set it with the CLI:

```bash
$ pulumi env set myorg/aws-dev pulumiConfig.aws:assumeRoles[0].roleArn arn:aws:iam::123456789012:role/OrganizationAccountAccessRole
$ pulumi env set myorg/aws-dev pulumiConfig.aws:assumeRoles[0].sessionName pulumi-session
```

That produces the following environment definition in `environments/aws-dev.yaml`:

```yaml
values:
  pulumiConfig:
    aws:region: us-west-2
    aws:assumeRoles:
      - roleArn: arn:aws:iam::123456789012:role/OrganizationAccountAccessRole
        sessionName: pulumi-session
```

{{% /choosable %}}

{{% choosable configsource stack %}}

Set it with the CLI:

```bash
$ pulumi config set --path aws:assumeRoles[0].roleArn arn:aws:iam::123456789012:role/OrganizationAccountAccessRole
$ pulumi config set --path aws:assumeRoles[0].sessionName pulumi-session
```

That writes the following into `Pulumi.dev.yaml`:

```yaml
config:
  aws:region: us-west-2
  aws:assumeRoles:
    - roleArn: arn:aws:iam::123456789012:role/OrganizationAccountAccessRole
      sessionName: pulumi-session
```

{{% /choosable %}}

#### Web identity and OpenID Connect (OIDC)

Configure a role to assume plus a source for a web identity token — an OIDC ID token — and the provider exchanges it for temporary credentials. This lets Pulumi run on GitHub, GitLab or Azure DevOps without storing any long-lived AWS credentials. See AWS's [web identity federation](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_oidc.html) docs; configuring the OIDC trust relationship is a one-time setup.

On GitHub Actions you usually don't configure the token at all — Amazon's [`configure-aws-credentials`](https://github.com/aws-actions/configure-aws-credentials/tree/v4/#OIDC) action handles it:

```yaml
- uses: aws-actions/configure-aws-credentials@v4
  with:
    aws-region: us-west-2
    role-session-name: pulumi-session
    role-to-assume: arn:aws:iam::123456789012:role/GitHubActions
```

Elsewhere, configure `assumeRoleWithWebIdentity` with the role and the token source.

{{< chooser configsource "esc,stack" / >}}

{{% choosable configsource esc %}}

Set it with the CLI:

```bash
$ pulumi env set myorg/aws-dev pulumiConfig.aws:assumeRoleWithWebIdentity.roleArn arn:aws:iam::123456789012:role/GitHubActions
$ pulumi env set myorg/aws-dev pulumiConfig.aws:assumeRoleWithWebIdentity.webIdentityTokenFile /var/run/secrets/token
```

That produces the following environment definition in `environments/aws-dev.yaml`:

```yaml
values:
  pulumiConfig:
    aws:region: us-west-2
    aws:assumeRoleWithWebIdentity:
      roleArn: arn:aws:iam::123456789012:role/GitHubActions
      sessionName: pulumi-session
      webIdentityTokenFile: /var/run/secrets/token
```

{{% /choosable %}}

{{% choosable configsource stack %}}

Set it with the CLI:

```bash
$ pulumi config set --path aws:assumeRoleWithWebIdentity.roleArn arn:aws:iam::123456789012:role/GitHubActions
$ pulumi config set --path aws:assumeRoleWithWebIdentity.webIdentityTokenFile /var/run/secrets/token
```

That writes the following into `Pulumi.dev.yaml`:

```yaml
config:
  aws:region: us-west-2
  aws:assumeRoleWithWebIdentity:
    roleArn: arn:aws:iam::123456789012:role/GitHubActions
    sessionName: pulumi-session
    webIdentityTokenFile: /var/run/secrets/token
```

{{% /choosable %}}

Set either `webIdentityToken` or `webIdentityTokenFile`, not both.

#### Dynamically generated credentials

The one case with no stack-configuration equivalent: rather than storing credentials anywhere, have ESC mint short-lived ones on every run. First set up [OIDC between Pulumi and AWS](/docs/pulumi-cloud/oidc/aws/), then define an environment:

`environments/aws-dev.yaml`

```yaml
values:
  aws:
    login:
      fn::open::aws-login:
        oidc:
          duration: 1h
          roleArn: <your-oidc-iam-role-arn>
          sessionName: pulumi-environments-session
  pulumiConfig:
    aws:region: us-west-2
    aws:accessKey: ${aws.login.accessKeyId}
    aws:secretKey: ${aws.login.secretAccessKey}
    aws:token: ${aws.login.sessionToken}
```

Import it from your stack's configuration file:

`Pulumi.dev.yaml`

```yaml
environment:
  - aws-dev
```

Values under `pulumiConfig` are scoped to the `pulumi` run. If you need them exported to child processes as well, put them under `environmentVariables` instead — see [projecting environment variables](/docs/pulumi-cloud/esc/environments/#projecting-environment-variables).

Verify with `pulumi preview`. Make sure your local environment has no AWS credentials configured first; `aws s3 ls` should fail with *"Unable to locate credentials"*.

### Reference

Every configuration option the provider accepts — 50 of them, including the nested `assumeRoles`, `assumeRoleWithWebIdentity`, `defaultTags` and `ignoreTags` objects — lives on its own page:

**[AWS configuration reference →](/registry/packages/aws/configuration-reference/)**

Options are listed alphabetically, and each carries its own anchor — so you can link straight to one, and the page's "On this page" rail doubles as an index of every option name.


## Related packages

Pulumi offers components with simpler, higher-productivity APIs for several areas of AWS:

- [Amazon EKS](/registry/packages/eks/) — managed Kubernetes clusters
- [AWSx](/registry/packages/awsx/) — Crosswalk for AWS, covering API Gateway, CloudWatch, ECR, ECS, ELB, IAM, Lambda, VPC and more

## Migration guides

Breaking changes and upgrade instructions for each major version:

- [Migrating to v6.x.x](/registry/packages/aws/how-to-guides/6-0-migration/)
- [Migrating to v7.x.x](/registry/packages/aws/how-to-guides/7-0-migration/)
