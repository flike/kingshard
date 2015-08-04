# 管理端命令

kingshard的管理端口复用了工作端口，通过特定的关键字来标示，目前支持对后端DB常用的管理操作。

## 平滑上（下）线后端DB

```
#添加一个新的slave到node1
admin node(opt,node,k,v) values(‘add’,’node1’,’slave’,’127.0.0.1:3306’)

#删除node1上的一个slave。注意：只能删除slave，不能删除master
admin node(opt,node,k,v) values(‘del’,’node1’,’slave’,’127.0.0.1:3306’)

#将一个slave设置为下线状态
admin node(opt,node,k,v) values(‘down’,’node1’,’slave’,’127.0.0.1:3306’)

#将一个slave设置为上线状态
admin node(opt,node,k,v) values(‘up’,’node1’,’slave’,’127.0.0.1:3306’)

#将master设置为下线状态
admin node(opt,node,k,v) values(‘down’,’node1’,’master’,’127.0.0.1:3306’)

#将master设置为上线状态
admin node(opt,node,k,v) values(‘up’,’node1’,’master’,’127.0.0.1:3306’)

```

## 查看kingshard配置

```
admin server(opt,k,v) values('show','proxy','config')

```