// Copyright 2016-2026, Pulumi Corporation.  All rights reserved.
//
// Versioned provider API docs — storage stack.
//
// This stack owns ONLY the permanent, write-once archive bucket and the GitHub-OIDC
// publisher role used to write to it. The main registry stack
// (infrastructure/index.ts) consumes the exports here via a StackReference and adds
// the CloudFront origin group + `/registry/packages/*@*` and `/registry/versioned/*`
// behaviors. Keeping the bucket in its own stack means its lifecycle is independent
// of the per-deploy atomic origin buckets.
//
// See scripts/versioned-docs/README.md for the overall architecture.
//
// Stacks: pulumi/pulumi-registry-versioned/{testing,production}.

import * as aws from "@pulumi/aws";
import * as pulumi from "@pulumi/pulumi";

const config = new pulumi.Config();

// env drives the bucket name. Taken from the stack name (testing | production)
// unless explicitly overridden via config.
const env = config.get("env") || pulumi.getStack();

const bucketName = `pulumi-registry-versioned-${env}`;

// Permanent archive bucket. `protect: true` guards against accidental teardown.
// Website hosting is required because the registry distribution reaches S3 origins
// via their website endpoints, which is what resolves a directory URL like
// /registry/packages/aws@4.x/api-docs/ to .../api-docs/index.html.
const bucket = new aws.s3.Bucket("registry-versioned-docs", {
    bucket: bucketName,
}, {
    protect: true,
});

const bucketWebsite = new aws.s3.BucketWebsiteConfiguration("registry-versioned-docs-website", {
    bucket: bucket.id,
    indexDocument: { suffix: "index.html" },
    errorDocument: { key: "404.html" },
});

// Versioning lets us recover an accidental overwrite and — unlike S3 Object Lock in
// compliance mode — does NOT block redaction. Immutability is operational: only the
// publisher role writes, and publish-version.sh refuses to overwrite an existing
// registry/packages/<pkg>@<slug>/ prefix without --force.
new aws.s3.BucketVersioning("registry-versioned-docs-versioning", {
    bucket: bucket.id,
    versioningConfiguration: { status: "Enabled" },
});

// Abort dangling multipart uploads after 7 days. No object expiration — archives are
// permanent.
new aws.s3.BucketLifecycleConfiguration("registry-versioned-docs-lifecycle", {
    bucket: bucket.id,
    rules: [{
        id: "abort-incomplete-multipart",
        status: "Enabled",
        filter: {},
        abortIncompleteMultipartUpload: { daysAfterInitiation: 7 },
    }],
});

// Public-read (GetObject) like the existing origin buckets, but deliberately NOT
// ListBucket — the archive is browsable by URL, not enumerable.
const ownership = new aws.s3.BucketOwnershipControls("registry-versioned-docs-ownership", {
    bucket: bucket.id,
    rule: { objectOwnership: "ObjectWriter" },
});

const publicAccessBlock = new aws.s3.BucketPublicAccessBlock("registry-versioned-docs-public-access-block", {
    bucket: bucket.id,
    blockPublicAcls: false,
    blockPublicPolicy: false,
    ignorePublicAcls: false,
    restrictPublicBuckets: false,
});

new aws.s3.BucketPolicy("registry-versioned-docs-policy", {
    bucket: bucket.id,
    policy: bucket.arn.apply(arn => JSON.stringify({
        Version: "2012-10-17",
        Statement: [{
            Sid: "PublicReadObjects",
            Effect: "Allow",
            Principal: "*",
            Action: "s3:GetObject",
            Resource: `${arn}/*`,
        }],
    })),
}, { dependsOn: [publicAccessBlock, ownership] });

// --- Publisher role: assumed by pulumi/registry GitHub Actions via OIDC ---
//
// The archive workflows assume ONLY this narrow role to publish snapshots; the broad
// deploy role used by the main site build is not reused.
//
// The GitHub OIDC provider already exists in both accounts, so we reference it by its
// well-known ARN rather than creating it (CreateOpenIDConnectProvider would fail if
// one already exists).
const accountId = aws.getCallerIdentityOutput().accountId;
const oidcProviderArn = accountId.apply(
    id => `arn:aws:iam::${id}:oidc-provider/token.actions.githubusercontent.com`,
);

// distributionArn (optional): when set, scopes cloudfront:CreateInvalidation to that
// distribution instead of "*". The distribution lives in the main stack, so by default
// we avoid coupling its ARN here — CreateInvalidation is low-risk.
const distributionArn = config.get("distributionArn");

const publisherRole = new aws.iam.Role("registry-versioned-docs-publisher", {
    name: `registry-versioned-docs-publisher-${env}`,
    description: "Assumed by pulumi/registry GitHub Actions to publish versioned provider API doc snapshots.",
    assumeRolePolicy: oidcProviderArn.apply(arn => JSON.stringify({
        Version: "2012-10-17",
        Statement: [{
            Effect: "Allow",
            Principal: { Federated: arn },
            Action: "sts:AssumeRoleWithWebIdentity",
            Condition: {
                StringEquals: {
                    "token.actions.githubusercontent.com:aud": "sts.amazonaws.com",
                },
                StringLike: {
                    // Restrict to the pulumi/registry repository (any ref).
                    "token.actions.githubusercontent.com:sub": "repo:pulumi/registry:*",
                },
            },
        }],
    })),
});

new aws.iam.RolePolicy("registry-versioned-docs-publisher-policy", {
    role: publisherRole.id,
    policy: pulumi.all([bucket.arn, distributionArn]).apply(([arn, distArn]) => JSON.stringify({
        Version: "2012-10-17",
        Statement: [
            {
                Sid: "BucketObjectReadWrite",
                Effect: "Allow",
                Action: ["s3:PutObject", "s3:GetObject", "s3:DeleteObject", "s3:DeleteObjectVersion", "s3:ListBucketVersions"],
                Resource: `${arn}/*`,
            },
            {
                Sid: "BucketList",
                Effect: "Allow",
                Action: ["s3:ListBucket", "s3:GetBucketLocation", "s3:ListBucketVersions"],
                Resource: arn,
            },
            {
                Sid: "Invalidate",
                Effect: "Allow",
                Action: "cloudfront:CreateInvalidation",
                Resource: distArn || "*",
            },
        ],
    })),
});

// Consumed by the main stack's StackReference and by scripts/versioned-docs/*.
export const bucketNameOutput = bucket.bucket;
export const bucketArn = bucket.arn;
export const bucketWebsiteEndpoint = bucketWebsite.websiteEndpoint;
export const publisherRoleArn = publisherRole.arn;
