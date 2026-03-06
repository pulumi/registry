// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

import * as aws from "@pulumi/aws";
import * as pulumi from "@pulumi/pulumi";
import * as fs from "fs";

const stackConfig = new pulumi.Config();

const config = {
    originBucketNameOverride: stackConfig.get("originBucketNameOverride") || undefined,
    // pathToOriginBucketMetadata is the path to the file produced at the end of the
    // sync-and-test-bucket script (i.e., scripts/sync-and-test-bucket.sh).
    pathToOriginBucketMetadata: stackConfig.require("pathToOriginBucketMetadata"),
    // websiteLogsBucketName is the name of the S3 bucket used for storing access logs.
    websiteLogsBucketName: stackConfig.require("websiteLogsBucketName"),
    // e2eTestsBucketName is the S3 bucket stores the e2e test results.
    e2eTestsBucketName: stackConfig.require("e2eTestsBucketName"),
};

// Get role arns from datawarehouse stack.
const airflowStackRef = new pulumi.StackReference(`pulumi/dwh-workflows-orchestrate-airflow/production`)
const airflowTasksRole = airflowStackRef.getOutput('airflowTaskRoleArn')

const bucketReaderStackRef = new pulumi.StackReference(`pulumi/dwh-workflows-loader-prodbuckets/production`)
const bucketReaderRole = bucketReaderStackRef.getOutput('dwhBucketReaderRole')

// originBucketName is the name of the S3 bucket to use as the CloudFront origin for the
// website. This bucket is presumed to exist prior to the Pulumi run; if it doesn't, this
// program will fail.
export let originBucketName: string | undefined;

// If a build metadata file is present and contains valid content, use that for the bucket
// name by default. Will fail if there's a file present that isn't parsable as expected.
if (fs.existsSync(config.pathToOriginBucketMetadata)) {
    originBucketName = JSON.parse(fs.readFileSync(config.pathToOriginBucketMetadata).toString()).bucket;
}

// However, if the bucket's been configured manually, use that instead. A manually
// configured bucket means that someone's decided to pin it.
if (config.originBucketNameOverride) {
    originBucketName = config.originBucketNameOverride;
}

// If there's still no bucket selected, it's an error.
if (!originBucketName) {
    throw new Error("An origin bucket was not specified.");
}

// Fetch the bucket we'll use for the preview or update.
const originBucket = pulumi.output(aws.s3.getBucket({
    bucket: originBucketName,
}));

// We deny the s3:ListBucket permission to anyone but account users to prevent unintended 
// disclosure of the bucket's contents.
const originBucketPolicy = new aws.s3.BucketPolicy("origin-bucket-policy", {
    bucket: originBucket.bucket,
    policy: pulumi.all([originBucket.arn, aws.getCallerIdentity()])
        .apply(([arn, awsCallerIdentityResult]) => JSON.stringify({
            Version: "2008-10-17",
            Statement: [
                {
                    Effect: "Deny",
                    Principal: "*",
                    Action: "s3:ListBucket",
                    Resource: arn,
                    Condition: {
                        StringNotEquals: {
                            "aws:PrincipalAccount": awsCallerIdentityResult.accountId,
                        },
                    },
                },
            ],
    })),
});

const e2eTestsBucket = new aws.s3.Bucket("api-docs-e2e-test-results",
    {
        bucket: config.e2eTestsBucketName,
    },
    {
        protect: true,
    }
);

const e2eTestsBucketPolicy = new aws.s3.BucketPolicy("e2e-tests-bucket-policy", {
    bucket: e2eTestsBucket.bucket,
    policy: pulumi.all([e2eTestsBucket.bucket, airflowTasksRole, bucketReaderRole])
    .apply(([bucketName, airflowRole, readerRole]) => JSON.stringify({
        Version: "2008-10-17",
        Statement: [
            ...[airflowRole, readerRole].map(role => {
                return {
                    Effect: "Allow",
                    Principal: {
                        AWS: role
                    },
                    Action: [
                        "s3:GetObject",
                        "s3:GetObjectVersion"
                    ],
                    Resource: [
                        `arn:aws:s3:::${bucketName}/*`,
                    ],
                }
            }),
            ...[airflowRole, readerRole].map(role => {
                return {
                    Effect: "Allow",
                    Principal: {
                        AWS: role
                    },
                    Action: [
                        "s3:ListBucket",
                        "s3:GetBucketLocation",
                    ],
                    Resource: [
                        `arn:aws:s3:::${bucketName}`,
                    ],
                }
            })
        ]
    }))
});

// websiteLogsBucket stores the request logs for incoming requests.
const websiteLogsBucket = new aws.s3.Bucket(
    "website-logs-bucket",
    {
        bucket: config.websiteLogsBucketName,
    },
    {
        protect: true,
    },
);

// This needs to be set in order to allow the use of ACLs. This was added to update our infrastructure to be
// compatible with the default S3 settings from AWS' April update. `ObjectWriter` was the prior default, so
// changing it to that here to match the configuration prior to the update.
// https://aws.amazon.com/blogs/aws/heads-up-amazon-s3-security-changes-are-coming-in-april-of-2023/
const logsBucketOwnershipControls = new aws.s3.BucketOwnershipControls("logs-bucket-ownership-controls", {
    bucket: websiteLogsBucket.id,
    rule: {
        objectOwnership: "ObjectWriter"
    }
});

