package main

import (
	"common"
	"etcd"
	"fmt"
	"kafka"
	"sync"
	"tailog"
	"util"
)

func main() {
	var err error
	// 初始化kafka连接----↓---
	err = kafka.Init([]string{"127.0.0.1:9092"}, 100000)
	common.ErrorHandle(err, "kafka.Init")
	fmt.Println("init kafka success")
	//初始化etcd
	err = etcd.Init()
	common.ErrorHandle(err, "etcd.Init")
	fmt.Println("init etcd success")
	// 为每一个etcd获取配置信息
	ipStr, err := util.GetOutBoundIP()
	// 每一个机器的配置信息为"logAgentConf"+ipStr
	etcdConfKey := fmt.Sprintf("logAgentConf%s", ipStr)
	// 从etcd读取配置配置文件信息
	logConf, err := etcd.GetLogConf(etcdConfKey)
	//通过配置信息监听对应的日志文件
	tailog.Init(logConf)
	//派一个哨兵监听etcd的配置变动
	newConfChan := tailog.SetNewConf()
	go etcd.WatchConf(etcdConfKey, newConfChan)
	//防止主线程退出
	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
