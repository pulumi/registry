const fs = require("fs");
const Beasties = require("beasties");

// Pages to inline critical CSS for. These are the highest-traffic
// entry points in the registry.
const pages = [
    "public/registry/index.html",
    "public/registry/packages/aws/index.html",
    "public/registry/packages/gcp/index.html",
    "public/registry/packages/azure/index.html",
    "public/registry/packages/azure-native/index.html",
    "public/registry/packages/aws-native/index.html",
    "public/registry/packages/kubernetes/index.html",
    "public/registry/packages/docker/index.html",
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
