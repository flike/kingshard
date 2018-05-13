[![Author](https://img.shields.io/badge/author-@flike-blue.svg?style=flat)](http://weibo.com/chenfei001) [![Build Status](https://travis-ci.org/flike/kingshard.svg?branch=master)](https://travis-ci.org/flike/kingshard) [![Foundation](https://img.shields.io/badge/Golang-Foundation-green.svg)](http://golangfoundation.org)

# kingshard简介

kingshard是一个由Go开发高性能MySQL Proxy项目，kingshard在满足基本的读写分离的功能上，致力于简化MySQL分库分表操作；能够让DBA通过kingshard轻松平滑地实现MySQL数据库扩容。 **kingshard的性能是直连MySQL性能的80%以上**。

## 主要功能：

### 1. 基础功能

- 支持SQL读写分离。
- 支持透明的MySQL连接池，不必每次新建连接。
- 支持平滑上线DB或下线DB，前端应用无感知。
- 支持多个slave，slave之间通过权值进行负载均衡。
- 支持强制读主库。
- 支持主流语言（java,php,python,C/C++,Go)SDK的mysql的prepare特性。
- 支持到后端DB的最大连接数限制。
- 支持SQL日志及慢日志输出。
- 支持SQL黑名单机制。
- 支持客户端IP访问白名单机制，只有白名单中的IP才能访问kingshard。
- 支持字符集设置。
- 支持last_insert_id功能。
- 支持热加载配置文件，动态修改kingshard配置项（具体参考管理端命令）。
- 支持以Web API调用的方式管理kingshard。
- 支持多用户模式，不同用户之间的表是权限隔离的，互补感知。

### 2. sharding功能

- 支持按整数的hash和range分表方式。
- 支持按年、月、日维度的时间分表方式。
- 支持跨节点分表，子表可以分布在不同的节点。
- 支持跨节点的count,sum,max和min等聚合函数。
- 支持单个分表的join操作，即支持分表和另一张不分表的join操作。
- 支持跨节点的order by,group by,limit等操作。
- 支持将sql发送到特定的节点执行。
- 支持在单个节点上执行事务，不支持跨多节点的分布式事务。
- 支持非事务方式更新（insert,delete,update,replace）多个node上的子表。

## kinshard文档

### kingshard安装和使用

[1.安装kingshard](./doc/KingDoc/kingshard_install_document.md)

[2.如何利用一个数据库中间件扩展MySQL集群——kingshard使用指南](./doc/KingDoc/how_to_use_kingshard.md)

[3.kingshard sharding介绍](./doc/KingDoc/kingshard_sharding_introduce.md)

[4.kingshard按时间分表功能介绍](./doc/KingDoc/kingshard_date_sharding.md)

[5.kingshard 快速入门](./doc/KingDoc/kingshard_quick_try.md)

[6.管理端命令介绍](./doc/KingDoc/admin_command_introduce.md)

[7.管理端Web API接口介绍](./doc/KingDoc/kingshard_admin_api.md)

[8.kingshard SQL黑名单功能介绍](./doc/KingDoc/sql_blacklist_introduce.md)

[9.kingshard的FAQ](./doc/KingDoc/function_FAQ.md)

[10.kingshard SQL支持范围](./doc/KingDoc/kingshard_support_sql.md)

[11.如何配合LVS实现集群部署](./doc/KingDoc/how_to_use_lvs.md)

### kingshard架构与设计

[1.kingshard架构设计和功能实现](./doc/KingDoc/architecture_of_kingshard_CN.md)

[2.kingshard性能优化之网络篇](./doc/KingDoc/kingshard_performance_profiling.md)

[3.kingshard性能测试报告](./doc/KingDoc/kingshard_performance_test.md)
## 鸣谢
- 感谢[mixer](https://github.com/siddontang/mixer)作者siddontang, kingshard最初的版本正是基于mixer开发而来的。
- 感谢[bigpyer](https://github.com/bigpyer)，他对kingshard做了详细的性能测试，并撰写了一份非常详细的测试报告。
- 感谢以下[开源爱好者](https://github.com/flike/kingshard/graphs/contributors)为kingshard做出的贡献。

## kingshard用户列表

https://github.com/flike/kingshard/issues/148

## 技术支持Plus
Kingshard作为开源软件，会一直开源下去。但为了将kingshard项目更好地维护和发展下去，为有需要的用户提供更加全面的技术支撑服务，kingshard**推出了有偿技术服务**，主要包括但不限于以下几类：

* 咨询服务，为用户提供全方位的针对kingshard数据库中间件部署、使用、监控和告警等方面的咨询服务。
* 定制化开发服务，为用户基于kingshard提供可靠地定制化数据库中间件方案。

详情请邮件咨询:flikechen#qq.com

## 反馈
kingshard开源以来，经过不断地迭代开发，功能较为完善，稳定性有较大提升。 **目前已有超过50家公司在生产环境使用kingshard作为MySQL代理。** 如果您在使用kingshard的过程中发现BUG或者有新的功能需求，请发邮件至flikechen#qq.com与作者取得联系，或者加入QQ群(147926796)交流。
欢迎关注**后端技术快讯**公众号，有关kingshard的最新消息与后端架构设计类的文章，都会在这个公众号分享。

<img src="./doc/KingDoc/wechat_pic.png" width="20%" height="20%">

## License

kingshard采用Apache 2.0协议，相关协议请参看[目录](./doc/License)
