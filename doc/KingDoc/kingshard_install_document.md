## 安装kingshard
```
	1. 安装Go语言环境，具体步骤请Google。
	2. git clone https://github.com/flike/kingshard.git src/github.com/flike/kingshard
	3. cd src/github.com/flike/kingshard
	4. source ./dev.sh
	5. make
	6. 设置配置文件
	7. 运行kingshard。./bin/kingshard -config=etc/ks.yaml
```

##配置文件说明

```
# kingshard的地址和端口
addr : 0.0.0.0:9696

# 连接kingshard的用户名和密码
user :  kingshard
password : kingshard

# log级别，[debug|info|warn|error],默认是error
log_level : debug
# 打开SQL日志，设置为on;关闭SQL日志，设置为off
log_sql : on
#日志文件路径，如果不配置则会输出到终端。
log_path : /Users/flike/log
# 只允许下面的IP列表连接kingshard，如果不配置则对连接kingshard的IP不做限制。
allow_ips: 127.0.0.1

# 一个node节点表示mysql集群的一个数据分片，包括一主多从（可以不配置从库）
nodes :
    #node节点名字
    name : node1 

    # 连接池中默认的空闲连接数
    idle_conns : 16

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
    idle_conns : 16
    user :  kingshard 
    password : kingshard

    master : 192.168.59.103:3307
    slave : 
    down_after_noalive: 100

# 分表规则
schemas :
-
    #分表使用的db，所有的分表必须都在这个db中。
    db : kingshard
    #分表分布的node名字
    nodes: [node1,node2]
    rules:
    	#所有未分表的SQL，都会发往默认node。
        default: node1
        shard:
        -   
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
            table: test_shard_range
            key: id
            type: range
            nodes: [node1, node2]
            locations: [4,4]
            #表示每个子表包含的最大记录数，也就是说每				  
			#个子表最多包好10000条记录。
            table_row_limit: 10000


```

##Tips
**kingshard采用的是yaml方式解析配置文件，需要注意的是yaml配置文件不允许出现tab键，且冒号后面需要跟一个空格。配置文件编写完成后，可以在[yaml lint](http://www.yamllint.com/)网站验证是否有格式错误。**
