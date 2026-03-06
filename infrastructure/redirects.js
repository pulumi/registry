// CloudFront Function for handling legacy versioned package URL redirects.
// Maps old naming convention (pkg-vN) to new format (pkg@N.x).

var REDIRECTS = [
    { match: /^\/registry\/packages\/aws-v6\/(.*)/, replace: "/registry/packages/aws@6.x/$1" },
    { match: /^\/registry\/packages\/azure-native-v2\/(.*)/, replace: "/registry/packages/azure-native@2.x/$1" },
];

function handler(event) {
    var uri = event.request.uri;

    for (var i = 0; i < REDIRECTS.length; i++) {
        var redirect = REDIRECTS[i];
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
