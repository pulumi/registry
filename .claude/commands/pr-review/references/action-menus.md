---
user-invocable: false
description: Action menu options for bot and non-bot PRs
---

# Action Menus

**Decision Point**: Select the appropriate section based on contributor type.

## Bot PRs (pulumi-bot, renovate, github-actions, etc.)

### Display Header

```text
🤖 Bot: @username
[If automation/tfgen-provider-docs label] ✓ Automated metadata update
```

### Options (Max 4)

1. **Approve and merge** (Recommended for metadata-only PRs) - Ready for immediate merge
2. **Approve** - Changes acceptable, manual merge later
3. **Request changes** - Issues need addressing
4. **Do nothing yet** - Need investigation

### Important Notes

- Default to "Approve and merge" for pulumi-bot PRs that only touch package YAML files
- If "Close PR" is needed, user can select "Do nothing yet" and close manually
- "Make changes and approve" excluded for bots (editing breaks automation; bot PRs regenerated, not manually edited)

## Standard Action Menu (Non-Bot Contributors)

**IMPORTANT**: AskUserQuestion only supports max 4 options. Use adaptive menus based on review findings.

### Adaptive Menu Selection

Choose the appropriate menu based on review findings:

- **Scenario A**: Issues found → Request changes menu
- **Scenario B**: Clean review → Approve menu
- **Scenario C**: Should close → Close menu

### Scenario A: Issues Found - Request Changes Recommended

**Options**:
1. **Request changes** (Recommended) - Author should address issues
2. **Make changes and approve** - Fix issues yourself + approve
3. **Approve** - Approve despite issues (non-blocking feedback)
4. **Do nothing yet** - Need more time/discussion

### Scenario B: Clean Review - Approve Recommended

**Options**:
1. **Approve** (Recommended) - Approve PR with feedback
2. **Approve and merge** - Approve + merge (squash) for immediate merge
3. **Make changes and approve** - Minor edits (typos, formatting) + approve
4. **Do nothing yet** - Need more time/discussion

### Scenario C: Should Close - Close PR Recommended

**Options**:
1. **Close PR** (Recommended) - Close with explanation
2. **Request changes** - Give author chance to address issues
3. **Approve** - Override concerns and approve anyway
4. **Do nothing yet** - Need discussion before closing

## Implementation Notes

- Always use AskUserQuestion tool with max 4 options
- Select appropriate menu based on context (review findings, contributor type, PR state)
- Tone adjusts based on contributor type (external vs internal)
- "Make changes and approve" preserves contributor credit for minor fixes
- Bot PRs exclude "Make changes and approve" (breaks automation)
