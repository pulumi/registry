---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-postgresql/v3.16.2/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi/pulumi-postgresql/blob/v3.16.2/docs/_index.md
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Postgresql Provider
meta_desc: Provides an overview on how to configure the Pulumi Postgresql provider.
layout: package
---

## Installation

The Postgresql provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/postgresql`](https://www.npmjs.com/package/@pulumi/postgresql)
* Python: [`pulumi-postgresql`](https://pypi.org/project/pulumi-postgresql/)
* Go: [`github.com/pulumi/pulumi-postgresql/sdk/v3/go/postgresql`](https://github.com/pulumi/pulumi-postgresql)
* .NET: [`Pulumi.Postgresql`](https://www.nuget.org/packages/Pulumi.Postgresql)
* Java: [`com.pulumi/postgresql`](https://central.sonatype.com/artifact/com.pulumi/postgresql)

## Overview

The PostgreSQL provider gives the ability to deploy and configure resources in a PostgreSQL server.

Use the navigation to the left to read about the available resources.
## Usage

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    postgresql:connectTimeout:
        value: 15
    postgresql:database:
        value: postgres
    postgresql:host:
        value: postgres_server_ip
    postgresql:password:
        value: postgres_password
    postgresql:port:
        value: 5432
    postgresql:sslmode:
        value: require
    postgresql:username:
        value: postgres_user

```

An SSL client certificate can be configured using the `clientcert` sub-resource.

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    postgresql:database:
        value: postgres
    postgresql:host:
        value: postgres_server_ip
    postgresql:password:
        value: postgres_password
    postgresql:port:
        value: 5432
    postgresql:sslmode:
        value: require
    postgresql:username:
        value: postgres_user

```

Configuring multiple servers can be done by specifying the alias option.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as postgresql from "@pulumi/postgresql";

const myDb1 = new postgresql.Database("my_db1", {name: "my_db1"});
const myDb2 = new postgresql.Database("my_db2", {name: "my_db2"});
```

