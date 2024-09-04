const glob = require("glob");
const fs = require("fs");

glob.sync("themes/default/content/registry/packages/**/_index.md").forEach(path => {

    // https://github.com/pulumi/registry/issues/5095
    // Finds and replaces an errant set of backticks in the TF docs.
    // We should be able to remove this once https://github.com/rancher/terraform-provider-rancher2/pull/1394 merges and is released.
    if (path === "themes/default/content/registry/packages/rancher2/api-docs/nodetemplate/_index.md") {
        const content = fs.readFileSync(path, "utf8");
        const lines = content.split("\n");
        let precedingLine, maybeBadLine = undefined;

        lines.forEach((line, i) => {
            if (line.endsWith("implements `fixNodeTemplateID()` that will update tfstate with proper id.")) {
                precedingLine = i;
                maybeBadLine = i + 1;
            }
        });

        // If the line looks like the bad line, and the number of code fences is odd, delete the line.
        const fenceCount = lines.map(line => line.startsWith("```")).length;
        if (precedingLine && maybeBadLine && lines[maybeBadLine].trim() === "```" && fenceCount % 2 !== 0) {
            delete lines[maybeBadLine];
            fs.writeFileSync(path, lines.join("\n"), "utf8");
        }
    }
});
