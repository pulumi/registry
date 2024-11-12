---
# WARNING: this file was fetched from https://raw.githubusercontent.com/rootlyhq/pulumi-rootly/v1.5.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Rootly
meta_desc: Provides an overview of the Rootly Provider for Pulumi.
layout: package
---

The Rootly provider for Pulumi can be used to provision resources available in [Rootly](https://rootly.com/).
The Rootly provider must be configured with the an API token.

## Example

```javascript
const rootly = require("@rootly/pulumi")

new rootly.Severity("sev0", {
  name: "SEV0",
  color: "#FF0000"
})

new rootly.Service("elasticsearch_prod", {
  name: "elasticsearch-prod",
  color: "#800080"
})

new rootly.Functionality("add_items_to_card", {
  name: "Add items to card",
  color: "#FFFFFF"
})
```

> You could find more complete and detailed examples in the [pulumi-rootly repository](https://github.com/rootlyhq/pulumi-rootly/tree/main/examples)
