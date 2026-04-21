// CloudFront Function (viewer-request) for registry URL rewriting.
//
// Responsibilities, in order:
//   1. Legacy versioned package URL 301 redirects (e.g. aws-v6 -> aws@6.x).
//   2. Markdown content negotiation for hand-edited /registry/ pages:
//        - /foo.md -> internally rewrite to Hugo's /foo/index.md output.
//        - Accept: text/markdown on /foo/ or /foo/index.html -> 301 to /foo.md.
//      API-generated pages under /api-docs/ are skipped: they have no backing
//      index.md, so rewrites would produce 404s.
//
// Why 301 instead of URI rewrite for the Accept case: this distribution uses
// legacy forwardedValues with no Accept header in the cache key, so the cache
// key for /foo/ is the same regardless of what Accept was. If the function
// rewrote /foo/ -> /foo/index.md based on Accept, CloudFront would still key
// the cached response under the original /foo/ and pollute subsequent
// requests with a different Accept header. A 301 with cache-control: no-store
// avoids pollution -- the function runs fresh every time, and the underlying
// /foo/ HTML cache entry is untouched. minTtl=0 on the cache behavior ensures
// no-store is honored.

var EXACT_REDIRECTS = [
    { match: "/registry/packages/aws-v6/how-to-guides/aws-py-oidc-provider-pulumi-cloud/index.html", target: "/docs/esc/guides/configuring-oidc/aws/" },
    { match: "/registry/packages/azure-native-v2/how-to-guides/azure-py-oidc-provider-pulumi-cloud/index.html", target: "/docs/esc/guides/configuring-oidc/azure/" },
];

var PREFIX_REDIRECTS = [
    { match: /^\/registry\/packages\/aws-v6\/?(.*)/, replace: "/registry/packages/aws@6.x/$1" },
    { match: /^\/registry\/packages\/azure-native-v2\/?(.*)/, replace: "/registry/packages/azure-native@2.x/$1" },
];

function redirect(location, cacheControl) {
    return {
        statusCode: 301,
        statusDescription: "Moved Permanently",
        headers: {
            location: { value: location },
            "cache-control": { value: cacheControl },
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
            return redirect(rule.target, "max-age=604800");
        }
    }

    for (i = 0; i < PREFIX_REDIRECTS.length; i++) {
        rule = PREFIX_REDIRECTS[i];
        if (rule.match.test(uri)) {
            return redirect(uri.replace(rule.match, rule.replace), "max-age=604800");
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

    // Accept: text/markdown on a directory or /index.html URL -> 301 to the
    // canonical /foo.md. Use no-store so this 301 is not cached under the
    // original URI (see top-of-file comment).
    var accept = request.headers["accept"] ? request.headers["accept"].value : "";
    if (accept.indexOf("text/markdown") !== -1) {
        if (uri.endsWith("/index.html")) {
            return redirect(uri.slice(0, -"/index.html".length) + ".md", "no-store");
        }
        if (uri.endsWith("/")) {
            return redirect(uri.slice(0, -1) + ".md", "no-store");
        }
    }

    return request;
}
