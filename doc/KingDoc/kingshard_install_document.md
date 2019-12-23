## 安装kingshard
```
	1. 安装Go语言环境（请使用最新版），具体步骤请Google。
	2. git clone https://github.com/flike/kingshard.git $GOPATH/src/github.com/flike/kingshard
	3. cd src/github.com/flike/kingshard
	4. source ./dev.sh
	5. make
	6. 设置配置文件
	7. 运行kingshard。./bin/kingshard -config=etc/ks.yaml
```

## 配置文件说明

```
# kingshard的地址和端口
addr : 0.0.0.0:9696

# 连接kingshard的用户名和密码的用户列表
user_list:
-
    user :  kingshard
    password : kingshard
#kingshard的web API 端口
web_addr : 0.0.0.0:9797
#调用API的用户名和密码
web_user : admin
web_password : admin

# log级别，[debug|info|warn|error],默认是error
log_level : debug
# 打开SQL日志，设置为on;关闭SQL日志，设置为off
log_sql : on
#如果设置了该项，则只输出SQL执行时间超过slow_log_time(单位毫秒)的SQL日志，不设置则输出全部SQL日志
slow_log_time : 100
#日志文件路径，如果不配置则会输出到终端。
log_path : /Users/flike/log
# sql黑名单文件路径
# 所有在该文件中的sql都会被kingshard拒绝转发
#blacklist_sql_file: /Users/flike/blacklist
# 只允许下面的IP列表连接kingshard，如果不配置则对连接kingshard的IP不做限制。
allow_ips: 127.0.0.1
# kingshard使用的字符集，如果不设置该选项，则kingshard使用utf8作为默认字符集
#proxy_charset: utf8mb4

# 一个node节点表示mysql集群的一个数据分片，包括一主多从（可以不配置从库）
nodes :
-
    #node节点名字
    name : node1

    # 连接池中最大空闲连接数，也就是最多与后端DB建立max_conns_limit个连接
    max_conns_limit : 16

    # kingshard连接该node中mysql的用户名和密码，master和slave的用户名和密码必须一致
    user :  kingshard
    password : kingshard

    # master的地址和端口
    master : 127.0.0.1:3306

    # slave的地址、端口和读权重，@后面的表示该slave的读权重。可不配置slave
    #slave : 192.168.0.12@2,192.168.0.13@3
    #kingshard在300秒内都连接不上mysql，kingshard则会下线该mysql
    down_after_noalive : 300
-
    name : node2
    max_conns_limit : 16
    user :  kingshard
    password : kingshard

    master : 192.168.59.103:3307
    slave :
    down_after_noalive: 100

# 各用户的分表规则
schema_list :
-
    #schema的所属用户名
    user: kingshard
    nodes: [node1,node2]
    #分表分布的node名字
    nodes: [node1,node2]
    #所有未分表的SQL，都会发往默认node。
    default: node1
    shard:
    -
        #分表使用的db
        db : kingshard
	#分表名字
        table: test_shard_hash
        #分表字段
        key: id
        #分表分布的node
        nodes: [node1, node2]
        #分表类型
        type: hash
        #子表个数分布，表示node1有4个子表，
        #node2有4个子表。
        locations: [4,4]

    -
	#分表使用的db
        db : kingshard
	#分表名字
        table: test_shard_range
	#分表字段
        key: id
	#分表类型
        type: range
	#分表分布的node
        nodes: [node1, node2]
	#子表个数分布，表示node1有4个子表，
	#node2有4个子表。
        locations: [4,4]
        #表示每个子表包含的最大记录数，也就是说每
	#个子表最多包好10000条记录。即子表1对应的id为[0,10000),子表2[10000,20000)....
        table_row_limit: 10000


```

## 注意

**1. kingshard会响应SIGINT,SIGTERM,SIGQUIT这三个信号，平滑退出。在部署kingshard机器上应避免产生这三个信号，以免造成kingshard非正常退出！后台运行kingshard建议使用supervisor工具**

**2. kingshard采用的是yaml方式解析配置文件，需要注意的是yaml配置文件不允许出现tab键，且冒号后面需要跟一个空格。配置文件编写完成后，可以在[yaml lint](http://www.yamllint.com/)网站验证是否有格式错误。**

**3. windows下安装kingshard，参考[文档](https://github.com/flike/kingshard/wiki/%E5%9C%A8window%E4%B8%8B%E5%AE%89%E8%A3%85kingshard)**

**4. 可以通过`./bin/kingshard -v`来查看kingshard的commit hash和编译时间，从而维持kingshard的版本。**


