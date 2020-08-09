package tailog

import (
	"etcd"
	"fmt"
	"time"
)

var tskMgr *tailLogMgr

type tailLogMgr struct {
	logEntry    []*etcd.LogEntry
	tskMap      map[string]*TailTask
	newConfChan chan []*etcd.LogEntry
}

func Init(logConf []*etcd.LogEntry) {
	tskMgr = &tailLogMgr{
		logEntry:    logConf,
		tskMap:      make(map[string]*TailTask, 16),
		newConfChan: make(chan []*etcd.LogEntry), //无缓冲区通道
	}
	for _, item := range logConf {
		tailTask := NewTailTask(item.Path, item.Topic)
		//将初始从etc里面读取到用于监听的tailTask存到tskMap里面，记录监听了多少个日志文件
		mk := fmt.Sprintf("%s_%s", tailTask.topic, tailTask.path)
		tskMgr.tskMap[mk] = tailTask
	}

	go tskMgr.run()
}

// 有了新的配置修改，更新到newConfChan，则做修改
func (t *tailLogMgr) run() {
	for {
		select {
		case newConf := <-t.newConfChan:
			fmt.Println("读取到新的配置", newConf)
			for _, conf := range newConf {
				mk := fmt.Sprintf("%s_%s", conf.Topic, conf.Path)
				_, ok := t.tskMap[mk]
				if ok {
					//原来就有，不需要操作
					continue
				}
				//如果是新增的或是修改的
				tailTask := NewTailTask(conf.Path, conf.Topic)
				//将初始从etc里面读取到用于监听的tailTask存到tskMap里面，记录监听了多少个日志文件
				t.tskMap[mk] = tailTask
			}
			// 找出原来有，现在conf没有的配置项，要删除掉
			// 找出不同项
			for _, task1 := range t.logEntry {
				isDelete := true
				for _, task2 := range newConf {
					if task1.Topic == task2.Topic && task1.Path == task2.Path {
						isDelete = false
						break
					}
				}
				if isDelete {
					// 删除原来的配置信息
					mk := fmt.Sprintf("%s_%s", task1.Topic, task1.Path)
					t.tskMap[mk].cancelFunc()
				}
			}
			//1、配置新增
			//2、配置删除
			//3、配置变更
		default:
			time.Sleep(time.Millisecond * 500)
		}
	}
}

//向外暴露newConfChan
func SetNewConf() chan<- []*etcd.LogEntry {
	return tskMgr.newConfChan
}
