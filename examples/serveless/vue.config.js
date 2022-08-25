module.exports = {
  devServer: {
    historyApiFallback: true,
    disableHostCheck: true,
    //public: "serverless.snyht4.heiyu.space", //盒子里对应app的实际域名
  },
  chainWebpack: config => config.resolve.symlinks(false)
};
