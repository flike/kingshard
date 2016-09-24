#kingshard 管理端API设计

##接口定义和参数说明
- 采用restful风格API设计。
- 假设kingshard的Web Server IP和端口是：127.0.0.1：9797，web用户名：admin，密码：admin

###查看node的状态

```
Action:GET
URL:http://127.0.0.1:9797/api/v1/nodes/status
参数：无
返回结果：node数组，node中包含Master和Slave信息，
字段意思参考配置文件说明
```
####示例
```
curl -X GET \
  -H 'Content-Type: application/json' \
  -u admin:admin \
  127.0.0.1:9797/api/v1/nodes/status
 返回结果：
 [
    {
        "node": "node1",
        "address": "127.0.0.1:3307",
        "type": "master",
        "status": "up",
        "laste_ping": "2016-09-24 17:17:52 +0800 CST",
        "max_conn": 32,
        "idle_conn": 8
    },
    {
        "node": "node2",
        "address": "127.0.0.1:3309",
        "type": "master",
        "status": "up",
        "laste_ping": "2016-09-24 17:17:52 +0800 CST",
        "max_conn": 32,
        "idle_conn": 8
    }
]
```

###添加slave

```
Action:POST
URL:http://127.0.0.1:9797/api/v1/nodes/slaves
参数：
- node：节点名字
- addr：slave的IP和端口
返回结果：成功:"ok",失败："error message"
```
####示例
```
 curl -X POST \
  -H 'Content-Type: application/json' \
  -u admin:admin \
  -d '{"node":"node1","addr":"127.0.0.1:3309"}' \
  127.0.0.1:9797/api/v1/nodes/slaves
  返回结果："ok"
```
###删除slave

```
Action:DELETE
URL:http://127.0.0.1:9797/api/v1/nodes/slaves
参数：
- node：节点名字
- addr：slave的IP和端口
返回结果：成功:"ok",失败："error message"
```
####示例
```
 curl -X DELETE \
  -H 'Content-Type: application/json' \
  -u admin:admin \
  -d '{"node":"node1","addr":"127.0.0.1:3309"}' \
  127.0.0.1:9797/api/v1/nodes/slaves
  返回结果："ok"
```
###设置slave状态

```
Action:PUT
URL：http://127.0.0.1:9797/api/v1/nodes/slaves/status
参数：
- opt："up" or "down"，上线或下线slave
- node：节点名字
- addr：slave的IP和端口
返回结果：成功:"ok",失败："error message"
```

####示例
```
 curl -X PUT \
  -H 'Content-Type: application/json' \
  -u admin:admin \
  -d '{"opt":"up","node":"node1","addr":"127.0.0.1:3309"}' \
  http://127.0.0.1:9797/api/v1/nodes/slaves/status
  返回结果："ok"

```

###设置master状态

```
Action:PUT
URL：http://127.0.0.1:9797/api/v1/nodes/masters/status
参数：
- opt："up" or "down"，上线或下线slave
- node：节点名字
- addr：slave的IP和端口
返回结果：成功:"ok",失败："error message"
```

####示例
```
 curl -X PUT \
  -H 'Content-Type: application/json' \
  -u admin:admin \
  -d '{"opt":"down","node":"node2","addr":"127.0.0.1:3309"}' \
  http://127.0.0.1:9797/api/v1/nodes/masters/status
  返回结果："ok"
```


###查看proxy状态

```
Action:GET
URL:http://127.0.0.1:9797/api/v1/proxy/status
参数：无
返回结果："online"或者"offline"
注意：该API主要用于配合LVS平滑下线
```

####示例
```
curl -X GET \
  -H 'Content-Type: application/json' \
  -u admin:admin \
  http://127.0.0.1:9797/api/v1/proxy/status
 返回结果:"online"
```
###设置proxy状态

```
Action:PUT
URL:http://127.0.0.1:9797/api/v1/proxy/status
参数：opt:"online"或者"offline"
返回结果：成功:"ok",失败："error message"
注意：该API主要用于配合LVS平滑下线

```

####示例
```
curl -X PUT \
  -H 'Content-Type: application/json' \
  -u admin:admin \
   -d '{"opt":"offline"}' \
  http://127.0.0.1:9797/api/v1/proxy/status
  返回结果："ok"
```
###查看proxy的schema

```
Action：GET
http://127.0.0.1:9797/api/v1/proxy/schemas
参数：无
返回结果:schema数组，但只有一项。字段意思参考配置文件
```
####示例
```
curl -X GET \
  -H 'Content-Type: application/json' \
  -u admin:admin \
  127.0.0.1:9797/api/v1/proxy/schema
返回结果：
[
    {
        "db": "kingshard",
        "Table": "",
        "Key": "",
        "Nodes": [
            "node1",
            "node2"
        ],
        "Locations": null,
        "Type": "default",
        "TableRowLimit": 0,
        "DateRange": null
    },
    {
        "db": "kingshard",
        "Table": "test_shard_hash",
        "Key": "id",
        "Nodes": [
            "node1",
            "node2"
        ],
        "Locations": [
            4,
            4
        ],
        "Type": "hash",
        "TableRowLimit": 0,
        "DateRange": null
    },
    {
        "db": "kingshard",
        "Table": "test_shard_range",
        "Key": "id",
        "Nodes": [
            "node1",
            "node2"
        ],
        "Locations": [
            4,
            4
        ],
        "Type": "range",
        "TableRowLimit": 10000,
        "DateRange": null
    },
    {
        "db": "kingshard",
        "Table": "test_shard_time",
        "Key": "id",
        "Nodes": [
            "node1",
            "node2"
        ],
        "Locations": [
            2,
            2
        ],
        "Type": "hash",
        "TableRowLimit": 0,
        "DateRange": null
    },
    {
        "db": "kingshard",
        "Table": "test_shard_month",
        "Key": "dtime",
        "Nodes": [
            "node1",
            "node2"
        ],
        "Locations": null,
        "Type": "date_month",
        "TableRowLimit": 0,
        "DateRange": [
            "201603-201605",
            "201609-201612"
        ]
    },
    {
        "db": "kingshard",
        "Table": "test_shard_day",
        "Key": "mtime",
        "Nodes": [
            "node1",
            "node2"
        ],
        "Locations": null,
        "Type": "date_day",
        "TableRowLimit": 0,
        "DateRange": [
            "20160306-20160307",
            "20160308-20160309"
        ]
    }
]
```

