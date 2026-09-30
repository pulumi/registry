---
# WARNING: this file was fetched from https://raw.githubusercontent.com/jschady/pulumi-filescom/v0.2.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/jschady/pulumi-filescom/blob/v0.2.0/docs/_index.md
title: Files.com
meta_desc: Provides an overview of the Files.com provider for Pulumi.
layout: package
---

The Files.com provider creates and manages the objects in a Files.com account:

- folders
- folder settings that Files.com calls behaviors
- remote servers
- groups and the users
- API keys

This provider wraps the
[Files.com provider for Terraform](https://github.com/Files-com/terraform-provider-files). Every
resource and every data source that provider carries is available here, under a Pulumi name.

## Example

The program below creates one group and exports the id that Files.com assigns to it.

{{< chooser language "typescript,python,csharp,go" >}}

{{% choosable language typescript %}}

```typescript
import * as filescom from "pulumi-filescom";

const engineers = new filescom.Group("engineers", {
    name: "engineers",
    notes: "The engineering team.",
});

export const groupId = engineers.id;
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_filescom as filescom

engineers = filescom.Group("engineers",
    name="engineers",
    notes="The engineering team.")

pulumi.export("group_id", engineers.id)
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using Pulumi;
using Filescom = Jschady.Filescom;

return await Deployment.RunAsync(() =>
{
    var engineers = new Filescom.Group("engineers", new()
    {
        Name = "engineers",
        Notes = "The engineering team.",
    });

    return new Dictionary<string, object?>
    {
        ["groupId"] = engineers.Id,
    };
});
```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"github.com/jschady/pulumi-filescom/sdk/go/filescom"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		engineers, err := filescom.NewGroup(ctx, "engineers", &filescom.GroupArgs{
			Name:  pulumi.String("engineers"),
			Notes: pulumi.String("The engineering team."),
		})
		if err != nil {
			return err
		}
		ctx.Export("groupId", engineers.ID())
		return nil
	})
}
```

{{% /choosable %}}

{{< /chooser >}}

## API key

The provider reads the API key from the `FILES_API_KEY` environment variable. You can set the key in
the stack configuration instead. The
[installation and configuration page](/registry/packages/filescom/installation-configuration/)
carries both commands and the rest of the provider configuration.

## Resource id

Files.com assigns a decimal id to each object it stores. The provider reports that id as the Pulumi
`id` output, and the value matches the id the account holds.

## Import of an existing object

You can adopt a Files.com object that Pulumi did not create, except some
[behaviors](#import-of-a-behavior). The provider reads the object, and
`pulumi import` writes the properties of the answer into a generated declaration.

**Warning:** `pulumi import` protects the resource, and `pulumi destroy` refuses a protected
resource. Set `--protect=false` when you plan to destroy the stack later.

Set the API key in the environment, then import the object by its decimal id:

```bash
export FILES_API_KEY=your-api-key
pulumi import filescom:index/group:Group adopted 12345 --protect=false --out adopted.yaml
```

Copy the declaration from `adopted.yaml` into your program, then run `pulumi preview`. The plan
reports no change.

The generated declaration carries each property that Files.com reports with a value. It leaves out
a property whose value is one of these, and the plan after the import stays empty:

- empty
- zero
- null

## Upgrade from 0.1.x to 0.2.0

Files.com retired projects and messages, so 0.2.0 drops these resources and their `get` functions:

- `Message`
- `MessageComment`
- `MessageCommentReaction`
- `MessageReaction`
- `Project`

Pulumi cannot delete these resources: 0.1.x gets a 404 from Files.com, and 0.2.0 no longer knows
the types. Nothing is left on Files.com, so remove each one from the state instead:

1. Remove the resource from your program.
2. Find its URN:

   ```bash
   pulumi stack --show-urns
   ```

3. Remove it from the state, dependents first (a message before its project), or add
   `--target-dependents`:

   ```bash
   pulumi state delete '<urn>'
   ```

Version 0.2.0 also warns about every unwrapped behavior `value`.
A `value` that 0.1.x wrote as a JSON string fails every preview until you wrap it. Either way,
wrap the value as [Value of a behavior](#value-of-a-behavior) shows.
Pulumi updates the behavior in place.

## Limitations

### Value of a behavior

Wrap `value` under the behavior name and keep the API's snake_case keys:

```typescript
new filescom.Behavior("retention", {
    path: "reports",
    behavior: "file_expiration",
    value: { file_expiration: { days_to_retain: 30, delete_empty_folders: false } },
});
```

An unwrapped value or a JSON string still works, with a warning. Files.com plans to drop both on
March 1, 2027.

If 0.2.0 wrote `value` as a JSON string, changing it to an object fails every preview with
`can't unmarshal ... into *string`
([pulumi/pulumi-terraform-bridge#3122](https://github.com/pulumi/pulumi-terraform-bridge/issues/3122)).
To make the change:

1. Change `value` to the wrapped object.
2. Add the `deleteBeforeReplace` resource option (`delete_before_replace` in Python). Without it,
   Files.com can refuse the replacement with `Behavior is already set for this folder`.
3. Replace the behavior. It gets a new id.

   ```bash
   pulumi up --replace '<urn>'
   ```

4. Optionally, remove the option. The next preview shows no change.

### Plan of a behavior

Every update plan and every replace plan for a behavior drops 3 computed outputs: `inherited`,
`managed`, and `rootBehaviorSiteAdminOnly`. The apply returns all 3. The preview right after an
apply is clean, so this is not drift.

### Import of a behavior

Importing a behavior crashes the provider if its `value` holds a nested object, such as webhook
`headers`. Other behaviors import, but the generated declaration leaves `value` unwrapped. Wrap it,
and Pulumi updates the behavior in place.

### Path of an API key

Pulumi sends the `path` property to Files.com and never stores it. Pulumi holds no stored value to
compare, so a change to `path` plans no update at all.

**Warning:** If you delete an API key, every program that holds it stops working.

To restrict a different path, follow these steps.

1. Delete the API key.
2. Create a new key with the path you want.

A replace that another property forces does send the current `path` again.

### Order of the group members

If you change the order of `userIds` on a group, Pulumi plans an update. The set of members is the
same, and the plan still shows the change. The plan path in the upstream framework never calls the
semantic-equality check that treats the two orders as equal. Write `userIds` in a stable order to
keep the plan empty.

### Properties without a description

The registry shows 312 properties with an empty description. Files.com declares 310 of them with no
text. The code generator adds the other 2. This provider invents no replacement.

| Where the description is empty | Count |
| --- | --- |
| A resource output | 4 |
| A data source output | 6 |
| A provider configuration property | 2 |
| A property inside a nested object | 298 |
| A property the code generator adds | 2 |

The 298 nested properties sit under two trees. The definition of an automation holds 280, and the
definition of a holiday calendar holds 18.
