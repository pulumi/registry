/**
 * Handle uncaught exceptions from third-party scripts.
 * 
 * The vector.co pixel.js script throws an uncaught exception when it detects a cloud provider environment.
 * This happens in the GitHub Actions runner but not locally, causing tests to fail in CI.
 * 
 * We catch and ignore this specific error while still allowing other legitimate errors to fail the tests.
 */
Cypress.on('uncaught:exception', (err) => {
    // Return false to prevent the error from failing the test
    if (err.message.includes('Cloud provider detected')) {
        return false;
    }
    // Return true for other errors to fail the test
    return true;
});

describe("www.pulumi.com/registry", () => {

    describe("registry home page", () => {
        beforeEach(() => {
            // Pin the color theme to light. The default preference is "system", so
            // without this the assertion below depends on the OS appearance of
            // whatever machine is running the suite.
            cy.visit("/registry/", {
                onBeforeLoad(win) {
                    win.localStorage.setItem("pulumi-docs-theme", "light");
                },
            });
        });

        it("loads and applies CSS", () => {
            // Checking the computed background-color value validates that the CSS bundle
            // was properly loaded and applied.
            cy.get("header").invoke("css", "background-color").should("equal", "rgb(255, 255, 255)");
        });

        it("loads and applies JavaScript", () => {
            // docs-main.ts sets an inline height on .docs-main-nav when the bundle runs,
            // so an inline style here proves the JS bundle was loaded and applied
            // (excluding Stencil components, which are bundled separately).
            cy.get(".docs-main-nav").should("have.attr", "style").and("match", /height:\s*\d+px/);
        });
    });

    describe("package list filter", () => {
        beforeEach(() => {
            cy.visit("/registry/");
        });

        const filterBy = (query) => {
            cy.get(".registry-filter-input").clear().type(query);
            // The search component debounces its event. The random package matches
            // none of the queries below, so wait until it is hidden before asserting
            // on the result.
            cy.get(".all-packages .package.hidden [data-name='random']").should("exist");
        };

        const expectShown = (name) => {
            cy.get(`.all-packages .package:not(.hidden) [data-name='${name}']`).should("exist");
        };

        it("matches the package name", () => {
            filterBy("scm");
            expectShown("scm");
        });

        it("matches every token against the package keywords", () => {
            filterBy("palo alto");
            expectShown("scm");
            filterBy("alto palo");
            expectShown("scm");
        });

        it("matches the title after the native hack strips the word", () => {
            filterBy("native aws");
            expectShown("aws-native");
            expectShown("aws");
        });

        it("treats amazon as a synonym for aws", () => {
            filterBy("amazon");
            expectShown("aws");
            expectShown("eks");
        });

        it("hides every package when nothing matches", () => {
            filterBy("no-such-package");
            cy.get(".all-packages .package:not(.hidden)").should("have.length", 0);
        });
    });

    // When the filter comes up empty, we look the query up in the OpenTofu registry
    // and offer `pulumi package add terraform-provider ...` for what we find, since
    // Pulumi can run any Terraform/OpenTofu provider. See issue #7299.
    describe("no-results OpenTofu suggestions", () => {
        const TOFU_API = "https://api.opentofu.org/registry/docs/search*";

        // Deliberately returned out of rank order, and with a non-provider row, so
        // the assertions below prove we re-sort and filter rather than echoing the API.
        const twoProviders = [
            { type: "provider", addr: "nobbs/sops", version: "v0.3.3", popularity: 18 },
            { type: "module", addr: "someone/sops/aws", version: "v1.0.0", popularity: 99999 },
            { type: "provider/resource", addr: "carlpett/sops", version: "v1.4.1", popularity: 99999 },
            { type: "provider", addr: "carlpett/sops", version: "v1.4.1", popularity: 584 },
        ];

        const stubTofu = (response) => {
            cy.intercept("GET", TOFU_API, response).as("tofu");
        };

        // Types into the filter and waits for the empty state to actually be reached,
        // so assertions don't race the search component's 300ms debounce. The
        // zero-visible-packages assertion is retried, and it can't pass until
        // filterByTextAndTags -> updateAllCount -> syncEmptyState has run, which is
        // what makes the "did we fetch?" assertions below deterministic.
        const filterToEmpty = (query) => {
            cy.get(".registry-filter-input").clear().type(query);
            cy.get(".all-packages .package:not(.hidden)").should("have.length", 0);
        };

        // Ticks a tag filter. The option checkboxes live in shadow DOM behind an
        // opacity-transitioned menu, so this drives pulumi-filter-select's own
        // select() @Method() instead. The should() retries until Stencil has
        // upgraded the element and put the method on it.
        const selectFilter = (value) => {
            cy.get("pulumi-filter-select")
                .should(($el) => expect($el[0].select).to.be.a("function"))
                .then(($el) => $el[0].select({ value }));
        };

        beforeEach(() => {
            cy.visit("/registry/");
        });

        it("ranks providers by stars, not by API order", () => {
            stubTofu({ body: twoProviders });
            filterToEmpty("sops");
            cy.wait("@tofu");

            cy.get(".tf-suggest .tf-suggest-item").should("have.length", 2);
            cy.get(".tf-suggest .tf-suggest-item")
                .first()
                .should("contain.text", "carlpett/sops");
        });

        it("drops non-provider result types", () => {
            stubTofu({ body: twoProviders });
            filterToEmpty("sops");
            cy.wait("@tofu");

            // The module and the provider/resource rows must not appear.
            cy.get(".tf-suggest").should("not.contain.text", "someone/sops/aws");
            cy.get(".tf-suggest .tf-suggest-item").should("have.length", 2);
        });

        it("pins the version in the command, without the leading v", () => {
            stubTofu({ body: twoProviders });
            filterToEmpty("sops");
            cy.wait("@tofu");

            cy.get(".tf-suggest .tf-suggest-item")
                .first()
                .find("code")
                .should(
                    "have.text",
                    "$ pulumi package add terraform-provider carlpett/sops 1.4.1",
                );
        });

        it("collapses repeated rows for the same provider", () => {
            // The endpoint looks like one row per documentation page, so a provider
            // could plausibly return several type:"provider" rows and eat several of
            // the five slots. Only the best-ranked row for each addr survives.
            stubTofu({
                body: [
                    { type: "provider", addr: "carlpett/sops", version: "v1.4.1", popularity: 584 },
                    { type: "provider", addr: "carlpett/sops", version: "v1.3.0", popularity: 584 },
                    { type: "provider", addr: "nobbs/sops", version: "v0.3.3", popularity: 18 },
                ],
            });
            filterToEmpty("sops");
            cy.wait("@tofu");

            cy.get(".tf-suggest .tf-suggest-item").should("have.length", 2);
            cy.get(".tf-suggest .tf-suggest-item")
                .first()
                .find("code")
                .should("contain.text", "carlpett/sops 1.4.1");
        });

        it("announces the search only once the lookup is visibly slow", () => {
            // The loading heading is deferred so a lookup that resolves at typing
            // speed never paints an intermediate state into the aria-live region.
            // A slow one still says what it is doing.
            cy.intercept("GET", TOFU_API, { body: twoProviders, delay: 1500 }).as("tofu");
            filterToEmpty("sops");

            cy.get(".tf-suggest .tf-suggest-message").should(
                "contain.text",
                "Searching the OpenTofu registry",
            );

            cy.wait("@tofu");
            cy.get(".tf-suggest .tf-suggest-message").should(
                "contain.text",
                "2 providers in the OpenTofu registry",
            );
        });

        it("prefers the vendor's own namespace when stars tie", () => {
            stubTofu({
                body: [
                    { type: "provider", addr: "adamdecaf/namecheap", version: "v2.9.2", popularity: 168 },
                    { type: "provider", addr: "namecheap/namecheap", version: "v2.9.2", popularity: 168 },
                ],
            });
            filterToEmpty("namecheap");
            cy.wait("@tofu");

            cy.get(".tf-suggest .tf-suggest-item")
                .first()
                .should("contain.text", "namecheap/namecheap");
        });

        it("links each provider to the OpenTofu registry in a new tab", () => {
            stubTofu({ body: twoProviders });
            filterToEmpty("sops");
            cy.wait("@tofu");

            cy.get(".tf-suggest .tf-suggest-item")
                .first()
                .find("a")
                .should(
                    "have.attr",
                    "href",
                    "https://search.opentofu.org/provider/carlpett/sops/v1.4.1",
                )
                .and("have.attr", "target", "_blank")
                .and("have.attr", "rel")
                .and("include", "noopener");
        });

        it("falls back to the docs link when OpenTofu returns nothing", () => {
            stubTofu({ body: [] });
            filterToEmpty("zzzznope");
            cy.wait("@tofu");

            cy.get(".tf-suggest .tf-suggest-item").should("not.exist");
            cy.get(".no-results .tf-suggest-fallback a")
                .should("be.visible")
                .and(
                    "have.attr",
                    "href",
                    "/docs/iac/concepts/providers/any-terraform-provider/",
                );
        });

        it("falls back to the docs link when the API fails", () => {
            cy.intercept("GET", TOFU_API, { forceNetworkError: true }).as("tofu");
            filterToEmpty("zzzznope");

            cy.get(".tf-suggest .tf-suggest-item").should("not.exist");
            cy.get(".no-results .tf-suggest-fallback a").should("be.visible");
        });

        it("drops entries whose addr is not a plain namespace/name", () => {
            stubTofu({
                body: [
                    { type: "provider", addr: "evil/<img src=x onerror=alert(1)>", version: "v1.0.0", popularity: 99999 },
                    { type: "provider", addr: "carlpett/sops", version: "v1.4.1", popularity: 1 },
                ],
            });
            filterToEmpty("sops");
            cy.wait("@tofu");

            cy.get(".no-results img").should("not.exist");
            cy.get(".tf-suggest .tf-suggest-item").should("have.length", 1);
        });

        it("attributes each message to its own registry, split by a rule", () => {
            stubTofu({ body: twoProviders });
            filterToEmpty("sops");
            cy.wait("@tofu");

            // Pulumi's message names Pulumi; OpenTofu's names OpenTofu and counts.
            cy.get(".no-results .no-results-message")
                .should("contain.text", "Pulumi Registry")
                .and("contain.text", "sops");
            cy.get(".tf-suggest .tf-suggest-message")
                .should("contain.text", "2 providers in the OpenTofu registry")
                .and("contain.text", "sops");
            cy.get(".tf-suggest hr.tf-suggest-divider").should("exist");
        });

        it("uses the singular message for a single provider", () => {
            stubTofu({
                body: [
                    { type: "provider", addr: "carlpett/sops", version: "v1.4.1", popularity: 584 },
                ],
            });
            filterToEmpty("sops");
            cy.wait("@tofu");

            cy.get(".tf-suggest .tf-suggest-message").should(
                "contain.text",
                "1 provider in the OpenTofu registry matches",
            );
        });

        it("says so explicitly when OpenTofu has no match either", () => {
            stubTofu({ body: [] });
            filterToEmpty("zzzznope");
            cy.wait("@tofu");

            cy.get(".tf-suggest .tf-suggest-message").should(
                "contain.text",
                "No providers in the OpenTofu registry match",
            );
            cy.get(".tf-suggest hr.tf-suggest-divider").should("exist");
        });

        it("does not query OpenTofu for a single-character query", () => {
            cy.intercept("GET", TOFU_API, { body: [] }).as("tofu");

            // "~" appears in no package name, title, or keyword, so the empty state
            // is genuinely reached and MIN_QUERY_LENGTH is what suppresses the
            // lookup. A letter would not test that: "q" alone matches equinix,
            // mssql, mysql, postgresql, qdrant-cloud and five more, so the panel
            // would stay CSS-hidden and syncEmptyState would bail on visibleCount
            // long before reaching the guard.
            filterToEmpty("~");

            cy.get(".tf-suggest").should("be.empty");
            cy.get("@tofu.all").should("have.length", 0);
        });

        it("does not query OpenTofu when only tag filters are applied", () => {
            cy.intercept("GET", TOFU_API, { body: [] }).as("tofu");

            // Deprecated + Version Control: both version-control packages are
            // current, so the pair yields an empty list while the search box stays
            // empty -- there is no query to send.
            selectFilter("deprecated");
            selectFilter("version control system");

            cy.get(".all-packages .package:not(.hidden)").should("have.length", 0);
            cy.get(".registry-filter-input").should("have.value", "");
            cy.get(".tf-suggest").should("be.empty");
            cy.get("@tofu.all").should("have.length", 0);
        });

        it("keeps Clear all filters working once suggestions are rendered", () => {
            stubTofu({ body: twoProviders });
            filterToEmpty("sops");
            cy.wait("@tofu");
            cy.get(".tf-suggest .tf-suggest-item").should("have.length", 2);

            cy.get(".no-results .reset").click();

            cy.get(".all-packages .package:not(.hidden)").should("have.length.greaterThan", 0);
            cy.get(".tf-suggest").should("be.empty");
        });
    });
});
