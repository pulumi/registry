---
title: NGINX Ingress Controller 
meta_desc: Use Pulumi's Component for managing NGINX Ingress Controller installations using infrastructure as code.
layout: package
---

Easily manage NGINX Ingress Controller installations as a package available in all Pulumi languages.

## Example

{{< chooser language "typescript,python,go,csharp,java,yaml,hcl" >}}

{{% choosable language typescript %}}

```typescript
import * as k8s from "@pulumi/kubernetes";
import * as nginx from "@pulumi/kubernetes-ingress-nginx";

// Install the NGINX ingress controller to our cluster. The controller
// consists of a Pod and a Service. Install it and configure the controller
// to publish the load balancer IP address on each Ingress so that
// applications can depend on the IP address of the load balancer if needed.
const ctrl = new nginx.IngressController("myctrl", {
    controller: {
        publishService: {
            enabled: true,
        },
    },
});

// Now let's deploy two applications which are identical except for the
// names. We will later configure the ingress to direct traffic to them,
// one domain name per application instance.
const apps = [];
const appBase = "hello-k8s";
const appNames = [ `${appBase}-first`, `${appBase}-second` ];
for (const appName of appNames) {
    const appSvc = new k8s.core.v1.Service(`${appName}-svc`, {
        metadata: { name: appName },
        spec: {
            type: "ClusterIP",
            ports: [{ port: 80, targetPort: 8080 }],
            selector: { app: appName },
        },
    });
    const appDep = new k8s.apps.v1.Deployment(`${appName}-dep`, {
        metadata: { name: appName },
        spec: {
            replicas: 3,
            selector: {
                matchLabels: { app: appName },
            },
            template: {
                metadata: {
                    labels: { app: appName },
                },
                spec: {
                    containers: [{
                        name: appName,
                        image: "paulbouwer/hello-kubernetes:1.8",
                        ports: [{ containerPort: 8080 }],
                        env: [{ name: "MESSAGE", value: "Hello K8s!" }],
                    }],
                },
            },
        },
    });
    apps.push(appSvc.status);
}

// Next, expose the app using an Ingress.
const appIngress = new k8s.networking.v1.Ingress(`${appBase}-ingress`, {
    metadata: {
        name: "hello-k8s-ingress",
        annotations: {
            "kubernetes.io/ingress.class": "nginx",
        },
    },
    spec: {
        rules: [
            {
                // Replace this with your own domain!
                host: "myservicea.foo.org",
                http: {
                    paths: [{
                        pathType: "Prefix",
                        path: "/",
                        backend: {
                            service: {
                                name: appNames[0],
                                port: { number: 80 },
                            },
                        },
                    }],
                },
            },
            {
                // Replace this with your own domain!
                host: "myserviceb.foo.org",
                http: {
                    paths: [{
                        pathType: "Prefix",
                        path: "/",
                        backend: {
                            service: {
                                name: appNames[1],
                                port: { number: 80 },
                            },
                        },
                    }],
                },
            },
        ],
    },
});

export const appStatuses = apps;
export const controllerStatus = ctrl.status;
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
from pulumi_kubernetes.apps.v1 import Deployment
from pulumi_kubernetes.core.v1 import Service
from pulumi_kubernetes.networking.v1 import Ingress
from pulumi_kubernetes_ingress_nginx import IngressController, ControllerArgs, ControllerPublishServiceArgs

# Install the NGINX ingress controller to our cluster. The controller
# consists of a Pod and a Service. Install it and configure the controller
# to publish the load balancer IP address on each Ingress so that
# applications can depend on the IP address of the load balancer if needed.
ctrl = IngressController('myctrl',
                         controller=ControllerArgs(
                             publish_service=ControllerPublishServiceArgs(
                                 enabled=True,
                             ),
                         ),
                         )

# Now let's deploy two applications which are identical except for the
# names. We will later configure the ingress to direct traffic to them,
# one domain name per application instance.
apps = []
app_base = 'hello-k8s'
app_names = [ f'{app_base}-first', f'{app_base}-second' ]
for app_name in app_names:
    app_svc = Service(f'{app_name}-svc',
                      metadata={ 'name': app_name },
                      spec={
                          'type': 'ClusterIP',
                          'ports': [{ 'port': 80, 'targetPort': 8080 }],
                          'selector': { 'app': app_name },
                      },
                      )
    app_dep = Deployment(f'{app_name}-dep',
                         metadata={ 'name': app_name },
                         spec={
                             'replicas': 3,
                             'selector': {
                                 'matchLabels': { 'app': app_name },
                             },
                             'template': {
                                 'metadata': {
                                     'labels': { 'app': app_name },
                                 },
                                 'spec': {
                                     'containers': [{
                                         'name': app_name,
                                         'image': 'paulbouwer/hello-kubernetes:1.8',
                                         'ports': [{ 'containerPort': 8080 }],
                                         'env': [{ 'name': 'MESSAGE', 'value': 'Hello K8s!' }],
                                     }],
                                 },
                             },
                         },
                         )
    apps.append(app_svc.status)

# Next, expose the app using an Ingress.
app_ingress = Ingress(f'{app_base}-ingress',
                      metadata={
                          'name': 'hello-k8s-ingress',
                          'annotations': {
                              'kubernetes.io/ingress.class': 'nginx',
                          },
                      },
                      spec={
                          'rules': [
                              {
                                  # Replace this with your own domain!
                                  'host': 'myservicea.foo.org',
                                  'http': {
                                      'paths': [{
                                          'pathType': 'Prefix',
                                          'path': '/',
                                          'backend': {
                                              'service': {
                                                  'name': app_names[0],
                                                  'port': { 'number': 80 },
                                              },
                                          },
                                      }],
                                  },
                              },
                              {
                                  # Replace this with your own domain!
                                  'host': 'myserviceb.foo.org',
                                  'http': {
                                      'paths': [{
                                          'pathType': 'Prefix',
                                          'path': '/',
                                          'backend': {
                                              'service': {
                                                  'name': app_names[1],
                                                  'port': { 'number': 80 },
                                              },
                                          },
                                      }],
                                  },
                              },
                          ],
                      },
                      )
```

