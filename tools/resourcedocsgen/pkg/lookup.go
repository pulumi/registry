// Copyright 2024, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pkg

var CategoryNameMap = map[string]PackageCategory{
	"cloud":          PackageCategoryCloud,
	"database":       PackageCategoryDatabase,
	"infrastructure": PackageCategoryInfrastructure,
	"monitoring":     PackageCategoryMonitoring,
	"network":        PackageCategoryNetwork,
	"utility":        PackageCategoryUtility,
	"vcs":            PackageCategoryVCS,
}

var CategoryLookup = map[string]PackageCategory{
	"aiven":                               PackageCategoryInfrastructure,
	"akamai":                              PackageCategoryNetwork,
	"alicloud":                            PackageCategoryCloud,
	"auth0":                               PackageCategoryInfrastructure,
	"aws":                                 PackageCategoryCloud,
	"aws-apigateway":                      PackageCategoryCloud,
	"aws-miniflux":                        PackageCategoryCloud,
	"aws-native":                          PackageCategoryCloud,
	"aws-quickstart-aurora-mysql":         PackageCategoryCloud,
	"aws-quickstart-aurora-postgresql":    PackageCategoryCloud,
	"aws-quickstart-redshift":             PackageCategoryCloud,
	"aws-serverless":                      PackageCategoryCloud,
	"aws-quickstart-vpc":                  PackageCategoryCloud,
	"aws-s3-replicated-bucket":            PackageCategoryCloud,
	"azure":                               PackageCategoryCloud,
	"azure-native":                        PackageCategoryCloud,
	"azure-quickstart-acr-geo-replicated": PackageCategoryCloud,
	"azure-quickstart-aks":                PackageCategoryCloud,
	"azure-quickstart-compute":            PackageCategoryCloud,
	"azure-quickstart-sql":                PackageCategoryCloud,
	"azuread":                             PackageCategoryInfrastructure,
	"azuredevops":                         PackageCategoryInfrastructure,
	"civo":                                PackageCategoryCloud,
	"cloudamqp":                           PackageCategoryCloud,
	"cloudflare":                          PackageCategoryNetwork,
	"cloudinit":                           PackageCategoryUtility,
	"confluent":                           PackageCategoryInfrastructure,
	"consul":                              PackageCategoryInfrastructure,
	"coredns-helm":                        PackageCategoryNetwork,
	"datadog":                             PackageCategoryMonitoring,
	"digitalocean":                        PackageCategoryCloud,
	"dnsimple":                            PackageCategoryNetwork,
	"docker":                              PackageCategoryInfrastructure,
	"docker-buildkit":                     PackageCategoryInfrastructure,
	"eks":                                 PackageCategoryCloud,
	"equinix-metal":                       PackageCategoryCloud,
	"f5bigip":                             PackageCategoryNetwork,
	"fastly":                              PackageCategoryCloud,
	"gcp":                                 PackageCategoryCloud,
	"gcp-cloudrun-multi-region":           PackageCategoryCloud,
	"gcp-project-scaffold":                PackageCategoryCloud,
	"google-native":                       PackageCategoryCloud,
	"github":                              PackageCategoryVCS,
	"github-serverless-webhook":           PackageCategoryVCS,
	"gitlab":                              PackageCategoryVCS,
	"hcloud":                              PackageCategoryCloud,
	"istio-helm":                          PackageCategoryInfrastructure,
	"jaeger-helm":                         PackageCategoryMonitoring,
	"kafka":                               PackageCategoryInfrastructure,
	"keycloak":                            PackageCategoryInfrastructure,
	"kong":                                PackageCategoryInfrastructure,
	"kubernetes":                          PackageCategoryCloud,
	"libvirt":                             PackageCategoryUtility,
	"linode":                              PackageCategoryCloud,
	"mailgun":                             PackageCategoryInfrastructure,
	"minio":                               PackageCategoryInfrastructure,
	"mongodbatlas":                        PackageCategoryDatabase,
	"mysql":                               PackageCategoryDatabase,
	"newrelic":                            PackageCategoryMonitoring,
	"nginx-ingress-controller-helm":       PackageCategoryNetwork,
	"nomad":                               PackageCategoryInfrastructure,
	"ns1":                                 PackageCategoryNetwork,
	"okta":                                PackageCategoryInfrastructure,
	"onelogin":                            PackageCategoryInfrastructure,
	"openstack":                           PackageCategoryCloud,
	"opsgenie":                            PackageCategoryInfrastructure,
	"pagerduty":                           PackageCategoryInfrastructure,
	"postgresql":                          PackageCategoryDatabase,
	"prometheus-helm":                     PackageCategoryMonitoring,
	"rabbitmq":                            PackageCategoryInfrastructure,
	"rancher2":                            PackageCategoryInfrastructure,
	"random":                              PackageCategoryUtility,
	"rke":                                 PackageCategoryInfrastructure,
	"shipa":                               PackageCategoryCloud,
	"signalfx":                            PackageCategoryMonitoring,
	"snowflake":                           PackageCategoryInfrastructure,
	"splunk":                              PackageCategoryInfrastructure,
	"spotinst":                            PackageCategoryInfrastructure,
	"sumologic":                           PackageCategoryMonitoring,
	"tls":                                 PackageCategoryUtility,
	"vault":                               PackageCategoryInfrastructure,
	"venafi":                              PackageCategoryInfrastructure,
	"vsphere":                             PackageCategoryCloud,
	"wavefront":                           PackageCategoryMonitoring,
	"yandex":                              PackageCategoryCloud,
}

