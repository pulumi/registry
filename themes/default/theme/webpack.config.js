const path = require("path");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const TerserPlugin = require("terser-webpack-plugin");
const WebpackShellPluginNext = require("webpack-shell-plugin-next");
const { WebpackManifestPlugin } = require("webpack-manifest-plugin");

module.exports = function (env, { mode }) {
    return {
        mode: mode || "development",
        entry: {
            bundle: "./src/ts/main.ts",
            marketing: "./src/ts/marketing.ts",
        },
        output: {
            filename: "[name]-registry.[contenthash:8].js",
            chunkFilename: "chunk-[id]-registry.[contenthash:8].js",
            path: path.resolve(__dirname, "../static/registry/js"),
            publicPath: "/registry/js/",
            clean: true,
        },
        resolve: {
            extensions: [".ts", ".js"],
            modules: ["src", "node_modules"],
        },
        devServer: {
            writeToDisk: true,
        },
        module: {
            rules: [
                {
                    test: /\.ts$/i,
                    use: [
                        {
                            loader: "ts-loader",
                            options: {
                                transpileOnly: true,
                            },
                        },
                    ],
                    exclude: /node_modules/,
                },
                {
                    test: /\.s[ac]ss$/i,
                    use: [
                        MiniCssExtractPlugin.loader,
                        {
                            loader: "css-loader",
                            options: {
                                sourceMap: false,
                                url: false,
                            },
                        },
                        {
                            loader: "postcss-loader",
                        },
                        {
                            loader: "sass-loader",
                            options: {
                                sourceMap: false,
                                sassOptions: {
                                    outputStyle: "compressed",
                                },
                            },
                        },
                    ],
                },
            ],
        },
        plugins: [
            new MiniCssExtractPlugin({
                filename: "../css/[name]-registry.[contenthash:8].css",
            }),
            new WebpackManifestPlugin({
                fileName: path.resolve(__dirname, "../data/asset_manifest.json"),
                publicPath: "/registry/js/",
                generate: (seed, files) => {
                    const manifest = {};
                    for (const file of files) {
                        const name = file.name;
                        // Normalize CSS paths: /js/../css/ -> /css/
                        const filePath = file.path.replace("/js/../css/", "/css/");
                        manifest[name] = filePath;
                    }
                    return manifest;
                },
            }),
            new WebpackShellPluginNext({
                onBuildStart: {
                    blocking: true,
                    parallel: false,
                    scripts: ["yarn --cwd stencil run build"],
                },
            }),
        ],
        optimization: {
            minimize: true,
            minimizer: [
                new TerserPlugin({
                    terserOptions: {
                        format: {
                            comments: false,
                        },
                    },
                    extractComments: false,
                }),
            ],
        },
    };
};
