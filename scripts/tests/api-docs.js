const chai = require("chai");
const glob = require("glob");
const fs = require("fs");
const path = require("path");
const cheerio = require("cheerio");
const htmlparser2 = require("htmlparser2");
const process = require("process");

const expect = chai.expect;

const pkg = process.argv[5].replace("--pkg=", "");

if (!pkg) {
    console.error("package not set", 1);
}

const modules = getModules(pkg);
    const paths = glob.sync(
        `./public/registry/packages/${pkg}/api-docs/**/*.html`,
        {
            ignore: [
                `./public/registry/packages/${pkg}/api-docs/index.html`,
            ],
        },
    );

    // Iterate over pages and build a list of all the functions. This way we know which pages to skip
    // since these tests are only built for resource pages.
    const functions = [];
    paths.forEach((p) => {
        const fileContent = fs.readFileSync(p, "utf-8").toString();
        const dom = htmlparser2.parseDocument(fileContent);
        const $ = cheerio.load(dom);
        if (isModuleFile($) || isResourceListFile($)) {
            getFunctions($);
            return;
        }
    });

    paths.forEach((p) => {

        // Some of our non-bridged providers maintain a different directory structure 
        // where the module structure is more deep or shallowly nested. This checks the
        // html file to see if it is a module page, and if it is, then skip it as the
        // tests that follow are to be run against resource pages.

        const fileContent = fs.readFileSync(p, "utf-8").toString();
        const dom = htmlparser2.parseDocument(fileContent);
        const $ = cheerio.load(dom);

        if (isModuleFile($) || isResourceListFile($)) {
            return;
        }
        
        
        // Ignore function pages for now.
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
                    // it("contains the module name", function () {
                    //     expect(h1s.text().toLowerCase()).to.have.string(mod);
                    // });
                });

                // Verify sections exist in correct order
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
                    it("outputs section exists", () => {
                        const outputs = $("h3#outputs");
                        expect(outputs.length).to.equal(1);
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

                    const examples = $(
                        "h2#example-usage ~ div > pulumi-chooser",
                    );
                    it("contains at least one example", () => {
                        expect(examples.length).to.be.at.least(1);
                    });
                    it("contains at least 1 example with code block", () => {
                        const code = $("h2#example-usage ~ div > pulumi-choosable");
                        expect(code.length).to.be.at.least(1);
                    })
                });

                // Verify the page contains an Import section and that it comes somewhere
                // after the Properties section.
                describe("Import section", () => {
                    const heading = $("h2#import");
                    if (heading.length > 0) {
                        it("comes after Properties list", () => {
                            const importSection = $(
                                "h2#properties ~ h2#import",
                            );
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
        }
    });

function getModules(pkg) {
    const mods = fs
        .readdirSync(`./public/registry/packages/${pkg}/api-docs`, {
            withFileTypes: true,
        })
        .filter((dirent) => dirent.isDirectory())
        .map((dirent) => dirent.name);

    return mods;
}

function constructPageRoute(filePath) {
    return filePath
        .replace("./public", "https://pulumi.com")
        .replace("index.html", "");
}

// Check if this page is a module list page.
function isModuleFile($) {
  const moduleHeading = $("h2#modules")
  return moduleHeading.length > 0;
}

// Check if the page is a the resource list page.
function isResourceListFile($) {
  const resourcesHeading = $("h2#resources");
  const functionsHeading = $("h2#functions");
  return resourcesHeading.length > 0 || functionsHeading.length > 0;
}

// Generate a list of the function names to add to a list so we know which pages
// to skip over.
function getFunctions($) {
    const containsFunctions = $("h2#functions").length > 0;
    if (containsFunctions) {
        const list = $("h2#functions + ul > li");
        list.each((i, elm) => {
            functions.push($(elm).text().toLowerCase())
        });
    }
}

// Checks if the page is a function page. There is not a distinction on the page itself
// other than it having a heading for the return result of the function.
function isFunctionPage($) {
    const functionPageHeading = $("h2#result");
    return functionPageHeading.length > 0;
}