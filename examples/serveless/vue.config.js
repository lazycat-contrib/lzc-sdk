module.exports = {
  devServer: {
    historyApiFallback: true,
  },
  chainWebpack: config => config.resolve.symlinks(false)
};
