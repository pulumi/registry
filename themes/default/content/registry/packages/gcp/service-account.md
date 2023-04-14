---
title: Installation & Configuration Service Account
meta_desc: Provides an overview of how set up a Google Cloud Platform Service Account
           with Pulumi.
---

[Google Cloud Platform Credentials]: https://console.cloud.google.com/apis/credentials

Using a Google service account allows you to use Pulumi in a non-interactive setting (for example CI/CD systems, where a person can not complete the normal `gcloud auth application-default login` flow). A service account can also be used when developing locally to ensure a specific set of scoped credentials not tied to a user account are used. This can be useful even when developing locally to give you more control over the account role used for deployment.

To use a service account with Pulumi you will need to provide the Google Cloud Platform Provider with a [Google service account private key in JSON format].
You can create and download credentials using the [Google Cloud Platform Credentials] page on the Google Cloud Platform Console.

> Your credentials are only used to authenticate with Google Cloud APIs on your behalf. Your credentials are never sent to pulumi.com.

## Using your credentials

To communicate your credentials to the Pulumi Google Cloud Platform Provider,
you can do any of the following:

### Set the `GOOGLE_CREDENTIALS` environment variable

The `GOOGLE_CREDENTIALS` variable can accept _either_ the contents _or_ the filepath pointing at the location of your credentials file.

{{< chooser os "linux,macos,windows" >}}
{{% choosable os linux %}}

```bash
export GOOGLE_CREDENTIALS=$(cat credentials.json)
```

or

```bash
export GOOGLE_CREDENTIALS=<path-to-credentials.json>
```

{{% /choosable %}}

{{% choosable os macos %}}

```bash
export GOOGLE_CREDENTIALS=$(cat credentials.json)
```

or

```bash
export GOOGLE_CREDENTIALS=<path-to-credentials.json>
```

{{% /choosable %}}

{{% choosable os windows %}}

Windows Powershell

```powershell
$env:GOOGLE_CREDENTIALS=cat credentials.json
```

or

```powershell
$env:GOOGLE_CREDENTIALS=<path-tocredentials.json>
```

{{% /choosable %}}
{{< /chooser >}}

### Set the credentials in your Pulumi.yaml

You may also set your credentials in your Pulumi.yaml config file:

{{< chooser os "linux,macos,windows" >}}
{{% choosable os linux %}}

```bash
pulumi config set gcp:credentials <your-credentials> # contents or filepath
```

{{% /choosable %}}

{{% choosable os macos %}}

```bash
pulumi config set gcp:credentials <your-credentials> # contents or filepath
```

{{% /choosable %}}

{{% choosable os windows %}}

Windows Powershell

```powershell
pulumi config set gcp:credentials <your-credentials> # contents or filepath
```

{{% /choosable %}}
{{< /chooser >}}
