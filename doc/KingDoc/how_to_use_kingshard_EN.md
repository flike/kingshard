# How to use kingshard building a MySQL cluster

## 1. The application scenario of kingshard

Now more and more Internet companies still in heavy use MySQL to store various types of relational data. With the amount of data and traffic increasing, developers had to consider some new MySQL-related problems.

1. Read/Write Splitting. With the increasing traffic sent by the front-end applications, one instance of MySQL can not hold all the queries. At this time, we have to send the read queries to the slaves for load balance.
2. The capacity of one table in MySQL. If in the begining of system design, you have not considered the table sharding, that will make it difficult to keep your system high-performance.
3. MySQL maintenance operation. Without proxy you should configure the master and slaves host in your source code. When you upgrade the MySQL server, the front-end applications have to make relevant regulation.
4. Connection pool. The front-end applications send queries by creating a new connection with MySQL, and close the connection when they don't need to send queries any more. The extra performance cost of these operations can not be ignored. If a connection pool is added between the front-end applications and MySQL, and the front-end applications can pick a connection from the connection pool, it will enhance the performance of your system.
5. SQL logs. When the program has problems, usually we want to get some SQL logs sent by the program. For example, We want to know which SQL was sent to which DB backend. By checking the log, it can help us locate the problem more quickly.

Faced with these problems, we can implement every function in the client code. But this also makes the client less flexible. I have been working on database development for years, and I believe we can use a MySQL proxy to solve the problems more effectively, which is why I created this project. In this document, I will show you how kingshard solve the above problems. 

## 2. Install kingshard
### (1). set the config file

kingshard run with a configuration file (ks.yaml). Before running kingshard, the file needs to be configured. Here I give a configure file as a demo, we only need to modify some configuration options inside, not need to rewrite a configuration file from scratch.

