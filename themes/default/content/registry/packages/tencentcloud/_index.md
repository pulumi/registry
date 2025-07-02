---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/tencentcloudstack/tencentcloud/1.82.6/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Tencentcloud Provider
meta_desc: Provides an overview on how to configure the Pulumi Tencentcloud provider.
layout: package
---

## Generate Provider

The Tencentcloud provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider tencentcloudstack/tencentcloud
```
## Overview

The TencentCloud provider is used to interact with many resources supported by [TencentCloud](https://intl.cloud.tencent.com).
The provider needs to be configured with the proper credentials before it can be used.

Use the navigation on the left to read about the available resources.

> **Note:** Pulumi 0.12.x supported began with provider version 1.9.0 (June 18, 2019).
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    tencentcloud:region:
        value: ap-guangzhou
    tencentcloud:secretId:
        value: my-secret-id
    tencentcloud:secretKey:
        value: my-secret-key

```
```typescript
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    tencentcloud:region:
        value: ap-guangzhou
    tencentcloud:secretId:
        value: my-secret-id
    tencentcloud:secretKey:
        value: my-secret-key

```
```python
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    tencentcloud:region:
        value: ap-guangzhou
    tencentcloud:secretId:
        value: my-secret-id
    tencentcloud:secretKey:
        value: my-secret-key

```
```csharp
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    tencentcloud:region:
        value: ap-guangzhou
    tencentcloud:secretId:
        value: my-secret-id
    tencentcloud:secretKey:
        value: my-secret-key

```
```go
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    tencentcloud:region:
        value: ap-guangzhou
    tencentcloud:secretId:
        value: my-secret-id
    tencentcloud:secretKey:
        value: my-secret-key

```
```yaml
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    tencentcloud:region:
        value: ap-guangzhou
    tencentcloud:secretId:
        value: my-secret-id
    tencentcloud:secretKey:
        value: my-secret-key

```
```java
Example currently unavailable in this language
```
{{% /choosable %}}
{{< /chooser >}}
## Authentication

The TencentCloud provider offers a flexible means of providing credentials for authentication.
The following methods are supported, in this order, and explained below:

- Static credentials
- Environment variables
- Assume role
- Assume role with SAML
- Assume role with OIDC
- Shared credentials
- Enable pod OIDC
- Cam role name
- MFA certification
### Static credentials

!> **Warning:** Hard-coding credentials into any Pulumi configuration is not
recommended, and risks secret leakage should this file ever be committed to a
public version control system.

Static credentials can be provided by adding an `secretId` `secretKey` and `region` in-line in the tencentcloud provider configuration:

Usage:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    tencentcloud:region:
        value: ap-guangzhou
    tencentcloud:secretId:
        value: my-secret-id
    tencentcloud:secretKey:
        value: my-secret-key

```

Use `allowedAccountIds` or `forbiddenAccountIds`

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    tencentcloud:allowedAccountIds:
        value:
            - "100023201586"
            - "100023201349"
    tencentcloud:region:
        value: ap-guangzhou
    tencentcloud:secretId:
        value: my-secret-id
    tencentcloud:secretKey:
        value: my-secret-key

```

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    tencentcloud:forbiddenAccountIds:
        value:
            - "100023201223"
    tencentcloud:region:
        value: ap-guangzhou
    tencentcloud:secretId:
        value: my-secret-id
    tencentcloud:secretKey:
        value: my-secret-key

```
### Environment variables

You can provide your credentials via `TENCENTCLOUD_SECRET_ID` and `TENCENTCLOUD_SECRET_KEY` environment variables,
representing your TencentCloud Secret Id and Secret Key respectively. `TENCENTCLOUD_REGION` is also used, if applicable:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:

```

Usage:

```shell
$ export TENCENTCLOUD_SECRET_ID="my-secret-id"
$ export TENCENTCLOUD_SECRET_KEY="my-secret-key"
$ export TENCENTCLOUD_REGION="ap-guangzhou"
$ pulumi preview
```
### Assume role

If provided with an assume role, Pulumi will attempt to assume this role using the supplied credentials. Assume role can be provided by adding an `roleArn`, `sessionName`, `sessionDuration`, `policy`(optional) and `externalId`(optional) in-line in the tencentcloud provider configuration:

Usage:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    tencentcloud:region:
        value: ap-guangzhou
    tencentcloud:secretId:
        value: my-secret-id
    tencentcloud:secretKey:
        value: my-secret-key

