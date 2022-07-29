LZC Package
============
- lpk本身是一个zip格式
- 根目录下有一个manifest.yml文件，格式为 [manifest.yml](./manifest.yml)
- 推荐将icon,vue dist, program bin等都直接嵌入到lpk内部，而非通过docker image的形式传播
- lpk的内容会原封不动的解压到app container内的/lzcapp/pkg/目录下。因此启动参数等可以依赖与这个路径进行配置
- lpk不仅仅可以作为lzcapp的载体格式，将来还可以作为系统插件等组件的载体格式。
