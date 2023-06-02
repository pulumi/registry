
---
title: impart
meta_desc: Provides an overview of the Impart Provider for Pulumi.
layout: overview
---

The Impart Resource Provider lets you manage Impart resources.

## Example

{{< chooser language "javascript,typescript,go" >}}

{{% choosable language javascript %}}

```js
const pulumi = require("@pulumi/pulumi");
const impart = require("@impart-security/pulumi-impart");
const crypto = require("crypto");
const fs = require("fs");

const hashSum = crypto.createHash("sha256");
const specSource = fs.readFileSync(`./spec.yaml`).toString();

const spec = new impart.Spec("spec", {
  name: "example_spec",
  sourceFile: "spec.yaml",
  //optional to track source files changes
  sourceHash: hashSum.update(specSource).digest("hex"),
});

const binding = new impart.Binding("binding", {
  name: "example_binding",
  port: 8080,
  hostname: "example.com",
  basePath: "/",
  specId: spec.id,
});
```

{{% /choosable %}}

{{% choosable language typescript %}}

```ts
import * as pulumi from "@pulumi/pulumi";
import * as impart from "@impart-security/pulumi-impart";
import * as crypto from "crypto";
import * as fs from "fs";

const hashSum = crypto.createHash("sha256");
const specSource = fs.readFileSync(`./spec.yaml`).toString();

const spec = new impart.Spec("spec", {
  name: "example_spec",
  sourceFile: "spec.yaml",
  //optional to track source files changes
  sourceHash: hashSum.update(specSource).digest("hex"),
});

const binding = new impart.Binding("binding", {
  name: "example_binding",
  port: 8080,
  hostname: "example.com",
  basePath: "/",
  specId: spec.id,
});
```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"

	"github.com/impart-security/pulumi-impart/sdk/go/impart"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		specFile := "./spec.yaml"
		specHash, err := getFileHash(specFile)
		if err != nil {
			return err
		}
		spec, err := impart.NewSpec(ctx, "example", &impart.SpecArgs{
			Name:       pulumi.String("spec_example"),
			SourceFile: pulumi.String(specFile),
			SourceHash: pulumi.String(specHash),
		})
		if err != nil {
			return err
		}

		_, err = impart.NewBinding(ctx, "example", &impart.BindingArgs{
			Name:     pulumi.String("binding_example"),
			Port:     pulumi.Int(443),
			SpecId:   spec.ID().ToStringOutput(),
			Hostname: pulumi.String("example.com"),
			BasePath: pulumi.String("/"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

func getFileHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	sum := fmt.Sprintf("%x", hash.Sum(nil))
	return sum, nil
}
```

{{% /choosable %}}

{{< /chooser >}}

