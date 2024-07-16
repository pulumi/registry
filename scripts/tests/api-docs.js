var chai = require('chai')
const glob = require("glob");
const fs = require("fs");
const parser = require("node-html-parser");
var path = require("path");

const expect = chai.expect

const packages = getPackagesList();


packages.forEach(pkg => {

  const modules = getModules(pkg);
  modules.forEach(mod => {
    const paths = glob.sync(`./public/registry/packages/${pkg}/api-docs/${mod}/**/*.html`, {
      ignore: [
        `./public/registry/packages/${pkg}/api-docs/${mod}/index.html`
      ],
    });
    paths.forEach(p => {
      const fileContent = fs.readFileSync(p, "utf-8").toString();
      const doc = parser.parse(fileContent);

      if (!path.basename(path.dirname(p)).startsWith("get")) {

        describe(constructPageRoute(p), function () {
          // Verify page has a title and it is an h1 that contains the package and
          // module name.
          describe("h1 title", function () {
            const h1s = doc.querySelectorAll("h1");
            it("contains exactly 1 h1", function () {
              expect(h1s.length).to.equal(1);
            })
            it("contains the package name", function () {
              expect(h1s[0].innerHTML.toLowerCase()).to.have.string(pkg);
            })
            it("contains the module name", function () {
              expect(h1s[0].innerHTML.toLowerCase()).to.have.string(mod);
            })
          })

          // Verify it contains a description paragraph.
          describe("Description paragraph", function () {
            const paragraph = doc.querySelectorAll("section.docs-content > p");
            it('paragraph exists', function () {
              expect(paragraph.length).to.be.at.least(1);
            })
          })

          // Verify constructor section exists.
          describe("Constructor", function () {
            it("constructor section exists", function () {
              const constructorSection = doc.querySelectorAll("h2#create");
              expect(constructorSection.length).to.equal(1);
            })
            it("syntax section exists", function () {
              const syntax = doc.querySelectorAll("h3#constructor-syntax");
              expect(syntax.length).to.equal(1);
            })
            it('contains syntax code', function () {
              const syntaxCode = doc.querySelectorAll("h3#constructor-syntax + div > pulumi-chooser");
              expect(syntaxCode.length).to.be.at.least(1);
            })
          })

          // Verify Input/Output Properties section.
          describe("Properties", function () {
            it("properties section exists", function () {
              const properties = doc.querySelectorAll("h2#properties");
              expect(properties.length).to.equal(1);
            })
            it("inputs section exists", function () {
              const inputs = doc.querySelectorAll("h3#inputs");
              expect(inputs.length).to.equal(1);
            })
            it("outputs section exists", function () {
              const outputs = doc.querySelectorAll("h3#outputs");
              expect(outputs.length).to.equal(1);
            })
          })

          // Verify the page contains an Example Usage section and contains
          // at least 1 code example.
          if (shouldContainExamples(pkg)) {
            describe("Examples section", () => {
              it("contains Example Usage heading", () => {
                const heading = doc.querySelectorAll("h2#example-usage");
                expect(heading.length).to.equal(1);
                expect(heading[0].innerHTML).to.have.string("Example Usage");
              });

              const examples = doc.querySelectorAll("h2#example-usage ~ div > pulumi-chooser");
              it("contains at least one example", () => {
                expect(examples.length).to.be.at.least(1);
              });
            })
          }

          // Verify the page contains an Import section and that it comes somewhere
          // after the Properties section.
          describe("Import section", () => {
            const heading = doc.querySelectorAll("h2#import");
            if (heading.length > 0) {
              it("comes after Properties list", () => {
                const importSection = doc.querySelectorAll("h2#properties ~ h2#import");
                expect(importSection.length).to.equal(1);
              });
            } else {
              it("contains Import section", () => {
                expect(heading.length).to.equal(1);
                expect(heading[0].innerHTML).to.have.string("Import");
              });
            }
          });
        });
      }
    });
  })
})


function getPackagesList() {
  // const pkgs = fs.readdirSync(`./public/registry/packages/`, { withFileTypes: true })
  //   .filter(dirent => dirent.isDirectory())
  //   .map(dirent => dirent.name);

  // return pkgs;
  return [
    // "acme",
    // "aiven",
    "aws",
    // "azure"
  ];
}

function shouldContainExamples(pkg) {
  return [
    "aws",
    "azure",
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