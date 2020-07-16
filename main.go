package main

import (
	"common"
	"fmt"
	"kafka"
	"tailog"
	"time"
)

func run() {
	// 读取日志
	for {
		select {
		case line := <-tailog.ReadChan():
			kafka.SendToKafka("web_log", line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}

func main() {
	// 初始化kafka连接
	err := kafka.Init([]string{"127.0.0.1:9092"})
	common.ErrorHandle(err, "kafka.Init")
	fmt.Println("init kafka success")
	// 打开日志文件准备收集日志
	err = tailog.Init("./my.log")
	common.ErrorHandle(err, "tailog.Init")
	fmt.Println("init tailog success")
	run()
}
