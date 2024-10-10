---
title: Migrating aws-auth ConfigMap entries to access entries
meta_desc: How to update from the deprecated aws-auth ConfigMap to Access Entries
layout: package
---

AWS introduced a new method for granting IAM principals access to Kubernetes resources called Access Entries. In contrast to the existing, deprecated approach using the `aws-auth` ConfigMap, this solely relies on AWS resources for managing Kubernetes auth.

Previously, the aws-auth ConfigMap was the sole method for mapping IAM principals to Kubernetes RBAC. Now, users can choose between using the ConfigMap (access mode `CONFIG_MAP`), Access Entries (access mode `API`), or both (access mode `API_AND_CONFIG_MAP`). If no authentication mode is configured for a cluster, AWS defaults to using the aws-auth ConfigMap.

> Caution! Once the access entry method is enabled, it cannot be disabled. If the ConfigMap method is not enabled during cluster creation, it cannot be enabled later. All clusters created before the introduction of access entries have the ConfigMap method enabled.

This guide is based on the [AWS Docs](https://docs.aws.amazon.com/eks/latest/userguide/migrating-access-entries.html) for migrating existing aws-auth ConfigMap entries to access entries.

## Migrate to Access Entries

Before attempting to migrate the authentication mode to access entries you need to ensure your EKS cluster's platform version is compatible with that authentication mode. To do that you can follow the "Prerequisites" section in the [AWS Docs](https://docs.aws.amazon.com/eks/latest/userguide/access-entries.html).

After confirming your cluster supports Access Entries you need to perform the following steps to migrate:

### 1. Switch the Authentication Mode to `API_AND_CONFIG_MAP`
To do this, you need to set the `authenticationMode` parameter of your cluster to `API_AND_CONFIG_MAP`. This step does not modify the cluster's existing authentication configuration. It only enables using the `CONFIG_MAP` and `API` in parallel.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}

{{% choosable language typescript %}}

```ts
new eks.Cluster(`cluster`, {
    authenticationMode: eks.AuthenticationMode.ApiAndConfigMap,
});
```

{{% /choosable %}}

{{% choosable language python %}}

```python
eks.Cluster(
    "cluster",
    authentication_mode=eks.AuthenticationMode.API_AND_CONFIG_MAP,
)
```

{{% /choosable %}}

{{% choosable language go %}}

```go
authMode := eks.AuthenticationModeApiAndConfigMap
eks.NewCluster(ctx, "cluster", &eks.ClusterArgs{
    AuthenticationMode: &authMode,
})
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
new Eks.Cluster("cluster", new Eks.ClusterArgs
{
    AuthenticationMode = Eks.AuthenticationMode.ApiAndConfigMap,
});
```

{{% /choosable %}}

{{% choosable language java %}}

```java
Cluster("cluster", ClusterArgs.builder()
    .authenticationMode(AuthenticationMode.ApiAndConfigMap)
    .build());
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
cluster:
  type: eks:index:Cluster
  properties:
    authenticationMode: API_AND_CONFIG_MAP
```

{{% /choosable %}}

{{< /chooser >}}

### 2. Add Access Entries for entries in the `aws-auth` ConfigMap
The `aws-auth` ConfigMap gets populated based on the following parameters:
- `roleMappings`
- `userMappings`
- `instanceRoles`

To create Access Entries for those you need to choose the right Access Entry type.

#### roleMappings & userMappings
These map to Access Entries of type `STANDARD`, which is the default.

The following shows how you'd add access entries for the role and user mappings:
{{< chooser language "typescript,python,go,csharp,java,yaml" >}}

{{% choosable language typescript %}}

```ts
new eks.Cluster(`cluster`, {
    authenticationMode: eks.AuthenticationMode.ApiAndConfigMap,
    roleMappings: [
        {
            roleArn: iamRole.arn,
            groups: ["test-group"],
            username: "test-role",
        }
    ],
    userMappings: [
        {
            userArn: iamUser.arn,
            groups: ["user-group"],
            username: "test-user",
        }
    ]
    accessEntries: {
        myRole: {
            principalArn: iamRole.arn,
            kubernetesGroups: ["test-group"],
            username: "test-role",
            type: eks.AccessEntryType.STANDARD
        },
        myUser: {
            principalArn: iamUser.arn,
            kubernetesGroups: ["user-group"],
            username: "test-user",
            type: eks.AccessEntryType.STANDARD
        }
    }
});
```

{{% /choosable %}}

{{% choosable language python %}}

```python
eks.Cluster("cluster",
    authentication_mode=eks.AuthenticationMode.API_AND_CONFIG_MAP,
    role_mappings=[{
        "role_arn": "arn:aws:iam::123456789012:role/eks-cluster-role",
        "username": "test-role",
        "groups": ["test-group"],
    }],
    user_mappings=[{
        "user_arn": "arn:aws:iam::123456789012:user/test-user",
        "username": "test-user",
        "groups": ["test-users"],
    }],
    access_entries={
        "myRole": {
            "principal_arn": "arn:aws:iam::123456789012:role/eks-cluster-role",
            "username": "test-role",
            "kubernetes_groups": ["test-group"],
            "type": eks.AccessEntryType.STANDARD,
        },
        "myUser": {
            "principal_arn": "arn:aws:iam::123456789012:user/test-user",
            "username": "test-user",
            "kubernetes_groups": ["test-users"],
            "type": eks.AccessEntryType.STANDARD,
        },
    })
```

{{% /choosable %}}

{{% choosable language go %}}

```go
authMode := eks.AuthenticationModeApiAndConfigMap
eks.NewCluster(ctx, "cluster", &eks.ClusterArgs{
    AuthenticationMode: &authMode,
	RoleMappings: eks.RoleMappingArray{
		&eks.RoleMappingArgs{
			RoleArn:  pulumi.String("arn:aws:iam::123456789012:role/eks-cluster-role"),
			Username: pulumi.String("test-role"),
			Groups: pulumi.StringArray{
				pulumi.String("test-group"),
			},
		},
	},
	UserMappings: eks.UserMappingArray{
		&eks.UserMappingArgs{
			UserArn:  pulumi.String("arn:aws:iam::123456789012:user/test-user"),
			Username: pulumi.String("test-user"),
			Groups: pulumi.StringArray{
				pulumi.String("test-users"),
			},
		},
	},
	AccessEntries: map[string]eks.AccessEntryArgs{
		"myRole": &eks.AccessEntryArgs{
			PrincipalArn: pulumi.String("arn:aws:iam::123456789012:role/eks-cluster-role"),
			Username:     pulumi.String("test-role"),
			KubernetesGroups: pulumi.StringArray{
				pulumi.String("test-group"),
			},
			Type: eks.AccessEntryTypeStandard,
		},
		"myUser": &eks.AccessEntryArgs{
			PrincipalArn: pulumi.String("arn:aws:iam::123456789012:user/test-user"),
			Username:     pulumi.String("test-user"),
			KubernetesGroups: pulumi.StringArray{
				pulumi.String("test-users"),
			},
			Type: eks.AccessEntryTypeStandard,
		},
	},
})
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
new Eks.Cluster("cluster", new()
{
    AuthenticationMode = Eks.AuthenticationMode.ApiAndConfigMap,
    RoleMappings = new[]
    {
        new Eks.Inputs.RoleMappingArgs
        {
            RoleArn = "arn:aws:iam::123456789012:role/eks-cluster-role",
            Username = "test-role",
            Groups = new[]
            {
                "test-group",
            },
        },
    },
    UserMappings = new[]
    {
        new Eks.Inputs.UserMappingArgs
        {
            UserArn = "arn:aws:iam::123456789012:user/test-user",
            Username = "test-user",
            Groups = new[]
            {
                "test-users",
            },
        },
    },
    AccessEntries = 
    {
        { "myRole", new Eks.Inputs.AccessEntryArgs
        {
            PrincipalArn = "arn:aws:iam::123456789012:role/eks-cluster-role",
            Username = "test-role",
            KubernetesGroups = new()
            {
                "test-group",
            },
            Type = Eks.AccessEntryType.Standard,
        } },
        { "myUser", new Eks.Inputs.AccessEntryArgs
        {
            PrincipalArn = "arn:aws:iam::123456789012:user/test-user",
            Username = "test-user",
            KubernetesGroups = new()
            {
                "test-users",
            },
            Type = Eks.AccessEntryType.Standard,
        } },
    },
});
```

{{% /choosable %}}

{{% choosable language java %}}

```java
new Cluster("cluster", ClusterArgs.builder()
    .authenticationMode(AuthenticationMode.ApiAndConfigMap)
    .roleMappings(RoleMappingArgs.builder()
        .roleArn("arn:aws:iam::123456789012:role/eks-cluster-role")
        .username("test-role")
        .groups("test-group")
        .build())
    .userMappings(UserMappingArgs.builder()
        .userArn("arn:aws:iam::123456789012:user/test-user")
        .username("test-user")
        .groups("test-users")
        .build())
    .accessEntries(Map.ofEntries(
        Map.entry("myRole", Map.ofEntries(
            Map.entry("principalArn", "arn:aws:iam::123456789012:role/eks-cluster-role"),
            Map.entry("username", "test-role"),
            Map.entry("kubernetesGroups", "test-group"),
            Map.entry("type", "STANDARD")
        )),
        Map.entry("myUser", Map.ofEntries(
            Map.entry("principalArn", "arn:aws:iam::123456789012:user/test-user"),
            Map.entry("username", "test-user"),
            Map.entry("kubernetesGroups", "test-users"),
            Map.entry("type", "STANDARD")
        ))
    ))
    .build());
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
eksCluster:
  type: eks:index:Cluster
  properties:
    authenticationMode: API_AND_CONFIG_MAP
    roleMappings:
      - roleArn: arn:aws:iam::123456789012:role/eks-cluster-role
        username: test-role
        groups:
          - test-group
    userMappings:
      - userArn: arn:aws:iam::123456789012:user/test-user
        username: test-user
        groups:
          - test-users
    accessEntries:
      myRole:
        principalArn: arn:aws:iam::123456789012:role/eks-cluster-role
        username: test-role
        kubernetesGroups:
          - test-group
        type: STANDARD
      myUser:
        principalArn: arn:aws:iam::123456789012:user/test-user
        username: test-user
        kubernetesGroups:
          - test-users
        type: STANDARD
```

{{% /choosable %}}

{{< /chooser >}}

#### instanceRoles
These map to Access Entries of type `EC2_LINUX`.

The following shows how you'd add access entries for the instance roles:
```ts
const cluster = new eks.Cluster(`${projectName}-cluster`, {
    ...
    authenticationMode: eks.AuthenticationMode.API_AND_CONFIG_MAP,
    instanceRoles: [
        ec2IamRole
    ],
    accessEntries: {
        myEC2InstanceRole: {
            principalArn: iamRole.arn,
            type: eks.AccessEntryType.EC2_LINUX
        }
    }
    ...
});
```

### 3. Confirm all entries of the `aws-auth` ConfigMap have corresponding access entries
You can either inspect the `aws-auth` ConfigMap in the `kube-system` namespace directly or use the [eksctl](https://eksctl.io/installation/) command line utility to analyze it by running `eksctl get iamidentitymapping --cluster my-cluster`.

### 4. Remove all entries of the `aws-auth` ConfigMap
After you've confirmed that all entries of the `aws-auth` ConfigMap have corresponding access entries you can remove these parameters:
- `roleMappings`
- `userMappings`
- `instanceRoles`

> Caution! This can cause disruptions if any of the parameters are not represented as Access Entries.

#### 5. Switch the Authentication Mode to `API`
Now that the authentication is fully driven by access entries you can switch to the `API` authentication mode. This will delete the `aws-auth` ConfigMap.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}

{{% choosable language typescript %}}

```ts
new eks.Cluster(`cluster`, {
    authenticationMode: eks.AuthenticationMode.Api,
});
```

{{% /choosable %}}

{{% choosable language python %}}

```python
eks.Cluster(
    "cluster",
    authentication_mode=eks.AuthenticationMode.API,
)
```

{{% /choosable %}}

{{% choosable language go %}}

```go
authMode := eks.AuthenticationModeApi
eks.NewCluster(ctx, "cluster", &eks.ClusterArgs{
    AuthenticationMode: &authMode,
})
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
new Eks.Cluster("cluster", new Eks.ClusterArgs
{
    AuthenticationMode = Eks.AuthenticationMode.Api,
});
```

{{% /choosable %}}

{{% choosable language java %}}

```java
Cluster("cluster", ClusterArgs.builder()
    .authenticationMode(AuthenticationMode.Api)
    .build());
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
cluster:
  type: eks:index:Cluster
  properties:
    authenticationMode: API
```

{{% /choosable %}}

{{< /chooser >}}
