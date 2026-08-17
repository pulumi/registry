import {LocalStorageService} from "./state";


const navigationState = new LocalStorageService("navigation-toggle-state");
loadToggleStates();

// Badge icons for the API-docs property lists (see the injection loop at the bottom of
// this file). Both are inlined as `currentColor` SVGs rather than <img src="/icons/...">
// so they tint from the badge's `color` and therefore follow the dark theme; an <img>
// paints its own baked-in fill. The `ph-icon` class is what sizes and baseline-aligns
// them (_phosphor.scss: 1.25em square, which is ~18px against the property list's
// text-sm) -- the same thing copybutton.ts does for the copy glyph.

// Two arced arrows around an exclamation mark: the art from
// static/icons/replacement-property.svg, with its hardcoded #595A5E fills dropped. That
// file stays put even though nothing here loads it any more -- pulumi/docs renders the
// same badge from its own misc.ts against the same /icons/ path on the merged site.
const replacementIconSvg =
    '<svg xmlns="http://www.w3.org/2000/svg" class="ph-icon" fill="currentColor" viewBox="0 0 58 58" aria-hidden="true" focusable="false">' +
    '<path d="M28.8574 49.1586C34.4619 49.1586 39.5384 46.8831 43.2137 43.2056L38.4286 38.4145L57.2857 32.9045L53.5143 53.5143L48.9603 48.9555C43.8229 54.0958 36.8584 57.2857 29 57.2857C14.6843 57.2857 2.72372 46.6799 0.754046 32.9045H8.95764C10.8461 42.1694 19.0294 49.1586 28.8574 49.1586Z"/>' +
    '<path d="M29.1426 8.84135C23.5381 8.84135 18.4617 11.1169 14.7863 14.7944L19.5714 19.5855L0.714294 25.0955L4.48572 4.48571L9.03972 9.04452C14.1771 3.90415 21.2842 0.714277 29.1426 0.714277C43.4582 0.714277 55.2763 11.3201 57.246 25.0955H49.0424C47.1539 15.8306 38.9706 8.84135 29.1426 8.84135Z"/>' +
    '<rect x="25.8571" y="16.4286" width="6.28572" height="12.5714"/>' +
    '<rect x="25.8571" y="35.2857" width="6.28572" height="6.28571"/>' +
    "</svg>";

// Phosphor "lock", bold weight -- the same symbol the rest of the site draws through the
// icon sprite (#p-lock-bold in assets/icons/sprite.svg). Inlined rather than referenced
// with <use>, because that sprite's URL is fingerprinted at build time by
// partials/icon-context.html and this bundle has no way to know it. Bold rather than
// regular so it doesn't read spindly next to the chunky replacement mark.
const secretIconSvg =
    '<svg xmlns="http://www.w3.org/2000/svg" class="ph-icon" fill="currentColor" viewBox="0 0 256 256" aria-hidden="true" focusable="false">' +
    '<path d="M208,76H180V56A52,52,0,0,0,76,56V76H48A20,20,0,0,0,28,96V208a20,20,0,0,0,20,20H208a20,20,0,0,0,20-20V96A20,20,0,0,0,208,76ZM100,56a28,28,0,0,1,56,0V76H100ZM204,204H52V100H204Zm-76-92a32,32,0,0,0-12,61.66V180a12,12,0,0,0,24,0v-6.34A32,32,0,0,0,128,112Zm0,24a8,8,0,1,1-8,8A8,8,0,0,1,128,136Z"/>' +
    "</svg>";

// The flags the docs generator can stamp on a property <dt>, in the order they render.
// One entry per flag rather than one query per combination of flags: three flags would
// otherwise need seven selectors, and a fourth would need fifteen.
const propertyBadges = [
    {
        // The asterisk itself is drawn in CSS, by `dt.property-required .property-indicator`
        // in _lists.scss. Keep the class name here so that rule still matches.
        className: "property-required",
        markup: '<span class="property-indicator"></span>',
        tooltip: "This property is required.",
    },
    {
        className: "property-replacement",
        markup: '<span class="property-indicator-replacement">' + replacementIconSvg + "</span>",
        tooltip: "Changes to this property will trigger replacement.",
    },
    {
        className: "property-secret",
        markup: '<span class="property-indicator-secret">' + secretIconSvg + "</span>",
        tooltip: "This property's value is encrypted in state and masked in CLI and Console output.",
    },
];

