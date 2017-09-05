# Kingshard支持SQL的范围

## 1.简要说明
kingshard在非分表的情况下支持绝大部分MySQL语法和协议，包括类似SHOW DATABASES, SHOW TABLES, 以及各种DML语句和DDL语句。在分表的情况下，目前只支持有限的DML语句，主要包含：SELECT,UPDATE,INSERT,REPLACE, DELETE这五种SQL操作。并且不支持自动建子表功能。以及有限的kingshard自定义管理端命令。在分表和非分表的情况下，都不支持以下情形：
- 暂不支持用户自定义数据类型、自定义函数。
- 暂不支持视图、存储过程、触发器、游标。
- 暂不支持类似 BEGIN…END，LOOP...END LOOP，REPEAT...UNTIL...END REPEAT，WHILE...DO...END WHILE 等的复合语句。
- 暂不支类似 IF,WHILE 等流程控制类语句。
下面分两部分介绍kingshard支持SQL的情况：非分表情况下SQL支持范围和分表情况下SQL支持范围。

## 2.非分表情况下SQL的支持范围
以下说明都是基于非分表的情况下，SQL的支持情况。
2.1 数据库DDL语法
- CREATE TABLE Syntax
- CREATE INDEX Syntax
- DROP TABLE Syntax
- DROP INDEX Syntax
- ALTER TABLE Syntax
- TRUNCATE TABLE Syntax

### 2.2 数据库DML语法
- INSERT Syntax
- INSERT DELAYED Syntax 暂不支持
- REPLACE Syntax
- UPDATE Syntax
- DELETE Syntax
- Subquery Syntax
- Scalar Subquery
- Comparisons Subquery
- Subqueries with ANY, IN, or SOME
- Subqueries with ALL
- Row Subqueries
- Subqueries with EXISTS or NOT EXISTS
- Subqueries in the FROM Clause
- SELECT Syntax
- SELECT INTO OUTFILE/INTO DUMPFILE/INTO var_name 暂不支持
- Last_insert_id特性

### 2.3 事务的支持
- START TRANSACTION, COMMIT, and ROLLBACK Syntax
- 暂不支持transaction_characteristic定义
- 暂不支持savepoint嵌套事务的相关语法
- 暂不支持XA事务的相关语法
- 支持set autocommit=0/1方式设置事务.
- 支持begin/commit方式设置事务
- 支持start transaction方式设置事务
- SET TRANSACTION Syntax
- 暂不支持对global的事务隔离级别进行调整

### 2.4 预处理的支持
- Prepared Statements
支持主流语言（java,php,python,C/C++,Go)SDK的MySQL的Prepare语法。

### 2.5 数据库管理语法的支持
- SET Syntax
只支持字符集和set autocommit相关语法，其他set语法未测试过。
- Show Syntax
默认show操作会转发到默认DB，需要查看其他DB的内容，通过在SQL中加注释的方式。
- KILL Syntax
目前不支持KILL QUERY processlist_id

### 2.6 数据库管理语法的支持
- DESCRIBE Syntax
- EXPLAIN Syntax
- USE Syntax

### 2.7 数据库系统函数的支持
默认都支持（未测试）

## 3.分表的情况下SQL的支持范围

### 3.1 数据库DDL语法
- CREATE TABLE Syntax
- CREATE INDEX Syntax
- DROP TABLE Syntax
- DROP INDEX Syntax
- ALTER TABLE Syntax
- TRUNCATE TABLE Syntax
 
分表的情况下支持这些语法，但需要在SQL中加注释，例如：
`/*node1*/create table stu_0000(id int, name char(20));`
这样kingshard就会将该SQL转发到node1节点的Master上。

**注：**
`truncate`如果不指定节点注释则会将所有分表都清空，例如：`truncate stu`
### 3.2 数据库DML语法
- INSERT Syntax
- INSERT DELAYED Syntax 不支持
- INSERT INTO SELECT 不支持
- REPLACE Syntax
- UPDATE Syntax
//分表使用的字段无论何种分表类型都不能作为被更新的字段。
- UPDATE SET xx=REPLACE(xx,'a','b') Syntax 不支持
- DELETE Syntax
- Subquery Syntax
- SELECT Syntax
对于UPDATE，DELETE和SELECT三种SQL中WHERE后面的条件不能包含子查询，函数等。只能是字段名。

### 3.3 数据库管理语法的支持
- DESCRIBE Syntax
通过SQL语句hint方式支持，例如:` /*node2*/describe table_name`
- EXPLAIN Syntax
通过SQL语句hint方式支持，例如:` /*node2*/explain select * from xxxx`
- USE Syntax

### 3.4 分表聚合函数的支持
支持以下聚合函数：
- sum函数
- max函数
- count函数
- min函数
不支持distinct后聚合，例如:` select count(distinct id) from xxxx`

### 3.4 分表group by,order by,limit支持
支持分表情况下的group by, order by, limit

### 3.5 其他情形说明
- 不支持分布式事务，支持以非事务的方式更新多node上的数据。
- 不支持预处理。
- 不支持数据库管理语法。
