const fs = require("fs");
const path = require("path");
const { SitemapAndIndexStream, SitemapStream } = require("sitemap");

const SITEMAP_PATH = path.resolve("public/sitemap.xml");
const OUTPUT_DIR = path.resolve("public/registry");
const CANONICAL_URL = "https://www.pulumi.com";

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

    // Track created sub-sitemap filenames for logging.
    const sitemapFiles = [];

    // SitemapAndIndexStream calls getSitemapStream each time it needs a new
    // sub-sitemap file (i.e. when the item limit is reached).
    const sms = new SitemapAndIndexStream({
        getSitemapStream(i) {
            const filename = `sitemap-${i}.xml`;
            const filePath = path.join(OUTPUT_DIR, filename);
            const ws = fs.createWriteStream(filePath);
            const smStream = new SitemapStream({
                hostname: CANONICAL_URL,
            });
            smStream.pipe(ws);
            sitemapFiles.push(filename);
            return [`${CANONICAL_URL}/registry/${filename}`, smStream, ws];
        },
    });

    // Pipe the sitemap index output to replace the original sitemap file.
    const indexWs = fs.createWriteStream(SITEMAP_PATH);
    sms.pipe(indexWs);

    // Feed all parsed items into the stream.
    for (const item of items) {
        sms.write(item);
    }

    // Wait for everything to finish.
    await new Promise((resolve, reject) => {
        indexWs.on("finish", resolve);
        indexWs.on("error", reject);
        sms.end();
    });

    console.log(
        `Split ${items.length} URLs into ${sitemapFiles.length} sitemap(s): ${sitemapFiles.join(", ")}`,
    );
}

splitSitemap().catch((err) => {
    console.error("Failed to split sitemap:", err);
    process.exit(1);
});
