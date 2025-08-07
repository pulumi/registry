---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/sumologic/sumologic/3.1.3/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Sumologic Provider
meta_desc: Provides an overview on how to configure the Pulumi Sumologic provider.
layout: package
---

## Generate Provider

The Sumologic provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider sumologic/sumologic
```
## Overview

This provider is used to manage resources supported by Sumo Logic. The provider needs to be configured with the proper credentials before it can be used.
## Authentication
The Sumo Logic Provider offers a flexible means of providing credentials for authentication. The following methods are supported and explained below:

- Static credentials
- Environment variables
### Static credentials
Static credentials can be provided by adding an `accessId` and `accessKey` in-line in the Sumo Logic provider configuration:

Usage:
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    sumologic:accessId:
        value: your-access-id
    sumologic:accessKey:
        value: your-access-key
    sumologic:environment:
        value: us2

```
### Environment variables
You can provide your credentials via the `SUMOLOGIC_ACCESSID` and `SUMOLOGIC_ACCESSKEY` environment variables, representing your Sumo Logic Access ID and Sumo Logic Access Key, respectively.

Usage:
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:

```

```bash
$ export SUMOLOGIC_ACCESSID="your-access-id"
$ export SUMOLOGIC_ACCESSKEY="your-access-key"
$ export SUMOLOGIC_ENVIRONMENT=us2
$ pulumi preview
```
## Configuration Reference
- `accessId` - (Required) This is the Sumo Logic Access ID. It must be provided, but it can also be source from the SUMOLOGIC_ACCESSID environment variable.
- `accessKey` - (Required) This is the Sumo Logic Access Key. It must be provided, but it can also be sourced from the SUMOLOGIC_ACCESSKEY variable.
- `environment` - (Required) This is the API endpoint to use. See the [Sumo Logic documentation](https://help.sumologic.com/APIs/General_API_Information/Sumo_Logic_Endpoints_and_Firewall_Security) for details on which environment you should use. It must be provided, but it can be sourced from the SUMOLOGIC_ENVIRONMENT variable.
## Common Source Properties

The following properties are common to ALL sources and can be used to configure each source.

- `collectorId` - (Required) The ID of the collector to attach this source to.
- `name` - (Required) The name of the source. This is required, and has to be unique in the scope of the collector. Changing this will force recreation the source.
- `description` - (Optional) Description of the source.
- `category` - (Optional) The source category this source logs to.
- `hostName` - (Optional) The source host this source logs to.
- `timezone` - (Optional) The timezone assigned to the source. The value follows the [tzdata](https://en.wikipedia.org/wiki/Tz_database) naming convention.
- `automaticDateParsing` - (Optional) Determines if timestamp information is parsed or not. Type true to enable automatic parsing of dates (the default setting); type false to disable. If disabled, no timestamp information is parsed at all.
- `multilineProcessingEnabled` - (Optional) Type true to enable; type false to disable. The default setting is true. Consider setting to false to avoid unnecessary processing if you are collecting single message per line files (for example, Linux system.log). If you're working with multiline messages (for example, log4J or exception stack traces), keep this setting enabled.
- `useAutolineMatching` - (Optional) Type true to enable if you'd like message boundaries to be inferred automatically; type false to prevent message boundaries from being automatically inferred (equivalent to the Infer Boundaries option in the UI). The default setting is true.
- `manualPrefixRegexp` - (Optional) When using useAutolineMatching=false, type a regular expression that matches the first line of the message to manually create the boundary. Note that any special characters in the regex, such as backslashes or double quotes, must be escaped.
- `forceTimezone` - (Optional) Type true to force the source to use a specific time zone, otherwise type false to use the time zone found in the logs. The default setting is false.
- `defaultDateFormats` - (Optional) Define the format for the timestamps present in your log messages. You can specify a locator regex to identify where timestamps appear in log lines. Requires 'automatic_date_parsing' set to True.
  + `format` - (Required) The timestamp format supplied as a Java SimpleDateFormat, or "epoch" if the timestamp is in epoch format.
  + `locator` - (Optional) Regular expression to locate the timestamp within the messages.

  Usage:
  ``
- `filters` - (Optional) If you'd like to add a filter to the source.
  + `filterType` - (Required) The type of filter to apply. (Exclude, Include, Mask, or Hash)
  + `name` - (Required) The Name for the filter.
  + `regexp` - (Required) Regular expression to match within the messages. When used with Incude/Exclude the expression must match the entire message. When used with Mask/Hash rules the expression must contain an unnamed capture group to hash/mask.
  + `mask` - (Optional) When applying a Mask rule, replaces the detected expression with this string.

  Usage:
  ``
- `hashAlgorithm` - (Optional) Define the hash algorithm used for Hash type filters. Available values are "MD5" and "SHA-256". The default value will be "MD5".
- `cutoffTimestamp` - (Optional) Only collect data more recent than this timestamp, specified as milliseconds since epoch (13 digit). This maps to the `Collection should begin` field on the UI. Example: using `1663786159000` will set the cutoff timestamp to `Wednesday, September 21, 2022 6:49:19 PM GMT`
- `cutoffRelativeTime` - (Optional) Can be specified instead of cutoffTimestamp to provide a relative offset with respect to the current time.This maps to the `Collection should begin` field on the UI. Example: use -1h, -1d, or -1w to collect data that's less than one hour, one day, or one week old, respectively.
- `fields` - (Optional) Map containing key/value pairs.

  Usage:
  ``
## Configuring SNS Subscription
This is supported in the following resources.
- `sumologic.CloudfrontSource`
- `sumologic.CloudtrailSource`
- `sumologic.ElbSource`
- `sumologic.S3AuditSource`
- `sumologic.S3Source`

Steps to configure SNS subscription and sync the state in terrform:
- Step 1: Create the source via pulumi.
- Step 2: Setup [SNS subscription](https://help.sumologic.com/03Send-Data/Sources/02Sources-for-Hosted-Collectors/Amazon-Web-Services/AWS_Sources#set-up-sns-in-aws-highly-recommended) outside of pulumi on Sumologic UI.
- Step 3: Run `pulumi preview -refresh-only` to review the changes and verify the state with the SNS subscription information. Make sure only `snsTopicOrSubscriptionArn` is updated. If SNS has been successfully configured and has received a subscription confirmation request `isSuccess` parameter will be true.
- Step 4: Apply the changes with `pulumi up -refresh-only`.