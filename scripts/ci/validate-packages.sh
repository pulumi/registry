#! /bin/bash

set -e

# capture all existing packages in the content dir.
pkg_dirs=$(ls -l "themes/default/content/registry/packages/" | tail -n +2 | awk '{print $9}')

# validates that all packages also have related content directory present.
ls -l "themes/default/data/registry/packages" | tail -n +2 | awk '{print $9}' |  while read line; do
    # extract the name from the frontmatter
    pkg=$(cat "themes/default/data/registry/packages/${line}" | yq ".name")
    # check to see that a dir exists for the package
    if [[ "$pkg_dirs" != *"$pkg"* ]]; then
        echo "ERROR: no content files found for package, $pkg"
        exit 1
    fi

    # check to see that package _index.md file exists.
    overview_path="themes/default/content/registry/packages/$pkg/_index.md"
    if [ ! -f $overview_path ]; then
        echo "ERROR: '$overview_path' does not exist."
        exit 1
    fi

    install_path="themes/default/content/registry/packages/$pkg/installation-configuration.md"
    if [ ! -f $install_path ]; then
        echo "WARN: '$install_path' does not exist. Package will be displayed with a single overview file."
    fi

done;

echo "packages validated successfully!!"
