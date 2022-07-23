背景
========
从技术上懒猫云系统从下到上分为3部分

1. hportal system 一个独立与懒猫云的去中心化虚拟隧道方案，提供家庭网络的可访问性，以及去中心化下终端设备与用户管理体系。

2. 一个精简的linux容器操作系统，以容器为单位运行服务程序。
   此部分主要是提供了
   - 应用程序的运行环境(统一的永久存储，网络IP分配)
   - 客户端的访问环境(统一的ingress入口;自动化的域名、证书机制)
   - 应用程序的生命周期管理(启动、休眠、备份、还原)
   - 统一的应用软件分发、部署机制;
   基于以上机制对传统selfhost应用(如nextcloud,jellfin,drawio)提供了完整的兼容能力

3. 基于2提供的服务，结合1提供的终端设备与用户管理体系建立一套全新的应用开发生态，称作LZCAPP。

LZCAPP
==========
lzcapp运行于家庭云中，界面呈现于各类终端设备中。与传统webapp类似，但也存在很多明显区别，基于这些区别也带来了完全不同的思维模式。

- 自动身份认证

   lzcapp的身份登陆认证由虚拟网络隧道强制完成。 记事本等小工具，不需要考虑任何用户认证等多用户必备的组件。

- 无限IP资源

  每个应用程序可以单独分配一个唯一IP和域名，可以轻松提供非http服务。每个终端设备也可以配备一个唯一IP和域名，来
  实现客户端离线时的服务端PUSH/PULL等操作。比如文件同步、剪贴板共享等可以发生在终端设备没有打开webapp的情况下完成。


- 环境感知
  1. 终端设备信息
     app可以轻松感知到当前登陆UID、登陆设备ID。以及当前UID的其他终端设备。比如应用接力这类功能不需要任何额外机制。
  2. 多用户信息
     app可以轻松建立用户之间的联系。比如在线白板应用，隔空传送等不需要任何帐号体系的维护。
  3. 周边设备感知
     移动硬盘、音箱、摄像头、智能家居硬件等由系统统一接入管理。

- 开发者0运维成本
  实现你的小点子即可。不需要考虑任何部署、带宽、用户量、灾备、分发等传统webapp的琐事/难题。
  这一点也体现在，单个lzcapp实例仅需考虑最多给几十人规模的用户使用

对selfhost的态度
============

目前开源的一些selfhost可以满足[本地优先](https://www.infoq.cn/article/kpiK-qYmJGuzX12ejvaG)需求。

但普遍是传统偏技术性的软件(大概300款)，与普通用户相关的只占少部分(大概50款)。

且这些软件普遍是在传统服务器的背景下设计的，对多设备(非多用户)、移动端访问等。

传统selfhost概念已经过时，仅仅把传统服务端软件部署到本地无法支撑新的生态。需要在新背景下开发新的应用程序来适配当今的用户使用习惯。

懒猫云在设计时应该重点考虑新模式下的开发方式(lzcapp)，在近几年内对传统selfhost提供足够的兼容性即可，长期看仅会有少量几款传统软件保留下来。

单独一款应用软件是解决不了“本地优先”的需求，需要懒猫云这类基础平台的出现后才能一键式解决其中的诸多问题。这也是本产品可能成功的突围点，如果思维仅停留在传统软件上是不可能成功的。


lzcapp文件系统结构
==============

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


用户与文件权限
=======
- 所有帐号均创建相同的UID和GID
- UID:BOX=10000, GID:SYSGROUP=10001(系统APP), GID:APPGROUP=10002(普通APP)
- 系统用户BOX的UID使用<1000的低UID
- 为每个普通用户分配一个1000~9999范围内的UID,称作$UID. (超级管理员UID为1000)
- 为每个运行的APP分配一个20000+的UID,称作$APPUID. 同一个APP内的不同service均使用同一$APPUID
- 默认会将$APPUID加入到APPGROUP组中，以便有权限访问$UID的home数据。(至于是否挂载对应目录，则根据其他条件决定)

- home                                            $UID:APPGROUP    rwxrwx---
- /lzcapp/var; /lzcapp/cache; /lzcapp/hooks;     $APPUID:$APPUID   rwx------
- /tmp                                            BOX:BOX          rwxrwxrwxt
- /lzcapp/run;                                    $APPUID:$APPUID  rwx------



APP配置清单
==========
TODO

serverless服务
===========

- /lzcapp/var/user/$uid/*filepath

  app存放的任意用户数据，需要严格按照UID来存放, serveless基于此规范做自动权限鉴定

  导出到 https://$appdomain/_lzc/files/app/user/*filepath

- /lzcapp/run/mnt/home/$uid/*filepath

  导出到https://$appdomain/_lzc/files/home/*filepath

/lzcapp/var下的具体结构，由serveless功能的完善来逐步建立，与系统底层的存储、备份、快照无关。


hooks
======

在系统或app状态变动时，由系统统一进入对应容器内执行对应事件的脚本，以便辅助完成特定功能。

计划有:
- AppStoping
- AppResumed
- APPStarted

- BeforeBackup
- EndBackup
- CacheCleand
