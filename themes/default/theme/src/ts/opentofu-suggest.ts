// Suggests Terraform/OpenTofu providers when the registry package filter comes up
// empty. Pulumi can run any provider in the OpenTofu registry via Any Terraform
// Provider, so a zero-result filter is the highest-intent moment to say so.
//
// The panel itself is shown and hidden entirely by CSS (see _packages.scss: a
// general sibling combinator reveals .no-results once every .package is hidden).
// This module never touches that. It owns exactly two things inside the panel:
// the .no-results-message text and the children of .tf-suggest.
//
// It deliberately does NOT rebuild .no-results itself — packages.ts binds the
// "Clear all filters" listener to that button once at module load, so replacing
// the panel's innerHTML would silently break it.

import { addCopyButton } from "./copybutton";

const API = "https://api.opentofu.org/registry/docs/search?q=";
const PROVIDER_BASE = "https://search.opentofu.org/provider/";
const MIN_QUERY_LENGTH = 2;
const MAX_RESULTS = 5;
const TIMEOUT_MS = 5000;

const DEFAULT_MESSAGE =
    "Looks like we don't have any packages that match your filters. " +
    "Adjust or clear the filters and try again.";

// `addr` and `version` end up in a URL and in a command we ask people to paste into
// a shell. Anyone can publish a provider to the OpenTofu registry under any name, so
// these are allowlists: an entry that doesn't match is dropped, not escaped.
const ADDR_RE = /^[A-Za-z0-9][A-Za-z0-9._-]*\/[A-Za-z0-9][A-Za-z0-9._-]*$/;
const VERSION_RE = /^[0-9A-Za-z][0-9A-Za-z.+-]*$/;
// The registry keys versions with their leading "v" ("v1.4.1"), which is what the
// provider page URL needs; `pulumi package add` wants them without it. Keep both.
const RAW_VERSION_RE = /^v?[0-9][0-9A-Za-z.+-]*$/;

interface TofuProvider {
    addr: string;
    version: string; // stripped of its leading "v", for the command; "" when unusable
    urlVersion: string; // exactly as the registry keys it ("v1.4.1"), or "latest"
    stars: number;
}

const cache = new Map<string, TofuProvider[]>();
let inFlight: AbortController | null = null;
let currentKey = ""; // the query that owns the panel right now
let renderedKey: string | null = null; // what is currently painted

// Matches the compact form the site already uses for the pulumi/pulumi star count
// in docs-top-nav.html (via scripts/fetch-github-stars.js).
const formatStars = (n: number): string => {
    if (n < 1000) return String(n);
    if (n < 100000) return `${(n / 1000).toFixed(1).replace(/\.0$/, "")}K`;
    if (n < 1000000) return `${Math.round(n / 1000)}K`;
    return `${(n / 1000000).toFixed(1).replace(/\.0$/, "")}M`;
};

const truncate = (s: string, max: number): string =>
    s.length > max ? `${s.slice(0, max)}…` : s;

const clear = (el: HTMLElement): void => {
    while (el.firstChild) el.removeChild(el.firstChild);
};

// True for providers published under a namespace matching their own name, e.g.
// zerotier/zerotier — the vendor's own publication rather than someone's fork.
const isVendorNamespace = (p: TofuProvider): boolean => {
    const parts = p.addr.split("/");
    return parts[0].toLowerCase() === parts[1].toLowerCase();
};

