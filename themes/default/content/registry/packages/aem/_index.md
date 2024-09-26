---
title: Adobe Experience Manager Provider
meta_desc: Provides an overview of the Adobe Experience Manager Provider for Pulumi.
layout: package
---

## Adobe Experience Manager Provider for Pulumi

This provider allows development teams to easily set up [Adobe Experience Manager](https://business.adobe.com/products/experience-manager/adobe-experience-manager.html) instances on virtual machines in the cloud (AWS, Azure, GCP, etc.) or bare metal machines.
It's based on the [AEM Compose](https://github.com/wttech/aemc) tool and aims to simplify the process of creating AEM environments without requiring deep DevOps knowledge.

## Example

{{< chooser language "typescript,python,go,csharp" >}}
{{% choosable language typescript %}}

```typescript
import * as aem from "@wttech/aem";
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const workspace = "aemc"
const env = pulumi.getStack()
const envType = "tf-minimal"
const host = "aem-single"
const dataDevice = "/dev/nvme1n1"
const dataDir = "/data"
const composeDir = `${dataDir}/aemc`

const tags = {
    "Workspace": workspace,
    "Env": env,
    "EnvType": envType,
    "Host": host,
    "Name": `${workspace}_${env}_${host}`,
}

const role = new aws.iam.Role("aem_ec2", {
    name: `${workspace}_${env}_aem_ec2`,
    assumeRolePolicy: JSON.stringify({
        "Version": "2012-10-17",
        "Statement": {
            "Effect": "Allow",
            "Principal": {"Service": "ec2.amazonaws.com"},
            "Action": "sts:AssumeRole"
        }
    }),
    tags: tags,
});

new aws.iam.RolePolicyAttachment("ssm", {
    role: role.name,
    policyArn: "arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore",
});

new aws.iam.RolePolicyAttachment("s3", {
    role: role.name,
    policyArn: "arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess",
});

const instanceProfile = new aws.iam.InstanceProfile("aem_ec2", {
    name: `${workspace}_${env}_aem_ec2`,
    role: role.name,
    tags: tags,
});

const instance = new aws.ec2.Instance("aem_single", {
    ami: "ami-043e06a423cbdca17", // RHEL 8
    instanceType: "m5.xlarge",
    iamInstanceProfile: instanceProfile.name,
    tags: tags,
    userData: `#!/bin/bash
sudo dnf install -y https://s3.amazonaws.com/ec2-downloads-windows/SSMAgent/latest/linux_amd64/amazon-ssm-agent.rpm`,
});

const volume = new aws.ebs.Volume("aem_single_data", {
    availabilityZone: instance.availabilityZone,
    size: 128,
    type: "gp2",
    tags: tags,
});

const volumeAttachment = new aws.ec2.VolumeAttachment("aem_single_data", {
    deviceName: "/dev/xvdf",
    volumeId: volume.id,
    instanceId: instance.id,
});

const aemInstance = new aem.compose.Instance("aem_instance", {
    client: {
        type: "aws-ssm",
        settings: {
            instance_id: instance.id,
        },
    },
    system: {
        data_dir: composeDir,
        bootstrap: {
            inline: [
                `sudo mkfs -t ext4 ${dataDevice}`,
                `sudo mkdir -p ${dataDir}`,
                `sudo mount ${dataDevice} ${dataDir}`,
                `echo '${dataDevice} ${dataDir} ext4 defaults 0 0' | sudo tee -a /etc/fstab`,
                "sudo yum install -y unzip",
                "curl 'https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip' -o 'awscliv2.zip'",
                "unzip -q awscliv2.zip",
                "sudo ./aws/install --update",
            ],
        },
    },
    compose: {
        create: {
            inline: [
                `mkdir -p '${composeDir}/aem/home/lib'`,
                `aws s3 cp --recursive --no-progress 's3://aemc/instance/classic/' '${composeDir}/aem/home/lib'`,
                "sh aemw instance init",
                "sh aemw instance create",
            ],
        },
        configure: {
            inline: [
                "sh aemw osgi config save --pid 'org.apache.sling.jcr.davex.impl.servlets.SlingDavExServlet' --input-string 'alias: /crx/server'",
                "sh aemw repl agent setup -A --location 'author' --name 'publish' --input-string '{enabled: true, transportUri: \"http://localhost:4503/bin/receive?sling:authRequestLogin=1\", transportUser: admin, transportPassword: admin, userId: admin}'",
                "sh aemw package deploy --file 'aem/home/lib/aem-service-pkg-6.5.*.0.zip'",
            ],
        },
    }
}, {dependsOn: [instance, volumeAttachment]});

