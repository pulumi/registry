---
description: Finalize and ship changes with quality checks, commits, and PR creation
---

# Ship It!

**Use this when:** You're ready to commit your work and create a pull request.

Performs comprehensive quality checks (lint, build, code compilation), creates meaningful commits, and handles the complete workflow from staging to PR creation with multiple approval checkpoints.

---

## Usage

`/shipit`

Finalizes your current work by running quality checks, committing changes, pushing to remote, and creating a pull request.

**Required**: Uncommitted changes in the working directory

**Works with**: All file types (Go code, YAML, CI workflows, scripts, theme, infrastructure)

**Special handling**: Detects and prevents master branch commits, validates Go code and YAML, and scans conversation history to avoid redundant quality checks.

---

## Process

**CRITICAL SUCCESS CRITERIA**: Complete all 8 steps in sequence. Every step is mandatory and serves a critical purpose in the workflow. **DO NOT SKIP ANY STEP OR END THE WORKFLOW PREMATURELY!**

**Step Counter**: Display progress before each step as: **[Step X/8]** followed by the step heading. This helps users track progress through the workflow.

**References**: This skill uses detailed reference files. Always follow the detailed instructions in these referenced documents when applicable:
- `shipit:references:quality-checks` - Quality check procedures and code testing
- `shipit:references:commit-messages` - Commit message style guidelines

---

## **[Step 1/8] Context Gathering**

**Purpose**: Understand what's changed and check for next steps from previous work.

**Actions**:
1. Run `git status --short` to see all changes
2. Run `git branch --show-current` to verify current branch
3. Scan last 5 conversation messages for "next steps" or "TODO" mentions
4. Categorize changed files:
   - Package YAML (`themes/default/data/registry/packages/`)
   - Go tools (`tools/resourcedocsgen/`, `tools/mktutorial/`)
   - CI workflows (`.github/workflows/`)
   - Build scripts (`Makefile`, `scripts/`)
   - Infrastructure (`infrastructure/`)
   - Theme/layout (`themes/default/layouts/`, `themes/default/theme/`)
   - Community packages (`community-packages/`)
   - Other

**Output**: Display summary of changes with file counts per category

---

## **[Step 2/8] Quality Checks**

Run context-aware quality checks. Scan conversation history to skip redundant checks (lint/build), and run appropriate checks based on file types changed.

See `shipit:references:quality-checks` for complete procedures, detection patterns, and error handling.

Display summary and ask user to proceed with `AskUserQuestion`.

---

## **[Step 3/8] Branch Verification**

**Purpose**: Prevent accidental commits to master branch.

**Actions**:

1. Check current branch: `git branch --show-current`

