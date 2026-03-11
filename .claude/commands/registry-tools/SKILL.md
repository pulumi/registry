---
description: List available registry tools and commands.
user-invocable: true
---

# Registry Tools

**Use this when:** You want to see all available Claude Code commands for the Pulumi Registry.

Lists all available commands grouped by category.

---

## Usage

`/registry-tools`

---

## Implementation

1. Run the metadata scraper to discover all available commands:

   ```bash
   python3 .claude/commands/registry-tools/scripts/scrape-metadata.py
   ```

2. Parse the JSON output and group commands by category:

   **Review & Quality**:
   - `/registry-review` - Review changes for quality
   - `/pr-review <PR_NUMBER>` - Full PR review workflow with approve/merge/close
   - `/validate-package [name]` - Validate provider package YAML

   **Package Management**:
   - `/new-package` - Guided workflow to add a new provider package
   - `/validate-package [name]` - Validate provider package YAML

   **Shipping & PRs**:
   - `/shipit` - Commit, push, and create PR with quality checks

   **Issue Management**:
   - `/fix-issue <issue-number>` - Analyze issue and create implementation plan

   **Repository Health**:
   - `/dashboard [mode]` - Repository status, PRs, issues, CI health

   **Discovery**:
   - `/registry-tools` - This command (list all available tools)

3. Display the grouped list in a terminal-friendly format:

   ```
   ════════════════════════════════════════════════════════════════════════════════
   📦 PULUMI REGISTRY - Available Commands
   ════════════════════════════════════════════════════════════════════════════════

   Review & Quality:
     /registry-review        Review changes for quality before committing
     /pr-review <PR>         Review and approve/merge pull requests
     /validate-package       Validate provider package YAML

   Package Management:
     /new-package            Add a new provider package to the registry

   Shipping & PRs:
     /shipit                 Commit, push, and create PR with quality checks

   Issue Management:
     /fix-issue <issue>      Analyze issue and create implementation plan

   Repository Health:
     /dashboard [mode]       Repository status and health metrics

   Discovery:
     /registry-tools         List all available commands (this command)

   ════════════════════════════════════════════════════════════════════════════════
   Tip: Run any command by typing it in the chat (e.g., /dashboard)
   ```

4. If the scraper discovers additional commands not in the hardcoded list above, include them in an "Other" category.