```

Combining MFA

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    tencentcloud:region:
        value: ap-guangzhou
    tencentcloud:secretId:
        value: my-secret-id
    tencentcloud:secretKey:
        value: my-secret-key

```

The `roleArn`, `sessionName`, `sessionDuration` and `externalId` can also provided via `TENCENTCLOUD_ASSUME_ROLE_ARN`, `TENCENTCLOUD_ASSUME_ROLE_SESSION_NAME`, `TENCENTCLOUD_ASSUME_ROLE_SESSION_DURATION` and `TENCENTCLOUD_ASSUME_ROLE_EXTERNAL_ID` environment variables.

The `serialNumber`, `tokenCode` can also provided via `TENCENTCLOUD_ASSUME_ROLE_SERIAL_NUMBER`, `TENCENTCLOUD_ASSUME_ROLE_TOKEN_CODE` environment variables.

Usage:

```shell
$ export TENCENTCLOUD_SECRET_ID="my-secret-id"
$ export TENCENTCLOUD_SECRET_KEY="my-secret-key"
$ export TENCENTCLOUD_REGION="ap-guangzhou"
$ export TENCENTCLOUD_ASSUME_ROLE_ARN="my-role-arn"
$ export TENCENTCLOUD_ASSUME_ROLE_SESSION_NAME="my-session-name"
$ export TENCENTCLOUD_ASSUME_ROLE_SESSION_DURATION=3600

$ export TENCENTCLOUD_ASSUME_ROLE_SERIAL_NUMBER="my-serial-number"
$ export TENCENTCLOUD_ASSUME_ROLE_TOKEN_CODE="my-token-code"
$ pulumi preview
```
### Assume role with SAML

If provided with an assume role with SAML, Pulumi will attempt to assume this role using the supplied credentials. Assume role can be provided by adding an `roleArn`, `sessionName`, `sessionDuration`, `samlAssertion` and `principalArn` in-line in the tencentcloud provider configuration:

> **Note:** Assume-role-with-SAML is a no-AK auth type, and there is no need setting secretId and secretKey while using it.

Usage:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:

```

The `roleArn`, `sessionName`, `sessionDuration`, `samlAssertion`, `principalArn` can also provided via `TENCENTCLOUD_ASSUME_ROLE_ARN`, `TENCENTCLOUD_ASSUME_ROLE_SESSION_NAME`, `TENCENTCLOUD_ASSUME_ROLE_SESSION_DURATION`, `TENCENTCLOUD_ASSUME_ROLE_SAML_ASSERTION` and `TENCENTCLOUD_ASSUME_ROLE_PRINCIPAL_ARN` environment variables.

Usage:

```shell
$ export TENCENTCLOUD_ASSUME_ROLE_ARN="my-role-arn"
$ export TENCENTCLOUD_ASSUME_ROLE_SESSION_NAME="my-session-name"
$ export TENCENTCLOUD_ASSUME_ROLE_SESSION_DURATION=3600
$ export TENCENTCLOUD_ASSUME_ROLE_SAML_ASSERTION="my-saml-assertion"
$ export TENCENTCLOUD_ASSUME_ROLE_PRINCIPAL_ARN="my-principal-arn"
$ pulumi preview
```
### Assume role with OIDC

If provided with an assume role with OIDC, Pulumi will attempt to assume this role using the supplied credentials. Assume role can be provided by adding an `roleArn`, `sessionName`, `sessionDuration` and `webIdentityToken` in-line in the tencentcloud provider configuration:

> **Note:** Assume-role-with-OIDC is a no-AK auth type, and there is no need setting secretId and secretKey while using it.

Usage:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:

```

The `providerId`, `roleArn`, `sessionName`, `sessionDuration`, `webIdentityToken` can also provided via `TENCENTCLOUD_ASSUME_ROLE_PROVIDER_ID`, `TENCENTCLOUD_ASSUME_ROLE_ARN`, `TENCENTCLOUD_ASSUME_ROLE_SESSION_NAME`, `TENCENTCLOUD_ASSUME_ROLE_SESSION_DURATION` and `TENCENTCLOUD_ASSUME_ROLE_WEB_IDENTITY_TOKEN` environment variables.

Usage:

```shell
$ export TENCENTCLOUD_SECRET_ID="my-secret-id"
$ export TENCENTCLOUD_SECRET_KEY="my-secret-key"
$ export TENCENTCLOUD_ASSUME_ROLE_SESSION_DURATION=3600
$ export TENCENTCLOUD_ASSUME_ROLE_WEB_IDENTITY_TOKEN="my-web-identity-token"
$ export TENCENTCLOUD_ASSUME_ROLE_PROVIDER_ID="OIDC"
$ pulumi preview
```
### Enable pod OIDC

