---
# WARNING: this file was fetched from https://raw.githubusercontent.com/stackitcloud/pulumi-stackit/v0.0.4/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/stackitcloud/pulumi-stackit/blob/v0.0.4/docs/_index.md
title: stackit
meta_desc: Learn how to use Pulumi's STACKIT provider to manage STACKIT resources.
layout: package
---

<div align="center">
<br>
<img src="https://raw.githubusercontent.com/stackitcloud/pulumi-stackit/refs/heads/main/.github/images/stackit-logo.svg" alt="STACKIT logo" width="50%"/>
<br>
<br>
</div>

# STACKIT Pulumi Provider (ALPHA)

[![GitHub License](https://img.shields.io/github/license/stackitcloud/pulumi-stackit)](https://www.apache.org/licenses/LICENSE-2.0)

> [!WARNING]
> This STACKIT Pulumi Provider is in an ALPHA state. Currently, only IaaS and Resource Manager services in TypeScript are supported. More functionality will be supported soon. Your feedback is appreciated! Feel free to open GitHub issues to provide feature requests and bug reports.

The STACKIT Pulumi Provider lets you manage [STACKIT](https://www.stackit.de/en/) resources.


## Reporting issues

If you encounter any issues or have suggestions for improvements, please open an issue in the repository or create a ticket in the [STACKIT Help Center](https://support.stackit.cloud/).

## License

Apache 2.0

## Installing

> [!NOTE]
> Currently only the npm package is available via the mentioned packages managers but the other ones can be installed locally by using the generated sdks directly.
> Keep in mind that this is in ALPHA state and only the Resource Manager and IaaS service for Node.js is supported.
> E.g. in order to create and install the TypeScript SDK you can use the following [guide](https://github.com/stackitcloud/pulumi-stackit/blob/main/examples/ts/README.md).

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @stackitcloud/pulumi-stackit
```

or `yarn`:

```bash
yarn add @stackitcloud/pulumi-stackit
```

## Example

{{< chooser language "typescript" >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as stackit from "@stackitcloud/pulumi-stackit";

const exampleWithName = new stackit.Network("example_with_name", {
    projectId: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",                          // Replace with your actual project ID
    name: "example-with-name",
});
```

{{% /choosable %}}
{{< /chooser >}}

For more examples, please have a look at the [Examples](https://github.com/stackitcloud/pulumi-stackit/blob/main/examples).

A detailed authentication guide, information on how to use beta and experimental resources, and available configuration options can be found under [Installation Configuration](./installation-configuration).

## Reference

For detailed reference documentation, please visit [the Pulumi registry](./api-docs/).
