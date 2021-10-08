module github.com/pulumi/registry

go 1.16

require (
	github.com/pulumi/pulumi-hugo/themes/default v0.0.0-20211006183715-8bfe690c3b3f // indirect
	github.com/pulumi/registry/themes/default v0.0.0-20210927233147-6d9e62ba8825 // indirect
	github.com/pulumi/theme v0.0.0-20211008234130-0ca1a3506739 // indirect
)

// The override is needed because this repo is currently private and module at themes/default
// will be considered a private Go module as well. We could configure an SSH key to get around
// that but this is simpler for the time being.
replace github.com/pulumi/registry/themes/default => ./themes/default

replace github.com/pulumi/theme => ../theme