```
# server listen addr
addr : 0.0.0.0:9696
# the web api server
web_addr : 0.0.0.0:9797
#HTTP Basic Auth
web_user : admin
web_password : admin

# user list with user name and password
user_list:
-
user :  kingshard
password : kingshard

#if set log_path, the sql log will write into log_path/sql.log,the system log
#will write into log_path/sys.log
#log_path : /Users/flike/log

# log level[debug|info|warn|error],default error
log_level : debug
#if set log_sql(on|off) off,the sql log will not output
#log_sql: off 
#only log the query that take more than slow_log_time ms
#slow_log_time : 100
# the path of blacklist sql file
# all these sqls in the file will been forbidden by kingshard
#blacklist_sql_file: /Users/flike/blacklist
# only allow this ip list ip to connect kingshard
#allow_ips: 127.0.0.1
# the default charset of kingshard is utf8.
#proxy_charset: utf8mb4

# node is an agenda for real remote mysql server.
nodes :
- 
    name : node1 

    # default max conns for mysql server
    max_conns_limit : 8

    # all mysql in a node must have the same user and password
    user :  kingshard 
    password : kingshard

    # master represents a real mysql master server 
    master : 127.0.0.1:3306

    # slave represents a real mysql slave server,and the number after '@' is 
    # read load weight of this slave.
    slave : 
    down_after_noalive : 32
- 
    name : node2 

    # default max conns for mysql server
    max_conns_limit : 8

    # all mysql in a node must have the same user and password
    user :  kingshard 
    password : kingshard

    # master represents a real mysql master server 
    master : 192.168.59.103:3307

    # slave represents a real mysql slave server 
    slave : 

    # down mysql after N seconds noalive
    # 0 will no down
    down_after_noalive: 32

# schema list include all user's schema
# schema defines which db can be used by client and this db's sql will be executed in which nodes
schema_list :
-
    user: kingshard
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

**Sharding Notes:**

* kingshard supports two sharding type: range and hash.
* the sub-table needs to be created in the right database manually. The format is `table_%4d`. In other words, the index of the sub-table is a integer made by four digits. Such as `table_name_0000,table_name_0012`.
* All SQLs that handle the unshading table will be sent to the default node.

### (2). Install And Run

```
1. Install Go
2. git clone https://github.com/flike/kingshard.git src/github.com/flike/kingshard
3. cd src/github.com/flike/kingshard
4. source ./dev.sh
5. make
6. set the config file (etc/ks.yaml)
7. run kingshard (./bin/kingshard -config=etc/multi.yaml)
```

## 3. Sharding

I build a mysql cluster with kingshard, and the topology is shown below.
![topology](./kingshard_access_node_arch.jpg)

### 3.1 The hands-on Examples of sharding
### 3.1.1 Create sub table manually

I create eight sub tables in node1 and node2, each node have four sub-tables. `test_shard_hash_0000, test_shard_hash_0001, test_shard_hash_0002, test_shard_hash_0003` in node1, and `test_shard_hash_0004, test_shard_hash_0005, test_shard_hash_0006, test_shard_hash_0007` in node2. The create table sql as below:

```
CREATE TABLE `test_shard_hash_0000` (
  `id` bigint(64) unsigned NOT NULL,
  `str` varchar(256) DEFAULT NULL,
  `f` double DEFAULT NULL,
  `e` enum('test1','test2') DEFAULT NULL,
  `u` tinyint(3) unsigned DEFAULT NULL,
  `i` tinyint(4) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
```

### 3.1.2 Insert and select operation of sharding

The select SQL queries will be sent to a proper database or multiple databases based on the conditions. The insert SQLs will also be sent to multiple databases, if the insert operation is across multiple databases, kingshard will send the queries to multiple databases. The queries as below:

```
mysql> insert into test_shard_hash(id,str,f,e,u,i) values(15,"flike",3.14,'test2',2,3);
Query OK, 1 row affected (0.01 sec)

mysql> mysql> insert into test_shard_hash(id,str,f,e,u,i) values(7,"chen",2.1,'test1',32,3);
Query OK, 1 row affected (0.01 sec)

mysql> insert into test_shard_hash(id,str,f,e,u,i) values(17,"github",2.5,'test1',32,3);
Query OK, 1 row affected (0.00 sec)

mysql> insert into test_shard_hash(id,str,f,e,u,i) values(18,"kingshard",7.3,'test1',32,3);
Query OK, 1 row affected (0.01 sec)

``` 
And the corresponding log below:

```
2015/09/02 18:48:24 - INFO - 127.0.0.1:55003->192.168.59.103:3307:insert into test_shard_hash_0007(id, str, f, e, u, i) values (15, 'flike', 3.14, 'test2', 2, 3)
2015/09/02 18:49:05 - INFO - 127.0.0.1:55003->192.168.59.103:3307:insert into test_shard_hash_0007(id, str, f, e, u, i) values (7, 'chen', 2.1, 'test1', 32, 3)
2015/09/02 18:49:51 - INFO - 127.0.0.1:55003->127.0.0.1:3306:insert into test_shard_hash_0001(id, str, f, e, u, i) values (17, 'github', 2.5, 'test1', 32, 3)
2015/09/02 18:50:21 - INFO - 127.0.0.1:55003->127.0.0.1:3306:insert into test_shard_hash_0002(id, str, f, e, u, i) values (18, 'kingshard', 7.3, 'test1', 32, 3)
```

Notice that the first two queries have been sent to the master in node2, and the last two SQLs have been sent to the master in node1.

Then we send the select queries to get the records-kingshard supports select operation across nodes. The select queries are below:

```
mysql> select * from test_shard_hash where id < 18;
+----+--------+------+-------+------+------+
| id | str    | f    | e     | u    | i    |
+----+--------+------+-------+------+------+
| 17 | github |  2.5 | test1 |   32 |    3 |
|  7 | chen   |  2.1 | test1 |   32 |    3 |
| 15 | flike  | 3.14 | test2 |    2 |    3 |
+----+--------+------+-------+------+------+
3 rows in set (0.02 sec)
```

As the sharding type is hash, the select operation will query all databases. And the corresponding SQLs log as below:

```
2015/09/02 18:55:01 - INFO - 127.0.0.1:55003->127.0.0.1:3306:select * from test_shard_hash_0000 where id < 18
2015/09/02 18:55:01 - INFO - 127.0.0.1:55003->127.0.0.1:3306:select * from test_shard_hash_0001 where id < 18
2015/09/02 18:55:01 - INFO - 127.0.0.1:55003->127.0.0.1:3306:select * from test_shard_hash_0002 where id < 18
2015/09/02 18:55:01 - INFO - 127.0.0.1:55003->127.0.0.1:3306:select * from test_shard_hash_0003 where id < 18
2015/09/02 18:55:01 - INFO - 127.0.0.1:55003->192.168.59.103:3307:select * from test_shard_hash_0004 where id < 18
2015/09/02 18:55:01 - INFO - 127.0.0.1:55003->192.168.59.103:3307:select * from test_shard_hash_0005 where id < 18
2015/09/02 18:55:01 - INFO - 127.0.0.1:55003->192.168.59.103:3307:select * from test_shard_hash_0006 where id < 18
2015/09/02 18:55:01 - INFO - 127.0.0.1:55003->192.168.59.103:3307:select * from test_shard_hash_0007 where id < 18
```

If the query criteria of select SQLs is equal, kingshard will calculate the index of sub-table based on a condition. For example:

```
mysql> select * from test_shard_hash where id = 18;
+----+-----------+------+-------+------+------+
| id | str       | f    | e     | u    | i    |
+----+-----------+------+-------+------+------+
| 18 | kingshard |  7.3 | test1 |   32 |    3 |
+----+-----------+------+-------+------+------+
1 row in set (0.00 sec)
```
And the SQL log is:

```
2015/09/02 18:59:37 - INFO - 127.0.0.1:55003->127.0.0.1:3306:select * from test_shard_hash_0002 where id = 18
```

### 3.1.3 Update operation of sharding

The update SQLs will only be sent to one database, if the update operation accoss multi databases, kingshard will response error messages. 

Kingshard supports the update operation accoss sub tables that in one node. For example as below:

```
mysql> update test_shard_hash set i=23 where id = 17 or id = 18;
Query OK, 2 rows affected (0.00 sec)
```

The corresponding SQL log is below:

```
2015/09/02 19:24:46 - INFO - 127.0.0.1:55003->127.0.0.1:3306:update test_shard_hash_0001 set i = 23 where id = 17 or id = 18
2015/09/02 19:24:46 - INFO - 127.0.0.1:55003->127.0.0.1:3306:update test_shard_hash_0002 set i = 23 where id = 17 or id = 18
```
If the records that need be updated are located in multiple databases, kingshard will response with a error message. For example as below:

```
mysql> update test_shard_hash set i=23 where id = 15 or id = 18;
ERROR 1105 (HY000): no route node
```
The corresponding SQLs log is below:

```
2015/09/02 19:24:24 - ERROR - router.go:[483] - [Router] "generateUpdateSql" "update in multi node" "RouteNodeIndexs=[0 1]" conn_id=0
```

## 3.2 Send query to the specified node

Sometimes the table we want to query is not in the default node. We can use the feature provided by kingshard to solve this problem. You only need add comment to specify the node ahead the query. For example as below :

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

## 3.3 Send the reading query to master

Kingshard can split reading/writing queries, but sometimes we may want send the reading query to master, but don't use transaction. You can add a comment(`/*master*/`) in select SQL, like `select /*master*/ * from stu;`,and this sql will send to the master.When you use mysql client to test this function, you need use the parameter:'-c' to connect the mysql server in order to retain the comment. For example as below:

```
mysql> select/*master*/ * from kingshard_test_conn;
+----+----------+------+-------+------+------+
| id | str      | f    | e     | u    | i    |
+----+----------+------+-------+------+------+
|  1 | a        | 3.14 | test1 | NULL | NULL |
|  5 | ""''\abc | NULL | NULL  | NULL | NULL |
|  6 | 中国     | NULL | NULL  | NULL | NULL |
+----+----------+------+-------+------+------+
3 rows in set (0.01 sec)
```

## 3.4 Function support

Kingshard also support most commonly used functions, such as `max, min, count, sum`, and also support `order by`. For example as below:

```
mysql> select count(id) from test_shard_hash where id > 1;
+-----------+
| count(id) |
+-----------+
|         4 |
+-----------+
1 row in set (0.02 sec)

mysql> select sum(id) from test_shard_hash where id > 1;
+---------+
| sum(id) |
+---------+
|      57 |
+---------+
1 row in set (0.02 sec)
```

The corresponding SQLs log as below:

```
2015/09/03 14:49:01 - INFO - 127.0.0.1:55768->127.0.0.1:3306:select count(id) from test_shard_hash_0000 where id > 1
2015/09/03 14:49:01 - INFO - 127.0.0.1:55768->127.0.0.1:3306:select count(id) from test_shard_hash_0001 where id > 1
2015/09/03 14:49:01 - INFO - 127.0.0.1:55768->127.0.0.1:3306:select count(id) from test_shard_hash_0002 where id > 1
2015/09/03 14:49:01 - INFO - 127.0.0.1:55768->127.0.0.1:3306:select count(id) from test_shard_hash_0003 where id > 1
2015/09/03 14:49:01 - INFO - 127.0.0.1:55768->192.168.59.103:3307:select count(id) from test_shard_hash_0004 where id > 1
2015/09/03 14:49:01 - INFO - 127.0.0.1:55768->192.168.59.103:3307:select count(id) from test_shard_hash_0005 where id > 1
2015/09/03 14:49:01 - INFO - 127.0.0.1:55768->192.168.59.103:3307:select count(id) from test_shard_hash_0006 where id > 1
2015/09/03 14:49:01 - INFO - 127.0.0.1:55768->192.168.59.103:3307:select count(id) from test_shard_hash_0007 where id > 1
2015/09/03 14:49:14 - INFO - 127.0.0.1:55768->127.0.0.1:3306:select sum(id) from test_shard_hash_0000 where id > 1
2015/09/03 14:49:14 - INFO - 127.0.0.1:55768->127.0.0.1:3306:select sum(id) from test_shard_hash_0001 where id > 1
2015/09/03 14:49:14 - INFO - 127.0.0.1:55768->127.0.0.1:3306:select sum(id) from test_shard_hash_0002 where id > 1
2015/09/03 14:49:14 - INFO - 127.0.0.1:55768->127.0.0.1:3306:select sum(id) from test_shard_hash_0003 where id > 1
2015/09/03 14:49:14 - INFO - 127.0.0.1:55768->192.168.59.103:3307:select sum(id) from test_shard_hash_0004 where id > 1
2015/09/03 14:49:14 - INFO - 127.0.0.1:55768->192.168.59.103:3307:select sum(id) from test_shard_hash_0005 where id > 1
2015/09/03 14:49:14 - INFO - 127.0.0.1:55768->192.168.59.103:3307:select sum(id) from test_shard_hash_0006 where id > 1
2015/09/03 14:49:14 - INFO - 127.0.0.1:55768->192.168.59.103:3307:select sum(id) from test_shard_hash_0007 where id > 1
```

```
mysql> select * from test_shard_hash where id > 1 order by id;
+----+-----------+------+-------+------+------+
| id | str       | f    | e     | u    | i    |
+----+-----------+------+-------+------+------+
|  7 | chen      |  2.1 | test1 |  123 |    3 |
| 15 | flike     | 3.14 | test2 |  123 |    3 |
| 17 | github    |  2.5 | test1 |   32 |   23 |
| 18 | kingshard |  7.3 | test1 |   32 |   23 |
+----+-----------+------+-------+------+------+
4 rows in set (0.02 sec)
```

The corresponding SQLs log as below:

```
2015/09/03 14:54:11 - INFO - 127.0.0.1:55768->127.0.0.1:3306:select * from test_shard_hash_0000 where id > 1 order by id asc
2015/09/03 14:54:11 - INFO - 127.0.0.1:55768->127.0.0.1:3306:select * from test_shard_hash_0001 where id > 1 order by id asc
2015/09/03 14:54:11 - INFO - 127.0.0.1:55768->127.0.0.1:3306:select * from test_shard_hash_0002 where id > 1 order by id asc
2015/09/03 14:54:11 - INFO - 127.0.0.1:55768->127.0.0.1:3306:select * from test_shard_hash_0003 where id > 1 order by id asc
2015/09/03 14:54:11 - INFO - 127.0.0.1:55768->192.168.59.103:3307:select * from test_shard_hash_0004 where id > 1 order by id asc
2015/09/03 14:54:11 - INFO - 127.0.0.1:55768->192.168.59.103:3307:select * from test_shard_hash_0005 where id > 1 order by id asc
2015/09/03 14:54:11 - INFO - 127.0.0.1:55768->192.168.59.103:3307:select * from test_shard_hash_0006 where id > 1 order by id asc
2015/09/03 14:54:11 - INFO - 127.0.0.1:55768->192.168.59.103:3307:select * from test_shard_hash_0007 where id > 1 order by id asc
```

## 4. Transaction in one node

kingshard support executing a transaction only in one database. If the transaction accoss multi databases, kingshard will response error messages. For example as below:

```
mysql> begin;
Query OK, 0 rows affected (0.00 sec)

mysql> insert into test_shard_hash(id,str,f,e,u,i) values(23,'proxy',9.2,'test1',12,3);
Query OK, 1 row affected (0.00 sec)

mysql> commit;
Query OK, 0 rows affected (0.01 sec)
```

## 5. The admin command of kingshard

The admin command of kingshard show as below:

### 5.1 Database operation

```
#add a slave in node1
admin node(opt,node,k,v) values(‘add’,’node1’,’slave’,’127.0.0.1:3306’)

#delete slave in node1, ps:master can't been remove.
admin node(opt,node,k,v) values(‘del’,’node1’,’slave’,’127.0.0.1:3306’)

#set slave down
admin node(opt,node,k,v) values(‘down’,’node1’,’slave’,’127.0.0.1:3306’)

#set slave up
admin node(opt,node,k,v) values(‘up’,’node1’,’slave’,’127.0.0.1:3306’)

#set master down
admin node(opt,node,k,v) values(‘down’,’node1’,’master’,’127.0.0.1:3306’)

#set master up
admin node(opt,node,k,v) values(‘up’,’node1’,’master’,’127.0.0.1:3306’)
```

## 5.2 View the status of kingshard

```
#view the config of kingshard
mysql> admin server(opt,k,v) values('show','proxy','config');
+-------------+----------------+
| Key         | Value          |
+-------------+----------------+
| Addr        | 127.0.0.1:9696 |
| User        | kingshard      |
| Password    | kingshard      |
| LogLevel    | debug          |
| Nodes_Count | 2              |
| Nodes_List  | node1,node2    |
+-------------+----------------+
6 rows in set (0.00 sec)

#view the status of node
mysql> admin server(opt,k,v) values('show','node','config');
+-------+---------------------+--------+-------+-------------------------------+-------------+----------+
| Node  | Address             | Type   | State | LastPing                      | MaxIdleConn | IdleConn |
+-------+---------------------+--------+-------+-------------------------------+-------------+----------+
| node1 | 127.0.0.1:3306      | master | up    | 2015-08-07 15:54:44 +0800 CST | 16          | 1        |
| node2 | 192.168.59.103:3307 | master | up    | 2015-08-07 15:54:44 +0800 CST | 16          | 1        |
+-------+---------------------+--------+-------+-------------------------------+-------------+----------+
2 rows in set (0.00 sec)

#view the config of schema
mysql> admin server(opt,k,v) values('show','schema','config');
+-----------+------------------+---------+------+--------------+-----------+---------------+
| DB        | Table            | Type    | Key  | Nodes_List   | Locations | TableRowLimit |
+-----------+------------------+---------+------+--------------+-----------+---------------+
| kingshard |                  | default |      | node1        |           | 0             |
| kingshard | test_shard_hash  | hash    | id   | node1, node2 | 4, 4      | 0             |
| kingshard | test_shard_range | range   | id   | node1, node2 | 4, 4      | 10000         |
+-----------+------------------+---------+------+--------------+-----------+---------------+
3 rows in set (0.00 sec)

#view the config of white list ip
mysql> admin server(opt,k,v) values('show','allow_ip','config');
+--------------+
| AllowIP      |
+--------------+
| 127.0.0.1    |
| 192.168.10.1 |
+--------------+
2 rows in set (0.00 sec)

#view the config of black list sql
mysql> admin server(opt,k,v) values('show','black_sql','config');
+-------------------------------+
| BlackListSql                  |
+-------------------------------+
| select * from sbtest1         |
| select * from sbtest1 limit ? |
+-------------------------------+
2 rows in set (0.00 sec)

```

### 5.3 Change the config of kingshard

```
#turn off the sql log
admin server(opt,k,v) values('change','log_sql','off')

#turn on the sql log
admin server(opt,k,v) values('change','log_sql','on')

#change the threshold of slow log time
admin server(opt,k,v) values('change','slow_log_time','50');

#add white list ip
admin server(opt,k,v) values('add','allow_ip','127.0.0.1');

#delete white list ip
admin server(opt,k,v) values('del','allow_ip','127.0.0.1');

#add black list sql
admin server(opt,k,v) values('add','black_sql','select count(*) from sbtest1')

#delete black list sql
admin server(opt,k,v) values('del','black_sql','select count(*) from sbtest1')

#save config
admin server(opt,k,v) values('save','proxy','config')
```

### 5.4 support LVS/Keepalived

```
#show status of kingshard
admin server(opt,k,v) values('show','proxy','status')

#change status of kingshard online/offline
admin server(opt,k,v) values('change','proxy','online')

`````

## 6.Requirement and feedback

If You have new functional requirements about kingshard in the production environment, or find a bug in the process of using kingshard. Welcome to send a mail to `flikecn#126.com`, I will reply you as soon as possible.


