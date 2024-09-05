const glob = require("glob");
const fs = require("fs");

/**
 * This purpose of this script is to allow us to apply targeted, temporary fixes
 * to API docs pages that need it, usually in response to a major rendering
 * issue involving upstream provider docs.
 *
 * Any fix listed here should be accompanied by a comment explaining the issue
 * as well as the fix, along with links to any related issues. Once those issues
 * are resolved and the problem is no longer reproducible in production, the fix
 * should be removed.
 */

glob.sync("themes/default/content/registry/packages/**/_index.md").forEach(path => {

    /**
     * Finds and replaces an errant set of backticks in the TF docs, breaking the layout
     * of the page and rendering it unusable.
     * https://github.com/pulumi/registry/issues/5095
     *
     * The script looks specifically for the following two-line pattern in the `nodetemplate`
     * page of the `rancher2` docs:
     *
     * implements `fixNodeTemplateID()` that will update tfstate with proper id.
     * ```
     *
     * The unmatched code fence above is what's breaking the page. The script just finds and
     * removes this line.
     *
     * The following PR on the upstream provider fixes the issue:
     * https://github.com/rancher/terraform-provider-rancher2/pull/1394
     *
     * Once the fix above has been merged and the provider republished to the Registry,
     * this fix can be removed.
     */
    if (path === "themes/default/content/registry/packages/rancher2/api-docs/nodetemplate/_index.md") {
        const content = fs.readFileSync(path, "utf8");
        const lines = content.split("\n");
        let precedingLine, maybeBadLine;

        lines.forEach((line, i) => {
            if (line.endsWith("implements `fixNodeTemplateID()` that will update tfstate with proper id.")) {
                precedingLine = i;
                maybeBadLine = i + 1;
            }
        });

        // If the line looks like the bad one and the number of code fences is
        // odd (indicating the page contains unmatched pairs), delete the line.
        const fenceCount = lines.map(line => line.startsWith("```")).length;
        if (precedingLine && maybeBadLine && lines[maybeBadLine].trim() === "```" && fenceCount % 2 !== 0) {
            delete lines[maybeBadLine];
            fs.writeFileSync(path, lines.join("\n"), "utf8");
        }
    }
});
