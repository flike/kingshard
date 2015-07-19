## 安装kingshard
```
	1.	安装Go语言环境，具体步骤请Google。
	2.	git clone https://github.com/flike/kingshard.git src/github.com/flike/kingshard
	3. cd src/github.com/flike/kingshard
	4. sh ./dev.sh
	5. make
	6.设置配置文件
	7.运行kingshard。./bin/kingshard -config=etc/multi.yaml
```

##配置文件说明
```
# kingshard的地址和端口
addr : 127.0.0.1:9696

# 连接kingshard的用户名和密码
user :  kingshard
password : kingshard

# log级别，[debug|info|warn|error],默认是error
log_level : debug
# 只允许下面的IP列表连接kingshard
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

    # slave的地址和端口，可不配置
    slave : 
    #kingshard在300秒内都连接不上mysql，则会下线该mysql
    down_after_noalive : 300
- 
    name : node2 
    idle_conns : 16
    rw_split: true
    user :  kingshard 
    password : kingshard

    master : 192.168.59.103:3307
    slave : 
    down_after_noalive: 100

# 分表规则
schemas :
-
    db : kingshard
    nodes: [node1,node2]
    rules:
        default: node1
        shard:
        -   
            table: test_shard_hash
            key: id
            nodes: [node1, node2]
            type: hash
            locations: [4,4]

        -   
            table: test_shard_range
            key: id
            type: range
            nodes: [node1, node2]
            locations: [4,4]
            table_row_limit: 10000

```

##Tips
**kingshard采用的是yaml方式解析配置文件，需要注意的是yaml配置文件不允许出现tab键，且冒号后面需要跟一个空格。配置文件编写完成后，可以在[yaml lint](http://www.yamllint.com/)网站验证是否有格式错误。**