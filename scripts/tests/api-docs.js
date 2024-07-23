const chai = require('chai')
const glob = require("glob");
const fs = require("fs");
const path = require("path");
const cheerio = require('cheerio');
const htmlparser2 = require('htmlparser2');
const process = require('process')

const expect = chai.expect

const pkg = process.argv[5].replace("--pkg=", "");

console.log("pkg", pkg);

if (!pkg) {
  console.error("package not set", 1);
}

const modules = getModules(pkg);
modules.forEach(mod => {
  const paths = glob.sync(`./public/registry/packages/${pkg}/api-docs/${mod}/**/*.html`, {
    ignore: [
      `./public/registry/packages/${pkg}/api-docs/${mod}/index.html`
    ],
  });
  paths.forEach(p => {
    const fileContent = fs.readFileSync(p, "utf-8").toString();
    const dom = htmlparser2.parseDocument(fileContent);
    const $ = cheerio.load(dom);

    // Ignore funtion pages for now.
    if (!path.basename(path.dirname(p)).startsWith("get")) {

      describe(constructPageRoute(p), function () {
        // Verify page has a title and it is an h1 that contains the package and
        // module name.
        describe("h1 title", function () {
          const h1s = $("h1");

          it("contains exactly 1 h1", function () {
            expect(h1s.length).to.equal(1);
          })
          it("contains the package name", function () {
            expect(h1s.text().toLowerCase()).to.have.string(pkg);
          })
          it("contains the module name", function () {
            expect(h1s.text().toLowerCase()).to.have.string(mod);
          })
        })

        // Verify sections exist in correct order
        describe("Sections", function () {
          const headings = [];
          $("section.docs-content > h2").each((i, elm) => {
            if (elm) {
              headings.push($(elm).text());
            }
          })
          it('exist in correct order', function () {
            const possibleHeadings = [
              "Example Usage",
              "Create",
              "Resource Properties",
              "Look up Existing",
              "Supporting Types",
              "Import",
              "Package Details",
            ];
            const actual = possibleHeadings.filter(ph => {
              return headings.some(h => {
                return ph.includes(h);
              })
            });
            const expected = possibleHeadings.filter(h => actual.includes(h))
            expect(JSON.stringify(actual)).to.equal(JSON.stringify(expected));
          })
        })

        // Verify it contains a description paragraph.
        describe("Description paragraph", function () {
          const paragraph = $(".docs-main-content .docs-content > p").text();
          it('paragraph exists', function () {
            expect(paragraph.length).to.be.at.least(1);
          })
        })

        // Verify constructor section exists.
        describe("Constructor", function () {
          it("constructor section exists", function () {
            const constructorSection = $("h2#create");
            expect(constructorSection.length).to.equal(1);
          })
          it("syntax section exists", function () {
            const syntax = $("h3#constructor-syntax");
            expect(syntax.length).to.equal(1);
          })
          it('contains syntax code', function () {
            const syntaxCode = $("h3#constructor-syntax + div > pulumi-chooser");
            expect(syntaxCode.length).to.be.at.least(1);
          })
        })

        // Verify Input/Output Properties section.
        describe("Properties", function () {
          it("properties section exists", function () {
            const properties = $("h2#properties");
            expect(properties.length).to.equal(1);
          })
          it("inputs section exists", function () {
            const inputs = $("h3#inputs");
            expect(inputs.length).to.equal(1);
          })
          it("outputs section exists", function () {
            const outputs = $("h3#outputs");
            expect(outputs.length).to.equal(1);
          })
        })

        // Verify the page contains an Example Usage section and contains
        // at least 1 code example.
        if (shouldContainExamples(pkg)) {
          describe("Examples section", () => {
            it("contains Example Usage heading", () => {
              const heading = $("h2#example-usage");
              expect(heading.length).to.equal(1);
              expect(heading.text()).to.have.string("Example Usage");
            });

            const examples = $("h2#example-usage ~ div > pulumi-chooser");
            it("contains at least one example", () => {
              expect(examples.length).to.be.at.least(1);
            });
          })
        }

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
            it("contains Import section", () => {
              expect(heading.length).to.equal(1);
              expect(heading.text()).to.have.string("Import");
            });
          }
        });

        // Verify the page contains a Package Details section
        describe("Package Details section", () => {       
          it("it exists", () => {
            const packageSection = $("h2#package-details");
            expect(packageSection.length).to.equal(1);
          });
        });
      });
    };
  });
});

function shouldContainExamples(pkg) {
  return [
    "aws",
    "azure",
    "gcp",
  ].includes(pkg);
}

function getModules(pkg) {
  const mods = fs.readdirSync(`./public/registry/packages/${pkg}/api-docs`, { withFileTypes: true })
    .filter(dirent => dirent.isDirectory())
    .map(dirent => dirent.name);

  return mods;
}

function constructPageRoute(filePath) {
  return filePath.replace("./public", "https://pulumi.com").replace("index.html", "");
}
