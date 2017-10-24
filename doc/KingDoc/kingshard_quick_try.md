# KingShard快速入门



## 环境说明

本文仅作为最小实验环境，因此不使用master, slave模式. 单机上使用mysqld_multi运行两个mysql实列


## 初始化数据目录

    # mysql_install_db --datadir=/var/lib/mysql2/ --user=mysql
    # mysql_install_db --datadir=/var/lib/mysql3/ --user=mysql


## 生成配置文件

利用mysqld_multi工具生成配置文件

    # mysqld_multi --example > mysqld_multi.conf

修改根据自己的需求修改mysqld_multi.conf

例：

    [mysqld_multi]
    mysqld     = /usr/bin/mysqld_safe
    mysqladmin = /usr/bin/mysqladmin
    user       = multi_admin
    password   = my_password

    [mysqld2]
    socket     = /var/lib/mysql2/mysql.sock2
    port       = 3307
    pid-file   = /var/lib/mysql2/hostname.pid2
    datadir    = /var/lib/mysql2
    #language   = /usr/share/mysql/english
    user       = unix_user1

    [mysqld3]
    socket     = /var/lib/mysql3/mysql.sock3
    port       = 3308
    pid-file   = /var/lib/mysql3/hostname.pid3
    datadir    = /var/lib/mysql3
    #language   = /usr/share/mysql/swedish
    user       = unix_user2


### 启动多个实例   

    # mysqld_multi --defaults-extra-file=./mysqld_multi.conf start
    或者 mysqld_multi --defaults-extra-file=./mysqld_multi.conf start 2; mysqld_multi --defaults-extra-file=./mysqld_multi.conf start 3(分别启动)


注意这里的2、3对应conf配置文件 mysqld2、mysqld3，以此来区分。


查看实例状态

    [root@testnode kingshard]# mysqld_multi --defaults-extra-file=./mysqld_multi.conf report
    Reporting MySQL servers
    MySQL server from group: mysqld2 is running
    MySQL server from group: mysqld3 is running

说明２个实例都已经启动了。

## 安装Kingshard

