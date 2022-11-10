module github.com/pulumi/registry

go 1.16

require (
	github.com/pulumi/pulumi-hugo v0.0.0-20221213210121-e06208c33c16 // indirect
	github.com/pulumi/pulumi-hugo/themes/default v0.0.0-20221213121308-7111f19e2275 // indirect
	github.com/pulumi/theme v0.0.0-20221220211425-e19758b7aade
)

// The override is needed because this repo is currently private and module at themes/default
// will be considered a private Go module as well. We could configure an SSH key to get around
// that but this is simpler for the time being.
replace github.com/pulumi/registry/themes/default => ./themes/default