{{% /choosable %}}

{{% choosable language go %}}

```go
import (
	"fmt"

	nginx "github.com/pulumi/pulumi-kubernetes-ingress-nginx/sdk/go/kubernetes-ingress-nginx"
	appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/apps/v1"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	networkingv1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/networking/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Install the NGINX ingress controller to our cluster. The controller
		// consists of a Pod and a Service. Install it and configure the controller
		// to publish the load balancer IP address on each Ingress so that
		// applications can depend on the IP address of the load balancer if needed.
		ctrl, err := nginx.NewIngressController(ctx, "myctrl", &nginx.IngressControllerArgs{
			Controller: &nginx.ControllerArgs{
				PublishService: &nginx.ControllerPublishServiceArgs{
					Enabled: pulumi.Bool(true),
				},
			},
		})
		if err != nil {
			return err
		}

		// Now let's deploy two applications which are identical except for the
		// names. We will later configure the ingress to direct traffic to them,
		// one domain name per application instance.
		var apps pulumi.Array
		appBase := "hello-k8s"
		appNames := []string{appBase + "-first", appBase + "-second"}
		for _, appName := range appNames {
			appLabels := pulumi.StringMap{"app": pulumi.String(appName)}
			appSvc, err := corev1.NewService(ctx, fmt.Sprintf("%s-svc", appName), &corev1.ServiceArgs{
				Metadata: &metav1.ObjectMetaArgs{
					Name: pulumi.String(appName),
				},
				Spec: &corev1.ServiceSpecArgs{
					Type: pulumi.String("ClusterIP"),
					Ports: corev1.ServicePortArray{
						&corev1.ServicePortArgs{
							Port:       pulumi.Int(80),
							TargetPort: pulumi.Int(8080),
						},
					},
					Selector: appLabels,
				},
			})
			if err != nil {
				return err
			}
			_, err = appsv1.NewDeployment(ctx, fmt.Sprintf("%s-dep", appName), &appsv1.DeploymentArgs{
				Metadata: &metav1.ObjectMetaArgs{
					Name: pulumi.String(appName),
				},
				Spec: &appsv1.DeploymentSpecArgs{
					Replicas: pulumi.Int(3),
					Selector: &metav1.LabelSelectorArgs{
						MatchLabels: appLabels,
					},
					Template: &corev1.PodTemplateSpecArgs{
						Metadata: &metav1.ObjectMetaArgs{
							Labels: appLabels,
						},
						Spec: &corev1.PodSpecArgs{
							Containers: corev1.ContainerArray{
								&corev1.ContainerArgs{
									Name:  pulumi.String(appName),
									Image: pulumi.String("paulbouwer/hello-kubernetes:1.8"),
									Ports: corev1.ContainerPortArray{
										&corev1.ContainerPortArgs{
											ContainerPort: pulumi.Int(8080),
										},
									},
									Env: corev1.EnvVarArray{
										&corev1.EnvVarArgs{
											Name:  pulumi.String("MESSAGE"),
											Value: pulumi.String("Hello K8s!"),
										},
									},
								},
							},
						},
					},
				},
			})
			if err != nil {
				return err
			}
			apps = append(apps, appSvc.Status)
		}

		// Next, expose the app using an Ingress.
		rule := func(host, serviceName string) networkingv1.IngressRuleInput {
			return &networkingv1.IngressRuleArgs{
				Host: pulumi.String(host),
				Http: &networkingv1.HTTPIngressRuleValueArgs{
					Paths: networkingv1.HTTPIngressPathArray{
						&networkingv1.HTTPIngressPathArgs{
							PathType: pulumi.String("Prefix"),
							Path:     pulumi.String("/"),
							Backend: &networkingv1.IngressBackendArgs{
								Service: &networkingv1.IngressServiceBackendArgs{
									Name: pulumi.String(serviceName),
									Port: &networkingv1.ServiceBackendPortArgs{
										Number: pulumi.Int(80),
									},
								},
							},
						},
					},
				},
			}
		}
		_, err = networkingv1.NewIngress(ctx, fmt.Sprintf("%s-ingress", appBase), &networkingv1.IngressArgs{
			Metadata: &metav1.ObjectMetaArgs{
				Name: pulumi.String("hello-k8s-ingress"),
				Annotations: pulumi.StringMap{
					"kubernetes.io/ingress.class": pulumi.String("nginx"),
				},
			},
			Spec: &networkingv1.IngressSpecArgs{
				Rules: networkingv1.IngressRuleArray{
					// Replace these with your own domains!
					rule("myservicea.foo.org", appNames[0]),
					rule("myserviceb.foo.org", appNames[1]),
				},
			},
		})
		if err != nil {
			return err
		}

		ctx.Export("appStatuses", apps)
		ctx.Export("controllerStatus", ctrl.Status)
		return nil
	})
}
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using Pulumi;
using Pulumi.Kubernetes.Core.V1;
using Pulumi.Kubernetes.Networking.V1;
using Pulumi.Kubernetes.Types.Inputs.Apps.V1;
using Pulumi.Kubernetes.Types.Inputs.Core.V1;
using Pulumi.Kubernetes.Types.Inputs.Meta.V1;
using Pulumi.Kubernetes.Types.Inputs.Networking.V1;
using Pulumi.KubernetesIngressNginx;
using Pulumi.KubernetesIngressNginx.Inputs;

await Deployment.RunAsync(() =>
{
    // Install the NGINX ingress controller to our cluster. The controller
    // consists of a Pod and a Service. Install it and configure the controller
    // to publish the load balancer IP address on each Ingress so that
    // applications can depend on the IP address of the load balancer if needed.
    var ctrl = new IngressController("myctrl", new IngressControllerArgs
    {
        Controller = new ControllerArgs
        {
            PublishService = new ControllerPublishServiceArgs
            {
                Enabled = true,
            },
        },
    });

    // Now let's deploy two applications which are identical except for the
    // names. We will later configure the ingress to direct traffic to them,
    // one domain name per application instance.
    var apps = new List<Output<Pulumi.Kubernetes.Types.Outputs.Core.V1.ServiceStatus>>();
    var appBase = "hello-k8s";
    var appNames = new[] { $"{appBase}-first", $"{appBase}-second" };
    foreach (var appName in appNames)
    {
        var appLabels = new InputMap<string> { { "app", appName } };
        var appSvc = new Service($"{appName}-svc", new ServiceArgs
        {
            Metadata = new ObjectMetaArgs { Name = appName },
            Spec = new ServiceSpecArgs
            {
                Type = "ClusterIP",
                Ports = { new ServicePortArgs { Port = 80, TargetPort = 8080 } },
                Selector = appLabels,
            },
        });
        var appDep = new Pulumi.Kubernetes.Apps.V1.Deployment($"{appName}-dep", new DeploymentArgs
        {
            Metadata = new ObjectMetaArgs { Name = appName },
            Spec = new DeploymentSpecArgs
            {
                Replicas = 3,
                Selector = new LabelSelectorArgs { MatchLabels = appLabels },
                Template = new PodTemplateSpecArgs
                {
                    Metadata = new ObjectMetaArgs { Labels = appLabels },
                    Spec = new PodSpecArgs
                    {
                        Containers =
                        {
                            new ContainerArgs
                            {
                                Name = appName,
                                Image = "paulbouwer/hello-kubernetes:1.8",
                                Ports = { new ContainerPortArgs { ContainerPortValue = 8080 } },
                                Env = { new EnvVarArgs { Name = "MESSAGE", Value = "Hello K8s!" } },
                            },
                        },
                    },
                },
            },
        });
        apps.Add(appSvc.Status);
    }

    // Next, expose the app using an Ingress.
    IngressRuleArgs Rule(string host, string serviceName) => new IngressRuleArgs
    {
        Host = host,
        Http = new HTTPIngressRuleValueArgs
        {
            Paths =
            {
                new HTTPIngressPathArgs
                {
                    PathType = "Prefix",
                    Path = "/",
                    Backend = new IngressBackendArgs
                    {
                        Service = new IngressServiceBackendArgs
                        {
                            Name = serviceName,
                            Port = new ServiceBackendPortArgs { Number = 80 },
                        },
                    },
                },
            },
        },
    };
    var appIngress = new Ingress($"{appBase}-ingress", new IngressArgs
    {
        Metadata = new ObjectMetaArgs
        {
            Name = "hello-k8s-ingress",
            Annotations = { { "kubernetes.io/ingress.class", "nginx" } },
        },
        Spec = new IngressSpecArgs
        {
            Rules =
            {
                // Replace these with your own domains!
                Rule("myservicea.foo.org", appNames[0]),
                Rule("myserviceb.foo.org", appNames[1]),
            },
        },
    });

    return new Dictionary<string, object?>
    {
        ["appStatuses"] = Output.All(apps),
        ["controllerStatus"] = ctrl.Status,
    };
});
```

