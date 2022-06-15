---
title: Checkly
meta_desc: Tutorials for using infrastructure as code in Checkly Pulimi Provider.
layout: how-to
---

## Creating your first Checks with JavaScript

1. Open your terminal and run `$ pulumi new` to create a new Pulumi project, chose the `javascript` template and the target stack.
1. Install the Checkly Pulumi provider using npm: `$ npm i @checkly/pulumi`.
1. Create an `index.js` file in the root of your project and paste the following code:

    ```javascript
    const checkly = require("@checkly/pulumi")

    new checkly.Check("api-check", {
      type: "API",
      name: "Public SpaceX API",
      activated: true,
      frequency: 10,
      locations: ["us-east-1"],
      request: {
        method: "GET",
        url: "https://api.spacexdata.com/v3",
        assertions: [
          {
            source: 'STATUS_CODE',
            comparison: 'EQUALS',
            target: 200
          },
          {
            source: 'JSON_BODY',
            property: '$.project_name',
            comparison: 'EQUALS',
            target: 'SpaceX-API'
          }
        ]
      }
    })

    new checkly.Check("browser-check", {
      type: "BROWSER",
      name: "Google.com Playwright check",
      activated: true,
      frequency: 10,
      locations: ["us-east-1"],
      script: `const { chromium } = require('playwright')

    async function run () {
      const browser = await chromium.launch()
      const page = await browser.newPage()

      const response = await page.goto('https://google.com')

      if (response.status() > 399) {
        throw new Error('Failed with response code \${response.status()}')
      }

      await page.screenshot({ path: 'screenshot.jpg' })

      await page.close()
      await browser.close()
    }

    run()`
    })
    ```
1. Setup you Checkly API Key and Account id:
    ```bash
    $ pulumi config set checkly:apiKey cu_xxx --secret
    $ pulumi config set checkly:accountId xxx
    ```
1. You are ready to go, run `$ pulumi up` to deploy your stack 🚀

> Check the [examples directory](https://github.com/checkly/pulumi-checkly/tree/main/examples) for more detailed code samples.