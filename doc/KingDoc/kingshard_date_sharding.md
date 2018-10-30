# kingshard按时间分表功能介绍

在[文档](https://github.com/flike/kingshard/blob/master/doc/KingDoc/how_to_use_kingshard.md)中主要介绍了kingshard的Hash和Range方式的分表，最近又开发了按时间维度的分表方式。按时间维度分表的场景非常普遍，下面介绍一下kingshard的时间分表功能

## 1. 支持的时间类型

kingshard中的分表字段支持MySQL中三种类型的时间格式

- date类型，格式：YYYY-MM-DD，例如:2016-03-04,注意：2016-3-04，2016-03-4，2016-3-4等格式kingshard都是不支持的。
- datetime类型，格式：YYYY-MM-DD HH:MM:SS，例如:2016-03-04 13:23:43,注意：2016-3-04 13:23:43，2016-03-4  13:23:43，2016-3-4  13:23:43等格式kingshard都是不支持的，必须严格按照规定的格式，kingshard才支持。
- timestamp类型，整数类型，例如：1457165568，对应的是：2016-3-5 16:12:48。

## 2. 支持的时间分表类型

kingshard支持MySQL中三种格式的时间类型

- date类型，格式：YYYY-MM-DD，例如:2016-03-04,注意：2016-3-04，2016-03-4，2016-3-4等格式kingshard都是不支持的。
- datetime，格式：YYYY-MM-DD HH:MM:SS，例如:2016-03-04 13:23:43,注意：2016-3-04 13:23:43，2016-03-4  13:23:43，2016-3-4  13:23:43等格式kingshard都是不支持的，必须严格按照规定的格式，kingshard才支持。
- timestamp，整数类型。

## 3. 功能演示

kingshard的配置文件如下所示：

```
# server listen addr
addr : 0.0.0.0:9696

# user list with user name and password
user_list:
-
user :  kingshard
password : kingshard

# the web api server
web_addr : 0.0.0.0:9797
#HTTP Basic Auth
web_user : admin
web_password : admin

# if set log_path, the sql log will write into log_path/sql.log,the system log
# will write into log_path/sys.log
#log_path : /Users/flike/log

# log level[debug|info|warn|error],default error
log_level : debug

# if set log_sql(on|off) off,the sql log will not output
log_sql: on

# only log the query that take more than slow_log_time ms
#slow_log_time : 100

# the path of blacklist sql file
# all these sqls in the file will been forbidden by kingshard
#blacklist_sql_file: /Users/flike/blacklist

# only allow this ip list ip to connect kingshard
#allow_ips: 127.0.0.1

# the charset of kingshard, if you don't set this item
# the default charset of kingshard is utf8.
#proxy_charset: gbk

# node is an agenda for real remote mysql server.
nodes :
-
    name : node1

    # default max conns for mysql server
    max_conns_limit : 32

    # all mysql in a node must have the same user and password
    user :  kingshard
    password : kingshard

    # master represents a real mysql master server
    master : 127.0.0.1:3306

    # slave represents a real mysql salve server,and the number after '@' is
    # read load weight of this slave.
    #slave : 192.168.59.101:3307@2,192.168.59.101:3307@3
    down_after_noalive : 32
-
    name : node2

    # default max conns for mysql server
    max_conns_limit : 32

    # all mysql in a node must have the same user and password
    user :  kingshard
    password : kingshard

    # master represents a real mysql master server
    master : 192.168.59.103:3307

    # slave represents a real mysql salve server
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
       table: test_shard_year
       key: ctime
       type: date_year
       nodes: [node1,node2]
       date_range: [2015-2016,2017-2018]

```

### 3.1 按年分表

#### 3.1.1 配置说明
按年分表的配置项设置如下：

```
       table: test_shard_year
       key: ctime
       type: date_year
       nodes: [node1,node2]
       date_range: [2015-2016,2017-2018]
```
该配置表示：

- sharding key是ctime。
- 按年的分表类型是:`date_year`。
- `test_shard_year_2015, test_shard_year_2016`两个子表落在node1上，`test_shard_year_2017，test_shard_year_2018`两个子表落在node2上。
- 如果你一个node上只包含一张子表，你可以这样配置`date_range[2015,2017-2018]`。

注意：子表的命名格式必须是:shard_table_YYYY,shard_table是分表名，后面接具体的年。**传入范围必须是有序递增的,不能是[2016,2013-2014]**

#### 3.1.2 功能演示
在node1上创建两张子表`test_shard_year_2015, test_shard_year_2016`，在node2上创建两种子表`test_shard_year_2017，test_shard_year_2018`。建表SQL如下

```
CREATE TABLE `test_shard_year_2016` (
  `id` int(10) NOT NULL,
  `name` varchar(40) DEFAULT NULL,
  `ctime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```
插入数据：

```
mysql> insert into test_shard_year(id,name,ctime) values(12,"hello","2015-02-22 13:23:45");
Query OK, 1 row affected (0.01 sec)

mysql> insert into test_shard_year(id,name,ctime) values(13,"world","2016-03-22");
Query OK, 1 row affected (0.00 sec)

mysql> select * from test_shard_year where ctime < "2016-03-23";
+----+-------+---------------------+
| id | name  | ctime               |
+----+-------+---------------------+
| 12 | hello | 2015-02-22 13:23:45 |
| 13 | world | 2016-03-22 00:00:00 |
+----+-------+---------------------+
2 rows in set (0.00 sec)

```
对应的SQL log信息是：

```
2016/03/05 12:06:32 - OK - 1.2ms - 127.0.0.1:56597->127.0.0.1:3306:insert into test_shard_year_2015(id, name, ctime) values (12, 'hello', '2015-02-22 13:23:45')
2016/03/05 12:06:59 - OK - 2.0ms - 127.0.0.1:56597->127.0.0.1:3306:insert into test_shard_year_2016(id, name, ctime) values (13, 'world', '2016-03-22')
2016/03/05 12:08:30 - OK - 1.6ms - 127.0.0.1:56597->127.0.0.1:3306:select * from test_shard_year_2015 where ctime < '2016-03-23'
2016/03/05 12:08:30 - OK - 0.3ms - 127.0.0.1:56597->127.0.0.1:3306:select * from test_shard_year_2016 where ctime < '2016-03-23'
```

当然如果你把id作为一个unix时间戳，来分表的话，kingshard也是支持的。具体配置就是这样的：

```
       table: test_shard_year
       key: id
       type: date_year
       nodes: [node1,node2]
       date_range: [2015-2016,2017-2018]

```

插入数据:

```
mysql> insert into test_shard_year(id,name,ctime) values(1457410310,"world","2018-03-22");
Query OK, 1 row affected (0.01 sec)

mysql> select * from test_shard_year where id = 1457410310;
+------------+-------+---------------------+
| id         | name  | ctime               |
+------------+-------+---------------------+
| 1457410310 | world | 2018-03-22 00:00:00 |
+------------+-------+---------------------+
1 row in set (0.00 sec)

```

1457410310 这个unix时间戳对应的日期是：2016-3-8 12:11:50。kingshard准确地将这条记录路由到了`test_shard_year_2016`这张子表中了。
对应的SQL log是：

```
2016/03/08 12:12:49 - OK - 1.0ms - 127.0.0.1:56669->127.0.0.1:3306:insert into test_shard_year_2016(id, name, ctime) values (1457410310, 'world', '2018-03-22')
2016/03/08 12:13:23 - OK - 0.4ms - 127.0.0.1:56669->127.0.0.1:3306:select * from test_shard_year_2016 where id = 1457410310

```

### 3.2 按月分表

#### 配置说明
按月分表的配置项设置如下：

```
       table: test_shard_month
       key: ctime
       type: date_month
       nodes: [node1,node2]
       date_range: [201512-201602,201609-2016010]
```
该配置表示：

- sharding key是ctime。
- 按月的分表类型是:`date_month`。
- `test_shard_month_201512, test_shard_month_201601, test_shard_month_201602`两个子表落在node1上，`test_shard_month_201609，test_shard_month_201610`两个子表落在node2上。
- 如果你一个node上只包含一张子表，你可以这样配置`date_range[201501,201609-201610]`。

注意：子表的命名格式必须是:`shard_table_YYYYMM,shard_table`是分表名，后面接具体的年和月。**传入范围必须是有序递增的,不能是[201609-201610,201501]**

功能演示参考按年分表的操作。

### 3.3 按天分表

#### 配置说明
按天分表的配置项设置如下：

```
       table: test_shard_day
       key: ctime
       type: date_day
       nodes: [node1,node2]
       date_range: [20151222-20151224,20160901-20160902]
```
该配置表示：

- sharding key是ctime。
- 按天的分表类型是:`date_day`。
- `test_shard_day_20151222, test_shard_day_20151223, test_shard_day_20151224`两个子表落在node1上，`test_shard_day_20160901，test_shard_day_20160902`两个子表落在node2上。
- 如果你一个node上只包含一张子表，你可以这样配置`date_range[20150101,20160901-20161010]`。

注意：子表的命名格式必须是:`shard_table_YYYYMMDD,shard_table`是分表名，后面接具体的年,月和日。**传入范围必须是有序递增的,不能是[20160901-20161010,20150101]**

功能演示参考按年分表的操作。
