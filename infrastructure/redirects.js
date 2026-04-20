// CloudFront Function (viewer-request) for registry URL rewriting.
//
// Responsibilities, in order:
//   1. Legacy versioned package URL 301 redirects (e.g. aws-v6 -> aws@6.x).
//   2. Markdown content negotiation for hand-edited /registry/ pages:
//        - Hide /foo/index.md by 301'ing to /foo.md.
//        - Internally rewrite /foo.md to /foo/index.md at origin.
//        - Honor Accept: text/markdown by 301'ing to the canonical /foo.md.
//      API-generated pages under /api-docs/ are skipped: they have no backing
//      index.md, so rewrites would produce 404s.

var EXACT_REDIRECTS = [
    { match: "/registry/packages/aws-v6/how-to-guides/aws-py-oidc-provider-pulumi-cloud/index.html", target: "/docs/esc/guides/configuring-oidc/aws/" },
    { match: "/registry/packages/azure-native-v2/how-to-guides/azure-py-oidc-provider-pulumi-cloud/index.html", target: "/docs/esc/guides/configuring-oidc/azure/" },
];

var PREFIX_REDIRECTS = [
    { match: /^\/registry\/packages\/aws-v6\/?(.*)/, replace: "/registry/packages/aws@6.x/$1" },
    { match: /^\/registry\/packages\/azure-native-v2\/?(.*)/, replace: "/registry/packages/azure-native@2.x/$1" },
];

function redirect(location) {
    return {
        statusCode: 301,
        statusDescription: "Moved Permanently",
        headers: {
            location: { value: location },
            "cache-control": { value: "max-age=604800" },
        },
    };
}

function handler(event) {
    var request = event.request;
    var uri = request.uri;
    var i, rule;

    for (i = 0; i < EXACT_REDIRECTS.length; i++) {
        rule = EXACT_REDIRECTS[i];
        if (uri === rule.match) {
            return redirect(rule.target);
        }
    }

    for (i = 0; i < PREFIX_REDIRECTS.length; i++) {
        rule = PREFIX_REDIRECTS[i];
        if (rule.match.test(uri)) {
            return redirect(uri.replace(rule.match, rule.replace));
        }
    }

    // Markdown handling is scoped to /registry/* (plus the /registry.md alias
    // for the registry landing). Other paths on this distribution (/uploads/,
    // /fingerprinted/, etc.) are unrelated.
    if (!uri.startsWith("/registry/") && uri !== "/registry.md") {
        return request;
    }

    // API-generated pages have no backing index.md. Leave them alone.
    if (uri.indexOf("/api-docs/") !== -1) {
        return request;
    }

    if (uri.endsWith("/index.md")) {
        return redirect(uri.slice(0, -"/index.md".length) + ".md");
    }

    if (uri.endsWith(".md")) {
        request.uri = uri.slice(0, -".md".length) + "/index.md";
        return request;
    }

    var accept = request.headers["accept"] ? request.headers["accept"].value : "";
    if (accept.indexOf("text/markdown") !== -1) {
        if (uri.endsWith("/index.html")) {
            return redirect(uri.slice(0, -"/index.html".length) + ".md");
        }
        if (uri.endsWith("/")) {
            return redirect(uri.slice(0, -1) + ".md");
        }
    }

    return request;
}
