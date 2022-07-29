LZCAPP运行时
==============

app container的文件系统视角里，除了正常的linux rootfs外，还存在一个特殊的/lzcapp/根目录，

- /lzcapp/run/
  - /lzcapp/run/sys/lzc-api.socket
    LZCAPI的入口文件，文件权限777。平台提供的所有API服务都通过这个socket暴露
  - /lzcapp/run/app.crt,app.key,box.crt
     app证明自身的文件，其中app.key为400，其他为444
  - /lzcapp/run/mnt/
     挂载点：如移动硬盘，用户家目录等
    - /lzcapp/run/mnt/home/$uid/{*Shared,Pictures,Movies,Musics,Download,Documents,...}
    - /lzcapp/run/mnt/media/$disk_id/
     $uid对应的文稿数据，动态挂载。app需要拿用户登陆凭证去请求挂载对应数据。请求通过后会自动挂载到对应目录中。

- /lzcapp/var/                #app存放的任意数据，属于备份数据，仅在程序被卸载时会被删清理掉

- /lzcapp/cache/              #app存放的任意数据，不属于备份数据，在空间不足时可能会被清理掉

- /lzcapp/hooks/              #主要是为了方便第三方软件的适配，具体见hooks章节

- /tmp                        #app存放的任意数据，不属于备份数据，在程序重启时一定会被清理掉

- /lzcapp/pkg/                #只读目录，内容为app pkg解压后的内容


1. 其他目录均为只读文件系统，如需修改，需要自行修改APP配置清单并获取对应权限。
2. 由用户产生或用户应该期望在文件管理器中看到的文稿数据，应该统一放到/lzcapp/run/mnt/home/$uid/目录下。
3. 系统视角不存在公共目录。文件放到/lzcapp/run/mnt/home/$uid/shared目录即表示共享，由系统组件抽取并由UI模拟公共目录效果。
4. /tmp目录不放置到/lzcapp是考虑到大部分软件组件会假设/tmp的存在且由权限写入。
