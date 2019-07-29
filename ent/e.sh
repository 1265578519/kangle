#!/bin/bash
VERSION="3.5.16.3"
if test $# != 1;then
        echo "Usage: $0 dir"
        exit 1;
fi
PREFIX=$1
ARCH="-6"
if test `ldd --version|head -1|awk '{print $NF;}'` = "2.5" ; then
        ARCH="-5"
fi
if test `arch` = "x86_64"; then
        ARCH="$ARCH-x64"
fi
#stop httpd nginx
service httpd stop 2&> /dev/null
service nginx stop 2&> /dev/null
chkconfig --level 2345 httpd off 2&> /dev/null
chkconfig --level 2345 nginx off 2&> /dev/null
URL="http://github.itzmx.com/1265578519/kangle/master/ent/kangle-ent-$VERSION$ARCH.tar.gz"
wget $URL -O kangle.tar.gz
tar xzf kangle.tar.gz
cd kangle
$PREFIX/bin/kangle -q
killall -9 kangle
sleep 3
mkdir -p $PREFIX
wget http://github.itzmx.com/1265578519/kangle/master/ent/license/Ultimate/license.txt -O $PREFIX/license.txt
./install.sh $PREFIX
$PREFIX/bin/kangle
echo "$PREFIX/bin/kangle" >> /etc/rc.d/rc.local
chmod +x /etc/rc.d/rc.local
/sbin/iptables -I INPUT -p tcp --dport 80 -j ACCEPT
/sbin/iptables -I INPUT -p tcp --dport 443 -j ACCEPT
/sbin/iptables -I INPUT -p tcp --dport 3311 -j ACCEPT
/sbin/iptables -I INPUT -p tcp --dport 3312 -j ACCEPT
/sbin/iptables -I INPUT -p tcp --dport 3313 -j ACCEPT
/sbin/iptables -I INPUT -p tcp --dport 21 -j ACCEPT
/etc/rc.d/init.d/iptables save
service ip6tables stop 2&> /dev/null
chkconfig ip6tables off 2&> /dev/null
systemctl stop firewalld 2&> /dev/null
systemctl disable firewalld 2&> /dev/null
rm -rf $PREFIX/www/index.html
wget http://github.itzmx.com/1265578519/kangle/master/easypanel/index.html -O $PREFIX/www/index.html
$PREFIX/bin/kangle -q
killall -9 kangle
sleep 3
$PREFIX/bin/kangle -z /var/cache/kangle
$PREFIX/bin/kangle
