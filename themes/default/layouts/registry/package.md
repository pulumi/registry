---
title: {{ .Title }}
{{ with .Description }}description: {{ . }}
{{ end }}url: {{ .RelPermalink }}
---
{{- if strings.HasSuffix .Permalink "how-to-guides/" -}}
{{- /* Generate a listing of child guide pages, mirroring the HTML layout */ -}}
{{- if eq (len .Pages) 0 }}
There are no guides yet for this package.
{{- else -}}
{{- $featured := where .Pages "Params.language" "==" nil -}}
{{- if $featured }}

## Featured Guides
{{ range $featured }}
- [{{ with .Params.h1 }}{{ . }}{{ else }}{{ .Title }}{{ end }}]({{ .RelPermalink }})
{{- end }}

## How-to Guides by Language
{{- end }}

<!-- chooser: language -->
{{- $langMap := dict "ts" "typescript" "py" "python" "go" "go" "cs" "csharp" "java" "java" "fs" "fsharp" -}}
{{- range $short, $tag := $langMap -}}
{{- $guides := where $.Pages "Params.language" $short -}}
{{- if $guides }}

<!-- option: {{ $tag }} -->
{{ range $guides }}
- [{{ with .Params.h1 }}{{ . }}{{ else }}{{ .Title }}{{ end }}]({{ .RelPermalink }})
{{- end }}

<!-- /option -->
{{- end -}}
{{- end }}

<!-- /chooser -->
{{- end -}}
{{- else -}}
{{- $content := .RenderShortcodes -}}
{{- $content = partial "registry/markdown-pipeline.md" $content -}}

{{ $content }}
{{- end -}}
