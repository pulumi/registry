---
description: Analyze a registry issue and create an implementation plan (investigates problem and plans solution).
---

# Usage

`/fix-issue <issue-number>`

**Use this when:** You need to address a GitHub issue and want help investigating the problem and creating a step-by-step implementation plan.

Analyzes GitHub issue #{{arg}} and creates a comprehensive plan to address it, acting as a technical engineer focused on improving the Pulumi Registry.

---

## Process

### 1. Gather issue details

Use `gh issue view {{arg}}` to retrieve the complete issue information:

- Read the full issue description to understand the problem
- Review labels to identify issue type (bug, enhancement, clarification, etc.)
- Read all comments and discussion to understand user pain points
- Note any related issues or PRs mentioned
- Identify the affected area (provider YAML, Go tools, CI/CD, build system, infrastructure, theme/layout)

### 2. Research current state

Use semantic search and file tools to thoroughly investigate:

- Locate the specific files mentioned or affected
- Read the current content to understand what exists today
- Identify what's missing, incorrect, outdated, or broken
- Review related files for consistency
- Search for similar patterns elsewhere in the codebase that might have the same issue
- Check if the issue affects generated content vs. source files

### 3. Review guidelines

Before proposing changes, verify compliance with:

- `AGENTS.md` — Repository conventions, build/test workflow, architecture, file paths
- `BUILD-AND-DEPLOY.md` — Build system details, CI/CD pipeline, deployment model (read only the sections relevant to the issue)
- Existing patterns in the same area of the codebase to maintain consistency

### 4. Create implementation plan

Develop a detailed plan with these components:

1. **Analysis** — Clearly explain:
   - What the current state is (or what's missing)
   - Why it's problematic
   - What impact this has on users or the system
   - The root cause of the issue

2. **Solution** — Describe the specific changes:
   - Code to write or modify
   - Configuration to update
   - YAML metadata to fix
   - CI workflows to adjust

3. **Files to modify** — List exact file paths and locations:
   - Full paths from repository root (e.g., `tools/resourcedocsgen/cmd/generate.go`)
   - Specific functions, sections, or line numbers where changes occur
   - New files to create (if any)

4. **Implementation steps** — Break work into logical, ordered tasks:
   - Update source files
   - Modify configurations
   - Add or fix tests
   - Update documentation if needed

5. **Verification** — Define how to validate the fix:
   - Run `make lint` to check for lint errors
   - Run `make test` to verify Go tests pass
   - Run `make lint-go` for Go-specific linting
   - Run `make api-docs/<package>` if provider changes are involved
   - Run `make build-assets` if theme changes are involved
   - Run `make lint-markdown` if markdown files changed
   - Verify all links work (internal and external)

6. **Considerations** — Note special requirements or impacts:
   - Does this affect the nightly package update workflow (`generate-package-metadata.yml`)?
   - Does this affect `push-registry.py` (live registry publication)?
   - Does `BUILD-AND-DEPLOY.md` need updating?
   - Does this affect the CI/CD pipeline or deployment infrastructure?
   - Could this break existing provider documentation generation?
   - Are there related issues that should be addressed together?

### 5. Present for approval

**IMPORTANT: Do NOT implement without user approval.**

Present the complete plan to the user including:

- Clear, numbered steps in logical order
- Exact file paths with section headings or line numbers
- Specific commands that will be run
- Any assumptions made or decisions that need user input
- Estimated scope (minor fix vs. major restructuring)

If the issue is ambiguous, ask clarifying questions before creating the plan. If multiple valid approaches exist, present options with pros/cons and ask for the user's preference.
