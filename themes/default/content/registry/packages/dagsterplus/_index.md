---
# WARNING: this file was fetched from https://raw.githubusercontent.com/dagster-io/pulumi-dagsterplus/v0.0.4/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/dagster-io/pulumi-dagsterplus/blob/v0.0.4/docs/_index.md
title: Dagster+
meta_desc: Provides an overview of the Dagster+ provider for Pulumi.
layout: package
---

## Overview

The Dagster+ provider for Pulumi can be used to provision and manage resources in [Dagster+](https://dagster.io/cloud), the managed cloud platform for Dagster. Use it to manage deployments, code locations, agents, teams, secrets, and more — all as code.

## Installation

The Dagster+ provider is available as a package in the following languages:

* Python: [`pulumi-dagsterplus`](https://pypi.org/project/pulumi-dagsterplus/)
* YAML: Install the [Pulumi CLI](https://www.pulumi.com/docs/install/) and reference the provider directly in your YAML programs.

## Authentication

The Dagster+ provider requires a Dagster+ user token and your organization name to authenticate. You can generate a user token from your [Dagster+ account settings](https://dagster.cloud).

Configure the provider via environment variables:

```bash
export DAGSTER_CLOUD_API_TOKEN=your-user-token
export DAGSTER_CLOUD_ORGANIZATION=your-org-name
```

Or pass them explicitly in your Pulumi program:

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

## Example Usage

Create a Dagster+ team:

{{< chooser language "python,yaml" >}}

{{% choosable language python %}}

```python
import pulumi_dagsterplus as dagsterplus

team = dagsterplus.Team("my-team",
    name="platform-engineers",
)
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
resources:
  my-team:
    type: dagsterplus:Team
    properties:
      name: platform-engineers
```

{{% /choosable %}}

{{< /chooser >}}
