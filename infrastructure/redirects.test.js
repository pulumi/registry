// Unit tests for the CloudFront viewer-request function in redirects.js.
// Run with: node --test redirects.test.js

const { test } = require("node:test");
const assert = require("node:assert/strict");
const fs = require("node:fs");
const path = require("node:path");
const vm = require("node:vm");

// CloudFront Functions use a synchronous, top-level `handler` function rather
// than module exports. Load the source and evaluate it in a fresh context to
// get a reference we can call.
const source = fs.readFileSync(path.join(__dirname, "redirects.js"), "utf8");
const context = vm.createContext({});
vm.runInContext(source + "\n;this.handler = handler;", context);
const handler = context.handler;

function request(uri, acceptHeader) {
    const headers = {};
    if (acceptHeader !== undefined) {
        headers["accept"] = { value: acceptHeader };
    }
    return { request: { uri, headers } };
}

function assertRedirect(result, expectedLocation, expectedCacheControl) {
    assert.equal(result.statusCode, 301);
    assert.equal(result.headers.location.value, expectedLocation);
    assert.equal(
        result.headers["cache-control"].value,
        expectedCacheControl || "max-age=604800",
    );
}

function assertPassThrough(result, expectedUri) {
    // Pass-through returns the request object (possibly with rewritten uri).
    assert.ok(result.uri !== undefined, "expected a request, got a response");
    assert.equal(result.uri, expectedUri);
}

test("legacy exact redirect: aws-v6 oidc guide", () => {
    const result = handler(request("/registry/packages/aws-v6/how-to-guides/aws-py-oidc-provider-pulumi-cloud/index.html"));
    assertRedirect(result, "/docs/esc/guides/configuring-oidc/aws/");
});

test("legacy prefix redirect: aws-v6 -> aws@6.x", () => {
    const result = handler(request("/registry/packages/aws-v6/resources/s3/bucket/"));
    assertRedirect(result, "/registry/packages/aws@6.x/resources/s3/bucket/");
});

test("legacy prefix redirect: azure-native-v2 -> azure-native@2.x", () => {
    const result = handler(request("/registry/packages/azure-native-v2/"));
    assertRedirect(result, "/registry/packages/azure-native@2.x/");
});

test("Accept: text/markdown on directory URL -> URI rewrite to /index.md", () => {
    const result = handler(request("/registry/packages/aws/", "text/markdown"));
    assertPassThrough(result, "/registry/packages/aws/index.md");
});

test("Accept: text/markdown on /index.html -> URI rewrite to /index.md", () => {
    const result = handler(request("/registry/packages/aws/index.html", "text/markdown"));
    assertPassThrough(result, "/registry/packages/aws/index.md");
});

test("Accept: text/markdown with q-value still matches", () => {
    const result = handler(request("/registry/packages/aws/", "text/html, text/markdown;q=0.9"));
    assertPassThrough(result, "/registry/packages/aws/index.md");
});

test("/foo/index.md -> pass through (served directly from origin)", () => {
    const result = handler(request("/registry/packages/aws/index.md"));
    assertPassThrough(result, "/registry/packages/aws/index.md");
});

test("/foo/index.md + Accept:md -> no double rewrite", () => {
    const result = handler(request("/registry/packages/aws/index.md", "text/markdown"));
    assertPassThrough(result, "/registry/packages/aws/index.md");
});

test("/foo.md -> internal rewrite to /foo/index.md", () => {
    const result = handler(request("/registry/packages/aws.md"));
    assertPassThrough(result, "/registry/packages/aws/index.md");
});

test("nested leaf: .md suffix on a how-to-guide", () => {
    const result = handler(request("/registry/packages/aws/how-to-guides/6-0-migration.md"));
    assertPassThrough(result, "/registry/packages/aws/how-to-guides/6-0-migration/index.md");
});

test("registry root: /registry/ + Accept -> URI rewrite to /registry/index.md", () => {
    const result = handler(request("/registry/", "text/markdown"));
    assertPassThrough(result, "/registry/index.md");
});

test("registry root: /registry.md -> internal rewrite to /registry/index.md", () => {
    const result = handler(request("/registry.md"));
    assertPassThrough(result, "/registry/index.md");
});

test("api-docs directory + Accept: text/markdown -> untouched", () => {
    const result = handler(request("/registry/packages/aws/api-docs/s3/bucket/", "text/markdown"));
    assertPassThrough(result, "/registry/packages/aws/api-docs/s3/bucket/");
});

test("api-docs .md hand-typed -> untouched (origin will 404)", () => {
    const result = handler(request("/registry/packages/aws/api-docs/s3/bucket.md"));
    assertPassThrough(result, "/registry/packages/aws/api-docs/s3/bucket.md");
});

test("no Accept, no .md suffix -> untouched", () => {
    const result = handler(request("/registry/packages/aws/"));
    assertPassThrough(result, "/registry/packages/aws/");
});

test("Accept: application/markdown (non-standard) -> no rewrite", () => {
    const result = handler(request("/registry/packages/aws/", "application/markdown"));
    assertPassThrough(result, "/registry/packages/aws/");
});

test("Accept: text/html only -> no rewrite", () => {
    const result = handler(request("/registry/packages/aws/", "text/html"));
    assertPassThrough(result, "/registry/packages/aws/");
});

test("missing Accept header on directory URL -> no rewrite", () => {
    const result = handler(request("/registry/packages/aws/"));
    assertPassThrough(result, "/registry/packages/aws/");
});

test("out-of-scope path /uploads/foo.md -> untouched", () => {
    // Markdown logic scoped to /registry/*; other cache behaviors should pass through.
    const result = handler(request("/uploads/something.md"));
    assertPassThrough(result, "/uploads/something.md");
});

test("out-of-scope path /fingerprinted/x + Accept -> untouched", () => {
    const result = handler(request("/fingerprinted/logos/aws.abc123.svg", "text/markdown"));
    assertPassThrough(result, "/fingerprinted/logos/aws.abc123.svg");
});

test("legacy redirect wins over markdown rewrite", () => {
    // A legacy aws-v6 .md URL should hit the prefix 301, not the .md rewrite.
    const result = handler(request("/registry/packages/aws-v6/index.md"));
    assertRedirect(result, "/registry/packages/aws@6.x/index.md");
});

test("Accept normalization: text/markdown bucket collapses to 'text/markdown'", () => {
    const event = request("/registry/packages/aws/", "text/html, text/markdown;q=0.9");
    handler(event);
    assert.equal(event.request.headers["accept"].value, "text/markdown");
});

test("Accept normalization: non-markdown Accept is stripped", () => {
    const event = request("/registry/packages/aws/", "text/html,application/xhtml+xml");
    handler(event);
    assert.equal(event.request.headers["accept"], undefined);
});

test("Accept normalization: missing Accept stays missing", () => {
    const event = request("/registry/packages/aws/");
    handler(event);
    assert.equal(event.request.headers["accept"], undefined);
});
