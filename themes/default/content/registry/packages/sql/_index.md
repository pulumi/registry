---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/paultyng/sql/0.5.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Sql Provider
meta_desc: Provides an overview on how to configure the Pulumi Sql provider.
layout: package
---

## Generate Provider

The Sql provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider paultyng/sql
```
## Configuration Reference

- `maxIdleConns` (Number) Sets the maximum number of connections in the idle connection pool. Default is `2`. See Go's documentation on [DB.SetMaxIdleConns](https://golang.org/pkg/database/sql/#DB.SetMaxIdleConns).
- `maxOpenConns` (Number) Sets the maximum number of open connections to the database. Default is `0` (unlimited). See Go's documentation on [DB.SetMaxOpenConns](https://golang.org/pkg/database/sql/#DB.SetMaxOpenConns).
- `url` (String) Database connection strings are specified via URLs. The URL format is driver dependent but generally has the form: `dbdriver://username:password@host:port/dbname?param1=true&param2=false`. You can optionally set the `SQL_URL` environment variable instead.