module github.com/pulumi/registry

go 1.16

require (
	github.com/pulumi/pulumi-hugo/themes/default v0.0.0-20211103173548-5bdbe809c432 // indirect
	github.com/pulumi/registry/themes/default v0.0.0-20211011171710-45eb4e243ab7 // indirect
	github.com/pulumi/theme v0.0.0-20211101174744-08cd7b7eb0cc // indirect
)

// The override is needed because this repo is currently private and module at themes/default
// will be considered a private Go module as well. We could configure an SSH key to get around
// that but this is simpler for the time being.
replace github.com/pulumi/registry/themes/default => ./themes/default
