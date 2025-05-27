---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/nats-io/jetstream/0.2.1/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Jetstream Provider
meta_desc: Provides an overview on how to configure the Pulumi Jetstream provider.
layout: package
---

## Generate Provider

The Jetstream provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider nats-io/jetstream
```
## Overview

The JetStream Provider manages [NATS](https://nats.io) JetStream that enables persistence and Stream
Processing for the NATS eco system.

For background information please review the [Technical Preview](https://github.com/nats-io/jetstream#readme) notes.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    jetstream:credentials:
        value: /home/you/ngs_stream_admin.creds
    jetstream:servers:
        value: connect.ngs.global:4222

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as jetstream from "@pulumi/jetstream";

const ORDERS = new jetstream.Stream("ORDERS", {
    name: "ORDERS",
    subjects: ["ORDERS.*"],
    storage: "file",
    maxAge: 60 * 60 * 24 * 365,
});
const ORDERS_NEW = new jetstream.Consumer("ORDERS_NEW", {
    streamId: ORDERS.streamId,
    durableName: "NEW",
    deliverAll: true,
    filterSubject: "ORDERS.received",
    sampleFreq: 100,
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    jetstream:credentials:
        value: /home/you/ngs_stream_admin.creds
    jetstream:servers:
        value: connect.ngs.global:4222

```
```python
import pulumi
import pulumi_jetstream as jetstream

orders = jetstream.Stream("ORDERS",
    name="ORDERS",
    subjects=["ORDERS.*"],
    storage="file",
    max_age=60 * 60 * 24 * 365)
order_s__new = jetstream.Consumer("ORDERS_NEW",
    stream_id=orders.stream_id,
    durable_name="NEW",
    deliver_all=True,
    filter_subject="ORDERS.received",
    sample_freq=100)
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    jetstream:credentials:
        value: /home/you/ngs_stream_admin.creds
    jetstream:servers:
        value: connect.ngs.global:4222

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Jetstream = Pulumi.Jetstream;

return await Deployment.RunAsync(() =>
{
    var ORDERS = new Jetstream.Stream("ORDERS", new()
    {
        Name = "ORDERS",
        Subjects = new[]
        {
            "ORDERS.*",
        },
        Storage = "file",
        MaxAge = 60 * 60 * 24 * 365,
    });

    var ORDERS_NEW = new Jetstream.Consumer("ORDERS_NEW", new()
    {
        StreamId = ORDERS.StreamId,
        DurableName = "NEW",
        DeliverAll = true,
        FilterSubject = "ORDERS.received",
        SampleFreq = 100,
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
    jetstream:credentials:
        value: /home/you/ngs_stream_admin.creds
    jetstream:servers:
        value: connect.ngs.global:4222

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/jetstream/jetstream"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		ORDERS, err := jetstream.NewStream(ctx, "ORDERS", &jetstream.StreamArgs{
			Name: pulumi.String("ORDERS"),
			Subjects: pulumi.StringArray{
				pulumi.String("ORDERS.*"),
			},
			Storage: pulumi.String("file"),
			MaxAge:  60 * 60 * 24 * 365,
		})
		if err != nil {
			return err
		}
		_, err = jetstream.NewConsumer(ctx, "ORDERS_NEW", &jetstream.ConsumerArgs{
			StreamId:      ORDERS.StreamId,
			DurableName:   pulumi.String("NEW"),
			DeliverAll:    pulumi.Bool(true),
			FilterSubject: pulumi.String("ORDERS.received"),
			SampleFreq:    pulumi.Float64(100),
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
    jetstream:credentials:
        value: /home/you/ngs_stream_admin.creds
    jetstream:servers:
        value: connect.ngs.global:4222

```
```yaml
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    jetstream:credentials:
        value: /home/you/ngs_stream_admin.creds
    jetstream:servers:
        value: connect.ngs.global:4222

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.jetstream.Stream;
import com.pulumi.jetstream.StreamArgs;
import com.pulumi.jetstream.Consumer;
import com.pulumi.jetstream.ConsumerArgs;
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
        var oRDERS = new Stream("ORDERS", StreamArgs.builder()
            .name("ORDERS")
            .subjects("ORDERS.*")
            .storage("file")
            .maxAge(60 * 60 * 24 * 365)
            .build());

        var oRDERS_NEW = new Consumer("ORDERS_NEW", ConsumerArgs.builder()
            .streamId(ORDERS.streamId())
            .durableName("NEW")
            .deliverAll(true)
            .filterSubject("ORDERS.received")
            .sampleFreq(100)
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

* `servers` - The list of servers to connect to in a comma seperated list.
* `credentials` - (optional) Fully Qualified Path to a file holding NATS credentials.
* `credentialData` - (optional) The NATS credentials as a string, intended to use with data providers.
* `user` - (optional) Connects using a username, when no password is set this is assumed to be a Token.
* `password` - (optional) Connects using a password.
* `nkey` - (optional) Connects using an nkey stored in a file.
* `tls.ca_file` - (optional) Fully Qualified Path to a file containing Root CA (PEM format). Use when the server has certs signed by an unknown authority.
* `tls.ca_file_data` - (optional) The Root CA PEM as a string, intended to use with data providers. Use when the server has certs signed by an unknown authority.
* `tls.cert_file` - (optional) The certificate to authenticate with.
* `tls.cert_file_data` - (optional) The certificate to authenticate with, intended to use with data providers.
* `tls.key_file` - (optional) The private key to authenticate with.
* `tls.key_file_data` - (optional) The private key to authenticate with, intended to use with data providers.
## Resources

* `jetstream.Stream` - Manage a Stream that persistently stores messages
* `jetstream.Consumer` - Creates a Consumer that defines how Stream messages can be consumed by clients
* `jetstream.KvBucket` - Creates a Key-Value store