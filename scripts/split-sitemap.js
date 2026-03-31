const fs = require("fs");
const path = require("path");
const {
    parseSitemap,
    SitemapAndIndexStream,
    SitemapStream,
} = require("sitemap");

const SITEMAP_PATH = path.resolve("public/sitemap.xml");
const OUTPUT_DIR = path.resolve("public/registry");
const CANONICAL_URL = "https://www.pulumi.com";

async function splitSitemap() {
    if (!fs.existsSync(SITEMAP_PATH)) {
        console.log("No sitemap found at", SITEMAP_PATH);
        return;
    }

    // Parse the Hugo-generated sitemap into structured items.
    const xml = fs.createReadStream(SITEMAP_PATH);
    const items = await parseSitemap(xml);
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
