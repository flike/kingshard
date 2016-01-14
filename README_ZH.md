[![Build Status](https://travis-ci.org/flike/kingshard.svg?branch=master)](https://travis-ci.org/flike/kingshard)

# kingshard简介

kingshard是一个由Go开发高性能MySQL Proxy项目，kingshard在满足基本的读写分离的功能上，致力于简化MySQL分库分表操作；能够让DBA通过kingshard轻松平滑地实现MySQL数据库扩容。 **kingshard的性能大约是直连MySQL性能的80%以上**。

## 主要功能：

	1. 读写分离。
	2. 跨节点分表。
	3. 支持透明的MySQL连接池，不必每次新建连接。
	4. 支持客户端IP访问控制。
	5. 平滑上线DB或下线DB，前端应用无感知。
	6. 支持多个slave，slave之间通过权值进行负载均衡。
	7. 支持强制读主库。
	8. 支持将sql发送到特定的node。
	9. 支持在单个node上执行事务，不支持跨多个node执行事务。
	10. 支持跨node的count,sum,max和min等函数。
	11. 支持单个分表的join操作，即支持分表和另一张不分表的join操作。
	12. 支持order by,group by,limit等操作。
	13. 支持主流语言（java,php,python,C/C++,Go)SDK的mysql的prepare特性。
	14. 支持到后端DB的最大连接数限制。
	15. 支持SQL日志及慢日志输出。

## kinshard相关文档

### kingshard安装和使用

[1.安装kingshard](./doc/KingDoc/kingshard_install_document.md)

[2.如何利用一个数据库中间件扩展MySQL集群——kingshard使用指南](./doc/KingDoc/how_to_use_kingshard.md)

[3.kingshard sharding介绍](./doc/KingDoc/kingshard_sharding_introduce.md)

[4.kingshard 快速入门](./doc/KingDoc/kingshard_quick_try.md)

[5.管理端命令介绍](./doc/KingDoc/admin_command_introduce.md)

[6.kingshard性能测试报告](./doc/KingDoc/kingshard_performance_test.md)

[7.kingshard的FAQ](./doc/KingDoc/function_FAQ.md)

### kingshard架构与设计

[1.kingshard架构设计和功能实现](./doc/KingDoc/architecture_of_kingshard_CN.md)

[2.kingshard性能优化之网络篇](./doc/KingDoc/kingshard_performance_profiling.md)

## License

kingshard采用Apache 2.0协议，相关协议请参看[目录](./doc/License)

## 鸣谢
- 感谢[mixer](https://github.com/siddontang/mixer)作者siddontang, kingshard最初的版本正是基于mixer开发而来的。
- 感谢[bigpyer](https://github.com/bigpyer)，他对kingshard做了详细的性能测试，并撰写了一份非常详细的测试报告。
- 感谢以下[开源爱好者](https://github.com/flike/kingshard/graphs/contributors)为kingshard做出的贡献。

## 反馈
kingshard开源以来，经过不断地迭代开发，功能较为完善，稳定性有较大提升。 **目前已有几家公司在生产环境使用kingshard作为MySQL代理。** 如果您在使用kingshard的过程中发现BUG或者有新的功能需求，非常欢迎您发邮件至hiflike#gmail.com与作者取得联系，或者加入QQ群(147926796)交流。
