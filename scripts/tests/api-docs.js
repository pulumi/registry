var chai = require('chai')
const glob = require("glob");
const fs = require("fs");
const parser = require("node-html-parser");
var path = require("path");

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

      describe(constructPageRoute(p), function () {
        describe('h1 title', function () {
          const h1s = doc.querySelectorAll("h1");
          it('contains exactly 1 "h1"', function () {
            chai.expect(h1s.length).to.equal(1);
          })
          it('contains the package name', function () {
            chai.expect(h1s[0].innerHTML.toLowerCase()).to.have.string(pkg);
          })
          it('contains the module name', function () {
            chai.expect(h1s[0].innerHTML.toLowerCase()).to.have.string(mod);
          })
        })
        
        if (shouldContainExamples(pkg)) {
          describe("Examples section", () => {
            it("contains Example Usage heading", () => {
              const heading = doc.querySelectorAll("h2#example-usage");
              chai.expect(heading.length).to.equal(1);
              chai.expect(heading[0].innerHTML).to.have.string("Example Usage");
            });
  
            const examples = doc.querySelectorAll("h2#example-usage ~ div > pulumi-chooser");
            it("contains at least one example", () => {
              if (!path.basename(path.dirname(p)).startsWith("get")) {
                chai.expect(examples.length).to.be.at.least(1);
              }
            });
          })
        }

        describe("Import section", () => {
          const heading = doc.querySelectorAll("h2#import");
          if (heading.length > 0) {
            it("comes after Properties list", () => {
              const importSection = doc.querySelectorAll("h2#properties ~ h2#import");
              chai.expect(importSection.length).to.equal(1);
            });
          } else {
            it("contains Import section", () => {
              chai.expect(heading.length).to.equal(1);
              chai.expect(heading[0].innerHTML).to.have.string("Import");
            });
          }
        })
      });
    });
  })
})


function getPackagesList() {
  return [
    "acme",
    "aiven",
    "aws",
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