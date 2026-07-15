## Add a community package

> **Dynamically bridged Terraform provider?** If you consume it with
> `pulumi package add terraform-provider <name>` and it has no provider repo or committed
> `schema.json`, this flow does not apply. Those are listed through a separate Pulumi pipeline,
> so open a [New Package issue](https://github.com/pulumi/registry/issues/new?template=new-package.yml)
> to request one instead of this PR.

This PR should add exactly one entry to `community-packages/package-list.json` and change
nothing else. That one entry is the whole registration.

```json
{
  "repoSlug": "<owner>/<repo>",
  "schemaFile": "provider/cmd/pulumi-resource-<name>/schema.json"
}
```

### What happens next

- Automated checks run on this PR and post a fact-sheet comment: they pin your latest release,
  install the SDKs you advertise, and validate your docs. They hold no secrets, so they run on
  your fork.
- If something is red, fix it in your provider repo (cut a release, publish an SDK, correct the
  schema path), then comment `/check` to re-run. The checks read your live upstream, not this
  diff, so you do not push a new commit here to re-validate.
- A Pulumi maintainer reads the fact-sheet and approves. Nothing merges automatically. After
  merge, your docs are generated and published for you, so you never commit generated files here.

### Before you open this

- [ ] Your provider has a `v`-prefixed [Semver 2.0](https://semver.org) release (`vX.Y.Z`).
- [ ] The `schemaFile` path resolves in your repo at that release.
- [ ] Your repo has `docs/_index.md` (a summary plus a code example per language) and
      `docs/installation-configuration.md` (SDK links plus configuration).
- [ ] The SDKs you advertise are actually published.

Not sure about one of these? Open the PR anyway; the fact-sheet will tell you exactly what is
missing.

### Your connection to the provider

One line on how you're connected (maintainer, contributor, user). Your GitHub handle is enough
for an individually maintained provider. If the provider is maintained by a company, please also
complete the [Pulumi Registry Onboarding form](https://docs.google.com/forms/d/e/1FAIpQLSe8ohRL1PP00KTXEGPYawjY-PgrbxCH6Z88wBqxEDMa4hw2qw/viewform).
