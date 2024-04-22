describe("Registry", () => {

    // By default, every Cypress test runs in "isolation", which is essentially like opening each call to `it()` in entirely
    // new incognito browser tab. In general this is a good thing, but it comes with a big performance cost, so for tests like these
    // that don't rely on clean state, it's much faster to turn isolation off, and just reuse the existing browser context between
    // page visits. (Pages are still requested individually, but things like cookies and localStorage are persisted across requests,)
    // More on how this all works at https://docs.cypress.io/guides/core-concepts/test-isolation.
    describe("API docs", { testIsolation: false }, () => {

        const resources = [
            [ "aws.lambda.Function", "lambda/function" ],
            [ "aws.s3.Bucket", "s3/bucket" ],
            [ "aws.iam.Role", "iam/role" ],
            [ "aws.ec2.Instance", "ec2/instance" ],
            [ "aws.rds.Instance", "rds/instance" ],
            [ "aws.ec2.SecurityGroup", "ec2/securitygroup" ],
            [ "kubernetes.helm.sh.v3.Chart", "helm/v3/chart", "kubernetes.helm.sh/v3.Chart" ],
            [ "aws.ec2.Vpc", "ec2/vpc" ],
            [ "kubernetes.helm.sh.v3.Release", "helm/v3/release", "kubernetes.helm.sh/v3.Release" ],
            [ "aws.alb.LoadBalancer", "alb/loadbalancer", ],
            [ "aws.cloudfront.Distribution", "cloudfront/distribution" ],
            [ "aws.ecs.TaskDefinition", "ecs/taskdefinition" ],
            [ "aws.route53.Record", "route53/record" ],
            [ "aws.ec2.LaunchTemplate", "ec2/launchtemplate" ],
            [ "gcp.storage.Bucket", "storage/bucket" ],
            [ "kubernetes.apps.v1.Deployment ", "apps/v1/deployment", "kubernetes.apps/v1.Deployment" ],
            [ "aws.rds.Cluster", "rds/cluster" ],
            [ "aws.iam.Policy", "iam/policy" ],
            [ "aws.eks.Cluster", "eks/cluster" ],
            [ "azure-native.storage.StorageAccount", "storage/storageaccount" ],
            [ "gcp.serviceaccount.Account", "serviceaccount/account" ],
            [ "aws.secretsmanager.Secret", "secretsmanager/secret" ],
            [ "aws.ec2.Subnet", "ec2/subnet" ],
            [ "kubernetes.networking.k8s.io.v1.Ingress", "networking/v1/ingress", "kubernetes.networking.k8s.io/v1.Ingress" ],
            [ "aws.dynamodb.Table", "dynamodb/table" ],
            [ "kubernetes.autoscaling.v2beta2.HorizontalPodAutoscaler", "autoscaling/v2beta2/horizontalpodautoscaler", "kubernetes.autoscaling/v2beta2.HorizontalPodAutoscaler" ],
            [ "aws.ecs.Service", "ecs/service" ],
            [ "aws.acm.Certificate", "acm/certificate" ],
            [ "aws.apigateway.RestApi", "apigateway/restapi" ],
            [ "gcp.container.Cluster", "container/cluster" ],
            [ "aws.autoscaling.Group", "autoscaling/group" ],
            [ "gcp.cloudrun.Service", "cloudrun/service" ],
            [ "azure-native.resources.ResourceGroup", "resources/resourcegroup" ],
            [ "aws.iam.RolePolicyAttachment ", "iam/rolepolicyattachment" ],
            [ "azure-native.web.WebApp ", "web/webapp" ],
            [ "aws.lb.TargetGroup", "lb/targetgroup" ],
            [ "awsx.ec2.Vpc", "ec2/vpc" ],
            [ "kubernetes.core.v1.Secret", "core/v1/secret", "kubernetes.core/v1.Secret" ],
            [ "aws.sqs.Queue", "sqs/queue" ],
            [ "gcp.compute.Instance", "compute/instance" ],
            [ "awsx.ecs.FargateService", "ecs/fargateservice" ],
            [ "kubernetes.yaml.ConfigFile", "yaml/configfile" ],
            [ "aws.alb.Listener", "alb/listener" ],
            [ "aws.eks.NodeGroup", "eks/nodegroup" ],
            [ "aws.ecs.Cluster", "ecs/cluster" ],
            [ "aws.iam.getPolicyDocument", "iam/getpolicydocument" ],
            [ "aws.s3.BucketPolicy", "s3/bucketpolicy" ],
            [ "aws.cloudwatch.MetricAlarm", "cloudwatch/metricalarm" ],
            [ "azure-native.keyvault.Vault ", "keyvault/vault" ],
            [ "aws.sns.Topic", "sns/topic" ],
        ];

        resources.forEach(resource => {
            const [ token, path, alternateTitle ] = resource;
            const title = token;
            const provider = token.split(".").slice(0)[0];
            const resourceName = token.split(".").slice(-1)[0];

            const page = {
                title,
                alternateTitle,
                provider,
                resourceName,
                path: `/registry/packages/${provider}/api-docs/${path}/`,
            };

            describe(`${page.title} resource (${[ Cypress.config().baseUrl, page.path ].join("")})`, () => {

                beforeEach(() => {
                    cy.visit(page.path);
                });

                it("has the correct page title", () => {
                    // We verify that the title includes both the resource title and the site title. Sometimes there is
                    // a trailing space after the resource title that messes with the logic, so we just check to make sure
                    // it includes both pieces of text.
                    cy.get("head title").should("include.text", page.title).and("include.text", "| Pulumi Registry");
                });

                describe("main content", () => {
                    const container = "main div.docs-main-content";
                    const expectedTitle = (page.alternateTitle || page.title).trim();

                    it("has the correct H1", () => {
                        cy.get(container)
                            .find("h1")
                            .should("be.visible")
                            .should("have.text", expectedTitle.trim());
                    });

                    describe("description", () => {
                        it("leads with a non-empty paragraph tag", () => {
                            cy.get(container)
                                .find("section.docs-content > p")
                                .first()
                                .should("exist")
                                .invoke("text")
                                .should("match", /^\w+/);
                        });
                    });

                    it("renders the correct set of subsections, in the correct order", () => {
                        cy.get(container).find("section.docs-content h2").as("sections");

                        // possibleHeadings is the set of H2 headings that can appear on an API docs page.
                        // Some pages won't contain *all* of these headings, but the headings that do appear
                        // should always be presented in the following order.
                        //
                        // Note that this test intentionally fails on pages that contain unexpected H2 headings,
                        // such as those that are occasionally baked into upstream provider docs.
                        const possibleHeadings = [
                            "Example Usage",
                            `Create ${page.resourceName} Resource`,
                            `${page.resourceName} Resource Properties`,
                            `Look up Existing ${page.resourceName} Resource`,
                            "Supporting Types",
                            "Import",
                            "Package Details",
                        ];

                        cy.get("@sections").then(headings => {

                            const actualHeadings = headings
                                .map((_, heading) => heading.textContent)
                                .filter(heading => possibleHeadings.includes(heading))
                                .get();

                            const expectedHeadings = possibleHeadings
                                .filter(heading => actualHeadings.includes(heading));

                            // JSON.stringify() gives us an easy way to compare the two lists.
                            expect(JSON.stringify(actualHeadings)).to.equal(JSON.stringify(expectedHeadings));
                        });
                    });

                    describe("Examples section", () => {
                        it("contains at least one example", () => {
                            cy.get(container)
                                .find("h2 ~ h3")
                                .should("have.length.of.at.least", 1);
                        });

                        describe("examples", () => {
                            it("all have language choosers", () => {
                                cy.get("#example-usage + h3 + div > pulumi-chooser")
                            });
                        });
                    });

                    describe("Create section", () => {
                        it("renders a minimal example in all languages", () => {
                            // renders the minimal example right after the `Constructor syntax` heading
                            cy.get(container).find("#constructor-syntax + div > pulumi-chooser")
                        });

                        it.skip("renders a full example in all languages", () => {
                          // We may disable this section for pages that fail to convert, so some
                          // examples may not have this section once this is done.
                          // https://github.com/pulumi/registry/issues/4320
                        });
                    });

                    describe("Inputs and Outputs sections", () => {
                        const propertyLists = ".resources-properties";

                        it("render deprecation notices properly", () => {
                            cy.get(container).find(propertyLists).as("propertyLists");

                            cy.get("@propertyLists").then(lists => {
                                const deprecatedProperties = "dt.property-deprecated + dd > p.property-message";

                                if (lists.find(deprecatedProperties).length > 0) {
                                    cy.get("@propertyLists")
                                        .find(deprecatedProperties)
                                        .first()
                                        .invoke("text")
                                        .should("not.equal", "Deprecated:");
                                }
                            });
                        });

                        describe("type links", () => {
                            it("all point to an on-page supporting type", () => {
                                cy.get(container).find(propertyLists).find("dt .property-type a[href*='#']").then(links => {
                                    links.each((i, link) => {
                                        if (link.href.startsWith("#")) {
                                            cy.get(`#supporting-types ~ ${link.href}`).should("exist");
                                        }
                                    });
                                });
                            });
                        })
                    });

                    describe("Lookup section", () => {
                        it("renders an example for all languages", () => {
                            // Should contain a code example.
                            cy.get("#look-up + p + div > pulumi-chooser").should("exist");
                        });
                    });

                    describe("Supporting Types section", () => {
                        describe("type lists", () => {      
                            it("are visible for all languages", () => {
                                cy.get(container).find("section.docs-content h2").as("sections");
                                let pageHeadings = [];
                                cy.get("@sections").then(headings => {
                                    pageHeadings = headings
                                    .map((_, heading) => heading.textContent)
                                    .get();
                                });

                                // Some pages won't contain the `Supporting Types` heading, so we only run the test if it's present.
                                if (pageHeadings.includes("Supporting Types")) {
                                    const languages = [ "TypeScript", "Python", "Go", "C#", "Java", "YAML" ];
                                    const headings = "#supporting-types ~ h4";
                                    const lists = "#supporting-types + h4 ~ div pulumi-choosable div.active";

                                    languages.forEach(language => {
                                        cy.get("pulumi-chooser li a").contains(language).first().click();
                                        cy.log(`Checking ${language} supporting types...`);
                                        cy.get(headings).then(headings => {
                                            cy.get(lists).then(lists => {
                                                expect(lists.length).to.equal(headings.length);
                                            });
                                        });
                                    });
                                }
                            });
                        });
                    });

                    describe("Import section", () => {
                        it("has description text below heading", () => {
                            cy.get(container).find("#import + p").then(p => {
                                expect(p).to.not.be.empty;
                            });
                        });
                    });

                    describe("Package Details section", () => {
                        it("exists", () => {
                            cy.get("#package-details").should("exist");
                        });
                    });
                });

                describe("left column", () => {
                    const container = "#docs-main-nav";

                    it("has an API docs menu", () => {
                        cy.get(container)
                            .find("pulumi-api-doc-filterable-nav")
                            .should("exist");
                    });
                });

                describe("right column", () => {
                    const container = ".docs-table-of-contents .table-of-contents";

                    it("has a package card", () => {
                        cy.get(container)
                            .find("#accordion-table-of-contents")
                            .should("exist");
                    });

                    it("has a table of contents", () => {
                        cy.get(container)
                            .find("#accordion-table-of-contents")
                            .should("exist");
                    });
                });
            });
        });
    });
});
