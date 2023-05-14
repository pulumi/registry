---
title: ngrok
meta_desc: Provides an overview of the ngrok Provider for Pulumi.
layout: overview
---

The ngrok provider for Pulumi can be used to provision any of the cloud resources available in [ngrok](https://ngrok.com/).
The ngrok provider must be configured with credentials to deploy and update resources in ngrok.

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const ngrok = require("@pierskarsenbarg/ngrok")

const apikey = new ngrok.ApiKey("apikey");
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as ngrok from "@pierskarsenbarg/ngrok";

const apikey = new ngrok.ApiKey("apikey");
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pierskarsenbarg_pulumi_ngrok as ngrok

api_key = ngrok.ApiKey("apikey")
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pierskarsenbarg/pulumi-ngrok/ngrok/go/ngrok"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
        _, err := ngrok.NewApiKey(ctx, "apikey")
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
using Pulumi;
using PiersKarsenbarg.Ngrok;
using PiersKarsenbarg.Ngrok.Inputs;

return await Deployment.RunAsync(() =>
{
   var apiKey = new ApiKey("apikey");
});
```

{{% /choosable %}}

{{< /chooser >}}