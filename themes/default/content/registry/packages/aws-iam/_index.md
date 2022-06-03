---
title: AWS IAM
meta_desc: Use Pulumi's Component for creating an AWS IAM resources using infrastructure as code.
layout: overview
---

You can use the Pulumi AWS IAM Component to help you create AWS IAM roles in all supported Pulumi Languages. The code below
will show you examples of each resources supported in this Component, but please refer to the API Docs for more detailed
descriptions and information about each resource.

This Component was heavily inspired by the [Terraform AWS IAM Module](https://github.com/terraform-aws-modules/terraform-aws-iam/) and
should provide users with equal functionality.

## Examples

{{< chooser language "typescript,python,go,csharp,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as iam from "@pulumi/aws-iam";

// Account
export const account = new iam.Account("account", {
    accountAlias: "cool-alias",
    passwordPolicy: {
        minimumLength: 37,
        requireNumbers: false,
        allowUsersToChange: true,
        hardExpiry: true,
        requireSymbols: true,
        requireLowercaseCharacters: true,
        requireUppercaseCharacters: true,
    },
});

// User
export const user = new iam.User("aws-iam-example-user", {
    name: "pulumipus",
    forceDestroy: true,
    pgpKey: "keybase:test",
    passwordResetRequired: false,
});

// Assumable Role
export const assumableRole = new iam.AssumableRole("aws-iam-example-assumable-role", {
    trustedRoleArns: [ "arn:aws:iam::307990089504:root", "arn:aws:iam::835367859851:user/pulumipus" ],
    role: {
        name: "custom",
        requiresMfa: true,
        policyArns: [ "arn:aws:iam::aws:policy/AmazonCognitoReadOnly","arn:aws:iam::aws:policy/AlexaForBusinessFullAccess" ],
    },
});

// // Assumable Role With OIDC
export const assumableRoleWithOidc = new iam.AssumableRoleWithOIDC("aws-iam-example-assumable-role-with-oidc", {
    providerUrls: ["oidc.eks.eu-west-1.amazonaws.com/id/BA9E170D464AF7B92084EF72A69B9DC8"],
    role: {
        name: "oidc-role",
        policyArns: [ "arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy" ],
    },
    tags: {
        Role: "oidc-role",
    },
});

// // Assumable Role With SAML
export const assumableRoleWithSaml = new iam.AssumableRoleWithSAML("aws-iam-example-assumable-role-with-saml", {
    providerIds: [ "arn:aws:iam::235367859851:saml-provider/idp_saml" ],
    role: {
        name: "saml-role",
        policyArns: [ "arn:aws:iam::aws:policy/ReadOnlyAccess" ],
    },
    tags: {
        Role: "saml-role",
    },
});

// // Assumable Roles
export const assumableRoles = new iam.AssumableRoles("aws-iam-example-assumable-roles", {
    trustedRoleArns: [ "arn:aws:iam::307990089504:root", "arn:aws:iam::835367859851:user/pulumipus" ],
    admin: {},
    poweruser: {
        name: "developer",
    },
    readonly: {
        requiresMfa: true,
    },
});

// // Assumable Roles With SAML
export const assumableRolesWithSaml = new iam.AssumableRolesWithSAML("aws-iam-example-assumable-role-with-saml", {
    providerIds: [ "arn:aws:iam::235367859851:saml-provider/idp_saml" ],
    admin: {},
    poweruser: {
        name: "developer",
    },
    readonly: {},
});

// // EKS Role
export const eksRole = new iam.EKSRole("aws-iam-example-eks-role", {
    role: {
        name: "eks-role",
        policyArns: [ "arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy" ],
    },
    tags: {
        Name: "eks-role",
    },
    clusterServiceAccounts: {
        "cluster1": [ "default:my-app" ],
        "cluster2": [ "default:my-app", "canary:my-app" ],
    },
});

// // Group With Assumable Roles Policy
export const groupWithAssumableRolesPolicy = new iam.GroupWithAssumableRolesPolicy("aws-iam-example-group-with-assumable-roles-policy", {
    name: "production-readonly",
    assumableRoles: [ "arn:aws:iam::835367859855:role/readonly" ],
    groupUsers: [ "pulumipus" ],
}, { dependsOn: [user] });

// // Group With Policies
export const groupWithPolicies = new iam.GroupWithPolicies("aws-iam-example-group-with-policies", {
    name: "superadmins",
    groupUsers: [ "pulumipus" ],
    attachIamSelfManagementPolicy: true,
    customGroupPolicyArns: [ "arn:aws:iam::aws:policy/AdministratorAccess" ],
    customGroupPolicies: [{
        "name": "AllowS3Listing",
        "policy": "{}",
    }],
}, { dependsOn: [user] });

// // Policy
export const policy = new iam.Policy("aws-iam-example-policy", {
    name: "aws-iam-example-policy",
    path: "/",
    description: "My example policy",
    policyDocument: `{
        "Version": "2012-10-17",
        "Statement": [
          {
            "Action": [
              "ec2:Describe*"
            ],
            "Effect": "Allow",
            "Resource": "*"
          }
        ]
    }`,
});

// // Read Only Policy
export const readOnlyPolicy = new iam.ReadOnlyPolicy("aws-iam-example-read-only-policy", {
    name: "aws-iam-example-read-only",
    path: "/",
    description: "My example read only policy",
    allowedServices: [ "rds", "dynamodb" ],
});

// // Role For Service Accounts EKS
export const roleForServiceAccountsEks = new iam.RoleForServiceAccountsEks("aws-iam-example-role-for-service-accounts-eks", {
    role: {
        name: "vpc-cni"
    },
    tags: {
        Name: "vpc-cni-irsa",
    },
    oidcProviders: {
        main: {
            providerArn: "arn:aws:iam::012345678901:oidc-provider/oidc.eks.us-east-1.amazonaws.com/id/5C54DDF35ER19312844C7333374CC09D",
            namespaceServiceAccounts: ["default:my-app", "canary:my-app"],
        }
    },
    policies: {
        vpnCni: {
            attach: true,
            enableIpv4: true,
        },
    },
});

```

{{% /choosable %}}

{{% choosable language python %}}

```python
"""An AWS Python Pulumi program"""

import json
import pulumi
from pulumi_aws import s3
import pulumi_aws_iam as iam

# Account
account = iam.Account(
    'account',
    account_alias='cool-alias',
    password_policy=iam.AccountPasswordPolicyArgs(
        minimum_length=37,
        require_numbers=False,
        allow_users_to_change=True,
        hard_expiry=True,
        require_symbols=True,
        require_lowercase_characters=True,
        require_uppercase_characters=True,
    )
)

pulumi.export('account', account)

# Assumable Role
assumable_role = iam.AssumableRole(
    'assumable_role',
    trusted_role_arns=['arn:aws:iam::307990089504:root','arn:aws:iam::835367859851:user/pulumipus'],
    role=iam.RoleWithMFAArgs(
        name='custom',
        requires_mfa=True,
        policy_arns=['arn:aws:iam::aws:policy/AmazonCognitoReadOnly','arn:aws:iam::aws:policy/AlexaForBusinessFullAccess'],
    ),
)

pulumi.export('assumable_role', assumable_role)

# Assumable Role With OIDC
assumable_role_with_oidc = iam.AssumableRoleWithOIDC(
    'assumable_role_with_oidc',
    role=iam.RoleArgs(
        name='oidc-role',
        policy_arns=['arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy']
    ),
    tags={
        'Role': 'oidc-role',
    },
    provider_urls=['oidc.eks.eu-west-1.amazonaws.com/id/BA9E170D464AF7B92084EF72A69B9DC8']
)

pulumi.export('assumable_role_with_oidc', assumable_role_with_oidc)

# Assumable Role With SAML
assumable_role_with_saml = iam.AssumableRoleWithSAML(
    'assumable_role_with_saml',
    role=iam.RoleArgs(
        name='saml-role',
        policy_arns=['arn:aws:iam::aws:policy/ReadOnlyAccess'],
    ),
    tags={
        'Role': 'saml-role',
    },
    provider_ids=['arn:aws:iam::235367859851:saml-provider/idp_saml']
)

pulumi.export('assumable_role_with_saml', assumable_role_with_saml)

# Assumable Roles
assumable_roles = iam.AssumableRoles(
    'assumable_roles',
    trusted_role_arns=['arn:aws:iam::307990089504:root','arn:aws:iam::835367859851:user/pulumipus'],
    admin=iam.AdminRoleArgs(),
    poweruser=iam.PoweruserRoleArgs(
        name='developer',
    ),
    readonly=iam.ReadonlyRoleWithMFAArgs(
        requires_mfa=True,
    ),
)

pulumi.export('assumable_roles', assumable_roles)

# Assumable Roles With SAML
assumable_roles_with_saml = iam.AssumableRolesWithSAML(
    'assumable_roles_with_saml',
    provider_ids=['arn:aws:iam::235367859851:saml-provider/idp_saml'],
    admin=iam.AdminRoleArgs(),
    readonly=iam.ReadonlyRoleArgs(),
    poweruser=iam.PoweruserRoleArgs(
        name='developer',
    ),
)

pulumi.export('assumable_roles_with_saml', assumable_roles_with_saml)

# EKS Role
eks_role = iam.EKSRole(
    'eks_role',
    role=iam.RoleArgs(
        name='eks-role',
        policy_arns=['arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy'],
    ),
    tags={
        'Name': 'eks-role',
    },
    # cluster_service_acccounts={
    #     'cluster1': [ 'default:my-app' ],
    #     'cluster2': [ 'default:my-app', 'canary:my-app' ],
    # },
)

pulumi.export('eks_role', eks_role)

# Group With Assumable Roles Policy
group_with_assume_roles_policy = iam.GroupWithAssumableRolesPolicy(
    'group_with_assume_roles_policy',
    name='production-readonly',
    assumable_roles=['arn:aws:iam::835367859855:role/readonly'],
    group_users=['user1','user2'],
)

pulumi.export('group_with_assume_roles_policy', group_with_assume_roles_policy)

# Group With Policies
group_with_policies = iam.GroupWithPolicies(
    'group_with_policies',
    name='superadmins',
    group_users=['user1','user2'],
    attach_iam_self_management_policy=True,
    custom_group_policy_arns=['arn:aws:iam::aws:policy/AdministratorAccess'],
    custom_group_policies=[{
        'name': 'AllowS3Listing',
        'policy': '{}',
    }],
)

pulumi.export('group_with_policies', group_with_policies)

# Policy
policy = iam.Policy(
    'policy',
    name='example',
    path='/',
    description='My example policy',
    policy_document=json.dumps({
        "Version": "2012-10-17",
        "Statement": [
          {
            "Action": [
              "ec2:Describe*"
            ],
            "Effect": "Allow",
            "Resource": "*"
          }
        ]
    })
)

pulumi.export('policy', policy)

# Read Only Policy
read_only_policy = iam.ReadOnlyPolicy(
    'read_only_policy',
    name='example',
    path='/',
    description='My example read only policy',
    allowed_services=['rds','dynamo'],
)

pulumi.export('read_only_policy', read_only_policy)

# Role For Service Accounts EKS
role_for_service_account_eks = iam.RoleForServiceAccountsEks(
    'role_for_service_account_eks',
    role=iam.RoleArgs(
        name='vpc-cni'
    ),
    tags={
        'Name': 'vpc-cni-irsa',
    },
    oidc_providers={
        'main': iam.OIDCProviderArgs(
            provider_arn='arn:aws:iam::012345678901:oidc-provider/oidc.eks.us-east-1.amazonaws.com/id/5C54DDF35ER19312844C7333374CC09D',
            namespace_service_accounts=['default:my-app', 'canary:my-app'],
        ),
    },
    policies=iam.EKSRolePoliciesArgs(
        vpn_cni=iam.EKSVPNCNIPolicyArgs(
            attach=True,
            enable_ipv4=True,
        ),
    ),
)

pulumi.export('role_for_service_account_eks', role_for_service_account_eks)

# User
user = iam.User(
    'user',
    name='pulumipus',
    force_destroy=True,
    pgp_key='keybase:test',
    password_reset_required=False,
)

pulumi.export('user', user)
```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"encoding/json"

	iam "github.com/pulumi/pulumi-aws-iam/sdk/go/aws-iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Account
		account, err := iam.NewAccount(ctx, "account", &iam.AccountArgs{
			AccountAlias: pulumi.String("cool-alias"),
			PasswordPolicy: iam.AccountPasswordPolicyArgs{
				MinimumLength:              pulumi.IntPtr(37),
				RequireNumbers:             pulumi.Bool(false),
				AllowUsersToChange:         pulumi.Bool(true),
				HardExpiry:                 pulumi.Bool(true),
				RequireSymbols:             pulumi.Bool(true),
				RequireLowercaseCharacters: pulumi.Bool(true),
				RequireUppercaseCharacters: pulumi.Bool(true),
			},
		})
		if err != nil {
			return err
		}

		ctx.Export("account", account)

		// Assumable Role
		assumableRole, err := iam.NewAssumableRole(ctx, "assumable-role", &iam.AssumableRoleArgs{
			TrustedRoleArns: pulumi.ToStringArray([]string{"arn:aws:iam::307990089504:root", "arn:aws:iam::835367859851:user/pulumipus"}),
			Role: &iam.RoleWithMFAArgs{
				Name:        pulumi.String("custom"),
				RequiresMfa: pulumi.BoolPtr(true),
				PolicyArns:  pulumi.ToStringArray([]string{"arn:aws:iam::aws:policy/AmazonCognitoReadOnly", "arn:aws:iam::aws:policy/AlexaForBusinessFullAccess"}),
			},
		})
		if err != nil {
			return err
		}

		ctx.Export("assumableRole", assumableRole)

		// Assumable Role With OIDC
		assumableRoleWithOIDC, err := iam.NewAssumableRoleWithOIDC(ctx, "assumable-role-with-oidc", &iam.AssumableRoleWithOIDCArgs{
			Role: iam.RoleArgs{
				Name:       pulumi.String("oidc-role"),
				PolicyArns: pulumi.ToStringArray([]string{"arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy"}),
			},
			Tags: pulumi.ToStringMap(map[string]string{
				"Role": "oidc-role",
			}),
			ProviderUrls: pulumi.ToStringArray([]string{"oidc.eks.eu-west-1.amazonaws.com/id/BA9E170D464AF7B92084EF72A69B9DC8"}),
		})
		if err != nil {
			return err
		}

		ctx.Export("assumableRoleWithOIDC", assumableRoleWithOIDC)

		// Assumable Role With SAML
		assumableRoleWithSAML, err := iam.NewAssumableRoleWithSAML(ctx, "assumable-role-with-saml", &iam.AssumableRoleWithSAMLArgs{
			Role: iam.RoleArgs{
				Name:       pulumi.String("saml-role"),
				PolicyArns: pulumi.ToStringArray([]string{"arn:aws:iam::aws:policy/ReadOnlyAccess"}),
			},
			Tags: pulumi.ToStringMap(map[string]string{
				"Role": "saml-role",
			}),
			ProviderIds: pulumi.ToStringArray([]string{"arn:aws:iam::235367859851:saml-provider/idp_saml"}),
		})
		if err != nil {
			return err
		}

		ctx.Export("assumableRoleWithSAML", assumableRoleWithSAML)

		// Assumable Roles
		assumableRoles, err := iam.NewAssumableRoles(ctx, "assumable-roles", &iam.AssumableRolesArgs{
			TrustedRoleArns: pulumi.ToStringArray([]string{"arn:aws:iam::307990089504:root", "arn:aws:iam::835367859851:user/pulumipus"}),
			Admin:           iam.AdminRoleWithMFAArgs{},
			Poweruser: iam.PoweruserRoleWithMFAArgs{
				Name: pulumi.String("developer"),
			},
			Readonly: iam.ReadonlyRoleWithMFAArgs{
				RequiresMfa: pulumi.BoolPtr(true),
			},
		})
		if err != nil {
			return err
		}

		ctx.Export("assumableRoles", assumableRoles)

		// Assumable Roles With SAML
		assumableRolesWithSAML, err := iam.NewAssumableRolesWithSAML(ctx, "assumable-roles-with-saml", &iam.AssumableRolesWithSAMLArgs{
			ProviderIds: pulumi.ToStringArray([]string{"arn:aws:iam::235367859851:saml-provider/idp_saml"}),
			Admin:       iam.AdminRoleArgs{},
			Readonly:    iam.ReadonlyRoleArgs{},
			Poweruser: iam.PoweruserRoleArgs{
				Name: pulumi.String("developer"),
			},
		})
		if err != nil {
			return err
		}

		ctx.Export("assumableRolesWithSAML", assumableRolesWithSAML)

		// EKS Role
		eksRole, err := iam.NewEKSRole(ctx, "eks-role", &iam.EKSRoleArgs{
			Role: iam.RoleArgs{
				Name:       pulumi.String("eks-role"),
				PolicyArns: pulumi.ToStringArray([]string{"arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy"}),
			},
			Tags: pulumi.ToStringMap(map[string]string{
				"Role": "eks-role",
			}),
			// Uncomment the below and replace actual cluster values.
			// ClusterServiceAccounts: pulumi.ToStringArrayMap(map[string][]string{
			// 	"cluster1": {"default:my-app"},
			// 	"cluster2": {"default:my-app", "canary:my-app"},
			// }),
		})
		if err != nil {
			return err
		}

		ctx.Export("eksRole", eksRole)

		// Group With Assumable Roles Policy
		groupWithAssumableRolesPolicy, err := iam.NewGroupWithAssumableRolesPolicy(ctx, "group-with-assumable-roles-policy", &iam.GroupWithAssumableRolesPolicyArgs{
			Name:           pulumi.String("production-readonly"),
			AssumableRoles: pulumi.ToStringArray([]string{"arn:aws:iam::835367859855:role/readonly"}),
			GroupUsers:     pulumi.ToStringArray([]string{"user1", "user2"}),
		})
		if err != nil {
			return err
		}

		ctx.Export("groupWithAssumableRolesPolicy", groupWithAssumableRolesPolicy)

		// Group With Policies
		groupWithPolicies, err := iam.NewGroupWithPolicies(ctx, "group-with-policies", &iam.GroupWithPoliciesArgs{
			Name:                          pulumi.String("superadmins"),
			GroupUsers:                    pulumi.ToStringArray([]string{"user1", "user2"}),
			AttachIamSelfManagementPolicy: pulumi.BoolPtr(true),
			CustomGroupPolicyArns:         pulumi.ToStringArray([]string{"arn:aws:iam::aws:policy/AdministratorAccess"}),
			CustomGroupPolicies: pulumi.ToStringMapArray([]map[string]string{
				{
					"name":   "AllowS3Listing",
					"policy": "{}",
				},
			}),
		})
		if err != nil {
			return err
		}

		ctx.Export("groupWithPolicies", groupWithPolicies)

		// Policy
		policyJSON, err := json.Marshal(map[string]interface{}{
			"Version": "2012-10-17",
			"Statement": []interface{}{
				map[string]interface{}{
					"Effect":   "Allow",
					"Action":   []string{"ec2:Describe"},
					"Resource": []string{"*"},
				},
			},
		})
		if err != nil {
			return err
		}

		policy, err := iam.NewPolicy(ctx, "policy", &iam.PolicyArgs{
			Name:           pulumi.String("example"),
			Path:           pulumi.String("/"),
			Description:    pulumi.String("My example policy"),
			PolicyDocument: pulumi.String(string(policyJSON)),
		})
		if err != nil {
			return err
		}

		ctx.Export("policy", policy)

		// Read Only Policy
		readOnlyPolicy, err := iam.NewReadOnlyPolicy(ctx, "read-only-policy", &iam.ReadOnlyPolicyArgs{
			Name:            pulumi.String("example"),
			Path:            pulumi.String("/"),
			Description:     pulumi.String("My example policy"),
			AllowedServices: pulumi.ToStringArray([]string{"rds", "dynamo"}),
		})
		if err != nil {
			return err
		}

		ctx.Export("readOnlyPolicy", readOnlyPolicy)

		// Role For Service Accounts EKS
		roleForServiceAccountsEKS, err := iam.NewRoleForServiceAccountsEks(ctx, "role-for-service-accounts-eks", &iam.RoleForServiceAccountsEksArgs{
			Role: iam.EKSServiceAccountRolePtr(&iam.EKSServiceAccountRoleArgs{
				Name: pulumi.String("vpc-cni"),
			}),
			Tags: pulumi.ToStringMap(map[string]string{
				"Name": "vpc-cni-irsa",
			}),
			OidcProviders: iam.OIDCProviderMap{
				"main": iam.OIDCProviderArgs{
					ProviderArn:              pulumi.String("arn:aws:iam::012345678901:oidc-provider/oidc.eks.us-east-1.amazonaws.com/id/5C54DDF35ER19312844C7333374CC09D"),
					NamespaceServiceAccounts: pulumi.ToStringArray([]string{"default:my-app", "canary:my-app"}),
				},
			},
			Policies: iam.EKSRolePoliciesPtr(&iam.EKSRolePoliciesArgs{
				VpnCni: iam.EKSVPNCNIPolicyPtr(&iam.EKSVPNCNIPolicyArgs{
					Attach:     pulumi.Bool(true),
					EnableIpv4: pulumi.BoolPtr(true),
				}),
			}),
		})
		if err != nil {
			return err
		}

		ctx.Export("roleForServiceAccountsEKS", roleForServiceAccountsEKS)

		// User
		user, err := iam.NewUser(ctx, "user", &iam.UserArgs{
			Name:                  pulumi.String("pulumipus"),
			ForceDestroy:          pulumi.BoolPtr(true),
			PgpKey:                pulumi.String("keybase:test"),
			PasswordResetRequired: pulumi.BoolPtr(false),
		})
		if err != nil {
			return err
		}

		ctx.Export("user", user)

		return nil
	})
}
```

{{% /choosable %}}

{{% choosable language csharp %}}
```csharp
using Pulumi;
using Pulumi.AwsIam;
using Pulumi.AwsIam.Inputs;
using System.Collections.Immutable;

class MyStack : Stack
{
    public MyStack()
    {
        // Account
        var account = new Account("account", new AccountArgs
        {
            AccountAlias = "cool-alias",
            PasswordPolicy=new AccountPasswordPolicyArgs
            {
                MinimumLength = 37,
                RequireNumbers = false,
                AllowUsersToChange = true,
                HardExpiry = true,
                RequireSymbols = true,
                RequireLowercaseCharacters = true,
                RequireUppercaseCharacters = true,
            }

        });

        this.Account = Output.Create<Account>(account);

        // Assumable Role
        var assumableRole = new AssumableRole("assumable-role", new AssumableRoleArgs
        {
            TrustedRoleArns = {"arn:aws:iam::307990089504:root", "arn:aws:iam::835367859851:user/pulumipus"},
            Role = new RoleWithMFAArgs
            {
                Name = "custom",
                RequiresMfa = true,
                PolicyArns = {"arn:aws:iam::aws:policy/AmazonCognitoReadOnly","arn:aws:iam::aws:policy/AlexaForBusinessFullAccess"},
            },
        });

        this.AssumableRole = Output.Create<AssumableRole>(assumableRole);

        // Assumable Role With OIDC
        var assumableRoleWithOidc = new AssumableRoleWithOIDC("assumable-role-with-oidc", new AssumableRoleWithOIDCArgs
        {
            Role = new RoleArgs
            {
                Name = "oidc-role",
                PolicyArns = {"arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy"},
            },
            Tags = new InputMap<string>
            {
                {"Role", "odic-role"},
            },
            ProviderUrls = {"oidc.eks.eu-west-1.amazonaws.com/id/BA9E170D464AF7B92084EF72A69B9DC8"},
        });

        this.AssumableRoleWithOidc = Output.Create<AssumableRoleWithOIDC>(assumableRoleWithOidc);

        // Assumable Role With SAML
        var assumableRoleWithSaml = new AssumableRoleWithSAML("assumable-role-with-saml", new AssumableRoleWithSAMLArgs
        {
            Role = new RoleArgs
            {
                Name = "saml-role",
                PolicyArns = {"arn:aws:iam::aws:policy/ReadOnlyAccess"},
            },
            Tags = new InputMap<string>
            {
                {"Role", "saml-role"},
            },
            ProviderIds = {"arn:aws:iam::235367859851:saml-provider/idp_saml"},
        });

        this.AssumableRoleWithSaml = Output.Create<AssumableRoleWithSAML>(assumableRoleWithSaml);

        // Assumable Roles
        var assumableRoles = new AssumableRoles("assumable-roles", new AssumableRolesArgs
        {
            TrustedRoleArns = {"arn:aws:iam::307990089504:root", "arn:aws:iam::835367859851:user/pulumipus"},
            Admin = new AdminRoleWithMFAArgs(),
            Poweruser = new PoweruserRoleWithMFAArgs
            {
                Name = "developer",
            },
            Readonly = new ReadonlyRoleWithMFAArgs
            {
                RequiresMfa = true,
            },
        });

        this.AssumableRoles = Output.Create<AssumableRoles>(assumableRoles);

        // Assumable Roles With SAML
        var assumableRolesWithSaml = new AssumableRolesWithSAML("assumable-roles-with-saml", new AssumableRolesWithSAMLArgs
        {
            ProviderIds = {"arn:aws:iam::235367859851:saml-provider/idp_saml"},
            Admin = new AdminRoleArgs(),
            Readonly = new ReadonlyRoleArgs(),
            Poweruser = new PoweruserRoleArgs
            {
                Name = "developer",
            },
        });

        this.AssumableRolesWithSaml = Output.Create<AssumableRolesWithSAML>(assumableRolesWithSaml);

        // EKS Role
        var eksRole = new EKSRole("eks-role", new EKSRoleArgs
        {
            Role = new RoleArgs
            {
                Name = "eks-role",
                PolicyArns = {"arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy"},
            },
            Tags = new InputMap<string>
            {
                {"Name", "eks-role"},
            },
            // Uncomment the below and replace actual cluster values.
            // ClusterServiceAccounts = {
            //     {"cluster1", ImmutableArray.Create<string>(new string[] {"default:my-app"})},
            //     {"cluster2", ImmutableArray.Create<string>(new string[] {"default:my-app", "canary:my-app"})}
            // },
        });

        this.EksRole = Output.Create<EKSRole>(eksRole);

        // Group With Assumable Roles Policy
        var groupWithAssumableRolePolicy = new GroupWithAssumableRolesPolicy("group-with-assumable-roles-policy", new GroupWithAssumableRolesPolicyArgs
        {
            Name = "production-readonly",
            AssumableRoles = {"arn:aws:iam::835367859855:role/readonly"},
            GroupUsers = {"user1", "user2"},
        });

        this.GroupWithAssumableRolesPolicy = Output.Create<GroupWithAssumableRolesPolicy>(groupWithAssumableRolePolicy);

        // Group With Policies
        var groupWithPolicies = new GroupWithPolicies("group-with-policies", new GroupWithPoliciesArgs
        {
            Name = "superadmins",
            GroupUsers = {"user1", "user2"},
            AttachIamSelfManagementPolicy = true,
            CustomGroupPolicyArns = {"arn:aws:iam::aws:policy/AdministratorAccess"},
            CustomGroupPolicies = new InputList<ImmutableDictionary<string, string>>
            {
                ImmutableDictionary.Create<string, string>()
                    .Add("name", "AllowS3Listing")
                    .Add("policy", "{}"),
            },
        });

        this.GroupWithPolicies = Output.Create<GroupWithPolicies>(groupWithPolicies);

        // Policy
        var policy = new Policy("policy", new PolicyArgs
        {
            Name = "example",
            Path = "/",
            Description = "My example policy",
            PolicyDocument =
                @"{
                ""Version"": ""2012-10-17"",
                ""Statement"": [
                {
                    ""Action"": [
                    ""ec2:Describe*""
                    ],
                    ""Effect"": ""Allow"",
                    ""Resource"": ""*""
                }
                ]
            }"
        });

        this.Policy = Output.Create<Policy>(policy);

        // Read Only Policy
        var readOnlyPolicy = new ReadOnlyPolicy("read-only-policy", new ReadOnlyPolicyArgs
        {
            Name = "example",
            Path = "/",
            Description = "My example read only policy",
            AllowedServices = {"rds", "dynamo"},
        });

        this.ReadOnlyPolicy = Output.Create<ReadOnlyPolicy>(readOnlyPolicy);

        // Role For Service Accounts EKS
        var roleForServiceAccountEks = new RoleForServiceAccountsEks("role-for-service-account-eks", new RoleForServiceAccountsEksArgs
        {
            Role = new EKSServiceAccountRoleArgs
            {
                Name = "vpn-cni",
            },
            Tags = {
                {"Name", "vpc-cni-irsa"},
            },
            OidcProviders = {
                {"main", new OIDCProviderArgs
                {
                    ProviderArn = "arn:aws:iam::012345678901:oidc-provider/oidc.eks.us-east-1.amazonaws.com/id/5C54DDF35ER19312844C7333374CC09D",
                    NamespaceServiceAccounts = {"default:my-app", "canary:my-app"},
                }},
            },
            Policies = new EKSRolePoliciesArgs
            {
                VpnCni = new EKSVPNCNIPolicyArgs
                {
                    Attach = true,
                    EnableIpv4 = true,
                },
            },
        });

        this.RoleForServiceAccountEks = Output.Create<RoleForServiceAccountsEks>(roleForServiceAccountEks);

        // User
        var user = new User("user", new UserArgs
        {
            Name = "pulumipus",
            ForceDestroy = true,
            PgpKey = "keybase:test",
            PasswordResetRequired = false,
        });

        this.User = Output.Create<User>(user);
    }

    [Output]
    public Output<Account> Account { get; set; }

    [Output]
    public Output<AssumableRole> AssumableRole { get; set; }

    [Output]
    public Output<AssumableRoleWithOIDC> AssumableRoleWithOidc { get; set; }

    [Output]
    public Output<AssumableRoleWithSAML> AssumableRoleWithSaml { get; set; }

    [Output]
    public Output<AssumableRoles> AssumableRoles { get; set; }

    [Output]
    public Output<AssumableRolesWithSAML> AssumableRolesWithSaml { get; set; }

    [Output]
    public Output<EKSRole> EksRole { get; set; }

    [Output]
    public Output<GroupWithAssumableRolesPolicy> GroupWithAssumableRolesPolicy { get; set; }

    [Output]
    public Output<GroupWithPolicies> GroupWithPolicies { get; set; }

    [Output]
    public Output<Policy> Policy { get; set; }

    [Output]
    public Output<ReadOnlyPolicy> ReadOnlyPolicy { get; set; }

    [Output]
    public Output<RoleForServiceAccountsEks> RoleForServiceAccountEks { get; set; }

    [Output]
    public Output<User> User { get; set; }
}

```
{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: awsiam-yaml
runtime: yaml
resources:
    account:
        type: "aws-iam:index:Account"
        properties:
            accountAlias: "cool-alias"
            passwordPolicy:
                minimumLength: 37
                requireNumbers: false
                allowUsersToChange: true
                hardExpiry: true
                requireSymbols: true
                requireLowercaseCharacters: true
                requireUppercaseCharacters: true

    assumableRole:
        type: "aws-iam:index:AssumableRole"
        properties:
            trustedRoleArns:
                - "arn:aws:iam::307990089504:root"
                - "arn:aws:iam::835367859851:user/pulumipus"
            role:
                name: "custom"
                requiresMfa: true
                policyArns:
                    - "arn:aws:iam::aws:policy/AmazonCognitoReadOnly"
                    - "arn:aws:iam::aws:policy/AlexaForBusinessFullAccess"

    assumableRoleWithOidc:
        type: "aws-iam:index:AssumableRoleWithOIDC"
        properties:
            role:
                name: "oidc-role"
                policyArns:
                    - "arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy"
            tags:
                Role: "oidc-role"
            providerUrls:
                - "oidc.eks.eu-west-1.amazonaws.com/id/BA9E170D464AF7B92084EF72A69B9DC8"

    assumableRoleWithSaml:
        type: "aws-iam:index:AssumableRoleWithSAML"
        properties:
            role:
                name: "saml-role"
                policyArns:
                    - "arn:aws:iam::aws:policy/ReadOnlyAccess"
            tags:
                Role: "saml-role"
            providerIds:
                - "arn:aws:iam::235367859851:saml-provider/idp_saml"

    assumableRoles:
        type: "aws-iam:index:AssumableRoles"
        properties:
            trustedRoleArns:
                - "arn:aws:iam::307990089504:root"
                - "arn:aws:iam::835367859851:user/pulumipus"
            poweruser:
                name: "developer"
            readonly:
                requiresMfa: true

    assumableRolesWithSaml:
        type: "aws-iam:index:AssumableRolesWithSAML"
        properties:
            providerIds:
                - "arn:aws:iam::235367859851:saml-provider/idp_saml"
            poweruser:
                name: "developer"

    eksRole:
        type: "aws-iam:index:EKSRole"
        properties:
            role:
                name: "eks-role"
                policyArns:
                    - "arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy"
            tags:
                Name: "eks-role"
            # Uncomment the below and replace actual cluster values.
            # clusterServiceAccounts:
            #     cluster1:
            #         - "default:my-app"
            #     cluster2:
            #         - "default:my-app"
            #         - "canary:my-app"

    groupWithAssumableRolesPolicy:
        type: "aws-iam:index:GroupWithAssumableRolesPolicy"
        properties:
            name: "production-readonly"
            assumableRoles:
                - "arn:aws:iam::835367859855:role/readonly"
            groupUsers:
                - "user1"
                - "user2"

    groupWithPolicies:
        type: "aws-iam:index:GroupWithPolicies"
        properties:
            name: "superadmins"
            groupUsers:
                - "user1"
                - "user2"
            attachIamSelfManagementPolicy: true
            customGroupPolicyArns:
                - "arn:aws:iam::aws:policy/AdministratorAccess"
            customGroupPolicies:
                - name: "AllowS3Listing"
                  policy: "{}"

    policy:
        type: "aws-iam:index:Policy"
        properties:
            name: "example"
            path: "/"
            description: "My example policy"
            policyDocument: |
                {
                    "Version": "2012-10-17",
                    "Statement": [
                        {
                            "Action": [
                                "ec2:Describe*"
                            ],
                            "Effect": "Allow",
                            "Resource": "*"
                        }
                    ]
                }

    readOnlyPolicy:
        type: "aws-iam:index:ReadOnlyPolicy"
        properties:
            name: "example"
            path: "/"
            description: "My example read only policy"
            allowedServices:
                - "rds"
                - "dynamodb"

    roleForServiceAccountsEks:
        type: "aws-iam:index:RoleForServiceAccountsEks"
        properties:
            role:
                name: "vpc-cni"
            tags:
                Name: "vpc-cni-irsa"
            oidcProviders:
                main:
                    providerArn: "arn:aws:iam::012345678901:oidc-provider/oidc.eks.us-east-1.amazonaws.com/id/5C54DDF35ER19312844C7333374CC09D"
                    namespaceServiceAccounts:
                        - "default:my-app"
                        - "canary:my-app"
            policies:
                vpnCni:
                    attach: true
                    enableIpv4: true

    user:
        type: "aws-iam:index:User"
        properties:
            name: "pulumipus"
            forceDestroy: true
            pgpKey: "keybase:test"
            passwordResetRequired: false

outputs:
    account: ${account}
    assumableRole: ${assumableRole}
    assumableRoleWithOidc: ${assumableRoleWithOidc}
    assumableRoleWithSaml: ${assumableRoleWithSaml}
    assumableRoles: ${assumableRoles}
    assumableRolesWithSaml: ${assumableRolesWithSaml}
    eksRole: ${eksRole}
    groupWithAssumableRolesPolicy: ${groupWithAssumableRolesPolicy}
    groupWithPolicies: ${groupWithPolicies}
    policy: ${policy}
    readOnlyPolicy: ${readOnlyPolicy}
    roleForServiceAccountsEks: ${roleForServiceAccountsEks}
    user: ${user}

```

{{% /choosable %}}


{{< /chooser >}}
