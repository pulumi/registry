function setDocsMainNavPosition() {
    const docsMainNavToggleWrapper = document.querySelector(".docs-main-nav-toggle-wrapper");
    const docsNavToggleIcon = document.querySelector(".docs-nav-toggle-icon");

    if (window.innerWidth <= 1280) {
        if (docsMainNavToggleWrapper?.classList.contains("docs-nav-show")) {
            docsNavToggleIcon?.classList.remove("open-docs-main-nav");
            docsNavToggleIcon?.classList.add("close-docs-main-nav");
        } else if (docsMainNavToggleWrapper?.classList.contains("docs-nav-hide")) {
            docsNavToggleIcon?.classList.remove("close-docs-main-nav");
            docsNavToggleIcon?.classList.add("open-docs-main-nav");
        }
    }

    const mainNav = document.querySelector<HTMLElement>(".main-nav");
    const mainNavToggle = document.querySelector<HTMLElement>(".docs-nav-toggle");
    const docsTypeNavSearch = document.querySelector<HTMLElement>(".docs-type-nav-search");
    const docsToggleOffset = 94;

    const docsListMain = document.querySelector(".section-docs .docs-list-main") as HTMLElement;
    if (docsListMain) {
        const topNavContainer = document.querySelector(".top-nav-container") as HTMLElement;
        if (docsListMain.getBoundingClientRect().y <= 0) {
            const searchHeight = docsTypeNavSearch?.offsetHeight || 0;
            const topNavY = topNavContainer ? Math.max(topNavContainer.getBoundingClientRect().y, 0) : 0;
            if (mainNav) mainNav.style.marginTop = (searchHeight - topNavY) + "px";
            if (mainNavToggle) mainNavToggle.style.top = (docsToggleOffset + searchHeight - topNavY) + "px";
        } else {
            if (mainNav) mainNav.style.marginTop = "0";
        }
    }

    if (window.innerWidth > 1280) {
        docsMainNavToggleWrapper?.classList.remove("docs-nav-show");
        docsMainNavToggleWrapper?.classList.remove("docs-nav-hide");
    } else if (!docsMainNavToggleWrapper?.classList.contains("docs-nav-hide") && !docsMainNavToggleWrapper?.classList.contains("docs-nav-show")) {
        docsMainNavToggleWrapper?.classList.add("docs-nav-hide");
    }
}

function setTableOfContentsVisibility() {
    const docsTableOfContents = document.querySelector<HTMLElement>(".docs-toc-desktop");
    const docsMainNavToggleWrapper = document.querySelector(".docs-main-nav-toggle-wrapper");
    if (docsTableOfContents) {
        if (window.innerWidth > 1024 && window.innerWidth <= 1280) {
            docsTableOfContents.style.display = docsMainNavToggleWrapper?.classList.contains("docs-nav-show") ? "none" : "";
        } else if (window.innerWidth > 1280) {
            docsTableOfContents.style.display = "";
        } else {
            docsTableOfContents.style.display = "none";
        }
    }
}

// Document-relative bottom edge of the top nav. The bar isn't sticky, so its
// viewport-relative bottom is just this minus the scroll offset — which keeps
// getBoundingClientRect(), and the layout flush it forces, off the scroll path.
// Re-measured on resize and load, the only times the bar's own geometry moves.
let topNavBottomInDocument = 0;

function measureTopNav() {
    const topNav = document.querySelector<HTMLElement>("header.docs-top-nav");
    topNavBottomInDocument = topNav ? topNav.getBoundingClientRect().bottom + window.scrollY : 0;
}

// Size the sticky sidebar to exactly the viewport space below the top nav, so its
// bottom edge always lands on the fold. The nav bar scrolls away, so the amount to
// subtract shrinks to zero as the page scrolls — hence the recalculation on scroll
// as well as on resize. Mirrors the `height: calc(100vh - 65px)` fallback in
// docs/_docs-main.scss, which holds until this runs.
//
// This has to be exact rather than generous: the sidebar is a flex column whose
// last child is the theme toggle (partials/docs/theme-toggle.html), and any excess
// height pushes that control off the bottom of the screen.
function setMainNavHeight() {
    const docsMainNav = document.querySelector<HTMLElement>(".docs-main-nav");
    if (!docsMainNav) {
        return;
    }
    const topNavBottom = Math.max(topNavBottomInDocument - window.scrollY, 0);
    const height = (window.innerHeight - topNavBottom) + "px";
    // Past the nav bar the value stops changing, so skip the write (and the style
    // invalidation it costs) for the rest of the page.
    if (docsMainNav.style.height !== height) {
        docsMainNav.style.height = height;
    }
}

