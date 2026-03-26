{{ $type := default "gcp" (.Get "type") }}

When developing locally, we recommend that you install the [Google Cloud SDK](https://cloud.google.com/sdk/install) and then [authorize access with a user account](https://cloud.google.com/sdk/docs/authorizing#authorizing_with_a_user_account). Next, Pulumi requires default application credentials to interact with your Google Cloud resources, so run `auth application-default login` command to obtain those credentials:

```bash
$ gcloud auth application-default login
```

To configure Pulumi to interact with your Google Cloud project, set it with the `pulumi config` command using the project's ID:

```bash
$ pulumi config set {{ $type }}:project your-gcp-project-id
```

You may also set your Google Cloud Project via environment variable (listed in order of precedence):

- `GOOGLE_PROJECT`
- `GOOGLE_CLOUD_PROJECT`
- `GCLOUD_PROJECT`
- `CLOUDSDK_CORE_PROJECT`

```bash
$ export GOOGLE_PROJECT=your-gcp-project-id
```
