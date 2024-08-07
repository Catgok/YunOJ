const CompressionPlugin = require("compression-webpack-plugin")
const {defineConfig} = require('@vue/cli-service')
module.exports = {
    transpileDependencies: true,
    chainWebpack: config => {
        config.plugin('html').tap(args => {
            args[0].title = 'YunOJ' // 设置全局页面标题
            return args
        })
    },
    productionSourceMap: false, //源码
    configureWebpack: {
        plugins: [
            new CompressionPlugin({
                test: /\.(js|css)(\?.*)?$/i,//需要压缩的文件正则
                threshold: 10240000,//文件大小大于这个值时启用压缩
                deleteOriginalAssets: true//压缩后保留原文件
            })
        ],
        resolve: {
            fallback: {
                http: require.resolve("stream-http"),
                url: require.resolve("url/")
            },
        }
        // devtool: "source-map"
    },
    lintOnSave: false,
    devServer: {
        port: 8081
    },
}