{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_postgresql as postgresql

my_db1 = postgresql.Database("my_db1", name="my_db1")
my_db2 = postgresql.Database("my_db2", name="my_db2")
```

{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using PostgreSql = Pulumi.PostgreSql;

return await Deployment.RunAsync(() =>
{
    var myDb1 = new PostgreSql.Database("my_db1", new()
    {
        Name = "my_db1",
    });

    var myDb2 = new PostgreSql.Database("my_db2", new()
    {
        Name = "my_db2",
    });

});

```

{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-postgresql/sdk/v3/go/postgresql"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := postgresql.NewDatabase(ctx, "my_db1", &postgresql.DatabaseArgs{
			Name: pulumi.String("my_db1"),
		})
		if err != nil {
			return err
		}
		_, err = postgresql.NewDatabase(ctx, "my_db2", &postgresql.DatabaseArgs{
			Name: pulumi.String("my_db2"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
resources:
  myDb1:
    type: postgresql:Database
    name: my_db1
    properties:
      name: my_db1
  myDb2:
    type: postgresql:Database
    name: my_db2
    properties:
      name: my_db2
```

{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.postgresql.Database;
import com.pulumi.postgresql.DatabaseArgs;
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
        var myDb1 = new Database("myDb1", DatabaseArgs.builder()
            .name("my_db1")
            .build());

        var myDb2 = new Database("myDb2", DatabaseArgs.builder()
            .name("my_db2")
            .build());

    }
}
```

{{% /choosable %}}
{{< /chooser >}}
## Injecting Credentials
There are several methods of providing credentials to the provider without hardcoding them.
### Environment Variables
Provider settings can be specified via environment variables as follows:

```shell
export PGHOST=localhost
export PGPORT=5432
export PGUSER=postgres
export PGPASSWORD=postgres
```
## Configuration Reference

The following configuration inputs are supported:

* `scheme` - (Optional) The driver to use. Valid values are:
  * `postgres`: Default value, use [`lib/pq`](https://pkg.go.dev/github.com/lib/pq)
  * `awspostgres`: Use GoCloud for AWS
  * `gcppostgres`: Use GoCloud for GCP
* `host` - (Required) The address for the postgresql server connection, see GoCloud for specific format.
* `port` - (Optional) The port for the postgresql server connection. The default is `5432`.
* `database` - (Optional) Database to connect to. The default is `postgres`.
* `username` - (Required) Username for the server connection.
* `password` - (Optional) Password for the server connection.
* `databaseUsername` - (Optional) Username of the user in the database if different than connection username (See [user name maps](https://www.postgresql.org/docs/current/auth-username-maps.html)).
* `superuser` - (Optional) Should be set to `false` if the user to connect is not a PostgreSQL superuser (as is the case in AWS RDS or GCP SQL).
  In this case, some features might be disabled (e.g.: Refreshing state password from database).
* `sslmode` - (Optional) Set the priority for an SSL connection to the server.
  Valid values for `sslmode` are (note: `prefer` is not supported by Go's
  [`lib/pq`](https://pkg.go.dev/github.com/lib/pq)):
  * disable - No SSL
  * require - Always SSL (the default, also skip verification)
  * verify-ca - Always SSL (verify that the certificate presented by the server was signed by a trusted CA)
  * verify-full - Always SSL (verify that the certification presented by the server was signed by a trusted CA and the server host name matches the one in the certificate)
    Additional information on the options and their implications can be seen
    [in the `libpq(3)` SSL guide](http://www.postgresql.org/docs/current/static/libpq-ssl.html#LIBPQ-SSL-PROTECTION).
* `clientcert` - (Optional) - Configure the SSL client certificate.
  * `cert` - (Required) - The SSL client certificate file path. The file must contain PEM encoded data.
  * `key` - (Required) - The SSL client certificate private key file path. The file must contain PEM encoded data.
  * `sslinline` - (Optional) - If set to `true`, arguments accept inline ssl cert and key rather than a filename. Defaults to `false`.
* `sslrootcert` - (Optional) - The SSL server root certificate file path. The file must contain PEM encoded data.
* `connectTimeout` - (Optional) Maximum wait for connection, in seconds. The
  default is `180s`.  Zero or not specified means wait indefinitely.
* `maxConnections` - (Optional) Set the maximum number of open connections to
  the database. The default is `20`.  Zero means unlimited open connections.
* `expectedVersion` - (Optional) Specify a hint to Pulumi regarding the
  expected version that the provider will be talking with.  This is a required
  hint in order for Pulumi to talk with an ancient version of PostgreSQL.
  This parameter is expected to be a [PostgreSQL
  Version](https://www.postgresql.org/support/versioning/) or `current`.  Once a
  connection has been established, Pulumi will fingerprint the actual
  version.  Default: `9.0.0`.
* `awsRdsIamAuth` - (Optional) If set to `true`, call the AWS RDS API to grab a temporary password, using AWS Credentials
  from the environment (or the given profile, see `awsRdsIamProfile`)
* `awsRdsIamProfile` - (Optional) The AWS IAM Profile to use while using AWS RDS IAM Auth.
* `awsRdsIamRegion` - (Optional) The AWS region to use while using AWS RDS IAM Auth.
* `awsRdsIamProviderRoleArn` - (Optional) AWS IAM role to assume while using AWS RDS IAM Auth.
* `azureIdentityAuth` - (Optional) If set to `true`, call the Azure OAuth token endpoint for temporary token
* `azureTenantId` - (Optional) (Required if `azureIdentityAuth` is `true`) Azure tenant ID read more
## GoCloud

By default, the provider uses the [lib/pq](https://pkg.go.dev/github.com/lib/pq) library to directly connect to PostgreSQL host instance. For connections to AWS/GCP hosted instances, the provider can connect through the [GoCloud](https://gocloud.dev/howto/sql/) library. GoCloud simplifies connecting to AWS/GCP hosted databases, managing any proxy or custom authentication details.
### AWS

To enable GoCloud based connections to AWS RDS instances, set `scheme` to `awspostgres` and `host` to the RDS database's endpoint value.
(e.g.: `instance.xxxxxx.region.rds.amazonaws.com`)

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    postgresql:host:
        value: test-instance.cvvrsv6scpgd.eu-central-1.rds.amazonaws.com
    postgresql:password:
        value: test1234
    postgresql:port:
        value: 5432
    postgresql:scheme:
        value: awspostgres
    postgresql:superuser:
        value: false
    postgresql:username:
        value: postgres

```
### GCP

To enable GoCloud for GCP SQL, set `scheme` to `gcppostgres` and `host` to the connection name of the instance in following format: `project/region/instance` (or `project:region:instance`).

For GCP, GoCloud also requires the `GOOGLE_APPLICATION_CREDENTIALS` environment variable to be set to the service account credentials file.
These credentials can be created here: <https://console.cloud.google.com/iam-admin/serviceaccounts>

In addition, the provider supports service account impersonation with the `gcpIamImpersonateServiceAccount` option. You must ensure:

- The IAM database user has sufficient permissions to connect to the database, e.g., `roles/cloudsql.instanceUser`
- The principal (IAM user or IAM service account) behind the `GOOGLE_APPLICATION_CREDENTIALS` has sufficient permissions to impersonate the provided service account. Learn more from [roles for service account authentication](https://cloud.google.com/iam/docs/service-account-permissions).

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    postgresql:gcpIamImpersonateServiceAccount:
        value: service_account_id@$project_id.iam.gserviceaccount.com
    postgresql:host:
        value: test-project/europe-west3/test-instance
    postgresql:port:
        value: 5432
    postgresql:scheme:
        value: gcppostgres
    postgresql:superuser:
        value: false
    postgresql:username:
        value: service_account_id@$project_id.iam

```

See also:

- <https://cloud.google.com/docs/authentication/production>
- <https://cloud.google.com/sql/docs/postgres/iam-logins>

---
**Note**

[Cloud SQL API](https://console.developers.google.com/apis/api/sqladmin.googleapis.com/overview) needs to be enabled for GoCloud to connect to your instance.

---

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    postgresql:host:
        value: test-project/europe-west3/test-instance
    postgresql:password:
        value: test1234
    postgresql:port:
        value: 5432
    postgresql:scheme:
        value: gcppostgres
    postgresql:superuser:
        value: false
    postgresql:username:
        value: postgres

```

Example with GCP resources:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    postgresql:host:
        value: 'TODO: google_sql_database_instance.test.connection_name'
    postgresql:password:
        value: 'TODO: google_sql_user.postgres.password'
    postgresql:scheme:
        value: gcppostgres
    postgresql:username:
        value: 'TODO: google_sql_user.postgres.name'

```

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as gcp from "@pulumi/gcp";
import * as postgresql from "@pulumi/postgresql";

const test = new gcp.sql.DatabaseInstance("test", {
    project: "test-project",
    name: "test-instance",
    databaseVersion: "POSTGRES_13",
    region: "europe-west3",
    settings: {
        tier: "db-f1-micro",
    },
});
const postgres = new gcp.sql.User("postgres", {
    project: "test-project",
    name: "postgres",
    instance: test.name,
    password: "xxxxxxxx",
});
const testDb = new postgresql.Database("test_db", {name: "test_db"});
```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    postgresql:host:
        value: 'TODO: google_sql_database_instance.test.connection_name'
    postgresql:password:
        value: 'TODO: google_sql_user.postgres.password'
    postgresql:scheme:
        value: gcppostgres
    postgresql:username:
        value: 'TODO: google_sql_user.postgres.name'

```

```python
import pulumi
import pulumi_gcp as gcp
import pulumi_postgresql as postgresql

test = gcp.sql.DatabaseInstance("test",
    project="test-project",
    name="test-instance",
    database_version="POSTGRES_13",
    region="europe-west3",
    settings={
        "tier": "db-f1-micro",
    })
postgres = gcp.sql.User("postgres",
    project="test-project",
    name="postgres",
    instance=test.name,
    password="xxxxxxxx")
test_db = postgresql.Database("test_db", name="test_db")
```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    postgresql:host:
        value: 'TODO: google_sql_database_instance.test.connection_name'
    postgresql:password:
        value: 'TODO: google_sql_user.postgres.password'
    postgresql:scheme:
        value: gcppostgres
    postgresql:username:
        value: 'TODO: google_sql_user.postgres.name'

```

```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Gcp = Pulumi.Gcp;
using PostgreSql = Pulumi.PostgreSql;

return await Deployment.RunAsync(() =>
{
    var test = new Gcp.Sql.DatabaseInstance("test", new()
    {
        Project = "test-project",
        Name = "test-instance",
        DatabaseVersion = "POSTGRES_13",
        Region = "europe-west3",
        Settings = new Gcp.Sql.Inputs.DatabaseInstanceSettingsArgs
        {
            Tier = "db-f1-micro",
        },
    });

    var postgres = new Gcp.Sql.User("postgres", new()
    {
        Project = "test-project",
        Name = "postgres",
        Instance = test.Name,
        Password = "xxxxxxxx",
    });

    var testDb = new PostgreSql.Database("test_db", new()
    {
        Name = "test_db",
    });

});

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    postgresql:host:
        value: 'TODO: google_sql_database_instance.test.connection_name'
    postgresql:password:
        value: 'TODO: google_sql_user.postgres.password'
    postgresql:scheme:
        value: gcppostgres
    postgresql:username:
        value: 'TODO: google_sql_user.postgres.name'

```

```go
package main

import (
	"github.com/pulumi/pulumi-gcp/sdk/v9/go/gcp/sql"
	"github.com/pulumi/pulumi-postgresql/sdk/v3/go/postgresql"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		test, err := sql.NewDatabaseInstance(ctx, "test", &sql.DatabaseInstanceArgs{
			Project:         pulumi.String("test-project"),
			Name:            pulumi.String("test-instance"),
			DatabaseVersion: pulumi.String("POSTGRES_13"),
			Region:          pulumi.String("europe-west3"),
			Settings: &sql.DatabaseInstanceSettingsArgs{
				Tier: pulumi.String("db-f1-micro"),
			},
		})
		if err != nil {
			return err
		}
		_, err = sql.NewUser(ctx, "postgres", &sql.UserArgs{
			Project:  pulumi.String("test-project"),
			Name:     pulumi.String("postgres"),
			Instance: test.Name,
			Password: pulumi.String("xxxxxxxx"),
		})
		if err != nil {
			return err
		}
		_, err = postgresql.NewDatabase(ctx, "test_db", &postgresql.DatabaseArgs{
			Name: pulumi.String("test_db"),
		})
		if err != nil {
			return err
		}
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
    postgresql:host:
        value: 'TODO: google_sql_database_instance.test.connection_name'
    postgresql:password:
        value: 'TODO: google_sql_user.postgres.password'
    postgresql:scheme:
        value: gcppostgres
    postgresql:username:
        value: 'TODO: google_sql_user.postgres.name'

```

```yaml
resources:
  test:
    type: gcp:sql:DatabaseInstance
    properties:
      project: test-project
      name: test-instance
      databaseVersion: POSTGRES_13
      region: europe-west3
      settings:
        tier: db-f1-micro
  postgres:
    type: gcp:sql:User
    properties:
      project: test-project
      name: postgres
      instance: ${test.name}
      password: xxxxxxxx
  testDb:
    type: postgresql:Database
    name: test_db
    properties:
      name: test_db
```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    postgresql:host:
        value: 'TODO: google_sql_database_instance.test.connection_name'
    postgresql:password:
        value: 'TODO: google_sql_user.postgres.password'
    postgresql:scheme:
        value: gcppostgres
    postgresql:username:
        value: 'TODO: google_sql_user.postgres.name'

```

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.gcp.sql.DatabaseInstance;
import com.pulumi.gcp.sql.DatabaseInstanceArgs;
import com.pulumi.gcp.sql.inputs.DatabaseInstanceSettingsArgs;
import com.pulumi.gcp.sql.User;
import com.pulumi.gcp.sql.UserArgs;
import com.pulumi.postgresql.Database;
import com.pulumi.postgresql.DatabaseArgs;
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
        var test = new DatabaseInstance("test", DatabaseInstanceArgs.builder()
            .project("test-project")
            .name("test-instance")
            .databaseVersion("POSTGRES_13")
            .region("europe-west3")
            .settings(DatabaseInstanceSettingsArgs.builder()
                .tier("db-f1-micro")
                .build())
            .build());

        var postgres = new User("postgres", UserArgs.builder()
            .project("test-project")
            .name("postgres")
            .instance(test.name())
            .password("xxxxxxxx")
            .build());

        var testDb = new Database("testDb", DatabaseArgs.builder()
            .name("test_db")
            .build());

    }
}
```

{{% /choosable %}}
{{< /chooser >}}
### Azure

To enable [passwordless authentication](https://learn.microsoft.com/en-us/azure/postgresql/flexible-server/how-to-configure-sign-in-azure-ad-authentication) with MS Azure set `azureIdentityAuth` to `true` and provide `azureTenantId`

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    postgresql:azureIdentityAuth:
        value: true
    postgresql:azureTenantId:
        value: 'TODO: data.azurerm_client_config.current.tenant_id'
    postgresql:database:
        value: postgres
    postgresql:host:
        value: 'TODO: azurerm_postgresql_flexible_server.pgsql.fqdn'
    postgresql:port:
        value: 5432
    postgresql:sslmode:
        value: require
    postgresql:username:
        value: 'TODO: azurerm_postgresql_flexible_server_active_directory_administrator.administrators.principal_name'

```

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as azure from "@pulumi/azure";

const current = azure.core.getClientConfig({});
// https://registry.pulumi.io/providers/pulumi/azurerm/latest/docs/resources/postgresql_flexible_server
const pgsql = new azure.postgresql.FlexibleServer("pgsql", {authentication: {
    activeDirectoryAuthEnabled: true,
    passwordAuthEnabled: false,
    tenantId: current.then(current => current.tenantId),
}});
// https://registry.pulumi.io/providers/pulumi/azurerm/latest/docs/resources/postgresql_flexible_server_active_directory_administrator
const administrators = new azure.postgresql.FlexibleServerActiveDirectoryAdministrator("administrators", {
    objectId: "00000000-0000-0000-0000-000000000000",
    principalName: "Azure AD Admin Group",
    principalType: "Group",
    resourceGroupName: rgName,
    serverName: pgsql.name,
    tenantId: current.then(current => current.tenantId),
});
```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    postgresql:azureIdentityAuth:
        value: true
    postgresql:azureTenantId:
        value: 'TODO: data.azurerm_client_config.current.tenant_id'
    postgresql:database:
        value: postgres
    postgresql:host:
        value: 'TODO: azurerm_postgresql_flexible_server.pgsql.fqdn'
    postgresql:port:
        value: 5432
    postgresql:sslmode:
        value: require
    postgresql:username:
        value: 'TODO: azurerm_postgresql_flexible_server_active_directory_administrator.administrators.principal_name'

```

```python
import pulumi
import pulumi_azure as azure

current = azure.core.get_client_config()
# https://registry.pulumi.io/providers/pulumi/azurerm/latest/docs/resources/postgresql_flexible_server
pgsql = azure.postgresql.FlexibleServer("pgsql", authentication={
    "active_directory_auth_enabled": True,
    "password_auth_enabled": False,
    "tenant_id": current.tenant_id,
})
# https://registry.pulumi.io/providers/pulumi/azurerm/latest/docs/resources/postgresql_flexible_server_active_directory_administrator
administrators = azure.postgresql.FlexibleServerActiveDirectoryAdministrator("administrators",
    object_id="00000000-0000-0000-0000-000000000000",
    principal_name="Azure AD Admin Group",
    principal_type="Group",
    resource_group_name=rg_name,
    server_name=pgsql.name,
    tenant_id=current.tenant_id)
```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    postgresql:azureIdentityAuth:
        value: true
    postgresql:azureTenantId:
        value: 'TODO: data.azurerm_client_config.current.tenant_id'
    postgresql:database:
        value: postgres
    postgresql:host:
        value: 'TODO: azurerm_postgresql_flexible_server.pgsql.fqdn'
    postgresql:port:
        value: 5432
    postgresql:sslmode:
        value: require
    postgresql:username:
        value: 'TODO: azurerm_postgresql_flexible_server_active_directory_administrator.administrators.principal_name'

```

```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Azure = Pulumi.Azure;

return await Deployment.RunAsync(() =>
{
    var current = Azure.Core.GetClientConfig.Invoke();

    // https://registry.pulumi.io/providers/pulumi/azurerm/latest/docs/resources/postgresql_flexible_server
    var pgsql = new Azure.PostgreSql.FlexibleServer("pgsql", new()
    {
        Authentication = new Azure.PostgreSql.Inputs.FlexibleServerAuthenticationArgs
        {
            ActiveDirectoryAuthEnabled = true,
            PasswordAuthEnabled = false,
            TenantId = current.Apply(getClientConfigResult => getClientConfigResult.TenantId),
        },
    });

    // https://registry.pulumi.io/providers/pulumi/azurerm/latest/docs/resources/postgresql_flexible_server_active_directory_administrator
    var administrators = new Azure.PostgreSql.FlexibleServerActiveDirectoryAdministrator("administrators", new()
    {
        ObjectId = "00000000-0000-0000-0000-000000000000",
        PrincipalName = "Azure AD Admin Group",
        PrincipalType = "Group",
        ResourceGroupName = rgName,
        ServerName = pgsql.Name,
        TenantId = current.Apply(getClientConfigResult => getClientConfigResult.TenantId),
    });

});

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    postgresql:azureIdentityAuth:
        value: true
    postgresql:azureTenantId:
        value: 'TODO: data.azurerm_client_config.current.tenant_id'
    postgresql:database:
        value: postgres
    postgresql:host:
        value: 'TODO: azurerm_postgresql_flexible_server.pgsql.fqdn'
    postgresql:port:
        value: 5432
    postgresql:sslmode:
        value: require
    postgresql:username:
        value: 'TODO: azurerm_postgresql_flexible_server_active_directory_administrator.administrators.principal_name'

```

```go
package main

import (
	"github.com/pulumi/pulumi-azure/sdk/v6/go/azure/core"
	"github.com/pulumi/pulumi-azure/sdk/v6/go/azure/postgresql"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		current, err := core.GetClientConfig(ctx, map[string]interface{}{}, nil)
		if err != nil {
			return err
		}
		// https://registry.pulumi.io/providers/pulumi/azurerm/latest/docs/resources/postgresql_flexible_server
		pgsql, err := postgresql.NewFlexibleServer(ctx, "pgsql", &postgresql.FlexibleServerArgs{
			Authentication: &postgresql.FlexibleServerAuthenticationArgs{
				ActiveDirectoryAuthEnabled: pulumi.Bool(true),
				PasswordAuthEnabled:        pulumi.Bool(false),
				TenantId:                   pulumi.String(current.TenantId),
			},
		})
		if err != nil {
			return err
		}
		// https://registry.pulumi.io/providers/pulumi/azurerm/latest/docs/resources/postgresql_flexible_server_active_directory_administrator
		_, err = postgresql.NewFlexibleServerActiveDirectoryAdministrator(ctx, "administrators", &postgresql.FlexibleServerActiveDirectoryAdministratorArgs{
			ObjectId:          pulumi.String("00000000-0000-0000-0000-000000000000"),
			PrincipalName:     pulumi.String("Azure AD Admin Group"),
			PrincipalType:     pulumi.String("Group"),
			ResourceGroupName: pulumi.Any(rgName),
			ServerName:        pgsql.Name,
			TenantId:          pulumi.String(current.TenantId),
		})
		if err != nil {
			return err
		}
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
    postgresql:azureIdentityAuth:
        value: true
    postgresql:azureTenantId:
        value: 'TODO: data.azurerm_client_config.current.tenant_id'
    postgresql:database:
        value: postgres
    postgresql:host:
        value: 'TODO: azurerm_postgresql_flexible_server.pgsql.fqdn'
    postgresql:port:
        value: 5432
    postgresql:sslmode:
        value: require
    postgresql:username:
        value: 'TODO: azurerm_postgresql_flexible_server_active_directory_administrator.administrators.principal_name'

```

```yaml
resources:
  # https://registry.pulumi.io/providers/pulumi/azurerm/latest/docs/resources/postgresql_flexible_server
  pgsql:
    type: azure:postgresql:FlexibleServer
    properties:
      authentication:
        activeDirectoryAuthEnabled: true
        passwordAuthEnabled: false
        tenantId: ${current.tenantId}
  # https://registry.pulumi.io/providers/pulumi/azurerm/latest/docs/resources/postgresql_flexible_server_active_directory_administrator
  administrators:
    type: azure:postgresql:FlexibleServerActiveDirectoryAdministrator
    properties:
      objectId: 00000000-0000-0000-0000-000000000000
      principalName: Azure AD Admin Group
      principalType: Group
      resourceGroupName: ${rgName}
      serverName: ${pgsql.name}
      tenantId: ${current.tenantId}
variables:
  current:
    fn::invoke:
      function: azure:core:getClientConfig
      arguments: {}
```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    postgresql:azureIdentityAuth:
        value: true
    postgresql:azureTenantId:
        value: 'TODO: data.azurerm_client_config.current.tenant_id'
    postgresql:database:
        value: postgres
    postgresql:host:
        value: 'TODO: azurerm_postgresql_flexible_server.pgsql.fqdn'
    postgresql:port:
        value: 5432
    postgresql:sslmode:
        value: require
    postgresql:username:
        value: 'TODO: azurerm_postgresql_flexible_server_active_directory_administrator.administrators.principal_name'

```

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.azure.core.CoreFunctions;
import com.pulumi.azure.postgresql.FlexibleServer;
import com.pulumi.azure.postgresql.FlexibleServerArgs;
import com.pulumi.azure.postgresql.inputs.FlexibleServerAuthenticationArgs;
import com.pulumi.azure.postgresql.FlexibleServerActiveDirectoryAdministrator;
import com.pulumi.azure.postgresql.FlexibleServerActiveDirectoryAdministratorArgs;
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
        final var current = CoreFunctions.getClientConfig(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference);

        // https://registry.pulumi.io/providers/pulumi/azurerm/latest/docs/resources/postgresql_flexible_server
        var pgsql = new FlexibleServer("pgsql", FlexibleServerArgs.builder()
            .authentication(FlexibleServerAuthenticationArgs.builder()
                .activeDirectoryAuthEnabled(true)
                .passwordAuthEnabled(false)
                .tenantId(current.tenantId())
                .build())
            .build());

        // https://registry.pulumi.io/providers/pulumi/azurerm/latest/docs/resources/postgresql_flexible_server_active_directory_administrator
        var administrators = new FlexibleServerActiveDirectoryAdministrator("administrators", FlexibleServerActiveDirectoryAdministratorArgs.builder()
            .objectId("00000000-0000-0000-0000-000000000000")
            .principalName("Azure AD Admin Group")
            .principalType("Group")
            .resourceGroupName(rgName)
            .serverName(pgsql.name())
            .tenantId(current.tenantId())
            .build());

    }
}
```

{{% /choosable %}}
{{< /chooser >}}
### SOCKS5 Proxy Support

The provider supports connecting via a SOCKS5 proxy, but when the `postgres` scheme is used. It can be configured by setting the `ALL_PROXY` or `allProxy` environment variable to a value like `socks5://127.0.0.1:1080`.

The `NO_PROXY` or `noProxy` environment can also be set to opt out of proxying for specific hostnames or ports.