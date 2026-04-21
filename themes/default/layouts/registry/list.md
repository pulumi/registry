---
title: {{ .Title }}
{{ with .Description }}description: {{ . }}
{{ end }}url: {{ .RelPermalink }}
---
{{ if .RawContent -}}
{{ $content := .RenderShortcodes -}}
{{ $content = partial "registry/markdown-pipeline.md" $content -}}
{{ $content }}
{{- else }}
{{ range .Pages -}}
- [{{ .Title }}]({{ .RelPermalink }})
{{ end -}}
{{- end }}
