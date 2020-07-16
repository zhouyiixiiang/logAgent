package tailog

import (
	"fmt"
	"github.com/hpcloud/tail"
)

var (
	tailItem *tail.Tail
	LogChan  chan string
)

func Init(filename string) (err error) {
	tailItem, err = tail.TailFile(filename, tail.Config{
		ReOpen:    true,                                 // 文件被移除或被打包，需要重新打开,基础库会检测，如果文件有改变，会重新打开
		Follow:    true,                                 // 实时跟踪
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 如果程序出现异常，保存上次读取的位置，避免重新读取
		MustExist: false,                                //flase日志文件不存在也监控
		Poll:      true,                                 //不断的去查询
	})
	if err != nil {
		fmt.Println("tail file err:", err)
		return err
	}
	return
}

func ReadChan() <-chan *tail.Line {
	return tailItem.Lines
}
