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
