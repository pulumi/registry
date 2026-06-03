---
# WARNING: this file was fetched from https://raw.githubusercontent.com/rpothin/pulumi-powerplatform/v0.3.3/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/rpothin/pulumi-powerplatform/blob/v0.3.3/docs/installation-configuration.md
title: Power Platform Installation & Configuration
meta_desc: How to install and configure the Pulumi Power Platform provider.
layout: package
---

To provision resources with the Pulumi Power Platform provider, you need Azure Active Directory credentials. Your credentials are never sent to Pulumi — they are used locally by the provider to authenticate against the Power Platform admin APIs.

## Installation

The Power Platform provider is available as a package in the following
Pulumi-supported languages:

### Python

[rpothin-powerplatform on PyPI](https://pypi.org/project/rpothin-powerplatform/)

```bash
pip install rpothin-powerplatform
```

### Node.js

[@rpothin/powerplatform on npm](https://www.npmjs.com/package/@rpothin/powerplatform)

```bash
npm install @rpothin/powerplatform
```

### .NET

[Rpothin.Powerplatform on NuGet](https://www.nuget.org/packages/Rpothin.Powerplatform/)

```bash
dotnet add package Rpothin.Powerplatform
```

### Go

[powerplatform on pkg.go.dev](https://pkg.go.dev/github.com/rpothin/pulumi-powerplatform/sdk/go/powerplatform)

```bash
go get github.com/rpothin/pulumi-powerplatform/sdk/go/powerplatform
```

### Java

[io.github.rpothin:powerplatform on Maven Central](https://central.sonatype.com/artifact/io.github.rpothin/powerplatform)

**Gradle**

```groovy
implementation "io.github.rpothin:powerplatform:<version>"
```

**Maven**

```xml
<dependency>
  <groupId>io.github.rpothin</groupId>
  <artifactId>powerplatform</artifactId>
  <version>&lt;version&gt;</version>
</dependency>
```

### Plugin installation

```bash
VERSION=$(curl -sL https://api.github.com/repos/rpothin/pulumi-powerplatform/releases/latest | grep '"tag_name"' | cut -d'"' -f4)
pulumi plugin install resource powerplatform $VERSION --server github://api.github.com/rpothin
```

## Authentication

The provider supports several authentication methods. For local development, the Azure CLI is the simplest option. For CI/CD pipelines, OIDC / Workload Identity is recommended.

### Azure CLI (recommended for local development)

Install the [Azure CLI](https://learn.microsoft.com/cli/azure/install-azure-cli), then log in:

```bash
az login
```

The provider automatically discovers CLI credentials via `DefaultAzureCredential`. No additional Pulumi configuration is required when using the CLI.

### DefaultAzureCredential (hosted environments)

When no explicit credentials are configured, the provider uses
[`DefaultAzureCredential`](https://learn.microsoft.com/python/api/azure-identity/azure.identity.defaultazurecredential),
which automatically tries multiple credential sources (environment variables, workload identity, managed identity, Azure CLI, and others) in a well-defined order. See the
[official documentation](https://learn.microsoft.com/python/api/azure-identity/azure.identity.defaultazurecredential)
for the full chain.

This makes the provider work automatically in most Azure-hosted environments without any explicit configuration.

### OIDC / Workload Identity (recommended for CI/CD)

Use federated credentials to authenticate from GitHub Actions (or other OIDC providers) without storing a client secret.

**1. Configure federated credentials in Azure:**

In your Azure AD app registration, add a federated identity credential for your GitHub repository:
- Issuer: `https://token.actions.githubusercontent.com`
- Subject: `repo:<owner>/<repo>:ref:refs/heads/main` (adjust for your branch or environment)
- Audience: `api://AzureADTokenExchange`

**2. Add permissions in your GitHub Actions workflow:**

```yaml
permissions:
  id-token: write
  contents: read
```

**3. Login using the `azure/login` action:**

```yaml
- uses: azure/login@v2
  with:
    client-id: ${{ secrets.AZURE_CLIENT_ID }}
    tenant-id: ${{ secrets.AZURE_TENANT_ID }}
    allow-no-subscriptions: true
```

After this step, the provider authenticates via `DefaultAzureCredential`, which picks up the Azure CLI session established by `azure/login` — no `clientSecret` is needed.

### Service Principal (explicit configuration)

For environments where you cannot use CLI or OIDC, configure a service principal explicitly.

**Via `pulumi config`:**

```bash
pulumi config set powerplatform:tenantId <AZURE_TENANT_ID>
pulumi config set powerplatform:clientId <AZURE_CLIENT_ID>
pulumi config set --secret powerplatform:clientSecret <AZURE_CLIENT_SECRET>
```

**Via environment variables:**

| Variable              | Description                       |
|-----------------------|-----------------------------------|
| `AZURE_TENANT_ID`    | Azure AD tenant (directory) ID    |
| `AZURE_CLIENT_ID`    | Application (client) ID           |
| `AZURE_CLIENT_SECRET` | Client secret value              |

## Configuration Reference

| Property       | Type   | Required | Description |
|---------------|--------|----------|-------------|
| `tenantId`    | string | No       | Azure AD Tenant ID. Falls back to `AZURE_TENANT_ID`. |
| `clientId`    | string | No       | Azure AD Application (Client) ID. Falls back to `AZURE_CLIENT_ID`. |
| `clientSecret`| string | No       | Azure AD Client Secret. Falls back to `AZURE_CLIENT_SECRET`. |

## Prerequisites

- A Microsoft Power Platform tenant with admin access
- An Azure AD app registration with the Power Platform API permission:
  `https://api.powerplatform.com/.default`
- Pulumi CLI v3 or later
- Python 3.10 or later available on `PATH`
