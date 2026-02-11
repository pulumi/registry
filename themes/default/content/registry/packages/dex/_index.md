---
# WARNING: this file was fetched from https://raw.githubusercontent.com/kotaicode/pulumi-dex/v0.7.2/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/kotaicode/pulumi-dex/blob/v0.7.2/docs/_index.md
title: Dex Provider
meta_desc: A Pulumi provider for managing Dex (dexidp.io) resources via infrastructure as code.
layout: package
---

# Dex Provider

The Dex provider for Pulumi allows you to manage Dex (https://dexidp.io/) resources using infrastructure as code. This provider uses the Dex gRPC Admin API to manage OAuth2 clients and identity provider (IdP) connectors.

## Features

- **OAuth2 Client Management**: Create, update, and delete Dex OAuth2 clients
- **Generic Connector Support**: Manage any Dex connector type (OIDC, LDAP, SAML, etc.)
- **OIDC Connector Support**: First-class support for OIDC connectors with typed configuration
- **Azure/Entra ID Integration**: 
  - `AzureOidcConnector` - Uses generic OIDC connector (type: `oidc`)
  - `AzureMicrosoftConnector` - Uses Dex's Microsoft-specific connector (type: `microsoft`)
- **AWS Cognito Integration**: `CognitoOidcConnector` for managing Cognito user pools as IdPs

## Example Usage

### TypeScript

```typescript
import * as dex from "@kotaicode/pulumi-dex";
import * as fs from "fs";

// Configure the provider
const provider = new dex.Provider("dex", {
    host: "dex.example.com:5557",
    // For production, use TLS/mTLS:
    // caCert: fs.readFileSync("certs/ca.crt", "utf-8"),
    // clientCert: fs.readFileSync("certs/client.crt", "utf-8"),
    // clientKey: fs.readFileSync("certs/client.key", "utf-8"),
    // For development:
    insecureSkipVerify: true,
});

// Create an OAuth2 client
const client = new dex.Client("my-client", {
    clientId: "my-web-app",
    name: "My Web Application",
    redirectUris: ["https://app.example.com/callback"],
}, { provider });

// Create an Azure OIDC connector
const azureConnector = new dex.AzureOidcConnector("azure-oidc", {
    connectorId: "azure-oidc",
    name: "Azure AD",
    tenantId: "your-tenant-id",
    clientId: "your-azure-app-id",
    clientSecret: "your-azure-app-secret",
    redirectUri: "https://dex.example.com/dex/callback",
    scopes: ["openid", "email", "profile"],
}, { provider });

export const clientId = client.clientId;
export const clientSecret = client.secret;
```

### Go

```go
package main

import (
    "github.com/kotaicode/pulumi-provider-dex/sdk/go/dex"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
        // Configure the provider
        provider, err := dex.NewProvider(ctx, "dex", &dex.ProviderArgs{
            Host:              pulumi.String("dex.example.com:5557"),
            InsecureSkipVerify: pulumi.Bool(true),
        })
        if err != nil {
            return err
        }

        // Create an OAuth2 client
        client, err := dex.NewClient(ctx, "my-client", &dex.ClientArgs{
            ClientId:     pulumi.String("my-web-app"),
            Name:         pulumi.String("My Web Application"),
            RedirectUris: pulumi.StringArray{pulumi.String("https://app.example.com/callback")},
        }, pulumi.Provider(provider))
        if err != nil {
            return err
        }

        // Create an Azure OIDC connector
        azureConnector, err := dex.NewAzureOidcConnector(ctx, "azure-oidc", &dex.AzureOidcConnectorArgs{
            ConnectorId:  pulumi.String("azure-oidc"),
            Name:         pulumi.String("Azure AD"),
            TenantId:     pulumi.String("your-tenant-id"),
            ClientId:     pulumi.String("your-azure-app-id"),
            ClientSecret: pulumi.String("your-azure-app-secret"),
            RedirectUri:  pulumi.String("https://dex.example.com/dex/callback"),
            Scopes:       pulumi.StringArray{pulumi.String("openid"), pulumi.String("email"), pulumi.String("profile")},
        }, pulumi.Provider(provider))
        if err != nil {
            return err
        }

        ctx.Export("clientId", client.ClientId)
        ctx.Export("clientSecret", client.Secret)
        
        return nil
    })
}
```

### Python

```python
import pulumi
import pulumi_dex as dex

# Configure the provider
provider = dex.Provider("dex",
    host="dex.example.com:5557",
    insecure_skip_verify=True
)

# Create an OAuth2 client
client = dex.Client("my-client",
    client_id="my-web-app",
    name="My Web Application",
    redirect_uris=["https://app.example.com/callback"],
    opts=pulumi.ResourceOptions(provider=provider)
)

# Create an Azure OIDC connector
azure_connector = dex.AzureOidcConnector("azure-oidc",
    connector_id="azure-oidc",
    name="Azure AD",
    tenant_id="your-tenant-id",
    client_id="your-azure-app-id",
    client_secret="your-azure-app-secret",
    redirect_uri="https://dex.example.com/dex/callback",
    scopes=["openid", "email", "profile"],
    opts=pulumi.ResourceOptions(provider=provider)
)

pulumi.export("client_id", client.client_id)
pulumi.export("client_secret", client.secret)
```

## Resources

- **dex.Client** - Manage OAuth2 clients
- **dex.Connector** - Manage generic connectors (OIDC, LDAP, SAML, etc.)
- **dex.AzureOidcConnector** - Azure AD/Entra ID connector (OIDC)
- **dex.AzureMicrosoftConnector** - Azure AD/Entra ID connector (Microsoft-specific)
- **dex.CognitoOidcConnector** - AWS Cognito user pool connector

## Prerequisites

- Dex instance with gRPC API enabled
- For connector management: Dex must be started with `DEX_API_CONNECTORS_CRUD=true`
- Network access to Dex's gRPC endpoint (default port: 5557)

## Learn More

- [Installation and Configuration](installation-configuration.md)
- [GitHub Repository](https://github.com/kotaicode/pulumi-provider-dex)
- [Dex Documentation](https://dexidp.io/docs/)

