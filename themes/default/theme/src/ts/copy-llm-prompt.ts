const SELECTOR = '[data-component="copy-llm-prompt"]';

const initCopyPromptButtons = () => {
    document.querySelectorAll<HTMLElement>(SELECTOR).forEach((root) => {
        if (root.hasAttribute("data-initialized")) {
            return;
        }
        root.setAttribute("data-initialized", "");

        const template = root.querySelector<HTMLTemplateElement>("template");
        const button = root.querySelector<HTMLButtonElement>("button");
        if (!template || !button) {
            return;
        }

        const prompt = template.content.textContent?.trim() ?? "";
        const textEl = button.querySelector<HTMLElement>(".copy-llm-prompt__text");
        const iconEl = button.querySelector<HTMLElement>(".copy-llm-prompt__icon");
        const originalText = textEl?.textContent ?? "";

        button.addEventListener("click", async () => {
            await navigator.clipboard.writeText(prompt);

            button.classList.add("is-copied");
            if (textEl) {
                textEl.textContent = "Copied!";
            }
            if (iconEl) {
                iconEl.classList.remove("fa-copy");
                iconEl.classList.add("fa-check");
            }

            setTimeout(() => {
                button.classList.remove("is-copied");
                if (textEl) {
                    textEl.textContent = originalText;
                }
                if (iconEl) {
                    iconEl.classList.remove("fa-check");
                    iconEl.classList.add("fa-copy");
                }
            }, 2000);
        });
    });
};

if (document.readyState === "loading") {
    document.addEventListener("DOMContentLoaded", initCopyPromptButtons);
} else {
    initCopyPromptButtons();
}

export { initCopyPromptButtons };