const normalizeResults = (body: any[]): TofuProvider[] =>
    body
        .filter(
            r =>
                r &&
                r.type === "provider" &&
                typeof r.addr === "string" &&
                ADDR_RE.test(r.addr),
        )
        .map(r => {
            const raw = typeof r.version === "string" ? r.version : "";
            const bare = raw.replace(/^v/, "");
            const stars = Number(r.popularity);
            return {
                addr: r.addr as string,
                version: VERSION_RE.test(bare) ? bare : "",
                // Pass the registry's own key through untouched rather than
                // re-deriving it: "1.4.1" 404s where "v1.4.1" resolves, and not every
                // provider is guaranteed to carry the prefix.
                urlVersion: RAW_VERSION_RE.test(raw) ? raw : "latest",
                stars: isFinite(stars) ? stars : 0,
            };
        })
        // The API's own ranking is not trustworthy — for "random" it puts a 0-star
        // fork above hashicorp/random. Sort by:
        //   1. stars, which is what separates a real provider from its forks;
        //   2. a namespace matching the provider name (zerotier/zerotier,
        //      namecheap/namecheap), which marks the vendor's own publication. Real
        //      ties happen — adamdecaf/namecheap and namecheap/namecheap both sit at
        //      168 stars, and every zerotier provider has 0 — and without this the
        //      fork wins the top slot on alphabetical order alone;
        //   3. addr, purely so equal entries don't reshuffle between renders.
        .sort(
            (a, b) =>
                b.stars - a.stars ||
                Number(isVendorNamespace(b)) - Number(isVendorNamespace(a)) ||
                a.addr.localeCompare(b.addr),
        )
        .slice(0, MAX_RESULTS);

const buildCommand = (provider: TofuProvider): HTMLElement => {
    const highlight = document.createElement("div");
    highlight.className = "highlight";
    const pre = document.createElement("pre");
    pre.className = "chroma";
    const code = document.createElement("code");
    code.className = "language-bash";
    code.setAttribute("data-lang", "bash");
    code.setAttribute("data-track", "registry-no-results-tf-provider");

    // textContent, never innerHTML — addr/version are third-party strings. The "$ "
    // prompt is the site convention; copybutton's normalizeText strips it on copy.
    const base = `$ pulumi package add terraform-provider ${provider.addr}`;
    code.textContent = provider.version ? `${base} ${provider.version}` : base;

    pre.appendChild(code);
    highlight.appendChild(pre);
    // Reuses the sitewide copy button so this looks and behaves like every other
    // command on pulumi.com. Its internal insertAdjacentHTML takes a static literal.
    addCopyButton(highlight);
    return highlight;
};

const buildRow = (provider: TofuProvider): HTMLElement => {
    const li = document.createElement("li");
    li.className = "tf-suggest-item";

    const header = document.createElement("div");
    header.className = "tf-suggest-header";

    const [namespace, name] = provider.addr.split("/");
    const link = document.createElement("a");
    // Origin is a hardcoded literal and both segments already passed ADDR_RE, so a
    // javascript: href is structurally impossible; encode anyway.
    link.href =
        PROVIDER_BASE +
        `${encodeURIComponent(namespace)}/${encodeURIComponent(name)}/` +
        encodeURIComponent(provider.urlVersion);
    link.target = "_blank";
    link.rel = "noopener noreferrer";
    link.textContent = provider.addr;
    link.setAttribute(
        "aria-label",
        `${provider.addr} on the OpenTofu registry (opens in a new tab)`,
    );
    header.appendChild(link);

    const meta = document.createElement("span");
    meta.className = "tf-suggest-meta";
    meta.textContent = `★ ${formatStars(provider.stars)}`;
    meta.setAttribute(
        "aria-label",
        `${provider.stars} GitHub ${provider.stars === 1 ? "star" : "stars"}`,
    );
    header.appendChild(meta);

    li.appendChild(header);
    li.appendChild(buildCommand(provider));
    return li;
};

// Everything the OpenTofu lookup renders sits below a rule, so it reads as coming
// from a different registry than the (empty) Pulumi results above it.
const openSection = (region: HTMLElement, message: string): void => {
    clear(region);

    const divider = document.createElement("hr");
    divider.className = "tf-suggest-divider";
    region.appendChild(divider);

    const heading = document.createElement("h3");
    heading.className = "tf-suggest-message";
    heading.id = "tf-suggest-heading";
    heading.textContent = message;
    region.appendChild(heading);
};

