#!/bin/sh
yum -y install bzip2-devel libxml2-devel curl-devel db4-devel libjpeg-devel libpng-devel freetype-devel mysql-devel pcre-devel zlib-devel sqlite-devel libmcrypt-devel unzip
yum -y install mhash-devel openssl-devel
yum -y install libtool-ltdl libtool-ltdl-devel
PREFIX="/vhs/kangle/ext/tpl_php5217"
ZEND_ARCH="i386"
LIB="lib"
if test `arch` = "x86_64"; then
        LIB="lib64"
        ZEND_ARCH="x86_64"
fi

wget -c https://raw.githubusercontent.com/1265578519/kangle/master/php/5.2/5217/php-5.2.17.tar.bz2
tar xjf php-5.2.17.tar.bz2
cd php-5.2.17
CONFIG_CMD="./configure --prefix=$PREFIX --with-config-file-scan-dir=$PREFIX/etc/php.d --with-libdir=$LIB --enable-fastcgi --with-mysql --with-mysqli --with-pdo-mysql --with-iconv-dir --with-freetype-dir --with-jpeg-dir --with-png-dir --with-zlib --with-libxml-dir=/usr/include/libxml2/libxml --enable-xml --disable-fileinfo --enable-magic-quotes --enable-safe-mode --enable-bcmath --enable-shmop --enable-sysvsem --enable-inline-optimization --with-curl --with-curlwrappers --enable-mbregex --enable-mbstring --enable-ftp --with-gd --enable-gd-native-ttf --with-openssl --enable-pcntl --enable-sockets --with-xmlrpc --enable-zip --enable-soap --with-pear --with-gettext --enable-calendar --with-openssl"
if [ -f /usr/include/mcrypt.h ]; then
        CONFIG_CMD="$CONFIG_CMD --with-mcrypt"
fi
#'./configure' --prefix=$PREFIX --with-config-file-scan-dir=$PREFIX/etc/php.d --with-libdir=$LIB '--enable-fastcgi' '--with-mysql' '--with-mysqli' --with-pdo-mysql '--with-iconv-dir' '--with-freetype-dir' '--with-jpeg-dir' '--with-png-dir' '--with-zlib' '--with-libxml-dir=/usr/include/libxml2/libxml' '--enable-xml' '--disable-fileinfo' '--enable-magic-quotes' '--enable-safe-mode' '--enable-bcmath' '--enable-shmop' '--enable-sysvsem' '--enable-inline-optimization' '--with-curl' '--with-curlwrappers' '--enable-mbregex' '--enable-mbstring' '--enable-ftp' '--with-gd' '--enable-gd-native-ttf' '--with-openssl' '--enable-pcntl' '--enable-sockets' '--with-xmlrpc' '--enable-zip' '--enable-soap' '--with-pear' '--with-gettext' '--enable-calendar'
#'./configure' --prefix=$PREFIX --with-config-file-scan-dir=$PREFIX/etc/php.d --with-libdir=$LIB '--enable-fastcgi' '--with-mysql' '--with-mysqli' --with-pdo-mysql '--with-iconv-dir' '--with-freetype-dir' '--with-jpeg-dir' '--with-png-dir' '--with-zlib' '--with-libxml-dir=/usr/include/libxml2/libxml' '--enable-xml' '--disable-fileinfo' '--enable-magic-quotes' '--enable-safe-mode' '--enable-bcmath' '--enable-shmop' '--enable-sysvsem' '--enable-inline-optimization' '--with-curl' '--with-curlwrappers' '--enable-mbregex' '--enable-mbstring' '--with-mcrypt' '--enable-ftp' '--with-gd' '--enable-gd-native-ttf' '--with-openssl' '--with-mhash' '--enable-pcntl' '--enable-sockets' '--with-xmlrpc' '--enable-zip' '--enable-soap' '--with-pear' '--with-gettext' '--enable-calendar'
$CONFIG_CMD
if test $? != 0; then
	echo $CONFIG_CMD
	echo "configure php error";
	exit 1
fi
make
make install
mkdir -p $PREFIX/etc/php.d
if [ ! -f $PREFIX/php-templete.ini ]; then
        cp php.ini-dist $PREFIX/php-templete.ini
fi
if [ ! -f $PREFIX/config.xml ]; then
        wget https://raw.githubusercontent.com/1265578519/kangle/master/php/5.2/5217/config.xml -O $PREFIX/config.xml
fi
cd ..
#install zend
wget -c https://raw.githubusercontent.com/1265578519/kangle/master/php/5.2/5217/zend-3.3.9-$ZEND_ARCH.tar.bz2
tar xjf zend-3.3.9-$ZEND_ARCH.tar.bz2
mv zend-3.3.9-$ZEND_ARCH $PREFIX/zend
wget -c https://raw.githubusercontent.com/1265578519/kangle/master/php/5.2/5217/zend.ini -O $PREFIX/etc/php.d/zend.ini

#install ioncube
wget -c https://raw.githubusercontent.com/1265578519/kangle/master/php/5.2/5217/ioncube-$ZEND_ARCH-5.2.zip
unzip ioncube-$ZEND_ARCH-5.2.zip
mkdir -p $PREFIX/ioncube
mv ioncube_loader_lin_5.2.so $PREFIX/ioncube/ioncube_loader_lin_5.2.so
wget -c https://raw.githubusercontent.com/1265578519/kangle/master/php/5.2/5217/ioncube.ini -O $PREFIX/etc/php.d/ioncube.ini
rm -rf $PREFIX/php-templete.ini
wget https://raw.githubusercontent.com/1265578519/kangle/master/php/5.2/5217/php-templete.ini -O $PREFIX/php-templete.ini
rm -rf /tmp/*
/vhs/kangle/bin/kangle -r
