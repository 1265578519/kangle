最低配置安装需求
本教程至少需要以下配置进行安装
CPU：单核
内存：256M
硬盘：5G或者以上
网络：独立公网IP
操作系统：centos 6 x64
安装时间：普通VPS大约在2-5分钟左右


默认登录密码是什么?
easypanel控制面板管理员后台：ip:3312/admin
帐号：admin
密码：kangle

mysql数据库管理员后台：ip:3313/mysql
帐号：root
密码：空

想要用mysql的话，一定要先进去3312/admin登录，左边有个服务器设置，把数据库帐号密码填入，并且初始化服务器后即可正常登录mysql使用。


一键安装包是什么?
kangle web server一键安装包是一个用Linux Shell编写的可以为CentOS 6 VPS(VDS)或独立主机安装kangle web server(kangle,easypanel,proftpd,mysql,php,apc,Memcached,safedog)生产环境的Shell程序。


我们为什么需要它?
编译安装需要输入大量的命令，如果是配置生产环境需要耗费大量的时间。
不会Linux的站长或Linux新手想使用Linux作为生产环境……


它有什么优势?
无需一个一个的输入命令，无需值守，编译安装优化编译参数，提高性能，解决不必要的软件间依赖，特别针对VPS用户进行了优化。


如何获取它?
你可以自由 下载 并使用它在VPS(VDS)或独立服务器上，做为真正的生产环境或测试环境。


我们为什么采用kangle这种架构?
采用Linux、PHP、MySQL的优点我们不必多说。
kangle是一个小巧而高效的Linux下的Web服务器软件，是由 kanglesoft.com 站点开发的高并发服务器软件，已经在一些国内的大型网站上运行多年，目前很多国内外的门户网站、行业网站也都在是使用kangle，相当的稳定。
kangle相当的稳定、功能丰富、并发性能强、安装配置简单、低系统资源……

kangle web server 8核心8G内存VPS实测跑60W并发连接数妥妥的




2015-12-31 3.5.5 更新：
修复一个处理上游chunked的bug
删除重试次数设置，将更加智能的方式判断是否进行重试，针对上游新连接不进行重试，长连接则以新连接重试一次。
多节点服务器扩展，可以显示的节点统计数据。
多节点服务器扩展，在连续错误次数设置为0并且所有节点的权重设置为0的情况下，kangle将自动监控所有节点,并自动选择一个最快的节点。
修复websocket无法识别客户端发送Connection: keep-alive, Upgrade头的bug,即keep-alive和Upgrade同时存在于Connection中。

2015-11-26 3.5.4 更新：
支持chunked方式post数据
支持Etag方式缓存

2015-06-18 3.5.1 更新：
支持websocket
增加path_sign，防盗链签名模块。
增加try_file匹配模块
新增http10的标记模块
新增cname绑定模式(详细文档介绍之后介绍)
