# Checkly

The Checkly provider for Pulumi can be used to provision any of the monitoring resources available in [Checkly](https://checklhyhq.com/).
The Checkly provider must be configured with the an `API Key` and also set the target `Account ID` in order to deploy Checkly resources.

## Example

{{< chooser language "javascript,typescript" >}}

{{% choosable language javascript %}}

```javascript
const checkly = require("@checkly/pulumi")

new checkly.Check("api-check", {
  activated: true,
  frequency: 10,
  type: "API",
  request: {
    method: "GET",
    url: "https://checklyhq.com",
  }
})

new checkly.Check("brwoser-check", {
  activated: true,
  frequency: 10,
  type: "BROWSER",
  script: 'console.log("Hello World!")'
})
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as checkly from "@pulumi/checkly";

new checkly.Check( "api-check", {
  activated: true,
  frequency: 10,
  type: "API",
  request: {
    method: "GET",
    url: "https://checklyhq.com",
  }
})

new checkly.Check( "brwoser-check", {
  activated: true,
  frequency: 10,
  type: "BROWSER",
  script: 'console.log("Hello World!")'
})
```
{{% /choosable %}}
{{< /chooser >}}
