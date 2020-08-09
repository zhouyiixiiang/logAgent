# kafka
kafka集群架构：

* broker：集群里的某台节点，可能是一台服务器上的不同端口，也可能是不同ip地址
* topic：主题
* partition：分区，把数据分为好几个分区从而提高负载
	* leader：分区主节点
	* follower：分区从节点
* consumer group：消费者组

生产者往kafka发送数据流程：

* 生产者与集群leader建立连接
* 生产者往leader发送sync
* leader将数据保存到本地
* follower主动向leader拉取数据
* follower保存数据到本地，并返回ack给leader
* leader返回ack给生产者（根据发送级别0，1，all）

分区存储数据原理：

* 把磁盘随机读变为顺序读，基于时间index、消息index、segment文件、offset偏移量来定位消息具体log文件中在哪里

消费者组消费数据原理：

* 同一个partion只能被一个消费者组中的一个消费者消费
* 同一个partion会被多个消费者组中的一个消费者以轮询的方式消费

## kafka & zookeeper

### 启动zookeeper

配置文件在 ``` C:\softwares\kafka\kafka_2.12-2.5.0\config\zookeeper.properties```

将里面的```dataDir=/tmp/zookeeper```(数据存放路径)改为自己定义的数据存放路径```D:/tmp/zookeeper```

定义zookeeper的端口号```clientPort=2181```

```
dataDir=D:/tmp/zookeeper
# the port at which the clients will connect
clientPort=2181
# disable the per-ip limit on the number of connections since this is a non-production config
maxClientCnxns=0
# Disable the adminserver by default to avoid port conflicts.
# Set the port to something non-conflicting if choosing to enable this
admin.enableServer=false
```

配置完成后，可以启动zookeeper：

```
cd C:\softwares\kafka\kafka_2.12-2.5.0
bin\windows\zookeeper-server-start.bat config\zookeeper.properties
```



启动kafka 

kafka 的配置文件路径为 ``` C:\softwares\kafka\kafka_2.12-2.5.0\config\server.properties```

```
# A comma separated list of directories under which to store log files
log.dirs=D:/tmp/kafka-logs

# The default number of log partitions per topic. More partitions allow greater
# parallelism for consumption, but this will also result in more files across
# the brokers.
num.partitions=1

zookeeper.connect=localhost:2181
zookeeper.connection.timeout.ms=18000

#listeners=PLAINTEXT://:9092
```

以上是配置kafka的日志文件以及分区数，zookeeper的ip地址、最大超时时间，监听地址默认是9092端口

配置完成后启动kafka

```
cd C:\softwares\kafka\kafka_2.12-2.5.0
bin\windows\kafka-server-start.bat config\server.properties
```
## mac版kafka
```
# 启动zookeeper
/usr/local/bin/zookeeper-server-start /usr/local/etc/kafka/zookeeper.properties
# 启动kafka
/usr/local/bin/kafka-server-start /usr/local/etc/kafka/server.properties
# 启动消费者
/usr/local/bin/kafka-console-consumer --topic=checkService --bootstrap-server=127.0.0.1:9092 --from-beginning
```
## ETCD
```
# 启动etcd
cd /Users/zhouyixiang/Documents/softwares/etcd/etcd-v3.4.10-darwin-amd64
etcd
# 使用etcdctl
etcdctl --endpoints=http://localhost:2379 put logAgentConf "newconf"
etcdctl --endpoints=http://localhost:2379 get logAgentConf
```


### 快速拉取go仓库：设置代理

设置go代理，获取tail库

```SET GOPROXY=https://goproxy.cn```

``` go get github.com/hpcloud/tail```

### Go版本是1.13及以上

```
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.io,direct

# 设置不走proxy的私有仓库，多个用逗号相隔
go env -w GOPRIVATE=
```

##  GO Mod使用

### 初始化

```
cd project
go mod init project
```

此时会在project目录下生成go.mod，文件内容以project为根目录

go build 会自动查找文件依赖并下载

### 下载依赖

```
go mod download
```

下载依赖并再my_project下生成go.lock

ide支持goland：

* 在preferences->go->go mode 选择enable
* 再preferences->go->gopath去掉所有gopath

### 运行

prod 环境不需要指定gopath

dev需要指定gopath，用于查找依赖
查看gopath
echo $GOPATH
指定临时gopath
export GOPATH=/Users/ocean/go



# github

```
cd project
git init #在这个目录里初始化git仓库
git commit -m "first commit"
git remote add origin https://github.com/zhouyiixiiang/casitworkspace.git
git push -u origin master
```

### github操作

```
# 查看本地和远程分支
git branch -a
# 在本地新建一个分支
git branch yym
# git checkout -b 本地分支名x origin/远程分支名x 拉取远程分支并同时创建对应的本地分支
git checkout -b yym origin/unsy

# 删除本地分支
git branch -d yym
# 放弃本地所有修改
git checkout .
```



## 测试kafka效果

路径： ```/usr/local/bin/kafka-console-consumer```

使用kafka自带的终端消费者读取消息到终端

```
cd  /usr/local/bin/
# 查看帮助文档
kafka-console-consumer -help

kafka-console-consumer --bootstrap-server=127.0.0.1:9092 --topic=web_log --from-beginning

```
### context
goroutine 管理

context.Context

两个根节点：```context.TODO()```、```context.Background()```

四个方法：

* ```context.withCancel```：把context.Context变量调用一下withCancel，返回一个context对象和一
* ```context.withTimeout()``` 设置相对时间
* ```context.withDeadline()```设置绝对时间
* ```context.withValue()```用在请求，一般web框架来一个请求，启动一个goroutine来处理，这个goroutine下可能衍生其他goroutine，使用其，可以实现父goroutine下的多个子goroutine的传值

用法回顾

```
func f(ctx context.Context){
	defer wg.Done
loop:
	for {
		select{
			case:<-ctx.Done
				break loop
			default:
				time.sleep(time.Second)
		}	
	}
}

var wg sync.WaitGroup
func main(){
	ctx,cancel:=context.WithCancel(context.Background())
	go f(ctx)
	time.sleep(3*time.second)
	cancel()
	wg.Wait()
}
```
## 日志收集项目回顾
为什么自己写不用ELK？

* ELK：部署的时候麻烦，每一个filebeat都需要配置一个配置文件

因此：

* 使用etcd来管理被收集的日志项

### 日志架构

流程：

* 发起kafka和etcd，在多台服务器上跑起来，形成集群，利用etcd的raft算法保持配置一致性
* 每一个要进行监听日志的服务前往etcd注册一下日志地址以及Topic。
* tailLog服务根据etcd上登记的需要监听日志的对象，追踪每一个对象的日志信息，并派一个watcher观察追踪日志信息是否发生改变
* 如果监听的日志新增日志信息，tailLog将新增信息根据topic发送至kafka消息队列
* etcd watch实现方法：底层通过webSocket实现给客户端发送通知
* 如果有日志信息改变，例如删除了，可以通过context.Cancel退出该goroutine

整体回顾：

* 不同机器的配置不同，但需要同一个etcd配置
	* 可以从etcd起来的时候，获取本机的ip地址，用于不同配置的实现	



