---
user-invocable: false
---

# Quality Checks Reference

Quick reference for quality check procedures in the `/shipit` skill.

## Check Detection Matrix

| Content Changed | Mandatory Checks | Optional Suggestions |
|-------------------|------------------|---------------------|
| Go code (`tools/resourcedocsgen/`, `tools/mktutorial/`) | `make lint-go`, `make test` | - |
| Package YAML (`themes/default/data/registry/packages/`) | YAML validity, required fields check | `/validate-package <name>` |
| Markdown files | `make lint-markdown` | `/registry-review` |
| Theme JS/CSS (`themes/default/theme/`) | `make build-assets` | - |
| CI workflows (`.github/workflows/`) | YAML validity check | `/registry-review` |
| Community package list (`community-packages/`) | JSON syntax check (`python3 -m json.tool`) | - |
| Build scripts (`Makefile`, `scripts/`) | `make lint` | `/registry-review` |
| Other changes | `make lint` | Use your best judgment |

Any check above may be skipped at your discretion, e.g. if it was run recently or is unnecessary in context.

## Detecting Already-Run Checks

Scan last 10 conversation messages for:

**make lint patterns**: `"make lint"`, `"golangci-lint"`, `"markdownlint"`, `"lint passed"`

**make test patterns**: `"make test"`, `"go test"`, `"PASS"`, `"test passed"`

**make build-assets patterns**: `"make build-assets"`, `"yarn run build"`, `"build-assets"`

**/registry-review patterns**: `"/registry-review"`, `"[registry-review]"`

**Display if detected**: `[Step 2/8] Detected previous checks: lint ✓, test ✓ (skipping)`

## YAML Validation (Quick Check)

For package YAML files, perform a quick inline check:

1. Verify required fields are present: `name`, `title`, `description`, `publisher`, `repo_url`, `version`, `schema_file_url`, `category`, `component`, `native`, `package_status`
2. Verify `category` is one of: `Cloud`, `Database`, `Infrastructure`, `Monitoring`, `Network`, `Utility`, `Version Control System`
3. Verify `version` matches semver: `v\d+\.\d+\.\d+`
4. Verify `package_status` is `ga` or `public_preview`

### Results

- **Pass** ✓: All checks passed
- **Failed** ✗: Issues found (mandatory blocker for invalid YAML, warning for missing optional fields)

## JSON Validation (Community Package List)

```bash
python3 -m json.tool community-packages/package-list.json > /dev/null
```

## Summary Format

```
═══════════════════════════════════════════════════
Quality Checks Summary
═══════════════════════════════════════════════════
Mandatory Checks:
  ✓ make lint-go passed
  ✓ make test passed
  ✓ Package YAML validation passed
  ✓ JSON syntax check passed

Optional Suggestions:
  💡 Consider running /validate-package for full validation
  💡 Consider running /registry-review for content quality
═══════════════════════════════════════════════════
```

## Key Principles

- **Go code**: Strict (lint + tests must pass)
- **Package YAML**: Required fields must be present, values must be valid
- **Previous checks**: Skip if already run in this session
- **User control**: Allow overrides with acknowledgment