const logsBucketACL = new aws.s3.BucketAclV2("logs-bucket-acl", {
    bucket: websiteLogsBucket.id,
    acl: aws.s3.CannedAcl.Private,
}, {
    dependsOn: [logsBucketOwnershipControls],
});

// The canonical user ID for the account.
const owner = aws.s3.getCanonicalUserId({});
// Grant the CloudFront log delivery account permission to write to the bucket.
const logsBucketDeliveryACL = new aws.s3.BucketAclV2("logs-bucket-delivery-acl", {
    bucket: websiteLogsBucket.id,
    accessControlPolicy: {
        grants: [
            {
                grantee: {
                    // The canconical ID for the `awslogsdelivery` account.
                    // see: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/AccessLogs.html#AccessLogsOverview
                    id: "c4c1ede66af53448b93c283ce9448c4ba468c9432aa01d700d3878632f77d2d0",
                    type: "CanonicalUser",
                },
                permission: "WRITE",
            },
        ],
        owner: {
            id: owner.then(owner => owner.id),
        },
    },
}, {
    dependsOn: [logsBucketOwnershipControls],
});

const fiveMinutes = 60 * 5;
const oneHour = fiveMinutes * 12;
const oneWeek = oneHour * 24 * 7;
const oneYear = oneWeek * 52;

const baseCacheBehavior = {
    targetOriginId: originBucket.arn,
    compress: true,

    viewerProtocolPolicy: "redirect-to-https",

    allowedMethods: ["GET", "HEAD", "OPTIONS"],
    cachedMethods: ["GET", "HEAD", "OPTIONS"],

    // S3 doesn't need take any of these values into account when serving content.
    forwardedValues: {
        cookies: {
            forward: "none",
        },
        queryString: false,
    },

    minTtl: 0,
    defaultTtl: fiveMinutes,
    maxTtl: fiveMinutes,

    // https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-managed-response-headers-policies.html#managed-response-headers-policies-security
    responseHeadersPolicyId: "67f7725c-6f97-4210-82d7-5512b31e9d03", // SecurityHeadersPolicy
};

// distributionArgs configures the CloudFront distribution. Relevant documentation:
// https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html
// https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html
// https://docs.aws.amazon.com/general/latest/gr/aws_service_limits.html#limits_cloudfront
const distributionArgs: aws.cloudfront.DistributionArgs = {
    enabled: true,

    // We only specify one origin for this distribution: the S3 content bucket.
    origins: [
        {
            originId: originBucket.arn,
            domainName: originBucket.websiteEndpoint,
            customOriginConfig: {
                // > If your Amazon S3 bucket is configured as a website endpoint, [like we have here] you must specify
                // > HTTP Only. Amazon S3 doesn't support HTTPS connections in that configuration.
                // tslint:disable-next-line: max-line-length
                // https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#DownloadDistValuesOriginProtocolPolicy
                originProtocolPolicy: "http-only",
                httpPort: 80,
                httpsPort: 443,
                originSslProtocols: ["TLSv1.2"],
            },
        },
    ],

    // Default object to serve when no path is given.
    // https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/DefaultRootObject.html
    defaultRootObject: "index.html",

    defaultCacheBehavior: {
        ...baseCacheBehavior,
    },

    orderedCacheBehaviors: [
        // Build-metadata files are never cached.
        {
            ...baseCacheBehavior,
            pathPattern: "/metadata.json",
            defaultTtl: 0,
            minTtl: 0,
            maxTtl: 0,
        },
    ],

    // "All" is the most broad distribution, and also the most expensive.
    // "100" is the least broad, and also the least expensive.
    priceClass: "PriceClass_All",

    // Customize error pages.
    // https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/custom-error-pages.html
    customErrorResponses: [
        {
            errorCode: 404,
            responseCode: 404,
            errorCachingMinTtl: fiveMinutes,
            responsePagePath: "/404.html",
        },
    ],

    restrictions: {
        geoRestriction: {
            restrictionType: "none",
        },
    },

    // CloudFront certs must be in us-east-1, just like API Gateway.
    viewerCertificate: {
        cloudfrontDefaultCertificate: true,
        sslSupportMethod: "sni-only",
        minimumProtocolVersion: "TLSv1.2_2018",
    },

    loggingConfig: {
        bucket: websiteLogsBucket.bucketDomainName,
        includeCookies: false,
        prefix: `pulumi-registry/`,
    },
};

// cdn is the CloudFront distribution that serves the content of the website.
const cdn = new aws.cloudfront.Distribution(
    "cdn",
    distributionArgs,
    {
        protect: true,
        dependsOn: [ websiteLogsBucket ],
        // Work-around https://github.com/pulumi/pulumi-aws/issues/5229 to allow upgrading the version of pulumi-aws used.
        //
        // Before upgrading to v6, this has no effect.
        ignoreChanges: ["staging"],
    },
);

export const originBucketWebsiteDomain = originBucket.websiteDomain;
export const originBucketWebsiteEndpoint = originBucket.websiteEndpoint;
export const cloudFrontDomain = cdn.domainName;
export const e2eTestsBucketName = e2eTestsBucket.bucket;
