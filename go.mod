module github.com/pulumi/registry

go 1.16

require (
	github.com/pulumi/pulumi-hugo/themes/default v0.0.0-20210920175349-6ece72b1e55a // indirect
	github.com/pulumi/registry/themes/default v0.0.0-20210916205127-04dd45dfff38 // indirect
	github.com/pulumi/theme v0.0.0-20210924215439-9cff11323c57 // indirect
)

// The override is needed because this repo is currently private and module at themes/default
// will be considered a private Go module as well. We could configure an SSH key to get around
// that but this is simpler for the time being.
replace github.com/pulumi/registry/themes/default => ./themes/default
