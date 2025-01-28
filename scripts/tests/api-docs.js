const chai = require("chai");
const glob = require("glob");
const fs = require("fs");
const path = require("path");
const cheerio = require("cheerio");
const htmlparser2 = require("htmlparser2");
const process = require("process");

const expect = chai.expect;

const pkg = process.argv[6].replace("--pkg=", "");
const split = process.argv[8].replace("--split=", "");

if (!pkg) {
    console.error("package not set", 1);
}

if (!fs.existsSync(`./public/registry/packages/${pkg}/api-docs/`)) {
    console.log("pkg skipped:", pkg);
    return;
}

const paths = getPaths(split);

let processed = 0;

paths.forEach((p) => {
    let fileContent = fs.readFileSync(p, "utf-8").toString();
    let dom = htmlparser2.parseDocument(fileContent);
    const $ = cheerio.load(dom);

    // Some of our non-bridged providers maintain a different directory structure
    // where the module structure is more deep or shallowly nested. This checks the
    // HTML file to see if it is a module page, and if it is, then skip it as the
    // tests that follow are to be run against resource pages.
    if (isModuleFile($) || isResourceListFile($)) {
        return;
    }

    if (!isFunctionPage($)) {
        describe(constructPageRoute(p), function () {
            // Verify page has a title and it is an h1 that contains the package and
            // module name.
            describe("h1 title", () => {
                const h1s = $("h1");

                it("contains exactly 1 h1", function () {
                    expect(h1s.length).to.equal(1);
                });
                it("contains the package name", function () {
                    expect(h1s.text().toLowerCase()).to.have.string(pkg);
                });
            });

            describe("Structure", () => {
                it("is nested correctly within docs-content", () => {
                    // Check if correctly nested withing the docs-main-content and docs-content.
                    // This is kind of hard to verify 100%, but package details section is
                    // a section that is ~always reliably there since it is not something
                    // conditionally there or relies on upstream parsing, so is an indicator
                    // something is wildly off it there is a failure here.
                    const structures = $(
                        `.docs-main-content .docs-content > h2#package-details`,
                    );
                    expect(structures.length).to.be.at.least(1);
                });
            });

            // Verify sections exist in correct order.
            describe("Sections", () => {
                const headings = [];
                $("section.docs-content > h2").each((i, h) => {
                    if (h) {
                        headings.push($(h).text());
                    }
                });
                it("exist in correct order", () => {
                    const possibleHeadings = [
                        "Example Usage",
                        "Create",
                        "Resource Properties",
                        "Look up Existing",
                        "Supporting Types",
                        "Import",
                        "Package Details",
                    ];
                    const actual = possibleHeadings.filter((ph) => {
                        return headings.some((h) => {
                            return ph.includes(h);
                        });
                    });
                    const expected = possibleHeadings.filter((h) =>
                        actual.includes(h),
                    );
                    expect(JSON.stringify(actual)).to.equal(
                        JSON.stringify(expected),
                    );
                });
            });

            // If contains a supporting types section check properties contain descriptions.
            if ($("h2#supporting-types").length > 0) {
                describe("Supporting Types", () => {
                    it("all props should contain descriptions", () => {
                        expect(
                            checkSupportingTypePropertyDescriptions($).length,
                        ).to.equal(0);
                    });
                });
            }

            // Verify it contains a description paragraph.
            describe("Description paragraph", () => {
                const paragraph = $(
                    ".docs-main-content .docs-content > p",
                ).text();
                it("exists", () => {
                    expect(paragraph.length).to.be.at.least(1);
                });
            });

            // Verify constructor section exists.
            describe("Constructor", () => {
                it("has a the constructor section heading", () => {
                    const constructorSection = $("h2#create");
                    expect(constructorSection.length).to.equal(1);
                });
                it("has a description following the heading", () => {
                    const constructorSection = $("h2#create + p");
                    expect(constructorSection.length).to.equal(1);
                });
                it("has a constructor syntax section", () => {
                    const syntax = $("h3#constructor-syntax");
                    expect(syntax.length).to.equal(1);
                });
                it("has a code chooser block", () => {
                    const codeChooser = $(
                        "h3#constructor-syntax + div > pulumi-chooser",
                    );
                    expect(codeChooser.length).to.be.at.least(1);
                });
            });

            // Verify Input/Output Properties section.
            describe("Properties", () => {
                it("properties section exists", () => {
                    const properties = $("h2#properties");
                    expect(properties.length).to.equal(1);
                });
                it("inputs section exists", () => {
                    const inputs = $("h3#inputs");
                    expect(inputs.length).to.equal(1);
                });
                it("all input properties contain descriptions", () => {
                    expect(checkInputPropertyDescriptions($).length).to.equal(
                        0,
                    );
                });
                it("outputs section exists", () => {
                    const outputs = $("h3#outputs");
                    expect(outputs.length).to.equal(1);
                });
                it("all output properties contain descriptions", () => {
                    expect(checkOutputPropertyDescriptions($).length).to.equal(
                        0,
                    );
                });
            });

            // Verify the page contains an Example Usage section and contains
            // at least 1 code example.
            describe("Examples section", () => {
                it("contains Example Usage heading", () => {
                    const heading = $("h2#example-usage");
                    expect(heading.length).to.equal(1);
                    expect(heading.text()).to.have.string("Example Usage");
                });
                const examplesChooser = $(
                    "h2#example-usage ~ div > pulumi-chooser",
                );
                it("contains at least one example with chooser", () => {
                    expect(examplesChooser.length).to.be.at.least(1);
                });
                it("contains at least 1 example with code block", () => {
                    const code = $("h2#example-usage ~ div > pulumi-choosable");
                    expect(code.length).to.be.at.least(1);
                });
                it("contains at least 1 example with code block", () => {
                    const code = $("h2#example-usage ~ div > pulumi-choosable");
                    expect(code.length).to.be.at.least(1);
                });
            });

            // Verify the page contains an Import section and that it comes somewhere
            // after the Properties section.
            describe("Import section", () => {
                const heading = $("h2#import");
                if (heading.length > 0) {
                    it("comes after Properties list", () => {
                        const importSection = $("h2#properties ~ h2#import");
                        expect(importSection.length).to.equal(1);
                    });
                } else {
                    it("exists", () => {
                        expect(heading.length).to.equal(1);
                        expect(heading.text()).to.have.string("Import");
                    });
                }
            });

            // Verify the page contains a Package Details section
            describe("Package Details section", () => {
                it("exists", () => {
                    const packageSection = $("h2#package-details");
                    expect(packageSection.length).to.equal(1);
                });
                it("contains details table", () => {
                    const table = $("h2#package-details + dl");
                    expect(table.length).to.equal(1);
                });
            });
        });
    } else {
        describe(constructPageRoute(p) + " (Fn)", function () {
            describe("h1 title", () => {
                const h1s = $("h1");

                it("contains exactly 1 h1", function () {
                    expect(h1s.length).to.equal(1);
                });
                it("contains the package name", function () {
                    expect(h1s.text().toLowerCase()).to.have.string(pkg);
                });
            });
            // Verify it contains a description paragraph.
            describe("Description paragraph", () => {
                const paragraph = $(
                    ".docs-main-content .docs-content > p",
                ).text();
                it("exists", () => {
                    expect(paragraph.length).to.be.at.least(1);
                });
            });
            // Contains a section describing the function syntax
            describe("Syntax usage", () => {
                it("has a usage section heading", () => {
                    const constructorSection = $("h2#using");
                    expect(constructorSection.length).to.equal(1);
                });
                it("has a description paragraph following the heading", () => {
                    const constructorSection = $("h2#using + p");
                    expect(constructorSection.length).to.equal(1);
                });
                it("contains code syntax", () => {
                    const syntax = $("h2#using + p + div > pulumi-chooser");
                    expect(syntax.length).to.equal(1);
                });
            });

            // Contains a section describing the output of the function.
            describe("Result section", () => {
                it("has a result heading", () => {
                    const resultHeading = $("h2#result");
                    expect(resultHeading.length).to.equal(1);
                });
                it("has a properties table", () => {
                    const propsTable = $(
                        "h2#result ~ div > pulumi-choosable > dl",
                    );
                    expect(propsTable.length).to.be.at.least(1);
                });
            });

            // Verify the page contains an Example Usage section and contains
            // at least 1 code example.
            describe("Examples section", () => {
                it("contains Example Usage heading", () => {
                    const heading = $("h2#example-usage");
                    expect(heading.length).to.equal(1);
                    expect(heading.text()).to.have.string("Example Usage");
                });
                const examplesChooser = $(
                    "h2#example-usage ~ div > pulumi-chooser",
                );
                it("contains at least one example with chooser", () => {
                    expect(examplesChooser.length).to.be.at.least(1);
                });
                it("contains at least 1 example with code block", () => {
                    const code = $("h2#example-usage ~ div > pulumi-choosable");
                    expect(code.length).to.be.at.least(1);
                });
            });

            // Contains a Package Details section
            describe("Package Details section", () => {
                it("exists", () => {
                    const packageSection = $("h2#package-details");
                    expect(packageSection.length).to.equal(1);
                });
                it("contains details table", () => {
                    const table = $("h2#package-details + dl");
                    expect(table.length).to.equal(1);
                });
            });
        });
    }

    // By default node runs garbage collection very lazily since it needs to halt
    // execution in order to run it, so this manually invokes the garbage collector
    // every 200 pages and tends to slightly help the memory issues.
    processed++;
    if (processed > 200) {
        global.gc();
        processed = 0;
    }
});

