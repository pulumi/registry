// CloudFront Function (viewer-request) for registry URL rewriting.
//
// Responsibilities, in order:
//   1. 301 redirects for package URLs that no longer exist: legacy versioned
//      packages (e.g. aws-v6 -> aws@6.x) and removed Install & config pages.
//   2. Normalize the Accept header to two buckets -- "text/markdown" if the
//      client accepts markdown, otherwise the header is removed. The default
//      cache behavior includes Accept in the cache key, so normalization
//      prevents every unique browser Accept string from creating its own
//      cache entry.
//   3. Markdown content negotiation for hand-edited /registry/ pages:
//        - /foo.md -> internally rewrite to Hugo's /foo/index.md output.
//        - Accept: text/markdown on /foo/ or /foo/index.html -> internally
//          rewrite to /foo/index.md. The client's URL stays put; Accept is
//          in the cache key, so the HTML and markdown variants do not
//          pollute each other's cache entries.
//      API-generated pages under /api-docs/ are skipped: they have no backing
//      index.md, so rewrites would produce 404s.

var EXACT_REDIRECTS = [
    { match: "/registry/packages/aws-v6/how-to-guides/aws-py-oidc-provider-pulumi-cloud/index.html", target: "/docs/esc/guides/configuring-oidc/aws/" },
    { match: "/registry/packages/azure-native-v2/how-to-guides/azure-py-oidc-provider-pulumi-cloud/index.html", target: "/docs/esc/guides/configuring-oidc/azure/" },
];

var PREFIX_REDIRECTS = [
    { match: /^\/registry\/packages\/aws-v6\/?(.*)/, replace: "/registry/packages/aws@6.x/$1" },
    { match: /^\/registry\/packages\/azure-native-v2\/?(.*)/, replace: "/registry/packages/azure-native@2.x/$1" },

    // Packages whose Install & config page was deleted because the Overview
    // beside it already carried every section, near-verbatim and fresher --
    // see pulumi/registry#12225. The deleted pages were hand-written in
    // pulumi/registry (no upstream docs/installation-configuration.md exists
    // for any of these providers), so the URLs are gone for good rather than
    // pending regeneration. Send them to the package Overview.
    //
    // Hugo `aliases:` would be the usual mechanism, but resourcedocsgen
    // rewrites each _index.md wholesale on every provider release, so an
    // alias in front matter is erased at the next release. This table is not
    // generated, so a rule here survives.
    //
    // The suffix group matches every form the page's URL can arrive in: bare,
    // trailing slash, /index.html, and /index.md or .md from markdown content
    // negotiation. Markdown clients keep their Accept header across the 301
    // and get the Overview's markdown from the redirect target.
    {
        match: /^\/registry\/packages\/(aiven|alicloud|artifactory|auth0|cloudinit|consul|dbtcloud|digitalocean|random|tailscale|tls)\/installation-configuration(?:\.md|\/.*)?$/,
        replace: "/registry/packages/$1/",
    },
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

    // Normalize Accept for cache-key stability. Must happen before any
    // decision that reads Accept, since we also use the normalized value
    // for the markdown-negotiation branch below.
    var rawAccept = request.headers["accept"] ? request.headers["accept"].value : "";
    var wantsMarkdown = rawAccept.indexOf("text/markdown") !== -1;
    if (wantsMarkdown) {
        request.headers["accept"] = { value: "text/markdown" };
    } else if (request.headers["accept"]) {
        delete request.headers["accept"];
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
    // sibling /index.md from origin. Cache key includes the normalized Accept,
    // so HTML and markdown variants of the same URI get separate cache entries.
    if (wantsMarkdown) {
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
