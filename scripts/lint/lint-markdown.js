#!/usr/bin/env node

const fs = require("fs");
const yaml = require("js-yaml");
const markdownlint = require("markdownlint");
const path = require("path");

/**
 * REGEX for grabbing the the front matter of a Hugo markdown file. Example:
 *
 *     ---
 *     ...props
 *     ---
 */
const FRONT_MATTER_REGEX =
    /((^---\s*$[^]*?^---\s*$)|(^\+\+\+\s*$[^]*?^(\+\+\+|\.\.\.)\s*$))(\r\n|\r|\n|$)/m;
const AUTO_GENERATED_HEADING_REGEX =
    /###### Auto generated by ([a-z0-9]\w+)[/]([a-z0-9]\w+) on ([0-9]+)-([a-zA-z]\w+)-([0-9]\w+)/g;
const GENERATED_FILE_HEADER_REGEX = /# WARNING: this file was fetched from/;

/**
 * REGEX for matching top-level provider index files only. Example:
 *
 * registry/themes/default/content/registry/packages/auth0/_index.md
 *
 */
const PACKAGE_INDEX_REGEX = /packages\/[a-z0-9-]*\/_index.md/;

/**
 * Validates if a title exists, has length, and does not have a length over 60 characters.
 * More info: https://moz.com/learn/seo/title-tag
 *
 * @param {string} title The title tag value for a given page.
 */
function checkPageTitle(title) {
    if (!title) {
        return "Missing page title";
    } else if (typeof title === "string") {
        const titleLength = title.trim().length;
        if (titleLength === 0) {
            return "Page title is empty";
        }
    } else {
        return "Page title is not a valid string";
    }
    return null;
}

/**
 * Validates that a meta description exists, has length, is not too short,
 * and is not too long.
 * More info: https://moz.com/learn/seo/meta-description
 *
 * @param {string} meta The meta description for a given page
 */
function checkPageMetaDescription(meta) {
    if (!meta) {
        return "Missing meta description";
    } else if (typeof meta === "string") {
        const metaLength = meta.trim().length;
        if (metaLength === 0) {
            return "Meta description is empty";
        }
    } else {
        return "Meta description is not a valid string";
    }
    return null;
}

/**
 * Builds an array of markdown files to lint and checks each file's front matter
 * for formatting errors.
 *
 * @param {string[]} paths An array of paths to search for markdown files.
 * @param {Object} [markdownFiles] The markdownFiles object returned after finishing searching for non-index.md
 * markdown files.
 * @param {Object} [indexFiles] The indexFiles object returned after finishing searching for _index.md markdown files.
 * @returns [{Object}, {Object}] The markdown file objects for the two separate types of markdown files,
 * each including an error object for the files' front matter.
 */
