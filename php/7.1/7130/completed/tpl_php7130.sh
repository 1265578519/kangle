#!/bin/sh
yum -y install bzip2-devel libxml2-devel curl-devel db4-devel libjpeg-devel libpng-devel freetype-devel pcre-devel zlib-devel sqlite-devel libmcrypt-devel unzip bzip2
yum -y install mhash-devel openssl-devel
yum -y install libtool-ltdl libtool-ltdl-devel
PREFIX="/vhs/kangle/ext"
wget -c http://github.itzmx.com/1265578519/kangle/master/php/7.1/7130/completed/tpl_php7130.tar.bz2 -O tpl_php7130.tar.bz2
tar xjf tpl_php7130.tar.bz2
mv tpl_php7130 $PREFIX
rm -rf /tmp/*
/vhs/kangle/bin/kangle -r
