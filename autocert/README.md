# autocert kangle web server自动获取和部署ssl证书工具

#### 介绍
kangle web server的自动创建ssl证书，仅支持kangle 3.5.21以上版本。


#### 安装教程

1.  下载源码,`git clone https://gitee.com/keengo/autocert`
2.  `go build autocert`
3.  编译后生成的autocert放入kangle的bin目录。
4.  如遇到编译报错，需要另外下载go.sum

#### 使用说明

1.  执行权限:
`chmod +x /vhs/kangle/bin/autocert`
2.  增加域名:
`/vhs/kangle/bin/autocert -a domain1,domain2,...`
3.  删除域名
`/vhs/kangle/bin/autocert -d domain1,domain2,...`
4.  列出域名
`/vhs/kangle/bin/autocert -l`
5.  申请成功的证书文件在 /vhs/kangle/etc/ssl 目录中

注：在网站解析的节点中操作申请，申请到的ssl文件下载回本地电脑中，然后记事本打开复制内容填入主控上面即可
我域名解析有100个节点怎么使用这个自动申请SSL：你先把他需要申请得域名进行分组，创建一个新的SSL解析分组，这个SSL分组只解析到一个节点IP中，用于申请SSL，申请完毕后，在把分组解析切回去

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request

