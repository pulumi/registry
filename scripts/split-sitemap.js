const fs = require("fs");
const path = require("path");
const { SitemapIndexStream, SitemapStream } = require("sitemap");

const SITEMAP_PATH = path.resolve("public/sitemap.xml");
const OUTPUT_DIR = path.resolve("public/registry");
const CANONICAL_URL = "https://www.pulumi.com";
// Matches the "sitemap" package's own DEFAULT_SITEMAP_ITEM_LIMIT. Chunking is
// done by hand here (rather than via SitemapAndIndexStream) so that each
// shard's <lastmod> in the index can be computed from the items actually
// written into it -- SitemapAndIndexStream asks for a shard's IndexItem
// before any of that shard's items are known, so it cannot supply this.
const ITEMS_PER_SITEMAP = 45000;

function waitForFinish(stream) {
    return new Promise((resolve, reject) => {
        stream.on("finish", resolve);
        stream.on("error", reject);
    });
}

// Newest ISO-8601 lastmod across a shard's items, or undefined if none of
// the items in the shard carry a lastmod. A shard's own modification time is
// exactly when its newest contained URL last changed, so this is an honest
// (not fabricated) value -- consistent with only emitting <lastmod> where it
// can be computed truthfully.
function maxLastmod(items) {
    let max;
    for (const item of items) {
        if (!item.lastmod) {
            continue;
        }
        const d = new Date(item.lastmod);
        if (Number.isNaN(d.getTime())) {
            continue;
        }
        if (!max || d > max) {
            max = d;
        }
    }
    return max ? max.toISOString() : undefined;
}

async function writeShard(filename, items) {
    const filePath = path.join(OUTPUT_DIR, filename);
    const ws = fs.createWriteStream(filePath);
    const smStream = new SitemapStream({ hostname: CANONICAL_URL });
    smStream.pipe(ws);
    for (const item of items) {
        smStream.write(item);
    }
    smStream.end();
    await Promise.all([waitForFinish(smStream), waitForFinish(ws)]);
}

async function splitSitemap() {
    if (!fs.existsSync(SITEMAP_PATH)) {
        console.log("No sitemap found at", SITEMAP_PATH);
        return;
    }

    // Parse URL entries directly from the Hugo-generated XML. We avoid the
    // sitemap library's parseSitemap because it enforces a 50K-entry limit
    // on the *input* — which is the very problem we're trying to solve.
    const xml = fs.readFileSync(SITEMAP_PATH, "utf-8");
    const urlPattern =
        /<url>\s*<loc>([^<]+)<\/loc>(?:\s*<lastmod>([^<]+)<\/lastmod>)?/g;
    const items = [];
    let droppedNonCanonical = 0;
    let match;
    while ((match = urlPattern.exec(xml)) !== null) {
        const url = match[1];
        // Defense in depth for #10639: drop versioned package URLs
        // (e.g. /registry/packages/aws@6.x/...) whose canonical tag
        // points to the unversioned sibling. The Hugo sitemap template
        // should already exclude these, but filtering here catches any
        // template regression before it reaches search engines.
        if (url.includes("@") || url.includes("%40")) {
            droppedNonCanonical++;
            continue;
        }
        const item = { url };
        if (match[2]) {
            item.lastmod = match[2];
        }
        items.push(item);
    }

    if (droppedNonCanonical > 0) {
        console.log(
            `Dropped ${droppedNonCanonical} non-canonical (@-versioned) URLs.`,
        );
    }
    console.log(`Sitemap contains ${items.length} URLs.`);

    if (items.length === 0) {
        console.log("Nothing to split.");
        return;
    }

    fs.mkdirSync(OUTPUT_DIR, { recursive: true });

    // Chunk items into shards up front so each shard's max lastmod is known
    // before we write its index entry.
    const chunks = [];
    for (let i = 0; i < items.length; i += ITEMS_PER_SITEMAP) {
        chunks.push(items.slice(i, i + ITEMS_PER_SITEMAP));
    }

    const sitemapFiles = [];
    const indexItems = [];
    for (let i = 0; i < chunks.length; i++) {
        const filename = `sitemap-${i}.xml`;
        await writeShard(filename, chunks[i]);
        sitemapFiles.push(filename);
        const indexItem = { url: `${CANONICAL_URL}/registry/${filename}` };
        const lastmod = maxLastmod(chunks[i]);
        if (lastmod) {
            indexItem.lastmod = lastmod;
        }
        indexItems.push(indexItem);
    }

    // Write the sitemap index, replacing the original flat sitemap file.
    const indexStream = new SitemapIndexStream();
    const indexWs = fs.createWriteStream(SITEMAP_PATH);
    indexStream.pipe(indexWs);
    for (const indexItem of indexItems) {
        indexStream.write(indexItem);
    }
    indexStream.end();
    await Promise.all([waitForFinish(indexStream), waitForFinish(indexWs)]);

    console.log(
        `Split ${items.length} URLs into ${sitemapFiles.length} sitemap(s): ${sitemapFiles.join(", ")}`,
    );
}

splitSitemap().catch((err) => {
    console.error("Failed to split sitemap:", err);
    process.exit(1);
});
