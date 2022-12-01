---
title: OpenStack Installation & Configuration
meta_desc: Provides an overview on how to set up the Pulumi OpenStack Provider with proper credentials.
layout: installation
---

<!-- markdownlint-disable url -->
[Pulumi OpenStack Provider]: {{< relref "./" >}}
[Download your OpenStack credentials]: https://docs.openstack.org/newton/user-guide/common/cli-set-environment-variables-using-openstack-rc.html
<!-- markdownlint-enable url -->

The [Pulumi OpenStack Provider] needs to be configured with OpenStack credentials
before it can be used to create resources.
You will need to provide the OpenStack Provider with your OpenStack account details. You can [Download your OpenStack credentials] as a sourceable rc file from the OpenStack dashboard.
Your credentials are only used to authenticate with OpenStack APIs on your behalf. Your credentials are never sent to pulumi.com.

To communicate your credentials to the Pulumi OpenStack Provider source the file downloaded from OpenStack: `source openrc.sh`.

## Installation

The OpenStack provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/openstack`](https://www.npmjs.com/package/@pulumi/openstack)
* Python: [`pulumi-openstack`](https://pypi.org/project/pulumi-openstack/)
* Go: [`github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack`](https://github.com/pulumi/pulumi-openstack)
* .NET: [`Pulumi.Openstack`](https://www.nuget.org/packages/Pulumi.Openstack)
