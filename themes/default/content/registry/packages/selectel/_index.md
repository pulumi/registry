---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/selectel/selectel/7.4.1/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Selectel Provider
meta_desc: Provides an overview on how to configure the Pulumi Selectel provider.
layout: package
---

## Generate Provider

The Selectel provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider selectel/selectel
```
## Overview

Use the Selectel provider to manage Selectel products.

To manage resources available via OpenStack API, use OpenStack Pulumi provider.
## Authentication (6.0.0 and later)

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    selectel:authRegion:
        value: pool
    selectel:authUrl:
        value: https://cloud.api.selcloud.ru/identity/v3/
    selectel:domainName:
        value: "123456"
    selectel:password:
        value: password
    selectel:username:
        value: user

```
## Configuration Reference (6.0.0 and later)

* `domainName` - (Required) Selectel account ID. The account ID is in the top right corner of the [Control panel](https://my.selectel.ru/). For import, use the value in the `OS_DOMAIN_NAME` environment variable. Learn more about [Registration](https://docs.selectel.ru/en/control-panel-actions/account/registration/).

* `username` - (Required) Name of the service user. To get the name, in the [Control panel](https://my.selectel.ru/iam/users_management/users?type=service), go to **Identity & Access Management** ⟶ **User management** ⟶ the **Service users** tab ⟶ copy the name of the required user. For import, use the value in the `OS_USERNAME` environment variable. Learn more about [Service users](https://docs.selectel.ru/en/control-panel-actions/users-and-roles/user-types-and-roles/) and [how to create service user](https://docs.selectel.ru/en/control-panel-actions/users-and-roles/add-user/#add-service-user).

* `password` - (Required, Sensitive) Password of the service user. For import, use the value in the `OS_PASSWORD` environment variable.

* `authUrl`- (Required) Keystone Identity authentication URL for authentication via user credentials. For import, use the value in the `OS_AUTH_URL` environment variable.

* `authRegion` - (Required) Pool where the endpoint for Keystone API and Resell API is located. For import, use the value in the `OS_REGION_NAME` environment variable. Learn more about available pools in the [Availability matrix](https://docs.selectel.ru/en/control-panel-actions/availability-matrix/).

* `userDomainName` - (Optional) Selectel account ID. Use only for users that were created and assigned a role in a different account. Applicable only to public cloud. The account ID is in the top right corner of the [Control panel](https://my.selectel.ru/). For import, use the value in the `OS_USER_DOMAIN_NAME` environment variable.

* `projectId` - (Optional) Unique identifier of the project. Use only to import resources that are associated with the specific project. To get the ID, in the [Control panel](https://my.selectel.ru/), go to the product section in the navigation menu ⟶ project name ⟶ copy the ID of the required project. As an alternative, you can retrieve project ID from the selectel.VpcProjectV2 resource. If skipped, use the `INFRA_PROJECT_ID` environment variable. Learn more about [Projects](https://docs.selectel.ru/en/control-panel-actions/projects/about-projects/).

* `region` - (Optional) Pool, for example, `ru-3`. Use only to import resources from the specific pool. If skipped, use the `INFRA_REGION` environment variable. Learn more about available pools in the [Availability matrix](https://docs.selectel.ru/en/control-panel-actions/availability-matrix/).
## Authentication (4.0.0 up to 5.*)

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    selectel:domainName:
        value: "123456"
    selectel:password:
        value: password
    selectel:username:
        value: user

```
## Configuration Reference (4.0.0 up to 5.*)

* `domainName` - (Required) Selectel account ID. The account ID is in the top right corner of the [Control panel](https://my.selectel.ru/). For import, use the value in the `OS_DOMAIN_NAME` environment variable. Learn more about [Registration](https://docs.selectel.ru/en/control-panel-actions/account/registration/).

* `username` - (Required) Name of the service user. To get the name, in the [Control panel](https://my.selectel.ru/iam/users_management/users?type=service), go to **Identity & Access Management** ⟶ **User management** ⟶ the **Service users** tab ⟶ copy the name of the required user. For import, use the value in the `OS_USERNAME` environment variable. Learn more about [Service users](https://docs.selectel.ru/en/control-panel-actions/users-and-roles/user-types-and-roles/) and [how to create service user](https://docs.selectel.ru/en/control-panel-actions/users-and-roles/add-user/#add-service-user).

* `password` - (Required, Sensitive) Password of the service user. For import, use the value in the `OS_PASSWORD` environment variable.

* `userDomainName` - (Optional) Selectel account ID. Use only for users that were created and assigned a role in a different account. Applicable only to public cloud. The account ID is in the top right corner of the [Control panel](https://my.selectel.ru/). For import, use the value in the `OS_USER_DOMAIN_NAME` environment variable.

* `authUrl`- (Optional) Keystone Identity authentication URL for authentication via user credentials. If skipped, the provider uses the default endpoint `https://cloud.api.selcloud.ru/identity/v3/`. For import, use the value in the `OS_AUTH_URL` environment variable.

* `authRegion` - (Optional) Pool where the endpoint for Keystone API and Resell API is located, for example, `ru-3`. If skipped, the provider uses the default pool `ru-1`. Does not affect the region parameter in the resources, but it is preferable to use one pool in a manifest. For import, use the value in the `OS_REGION_NAME` environment variable. Learn more about available pools in the [Availability matrix](https://docs.selectel.ru/en/control-panel-actions/availability-matrix/).

* `projectId` - (Optional) Unique identifier of the project. Use only to import resources that are associated with the specific project. To get the ID, in the [Control panel](https://my.selectel.ru/), go to the product section in the navigation menu ⟶ project name ⟶ copy the ID of the required project. As an alternative, you can retrieve project ID from the selectel.VpcProjectV2 resource. If skipped, use the `SEL_PROJECT_ID` environment variable. Learn more about [Projects](https://docs.selectel.ru/en/control-panel-actions/projects/about-projects/).

* `region` - (Optional) Pool, for example, `ru-3`. Use only to import resources from the specific pool. If skipped, use the `SEL_REGION` environment variable. Learn more about available pools in the [Availability matrix](https://docs.selectel.ru/en/control-panel-actions/availability-matrix/).
## Authentication (up to 3.11.0)

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    selectel:token:
        value: Kfpfdf7fjdv0_123456

```
## Configuration Reference (up to 3.11.0)

* `token` - (Required) Selectel token. To get the token, in the top right corner of the [Control panel](https://my.selectel.ru/profile/apikeys), go to the account menu ⟶ **Profile and Settings** ⟶ **API keys** ⟶ copy the token. Learn more about [Selectel token](https://developers.selectel.ru/docs/control-panel/authorization/).

* `endpoint` - (Optional) Selectel API endpoint. Use only for test environments. If skipped, the provider automatically uses the official Selectel endpoint.

* `projectId` - (Optional) Unique identifier of the project. Use only to import resources that are associated with the specific project. To get the ID, in the [Control panel](https://my.selectel.ru/), go to the product section in the navigation menu ⟶ project name ⟶ copy the ID of the required project. As an alternative, you can retrieve project ID from the selectel.VpcProjectV2 resource. Learn more about [Projects](https://docs.selectel.ru/en/control-panel-actions/projects/about-projects/). If skipped, use the `SEL_PROJECT_ID` environment variable.

* `region` - (Optional) Pool, for example, `ru-3`. Use only to import resources from the specific pool. Learn more about available pools in the [Availability matrix](https://docs.selectel.ru/en/control-panel-actions/availability-matrix/). If skipped, use the `SEL_REGION` environment variable.