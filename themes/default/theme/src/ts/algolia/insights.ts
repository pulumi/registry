// Algolia search analytics (the "viewed" and "clicked" events Autocomplete sends
// to insights.algolia.io) is Performance tracking, so it waits for C0002. Until
// then, setup calls are queued and events are dropped; search itself doesn't
// depend on it. Under the Segment provider runWithPulumiConsent runs right away,
// so behavior there is unchanged. createConsentGatedInsightsClient is a copy of
// the one in pulumi/docs; keep the two in sync.

type InsightsClient = ((...args: any[]) => unknown) & { version?: string };

const EVENT_METHOD = /ObjectIDs|Filters|sendEvents/;

// The search-insights build the Autocomplete insights plugin loads by default.
export const SEARCH_INSIGHTS_VERSION = "2.15.0";
export const SEARCH_INSIGHTS_SRC = `https://cdn.jsdelivr.net/npm/search-insights@${SEARCH_INSIGHTS_VERSION}/dist/search-insights.min.js`;

// What the Autocomplete insights plugin does when it isn't given a client: a
// window.aa stub that queues calls until the search-insights script replaces it
// and replays the queue. The difference is that the script is only fetched once
// C0002 is granted. The returned client calls whatever window.aa is at the
// time, so it reaches the real one after the swap.
export function createSearchInsightsClient(): InsightsClient {
    const w = window as any;
    w.AlgoliaAnalyticsObject = "aa";
    w.aa = (...args: unknown[]) => {
        (w.aa.queue = w.aa.queue || []).push(args);
    };

    w.runWithPulumiConsent("C0002", () => {
        const script = document.createElement("script");
        script.async = true;
        script.src = SEARCH_INSIGHTS_SRC;
        document.body.appendChild(script);
    });

    const client: InsightsClient = (...args) => w.aa(...args);
    client.version = SEARCH_INSIGHTS_VERSION;
    return client;
}

export function createConsentGatedInsightsClient(client: InsightsClient): InsightsClient {
    let allowed = false;
    const queue: unknown[][] = [];

    const gated: InsightsClient = (method, ...args) => {
        if (allowed) return client(method, ...args);
        if (!EVENT_METHOD.test(method)) queue.push([method, ...args]);
    };
    // Autocomplete reads the version to decide how to pass credentials.
    gated.version = client.version;

    (window as any).runWithPulumiConsent("C0002", () => {
        allowed = true;
        for (const [method, ...args] of queue.splice(0)) client(method as string, ...args);
    });

    return gated;
}
