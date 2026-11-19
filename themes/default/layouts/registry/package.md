---
title: {{ .Title }}
{{ with .Description }}description: {{ . }}
{{ end }}url: {{ .RelPermalink }}
---
{{- $segments := split .File.Path "/" -}}
{{- if and (eq (len $segments) 5) (eq (index $segments 3) "guides") (eq (index $segments 4) "_index.md") -}}
{{- /* Generate a listing of child guide pages, mirroring the HTML layout */}}
{{ .RenderShortcodes }}
{{ range .Pages }}
- [{{ with .Params.h1 }}{{ . }}{{ else }}{{ .Title }}{{ end }}]({{ .RelPermalink }}){{ with .Description }}: {{ . }}{{ end }}
{{- else }}
There are no guides yet for this package.
{{- end }}
{{- else -}}
{{- $content := .RenderShortcodes -}}
{{- $content = partial "registry/markdown-pipeline.md" $content -}}

{{ $content }}
{{- end -}}
