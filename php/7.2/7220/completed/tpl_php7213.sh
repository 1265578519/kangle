#!/bin/sh
yum -y install bzip2-devel libxml2-devel curl-devel db4-devel libjpeg-devel libpng-devel freetype-devel pcre-devel zlib-devel sqlite-devel libmcrypt-devel unzip bzip2
yum -y install mhash-devel openssl-devel
yum -y install libtool-ltdl libtool-ltdl-devel
PREFIX="/vhs/kangle/ext"
wget -c http://github.itzmx.com/1265578519/kangle/master/php/7.2/7213/completed/tpl_php7213.tar.bz2 -O tpl_php7213.tar.bz2
tar xjf tpl_php7213.tar.bz2
mv tpl_php7213 $PREFIX
rm -rf /tmp/*
/vhs/kangle/bin/kangle -r
