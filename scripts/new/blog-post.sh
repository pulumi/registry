#!/bin/bash

# By default, Hugo will attempt to create the new post in a directory relative to the location
# of the archetype, which we fetch from pulumi-hugo and vendor at ./_vendor.
# We want it created in _this_ repo, though, so we pass in an explicit path.
hugo new --kind blog-post "../../../../../../../themes/default/content/blog/${1}"
