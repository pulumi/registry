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
});
