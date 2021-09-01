import { initDesignSystem, disclosure } from "@pulumi/facet";

initDesignSystem({
    prefix: "pulumi",
    components: [
        disclosure(),
    ],
});
