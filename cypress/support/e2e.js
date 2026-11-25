// ***********************************************************
// This example support/index.js is processed and
// loaded automatically before your test files.
//
// This is a great place to put global configuration and
// behavior that modifies Cypress.
//
// You can change the location of this file or turn off
// automatically serving support files with the
// 'supportFile' configuration option.
//
// You can read more here:
// https://on.cypress.io/configuration
// ***********************************************************

// Import commands.js using ES2015 syntax:
import "./commands";

// Alternatively you can use CommonJS syntax:
// require('./commands')

// Every build loads the OneTrust consent banner, which covers controls at the
// bottom of the page. Block its scripts so the banner never renders: OneTrust
// picks a banner by the visitor's location, and some of them still showed up
// for runners that had marked the banner as already dismissed.
beforeEach(() => {
    cy.intercept("https://cdn.cookielaw.org/**", { statusCode: 204 });
});
