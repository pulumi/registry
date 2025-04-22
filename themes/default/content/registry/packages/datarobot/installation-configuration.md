---
# WARNING: this file was fetched from https://raw.githubusercontent.com/datarobot-community/pulumi-datarobot/v0.9.2/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: DataRobot Installation & Configuration
meta_desc: Information on how to install the DataRobot provider.
layout: package
---

## Installing

This package is available for several languages/platforms:

- Python: [`pulumi-datarobot`](https://pypi.org/project/pulumi-datarobot/)
- JavaScript/TypeScript: [`@datarobot/pulumi-datarobot`](https://www.npmjs.com/package/@datarobot/pulumi-datarobot)
- Go: [`github.com/datarobot-community/pulumi-datarobot/sdk`](https://pkg.go.dev/github.com/datarobot-community/pulumi-datarobot/sdk)
- .NET: [`DataRobotPulumi.Datarobot`](https://www.nuget.org/packages/DataRobotPulumi.Datarobot)

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi_datarobot
```

### Javscript/Typescript

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @datarobot/pulumi-datarobot
```

or `yarn`:

```bash
yarn add @datarobot/pulumi-datarobot
```

### Go

```
go get github.com/datarobot-community/pulumi-datarobot/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```
dotnet add package DataRobotPulumi.Datarobot
```

### YAML

No install necessary, just run `pulumi up`.

### Provider Binary

The DataRobot provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```sh
pulumi plugin install resource datarobot <version> --server github://api.github.com/datarobot-community/pulumi-datarobot
```

Replace the version string with your desired version.

## Configuration

You must configure the DataRobot provider for Pulumi with a [DataRobot API Key](https://docs.datarobot.com/en/docs/get-started/acct-mgmt/acct-settings/api-key-mgmt.html#api-key-management) before the provider can be used to access and manage items in your DataRobot account.

- `datarobot:apikey` (environment: `DATAROBOT_API_TOKEN`) - the API key for `datarobot`

Once you have your API Key, there are two ways to provide it to Pulumi:

1. Set the environment variable for the preferred method. For example, to set the environment variable for an API Key:

   ```sh
   $ export DATAROBOT_API_TOKEN=XXXXXXXXXXXXXX
   ```

2. If you prefer to store your API Key alongside your Pulumi stack for easy multi-user access, use configuration to set them.

   ```sh
   $ pulumi config set pulumi-datarobot:apikey --secret
   Value: <paste api key here>
   ```

Make sure to pass `--secret` when setting any sensitive data (in this example `pulumi-datarobot:apikey`) so that it's properly encrypted. The complete list of configuration parameters is in the [DataRobot provider for Pulumi README](https://github.com/datarobot-community/pulumi-datarobot/blob/main/README.md#configuration).
