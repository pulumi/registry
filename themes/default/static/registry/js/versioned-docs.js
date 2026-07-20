// Versioned provider API docs — version-selector loader.
//
// PERMANENT CONTRACT URL: every archived doc page references this file by the
// exact path /registry/js/versioned-docs.js. Rewrite this file freely, but
// NEVER move or rename it — archived snapshots are immutable and will request
// this URL forever. Styling for loader-created chrome lives in the sibling
// /registry/js/versioned-docs.css so archives can be restyled at any time
// without republishing. (webpack.config.js `output.clean.keep` protects both
// files from the theme build's output cleaning.)
//
// The script tag carries its config as data attributes:
//   data-vdocs-pkg        package name, e.g. "aws"
//   data-vdocs-mode       "latest" (a live Hugo page, emitted at build time by
//                         partials/registry/package/version-selector.html)
//                         | "archive" (a frozen snapshot, injected at publish
//                         time by scripts/versioned-docs/inject-version-switcher.sh)
//   data-vdocs-slug       this page's version slug ("" = latest, "6.x", ...)
//   data-vdocs-version    this page's full version (archive mode)
//   data-vdocs-live-root  canonical live root, e.g. "/registry/packages/aws/"
//   data-vdocs-rel        this page's path within its version root (archive mode)
//
// It fetches the per-package manifest at /registry/versioned/<pkg>/versions.json
// and unifies the version dropdown:
//   * latest mode: the build-time <select> (see
//     partials/registry/package/version-selector.html) is authoritative for
//     in-Hugo versions; manifest-only (archived) versions are merged in. When
//     the build rendered no <select> (single-version package), one is created
//     at the [data-vdocs-mount] wrapper.
//   * archive mode: the baked <select> is frozen at snapshot time, so its
//     options are replaced wholesale from the manifest.
//
// Hard rule: fail SILENT. PR previews and any environment without a published
// archive have no versions.json; on any error we render nothing and the page
// is unaffected.
(function () {
    "use strict";
    try {
        if (window.__pulumiVdocsLoaded) return;
        window.__pulumiVdocsLoaded = true;

        var script =
            document.querySelector('script[data-vdocs-pkg][data-vdocs-mode="archive"]') ||
            document.querySelector("script[data-vdocs-pkg]");
        if (!script) return;
        var cfg = {
            pkg: script.getAttribute("data-vdocs-pkg"),
            mode: script.getAttribute("data-vdocs-mode") || "latest",
            slug: script.getAttribute("data-vdocs-slug") || "",
            version: script.getAttribute("data-vdocs-version") || "",
            liveRoot: script.getAttribute("data-vdocs-live-root") || "",
        };
        if (!cfg.pkg) return;

        fetch("/registry/versioned/" + encodeURIComponent(cfg.pkg) + "/versions.json", { credentials: "omit" })
            .then(function (r) {
                if (!r.ok) throw new Error("manifest " + r.status);
                return r.json();
            })
            .then(function (manifest) {
                if (!manifest || !Array.isArray(manifest.versions) || manifest.versions.length < 2) return;
                onReady(function () {
                    try {
                        render(cfg, manifest);
                    } catch (e) {
                        /* fail silent */
                    }
                });
            })
            .catch(function () {
                /* fail silent — no versioned content here */
            });
    } catch (e) {
        /* fail silent */
    }

    function onReady(fn) {
        if (document.readyState === "loading") {
            document.addEventListener("DOMContentLoaded", fn, { once: true });
        } else {
            fn();
        }
    }

    function ensureCss() {
        if (document.querySelector("link[data-vdocs-css]")) return;
        var l = document.createElement("link");
        l.rel = "stylesheet";
        l.href = "/registry/js/versioned-docs.css";
        l.setAttribute("data-vdocs-css", "");
        document.head.appendChild(l);
    }

    // "v6.83.0" -> [6, 83, 0] for descending sort.
    function semverKey(version) {
        return String(version || "")
            .replace(/^v/, "")
            .split(".")
            .map(function (p) {
                var n = parseInt(p, 10);
                return isNaN(n) ? 0 : n;
            });
    }

    function compareVersionsDesc(a, b) {
        var ka = semverKey(a);
        var kb = semverKey(b);
        for (var i = 0; i < Math.max(ka.length, kb.length); i++) {
            var d = (kb[i] || 0) - (ka[i] || 0);
            if (d !== 0) return d;
        }
        return 0;
    }

    function labelFor(entry) {
        if (entry.latest) {
            return entry.version + " (" + String(entry.version).replace(/^v/, "").split(".")[0] + ".x, latest)";
        }
        return entry.version + " (" + entry.slug + ")";
    }

    // Slug of an existing <option>, derived from its target URL:
    // /registry/packages/aws/ -> "", /registry/packages/aws@6.x/ -> "6.x".
    function slugOfOption(opt) {
        var m = /\/registry\/packages\/[^/@]+(?:@([^/]+))?\//.exec(opt.value || "");
        return m && m[1] ? m[1] : "";
    }

    function makeOption(entry, selected) {
        var opt = document.createElement("option");
        opt.value = entry.path;
        opt.textContent = labelFor(entry);
        opt.setAttribute("data-vdocs-added", "");
        if (selected) opt.selected = true;
        return opt;
    }

    function render(cfg, manifest) {
        // A page can carry more than one mount (the top-of-page accordion card
        // and the sidebar card both render the selector partial).
        var mounts = document.querySelectorAll("[data-vdocs-mount]");
        if (!mounts.length) return;
        ensureCss();

        var entries = manifest.versions.slice().sort(function (a, b) {
            return compareVersionsDesc(a.version, b.version);
        });

        Array.prototype.forEach.call(mounts, function (mount) {
            renderMount(cfg, mount, entries);
        });
    }

    function renderMount(cfg, mount, entries) {
        var select = mount.querySelector("select.version-selector");

        if (cfg.mode === "archive") {
            // The baked dropdown is frozen at snapshot time; the manifest is
            // the source of truth.
            if (!select) select = createSelect(mount);
            while (select.firstChild) select.removeChild(select.firstChild);
            entries.forEach(function (entry) {
                select.appendChild(makeOption(entry, entry.slug === cfg.slug));
            });
            mount.hidden = false;
            return;
        }

        // latest mode: baked in-Hugo options win; merge in manifest-only
        // (archived) versions.
        if (!select) {
            select = createSelect(mount);
            entries.forEach(function (entry) {
                select.appendChild(makeOption(entry, entry.slug === cfg.slug));
            });
            mount.hidden = false;
            return;
        }

        var bakedSlugs = {};
        Array.prototype.forEach.call(select.options, function (opt) {
            bakedSlugs[slugOfOption(opt)] = true;
        });

        entries.forEach(function (entry) {
            if (bakedSlugs[entry.slug]) return;
            var added = makeOption(entry, false);
            // Insert in version order (options are sorted newest-first).
            var before = null;
            for (var i = 0; i < select.options.length; i++) {
                var existing = select.options[i];
                var existingVersion = (existing.textContent || "").split(" ")[0];
                if (compareVersionsDesc(entry.version, existingVersion) < 0) {
                    before = existing;
                    break;
                }
            }
            select.insertBefore(added, before);
        });
        mount.hidden = false;
    }

    function createSelect(mount) {
        var select = document.createElement("select");
        select.className = "version-selector";
        select.addEventListener("change", function () {
            if (select.value) window.location.href = select.value;
        });
        mount.appendChild(select);
        return select;
    }
})();
