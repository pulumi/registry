module github.com/pulumi/registry

go 1.16

require (
	github.com/pulumi/pulumi-hugo/themes/default v0.0.0-20210827165408-f0c2264c9cf4 // indirect
	github.com/pulumi/registry/themes/default v0.0.0-20210827215118-ba028ca54298 // indirect
	github.com/pulumi/theme v0.0.0-20210926220603-a8786cb11ce2 // indirect
)

// The override is needed because this repo is currently private and module at themes/default
// will be considered a private Go module as well. We could configure an SSH key to get around
// that but this is simpler for the time being.
replace github.com/pulumi/registry/themes/default => ./themes/default
