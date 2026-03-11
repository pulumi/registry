---
user-invocable: false
description: JSON entry format for community packages list
---

# Community Package List Format

## File Location

`community-packages/package-list.json`

## Structure

The file is a JSON object with an `include` array:

```json
{
  "include": [
    {
      "repoSlug": "{owner}/{repo}",
      "schemaFile": "{path-to-schema-file}"
    }
  ]
}
```

## Entry Format

Each entry has two required fields:

| Field | Description | Example |
|---|---|---|
| `repoSlug` | GitHub `{owner}/{repo}` format (without `https://github.com/`) | `pulumiverse/pulumi-example` |
| `schemaFile` | Path to the schema JSON file within the repository | `provider/cmd/pulumi-resource-example/schema.json` |

## Example Entry

```json
{
  "repoSlug": "koyeb/pulumi-koyeb",
  "schemaFile": "provider/cmd/pulumi-resource-koyeb/schema.json"
}
```

## Insertion Rules

1. Add the new entry to the `include` array
2. Maintain alphabetical order by `repoSlug`
3. Ensure valid JSON (proper comma placement)
4. No duplicate `repoSlug` entries

## Validation

After adding an entry, verify:
```bash
python3 -m json.tool community-packages/package-list.json > /dev/null
```
