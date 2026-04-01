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
- If the package already has `publisher: DEPRECATED`, inform the user it's already delisted

### Step 2: Update Package YAML

Set the `publisher` field to `DEPRECATED` in the package YAML file.

### Step 3: Remove from Community Packages List

Read `community-packages/package-list.json` and check if the package's repo slug appears in the `include` array.

- If found, remove the entry from the JSON array
- If not found, skip this step

### Step 4: Verify Changes

Run `git diff` to confirm:

- The `publisher` field in the YAML was changed to `DEPRECATED`
- The community packages list entry was removed (if applicable)

### Step 5: Report Results

Present a summary:

```
Package Delisted: <package-name>
═══════════════════════════════════════════════════

Changes:
  ✓ Set publisher to DEPRECATED in <package-name>.yaml
  ✓ Removed from community-packages/package-list.json (if applicable)

Note: The package YAML file is kept in the repo but will be
skipped by the publishing scripts (push-registry.py and
publish_to_registry.py) on the next deploy.
═══════════════════════════════════════════════════
```

## Background

Delisting works by setting `publisher: DEPRECATED` in the package YAML. Both publishing scripts (`scripts/ci/push-registry.py` and `scripts/ci/publish_to_registry.py`) skip packages with this publisher value, effectively removing them from the live registry without deleting the metadata files.