// TitleLookup is a map pf package name to the desired display name
// for display in the TOC menu under API Reference.
//
// NOTE: For the time being this lookup map and the one used by the docs
// generator in `pulumi/pulumi` must be kept up-to-date.
//
// TODO[pulumi/pulumi#7813]: Remove this lookup once display name is available in the
// Pulumi schema. To remove safely, we must show that .DisplayName is set in the latest
// *released* version of each of these providers.
//
// Safe removal is tracked in pulumi/registry#4672.
//
//nolint:lll
var TitleLookup = map[string]string{
	"aiven":                               "Aiven",           // https://github.com/pulumi/pulumi-aiven/pull/761
	"akamai":                              "Akamai",          // https://github.com/pulumi/pulumi-akamai/pull/642
	"alicloud":                            "Alibaba Cloud",   // https://github.com/pulumi/pulumi-alicloud/pull/839
	"auth0":                               "Auth0",           // https://github.com/pulumi/pulumi-auth0/pull/695
	"aws":                                 "AWS",             // https://github.com/pulumi/pulumi-aws/pull/4765
	"aws-apigateway":                      "AWS API Gateway", // https://github.com/pulumi/pulumi-aws-apigateway/pull/43
	"aws-miniflux":                        "Miniflux",
	"aws-native":                          "AWS Cloud Control",
	"aws-quickstart-aurora-mysql":         "AWS QuickStart Aurora MySQL",
	"aws-quickstart-aurora-postgresql":    "AWS QuickStart Aurora PostgreSQL",
	"aws-quickstart-redshift":             "AWS QuickStart Redshift",
	"aws-serverless":                      "AWS Serverless",
	"aws-quickstart-vpc":                  "AWS QuickStart VPC",
	"aws-s3-replicated-bucket":            "AWS S3 Replicated Bucket",
	"azure":                               "Azure Classic", // https://github.com/pulumi/pulumi-azure/pull/2630
	"azure-native":                        "Azure Native",
	"azure-quickstart-acr-geo-replicated": "Azure QuickStart ACR Geo Replicated",
	"azure-quickstart-aks":                "Azure QuickStart AKS",
	"azure-quickstart-compute":            "Azure QuickStart Compute",
	"azure-quickstart-sql":                "Azure QuickStart SQL",
	"azuread":                             "Azure Active Directory", // https://github.com/pulumi/pulumi-azuread/pull/321
	"azuredevops":                         "Azure DevOps",           // https://github.com/pulumi/pulumi-azuredevops/pull/519
	"azuresel":                            "Azure",                  // https://github.com/pulumi/pulumi-azure/pull/2630
	"civo":                                "Civo",                   // https://github.com/pulumi/pulumi-civo/pull/647
	"cloudamqp":                           "CloudAMQP",              // https://github.com/pulumi/pulumi-cloudamqp/pull/589
	"cloudflare":                          "Cloudflare",             // https://github.com/pulumi/pulumi-cloudflare/pull/1008
	"cloudinit":                           "cloud-init",             // https://github.com/pulumi/pulumi-cloudinit/pull/465
	"confluent":                           "Confluent Cloud",        // https://github.com/pulumi/pulumi-confluentcloud/blob/0f4313de2fc3a025225950279f1ac68c6116cad9/provider/resources.go#L58
	"consul":                              "Consul",                 // https://github.com/pulumi/pulumi-consul/pull/574
	"coredns-helm":                        "CoreDNS (Helm)",
	"datadog":                             "Datadog",      // https://github.com/pulumi/pulumi-datadog/pull/729
	"digitalocean":                        "DigitalOcean", // https://github.com/pulumi/pulumi-digitalocean/pull/908
	"dnsimple":                            "DNSimple",     // https://github.com/pulumi/pulumi-dnsimple/pull/558
	"docker":                              "Docker",
	"docker-buildkit":                     "Docker BuildKit",
	"eks":                                 "Amazon EKS",
	"equinix-metal":                       "Equinix Metal",
	"f5bigip":                             "f5 BIG-IP",
	"fastly":                              "Fastly",               // https://github.com/pulumi/pulumi-fastly/pull/693
	"gcp":                                 "Google Cloud Classic", // https://github.com/pulumi/pulumi-gcp/pull/2745
	"gcp-cloudrun-multi-region":           "Google Cloud Run Multi-Region",
	"gcp-project-scaffold":                "Google Project Scaffolding",
	"google-native":                       "Google Cloud Native",
	"github":                              "GitHub",
	"github-serverless-webhook":           "GitHub Serverless Webhook",
	"gitlab":                              "GitLab",
	"hcloud":                              "Hetzner Cloud",
	"istio-helm":                          "Istio (Helm)",
	"jaeger-helm":                         "Jaeger (Helm)",
	"kafka":                               "Kafka",
	"keycloak":                            "Keycloak",
	"kong":                                "Kong",
	"kubernetes":                          "Kubernetes",
	"libvirt":                             "libvirt",
	"linode":                              "Linode",
	"mailgun":                             "Mailgun",
	"minio":                               "MinIO",
	"mongodbatlas":                        "MongoDB Atlas",
	"mysql":                               "MySQL",
	"newrelic":                            "New Relic",
	"nginx-ingress-controller-helm":       "NGINX Ingress Controller (Helm)",
	"nomad":                               "Nomad",
	"ns1":                                 "NS1",
	"okta":                                "Okta",
	"openstack":                           "OpenStack",
	"opsgenie":                            "Opsgenie",
	"packet":                              "Packet",
	"pagerduty":                           "PagerDuty",
	"postgresql":                          "PostgreSQL",
	"prometheus-helm":                     "Prometheus (Helm)",
	"rabbitmq":                            "RabbitMQ",
	"rancher2":                            "Rancher 2",
	"random":                              "random",
	"rke":                                 "Rancher RKE",
	"shipa":                               "Shipa",
	"signalfx":                            "SignalFx",
	"snowflake":                           "Snowflake",
	"splunk":                              "Splunk",
	"spotinst":                            "Spotinst",
	"sumologic":                           "Sumo Logic",
	"tls":                                 "TLS",
	"vault":                               "Vault",
	"venafi":                              "Venafi",
	"vsphere":                             "vSphere",
	"wavefront":                           "Wavefront",
	"yandex":                              "Yandex",
}
