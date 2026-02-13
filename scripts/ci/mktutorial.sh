#!/bin/bash
set -e

TUTORIAL_OUT=$(mktemp -d)
make bin/mktutorial
bin/mktutorial generate "$1" "$TUTORIAL_OUT"

for cloud in "aws-apigateway" "aws" "classic-azure" "azure" "gcp" "kubernetes" ; do
    ## Temporary hack to address mismatched package names in pulumi/examples
    if [[ "$cloud" == "azure" ]]; then
        dest="azure-native"
    elif [[ "$cloud" == "classic-azure" ]]; then
        dest="azure"
    else
        dest="$cloud"
    fi

    dest_dir="./themes/default/content/registry/packages/$dest/how-to-guides"
    src_dir="$TUTORIAL_OUT/tutorials/$cloud"

    mkdir -p "$dest_dir"

    # Copy generated tutorials if any exist for this cloud.
    if [ -d "$src_dir" ]; then
        cp "$src_dir/"* "$dest_dir/"
    fi

    # Clean up stale tutorials.
    bin/mktutorial cleanup "$src_dir" "$dest_dir"
done

# Also clean up stale tutorials in versioned package directories.
for pkg in "aws-v6" "azure-native-v2" ; do
    dest_dir="./themes/default/content/registry/packages/$pkg/how-to-guides"
    [ -d "$dest_dir" ] || continue

    # Determine which cloud this versioned package corresponds to.
    if [[ "$pkg" == "aws-v6" ]]; then
        src_dir="$TUTORIAL_OUT/tutorials/aws"
    elif [[ "$pkg" == "azure-native-v2" ]]; then
        src_dir="$TUTORIAL_OUT/tutorials/azure"
    fi

    bin/mktutorial cleanup "$src_dir" "$dest_dir"
done

rm -rf "$TUTORIAL_OUT"
