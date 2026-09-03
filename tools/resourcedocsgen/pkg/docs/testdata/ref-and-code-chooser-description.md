Provides a demonstration of refs alongside code chooser blocks.

See also {{% ref #/resources/prov:module%2Fresource:Resource %}} for the primary resource.

<!--Start PulumiCodeChooser -->
```typescript
import * as prov from "@pulumi/prov";
const example = new prov.module.Resource("example", {});
```
```python
import pulumi_prov as prov
example = prov.module.Resource("example")
```
```go
example, err := module.NewResource(ctx, "example", nil)
```
```csharp
var example = new Prov.Module.Resource("example");
```
<!--End PulumiCodeChooser -->

The {{% ref #/resources/prov:module%2Fresource:Resource/inputProperties/integerProp %}} property
controls the integer value.
