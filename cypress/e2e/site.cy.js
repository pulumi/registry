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

describe("www.pulumi.com", () => {

    // This should probably be testing the Registry home page and not the home-home page.
    describe("home page", () => {
        beforeEach(() => {
            cy.visit("/");
        });

        it("loads and applies CSS", () => {
            // Checking the computed background-color value validates that the CSS bundle
            // was properly loaded and applied.
            cy.get(".header-container").invoke("css", "background-color").should("equal", "rgb(255, 255, 255)");
        });

        it("loads and applies JavaScript", () => {
            // Checking the carousel validates that the JS bundle was loaded and applied
            // (excluding Stencil components, which are bundled separately).
            cy.get(".header-container").should("not.have.class", "is-pinned");

            cy.scrollTo(0, 250);

            cy.get(".header-container").should("have.class", "is-pinned");
        });
    });
});
