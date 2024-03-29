---
title: "API Gateway V1 to EventBridge | TypeScript"
h1: "API Gateway V1 to EventBridge"
linktitle: "API Gateway V1 to EventBridge"
meta_desc: "API Gateway V1 to EventBridge How-to Guide using TypeScript"
no_edit_this_page: true
cloud: aws
language: ts
layout: package
---

<!-- WARNING: this page was generated by a tool. Do not edit it by hand. -->
<!-- To change it, please see https://github.com/pulumi/docs/tree/master/tools/mktutorial. -->

<p class="mb-4 flex">
    <a class="flex flex-wrap items-center rounded-md font-display text-lg text-white bg-blue-600 border-2 border-blue-600 px-2 mr-2 whitespace-no-wrap hover:text-white" style="height: 45px;" href="https://github.com/pulumi/examples/tree/master/aws-ts-apigateway-eventbridge" target="_blank">
        <span><i class="fab fa-github pr-2"></i> View Code</span>
    </a>
</p>


[![Deploy with Pulumi](https://get.pulumi.com/new/button.svg)](https://app.pulumi.com/new?template=https://github.com/pulumi/examples/tree/master/aws-ts-apigateway-eventbridge)

This example demonstrates an API Gateway V1 integration with EventBridge and Lambda that also validates request bodies (using an API Gateway model) and returns a custom HTTP response.

## Prerequisites

1. [Install Pulumi](https://www.pulumi.com/docs/get-started/install/).
1. [Install Node.js](https://www.pulumi.com/docs/intro/languages/javascript/).
1. Configure your [AWS credentials](https://www.pulumi.com/docs/intro/cloud-providers/aws/setup/).

### Deploying the App

1. Clone this repo, change to this directory, then create a new [stack](https://www.pulumi.com/docs/intro/concepts/stack/) for the project:

    ```bash
    pulumi stack init
    ```

1. Specify an AWS region to deploy into:

    ```bash
    pulumi config set aws:region us-west-2
    ```

1. Install Node dependencies and run Pulumi:

    ```bash
    npm install
    pulumi up
    ```

1. In a few moments, the API Gateway instance service will be up and running and its public URL emitted as a Pulumi [stack output](https://www.pulumi.com/docs/intro/concepts/stack/#outputs).

    ```bash
    ...
    Outputs:
        url: "https://andchh8hg8.execute-api.us-west-2.amazonaws.com/dev"
    ```

1. Verify the deployment with `curl`:

    With invalid POST data:

    ```bash
    curl --data '{"some-invalid-property-name": "Chris"}' --header "Content-Type: application/json" "$(pulumi stack output url)/uploads"

    HTTP/2 400
    {"message": "Invalid request body"}
    ```

    With valid POST data:

    ```bash
    curl --data '{"name": "Chris"}' --header "Content-Type: application/json" "$(pulumi stack output url)/uploads"

    HTTP/2 201
    {"accepted":true}
    ```

1. Verify the Lambda was invoked with `pulumi logs`:

    ```bash
    pulumi logs --follow

    Collecting logs for stack dev since 2022-01-06T16:18:48.000-08:00.
    ...

    {
        source: 'my-event-source',
        detail: { 'name': 'Chris' }
    }
    ```

1. When you're ready, destroy your stack and remove it:

    ```bash
    pulumi destroy --yes
    pulumi stack rm --yes
    ```

