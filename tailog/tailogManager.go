package tailog

import "etcd"

var tskMgr  *tailLogMgr

type tailLogMgr struct{
	logEntry []*etcd.LogEntry
	tsk map[string]*TailTask
}

func Init(logConf []*etcd.LogEntry){
	tskMgr=&tailLogMgr{
		logEntry: logConf,
	}
	for _,item :=range(logConf){
		NewTailTask(item.Path,item.Topic)
	}
}
