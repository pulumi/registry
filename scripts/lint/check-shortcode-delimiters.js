#!/usr/bin/env node
"use strict";

/**
 * A stray space in a shortcode delimiter ("{{ <" instead of "{{<") aborts the
 * entire Hugo build. The front-matter linter skips the auto-generated provider
 * `_index.md` files where this tends to appear, so scan every markdown file.
 * Mirrors resourcedocsgen's sanitizeShortcodeDelimiters guard.
 */

const fs = require("fs");
const path = require("path");

const MALFORMED_DELIMITER = /\{\{[ \t]+[<%]/;

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
        const hits = findMalformedShortcodeDelimiters(
            fs.readFileSync(file, "utf8"),
        );
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