export const output = {
    instanceIp: instance.publicIp,
    aemInstances: aemInstance.instances,
};
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
import pulumi_aws as aws
import wttech_aem as aem

workspace = "aemc"
env = pulumi.get_stack()
envType = "tf-minimal"
host = "aem-single"
dataDevice = "/dev/nvme1n1"
dataDir = "/data"
composeDir = f"{dataDir}/aemc"

tags = {
    "Workspace": workspace,
    "Env": env,
    "EnvType": envType,
    "Host": host,
    "Name": f"{workspace}_{env}_{host}",
}

role = aws.iam.Role("aem_ec2",
                    name=f"{workspace}_{env}_aem_ec2",
                    assume_role_policy="""{
    "Version": "2012-10-17",
    "Statement": {
        "Effect": "Allow",
        "Principal": {"Service": "ec2.amazonaws.com"},
        "Action": "sts:AssumeRole"
    }
}""",
                    tags=tags,
                    )

aws.iam.RolePolicyAttachment("ssm",
                             role=role.name,
                             policy_arn="arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore",
                             )

aws.iam.RolePolicyAttachment("s3",
                             role=role.name,
                             policy_arn="arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess",
                             )

instanceProfile = aws.iam.InstanceProfile("aem_ec2",
                                          name=f"{workspace}_{env}_aem_ec2",
                                          role=role.name,
                                          tags=tags,
                                          )

instance = aws.ec2.Instance("aem_single",
                            ami="ami-043e06a423cbdca17",  # RHEL 8
                            instance_type="m5.xlarge",
                            iam_instance_profile=instanceProfile.name,
                            tags=tags,
                            user_data=pulumi.String("""#!/bin/bash
sudo dnf install -y https://s3.amazonaws.com/ec2-downloads-windows/SSMAgent/latest/linux_amd64/amazon-ssm-agent.rpm"""),
                            )

volume = aws.ebs.Volume("aem_single_data",
                        availability_zone=instance.availability_zone,
                        size=128,
                        type="gp2",
                        tags=tags,
                        )

volumeAttachment = aws.ec2.VolumeAttachment("aem_single_data",
                                            device_name="/dev/xvdf",
                                            volume_id=volume.id,
                                            instance_id=instance.id,
                                            )

aemInstance = aem.compose.Instance("aem_instance",
                                   client={
                                       "type": "aws-ssm",
                                       "settings": {
                                           "instance_id": instance.id,
                                       },
                                   },
                                   system={
                                       "data_dir": composeDir,
                                       "bootstrap": {
                                           "inline": [
                                               f"sudo mkfs -t ext4 {dataDevice}",
                                               f"sudo mkdir -p {dataDir}",
                                               f"sudo mount {dataDevice} {dataDir}",
                                               f"echo '{dataDevice} {dataDir} ext4 defaults 0 0' | sudo tee -a /etc/fstab",
                                               "sudo yum install -y unzip",
                                               "curl 'https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip' -o 'awscliv2.zip'",
                                               "unzip -q awscliv2.zip",
                                               "sudo ./aws/install --update",
                                           ],
                                       },
                                   },
                                   compose={
                                       "create": {
                                           "inline": [
                                               f"mkdir -p '{composeDir}/aem/home/lib'",
                                               f"aws s3 cp --recursive --no-progress 's3://aemc/instance/classic/' '{composeDir}/aem/home/lib'",
                                               "sh aemw instance init",
                                               "sh aemw instance create",
                                           ],
                                       },
                                       "configure": {
                                           "inline": [
                                               "sh aemw osgi config save --pid 'org.apache.sling.jcr.davex.impl.servlets.SlingDavExServlet' --input-string 'alias: /crx/server'",
                                               "sh aemw repl agent setup -A --location 'author' --name 'publish' --input-string '{enabled: true, transportUri: \"http://localhost:4503/bin/receive?sling:authRequestLogin=1\", transportUser: admin, transportPassword: admin, userId: admin}'",
                                               "sh aemw package deploy --file 'aem/home/lib/aem-service-pkg-6.5.*.0.zip'",
                                           ],
                                       },
                                   },
                                   depends_on=[instance, volumeAttachment],
                                   )

