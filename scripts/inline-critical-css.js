const fs = require("fs");
const Beasties = require("beasties");

// Pages to inline critical CSS for, based on CDN traffic data (last 30 days).
// API docs index pages dominate registry traffic.
const pages = [
    "public/registry/index.html",
    "public/registry/packages/cloudflare/api-docs/index.html",
    "public/registry/packages/akamai/api-docs/index.html",
    "public/registry/packages/aws/api-docs/index.html",
    "public/registry/packages/digitalocean/api-docs/index.html",
    "public/registry/packages/gcp/api-docs/index.html",
    "public/registry/packages/azure/api-docs/index.html",
    "public/registry/packages/datadog/api-docs/index.html",
    "public/registry/packages/gitlab/api-docs/index.html",
];

async function inlineCriticalCSS() {
    const beasties = new Beasties({
        path: "public/",
        preload: "swap",
        reduceInlineStyles: true,
        // Keep original CSS files intact; the full stylesheet is still loaded
        // async so uncovered rules still apply after page load.
        pruneSource: false,
    });

    for (const page of pages) {
        if (!fs.existsSync(page)) {
            console.warn(`Skipping ${page} (not found)`);
            continue;
        }

        const html = fs.readFileSync(page, "utf-8");
        const inlined = await beasties.process(html);
        fs.writeFileSync(page, inlined);
        console.log(`Inlined critical CSS: ${page}`);
    }
}

inlineCriticalCSS().catch((err) => {
    console.error("Failed to inline critical CSS:", err);
    process.exit(1);
});
