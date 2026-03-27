import { getQueryVariable } from "./util";

document.addEventListener("rendered", function () {
    const el = document.querySelector("#pulumi-ai pulumi-ai");
    if (el) {
        const prompt = getQueryVariable("prompt");
        const language = getQueryVariable("language");
        if (prompt) el.setAttribute("prompt", prompt);
        if (language) el.setAttribute("language", language);
    }
});