2. **If on `master` or `main`**:
   - Display warning with current git status
   - Use `AskUserQuestion` with options:
     1. Create new branch from current state (Recommended)
        - Prompts for branch name with format: `CamSoper/{descriptive-name}`
        - Runs: `git checkout -b CamSoper/{name}`
     2. Reset to origin/master and create clean branch
        - Warns this will discard all changes
        - Runs: `git reset --hard origin/master`
        - Then prompts for branch name and creates it
     3. Cancel (I'll handle manually)
        - Exits the skill

3. **If on feature branch**:
   - Display: "[Step 3/8] Skipped - already on feature branch `{branch-name}`"
   - Continue to Step 4

**Safety**: Always preview destructive operations (like reset --hard) before executing.

---

## **[Step 4/8] Commit Preparation**

Generate 3 meaningful commit message suggestions based on:
- `git diff --stat` output
- Conversation context (user's stated goal from last 10 messages)
- Recent commit style (`git log --oneline -5`)
- File types changed

See `shipit:references:commit-messages` for message format, prefix guidelines, and generation strategies.

Present suggestions with `AskUserQuestion` (3 options + custom message).

All messages include: `Co-Authored-By: Claude Opus 4.6 <noreply@anthropic.com>`

---

## **[Step 5/8] Commit Preview**

**Purpose**: Preview exactly what will be committed before execution.

**Display preview**:

```
Ready to commit:
Branch: CamSoper/fix-yaml-validation

Files to be committed:
  M themes/default/data/registry/packages/example.yaml
  M tools/resourcedocsgen/cmd/generate.go
  A tools/resourcedocsgen/cmd/generate_test.go

Commit message:
─────────────────────────────────────────
Fix YAML validation for community packages

Corrects category field validation in resourcedocsgen.

Co-Authored-By: Claude Opus 4.6 <noreply@anthropic.com>
─────────────────────────────────────────

Commands that will run:
  git add themes/default/data/registry/packages/example.yaml tools/resourcedocsgen/cmd/generate.go tools/resourcedocsgen/cmd/generate_test.go
  git commit -m "$(cat <<'EOF'
  Fix YAML validation for community packages

  Corrects category field validation in resourcedocsgen.

  Co-Authored-By: Claude Opus 4.6 <noreply@anthropic.com>
  EOF
  )"
```

**Ask for confirmation**:
- Use `AskUserQuestion` with options:
  1. Commit (Recommended) - Proceed to Step 6
  2. Edit message - Return to Step 4 to select different message
  3. Cancel - Exit without committing

**If "Edit message"**: Loop back to Step 4

**If "Cancel"**: Exit skill with message about uncommitted changes

---

## **[Step 6/8] Push Changes**

**Purpose**: Commit and push changes to remote.

**Actions** (execute sequentially):

1. **Stage specific files**:
   ```bash
   git add file1.yaml file2.go file3_test.go ...
   ```
   Never use `git add .` or `git add -A` - always list specific files

2. **Commit with heredoc**:
   ```bash
   git commit -m "$(cat <<'EOF'
   {commit message from Step 5}

   Co-Authored-By: Claude Opus 4.6 <noreply@anthropic.com>
   EOF
   )"
   ```

3. **Check remote tracking**:
   ```bash
   git rev-parse --abbrev-ref @{u} 2>/dev/null
   ```
   If fails (no tracking branch): Need to push with `-u`

4. **Push to remote**:
   ```bash
   # If no tracking branch:
   git push -u origin {branch-name}

   # If tracking branch exists:
   git push
   ```

5. **Verify push**:
   ```bash
   git status
   ```
   Should show "Your branch is up to date with 'origin/{branch}'"

**Display confirmation**:
```
✓ Committed changes (abc123d)
✓ Pushed to origin/CamSoper/fix-yaml-validation
```

**If any step fails**:
- Display error
- Offer to retry or cancel
- Don't proceed to Step 7 if push failed

---

## **[Step 7/8] Create Pull Request**

**Purpose**: Generate and create a pull request with appropriate description.

**Actions**:

1. **Generate PR description**:

   **Title**: Use first line of commit message (without co-author line)

   **Body**:
   - Brief summary (1-2 sentences from commit message)
   - Changes included (from `git diff --stat master...HEAD`)
   - Keep concise (aim for < 500 characters)
   - Include issue references if applicable (e.g., "Fixes #123")
   - No "Generated with Claude Code" footer

   Example:
   ```markdown
   Corrects category field validation in resourcedocsgen for community packages.

   ## Changes
   - Fixed YAML validation for category enum values
   - Added unit tests for edge cases
   - Updated example package YAML
   ```

2. **Display PR preview**:
   ```
   Pull Request Preview:
   ─────────────────────────────────────────
   Title: Fix YAML validation for community packages

   Body:
   Corrects category field validation in resourcedocsgen for community packages.

   ## Changes
   - Fixed YAML validation for category enum values
   - Added unit tests for edge cases
   - Updated example package YAML
   ─────────────────────────────────────────
   ```

3. **Ask for confirmation**:
   - Use `AskUserQuestion` with options:
     1. Create PR (Recommended) - Execute pr creation
     2. Edit description - Allow custom title/body input, show preview again
     3. Cancel - Exit without creating PR (changes are still committed/pushed)

4. **If "Edit description"**:
   - Ask for custom title (or keep current)
   - Ask for custom body (or keep current)
   - Show updated preview
   - Loop back to confirmation

5. **If "Create PR"**:
   ```bash
   gh pr create --title "{title}" --body "$(cat <<'EOF'
   {body}
   EOF
   )"
   ```

**If PR creation fails**:
- Display error (auth issues, network, etc.)
- Note that changes are still committed and pushed
- Suggest manual PR creation with URL:
  `https://github.com/pulumi/registry/pull/new/{branch-name}`

---

## **[Step 8/8] Completion Report**

**Purpose**: Confirm successful completion and provide next steps.

**Actions**:

1. **Fetch PR details**:
   ```bash
   gh pr view --json url,number,title
   ```

2. **Check for "next steps" in conversation**:
   - Scan last 10 messages for mentions of "next steps", "TODO", "follow-up"
   - Extract any action items that were mentioned

3. **Display success message**:
   ```
   Successfully shipped your changes!

   Pull Request: #1234
   URL: https://github.com/pulumi/registry/pull/1234
   Title: Fix YAML validation for community packages
   Branch: CamSoper/fix-yaml-validation

   Next steps:
   - Monitor PR for CI/CD status
   - {any next steps from conversation}
   ```

4. **Celebrate**:
   ```
   Ship it! Your changes are ready for review.
   ```

**End of skill**

---

## Critical Workflow Rules

1. **Progress Display**: Always display step number "[Step X/8]" before each heading
2. **Sequential Execution**: Never skip ahead - complete each step before moving to next
3. **User Approval**: Get explicit approval before destructive actions (Steps 3, 5, 7)
4. **Error Handling**: If a step fails, don't proceed - offer retry or cancel
5. **Skip Display**: If skipping a step, show "[Step X/8] Skipped - {reason}"
6. **Context Preservation**: Store decisions from previous steps to avoid re-asking

## Notes

- This skill respects the `/AGENTS.md` repository guidelines
- Commit messages always include co-author attribution for Claude
- PR descriptions never include "Generated with Claude Code"
- Branch names follow `CamSoper/{descriptive-name}` format
- All git operations use safe practices (no force push, no destructive defaults)
