const CompressionWebpackPlugin = require('compression-webpack-plugin');
const productionGzipExtensions = ['js', 'css'];
const isProduction = process.env.NODE_ENV === 'production';

const { defineConfig } = require('@vue/cli-service');

module.exports = defineConfig({
  transpileDependencies: true,
  configureWebpack: config => {
    if (isProduction) {
      // 开启gzip压缩
      config.plugins.push(new CompressionWebpackPlugin({
        algorithm: 'gzip',
        test: new RegExp(`\\.(${productionGzipExtensions.join('|')})$`),
        threshold: 10240,
        minRatio: 0.8
      }));
    }
  },
  chainWebpack: config => {
    config.plugin('html').tap(args => {
      args[0].title = 'yyym的秘密基地';
      return args;
    })
  }
});