function handleResize() {
    setDocsMainNavPosition();
    setTableOfContentsVisibility();
    measureTopNav();
    setMainNavHeight();
}

function handleScroll() {
    setDocsMainNavPosition();
    setMainNavHeight();
}

window.addEventListener("resize", handleResize);
window.addEventListener("scroll", handleScroll);
window.addEventListener("load", handleResize);
handleResize();

(function () {
    const docsMainNavToggleWrapper = document.querySelector(".docs-main-nav-toggle-wrapper");
    const docsNavToggleIcon = document.querySelector(".docs-nav-toggle-icon");

    document.querySelector(".docs-nav-toggle")?.addEventListener("click", function () {
        docsMainNavToggleWrapper?.classList.toggle("docs-nav-show");
        docsMainNavToggleWrapper?.classList.toggle("docs-nav-hide");
        docsNavToggleIcon?.classList.toggle("close-docs-main-nav");
        docsNavToggleIcon?.classList.toggle("open-docs-main-nav");
        setTableOfContentsVisibility();
    });

    const packageCardCheckbox = document.getElementById("accordion-checkbox-package-card") as HTMLInputElement;
    const packageCardBackground = document.getElementById("accordion-package-card") as HTMLElement;

    if (packageCardCheckbox && packageCardBackground) {
        // Written as custom properties rather than literal hexes so the accordion
        // follows the color theme: docs/_docs-theme.scss declares both under
        // html[data-theme="dark"], and light mode leaves them undefined and falls
        // back to the original values here. An inline style can't be overridden
        // from a stylesheet, so the indirection has to live on this side.
        packageCardCheckbox.addEventListener("change", function () {
            packageCardBackground.style.background = packageCardCheckbox.checked
                ? "var(--docs-accordion-bg-open, #fff)"
                : "var(--docs-accordion-bg-closed, #f9f9f9)";
        });
    }

    function loadContentWidthState() {
        const contentWidthState = window.localStorage.getItem("content-width-state");
        if (contentWidthState === "expanded") {
            expandContentWidth();
        } else {
            collapseContentWidth();
        }
    }

    const collapseContentButton = document.getElementById("collapse-content-button");
    const expandContentButton = document.getElementById("expand-content-button");

    function expandContentWidth() {
        document.querySelectorAll(".docs-main-content").forEach(el => {
            el.classList.add("docs-content-width-expanded");
            if (window.location.pathname.startsWith("/registry")) {
                el.classList.add("expand-registry");
            }
        });
        const banner = document.getElementById("docs-home-banner");
        if (banner) {
            banner.querySelectorAll("p").forEach(p => p.classList.add("wider"));
            banner.style.backgroundImage = `url("/images/docs/docs-home-header-background-desktop-wide.svg")`;
        }
        collapseContentButton?.classList.remove("hide");
        expandContentButton?.classList.add("hide");
        window.localStorage.setItem("content-width-state", "expanded");
    }

    function collapseContentWidth() {
        document.querySelectorAll(".docs-main-content").forEach(el => {
            el.classList.remove("docs-content-width-expanded");
        });
        const banner = document.getElementById("docs-home-banner");
        if (banner) {
            banner.querySelectorAll("p").forEach(p => p.classList.remove("wider"));
            banner.style.backgroundImage = `url("/images/docs/docs-home-header-background-desktop.svg")`;
        }
        collapseContentButton?.classList.add("hide");
        expandContentButton?.classList.remove("hide");
        window.localStorage.setItem("content-width-state", "collapsed");
    }

    expandContentButton?.addEventListener("click", expandContentWidth);
    collapseContentButton?.addEventListener("click", collapseContentWidth);

    loadContentWidthState();
})();