{{% /choosable %}}

{{% choosable language java %}}

```java
import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.kubernetes.apps.v1.Deployment;
import com.pulumi.kubernetes.apps.v1.DeploymentArgs;
import com.pulumi.kubernetes.apps.v1.inputs.DeploymentSpecArgs;
import com.pulumi.kubernetes.core.v1.Service;
import com.pulumi.kubernetes.core.v1.ServiceArgs;
import com.pulumi.kubernetes.core.v1.inputs.ContainerArgs;
import com.pulumi.kubernetes.core.v1.inputs.ContainerPortArgs;
import com.pulumi.kubernetes.core.v1.inputs.EnvVarArgs;
import com.pulumi.kubernetes.core.v1.inputs.PodSpecArgs;
import com.pulumi.kubernetes.core.v1.inputs.PodTemplateSpecArgs;
import com.pulumi.kubernetes.core.v1.inputs.ServicePortArgs;
import com.pulumi.kubernetes.core.v1.inputs.ServiceSpecArgs;
import com.pulumi.kubernetes.core.v1.outputs.ServiceStatus;
import com.pulumi.kubernetes.meta.v1.inputs.LabelSelectorArgs;
import com.pulumi.kubernetes.meta.v1.inputs.ObjectMetaArgs;
import com.pulumi.kubernetes.networking.v1.Ingress;
import com.pulumi.kubernetes.networking.v1.IngressArgs;
import com.pulumi.kubernetes.networking.v1.inputs.HTTPIngressPathArgs;
import com.pulumi.kubernetes.networking.v1.inputs.HTTPIngressRuleValueArgs;
import com.pulumi.kubernetes.networking.v1.inputs.IngressBackendArgs;
import com.pulumi.kubernetes.networking.v1.inputs.IngressRuleArgs;
import com.pulumi.kubernetes.networking.v1.inputs.IngressServiceBackendArgs;
import com.pulumi.kubernetes.networking.v1.inputs.IngressSpecArgs;
import com.pulumi.kubernetes.networking.v1.inputs.ServiceBackendPortArgs;
import com.pulumi.kubernetesingressnginx.IngressController;
import com.pulumi.kubernetesingressnginx.IngressControllerArgs;
import com.pulumi.kubernetesingressnginx.inputs.ControllerArgs;
import com.pulumi.kubernetesingressnginx.inputs.ControllerPublishServiceArgs;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.Optional;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        // Install the NGINX ingress controller to our cluster. The controller
        // consists of a Pod and a Service. Install it and configure the controller
        // to publish the load balancer IP address on each Ingress so that
        // applications can depend on the IP address of the load balancer if needed.
        var ctrl = new IngressController("myctrl", IngressControllerArgs.builder()
            .controller(ControllerArgs.builder()
                .publishService(ControllerPublishServiceArgs.builder()
                    .enabled(true)
                    .build())
                .build())
            .build());

        // Now let's deploy two applications which are identical except for the
        // names. We will later configure the ingress to direct traffic to them,
        // one domain name per application instance.
        var apps = new ArrayList<Output<Optional<ServiceStatus>>>();
        var appBase = "hello-k8s";
        var appNames = List.of(appBase + "-first", appBase + "-second");
        for (var appName : appNames) {
            var appLabels = Map.of("app", appName);
            var appSvc = new Service(appName + "-svc", ServiceArgs.builder()
                .metadata(ObjectMetaArgs.builder()
                    .name(appName)
                    .build())
                .spec(ServiceSpecArgs.builder()
                    .type("ClusterIP")
                    .ports(ServicePortArgs.builder()
                        .port(80)
                        .targetPort(8080)
                        .build())
                    .selector(appLabels)
                    .build())
                .build());
            new Deployment(appName + "-dep", DeploymentArgs.builder()
                .metadata(ObjectMetaArgs.builder()
                    .name(appName)
                    .build())
                .spec(DeploymentSpecArgs.builder()
                    .replicas(3)
                    .selector(LabelSelectorArgs.builder()
                        .matchLabels(appLabels)
                        .build())
                    .template(PodTemplateSpecArgs.builder()
                        .metadata(ObjectMetaArgs.builder()
                            .labels(appLabels)
                            .build())
                        .spec(PodSpecArgs.builder()
                            .containers(ContainerArgs.builder()
                                .name(appName)
                                .image("paulbouwer/hello-kubernetes:1.8")
                                .ports(ContainerPortArgs.builder()
                                    .containerPort(8080)
                                    .build())
                                .env(EnvVarArgs.builder()
                                    .name("MESSAGE")
                                    .value("Hello K8s!")
                                    .build())
                                .build())
                            .build())
                        .build())
                    .build())
                .build());
            apps.add(appSvc.status());
        }

        // Next, expose the app using an Ingress.
        new Ingress(appBase + "-ingress", IngressArgs.builder()
            .metadata(ObjectMetaArgs.builder()
                .name("hello-k8s-ingress")
                .annotations(Map.of("kubernetes.io/ingress.class", "nginx"))
                .build())
            .spec(IngressSpecArgs.builder()
                .rules(
                    // Replace these with your own domains!
                    rule("myservicea.foo.org", appNames.get(0)),
                    rule("myserviceb.foo.org", appNames.get(1)))
                .build())
            .build());

        ctx.export("appStatuses", Output.all(apps));
        ctx.export("controllerStatus", ctrl.status());
    }

    private static IngressRuleArgs rule(String host, String serviceName) {
        return IngressRuleArgs.builder()
            .host(host)
            .http(HTTPIngressRuleValueArgs.builder()
                .paths(HTTPIngressPathArgs.builder()
                    .pathType("Prefix")
                    .path("/")
                    .backend(IngressBackendArgs.builder()
                        .service(IngressServiceBackendArgs.builder()
                            .name(serviceName)
                            .port(ServiceBackendPortArgs.builder()
                                .number(80)
                                .build())
                            .build())
                        .build())
                    .build())
                .build())
            .build();
    }
}
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
resources:
  # Install the NGINX ingress controller to our cluster. The controller
  # consists of a Pod and a Service. Install it and configure the controller
  # to publish the load balancer IP address on each Ingress so that
  # applications can depend on the IP address of the load balancer if needed.
  myctrl:
    type: kubernetes-ingress-nginx:IngressController
    properties:
      controller:
        publishService:
          enabled: true

  # Now let's deploy two applications which are identical except for the
  # names. We will later configure the ingress to direct traffic to them,
  # one domain name per application instance.
  hello-k8s-first-svc:
    type: kubernetes:core/v1:Service
    properties:
      metadata:
        name: hello-k8s-first
      spec:
        type: ClusterIP
        ports:
          - port: 80
            targetPort: 8080
        selector:
          app: hello-k8s-first
  hello-k8s-first-dep:
    type: kubernetes:apps/v1:Deployment
    properties:
      metadata:
        name: hello-k8s-first
      spec:
        replicas: 3
        selector:
          matchLabels:
            app: hello-k8s-first
        template:
          metadata:
            labels:
              app: hello-k8s-first
          spec:
            containers:
              - name: hello-k8s-first
                image: paulbouwer/hello-kubernetes:1.8
                ports:
                  - containerPort: 8080
                env:
                  - name: MESSAGE
                    value: Hello K8s!
  hello-k8s-second-svc:
    type: kubernetes:core/v1:Service
    properties:
      metadata:
        name: hello-k8s-second
      spec:
        type: ClusterIP
        ports:
          - port: 80
            targetPort: 8080
        selector:
          app: hello-k8s-second
  hello-k8s-second-dep:
    type: kubernetes:apps/v1:Deployment
    properties:
      metadata:
        name: hello-k8s-second
      spec:
        replicas: 3
        selector:
          matchLabels:
            app: hello-k8s-second
        template:
          metadata:
            labels:
              app: hello-k8s-second
          spec:
            containers:
              - name: hello-k8s-second
                image: paulbouwer/hello-kubernetes:1.8
                ports:
                  - containerPort: 8080
                env:
                  - name: MESSAGE
                    value: Hello K8s!

  # Next, expose the app using an Ingress.
  hello-k8s-ingress:
    type: kubernetes:networking.k8s.io/v1:Ingress
    properties:
      metadata:
        name: hello-k8s-ingress
        annotations:
          kubernetes.io/ingress.class: nginx
      spec:
        rules:
          # Replace this with your own domain!
          - host: myservicea.foo.org
            http:
              paths:
                - pathType: Prefix
                  path: /
                  backend:
                    service:
                      name: hello-k8s-first
                      port:
                        number: 80
          # Replace this with your own domain!
          - host: myserviceb.foo.org
            http:
              paths:
                - pathType: Prefix
                  path: /
                  backend:
                    service:
                      name: hello-k8s-second
                      port:
                        number: 80

outputs:
  appStatuses:
    - ${hello-k8s-first-svc.status}
    - ${hello-k8s-second-svc.status}
  controllerStatus: ${myctrl.status}
```

