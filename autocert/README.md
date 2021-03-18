# autocert kangle web server自动获取和部署ssl证书工具

#### 介绍
kangle web server的自动创建ssl证书，仅支持kangle 3.5.21以上版本。


#### 安装教程

1.  下载源码,`git clone https://gitee.com/keengo/autocert`
2.  `go build autocert`
3.  编译后生成的autocert放入kangle的bin目录。

#### 使用说明

1.  增加域名:
`./autocert -a domain1,domain2,..`
2.  删除域名
`./autocert -d domain1,domain2,...`
3.  列出域名
`./autocert -l`

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request

