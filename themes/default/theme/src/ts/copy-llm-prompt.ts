type MenuElements = {
    root: HTMLElement;
    trigger: HTMLButtonElement;
    dropdown: HTMLElement;
    items: HTMLElement[];
    copyButton: HTMLButtonElement | null;
    prompt: string;
};

const SELECTOR = '[data-component="copy-llm-prompt"]';
const TRIGGER_SELECTOR = '[data-role="trigger"]';
const DROPDOWN_SELECTOR = '[data-role="dropdown"]';
const COPY_SELECTOR = '[data-role="copy"]';

const collectMenus = (): MenuElements[] => {
    return Array.from(document.querySelectorAll<HTMLElement>(SELECTOR))
        .map((root) => {
            const trigger = root.querySelector<HTMLButtonElement>(TRIGGER_SELECTOR);
            const dropdown = root.querySelector<HTMLElement>(DROPDOWN_SELECTOR);
            const template = root.querySelector<HTMLTemplateElement>("template");
            if (!trigger || !dropdown || !template) {
                return null;
            }

            const items = Array.from(
                dropdown.querySelectorAll<HTMLElement>('[role="menuitem"]'),
            );
            const copyButton = dropdown.querySelector<HTMLButtonElement>(COPY_SELECTOR);
            const prompt = template.content.textContent?.trim() ?? "";

            return { root, trigger, dropdown, items, copyButton, prompt };
        })
        .filter((value): value is MenuElements => value !== null);
};

const closeMenu = ({ root, trigger, dropdown }: MenuElements) => {
    root.classList.remove("is-open");
    trigger.setAttribute("aria-expanded", "false");
    dropdown.setAttribute("hidden", "");
};

const openMenu = ({ root, trigger, dropdown, items }: MenuElements) => {
    root.classList.add("is-open");
    trigger.setAttribute("aria-expanded", "true");
    dropdown.removeAttribute("hidden");

    const firstItem = items[0];
    if (document.activeElement === trigger && firstItem) {
        firstItem.focus();
    }
};

const bindCopyButton = (copyButton: HTMLButtonElement, prompt: string) => {
    const label = copyButton.querySelector<HTMLElement>(".copy-llm-prompt__item-label");
    const icon = copyButton.querySelector<HTMLElement>(".copy-llm-prompt__item-icon i");
    const originalLabel = label?.textContent ?? "";

    copyButton.addEventListener("click", async () => {
        await navigator.clipboard.writeText(prompt);

        if (label) label.textContent = "Copied!";
        icon?.classList.replace("fa-copy", "fa-check");

        setTimeout(() => {
            if (label) label.textContent = originalLabel;
            icon?.classList.replace("fa-check", "fa-copy");
        }, 2000);
    });
};

const initCopyPromptButtons = () => {
    const menus = collectMenus();
    if (!menus.length) {
        return;
    }

    document.addEventListener("click", (event) => {
        const target = event.target as Node | null;
        if (!target) return;
        menus.forEach((menu) => {
            if (!menu.root.contains(target)) {
                closeMenu(menu);
            }
        });
    });

    document.addEventListener("keydown", (event) => {
        if (event.key !== "Escape") return;
        menus.forEach((menu) => {
            if (menu.root.classList.contains("is-open")) {
                closeMenu(menu);
                menu.trigger.focus();
            }
        });
    });

    menus.forEach((menu) => {
        menu.trigger.addEventListener("click", (event) => {
            event.preventDefault();
            if (menu.root.classList.contains("is-open")) {
                closeMenu(menu);
            } else {
                openMenu(menu);
            }
        });

        menu.trigger.addEventListener("keydown", (event) => {
            if (event.key === "ArrowDown") {
                event.preventDefault();
                openMenu(menu);
            } else if (event.key === "Escape") {
                closeMenu(menu);
            }
        });

        if (menu.copyButton) {
            bindCopyButton(menu.copyButton, menu.prompt);
        }
    });
};

if (document.readyState === "loading") {
    document.addEventListener("DOMContentLoaded", initCopyPromptButtons);
} else {
    initCopyPromptButtons();
}

export { initCopyPromptButtons };
