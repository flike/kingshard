[![Build Status](https://travis-ci.org/flike/kingshard.svg?branch=master)](https://travis-ci.org/flike/kingshard)

# kingshard简介

kingshard是一个由Go开发高性能MySQL Proxy项目，kingshard在满足基本的读写分离的功能上，致力于简化MySQL分库分表操作；能够让DBA通过kingshard轻松平滑地实现MySQL数据库扩容。

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
	10. 支持跨node的count和sum操作。
	11. 支持order by,limit等操作。
	12. 支持主流语言（java,php,python,C/C++,Go)SDK的mysql的prepare特性。
	
## kinshard详细说明

[1.安装kingshard](./doc/KingDoc/kingshard_install_document.md)

[2.kingshard 快速入门](./doc/KingDoc/kingshard_quick_try.md)

[3.kingshard架构设计和功能实现](./doc/KingDoc/architecture_of_kingshard_CN.md)

[4.kingshard sharding介绍](./doc/KingDoc/kingshard_sharding_introduce.md)

[5.功能FAQ](./doc/KingDoc/function_FAQ.md)

[6.管理端命令介绍](./doc/KingDoc/admin_command_introduce.md)

[7.ChangeLog](./doc/KingDoc/change_log_CN.md)

## 反馈
目前kingshard还是1.0版本，比较核心的功能已经实现了。但还有很多地方不完善。如果您在使用kingshard的过程中发现BUG或者有新的功能需求，非常欢迎您发邮件至flikecn#126.com与作者取得联系，或者加入QQ群(147926796)交流。
