---
title: NGINX Ingress Controller 
meta_desc: Use Pulumi's Component for managing NGINX Ingress Controller installations using infrastructure as code.
layout: overview
---

Easily manage NGINX Ingress Controller installations as a package available in all Pulumi languages.

## Example

{{< chooser language "typescript,python" >}}

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

{{< /chooser >}}
