---
description: Guided workflow to add a new provider package to the Pulumi Registry
user-invocable: true
---

# New Package

**Use this when:** You want to add a new provider package to the Pulumi Registry, either as a community package or a first-party Pulumi provider.

---

## Usage

`/new-package`

Walks through a guided process to create the package YAML metadata, validate the provider, and optionally add it to the community package list.

---

## Process

**CRITICAL SUCCESS CRITERIA**: Complete all 7 steps in sequence. Every step is mandatory. **DO NOT SKIP ANY STEP OR END THE WORKFLOW PREMATURELY!**

**Step Counter**: Display progress before each step as: **[Step X/7]** followed by the step heading.

**References**: This skill uses reference files for templates and validation:
- `new-package:references:yaml-template` - Canonical YAML structure
- `new-package:references:package-list-format` - Community package list entry format
- `new-package:references:first-party-checklist` - First-party provider setup checklist

---

## **[Step 1/7] Determine Package Type**

Use `AskUserQuestion` to ask:

**Question**: "What type of package are you adding?"

**Options**:
1. **Community package** - Third-party provider published by the community (e.g., pulumiverse)
2. **First-party Pulumi provider** - Official Pulumi provider (published by Pulumi)

Store the selection for later steps.

---

## **[Step 2/7] Gather Package Information**

Use `AskUserQuestion` to collect the following (ask in sequence, one at a time):

1. **Package name** — The short name for the provider (lowercase, hyphens allowed). Example: `datadog`, `cloudflare`, `my-provider`
2. **Repository URL** — The GitHub repository URL. Example: `https://github.com/pulumiverse/pulumi-example`
3. **Display title** — The human-readable title. Example: `Datadog`, `Cloudflare`, `My Provider`
4. **Description** — Brief description of what the provider does (1-2 sentences)
5. **Publisher** — The publisher name. Example: `Pulumi`, `pulumiverse`, `MyOrg`
6. **Category** — One of: `Cloud`, `Database`, `Infrastructure`, `Monitoring`, `Network`, `Utility`, `Version Control System`
7. **Schema file path** — Path to the schema JSON file within the repo. Example: `provider/cmd/pulumi-resource-example/schema.json`

For community packages, also ask:
8. **Is this a component provider?** (yes/no)
9. **Is this a native provider?** (yes/no)

For first-party packages, default `component: false` and `native: false` (can be overridden).

---

## **[Step 3/7] Validate Repository**

Verify the repository exists and is accessible:

```bash
gh api repos/{owner}/{repo} --jq '.full_name'
```

- If the repo exists, display confirmation with repo details (description, stars, language)
- If the repo doesn't exist or is inaccessible, report error and ask user to verify the URL
- Check if the repo is public (required for community packages)

---

## **[Step 4/7] Validate Schema File**

Construct the schema file URL and verify it's reachable:

For GitHub repos, construct the raw URL:
```
https://raw.githubusercontent.com/{owner}/{repo}/{default-branch}/{schema-file-path}
```

Use WebFetch to check if the URL returns valid content:
- If reachable, display confirmation
- If not reachable, warn the user and ask if they want to continue with a different URL or path
- Try common default branches: `main`, `master`

---

## **[Step 5/7] Create Package YAML**

Generate the YAML file using the template from `new-package:references:yaml-template`.

**Actions**:
1. Determine the latest version by checking the repo's releases:
   ```bash
   gh api repos/{owner}/{repo}/releases/latest --jq '.tag_name' 2>/dev/null || echo "v0.0.1"
   ```

2. Generate the YAML content with all gathered information

3. Display the complete YAML for review:
   ```
   Package YAML Preview:
   ─────────────────────────────────────────
   [YAML content]
   ─────────────────────────────────────────

   File: themes/default/data/registry/packages/{name}.yaml
   ```

4. Use `AskUserQuestion` to confirm:
   - **Create file** (Recommended) - Write the YAML file
   - **Edit values** - Modify specific fields
   - **Cancel** - Exit without creating

5. Write the file to `themes/default/data/registry/packages/{name}.yaml`

---

## **[Step 6/7] Update Package List or Note First-Party Setup**

### Community Package Path

Add an entry to `community-packages/package-list.json`:

1. Read the current file
2. Add the new entry in the correct alphabetical position within the `include` array
3. Display the change for confirmation
4. Write the updated file

See `new-package:references:package-list-format` for the entry format.

### First-Party Package Path

Skip `package-list.json` modification. Instead, display the first-party setup checklist:

See `new-package:references:first-party-checklist` for the complete checklist.

**Key message**: "First-party providers use `repository_dispatch` events. The provider repository needs to be configured to send `resource-provider` dispatch events to `pulumi/registry` via the `publish-provider-update.yml` workflow."

---

## **[Step 7/7] Validate and Next Steps**

1. Run `/validate-package {name}` to verify the new package

2. Display completion summary:
   ```
   New package added successfully!

   Package: {name}
   Type: {community|first-party}
   File: themes/default/data/registry/packages/{name}.yaml
   {If community: Added to community-packages/package-list.json}

   Next steps:
   - Review the generated YAML file
   - Run /shipit to commit and create a PR
   {If first-party: - Configure repository_dispatch in the provider repo}
   ```

3. Offer to run `/shipit` to commit and create a PR

---

## Critical Workflow Rules

1. **Progress Display**: Always display step number "[Step X/7]" before each heading
2. **Sequential Execution**: Never skip ahead - complete each step before moving to next
3. **User Approval**: Get explicit approval before writing files (Step 5)
4. **Validation**: Always run validate-package at the end
5. **Error Handling**: If a step fails, don't proceed - offer retry or cancel
