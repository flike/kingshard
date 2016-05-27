# kingshard SQL黑名单功能介绍

## 1. 应用场景介绍
在kingshard开源之后，有用户多次提到能不能在kingshard中加入SQL黑名单机制，让kingshard能够根据特定的规则来拦截在黑名单中的SQL。有几个比较典型的应用场景：

1. DBA定义一些比较危险的SQL，放在SQL黑名单文件中。可以避免前端应用发过来的SQL对数据库造成危害。这种SQL有可能是开发者粗心编写的，也有可能是被SQL注入生成的SQL。例如：`delete from mytable`，这种不带where条件的SQL，会把整个表删除。
2. 在kingshard项目上线后，通过log发现存在大量某条SQL给DB造成了很大的压力。这时候可以动态地将这条SQL加入黑名单，阻止该SQL的执行，从而使数据库压力降低。例如:`select count(*) from mytable where xxxx`,这类SQL如果没有优化得当，是很容易造成系统的IO过高的。

## 2. 功能介绍
在kingshard如果想使用SQL黑名单功能，只需要在配置中添加：
```
blacklist_sql_file: /Users/flike/blacklist
```
然后我们在blacklist定义SQL黑名单，这样kingshard在转发的时候，就会阻止黑名单中SQL的转发。

黑名单SQL以正则表达式的形式定义。对于SQL中的值用`?`或`?+`代替。为保证黑名单有效，最好手动验证一下，kingshard是否正确拦截了黑名单中的SQL。定义规则（上一条是原SQL，对应的下一条是黑名单形式的SQL）可以参考下列例子：

```
SELECT c FROM t WHERE id=1
select c from t where id=?

SELECT * FROM prices.rt_5min where id=1
select * from prices.rt_5min where id=?

select null, 5.001, 5001. from foo
select ?, ?, ? from foo

select 'hello', '\nhello\n', \"hello\", '\\'' from foo
select ?, ?, ?, ? from foo

select 'hello'\n
select ?

select * from t where (base.nid IN  ('1412', '1410', '1411'))
select * from t where (base.nid in(?+))

select * from foo where a in (5) and b in (5, 8,9 ,9 , 10)
select * from foo where a in(?+) and b in(?+)

select * from foo limit 5
select * from foo limit ?

select * from foo limit 5, 10
select * from foo limit ?, ?

select * from foo limit 5 offset 10
select * from foo limit ? offset ?

INSERT INTO t (ts) VALUES (NOW())
insert into t (ts) values(?+)

insert into foo(a, b, c) values(2, 4, 5)
insert into foo(a, b, c) values(?+)

CALL foo(1, 2, 3)
call foo

LOAD DATA INFILE '/tmp/foo.txt' INTO db.tbl
load data infile ? into db.tbl

administrator command: Init DB
administrator command: Init DB

use `foo`
use ?

```

## 3.功能演示
在blacklist加入如下SQL:

```
select count(*) from test_shard_hash where id > ?
select count(*) from test_shard_range
SELECT * FROM WORLD
DELETE FROM WORLD
```

连接kingshard，执行SQL显示如下：

```
mysql> select * from world;
ERROR 1105 (HY000): sql in blacklist.
mysql> select * from world where a > 0;
+------+------+
| a    | b    |
+------+------+
|   10 |   23 |
|   45 |  565 |
+------+------+
2 rows in set (0.00 sec)

mysql> delete from world;
ERROR 1105 (HY000): sql in blacklist.
mysql> delete from world where a =10;
Query OK, 1 row affected (0.00 sec)
#注意在SQL黑名单中该SQL是大于后面有个空格，必须要严格匹配，否则#kingshard不会认为是黑名单SQL
mysql> select count(*) from test_shard_hash where id >1;
+----------+
| count(*) |
+----------+
|       24 |
+----------+
1 row in set (0.02 sec)

mysql> select count(*) from test_shard_hash where id > 1;
ERROR 1105 (HY000): sql in blacklist.
```

用sysbench测试了一下存在blacklist时kingshad的性能，发现性能并没有明显下降，所以可以放心使用该功能。