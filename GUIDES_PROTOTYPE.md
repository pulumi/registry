# Guides Section Prototype

This branch (`resource-examples-export`) contains a prototype implementation of a new "Guides" section for the Pulumi Registry, specifically for AWS packages.

## What Was Done

### 1. Created Guides Structure
- **Location**: `content/registry/packages/aws/guides/`
- Created a new section within the AWS package for guides
- Added a `_index.md` list page that explains what guides are
- Note: Content lives in root `content/` directory, not `themes/default/content/`

### 2. Ported EC2 Auto-Scaling Guide
- **File**: `content/registry/packages/aws/guides/ec2-auto-scaling.md`
- Ported from the pulumi/docs PR #16229
- Contains:
  - Multi-language code examples (TypeScript, Python, Go, C#, Java)
  - FAQ section with common questions about Auto Scaling
  - Frontmatter with metadata (category, tags, etc.)

### 3. Created Layout Templates
- **List Page**: `themes/default/layouts/guides/list.html`
  - Shows all guides grouped by category
  - Includes hero section, category icons, and CTA

- **Single Page**: `themes/default/layouts/guides/single.html`
  - Individual guide page with sidebar
  - Table of contents, quick links, CTAs
  - Breadcrumb navigation

## URLs Generated

With this structure, the following URLs are created:
- List page: `/registry/packages/aws/guides/`
- Individual guide: `/registry/packages/aws/guides/ec2-auto-scaling/`

## Testing

The prototype successfully builds with Hugo:
```bash
hugo --quiet --buildDrafts
```

Generated pages are in `public/registry/packages/aws/guides/`

## Next Steps for Resource Examples Export

The original goal was to create a system where guides can reference resource examples from the API docs. This would involve:

1. **Export Examples from resourcedocsgen**
   - Modify `tools/resourcedocsgen/pkg/docs/gen.go` to add `GenerateExamplesIndex()`
   - Export examples as JSON to `static/registry/packages/examples/{provider}/{resource}.json`

2. **Create Shortcode**
   - Add `themes/default/layouts/shortcodes/resource-example.html`
   - Reads the exported JSON and renders multi-language examples
   - Usage: `{{< resource-example provider="aws" resource="autoscaling/group" >}}`

3. **Integrate into Guides**
   - Guides can use the shortcode instead of manually embedding code
   - Examples stay in sync with API docs generation

## File Structure

```
content/registry/packages/aws/guides/
├── _index.md                         # Guides list page
└── ec2-auto-scaling.md               # Individual guide

themes/default/layouts/guides/
├── list.html                         # Layout for guides list
└── single.html                       # Layout for individual guides
```

## Design Decisions

1. **Placed under AWS package** - Guides are specific to AWS, so they live at `/registry/packages/aws/guides/` rather than a top-level `/guides/` section

2. **Separate from how-to-guides** - The existing `how-to-guides` section contains tutorials, while these guides are meant to be quick, copy-paste solutions

3. **Layout inheritance** - The layouts reference the standard Hugo layout blocks (`hero`, `main`) but with custom styling from the docs repo

## Notes

- The `content/` directory at the root is git-ignored by default (see .gitignore)
- Content should be added to `content/` (not `themes/default/content/`)
- Files will need to be added with `git add -f` to include them since content/ is ignored
- The layouts use TailwindCSS classes that may need verification against the registry's CSS setup
