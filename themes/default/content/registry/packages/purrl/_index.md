---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumiverse/pulumi-purrl/v0.6.2/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Purrl
meta_desc: Provides an overview of the Purrl Provider for Pulumi.
layout: overview
---

This provider is designed to be a flexible extension of your Pulumi code to make API calls to your target endpoint. `Purrl` is useful when a provider does not have a resource or data source that you require, so `Purrl` can be used to make substitute API calls.

## Example

{{< chooser language "typescript,python,go,csharp" >}}

{{% choosable language typescript %}}

```typescript
import * as purrl from "@pulumiverse/purrl";

let purrlCommand = new purrl.Purrl("purrl", {
    name: "httpbin",
    url: "https://httpbin.org/get",
    method: "GET",
    headers: {
        "test": "test",
    },
    responseCodes: [
        "200"
    ],
    deleteMethod: "DELETE",
    deleteUrl: "https://httpbin.org/delete",
    deleteResponseCodes: [
        "200"
    ],
});

export const url = purrlCommand.response
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumiverse_purrl as purrl
import pulumi

purrl_command = purrl.Purrl("purrl-python", name="purrl-python",
                            method="GET",
                            headers={
                                "test": "test"
                            },
                            url="https://httpbin.org/get",
                            response_codes=[
                                "200"],
                            delete_method="DELETE",
                            delete_url="https://httpbin.org/delete",
                            delete_response_codes=["200"]
                            )

pulumi.export("response", purrl_command.response)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-purrl/sdk/go/purrl"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		purrl, err := purrl.NewPurrl(ctx, "purrl", &purrl.PurrlArgs{
			Url:  pulumi.String("https://httpbin.org/get"),
			Name: pulumi.String("httpbin"),
			ResponseCodes: pulumi.StringArray{
				pulumi.String("200"),
			},
			Method: pulumi.String("GET"),
			Headers: pulumi.StringMap{
				"test": pulumi.String("test"),
			},
			DeleteMethod: pulumi.String("DELETE"),
			DeleteUrl:    pulumi.String("https://httpbin.org/delete"),
			DeleteResponseCodes: pulumi.StringArray{
				pulumi.String("200"),
			},
		})
		if err != nil {
			return err
		}
		ctx.Export("response", purrl.Response)
		return nil
	})
}
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using Pulumi;
using Pulumiverse.Purrl;

return await Deployment.RunAsync(() =>
{
   var purrl =new Purrl("purrl", new PurrlArgs
   {
      Name = "httpbin",
      Url = "https://httpbin.org/get",
      ResponseCodes = new List<string> { "200" },
      Method = "GET",
      Headers = new Dictionary<string, string> { { "test", "test" } },
      DeleteMethod = "DELETE",
      DeleteUrl = "https://httpbin.org/delete",
      DeleteResponseCodes = new List<string> { "200" },
   });

   // Export outputs here
   return new Dictionary<string, object?>
   {
      ["response"] =purrl.Response
   };
});
```

{{% /choosable %}}

{{< /chooser >}}