const renderLoading = (region: HTMLElement, query: string): void => {
    openSection(region, `Searching the OpenTofu registry for "${truncate(query, 60)}"…`);
};

const render = (
    region: HTMLElement,
    providers: TofuProvider[],
    key: string,
    query: string,
): void => {
    renderedKey = key;
    const shown = `"${truncate(query, 60)}"`;

    if (providers.length === 0) {
        // Said explicitly rather than left blank: we told the user we were looking,
        // so report the result. The fallback paragraph below still offers the docs.
        openSection(region, `No providers in the OpenTofu registry match ${shown}.`);
        return;
    }

    openSection(
        region,
        providers.length === 1
            ? `1 provider in the OpenTofu registry matches ${shown}:`
            : `${providers.length} providers in the OpenTofu registry match ${shown}:`,
    );

    const list = document.createElement("ul");
    list.className = "tf-suggest-list";
    list.setAttribute("aria-labelledby", "tf-suggest-heading");
    providers.forEach(p => list.appendChild(buildRow(p)));
    region.appendChild(list);
};

const fetchAndRender = async (
    region: HTMLElement,
    query: string,
    key: string,
): Promise<void> => {
    const controller = new AbortController();
    inFlight = controller;
    const timer = window.setTimeout(() => controller.abort(), TIMEOUT_MS);

    renderLoading(region, query);
    renderedKey = null;

    try {
        const res = await fetch(API + encodeURIComponent(query), {
            signal: controller.signal,
            headers: { Accept: "application/json" },
        });
        if (!res.ok) throw new Error(String(res.status));
        const body = await res.json();
        if (!Array.isArray(body)) throw new Error("unexpected shape");

        const providers = normalizeResults(body);
        cache.set(key, providers);
        if (key !== currentKey) return; // a newer query owns the panel now
        render(region, providers, key, query);
    } catch (e) {
        // Aborted, offline, non-200, or malformed. Degrade quietly: the static
        // fallback paragraph below is server-rendered and still on screen. An error
        // banner in an already-frustrating dead end is worse than nothing. Failures
        // are deliberately not cached, so the next keystroke retries.
        if (key !== currentKey) return;
        clear(region);
        renderedKey = key;
    } finally {
        window.clearTimeout(timer);
        if (inFlight === controller) inFlight = null;
    }
};

// Called from updateAllCount() in packages.ts on every filter path. `filterText` is
// passed in rather than read from the DOM because the Stencil search component
// re-renders asynchronously — reading .registry-filter-input right after reset()
// yields the query the user just cleared.
export const syncEmptyState = (visibleCount: number, filterText: string): void => {
    const region = document.querySelector<HTMLElement>(
        ".all-packages .no-results .tf-suggest",
    );
    if (!region) return;

    const message = document.querySelector<HTMLElement>(
        ".all-packages .no-results .no-results-message",
    );

    const reset = (): void => {
        if (inFlight) {
            inFlight.abort();
            inFlight = null;
        }
        currentKey = "";
        renderedKey = null;
        clear(region);
    };

    // Results exist: the panel is CSS-hidden, so there is nothing to look up.
    if (visibleCount > 0) {
        reset();
        return;
    }

    const query = (filterText || "").trim();

    if (message) {
        message.textContent =
            query.length > 0
                ? `No packages in the Pulumi Registry match "${truncate(query, 80)}".`
                : DEFAULT_MESSAGE;
    }

    // Tag-only filters, or a single character: no meaningful query to send.
    if (query.length < MIN_QUERY_LENGTH) {
        reset();
        return;
    }

    const key = query.toLowerCase();
    if (key === renderedKey) return; // already painted for this query
    currentKey = key;

    const cached = cache.get(key);
    if (cached) {
        render(region, cached, key, query);
        return;
    }

    if (inFlight) inFlight.abort();
    void fetchAndRender(region, query, key);
};