function bindToggle(el: HTMLElement) {
    el.querySelectorAll(".toggleButton").forEach(btn => {
        btn.addEventListener("click", function () {
            if (this.closest(".toggle, .toggleVisible") !== el) {
                return;
            }

            if (el.classList.contains("toggle")) {
                el.classList.add("toggleVisible");
                el.classList.remove("toggle");
            } else {
                el.classList.add("toggle");
                el.classList.remove("toggleVisible");
            }
        });
    });
}

function loadToggleStates() {
    const isCurrentPage = (el: HTMLElement) => {
        const browserUrl = window.location.href;
        const anchor = el.querySelector('a');
        const anchorRef = anchor ? anchor.getAttribute('href') : '';
        return browserUrl.includes(anchorRef);
    };

    document.querySelectorAll(".toggle-topLevel, .toggleVisible-topLevel").forEach((el: HTMLElement) => {
        if (navigationState.getKey(el.id) == "expanded" || isCurrentPage(el)) {
            el.classList.add("toggleVisible");
            el.classList.remove("toggle");
        } else if (navigationState.getKey(el.id) == "collapsed") {
            el.classList.add("toggle");
            el.classList.remove("toggleVisible");
        }

        el.addEventListener("click", function () {
            const folderOpenIcon = el.querySelector(".folder-open");
            const folderClosedIcon = el.querySelector(".folder");
            if (folderOpenIcon) {
                folderOpenIcon.classList.add("folder");
                folderOpenIcon.classList.remove("folder-open");
            } else if (folderClosedIcon) {
                folderClosedIcon.classList.add("folder-open");
                folderClosedIcon.classList.remove("folder");
            }
        });
    });

    document.querySelectorAll(".toggleVisible, .toggleVisible-topLevel").forEach((el: HTMLElement) => {
        if (isCurrentPage(el)) {
            const leftNav = document.getElementById("left-nav");
            if (leftNav) {
                leftNav.scrollTop = el.getBoundingClientRect().top + window.scrollY - 145;
            }
        }
    });
}

function updateToggleState(el: HTMLElement, toggleState: string) {
    navigationState.updateKey(el.id, toggleState)
}

function bindTopLevelToggle(el: HTMLElement) {
    el.querySelectorAll(".toggleButton-topLevel").forEach(btn => {
        btn.addEventListener("click", function () {
            if (this.closest(".toggle-topLevel, .toggleVisible-topLevel") !== el) {
                return;
            }

            if (el.classList.contains("toggle")) {
                el.classList.add("toggleVisible");
                el.classList.remove("toggle");
                updateToggleState(el, "expanded");
            } else {
                el.classList.add("toggle");
                el.classList.remove("toggleVisible");
                updateToggleState(el, "collapsed");
            }
        });
    });
}

function bindTopLevelToggles(selector: string) {
    document.querySelectorAll(selector).forEach((el: HTMLElement) => {
        bindTopLevelToggle(el);
    });
}

function bindToggles(selector: string) {
    document.querySelectorAll(selector).forEach((el: HTMLElement) => {
        bindToggle(el);
    });
}

export function generateOnThisPage() {
    const tocs = document.querySelectorAll<HTMLElement>(".table-of-contents");
    tocs.forEach(toc => toc.style.display = "none");

    const uls = document.querySelectorAll(".table-of-contents .content ul.table-of-contents-list");
    if (uls.length === 0) return;

    let found = false;
    const headingItems: { element: HTMLElement, listItems: HTMLElement[] }[] = [];

    document.querySelectorAll("h2, h3").forEach((el: HTMLElement) => {
        if (el.closest('.hidden')) {
            return;
        }
        const id = el.getAttribute("id");
        const text = el.textContent;
        const linkTitle = el.dataset.linkTitle;
        const tag = el.tagName.toLowerCase();

        if (id && text) {
            found = true;
            const listItems: HTMLElement[] = [];
            uls.forEach(ul => {
                const li = document.createElement("li");
                li.className = tag;
                const a = document.createElement("a");
                a.href = '#' + id;
                a.textContent = linkTitle || text;
                li.appendChild(a);
                ul.appendChild(li);
                listItems.push(li);
            });

            headingItems.push({ element: el, listItems });
        }
    });

    if (found) {
        tocs.forEach(toc => toc.style.display = "");

        const setActiveItem = () => {
            let active = null;
            for (const heading of headingItems) {
                if (!active && heading.element.getBoundingClientRect().top >= 0) {
                    active = heading;
                }
                heading.listItems.forEach(li => li.classList.toggle("active", heading === active));
            }
        };

        window.addEventListener("scroll", setActiveItem);
        setActiveItem();
    }
}

