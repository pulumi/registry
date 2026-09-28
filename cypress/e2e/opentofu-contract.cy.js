// Contract check against the live OpenTofu search API that backs the no-results
// suggestions in themes/default/theme/src/ts/opentofu-suggest.ts.
//
// site.cy.js stubs this endpoint with cy.intercept, so it can never notice the API
// changing shape. This spec hits the real endpoint instead, and runs on a schedule
// (run-browser-tests.yml) rather than on PRs, since its failures are about OpenTofu,
// not about the change under review.
//
// The field that matters most is `popularity` (GitHub stars). It's the only signal
// that separates a real provider from its zero-star forks, and if it disappears,
// `Number(undefined)` is NaN, every provider scores 0, and the ranking silently
// degrades to alphabetical. Nothing in the UI fails when that happens.
const API = "https://api.opentofu.org/registry/docs/search?q=";

// Mirrors ADDR_RE in opentofu-suggest.ts.
const ADDR_RE = /^[A-Za-z0-9][A-Za-z0-9._-]*\/[A-Za-z0-9][A-Za-z0-9._-]*$/;

describe("OpenTofu search API contract", () => {
    let providers;

    before(() => {
        cy.request(`${API}random`).then(response => {
            expect(response.status).to.eq(200);
            expect(response.body).to.be.an("array").that.is.not.empty;
            providers = response.body.filter(r => r && r.type === "provider");
        });
    });

    it("returns provider rows", () => {
        expect(providers).to.not.be.empty;
    });

    it("gives every provider a plain namespace/name addr", () => {
        providers.forEach(p => {
            expect(p.addr, JSON.stringify(p)).to.be.a("string").and.match(ADDR_RE);
        });
    });

    it("gives every provider a string version", () => {
        providers.forEach(p => {
            expect(p.version, p.addr).to.be.a("string").and.not.be.empty;
        });
    });

    it("gives every provider a numeric popularity", () => {
        providers.forEach(p => {
            expect(p.popularity, p.addr).to.be.a("number").and.satisfy(Number.isFinite);
        });
    });

    it("ranks hashicorp/random above its forks by popularity", () => {
        const official = providers.find(p => p.addr === "hashicorp/random");
        expect(official, "hashicorp/random in results").to.exist;
        expect(official.popularity).to.be.greaterThan(0);

        const top = Math.max(...providers.map(p => p.popularity));
        expect(official.popularity).to.eq(top);
    });
});
