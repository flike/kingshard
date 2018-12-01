# 管理端命令 [Web API版本](./kingshard_admin_api.md)

kingshard的管理端口复用了工作端口，通过特定的关键字来标示，目前支持对后端DB常用的管理操作。kingshard支持了多用户，
只有root用户才有权限操作admin相关命令。

## 平滑上（下）线后端DB

```
#添加一个新的slave到node1
admin node(opt,node,k,v) values('add','node1','slave','127.0.0.1:3306')

#删除node1上的一个slave。注意：只能删除slave，不能删除master
admin node(opt,node,k,v) values('del','node1','slave','127.0.0.1:3306')

#将一个slave设置为下线状态
admin node(opt,node,k,v) values('down','node1','slave','127.0.0.1:3306')

#将一个slave设置为上线状态
admin node(opt,node,k,v) values('up','node1','slave','127.0.0.1:3306')

#将master设置为下线状态
admin node(opt,node,k,v) values('down','node1','master','127.0.0.1:3306')

#将master设置为上线状态
admin node(opt,node,k,v) values('up','node1','master','127.0.0.1:3306')

```

## 查看kingshard配置

```
#查看kingshard全局配置
mysql> admin server(opt,k,v) values('show','proxy','config');
+--------------+----------------+
| Key          |   Value        |
+--------------+----------------+
| Addr         | 127.0.0.1:9696 |
| User_List    | root,kingshard |
| LogPath      | ./             |
| LogLevel     | debug          |
| LogSql       | on             |
| SlowLogTime  | 10             |
| Nodes_Count  | 2              |
| Nodes_List   | node1,node2    |
| ClientConns  | 32             |
| ClientQPS    | 15             |
| ErrLogTotal  | 12             |
| SlowLogTotal | 26             |
+--------------+----------------+
6 rows in set (0.00 sec)

ClientConns:客户端连接数
ClientQPS:客户端的QPS大小
ErrLogTotal:kingshard启动以来产生的错误日志个数
SlowLogTotal:kingshard启动以来产生的慢日志个数

#查看node状态
mysql> admin server(opt,k,v) values('show','node','config');
+-------+--------------------+--------+-------+-------------------------------+---------+----------+------------+---------------+--------------+
| Node  | Address            | Type   | State | LastPing                      | MaxConn | IdleConn | CacheConns | PushConnCount | PopConnCount |
+-------+--------------------+--------+-------+-------------------------------+---------+----------+------------+---------------+--------------+
| node1 | 127.0.0.1:3306      | master | up    | 2015-08-07 15:54:44 +0800 CST  | 512     | 509      | 2          | 6301447       | 6300936      |
| node2 | 192.168.59.103:3307 | master | up    | 2015-08-07 15:54:44 +0800 CST  | 512     | 509      | 2          | 6301447       | 6300936      |
+-------+--------------------+--------+-------+-------------------------------+---------+----------+------------+---------------+--------------+
2 rows in set (0.00 sec)

#查看schema配置

mysql> admin server(opt,k,v) values('show','schema','config');
+-----------+-----------+------------------+---------+------+--------------+-----------+---------------+
| User      | DB        | Table            | Type    | Key  | Nodes_List   | Locations | TableRowLimit |
+-----------+-----------+------------------+---------+------+--------------+-----------+---------------+
| kingshard | kingshard |                  | default |      | node1        |           | 0             |
| kingshard | kingshard | test_shard_hash  | hash    | id   | node1, node2 | 4, 4      | 0             |
| kingshard | kingshard | test_shard_range | range   | id   | node1, node2 | 4, 4      | 10000         |
| root      | kingshard |                  | default |      | node1        |           | 0             |
+-----------+-----------+------------------+---------+------+--------------+-----------+---------------+

3 rows in set (0.00 sec)

#查看白名单ip
mysql> admin server(opt,k,v) values('show','allow_ip','config');
+--------------+
| AllowIP      |
+--------------+
| 127.0.0.1    |
| 192.168.10.1 |
+--------------+
2 rows in set (0.00 sec)

#查看黑名单sql
mysql> admin server(opt,k,v) values('show','black_sql','config');
+-------------------------------+
| BlackListSql                  |
+-------------------------------+
| select * from sbtest1         |
| select * from sbtest1 limit ? |
+-------------------------------+
2 rows in set (0.00 sec)

```

## 修改kingshard配置

为保证kingshard的安全性，管理端命令，只能通过root用户来操作，其他用户不能操作。也就是说用户列表中，需要配置一个root用户

```
#关闭sql日志打印
admin server(opt,k,v) values('change','log_sql','off')

#开启sql日志打印
admin server(opt,k,v) values('change','log_sql','on')

#修改慢sql日志时间, 单位ms
admin server(opt,k,v) values('change','slow_log_time','50');

#添加白名单IP
#支持IP或IP段
admin server(opt,k,v) values('add','allow_ip','192.168.14.0/24');
admin server(opt,k,v) values('add','allow_ip','192.168.15.1');

#删除白名单IP
admin server(opt,k,v) values('del','allow_ip','127.0.0.1');

#添加黑名单sql语句
admin server(opt,k,v) values('add','black_sql','select count(*) from sbtest1')

#删除黑名单sql语句
admin server(opt,k,v) values('del','black_sql','select count(*) from sbtest1')

#保存当前配置
admin server(opt,k,v) values('save','proxy','config')

```

上面是通过命令方式修改部分kingshard的配置，但更推荐通过修改配置文件，然后动态加载的方式来
热加载配置文件，步骤如下：

1.修改kingshard **正在使用的配置文件**，不能是kingshard未使用的文件

2.向kingshard发送USR1信号，kingshard就会加载新配置文件的全部内容

```
kill -USR1 pid_of_kingshard
```

## 支持LVS/Keepalived

```
#查看kingshard运行状态
admin server(opt,k,v) values('show','proxy','status')

#改变kingshard运行状态 online: 在线 offline: 下线
admin server(opt,k,v) values('change','proxy','online')

```
