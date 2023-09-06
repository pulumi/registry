/**
 * Configuration file for PostCSS. This is only used as part of the Hugo pipeline
 * build process for our Sass files, since we run the resulting CSS through PostCSS
 * at the end to do more transformations.
 */

module.exports = {
    plugins: [
        // TailwindCSS
        require("tailwindcss")("./tailwind.config.js"),
    ],
};
