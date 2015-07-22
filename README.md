[![Build Status](https://travis-ci.org/flike/kingshard.svg?branch=master)](https://travis-ci.org/flike/kingshard)

#Overview
kingshard is a high-performance proxy for MySQL powered by Go. Just like other mysql proxies, you can use it to split the read/write sqls. Now it supports basic SQL statements (select, insert, update, replace, delete). The most important feature is the sharding function. Kingshard aims to simplify the sharding solution of MySQL.

#Feature
- splits reads and writes
- sharding table across multiple nodes
- client's ip ACL control.
- supports prepared statement: COM_STMT_PREPARE, COM_STMT_EXECUTE, etc.
- MySQL HA

#Install
    1. Install Go
    2. git clone https://github.com/flike/kingshard.git src/github.com/flike/kingshard
    3. cd src/github.com/flike/kingshard
    4. source ./dev.sh
    5. make
    6. set the config file (etc/multi.yaml)
    7. run kingshard (./bin/kingshard -config=etc/multi.yaml)

#Other language version

[简体中文](README_ZH.md)