function searchForMarkdown(paths, markdownFiles, indexFiles) {
    // If the markdownFiles result arg does not exist, initialize it
    if (!markdownFiles) {
        markdownFiles = {
            files: [],
            frontMatter: {},
        };
    }
    // If the indexFiles result arg does not exist, initialize it
    if (!indexFiles) {
        indexFiles = {
            files: [],
            frontMatter: {},
        };
    }

    // Grab the first file in the list and generate
    // its full path.
    const file = paths[0];
    const fullPath = path.resolve(__dirname, file);

    // Check if the path is a directory
    const isDirectory = fs.statSync(fullPath).isDirectory();

    // Get the file suffix so we can grab the markdown files.
    const fileParts = file.split(".");
    const fileSuffix = fileParts[fileParts.length - 1];

    // Ignore auto generated docs.
    if (file.indexOf("/content/docs/reference/pkg") > -1) {
        const remaining = paths.slice(1, paths.length);
        return searchForMarkdown(remaining, markdownFiles, indexFiles);
    }
    // If the path is a directory we want to add the contents of the directory
    // to the list.
    // azure-native-v2 is a static copy that's not updated anymore and whose
    // number of files causes stack overflow, so we skip it.
    if (isDirectory && file.indexOf("azure-native-v2") === -1) {
        const contents = fs.readdirSync(fullPath).map(function (file) {
            return fullPath + "/" + file;
        });
        paths[0] = contents;

        // Flatten the array.
        const newPaths = [].concat.apply([], paths);
        return searchForMarkdown(newPaths, markdownFiles, indexFiles);
        // Else check if the file suffix is a markdown
        // and add it the resulting file list.
    }
    if (fileSuffix === "md") {
        try {
            // Read the file contents so we can grab the file header.
            const content = fs.readFileSync(fullPath, "utf8");

            // Grab the file header.
            const frontMatter = content.match(FRONT_MATTER_REGEX);

            // Remove the dash blocks around the file header.
            const fContent = frontMatter[0].split("---").join("");

            // Read the yaml.
            const yamlObj = yaml.load(fContent);

            // If the page is auto generated, a redirect, or not indexed do not parse the front matter.
            const autoGenerated =
                yamlObj.no_edit_this_page === true ||
                content.match(AUTO_GENERATED_HEADING_REGEX) ||
                content.match(GENERATED_FILE_HEADER_REGEX);
            const redirectPassthrough = typeof yamlObj.redirect_to === "string";
            const noIndex = yamlObj.block_external_search_index === true;
            const allowLongTitle = !!yamlObj.allow_long_title;
            const shouldCheckFrontMatter =
                !autoGenerated &&
                !redirectPassthrough &&
                !noIndex &&
                !allowLongTitle;

            if (shouldCheckFrontMatter) {
                // Write the found file with information to the requisite return object.
                if (file.match(PACKAGE_INDEX_REGEX)) {
                    indexFiles = pushResults(indexFiles, yamlObj, fullPath);
                } else {
                    markdownFiles = pushResults(
                        markdownFiles,
                        yamlObj,
                        fullPath,
                    );
                }
            }
        } catch (e) {
            // Write the found file with error to the appropriate return object.
            if (file.match(PACKAGE_INDEX_REGEX)) {
                indexFiles = pushResults(indexFiles, null, fullPath, e.message);
            } else {
                markdownFiles = pushResults(
                    markdownFiles,
                    null,
                    fullPath,
                    e.message,
                );
            }
        }
    }

    // If there are remaining paths in the list, keep going.
    const remaining = paths.slice(1, paths.length);
    if (remaining.length > 0) {
        return searchForMarkdown(remaining, markdownFiles, indexFiles);
    }
    return [markdownFiles, indexFiles];
}

/**
 * Pushes a file information object to the collection of files given.
 * @param {Object} [fileCollection] The file information object we want to add a new file to.
 * markdown files.
 * @param {Object} [yamlObj] The indexFiles object returned after finishing searching for _index.md markdown files.
 * @param {string} fullPath The full path to each file.
 * @param {error} e Any error encountered when processing the file.
 * @returns [{Object}, {Object}] The markdown file objects for the two separate types of markdown files,
 * each including an error object for the files' front matter.
 */
function pushResults(fileCollection, yamlObj, fullPath, e) {
    // Build the front matter error object and add the file path.
    if (!yamlObj) {
        fileCollection.frontMatter[fullPath] = {
            error: e,
        };
    } else {
        fileCollection.frontMatter[fullPath] = {
            error: e,
            title: checkPageTitle(yamlObj.title),
            metaDescription: checkPageMetaDescription(yamlObj.meta_desc),
        };
    }
    fileCollection.files.push(fullPath);
    return fileCollection;
}

/**
 * Builds an array of markdown files to search through from a
 * given path.
 *
 * @param {string} parentPath The path to search for markdown files
 */
function getMarkdownFiles(parentPath) {
    const fullParentPath = path.resolve(__dirname, parentPath);
    const dirs = fs.readdirSync(fullParentPath).map(function (dir) {
        return path.join(parentPath, dir);
    });
    return searchForMarkdown(dirs);
}

/**
 * Groups the result of linting a file for markdown errors.
 *
 * @param {Object} result Results of lint errors. See: https://github.com/DavidAnson/markdownlint#usage
 * @return {Object} An object containing the file path and lint errors.
 * @return {string} result.path The full path of the linted file.
 * @return {Object[]} result.errors An array of error objects. Same as the result param.
 */
function groupLintErrorOutput(result, filesToLint) {
    // Grab the keys of the result object.
    const keys = Object.keys(result);

    // Map over the key array so we can combine front matter errors
    // with the markdown lint errors.
    const combinedErrors = keys.map(function (key) {
        // Get the lint and front matter errors.
        const lintErrors = result[key];
        const frontMatterErrors = filesToLint.frontMatter[key];

        // If the front matter error check threw an error add it to the lint
        // error array. Else add title and meta descriptions if they exist.
        if (frontMatterErrors.error) {
            lintErrors.push({
                lineNumber: "File Header",
                ruleDescription: frontMatterErrors.error,
            });
        } else {
            if (frontMatterErrors.title) {
                lintErrors.push({
                    lineNumber: "File Header",
                    ruleDescription: frontMatterErrors.title,
                });
            }
            if (frontMatterErrors.metaDescription) {
                lintErrors.push({
                    lineNumber: "File Header",
                    ruleDescription: frontMatterErrors.metaDescription,
                });
            }
        }

        if (lintErrors.length > 0) {
            return { path: key, errors: lintErrors };
        }
        return null;
    });

    // Filter out all null values from the combined result array.
    const filteredErrors = combinedErrors.filter(function (err) {
        return err !== null;
    });
    return filteredErrors;
}

