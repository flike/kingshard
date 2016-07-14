#!/bin/sh
#
# chkconfig: 2345 64 36
# description: ucdbproxy startup scripts
#
ucdbproxy_root=/uc/sbin
# each config file for one instance
configs="/uc/etc/ks.yaml"
pidfile="/uc/etc/ucdbproxy.pid"
nohupfile="/uc/etc/nohup.ucdbproxy.out"

ulimit -c unlimited
echo 1 > /proc/sys/fs/suid_dumpable
echo  "/uc/share/%e-%p-%s-%t.core" >/proc/sys/kernel/core_pattern

ulimit -n 1024000 
 
start() {
	test -f $pidfile && echo "ucdbproxy is running" 
	nohup $ucdbproxy_root/ucdbproxy -config $configs >> $nohupfile &
	sleep 1
	test -f $pidfile && echo "Start ucdbproxy success"
}
 
stop() {
	kill -9 `cat $pidfile`
	test -f $pidfile && rm $pidfile
	echo "Stop ucdbproxy success";
}
 
# See how we were called.
case "$1" in
    start)
        start
        ;;
    stop)
        stop
        ;;
    restart)
        stop
        start
        ;;
    *)
        echo $"Usage: $0 {start|stop|restart}"
        ;;
esac
exit 0 
