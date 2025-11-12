type MenuElements = {
    root: HTMLElement;
    trigger: HTMLButtonElement;
    dropdown: HTMLElement;
    items: HTMLElement[];
};

const SELECTOR = '[data-component="pulumi-ai-menu"]';
const TRIGGER_SELECTOR = '[data-role="trigger"]';
const DROPDOWN_SELECTOR = '[data-role="dropdown"]';

const collectMenus = (): MenuElements[] => {
    return Array.from(document.querySelectorAll<HTMLElement>(SELECTOR))
        .map((root) => {
            const trigger = root.querySelector<HTMLButtonElement>(TRIGGER_SELECTOR);
            const dropdown = root.querySelector<HTMLElement>(DROPDOWN_SELECTOR);
            if (!trigger || !dropdown) {
                return null;
            }

            const items = Array.from(
                dropdown.querySelectorAll<HTMLElement>('[role="menuitem"]'),
            );

            return { root, trigger, dropdown, items };
        })
        .filter((value): value is MenuElements => value !== null);
};

const closeMenu = ({ root, trigger, dropdown }: MenuElements) => {
    root.classList.remove("is-open");
    trigger.setAttribute("aria-expanded", "false");
    if (!dropdown.hasAttribute("hidden")) {
        dropdown.setAttribute("hidden", "");
    }
};

const openMenu = ({ root, trigger, dropdown, items }: MenuElements) => {
    root.classList.add("is-open");
    trigger.setAttribute("aria-expanded", "true");
    dropdown.removeAttribute("hidden");

    // Move focus to the first interactive element when opened via keyboard.
    const firstItem = items[0];
    if (document.activeElement === trigger && firstItem) {
        firstItem.focus();
    }
};

const toggleMenu = (menu: MenuElements, menus: MenuElements[]) => {
    const isOpen = menu.root.classList.contains("is-open");
    menus.forEach((m) => {
        if (m !== menu) {
            closeMenu(m);
        }
    });
    if (isOpen) {
        closeMenu(menu);
    } else {
        openMenu(menu);
    }
};

const handleDocumentClick = (event: MouseEvent, menus: MenuElements[]) => {
    const target = event.target as Node | null;
    if (!target) {
        return;
    }

    menus.forEach((menu) => {
        if (!menu.root.contains(target)) {
            closeMenu(menu);
        }
    });
};

const initMenus = () => {
    const menus = collectMenus();
    if (!menus.length) {
        return;
    }

    menus.forEach((menu) => {
        menu.trigger.addEventListener("click", (event) => {
            event.preventDefault();
            toggleMenu(menu, menus);
        });

        menu.trigger.addEventListener("keydown", (event) => {
            if (event.key === "ArrowDown") {
                event.preventDefault();
                openMenu(menu);
            } else if (event.key === "Escape") {
                closeMenu(menu);
            }
        });

        menu.dropdown.addEventListener("keydown", (event) => {
            if (event.key === "Escape") {
                event.preventDefault();
                closeMenu(menu);
                menu.trigger.focus();
            }
        });
    });

    document.addEventListener("click", (event) => handleDocumentClick(event, menus));
    document.addEventListener("keydown", (event) => {
        if (event.key === "Escape") {
            menus.forEach((menu) => {
                if (menu.root.classList.contains("is-open")) {
                    closeMenu(menu);
                    menu.trigger.focus();
                }
            });
        }
    });
};

if (document.readyState === "loading") {
    document.addEventListener("DOMContentLoaded", initMenus);
} else {
    initMenus();
}

