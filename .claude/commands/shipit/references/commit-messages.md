---
user-invocable: false
---

# Commit Message Guidelines

Quick reference for generating meaningful commit messages that follow repository conventions.

## Format

```
{prefix}: {brief description}

{optional longer description}

Co-Authored-By: Claude Opus 4.6 <noreply@anthropic.com>
```

## General Principles

1. Use imperative mood: "Add feature" not "Added feature"
2. Keep first line under 72 characters
3. Be specific but concise
4. Always include co-author line at end

## Prefix Guidelines

| Prefix | When to Use | Example |
|--------|-------------|---------|
| `feat:` | New feature or capability | `feat: Add package validation command` |
| `fix:` | Bug fix | `fix: Correct category enum in YAML validation` |
| `docs:` | Documentation only changes | `docs: Update BUILD-AND-DEPLOY.md for new workflow` |
| `refactor:` | Code refactoring (no behavior change) | `refactor: Simplify schema URL resolution` |
| `test:` | Adding or updating tests | `test: Add unit tests for YAML parser` |
| `chore:` | Maintenance, dependencies, build | `chore: Update Go dependencies` |
| `ci:` | CI/CD workflow changes | `ci: Add Claude code review workflow` |

**No prefix**: For registry-specific changes, prefixes are often optional:

- `Update aws to v7.21.0`
- `Add new community provider: example-provider`
- `Fix broken schema URL for random provider`

## Context-Based Patterns

**Package YAML** (`themes/default/data/registry/packages/`):

- `Update {package} to v{version}`
- `Add {package} provider to registry`
- `Fix {package} metadata`
- Example: `Update datarobot to v0.10.30`

**Go tools** (`tools/resourcedocsgen/`, `tools/mktutorial/`):

- `Fix {issue} in resourcedocsgen`
- `Add {feature} to doc generation`
- Example: `Fix schema parsing for component providers`

**CI workflows** (`.github/workflows/`):

- `ci: Add {workflow} workflow`
- `ci: Fix {workflow} trigger conditions`
- Example: `ci: Add Claude code review on PR open`

**Infrastructure** (`infrastructure/`):

- `Update {resource} configuration`
- `Fix CloudFront origin configuration`

**Community packages** (`community-packages/`):

- `Add {package} to community package list`
- `Update community package list`

## When to Add Body

Add longer description when:

- Change affects multiple areas
- Context is important for reviewers
- Breaking changes introduced
- Non-obvious reasons for change

Example:

```
Fix schema URL resolution for native providers

Native providers use a different schema file path convention.
Updated resourcedocsgen to handle both native and bridge patterns.

Co-Authored-By: Claude Opus 4.6 <noreply@anthropic.com>
```

## Issue References

- `(#1234)` - PR number
- `Fixes #1234` - Closes issue
- `Closes #1234` - Closes issue

## Generation Strategy

Generate 3 options:

**Option 1**: Conventional commit with prefix (if applicable) + body if helpful
**Option 2**: Repository style (plain imperative, matches recent commits)
**Option 3**: Alternative phrasing (different emphasis, same meaning)

Consider:

- File types changed (YAML/Go/CI/scripts)
- Conversation context (what was the goal?)
- Recent commit patterns (`git log --oneline -5`)
- Scope (single file vs. feature-wide)

## Anti-Patterns

Avoid:

- Past tense: "Added feature" (use imperative: "Add feature")
- Vague: "Updated files", "Fixed stuff", "Changes"
- Too long: First line > 72 chars
- Incomplete: "WIP", "TODO", "test"

## Co-Author Line

**Always include at the end:**

```
Co-Authored-By: Claude <<Model and version>> <noreply@anthropic.com>
```
