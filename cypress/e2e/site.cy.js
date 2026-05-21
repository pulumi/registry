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
            cy.visit("/registry/");
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
});
