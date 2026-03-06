// CloudFront Function for handling legacy versioned package URL redirects.
// Maps old naming convention (pkg-vN) to new format (pkg@N.x).

const EXACT_REDIRECTS = [
    { match: "/registry/packages/aws-v6/how-to-guides/aws-py-oidc-provider-pulumi-cloud/index.html", target: "/docs/esc/guides/configuring-oidc/aws/" },
    { match: "/registry/packages/azure-native-v2/how-to-guides/azure-py-oidc-provider-pulumi-cloud/index.html", target: "/docs/esc/guides/configuring-oidc/azure/" },
];

const PREFIX_REDIRECTS = [
    { match: /^\/registry\/packages\/aws-v6\/(.*)/, replace: "/registry/packages/aws@6.x/$1" },
    { match: /^\/registry\/packages\/azure-native-v2\/(.*)/, replace: "/registry/packages/azure-native@2.x/$1" },
];

function handler(event) {
    const uri = event.request.uri;

    for (const redirect of EXACT_REDIRECTS) {
        if (uri === redirect.match) {
            return {
                statusCode: 301,
                statusDescription: "Moved Permanently",
                headers: {
                    location: { value: redirect.target },
                    "cache-control": { value: "max-age=604800" },
                },
            };
        }
    }

    for (const redirect of PREFIX_REDIRECTS) {
        if (redirect.match.test(uri)) {
            return {
                statusCode: 301,
                statusDescription: "Moved Permanently",
                headers: {
                    location: { value: uri.replace(redirect.match, redirect.replace) },
                    "cache-control": { value: "max-age=604800" },
                },
            };
        }
    }

    return event.request;
}
