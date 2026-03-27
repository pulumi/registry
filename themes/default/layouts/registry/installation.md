{{- $directories := (split .Page.File.Path "/") -}}
{{- $packageName := index $directories 2 -}}
{{- $packageHome := path.Join "registry/packages" $packageName -}}
---
title: {{ with .GetPage $packageHome }}{{ .Title }}{{ end }}: Installation & Configuration
url: {{ .RelPermalink }}
---
{{- $content := .RenderShortcodes -}}
{{- $content = partial "registry/markdown-pipeline.md" $content -}}

{{ $content }}
