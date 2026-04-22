---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-newrelic/v5.65.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi/pulumi-newrelic/blob/v5.65.0/docs/_index.md
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Newrelic Provider
meta_desc: Provides an overview on how to configure the Pulumi Newrelic provider.
layout: package
---

## Installation

The Newrelic provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/newrelic`](https://www.npmjs.com/package/@pulumi/newrelic)
* Python: [`pulumi-newrelic`](https://pypi.org/project/pulumi-newrelic/)
* Go: [`github.com/pulumi/pulumi-newrelic/sdk/v5/go/newrelic`](https://github.com/pulumi/pulumi-newrelic)
* .NET: [`Pulumi.Newrelic`](https://www.nuget.org/packages/Pulumi.Newrelic)
* Java: [`com.pulumi/newrelic`](https://central.sonatype.com/artifact/com.pulumi/newrelic)

## Overview

[New Relic](https://newrelic.com/) offers tools that help you fix problems
quickly, maintain complex systems, improve your code, and accelerate your
digital transformation.

Use the navigation to the left to read about the available resources.
## Configuration Reference

The following configuration inputs are supported.

|        Argument        | Required? |                                                                                                Description                                                                                                |
|------------------------|-----------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `accountId`           | Required  | Your New Relic account ID. The `NEW_RELIC_ACCOUNT_ID` environment variable can also be used.                                                                                                              |
| `apiKey`              | Required  | Your New Relic Personal API key (usually prefixed with `NRAK`). The `NEW_RELIC_API_KEY` environment variable can also be used.                                                                            |
| `region`               | Optional  | The region for the data center for which your New Relic account is configured. The `NEW_RELIC_REGION` environment variable can also be used. Valid values are `US`, `EU`, or `JP`. Default value is `US`. |
| `insecureSkipVerify` | Optional  | Trust self-signed SSL certificates. If omitted, the `NEW_RELIC_API_SKIP_VERIFY` environment variable is used.                                                                                             |
| `insightsInsertKey`  | Optional  | Your Insights insert key used when inserting Insights events via the `newrelic.insights.Event` resource. Can also use `NEW_RELIC_INSIGHTS_INSERT_KEY` environment variable.                               |
| `cacertFile`          | Optional  | A path to a PEM-encoded certificate authority used to verify the remote agent's certificate. The `NEW_RELIC_API_CACERT` environment variable can also be used.                                            |
## Authentication Requirements

This provider is in the midst of migrating away from our older REST based APIs
to a newer GraphQL based API that we lovingly call NerdGraph. During this
transition, the provider will be using different endpoints depending on which
resource is in use. Below is a table that reflects the current state of the
resources compared to which endpoint is in use.
### Resources

|                      Resource                       |        Endpoint         |    Authentication     |
|-----------------------------------------------------|-------------------------|-----------------------|
| `newrelic.AccountManagement`                       | NerdGraph               | `apiKey`             |
| `newrelic.AlertChannel`                            | RESTv2                  | `apiKey`             |
| `newrelic.AlertCondition`                          | RESTv2                  | `apiKey`             |
| `newrelic.AlertMutingRule`                        | NerdGraph               | `apiKey`             |
| `newrelic.AlertPolicy`                             | NerdGraph               | `apiKey`             |
| `newrelic.AlertPolicyChannel`                     | RESTv2                  | `apiKey`             |
| `newrelic.ApiAccessKey`                           | NerdGraph               | `apiKey`             |
| `newrelic.plugins.ApplicationSettings`                     | RESTv2                  | `apiKey`             |
| `newrelic.BrowserApplication`                      | NerdGraph               | `apiKey`             |
| `newrelic.cloud.AwsGovcloudIntegrations`          | NerdGraph               | `apiKey`             |
| `newrelic.cloud.AwsGovcloudLinkAccount`          | NerdGraph               | `apiKey`             |
| `newrelic.cloud.AwsIntegrations`                   | NerdGraph               | `apiKey`             |
| `newrelic.cloud.AwsLinkAccount`                   | NerdGraph               | `apiKey`             |
| `newrelic.cloud.AzureIntegrations`                 | NerdGraph               | `apiKey`             |
| `newrelic.cloud.AzureLinkAccount`                 | NerdGraph               | `apiKey`             |
| `newrelic.cloud.GcpIntegrations`                   | NerdGraph               | `apiKey`             |
| `newrelic.cloud.GcpLinkAccount`                   | NerdGraph               | `apiKey`             |
| `newrelic.DataPartitionRule`                      | NerdGraph               | `apiKey`             |
| `newrelic.EntityTags`                              | NerdGraph               | `apiKey`             |
| `newrelic.EventsToMetricsRule`                   | NerdGraph               | `apiKey`             |
| `newrelic.Group`                                    | NerdGraph               | `apiKey`             |
| `newrelic.InfraAlertCondition`                    | Infrastructure REST API | `apiKey`             |
| `newrelic.insights.Event`                           | Insights API            | `insightsInsertKey` |
| `newrelic.KeyTransaction`                          | NerdGraph               | `apiKey`             |
| `newrelic.LogParsingRule`                         | NerdGraph               | `apiKey`             |
| `newrelic.NotificationChannel`                     | NerdGraph               | `apiKey`             |
| `newrelic.NotificationDestination`                 | NerdGraph               | `apiKey`             |
| `newrelic.NrqlAlertCondition`                     | NerdGraph               | `apiKey`             |
| `newrelic.NrqlDropRule`                           | NerdGraph               | `apiKey`             |
| `newrelic.ObfuscationExpression`                   | NerdGraph               | `apiKey`             |
| `newrelic.ObfuscationRule`                         | NerdGraph               | `apiKey`             |
| `newrelic.OneDashboard`                            | NerdGraph               | `apiKey`             |
| `newrelic.OneDashboardJson`                       | NerdGraph               | `apiKey`             |
| `newrelic.OneDashboardRaw`                        | NerdGraph               | `apiKey`             |
| `newrelic.ServiceLevel`                            | NerdGraph               | `apiKey`             |
| `newrelic.synthetics.AlertCondition`               | RESTv2                  | `apiKey`             |
| `newrelic.synthetics.BrokenLinksMonitor`          | NerdGraph               | `apiKey`             |
| `newrelic.synthetics.CertCheckMonitor`            | NerdGraph               | `apiKey`             |
| `newrelic.synthetics.Monitor`                       | NerdGraph               | `apiKey`             |
| `newrelic.synthetics.MultiLocationAlertCondition` | RESTv2                  | `apiKey`             |
| `newrelic.synthetics.PrivateLocation`              | NerdGraph               | `apiKey`             |
| `newrelic.synthetics.ScriptMonitor`                | NerdGraph               | `apiKey`             |
| `newrelic.synthetics.SecureCredential`             | NerdGraph               | `apiKey`             |
| `newrelic.synthetics.StepMonitor`                  | NerdGraph               | `apiKey`             |
| `newrelic.User`                                     | NerdGraph               | `apiKey`             |
| `newrelic.Workflow`                                 | NerdGraph               | `apiKey`             |
| `newrelic.plugins.Workload`                                 | NerdGraph               | `apiKey`             |
### Functions

|               Function               | Endpoint  | Authentication |
|-----------------------------------------|-----------|----------------|
| `newrelic.getAccount`                      | NerdGraph | `apiKey`      |
| `newrelic.AlertChannel`                | RESTv2    | `apiKey`      |
| `newrelic.AlertPolicy`                 | NerdGraph | `apiKey`      |
| `newrelic.getApplication`                  | RESTv2    | `apiKey`      |
| `newrelic.getCloudAccount`                | NerdGraph | `apiKey`      |
| `newrelic.getEntity`                       | NerdGraph | `apiKey`      |
| `newrelic.KeyTransaction`              | NerdGraph | `apiKey`      |
| `newrelic.NotificationDestination`     | NerdGraph | `apiKey`      |
| `newrelic.ObfuscationExpression`       | NerdGraph | `apiKey`      |
| `newrelic.synthetics.PrivateLocation`  | NerdGraph | `apiKey`      |
| `newrelic.synthetics.SecureCredential` | NerdGraph | `apiKey`      |
| `newrelic.getTestGrokPattern`            | NerdGraph | `apiKey`      |