pulumi.export("output", {
    "instanceIp": instance.public_ip,
    "aemInstances": aemInstance.instances,
})
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
    "fmt"
    "github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ebs"
    "github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
    "github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
    "github.com/wttech/pulumi-aem/sdk/go/aem/compose"
)

func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
        workspace := "aemc"
        env := ctx.Stack()
        envType := "tf-minimal"
        host := "aem-single"
        dataDevice := "/dev/nvme1n1"
        dataDir := "/data"
        composeDir := fmt.Sprintf("%s/aemc", dataDir)

        tags := pulumi.StringMap{
            "Workspace": pulumi.String(workspace),
            "Env":       pulumi.String(env),
            "EnvType":   pulumi.String(envType),
            "Host":      pulumi.String(host),
            "Name":      pulumi.Sprintf("%s_%s_%s", workspace, env, host),
        }

        role, err := iam.NewRole(ctx, "aem_ec2", &iam.RoleArgs{
            Name: pulumi.Sprintf("%s_%s_aem_ec2", workspace, env),
            AssumeRolePolicy: pulumi.String(`{
    "Version": "2012-10-17",
    "Statement": {
        "Effect": "Allow",
        "Principal": {"Service": "ec2.amazonaws.com"},
        "Action": "sts:AssumeRole"
    }
}`),
            Tags: tags,
        })
        if err != nil {
            return err
        }

        _, err = iam.NewRolePolicyAttachment(ctx, "ssm", &iam.RolePolicyAttachmentArgs{
            Role:      role.Name,
            PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore"),
        })
        if err != nil {
            return err
        }

        _, err = iam.NewRolePolicyAttachment(ctx, "s3", &iam.RolePolicyAttachmentArgs{
            Role:      role.Name,
            PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess"),
        })
        if err != nil {
            return err
        }

        instanceProfile, err := iam.NewInstanceProfile(ctx, "aem_ec2", &iam.InstanceProfileArgs{
            Name: pulumi.Sprintf("%s_%s_aem_ec2", workspace, env),
            Role: role.Name,
            Tags: tags,
        })
        if err != nil {
            return err
        }

        instance, err := ec2.NewInstance(ctx, "aem_single", &ec2.InstanceArgs{
            Ami:                pulumi.String("ami-043e06a423cbdca17"), // RHEL 8
            InstanceType:       pulumi.String("m5.xlarge"),
            IamInstanceProfile: instanceProfile.Name,
            Tags:               tags,
            UserData: pulumi.String(`#!/bin/bash
sudo dnf install -y https://s3.amazonaws.com/ec2-downloads-windows/SSMAgent/latest/linux_amd64/amazon-ssm-agent.rpm`),
        })
        if err != nil {
            return err
        }

        volume, err := ebs.NewVolume(ctx, "aem_single_data", &ebs.VolumeArgs{
            AvailabilityZone: instance.AvailabilityZone,
            Size:             pulumi.Int(128),
            Type:             pulumi.String("gp2"),
            Tags:             tags,
        })
        if err != nil {
            return err
        }

        volumeAttachment, err := ec2.NewVolumeAttachment(ctx, "aem_single_data", &ec2.VolumeAttachmentArgs{
            DeviceName: pulumi.String("/dev/xvdf"),
            VolumeId:   volume.ID(),
            InstanceId: instance.ID(),
        })
        if err != nil {
            return err
        }

        aemInstance, err := compose.NewInstance(ctx, "aem_instance", &compose.InstanceArgs{
            Client: compose.ClientArgs{
                Type: pulumi.String("aws-ssm"),
                Settings: pulumi.StringMap{
                    "instance_id": instance.ID(),
                },
            },
            System: compose.SystemArgs{
                Data_dir: pulumi.String(composeDir),
                Bootstrap: compose.InstanceScriptArgs{
                    Inline: pulumi.StringArray{
                        pulumi.Sprintf("sudo mkfs -t ext4 %s", dataDevice),
                        pulumi.Sprintf("sudo mkdir -p %s", dataDir),
                        pulumi.Sprintf("sudo mount %s %s", dataDevice, dataDir),
                        pulumi.Sprintf("echo '%s %s ext4 defaults 0 0' | sudo tee -a /etc/fstab", dataDevice, dataDir),
                        pulumi.String("sudo yum install -y unzip"),
                        pulumi.String("curl 'https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip' -o 'awscliv2.zip'"),
                        pulumi.String("unzip -q awscliv2.zip"),
                        pulumi.String("sudo ./aws/install --update"),
                    },
                },
            },
            Compose: compose.ComposeArgs{
                Create: compose.InstanceScriptArgs{
                    Inline: pulumi.StringArray{
                        pulumi.Sprintf("mkdir -p '%s/aem/home/lib'", composeDir),
                        pulumi.Sprintf("aws s3 cp --recursive --no-progress 's3://aemc/instance/classic/' '%s/aem/home/lib'", composeDir),
                        pulumi.String("sh aemw instance init"),
                        pulumi.String("sh aemw instance create"),
                    },
                },
                Configure: compose.InstanceScriptArgs{
                    Inline: pulumi.StringArray{
                        pulumi.String("sh aemw osgi config save --pid 'org.apache.sling.jcr.davex.impl.servlets.SlingDavExServlet' --input-string 'alias: /crx/server'"),
                        pulumi.String("sh aemw repl agent setup -A --location 'author' --name 'publish' --input-string '{enabled: true, transportUri: \"http://localhost:4503/bin/receive?sling:authRequestLogin=1\", transportUser: admin, transportPassword: admin, userId: admin}'"),
                        pulumi.String("sh aemw package deploy --file 'aem/home/lib/aem-service-pkg-6.5.*.0.zip'"),
                    },
                },
            },
        }, pulumi.DependsOn([]pulumi.Resource{instance, volumeAttachment}))
        if err != nil {
            return err
        }

        ctx.Export("output", pulumi.Map{
            "instanceIp":   instance.PublicIp,
            "aemInstances": aemInstance.Instances,
        })
        return nil
    })
}
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using System.IO;
using Pulumi.Aws.Ec2;
using Pulumi.Aws.Ebs;
using Pulumi.Aws.Iam;
using Aem = WTTech.Aem;

