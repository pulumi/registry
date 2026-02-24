import { Config } from "@stencil/core";
import { sass } from "@stencil/sass";

export const config: Config = {
    namespace: "pulumi-registry",
    enableCache: true,
    buildDist: true,
    outputTargets: [
        {
            type: "dist",
            dir: "./dist",
        },
        {
            type: "www",
            buildDir: "./build",
        },
    ],
    plugins: [sass()],
};
