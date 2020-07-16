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

## 测试kafka效果

路径： ```/usr/local/bin/kafka-console-consumer```

使用kafka自带的终端消费者读取消息到终端

```
cd  /usr/local/bin/
# 查看帮助文档
kafka-console-consumer -help

kafka-console-consumer --bootstrap-server=127.0.0.1:9092 --topic=web_log --from-beginning

```