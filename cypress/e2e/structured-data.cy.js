describe("Registry", () => {

    // Regression coverage for the JSON-LD structured-data partial at
    // themes/default/layouts/partials/registry/structured-data.html.
    //
    // These tests exist primarily to catch two classes of bug that have
    // shipped before:
    //
    //   1. Double-encoding: if values are piped through `jsonify` without
    //      `safeJS` inside a <script> context, Go's html/template wraps the
    //      already-quoted output in another set of quotes, producing field
    //      values like "\"Registry\"" (literal quote characters inside the
    //      string). Google and LLMs would ingest the broken name.
    //
    //   2. The api-docs section index emitting zero blocks: earlier revisions
    //      read .Page.File.Path, which is nil for dynamically-generated pages
    //      like api-docs/_index.md, silently skipping every JSON-LD block.
    //
    // The assertions below run against a representative package (`aws`) and
    // verify structural correctness and the anti-double-encoding invariant
    // (no string field anywhere contains a literal `"` character).
    describe("Structured data", { testIsolation: false }, () => {

        // Recursively walk a parsed JSON-LD object and return the first string
        // value that contains an embedded literal `"` character, or null if
        // everything is clean. Used as the double-encoding canary.
        const findDoubleEncoded = (value, path = "$") => {
            if (typeof value === "string") {
                return value.includes('"') ? { path, value } : null;
            }
            if (Array.isArray(value)) {
                for (let i = 0; i < value.length; i++) {
                    const hit = findDoubleEncoded(value[i], `${path}[${i}]`);
                    if (hit) return hit;
                }
                return null;
            }
            if (value && typeof value === "object") {
                for (const key of Object.keys(value)) {
                    const hit = findDoubleEncoded(value[key], `${path}.${key}`);
                    if (hit) return hit;
                }
                return null;
            }
            return null;
        };

        // Collect all application/ld+json blocks on the current page, parse
        // them, and hand them off to the caller for assertions.
        const getJsonLdBlocks = () => {
            return cy.get('script[type="application/ld+json"]').then($scripts => {
                return Array.from($scripts).map(s => {
                    try {
                        return JSON.parse(s.textContent);
                    } catch (e) {
                        throw new Error(`Invalid JSON in ld+json block: ${e.message}\n${s.textContent}`);
                    }
                });
            });
        };

        // Registry pages can include an unrelated Organization schema block
        // from the `include_organization_schema` parameter. Filter it out so
        // the tests below are stable even if that schema is toggled elsewhere.
        const filterRegistryBlocks = blocks => {
            return blocks.filter(b => b["@type"] !== "Corporation");
        };

        describe("Package overview page (/registry/packages/aws/)", () => {
            beforeEach(() => {
                cy.visit("/registry/packages/aws/");
            });

            it("emits exactly BreadcrumbList and SoftwareSourceCode", () => {
                getJsonLdBlocks().then(blocks => {
                    const registryBlocks = filterRegistryBlocks(blocks);
                    const types = registryBlocks.map(b => b["@type"]);
                    expect(types).to.include("BreadcrumbList");
                    expect(types).to.include("SoftwareSourceCode");
                });
            });

            it("has no double-encoded string values", () => {
                getJsonLdBlocks().then(blocks => {
                    filterRegistryBlocks(blocks).forEach(block => {
                        const hit = findDoubleEncoded(block);
                        expect(hit, `double-encoded value at ${hit ? hit.path : "none"}`).to.be.null;
                    });
                });
            });

            it("renders the SoftwareSourceCode fields correctly", () => {
                getJsonLdBlocks().then(blocks => {
                    const ssc = filterRegistryBlocks(blocks).find(b => b["@type"] === "SoftwareSourceCode");
                    expect(ssc, "SoftwareSourceCode block").to.exist;
                    expect(ssc.name).to.equal("AWS Pulumi Provider");
                    expect(ssc.alternateName).to.equal("AWS");
                    expect(ssc.publisher).to.deep.include({ "@type": "Organization", name: "Pulumi" });
                    expect(ssc.dateModified, "dateModified should be present and parseable").to.match(/^\d{4}-\d{2}-\d{2}T/);
                    expect(new Date(ssc.dateModified).toString()).to.not.equal("Invalid Date");
                });
            });

            it("renders the BreadcrumbList with unwrapped name values", () => {
                getJsonLdBlocks().then(blocks => {
                    const crumb = filterRegistryBlocks(blocks).find(b => b["@type"] === "BreadcrumbList");
                    expect(crumb, "BreadcrumbList block").to.exist;
                    const names = crumb.itemListElement.map(i => i.name);
                    // Literal equality proves no embedded quote characters.
                    expect(names[0]).to.equal("Registry");
                    expect(names[1]).to.equal("AWS");
                });
            });
        });

        describe("Installation & configuration page (/registry/packages/aws/installation-configuration/)", () => {
            beforeEach(() => {
                cy.visit("/registry/packages/aws/installation-configuration/");
            });

            it("emits BreadcrumbList and HowTo", () => {
                getJsonLdBlocks().then(blocks => {
                    const types = filterRegistryBlocks(blocks).map(b => b["@type"]);
                    expect(types).to.include("BreadcrumbList");
                    expect(types).to.include("HowTo");
                });
            });

            it("has no double-encoded string values", () => {
                getJsonLdBlocks().then(blocks => {
                    filterRegistryBlocks(blocks).forEach(block => {
                        const hit = findDoubleEncoded(block);
                        expect(hit, `double-encoded value at ${hit ? hit.path : "none"}`).to.be.null;
                    });
                });
            });

            it("HowTo.name references the AWS provider literally", () => {
                getJsonLdBlocks().then(blocks => {
                    const howto = filterRegistryBlocks(blocks).find(b => b["@type"] === "HowTo");
                    expect(howto).to.exist;
                    expect(howto.name).to.equal("Install and configure the AWS Pulumi provider");
                });
            });
        });

        describe("API docs section index (/registry/packages/aws/api-docs/)", () => {
            beforeEach(() => {
                cy.visit("/registry/packages/aws/api-docs/");
            });

            // Regression: before the .RelPermalink fix, the api-docs section
            // index emitted zero JSON-LD blocks because .File was nil on that
            // dynamically-generated page.
            it("still emits a BreadcrumbList (regression test for empty-page bug)", () => {
                getJsonLdBlocks().then(blocks => {
                    const types = filterRegistryBlocks(blocks).map(b => b["@type"]);
                    expect(types).to.include("BreadcrumbList");
                });
            });

            it("does not emit a SoftwareSourceCode or TechArticle block", () => {
                getJsonLdBlocks().then(blocks => {
                    const types = filterRegistryBlocks(blocks).map(b => JSON.stringify(b["@type"]));
                    expect(types).to.not.include(JSON.stringify("SoftwareSourceCode"));
                    expect(types).to.not.include(JSON.stringify(["TechArticle", "APIReference"]));
                });
            });

            it("has no double-encoded string values", () => {
                getJsonLdBlocks().then(blocks => {
                    filterRegistryBlocks(blocks).forEach(block => {
                        const hit = findDoubleEncoded(block);
                        expect(hit, `double-encoded value at ${hit ? hit.path : "none"}`).to.be.null;
                    });
                });
            });
        });

        describe("API docs leaf page (/registry/packages/aws/api-docs/s3/bucket/)", () => {
            beforeEach(() => {
                cy.visit("/registry/packages/aws/api-docs/s3/bucket/");
            });

            it("emits BreadcrumbList and TechArticle/APIReference", () => {
                getJsonLdBlocks().then(blocks => {
                    const registryBlocks = filterRegistryBlocks(blocks);
                    const hasBreadcrumb = registryBlocks.some(b => b["@type"] === "BreadcrumbList");
                    const hasTechArticle = registryBlocks.some(b =>
                        Array.isArray(b["@type"]) &&
                        b["@type"].includes("TechArticle") &&
                        b["@type"].includes("APIReference"));
                    expect(hasBreadcrumb, "BreadcrumbList block").to.be.true;
                    expect(hasTechArticle, "TechArticle/APIReference block").to.be.true;
                });
            });

            it("has no double-encoded string values", () => {
                getJsonLdBlocks().then(blocks => {
                    filterRegistryBlocks(blocks).forEach(block => {
                        const hit = findDoubleEncoded(block);
                        expect(hit, `double-encoded value at ${hit ? hit.path : "none"}`).to.be.null;
                    });
                });
            });
        });
    });
});
