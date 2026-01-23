---
# WARNING: this file was fetched from https://raw.githubusercontent.com/runpod/pulumi-runpod-native/v1.9.99/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/runpod/pulumi-runpod-native/blob/v1.9.99/docs/_index.md
title: RunPod
meta_desc: Provides an overview of the RunPod Provider for Pulumi.
layout: package
---

The RunPod provider for Pulumi can be used to provision [RunPod](https://www.runpod.io) resources. The RunPod provider must be configured with RunPod's API keys to deploy and update resources in RunPod.

## Config

To begin with, please set your RunPod API key using Pulumi.

```bash
pulumi config set --secret runpod:token
```

## Note

Please make sure that you are inside the Python virtual environment created by Pulumi when using the Python SDK.

## Example

{{< chooser language "typescript,go,python,yaml" />}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as runpod from "@runpod-infra/pulumi";

const myTemplate = new runpod.Template("testTemplate", {
  containerDiskInGb: 5,
  dockerArgs: "python handler.py",
  env: [
    {
      key: "key1",
      value: "value1",
    },
    {
      key: "key2",
      value: "value2",
    },
  ],
  imageName: "runpod/serverless-hello-world:latest",
  isServerless: true,
  name: "Testing Pulumi V1",
  readme: "## Hello, World!",
  volumeInGb: 0,
});

const testNetworkStorage = new runpod.NetworkStorage("testNetworkStorage", {
  name: "testStorage1",
  size: 10,
  dataCenterId: "US-OR-1",
});

const myRandomPod = new runpod.Pod("myRandomPod", {
  cloudType: "ALL",
  networkVolumeId: testNetworkStorage.networkStorage.apply(
    // @ts-ignore
    (networkStorage) => networkStorage.id
  ),
  gpuCount: 1,
  volumeInGb: 50,
  containerDiskInGb: 50,
  minVcpuCount: 2,
  minMemoryInGb: 15,
  gpuTypeId: "NVIDIA GeForce RTX 4090",
  name: "RunPod Pytorch",
  imageName: "runpod/pytorch:latest",
  dockerArgs: "",
  ports: "8888/http",
  volumeMountPath: "/workspace",
  env: [
    {
      key: "JUPYTER_PASSWORD",
      value: "rns1hunbsstltcpad22d",
    },
  ],
});

const myRandomEndpoint = new runpod.Endpoint("myRandomEndpoint", {
  gpuIds: "AMPERE_16,AMPERE_24,-NVIDIA L4",
  idleTimeout: 100,
  locations: "CA-MTL-2,CA-MTL-3,EU-RO-1,US-CA-1,US-GA-1,US-KS-2,US-OR-1,CA-MTL-1,US-TX-3,EUR-IS-1,EUR-IS-2,SEA-SG-1",
  name: "myRandomEndpoint",
  networkVolumeId: testNetworkStorage.networkStorage.apply(
    // @ts-ignore
    (networkStorage) => networkStorage.id
  ),
  scalerType: 'REQUEST_COUNT',
  scalerValue: 2,
  templateId: myTemplate.template.apply(t => t.id),
  workersMax: 2,
  workersMin: 1,
})

export const template = {
  value: myTemplate.template,
};

export const endpoint = {
  value: myRandomEndpoint.endpoint,
};

export const pod = {
  value: myRandomPod.pod,
};

export const networkStorage = {
  value: testNetworkStorage.networkStorage,
};
```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/runpod/pulumi-runpod-native/sdk/go/runpod"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testNetworkStorage, err := runpod.NewNetworkStorage(ctx, "testNetworkStorage", &runpod.NetworkStorageArgs{
			Name:         pulumi.String("testStorage1"),
			Size:         pulumi.Int(5),
			DataCenterId: pulumi.String("EU-RO-1"),
		})
		if err != nil {
			return err
		}
		myRandomPod, err := runpod.NewPod(ctx, "myRandomPod", &runpod.PodArgs{
			CloudType: pulumi.String("ALL"),
			NetworkVolumeId: testNetworkStorage.NetworkStorage.ApplyT(func(networkStorage runpod.NetworkStorageType) (*string, error) {
				return &networkStorage.Id, nil
			}).(pulumi.StringPtrOutput),
			GpuCount:          pulumi.Int(1),
			VolumeInGb:        pulumi.Int(50),
			ContainerDiskInGb: pulumi.Int(50),
			MinVcpuCount:      pulumi.Int(2),
			MinMemoryInGb:     pulumi.Int(15),
			GpuTypeId:         pulumi.String("NVIDIA GeForce RTX 4090"),
			Name:              pulumi.String("RunPod Pytorch"),
			ImageName:         pulumi.String("runpod/pytorch"),
			DockerArgs:        pulumi.String(""),
			Ports:             pulumi.String("8888/http"),
			VolumeMountPath:   pulumi.String("/workspace"),
			Env: runpod.PodEnvArray{
				&runpod.PodEnvArgs{
					Key:   pulumi.String("JUPYTER_PASSWORD"),
					Value: pulumi.String("rns1hunbsstltcpad22d"),
				},
			},
		})
		if err != nil {
			return err
		}
		myTemplate, err := runpod.NewTemplate(ctx, "myTemplate", &runpod.TemplateArgs{
			ContainerDiskInGb:       pulumi.Int(5),
			DockerArgs:              pulumi.String("python3 -m http.server 8080"),
		    Env: runpod.PodEnvArray{
				&runpod.PodEnvArgs{
					Key:   pulumi.String("JUPYTER_PASSWORD"),
					Value: pulumi.String("rns1hunbsstltcpad22d"),
				},
			},
			ImageName:       pulumi.String("runpod/serverless-hello-world:latest"),
			IsServerless:    pulumi.Bool(true),
			Name:            pulumi.String("Testing Pulumi V1"),
			Readme:          pulumi.String("## Hello, World!"),
			VolumeInGb:      pulumi.Int(0),
		})

		if err != nil {
			return err
		}

		myEndpoint, err := runpod.NewEndpoint(ctx, "myEndpoint", &runpod.EndpointArgs{
			GpuIds:         pulumi.String("AMPERE_16"),
			Name:           pulumi.String("Pulumi Endpoint Test V2 -fb"),
			TemplateId:     myTemplate.Template.Id(),
			WorkersMax:     pulumi.Int(2),
			WorkersMin:     pulumi.Int(1),
			IdleTimeout:    pulumi.Int(6),
			Locations:      pulumi.String("EU-RO-1"),
			NetworkVolumeId: testNetworkStorage.NetworkStorage.Id(),
			ScalerType:     pulumi.String("QUEUE_DELAY"),
			ScalerValue:    pulumi.Int(4),
		})

		if err != nil {
			return err
		}

		ctx.Export("pod", myRandomPod.Pod.Id())
		ctx.Export("networkStorage", testNetworkStorage.NetworkStorage.Id())
		ctx.Export("template", myTemplate.Template.Id())
		ctx.Export("endpoint", myEndpoint.Endpoint.Id())
		return nil
	})
}
```

{{% /choosable %}}

{{% choosable language python %}}

```
  source venv/bin/activate
```

```python
import pulumi
import runpodinfra as runpod
from loguru import logger
import json

def fetch_id(a):
    if type(a) == runpod.outputs.NetworkStorage:
        return a.id
    else:
        return a

def fetch_template_id(a):
    if type(a) == runpod.outputs.Template:
        return a.id
    else:
        return a

try:
    test_network_storage = runpod.NetworkStorage("testNetworkStorage",
        name="testStorage1",
        size=5,
        data_center_id="EU-RO-1")

    my_random_pod = runpod.Pod("myRandomPod",
        cloud_type="ALL",
        network_volume_id=test_network_storage.network_storage.apply(lambda x : fetch_id(x)),
        gpu_count=1,
        volume_in_gb=50,
        container_disk_in_gb=50,
        min_vcpu_count=2,
        min_memory_in_gb=15,
        gpu_type_id="NVIDIA GeForce RTX 4090",
        name="RunPod Pytorch",
        image_name="runpod/pytorch",
        docker_args="",
        ports="8888/http",
        volume_mount_path="/workspace",
        env=[runpod.PodEnvArgs(
            key="JUPYTER_PASSWORD",
            value="rns1hunbsstltcpad22d",
        )])

    my_random_template = runpod.Template("myRandomTemplate",
        container_disk_in_gb = 5,
        container_registry_auth_id = "",
        docker_args="python handler.py",
        env=[{"key": "hi", "value": "hello"}],
        image_name="runpod/serverless-hello-world:latest",
        is_public=False,
        is_serverless=True,
        name="Generated Serverless Template",
        ports="1293/http",
        readme="Some readme", # pass some value to this. Won't work otherwise
        start_jupyter=False,
        start_ssh=False,
        volume_in_gb=20,
        volume_mount_path="/workspace",
    )

    my_random_endpoint = runpod.Endpoint("myRandomEndpoint",
        gpu_ids="AMPERE_16,AMPERE_24,-NVIDIA L4",
        idle_timeout=100,
        locations="CA-MTL-2,CA-MTL-3,EU-RO-1,US-CA-1,US-GA-1,US-KS-2,US-OR-1,CA-MTL-1,US-TX-3,EUR-IS-1,EUR-IS-2,SEA-SG-1",
        name="myRandomEndpoint",
        network_volume_id=test_network_storage.network_storage.apply(lambda x : fetch_id(x)),
        scaler_type='REQUEST_COUNT',
        scaler_value=2,
        template_id=my_random_template.template.apply(lambda x : fetch_template_id(x)),
        workers_max=2,
        workers_min=1,
    )

    pulumi.export("pod", {
        "value": my_random_pod.pod,
    })
    pulumi.export("networkStorage", {
        "value": test_network_storage.network_storage,
    })
    pulumi.export("testTemplate", {
        "value": my_random_template.template,
    })
    pulumi.export("testEndpoint", {
        "value": my_random_endpoint.endpoint,
    })
except Exception as e:
    logger.exception(e)
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: provider-runpod-native
runtime: yaml
plugins:
  providers:
    - name: runpod
      path: ../../bin

resources:
  testNetworkStorage:
    type: runpod:NetworkStorage
    properties:
      name: testStorage1
      size: 5
      dataCenterId: EU-RO-1

  myRandomPod:
    type: runpod:Pod
    properties:
      cloudType: ALL
      networkVolumeId: ${testNetworkStorage.networkStorage.id}
      gpuCount: 1
      volumeInGb: 60
      containerDiskInGb: 50
      minVcpuCount: 2
      minMemoryInGb: 15
      gpuTypeId: "NVIDIA GeForce RTX 4090"
      name: "RunPod Pytorch"
      imageName: "runpod/pytorch"
      dockerArgs: ""
      ports: "8888/http"
      volumeMountPath: "/workspace"
      env:
        - key: "JUPYTER_PASSWORD"
          value: "rns1hunbsstltcpad22d"

  myRandomTemplate:
    type: runpod:Template
    properties:
      containerDiskInGb: 20
      containerRegistryAuthId: ""
      dockerArgs: "python3 -m http.server 8080"
      env: [{ key: "JUPYTER_PASSWORD", value: "rns1hunbsstltcpad22d" }]
      imageName: "nginx:latest"
      isPublic: false
      isServerless: true
      name: "RunPod Nginx"
      ports: "8080/http"
      readme: "Test template"
      startJupyter: false
      startSsh: false
      volumeInGb: 10
      volumeMountPath: "/workspace"

  myRandomEndpoint:
    type: runpod:Endpoint
    properties:
      gpuIds: "AMPERE_16,AMPERE_24,-NVIDIA L4"
      idleTimeout: 100
      locations: "CA-MTL-2,CA-MTL-3,EU-RO-1,US-CA-1,US-GA-1,US-KS-2,US-OR-1,CA-MTL-1,US-TX-3,EUR-IS-1,EUR-IS-2,SEA-SG-1"
      name: "myRandomEndpoint"
      networkVolumeId: ${testNetworkStorage.networkStorage.id}
      scalerType: "REQUEST_COUNT"
      scalerValue: 2
      templateId: ${myRandomTemplate.template.id}
      workersMax: 2
      workersMin: 1

outputs:
  pod:
    value: ${myRandomPod.pod}
  networkStorage:
    value: ${testNetworkStorage.networkStorage}
  template:
    value: ${myRandomTemplate.template}
  endpoint:
    value: ${myRandomEndpoint.endpoint}
```

{{% /choosable %}}
