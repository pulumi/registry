---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-docker-build/v0.0.15/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi/pulumi-docker-build/blob/v0.0.15/docs/_index.md
title: Docker Build
meta_desc: Provides an overview of the Docker Build Provider for Pulumi.
layout: package
---

The Docker Build provider leverages [buildx and BuildKit](https://docs.docker.com/build/architecture/) to build modern Docker images with Pulumi.

Not to be confused with the earlier
[Docker](../docker) provider, which is still
appropriate for managing resources unrelated to building images.

| Provider               | Use cases                                                               |
| ----------------       | ----------------------------------------------------------------------- |
| `@pulumi/docker-build` | Anything related to building images with `docker build`.                |
| `@pulumi/docker`       | Everything else -- including running containers and creating networks.  |

## Example

If your Pulumi program has a directory called `app` alongside it, containing a
file named "Dockerfile" (which can be as simple as `FROM alpine` for the
purpose of example), then the code below shows how to build a multi-platform
image, publish it to a remote AWS ECR registry, and use an [inline
cache](https://docs.docker.com/build/cache/backends/inline/) to speed up
subsequent builds.

{{< chooser language "typescript,python,csharp,go,yaml,java" / >}}

{{% choosable language typescript %}}

```typescript
import * as aws from "@pulumi/aws";
import * as docker_build from "@pulumi/docker-build";

// Create an ECR repository for pushing.
const ecrRepository = new aws.ecr.Repository("ecr-repository", {});

// Grab auth credentials for ECR.
const authToken = aws.ecr.getAuthorizationTokenOutput({
    registryId: ecrRepository.registryId,
});

// Build and push an image to ECR with inline caching.
const myImage = new docker_build.Image("my-image", {
    // Tag our image with our ECR repository's address.
    tags: [pulumi.interpolate`${ecrRepository.repositoryUrl}:latest`],
    context: {
        location: "./app",
    },
    // Use the pushed image as a cache source.
    cacheFrom: [{
        registry: {
            ref: pulumi.interpolate`${ecrRepository.repositoryUrl}:latest`,
        },
    }],
    // Include an inline cache with our pushed image.
    cacheTo: [{
        inline: {},
    }],
    // Build a multi-platform image manifest for ARM and AMD.
    platforms: [
        "linux/amd64",
        "linux/arm64",
    ],
    // Push the final result to ECR.
    push: true,
    // Provide our ECR credentials.
    registries: [{
        address: ecrRepository.repositoryUrl,
        password: authToken.password,
        username: authToken.userName,
    }],
});

// Export a ref for the pushed images so we can deploy it.
export const ref = myImage.ref;
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_aws as aws
import pulumi_docker_build as docker_build

# Create an ECR repository for pushing.
ecr_repository = aws.ecr.Repository("ecr-repository")

# Grab auth credentials for ECR.
auth_token = aws.ecr.get_authorization_token_output(registry_id=ecr_repository.registry_id)

# Build and push an image to ECR with inline caching.
my_image = docker_build.Image("my-image",
    # Tag our image with our ECR repository's address.
    tags=[ecr_repository.repository_url.apply(lambda repository_url: f"{repository_url}:latest")],
    context=docker_build.BuildContextArgs(
        location="./app",
    ),
    # Use the pushed image as a cache source.
    cache_from=[docker_build.CacheFromArgs(
        registry=docker_build.CacheFromRegistryArgs(
            ref=ecr_repository.repository_url.apply(lambda repository_url: f"{repository_url}:latest"),
        ),
    )],
    # Include an inline cache with our pushed image.
    cache_to=[docker_build.CacheToArgs(
        inline=docker_build.CacheToInlineArgs(),
    )],
    # Build a multi-platform image manifest for ARM and AMD.
    platforms=[
        docker_build.Platform.LINUX_AMD64,
        docker_build.Platform.LINUX_ARM64,
    ],
    # Push the final result to ECR.
    push=True,
    # Provide our ECR credentials.
    registries=[docker_build.RegistryArgs(
        address=ecr_repository.repository_url,
        password=auth_token.password,
        username=auth_token.user_name,
    )],
)

# Export a ref for the pushed images so we can deploy it.
pulumi.export("ref", my_image.ref)
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;
using DockerBuild = Pulumi.DockerBuild;

return await Deployment.RunAsync(() =>
{
    // Create an ECR repository for pushing.
    var ecrRepository = new Aws.Ecr.Repository("ecr-repository");

    // Grab auth credentials for ECR.
    var authToken = Aws.Ecr.GetAuthorizationToken.Invoke(new()
    {
        RegistryId = ecrRepository.RegistryId,
    });

    // Build and push an image to ECR with inline caching.
    var myImage = new DockerBuild.Image("my-image", new()
    {
        // Tag our image with our ECR repository's address.
        Tags = new[]
        {
            ecrRepository.RepositoryUrl.Apply(repositoryUrl => $"{repositoryUrl}:latest"),
        },
        Context = new DockerBuild.Inputs.BuildContextArgs
        {
            Location = "./app",
        },
        // Use the pushed image as a cache source.
        CacheFrom = new[]
        {
            new DockerBuild.Inputs.CacheFromArgs
            {
                Registry = new DockerBuild.Inputs.CacheFromRegistryArgs
                {
                    Ref = ecrRepository.RepositoryUrl.Apply(repositoryUrl => $"{repositoryUrl}:latest"),
                },
            },
        },
        // Include an inline cache with our pushed image.
        CacheTo = new[]
        {
            new DockerBuild.Inputs.CacheToArgs
            {
                Inline = null,
            },
        },
        // Build a multi-platform image manifest for ARM and AMD.
        Platforms = new[]
        {
            DockerBuild.Platform.Linux_amd64,
            DockerBuild.Platform.Linux_arm64,
        },
        // Push the final result to ECR.
        Push = true,
        // Provide our ECR credentials.
        Registries = new[]
        {
            new DockerBuild.Inputs.RegistryArgs
            {
                Address = ecrRepository.RepositoryUrl,
                Password = authToken.Apply(getAuthorizationTokenResult => getAuthorizationTokenResult.Password),
                Username = authToken.Apply(getAuthorizationTokenResult => getAuthorizationTokenResult.UserName),
            },
        },
    });

    // Export a ref for the pushed images so we can deploy it.
    return new Dictionary<string, object?>
    {
        ["ref"] = myImage.Ref,
    };
});

```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
    "fmt"

    "github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecr"
    "github.com/pulumi/pulumi-docker-build/sdk/go/dockerbuild"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
        // Create an ECR repository for pushing.
        ecrRepository, err := ecr.NewRepository(ctx, "ecr-repository", nil)
        if err != nil {
            return err
        }

        // Grab auth credentials for ECR.
        authToken := ecr.GetAuthorizationTokenOutput(ctx, ecr.GetAuthorizationTokenOutputArgs{
            RegistryId: ecrRepository.RegistryId,
        }, nil)

        // Build and push an image to ECR with inline caching.
        myImage, err := dockerbuild.NewImage(ctx, "my-image", &dockerbuild.ImageArgs{
            // Tag our image with our ECR repository's address.
            Tags: pulumi.StringArray{
                ecrRepository.RepositoryUrl.ApplyT(func(repositoryUrl string) (string, error) {
                    return fmt.Sprintf("%v:latest", repositoryUrl), nil
                }).(pulumi.StringOutput),
            },
            Context: &dockerbuild.BuildContextArgs{
                Location: pulumi.String("./app"),
            },
            // Use the pushed image as a cache source.
            CacheFrom: dockerbuild.CacheFromArray{
                &dockerbuild.CacheFromArgs{
                    Registry: &dockerbuild.CacheFromRegistryArgs{
                        Ref: ecrRepository.RepositoryUrl.ApplyT(func(repositoryUrl string) (string, error) {
                            return fmt.Sprintf("%v:latest", repositoryUrl), nil
                        }).(pulumi.StringOutput),
                    },
                },
            },
            // Include an inline cache with our pushed image.
            CacheTo: dockerbuild.CacheToArray{
                &dockerbuild.CacheToArgs{
                    Inline: nil,
                },
            },
            // Build a multi-platform image manifest for ARM and AMD.
            Platforms: dockerbuild.PlatformArray{
                dockerbuild.Platform_Linux_amd64,
                dockerbuild.Platform_Linux_arm64,
            },
            // Push the final result to ECR.
            Push: pulumi.Bool(true),
            // Provide our ECR credentials.
            Registries: dockerbuild.RegistryArray{
                &dockerbuild.RegistryArgs{
                    Address: ecrRepository.RepositoryUrl,
                    Password: authToken.ApplyT(func(authToken ecr.GetAuthorizationTokenResult) (*string, error) {
                        return &authToken.Password, nil
                    }).(pulumi.StringPtrOutput),
                    Username: authToken.ApplyT(func(authToken ecr.GetAuthorizationTokenResult) (*string, error) {
                        return &authToken.UserName, nil
                    }).(pulumi.StringPtrOutput),
                },
            },
        })
        if err != nil {
            return err
        }

        // Export a ref for the pushed images so we can deploy it.
        ctx.Export("ref", myImage.Ref)
        return nil
    })
}
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
description: Push to AWS ECR with caching
name: ecr
outputs:
    ref: ${my-image.ref}
