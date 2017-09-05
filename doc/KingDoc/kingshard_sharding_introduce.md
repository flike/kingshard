# kingshard sharding介绍
现在开源的MySQL Proxy已经有几款了，并且有的已经在生产环境上广泛应用。但这些proxy在sharding方面，都是不能分子表的。也就是说一个node节点只能分一张表。但我们的线上需求通常是这样的：

**我有一张非常大的表，行数超过十亿，需要进行拆分处理。假设拆分因子是512。
如果采用单node单数据库的分表方式，那其实这512个子表还是存在一个物理节点上，意义不大。
	如果采用他们的sharding功能，就需要512个物理节点，也不现实。
	面对这种需求，现有的proxy就不能很好地满足要求了。通常我们希望将512张子表均分在几个MySQL节点上，从而达到系统的横向扩展。**

然而kingshard较好地实现了这种典型的需求。简单来说，kingshard的分表方案采用两级映射的方式：

	1.kingshard将该表分成512张子表，例如：test_0000,test_0001,...
	test_511。
	2.将shardKey通过hash或range方式定位到其要操作的记录在哪张子表上。
	3.子表落在哪个node上通过配置文件设置。

## sharding支持的操作

目前kingshard sharding支持insert, delete, select, update和replace语句, 所有这五类操作都支持跨子表。但写操作仅支持单node上的跨子表，select操作则可以跨node，跨子表。

## sharding方式

### range方式
基于整数范围划分来得到子表下标。该方式的优点：基于范围的查询或更新速度快，因为查询（或更新）的范围有可能落在同一张子表中。这样可以避免全部子表的查询（更新）。缺点：数据热点问题。因为在一段时间内整个集群的写压力都会落在一张子表上。此时整个mysql集群的写能力受限于单台mysql server的性能。并且，当正在集中写的mysql 节点如果宕机的话，整个mysql集群处于不可写状态。基于range方式的分表字段类型受限。

### hash方式
kingshard采用（shardKey%子表个数）的方式得到子表下标。优点：数据分布均匀，写压力会比较平均地落在后端的每个MySQL节点上，整个集群的写性能不会受限于单个MySQL节点。并且当某个分片节点宕机，只会影响到写入该节点的请求，其他节点的写入请求不受影响。分表字段类型不受限。因为任何一个类型的分表字段，都可以通过一个hash函数计算得到一个整数。缺点：基于范围的查询或更新，都需要将请求发送到全部子表，对性能有一定影响。但如果不是基于范围的查询或更新，则性能不会受到影响。

## sharding相关的配置介绍
在配置文件中，有关sharding设置是通过schema设置：

 ```
 schema :
-
    nodes: [node1,node2]
    rules:
        default: node1
        shard:
        -   
            #分表所在的DB
			db : kingshard
			#分表名字
            table: test_shard_hash
            #sharding key
            key: id
            #子表分布的节点名字
            nodes: [node1, node2]
            #sharding类型
            type: hash
            #子表个数分布，表示[test_shard_hash_0000, test_shard_hash_0001, test_shard_hash_0002, test_shard_hash_003]在node1上。
            #[test_shard_hash_0004, test_shard_hash_0005, test_shard_hash_0006, test_shard_hash_007]在node2上
            locations: [4,4]

        -   
            #分表所在的DB
			db : kingshard
			#分表名字
            table: test_shard_range
            #sharding key
            key: id
            #sharding类型
            type: range
            #子表分布的节点名字
            nodes: [node1, node2]
            #子表个数分布，表示[test_shard_range_0000, test_shard_range_0001, test_shard_range_0002, test_shard_range_003]在node1上。
            #[test_shard_range_0004, test_shard_range_0005, test_shard_range_0006, test_shard_range_007]在node2上
            locations: [4,4]
            #每张子表的记录数。[0,10000)在test_shard_range_0000上，[10000,20000)在test_shard_range_0001上。....
            table_row_limit: 10000

 ```
一个kingshard实例只能有一个schemas，从上面的配置可以看出，schema可以分为三个部分：

	1.db，表示这个schemas使用的数据库。

	2.nodes，表示子表分布的节点名字。

	3.rules，sharding规则。其中rules又可以分为两个部分：
		- default，默认分表规则。所有操作不在shard（default规则下面的规则）中的表的SQL语句都会发向该node。
		- hash，hash分表方式。
		- range，range分表方式
## kingshard架构图

