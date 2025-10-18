# Resource Example Shortcode Demo

This document demonstrates the dummy `resource-example` shortcode implementation.

## Usage

In the EC2 Auto-Scaling guide (`content/registry/packages/aws/guides/ec2-auto-scaling.md`), we replaced ~160 lines of hardcoded multi-language examples with a single line:

```markdown
{{< resource-example provider="aws" resource="aws.autoscaling.Group" >}}
```

**Important**: The `resource` parameter uses the Pulumi SDK format (e.g., `aws.autoscaling.Group`), not the URL path format (`autoscaling/group`).

## What It Renders

The shortcode currently renders a **purple placeholder box** that clearly shows:

### Visual Elements
1. **Info icon** - SVG icon indicating this is a placeholder
2. **Title**: "🚧 Placeholder for resource-example shortcode"
3. **Parameters shown**:
   - Provider: `aws`
   - Resource: `autoscaling/group`
4. **Expected JSON path**: `/static/registry/packages/examples/aws/autoscaling-group.json`
5. **What will be rendered**: Description of language chooser and multi-language examples
6. **Implementation notes**: Technical details about how it will work

### Styling
- Border: 4px purple dashed border
- Background: Light purple (purple-50)
- Padding: 1.5rem
- Rounded corners
- Clear visual hierarchy with nested sections

## Benefits of This Approach

### For Research/Prototyping
- **Clearly visible**: The purple placeholder is impossible to miss
- **Informative**: Shows exactly what parameters are being used
- **Educational**: Explains how the final version will work
- **No broken pages**: Renders gracefully without the actual data

### For Final Implementation
The shortcode signature is already defined:
```go
{{ $provider := .Get "provider" }}
{{ $resource := .Get "resource" }}
```

To implement the real version, we just need to:
1. Read JSON from: `static/registry/packages/examples/{{ $provider }}/{{ replace $resource "/" "-" }}.json`
2. Parse the JSON structure
3. Render with `pulumi-chooser`/`pulumi-choosable` components

## Example JSON Structure (Future)

When implemented, the JSON files will look like:

```json
{
  "resource": "aws:autoscaling/group:Group",
  "title": "Basic Auto Scaling Group",
  "examples": [
    {
      "title": "Basic Usage",
      "typescript": "import * as aws from \"@pulumi/aws\";\n...",
      "python": "import pulumi_aws as aws\n...",
      "go": "package main\n...",
      "csharp": "using Pulumi.Aws;\n...",
      "java": "import com.pulumi.aws;\n...",
      "yaml": "resources:\n..."
    }
  ]
}
```

## Comparison: Before vs After

### Before (Manual Multi-Language Examples)
```markdown
{{< chooser language "typescript,python,go,csharp,java" >}}
{{% choosable language typescript %}}
```typescript
[~30 lines of TypeScript code]
```
{{% /choosable %}}

{{% choosable language python %}}
```python
[~20 lines of Python code]
```
{{% /choosable %}}

{{% choosable language go %}}
```go
[~30 lines of Go code]
```
{{% /choosable %}}

{{% choosable language csharp %}}
```csharp
[~25 lines of C# code]
```
{{% /choosable %}}

{{% choosable language java %}}
```java
[~30 lines of Java code]
```
{{% /choosable %}}
{{< /chooser >}}
```

**Total: ~160 lines**

### After (Shortcode)
```markdown
{{< resource-example provider="aws" resource="autoscaling/group" >}}
```

**Total: 1 line**

## Next Steps

1. **Generate the JSON files**: Modify `resourcedocsgen` to export examples
2. **Implement real shortcode**: Replace placeholder with actual rendering logic
3. **Add more guides**: Each guide can now reference resource examples easily
4. **Automatic sync**: Examples update whenever API docs regenerate

## Testing URLs

- **Guides list**: http://localhost:1313/registry/packages/aws/guides/
- **EC2 Guide with placeholder**: http://localhost:1313/registry/packages/aws/guides/ec2-auto-scaling/

Look for the purple box under "Example: Basic Auto Scaling Group" section!