{{% /choosable %}}

{{% choosable language hcl %}}

```hcl
terraform {
  required_providers {
    kubernetes = {
      source = "pulumi/kubernetes"
    }
    kubernetes-ingress-nginx = {
      source = "pulumi/kubernetes-ingress-nginx"
    }
  }
}

# Install the NGINX ingress controller to our cluster. The controller
# consists of a Pod and a Service. Install it and configure the controller
# to publish the load balancer IP address on each Ingress so that
# applications can depend on the IP address of the load balancer if needed.
resource "kubernetes-ingress-nginx_ingress_controller" "myctrl" {
  controller = {
    publish_service = {
      enabled = true
    }
  }
}

# Now let's deploy two applications which are identical except for the
# names. We will later configure the ingress to direct traffic to them,
# one domain name per application instance.
locals {
  app_base  = "hello-k8s"
  app_names = ["${local.app_base}-first", "${local.app_base}-second"]
}

resource "kubernetes_core_v1_service" "app_svc" {
  for_each = toset(local.app_names)

  metadata = {
    name = each.key
  }
  spec = {
    type = "ClusterIP"
    ports = [{
      port        = 80
      target_port = 8080
    }]
    selector = {
      app = each.key
    }
  }
}

resource "kubernetes_apps_v1_deployment" "app_dep" {
  for_each = toset(local.app_names)

  metadata = {
    name = each.key
  }
  spec = {
    replicas = 3
    selector = {
      match_labels = {
        app = each.key
      }
    }
    template = {
      metadata = {
        labels = {
          app = each.key
        }
      }
      spec = {
        containers = [{
          name  = each.key
          image = "paulbouwer/hello-kubernetes:1.8"
          ports = [{
            container_port = 8080
          }]
          env = [{
            name  = "MESSAGE"
            value = "Hello K8s!"
          }]
        }]
      }
    }
  }
}

# Next, expose the app using an Ingress.
resource "kubernetes_networking.k8s.io_v1_ingress" "app_ingress" {
  metadata = {
    name = "hello-k8s-ingress"
    annotations = {
      "kubernetes.io/ingress.class" = "nginx"
    }
  }
  spec = {
    rules = [
      {
        # Replace this with your own domain!
        host = "myservicea.foo.org"
        http = {
          paths = [{
            path_type = "Prefix"
            path      = "/"
            backend = {
              service = {
                name = local.app_names[0]
                port = {
                  number = 80
                }
              }
            }
          }]
        }
      },
      {
        # Replace this with your own domain!
        host = "myserviceb.foo.org"
        http = {
          paths = [{
            path_type = "Prefix"
            path      = "/"
            backend = {
              service = {
                name = local.app_names[1]
                port = {
                  number = 80
                }
              }
            }
          }]
        }
      },
    ]
  }
}

output "appStatuses" {
  value = [for svc in kubernetes_core_v1_service.app_svc : svc.status]
}

output "controllerStatus" {
  value = kubernetes-ingress-nginx_ingress_controller.myctrl.status
}
```

{{% /choosable %}}

{{< /chooser >}}
