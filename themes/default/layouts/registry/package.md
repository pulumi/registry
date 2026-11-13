---
title: {{ .Title }}
{{ with .Description }}description: {{ . }}
{{ end }}url: {{ .RelPermalink }}
---
{{- if strings.HasSuffix .Permalink "how-to-guides/" -}}
{{- /* Generate a listing of child guide pages, mirroring the HTML layout */}}
{{ range .Pages }}
- [{{ with .Params.h1 }}{{ . }}{{ else }}{{ .Title }}{{ end }}]({{ .RelPermalink }})
{{- else }}
There are no guides yet for this package.
{{- end }}
{{- else -}}
{{- $content := .RenderShortcodes -}}
{{- $content = partial "registry/markdown-pipeline.md" $content -}}

{{ $content }}
{{- end -}}
