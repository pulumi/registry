const fs = require("fs");
const glob = require("glob");
const postcss = require("postcss");
const purgecss = require("@fullhuman/postcss-purgecss");
const cssnano = require("cssnano");

const registryBundleSlug = "-registry";

function minifyCSS(filePath) {
    const bundlePath = glob.sync(filePath)[0];

    // If there is no matching bundle then we will skip minifying things.
    if (bundlePath === undefined) {
        return Promise.resolve();
    }

    const css = fs.readFileSync(bundlePath);

    return postcss([
        // PurgeCSS removes unused CSS by analyzing the files of the built website.
        // https://purgecss.com/
        purgecss({
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
            css: [bundlePath],
            safelist: {
                deep: [
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
                ],
            },

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
        .process(css, { from: bundlePath, to: bundlePath })
        .then((result) => {
            const css = result.css;

            // Make sure there's at least some valid-looking CSS in the result.
            if (!css || !css.match(/html|body/)) {
                throw new Error(`Unexpected PostCSS result: ${css}`);
            }

            // Write the result back to the file.
            fs.writeFileSync(bundlePath, css);
        });
}

minifyCSS(`public/registry/css/bundle${registryBundleSlug}.*.css`).then(() => {
    minifyCSS(`public/registry/css/marketing${registryBundleSlug}.*.css`).then(
        () => {
            console.log("CSS bundles minified successfully!");
        },
    );
});

// Exit non-zero when something goes wrong in the promise chain.
process.on("unhandledRejection", (error) => {
    throw new Error(error);
});