// Build the lint object for the content directory.
const [markdownToLint, indexToLint] = getMarkdownFiles(
    `../../themes/default/content`,
);

/**
 * The config options for lint markdown files. All rules
 * are enabled by default. The config object let us customize
 * what rules we enforce and how we enforce them.
 *
 * See: https://github.com/DavidAnson/markdownlint
 */
const markdownFileOpts = {
    // The array of markdown files to lint.
    files: markdownToLint.files,
    config: {
        // Allow inline HTML.
        MD033: false,
        // Do not enforce line length.
        MD013: false,
        // Don't force language specification on code blocks.
        MD040: false,
        // Allow hard tabs.
        MD010: false,
        // Allow punctuation in headers.
        MD026: false,
        // Allow dollars signs in code blocks without values
        // immediately below the command.
        MD014: false,
        // Allow all code block styles in a file. Code block styles
        // are created equal and we shall not discriminate.
        MD046: false,
        // Allow duplicate headings.
        MD024: false,
        // Allow headings to be indented.
        MD023: false,
        // Allow blank lines in blockquotes.
        MD028: false,
        // Allow indentation in unordered lists.
        MD007: false,
        // Allow bareURLs.
        MD034: false,

        // Turning off the following rules so we can get this linter into a passing state so
        // we can turn it on. We can follow up and remove these in the future if we feel the
        // need to lint these rules.

        // Headings should be surrounded by blank lines
        MD022: false,
        // Lists should be surrounded by blank lines
        MD032: false,
        // Fenced code blocks should be surrounded by blank lines
        MD031: false,
        // Header line length
        MD036: false,
        // Unordered list style
        MD004: false,
        // Trailing spaces
        MD009: false,
        // Consecutive blank lines
        MD012: false,
        // Files should end with a single newline character
        MD047: false,
    },
    customRules: [],
};

/**
 * The config options for _index markdown files.
 * Similar to markdownFileOpts but with a few more rules disabled.
 */
const indexFileOpts = {
    // The array of markdown files to lint.
    files: indexToLint.files,
    config: Object.assign(
        markdownFileOpts.config,
        // Due to upstream docs not being linted rigorously, we ignore the following linting rules
        // in addition to the ones for other Markdown files.
        {
            // Heading levels should only increment by one level at a time
            MD001: false,
            // Multiple top-level headings in the same document
            MD025: false,
            // Ordered list item prefix
            MD029: false,
            // Spaces inside code span elements
            MD038: false,
        },
    ),
    customRules: [],
};

// Lint the markdown files.
let markdownLintResults = markdownlint.sync(markdownFileOpts);

// Lint `_index.md files`
let indexLintResults = markdownlint.sync(indexFileOpts);

// Group the lint errors by file.
let errors = groupLintErrorOutput(markdownLintResults, markdownToLint);
let lintErrors = groupLintErrorOutput(indexLintResults, indexToLint);
errors = errors.concat(lintErrors);

// Get the total number of errors.
const errorsArray = errors.map(function (err) {
    return err.errors;
});
const errorsCount = [].concat.apply([], errorsArray).length;

// Create the error output string.
const errorOutput = errors
    .map(function (err) {
        let msg = err.path + ":\n";
        for (let i = 0; i < err.errors.length; i++) {
            const error = err.errors[i];
            msg += "Line " + error.lineNumber + ": " + error.ruleDescription;
            msg += error.errorDetail
                ? " [" + error.errorDetail + "].\n"
                : ".\n";
        }
        return msg;
    })
    .join("\n");

// If there are errors output the error string and exit
// the program with an error.
if (errors.length > 0) {
    console.log(`
Markdown Lint Results:
    - ${markdownToLint.files.length + indexToLint.files.length} files parsed.
    - ${errorsCount} errors found.

Errors:

${errorOutput}
    `);

    const noError = process.argv.indexOf("--no-error") > -1;
    process.exit(noError ? 0 : 1);
}

console.log(`
Markdown Lint Results:
    - ${markdownToLint.files.length + indexToLint.files.length} files parsed.
    - ${errorsCount} errors found.
`);
process.exit(0);
