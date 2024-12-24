#!/bin/sh
yum -y install bzip2-devel libxml2-devel curl-devel db4-devel libjpeg-devel libpng-devel freetype-devel pcre-devel zlib-devel libmcrypt-devel unzip bzip2
yum -y install mhash-devel openssl-devel
yum -y install libtool-ltdl libtool-ltdl-devel
yum -y remove libzip-devel sqlite-devel
wget -c http://github.itzmx.com/1265578519/kangle/master/php/7.4/7433/libzip-1.3.2.tar.gz -O libzip-1.3.2.tar.gz
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
wget -c http://github.itzmx.com/1265578519/kangle/master/php/7.4/7433/sqlite-autoconf-3470200.tar.gz
tar -zxvf sqlite-autoconf-3470200.tar.gz
cd sqlite-autoconf-3470200
./configure
make -j 4
make install
ln -s /usr/local/lib/pkgconfig/sqlite3.pc /usr/lib64/pkgconfig/sqlite3.pc
echo "/usr/local/lib" > /etc/ld.so.conf.d/sqlite3.conf
ldconfig -v
PREFIX="/vhs/kangle/ext/tpl_php7433"
ZEND_ARCH="i386"
LIB="lib"
if test `arch` = "x86_64"; then
        LIB="lib64"
        ZEND_ARCH="x86_64"
fi

wget -c http://github.itzmx.com/1265578519/kangle/master/php/7.4/7433/php-7.4.33.tar.bz2 -O php-7.4.33.tar.bz2
tar xjf php-7.4.33.tar.bz2
cd php-7.4.33
CONFIG_CMD="./configure --prefix=$PREFIX --with-config-file-scan-dir=$PREFIX/etc/php.d --with-libdir=$LIB --enable-fastcgi --with-mysql --with-mysqli --with-pdo-mysql --with-iconv-dir --with-freetype-dir --with-jpeg-dir --with-png-dir --with-zlib --with-libxml-dir=/usr/include/libxml2/libxml --enable-xml --disable-fileinfo --enable-magic-quotes --enable-safe-mode --enable-bcmath --enable-shmop --enable-sysvsem --enable-inline-optimization --with-curl --with-curlwrappers --disable-mbregex --enable-mbstring --enable-ftp --enable-gd --enable-gd-native-ttf --with-openssl --enable-pcntl --enable-sockets --with-xmlrpc --with-zip --enable-soap --with-pear --with-gettext --enable-calendar --with-openssl"
if [ -f /usr/include/mcrypt.h ]; then
        CONFIG_CMD="$CONFIG_CMD --with-mcrypt"
fi
#'./configure' --prefix=$PREFIX --with-config-file-scan-dir=$PREFIX/etc/php.d --with-libdir=$LIB '--enable-fastcgi' '--with-mysql' '--with-mysqli' --with-pdo-mysql '--with-iconv-dir' '--with-freetype-dir' '--with-jpeg-dir' '--with-png-dir' '--with-zlib' '--with-libxml-dir=/usr/include/libxml2/libxml' '--enable-xml' '--disable-fileinfo' '--enable-magic-quotes' '--enable-safe-mode' '--enable-bcmath' '--enable-shmop' '--enable-sysvsem' '--enable-inline-optimization' '--with-curl' '--with-curlwrappers' '--disable-mbregex' '--enable-mbstring' '--enable-ftp' '--enable-gd' '--enable-gd-native-ttf' '--with-openssl' '--enable-pcntl' '--enable-sockets' '--with-xmlrpc' '--with-zip' '--enable-soap' '--with-pear' '--with-gettext' '--enable-calendar'
#'./configure' --prefix=$PREFIX --with-config-file-scan-dir=$PREFIX/etc/php.d --with-libdir=$LIB '--enable-fastcgi' '--with-mysql' '--with-mysqli' --with-pdo-mysql '--with-iconv-dir' '--with-freetype-dir' '--with-jpeg-dir' '--with-png-dir' '--with-zlib' '--with-libxml-dir=/usr/include/libxml2/libxml' '--enable-xml' '--disable-fileinfo' '--enable-magic-quotes' '--enable-safe-mode' '--enable-bcmath' '--enable-shmop' '--enable-sysvsem' '--enable-inline-optimization' '--with-curl' '--with-curlwrappers' '--disable-mbregex' '--enable-mbstring' '--with-mcrypt' '--enable-ftp' '--enable-gd' '--enable-gd-native-ttf' '--with-openssl' '--with-mhash' '--enable-pcntl' '--enable-sockets' '--with-xmlrpc' '--with-zip' '--enable-soap' '--with-pear' '--with-gettext' '--enable-calendar'
$CONFIG_CMD
if test $? != 0; then
	echo $CONFIG_CMD
	echo "configure php error";
	exit 1
fi
make -j 4
make install
mkdir -p $PREFIX/etc/php.d
if [ ! -f $PREFIX/php-templete.ini ]; then
        cp php.ini-dist $PREFIX/php-templete.ini
fi
if [ ! -f $PREFIX/config.xml ]; then
        wget http://github.itzmx.com/1265578519/kangle/master/php/7.4/7433/config.xml -O $PREFIX/config.xml
fi
cd ..
wget http://github.itzmx.com/1265578519/kangle/master/php/7.4/7433/php-templete.ini -O $PREFIX/php-templete.ini
#install ioncube
wget -c http://github.itzmx.com/1265578519/kangle/master/php/7.4/7433/ioncube-$ZEND_ARCH-7.4.zip
unzip ioncube-$ZEND_ARCH-7.4.zip
mkdir -p $PREFIX/ioncube
mv ioncube_loader_lin_7.4.so $PREFIX/ioncube/ioncube_loader_lin_7.4.so
#install autoconf
wget -c http://github.itzmx.com/1265578519/kangle/master/php/7.4/7433/autoconf-2.69.tar.gz
tar zxf autoconf-2.69.tar.gz
cd autoconf-2.69
./configure
make -j 4
make install
cd ..
#install apcu
wget -c http://github.itzmx.com/1265578519/kangle/master/php/7.4/7433/apcu-5.1.17.tgz
tar zxf apcu-5.1.17.tgz
cd apcu-5.1.17
/vhs/kangle/ext/tpl_php7433/bin/phpize
./configure --with-php-config=/vhs/kangle/ext/tpl_php7433/bin/php-config
make -j 4
make install
cd ..
#install libmemcached
wget -c http://github.itzmx.com/1265578519/kangle/master/php/7.4/7433/libmemcached-1.0.18.tar.gz
tar -zxvf libmemcached-1.0.18.tar.gz
cd libmemcached-1.0.18
./configure
make -j 4
make install
cd ..
#install memcached
wget -c http://github.itzmx.com/1265578519/kangle/master/php/7.4/7433/php-memcached-3.1.3-dev.zip
unzip -o php-memcached-3.1.3-dev.zip
cd php-memcached-3.1.3-dev
/vhs/kangle/ext/tpl_php7433/bin/phpize
./configure --with-php-config=/vhs/kangle/ext/tpl_php7433/bin/php-config --disable-memcached-sasl
make -j 4
make install
cd ..
#install memcache
wget -c http://github.itzmx.com/1265578519/kangle/master/php/7.4/7433/php-memcache-8.2.zip
unzip -o php-memcache-8.2.zip
cd php-memcache-8.2
/vhs/kangle/ext/tpl_php7433/bin/phpize
./configure --with-php-config=/vhs/kangle/ext/tpl_php7433/bin/php-config
make -j 4
make install
cd ..
rm -rf /tmp/*
/vhs/kangle/bin/kangle -r