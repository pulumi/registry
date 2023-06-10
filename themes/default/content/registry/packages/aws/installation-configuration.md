---
title: AWS Classic Installation & Configuration
meta_desc: How to set up credentials to use the Pulumi AWS Classic Provider and choose configuration options to tailor the provider to suit your use case.
layout: package
---

{{< aws-resource-note >}}

## Installation

The AWS Classic provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/aws`](https://www.npmjs.com/package/@pulumi/aws)
* Python: [`pulumi-aws`](https://pypi.org/project/pulumi-aws/)
* Go: [`github.com/pulumi/pulumi-aws/sdk/v5`](https://github.com/pulumi/pulumi-aws#go)
* .NET: [`Pulumi.Aws`](https://www.nuget.org/packages/Pulumi.Aws)
* Java: [`com.pulumi.aws`](https://search.maven.org/search?q=com.pulumi.aws)

## Credentials

1. [Create an IAM user in the AWS console](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users_create.html#id_users_create_console) with programmatic access and ensure it has sufficient permissions to deploy and manage your Pulumi program’s resources.
2. [Set up AWS credentials](https://docs.aws.amazon.com/general/latest/gr/aws-sec-cred-types.html#access-keys-and-secret-access-keys) for your user.

{{% notes "info" %}}
If you are using [temporary security credentials](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_use-resources.html), you will also have to supply an `AWS_SESSION_TOKEN` value before you can use Pulumi to create resources on your behalf.
{{% /notes %}}

Your AWS credentials are never sent to pulumi.com. Pulumi uses the AWS SDK and the credentials in your environment to authenticate requests from your computer to AWS.

## Configuration

There are a few different ways you can configure your AWS credentials to work with Pulumi.

### Set credentials as environment variables

You can authenticate using environment variables.
Doing so will [temporarily override the settings in your credentials file](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-quickstart.html#cli-configure-quickstart-precedence).

{{< chooser os "linux,macos,windows" >}}
{{% choosable os linux %}}

```bash
$ export AWS_ACCESS_KEY_ID=<YOUR_ACCESS_KEY_ID>
$ export AWS_SECRET_ACCESS_KEY=<YOUR_SECRET_ACCESS_KEY>
$ export AWS_REGION=<YOUR_AWS_REGION> # e.g.`ap-south-1`
```

{{% /choosable %}}

{{% choosable os macos %}}

```bash
$ export AWS_ACCESS_KEY_ID=<YOUR_ACCESS_KEY_ID>
$ export AWS_SECRET_ACCESS_KEY=<YOUR_SECRET_ACCESS_KEY>
$ export AWS_REGION=<YOUR_AWS_REGION> # e.g.`ap-south-1`
```

{{% /choosable %}}

{{% choosable os windows %}}

```powershell
> $env:AWS_ACCESS_KEY_ID = "<YOUR_ACCESS_KEY_ID>"
> $env:AWS_SECRET_ACCESS_KEY = "<YOUR_SECRET_ACCESS_KEY>"
> $env:AWS_REGION = "<YOUR_AWS_REGION>"
```

{{% /choosable %}}
{{< /chooser >}}

You may alternatively set the AWS region in your Pulumi.yaml:

```bash
$ pulumi config set aws:region <your-region> # e.g.`ap-south-1`
```

### Create a shared credentials file using the AWS CLI

1. [Install the AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/installing.html)

2. Configure your [AWS credentials](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-quickstart.html#cli-configure-quickstart-config).

    ```bash
    $ aws configure
    AWS Access Key ID [None]: <YOUR_ACCESS_KEY_ID>
    AWS Secret Access Key [None]: <YOUR_SECRET_ACCESS_KEY>
    Default region name [None]: <YOUR_AWS_REGION>
    Default output format [None]:
    ```

Your AWS credentials file is now located in your home directory at `.aws/credentials`.

You can also create the shared credentials file by hand.  For example:

```ini
[default]
aws_access_key_id = <YOUR_ACCESS_KEY_ID>
aws_secret_access_key = <YOUR_SECRET_ACCESS_KEY>
```

### Set up multiple profiles

As an optional step, you can [set up multiple profiles](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-quickstart.html#cli-configure-quickstart-profiles)
Here’s what that looks like in your ~/.aws/credentials file:

```ini
[default]
aws_access_key_id = <YOUR_DEFAULT_ACCESS_KEY_ID>
aws_secret_access_key = <YOUR_DEFAULT_SECRET_ACCESS_KEY>

