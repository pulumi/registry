# Registry publish drift

Findings from an audit of every `push.yml` run between 2026-03-10, when `publish_to_registry.py` replaced `push-registry.py` as the primary publish path, and 2026-08-06. 1,493 runs.

## tl;dr

The legacy fallback was still running, for two unrelated reasons. One is upstream placeholder schema versions, which is being fixed at the source. The other is structural and has no fix yet: when a production build fails, the package version it was carrying is dropped and never retried, because the primary publish path only looks at the last commit. The legacy path is the only thing that has been repairing that, and it repaired it by accident.

## The two paths

| | primary | legacy fallback |
|---|---|---|
| script | `scripts/ci/publish_to_registry.py` | `scripts/ci/push-registry.py` |
| tool | `registry-mirror-publish` | `pulumi package publish` |
| scope | packages changed in `git diff HEAD~1` | every package YAML, publishing whatever the API 404s on |
| schema version | uploaded as fetched | rewritten to match the registry YAML |
| runs when | always | only when the primary exits non-zero |

Both are steps in the same `build` job, after `Build and deploy`.

## How often the fallback ran

28 of 1,493 runs. It never failed, so the "both paths failed" Slack alert has never fired.

- 17 runs in one cluster on 2026-03-23 and 2026-03-24, closed by `b18d82305` (a missing `repo_root` argument). No recurrence.
- 11 runs since, all the same cause: `terraform-provider` and `powerplatform` ship placeholder schema versions.

## Cause 1: placeholder schema versions

The registry publishes a package's schema straight from its release tag. The service rejects a schema whose own `version` field disagrees with the version being published:

```text
400 Bad Request: Schema version (0.0.0-dev) must match query parameter version (1.3.0)
```

`terraform-provider` declares `0.0.0-dev` and `powerplatform` declares `0.0.0`. Both fail the primary path on every release and land only because the fallback rewrites the version on the way out. Three more packages are in the same state but have not released since March, so they have never exercised the current path: `aem`, `kafka-connect` and `genesiscloud`. The last is only a stray `v` prefix.

Most providers avoid this by omitting the field entirely, which lets the service take the version from the publish request.

## Cause 2: failed builds drop a package version permanently

This is the one worth acting on.

Both publish steps sit after `Build and deploy` in the same job. When that step fails, the job stops and both publish steps are skipped. The primary is scoped to `git diff HEAD~1`, so the next run only considers its own commit. Nothing revisits the package, and the version is simply never published.

Nine runs failed in the window. Seven of them were `Publish Package Metadata` commits, so seven package versions were dropped:

| run | date | package version dropped | why the build failed |
|---|---|---|---|
| 23263250725 | 2026-03-18 | `equinix@0.29.1` | `Build and deploy` |
| 24817511189 | 2026-04-23 | `fastly@12.0.1` | `Build and deploy` |
| 24817652115 | 2026-04-23 | `mongodbatlas@4.8.0` | `Build and deploy` |
| 24824043787 | 2026-04-23 | `ctfd@2.4.3` | `Build and deploy` |
| 25906915426 | 2026-05-15 | `equinix@0.30.0` | build job never started |
| 29441117266 | 2026-07-15 | `rootly@3.7.0` | `Install s5cmd` download failed |
| 30673917138 | 2026-07-31 | `selectel@8.2.4` | `llm-docs.json` fetch returned 403 during `Build and deploy` |

None of these is a publish bug. They are ordinary infrastructure flakes: a GitHub release download that failed, a CDN object that was not ready, a runner that never picked up the job. The publish never even attempted.

## The fallback has been repairing this by accident

The legacy path sweeps all package YAMLs and publishes anything missing, so whenever it ran it also backfilled unrelated packages that earlier runs had dropped:

| run | primary's scope that run | what the fallback also published |
|---|---|---|
| 26512197085 | `terraform-provider.yaml` | `mongodbatlas@4.8.0`, `equinix@0.30.0` |
| 29748320654 | `terraform-provider.yaml` | `rootly@3.7.0` |
| 30930672296 | `terraform-provider.yaml` | `selectel@8.2.4` |

Those four packages had nothing to do with the commits that triggered those runs. They were repaired only because `terraform-provider` happened to fail in the same run and drag the fallback in. A version dropped in a run where the fallback did not fire stayed dropped.

## Current state

Every package's current version resolves. Probing all 278 published versions on `master` against `api.pulumi.com` returns 200 for each; the only 404 was a commit whose build was still in flight.

Three historical versions never made it and are still 404:

- `fastly@12.0.1`
- `ctfd@2.4.3`
- `selectel@8.2.4`

They are invisible in normal use because each package has since released a newer version that did publish. They matter to anyone pinning an exact version.

`equinix@0.29.1` was dropped the same way but has since been repaired.

## What this means for retiring the fallback

Removing `push-registry.py` removes the only mechanism that repairs a dropped version. Nothing needs repairing today, so the removal is safe in the immediate sense, but the failure mode does not go away: it becomes permanent instead of eventually self-healing.

Two things would close it:

1. Move the publish steps out of the `build` job, or make them run when the build fails, so an infrastructure flake in the site build does not take the publish with it.
2. Add a scheduled reconciliation sweep. `publish_to_registry.py --validate-all` already enumerates every package and builds its spec; publishing whatever the API 404s on is a small addition, and it is what the legacy path was doing.

Either removes the dependence on a fallback firing by coincidence.