return await Deployment.RunAsync(() =>
{
    var workspace = "aemc";
    var env = Deployment.Instance.StackName;
    var envType = "tf-minimal";
    var host = "aem-single";
    var dataDevice = "/dev/nvme1n1";
    var dataDir = "/data";
    var composeDir = $"{dataDir}/aemc";

    var tags = new InputMap<string>
    {
        { "Workspace", workspace },
        { "Env", env },
        { "EnvType", envType },
        { "Host", host },
        { "Name", $"{workspace}_{env}_{host}" },
    };

    var role = new Role("aem_ec2", new RoleArgs
    {
        Name = $"{workspace}_{env}_aem_ec2",
        AssumeRolePolicy = @"{
    ""Version"": ""2012-10-17"",
    ""Statement"": {
        ""Effect"": ""Allow"",
        ""Principal"": {""Service"": ""ec2.amazonaws.com""},
        ""Action"": ""sts:AssumeRole""
    }
}",
        Tags = tags,
    });

    var ssmPolicyAttachment = new RolePolicyAttachment("ssm", new RolePolicyAttachmentArgs
    {
        Role = role.Name,
        PolicyArn = "arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore",
    });

    var s3PolicyAttachment = new RolePolicyAttachment("s3", new RolePolicyAttachmentArgs
    {
        Role = role.Name,
        PolicyArn = "arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess",
    });

    var instanceProfile = new InstanceProfile("aem_ec2", new InstanceProfileArgs
    {
        Name = $"{workspace}_{env}_aem_ec2",
        Role = role.Name,
        Tags = tags,
    });

    var instance = new Instance("aem_single", new InstanceArgs
    {
        Ami = "ami-043e06a423cbdca17", // RHEL 8
        InstanceType = "m5.xlarge",
        IamInstanceProfile = instanceProfile.Name,
        Tags = tags,
        UserData = @"#!/bin/bash
sudo dnf install -y https://s3.amazonaws.com/ec2-downloads-windows/SSMAgent/latest/linux_amd64/amazon-ssm-agent.rpm",
    });

    var volume = new Volume("aem_single_data", new VolumeArgs
    {
        AvailabilityZone = instance.AvailabilityZone,
        Size = 128,
        Type = "gp2",
        Tags = tags,
    });

    var volumeAttachment = new VolumeAttachment("aem_single_data", new VolumeAttachmentArgs
    {
        DeviceName = "/dev/xvdf",
        VolumeId = volume.Id,
        InstanceId = instance.Id,
    });

    var aemInstance = new Aem.Compose.Instance("aem_instance", new()
    {
        Client = new Aem.Compose.Inputs.ClientArgs
        {
            Type = "aws-ssm",
            Settings = new InputMap<string>
            {
                { "instance_id", instance.Id },
            },
        },
        System = new Aem.Compose.Inputs.SystemArgs
        {
            Data_dir = composeDir,
            Bootstrap = new InstanceScriptArgs
            {
                Inline = new InputList<string>
                {
                    $"sudo mkfs -t ext4 {dataDevice}",
                    $"sudo mkdir -p {dataDir}",
                    $"sudo mount {dataDevice} {dataDir}",
                    $"echo '{dataDevice} {dataDir} ext4 defaults 0 0' | sudo tee -a /etc/fstab",
                    "sudo yum install -y unzip",
                    "curl 'https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip' -o 'awscliv2.zip'",
                    "unzip -q awscliv2.zip",
                    "sudo ./aws/install --update",
                },
            },
        },
        Compose = new Aem.Compose.Inputs.ComposeArgs
        {
            Create = new InstanceScriptArgs
            {
                Inline = new InputList<string>
                {
                    $"mkdir -p '{composeDir}/aem/home/lib'",
                    $"aws s3 cp --recursive --no-progress 's3://aemc/instance/classic/' '{composeDir}/aem/home/lib'",
                    "sh aemw instance init",
                    "sh aemw instance create",
                },
            },
            Configure = new InstanceScriptArgs
            {
                Inline = new InputList<string>
                {
                    "sh aemw osgi config save --pid 'org.apache.sling.jcr.davex.impl.servlets.SlingDavExServlet' --input-string 'alias: /crx/server'",
                    "sh aemw repl agent setup -A --location 'author' --name 'publish' --input-string '{enabled: true, transportUri: \"http://localhost:4503/bin/receive?sling:authRequestLogin=1\", transportUser: admin, transportPassword: admin, userId: admin}'",
                    "sh aemw package deploy --file 'aem/home/lib/aem-service-pkg-6.5.*.0.zip'",
                },
            },
        },
    }, new CustomResourceOptions
    {
        DependsOn = { instance, volumeAttachment },
    });

    return new Dictionary<string, object?>
    {
        ["output"] =
        {
            { "instanceIp", instance.PublicIp },
            { "aemInstances", aemInstance.Instances },
        },
    };
});
```

{{% /choosable %}}
{{< /chooser >}}
