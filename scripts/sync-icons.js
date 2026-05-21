#!/usr/bin/env node
/**
 * sync-icons.js
 *
 * Refreshes the vendored icon SVGs under themes/default/assets/icons/ from
 * upstream npm packages:
 *
 *   - @phosphor-icons/core   → phosphor/{regular,bold,fill,duotone}/<name>.svg
 *   - simple-icons           → brand/<name>.svg  (limited to BRAND_ICONS below)
 *
 * Custom Pulumi icons under assets/icons/custom/ are not touched.
 *
 * After running this, run `make build-icon-sprite` (or `make build-assets`)
 * to regenerate the sprite + manifest.
 *
 * Adapted from pulumi/docs scripts/sync-icons.js.
 */

"use strict";

const fs = require("fs");
const path = require("path");

const ROOT = path.resolve(__dirname, "..");
const THEME_ROOT = path.join(ROOT, "themes/default");
const ICONS_DIR = path.join(THEME_ROOT, "assets/icons");
const PHOSPHOR_DIR = path.join(ICONS_DIR, "phosphor");
const BRAND_DIR = path.join(ICONS_DIR, "brand");

const PHOSPHOR_WEIGHTS = ["regular", "bold", "fill", "duotone"];

// Brand icons sourced from simple-icons. Add a slug here (matching the
// simple-icons filename without extension) to vendor a new brand mark.
const BRAND_ICONS = ["bluesky", "github", "linkedin", "slack", "x", "youtube"];

// Resolve a package directory without going through require.resolve(), which
// fails for packages with restricted `exports` fields (e.g. recent
// @phosphor-icons/core releases).
function resolvePackageDir(name) {
    const dir = path.join(ROOT, "node_modules", name);
    if (!fs.existsSync(dir)) {
        throw new Error(`package not installed: ${name}`);
    }
    return dir;
}

function resolvePhosphorAssets() {
    const pkgDir = resolvePackageDir("@phosphor-icons/core");
    const candidates = [path.join(pkgDir, "assets"), path.join(pkgDir, "raw")];
    for (const dir of candidates) {
        if (fs.existsSync(dir)) return dir;
    }
    throw new Error(
        `@phosphor-icons/core: no SVG asset directory under ${pkgDir}`,
    );
}

function resolveSimpleIcons() {
    const pkgDir = resolvePackageDir("simple-icons");
    const dir = path.join(pkgDir, "icons");
    if (!fs.existsSync(dir)) {
        throw new Error(
            `simple-icons: icons directory not found under ${pkgDir}`,
        );
    }
    return dir;
}

function ensureCleanDir(dir) {
    fs.rmSync(dir, { recursive: true, force: true });
    fs.mkdirSync(dir, { recursive: true });
}

function stripPhosphorSuffix(file, weight) {
    if (weight === "regular") return file;
    return file.replace(new RegExp(`-${weight}\\.svg$`), ".svg");
}

// Duotone SVGs from upstream are plain two-path documents. The sprite styles
// the two layers via `.duotone-primary` / `.duotone-secondary`, so we annotate
// each path on sync: paths with opacity="0.2" become secondary, the rest
// primary. Idempotent — skips paths that already carry a class.
function annotateDuotone(svg) {
    return svg.replace(/<path\b([^>]*?)\/>/g, (match, attrs) => {
        if (/\bclass=/.test(attrs)) return match;
        const role = /opacity="0\.2"/.test(attrs)
            ? "duotone-secondary"
            : "duotone-primary";
        return `<path${attrs} class="${role}"/>`;
    });
}

function syncPhosphor() {
    const src = resolvePhosphorAssets();
    console.log(`  source: ${path.relative(ROOT, src)}`);

    let total = 0;
    for (const weight of PHOSPHOR_WEIGHTS) {
        const srcDir = path.join(src, weight);
        if (!fs.existsSync(srcDir)) {
            console.warn(
                `  warn: missing weight directory ${srcDir} — skipping`,
            );
            continue;
        }
        const destDir = path.join(PHOSPHOR_DIR, weight);
        ensureCleanDir(destDir);

        let count = 0;
        for (const file of fs.readdirSync(srcDir)) {
            if (!file.endsWith(".svg")) continue;
            const destName = stripPhosphorSuffix(file, weight);
            const destPath = path.join(destDir, destName);
            if (weight === "duotone") {
                const svg = fs.readFileSync(path.join(srcDir, file), "utf8");
                fs.writeFileSync(destPath, annotateDuotone(svg));
            } else {
                fs.copyFileSync(path.join(srcDir, file), destPath);
            }
            count++;
        }
        console.log(`  phosphor/${weight}: ${count} icons`);
        total += count;
    }
    return total;
}

function syncBrand() {
    const src = resolveSimpleIcons();
    console.log(`  source: ${path.relative(ROOT, src)}`);

    // Preserve any non-simple-icons brand SVGs (e.g. readme.svg). Only the
    // names in BRAND_ICONS are refreshed; everything else is left alone.
    fs.mkdirSync(BRAND_DIR, { recursive: true });

    let count = 0;
    const missing = [];
    for (const name of BRAND_ICONS) {
        const srcFile = path.join(src, `${name}.svg`);
        if (!fs.existsSync(srcFile)) {
            missing.push(name);
            continue;
        }
        // Strip the inline <title> simple-icons embeds; we rely on the
        // consuming partial's aria-label for the accessible name and
        // would otherwise duplicate it for screen readers.
        const svg = fs
            .readFileSync(srcFile, "utf8")
            .replace(/<title>[\s\S]*?<\/title>/, "");
        fs.writeFileSync(path.join(BRAND_DIR, `${name}.svg`), svg);
        count++;
    }
    console.log(
        `  brand: ${count} icons (${BRAND_ICONS.length - count} missing)`,
    );
    if (missing.length) {
        console.warn(
            `  warn: not found in simple-icons: ${missing.join(", ")}`,
        );
    }
    return count;
}

function main() {
    console.log("=== sync-icons ===");
    const phosphor = syncPhosphor();
    const brand = syncBrand();
    console.log(`  total: ${phosphor + brand} icons synced`);
    console.log(
        "  next: run `make build-icon-sprite` to regenerate the sprite",
    );
    console.log("=== done ===");
}

main();
