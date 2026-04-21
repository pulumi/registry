// CloudFront Function (viewer-request) for registry URL rewriting.
//
// Responsibilities, in order:
//   1. Legacy versioned package URL 301 redirects (e.g. aws-v6 -> aws@6.x).
//   2. Markdown content negotiation for hand-edited /registry/ pages:
//        - Serve /foo.md by rewriting the URI to Hugo's /foo/index.md output.
//        - Honor Accept: text/markdown by rewriting the URI to the same
//          /foo/index.md target; the client's URL stays put.
//      API-generated pages under /api-docs/ are skipped: they have no backing
//      index.md, so rewrites would produce 404s.
//
// All Accept-driven behavior is implemented as URI rewrites (not 301s) so that
// the rewritten URI becomes the cache key. A response-returning branch here
// would be cached under the original URI and pollute the cache for requests
// with a different Accept header.

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

    // /foo.md -> origin fetches /foo/index.md. Skip URIs that already end in
    // /index.md (direct-file access) so we don't rewrite to /foo/index/index.md.
    if (uri.endsWith(".md") && !uri.endsWith("/index.md")) {
        request.uri = uri.slice(0, -".md".length) + "/index.md";
        return request;
    }

    // Accept: text/markdown on a directory or /index.html URL -> serve the
    // sibling /index.md from origin. We rewrite the URI rather than 301
    // because CloudFront caches function responses under the original URI,
    // and Accept is not in the cache key.
    var accept = request.headers["accept"] ? request.headers["accept"].value : "";
    if (accept.indexOf("text/markdown") !== -1) {
        if (uri.endsWith("/index.html")) {
            request.uri = uri.slice(0, -"/index.html".length) + "/index.md";
            return request;
        }
        if (uri.endsWith("/")) {
            request.uri = uri + "index.md";
            return request;
        }
    }

    return request;
}
