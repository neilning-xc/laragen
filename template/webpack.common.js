const path = require('path');
const {CleanWebpackPlugin} = require('clean-webpack-plugin');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');

module.exports = {
    entry: {
        app: path.resolve("resources/assets/js/app.js")
    },
    module: {
        rules: [
            {
                test: /\.js$/,
                exclude: /node_modules/,
                use: {
                    loader: "babel-loader"
                }
            },
            {
                test: /\.css$/,
                use: [ 'style-loader', 'css-loader' ]
            },
            {
                test: /\.scss$/,
                use: [
                    MiniCssExtractPlugin.loader,
                    {

                        loader: "css-loader",
                        options: {sourceMap: true}
                    },
                    {
                        loader:"postcss-loader",
                        options: {
                            ident: "postcss",
                            sourceMap: true,
                            plugins: loader => [
                                require('autoprefixer')(),
                                require('postcss-cssnext')()
                            ]
                        }
                    },
                    {
                        loader: "sass-loader",
                        options: {sourceMap: true}
                    }
                ]
            },
            {
                test: /\.(png|svg|jpg|jpeg|gif)$/,
                use: [
                    {loader: "url-loader", options: {limit: 5000}},
                    {
                        loader: "image-webpack-loader",
                        options: {
                            mozjpeg: { progressive: true, quality: 65 },
                            optipng: { enabled: false },
                            pngquant: { quality: '65-90', speed: 4 },
                            gifsicle: { interlaced: false },
                            webp: { quality: 75 }
                        }
                    }
                ]
            },
            {
                test: /\.(woff|woff2|eot|ttf|otf)$/,
                // include: [path.resolve(__dirname, 'src/')],
                use: [ 'file-loader' ]
            }
        ]
    },
    optimization: {
        splitChunks: {
            chunks: "initial",
            cacheGroups: {
                vendor: {
                    test: /[\\/]node_modules[\\/]/,
                    name: "vendors",
                    chunks: "all"
                },
                commons: {
                    name: 'commons',
                    chunks: 'initial',
                    minChunks: 2
                }
            }
        }
    },
    plugins: [
        new CleanWebpackPlugin(),
        new MiniCssExtractPlugin({
            filename: '[name].css',
            chunkFilename: '[id].css'
        })
    ],
    resolve: {
        extensions: ['.js', '.jsx', '.json']
    }
};