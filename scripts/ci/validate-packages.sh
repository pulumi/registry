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

    # check to see that package _index.md file exists and is using the correct layout.
    overview_layout=$(cat "themes/default/content/registry/packages/$pkg/_index.md" | yq -f eval '.layout')
    if [[ "$overview_layout" != "package" ]]; then
        echo "ERROR: invalid layout, '$overview_layout' found for package, $pkg. Must use the 'package' layout."
        exit 1
    fi

    # check to see that package  install and configuration file exists and is using the correct layout.
    install_layout=$(cat "themes/default/content/registry/packages/$pkg/installation-configuration.md" | yq -f eval '.layout')
    if [[ "$install_layout" != "package" ]]; then
        echo "ERROR: invalid layout, '$install_layout' found for package, $pkg. Must use the 'package' layout."
        exit 1
    fi
done; 

echo "packages validated successfully!!"


