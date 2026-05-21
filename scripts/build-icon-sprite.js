#!/usr/bin/env node
/**
 * build-icon-sprite.js
 *
 * Builds an SVG sprite (themes/default/assets/icons/sprite.svg) plus a JSON
 * manifest (themes/default/assets/icons/sprite-manifest.json) from the per-
 * icon SVGs under themes/default/assets/icons/{phosphor,custom,brand}/.
 *
 * The set of icons included is determined by tokenising the source tree
 * (layouts, content, data, archetypes within themes/default) and intersecting
 * the resulting tokens with the available icon basenames. We err on the side
 * of including too many icons rather than too few — false positives are
 * cheap (the sprite is a single fingerprinted asset) and a missing icon
 * breaks rendering.
 *
 * Symbol id scheme:
 *   p-{name}-{weight}  → phosphor (4 weights)
 *   c-{name}-{weight}  → custom   (4 weights; weights without a dedicated
 *                                  file fall back to the base SVG content)
 *   b-{name}           → brand    (single weight)
 *
 * The icon source SVGs under phosphor/ and brand/ are refreshed from
 * @phosphor-icons/core and simple-icons by scripts/sync-icons.js. Run
 * `make sync-icons` followed by `make build-icon-sprite` after bumping
 * those upstream packages.
 */

"use strict";

const fs = require("fs");
const path = require("path");

const ROOT = path.resolve(__dirname, "..");
const THEME_ROOT = path.join(ROOT, "themes/default");
const ICONS_DIR = path.join(THEME_ROOT, "assets/icons");
const PHOSPHOR_DIR = path.join(ICONS_DIR, "phosphor");
const CUSTOM_DIR = path.join(ICONS_DIR, "custom");
const BRAND_DIR = path.join(ICONS_DIR, "brand");
const SPRITE_PATH = path.join(ICONS_DIR, "sprite.svg");
const MANIFEST_PATH = path.join(ICONS_DIR, "sprite-manifest.json");

const PHOSPHOR_WEIGHTS = ["regular", "bold", "fill", "duotone"];
const CUSTOM_WEIGHTS = ["regular", "bold", "fill", "duotone"];

// Directories whose contents define which icons are referenced. theme/src is
// scanned so JS/TS/SCSS files that reference symbol IDs directly (e.g.
// copy-llm-prompt.ts swapping between #p-copy-* and #p-check-*) are picked up.
const SCAN_DIRS = [
    path.join(THEME_ROOT, "layouts"),
    path.join(THEME_ROOT, "content"),
    path.join(THEME_ROOT, "data"),
    path.join(THEME_ROOT, "archetypes"),
    path.join(THEME_ROOT, "theme/src"),
];

const SCAN_EXTS = new Set([
    ".html",
    ".md",
    ".yml",
    ".yaml",
    ".toml",
    ".json",
    ".ts",
    ".tsx",
    ".js",
    ".scss",
]);

// Always-include list for icons referenced via patterns the tokeniser can't
// see: names built at runtime by string concatenation, names that collide with
// common words and get filtered out, or icons used by code outside SCAN_DIRS.
// Add a name here if a build error reports a missing icon symbol.
const FORCE_INCLUDE = new Set([
    // Brand icons (small list; some names like "x" are short common words).
]);

// ─── Utilities ───────────────────────────────────────────────────────────────

function listSvgs(dir) {
    if (!fs.existsSync(dir)) return [];
    return fs.readdirSync(dir).filter((f) => f.endsWith(".svg"));
}

function basenameNoExt(file) {
    return file.replace(/\.svg$/, "");
}

function stripWeightSuffix(name) {
    return name.replace(/-(bold|fill|duotone)$/, "");
}

function* walk(dir) {
    if (!fs.existsSync(dir)) return;
    for (const entry of fs.readdirSync(dir, { withFileTypes: true })) {
        const full = path.join(dir, entry.name);
        if (entry.isDirectory()) {
            yield* walk(full);
        } else if (entry.isFile() && SCAN_EXTS.has(path.extname(entry.name))) {
            yield full;
        }
    }
}

function tokenise(text) {
    const out = new Set();
    const re = /[a-z][a-z0-9]*(?:-[a-z0-9]+)*/g;
    const lower = text.toLowerCase();
    let m;
    while ((m = re.exec(lower)) !== null) {
        out.add(m[0]);
    }
    return out;
}

function buildTokenSet() {
    const tokens = new Set();
    for (const dir of SCAN_DIRS) {
        for (const file of walk(dir)) {
            const text = fs.readFileSync(file, "utf8");
            for (const tok of tokenise(text)) {
                tokens.add(tok);
            }
        }
    }
    return tokens;
}

