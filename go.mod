module github.com/pulumi/registry

go 1.16

require (
	github.com/pulumi/pulumi-hugo/themes/default v0.0.0-20210831201829-0ed9e751a6cb // indirect
	github.com/pulumi/registry/themes/default v0.0.0-20210831000804-ccb15fd93a15 // indirect
	github.com/pulumi/theme v0.0.0-20210901221031-a30a476fda89
)

// The override is needed because this repo is currently private and module at themes/default
// will be considered a private Go module as well. We could configure an SSH key to get around
// that but this is simpler for the time being.
replace github.com/pulumi/registry/themes/default => ./themes/default

// replace github.com/pulumi/theme => ../theme
