// e2e-run-tests.js
const cypress = require('cypress')
const cp = require('child_process');
const { WebClient, LogLevel } = require('@slack/web-api');
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

console.log("Argument 1:", baseUrl);
console.log("Argument 2:", spec);

cp.exec(`CYPRESS_BASE_URL="${baseUrl}" yarn run cypress run --headless --reporter json --spec "./cypress/e2e/${spec}"`, processResults)

function processResults(error, stdout, stderr) {
    const out = stdout;
    const startIndex = out.indexOf("{");
    const endIndex = out.lastIndexOf("}");
    const jsonString = out.substring(startIndex, endIndex + 1);
    const results = JSON.parse(jsonString);
    console.log(JSON.stringify(results, null, 2));
    const transformed = transformResults(results)
    console.log(transformed);
    pushS3(transformed);
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
        tests: result.tests/1
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
    // const workflow = process.env.GITHUB_WORKFLOW

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
function pushS3(obj) {
    console.log("pushing to S3")
    const s3 = new AWS.S3();

    const jsonData = JSON.stringify(obj);

    // extract yyyy, mm, dd from date.
    const {year, month, day} = expandDate(obj.start);

    // Write JSON to file
    const filename = 'results.json';
    fs.writeFileSync(filename, jsonData);

    // Upload JSON file to S3 bucket
    const bucketName = 'pulumi-testing-e2e-test-results';
    const key = `${year}/${month}/${day}/results.json`; 

    const uploadParams = {
        Bucket: bucketName,
        Key: key,
        Body: fs.createReadStream(filename)
    };

    s3.upload(uploadParams, (err, data) => {
        if (err) {
            console.error("Error uploading file to s3:", err);
        } else {
            console.log("File uploaded successfully:", data.Location);
        }
    });
}

