{{- $providerName := .Get 0 -}}
{{- $packagePath := printf "/docs/reference/pkg/%s" $providerName -}}
{{- if fileExists (printf "%s/_index.md" $packagePath) }}

> **Note:** This page documents the language specification for the {{ $providerName }} package. If you're looking for help working with the inputs, outputs, or functions of {{ $providerName }} resources in a Pulumi program, please see the [resource documentation]({{ $packagePath }}) for examples and API reference.
{{- end }}
