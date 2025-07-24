---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-mysql/v3.2.11/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Mysql Provider
meta_desc: Provides an overview on how to configure the Pulumi Mysql provider.
layout: package
---
## Installation

The Mysql provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/mysql`](https://www.npmjs.com/package/@pulumi/mysql)
* Python: [`pulumi-mysql`](https://pypi.org/project/pulumi-mysql/)
* Go: [`github.com/pulumi/pulumi-mysql/sdk/v3/go/mysql`](https://github.com/pulumi/pulumi-mysql)
* .NET: [`Pulumi.Mysql`](https://www.nuget.org/packages/Pulumi.Mysql)
* Java: [`com.pulumi/mysql`](https://central.sonatype.com/artifact/com.pulumi/mysql)
## Overview

[MySQL](http://www.mysql.com) is a relational database server. The MySQL
provider exposes resources used to manage the configuration of resources
in a MySQL server.

Use the navigation to the left to read about the available resources.
## Example Usage

The following is a minimal example:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    mysql:endpoint:
        value: my-database.example.com:3306
    mysql:password:
        value: app-password
    mysql:username:
        value: app-user

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as mysql from "@pulumi/mysql";

// Create a Database
const app = new mysql.Database("app", {name: "my_awesome_app"});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    mysql:endpoint:
        value: my-database.example.com:3306
    mysql:password:
        value: app-password
    mysql:username:
        value: app-user

```
```python
import pulumi
import pulumi_mysql as mysql

# Create a Database
app = mysql.Database("app", name="my_awesome_app")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    mysql:endpoint:
        value: my-database.example.com:3306
    mysql:password:
        value: app-password
    mysql:username:
        value: app-user

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using MySql = Pulumi.MySql;

return await Deployment.RunAsync(() =>
{
    // Create a Database
    var app = new MySql.Database("app", new()
    {
        Name = "my_awesome_app",
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
    mysql:endpoint:
        value: my-database.example.com:3306
    mysql:password:
        value: app-password
    mysql:username:
        value: app-user

```
```go
package main

import (
	"github.com/pulumi/pulumi-mysql/sdk/v3/go/mysql"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a Database
		_, err := mysql.NewDatabase(ctx, "app", &mysql.DatabaseArgs{
			Name: pulumi.String("my_awesome_app"),
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
    mysql:endpoint:
        value: my-database.example.com:3306
    mysql:password:
        value: app-password
    mysql:username:
        value: app-user

```
```yaml
resources:
  # Create a Database
  app:
    type: mysql:Database
    properties:
      name: my_awesome_app
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    mysql:endpoint:
        value: my-database.example.com:3306
    mysql:password:
        value: app-password
    mysql:username:
        value: app-user

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.mysql.Database;
import com.pulumi.mysql.DatabaseArgs;
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
        // Create a Database
        var app = new Database("app", DatabaseArgs.builder()
            .name("my_awesome_app")
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

This provider can be used in conjunction with other resources that create
MySQL servers. For example, `awsDbInstance` is able to create MySQL
servers in Amazon's RDS service.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    mysql:endpoint:
        value: 'TODO: "${aws_db_instance.default.endpoint}"'
    mysql:password:
        value: 'TODO: "${aws_db_instance.default.password}"'
    mysql:username:
        value: 'TODO: "${aws_db_instance.default.username}"'

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as mysql from "@pulumi/mysql";

// Create a database server
const _default = new aws.rds.Instance("default", {
    engine: "mysql",
    engineVersion: "5.6.17",
    instanceClass: "db.t1.micro",
    name: "initial_db",
    username: "rootuser",
    password: "rootpasswd",
});
// Create a second database, in addition to the "initial_db" created
// by the aws_db_instance resource above.
const app = new mysql.Database("app", {name: "another_db"});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    mysql:endpoint:
        value: 'TODO: "${aws_db_instance.default.endpoint}"'
    mysql:password:
        value: 'TODO: "${aws_db_instance.default.password}"'
    mysql:username:
        value: 'TODO: "${aws_db_instance.default.username}"'

```
```python
import pulumi
import pulumi_aws as aws
import pulumi_mysql as mysql

# Create a database server
default = aws.rds.Instance("default",
    engine="mysql",
    engine_version="5.6.17",
    instance_class="db.t1.micro",
    name="initial_db",
    username="rootuser",
    password="rootpasswd")
# Create a second database, in addition to the "initial_db" created
# by the aws_db_instance resource above.
app = mysql.Database("app", name="another_db")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    mysql:endpoint:
        value: 'TODO: "${aws_db_instance.default.endpoint}"'
    mysql:password:
        value: 'TODO: "${aws_db_instance.default.password}"'
    mysql:username:
        value: 'TODO: "${aws_db_instance.default.username}"'

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;
using MySql = Pulumi.MySql;

return await Deployment.RunAsync(() =>
{
    // Create a database server
    var @default = new Aws.Rds.Instance("default", new()
    {
        Engine = "mysql",
        EngineVersion = "5.6.17",
        InstanceClass = "db.t1.micro",
        Name = "initial_db",
        Username = "rootuser",
        Password = "rootpasswd",
    });

    // Create a second database, in addition to the "initial_db" created
    // by the aws_db_instance resource above.
    var app = new MySql.Database("app", new()
    {
        Name = "another_db",
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
    mysql:endpoint:
        value: 'TODO: "${aws_db_instance.default.endpoint}"'
    mysql:password:
        value: 'TODO: "${aws_db_instance.default.password}"'
    mysql:username:
        value: 'TODO: "${aws_db_instance.default.username}"'

```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
	"github.com/pulumi/pulumi-mysql/sdk/v3/go/mysql"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a database server
		_, err := rds.NewInstance(ctx, "default", &rds.InstanceArgs{
			Engine:        pulumi.String("mysql"),
			EngineVersion: pulumi.String("5.6.17"),
			InstanceClass: pulumi.String("db.t1.micro"),
			Name:          pulumi.String("initial_db"),
			Username:      pulumi.String("rootuser"),
			Password:      pulumi.String("rootpasswd"),
		})
		if err != nil {
			return err
		}
		// Create a second database, in addition to the "initial_db" created
		// by the aws_db_instance resource above.
		_, err = mysql.NewDatabase(ctx, "app", &mysql.DatabaseArgs{
			Name: pulumi.String("another_db"),
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
    mysql:endpoint:
        value: 'TODO: "${aws_db_instance.default.endpoint}"'
    mysql:password:
        value: 'TODO: "${aws_db_instance.default.password}"'
    mysql:username:
        value: 'TODO: "${aws_db_instance.default.username}"'

```
```yaml
resources:
  # Create a database server
  default:
    type: aws:rds:Instance
    properties:
      engine: mysql
      engineVersion: 5.6.17
      instanceClass: db.t1.micro
      name: initial_db
      username: rootuser
      password: rootpasswd
  # Create a second database, in addition to the "initial_db" created
  # by the aws_db_instance resource above.
  app:
    type: mysql:Database
    properties:
      name: another_db
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    mysql:endpoint:
        value: 'TODO: "${aws_db_instance.default.endpoint}"'
    mysql:password:
        value: 'TODO: "${aws_db_instance.default.password}"'
    mysql:username:
        value: 'TODO: "${aws_db_instance.default.username}"'

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.rds.Instance;
import com.pulumi.aws.rds.InstanceArgs;
import com.pulumi.mysql.Database;
import com.pulumi.mysql.DatabaseArgs;
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
        // Create a database server
        var default_ = new Instance("default", InstanceArgs.builder()
            .engine("mysql")
            .engineVersion("5.6.17")
            .instanceClass("db.t1.micro")
            .name("initial_db")
            .username("rootuser")
            .password("rootpasswd")
            .build());

        // Create a second database, in addition to the "initial_db" created
        // by the aws_db_instance resource above.
        var app = new Database("app", DatabaseArgs.builder()
            .name("another_db")
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## SOCKS5 Proxy Support

The MySQL provider respects the `ALL_PROXY` and/or `allProxy` environment variables.

```
$ export all_proxy="socks5://your.proxy:3306"
```
## Configuration Reference

The following configuration inputs are supported:

* `endpoint` - (Required) The address of the MySQL server to use. Most often a "hostname:port" pair, but may also be an absolute path to a Unix socket when the host OS is Unix-compatible. Can also be sourced from the `MYSQL_ENDPOINT` environment variable.
* `username` - (Required) Username to use to authenticate with the server, can also be sourced from the `MYSQL_USERNAME` environment variable.
* `password` - (Optional) Password for the given user, if that user has a password, can also be sourced from the `MYSQL_PASSWORD` environment variable.
* `proxy` - (Optional) Proxy socks url, can also be sourced from `ALL_PROXY` or `allProxy` environment variables.
* `tls` - (Optional) The TLS configuration. One of `false`, `true`, or `skip-verify`. Defaults to `false`. Can also be sourced from the `MYSQL_TLS_CONFIG` environment variable.
* `maxConnLifetimeSec` - (Optional) Sets the maximum amount of time a connection may be reused. If d <= 0, connections are reused forever.
* `maxOpenConns` - (Optional) Sets the maximum number of open connections to the database. If n <= 0, then there is no limit on the number of open connections.
* `authenticationPlugin` - (Optional) Sets the authentication plugin, it can be one of the following: `native` or `cleartext`. Defaults to `native`.