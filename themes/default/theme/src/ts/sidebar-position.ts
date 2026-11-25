// Keeps the package sidebar from jumping back to the top on every navigation.
//
// Each click in the sidebar is a full page load, and the sidebar (nav.main-nav,
// the scroll container) comes back scrolled to 0. On a large package like aws,
// that leaves the page you just opened thousands of pixels out of view. So when a
// sidebar link is clicked, we remember where it sat within the sidebar, and on the
// page it leads to we scroll the sidebar so that page's link sits in the same
// spot. We anchor on the link rather than restoring scrollTop, because the tree
// changes shape between pages (the previous page's module collapses and the new
// one expands), and a raw offset would land on the wrong row.
//
// Arriving any other way (search, breadcrumbs, a link in the content) has no saved
// position, so the current page's link is scrolled into the middle of the sidebar
// instead, if it isn't already visible.

const STORAGE_KEY = "pulumi-registry-sidebar-anchor";

interface SidebarAnchor {
    // Pathname of the link that was clicked, which is the page the anchor is for.
    path: string;
    // The link's distance from the top of the sidebar's visible area, in pixels.
    offset: number;
}

function sidebar(): HTMLElement | null {
    return document.querySelector<HTMLElement>("#docs-main-nav nav.main-nav");
}

function offsetWithin(container: HTMLElement, el: HTMLElement): number {
    return (
        el.getBoundingClientRect().top - container.getBoundingClientRect().top
    );
}

function saveAnchor(anchor: SidebarAnchor) {
    try {
        sessionStorage.setItem(STORAGE_KEY, JSON.stringify(anchor));
    } catch (e) {
        // Storage can be unavailable (private browsing, blocked site data); the
        // next page then falls back to revealing the current link.
    }
}

// Reads and clears the saved anchor, so that it applies to exactly one page load.
function takeAnchor(): SidebarAnchor | null {
    try {
        const raw = sessionStorage.getItem(STORAGE_KEY);
        sessionStorage.removeItem(STORAGE_KEY);
        return raw ? JSON.parse(raw) : null;
    } catch (e) {
        return null;
    }
}

function currentPageLink(nav: HTMLElement): HTMLElement | null {
    const links = Array.from(
        nav.querySelectorAll<HTMLAnchorElement>("a[href]")
    );
    return links.find((a) => a.pathname === window.location.pathname) || null;
}

function positionSidebar() {
    const anchor = takeAnchor();
    const nav = sidebar();
    const link = nav && currentPageLink(nav);
    if (!link) {
        return;
    }

    const top = offsetWithin(nav, link);
    if (anchor && anchor.path === window.location.pathname) {
        nav.scrollTop += top - anchor.offset;
        return;
    }

    // The search box is pinned over the top of the sidebar, so a link beneath it
    // is out of view too, and the band below it is what we center within.
    const pinned =
        nav.querySelector<HTMLElement>(".docs-search")?.offsetHeight ?? 0;
    if (top < pinned || top + link.offsetHeight > nav.clientHeight) {
        nav.scrollTop +=
            top - (pinned + nav.clientHeight - link.offsetHeight) / 2;
    }
}

function rememberClickedLink(e: MouseEvent) {
    const nav = sidebar();
    const link = (e.target as Element).closest<HTMLAnchorElement>("a[href]");
    if (!nav || !link || !nav.contains(link)) {
        return;
    }
    saveAnchor({ path: link.pathname, offset: offsetWithin(nav, link) });
}

// Capture phase, because the API tree's nodes navigate from their own click
// handlers and stop propagation before a bubbling listener would see the click.
document.addEventListener("click", rememberClickedLink, true);

// The API tree is fetched and rendered after the page loads, and on package pages
// the current link may live inside it, so wait for it when it's present.
if (document.querySelector("pulumi-api-doc-filterable-nav")) {
    document.addEventListener("apiDocNavReady", positionSidebar, {
        once: true,
    });
} else {
    positionSidebar();
}
