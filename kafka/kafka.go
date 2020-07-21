package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

var (
	clientKafka sarama.SyncProducer
)

func Init(addrs []string) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	// 连接kafka
	clientKafka, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		fmt.Println("sarama.NewSyncProducer err: ", err)
		return err
	}
	//defer  clientKafka.Close()
	return
}

func SendToKafka(topic, data string)(err error) {
	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)
	// 发送消息到kafka
	pid, offset, err := clientKafka.SendMessage(msg)
	if err !=nil{
		return err
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
	return
}