###查看proxy的客户端白名单

```
Action：GET
URL：http://127.0.0.1:9797/api/v1/proxy/allow_ips
参数：allow_ips，ip列表数组
返回结果：IP列表
```
####示例
```
curl -X GET \
  -H 'Content-Type: application/json' \
  -u admin:admin \
  localhost:9797/api/v1/proxy/allow_ips
 返回结果：["127.0.0.1","192.168.0.14"]
 
```

###添加proxy的客户端白名单

```
Action：POST
URL：http://127.0.0.1:9797/api/v1/proxy/allow_ips
参数：allow_ips，ip列表数组
返回结果：成功:"ok",失败："error message"
```
####示例
```
curl -X POST \
  -H 'Content-Type: application/json' \
  -u admin:admin \
  -d '{"allow_ips":["127.0.0.1","192.168.0.14","192.168.0.223"]}' \
  127.0.0.1:9797/api/v1/proxy/allow_ips
  返回结果："ok"
 
```
###删除proxy的客户端白名单

```
Action：DELETE
URL：http://127.0.0.1:9797/api/v1/proxy/allow_ips
参数：allow_ips，ip列表数组
返回结果：成功:"ok",失败："error message"
```
####示例
```
curl -X DELETE \
  -H 'Content-Type: application/json' \
  -u admin:admin \
  -d '{"allow_ips":["192.168.0.14","192.168.0.223"]}' \
  127.0.0.1:9797/api/v1/proxy/allow_ips
  返回结果："ok"
 
```

###查看proxy的black_sql

```
Action:GET
http://127.0.0.1：9797/api/v1/proxy/black_sqls
参数：无
返回结果:sql列表
```
####示例
```
curl -X GET \
  -H 'Content-Type: application/json' \
  -u admin:admin \
  127.0.0.1:9797/api/v1/proxy/black_sqls
返回结果：["delete from test_shard_range","delete from test_shard_hash"]
```
###增加proxy的black_sql

```
Action:POST
http://127.0.0.1:9797/api/v1/proxy/black_sqls
参数：sql（注意：一次只能添加一条SQL）
返回结果：成功:"ok",失败："error message"

```
####示例
```
curl -X POST \
  -H 'Content-Type: application/json' \
  -u admin:admin \
  -d '{"sql":"delete from test_shard_range"}' \
  127.0.0.1:9797/api/v1/proxy/black_sqls
  返回结果："ok"
```
###删除proxy的black_sql

```
Action:DELETE
http:// 127.0.0.1:9797/api/v1/proxy/black_sqls
参数：sql（注意：一次只能删除一条SQL）
返回结果：成功:"ok",失败："error message"
```
####示例
```
curl -X DELETE \
  -H 'Content-Type: application/json' \
  -u admin:admin \
  -d '{"sql":"delete from test_shard_range"}' \
  127.0.0.1:9797/api/v1/proxy/black_sqls
  返回结果："ok"
```

###设置proxy的slow sql开关

```
Action:PUT
URL:http://127.0.0.1:9797/api/v1/proxy/slow_sql/status
参数：opt(可选值："on","off")
返回结果：成功:"ok",失败："error message"
```
####示例
```
curl -X PUT \
  -H 'Content-Type: application/json' \
  -u admin:admin \
  -d '{"opt":"on"}' \
  127.0.0.1:9797/api/v1/proxy/slow_sql/status
  返回结果："ok"
```
###查看proxy的slow sql的时间

```
Action:GET
URL:http://127.0.0.1:9797/api/v1/proxy/slow_sql/time
参数：无
返回结果：slow_log 时间，单位ms
```
####示例
```
curl -X GET \
  -H 'Content-Type: application/json' \
  -u admin:admin \
  127.0.0.1:9797/api/v1/proxy/slow_sql/time
  返回结果：500
```

###设置proxy的slow sql的时间

```
Action:PUT
URL:http://127.0.0.1:9797/api/v1/proxy/slow_sql/time
参数：slow_time(慢日志时间，单位ms。执行时间超过该值的SQL会输出到慢日志文件)
返回结果：成功:"ok",失败："error message"
```
####示例
```
curl -X PUT \
  -H 'Content-Type: application/json' \
  -u admin:admin \
  -d '{"slow_time":500}' \
  127.0.0.1:9797/api/v1/proxy/slow_sql/time
  返回结果："ok"
```
###保存proxy的配置

```
Action:PUT
URL:http://127.0.0.1:9797/api/v1/proxy/config/save
参数：无
返回结果：成功:"ok",失败："error message"
说明：将kingshard的配置写入文件
```
####示例
```
curl -X PUT \
  -H 'Content-Type: application/json' \
  -u admin:admin \
  127.0.0.1:9797/api/v1/proxy/config/save
  返回结果："ok"
```
