---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumiverse/pulumi-mssql/v0.1.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: MSSQL
meta_desc: Provides an overview of the MSSQL Provider for Pulumi.
layout: package
---

The MSSQL provider for Pulumi can be used to provision and configure objects in
Microsoft SQL Server and Azure SQL instances. The MSSQL provider must be
configured with credentials to deploy and update resources in MSSQL.


## Example

{{< chooser language "typescript,python,go,csharp,yaml" >}}


{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as azure from "@pulumi/azure";
import * as mssql from "@pulumiverse/mssql";

const current = azure.core.getClientConfig({});

const resourceGroup = new azure.core.ResourceGroup("resource-group", {
    name: "pulumi-mssql-test",
    location: "eastus",
});

const server = new azure.mssql.Server("server", {
    name: "pulumi-mssql-test",
    resourceGroupName: resourceGroup.name,
    location: resourceGroup.location,
    version: "12.0",
    minimumTlsVersion: "1.2",
    administratorLogin: "sadmin",
    administratorLoginPassword: "Str0ngPa$word12",
    azureadAdministrator: {
        loginUsername: "pulumi-app",
        objectId: current.then(current => current.objectId),
        tenantId: current.then(current => current.tenantId),
    },
});

const databaseFirewallRule = new azure.mssql.FirewallRule("database-firewall-rule", {
    name: "azure",
    serverId: server.id,
    startIpAddress: "0.0.0.0",
    endIpAddress: "0.0.0.0",
});

const database = new azure.mssql.Database("database", {
    name: "db",
    serverId: server.id,
    licenseType: "LicenseIncluded",
    maxSizeGb: 2,
    skuName: "Basic",
});

const providerMssql = new mssql.Provider("provider-mssql", {
    hostname: server.fullyQualifiedDomainName,
    azureAuth: {},
}, {
    dependsOn: [
        database,
        databaseFirewallRule,
    ],
});

const databaseId = mssql.getDatabaseOutput({
    name: database.name,
}, {
    provider: providerMssql
});
const databaseRoleOwner = mssql.getDatabaseRoleOutput({
    databaseId: databaseId.id,
    name: "db_owner",
}, {
    provider: providerMssql
});

const dbUser = new mssql.SqlLogin("db-user", {
    name: "example",
    password: "Str0ngPa$word12",
}, {
    provider: providerMssql,
});
const dbRoleMember = new mssql.DatabaseRoleMember("db-role-member", {
    roleId: databaseRoleOwner.id,
    memberId: dbUser.id,
}, {
    provider: providerMssql,
});
```
{{% /choosable %}}

{{% choosable language python %}}
```python
import pulumi
import pulumi_azure as azure
import pulumiverse_mssql as mssql

current = azure.core.get_client_config()
resource_group = azure.core.ResourceGroup(
    "resource-group", name="pulumi-mssql-test", location="eastus"
)

server = azure.mssql.Server(
    "server",
    name="pulumi-mssql-test",
    resource_group_name=resource_group.name,
    location=resource_group.location,
    version="12.0",
    minimum_tls_version="1.2",
    administrator_login="sadmin",
    administrator_login_password="Str0ngPa$word12",
    azuread_administrator=azure.mssql.ServerAzureadAdministratorArgs(
        login_username="pulumi-app",
        object_id=current.object_id,
        tenant_id=current.tenant_id,
    ),
)

database = azure.mssql.Database(
    "database",
    name="db",
    server_id=server.id,
    license_type="LicenseIncluded",
    max_size_gb=2,
    sku_name="Basic",
)

database_firewall_rule = azure.mssql.FirewallRule(
    "database-firewall-rule",
    name="azure",
    server_id=server.id,
    start_ip_address="0.0.0.0",
    end_ip_address="0.0.0.0",
)

provider_mssql = mssql.Provider(
    "provider-mssql",
    hostname=server.fully_qualified_domain_name,
    azure_auth=mssql.ProviderAzureAuthArgs(),
    opts=pulumi.ResourceOptions(
        depends_on=[
            database,
            database_firewall_rule,
        ]
    ),
)

database_id = mssql.get_database_output(
    name=database.name,
    opts=pulumi.InvokeOptions(provider=provider_mssql),
)

database_role_owner = mssql.get_database_role_output(
    database_id=database_id.id,
    name="db_owner",
    opts=pulumi.InvokeOptions(provider=provider_mssql),
)

db_user = mssql.SqlLogin(
    "db-user",
    name="example",
    password="Str0ngPa$word12",
    opts=pulumi.ResourceOptions(provider=provider_mssql),
)

db_role_member = mssql.DatabaseRoleMember(
    "db-role-member",
    role_id=database_role_owner.id,
    member_id=db_user.id,
    opts=pulumi.ResourceOptions(provider=provider_mssql),
)
```
{{% /choosable %}}

{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-azure/sdk/v5/go/azure/core"
	azuresql "github.com/pulumi/pulumi-azure/sdk/v5/go/azure/mssql"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-mssql/sdk/go/mssql"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		current, err := core.GetClientConfig(ctx, nil, nil)
		if err != nil {
			return err
		}

		resourceGroup, err := core.NewResourceGroup(ctx, "resource-group", &core.ResourceGroupArgs{
			Name:     pulumi.String("pulumi-mssql-test"),
			Location: pulumi.String("eastus"),
		})
		if err != nil {
			return err
		}

		server, err := azuresql.NewServer(ctx, "server", &azuresql.ServerArgs{
			Name:                       pulumi.String("pulumi-mssql-test"),
			ResourceGroupName:          resourceGroup.Name,
			Location:                   resourceGroup.Location,
			Version:                    pulumi.String("12.0"),
			MinimumTlsVersion:          pulumi.String("1.2"),
			AdministratorLogin:         pulumi.String("sadmin"),
			AdministratorLoginPassword: pulumi.String("Str0ngPa$word12"),
			AzureadAdministrator: &azuresql.ServerAzureadAdministratorArgs{
				LoginUsername: pulumi.String("pulumi-app"),
				ObjectId:      pulumi.String(current.ObjectId),
				TenantId:      pulumi.String(current.TenantId),
			},
		})
		if err != nil {
			return err
		}

		database, err := azuresql.NewDatabase(ctx, "database", &azuresql.DatabaseArgs{
			Name:        pulumi.String("db"),
			ServerId:    server.ID(),
			LicenseType: pulumi.String("LicenseIncluded"),
			MaxSizeGb:   pulumi.Int(2),
			SkuName:     pulumi.String("Basic"),
		})
		if err != nil {
			return err
		}

		databaseFirewallRule, err := azuresql.NewFirewallRule(ctx, "database-firewall-rule", &azuresql.FirewallRuleArgs{
			Name:           pulumi.String("azure"),
			ServerId:       server.ID(),
			StartIpAddress: pulumi.String("0.0.0.0"),
			EndIpAddress:   pulumi.String("0.0.0.0"),
		})
		if err != nil {
			return err
		}
		providerMssql, err := mssql.NewProvider(ctx, "provider-mssql", &mssql.ProviderArgs{
			Hostname:  server.FullyQualifiedDomainName,
			AzureAuth: mssql.ProviderAzureAuthArgs{},
		}, pulumi.DependsOn([]pulumi.Resource{
			database,
			databaseFirewallRule,
		}))
		if err != nil {
			return err
		}

		databaseId := mssql.LookupDatabaseOutput(ctx, mssql.LookupDatabaseOutputArgs{
			Name: database.Name,
		}, pulumi.Provider(providerMssql))

		databaseRoleOwner := mssql.LookupDatabaseRoleOutput(ctx, mssql.LookupDatabaseRoleOutputArgs{
			DatabaseId: databaseId.Id(),
			Name:       pulumi.String("db_owner"),
		}, pulumi.Provider(providerMssql))

		dbUser, err := mssql.NewSqlLogin(ctx, "db-user", &mssql.SqlLoginArgs{
			Name:     pulumi.String("example"),
			Password: pulumi.String("Str0ngPa$word12"),
		}, pulumi.Provider(providerMssql))
		if err != nil {
			return err
		}

		_, err = mssql.NewDatabaseRoleMember(ctx, "db-role-member", &mssql.DatabaseRoleMemberArgs{
			RoleId:   databaseRoleOwner.Id(),
			MemberId: dbUser.ID(),
		}, pulumi.Provider(providerMssql))
		if err != nil {
			return err
		}
		return nil
	})
}
```
{{% /choosable %}}

{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Azure = Pulumi.Azure;
using Mssql = Pulumiverse.Mssql;

return await Deployment.RunAsync(async () =>
{
    var current = await Azure.Core.GetClientConfig.InvokeAsync();

    var resourceGroup = new Azure.Core.ResourceGroup(
        "resource-group",
        new() { Name = "pulumi-mssql-test", Location = "eastus", }
    );

    var server = new Azure.MSSql.Server(
        "server",
        new()
        {
            Name = "pulumi-mssql-test",
            ResourceGroupName = resourceGroup.Name,
            Location = resourceGroup.Location,
            Version = "12.0",
            MinimumTlsVersion = "1.2",
            AdministratorLogin = "sadmin",
            AdministratorLoginPassword = "Str0ngPa$word12",
            AzureadAdministrator = new Azure.MSSql.Inputs.ServerAzureadAdministratorArgs
            {
                LoginUsername = "pulumi-app",
                ObjectId = current.ObjectId,
                TenantId = current.TenantId,
            },
        }
    );

    var database = new Azure.MSSql.Database(
        "database",
        new()
        {
            Name = "db",
            ServerId = server.Id,
            LicenseType = "LicenseIncluded",
            MaxSizeGb = 2,
            SkuName = "Basic",
        }
    );

    var databaseFirewallRule = new Azure.MSSql.FirewallRule(
        "database-firewall-rule",
        new()
        {
            Name = "azure",
            ServerId = server.Id,
            StartIpAddress = "0.0.0.0",
            EndIpAddress = "0.0.0.0",
        }
    );

    var providerMssql = new Mssql.Provider(
        "provider-mssql",
        new()
        {
            Hostname = server.FullyQualifiedDomainName,
            AzureAuth = new Mssql.Inputs.ProviderAzureAuthArgs(),
        },
        new CustomResourceOptions
        {
            DependsOn = new Resource[] { database, databaseFirewallRule, },
        }
    );

    var databaseId = Mssql.GetDatabase.Invoke(
        new() { Name = database.Name, },
        new InvokeOptions { Provider = providerMssql }
    );

    var databaseRoleOwner = Mssql.GetDatabaseRole.Invoke(
        new()
        {
            DatabaseId = databaseId.Apply(getDatabaseResult => getDatabaseResult.Id),
            Name = "db_owner",
        },
        new InvokeOptions { Provider = providerMssql }
    );

    var dbUser = new Mssql.SqlLogin(
        "db-user",
        new() { Name = "example", Password = "Str0ngPa$word12", },
        new CustomResourceOptions { Provider = providerMssql, }
    );

    var dbRoleMember = new Mssql.DatabaseRoleMember(
        "db-role-member",
        new()
        {
            RoleId = databaseRoleOwner.Apply(getDatabaseRoleResult => getDatabaseRoleResult.Id),
            MemberId = dbUser.Id,
        },
        new CustomResourceOptions { Provider = providerMssql, }
    );
});
```
{{% /choosable %}}

{{% choosable language yaml %}}
```yaml
name: mssql-test
runtime: yaml
description: A Pulumi YAML project to test the new MSSQL provider

variables:
  current:
    fn::invoke:
      function: azure:core:getClientConfig
      options:
        provider: ${provider}

  database-id:
    fn::invoke:
      function: mssql:getDatabase
      arguments:
        name: ${database.name}
      options:
        provider: ${provider-mssql}

  database-role-owner:
    fn::invoke:
      function: mssql:getDatabaseRole
      arguments:
        databaseId: ${database-id.id}
        name: "db_owner"
      options:
        provider: ${provider-mssql}

outputs: {}

resources:
  provider:
    type: pulumi:providers:azure
    properties:
      skipProviderRegistration: true
    options:
      version: 5.43.0

  resource-group:
    type: azure:core:ResourceGroup
    properties:
      name: "pulumi-mssql-test"
      location: "eastus"

  server:
    type: azure:mssql:Server
    properties:
      name: "pulumi-mssql-test"
      resourceGroupName: ${resource-group.name}
      location: ${resource-group.location}
      version: '12.0'
      minimumTlsVersion: '1.2'
      administratorLogin: sadmin
      administratorLoginPassword: "Str0ngPa$$word12"
      azureadAdministrator:
        loginUsername: "pulumi-app"
        objectId: ${current.objectId}
        tenantId: ${current.tenantId}
    options:
      provider: ${provider}

  database-firewall-rule:
    type: azure:mssql:FirewallRule
    properties:
      name: azure
      serverId: ${server.id}
      startIpAddress: "0.0.0.0"
      endIpAddress: "0.0.0.0"
    options:
      provider: ${provider}

  database:
    type: azure:mssql:Database
    properties:
      name: "db"
      serverId: ${server.id}
      licenseType: LicenseIncluded
      maxSizeGb: 2
      skuName: Basic
    options:
      provider: ${provider}

  provider-mssql:
    type: pulumi:providers:mssql
    properties:
      hostname: ${server.fullyQualifiedDomainName}
      azureAuth: {}
    options:
      dependsOn:
        - ${database}
        - ${database-firewall-rule}

  db-user:
    type: mssql:SqlLogin
    properties:
      name: example
      password: "Str0ngPa$$word12"
    options:
      provider: ${provider-mssql}

  db-role-member:
    type: mssql:DatabaseRoleMember
    properties:
      roleId: ${database-role-owner.id}
      memberId: ${db-user.id}
    options:
      provider: ${provider-mssql}
```
{{% /choosable %}}

{{< /chooser >}}