function extractSvg(svg) {
    const vbMatch = svg.match(/viewBox="([^"]+)"/);
    const viewBox = vbMatch ? vbMatch[1] : "0 0 256 256";

    let inner = svg.replace(/^[\s\S]*?<svg\b[^>]*>/, "");
    inner = inner.replace(/<\/svg>\s*$/, "");
    return { inner: inner.trim(), viewBox };
}

function makeSymbol(id, viewBox, inner) {
    return `<symbol id="${id}" viewBox="${viewBox}">${inner}</symbol>`;
}

// ─── Build ───────────────────────────────────────────────────────────────────

function build() {
    console.log("=== build-icon-sprite ===");

    const tokens = buildTokenSet();
    console.log(`  scanned tokens: ${tokens.size}`);

    const phosphorAvailable = new Set(
        listSvgs(path.join(PHOSPHOR_DIR, "regular")).map(basenameNoExt),
    );
    const customAvailable = new Set(
        listSvgs(CUSTOM_DIR).map(basenameNoExt).map(stripWeightSuffix),
    );
    const brandAvailable = new Set(listSvgs(BRAND_DIR).map(basenameNoExt));

    const phosphorWanted = new Set(
        [...phosphorAvailable].filter(
            (n) => tokens.has(n) || FORCE_INCLUDE.has(n),
        ),
    );
    const customWanted = new Set(
        [...customAvailable].filter(
            (n) => tokens.has(n) || FORCE_INCLUDE.has(n),
        ),
    );
    // Brand: include all (small list, names are short common words).
    const brandWanted = new Set(brandAvailable);

    console.log(
        `  phosphor: ${phosphorWanted.size} / ${phosphorAvailable.size} included`,
    );
    console.log(
        `  custom:   ${customWanted.size} / ${customAvailable.size} included`,
    );
    console.log(
        `  brand:    ${brandWanted.size} / ${brandAvailable.size} included`,
    );

    const symbols = [];

    for (const name of [...phosphorWanted].sort()) {
        for (const weight of PHOSPHOR_WEIGHTS) {
            const file = path.join(PHOSPHOR_DIR, weight, `${name}.svg`);
            if (!fs.existsSync(file)) continue;
            const { inner, viewBox } = extractSvg(
                fs.readFileSync(file, "utf8"),
            );
            symbols.push(makeSymbol(`p-${name}-${weight}`, viewBox, inner));
        }
    }

    for (const name of [...customWanted].sort()) {
        const baseFile = path.join(CUSTOM_DIR, `${name}.svg`);
        const baseExists = fs.existsSync(baseFile);
        const baseSvg = baseExists ? fs.readFileSync(baseFile, "utf8") : null;

        for (const weight of CUSTOM_WEIGHTS) {
            const suffix = weight === "regular" ? "" : `-${weight}`;
            const file = path.join(CUSTOM_DIR, `${name}${suffix}.svg`);
            let raw;
            if (fs.existsSync(file)) {
                raw = fs.readFileSync(file, "utf8");
            } else if (baseSvg) {
                raw = baseSvg;
            } else {
                continue;
            }
            const { inner, viewBox } = extractSvg(raw);
            symbols.push(makeSymbol(`c-${name}-${weight}`, viewBox, inner));
        }
    }

    for (const name of [...brandWanted].sort()) {
        const file = path.join(BRAND_DIR, `${name}.svg`);
        if (!fs.existsSync(file)) continue;
        const { inner, viewBox } = extractSvg(fs.readFileSync(file, "utf8"));
        symbols.push(makeSymbol(`b-${name}`, viewBox, inner));
    }

    const sprite =
        '<svg xmlns="http://www.w3.org/2000/svg" style="display:none">' +
        symbols.join("") +
        "</svg>\n";

    fs.writeFileSync(SPRITE_PATH, sprite);

    const manifest = {
        phosphor: [...phosphorWanted].sort(),
        custom: [...customWanted].sort(),
        brand: [...brandWanted].sort(),
    };
    fs.writeFileSync(MANIFEST_PATH, JSON.stringify(manifest, null, 2) + "\n");

    const kb = (sprite.length / 1024).toFixed(1);
    console.log(
        `  wrote ${path.relative(ROOT, SPRITE_PATH)} (${symbols.length} symbols, ${kb} KB)`,
    );
    console.log(`  wrote ${path.relative(ROOT, MANIFEST_PATH)}`);
    console.log("=== done ===");
}

build();
