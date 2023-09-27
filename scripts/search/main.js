const fs = require("fs");
const registry = require("./registry");
const path = require('path');
const yaml = require('js-yaml');

// As part of the Registry build, we also generate a collection of JSON files that we use to power
// the API docs navigation. We use this file here as well to add resource listings, etc., to the
// index to make them more findable by their type names and the like.
const pathToRegistryPackagesJSON = "./public/registry/packages/navs";

const directoryPath = 'themes/default/data/registry/packages'; // Update with your directory path

function listFilesInDirectory(directoryPath) {
  try {
    const files = fs.readdirSync(directoryPath);

    return files
      .filter((filename) => filename.endsWith('.yaml') || filename.endsWith('.yml'))
      .map((filename) => path.join(directoryPath, filename));
  } catch (error) {
    console.error(`Error reading directory: ${error}`);
    return [];
  }
}

function parseYamlFile(filePath) {
  try {
    const fileContent = fs.readFileSync(filePath, 'utf-8');
    const yamlObject = yaml.load(fileContent);
    return yamlObject;
  } catch (error) {
    console.error(`Error parsing YAML file ${filePath}: ${error}`);
    return null;
  }
}

function parseYamlFilesInDirectory(directoryPath) {
  const yamlFiles = listFilesInDirectory(directoryPath);
  const parsedObjects = [];

  for (const filePath of yamlFiles) {
    const parsedObject = parseYamlFile(filePath);
    if (parsedObject) {
      parsedObjects.push(parsedObject);
    }
  }

  return parsedObjects;
}

const registryPackages = parseYamlFilesInDirectory(directoryPath);

// Generate a list of all Registry items -- modules, resources, functions, etc.
console.log(" ↳ Building Registry resource objects...");
const registryObjects = registry.getObjects(pathToRegistryPackagesJSON, registryPackages);

// Stitch these lists together into one tidy bundle.
let allObjects = [
    ...registryObjects,
];

// Temporary hack: Remove any references to `azure-native-v1`. This line can be
// removed once the azure-native-v1 package is removed from the Registry.
// https://github.com/pulumi/registry/issues/2879
allObjects = allObjects.filter(o => !o.href.includes("azure-native-v1"));

// Write the results, just so we have them.
console.log(" ↳ Writing results...");
fs.writeFileSync("./search-index.json", JSON.stringify(allObjects, null, 4));
console.log(" ↳ Done. ✨\n");
