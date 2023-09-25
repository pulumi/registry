const fs = require("fs");
const registry = require("./registry");
const settings = require("./settings");
const algoliasearch = require("algoliasearch");
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
console.log(registryPackages);

// Configuration values required for updating the Algolia index.
const config = {
    appID: process.env.ALGOLIA_APP_ID,
    searchAPIKey: process.env.ALGOLIA_APP_SEARCH_KEY,
    adminAPIKey: process.env.ALGOLIA_APP_ADMIN_KEY,
    indexName: process.argv[2],
};

// if (!config.appID || !config.searchAPIKey || !config.adminAPIKey || !config.indexName) {
//     throw new Error(`Missing one or more required configuration values. (Provided keys: [${Object.keys(config)}])`);
// }

// Initialize the Algolia search client.
const client = algoliasearch(config.appID, config.adminAPIKey);
const algoliaIndex = client.initIndex(config.indexName);

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

// Gather up index settings, synonyms, and rules.
const indexSettings = {
    searchableAttributes: settings.getSearchableAttributes(),
    attributesForFaceting: settings.getAttributesForFaceting(),
    attributesToHighlight: settings.getAttributesToHighlight(),
    customRanking: settings.getCustomRanking(),
    ignorePlurals: true,
};
const indexSynonyms = settings.getSynonyms();
const indexRules = settings.getRules();

// Write the results, just so we have them.
console.log(" ↳ Writing results...");
fs.writeFileSync("./public/search-index.json", JSON.stringify(allObjects, null, 4));
fs.writeFileSync("./public/search-index-settings.json", JSON.stringify({ indexSettings, indexSynonyms, indexRules }, null, 4));
console.log(" ↳ Done. ✨\n");

// Update the Algolia index, including all page objects and index settings (like searchable
// attributes, custom ranking, synonyms, etc.).
// async function updateIndex(objects) {
//     console.log("Updating search index...");

//     try {
//         console.log(` ↳ Replacing all records in the '${ config.indexName }' index...`);
//         const result = await algoliaIndex.partialUpdateObjects(objects, {
//                 createIfNotExists: true,
//             });
//         console.log(`   ↳ ${result.objectIDs.length} records updated.`);

//         console.log(` ↳ Updating index settings...`)
//         await algoliaIndex.setSettings(indexSettings);

//         console.log(" ↳ Updating synonyms...")
//         await algoliaIndex.saveSynonyms(indexSynonyms, { replaceExistingSynonyms: true });

//         console.log(" ↳ Updating rules...")
//         await algoliaIndex.replaceAllRules(indexRules);

//         console.log(" ↳ Done. ✨\n");
//     }
//     catch (error) {
//         console.error(error);
//     }
// }

// updateIndex(allObjects);
