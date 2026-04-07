---
title: {{ .Title }}
{{ with .Description }}description: {{ . }}
{{ end }}url: {{ .RelPermalink }}
---
{{- if strings.HasSuffix .Permalink "how-to-guides/" -}}
{{- /* Generate a listing of child guide pages */ -}}
{{- range .Pages }}
- [{{ .Title }}]({{ .RelPermalink }})
{{- end }}
{{- else -}}
{{- $content := .RenderShortcodes -}}
{{- $content = partial "registry/markdown-pipeline.md" $content -}}

{{ $content }}
{{- end -}}
