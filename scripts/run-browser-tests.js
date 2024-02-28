// e2e-run-tests.js
const cypress = require('cypress')
const cp = require('child_process');
const { WebClient, LogLevel } = require('@slack/web-api');
const fs = require('fs');
const AWS = require('aws-sdk');

const specs = ["api-docs.cy.js", "site.cy.js"]

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
    const jsonString = out.substring(startIndex, endIndex + 1);
    const results = JSON.parse(jsonString);
    // console.log("start log", jsonString, "end log");
    console.log(results);
    pushS3(results);
    // results.tests.forEach(t => {
    //     if (t.err) {
    //         console.log(JSON.stringify(t.err));
    //     }
    // });
}


function pushS3(obj) {
    console.log("pushing to S3")
    const s3 = new AWS.S3();

    const jsonData = JSON.stringify(obj);

    const dateString = obj.stats.start;
    const date = new Date(dateString);

    const year = date.getFullYear();
    const month = date.getMonth() + 1; // Month is zero-based, so we add 1
    const day = date.getDate();

    // Write JSON data to file
    const filename = 'results.json';
    fs.writeFileSync(filename, jsonData);

    // Upload JSON file to S3 bucket
    const bucketName = 'pulumi-testing-e2e-test-results';
    const key = `${year}/${month}/${day}/results.json`; // Key under which the file will be saved in the bucket

    const uploadParams = {
        Bucket: bucketName,
        Key: key,
        Body: fs.createReadStream(filename)
    };

    s3.upload(uploadParams, (err, data) => {
        if (err) {
            console.error("Error uploading file:", err);
        } else {
            console.log("File uploaded successfully. Location:", data.Location);
        }
    });

}

