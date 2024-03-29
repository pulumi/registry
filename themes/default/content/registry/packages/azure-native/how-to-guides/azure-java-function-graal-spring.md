---
title: "Azure Function with Spring Boot Native and GraalVM | Java"
h1: "Azure Function with Spring Boot Native and GraalVM"
linktitle: "Azure Function with Spring Boot Native and GraalVM"
meta_desc: "Azure Function with Spring Boot Native and GraalVM How-to Guide using Java"
no_edit_this_page: true
cloud: azure
language: java
layout: package
---

<!-- WARNING: this page was generated by a tool. Do not edit it by hand. -->
<!-- To change it, please see https://github.com/pulumi/docs/tree/master/tools/mktutorial. -->

<p class="mb-4 flex">
    <a class="flex flex-wrap items-center rounded-md font-display text-lg text-white bg-blue-600 border-2 border-blue-600 px-2 mr-2 whitespace-no-wrap hover:text-white" style="height: 45px;" href="https://github.com/pulumi/examples/tree/master/azure-java-function-graal-spring" target="_blank">
        <span><i class="fab fa-github pr-2"></i> View Code</span>
    </a>
</p>


Starting point for building Spring Native application hosted in Azure Function.

Inspired by [Julien Dubois](https://github.com/jdubois/azure-native-spring-function) 
and [Spring Native - Cloud Function Netty example](https://github.com/spring-projects-experimental/spring-native/tree/main/samples/cloud-function-netty).

Azure Functions custom handlers are used to run the GraalVM binary.

## Running the App in Azure

1.  Create a new stack:

    ```
    $ pulumi stack init dev
    ```

1.  Login to Azure CLI (you will be prompted to do this during deployment if you forget this step):

    ```
    $ az login
    ```

1. Set the Azure region location:

    ```
    $ pulumi config set azure-native:location westus
    ```

1.  Run `pulumi up` to preview and deploy changes:

    ```
    $ pulumi up
    Previewing changes:
    ...

    Performing changes:
    ...
    Resources:
        + X created
    Duration: ...
    ```

1.  Check the deployed website endpoint:

    ```
    $ pulumi stack output endpoint
    https://<identifier>.web.core.windows.net/api/hello
    $ curl "$(pulumi stack output endpoint)"
    {"message":"Hello from Spring, Pulumi!"}
    ```

## Running the App locally

1. Run Spring Boot application and send a request:

    ```
    gradle bootRun
    curl localhost:8080
    ```


