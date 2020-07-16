package kafka

import (
	"github.com/Shopify/sarama"
	"common"
)
var (
	clientKafka sarama.AsyncProducer
)
func Init(){
	config:=sarama.NewConfig()
	config.Producer.RequiredAcks=sarama.WaitForAll
	config.Producer.Partitioner=sarama.NewRandomPartitioner
	config.Producer.Return.Successes=true
	var err error
	clientKafka,err=sarama.NewAsyncProducer([]string{"127.0.0.1:9092"},config)
	common.ErrorHandle(err,"sarama.NewAsyncProducer")
	//defer  clientKafka.Close()
}

func Run(){
	clientKafka.Input()
}
