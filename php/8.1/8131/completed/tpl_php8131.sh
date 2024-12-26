#!/bin/sh
yum -y install bzip2-devel libxml2-devel curl-devel db4-devel libjpeg-devel libpng-devel freetype-devel pcre-devel zlib-devel sqlite-devel libmcrypt-devel unzip bzip2
yum -y install mhash-devel openssl-devel
yum -y install libtool-ltdl libtool-ltdl-devel perl-devel perl-core
yum -y remove libzip-devel
wget http://github.itzmx.com/1265578519/kangle/master/php/8.1/8131/openssl-1.1.1k.tar.gz
tar -zxvf openssl-1.1.1k.tar.gz
cd openssl-1.1.1k
./config shared zlib
make -j 4
make test
make install
cd ..
rm -rf /usr/bin/openssl.OFF
mv /usr/bin/openssl /usr/bin/openssl.OFF
ln -s /usr/local/lib64/libssl.so.1.1 /usr/lib64/libssl.so.1.1
ln -s /usr/local/lib64/libcrypto.so.1.1 /usr/lib64/libcrypto.so.1.1
ln -s /usr/local/bin/openssl /usr/bin/openssl
ldconfig -v
yum -y remove openssl-devel
ln -s /usr/local/lib64/pkgconfig/libcrypto.pc /usr/lib64/pkgconfig/libcrypto.pc
ln -s /usr/local/lib64/pkgconfig/libssl.pc /usr/lib64/pkgconfig/libssl.pc
ln -s /usr/local/lib64/pkgconfig/openssl.pc /usr/lib64/pkgconfig/openssl.pc
wget -c http://github.itzmx.com/1265578519/kangle/master/php/8.1/8131/libzip-1.3.2.tar.gz -O libzip-1.3.2.tar.gz
tar xvf libzip-1.3.2.tar.gz
cd libzip-1.3.2
./configure
make -j 4
make install
cd ..
mv -f /usr/lib64/libzip.so.5 /usr/lib64/libzip.so.5.bak
ln -s /usr/local/lib/libzip.so.5 /usr/lib64/libzip.so.5
ln -s /usr/local/lib/pkgconfig/libzip.pc /usr/lib64/pkgconfig/libzip.pc
echo '/usr/local/lib' > /etc/ld.so.conf.d/libzip.conf
ldconfig -v
wget -c http://github.itzmx.com/1265578519/kangle/master/php/8.1/8131/libxml2-2.9.0.tar.gz
tar -zxvf libxml2-2.9.0.tar.gz
cd libxml2-2.9.0
./configure
make -j 4
make install
cd ..
rm -rf /usr/bin/xml2-config.OFF
mv /usr/bin/xml2-config /usr/bin/xml2-config.OFF
ln -s /usr/local/lib/pkgconfig/libxml-2.0.pc /usr/lib64/pkgconfig/libxml-2.0.pc
ln -s /usr/local/bin/xml2-config /usr/bin/xml2-config
ldconfig -v
PREFIX="/vhs/kangle/ext"
wget -c http://github.itzmx.com/1265578519/kangle/master/php/8.1/8131/completed/tpl_php8131.tar.bz2 -O tpl_php8131.tar.bz2
tar xjf tpl_php8131.tar.bz2
mv tpl_php8131 $PREFIX
rm -rf /tmp/*
/vhs/kangle/bin/kangle -r
