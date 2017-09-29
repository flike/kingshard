# kingshard的FAQ

**1. kingshard的线上使用情况和稳定性如何？**

- 目前据开发者不完全统计，已经有上十家公司在生产环境使用。
- 很多MySQL Proxy调研用户都比较关心kingshard的稳定性，从反馈情况来看，目前kingshard已经较为稳定，
Bug出现的次数越来越收敛了。对于想生产环境使用kingshard的用户来说，为保证安全起见，还是需要自己利用kingshard部署一套
测试环境，自行测试一下。如果遇到BUG，可以在github的issue中提出，但希望用户在描述问题的时候，尽量详细，以便开发者快速定位问题。
- 很多用户利用kingshard的跨平台特性，在windows的开发环境部署kingshard用于开发调试。待开发稳定后将业务代码和kingshard都部署Linux服务器上。

**2. kingshard可以用来做什么？**

kingshard 可以用来对SQL进行读写分离，配置多从库并进行负载均衡。当然还可以利用kingshard的连接池功能，
从而保护后端MySQL不受前端应用影响。最重要的是可以利用kingshard分表，无需前端应用感知。这里只列举小部分功能，更多的功能还需要大家
自己动手去发掘。

**3. kingshard 对事务的支持如何？**

kingshard目前只能支持单机事务，也就是说：只能保证在一台DB上完整执行事务。如果事务跨越了多个DB，kingshard会给客户端返回错误。
所以在使用sharding功能的时候，所有的更新操作务必保证在同一台DB上，否则会由于跨DB而操作事务失败。

**4. kingshard的配置文件中schema中的DB的作用？**

在配置文件中，如果分表的话，会存在如下配置项：

```
schema :
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

这里的db只针对分表，也就是说该分表的子表只能放在这个db中。但是如果不分表，这你的表可以放在任何db
中。

**5. 我有两个表需要分表，且这个两个表在两个不同的database中，该怎么设置schema？**

目前在kingshard的schema规则中，所有分表在配置的时候需要指定database（参考上面的case），如果存在另一个分表在不同的database中，也只需要指定一下db即可。
对于不需要分表的情况，则可以放在任何db，且不需要配置。使用时通过`use db`命令来选择db即可。

**6. `/*node2*/select * from stu`等类似SQL怎么不起作用？**

在命令行执行带注释的SQL时，需要加-c参数避免客户端过滤注释。例如：

```
mysql -h127.0.0.1 -ukingshard -pkingshard -P9696 -c;
```

**7. 利用kingshard分表应该注意什么？**

- 支持跨node的count,sum,max和min等函数。
- 支持单个分表的join操作，即支持分表和另一张不分表的join操作
- 支持order by
- 支持group by

**8. etc目录下有两个配置文件(ks.yaml,unshard.yaml),我该使用哪一个？**

如果你需要分表功能，请基于ks.yaml修改你的配置。如果你不需要分表，请基于unshard.yaml修改你的配置。
当然你如果能够完全掌握如何配置kingshard，你也可以自己编写配置文件。

**9. kingshard如何配合LVS切流量？**

kingshard提供了查看和改变系统状态的管理命令，在高可用状态检测、LVS流量切换、系统无损升级都可以使用该命令，详情见[如何配合LVS实现集群部署](./how_to_use_lvs.md)

**10. 如何给kingshard贡献代码？**

kingshard欢迎大家提交pr，对于pr我建议：

- 代码必须用go fmt格式化过。否则，合并后开发者还需要格式化一下。
- 提交的代码，必须有充分的unit test。这样才能保证代码的质量，要不然代码质量没法保证。
- bugfix 可以直接提交到mater分支。对于新的feature，请创建一个新分支(feature-xxx)，然后提交过来。
- 提交注释使用英文。
