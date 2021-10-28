module github.com/pulumi/registry

go 1.16

require (
	github.com/pulumi/pulumi-hugo/themes/default v0.0.0-20211027195409-079e74aeecd0 // indirect
    github.com/pulumi/registry/themes/api-docs v0.0.0-20211011171710-45eb4e243ab7 // indirect
	github.com/pulumi/registry/themes/default v0.0.0-20211011171710-45eb4e243ab7 // indirect
    github.com/pulumi/theme v0.0.0-20211025172841-d7092250170b // indirect
)

replace github.com/pulumi/registry/themes/default => ./themes/default

replace github.com/pulumi/registry/themes/api-docs => ./themes/api-docs
