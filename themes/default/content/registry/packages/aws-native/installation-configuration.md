---
title: AWS Native Installation & Configuration
meta_desc: Set up credentials to use the Pulumi AWS Native Provider and choose configuration options to customize how the provider works.
layout: installation
---

<!-- markdownlint-disable url -->
[iam-user-console]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users_create.html#id_users_create_console
[iam-manage-keys]: https://docs.aws.amazon.com/general/latest/gr/managing-aws-access-keys.html
[configure-aws-cli]: https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html
<!-- markdownlint-enable url -->

{{< aws-resource-note >}}

## Installation

The AWS Native provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/aws-native`](https://www.npmjs.com/package/@pulumi/aws-native)
* Python: [`pulumi-aws-native`](https://pypi.org/project/pulumi-aws-native/)
* Go: [`github.com/pulumi/pulumi-aws-native/sdk/go/aws`](https://github.com/pulumi/pulumi-aws-native/tree/master/sdk/go/aws)
* .NET: [`Pulumi.AwsNative`](https://www.nuget.org/packages/Pulumi.AwsNative)
* Java: [`com.pulumi.awsnative`](https://search.maven.org/search?q=com.pulumi.aws-native)

## Setup

To provision resources with the Pulumi AWS provider, you need to have AWS credentials. You can use the instructions on if you plan to use AWS credentials from a shared credentials file (which the AWS CLI usually manages for you) or from an environment variable. For other credential options, see the [AWS documentation](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html).

Your AWS credentials are never sent to pulumi.com. Pulumi uses the AWS SDK and the credentials in your environment to authenticate requests from your computer to AWS.

### Get your credentials

First, make sure you have an IAM user in [the AWS console][iam-user-console] with **Programmatic access** and ensure it has sufficient permissions to deploy and manage your program's resources. If you know the precise resource types you wish to create and delete, we recommend restrciting your IAM user's access to just those types.

You'll also need [an access key][iam-manage-keys] for your user. There are two parts to each key, which you'll see in the IAM console after you create it:

* `<YOUR_ACCESS_KEY_ID>`: your access key's ID
* `<YOUR_SECRET_ACCESS_KEY>`: your access key's secret

{{% notes "info" %}}
If you are using [temporary security credentials](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_use-resources.html), you will also have to supply an `AWS_SESSION_TOKEN` value before you can use Pulumi to create resources on your behalf.
{{% /notes %}}

### Create a shared credentials file

A credentials file is a plaintext file on your machine that contains your access keys. The file must be named `credentials` and is located underneath `.aws/` directory in your home directory.  We recommend this approach because it supports Amazon's recommended approach for securely managing multiple roles.

#### Option 1: Use the CLI

To create this file using the CLI, [install the AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/installing.html). If you're using
[Homebrew](https://brew.sh/) on macOS, you can use the community-managed [awscli](http://formulae.brew.sh/formula/awscli) via `brew install awscli`.

After installing the CLI, configure it with your IAM credentials, typically using the `aws configure` command. For other configuration options, see the AWS article [Configuring the AWS CLI][configure-aws-cli].

```bash
$ aws configure
AWS Access Key ID [None]: <YOUR_ACCESS_KEY_ID>
AWS Secret Access Key [None]: <YOUR_SECRET_ACCESS_KEY>
Default region name [None]:
Default output format [None]:
```

Now you've created the `~/.aws/credentials` file and populated it with the expected settings.

#### Option 2: Create by hand

You can also create the shared credentials file by hand.  For example:

```ini
[default]
aws_access_key_id = <YOUR_ACCESS_KEY_ID>
aws_secret_access_key = <YOUR_SECRET_ACCESS_KEY>
```

If you want to specify multiple profiles, those are listed in different sections:

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

In this case, you will need to set the `AWS_PROFILE` environment variable to the name of the profile to use.

### Set environment variables

We recommend using a shared credentials file for most development. However, if you need to temporarily override your credentials file, you can use environment variables. You can do this to quickly switch to a different access key or to configure AWS access from within an environment that might not have an AWS CLI, such as a continuous integration/continuous delivery (CI/CD) system.

To authenticate using environment variable, set them in your terminal:

{{< chooser os "linux,macos,windows" >}}
{{% choosable os linux %}}

```bash
$ export AWS_ACCESS_KEY_ID=<YOUR_ACCESS_KEY_ID>
$ export AWS_SECRET_ACCESS_KEY=<YOUR_SECRET_ACCESS_KEY>
```

{{% /choosable %}}

{{% choosable os macos %}}

```bash
$ export AWS_ACCESS_KEY_ID=<YOUR_ACCESS_KEY_ID>
$ export AWS_SECRET_ACCESS_KEY=<YOUR_SECRET_ACCESS_KEY>
```

{{% /choosable %}}

{{% choosable os windows %}}

```powershell
> $env:AWS_ACCESS_KEY_ID = "<YOUR_ACCESS_KEY_ID>"
> $env:AWS_SECRET_ACCESS_KEY = "<YOUR_SECRET_ACCESS_KEY>"
```

{{% /choosable %}}
{{< /chooser >}}

### Set up multiple profiles

As an optional step, if you have multiple AWS profiles set up, you can specify a different profile to use with Pulumi through one of the following methods:

* Set `AWS_PROFILE` as an environment variable
* After creating your project, run `pulumi config set aws-native:profile <profilename>`

## Configuration options

{{% notes type="info" %}}
Some configuration options are not supported in AWS Native. You can check on the status of a configuration option you need by visiting our backlog of [Github issues](https://github.com/pulumi/pulumi-aws-native/issues?q=is%3Aissue+is%3Aopen+label%3Aprovider-config); feel free to add your vote to an issue by adding a 👍 reaction.
{{% /notes %}}

Use `pulumi config set aws-native:<option>` or pass options to the [constructor of `new aws-native.Provider`](/registry/packages/aws-native/api-docs/provider).

| Option                                                                                                        | Required/Optional | Description                                                                                                                                                                                                                                                                                                                                                                                                        |
|---------------------------------------------------------------------------------------------------------------|-------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `region`                                                                                                      | Required          | The region where AWS operations will take place. Examples are `us-east-1`, `us-west-2`, etc.                                                                                                                                                                                                                                                                                                                       |
| ~~`allowedAccountIds`~~ ([Not supported yet](https://github.com/pulumi/pulumi-aws-native/issues/105))         | Optional          | List of allowed AWS account IDs to prevent you from mistakenly using an incorrect one (and potentially end up destroying a live environment). Conflicts with `forbiddenAccountIds`.                                                                                                                                                                                                                                |
| `accessKey`                                                                                                   | Optional          | The access key for API operations. You can retrieve this from the ‘Security & Credentials’ section of the AWS console.                                                                                                                                                                                                                                                                                             |
| `assumeRole`                                                                                                  | Optional          | A JSON object representing an IAM role to assume.  To set these nested properties, see docs on [structured configuration](/docs/intro/concepts/config#structured-configuration), for example `pulumi config set --path aws-native:assumeRole.roleArn arn:aws:iam::058111598222:role/OrganizationAccountAccessRole`. The object contains the properties marked with a ↳ below:                             |
| ↳ `durationSeconds`                                                                                           | Optional          | Number of seconds to restrict the assume role session duration.                                                                                                                                                                                                                                                                                                                                                    |
| ↳ `externalId`                                                                                                | Optional          | External identifier to use when assuming the role.                                                                                                                                                                                                                                                                                                                                                                 |
| ↳ `policy`                                                                                                    | Optional          | IAM Policy JSON describing further restricting permissions for the IAM Role being assumed.                                                                                                                                                                                                                                                                                                                         |
| ↳ `policyArns`                                                                                                | Optional          | Set of Amazon Resource Names (ARNs) of IAM Policies describing further restricting permissions for the IAM Role being assumed.                                                                                                                                                                                                                                                                                     |
| ↳ `roleArn`                                                                                                   | Optional          | Amazon Resource Name (ARN) of the IAM Role to assume.                                                                                                                                                                                                                                                                                                                                                              |
| ↳ `sessionName`                                                                                               | Optional          | Session name to use when assuming the role.                                                                                                                                                                                                                                                                                                                                                                        |
| ↳ `tags`                                                                                                      | Optional          | Map of assume role session tags.                                                                                                                                                                                                                                                                                                                                                                                   |
| ↳ `transitiveTagKeys`                                                                                         | Optional          | Set of assume role session tag keys to pass to any subsequent sessions.                                                                                                                                                                                                                                                                                                                                            |
| `dynamodbEndpoint`                                                                                            | Optional          | Use this to override the default endpoint URL constructed from the `region`. It’s typically used to connect to dynamodb-local.                                                                                                                                                                                                                                                                                     |
| ~~`endpoints`~~ ([Not supported yet](https://github.com/pulumi/pulumi-aws-native/issues/108))                 | Optional          | Configuration block for customizing service endpoints.                                                                                                                                                                                                                                                                                                                                                             |
| ~~`forbiddenAccountIds`~~  ([Not supported yet](https://github.com/pulumi/pulumi-aws-native/issues/109))      | Optional          | List of forbidden AWS account IDs to prevent you from mistakenly using the wrong one (and potentially end up destroying a live environment). Conflicts with `allowedAccountIds`.                                                                                                                                                                                                                                   |
| ~~`defaultTags`~~ ([Not supported yet](https://github.com/pulumi/pulumi-aws-native/issues/107))               | Optional          | A JSON block with resource tag settings to apply across all resources handled by this provider. Additional tags can be added/overridden at a per resource level. The object contains the properties marked with a ↳ below:                                                                                                                                                                                         |
| ↳ ~~`tags`~~                                                                                                  | Optional          | A key value pair of tags to apply across all resources.                                                                                                                                                                                                                                                                                                                                                            |
| ~~`ignoreTags`~~ ([Not supported yet](https://github.com/pulumi/pulumi-aws-native/issues/110))                | Optional          | A JSON block with resource tag settings to ignore across all resources handled by this provider (except any individual service tag resources such as `aws.ec2.Tag`) for situations where external systems are managing certain resource tags. The object contains the properties marked with a ↳ below:                                                                                                            |
| ↳ ~~`keys`~~                                                                                                  | Optional          | A list of exact resource tag keys to ignore across all resources handled by this provider. This configuration prevents Pulumi from returning the tag in any `tags` properties and displaying any diffs for the tag value. If any resource still has this tag key configured in the `tags` argument, it will display a perpetual diff until the tag is removed from the argument or `ignoreChanges` is also used.   |
| ↳ ~~`keyPrefixes`~~                                                                                           | Optional          | A list of resource tag key prefixes to ignore across all resources handled by this provider. This configuration prevents Pulumi from returning the tag in any `tags` properties and displaying any diffs for the tag value. If any resource still has this tag key configured in the `tags` argument, it will display a perpetual diff until the tag is removed from the argument or `ignoreChanges` is also used. |
| ~~`insecure`~~ ([Not supported yet](https://github.com/pulumi/pulumi-aws-native/issues/111))                  | Optional          | Explicitly allow the provider to perform "insecure" SSL requests. If omitted, the default value is `false`.                                                                                                                                                                                                                                                                                                        |
| `kinesisEndpoint`                                                                                             | Optional          | Use this to override the default endpoint URL constructed from the `region`. It's typically used to connect to kinesalite.                                                                                                                                                                                                                                                                                         |
| ~~`maxRetries`~~ ([Not supported yet](https://github.com/pulumi/pulumi-aws-native/issues/112))                | Optional          | The maximum number of times an AWS API request is being executed. If the API request still fails, an error is thrown.                                                                                                                                                                                                                                                                                              |
| `profile`                                                                                                     | Optional          | The profile for API operations. If not set, the default profile created with `aws configure` will be used.                                                                                                                                                                                                                                                                                                         |
| ~~`s3ForcePathStyle`~~ ([Not supported yet](https://github.com/pulumi/pulumi-aws-native/issues/113))          | Optional          | Set this to true to force the request to use path-style addressing, i.e., `http://s3.amazonaws.com/BUCKET/KEY`. By default, the S3 client will use virtual hosted bucket addressing when possible (`http://BUCKET.s3.amazonaws.com/KEY`). Specific to the Amazon S3 service.                                                                                                                                       |
| `secretKey`                                                                                                   | Optional          | The secret key for API operations. You can retrieve this from the 'Security & Credentials' section of the AWS console.                                                                                                                                                                                                                                                                                             |
| `sharedCredentialsFile`                                                                                       | Optional          | The path to the shared credentials file. If not set this defaults to `~~/.aws/credentials`.                                                                                                                                                                                                                                                                                                                        |
| ~~`skipCredentialsValidation`~~ ([Not supported yet](https://github.com/pulumi/pulumi-aws-native/issues/114)) | Optional          | Skip the credentials validation via STS API. Used for AWS API implementations that do not have STS available/implemented.                                                                                                                                                                                                                                                                                          |
| ~~`skipGetEc2Platforms`~~ ([Not supported yet](https://github.com/pulumi/pulumi-aws-native/issues/115))       | Optional          | Skip getting the supported EC2 platforms. Used by users that don't have `ec2:DescribeAccountAttributes` permissions.                                                                                                                                                                                                                                                                                               |
| ~~`skipMetadataApiCheck`~~ ([Not supported yet](https://github.com/pulumi/pulumi-aws-native/issues/116))      | Optional          | Skip the AWS Metadata API check. Useful for AWS API implementations that do not have a metadata API endpoint. Setting to true prevents Pulumi from authenticating via the Metadata API. You may need to use other authentication methods like static credentials, configuration variables, or environment variables.                                                                                               |
| ~~`skipRegionValidation`~~ ([Not supported yet](https://github.com/pulumi/pulumi-aws-native/issues/117))      | Optional          | Skip static validation of region name. Used by users of alternative AWS-like APIs or users w/ access to regions that are not public (yet).                                                                                                                                                                                                                                                                         |
| ~~`skipRequestingAccountId`~~ ([Not supported yet](https://github.com/pulumi/pulumi-aws-native/issues/118))   | Optional          | Skip requesting the account ID. Used for AWS API implementations that do not have IAM/STS API and/or metadata API.                                                                                                                                                                                                                                                                                                 |
| `token`                                                                                                       | Optional          | Use this to set an MFA token. It can also be sourced from the `AWS_SESSION_TOKEN` environment variable.                                                                                                                                                                                                                                                                                                            |