Configure the TencentCloud Provider with TKE OIDC.

> **Note:** Must ensure CAM OIDC provider and WEBHOOK component are created successfully.

Usage:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    tencentcloud:enablePodOidc:
        value: true

```
### Cam role name

If provided with a Cam role name, Pulumi will just access the metadata URL: `http://metadata.tencentyun.com/latest/meta-data/cam/security-credentials/<cam_role_name>` to obtain the STS credential. The CVM Instance Role also can be set using the `TENCENTCLOUD_CAM_ROLE_NAME` environment variables.

> **Note:** Cam-role-name is used to grant the role entity the permissions to access services and resources and perform operations in Tencent Cloud. You can associate the CAM role with a CVM instance to call other Tencent Cloud APIs from the instance using the periodically updated temporary Security Token Service (STS) key.

> **Note:** Cam-role-name is a no-AK auth type, and there is no need setting secretId and secretKey while using it.

Usage:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    tencentcloud:camRoleName:
        value: my-cam-role-name

```

It can also be authenticated together with method Assume role. Authentication process: Perform CAM authentication first, then proceed with Assume role authentication.

Usage:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    tencentcloud:camRoleName:
        value: my-cam-role-name

```
### MFA certification

If provided with MFA certification, Pulumi will attempt to use the provided credentials for MFA authentication.

Usage:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    tencentcloud:region:
        value: ap-guangzhou
    tencentcloud:secretId:
        value: my-secret-id
    tencentcloud:secretKey:
        value: my-secret-key

```

The `serialNumber`, `tokenCode`, `durationSeconds` can also provided via `TENCENTCLOUD_MFA_CERTIFICATION_SERIAL_NUMBER`, `TENCENTCLOUD_MFA_CERTIFICATION_TOKEN_CODE`, `TENCENTCLOUD_MFA_CERTIFICATION_DURATION_SECONDS` environment variables.

Usage:

```shell
$ export TENCENTCLOUD_MFA_CERTIFICATION_SERIAL_NUMBER="my-serial-number"
$ export TENCENTCLOUD_MFA_CERTIFICATION_TOKEN_CODE="my-token-code"
$ export TENCENTCLOUD_MFA_CERTIFICATION_DURATION_SECONDS=1800
$ pulumi preview
```
### CDC cos usage

You can set the cos domain by setting the environment variable `TENCENTCLOUD_COS_DOMAIN`, and configure the cdc scenario as follows:

> **Note:** Please note that not all cos resources are supported. Please pay attention to the prompts for each resource.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    tencentcloud:cosDomain:
        value: https://cluster_xxx.cos-cdc.ap-guangzhou.myqcloud.com/
    tencentcloud:region:
        value: ap-guangzhou
    tencentcloud:secretId:
        value: my-secret-id
    tencentcloud:secretKey:
        value: my-secret-key

```
```typescript
import * as pulumi from "@pulumi/pulumi";

const region = "ap-guangzhou";
const cdcId = "cluster_xxx";
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    tencentcloud:cosDomain:
        value: https://cluster_xxx.cos-cdc.ap-guangzhou.myqcloud.com/
    tencentcloud:region:
        value: ap-guangzhou
    tencentcloud:secretId:
        value: my-secret-id
    tencentcloud:secretKey:
        value: my-secret-key

```
```python
import pulumi

region = "ap-guangzhou"
cdc_id = "cluster_xxx"
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    tencentcloud:cosDomain:
        value: https://cluster_xxx.cos-cdc.ap-guangzhou.myqcloud.com/
    tencentcloud:region:
        value: ap-guangzhou
    tencentcloud:secretId:
        value: my-secret-id
    tencentcloud:secretKey:
        value: my-secret-key

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;

return await Deployment.RunAsync(() =>
{
    var region = "ap-guangzhou";

    var cdcId = "cluster_xxx";

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    tencentcloud:cosDomain:
        value: https://cluster_xxx.cos-cdc.ap-guangzhou.myqcloud.com/
    tencentcloud:region:
        value: ap-guangzhou
    tencentcloud:secretId:
        value: my-secret-id
    tencentcloud:secretKey:
        value: my-secret-key

```
```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_ := "ap-guangzhou"
		_ := "cluster_xxx"
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
    tencentcloud:cosDomain:
        value: https://cluster_xxx.cos-cdc.ap-guangzhou.myqcloud.com/
    tencentcloud:region:
        value: ap-guangzhou
    tencentcloud:secretId:
        value: my-secret-id
    tencentcloud:secretKey:
        value: my-secret-key

```
```yaml
variables:
  region: ap-guangzhou
  cdcId: cluster_xxx
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    tencentcloud:cosDomain:
        value: https://cluster_xxx.cos-cdc.ap-guangzhou.myqcloud.com/
    tencentcloud:region:
        value: ap-guangzhou
    tencentcloud:secretId:
        value: my-secret-id
    tencentcloud:secretKey:
        value: my-secret-key

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
        final var region = "ap-guangzhou";

        final var cdcId = "cluster_xxx";

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

The `cosDomain` can also provided via `TENCENTCLOUD_COS_DOMAIN` environment variables.

Usage:

```shell
$ export TENCENTCLOUD_SECRET_ID="my-secret-id"
$ export TENCENTCLOUD_SECRET_KEY="my-secret-key"
$ export TENCENTCLOUD_REGION="ap-guangzhou"
$ export TENCENTCLOUD_COS_DOMAIN="https://cluster-xxxxxx.cos-cdc.ap-guangzhou.myqcloud.com/"
$ pulumi preview
```
### Shared credentials

You can use [Tencent Cloud credentials](https://www.tencentcloud.com/document/product/1013/33464) to specify your credentials. The default location is `$HOME/.tccli` on Linux and macOS, And `"%USERPROFILE%\.tccli"` on Windows. You can optionally specify a different location in the Pulumi configuration by providing the `sharedCredentialsDir` argument or using the `TENCENTCLOUD_SHARED_CREDENTIALS_DIR` environment variable. This method also supports a `profile` configuration and matching `TENCENTCLOUD_PROFILE` environment variable:

Usage:

On Linux/MacOS

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    tencentcloud:profile:
        value: default
    tencentcloud:sharedCredentialsDir:
        value: /Users/tf_user/.tccli

```

