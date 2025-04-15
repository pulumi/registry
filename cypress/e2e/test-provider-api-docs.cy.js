describe("Test Provider", () => {

    [
        // Resource page that uses <!--Start PulumiCodeChooser --> fences
        {
            path: "/registry/packages/test-provider/api-docs/testresourcecommentfence/",
            resourceName: "TestResourceCommentFence"
        },
        // Resource page that uses {{% examples %}} fences
        {
            path: "/registry/packages/test-provider/api-docs/testresourceexamplesfence/",
            resourceName: "TestResourceExamplesFence"
        }
    ].forEach(p => {
        describe(`${p.resourceName} API Docs`, { testIsolation: false }, () => {
            const page = {
                title: `testprovider.${p.resourceName}`,
                provider: "test-provider",
                resourceName: p.resourceName,
                path: p.path,
            };
            const container = "main div.docs-main-content";

            beforeEach(() => {
                cy.visit(page.path);
            });

            it("has the correct page title", () => {
                cy.get("head title").should("include.text", page.title).and("include.text", "| Pulumi Registry");
            });

            it("has the correct H1", () => {
                cy.get(container)
                    .find("h1")
                    .should("be.visible")
                    .should("have.text", page.title);
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

            // Verify that the page renders the set of subsections in the correct order
            it("renders the correct set of subsections, in the correct order", () => {
                cy.get(container).find("section.docs-content h2").as("sections");
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
                    const expectedHeadings = possibleHeadings.filter(heading => actualHeadings.includes(heading));
                    expect(JSON.stringify(actualHeadings)).to.equal(JSON.stringify(expectedHeadings));
                });
            });

            // Verify that the examples section exists and contains examples for all languages
            describe("Examples section", () => {
                // Verify that the examples section exists
                it("contains the example section", () => {
                    cy.get(container)
                        .find("h2")
                        .contains("Example Usage")
                        .should("exist");
                });

                // Verify that the examples section contains examples for all languages
                it("contains examples for all languages", () => {
                    const languages = ["TypeScript", "Python", "Go", "C#", "Java", "YAML"];

                    // First verify the language chooser exists and has all expected languages
                    cy.get("#example-usage + div > pulumi-chooser").should("exist")
                        .invoke('attr', 'options')
                        .should('include', 'typescript')
                        .and('include', 'python')
                        .and('include', 'go')
                        .and('include', 'csharp')
                        .and('include', 'java')
                        .and('include', 'yaml');

                    // Then check each an example exists for each language
                    languages.forEach(language => {
                        cy.get("pulumi-chooser li a").contains(language).first().click();

                        // Find the appropriate pulumi-choosable based on language
                        let selector;
                        if (language === "TypeScript") {
                            selector = "div pulumi-choosable[type='language'][values*='typescript']";
                        } else if (language === "C#") {
                            selector = "div pulumi-choosable[type='language'][values*='csharp']";
                        } else {
                            selector = `div pulumi-choosable[type='language'][values*='${language.toLowerCase()}']`;
                        }

                        // Verify both the section and its code example exist
                        cy.get(`#example-usage ~ ${selector}`).should("exist");
                        cy.get(`#example-usage ~ ${selector} pre`).should("exist")
                            .and("not.be.empty")
                            .invoke('text')
                            .should('have.length.gt', 0);
                    });
                });
            });

            // Verify that the inputs and outputs sections exist
            describe("Inputs and Outputs sections", () => {
                const propertyLists = ".resources-properties";

                it("renders deprecation notices properly (if any)", () => {
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

                // Verify that the inputs section contains the properties for all languages
                it("contains input types for all languages", () => {
                    const languages = ["TypeScript", "Python", "Go", "C#", "Java", "YAML"];

                    // First verify the inputs section exists and has the language chooser with all languages
                    cy.get("#inputs").should("exist");

                    // Then check each language
                    languages.forEach(language => {
                        cy.get("pulumi-chooser li a").contains(language).first().click();

                        // Get the appropriate pulumi-choosable based on language
                        let selector;
                        if (language === "TypeScript") {
                            selector = "div pulumi-choosable[type='language'][values*='typescript']";
                        } else if (language === "C#") {
                            selector = "div pulumi-choosable[type='language'][values*='csharp']";
                        } else {
                            selector = `div pulumi-choosable[type='language'][values*='${language.toLowerCase()}']`;
                        }

                        // Verify the language section exists
                        cy.get(`#inputs ~ ${selector}`).should("exist");

                        // Verify it contains property definitions
                        cy.get(`#inputs ~ ${selector} ${propertyLists}`).should("exist")
                            .find("dt").should("have.length.at.least", 1)
                            .first().should("exist");

                        // Verify each property has a type and description
                        cy.get(`#inputs ~ ${selector} ${propertyLists} dt`).each(($dt) => {
                            cy.wrap($dt)
                                .find(".property-type").should("exist")
                                .invoke('text')
                                .should('have.length.gt', 0);
                            cy.wrap($dt)
                                .next("dd").should("exist")
                                .invoke('text')
                                .should('have.length.gt', 0);
                        });
                    });
                });

                // Verify that the type links point to an on-page supporting type if there are any.
                describe("type links", () => {
                    it("all point to an on-page supporting type", () => {
                        cy.get(container).find(propertyLists).find("dt .property-type a[href*='#']").then(links => {
                            links.each((i, link) => {
                                const href = link.getAttribute('href');
                                if (href && href.startsWith('#')) {
                                    cy.get("#supporting-types").should("exist");
                                    cy.get(href).should("exist");
                                }
                            });
                        });
                    });
                });
            });

            // Verify that the Supporting Types section exists and contains the correct properties for all languages
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

                        if (pageHeadings.includes("Supporting Types")) {
                            const languages = ["TypeScript", "Python", "Go", "C#", "Java", "YAML"];

                            // First verify the language chooser exists and has all languages
                            cy.get("#supporting-types ~ div > pulumi-chooser").should("exist")
                                .invoke('attr', 'options')
                                .should('include', 'typescript')
                                .and('include', 'python')
                                .and('include', 'go')
                                .and('include', 'csharp')
                                .and('include', 'java')
                                .and('include', 'yaml');

                            languages.forEach(language => {
                                cy.get("pulumi-chooser li a").contains(language).first().click();
                                cy.log(`Checking ${language} supporting types...`);

                                // Get the appropriate pulumi-choosable based on language
                                let selector;
                                if (language === "TypeScript") {
                                    selector = "div pulumi-choosable[type='language'][values*='typescript']";
                                } else if (language === "C#") {
                                    selector = "div pulumi-choosable[type='language'][values*='csharp']";
                                } else {
                                    selector = `div pulumi-choosable[type='language'][values*='${language.toLowerCase()}']`;
                                }

                                // Verify the language section exists
                                cy.get(`#supporting-types ~ ${selector}`).should("exist");

                                // Verify it contains property definitions
                                cy.get(`#supporting-types ~ ${selector} ${propertyLists}`).should("exist")
                                    .find("dt").should("have.length.at.least", 1)
                                    .first().should("exist");

                                // Verify each property has a type and description
                                cy.get(`#supporting-types ~ ${selector} ${propertyLists} dt`).each(($dt) => {
                                    cy.wrap($dt)
                                        .find(".property-type").should("exist")
                                        .invoke('text')
                                        .should('have.length.gt', 0);
                                    cy.wrap($dt)
                                        .next("dd").should("exist")
                                        .invoke('text')
                                        .should('have.length.gt', 0);
                                });
                            });
                        }
                    });
                });
            });
        });
    });
});