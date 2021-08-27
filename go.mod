module github.com/pulumi/pulumi-hugo-internal

go 1.16

require (
	github.com/pulumi/pulumi-hugo-internal/themes/default v0.0.0-20210528133833-69e13c9b80fa // indirect
	github.com/pulumi/pulumi-hugo/themes/default v0.0.0-20210818221115-a11b214aa9de // indirect
)

replace github.com/pulumi/pulumi-hugo/themes/default => ../pulumi-hugo/themes/default

replace github.com/pulumi/pulumi-hugo-internal/themes/default => ./themes/default
