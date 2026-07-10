#!/usr/bin/env node
"use strict";

const { test } = require("node:test");
const assert = require("node:assert");
const {
    findMalformedShortcodeDelimiters,
} = require("./check-shortcode-delimiters.js");

test("flags a malformed opening chooser delimiter", () => {
    const content = 'intro\n\n{{ < chooser language "typescript" >}}\nbody\n';
    const hits = findMalformedShortcodeDelimiters(content);
    assert.deepStrictEqual(hits, [
        { line: 3, text: '{{ < chooser language "typescript" >}}' },
    ]);
});

test("flags a malformed opening choosable delimiter", () => {
    const hits = findMalformedShortcodeDelimiters(
        "{{ % choosable language go %}}\n",
    );
    assert.strictEqual(hits.length, 1);
    assert.strictEqual(hits[0].line, 1);
});

test("flags multiple malformed delimiters with correct line numbers", () => {
    const content =
        '{{< chooser language "typescript" >}}\n' +
        '{{ < chooser language "go" >}}\n' +
        "ok\n" +
        "{{ % choosable language python %}}\n";
    const hits = findMalformedShortcodeDelimiters(content);
    assert.deepStrictEqual(
        hits.map((h) => h.line),
        [2, 4],
    );
});

test("does not flag well-formed delimiters", () => {
    const content =
        '{{< chooser language "typescript,python,go,csharp,java,yaml" >}}\n' +
        "{{% choosable language typescript %}}\n" +
        "{{% /choosable %}}\n" +
        "{{< /chooser >}}\n";
    assert.deepStrictEqual(findMalformedShortcodeDelimiters(content), []);
});

test("does not flag a Go template expression in prose", () => {
    assert.deepStrictEqual(
        findMalformedShortcodeDelimiters(
            "Use {{ .Value }} in your template.\n",
        ),
        [],
    );
});

test("flags a malformed closing delimiter", () => {
    const hits = findMalformedShortcodeDelimiters("{{ < /chooser >}}\n");
    assert.strictEqual(hits.length, 1);
});
