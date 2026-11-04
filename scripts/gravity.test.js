// Tests for themes/default/layouts/partials/gravity.html: the Gravity pixel must load
// only through the shared OneTrust gate (runWithPulumiConsent "C0004"), after window.load, once.
// Run: make test-gravity
const { test } = require("node:test");
const assert = require("node:assert/strict");
const fs = require("node:fs");
const path = require("node:path");
const vm = require("node:vm");

const source = fs
    .readFileSync(path.join(__dirname, "../themes/default/layouts/partials/gravity.html"), "utf8")
    .replace(/\{\{-? \/\*[\s\S]*?\*\/ -?\}\}/, "")
    .replace(/<\/?script>/g, "");

function harness() {
    const listeners = {}, scripts = [], gates = [];
    let reloads = 0;
    const window = {
        location: { reload() { reloads++; } },
        addEventListener(name, fn) { listeners[name] = fn; },
        runWithPulumiConsent(category, fn) { gates.push({ category, fn }); },
    };
    window.window = window;
    const document = {
        createElement() { return {}; },
        getElementsByTagName() { return [{ parentNode: { insertBefore(s) { scripts.push(s); } } }]; },
    };
    assert.ok(!source.includes("{{"), "unrendered template tags left");
    vm.runInContext(source, vm.createContext({ window, document, Date }));
    return { window, listeners, scripts, gates, get reloads() { return reloads; } };
}

test("nothing runs before window.load and no direct OneTrust listener", () => {
    const h = harness();
    assert.equal(h.gates.length, 0);
    assert.equal(h.scripts.length, 0);
    assert.ok(!("OneTrustGroupsUpdated" in h.listeners));
});

test("after load it waits on the C0004 gate", () => {
    const h = harness();
    h.listeners.load();
    assert.equal(h.gates.length, 1);
    assert.equal(h.gates[0].category, "C0004");
    assert.equal(h.scripts.length, 0);
});

test("granting C0004 injects the pixel once with the advertiser id", () => {
    const h = harness();
    h.listeners.load();
    h.gates[0].fn();
    h.gates[0].fn();
    assert.equal(h.scripts.length, 1);
    assert.equal(h.scripts[0].src, "https://code.trygravity.com/gr-pix.js");
    assert.deepEqual(Array.from(h.window.gravity.q[0]), ["init", "f9505c8e-5d0b-4159-88b8-53bd2d5f2c8a"]);
    assert.equal(h.reloads, 0);
});