resources:
    # Create an ECR repository for pushing.
    ecr-repository:
        type: aws:ecr:Repository

    # Build and push an image to ECR with inline caching.
    my-image:
        type: docker-build:Image
        properties:
            # Tag our image with our ECR repository's address.
            tags:
                - ${ecr-repository.repositoryUrl}:latest
            context:
                location: ./app
            # Use the pushed image as a cache source.
            cacheFrom:
                - registry:
                    ref: ${ecr-repository.repositoryUrl}:latest
            # Include an inline cache with our pushed image.
            cacheTo:
                - inline: {}
            # Build a multi-platform image manifest for ARM and AMD.
            platforms:
                - linux/amd64
                - linux/arm64
            # Push the final result to ECR.
            push: true
            # Provide our ECR credentials.
            registries:
                - address: ${ecr-repository.repositoryUrl}
                  password: ${auth-token.password}
                  username: ${auth-token.userName}

runtime: yaml
variables:
    auth-token:
        # Grab auth credentials for ECR.
        fn::aws:ecr:getAuthorizationToken:
            registryId: ${ecr-repository.registryId}
```

{{% /choosable %}}

{{% choosable language java %}}

```java
package myapp;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.ecr.Repository;
import com.pulumi.aws.ecr.EcrFunctions;
import com.pulumi.aws.ecr.inputs.GetAuthorizationTokenArgs;
import com.pulumi.dockerbuild.Image;
import com.pulumi.dockerbuild.ImageArgs;
import com.pulumi.dockerbuild.inputs.CacheFromArgs;
import com.pulumi.dockerbuild.inputs.CacheFromRegistryArgs;
import com.pulumi.dockerbuild.inputs.CacheToArgs;
import com.pulumi.dockerbuild.inputs.CacheToInlineArgs;
import com.pulumi.dockerbuild.inputs.BuildContextArgs;
import com.pulumi.dockerbuild.inputs.RegistryArgs;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        // Create an ECR repository for pushing.
        var ecrRepository = new Repository("ecrRepository");

        // Grab auth credentials for ECR.
        final var authToken = EcrFunctions.getAuthorizationToken(GetAuthorizationTokenArgs.builder()
            .registryId(ecrRepository.registryId())
            .build());

        // Build and push an image to ECR with inline caching.
        var myImage = new Image("myImage", ImageArgs.builder()
            // Tag our image with our ECR repository's address.
            .tags(ecrRepository.repositoryUrl().applyValue(repositoryUrl -> String.format("%s:latest", repositoryUrl)))
            .context(BuildContextArgs.builder()
                .location("./app")
                .build())
            // Use the pushed image as a cache source.
            .cacheFrom(CacheFromArgs.builder()
                .registry(CacheFromRegistryArgs.builder()
                    .ref(ecrRepository.repositoryUrl().applyValue(repositoryUrl -> String.format("%s:latest", repositoryUrl)))
                    .build())
                .build())
            // Include an inline cache with our pushed image.
            .cacheTo(CacheToArgs.builder()
                .inline()
                .build())
            // Build a multi-platform image manifest for ARM and AMD.
            .platforms(
                "linux/amd64",
                "linux/arm64")
            // Push the final result to ECR.
            .push(true)
            // Provide our ECR credentials.
            .registries(RegistryArgs.builder()
                .address(ecrRepository.repositoryUrl())
                .password(authToken.applyValue(getAuthorizationTokenResult -> getAuthorizationTokenResult).applyValue(authToken -> authToken.applyValue(getAuthorizationTokenResult -> getAuthorizationTokenResult.password())))
                .username(authToken.applyValue(getAuthorizationTokenResult -> getAuthorizationTokenResult).applyValue(authToken -> authToken.applyValue(getAuthorizationTokenResult -> getAuthorizationTokenResult.userName())))
                .build())
            .build());

        ctx.export("ref", myImage.ref());
    }
}
```

{{% /choosable %}}

{{< /chooser >}}
