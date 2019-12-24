const path = require('path');
const OptimizeCssAssetsPlugin = require('optimize-css-assets-webpack-plugin');
const UglifyJsPlugin = require('uglifyjs-webpack-plugin');
const merge = require('webpack-merge');
const ManifestPlugin = require('webpack-manifest-plugin');

const common = require('./webpack.common.js');

const prodConfig = {
    mode: 'production',
    output: {
        path: path.resolve(__dirname, "public/static"),
        publicPath: '/static/',
        filename: "[name].[hash].js"
    },
    plugins: [
        new OptimizeCssAssetsPlugin({}),
        new UglifyJsPlugin({
            cache: true, parallel: true, sourceMap: true
        }),
        new ManifestPlugin({basePath: '/static/'})
    ]
};

module.exports = merge(common, prodConfig);