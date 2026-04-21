---
title: Pulumi Registry
notitle: true
linktitle: Pulumi Registry
meta_desc: The Pulumi Registry hosts Providers that give direct access to all of a cloud provider’s resources and Components for common cloud architectures.
no_on_this_page: true
menu:
  clouds:
    name: All clouds
    weight: 5
aliases:
    - /registry/packages
    - /docs/tutorials
    - /docs/reference/tutorials/
    - /docs/clouds/pkg/
outputs:
    - HTML
    - markdown
    - llmsitemap
cascade:
    # API docs are generated in bulk during CI and don't need markdown or
    # llmsitemap outputs.  Restricting them to HTML avoids a massive build-time
    # multiplier (tens of thousands of extra pages).
    - outputs:
          - HTML
      _target:
          path: "{**/api-docs,**/api-docs/**}"
    - outputs:
          - HTML
          - markdown
          - llmsitemap
      _target:
          kind: section
    - outputs:
          - HTML
          - markdown
      _target:
          kind: page
---

The Pulumi Registry is the public index of Pulumi packages. It lists first-party Pulumi providers, bridged Terraform providers, and community components for common cloud architectures.

For the full list of packages, see [/registry/packages.md](/registry/packages.md).

Popular providers:

- [AWS](/registry/packages/aws.md)
- [Azure Native](/registry/packages/azure-native.md)
- [Google Cloud](/registry/packages/gcp.md)
- [Kubernetes](/registry/packages/kubernetes.md)
- [Docker](/registry/packages/docker.md)
- [Random](/registry/packages/random.md)

Each package page has hand-authored overview and installation/configuration content available as markdown (append `.md` to the URL, or send `Accept: text/markdown`). API reference documentation for resources and functions is available at `/api-docs/` on each package page in HTML only.
