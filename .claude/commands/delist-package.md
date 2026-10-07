---
description: Delist (deprecate) a provider package from the Pulumi Registry.
---

# Delist Package

**Use this when:** You want to remove a package from the Pulumi Registry by marking it as deprecated.

---

## Usage

`/delist-package [package-name]`

If `package-name` is not provided, ask the user which package to delist.

---

## Process

### Step 1: Locate Package YAML

Look for `themes/default/data/registry/packages/{{arg}}.yaml`.

- If not found, report error and exit
- If the package already has `deprecated: true`, inform the user it's already delisted

### Step 2: Update Package YAML

Add `deprecated: true` to the package YAML file, keeping keys in alphabetical order (it goes after `component`). Leave `publisher` as it is.

Ask the user whether a registry package replaces this one. If so, add `superseded_by: <replacement-name>` to the YAML, where `<replacement-name>` is the replacement's YAML file name without `.yaml`. The package's pages render a deprecation banner from `deprecated: true` and link to the replacement named here; the build fails if `superseded_by` names a package that doesn't exist. Don't add "(Deprecated)" to the title in `_index.md`: the templates append it.

Ask the user why the package is deprecated. If they give a reason, add it as a single sentence in `deprecation_reason`, right after `deprecated: true`; quote the value if it contains a colon. It follows "This package is deprecated." in the banner on the package's pages and in the tooltip on the badge in the package list, so don't repeat that phrase or name the replacement in it.

### Step 3: Remove from Community Packages List

Read `community-packages/package-list.json` and check if the package's repo slug appears in the `include` array.

- If found, remove the entry from the JSON array
- If not found, skip this step

### Step 4: Check for a Dynamically Bridged Provider

If the YAML's `schema_file_url` contains `registry.opentofu.org` (its `description` will also read `A Pulumi provider dynamically bridged from <name>.`), the package is a dynamically bridged Terraform provider fed by the internal `pulumi/terraform-to-pulumi-registry-pipeline` repo. `resourcedocsgen metadata` leaves a delisted package's files alone, but as long as the provider stays on the pipeline's list, every new version it ships posts a warning in this repo's workflow runs.

- Tell the user the package also has to be removed from `watched-providers` in that repo's `Pulumi.yaml`, and offer to file an issue there
- If the URL doesn't contain `registry.opentofu.org`, skip this step

### Step 5: Verify Changes

Run `git diff` to confirm:

- The YAML has `deprecated: true`
- The community packages list entry was removed (if applicable)

### Step 6: Report Results

Present a summary:

```
Package Delisted: <package-name>
═══════════════════════════════════════════════════

Changes:
  ✓ Set deprecated: true in <package-name>.yaml
  ✓ Set superseded_by to <replacement> (if applicable)
  ✓ Set deprecation_reason (if applicable)
  ✓ Removed from community-packages/package-list.json (if applicable)
  ! Dynamically bridged: remove from watched-providers in
    pulumi/terraform-to-pulumi-registry-pipeline (if applicable)

Note: The package YAML file is kept in the repo but will be
skipped by the publishing scripts (push-registry.py and
publish_to_registry.py) on the next deploy.
═══════════════════════════════════════════════════
```

## Background

Delisting works by setting `deprecated: true` in the package YAML. `resourcedocsgen metadata` then skips the package with a warning instead of regenerating its YAML and docs, since a delisted package shouldn't receive updates. Both publishing scripts (`scripts/ci/push-registry.py` and `scripts/ci/publish_to_registry.py`) skip packages with this field set, effectively removing them from the live registry without deleting the metadata files. On the site, the package gets a "Deprecated" badge in the package list (where it's hidden unless the Deprecated filter is on), a deprecation banner on its pages, and " (Deprecated)" after its title.