![](http://ww3.sinaimg.cn/large/6e5705a5gw1eu7wfrubi3j20qo0k0ab4.jpg)

## 基于kingshard的子表迁移方案
通过kingshard可以非常方便地动态迁移子表，从而保证MySQL节点的不至于负载压力太大。大致步骤如下所述：

1. 通过自动数据迁移工具开始数据迁移。
2. 数据差异小于某一临界值，阻塞老子表写操作（read-only）
3. 等待新子表数据同步完毕
4. 更改kingshard配置文件中的对应子表的路由规则。
4. 删除老节点上的子表。

## Example
简单演示一下kingshard的相关操作，感兴趣的同学可以自己试一试。:)

```
#启动kingshard
kingshard git:(master) ✗ ./bin/kingshard -config=etc/ks.yaml
kingshard
2015/07/19 11:13:43 - INFO - server.go:[205] - [server] "NewServer" "Server running" "netProto=tcp|address=127.0.0.1:9696" conn_id=0

#另一个终端连接kingshard
mysql -ukingshard -pkingshard -h127.0.0.1 -P9696;
Warning: Using a password on the command line interface can be insecure.
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 10001
Server version: kingshard-1.0 Homebrew

Copyright (c) 2000, 2014, Oracle and/or its affiliates. All rights reserved.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql>use kingshard;
Database changed
mysql> select/*master*/ * from kingshard_test_conn;
+----+----------+------+-------+------+------+
| id | str      | f    | e     | u    | i    |
+----+----------+------+-------+------+------+
|  1 | a        | 3.14 | test1 | NULL | NULL |
|  5 | ""''\abc | NULL | NULL  | NULL | NULL |
|  6 | 中国     | NULL | NULL  | NULL | NULL |
+----+----------+------+-------+------+------+
3 rows in set (0.01 sec)

mysql> select * from test_shard_hash where id in(6,10);
+----+-------+------+-------+------+------+
| id | str   | f    | e     | u    | i    |
+----+-------+------+-------+------+------+
| 10 | world |  2.1 | test1 |    1 |    1 |
+----+-------+------+-------+------+------+
1 row in set (0.03 sec)

mysql> show tables;
+----------------------------+
| Tables_in_kingshard        |
+----------------------------+
| kingshard_test_conn        |
| kingshard_test_proxy_conn  |
| kingshard_test_proxy_stmt  |
| kingshard_test_shard_hash  |
| kingshard_test_shard_range |
| kingshard_test_stmt        |
| test_shard_hash_0000       |
| test_shard_hash_0001       |
| test_shard_hash_0002       |
| test_shard_hash_0003       |
| test_shard_range_0000      |
| test_shard_range_0001      |
| test_shard_range_0002      |
| test_shard_range_0003      |
+----------------------------+
14 rows in set (0.00 sec)

```

## 将SQL路由到指定node上

在kingshard中允许用户将特定的sql路由到指定的node上。只需要在sql语句前面加上包含node名称的注释。

```
mysql> /*node2*/show tables;
+-----------------------+
| Tables_in_kingshard   |
+-----------------------+
| kingshard_test_conn   |
| test_shard_hash_0004  |
| test_shard_hash_0005  |
| test_shard_hash_0006  |
| test_shard_hash_0007  |
| test_shard_range_0004 |
| test_shard_range_0005 |
| test_shard_range_0006 |
| test_shard_range_0007 |
+-----------------------+
9 rows in set (0.03 sec)

mysql> /*node2*/select * from kingshard_test_conn;
Empty set (0.01 sec)

mysql> /*node2*/desc kingshard_test_conn;
+-------+-----------------------+------+-----+---------+-------+
| Field | Type                  | Null | Key | Default | Extra |
+-------+-----------------------+------+-----+---------+-------+
| id    | bigint(20) unsigned   | NO   | PRI | NULL    |       |
| str   | varchar(256)          | YES  |     | NULL    |       |
| f     | double                | YES  |     | NULL    |       |
| e     | enum('test1','test2') | YES  |     | NULL    |       |
| u     | tinyint(3) unsigned   | YES  |     | NULL    |       |
| i     | tinyint(4)            | YES  |     | NULL    |       |
+-------+-----------------------+------+-----+---------+-------+
6 rows in set (0.00 sec)

mysql> /*node2*/insert into kingshard_test_conn values(10,"hello",10.2,'test1',1,1);
Query OK, 1 row affected (0.00 sec)

mysql> /*node2*/select * from kingshard_test_conn;
+----+-------+------+-------+------+------+
| id | str   | f    | e     | u    | i    |
+----+-------+------+-------+------+------+
| 10 | hello | 10.2 | test1 |    1 |    1 |
+----+-------+------+-------+------+------+
1 row in set (0.00 sec)

```
