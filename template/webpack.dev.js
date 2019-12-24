const path = require('path');
const merge = require('webpack-merge');

const common = require('./webpack.common');

const devConfig = {
    mode: 'development',
    devtool: 'inline-source-map',
    output: {
        path: path.resolve(__dirname, "public/static"),
        publicPath: '/static/',
        filename: "[name].js"
    }
};

module.exports = merge(common, devConfig);
