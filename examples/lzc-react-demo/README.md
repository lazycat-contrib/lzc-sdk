# lzc-sdk react demo

> 该demo用于快速调试`lzc-sdk`相关API

1. 在项目【lzc-sdk】根目录先执行 build.sh 编译`JavaScript`资源
2. 通过反代方式将本地demo项目代理到盒子
  - `cd example/lzc-react-demo`
  - `npm run dev`
  - `lzc-cli project devshell`
  - `lzc-cli app install cloud.lazycat.apps.sdk-v0.0.0.lpk`
3. 打开盒子中的应用或通过链接访问 `https://sdk.[boxName].heiyu.space`