On Windows

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    tencentcloud:profile:
        value: default
    tencentcloud:sharedCredentialsDir:
        value: C:\Users\tf_user\.tccli

```
## Configuration Reference

In addition to generic provider arguments (e.g. alias and version), the following arguments are supported in the TencentCloud provider configuration:

* `secretId` - (Optional) This is the TencentCloud secret id. It must be provided, but it can also be sourced from the `TENCENTCLOUD_SECRET_ID` environment variable.
* `secretKey` - (Optional) This is the TencentCloud secret key. It must be provided, but it can also be sourced from the `TENCENTCLOUD_SECRET_KEY` environment variable.
* `securityToken` - (Optional) TencentCloud security token of temporary access credentials. It can also be sourced from the `TENCENTCLOUD_SECURITY_TOKEN` environment variable. Notice: for supported products, please refer to: [temporary key supported products](https://intl.cloud.tencent.com/document/product/598/10588).
* `region` - (Optional) This is the TencentCloud region. It must be provided, but it can also be sourced from the `TENCENTCLOUD_REGION` environment variables. The default input value is `ap-guangzhou`.
* `sharedCredentialsDir` - (Optional) The directory of the shared credentials. It can also be sourced from the `TENCENTCLOUD_SHARED_CREDENTIALS_DIR` environment variable. If not set this defaults to ~/.tccli.
* `profile` - (Optional) The profile name as set in the shared credentials. It can also be sourced from the `TENCENTCLOUD_PROFILE` environment variable. If not set, the default profile created with `tccli configure` will be used.
* `assumeRole` - (Optional, Available in 1.33.1+) An `assumeRole` block (documented below). If provided, pulumi will attempt to assume this role using the supplied credentials. Only one `assumeRole` block may be in the configuration.
* `assumeRoleWithSaml` - (Optional, Available in 1.81.111+) An `assumeRoleWithSaml` block (documented below). If provided, pulumi will attempt to assume this role using the supplied credentials. Only one `assumeRoleWithSaml` block may be in the configuration.
* `enablePodOidc` - (Optional, Available in 1.81.117+) Whether to enable pod oidc.
* `assumeRoleWithWebIdentity` - (Optional, Available in 1.81.111+) An `assumeRoleWithWebIdentity` block (documented below). If provided, pulumi will attempt to assume this role using the supplied credentials. Only one `assumeRoleWithWebIdentity` block may be in the configuration.
* `protocol` - (Optional, Available in 1.37.0+) The protocol of the API request. Valid values: `HTTP` and `HTTPS`. Default is `HTTPS`.
* `domain` - (Optional, Available in 1.37.0+) The root domain of the API request, Default is `tencentcloudapi.com`.
* `camRoleName` - (Optional, Available in 1.81.117+) The name of the CVM instance CAM role. It can be sourced from the `TENCENTCLOUD_CAM_ROLE_NAME` environment variable.
* `allowedAccountIds` - (Optional) List of allowed TencentCloud account IDs to prevent you from mistakenly using the wrong one (and potentially end up destroying a live environment). Conflicts with `forbiddenAccountIds`, If use `assumeRoleWithSaml` or `assumeRoleWithWebIdentity`, it is not supported.
* `forbiddenAccountIds` - (Optional) List of forbidden TencentCloud account IDs to prevent you from mistakenly using the wrong one (and potentially end up destroying a live environment). Conflicts with `allowedAccountIds`, If use `assumeRoleWithSaml` or `assumeRoleWithWebIdentity`, it is not supported.

The nested `assumeRole` block supports the following:
* `roleArn` - (Required) The ARN of the role to assume. It can also be sourced from the `TENCENTCLOUD_ASSUME_ROLE_ARN` environment variable.
* `sessionName` - (Required) The session name to use when making the AssumeRole call. It can also be sourced from the `TENCENTCLOUD_ASSUME_ROLE_SESSION_NAME` environment variable.
* `sessionDuration` - (Required) The duration of the session when making the AssumeRole call. Its value ranges from 0 to 43200(seconds), and default is 7200 seconds. It can also be sourced from the `TENCENTCLOUD_ASSUME_ROLE_SESSION_DURATION` environment variable.
* `policy` - (Optional) A more restrictive policy to apply to the temporary credentials. This gives you a way to further restrict the permissions for the resulting temporary security credentials. You cannot use the passed policy to grant permissions that are in excess of those allowed by the access policy of the role that is being assumed.
* `externalId` - (Optional) External role ID, which can be obtained by clicking the role name in the CAM console. It can contain 2-128 letters, digits, and symbols (=,.@\:/-). Regex: [\\w+=,.@\:/-]*. It can be sourced from the `TENCENTCLOUD_ASSUME_ROLE_EXTERNAL_ID`.

The nested `assumeRoleWithSaml` block supports the following:
* `roleArn` - (Required) The ARN of the role to assume. It can also be sourced from the `TENCENTCLOUD_ASSUME_ROLE_ARN` environment variable.
* `sessionName` - (Required) The session name to use when making the AssumeRole call. It can also be sourced from the `TENCENTCLOUD_ASSUME_ROLE_SESSION_NAME` environment variable.
* `sessionDuration` - (Required) The duration of the session when making the AssumeRole call. Its value ranges from 0 to 43200(seconds), and default is 7200 seconds. It can also be sourced from the `TENCENTCLOUD_ASSUME_ROLE_SESSION_DURATION` environment variable.
* `samlAssertion` - (Required) SAML assertion information encoded in base64. It can be sourced from the `TENCENTCLOUD_ASSUME_ROLE_SAML_ASSERTION`.
* `principalArn` - (Required) Player Access Description Name. It can be sourced from the `TENCENTCLOUD_ASSUME_ROLE_PRINCIPAL_ARN`.

The nested `assumeRoleWithWebIdentity` block supports the following:
* `providerId` - (Optional) Identity provider name. It can be sourced from the `TENCENTCLOUD_ASSUME_ROLE_PROVIDER_ID`, Default is OIDC.
* `roleArn` - (Required) The ARN of the role to assume. It can also be sourced from the `TENCENTCLOUD_ASSUME_ROLE_ARN` environment variable.
* `sessionName` - (Required) The session name to use when making the AssumeRole call. It can also be sourced from the `TENCENTCLOUD_ASSUME_ROLE_SESSION_NAME` environment variable.
* `sessionDuration` - (Required) The duration of the session when making the AssumeRole call. Its value ranges from 0 to 43200(seconds), and default is 7200 seconds. It can also be sourced from the `TENCENTCLOUD_ASSUME_ROLE_SESSION_DURATION` environment variable.
* `webIdentityToken` - (Required) OIDC token issued by IdP. It can be sourced from the  `TENCENTCLOUD_ASSUME_ROLE_WEB_IDENTITY_TOKEN`.