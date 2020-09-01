package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const SrvName = "log_agent"

type Configure struct {
	MysqlSetting map[string]*MysqlConfig
	TCPSetting   map[string]*TCPConfig
	LocalSetting map[string]*LocalConfig
	KafkaSetting map[string]*KafkaConfig
	EtcdSetting  map[string]*EtcdConfig
}

type MysqlConfig struct {
	MysqlConn            string
	MysqlConnectPoolSize int
}

type KafkaConfig struct {
	Addrs             []string
	MaximumChanSize   int
	CheckServiceTopic string
}

type ElasticSearchConfig struct {
	Address string
}

type EtcdConfig struct {
	Addrs       []string
	DialTimeout int
}

type TCPConfig struct {
	ServerAddr            string
	ServerMaxOrderChanNum int
}

type LocalConfig struct {
	BookStoreDir string
}

var (
	Config *Configure
)

//从文件中读取配置信息，并Unmarshal到config结构体中

func Init(fileName string) error {
	_, err := os.Stat(fileName)
	if err != nil {
		fmt.Println("os.Stat(fileName) err: ", err)
		return err
	}
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("ioutil.ReadFile(fileName) err: ", err)
		return err
	}
	Config = new(Configure)
	err = json.Unmarshal(bytes, Config)
	if err != nil {
		fmt.Println("json.Unmarshal(bytes, Config) err: ", err)
		return err
	}
	return nil
}