[test-account]
aws_access_key_id = <YOUR_TEST_ACCESS_KEY_ID>
aws_secret_access_key = <YOUR_TEST_SECRET_ACCESS_KEY>

[prod-account]
aws_access_key_id = <YOUR_PROD_ACCESS_KEY_ID>
aws_secret_access_key = <YOUR_PROD_SECRET_ACCESS_KEY>
```

You can specify which profile to use with Pulumi through one of the following methods:

* Set AWS_PROFILE as an environment variable

    ```bash
    $ export AWS_PROFILE=<YOUR_PROFILE_NAME>
    ```

* Set `aws:profile` in your Pulumi.yaml

    ```bash
    pulumi config set aws:profile <profilename>
    ```

### Authenticating via EC2 Instance Metadata?

As of pulumi-aws v3.28.1, the default behaviour for the provider [was changed](https://github.com/pulumi/pulumi-aws/blob/master/CHANGELOG_OLD.md#3281-2021-02-10) to disable MetadataApiCheck by default. This means,
you need to do either of the following

1. When using the default provider:

    ```
    pulumi config set aws:skipMetadataApiCheck false
    ```

1. When using a named provider

    ```typescript
    const myProvider = new aws.Provider("named-provider", {
      // other config
      skipMetadataApiCheck: false,
    });
    ```

    ```csharp
    var provider = new Aws.Provider("named-provider", new Aws.ProviderArgs
    {
      // other config
      SkipMetadataApiCheck = false,
    });
    ```

    ```go
    provider, err := aws.NewProvider(ctx, "named-provider", &aws.ProviderArgs{
        // other config
        SkipMetadataApiCheck: pulumi.Bool(false),
    })
    ```

    ```python
    provider = pulumi_aws.Provider('named-provider', skip_metadata_api_check=False)
    ```

## Configuration options

Use `pulumi config set aws:<option>` or pass options to the [constructor of `new aws.Provider`](../api-docs/provider).

| Option | Required? | Description |
| - | - | - |
| `region` | Required | The region where AWS operations will take place. Examples are `us-east-1`, `us-west-2`, etc. |
| `allowedAccountIds` | Optional | List of allowed AWS account IDs to prevent you from mistakenly using an incorrect one (and potentially end up destroying a live environment). Conflicts with `forbiddenAccountIds`. |
| `accessKey` | Optional | The access key for API operations. You can retrieve this from the ‘Security & Credentials’ section of the AWS console. |
| `assumeRole` | Optional | A JSON object representing an IAM role to assume.  To set these nested properties, see docs on [structured configuration](/docs/concepts/config#structured-configuration), for example `pulumi config set --path aws:assumeRole.roleArn arn:aws:iam::058111598222:role/OrganizationAccountAccessRole`. The object contains the properties marked with a ↳ below: |
| ↳ `durationSeconds` | Optional |  Number of seconds to restrict the assume role session duration. |
| ↳ `externalId` | Optional | External identifier to use when assuming the role. |
| ↳ `policy` | Optional | IAM Policy JSON describing further restricting permissions for the IAM Role being assumed. |
| ↳ `policyArns` | Optional | Set of Amazon Resource Names (ARNs) of IAM Policies describing further restricting permissions for the IAM Role being assumed. |
| ↳ `roleArn` | Optional | Amazon Resource Name (ARN) of the IAM Role to assume. |
| ↳ `sessionName` | Optional | Session name to use when assuming the role. |
| ↳ `tags` | Optional | Map of assume role session tags. |
| ↳ `transitiveTagKeys` | Optional | Set of assume role session tag keys to pass to any subsequent sessions. |
| `dynamodbEndpoint` | Optional | Use this to override the default endpoint URL constructed from the `region`. It’s typically used to connect to dynamodb-local. |
| `forbiddenAccountIds` | Optional | List of forbidden AWS account IDs to prevent you from mistakenly using the wrong one (and potentially end up destroying a live environment). Conflicts with `allowedAccountIds`. |
| `defaultTags` | Optional | A JSON block with resource tag settings to apply across all resources handled by this provider. Additional tags can be added/overridden at a per resource level. The object contains the properties marked with a ↳ below: |
| ↳ `tags` | Optional | A key value pair of tags to apply across all resources. |
| `ignoreTags` | Optional | A JSON block with resource tag settings to ignore across all resources handled by this provider (except any individual service tag resources such as `aws.ec2.Tag`) for situations where external systems are managing certain resource tags. The object contains the properties marked with a ↳ below: |
| ↳ `keys` | Optional | A list of exact resource tag keys to ignore across all resources handled by this provider. This configuration prevents Pulumi from returning the tag in any `tags` properties and displaying any diffs for the tag value. If any resource still has this tag key configured in the `tags` argument, it will display a perpetual diff until the tag is removed from the argument or `ignoreChanges` is also used. |
| ↳ `keyPrefixes` | Optional | A list of resource tag key prefixes to ignore across all resources handled by this provider. This configuration prevents Pulumi from returning the tag in any `tags` properties and displaying any diffs for the tag value. If any resource still has this tag key configured in the `tags` argument, it will display a perpetual diff until the tag is removed from the argument or `ignoreChanges` is also used. |
| `insecure` | Optional | Explicitly allow the provider to perform "insecure" SSL requests. If omitted, the default value is `false`. |
| `kinesisEndpoint` | Optional | Use this to override the default endpoint URL constructed from the `region`. It's typically used to connect to kinesalite. |
| `maxRetries` | Optional | The maximum number of times an AWS API request is being executed. If the API request still fails, an error is thrown. |
| `profile` | Optional | The profile for API operations. If not set, the default profile created with `aws configure` will be used. |
| `s3ForcePathStyle` | Optional | Set this to true to force the request to use path-style addressing, i.e., `http://s3.amazonaws.com/BUCKET/KEY`. By default, the S3 client will use virtual hosted bucket addressing when possible (`http://BUCKET.s3.amazonaws.com/KEY`). Specific to the Amazon S3 service. |
| `secretKey` |  Optional | The secret key for API operations. You can retrieve this from the 'Security & Credentials' section of the AWS console. |
| `sharedCredentialsFile` | Optional | The path to the shared credentials file. If not set this defaults to `~/.aws/credentials`. |
| `skipCredentialsValidation` | Optional | Skip the credentials validation via STS API. Used for AWS API implementations that do not have STS available/implemented. |
| `skipGetEc2Platforms` | Optional | Skip getting the supported EC2 platforms. Used by users that don't have `ec2:DescribeAccountAttributes` permissions. |
| `skipMetadataApiCheck` | Optional | Skip the AWS Metadata API check. Useful for AWS API implementations that do not have a metadata API endpoint. Setting to true prevents Pulumi from authenticating via the Metadata API. You may need to use other authentication methods like static credentials, configuration variables, or environment variables. |
| `skipRegionValidation` | Optional | Skip static validation of region name. Used by users of alternative AWS-like APIs or users w/ access to regions that are not public (yet). |
| `skipRequestingAccountId` | Optional | Skip requesting the account ID. Used for AWS API implementations that do not have IAM/STS API and/or metadata API. |
| `token` | Optional | Use this to set an MFA token. It can also be sourced from the `AWS_SESSION_TOKEN` environment variable. |
