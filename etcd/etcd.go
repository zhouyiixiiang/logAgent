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
	Path string `json:"path"`
	Topic string `json:"topic"`
}

func Init()(err error){
	config := clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}
	clientEtcd, err = clientv3.New(config)
	if err !=nil{
		return err
	}
	return
}
func Put(key,val string)(err error){
	ctx,cancel:=context.WithTimeout(context.Background(),time.Duration(5 * time.Second))
	_,err=clientEtcd.Put(ctx,key,val)
	cancel()
	if err !=nil{
		return err
	}
	return
}
func Get(key string)(err error){
	ctx,cancel:=context.WithTimeout(context.Background(),time.Duration(5 * time.Second))
	rsp,err:=clientEtcd.Get(ctx,key)
	cancel()
	if err!=nil{
		return err
	}
	for _,item := range rsp.Kvs{
		fmt.Println(item.Key," : ",item.Value)
	}
	return
}
func GetLogConf(key string)(logConf []*LogEntry,err error){
	ctx,cancel:=context.WithTimeout(context.Background(),time.Duration(5 * time.Second))
	rsp,err:=clientEtcd.Get(ctx,key)
	cancel()
	if err!=nil{
		fmt.Println("clientEtcd.Get err : ",err)
		return
	}
	for _,item := range rsp.Kvs{
		json.Unmarshal(item.Value,&logConf)
	}
	return logConf,err
}