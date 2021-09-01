const path = require("path");

module.exports = (env, { mode }) => {
    return {
        mode: mode || "production",
        devtool: "inline-source-map",
        entry: {
            index: "./src/main.ts",
        },
        output: {
            filename: `./static/js/local-bundle.min.js`,
            path: process.cwd(),
        },
        resolve: {
            extensions: [".ts", ".js"],
            modules: [
                path.join(process.cwd(), "src"),
                "node_modules"
            ],
        },
        module: {
            rules: [
                {
                    test: /\.ts$/i,
                    use: [
                        {
                            loader: "ts-loader"
                        }
                    ],
                    exclude: /node_modules/
                }
            ]
        }
    }
}
