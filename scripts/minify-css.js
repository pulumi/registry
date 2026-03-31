const fs = require("fs");
const glob = require("glob");
const postcss = require("postcss");
const { purgeCSSPlugin: purgecss } = require("@fullhuman/postcss-purgecss");
const cssnano = require("cssnano");

const registryBundleSlug = "-registry";

const cssBundleId = process.env.ASSET_BUNDLE_ID;

if (!cssBundleId) {
    console.error("ERROR: ASSET_BUNDLE_ID environment variable not set");
    process.exit(1);
}

// Per-bundle PurgeCSS configuration.
const bundles = [
    {
        name: `bundle${registryBundleSlug}`,
        input: `public/registry/css/bundle${registryBundleSlug}.*.css`,
        content: [
            // Scan top-level pages and layouts, but skip the bulk of
            // API doc pages (they all share the same templates/classes).
            "public/*.html",
            "public/registry/*.html",
            "public/registry/packages/*.html",
            "public/registry/packages/*/index.html",
            "public/registry/packages/*/installation-configuration/index.html",
            "public/registry/packages/*/api-docs/index.html",
            // Scan a representative provider's full API docs for class coverage.
            "public/registry/packages/aws/api-docs/**/*.html",
            `public/registry/js/bundle${registryBundleSlug}.*.js`,
        ],
    },
    {
        name: `marketing${registryBundleSlug}`,
        input: `public/registry/css/marketing${registryBundleSlug}.*.css`,
        content: [
            "public/*.html",
            "public/registry/*.html",
            `public/registry/js/bundle${registryBundleSlug}.*.js`,
        ],
    },
];

const sharedSafelist = [
    /^hs-/,
    /^highlight$/,
    /^pagination$/,
    /^code-/,
    /^copy-/,
    /^carousel/,
    /^st-/,
    /^icon-/,
    /^package-details/,
    /^resources-properties/,
    /^tabular/,
    /^pulumi-/,
];

function minifyCSS(config) {
    const bundlePath = glob.sync(config.input)[0];

    if (bundlePath === undefined) {
        return Promise.resolve();
    }

    const css = fs.readFileSync(bundlePath);
    const outputPath = `public/registry/css/${config.name}.${cssBundleId}.css`;

    return postcss([
        // PurgeCSS removes unused CSS by analyzing the files of the built website.
        // https://purgecss.com/
        purgecss({
            content: config.content,
            css: [bundlePath],
            safelist: { deep: sharedSafelist },

            // We need to extract the Tailwind screen size selectors (e.g. sm, md, lg)
            // so that we do not strip them out. As long as a class name appears in the HTML
            // in its entirety, PurgeCSS will not remove it.
            // Ex. https://tailwindcss.com/docs/optimizing-for-production#writing-purgeable-html
            defaultExtractor: (content) =>
                content.match(/[\w-/.:]+(?<!:)/g) || [],
        }),

        // CSSNano minifies our rendered CSS by removing whitespace, comments, etc.
        // https://cssnano.co/
        cssnano({
            preset: [
                "default",
                {
                    discardComments: {
                        removeAll: true,
                    },
                },
            ],
        }),
    ])
        .process(css, { from: bundlePath, to: outputPath })
        .then((result) => {
            const css = result.css;

            // Make sure there's at least some valid-looking CSS in the result.
            if (!css || !css.match(/html|body/)) {
                throw new Error(`Unexpected PostCSS result: ${css}`);
            }

            fs.writeFileSync(outputPath, css);
            console.log(`Minified: ${outputPath}`);
        });
}

Promise.all(bundles.map((b) => minifyCSS(b))).then(() => {
    console.log("CSS bundles minified successfully!");
    bundles.forEach((b) => console.log(`  - ${b.name}.${cssBundleId}.css`));
});

// Exit non-zero when something goes wrong in the promise chain.
process.on("unhandledRejection", (error) => {
    throw new Error(error);
});
