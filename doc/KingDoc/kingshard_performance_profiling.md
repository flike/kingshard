# kingshard性能优化网络篇

最近kingshard的功能开发节奏慢了许多。一方面是工作确实比较忙，另一方面是我觉得kingshard的功能已经比较完善了，下一步的开发重点应该是性能优化。毕竟作为一个MySQL proxy,如果转发SQL的性能很差，再多的功能都无济于事。所以这个周末一直宅在家里优化kingshard的转发性能。经过两天的探索发现，将kingshard的转发SQL性能提升了18%左右，在这个过程中学到了一下知识。借此机会分享一下，同时也是督促一下自己写博客的积极性。：）

## 1. 发现kingshard的性能瓶颈

首选，对kingshard进行性能优化，我们必须要找到kingshard的性能瓶颈在哪里。Go语言在性能优化支持方面做的非常好，借助于go语言的pprof工具，我们可以通过简单的几个步骤，就能得到kingshard在转发SQL请求时的各个函数耗时情况。

### 1.1 环境搭建
根据[kingshard使用指南](https://github.com/flike/kingshard/blob/master/doc/KingDoc/how_to_use_kingshard.md)搭建一个kingshard代理环境。我是用macbook搭建的环境，硬件参数如下所示：

```
CPU： 2.2GHZ * 4
内存：16GB
硬盘: 256GB

```

### 1.2 性能测试步骤
具体步骤如下所述：

1.获取一个性能分析的封装库
```
go get github.com/pkg/profile
```

2.在工程内import这个组件

3.在kingshard/cmd/kingshard/main.go的main函数开始部分添加CPU监控的启动和停止入口

```
func main() {
	defer profile.Start(profile.CPUProfile).Stop()
	fmt.Print(banner)
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()
	....
}
```
4.重新编译工程, 运行kingshard

```
./bin/kingshard -config=etc/ks.yaml
```

5.kingshard启动后会在终端输出下面一段提示：

```
2015/10/31 10:28:06 profile: cpu profiling enabled, /var/folders/4q/zzb55sfj377b6vdyz2brt6sc0000gn/T/profile205276958/cpu.pprof
```
后面的路径就是pprof性能分析文件的位置，Ctrl+C中断服务器

6.这时候用sysbench对kingshard进行压力测试，得到QPS（有关sysbench的安装和使用，请自行Google解决）。具体的代码如下所示：

```
sysbench --test=oltp --num-threads=16 --max-requests=160000 --oltp-test-mode=nontrx --db-driver=mysql --mysql-db=kingshard --mysql-host=127.0.0.1 --mysql-port=9696 --mysql-table-engine=innodb --oltp-table-size=10000 --mysql-user=kingshard --mysql-password=kingshard --oltp-nontrx-mode=select --db-ps-mode=disable run

```
得到如下结果:

```
OLTP test statistics:
    queries performed:
        read:                            160071
        write:                           0
        other:                           0
        total:                           160071
    transactions:                        160071 (16552.58 per sec.)
    deadlocks:                           0      (0.00 per sec.)
    read/write requests:                 160071 (16552.58 per sec.)
    other operations:                    0      (0.00 per sec.)

Test execution summary:
    total time:                          9.6705s
    total number of events:              160071
    total time taken by event execution: 154.4474
    per-request statistics:
         min:                                  0.29ms
         avg:                                  0.96ms
         max:                                 14.17ms
         approx.  95 percentile:               1.37ms

Threads fairness:
    events (avg/stddev):           10004.4375/24.95
    execution time (avg/stddev):   9.6530/0.00

```
- 按照上述步骤测试三次（16552.58,16769.72,16550.16）取平均值，得到优化前kingshard的QPS是：16624.15

- 按照上述步骤，直连MySQL。测试直连MySQL的QPS，同样测试三次QPS（27730.90，28499.05，27119.20），得到直连MySQL的QPS是：27783.05。
- 从上述数据可以计算出kingshard转发SQL的性能是直连MySQL的59%左右。

7.将cpu.prof拷贝到bin/kingshard所在位置

8.调用go tool工具制作CPU耗时的PDF文档
```
go tool pprof -pdf ./kingshard cpu.pprof > report.pdf
```

## 2. 性能测试报告分析

通过上述命令，可以生成压测期间主要函数耗时情况。从[report](./report.pdf)来看，主要的耗时在TCP层数据包的收发上面。那我们应该主要考虑如何优化TCP层数据的收发方面。优化TCP传输效率，我首先想到了减少系统调用，每个数据包传输尽量多的数据。

在通过 TCP socket 进行通信时，数据都拆分成了数据块，这样它们就可以封装到给定连接的 TCP payload（指 TCP 数据包中的有效负荷）中了。TCP payload 的大小取决于几个因素（例如最大报文长度和路径），但是这些因素在连接发起时都是已知的。为了达到最好的性能，我们的目标是使用尽可能多的可用数据来填充每个报文。当没有足够的数据来填充 payload 时（也称为最大报文段长度（maximum segment size） 或 MSS），TCP 就会采用 Nagle 算法自动将一些小的缓冲区连接到一个报文段中。这样可以通过最小化所发送的报文的数量来提高应用程序的效率，并减轻整体的网络拥塞问题。

由于这种算法对数据进行合并，试图构成一个完整的 TCP 报文段，因此它会引入一些延时。但是这种算法可以最小化在线路上发送的报文的数量，因此可以最小化网络拥塞的问题。但是在需要最小化传输延时的情况中，GO语言中Sockets API 可以提供一种解决方案。就是通过：
```
func (c *TCPConn) SetNoDelay(noDelay bool) error
```
这个函数在Go中默认情况下，是设置为true，也就是未开启延迟选项。我们需要将其设置为false选项，来达到每个数据包传输尽量多的数据，减少系统调用的目的。

## 2.1 代码修改和性能测试

发现了性能瓶颈以后，修改proxy/server/server.go文件中的newClientConn函数和backend/backend_conn.go中的ReConnect函数，分别设置client与kingshard之间的连接和kingshard到MySQL之间的连接为最小化传输延时。具体的代码修改可以查看这个[commit](https://github.com/flike/kingshard/commit/6c175d127c7b15b527cedb02876634901f2b9be1)。

修改后我们利用sysbench重新测试，测试命令和上述测试一致。得到的结果如下所示：

```
OLTP test statistics:
    queries performed:
        read:                            160174
        write:                           0
        other:                           0
        total:                           160174
    transactions:                        160174 (21291.68 per sec.)
    deadlocks:                           0      (0.00 per sec.)
    read/write requests:                 160174 (21291.68 per sec.)
    other operations:                    0      (0.00 per sec.)

Test execution summary:
    total time:                          7.5228s
    total number of events:              160174
    total time taken by event execution: 119.9655
    per-request statistics:
         min:                                  0.26ms
         avg:                                  0.75ms
         max:                                 10.78ms
         approx.  95 percentile:               1.13ms

Threads fairness:
    events (avg/stddev):           10010.8750/38.65
    execution time (avg/stddev):   7.4978/0.00
```

测试三次得到的QPS为：21291.68,21670.85,21463.44。 **相当于直连MySQL性能的77%左右，通过这个优化性能提升了18%左右**。

## 总结

通过这篇文章，介绍了通过Go语言提供的pprof对kingshard进行性能分析的详细步骤。对于其他Go语言项目也可以通过类似步骤生成性能报告文档。性能优化的关键是发现性能瓶颈，再去找优化方案。有时候简单的优化，就可以达到预想不到的效果，希望本文能给Go开发者在性能优化方面提供一个思路。最后打个广告：kingshard作为一个支持sharding的开源MySQL中间件项目，目前已经比较稳定了，且经过性能优化后，转发SQL的性能提升了不少。后续我还会在锁和内存方面对kingshard进行优化，敬请期待。

