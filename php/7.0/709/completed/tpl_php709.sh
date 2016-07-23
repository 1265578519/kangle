#!/bin/sh
yum -y install bzip2-devel libxml2-devel curl-devel db4-devel libjpeg-devel libpng-devel freetype-devel pcre-devel zlib-devel sqlite-devel libmcrypt-devel unzip bzip2
yum -y install mhash-devel openssl-devel
yum -y install libtool-ltdl libtool-ltdl-devel
wget -c http://github.itzmx.com/1265578519/kangle/master/php/7.0/709/completed/tpl_php709.zip -O tpl_php709.zip
unzip tpl_php709.zip
mv tpl_php709 /vhs/kangle/ext
rm -rf /tmp/*
/vhs/kangle/bin/kangle -r
