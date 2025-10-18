# LLM-Generated Content Shortcode Demo

This document explains the `llm-explain` shortcode for generating dynamic content using Claude AI.

## Overview

The `llm-explain` shortcode allows guide authors to use an LLM (Claude) to generate contextual explanations for Pulumi resources, keeping documentation fresh and consistent with the registry.

## Usage

```markdown
{{< llm-explain
    resource="aws.autoscaling.Group"
    context="This guide explains how to set up EC2 Auto Scaling Groups with Pulumi"
    prompt="Explain the key configuration parameters for this resource. Return as a markdown list with bold parameter names followed by concise, actionable explanations. Focus on what users need to know for production deployments."
    example="**desiredCapacity**: The number of instances to maintain. Auto Scaling replaces failed instances to maintain this count. Set based on baseline load."
>}}
```

## Parameters

| Parameter | Required | Description |
|-----------|----------|-------------|
| `resource` | Yes | The Pulumi resource type (e.g., `aws.autoscaling.Group`) |
| `context` | No | Additional context about the guide or section |
| `prompt` | Yes | The specific question/task for the LLM |
| `example` | No | Example of desired output format |

## How It Works

### Build-Time Generation

1. **Fetch Resource Context**
   - Reads resource schema from registry
   - Includes properties, types, examples from API docs
   - Uses this as context for the LLM

2. **Build Prompt**
   ```
   System: You are a technical writer for Pulumi documentation.

   Resource Context:
   - Resource: aws.autoscaling.Group
   - [Schema, properties, examples from registry]

   Guide Context:
   - [User-provided context]

   Task:
   - [User-provided prompt]

   Example Output:
   - [User-provided example]
   ```

3. **Call Claude API**
   - Uses Anthropic Claude API (claude-3-5-sonnet-20241022)
   - Sends prompt with full context
   - Receives markdown response

4. **Cache Response**
   - Stores in: `static/registry/packages/llm-cache/{hash}.md`
   - Hash based on: resource + prompt + schema version
   - Avoids repeated API calls during builds
   - Regenerates when resource docs update

5. **Render Markdown**
   - LLM response is rendered as markdown in the page
   - Seamlessly integrated with manual content

### Fallback Strategy

```
┌─────────────────────────────────┐
│ Check for manual override first │
│ content/.../ llm-overrides/     │
└─────────────┬───────────────────┘
              │ Not found
              ▼
┌─────────────────────────────────┐
│ Check cache                     │
│ static/.../llm-cache/           │
└─────────────┬───────────────────┘
              │ Cache miss
              ▼
┌─────────────────────────────────┐
│ Call Claude API                 │
│ Cache response                  │
└─────────────┬───────────────────┘
              │ API error
              ▼
┌─────────────────────────────────┐
│ Show error message              │
│ "Content unavailable"           │
└─────────────────────────────────┘
```

## Example Use Cases

### 1. Configuration Details
```markdown
## Key configuration details

{{< llm-explain
    resource="aws.autoscaling.Group"
    prompt="Explain key configuration parameters as a markdown list."
    example="**parameter**: Brief explanation."
>}}
```

### 2. Best Practices
```markdown
## Best practices

{{< llm-explain
    resource="aws.autoscaling.Group"
    context="Production deployment of Auto Scaling Groups"
    prompt="List 5 best practices for production use. Focus on reliability, cost, and security."
>}}
```

### 3. Common Patterns
```markdown
## Common patterns

{{< llm-explain
    resource="aws.autoscaling.Group"
    prompt="Describe 3 common architectural patterns for using this resource. Include when to use each pattern."
>}}
```

### 4. Troubleshooting
```markdown
## Troubleshooting

{{< llm-explain
    resource="aws.autoscaling.Group"
    prompt="List common issues users face with this resource and how to resolve them."
>}}
```

## Benefits

### For Content Authors
- ✅ **No manual maintenance** - Content updates when resource changes
- ✅ **Consistent formatting** - LLM follows specified format
- ✅ **Contextual** - Based on actual registry docs
- ✅ **Flexible** - Can ask any question about the resource

### For Users
- ✅ **Always current** - Reflects latest resource capabilities
- ✅ **Comprehensive** - LLM can explain complex relationships
- ✅ **Actionable** - Focused on what users need to know

### For Pulumi
- ✅ **Scalable** - Can generate docs for thousands of resources
- ✅ **Quality** - Claude provides high-quality technical writing
- ✅ **Maintainable** - Changes propagate automatically

## Implementation Notes

### API Costs
- Build-time only (not per-pageview)
- Cached responses reused across builds
- Typical cost: $0.001-0.01 per section
- Budget: ~$100 can generate 10,000+ sections

### Cache Invalidation
Cache is invalidated when:
- Resource schema changes (version bump)
- Prompt changes
- Manual cache clear

### Manual Overrides
To override LLM content for a specific resource:

```bash
# Create override file
content/registry/packages/aws/llm-overrides/aws.autoscaling.Group-configuration.md
```

The shortcode will use this instead of calling the API.

### Environment Variables
```bash
# Required for real implementation
ANTHROPIC_API_KEY=sk-ant-...

# Optional
LLM_CACHE_DIR=static/registry/packages/llm-cache
LLM_MODEL=claude-3-5-sonnet-20241022
LLM_ENABLE=true
```

## Current Status: Prototype

The current implementation is a **dummy placeholder** that:
- Shows what parameters are being used
- Displays the prompt that would be sent
- Explains how caching will work
- Renders a blue placeholder box

To see it in action:
- Visit: http://localhost:1313/registry/packages/aws/guides/ec2-auto-scaling/
- Look for the blue box under "Key configuration details"

## Next Steps for Implementation

1. **API Integration**
   - Add Anthropic SDK dependency
   - Implement API client with rate limiting
   - Handle authentication

2. **Cache System**
   - Create cache directory structure
   - Implement hash-based keying
   - Add cache invalidation logic

3. **Schema Integration**
   - Read resource schemas from registry
   - Extract relevant properties/types
   - Format as context for LLM

4. **Error Handling**
   - Graceful degradation on API failures
   - Timeout handling
   - Retry logic

5. **Testing**
   - Test with various resource types
   - Verify output quality
   - Measure API costs

## Comparison: Before vs After

### Before (Manual)
```markdown
## Key configuration details

**Launch template**: Defines the instance configuration...
**Desired capacity**: The number of instances...
**Min and max size**: Define the boundaries...
[... 7 more manually written paragraphs ...]
```

- ❌ Manual effort for every resource
- ❌ Gets outdated quickly
- ❌ Inconsistent across resources
- ❌ No context about schema changes

### After (LLM-Generated)
```markdown
## Key configuration details

{{< llm-explain resource="aws.autoscaling.Group" prompt="..." >}}
```

- ✅ One line per section
- ✅ Always reflects latest docs
- ✅ Consistent quality
- ✅ Automatically regenerates
