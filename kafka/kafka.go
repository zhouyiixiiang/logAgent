package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

var (
	clientKafka sarama.SyncProducer
	msgChan     chan *KafkaMsg
)

type KafkaMsg struct {
	topic string
	data  string
}

func Init(addrs []string, kafkamaximumSize int) (err error) {
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
	msgChan = make(chan *KafkaMsg, kafkamaximumSize)
	go SendToKafka()
	//defer  clientKafka.Close()
	return
}

func WriteMsgToChan(topic, data string) {
	msg := &KafkaMsg{
		topic: topic,
		data:  data,
	}
	msgChan <- msg
}

func SendToKafka() (err error) {
	for {
		select {
		case mg := <-msgChan:
			// 构造一个消息
			msg := &sarama.ProducerMessage{}
			msg.Topic = mg.topic
			msg.Value = sarama.StringEncoder(mg.data)
			// 发送消息到kafka
			pid, offset, err := clientKafka.SendMessage(msg)
			if err != nil {
				return err
			}
			fmt.Printf("pid:%v offset:%v topic:%v value:%v\n", pid, offset, msg.Topic, msg.Value)
		default:
			time.Sleep(time.Millisecond * 50)
		}
	}
}