function constructPageRoute(filePath) {
    return filePath
        .replace("./public", "https://pulumi.com")
        .replace("index.html", "");
}

// Check if this page is a module list page.
function isModuleFile($) {
    const moduleHeading = $("h2#modules");
    return moduleHeading.length > 0;
}

// Check if the page is a the resource list page.
function isResourceListFile($) {
    const resourcesHeading = $("h2#resources");
    const functionsHeading = $("h2#functions");
    return resourcesHeading.length > 0 || functionsHeading.length > 0;
}

// Generate a list of the function names to add to a list so we know which pages
// are functions.
function getFunctions($) {
    const containsFunctions = $("h2#functions").length > 0;
    if (containsFunctions) {
        const list = $("h2#functions + ul > li");
        list.each((i, elm) => {
            functions.push($(elm).text().toLowerCase());
        });
    }
}

function checkInputPropertyDescriptions($) {
    const props = $("h3#inputs + p + div > pulumi-choosable > dl > dd");
    return props
        .filter((i, elm) => {
            if ($(elm).html().length < 1) {
                return true;
            }
            return false;
        })
        .toArray();
}

function checkOutputPropertyDescriptions($) {
    const props = $("h3#outputs + p + div > pulumi-choosable > dl > dd");
    return props
        .filter((i, elm) => {
            if ($(elm).html().length < 1) {
                return true;
            }
            return false;
        })
        .toArray();
}

function checkSupportingTypePropertyDescriptions($) {
    const props = $(
        "h2#supporting-types + h4 + div > pulumi-choosable > dl > dd",
    );
    return props
        .filter((i, elm) => {
            if ($(elm).html().length < 1) {
                return true;
            }
            return false;
        })
        .toArray();
}

// Checks if the page is a function page. There is not a distinction on the page itself
// other than it having a heading for the return result of the function.
function isFunctionPage($) {
    const functionPageHeading = $("h2#result");
    return functionPageHeading.length > 0;
}

function getPaths(split) {
    const paths = glob.sync(
        `./public/registry/packages/${pkg}/api-docs/**/*.html`,
        {
            ignore: [`./public/registry/packages/${pkg}/api-docs/index.html`],
        },
    );
    switch (split) {
        case "1":
            console.log("processing first half...");
            return paths.sort().slice(0, paths.length / 2);
            break;
        case "2":
            console.log("processing second half...");
            return paths.sort().slice(paths.length / 2);
            break;
        default:
            return paths;
    }
}
