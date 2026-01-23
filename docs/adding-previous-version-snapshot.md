# Adding a Previous Version Snapshot

When creating a new major release of a tier-1 provider, we might want to keep the previous major version available in the Pulumi Registry for users who are not yet ready to upgrade to the new major version.

1. Check out this repository locally.
1. Create a new package metadata file in `themes/default/data/registry/packages` suffixed with the major version number. E.g. `aws-v6.yaml`.
1. Copy the contents of the existing package metadata file (e.g. `aws.yaml`) into the new file
   - Add the suffix to the provider `name` and `title` fields.
   - Set the version in the `version` and `schema_file_url` fields to the latest patch version of the previous major version.
   - Ensure `featured` is set to `false`.
1. Make a copy of the logo file in `themes/default/static/registry/logos` with the new package name. E.g. `aws-v6.png`.
1. Copy the content directory e.g. `cp -r themes/default/content/registry/packages/aws themes/default/content/registry/packages/aws-v6`
1. Run `make .make/content/registry/packages/<package_name>/api-docs` to generate the API docs for the new package.
1. Run `make serve-all` to generate output files and preview them locally at [[http://localhost:1313/registry/]]
1. Rename `static/registry/packages/navs/aws.json` to `static/registry/packages/navs/aws-v6.json`.

Commit all changes which are ignored by default except the schema and the original metadata file as we won't be re-generating again:

```sh
git add -f content/registry/packages/$PACKAGE_NAME
git add -f static/registry/packages/navs/$PACKAGE_NAME.json
git add -f themes/default/content/registry/packages/$PACKAGE_NAME
git add -f themes/default/static/logos/pkg/$PACKAGE_NAME.*
```
