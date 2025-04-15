describe("Test Provider API docs", { testIsolation: false }, () => {
    const page = {
        title: "testprovider.TestResource",
        provider: "test-provider",
        resourceName: "TestResource",
        path: "/registry/packages/test-provider/api-docs/testresource/",
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

    describe("Examples section", () => {
        it("contains at least one example", () => {
            cy.get(container)
                .find("h2 ~ h3")
                .should("have.length.of.at.least", 1);
        });

        it("contains examples for all languages", () => {
            const languages = ["TypeScript", "Python", "Go", "C#", "Java", "YAML"];
            
            // First verify the language chooser exists
            cy.get("#example-usage + div pulumi-chooser").should("exist");
            
            // Then check each language
            languages.forEach(language => {
                cy.get("pulumi-chooser li a").contains(language).first().click();
                
                // Get the appropriate pulumi-choosable based on language
                let selector;
                if (language === "TypeScript") {
                    selector = "div pulumi-choosable[type='language'][values*='javascript']";
                } else if (language === "C#") {
                    selector = "div pulumi-choosable[type='language'][values*='csharp']";
                } else {
                    selector = `div pulumi-choosable[type='language'][values*='${language.toLowerCase()}']`;
                }
                
                cy.get(`#example-usage ~ ${selector}`).should("exist")
                    .find("pre").should("exist")
                    .and("not.be.empty");
            });
        });
    });

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

        it("contains input types for all languages", () => {
            const languages = ["TypeScript", "Python", "Go", "C#", "Java", "YAML"];
            
            // First verify the inputs section exists
            cy.get("#inputs").should("exist");
            
            // Then check each language
            languages.forEach(language => {
                cy.get("pulumi-chooser li a").contains(language).first().click();
                
                // Get the appropriate pulumi-choosable based on language
                let selector;
                if (language === "TypeScript") {
                    selector = "div pulumi-choosable[type='language'][values*='javascript']";
                } else if (language === "C#") {
                    selector = "div pulumi-choosable[type='language'][values*='csharp']";
                } else {
                    selector = `div pulumi-choosable[type='language'][values*='${language.toLowerCase()}']`;
                }
                
                // Verify the language section exists and contains property definitions
                cy.get(`#inputs ~ ${selector}`).should("exist")
                    .find(propertyLists).should("exist")
                    .find("dt").should("have.length.at.least", 1)
                    .first().should("exist");
                
                // Verify each property has a type and description
                cy.get(`#inputs ~ ${selector} ${propertyLists} dt`).each(($dt) => {
                    cy.wrap($dt)
                        .find(".property-type").should("exist")
                        .and("not.be.empty");
                    cy.wrap($dt)
                        .next("dd").should("exist")
                        .and("not.be.empty");
                });
            });
        });

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
}); 