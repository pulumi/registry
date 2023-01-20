module github.com/pulumi/registry

go 1.16

require (
	github.com/pulumi/pulumi-hugo/themes/default v0.0.0-20230120004318-a072de0e703f // indirect
	github.com/pulumi/registry/themes/default v0.0.0-20221220231854-958c0296acf4 // indirect
)

// The override is needed because this repo is currently private and module at themes/default
// will be considered a private Go module as well. We could configure an SSH key to get around
// that but this is simpler for the time being.
replace github.com/pulumi/registry/themes/default => ./themes/default
