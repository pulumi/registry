// Unit tests for the consent-gated Algolia insights client. From the theme
// directory:
//
//     yarn test

import * as assert from "assert";
import { test } from "node:test";

import { createConsentGatedInsightsClient, createSearchInsightsClient, SEARCH_INSIGHTS_SRC } from "./insights";

const { JSDOM } = require("jsdom");

// A stand-in for runWithPulumiConsent: grant() runs the queued C0002 callbacks.
function setup() {
    const pending: Array<() => void> = [];
    const win = new JSDOM("<!doctype html><body></body>").window;
    win.runWithPulumiConsent = (category: string, cb: () => void) => {
        assert.strictEqual(category, "C0002");
        pending.push(cb);
    };
    (globalThis as any).window = win;
    (globalThis as any).document = win.document;
    const calls: unknown[][] = [];
    const client = Object.assign((...args: unknown[]) => calls.push(args), { version: "2.17.3" });
    const gated = createConsentGatedInsightsClient(client);
    const scripts = () => Array.from(win.document.querySelectorAll("script")).map((s: any) => s.src);
    return { win, gated, calls, scripts, grant: () => pending.splice(0).forEach((cb) => cb()) };
}

test("keeps the real client's version so Autocomplete sends credentials as headers", () => {
    assert.strictEqual(setup().gated.version, "2.17.3");
});

test("sends nothing to Algolia before C0002, and drops events made then", () => {
    const { gated, calls, grant } = setup();
    gated("addAlgoliaAgent", "insights-plugin");
    gated("viewedObjectIDs", { objectIDs: ["a"] });
    gated("clickedObjectIDsAfterSearch", { objectIDs: ["a"], queryID: "q" });
    assert.deepStrictEqual(calls, []);

    grant();
    assert.deepStrictEqual(calls, [["addAlgoliaAgent", "insights-plugin"]]);
});

test("forwards every call once C0002 is granted", () => {
    const { gated, calls, grant } = setup();
    grant();
    gated("viewedObjectIDs", { objectIDs: ["b"] });
    assert.deepStrictEqual(calls, [["viewedObjectIDs", { objectIDs: ["b"] }]]);
});

test("search-insights is only fetched once C0002 is granted, and calls reach the real client after it loads", () => {
    const { win, scripts, grant } = setup();
    const client = createSearchInsightsClient();
    assert.strictEqual(client.version, "2.15.0");
    assert.strictEqual(win.AlgoliaAnalyticsObject, "aa");

    client("getUserToken", null, () => {});
    assert.deepStrictEqual(scripts(), [], "no script before consent");
    assert.strictEqual(win.aa.queue.length, 1, "calls queue on the window.aa stub");

    grant();
    assert.deepStrictEqual(scripts(), [SEARCH_INSIGHTS_SRC]);

    // search-insights replaces window.aa with the real client when it loads.
    const real: unknown[][] = [];
    win.aa = (...args: unknown[]) => real.push(args);
    client("viewedObjectIDs", { objectIDs: ["c"] });
    assert.deepStrictEqual(real, [["viewedObjectIDs", { objectIDs: ["c"] }]]);
});
