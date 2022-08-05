LZCAPP运行时
==============

app container的文件系统视角里，除了正常的linux rootfs外，还存在一个特殊的/lzcapp/根目录，

- /lzcapp/run/

  - /lzcapp/run/appdomain  #app分配到的域名全称

  - /lzcapp/run/sys/lzc-api.socket
    LZCAPI的入口文件，文件权限777。平台提供的所有API服务都通过这个socket暴露

  - /lzcapp/run/certs/app.crt,app.key,box.crt
     app证明自身的文件，其中app.key为400，其他为444

  - /lzcapp/run/mnt/
     挂载点：如移动硬盘，用户家目录等

    - /lzcapp/run/mnt/home/$uid/{*Shared,Pictures,Movies,Musics,Download,Documents,...}

    - /lzcapp/run/mnt/media/$disk_id/
     $uid对应的文稿数据，动态挂载。app需要拿用户登陆凭证去请求挂载对应数据。请求通过后会自动挂载到对应目录中。

- /lzcapp/var/                #app存放的任意数据，属于备份数据，仅在程序被卸载时会被删清理掉

- /lzcapp/cache/              #app存放的任意数据，不属于备份数据，在空间不足时可能会被清理掉

- /tmp                        #app存放的任意数据，不属于备份数据，在程序重启时一定会被清理掉

- /lzcapp/pkg/                #只读目录，内容为app pkg解压后的内容. 仅lpk格式类型的app会存在此目录，系统app等没有此目录
 - /lzcapp/pkg/manifest.yml   # lpk:/manifest.yml本身
 - /lzcapp/pkg/icon.png       # lpk:/icon.png
 - /lzcapp/pkg/content/       # lpk:/content.tar解压后的内容

1. 其他目录均为只读文件系统，如需修改，需要自行修改APP配置清单并获取对应权限。
2. 由用户产生或用户应该期望在文件管理器中看到的文稿数据，应该统一放到/lzcapp/run/mnt/home/$uid/目录下。
3. 系统视角不存在公共目录。文件放到/lzcapp/run/mnt/home/$uid/shared目录即表示共享，由系统组件抽取并由UI模拟公共目录效果。
4. /tmp目录不放置到/lzcapp是考虑到大部分软件组件会假设/tmp的存在且由权限写入。


LPK格式
=======
LPK为一个后缀为`.lpk`，以zip压缩方式，压缩了一个目录。其目录结构为

- /manifest.yml  # 对应[lzc app manifest.yml](./manifest.yml)
- /content.tar   # app本身的实际程序以及资源文件等任意文件，一般提供给manifest.yml:routes,icon等字段使用
- /icon.png      # 程序图标，必须有且只支持png格式

一个包签名后还会额外新增以下3个文件

- /sign.sha256   # 除了sbign开头外所有其他顶层目录下普通文件的文件名和对应的sha256值,类似`sha256sum $(find -type f -maxdepth 1)`
- /sign.by.crt   # 签名使用的证书，系统安装和运行时会检测对应证书来源是否安全
- /sign.data     # 使用/sign.by.crt的公钥对/sign.sha256签名后的数据


其他:

- 推荐将icon,vue dist, program bin等都直接嵌入到lpk内部，而非通过docker image的形式传播
- lpk的内容会原封不动的解压到app container内的/lzcapp/pkg/目录下,因此启动参数等可以依赖与这个路径进行配置
- lpk不仅仅可以作为lzcapp的载体格式，将来还可以作为系统插件等组件的载体格式。