(function () {
    const observer = new IntersectionObserver(
        ([e]) => {
            e.target.classList.toggle("is-pinned", e.intersectionRatio < 1);
            const pinnedSearchContainerEl = document.querySelector(".header-pinned") as HTMLElement;
            const dotOverlay = document.querySelector(".hide-on-pinned") as HTMLElement;
            const heroTitle = document.querySelector(".header-hero-title") as HTMLElement;

            if (e.isIntersecting) {
                pinnedSearchContainerEl?.classList.add("hidden");
                pinnedSearchContainerEl?.classList.remove("flex");

                dotOverlay?.classList.remove("hidden");
                dotOverlay?.classList.add("flex");

                heroTitle?.classList.remove("hidden");
                heroTitle?.classList.add("flex");

            } else {
                pinnedSearchContainerEl?.classList.remove("hidden");
                pinnedSearchContainerEl?.classList.add("flex");

                dotOverlay?.classList.add("hidden");
                dotOverlay?.classList.remove("flex");

                heroTitle?.classList.add("hidden");
                heroTitle?.classList.remove("flex");

            }
        },
        { threshold: [1] },
    );

    const headerContainerEl = document.querySelector(".header-container");
    if (!headerContainerEl) {
        const registryNavBar = document.querySelector(".top-nav-bar.registry");
        if (registryNavBar) {
            observer.observe(registryNavBar);
        }
    } else {
        observer.observe(headerContainerEl);
    }

    bindToggles(".toggle");
    bindToggles(".toggleVisible");

    bindTopLevelToggles(".toggle-topLevel");
    bindTopLevelToggles(".toggleVisible-topLevel");

    generateOnThisPage();

    document.querySelector(".nav-header-toggle")?.addEventListener("click", () => {
        document.querySelector(".nav-header-items")?.classList.toggle("hidden");
    });
    document.querySelector(".blog-sidebar-toggle")?.addEventListener("click", () => {
        document.querySelector(".blog-sidebar-content")?.classList.toggle("hidden");
    });
    document.querySelector(".docs-sidebar-toggle")?.addEventListener("click", () => {
        document.querySelector(".docs-sidebar-content")?.classList.toggle("hidden");
    });

    document.querySelectorAll("ul[data-shuffle='true']").forEach((list: HTMLElement) => {
        const items = list.querySelectorAll(":scope > li") as NodeListOf<HTMLElement>;
        items.forEach(item => {
            item.style.order = String(Math.ceil(Math.random() * items.length));
        });
        list.classList.remove("invisible");
    });

    const aboutNavItem = document.querySelector(`#about-nav li[data-filter-name="who-we-are"]`);
    if (aboutNavItem) {
        aboutNavItem.classList.add("active-about-nav-item");
    }

    document.querySelectorAll("#about-nav li").forEach(li => {
        li.addEventListener("click", function () {
            const activeClassName = "active-about-nav-item";
            this.classList.add(activeClassName);

            const activeLink = (this as HTMLElement).dataset.filterName;
            const allLinks = ["who-we-are", "what-we-believe", "community", "history", "awards", "newsroom", "join-us"];
            const inactiveLinks = allLinks.filter(link => link !== activeLink);

            inactiveLinks.forEach(link => {
                document.querySelector(`#about-nav li[data-filter-name="${link}"]`)?.classList.remove(activeClassName);
            });
        });
    });

    // Swap the docs generator's empty <span class="property-indicator"> placeholder for
    // one tooltip per flag the <dt> carries. Every flagged property gets the same
    // wrapper, so a single CSS rule lays out any combination of badges.
    document.querySelectorAll("dl.resources-properties dt").forEach(dt => {
        const badges = propertyBadges.filter(badge => dt.classList.contains(badge.className));
        if (badges.length === 0) {
            return;
        }

        // Read the placeholder before building the replacement: the required badge emits
        // a .property-indicator of its own, and we must not select that one.
        const indicator = dt.querySelector(".property-indicator");
        if (!indicator) {
            return;
        }

        // The generator also emits a native title ("Required", "Optional, Deprecated").
        // Drop it so the browser's own tooltip doesn't race the bubbles.
        dt.removeAttribute("title");

        const container = document.createElement("div");
        container.className = "property-badges";

        // Sibling tooltips, never nested: pulumi-tooltip binds its listeners to the first
        // descendant .tooltip-target it finds, so a nested one would hijack the outer.
        container.innerHTML = badges
            .map(badge =>
                "<pulumi-tooltip>" +
                badge.markup +
                '<span slot="content">' + badge.tooltip + "</span>" +
                "</pulumi-tooltip>")
            .join("");

        indicator.replaceWith(container);
    });
})();
