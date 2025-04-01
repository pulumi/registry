const fs = require("fs");
const registry = require("./registry");
const page = require("./page");

// As part of the Registry build, we also generate a collection of JSON files that we use to power
// the API docs navigation. We use this file here as well to add resource listings, etc., to the
// index to make them more findable by their type names and the like.
const pathToRegistryPackagesJSON = "./public/registry/packages/navs";

// As part of the Hugo build, we generate a JSON file at /index.json containing a record for every
// page of the website. We use this file to generate primary search records for Algolia.
const pathToFullSiteJSON = "./public/index.json";
const hugoPageItems = JSON.parse(
    fs.readFileSync(pathToFullSiteJSON, "utf-8").toString(),
);

// Generate a list of primary page objects. This list contains one record for every page of the site.
console.log("\nBuilding search index...");
console.log(" ↳ Building primary page objects...");
const primaryPageObjects = page.getPrimaryObjects(hugoPageItems);

// Generate a list of all Registry items -- modules, resources, functions, etc.
console.log(" ↳ Building Registry resource objects...");
const registryObjects = registry.getObjects(
    pathToRegistryPackagesJSON,
    hugoPageItems,
);

// Filter out only non-registry objects from primaryPageObjects and de-dupe registryObjects.
const filteredPageObjects = primaryPageObjects
    .filter((o) => o.href.startsWith("/registry"))
    .filter(
        (o) => registryObjects.find((ro) => ro.href === o.href) === undefined,
    );

// Stitch these lists together into one tidy bundle.
let allObjects = [...filteredPageObjects, ...registryObjects];

// Temporary hack: Remove any references to `azure-native-v2`. This line can be
// removed once the azure-native-v2 package is removed from the Registry.
// https://github.com/pulumi/registry/issues/2879
// https://github.com/pulumi/pulumi-azure-native/issues/3420
allObjects = allObjects.filter((o) => !o.href.includes("azure-native-v2"));

// Write the results, just so we have them.
console.log(" ↳ Writing results...");
fs.writeFileSync("./search-index.json", JSON.stringify(allObjects, null, 4));
console.log(" ↳ Done. ✨\n");
