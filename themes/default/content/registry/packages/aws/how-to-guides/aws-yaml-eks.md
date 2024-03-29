---
title: "Amazon EKS Cluster | YAML"
h1: "Amazon EKS Cluster"
linktitle: "Amazon EKS Cluster"
meta_desc: "Amazon EKS Cluster How-to Guide using YAML"
no_edit_this_page: true
cloud: aws
language: yaml
layout: package
---

<!-- WARNING: this page was generated by a tool. Do not edit it by hand. -->
<!-- To change it, please see https://github.com/pulumi/docs/tree/master/tools/mktutorial. -->

<p class="mb-4 flex">
    <a class="flex flex-wrap items-center rounded-md font-display text-lg text-white bg-blue-600 border-2 border-blue-600 px-2 mr-2 whitespace-no-wrap hover:text-white" style="height: 45px;" href="https://github.com/pulumi/examples/tree/master/aws-yaml-eks" target="_blank">
        <span><i class="fab fa-github pr-2"></i> View Code</span>
    </a>
</p>


This example deploys an EKS Kubernetes cluster inside the default AWS VPC.

## Deploying the App

To deploy your infrastructure, follow the below steps.

## Prerequisites

1. [Install Pulumi](https://www.pulumi.com/docs/get-started/install/)
1. [Configure Pulumi for AWS](https://www.pulumi.com/docs/intro/cloud-providers/aws/setup/)

## Deploying and running the program

1.  Create a new stack:

    ```
    $ pulumi stack init dev
    ```

1.  Set the AWS region:

    ```
    $ pulumi config set aws:region us-east-2
    ```

1.  Run `pulumi up` to preview and deploy changes:

    ```
    $ pulumi up
    Previewing changes:
    ...

    Performing changes:
    ...
    Resources:
        + 28 created

    Duration: 10m0s
    ```

1.  Check the deployed kubeconfig:

    ```
    $ pulumi stack output kubeconfig
    {"apiVersion":"v1","clusters":[{"cluster":{"certificate-authority-data":"LS0tLS1CRUdJTiBDR...
    ```

