package tailog

import (
	"fmt"
	"github.com/hpcloud/tail"
	"kafka"
)

// 一个日志收集的任务
type TailTask struct{
	path string
	topic string
	instance *tail.Tail
}

func NewTailTask(path,topic string)(tailObj *TailTask){
	tailObj=&TailTask{
		path :path,
		topic: topic,
	}
	tailObj.init()
	return tailObj
}
func (t *TailTask)init(){
	config:=  tail.Config{
		ReOpen:    true,                                 // 文件被移除或被打包，需要重新打开,基础库会检测，如果文件有改变，会重新打开
		Follow:    true,                                 // 实时跟踪
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 如果程序出现异常，保存上次读取的位置，避免重新读取
		MustExist: false,                                //flase日志文件不存在也监控
		Poll:      true,                                 //不断的去查询
	}
	var err error
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Println("tail file err:", err)
		return
	}
	go t.run()
}

func (t* TailTask)run(){
	//从tailItem.Lines中读
	for {
		select {
		case line:=<-t.instance.Lines:
			kafka.WriteMsgToChan(t.topic,line.Text)
			//kafka.SendToKafka(t.topic,line.Text)
		}
	}
}


