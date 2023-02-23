---
title: Cloud Toolkit AWS Installation & Configuration
meta_desc: How to set up credentials to use the Cloud Toolkit AWS component library.
layout: installation
---

## Installation

The Cloud Toolkit AWS provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@cloudtoolkit/aws`](https://www.npmjs.com/package/@cloudtoolkit/aws)
* Python: [`cloud-toolkit-aws`](https://pypi.org/project/cloud-toolkit-aws/)

## Configuration

To provision resources with the Cloud Toolkit AWS provider, you need to have AWS credentials. You can use the instructions on if you plan to use AWS credentials from a shared credentials file (which the AWS CLI usually manages for you) or from an environment variable. For more details, see the [AWS documentation](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html).
