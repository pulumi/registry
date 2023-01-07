module github.com/pulumi/registry

go 1.16

require (
	github.com/pulumi/pulumi-hugo/themes/default v0.0.0-20230106194421-9b1f96367754 // indirect
	github.com/pulumi/registry/themes/default v0.0.0-20221220231854-958c0296acf4 // indirect
	github.com/pulumi/theme v0.0.0-20230104232221-88239bb01cba
)

// The override is needed because this repo is currently private and module at themes/default
// will be considered a private Go module as well. We could configure an SSH key to get around
// that but this is simpler for the time being.
replace github.com/pulumi/registry/themes/default => ./themes/default
