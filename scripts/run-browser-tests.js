// e2e-run-tests.js
const cypress = require('cypress')
const cp = require('child_process');
const { WebClient, LogLevel } = require('@slack/web-api');

const apispec = "api-docs.cy.js"
const sitespec = "site.cy.js"

// Extract command-line arguments
const args = process.argv.slice(2); // Remove first two elements (node and script path)

// Assuming you're passing two arguments: arg1 and arg2
const baseUrl = args[0];
const spec = args[1];

console.log("Argument 1:", baseUrl);
console.log("Argument 2:", spec);

cp.exec(`CYPRESS_BASE_URL="${baseUrl}" yarn run cypress run --headless --reporter json --spec "./cypress/e2e/${spec}"`, processResults)

function processResults(error, stdout, stderr) {
    const out = stdout;
    const startIndex = out.indexOf("{");
    const endIndex = out.lastIndexOf("}");
    const jsonString = out.substring(startIndex, endIndex+1);
    const results = JSON.parse(jsonString);
    // console.log("start log", jsonString, "end log");
    console.log(results);
    // results.tests.forEach(t => {
    //     if (t.err) {
    //         console.log(JSON.stringify(t.err));
    //     }
    // });
}


