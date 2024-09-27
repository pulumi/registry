#!/bin/bash
set -e
TUTORIAL_OUT=$(mktemp -d)
pushd ./tools/mktutorial
go run *.go https://github.com/pulumi/examples $TUTORIAL_OUT
popd
for cloud in "aws-apigateway" "aws" "aws-native" "classic-azure" "azure" "gcp" "kubernetes" ; do
    ## Temporary hack to address mismatched package names in pulumi/examples
    if [[ "$cloud" == "azure" ]]
    then
        dest="azure-native"
    elif [[ "$cloud" == "classic-azure" ]]
    then
        dest="azure"
    else
        dest=$cloud
    fi

    mkdir -p ./themes/default/content/registry/packages/$dest
    mkdir -p ./themes/default/content/registry/packages/$dest/how-to-guides

    # Check if the directory exists before copying.
    # (meaning there are tutorials for that cloud provider, if not the dir would not get created)
    if [ -d "$TUTORIAL_OUT/tutorials/$cloud" ]; then
        cp $TUTORIAL_OUT/tutorials/$cloud/* ./themes/default/content/registry/packages/$dest/how-to-guides/
    fi
done
rm -rf $TUTORIAL_OUT
