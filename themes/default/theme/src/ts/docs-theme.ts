// Registry color-theme toggle (light / dark / system). Ported verbatim from
// pulumi/docs (theme/src/ts/docs-theme.ts) except for the section gate, which
// checks for the registry body class instead of the docs one.
//
// The storage key is intentionally the same on both sites, so a reader who picks
// dark on /docs lands on a dark registry and vice versa.
//
// The inline script in layouts/partials/head.html already applied the theme before
// first paint; this module only owns the toggle's click behavior, the aria-pressed
// state, and live re-render when the OS appearance changes under "system".

type ThemePref = "light" | "dark" | "system";

const STORAGE_KEY = "pulumi-docs-theme";
const DARK_QUERY = "(prefers-color-scheme: dark)";

function isRegistry(): boolean {
    return document.body.classList.contains("section-registry");
}

function readPref(): ThemePref {
    try {
        const stored = localStorage.getItem(STORAGE_KEY);
        if (stored === "light" || stored === "dark" || stored === "system") {
            return stored;
        }
    } catch (e) {
        return "system";
    }
    return "system";
}

function systemPrefersDark(): boolean {
    return (
        typeof window.matchMedia === "function" &&
        window.matchMedia(DARK_QUERY).matches
    );
}

function resolve(pref: ThemePref): "light" | "dark" {
    if (pref === "system") {
        return systemPrefersDark() ? "dark" : "light";
    }
    return pref;
}

function reflectButtons(pref: ThemePref): void {
    const buttons =
        document.querySelectorAll<HTMLButtonElement>("[data-theme-set]");
    buttons.forEach((button) => {
        button.setAttribute(
            "aria-pressed",
            String(button.dataset.themeSet === pref)
        );
    });
}

function apply(pref: ThemePref): void {
    const el = document.documentElement;
    el.dataset.themePref = pref;
    if (resolve(pref) === "dark") {
        el.dataset.theme = "dark";
        document.body.dataset.theme = "dark";
    } else {
        delete el.dataset.theme;
        delete document.body.dataset.theme;
    }
    reflectButtons(pref);
}

function persist(pref: ThemePref): void {
    try {
        localStorage.setItem(STORAGE_KEY, pref);
    } catch (e) {
        return;
    }
}

// Dark mode puts a light chip behind the external package logos that would
// otherwise vanish (see registry/package/icon.html), but only once the image has
// actually painted — a chip behind a pending or broken remote image is just a
// white sliver. The inline onload handler in the partial covers images that load
// after parse; this catches the ones the browser had cached and finished before
// that handler existed.
function markLoadedLogos(): void {
    const logos = document.querySelectorAll<HTMLImageElement>(
        "img[data-logo-chip]"
    );
    logos.forEach((logo) => {
        if (logo.complete && logo.naturalWidth > 0) {
            logo.dataset.logoLoaded = "";
        }
    });
}

function init(): void {
    if (!isRegistry()) {
        return;
    }

    apply(readPref());
    markLoadedLogos();

    const buttons = document.querySelectorAll<HTMLButtonElement>(
        ".docs-theme-toggle [data-theme-set]"
    );
    buttons.forEach((button) => {
        button.addEventListener("click", () => {
            const next = button.dataset.themeSet as ThemePref;
            persist(next);
            apply(next);
        });
    });

    if (typeof window.matchMedia === "function") {
        const media = window.matchMedia(DARK_QUERY);
        const onChange = (): void => {
            if (readPref() === "system") {
                apply("system");
            }
        };
        if (typeof media.addEventListener === "function") {
            media.addEventListener("change", onChange);
        } else if (typeof media.addListener === "function") {
            media.addListener(onChange);
        }
    }
}

if (document.readyState === "loading") {
    document.addEventListener("DOMContentLoaded", init);
} else {
    init();
}
