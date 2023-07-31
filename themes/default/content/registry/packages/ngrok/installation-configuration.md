---
title: ngrok Installation & Configuration
meta_desc: Information on how to install the ngrok provider.
layout: package
---

## Installation

The Pulumi ngrok provider is available as a package from the following repositories:

* JavaScript/TypeScript: [`@pierskarsenbarg/ngrok`](https://www.npmjs.com/package/@pierskarsenbarg/ngrok)
* Python: [`pierskarsenbarg-pulumi-ngrok`](https://pypi.org/project/pierskarsenbarg-pulumi-ngrok/)
* Go: [`github.com/pierskarsenbarg/pulumi-ngrok/sdk/go/ngrok`](https://github.com/pierskarsenbarg/pulumi-ngrok)
* .NET: [`PiersKarsenbarg.Ngrok`](https://www.nuget.org/packages/PiersKarsenbarg.Ngrok)

## Setup

To provision resources with the Pulumi ngrok provider, you need to provide the `apiKey`. 

Optionally you can specify the API URL to talk with ngrok. The default is `https://api.ngrok.com`.

## Configuration Options

Use `pulumi config set ngrok:<option> --secret`.

| Option     | Required/Optional | Description                                                                                     |
|------------|-------------------|-------------------------------------------------------------------------------------------------|
| `apiKey` | Required          | This is the ngrok API key (environment: `NGROK_API_KEY`).      |
| `apiBaseUrl` | Optional          | The API URL used to talk with ngrok. The default is `https://api.ngrok.com.` (environment: `NGROK_API_BASE_URL`). |

{{% notes type="warning" %}}
You should use the `--secret` flag to encrypt the config values using your secret provider. For more information on this, view the [Pulumi Configuration Secrets](https://www.pulumi.com/docs/intro/concepts/secrets/#secrets) section in the Pulumi docs.
{{% /notes %}}