参考[kingshard install](https://github.com/doumadou/kingshard/blob/master/doc/KingDoc/kingshard_install_document.md)

## 配置Kingshard

修改/etc/hosts文件, 添加如下二行

    127.0.0.1 node1
    127.0.0.1 node2



配置如下

```

    # server listen addr
    addr : 127.0.0.1:9696

    # server user and password
    user :  kingshard
    password : kingshard
	# the web api server
    web_addr : 0.0.0.0:9797
    #HTTP Basic Auth
    web_user : admin
    web_password : admin

    # log level[debug|info|warn|error],default error
    log_level : debug
    # only allow this ip list ip to connect kingshard
    #allow_ips: 127.0.0.1

    # node is an agenda for real remote mysql server.
    nodes :
    -
        name : node1

        # default max conns for mysql server
        max_conns_limit : 8

        # all mysql in a node must have the same user and password
        user :  root
        password : root

        # master represents a real mysql master server
        master : 127.0.0.1:3307

        # slave represents a real mysql salve server,and the number after '@' is
        #read load weight of this slave.
        #slave : 192.168.0.11:3307@2,192.168.0.12:3307@5
        slave :
        #down_after_noalive : 300
    -
        name : node2

        # default max conns for mysql server
        max_conns_limit : 8

        # all mysql in a node must have the same user and password
        user :  root
        password : root

        # master represents a real mysql master server
        master : 127.0.0.1:3308

        # slave represents a real mysql salve server
        slave :

        # down mysql after N seconds noalive
        # 0 will no down
        down_after_noalive: 100

    # schema defines which db can be used by client and this db's sql will be executed in which nodes
    schema :
        nodes: [node1,node2]
		default: node1
        shard:
        -   
            db : kingshard
            table: test_shard_hash
            key: id
            nodes: [node1, node2]
            type: hash
            locations: [4,4]
        -   
            db : kingshard
            table: test_shard_range
            key: id
            type: range
            nodes: [node1, node2]
            locations: [4,4]
            table_row_limit: 10000

```

## 设置mysql实例信息

### 设置用户
分类登陆mysqld2, mysqld3, 创建root用户(该用户是给kingshard管理的，测试为了方便所以直接使用root)
若用户存在，跳过此步

    /usr/bin/mysqladmin -h 127.0.0.1 -P 3307 -u root password 'root'
    /usr/bin/mysqladmin -h 127.0.0.1 -P 3308 -u root password 'root'

### 建数据库
    分类登陆mysqld2, mysqld2，创建kingshard数据库
    /usr/bin/mysql -h 127.0.0.1 -P 3307 -u root -proot -e "create database kingshard;"
    /usr/bin/mysql -h 127.0.0.1 -P 3308 -u root -proot -e "create database kingshard;"


## 启动Kingshard

    # ./bin/kingshard -config=etc/ks.yaml

## 测试shard功能

使用test_shard_hash测试 shard　hash分表功能.

### 创建分表

创建test_shard_hash分表(_0000~_0007), _0001~_0003在node1(mysqld2)上创建, _0004~_0007在node2(mysqld3)上创建。

    for i in `seq 0 3`;do /usr/bin/mysql -h 127.0.0.1 -P 3307 -u root -proot kingshard -e "CREATE TABLE IF NOT EXISTS test_shard_hash_000"${i}" ( id BIGINT(64) UNSIGNED  NOT NULL, str VARCHAR(256), f DOUBLE, e enum('test1', 'test2', 'test3', 'test4', 'test5', 'test6', 'test7', 'test8', 'test9', 'test10'), u tinyint unsigned, i tinyint, ni tinyint, PRIMARY KEY (id)) ENGINE=InnoDB DEFAULT CHARSET=utf8;";done
    for i in `seq 4 7`;do /usr/bin/mysql -h 127.0.0.1 -P 3308 -u root -proot kingshard -e "CREATE TABLE IF NOT EXISTS test_shard_hash_000"${i}" ( id BIGINT(64) UNSIGNED  NOT NULL, str VARCHAR(256), f DOUBLE, e enum('test1', 'test2', 'test3', 'test4', 'test5', 'test6', 'test7', 'test8', 'test9', 'test10'), u tinyint unsigned, i tinyint, ni tinyint, PRIMARY KEY (id)) ENGINE=InnoDB DEFAULT CHARSET=utf8;";done

### 插入数据

mysql连接到kingshard插入数据

    for i in `seq 1 10`;do mysql -h 127.0.0.1 -P 9696 -u kingshard -pkingshard -e "insert into test_shard_hash (id, str, f, e, u, i) values(${i}, 'abc$i', 3.14, 'test$i', 255, -127)";done

kingshard日志如下：

    2015/07/29 07:39:15 - INFO - 127.0.0.1:40135->127.0.0.1:3307:select @@version_comment limit 1
    2015/07/29 07:39:15 - INFO - 127.0.0.1:40135->127.0.0.1:3307:insert into test_shard_hash_0001(id, str, f, e, u, i) values (1, 'abc1', 3.14, 'test1', 255, -127)
    2015/07/29 07:39:15 - INFO - 127.0.0.1:40136->127.0.0.1:3307:select @@version_comment limit 1
    2015/07/29 07:39:15 - INFO - 127.0.0.1:40136->127.0.0.1:3307:insert into test_shard_hash_0002(id, str, f, e, u, i) values (2, 'abc2', 3.14, 'test2', 255, -127)
    2015/07/29 07:39:15 - INFO - 127.0.0.1:40137->127.0.0.1:3307:select @@version_comment limit 1
    2015/07/29 07:39:15 - INFO - 127.0.0.1:40137->127.0.0.1:3307:insert into test_shard_hash_0003(id, str, f, e, u, i) values (3, 'abc3', 3.14, 'test3', 255, -127)
    2015/07/29 07:39:15 - INFO - 127.0.0.1:40138->127.0.0.1:3307:select @@version_comment limit 1
    2015/07/29 07:39:15 - INFO - 127.0.0.1:40138->127.0.0.1:3308:insert into test_shard_hash_0004(id, str, f, e, u, i) values (4, 'abc4', 3.14, 'test4', 255, -127)
    2015/07/29 07:39:15 - INFO - 127.0.0.1:40139->127.0.0.1:3307:select @@version_comment limit 1
    2015/07/29 07:39:15 - INFO - 127.0.0.1:40139->127.0.0.1:3308:insert into test_shard_hash_0005(id, str, f, e, u, i) values (5, 'abc5', 3.14, 'test5', 255, -127)
    2015/07/29 07:39:15 - INFO - 127.0.0.1:40140->127.0.0.1:3307:select @@version_comment limit 1
    2015/07/29 07:39:15 - INFO - 127.0.0.1:40140->127.0.0.1:3308:insert into test_shard_hash_0006(id, str, f, e, u, i) values (6, 'abc6', 3.14, 'test6', 255, -127)
    2015/07/29 07:39:15 - INFO - 127.0.0.1:40141->127.0.0.1:3307:select @@version_comment limit 1
    2015/07/29 07:39:15 - INFO - 127.0.0.1:40141->127.0.0.1:3308:insert into test_shard_hash_0007(id, str, f, e, u, i) values (7, 'abc7', 3.14, 'test7', 255, -127)
    2015/07/29 07:39:15 - INFO - 127.0.0.1:40142->127.0.0.1:3307:select @@version_comment limit 1
    2015/07/29 07:39:15 - INFO - 127.0.0.1:40142->127.0.0.1:3307:insert into test_shard_hash_0000(id, str, f, e, u, i) values (8, 'abc8', 3.14, 'test8', 255, -127)
    2015/07/29 07:39:15 - INFO - 127.0.0.1:40143->127.0.0.1:3307:select @@version_comment limit 1
    2015/07/29 07:39:15 - INFO - 127.0.0.1:40143->127.0.0.1:3307:insert into test_shard_hash_0001(id, str, f, e, u, i) values (9, 'abc9', 3.14, 'test9', 255, -127)
    2015/07/29 07:39:15 - INFO - 127.0.0.1:40144->127.0.0.1:3307:select @@version_comment limit 1
    2015/07/29 07:39:15 - INFO - 127.0.0.1:40144->127.0.0.1:3307:insert into test_shard_hash_0002(id, str, f, e, u, i) values (10, 'abc10', 3.14, 'test10', 255, -127)


通过kingshard的日志可以看到数据插入时根据不同的hash值，插入到不同的子表里面去了。

### 查看数据

    [root@testnode kingshard]# mysql -h 127.0.0.1 -P 9696 -u kingshard -pkingshard -e "select * from test_shard_hash where id in (2, 3, 4, 5)"
    +----+------+------+-------+------+------+------+
    | id | str  | f    | e     | u    | i    | ni   |
    +----+------+------+-------+------+------+------+
    |  2 | abc2 | 3.14 | test2 |  255 | -127 | NULL |
    |  3 | abc3 | 3.14 |       |  255 | -127 | NULL |
    |  4 | abc4 | 3.14 |       |  255 | -127 | NULL |
    |  5 | abc5 | 3.14 |       |  255 | -127 | NULL |
    +----+------+------+-------+------+------+------+
