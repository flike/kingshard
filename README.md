[![Build Status](https://travis-ci.org/flike/kingshard.svg?branch=master)](https://travis-ci.org/flike/kingshard)

# Overview
kingshard is a high-performance proxy for MySQL powered by Go. Just like other mysql proxies, you can use it to split the read/write sqls. Now it supports basic SQL statements (select, insert, update, replace, delete). The most important feature is the sharding function. Kingshard aims to simplify the sharding solution of MySQL. **The Performance of kingshard is about 75% compared to connecting to MySQL directly.**

# Feature
- Splits reads and writes
- Sharding table across multiple nodes
- Client's ip ACL control.
- Transaction in single node.
- Support setting the backend database online or offline dynamically.
- Supports prepared statement: COM_STMT_PREPARE, COM_STMT_EXECUTE, etc.
- Support multi slaves, and loading banlance between slaves.
- Support reading master database forcely.
- Support sending sql to the specified node.
- Support most commonly used functions, such as `max, min, count, sum`, and also support `join, limit, order by`. 
- MySQL HA

# Install
```
  1. Install Go
  2. git clone https://github.com/flike/kingshard.git src/github.com/flike/kingshard
  3. cd src/github.com/flike/kingshard
  4. source ./dev.sh
  5. make
  6. set the config file (etc/multi.yaml)
  7. run kingshard (./bin/kingshard -config=etc/multi.yaml)
```
	
# Details of kingshard

[1.How to use kingshard building a MySQL cluster](./doc/KingDoc/how_to_use_kingshard_EN.md)

# License

kingshard is under the MIT license. See the [LICENSE](./doc/License) directory for details.

# Other language version

[简体中文](README_ZH.md)
