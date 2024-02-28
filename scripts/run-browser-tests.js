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
    const jsonString = out.substring(startIndex, endIndex+1);
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
    // Configure AWS SDK with assumed role credentials and session token
// AWS.config.update({
//     accessKeyId: 'YOUR_ACCESS_KEY_ID',
//     secretAccessKey: 'YOUR_SECRET_ACCESS_KEY',
//     sessionToken: 'YOUR_SESSION_TOKEN' // Include the session token here
//   });
  
  // Create an S3 service object
  const s3 = new AWS.S3();
  
  // Define the object to be converted to JSON
 
  
  // Convert object to JSON
  const jsonData = JSON.stringify(obj);
  
  // Write JSON data to a file
  const filename = 'data.json';
  fs.writeFileSync(filename, jsonData);
  
  // Upload JSON file to S3 bucket
  const bucketName = 'pulumi-testing-e2e-test-results';
  const key = 'data.json'; // Key under which the file will be saved in the bucket
  
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

