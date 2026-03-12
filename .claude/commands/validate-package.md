---
description: Validate a provider package YAML file for correctness and completeness.
---

# Validate Package

**Use this when:** You want to verify that a provider package YAML file has all required fields, valid values, and is properly configured.

---

## Usage

`/validate-package [package-name]`

If `package-name` is not provided, auto-detects from changed files (`git status` or branch diff) that touch `themes/default/data/registry/packages/`.

---

## Process

### Step 1: Locate Package YAML

**If package name provided** ({{arg}}):

- Look for `themes/default/data/registry/packages/{{arg}}.yaml`
- If not found, report error and exit

**If no package name provided**:

- Check `git status --short` for changed files matching `themes/default/data/registry/packages/*.yaml`
- Check `git diff --name-only master...HEAD` for branch changes
- If multiple packages found, list them and ask which to validate
- If none found, ask user to specify a package name

### Step 2: Read and Parse YAML

Read the package YAML file and verify it is well-formed YAML.

### Step 3: Validate Required Fields

Check that all required fields are present:

| Field | Required | Type |
|---|---|---|
| `name` | Yes | string |
| `title` | Yes | string |
| `description` | Yes | string |
| `publisher` | Yes | string |
| `repo_url` | Yes | string |
| `version` | Yes | string |
| `schema_file_url` | Yes | string |
| `category` | Yes | string |
| `component` | Yes | boolean |
| `native` | Yes | boolean |
| `package_status` | Yes | string |

Optional fields: `logo_url`, `featured`, `updated_on`

### Step 4: Validate Field Values

- **`category`**: Must be one of: `Cloud`, `Database`, `Infrastructure`, `Monitoring`, `Network`, `Utility`, `Version Control System`
- **`package_status`**: Must be `ga` or `public_preview`
- **`version`**: Must match semver pattern `v\d+\.\d+\.\d+` (optional pre-release suffix like `-alpha.1`)
- **`repo_url`**: Must be a valid GitHub URL (`https://github.com/{owner}/{repo}`)
- **`component`** and **`native`**: Must be boolean (`true` or `false`)
- **`name`**: Must be lowercase, alphanumeric with hyphens only
- **`schema_file_url`**: Must be a valid URL

### Step 5: Validate Schema URL Reachability

Use WebFetch to check if `schema_file_url` is reachable:

- Fetch the URL (HEAD request if possible, or read first few bytes)
- Report if the URL returns an error or is unreachable
- This is a warning, not a blocker (URL may require authentication or be temporarily unavailable)

### Step 6: Cross-Reference Community Packages

If the package appears to be a community package (publisher is not "Pulumi"):

- Read `community-packages/package-list.json`
- Check if the package's repo slug appears in the list
- If present, verify the `schemaFile` path looks reasonable
- If absent, note that the package may need to be added to the community package list

### Step 7: Optional API Doc Generation

Ask the user if they want to test API doc generation:

- If yes, run `make api-docs/<package-name>`
- Report success or failure
- This step is optional because it can take significant time

### Step 8: Report Results

Present a summary:

```
Package Validation: <package-name>
═══════════════════════════════════════════════════

Required Fields:
  ✓ name: <value>
  ✓ title: <value>
  ✓ category: <value>
  ✗ version: missing or invalid

Value Validation:
  ✓ category is valid
  ✓ version matches semver
  ⚠ schema_file_url: unreachable (warning)

Cross-References:
  ✓ Found in community-packages/package-list.json

Overall: [PASS / WARN / FAIL]
═══════════════════════════════════════════════════
```

- **PASS**: All required fields present, all values valid
- **WARN**: All required fields present, but warnings (e.g., unreachable URL)
- **FAIL**: Missing required fields or invalid values
