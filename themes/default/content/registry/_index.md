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
    - clisitemap
cascade:
    # API docs are generated in bulk during CI and don't need markdown or
    # clisitemap outputs.  Restricting them to HTML avoids a massive build-time
    # multiplier (tens of thousands of extra pages).
    - outputs:
          - HTML
      _target:
          path: "{**/api-docs,**/api-docs/**}"
    - outputs:
          - HTML
          - markdown
          - clisitemap
      _target:
          kind: section
    - outputs:
          - HTML
          - markdown
      _target:
          kind: page
---
