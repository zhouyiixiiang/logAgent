package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var (
	clientEtcd *clientv3.Client
)

type LogEntry struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}

func Init(addrs []string, timeOut int) (err error) {
	config := clientv3.Config{
		Endpoints:   addrs,
		DialTimeout: time.Duration(timeOut) * time.Second,
	}
	clientEtcd, err = clientv3.New(config)
	if err != nil {
		return err
	}
	return
}
func Put(key, val string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	_, err = clientEtcd.Put(ctx, key, val)
	cancel()
	if err != nil {
		return err
	}
	return
}
func Get(key string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	rsp, err := clientEtcd.Get(ctx, key)
	cancel()
	if err != nil {
		return err
	}
	for _, item := range rsp.Kvs {
		fmt.Println(item.Key, " : ", item.Value)
	}
	return
}
func GetLogConf(key string) (logConf []*LogEntry, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	rsp, err := clientEtcd.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Println("clientEtcd.Get err : ", err)
		return
	}
	for _, item := range rsp.Kvs {
		json.Unmarshal(item.Value, &logConf)
	}
	return logConf, err
}

func WatchConf(key string, newConfCh chan<- []*LogEntry) {
	ch := clientEtcd.Watch(context.Background(), key)
	for item := range ch {
		for _, event := range item.Events {
			fmt.Printf("type: %v value: %v\n", event.Type, event.Kv)
			/*判断type
			1.
			*/
			var newConf []*LogEntry
			if event.Type != clientv3.EventTypeDelete {
				//如果是删除操作，手动传一个空的配置项
				err := json.Unmarshal(event.Kv.Value, &newConf)
				if err != nil {
					fmt.Println("json.Unmarshal(event.Kv.Value,&newConf) err: ", err)
					continue
				}
			}
			fmt.Printf("get new conf: %v\n", newConf)
			//通知tailogManager
			newConfCh <- newConf
		}
	}
}
