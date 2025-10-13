---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-slack/v0.4.14/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi/pulumi-slack/blob/v0.4.14/docs/installation-configuration.md
title: Slack Setup
meta_desc: Installation and configuration for the Pulumi Slack provider.
layout: installation
---

## Installation

The Pulumi Slack provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/slack`](https://www.npmjs.com/package/@pulumi/slack)
* Python: [`pulumi_slack`](https://pypi.org/project/pulumi-slack/)
* Go: [`github.com/pulumi/pulumi-slack/sdk/go/slack`](https://github.com/pulumi/pulumi-slack/tree/main/sdk/go/slack)
* .NET: [`Pulumi.Slack`](https://www.nuget.org/packages/Pulumi.Slack)

## Setup

The Slack provider must be configured with a token to manage resources in a Slack Workspace. Documentation about how to create an app and token for your workspace here [here](https://api.slack.com/apps)

## Configuration Options

Use `pulumi config set slack:token --secret` or set as an environment variable: `export SLACK_TOKEN="my-token"`

| Option | Required/Optional | Description |
|-----|------|----|
| `token`| Required | [Slack token](https://api.slack.com/apps)
