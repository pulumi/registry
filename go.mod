module github.com/pulumi/registry

go 1.23

// The override is needed because this repo is currently private and module at themes/default
// will be considered a private Go module as well. We could configure an SSH key to get around
// that but this is simpler for the time being.
replace github.com/pulumi/registry/themes/default => ./themes/default

require github.com/pulumi/registry/themes/default v0.0.0-20260127153208-e3baf61d8776 // indirect
