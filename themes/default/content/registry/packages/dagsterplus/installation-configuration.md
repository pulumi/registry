---
# WARNING: this file was fetched from https://raw.githubusercontent.com/dagster-io/pulumi-dagsterplus/v0.0.4/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/dagster-io/pulumi-dagsterplus/blob/v0.0.4/docs/installation-configuration.md
title: Dagster+ Provider Installation & Configuration
meta_desc: Information on how to install and configure the Dagster+ provider for Pulumi.
layout: package
---

## Installation

The Dagster+ provider is available as a package in the following Pulumi languages:

* Python: [`pulumi-dagsterplus`](https://pypi.org/project/pulumi-dagsterplus/)
* YAML: Install the [Pulumi CLI](https://www.pulumi.com/docs/install/) and reference the provider directly in your YAML programs.

### Python

Add the package to your project with `pip`:

```bash
pip install pulumi-dagsterplus
```

…or add it to your `requirements.txt`:

```
pulumi-dagsterplus
```

The provider plugin is downloaded automatically from the
[`dagster-io/pulumi-dagsterplus`](https://github.com/dagster-io/pulumi-dagsterplus)
GitHub releases the first time you run `pulumi up` or `pulumi preview`.

## Configuration

The following configuration points are available for the `dagsterplus` provider:

- `dagsterplus:apiToken` (string, secret) — A Dagster+ API token. May also be set
  via the `DAGSTER_CLOUD_API_TOKEN` environment variable. Generate a token at
  Dagster+ → Account Settings → API Tokens.
- `dagsterplus:organization` (string) — The Dagster+ organization name (e.g.
  `my-org` for `my-org.dagster.cloud`). May also be set via the
  `DAGSTER_CLOUD_ORGANIZATION` environment variable.
- `dagsterplus:baseUrl` (string, optional) — Override the base URL for the
  Dagster+ API. Defaults to `https://{organization}.dagster.cloud`. Only needed
  for on-prem or custom deployments.

### Setting configuration via environment variables

The recommended way to configure the provider is through environment variables,
which keeps your API token out of your Pulumi state and program source:

```bash
export DAGSTER_CLOUD_API_TOKEN=your-user-token
export DAGSTER_CLOUD_ORGANIZATION=your-org-name
```

### Setting configuration via the Pulumi CLI

You can also set configuration with `pulumi config set`. Use `--secret` for the
API token so it is stored encrypted:

```bash
pulumi config set dagsterplus:organization your-org-name
pulumi config set --secret dagsterplus:apiToken your-user-token
```

### Setting configuration in code

Configuration may also be passed explicitly when constructing the provider:

{{< chooser language "python,yaml" >}}

{{% choosable language python %}}

```python
import pulumi_dagsterplus as dagsterplus

provider = dagsterplus.Provider("dagsterplus",
    api_token="your-user-token",
    organization="your-org-name",
)
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
resources:
  dagsterplus:
    type: pulumi:providers:dagsterplus
    properties:
      apiToken: your-user-token
      organization: your-org-name
```

{{% /choosable %}}

{{< /chooser >}}
