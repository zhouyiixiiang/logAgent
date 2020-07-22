package main

import (
	"common"
	"etcd"
	"fmt"
	"kafka"
	"sync"
	"tailog"
)

func main() {
	var err error
	// 初始化kafka连接
	err = kafka.Init([]string{"127.0.0.1:9092"})
	common.ErrorHandle(err, "kafka.Init")
	fmt.Println("init kafka success")
	//初始化etcd
	err=etcd.Init()
	common.ErrorHandle(err, "etcd.Init")
	fmt.Println("init etcd success")
	// 从etcd读取配置配置文件信息
	logConf,err:=etcd.GetLogConf("logAgentConf")
	//通过配置信息监听对应的日志文件
	tailog.Init(logConf)
	wg :=sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
