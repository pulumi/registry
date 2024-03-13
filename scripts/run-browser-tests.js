const cp = require('child_process');
const fs = require('fs');
const AWS = require('aws-sdk');

// Extract command-line arguments
const args = process.argv.slice(2); // Remove first two elements (node and script path)

// Assuming you're passing two arguments: arg1 and arg2
const baseUrl = args[0];
const spec = args[1];

if (!(args[0] && args[1])) {
    console.error("baseURL and spec file are required as arguments")
    process.exit(1)
}

cp.exec(`CYPRESS_BASE_URL="${baseUrl}" yarn run cypress run --headless --reporter cypress-multi-reporters --spec "./cypress/e2e/${spec}"`, processResults)

async function processResults(error, stdout, stderr) {
    // log full run output to runner console.
    console.log(stdout);

    // EXTREME HACK: Extracts JSON out of console output.
    const startIndex = stdout.indexOf("{");
    const endIndex = stdout.lastIndexOf("}");
    const jsonString = stdout.substring(startIndex, endIndex + 1);
    const results = JSON.parse(jsonString);
    const transformed = transformResults(results)
    // Log JSON structure to the console.
    console.log(transformed);
    await pushS3(transformed);

    // The cli command error is trapped here since it is a sub process of this script.
    // Checking if error here will enable us to mark this as a failed run by exiting
    // unsuccessfully. Otherwise the run will always pass successfully regardless of
    // the result of the cypress tests.
    if (error) {
        console.error("errors:", stderr);
        process.exit(1)
    }
}

function transformResults(results) {
    const result = results.stats;
    result["ghRunURL"] = getRunUrl();

    // Get list of failed pages.
    const failedPages = [];
    // iterate over test failures and add page url to failedPages array.
    results.failures.forEach(f => {
        failedPages.push(extractPageUrl(f.fullTitle))
    })

    // dedupe pages and keep track of a count for each page failure occurance to map
    // the page to the number of failure occurances for each page.
    const pageMap = {}
    failedPages.forEach(page => {
        pageMap[page] = (pageMap[page] || 0) + 1;
    })
    
    // Convert pageMap to an array of objects with failure counts. This 
    result["failedPages"] = Object.keys(pageMap).map(key => ({
        page: key,
        failures: pageMap[key],
        tests: result.tests/50,
    }));

    result["failedPageCount"] = Object.keys(pageMap).length;
    return result;
}

function extractPageUrl(msg) {
    const urlRegex = /(https?:\/\/[^)]+)/;
    const match = msg.match(urlRegex);
    const url = match ? match[0] : null;

    return url;
}

function getRunUrl() {
    const serverUrl = process.env.GITHUB_SERVER_URL;
    const repo = process.env.GITHUB_REPOSITORY;
    const runId = process.env.GITHUB_RUN_ID;
    const runAttempt = process.env.GITHUB_RUN_ATTEMPT;

    return serverUrl + "/" + repo + "/actions/runs/" + runId + "/attempts/" + runAttempt;
}

// expand date string to return yyyy, mm, dd
function expandDate(dateString) {
    const date = new Date(dateString);
    const year = date.getFullYear();
    const monthNumber = date.getMonth() + 1 // zero indexed, so add 1 to get month number
    const month = `${ monthNumber < 10 ? `0${monthNumber}` : monthNumber }`;
    const dayNumber = date.getDate();
    const day = `${ dayNumber < 10 ? `0${dayNumber}` : dayNumber }`;
    return {year, month, day};
}

// upload to S3 with yyyy/mm/dd prefix.
async function pushS3(obj) {
    const bucketName = 'pulumi-api-docs-e2e-test-results-prodution';
    
    const jsonData = JSON.stringify(obj);
    // extract yyyy, mm, dd from date.
    const {year, month, day} = expandDate(obj.start);
    
    // Write JSON to file
    const filename = 'results.json';
    fs.writeFileSync(filename, jsonData);
    
    // Upload JSON file to S3 bucket
    console.log("pushing to S3")
    const s3 = new AWS.S3();
    const key = `${year}/${month}/${day}/results.json`; 

    const uploadParams = {
        Bucket: bucketName,
        Key: key,
        Body: fs.createReadStream(filename)
    };

    return s3.upload(uploadParams).promise();
}

