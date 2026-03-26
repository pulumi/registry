---
title: {{ .Title }}
{{ with .Description }}description: {{ . }}
{{ end }}url: {{ .RelPermalink }}
---
{{- /* For list pages (sections), output a simple listing of child pages */ -}}
{{- range .Pages -}}
- [{{ .Title }}]({{ .RelPermalink }})
{{ end -}}
