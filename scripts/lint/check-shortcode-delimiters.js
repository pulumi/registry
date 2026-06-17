#!/usr/bin/env node
"use strict";

/**
 * Fails if any registry content contains a malformed Hugo shortcode opening
 * delimiter, e.g. "{{ < chooser" or "{{ % choosable" instead of "{{<" / "{{%".
 *
 * Hugo requires the marker ("<" for HTML/block shortcodes, "%" for markdown
 * shortcodes) to immediately follow "{{". A stray space makes Hugo stop
 * recognizing the opening tag while still recognizing the paired closing tag,
 * which aborts the ENTIRE site build:
 *
 *     shortcode "X" does not evaluate .Inner or .InnerDeindent, yet a closing
 *     tag was provided
 *
 * Because a single malformed page is fatal, and because the auto-generated
 * provider `_index.md` files (which carry the "fetched from" header) are skipped
 * by the front-matter linter, this dedicated check scans every markdown file so
 * the failure surfaces as a clear, local lint error instead of a cryptic
 * whole-site build abort. The matching mirrors resourcedocsgen's
 * sanitizeShortcodeDelimiters guard.
 */

const fs = require("fs");
const path = require("path");

// Matches "{{" followed by one or more spaces/tabs and then a "<" or "%" marker.
// Narrow on purpose: a space after "{{" followed by anything else (e.g. a Go
// template expression "{{ .Value }}") is not matched, and the closing side
// (the legitimate space in `... >}}`) is never touched.
const MALFORMED_DELIMITER = /\{\{[ \t]+[<%]/;

/**
 * Returns the malformed shortcode delimiters found in the given file content.
 * @param {string} content
 * @returns {{line: number, text: string}[]}
 */
function findMalformedShortcodeDelimiters(content) {
    const hits = [];
    const lines = content.split("\n");
    for (let i = 0; i < lines.length; i++) {
        if (MALFORMED_DELIMITER.test(lines[i])) {
            hits.push({ line: i + 1, text: lines[i].trim() });
        }
    }
    return hits;
}

/**
 * Recursively collects all `.md` files under a directory.
 * @param {string} dir
 * @returns {string[]}
 */
function collectMarkdownFiles(dir) {
    const out = [];
    for (const entry of fs.readdirSync(dir, { withFileTypes: true })) {
        const full = path.join(dir, entry.name);
        if (entry.isDirectory()) {
            out.push(...collectMarkdownFiles(full));
        } else if (entry.isFile() && entry.name.endsWith(".md")) {
            out.push(full);
        }
    }
    return out;
}

function main() {
    const contentDir = path.resolve(__dirname, "../../themes/default/content");
    const files = collectMarkdownFiles(contentDir);

    const problems = [];
    for (const file of files) {
        const hits = findMalformedShortcodeDelimiters(fs.readFileSync(file, "utf8"));
        for (const hit of hits) {
            problems.push(`${file}:${hit.line}: ${hit.text}`);
        }
    }

    if (problems.length > 0) {
        console.error(
            `\nMalformed Hugo shortcode delimiter check:\n` +
                `    - ${files.length} files scanned.\n` +
                `    - ${problems.length} malformed delimiter(s) found ` +
                `(expected "{{<" / "{{%", not "{{ <" / "{{ %"):\n\n` +
                problems.map((p) => `  ${p}`).join("\n") +
                `\n`,
        );
        process.exit(1);
    }

    console.log(
        `\nMalformed Hugo shortcode delimiter check:\n` +
            `    - ${files.length} files scanned.\n` +
            `    - 0 malformed delimiters found.\n`,
    );
}

if (require.main === module) {
    main();
}

module.exports = { findMalformedShortcodeDelimiters };
